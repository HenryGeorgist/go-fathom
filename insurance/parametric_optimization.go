package insurance

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/HenryGeorgist/go-fathom/hazard_providers"
	"github.com/USACE/go-consequences/consequences"
	"github.com/USACE/go-consequences/geography"
	"github.com/USACE/go-consequences/hazards"
	sp "github.com/USACE/go-consequences/structureprovider"
)

type insuranceresults struct {
	premium         float64
	structdamage    float64
	contdamage      float64
	insuredlosses   float64
	uninsuredlosses float64
}

func ComputeSingleEvent_NSIStreamMonteCarlo(ds hazard_providers.SQLDataSet, fips string, simulations int) map[string][]damages {
	rmap := make(map[string][]insuranceresults)
	fmt.Println("Downloading NSI by fips " + fips)

	// create array of random frequency events
	simarray := make([]float64, 1, simulations)
	for simnumber := 0; simnumber < simulations; simnumber++ {
		// set the seed
		rand.Seed(time.Now().UnixNano())
		// random Fathom Event
		randomnumber := rand.Float64()
		freq := 1 / randomnumber
		simarray[simnumber] = freq
	}
	// initialize the NSI
	nsp := sp.InitGPK("/workspaces/go-fathom/data/nsiv2_29.gpkg", "nsi")

	// Start time
	// start := time.Now()
	nsp.ByFips(fips, func(s consequences.Receptor) {

		for simnumber := 0; simnumber < simulations; simnumber++ {
			fe := hazard_providers.FathomEvent{Year: 2050, Frequency: int(simarray[simnumber]), Fluvial: true}
			loc := geography.Location{X: s.Location().X, Y: s.Location().Y, SRID: s.Location().SRID}
			fq := hazard_providers.FathomQuery{Location: loc, FathomEvent: fe}
			result, err := ds.ProvideHazard(fq)
			//var results consequences.Results
			if err == nil {
				//structure presumably exists?
				depthevent, okd := result.(hazards.DepthEvent)
				if okd {
					if depthevent.Depth() <= 0 {
						//skip
					} else {
						r := s.Compute(depthevent)
						//StructureDamage := r.Result[6].(float64) //based on convention - super risky
						//ContentDamage := r.Result[7].(float64)   //based on convention - super risky
						//if val, ok := rmap[r.Headers[0]]; ok {
						//val.StructureCount += 1
						//val.StructureDamage += r.Result.Result[0].(float64) //based on convention - super risky
						//val.ContentDamage += r.Result.Result[1].(float64)   //based on convention - super risky
						//rmap[str.DamCat] = val
						//} else {
						//rmap[str.DamCat] = comp.SimulationSummaryRow{RowHeader: str.DamCat, StructureCount: 1, StructureDamage: r.Result.Result[0].(float64), ContentDamage: r.Result.Result[1].(float64)}
						//}
						//results.AddResult(r)

						// need to put structure
						damages1 := insuranceresults{r.Result[6].(float64), r.Result[7].(float64)}

						if simnumber == 0 {
							damagesarray := make([]insuranceresults, 1, simulations)
							damagesarray[0] = damages1
							rmap[r.Result[0].(string)] = damagesarray
						} else {
							rmap[r.Result[0].(string)] = append(rmap[r.Result[0].(string)], damages1)
							//rmap[r.Result[0].(string)].Result[6] = rmap[r.Result[0].(string)].Result[6].(float64) + r.Result[6].(float64)
							//rmap[r.Result[0].(string)].Result[7] = rmap[r.Result[0].(string)].Result[7].(float64) + r.Result[7].(float64)
						}
					}
				}

			}
		}

	})

	//rows := make([]consequences.Result, len(rmap))
	//idx := 0
	//s := "COMPLETE FOR SIMULATION" + "\n"
	// for _, val := range rmap {
	// 	fmt.Println(fmt.Sprintf("for %s, there were structures with %f structure damages %f content damages for location %s", fips, val.Result[6], val.Result[7], val.Result[1]))
	// 	//s += fmt.Sprintf("for %s, there were %d structures with %f structure damages %f content damages for damage category %s", fips, val.StructureCount, val.StructureDamage, val.ContentDamage, val.RowHeader) + "\n"
	// 	rows[idx] = val
	// 	idx++
	// }

	fmt.Println("Complete for " + fips)
	return rmap
}
