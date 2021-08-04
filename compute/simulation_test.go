package compute

import (
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/HenryGeorgist/go-fathom/hazard_providers"
	"github.com/USACE/go-consequences/census"
	"github.com/USACE/go-consequences/structureprovider"
)

func TestCompute_UserProvidedList(t *testing.T) {
	ss := []string{"11"} //[]string{"02", "05", "15", "41", "20", "56", "45"}
	st := time.Now()
	var wg sync.WaitGroup
	wg.Add(len(ss))
	for _, s := range ss {
		go func(state string) {
			defer wg.Done()
			computeState(state)
		}(s)
	}
	wg.Wait()
	et := time.Now()
	diff := et.Sub(st)
	fmt.Println(diff)
}
func computeState(ss string) {
	path := fmt.Sprintf("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_Uncertainty\\NSI_Fathom_depths%v_feet.csv", ss)
	ds := hazard_providers.ReadFeetFile(path)
	outputpath := fmt.Sprintf("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_Uncertainty\\NSI_Fathom_damages_%v.csv", ss)
	outputFile, err := os.Create(outputpath)
	defer outputFile.Close()
	if err != nil {
		panic(err)
	}
	//compute
	sp, err := structureprovider.InitGPK("C:\\Examples\\go-tc-consequences\\data\\nsi.gpkg", "nsi")
	if err != nil {
		panic(err)
	}
	//sp := structureprovider.InitNSISP()
	frequencies := []float64{.5, .2, .05, .01, .004, .002}
	sds := hazard_providers.StochasticDataSet{Data: ds.Data, StandardDeviation: 3.28084, Frequencies: frequencies}
	iterations := 1
	ComputeEadByFips(sds, sp, ss, outputFile, iterations)
}
func Test_Compute_AllStates(t *testing.T) {
	fmap := census.StateToCountyFipsMap()
	states := make([]string, len(fmap))
	i := 0
	max := 7
	stateCounter := 0
	for ss, _ := range fmap {
		states[stateCounter] = ss
		stateCounter++
	}
	for i < len(states) {
		var limiter sync.WaitGroup
		for j := 0; j < max; j++ {
			if j == 0 {
				limiter.Add(max)
			}
			go func(state string) {
				defer limiter.Done()
				computeState(state)
			}(states[i])
			i++
		}
		limiter.Wait()
	}
}
