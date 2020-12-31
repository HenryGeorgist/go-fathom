package compute

import (
	"fmt"
	"testing"

	"github.com/HenryGeorgist/go-fathom/hazard_providers"
	"github.com/HenryGeorgist/go-fathom/store"
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
