package insurance

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/HenryGeorgist/go-fathom/hazard_providers"
	"github.com/USACE/go-consequences/consequences"
	"github.com/USACE/go-consequences/geography"
	"github.com/USACE/go-consequences/hazards"
	sp "github.com/USACE/go-consequences/structureprovider"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

type insuranceresults struct {
	premium         float64
	structdamage    float64
	contdamage      float64
	insuredlosses   float64
	uninsuredlosses float64
}

func ComputeOptimalTriggerPremium(ds hazard_providers.SQLDataSet, fips string, simulations int) map[string][]insuranceresults {
	rmap := make(map[string][]insuranceresults)
	fmt.Println("Downloading NSI by fips " + fips)

	// create array of random frequency events
	simarray := make([]float64, simulations, simulations)
	for simnumber := 0; simnumber < simulations; simnumber++ {
		// set the seed
		rand.Seed(time.Now().UnixNano())
		// random Fathom Event
		randomnumber := rand.Float64()
		simarray[simnumber] = randomnumber
	}
	structdamagesarray := make([]float64, simulations, simulations)

	// initialize the NSI
	nsp := sp.InitGPK("/workspaces/go-fathom/data/nsiv2_29.gpkg", "nsi")

	// Start time
	// start := time.Now()
	num := 0
	nsp.ByFips(fips, func(s consequences.Receptor) {

		errs := 0
		for simnumber := 0; simnumber < simulations; simnumber++ {
			fe := hazard_providers.FathomEvent{Year: 2020, Frequency: int(1 / simarray[simnumber]), Fluvial: true}
			loc := geography.Location{X: s.Location().X, Y: s.Location().Y, SRID: s.Location().SRID}
			fq := hazard_providers.FathomQuery{Location: loc, FathomEvent: fe}
			result, err := ds.ProvideHazard(fq)
			//fmt.Println(err)
			//var results consequences.Results
			if err == nil {
				//structure presumably exists?
				depthevent, okd := result.(hazards.DepthEvent)
				if okd {
					if depthevent.Depth() <= 0 {
						//skip
						structdamagesarray[simnumber] = 0.0
						//fmt.Printf("Damage for freq %d is %d", simarray[simnumber], structdamagesarray[simnumber])
						//fmt.Println()
					} else {
						r := s.Compute(depthevent)

						// need to put structure
						structdamages := r.Result[6].(float64)
						//fmt.Printf("Damage for freq %d is %d", simarray[simnumber], structdamages)
						//fmt.Println()
						//contdamages := r.Result[7].(float64)
						// optimization function

						//if simnumber == 0 {
						//structdamagesarray := make([]float64, 1, simulations)
						//structdamagesarray[0] = structdamages
						//damagesarray := make([]insuranceresults, 1, simulations)
						//damagesarray[0] = damages1
						//rmap[r.Result[0].(string)] = damagesarray
						//} else {
						structdamagesarray[simnumber] = structdamages
						//rmap[r.Result[0].(string)] = append(rmap[r.Result[0].(string)], damages1)
						//rmap[r.Result[0].(string)].Result[6] = rmap[r.Result[0].(string)].Result[6].(float64) + r.Result[6].(float64)
						//rmap[r.Result[0].(string)].Result[7] = rmap[r.Result[0].(string)].Result[7].(float64) + r.Result[7].(float64)
						//}
					}
				}

			} else {
				errs++
			}
		}

		// find optimum value - mean in this quick case
		// MeanOptimization(structdamagesarray, simarray)

		if errs != simulations {
			convertToPoints := func(n int) plotter.XYs {
				pts := make(plotter.XYs, n)
				for i := range pts {
					pts[i].X = simarray[i]
					pts[i].Y = structdamagesarray[i]
				}
				return pts
			}
			scatterData := convertToPoints(len(simarray))
			// scatter plot
			p := plot.New()
			p.Title.Text = "Damage - Frequency Curve"
			p.X.Label.Text = "Frequency"
			p.Y.Label.Text = "Damage ($)"
			p.Add(plotter.NewGrid())

			si, err := plotter.NewScatter(scatterData)
			if err != nil {
				log.Panic(err)
			}
			p.Add(si)
			// p.Legend.Add("scatter", si)

			err = p.Save(200, 200, fmt.Sprintf("img/scatter_%d.png", num))
			num++
			if err != nil {
				log.Panic(err)
			}
		}
		fmt.Println("One building finished")
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

func MeanOptimization(structdamages []float64, frequencies []float64) float64 {
	total := 0.0
	for i := 0; i < len(structdamages); i++ {
		total += structdamages[i]
	}
	mean := total / float64(len(structdamages))

	// now relate this mean damage to the frequency
	// somehow need to go backward here

	return mean
}

func GradientDescentOptimization(structdamages []float64, contdamages []float64, frequencies []float64, epochs int, learning_rate float64) {
	// plan
	// without taking into account others' preferences
	// without taking into account administrative costs
	// make a depth damage curve based on Ollie's frequencies provided
	// premium for just structure = frequency * structure damage
	// the above frequencies would be trigger values
	// WHERE IS THE OPTIMIZATION POTENTIAL HERE

	// gradient descent method

	for epoch := 0; epoch < epochs; epoch++ {
		uninsuredlosses := 0.0
		// grab average annual loss
		aal := 450.0 // a pre-defined number for now

		// random starting premium
		premium := 10.0

		// initialize gradients
		premium_gradient := 0.0

		for i := 0; i < len(structdamages); i++ {

			// I need to optimize both the trigger value and the premium value
			// Right now I am using the Gradient Descent method to optimize the premium value, while holding the trigger value constant
			// I need a joint formula so I can take the partial derivative of each of them

			// trigger

			// initialize the trigger value
			trigger := 10 // the 1 in 10 year flood (10% chance any given year)
			// initialize the payout
			payout := 0.0

			if int(1/frequencies[i]) >= trigger {
				payout = (premium / aal) * structdamages[i]
			}
			// the loss function here is the mean squared error of uninsured losses
			// we want the MSE of uninsured losses because we are a neutral party
			// we don't care who makes the profit... in the end we want the closest to net 0 loses
			uninsuredlosses += (structdamages[i] - payout) * (structdamages[i] - payout)

			two_over_n := 2.0 / float64(len(structdamages))

			// partial derivative wrt the premium
			premium_gradient += -two_over_n * (structdamages[i] - payout) * (structdamages[i] / aal)
		}
		// learning rate
		new_premium := premium - (learning_rate * premium_gradient)

		// set premium to new premium
		premium = new_premium
		fmt.Printf("Total loss of epoch %v was %v", epoch, uninsuredlosses)
		fmt.Println("----------------------------------------------------------------------------")
	}
	// how many times does there need to be a payout
	// # of households paid out, and damaged
	// historical data... how often do I hit those feet of floodimg
	// premiums need to cover full or partial amount of damage
	// how many people pay what premium to cover that?

	// next week, show PDF of flood freq.. damage for elevation step function
}
