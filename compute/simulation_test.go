package compute

import (
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/HenryGeorgist/go-fathom/hazard_providers"
	"github.com/USACE/go-consequences/census"
	"github.com/USACE/go-consequences/structureprovider"
)

func TestComputeNewFile(t *testing.T) {
	ss := []string{"11"}
	var wg sync.WaitGroup
	wg.Add(len(ss))
	for _, s := range ss {
		go func(state string) {
			defer wg.Done()
			computeState_New(state)
		}(s)
	}
	wg.Wait()
	//computeState_New(ss)
}
func computeState_New(ss string) {
	path := fmt.Sprintf("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_depths_Pluvial\\NSI_Fathom_depths%v_feet.csv", ss)
	ds := hazard_providers.ReadFeetFile(path)
	outputpath := fmt.Sprintf("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_depths_Pluvial\\NSI_Fathom_damages_%v.csv", ss)
	outputFile, err := os.Create(outputpath)
	defer outputFile.Close()
	if err != nil {
		panic(err)
	}
	//compute
	/*sp, err := structureprovider.InitGPK("C:\\Examples\\go-tc-consequences\\data\\nsi.gpkg", "nsi")
	if err != nil {
		panic(err)
	}*/
	sp := structureprovider.InitNSISP()
	ComputeMultiEvent_NSIStream_toFile_withNew(ds, sp, ss, outputFile, true)
}
func TestMulti_State_ComputeNewFile(t *testing.T) {
	fmap := census.StateToCountyFipsMap()
	states := make([]string, len(fmap))
	//var wg sync.WaitGroup
	//wg.Add(len(fmap))
	i := 0
	max := 7
	stateCounter := 0
	for ss, _ := range fmap {
		states[stateCounter] = ss
		stateCounter++
	}
	for i < len(states) {
		var limiter sync.WaitGroup
		//limiter.Add(max) // maybe?
		for j := 0; j < max; j++ {
			if j == 0 {
				limiter.Add(max)
			}
			go func(state string) {
				defer limiter.Done()
				computeState_New(state)
			}(states[i])
			i++
		}
		limiter.Wait()
	}
	//wg.Wait()
}
