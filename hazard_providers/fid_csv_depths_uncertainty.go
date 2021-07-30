package hazard_providers

import (
	"errors"

	"github.com/HydrologicEngineeringCenter/go-statistics/statistics"
)

type StochasticDataSet struct {
	Data              map[string]Record
	StandardDeviation float64
	Frequencies       []float64 //{.5, .2, .05, .01, .004, .002}
}

func (ds StochasticDataSet) ProvideStageFrequencyCurve(fd_id string, year int, fluvial bool) (StageFrequencyCurve, error) {
	r, found := ds.Data[fd_id]
	if found {
		if year == 2020 {
			if fluvial {
				return generateStageFrequencyCurve(r.CurrentFluvial, ds.StandardDeviation, ds.Frequencies)
			} else {
				return generateStageFrequencyCurve(r.CurrentPluvial, ds.StandardDeviation, ds.Frequencies)
			}
		} else {
			if fluvial {
				return generateStageFrequencyCurve(r.FutureFluvial, ds.StandardDeviation, ds.Frequencies)
			} else {
				return generateStageFrequencyCurve(r.FuturePluvial, ds.StandardDeviation, ds.Frequencies)
			}
		}
	}
	return StageFrequencyCurve{}, errors.New("nope.")
}
func generateStageFrequencyCurve(data FrequencyData, sd float64, frequencies []float64) (StageFrequencyCurve, error) {
	hs := make([]statistics.ContinuousDistribution, len(data.Values))
	for i, d := range data.Values {
		n := statistics.NormalDistribution{Mean: d, StandardDeviation: sd}
		hs[i] = n
	}
	return StageFrequencyCurve{Stages: hs, Frequencies: frequencies}, nil
}
