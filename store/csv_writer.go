package store

import (
	"fmt"
	"os"

	"github.com/USACE/go-consequences/nsi"
)

type Output struct {
	Name string
	X    float64
	Y    float64
}

func (o Output) String() string {
	return fmt.Sprintf("%v, %f, %f", o.Name, o.X, o.Y)
}
func ProcessByState(ss string, f *os.File) bool {
	fmt.Println("Computing " + ss)
	f.WriteString("FD_ID, X, Y\n")
	i := 0
	nsi.GetByFipsStream(ss, func(str nsi.NsiFeature) {
		i++
		o := Output{
			Name: str.Properties.Name,
			X:    str.Properties.X,
			Y:    str.Properties.Y,
		}
		f.WriteString(fmt.Sprint(o) + "\n") //write to file
	})
	fmt.Println(fmt.Sprintf("Processed %v structures in state %v.", i, ss))
	return true
}
