package compute

import (
	"errors"
	"math/rand"

	"github.com/HenryGeorgist/go-fathom/hazard_providers"
	"github.com/USACE/go-consequences/compute"
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
		//dsfc := sfc.Sample(stageFrequencyRand.Float64())
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
			if sfc.Frequencies[idx] < .4 { //no damage more frequently than the x year (actually 2 and 5 year.)
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
			/*totval := ds.StructVal + ds.ContVal
			dampercent := tdam / (totval)
			if totval == 0 {
				dampercent = 0
			}*/
			realizationDamages[idx] = tdam

		}
		eadEst := compute.ComputeSpecialEAD(realizationDamages, sfc.Frequencies)
		/*if math.IsNaN(eadEst) {
			fmt.Println(fmt.Sprintf("%v", eadEst))
		}
		if eadEst < 0 {
			fmt.Println(fmt.Sprintf("%v", eadEst))
		}
		if eadEst > 1 {
			fmt.Println(fmt.Sprintf("%v", eadEst))
		}*/
		//eaddist.AddObservation(eadEst)
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
