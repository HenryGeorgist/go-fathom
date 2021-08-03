package hazard_providers

import (
	"fmt"
	"sync"
	"testing"
)

func TestOpen(t *testing.T) {
	ConvertFile("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_depths.csv")
}
func TestFeetFile(t *testing.T) {
	ReadFeetFile("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_depths_Filtered_Feet.csv")
}

func TestWrite(t *testing.T) {
	WriteBackToDisk(DataSet{}, "C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\TestOld.csv", false)
	WriteBackToDisk(DataSet{}, "C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\TestNew.csv", true)
}
func TestConvert(t *testing.T) {
	WriteBackToDisk(ConvertFile("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_depths.csv"), "C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_depths_feet.csv", false)
}
func TestConvertNewFile(t *testing.T) {
	//fmap := census.StateToCountyFipsMap()
	fmap := []string{"11"}
	var wg sync.WaitGroup
	wg.Add(len(fmap))
	for _, ss := range fmap {
		//for ss, _ := range fmap {
		go func(ss string) {
			defer wg.Done()
			fmt.Println("Computing " + ss)
			inpath := fmt.Sprintf("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_Uncertainty\\NSI_Fathom_depths%v.csv", ss)
			outpath := fmt.Sprintf("C:\\Users\\Q0HECWPL\\Documents\\NSI\\NSI_Fathom_depths\\NSI_Fathom_Uncertainty\\NSI_Fathom_depths%v_feet.csv", ss)
			WriteBackToDisk(ConvertFile(inpath), outpath, true)
			fmt.Println(ss + " Complete")
		}(ss)

	}
	wg.Wait()
}
