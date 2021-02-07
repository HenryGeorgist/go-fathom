package store

import (
	"os"
	"sync"
	"testing"
)

func TestJoinAndWriteData(t *testing.T) {
	//get a fips map
	//fmap := census.StateToCountyFipsMap()
	//var wg sync.WaitGroup
	//wg.Add(len(fmap))
	//for ss := range fmap {
	ss := []string{"12", "48", "06"}
	var wg sync.WaitGroup
	wg.Add(len(ss))
	for _, s := range ss {
		go func(state string) {
			defer wg.Done()
			f, err := os.Create("C:\\Examples\\go-fathom\\states\\" + state + "_Attributes.csv")
			if err != nil {
				panic(err)
			}
			defer f.Close()
			ProcessByStateMoreAttributes(state, f)
			f.Sync()
		}(s)
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
