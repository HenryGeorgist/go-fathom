package main

import (
	"github.com/HenryGeorgist/go-fathom/compute"
	hp "github.com/HenryGeorgist/go-fathom/hazard_providers"
)

func main() {
	x := hp.OpenSQLDepthDataSet("./data/fathom-depths.db")
	compute.ComputeMultiFips_MultiEvent(x)
}
