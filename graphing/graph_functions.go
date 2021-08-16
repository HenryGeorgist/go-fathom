package graphing

import (
	"fmt"

	"github.com/HenryGeorgist/go-fathom/hazard_providers"
	"github.com/USACE/go-consequences/consequences"
	"github.com/USACE/go-consequences/geography"
	"github.com/USACE/go-consequences/hazards"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type InsuranceResults struct {
	Premium         float64
	Trigger         float64
	Totallosses     float64
	Uninsuredlosses float64
}

func InsuranceFrequencyHistogram(vals []float64, title string, outpath string) {
	// Historgram of sim array
	v := make(plotter.Values, len(vals))
	for i := range v {
		v[i] = 1.0 / vals[i]
		if v[i] > 200 {
			v[i] = 100
		}
	}
	// Make a plot and set its title.
	pp := plot.New()

	pp.Title.Text = title

	// Create a histogram of our values drawn
	// from the standard normal.
	h, err := plotter.NewHist(v, 16)
	if err != nil {
		panic(err)
	}
	// Normalize the area under the histogram to
	// sum to one.
	h.Normalize(1)
	pp.Add(h)

	// The normal distribution function
	// norm := plotter.NewFunction(distuv.UnitNormal.Prob)
	// norm.Color = color.RGBA{R: 255, A: 255}
	// norm.Width = vg.Points(2)
	// pp.Add(norm)

	// Save the plot to a PNG file.
	if err := pp.Save(4*vg.Inch, 4*vg.Inch, outpath); err != nil {
		panic(err)
	}
}

func InsuranceFrequencyHistogramFiltered(vals []float64, title string, outpath string) {
	// Historgram of sim array
	v := make(plotter.Values, len(vals))
	for i := range v {
		if vals[i] < 0.3 {
			v[i] = 1.0 / vals[i]
		}
	}
	// Make a plot and set its title.
	pp := plot.New()

	pp.Title.Text = title

	// Create a histogram of our values drawn
	// from the standard normal.
	h, err := plotter.NewHist(v, 32)
	if err != nil {
		panic(err)
	}
	// Normalize the area under the histogram to
	// sum to one.
	h.Normalize(1)
	pp.Add(h)

	// The normal distribution function
	// norm := plotter.NewFunction(distuv.UnitNormal.Prob)
	// norm.Color = color.RGBA{R: 255, A: 255}
	// norm.Width = vg.Points(2)
	// pp.Add(norm)

	// Save the plot to a PNG file.
	if err := pp.Save(4*vg.Inch, 4*vg.Inch, outpath); err != nil {
		panic(err)
	}
}

func InsuranceSummaryHistogram(vals map[string]InsuranceResults, category, title, outpath string) {
	// Historgram of sim array
	v := make(plotter.Values, len(vals))
	num := 0
	if category == "trigger" {
		for _, homeresults := range vals {
			if homeresults.Trigger < 0.3 {
				v[num] = homeresults.Trigger
				num++
			}
		}
	}
	if category == "premium" {
		for _, homeresults := range vals {
			if homeresults.Premium < 3000 {
				v[num] = homeresults.Premium
				num++
			}
		}
	}
	if category == "total losses" {
		for _, homeresults := range vals {
			if homeresults.Totallosses < 15000000 {
				v[num] = homeresults.Totallosses
				num++
			}
		}
	}
	if category == "uninsured losses" {
		for _, homeresults := range vals {
			v[num] = homeresults.Uninsuredlosses
			num++
		}
	}
	// Make a plot and set its title.
	pp := plot.New()

	pp.Title.Text = title

	// Create a histogram of our values drawn
	// from the standard normal.
	h, err := plotter.NewHist(v, 32)
	if err != nil {
		panic(err)
	}
	// Normalize the area under the histogram to
	// sum to one.
	h.Normalize(1)
	pp.Add(h)

	// The normal distribution function
	// norm := plotter.NewFunction(distuv.UnitNormal.Prob)
	// norm.Color = color.RGBA{R: 255, A: 255}
	// norm.Width = vg.Points(2)
	// pp.Add(norm)

	// Save the plot to a PNG file.
	if err := pp.Save(4*vg.Inch, 4*vg.Inch, outpath); err != nil {
		panic(err)
	}
}

func InsuranceSummaryHistogramFiltered(vals map[string]InsuranceResults, category, title, outpath string) {
	// Historgram of sim array
	//v := make(plotter.Values, len(vals))
	var v plotter.Values
	num := 0
	if category == "trigger" {
		for _, homeresults := range vals {
			if homeresults.Trigger < 0.5 {
				v[num] = homeresults.Trigger
				//v.append(homeresults.Trigger)
				num++
			}
		}
	}
	if category == "premium" {
		for _, homeresults := range vals {
			if homeresults.Premium > 1 {
				v[num] = homeresults.Premium
				num++
			}
		}
	}
	if category == "total losses" {
		for _, homeresults := range vals {
			if homeresults.Totallosses > 1 {
				v[num] = homeresults.Totallosses
				num++
			}
		}
	}
	if category == "uninsured losses" {
		for _, homeresults := range vals {
			if homeresults.Uninsuredlosses > 1 {
				v[num] = homeresults.Uninsuredlosses
				num++
			}
		}
	}
	// Make a plot and set its title.
	pp := plot.New()

	pp.Title.Text = title

	// Create a histogram of our values drawn
	// from the standard normal.
	h, err := plotter.NewHist(v, 32)
	if err != nil {
		panic(err)
	}
	// Normalize the area under the histogram to
	// sum to one.
	h.Normalize(1)
	pp.Add(h)

	// The normal distribution function
	// norm := plotter.NewFunction(distuv.UnitNormal.Prob)
	// norm.Color = color.RGBA{R: 255, A: 255}
	// norm.Width = vg.Points(2)
	// pp.Add(norm)

	// Save the plot to a PNG file.
	if err := pp.Save(4*vg.Inch, 4*vg.Inch, outpath); err != nil {
		panic(err)
	}
}

func ExampleGraphsPPTX(s consequences.Receptor, ds hazard_providers.SQLDataSet, fd_id string) {
	// Here I am choosing two examples to highlight the process in the PPTX
	// These are two FD_IDs from Camden County, MO, the most floodprone county in MO according to First Street Foundation
	// One (47920553) has extensive flooding depths and damage
	// The ohter (47926618) does not... it has essentially no flood risk/damage

	if fd_id == "47920553" || fd_id == "47926618" {
		events := [2]int{20, 100}
		for _, hazard_freq := range events {
			fe := hazard_providers.FathomEvent{Year: 2020, Frequency: hazard_freq, Fluvial: true}
			loc := geography.Location{X: s.Location().X, Y: s.Location().Y, SRID: s.Location().SRID}
			fq := hazard_providers.FathomQuery{Location: loc, FathomEvent: fe}
			result, err := ds.ProvideHazard(fq)

			if err == nil {
				//structure presumably exists?
				depthevent, okd := result.(hazards.DepthEvent)
				if okd {
					if depthevent.Depth() <= 0 {
						//skip

					} else {
						r, err := s.Compute(depthevent)
						if err == nil {
							depth := r.Result[3].(hazards.DepthEvent).Depth()
							structdamages := r.Result[6].(float64)
							fmt.Printf("Depth of flooding and corresponding damage for the 1 in %v year for fd_id %v flood is %v and %v, respectively", hazard_freq, fd_id, depth, structdamages)
							fmt.Println()
						}

					}
				}

			}
		}
	}

}
