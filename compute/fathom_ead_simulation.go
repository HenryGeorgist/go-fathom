package compute

import (
	"fmt"
	"os"

	"github.com/HenryGeorgist/go-fathom/hazard_providers"
	"github.com/USACE/go-consequences/consequences"
	"github.com/USACE/go-consequences/structures"
)

func ComputeEadByFips(ds hazard_providers.StochasticDataSet, sp consequences.StreamProvider, fips string, outputFile *os.File, newData bool) bool {
	fmt.Println("Computing by fips for " + fips)
	binWidth := .01
	binStart := 0.0
	binEnd := 1.0
	iterations := 20
	stringifiedHeaders := make([]string, 4)
	stringifiedHeaders[0] = writeHistoHeader("fluv", 2020, iterations)
	stringifiedHeaders[1] = writeHistoHeader("fluv", 2050, iterations)
	stringifiedHeaders[2] = writeHistoHeader("cstl", 2020, iterations)
	stringifiedHeaders[3] = writeHistoHeader("cstl", 2050, iterations)
	outputFile.WriteString(fmt.Sprintf("FD_ID,X,Y,County,CB,OccType,DamCat,foundHt,StructVal,ContVal,PopDay,PopNight%s%s%s%s\n", stringifiedHeaders[0], stringifiedHeaders[2], stringifiedHeaders[1], stringifiedHeaders[3]))
	years := [2]int{2020, 2050}
	fluvial := [2]bool{true, false}
	eads := make([][]float64, 4)
	eadsExists := make([]bool, 4)
	index := 0
	//otfdm := FoundationDistributionMap()
	sp.ByFips(fips, func(cr consequences.Receptor) {
		str, sok := cr.(structures.StructureStochastic)
		if sok {
			//check to see if the structure exists for a first "default event"
			index = 0
			_, exists := ds.Data[str.Name] //checking to see if data exists.
			if exists {
				//set foundation height distribution
				//fhdist := otfdm[str.OccType.Name]
				//str.FoundHt = consequences.ParameterValue{Value: fhdist}
				for _, flu := range fluvial {
					for _, y := range years {
						sfc, err := ds.ProvideStageFrequencyCurve(str.Name, y, flu)
						if err == nil {
							ih, err := ComputeEadDistribution(sfc, str, binWidth, binStart, binEnd, iterations)
							eads[index] = ih
							if err != nil {
								eadsExists[index] = false
							} else {
								eadsExists[index] = true
							}
						} else {
							eadsExists[index] = false
						}
						index += 1
					}
				}
				//write to output file.
				//results := make([]float64, 4)
				stringifiedhistos := make([]string, 4)
				eadsum := 0.0
				for idx, b := range eadsExists {
					if b {
						//results[idx] = eads[idx].Mean() * (str.ContVal.CentralTendency() + str.StructVal.CentralTendency())
						//eadsum += results[idx]
						stringifiedhistos[idx] = writeHistoValues(eads[idx])
					} else {
						//results[idx] = 0
						stringifiedhistos[idx] = writeNilHistoValues(eads[idx], iterations)
					}
				}
				//outputFile.WriteString("FD_ID,X,Y,County,CB,OccType,DamCat,foundHt,StructVal,ContVal,PopDay,PopNight,fluv_2020_EAD,cstl_2020_EAD,fluv_2050_EAD,cstl_2050_EAD\n")
				county := str.CBFips[0:5] //county is first five characters of the cb.
				if eadsum > 0 {
					outputFile.WriteString(fmt.Sprintf("%s,%f,%f,%s,%s,%s,%s,%f,%f,%f,%d,%d%s%s%s%s\n", str.Name, str.X, str.Y, county, str.CBFips, str.OccType.Name, str.DamCat, str.FoundHt.CentralTendency(), str.StructVal.CentralTendency(), str.ContVal.CentralTendency(), str.Pop2amu65+str.Pop2amo65, str.Pop2pmu65+str.Pop2pmo65, stringifiedhistos[0], stringifiedhistos[2], stringifiedhistos[1], stringifiedhistos[3]))
				}
			}
		}

	})
	fmt.Println("Completed Computing by fips " + fips)
	return true
}
func writeHistoHeader(fluv string, year int, iterations int) string {
	ret := ""
	for i := 0; i < iterations; i++ {
		ret += fmt.Sprintf(",%s_%v_%v", fluv, year, i+1)
	}
	return ret
}
func writeHistoValues(ih []float64) string {
	ret := ""
	for _, d := range ih {
		ret += fmt.Sprintf(",%f", d)
	}
	return ret
}
func writeNilHistoValues(ih []float64, iterations int) string {
	ret := ""
	for i := 0; i < iterations; i++ {
		ret += fmt.Sprintf(",%f", 0.0)
	}
	return ret
}
