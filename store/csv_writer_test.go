package store

import (
	"os"
	"sync"
	"testing"

	"github.com/USACE/go-consequences/census"
)

func TestJoinAndWriteData(t *testing.T) {
	//get a fips map
	fmap := census.StateToCountyFipsMap()
	var wg sync.WaitGroup
	wg.Add(len(fmap))
	for ss := range fmap {
		defer wg.Done()
		f, err := os.Create("C:\\Examples\\go-fathom\\states\\" + ss + ".csv")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		go ProcessByState(ss, f)
		f.Sync()
	}
	wg.Wait()
}
func TestJoinAndWriteData_State(t *testing.T) {
	ss := "17"
	f, err := os.Create("C:\\Examples\\go-fathom\\states\\" + ss + "_rerun.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	ProcessByState(ss, f)
	f.Sync()

}
