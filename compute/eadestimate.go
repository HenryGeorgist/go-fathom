package compute

import (
	"math/rand"

	"github.com/HenryGeorgist/go-fathom/hazard_providers"
	"github.com/HydrologicEngineeringCenter/go-statistics/data"
	"github.com/USACE/go-consequences/compute"
	"github.com/USACE/go-consequences/hazards"
	"github.com/USACE/go-consequences/structures"
)

func ComputeEadDistribution(sfc hazard_providers.StageFrequencyCurve, s structures.StructureStochastic) (*data.InlineHistogram, error) {
	eaddist := data.Init(.05, 0, 1) //percent of total value?

	structureSeed := 1234                                                     //create a seed sequence for the structure
	stageFrequencySeed := 4431                                                //create a seed sequence for the stage frequency
	structureRand := rand.New(rand.NewSource(int64(structureSeed)))           //not concurrent safe
	stageFrequencyRand := rand.New(rand.NewSource(int64(stageFrequencySeed))) //not concurrent safe.
	//for some number of iterations
	for i := 0; i < 100; i++ { //just a guess...
		ds := s.SampleStructure(structureRand.Int63()) // sample a structure
		dsfc := sfc.Sample(stageFrequencyRand.Float64())
		realizationDamages := make([]float64, len(dsfc))
		//for each deterministic ordinate compute damage for the sampled structure
		for idx, d := range dsfc {
			de := hazards.DepthEvent{}
			de.SetDepth(d)
			r, err := ds.Compute(de)
			if err != nil {
				panic(err)
			}
			stdam, err := r.Fetch("structure damage")
			if err != nil {
				panic(err)
			}
			condam, err := r.Fetch("content damage")
			if err != nil {
				panic(err)
			}
			sdam := stdam.(float64)
			cdam := condam.(float64)
			tdam := sdam + cdam
			dampercent := tdam / (ds.StructVal + ds.ContVal)
			realizationDamages[idx] = dampercent

		}
		eadEst := compute.ComputeSpecialEAD(realizationDamages, sfc.Frequencies)
		eaddist.AddObservation(eadEst)
	}
	return eaddist, nil
}
