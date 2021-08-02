package compute

import "github.com/HydrologicEngineeringCenter/go-statistics/statistics"

func FoundationDistributionMap() map[string]statistics.ContinuousDistribution {
	m := make(map[string]statistics.ContinuousDistribution)
	m["AGR1"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["COM1"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["COM2"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["COM3"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["COM4"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["COM5"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["COM6"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["COM7"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["COM8"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["COM9"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["COM10"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["EDU1"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["EDU2"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["GOV1"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["GOV2"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["IND1"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["IND2"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["IND3"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["IND4"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["IND5"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["IND6"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["REL1"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES1-1SNB"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES1-1SNB_MEDWAVE"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES1-1SNB_HIGHWAVE"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES1-1SNB-PIER"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES1-1SNB-PIER_MEDWAVE"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES1-1SNB-PIER_HIGHWAVE"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES1-1SWB"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES1-1SWB_MEDWAVE"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES1-1SNB_HIGHWAVE"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES1-2SNB"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES1-2SNB_MEDWAVE"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES1-2SNB_HIGHWAVE"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES1-2SNB-PIER"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES1-2SNB-PIER_MEDWAVE"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES1-2SNB-PIER_HIGHWAVE"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES1-2SWB"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES1-2SWB_MEDWAVE"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES1-2SNB_HIGHWAVE"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES1-3SNB"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES1-3SWB"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES1-SLNB"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES1-SLWB"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES2"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES3A"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES3B"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES3C"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES3D"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES3E"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES3F"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES4"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES5"] = statistics.UniformDistribution{Min: 0, Max: 1}
	m["RES6"] = statistics.UniformDistribution{Min: 0, Max: 1}
	return m
}
