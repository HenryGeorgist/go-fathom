package main

import (
	"fmt"
	"time"

	"github.com/HenryGeorgist/go-fathom/compute"
	hp "github.com/HenryGeorgist/go-fathom/hazard_providers"
)

func main() {
	// run main file
	//db := store.CreateWALDatabase()
	//defer db.Close()
	start := time.Now()
	ds := hp.MergeSQLDepthNSIDataSet("/workspaces/go-fathom/data/nsiv2_29.gpkg")
	elapsed := time.Since(start)
	fmt.Printf("The merge of NSI and depth data took %s", elapsed)
	fmt.Println()
	//compute.ComputeMultiFips_MultiEvent(ds)
	//compute.ComputeMultiEvent_NSIStream(ds, "29005", db)
	//compute.ComputeMultiEvent_NSIStream(ds, "29", db)
	//fe := hp.FathomEvent{Year: 2050, Frequency: 5, Fluvial: true}
	//	compute.ComputeSingleEvent_NSIStream(ds, "29005", fe)
	compute.ComputeSingleEvent_NSIStreamMonteCarlo(ds, "29005", 10)

}
