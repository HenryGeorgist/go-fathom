package hazard_providers

import (
	"fmt"
	"testing"

	"github.com/USACE/go-consequences/geography"
)

func TestOpenDepths(t *testing.T) {
	MergeSQLDepthNSIDataSet("/workspaces/go-fathom/data/nsiv2_29.gpkg")
}

func TestGetRecord(t *testing.T) {
	data := MergeSQLDepthNSIDataSet("/workspaces/go-fathom/data/nsiv2_29.gpkg")
	x := geography.Location{X: -92.4744865, Y: 40.279474500054, SRID: "47714983"}
	rec, _ := data.getRecord(x)
	fmt.Println(rec)
}
