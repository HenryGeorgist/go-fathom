package compute

import (
	"errors"
	"math/rand"

	"github.com/HenryGeorgist/go-fathom/hazard_providers"
	"github.com/USACE/go-consequences/hazards"
	"github.com/USACE/go-consequences/structures"
)

func ComputeEadDistribution(sfc hazard_providers.StageFrequencyCurve, s structures.StructureStochastic, iterations int) ([]float64, error) { //, binWidth float64, binStart float64, binEnd float64, iterations int) ([]float64, error) {
	//eaddist := data.Init(binWidth, binStart, binEnd) //percent of total value?
	if iterations == 1 {
		s.UseUncertainty = false
	} else {
		s.UseUncertainty = true
	}
	eadlist := make([]float64, iterations)
	structureSeed := 1234 //create a seed sequence for the structure
	//stageFrequencySeed := 4431                                                //create a seed sequence for the stage frequency
	structureRand := rand.New(rand.NewSource(int64(structureSeed))) //not concurrent safe
	//stageFrequencyRand := rand.New(rand.NewSource(int64(stageFrequencySeed))) //not concurrent safe.
	//for some number of iterations
	counter := 0
	for i := 0; i < iterations; i++ {
		ds := s.SampleStructure(structureRand.Int63()) // sample a structure
		if ds.NumStories > 3 {
			floorsval := ds.StructVal / float64(ds.NumStories)
			floorcval := ds.ContVal / float64(ds.NumStories)
			ds.StructVal = floorsval * 2
			ds.ContVal = floorcval * 2
		}
		dsfc := sfc.Sample(float64(float64(i)+0.5) / float64(iterations))
		realizationDamages := make([]float64, len(dsfc))
		//for each deterministic ordinate compute damage for the sampled structure
		for idx, d := range dsfc {
			var stdam interface{}
			var condam interface{}
			stdam = 0.0
			condam = 0.0
			if sfc.Frequencies[idx] < .5 { //no damage more frequently than the x year
				if d > 0 {
					de := hazards.DepthEvent{}
					de.SetDepth(d)
					r, err := ds.Compute(de)
					if err != nil {
						panic(err)
					}
					stdam, err = r.Fetch("structure damage")
					if err != nil {
						panic(err)
					}
					condam, err = r.Fetch("content damage")
					if err != nil {
						panic(err)
					}
				}

			}

			sdam := stdam.(float64)
			cdam := condam.(float64)
			tdam := sdam + cdam
			realizationDamages[idx] = tdam

		}
		eadEst := ComputeSpecialEAD(realizationDamages, sfc.Frequencies)
		if eadEst == 0 {
			counter += 1
		}
		eadlist[i] = eadEst
	}
	if counter == iterations {
		return eadlist, errors.New("no damages detected.")
	}
	return eadlist, nil
}

//ComputeSpecialEAD integrates under the damage frequency curve but does not calculate the first triangle between 1 and the first frequency.
func ComputeSpecialEAD(damages []float64, freq []float64) float64 {
	//this differs from computeEAD in that it specifically does not calculate the first triangle between 1 and the first frequency to interpolate damages to zero.
	if len(damages) != len(freq) {
		panic("frequency curve is unbalanced")
	}
	triangle := 0.0
	square := 0.0
	x1 := freq[0]
	y1 := damages[0]
	eadT := 0.0
	twentyYearIndex := -1
	twentyYearHasDamages := false
	if len(damages) > 1 {
		for i := 1; i < len(freq); i++ {
			xdelta := x1 - freq[i]

			square = xdelta * y1

			if freq[i] == .05 {
				twentyYearIndex = i
				if damages[i] > 0 {
					twentyYearHasDamages = true
				}
			}
			if square != 0.0 { //we dont know where damage really begins until we see it. we can guess it is inbetween ordinates, but who knows.
				triangle = ((xdelta) * -(y1 - damages[i])) / 2.0
			} else {
				triangle = 0.0
			}
			eadT += square + triangle
			x1 = freq[i]
			y1 = damages[i]
		}
	}
	if x1 != 0.0 {
		xdelta := x1 - 0.0
		eadT += xdelta * y1 //no extrapolation, just continue damages out as if it were truth for all remaining probability.
	}
	if twentyYearHasDamages {
		base := .1 - .05
		base = .5 * base //1/2 base times height.
		height := damages[twentyYearIndex]
		triangle := base * height
		eadT += triangle
	}
	return eadT
}
