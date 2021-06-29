package graphing

import (
	"github.com/HenryGeorgist/go-fathom/insurance"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

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

func InsuranceSummaryHistogram(vals map[string]insurance.InsuranceResults, category, title, outpath string) {
	// Historgram of sim array
	v := make(plotter.Values, len(vals))
	num := 0
	if category == "trigger" {
		for _, homeresults := range vals {
			v[num] = homeresults.Trigger
		}
	}
	if category == "premium" {
		for _, homeresults := range vals {
			v[num] = homeresults.Trigger
		}
	}
	if category == "total losses" {
		for _, homeresults := range vals {
			v[num] = homeresults.Trigger
		}
	}
	if category == "uninsured losses" {
		for _, homeresults := range vals {
			v[num] = homeresults.Trigger
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
