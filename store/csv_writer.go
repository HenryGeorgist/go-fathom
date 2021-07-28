package store

import (
	"fmt"
	"os"

	"github.com/USACE/go-consequences/consequences"
	"github.com/USACE/go-consequences/structureprovider"
	"github.com/USACE/go-consequences/structures"
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
	nsi := structureprovider.InitNSISP()
	nsi.ByFips(ss, func(str consequences.Receptor) {
		s, sok := str.(structures.StructureStochastic)
		if sok {
			i++
			o := Output{
				Name: s.Name,
				X:    s.X,
				Y:    s.Y,
			}
			f.WriteString(fmt.Sprint(o) + "\n") //write to file
		}

	})
	fmt.Println(fmt.Sprintf("Processed %v structures in state %v.", i, ss))
	return true
}
func ProcessByStateMoreAttributes(ss string, f *os.File) bool {
	fmt.Println("Computing " + ss)
	f.WriteString("FD_ID,X,Y,County,CB,OccType,DamCat,foundHt,StructVal,ContVal,PopDay,PopNight\n")
	i := 0
	nsi := structureprovider.InitNSISP()
	nsi.ByFips(ss, func(str consequences.Receptor) {
		s, sok := str.(structures.StructureStochastic)
		if sok {
			i++
			county := s.CBFips[0:5]
			f.WriteString(fmt.Sprintf("%s,%f,%f,%s,%s,%s,%s,%f,%f,%f,%d,%d\n", s.Name, s.X, s.Y, county, s.CBFips, s.OccType.Name, s.DamCat, s.FoundHt, s.StructVal, s.ContVal, s.Pop2amu65+s.Pop2amo65, s.Pop2pmu65+s.Pop2pmo65))
		}
	})
	fmt.Println(fmt.Sprintf("Processed %v structures in state %v.", i, ss))
	return true
}
