package compute

import (
	"fmt"
	"os"

	"github.com/HenryGeorgist/go-fathom/hazard_providers"
	fstructs "github.com/HenryGeorgist/go-fathom/structures"
	"github.com/HydrologicEngineeringCenter/go-statistics/data"
	"github.com/USACE/go-consequences/consequences"
	"github.com/USACE/go-consequences/structures"
)

func ComputeEadByFips(ds hazard_providers.StochasticDataSet, sp consequences.StreamProvider, fips string, outputFile *os.File, iterations int) bool {
	fmt.Println("Computing by fips for " + fips)
	stringifiedHeaders := make([]string, 4)
	stringifiedHeaders[0] = writeHistoHeader("fluv", 2020, iterations)
	stringifiedHeaders[1] = writeHistoHeader("fluv", 2050, iterations)
	stringifiedHeaders[2] = writeHistoHeader("cstl", 2020, iterations)
	stringifiedHeaders[3] = writeHistoHeader("cstl", 2050, iterations)
	outputFile.WriteString(fmt.Sprintf("FD_ID,X,Y,County,CB,OccType,DamCat,foundHt,StructVal,ContVal%s%s%s%s\n", stringifiedHeaders[0], stringifiedHeaders[2], stringifiedHeaders[1], stringifiedHeaders[3]))
	years := [2]int{2020, 2050}
	fluvial := [2]bool{true, false}
	eads := make([][]float64, 4)
	eadsExists := make([]bool, 4)
	index := 0
	otfdm := FoundationDistributionMap()
	focctypes := fstructs.OccupancyTypeMap()
	crawlfhd := otfdm["Craw"]
	fhd := otfdm["Slab"]
	ok := true
	sp.ByFips(fips, func(cr consequences.Receptor) {
		str, sok := cr.(structures.StructureStochastic)
		if iterations != 1 {
			focctype, ook := focctypes[str.OccType.Name]
			if ook {
				str.OccType = focctype
				str.ContVal = consequences.ParameterValue{Value: str.StructVal.CentralTendency() * .4}
			}
			fhd, ok = otfdm[str.FoundType]
			if ok {
				str.FoundHt = consequences.ParameterValue{Value: fhd}
			}
		}

		if sok {
			//check to see if the structure exists for a first "default event"
			index = 0
			_, exists := ds.Data[str.Name] //checking to see if data exists.
			if exists {
				for _, flu := range fluvial {
					if iterations != 1 {
						if !flu {
							if str.FoundType == "Pier" {
								str.FoundHt = consequences.ParameterValue{Value: fhd}
							} else if str.FoundType == "Pile" {
								str.FoundHt = consequences.ParameterValue{Value: fhd}
							}
						} else {
							if str.FoundType == "Pier" {
								str.FoundHt = consequences.ParameterValue{Value: crawlfhd}
							} else if str.FoundType == "Pile" {
								str.FoundHt = consequences.ParameterValue{Value: crawlfhd}
							}
						}
					}

					for _, y := range years {
						sfc, err := ds.ProvideStageFrequencyCurve(str.Name, y, flu)
						if err == nil {
							ih, err := ComputeEadDistribution(sfc, str, iterations) // binWidth, binStart, binEnd, iterations)
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
						eadsum += 1 //results[idx]
						stringifiedhistos[idx] = writeHistoValues(eads[idx])
					} else {
						//results[idx] = 0
						stringifiedhistos[idx] = writeNilHistoValues(eads[idx], iterations)
					}
				}
				//outputFile.WriteString("FD_ID,X,Y,County,CB,OccType,DamCat,foundHt,StructVal,ContVal,PopDay,PopNight,fluv_2020_EAD,cstl_2020_EAD,fluv_2050_EAD,cstl_2050_EAD\n")
				county := str.CBFips[0:5] //county is first five characters of the cb.
				if eadsum > 0 {
					outputFile.WriteString(fmt.Sprintf("%s,%f,%f,%s,%s,%s,%s,%f,%f,%f%s%s%s%s\n", str.Name, str.X, str.Y, county, str.CBFips, str.OccType.Name, str.DamCat, str.FoundHt.CentralTendency(), str.StructVal.CentralTendency(), str.ContVal.CentralTendency(), stringifiedhistos[0], stringifiedhistos[2], stringifiedhistos[1], stringifiedhistos[3]))
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
	ret += fmt.Sprintf(",%s_%v_MeanEAD", fluv, year)
	return ret
}
func writeHistoValues(ih []float64) string {
	ret := ""
	pm := data.CreateProductMoments()
	for _, d := range ih {
		ret += fmt.Sprintf(",%f", d)
		pm.AddObservation(d)
	}
	ret += fmt.Sprintf(",%f", pm.GetMean())
	return ret
}
func writeNilHistoValues(ih []float64, iterations int) string {
	ret := ""
	for i := 0; i < iterations; i++ {
		ret += fmt.Sprintf(",%f", 0.0)
	}
	ret += fmt.Sprintf(",%f", 0.0)
	return ret
}
