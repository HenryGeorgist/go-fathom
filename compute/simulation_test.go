package compute

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/HenryGeorgist/go-fathom/hazard_providers"
	"github.com/HenryGeorgist/go-fathom/store"
	"github.com/USACE/go-consequences/consequences"
	"github.com/USACE/go-consequences/hazards"
	sp "github.com/USACE/go-consequences/structureprovider"
)

// func TestSingleEvent(t *testing.T) {
// 	fmt.Println("Reading Depths")
// 	ds := hazard_providers.ReadFeetFile("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_depths_Filtered_Feet.csv")
// 	fmt.Println("Finished Reading Depths")
// 	fe := hazard_providers.FathomEvent{Year: 2050, Frequency: 5, Fluvial: true}
// 	ComputeSingleEvent_NSIStream(ds, "11", fe)
// }
// func TestMultiEvent(t *testing.T) {
// 	fmt.Println("Reading Depths")
// 	ds := hazard_providers.ReadFeetFile("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_depths_Filtered_Feet.csv")
// 	fmt.Println("Finished Reading Depths")
// 	db := store.CreateDatabase()
// 	defer db.Close()
// 	ComputeMultiEvent_NSIStream(ds, "11", db)
// }
// func TestMultiEvent_MultiState(t *testing.T) {
// 	fmt.Println("Reading Depths")
// 	ds := hazard_providers.ReadFeetFile("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_depths_Filtered_Feet.csv")
// 	fmt.Println("Finished Reading Depths")
// 	ComputeMultiFips_MultiEvent(ds)
// }
func TestSQLMultiEvent_SingleState(t *testing.T) {
	fmt.Println("Reading Depths")
	ds := hazard_providers.OpenSQLDepthDataSet("/workspaces/go-fathom/data/nsiv2_29.gpkg")
	fmt.Println("Finished Reading Depths")
	db := store.CreateDatabase()
	defer db.Close()
	ComputeMultiEvent_NSIStream(ds, "29005", db)
}
func TestSQL_MultiEvent_MultiState(t *testing.T) {
	fmt.Println("Reading Depths")
	ds := hazard_providers.OpenSQLDepthDataSet("/workspaces/go-fathom/data/nsiv2_29.gpkg")
	fmt.Println("Finished Reading Depths")
	ComputeMultiFips_MultiEvent(ds)
}

func TestGPKByFips(t *testing.T) {
	filepath := "/workspaces/go-fathom/data/nsiv2_29.gpkg"
	nsp := sp.InitGPK(filepath, "nsi")
	// take only the first 2000 structures to ensure it works in 30 seconds
	fmt.Println(nsp.FilePath)
	d := hazards.DepthEvent{}
	d.SetDepth(2.4)
	count := 0
	nsp.ByFips("29005", func(s consequences.Receptor) {
		r := s.Compute(d)
		b, _ := json.Marshal(r)
		fmt.Println(string(b))
		count++
		if count == 2000 {
			fmt.Println("2000 structures done")
		}
	})
	fmt.Println(count)
}

// func TestComputeNewFile(t *testing.T) {
// 	ss := []string{"51"}
// 	var wg sync.WaitGroup
// 	wg.Add(len(ss))
// 	for _, s := range ss {
// 		go func(state string) {
// 			defer wg.Done()
// 			computeState_New(state)
// 		}(s)
// 	}
// 	wg.Wait()
// 	//computeState_New(ss)
// }

// func computeState_New(ss string) {
// 	path := fmt.Sprintf("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_depths_Pluvial\\NSI_Fathom_depths%v_feet.csv", ss)
// 	ds := hazard_providers.ReadFeetFile(path)
// 	outputpath := fmt.Sprintf("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_depths_Pluvial\\NSI_Fathom_damages_%v.csv", ss)
// 	outputFile, error := os.Create(outputpath)
// 	defer outputFile.Close()
// 	if error != nil {
// 		panic(error)
// 	}
// 	//compute
// 	ComputeMultiEvent_NSIStream_toFile_withNew(ds, ss, outputFile, true)
// }
// func TestMulti_State_ComputeNewFile(t *testing.T) {
// 	fmap := census.StateToCountyFipsMap()
// 	states := make([]string, len(fmap))
// 	//var wg sync.WaitGroup
// 	//wg.Add(len(fmap))
// 	i := 0
// 	max := 7
// 	stateCounter := 0
// 	for ss, _ := range fmap {
// 		states[stateCounter] = ss
// 		stateCounter++
// 	}
// 	for i < len(states) {
// 		var limiter sync.WaitGroup
// 		//limiter.Add(max) // maybe?
// 		for j := 0; j < max; j++ {
// 			if j == 0 {
// 				limiter.Add(max)
// 			}
// 			go func(state string) {
// 				defer limiter.Done()
// 				computeState_New(state)
// 			}(states[i])
// 			i++
// 		}
// 		limiter.Wait()
// 	}
// 	//wg.Wait()
// }
