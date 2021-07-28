package compute

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	"github.com/HenryGeorgist/go-fathom/hazard_providers"
	"github.com/HenryGeorgist/go-fathom/store"
	"github.com/USACE/go-consequences/census"
	comp "github.com/USACE/go-consequences/compute"
	"github.com/USACE/go-consequences/consequences"
	"github.com/USACE/go-consequences/hazards"
	"github.com/USACE/go-consequences/structureprovider"
	"github.com/USACE/go-consequences/structures"
)

func ComputeMultiFips_MultiEvent(ds hazard_providers.DataSet) {
	db := store.CreateWALDatabase()
	defer db.Close()
	fmap := census.StateToCountyFipsMap()
	var wg sync.WaitGroup
	wg.Add(len(fmap))
	for ss, _ := range fmap {
		go func(state string) {
			defer wg.Done()
			ComputeMultiEvent_NSIStream(ds, state, db)
		}(ss)
	}
	wg.Wait()
}
func ComputeMultiEvent_NSIStream(ds hazard_providers.DataSet, fips string, db *sql.DB) bool {
	fmt.Println("Downloading NSI by fips " + fips)
	years := [2]int{2020, 2050}
	frequencies := [5]int{5, 20, 100, 250, 500} //for new data, this should be {2,5,20, 100, 250,500}//
	fluvial := [2]bool{true, false}
	tx, _ := db.Begin()
	index := 0
	maxTransaction := 1000
	//transaction := make([]interface{}, maxTransaction)
	nsi := structureprovider.InitNSISP()
	nsi.ByFips(fips, func(cr consequences.Receptor) {
		str, sok := cr.(structures.StructureStochastic)
		if sok {
			//check to see if the structure exists for a first "default event"
			fe := hazard_providers.FathomEvent{Year: 2050, Frequency: 500, Fluvial: true}
			fq := hazard_providers.FathomQuery{Fd_id: str.Name, FathomEvent: fe}
			_, err := ds.ProvideHazard(fq)
			if err == nil {
				//structure presumably exists?
				cfdam := make([]float64, 5) //for new data this needs to be 6//
				cpdam := make([]float64, 5)
				ffdam := make([]float64, 5)
				fpdam := make([]float64, 5)
				cfdamc := make([]float64, 5)
				cpdamc := make([]float64, 5)
				ffdamc := make([]float64, 5)
				fpdamc := make([]float64, 5)
				for _, flu := range fluvial {
					//hazard := "pluvial"
					//if flu {
					//hazard = "fluvial"
					//}
					for _, y := range years {
						for _, f := range frequencies {

							fe = hazard_providers.FathomEvent{Year: y, Frequency: f, Fluvial: flu}
							fq.FathomEvent = fe
							result, _ := ds.ProvideHazard(fq)
							depthevent, okd := result.(hazards.DepthEvent)
							if okd {
								if depthevent.Depth() <= 0 {
									//skip
									assignDamage(flu, y, f, 0, ffdam, fpdam, cfdam, cpdam, false)
									assignDamage(flu, y, f, 0, ffdamc, fpdamc, cfdamc, cpdamc, false)
								} else {
									r, err := str.Compute(depthevent)
									if err != nil {
										panic(err)
									}
									sd, err := r.Fetch("structure damage")
									StructureDamage := sd.(float64) //based on convention - super risky
									cd, err := r.Fetch("content damage")
									ContentDamage := cd.(float64) //based on convention - super risky
									assignDamage(flu, y, f, StructureDamage, ffdam, fpdam, cfdam, cpdam, false)
									assignDamage(flu, y, f, ContentDamage, ffdamc, fpdamc, cfdamc, cpdamc, false)
									//transaction[index] = store.CreateResult(str.Name, y, hazard, fmt.Sprint(f), StructureDamage, ContentDamage)
									//index++
									//store.WriteToDatabase(stmt, str.Name, y, hazard, fmt.Sprint(f), StructureDamage, ContentDamage)
									//store.WriteToTransaction(tx, str.Name, y, hazard, fmt.Sprint(f), StructureDamage, ContentDamage)
									//if index >= maxTransaction {
									//store.WriteArrayToDatabase(db, transaction)
									//index = 0
									//}
								}
							}
						}

					}
				}

				//compute ead's for each of the 4 caases for structure and content.
				freq := []float64{.2, .05, .01, .004, .002} //5,20,100,250,500
				//freq:= []float64{.5,.2,.05,.01,004,.002}//for newData
				cfead := comp.ComputeSpecialEAD(cfdam, freq)
				cpead := comp.ComputeSpecialEAD(cpdam, freq)
				ffead := comp.ComputeSpecialEAD(ffdam, freq)
				fpead := comp.ComputeSpecialEAD(fpdam, freq)

				cfeadc := comp.ComputeSpecialEAD(cfdamc, freq)
				cpeadc := comp.ComputeSpecialEAD(cpdamc, freq)
				ffeadc := comp.ComputeSpecialEAD(ffdamc, freq)
				fpeadc := comp.ComputeSpecialEAD(fpdamc, freq)
				if cfead > cpead {
					//transaction[index] = store.CreateResult(str.Name, str.X, str.Y, fips, years[0], "fluvial", "EAD", cfead, cfeadc)
					store.WriteToTransaction(tx, str.Name, str.X, str.Y, fips, years[0], "fluvial", "EAD", cfead, cfeadc)
					index++ //what if we exceed 500...
					if index >= maxTransaction {
						//store.WriteArrayToDatabase(db, transaction)
						//store.WriteArrayToTransaction(tx, transaction)
						tx.Commit()
						tx, _ = db.Begin()
						index = 0
					}
				} else {
					if cpead > 0.0 { //should we exclude zero ead for one year but not the other?
						//transaction[index] = store.CreateResult(str.Name, str.X, str.Y, fips, years[0], "pluvial", "EAD", cpead, cpeadc)
						store.WriteToTransaction(tx, str.Name, str.X, str.Y, fips, years[0], "pluvial", "EAD", cpead, cpeadc)
						index++
						if index >= maxTransaction {
							//store.WriteArrayToDatabase(db, transaction)
							//store.WriteArrayToTransaction(tx, transaction)
							tx.Commit()
							tx, _ = db.Begin()
							index = 0
						}
					}

				}
				if ffead > fpead {
					//transaction[index] = store.CreateResult(str.Name, str.X, str.Y, fips, years[1], "fluvial", "EAD", ffead, ffeadc)
					store.WriteToTransaction(tx, str.Name, str.X, str.Y, fips, years[1], "fluvial", "EAD", ffead, ffeadc)
					index++
					if index >= maxTransaction {
						//store.WriteArrayToDatabase(db, transaction)
						//store.WriteArrayToTransaction(tx, transaction)
						tx.Commit()
						tx, _ = db.Begin()
						index = 0
					}
				} else {
					if fpead > 0.0 {
						//transaction[index] = store.CreateResult(str.Name, str.X, str.Y, fips, years[1], "pluvial", "EAD", fpead, fpeadc)
						store.WriteToTransaction(tx, str.Name, str.X, str.Y, fips, years[1], "pluvial", "EAD", fpead, fpeadc)
						index++
						if index >= maxTransaction {
							//store.WriteArrayToDatabase(db, transaction)
							//store.WriteArrayToTransaction(tx, transaction)
							tx.Commit()
							tx, _ = db.Begin()
							index = 0
						}
					}

				}
			}
		}

	})
	if index > 0 {
		//smalltransaction := transaction[0 : index-1]
		//store.WriteArrayToDatabase(db, smalltransaction)
		tx.Commit()
		//store.WriteArrayToTransaction(tx, smalltransaction)
		index = 0
	}
	//tx.Commit()
	fmt.Println("Completed Computing by fips " + fips)
	return true
}
func ComputeMultiEvent_NSIStream_toFile_withNew(ds hazard_providers.DataSet, fips string, outputFile *os.File, newData bool) bool {
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
	nsi := structureprovider.InitNSISP()
	nsi.ByFips(fips, func(cr consequences.Receptor) {
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

/*
func ComputeSingleEvent_NSIStream(ds hazard_providers.DataSet, fips string, fe hazard_providers.FathomEvent) {
	rmap := make(map[string]comp.SimulationSummaryRow)
	fmt.Println("Downloading NSI by fips " + fips)
	nsi.GetByFipsStream(fips, func(feature nsi.NsiFeature) {
		m := structures.OccupancyTypeMap()
		defaultOcctype := m["RES1-1SNB"]
		str := comp.NsiFeaturetoStructure(feature, m, defaultOcctype)
		fq := hazard_providers.FathomQuery{Fd_id: str.Name, FathomEvent: fe}
		result, err := ds.ProvideHazard(fq)
		if err == nil {
			//structure presumably exists?
			depthevent, okd := result.(hazards.DepthEvent)
			if okd {
				if depthevent.Depth() <= 0 {
					//skip
				} else {
					r := str.Compute(depthevent)
					if val, ok := rmap[str.DamCat]; ok {
						val.StructureCount += 1
						val.StructureDamage += r.Result.Result[0].(float64) //based on convention - super risky
						val.ContentDamage += r.Result.Result[1].(float64)   //based on convention - super risky
						rmap[str.DamCat] = val
					} else {
						rmap[str.DamCat] = comp.SimulationSummaryRow{RowHeader: str.DamCat, StructureCount: 1, StructureDamage: r.Result.Result[0].(float64), ContentDamage: r.Result.Result[1].(float64)}
					}
				}
			}

		}

	})
	rows := make([]comp.SimulationSummaryRow, len(rmap))
	idx := 0
	//s := "COMPLETE FOR SIMULATION" + "\n"
	for _, val := range rmap {
		fmt.Println(fmt.Sprintf("for %s, there were %d structures with %f structure damages %f content damages for damage category %s", fips, val.StructureCount, val.StructureDamage, val.ContentDamage, val.RowHeader))
		//s += fmt.Sprintf("for %s, there were %d structures with %f structure damages %f content damages for damage category %s", fips, val.StructureCount, val.StructureDamage, val.ContentDamage, val.RowHeader) + "\n"
		rows[idx] = val
		idx++
	}

	fmt.Println("Complete for" + fips)
}*/
