package compute

import (
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/HenryGeorgist/go-fathom/hazard_providers"
	"github.com/HenryGeorgist/go-fathom/store"
	"github.com/USACE/go-consequences/census"
)

func TestSingleEvent(t *testing.T) {
	fmt.Println("Reading Depths")
	ds := hazard_providers.ReadFeetFile("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_depths_Filtered_Feet.csv")
	fmt.Println("Finished Reading Depths")
	fe := hazard_providers.FathomEvent{Year: 2050, Frequency: 5, Fluvial: true}
	ComputeSingleEvent_NSIStream(ds, "11", fe)
}
func TestMultiEvent(t *testing.T) {
	fmt.Println("Reading Depths")
	ds := hazard_providers.ReadFeetFile("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_depths_Filtered_Feet.csv")
	fmt.Println("Finished Reading Depths")
	db := store.CreateDatabase()
	defer db.Close()
	ComputeMultiEvent_NSIStream(ds, "11", db)
}
func TestMultiEvent_MultiState(t *testing.T) {
	fmt.Println("Reading Depths")
	ds := hazard_providers.ReadFeetFile("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_depths_Filtered_Feet.csv")
	fmt.Println("Finished Reading Depths")
	ComputeMultiFips_MultiEvent(ds)
}
func TestSQLMultiEvent_SingleState(t *testing.T) {
	fmt.Println("Reading Depths")
	ds := hazard_providers.OpenSQLDepthDataSet()
	fmt.Println("Finished Reading Depths")
	db := store.CreateDatabase()
	defer db.Close()
	ComputeMultiEvent_NSIStream(ds, "11", db)
}
func TestSQL_MultiEvent_MultiState(t *testing.T) {
	fmt.Println("Reading Depths")
	ds := hazard_providers.OpenSQLDepthDataSet()
	fmt.Println("Finished Reading Depths")
	ComputeMultiFips_MultiEvent(ds)
}

func TestComputeNewFile(t *testing.T) {
	ss := []string{"17"}
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
	path := fmt.Sprintf("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_depths\\NSI_Fathom_depths%v_feet.csv", ss)
	ds := hazard_providers.ReadFeetFile(path)
	outputpath := fmt.Sprintf("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_depths\\NSI_Fathom_damages_%v.csv", ss)
	outputFile, error := os.Create(outputpath)
	defer outputFile.Close()
	if error != nil {
		panic(error)
	}
	//compute
	ComputeMultiEvent_NSIStream_toFile_withNew(ds, ss, outputFile, true)
}
func TestMulti_State_ComputeNewFile(t *testing.T) {
	fmap := census.StateToCountyFipsMap()
	var wg sync.WaitGroup
	wg.Add(len(fmap))
	for ss, _ := range fmap {
		go func(state string) {
			defer wg.Done()
			computeState_New(state)
		}(ss)
	}
	wg.Wait()
}
