package insurance

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"
	"sort"
	"time"

	"github.com/HenryGeorgist/go-fathom/graphing"
	"github.com/HenryGeorgist/go-fathom/hazard_providers"
	"github.com/USACE/go-consequences/consequences"
	"github.com/USACE/go-consequences/geography"
	"github.com/USACE/go-consequences/hazards"
	"github.com/USACE/go-consequences/paireddata"
	sp "github.com/USACE/go-consequences/structureprovider"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

type InsuranceResults struct {
	Premium         float64
	Trigger         float64
	Totallosses     float64
	Uninsuredlosses float64
}

func ComputeOptimalTriggerPremium(ds hazard_providers.SQLDataSet, fips string, simulations int) map[string]InsuranceResults {
	rmap := make(map[string]InsuranceResults)
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

	// graph the frequencies with a histogram
	graphing.InsuranceFrequencyHistogram(simarray, "Flood Magnitude Histogram", "histogram_freqs.png")

	structdamagesarray := make([]float64, simulations, simulations)

	// initialize the NSI
	nsp := sp.InitGPK("/workspaces/go-fathom/data/nsiv2_29.gpkg", "nsi")

	// Start time
	// start := time.Now()
	num := 0
	economicdamages := 0.0
	uninsuredlosses := 0.0
	insuredlosses := 0.0
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
						structdamages := r.Result[6].(float64)
						structdamagesarray[simnumber] = structdamages
					}
				}

			} else {
				errs++
			}
		}
		//trigger := GradientDescentOptimization(structdamagesarray, simarray, 20, .001)
		payout, premium, trigger := MeanOptimization(structdamagesarray, simarray, s, ds)
		//var data = paireddata.PairedData{simarray, structdamagesarray}
		//payout := data.SampleValue(trigger)

		// find optimum value - mean in this quick case
		// MeanOptimization(structdamagesarray, simarray)
		simarraygraph := make([]float64, simulations, simulations)
		for i := 0; i < len(simarray); i++ {
			simarraygraph[i] = 1 - simarray[i]
		}
		if errs != simulations {
			convertToPoints := func(n int) plotter.XYs {
				pts := make(plotter.XYs, n)
				for i := range pts {
					pts[i].X = simarraygraph[i]
					pts[i].Y = structdamagesarray[i]
				}
				return pts
			}
			makePayout := func(n int) plotter.XYs {
				pts := make(plotter.XYs, n)
				for i := range pts {
					pts[i].X = trigger
					pts[i].Y = payout
				}
				return pts
			}
			scatterData := convertToPoints(len(simarray))
			payoutData := makePayout(1)
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
			py, err := plotter.NewScatter(payoutData)
			if err != nil {
				log.Panic(err)
			}
			py.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
			//py.GlyphStyle.Shape = py.Shape
			p.Add(py)

			err = p.Save(250, 250, fmt.Sprintf("img/scatter_%d.png", num))
			num++
			if err != nil {
				log.Panic(err)
			}
		}
		totallossessum := 0.0
		for i := 0; i < len(structdamagesarray); i++ {
			economicdamages += structdamagesarray[i]
			totallossessum += structdamagesarray[i]
		}
		uninsuredlossessum := 0.0
		for i := 0; i < len(simarray); i++ {
			if trigger <= simarray[i] {
				insuredlosses += structdamagesarray[i]
			} else {
				uninsuredlosses += structdamagesarray[i]
				uninsuredlossessum += structdamagesarray[i]
			}
		}
		rmap[fmt.Sprintf("structure_%d", num)] = InsuranceResults{Premium: premium, Trigger: trigger, Totallosses: totallossessum, Uninsuredlosses: uninsuredlossessum}

		fmt.Println("One building finished")
	})
	fmt.Printf("The Total Economic Damage of %v simulated events is %v", simulations, economicdamages)
	fmt.Println()
	fmt.Printf("The Total Uninsured losses are %v", uninsuredlosses)
	fmt.Println()
	fmt.Printf("The proportion of uninsured to total loses is %v", uninsuredlosses/economicdamages)
	fmt.Println()

	// aggregate graphs for each category
	graphing.InsuranceSummaryHistogram(rmap, "trigger", "Histgram of Trigger Values for All Structures", "histogram_trigger.png")
	graphing.InsuranceSummaryHistogram(rmap, "premium", "Histgram of Premium Values for All Structures", "histogram_premium.png")
	graphing.InsuranceSummaryHistogram(rmap, "total losses", "Histgram of Total Losses Values for All Structures", "histogram_totalloss.png")

	fmt.Println("Complete for " + fips)
	return rmap
}

func MeanOptimization(structdamages []float64, frequencies []float64, s consequences.Receptor, ds hazard_providers.SQLDataSet) (float64, float64, float64) {
	total := 0.0
	for i := 0; i < len(structdamages); i++ {
		total += structdamages[i]
	}
	mean_dam := total / float64(len(structdamages))

	frequencies_sorted := frequencies
	structdamages_sorted := structdamages
	sort.Float64s(frequencies_sorted[:])
	sort.Float64s(structdamages_sorted[:])

	data1 := paireddata.PairedData{Xvals: structdamages_sorted, Yvals: frequencies_sorted}
	// I am not sure why I do this but it works
	freq := 1 - data1.SampleValue(mean_dam)

	premium := mean_dam * freq
	return mean_dam, premium, freq
}

func GradientDescentOptimization(structdamages []float64, frequencies []float64, epochs int, learning_rate float64) float64 {
	// plan
	// without taking into account others' preferences
	// without taking into account administrative costs
	// make a depth damage curve based on Ollie's frequencies provided
	// premium for just structure = frequency * structure damage
	// the above frequencies would be trigger values
	// WHERE IS THE OPTIMIZATION POTENTIAL HERE

	// curve of trigger points (histogram of all 1000 homes) / premiums (histogram of all 1000 homes) / cumulative losses (all homes)
	// summary of the results of the simulation

	// cumulative losses per home -- below line - uninsured lossed per home (maybe see the effectiveness of the optimization)
	//    --- two data points per home
	//    --- two different graphs -- same x-axis

	// gradient descent method
	// initialize the trigger value
	trigger := 0.1 // the 1 in 10 year flood (10% chance any given year)
	for epoch := 0; epoch < epochs; epoch++ {
		uninsuredlosses := 0.0
		// grab average annual loss
		//aal := 450.0 // a pre-defined number for now

		// random starting premium
		//premium := 10.0

		// initialize gradients
		gradient := 0.0

		// need to fit a function to the simulated data points
		var data = paireddata.PairedData{Xvals: frequencies, Yvals: structdamages}

		// copies of structdamages and frequencies to use for sorting
		//frequencies_sorted := frequencies
		//structdamages_sorted := structdamages
		//sort.Float64s(frequencies_sorted[:])
		//sort.Float64s(structdamages_sorted[:])

		indx1 := 0
		indx2 := 0
		maxvaluearr := make([]float64, 2)
		maxvaluearr[0] = 1000000
		maxvaluearr[1] = 1000000
		for j := 0; j < len(frequencies); j++ {
			if math.Abs(frequencies[j]-trigger) < maxvaluearr[0] {
				maxvaluearr[0] = math.Abs(frequencies[j] - trigger)
				indx1 = j
			} else if math.Abs(frequencies[j]-trigger) < maxvaluearr[1] {
				maxvaluearr[1] = math.Abs(frequencies[j] - trigger)
				indx2 = j
			}
		}
		slope := (frequencies[indx1] - frequencies[indx2]) / 100000000.0
		if structdamages[indx1] != 0 || structdamages[indx2] != 0 {
			slope = (frequencies[indx1] - frequencies[indx2]) / (structdamages[indx1] - structdamages[indx2])
		}
		for i := 0; i < len(structdamages); i++ {

			// I need to optimize both the trigger value and the premium value
			// Right now I am using the Gradient Descent method to optimize the premium value, while holding the trigger value constant
			// I need a joint formula so I can take the partial derivative of each of them

			// initialize the payout
			payout := 0.0

			if frequencies[i] <= trigger {
				// payout = (premium / aal) * structdamages[i]
				//payout = structdamages // hmm it equals the structure damage at the point of the trigger value...
				payout = data.SampleValue(trigger)
			}
			// the loss function here is the mean squared error of uninsured losses
			// we want the MSE of uninsured losses because we are a neutral party
			// we don't care who makes the profit... in the end we want the closest to net 0 loses
			uninsuredlosses += (structdamages[i] - payout) * (structdamages[i] - payout)

			two_over_n := 2.0 / float64(len(structdamages))

			// estimate the derivate by averaging the slope of the closest (x,y) pair above the value in question and below the value in question
			// my assumption when I sort both arrays is that they are all in the same order (descending or ascending) so that the values are not jumbled

			//upper_x := sort.SearchFloat64s(frequencies_sorted, trigger)
			//lower_x := upper_x - 1
			// partial derivative wrt the damage-frequency curve
			gradient += -two_over_n * (structdamages[i] - payout) * (slope)
			//fmt.Println(gradient)
		}

		// learning rate
		new_trigger := trigger - (learning_rate * gradient)

		// set premium to new premium
		trigger = new_trigger
		//fmt.Printf("New Trigger Value is %v", new_trigger)
		//fmt.Println()
		//fmt.Printf("Total loss of epoch %v was %v", epoch, uninsuredlosses)
		//fmt.Println("----------------------------------------------------------------------------")
	}
	// how many times does there need to be a payout
	// # of households paid out, and damaged
	// historical data... how often do I hit those feet of floodimg
	// premiums need to cover full or partial amount of damage
	// how many people pay what premium to cover that?

	// next week, show PDF of flood freq.. damage for elevation step function
	return trigger
}
