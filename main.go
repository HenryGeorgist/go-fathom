package main

import (
	"github.com/HenryGeorgist/go-fathom/compute"
	hp "github.com/HenryGeorgist/go-fathom/hazard_providers"
)

func main() {
	x := hp.OpenSQLDepthDataSet("/workspaces/go-fathom/data/nsiv2_29.gpkg")
	compute.ComputeMultiFips_MultiEvent(x)
}
