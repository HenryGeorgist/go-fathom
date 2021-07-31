package hazard_providers

import (
	"github.com/HydrologicEngineeringCenter/go-statistics/statistics"
)

type StageFrequencyCurve struct {
	Stages      []statistics.ContinuousDistribution
	Frequencies []float64
}

func (s StageFrequencyCurve) Sample(randomValue float64) []float64 {
	stages := make([]float64, len(s.Stages))
	for idx, dist := range s.Stages {
		stages[idx] = dist.InvCDF(randomValue) //holding random value constant enforces consistentcy in stage frequency
	}
	return stages
}
