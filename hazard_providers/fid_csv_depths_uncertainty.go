package hazard_providers

import (
	"errors"

	"github.com/USACE/go-consequences/hazards"
)

type StochasticDataSet struct {
	Data              map[string]Record
	Mean              float64
	StandardDeviation float64
}

func (ds StochasticDataSet) ProvideFutureFluvialRatingCurve(fd_id string) (RatingCurve, error) {

	r, found := ds.Data[fd_id]
	if found {
		return generateRatingCurve(r.FutureFluvial, true)
	}
	return RatingCurve{}, errors.New("nope.")
}
func generateRatingCurve(data FrequencyData, newData bool) (RatingCurve, error) {
	hs := make([]hazards.DepthEvent, len(data.Values))
	fs := make([]float64, len(data.Values))
	for i, d := range data.Values {
		h := hazards.DepthEvent{}
		h.SetDepth(d)
		hs[i] = h
		fs[i] = .05 //just to start - need to define an array of frequencies.
	}
	return RatingCurve{Hazards: hs, Frequencies: fs}, nil
}
