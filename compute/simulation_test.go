package compute

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/HenryGeorgist/go-fathom/hazard_providers"
	"github.com/USACE/go-consequences/census"
	"github.com/USACE/go-consequences/structureprovider"
)

func TestCompute_UserProvidedList(t *testing.T) {
	ss := []string{"24"} //[]string{"02", "05", "15", "41", "20", "56", "45"}
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
	path := fmt.Sprintf("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_3m_Uncertainty\\NSI_Fathom_depths%v_feet.csv", ss)
	ds := hazard_providers.ReadFeetFile(path)
	outputpath := fmt.Sprintf("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_3m_Uncertainty\\NSI_Fathom_damages_%v_v4.csv", ss)
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
	//frequencies := []float64{.5, .2, .05, .01, .004, .002}
	frequencies := []float64{.2, .05, .01, .004, .002}
	sds := hazard_providers.StochasticDataSet{Data: ds.Data, StandardDeviation: 3.28084, Frequencies: frequencies}
	iterations := 100
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
				if i+max > len(states) {
					limiter.Add(len(states) - i - 1)
				} else {
					limiter.Add(max)
				}
			}
			if i < (len(states)) {
				go func(state string) {
					defer limiter.Done()
					computeState(state)
				}(states[i])
				i++
			}
		}
		limiter.Wait()
	}
}
func Test_Read_State_Output_For_EAD(t *testing.T) {
	ss := "11"
	fileA := fmt.Sprintf("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_Uncertainty\\NSI_Fathom_damages_%v_10y_stories_damagecurves.csv", ss)
	fileB := fmt.Sprintf("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_Uncertainty\\NSI_Fathom_damages_%v_triangles.csv", ss)
	ma, erra := createMapOfStructureDamages(fileA)
	if erra != nil {
		t.Error(erra)
	}
	mb, errb := createMapOfStructureDamages(fileB)
	if errb != nil {
		t.Error(errb)
	}
	for keya, valuea := range ma {
		valueb, inb := mb[keya]
		if inb {
			diff := valueb.Damage - valuea.Damage
			percent := math.Abs(diff / valueb.Damage)
			if diff < 0 {
				if percent > 0.1 {
					//fmt.Println(fmt.Sprintf("For structure %v with occtype %v damages of %v were greater in %v", keya, valueb.Occtype, valuea.Damage, fileA))
				}
			}

		} else {
			//fmt.Println(fmt.Sprintf("structure %v was not in %v", keya, fileB))
		}
	}
}

type Result struct {
	Occtype string
	Damage  float64
}

func createMapOfStructureDamages(file string) (map[string]Result, error) {
	m := make(map[string]Result)
	f, err := os.Open(file)
	if err != nil {
		return nil, errors.New("could not open file")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	//fmt.Println(scanner.Text()) //header row
	sum := 0.0
	for scanner.Scan() {
		lines := strings.Split(scanner.Text(), ",")
		ead, err := strconv.ParseFloat(lines[110], 64)
		if err != nil {
			return nil, errors.New("could not open file")
		}
		sum += ead
		m[lines[0]] = Result{Damage: ead, Occtype: lines[5]}
	}
	fmt.Println(fmt.Sprintf("EAD WAS %f", sum))
	return m, nil
}
