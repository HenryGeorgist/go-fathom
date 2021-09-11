package compute

import "github.com/HydrologicEngineeringCenter/go-statistics/statistics"

func FoundationDistributionMap() map[string]statistics.ContinuousDistribution {
	m := make(map[string]statistics.ContinuousDistribution)
	m["Slab"] = statistics.TriangularDistribution{Min: 0, Max: 1.5, MostLikely: .5}
	m["Craw"] = statistics.TriangularDistribution{Min: 0, Max: 4, MostLikely: 1.5}
	m["Base"] = statistics.TriangularDistribution{Min: 0, Max: 4, MostLikely: 1.5}
	m["Pier"] = statistics.TriangularDistribution{Min: 6, Max: 12, MostLikely: 9}
	m["Pile"] = statistics.TriangularDistribution{Min: 6, Max: 12, MostLikely: 9}
	return m
}
