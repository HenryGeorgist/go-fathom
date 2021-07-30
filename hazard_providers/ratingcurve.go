package hazard_providers

import "github.com/USACE/go-consequences/hazards"

type RatingCurve struct {
	Hazards     []hazards.DepthEvent
	Frequencies []float64
}
