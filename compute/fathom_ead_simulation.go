package compute

import (
	"fmt"
	"os"

	"github.com/HenryGeorgist/go-fathom/hazard_providers"
	comp "github.com/USACE/go-consequences/compute"
	"github.com/USACE/go-consequences/consequences"
	"github.com/USACE/go-consequences/hazards"
	"github.com/USACE/go-consequences/structures"
)

func ComputeMultiEvent_NSIStream_toFile_withNew(ds hazard_providers.DataSet, sp consequences.StreamProvider, fips string, outputFile *os.File, newData bool) bool {
	fmt.Println("Downloading NSI by fips " + fips)
	outputFile.WriteString("FD_ID,X,Y,County,CB,OccType,DamCat,foundHt,StructVal,ContVal,PopDay,PopNight,fluv_2020_EAD,cstl_2020_EAD,fluv_2050_EAD,cstl_2050_EAD\n")
	years := [2]int{2020, 2050}
	frequencies := []int{5, 20, 100, 250, 500}
	freq := []float64{.2, .05, .01, .004, .002}
	size := 5
	if newData {
		frequencies = []int{2, 5, 20, 100, 250, 500}
		size = 6
		freq = []float64{.5, .2, .05, .01, .004, .002}
	}
	fluvial := [2]bool{true, false}
	//index := 0
	sp.ByFips(fips, func(cr consequences.Receptor) {
		str, sok := cr.(structures.StructureStochastic)
		if sok {
			//check to see if the structure exists for a first "default event"
			fe := hazard_providers.FathomEvent{Year: 2050, Frequency: 500, Fluvial: true}
			fq := hazard_providers.FathomQuery{Fd_id: str.Name, FathomEvent: fe}
			_, err := ds.ProvideHazard(fq)
			if err == nil {
				cfdam := make([]float64, size)
				cpdam := make([]float64, size)
				ffdam := make([]float64, size)
				fpdam := make([]float64, size)
				cfdamc := make([]float64, size)
				cpdamc := make([]float64, size)
				ffdamc := make([]float64, size)
				fpdamc := make([]float64, size)
				for _, flu := range fluvial {
					for _, y := range years {
						for _, f := range frequencies {
							fe = hazard_providers.FathomEvent{Year: y, Frequency: f, Fluvial: flu}
							fq.FathomEvent = fe
							result, _ := ds.ProvideHazard(fq)
							depthevent, okd := result.(hazards.DepthEvent)
							if okd {
								if depthevent.Depth() <= 0 {
									//skip
									assignDamage(flu, y, f, 0, ffdam, fpdam, cfdam, cpdam, newData)
									assignDamage(flu, y, f, 0, ffdamc, fpdamc, cfdamc, cpdamc, newData)
								} else {
									r, err := str.Compute(depthevent)
									if err != nil {
										panic(err)
									}
									sd, err := r.Fetch("structure damage")
									StructureDamage := sd.(float64) //based on convention - super risky
									cd, err := r.Fetch("content damage")
									ContentDamage := cd.(float64) //based on convention - super risky
									assignDamage(flu, y, f, StructureDamage, ffdam, fpdam, cfdam, cpdam, newData)
									assignDamage(flu, y, f, ContentDamage, ffdamc, fpdamc, cfdamc, cpdamc, newData)
								}
							}
						}

					}
				}
				//compute ead's for each of the 4 caases for structure and content.
				cfead := comp.ComputeSpecialEAD(cfdam, freq)
				cpead := comp.ComputeSpecialEAD(cpdam, freq)
				ffead := comp.ComputeSpecialEAD(ffdam, freq)
				fpead := comp.ComputeSpecialEAD(fpdam, freq)

				cfeadc := comp.ComputeSpecialEAD(cfdamc, freq)
				cpeadc := comp.ComputeSpecialEAD(cpdamc, freq)
				ffeadc := comp.ComputeSpecialEAD(ffdamc, freq)
				fpeadc := comp.ComputeSpecialEAD(fpdamc, freq)
				//write to output file.
				//outputFile.WriteString("FD_ID,X,Y,County,CB,OccType,DamCat,foundHt,StructVal,ContVal,PopDay,PopNight,fluv_2020_EAD,cstl_2020_EAD,fluv_2050_EAD,cstl_2050_EAD\n")
				county := str.CBFips[0:5] //county is first five characters of the cb.
				outputFile.WriteString(fmt.Sprintf("%s,%f,%f,%s,%s,%s,%s,%f,%f,%f,%d,%d,%f,%f,%f,%f\n", str.Name, str.X, str.Y, county, str.CBFips, str.OccType.Name, str.DamCat, str.FoundHt, str.StructVal, str.ContVal, str.Pop2amu65+str.Pop2amo65, str.Pop2pmu65+str.Pop2pmo65, cfead+cfeadc, cpead+cpeadc, ffead+ffeadc, fpead+fpeadc))
			}
		}

	})

	fmt.Println("Completed Computing by fips " + fips)
	return true
}
func frequencyIndex(frequency int, newData bool) int {
	if newData {
		switch frequency {
		case 2:
			return 0
		case 5:
			return 1
		case 20:
			return 2
		case 100:
			return 3
		case 250:
			return 4
		case 500:
			return 5
		default:
			return -1 //bad frequency
		}
	}
	switch frequency {
	case 5:
		return 0
	case 20:
		return 1
	case 100:
		return 2
	case 250:
		return 3
	case 500:
		return 4
	default:
		return -1 //bad frequency
	}
}
func assignDamage(fluvial bool, year int, frequency int, damage float64, ffdam []float64, fpdam []float64, cfdam []float64, cpdam []float64, newData bool) {
	if fluvial {
		if year == 2020 {
			cfdam[frequencyIndex(frequency, newData)] = damage
		} else if year == 2050 {
			ffdam[frequencyIndex(frequency, newData)] = damage
		}
	} else {
		if year == 2020 {
			cpdam[frequencyIndex(frequency, newData)] = damage
		} else if year == 2050 {
			fpdam[frequencyIndex(frequency, newData)] = damage
		}
	}

}
