package main

import (
	"github.com/HenryGeorgist/go-fathom/compute"
	hp "github.com/HenryGeorgist/go-fathom/hazard_providers"
	"github.com/HenryGeorgist/go-fathom/store"
)

func main() {
	// run main file
	db := store.CreateWALDatabase()
	defer db.Close()
	ds := hp.OpenSQLDepthDataSet("/workspaces/go-fathom/data/nsiv2_29.gpkg")
	//compute.ComputeMultiFips_MultiEvent(ds)
	compute.ComputeMultiEvent_NSIStream(ds, "29005", db)
}
