package main

import (
	"fmt"
	"time"

	"github.com/HenryGeorgist/go-fathom/compute"
	hp "github.com/HenryGeorgist/go-fathom/hazard_providers"
	"github.com/HenryGeorgist/go-fathom/store"
)

func main() {
	// run main file
	db := store.CreateWALDatabase()
	defer db.Close()
	start := time.Now()
	ds := hp.MergeSQLDepthNSIDataSet("/workspaces/go-fathom/data/nsiv2_29.gpkg")
	elapsed := time.Since(start)
	fmt.Printf("The merge of NSI and depth data took %s", elapsed)
	fmt.Println()
	//compute.ComputeMultiFips_MultiEvent(ds)
	//compute.ComputeMultiEvent_NSIStream(ds, "29005", db)
	compute.ComputeMultiEvent_NSIStream(ds, "29", db)
}
