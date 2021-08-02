package compute

import (
	"fmt"
	"os"

	"github.com/HenryGeorgist/go-fathom/hazard_providers"
	"github.com/HydrologicEngineeringCenter/go-statistics/data"
	"github.com/USACE/go-consequences/consequences"
	"github.com/USACE/go-consequences/structures"
)

func ComputeEadByFips(ds hazard_providers.StochasticDataSet, sp consequences.StreamProvider, fips string, outputFile *os.File, newData bool) bool {
	fmt.Println("Computing by fips for " + fips)
	binWidth := .05
	binStart := 0.0
	binEnd := 1.0
	iterations := 100
	stringifiedHeaders := make([]string, 4)
	stringifiedHeaders[0] = writeHistoHeader("fluv", 2020, binWidth, binStart, binEnd)
	stringifiedHeaders[1] = writeHistoHeader("fluv", 2050, binWidth, binStart, binEnd)
	stringifiedHeaders[2] = writeHistoHeader("cstl", 2020, binWidth, binStart, binEnd)
	stringifiedHeaders[3] = writeHistoHeader("cstl", 2050, binWidth, binStart, binEnd)
	outputFile.WriteString(fmt.Sprintf("FD_ID,X,Y,County,CB,OccType,DamCat,foundHt,StructVal,ContVal,PopDay,PopNight,fluv_2020_Mean_EAD,cstl_2020_Mean_EAD,fluv_2050_Mean_EAD,cstl_2050_Mean_EAD%s%s%s%s\n", stringifiedHeaders[0], stringifiedHeaders[2], stringifiedHeaders[1], stringifiedHeaders[3]))
	years := [2]int{2020, 2050}
	fluvial := [2]bool{true, false}
	histograms := make([]*data.InlineHistogram, 4)
	histogramExists := make([]bool, 4)
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
							histograms[index] = ih
							if err != nil {
								histogramExists[index] = false
							} else {
								histogramExists[index] = true
							}
						} else {
							histogramExists[index] = false
						}
						index += 1
					}
				}
				//write to output file.
				results := make([]float64, 4)
				stringifiedhistos := make([]string, 4)
				for idx, b := range histogramExists {
					if b {
						results[idx] = histograms[idx].Mean() * (str.ContVal.CentralTendency() + str.StructVal.CentralTendency())
						stringifiedhistos[idx] = writeHistoValues(histograms[idx], iterations, str, binWidth, binStart, binEnd)
					} else {
						results[idx] = 0
						stringifiedhistos[idx] = writeNilHistoValues(histograms[idx], iterations, str, binWidth, binStart, binEnd)
					}
				}
				//outputFile.WriteString("FD_ID,X,Y,County,CB,OccType,DamCat,foundHt,StructVal,ContVal,PopDay,PopNight,fluv_2020_EAD,cstl_2020_EAD,fluv_2050_EAD,cstl_2050_EAD\n")
				county := str.CBFips[0:5] //county is first five characters of the cb.

				outputFile.WriteString(fmt.Sprintf("%s,%f,%f,%s,%s,%s,%s,%f,%f,%f,%d,%d,%f,%f,%f,%f%s%s%s%s\n", str.Name, str.X, str.Y, county, str.CBFips, str.OccType.Name, str.DamCat, str.FoundHt.CentralTendency(), str.StructVal.CentralTendency(), str.ContVal.CentralTendency(), str.Pop2amu65+str.Pop2amo65, str.Pop2pmu65+str.Pop2pmo65, results[0], results[2], results[1], results[3], stringifiedhistos[0], stringifiedhistos[2], stringifiedhistos[1], stringifiedhistos[3]))
			}
		}

	})
	fmt.Println("Completed Computing by fips " + fips)
	return true
}
func writeHistoHeader(fluv string, year int, binWidth float64, binStart float64, binEnd float64) string {
	loc := binStart + binWidth
	binEndEpsilon := binEnd + (binWidth * .5)
	ret := ""
	for loc <= binEndEpsilon {
		ret += fmt.Sprintf(",%s_%v_%f", fluv, year, loc)
		loc += binWidth
	}
	return ret
}
func writeHistoValues(ih *data.InlineHistogram, iterations int, s structures.StructureStochastic, binWidth float64, binStart float64, binEnd float64) string {
	ret := ""
	totval := s.StructVal.CentralTendency() + s.ContVal.CentralTendency()
	loc := binStart + binWidth
	binEndEpsilon := binEnd + (binWidth * .5)
	for loc <= binEndEpsilon {
		ret += fmt.Sprintf(",%f", totval*ih.InvCDF(loc))
		loc += binWidth
	}
	return ret
}
func writeNilHistoValues(ih *data.InlineHistogram, iterations int, s structures.StructureStochastic, binWidth float64, binStart float64, binEnd float64) string {
	ret := ""
	loc := binStart + binWidth
	binEndEpsilon := binEnd + (binWidth * .5)
	for loc <= binEndEpsilon {
		ret += fmt.Sprintf(",%f", 0.0)
		loc += binWidth
	}
	return ret
}
