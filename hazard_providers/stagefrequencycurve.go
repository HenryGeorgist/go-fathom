package hazard_providers

import (
	"github.com/HydrologicEngineeringCenter/go-statistics/statistics"
)

type StageFrequencyCurve struct {
	Stages      []statistics.ContinuousDistribution
	Frequencies []float64
}
