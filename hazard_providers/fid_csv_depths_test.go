package hazard_providers

import (
	"fmt"
	"testing"

	"github.com/USACE/go-consequences/geography"
	"github.com/USACE/go-consequences/hazards"
)

// import (
// 	"fmt"
// 	"sync"
// 	"testing"

// 	"github.com/USACE/go-consequences/hazards"
// )

// func TestOpen(t *testing.T) {
// 	ConvertFile("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_depths.csv")
// }
// func TestFeetFile(t *testing.T) {
// 	ReadFeetFile("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_depths_Filtered_Feet.csv")
// }
// func TestConvertToSqlite(t *testing.T) {
// 	ReadFeetFile("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_depths_Filtered_Feet.csv").WriteToSqlite()
// }
func TestReadSqliteEvent(t *testing.T) {
	db := OpenSQLDepthDataSet("/workspaces/go-fathom/data/nsiv2_29.gpkg")
	fe := FathomEvent{Year: 2050, Frequency: 500, Fluvial: true}
	fq := FathomQuery{Location: geography.Location{X: -92.632957458, Y: 40.208374023, SRID: "47715104"}, FathomEvent: fe}
	h, _ := db.ProvideHazard(fq)
	depthevent, _ := h.(hazards.DepthEvent)
	fmt.Println(depthevent.DepthValue)
}

// func TestWrite(t *testing.T) {
// 	WriteBackToDisk(DataSet{}, "C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\TestOld.csv", false)
// 	WriteBackToDisk(DataSet{}, "C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\TestNew.csv", true)
// }
// func TestConvert(t *testing.T) {
// 	WriteBackToDisk(ConvertFile("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_depths.csv"), "C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_depths_feet.csv", false)
// }
// func TestConvertNewFile(t *testing.T) {
// 	//fmap := census.StateToCountyFipsMap()
// 	fmap := []string{"17"}
// 	var wg sync.WaitGroup
// 	wg.Add(len(fmap))
// 	for _, ss := range fmap {
// 		//for ss, _ := range fmap {
// 		go func(ss string) {
// 			defer wg.Done()
// 			fmt.Println("Computing " + ss)
// 			inpath := fmt.Sprintf("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_depths_Pluvial\\NSI_Fathom_depths%v.csv", ss)
// 			outpath := fmt.Sprintf("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_depths_Pluvial\\NSI_Fathom_depths%v_feet.csv", ss)
// 			WriteBackToDisk(ConvertFile(inpath), outpath, true)
// 			fmt.Println(ss + " Complete")
// 		}(ss)

// 	}
// 	wg.Wait()
// }
