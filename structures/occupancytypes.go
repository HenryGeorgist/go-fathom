package structures

import (
	"github.com/HydrologicEngineeringCenter/go-statistics/statistics"
	"github.com/USACE/go-consequences/hazards"
	"github.com/USACE/go-consequences/paireddata"
	gstructs "github.com/USACE/go-consequences/structures"
)

//OccupancyTypeMap produces a map of all occupancy types as OccupancyTypeStochastic so they can be joined to the structure inventory to compute damage
func OccupancyTypeMap() map[string]gstructs.OccupancyTypeStochastic {
	m := make(map[string]gstructs.OccupancyTypeStochastic)
	m["RES1-1SNB"] = res11snb()
	m["RES1-1SWB"] = res11swb()
	m["RES1-2SNB"] = res12snb()
	m["RES1-2SWB"] = res12swb()
	m["RES1-3SNB"] = res13snb()
	m["RES1-3SWB"] = res13swb()
	m["RES1-SLNB"] = res1slnb()
	m["RES1-SLWB"] = res1slwb()
	m["RES3A"] = res3a()
	m["RES3B"] = res3b()
	m["RES3C"] = res3c()
	m["RES3D"] = res3d()
	m["RES3E"] = res3e()
	m["RES3F"] = res3f()

	return m
}
func res11snb() gstructs.OccupancyTypeStochastic {
	xvals := []float64{-8.0, -7.0, -6.0, -5.0, -4.0, -3.0, -2.0, -1.0, 0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0}
	structureydists := make([]statistics.ContinuousDistribution, 25)
	structureydists[0] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}                          //.28304
	structureydists[1] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}                          //1.683534
	structureydists[2] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}                          //3.084029
	structureydists[3] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}                          //4.484523103
	structureydists[4] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}                          //5.885018
	structureydists[5] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}                          //7.285512
	structureydists[6] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}                          //8.686007
	structureydists[7] = statistics.NormalDistribution{Mean: 3.52144, StandardDeviation: 0.30000001192092896}  //updated mean, not stdev, increased
	structureydists[8] = statistics.NormalDistribution{Mean: 9.020082, StandardDeviation: 1.2000000476837158}  //updated mean, not stdev, decreased
	structureydists[9] = statistics.NormalDistribution{Mean: 19.83513, StandardDeviation: 1.6000000238418579}  //updated mean, not stdev, decreased
	structureydists[10] = statistics.NormalDistribution{Mean: 30.97816, StandardDeviation: 1.6000000238418579} //updated mean, not stdev, decreased
	structureydists[11] = statistics.NormalDistribution{Mean: 38.14662, StandardDeviation: 1.7999999523162842} //updated mean, not stdev, decreased
	structureydists[12] = statistics.NormalDistribution{Mean: 44.10787, StandardDeviation: 1.8999999761581421} //updated mean, not stdev, decreased
	structureydists[13] = statistics.NormalDistribution{Mean: 49.34622, StandardDeviation: 2}                  //updated mean, not stdev, decreased
	structureydists[14] = statistics.NormalDistribution{Mean: 54.09353, StandardDeviation: 2.0999999046325684} //updated mean, not stdev, decreased
	structureydists[15] = statistics.NormalDistribution{Mean: 58.51472, StandardDeviation: 2.2000000476837158} //updated mean, not stdev, decreased
	structureydists[16] = statistics.NormalDistribution{Mean: 62.77186, StandardDeviation: 2.2999999523162842} //updated mean, not stdev, decreased
	structureydists[17] = statistics.NormalDistribution{Mean: 67.07379, StandardDeviation: 2.2999999523162842} //updated mean, not stdev, decreased
	structureydists[18] = statistics.NormalDistribution{Mean: 71.36232, StandardDeviation: 2.3499999046325684} //updated mean, not stdev, decreased
	structureydists[19] = statistics.NormalDistribution{Mean: 75.3777, StandardDeviation: 2.3900001049041748}  //updated mean, not stdev, decreased
	structureydists[20] = statistics.NormalDistribution{Mean: 78.39432, StandardDeviation: 2.4000000953674316} //updated mean, not stdev, increased
	structureydists[21] = statistics.NormalDistribution{Mean: 80.47807, StandardDeviation: 2.4100000858306885} //updated mean, not stdev, increased
	structureydists[22] = statistics.NormalDistribution{Mean: 82.56183, StandardDeviation: 2.4200000762939453} //updated mean, not stdev, increased
	structureydists[23] = statistics.NormalDistribution{Mean: 84.40445, StandardDeviation: 2.4300000667572021} //updated mean, not stdev, increased
	structureydists[24] = statistics.NormalDistribution{Mean: 85.46739, StandardDeviation: 2.4300000667572021} //updated mean, not stdev, increased
	contentydists := make([]statistics.ContinuousDistribution, 25)
	contentydists[0] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[1] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[2] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[3] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[4] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[5] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[6] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[7] = statistics.NormalDistribution{Mean: 5.77496, StandardDeviation: 0.20000000298023224}
	contentydists[8] = statistics.NormalDistribution{Mean: 14.79242, StandardDeviation: 0.800000011920929}
	contentydists[9] = statistics.NormalDistribution{Mean: 26.61195, StandardDeviation: 1.2999999523162842}
	contentydists[10] = statistics.NormalDistribution{Mean: 34.46171158, StandardDeviation: 1.7000000476837158}
	contentydists[11] = statistics.NormalDistribution{Mean: 40.37666457, StandardDeviation: 1.8999999761581421}
	contentydists[12] = statistics.NormalDistribution{Mean: 45.13015224, StandardDeviation: 2.1700000762939453}
	contentydists[13] = statistics.NormalDistribution{Mean: 49.07197792, StandardDeviation: 2.5}
	contentydists[14] = statistics.NormalDistribution{Mean: 52.40651485, StandardDeviation: 2.7999999523162842}
	contentydists[15] = statistics.NormalDistribution{Mean: 55.26772666, StandardDeviation: 2.9500000476837158}
	contentydists[16] = statistics.NormalDistribution{Mean: 57.75029694, StandardDeviation: 3.0999999046325684}
	contentydists[17] = statistics.NormalDistribution{Mean: 59.92361793, StandardDeviation: 3.2000000476837158}
	contentydists[18] = statistics.NormalDistribution{Mean: 61.79301203, StandardDeviation: 3.2999999523162842}
	contentydists[19] = statistics.NormalDistribution{Mean: 63.48169585, StandardDeviation: 3.4000000953674316}
	contentydists[20] = statistics.NormalDistribution{Mean: 64.9619038, StandardDeviation: 3.4000000953674316}
	contentydists[21] = statistics.NormalDistribution{Mean: 66.24739316, StandardDeviation: 3.4100000858306885}
	contentydists[22] = statistics.NormalDistribution{Mean: 67.53288252, StandardDeviation: 3.4100000858306885}
	contentydists[23] = statistics.NormalDistribution{Mean: 68.73556038, StandardDeviation: 3.4100000858306885}
	contentydists[24] = statistics.NormalDistribution{Mean: 69.6704811, StandardDeviation: 3.4100000858306885}
	var structuredamagefunctionStochastic = paireddata.UncertaintyPairedData{Xvals: xvals, Yvals: structureydists}
	var contentdamagefunctionStochastic = paireddata.UncertaintyPairedData{Xvals: xvals, Yvals: contentydists}

	sm := make(map[hazards.Parameter]interface{})
	var sdf = gstructs.DamageFunctionFamilyStochastic{DamageFunctions: sm}

	cm := make(map[hazards.Parameter]interface{})
	var cdf = gstructs.DamageFunctionFamilyStochastic{DamageFunctions: cm}
	//Default hazard.
	sdf.DamageFunctions[hazards.Default] = structuredamagefunctionStochastic
	cdf.DamageFunctions[hazards.Default] = contentdamagefunctionStochastic
	//Depth Hazard
	return gstructs.OccupancyTypeStochastic{Name: "RES1-1SNB", StructureDFF: sdf, ContentDFF: cdf}
}

func res11swb() gstructs.OccupancyTypeStochastic {
	xvals := []float64{-8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	structureydists := make([]statistics.ContinuousDistribution, 25)
	structureydists[0] = statistics.NormalDistribution{Mean: 0.143063909, StandardDeviation: 0}
	structureydists[1] = statistics.NormalDistribution{Mean: 0.85095156, StandardDeviation: 0.0099999997764825821}
	structureydists[2] = statistics.NormalDistribution{Mean: 1.558839211, StandardDeviation: 0.019999999552965164}
	structureydists[3] = statistics.NormalDistribution{Mean: 2.266726863, StandardDeviation: 0.10000000149011612}
	structureydists[4] = statistics.NormalDistribution{Mean: 2.974614514, StandardDeviation: 0.30000001192092896}
	structureydists[5] = statistics.NormalDistribution{Mean: 3.682502166, StandardDeviation: 0.699999988079071}
	structureydists[6] = statistics.NormalDistribution{Mean: 4.390389817, StandardDeviation: 0.85000002384185791}
	structureydists[7] = statistics.NormalDistribution{Mean: 5.098277468, StandardDeviation: 0.82999998331069946}
	structureydists[8] = statistics.NormalDistribution{Mean: 5.80616512, StandardDeviation: 0.85000002384185791}
	structureydists[9] = statistics.NormalDistribution{Mean: 12.57304737, StandardDeviation: 0.95999997854232788}
	structureydists[10] = statistics.NormalDistribution{Mean: 15.6052059, StandardDeviation: 1.1399999856948853}
	structureydists[11] = statistics.NormalDistribution{Mean: 16.7709469, StandardDeviation: 1.3700000047683716}
	structureydists[12] = statistics.NormalDistribution{Mean: 17.59148844, StandardDeviation: 1.6299999952316284}
	structureydists[13] = statistics.NormalDistribution{Mean: 18.3264792, StandardDeviation: 1.8899999856948853}
	structureydists[14] = statistics.NormalDistribution{Mean: 19.13620178, StandardDeviation: 1.8999999761581421}
	structureydists[15] = statistics.NormalDistribution{Mean: 20.23852106, StandardDeviation: 2.0199999809265137}
	structureydists[16] = statistics.NormalDistribution{Mean: 22.02277789, StandardDeviation: 2.0399999618530273}
	structureydists[17] = statistics.NormalDistribution{Mean: 25.19167289, StandardDeviation: 2.130000114440918}
	structureydists[18] = statistics.NormalDistribution{Mean: 32.00189443, StandardDeviation: 2.2000000476837158}
	structureydists[19] = statistics.NormalDistribution{Mean: 42.28177028, StandardDeviation: 2.2999999523162842}
	structureydists[20] = statistics.NormalDistribution{Mean: 53.0254879, StandardDeviation: 2.2999999523162842}
	structureydists[21] = statistics.NormalDistribution{Mean: 64.20243845, StandardDeviation: 2.2999999523162842}
	structureydists[22] = statistics.NormalDistribution{Mean: 75.37938899, StandardDeviation: 2.2999999523162842}
	structureydists[23] = statistics.NormalDistribution{Mean: 83.99558572, StandardDeviation: 2.2999999523162842}
	structureydists[24] = statistics.NormalDistribution{Mean: 84.33201177, StandardDeviation: 2.2999999523162842}
	contentydists := make([]statistics.ContinuousDistribution, 25)
	contentydists[0] = statistics.NormalDistribution{Mean: 0.40215542, StandardDeviation: 0}
	contentydists[1] = statistics.NormalDistribution{Mean: 2.392041331, StandardDeviation: 0.0099999997764825821}
	contentydists[2] = statistics.NormalDistribution{Mean: 4.381927241, StandardDeviation: 0.10000000149011612}
	contentydists[3] = statistics.NormalDistribution{Mean: 6.371813151, StandardDeviation: 0.30000001192092896}
	contentydists[4] = statistics.NormalDistribution{Mean: 8.361699062, StandardDeviation: 0.5}
	contentydists[5] = statistics.NormalDistribution{Mean: 10.35158497, StandardDeviation: 0.60000002384185791}
	contentydists[6] = statistics.NormalDistribution{Mean: 12.34147088, StandardDeviation: 0.74000000953674316}
	contentydists[7] = statistics.NormalDistribution{Mean: 14.33135679, StandardDeviation: 0.72000002861022949}
	contentydists[8] = statistics.NormalDistribution{Mean: 16.3212427, StandardDeviation: 0.74000000953674316}
	contentydists[9] = statistics.NormalDistribution{Mean: 22.1867589, StandardDeviation: 0.82999998331069946}
	contentydists[10] = statistics.NormalDistribution{Mean: 26.8400598, StandardDeviation: 0.98000001907348633}
	contentydists[11] = statistics.NormalDistribution{Mean: 31.77075781, StandardDeviation: 1.1699999570846558}
	contentydists[12] = statistics.NormalDistribution{Mean: 37.1854767, StandardDeviation: 1.3899999856948853}
	contentydists[13] = statistics.NormalDistribution{Mean: 42.83692138, StandardDeviation: 1.6000000238418579}
	contentydists[14] = statistics.NormalDistribution{Mean: 48.36619829, StandardDeviation: 1.8400000333786011}
	contentydists[15] = statistics.NormalDistribution{Mean: 53.47110577, StandardDeviation: 2}
	contentydists[16] = statistics.NormalDistribution{Mean: 57.97506284, StandardDeviation: 2.1600000858306885}
	contentydists[17] = statistics.NormalDistribution{Mean: 61.822085, StandardDeviation: 2.2999999523162842}
	contentydists[18] = statistics.NormalDistribution{Mean: 64.91563518, StandardDeviation: 2.4000000953674316}
	contentydists[19] = statistics.NormalDistribution{Mean: 67.54162193, StandardDeviation: 2.4500000476837158}
	contentydists[20] = statistics.NormalDistribution{Mean: 69.64996169, StandardDeviation: 2.4500000476837158}
	contentydists[21] = statistics.NormalDistribution{Mean: 71.27481391, StandardDeviation: 2.4500000476837158}
	contentydists[22] = statistics.NormalDistribution{Mean: 72.89966614, StandardDeviation: 2.4500000476837158}
	contentydists[23] = statistics.NormalDistribution{Mean: 74.35449039, StandardDeviation: 2.4500000476837158}
	contentydists[24] = statistics.NormalDistribution{Mean: 75.25955753, StandardDeviation: 2.4500000476837158}
	var structuredamagefunctionStochastic = paireddata.UncertaintyPairedData{Xvals: xvals, Yvals: structureydists}
	var contentdamagefunctionStochastic = paireddata.UncertaintyPairedData{Xvals: xvals, Yvals: contentydists}
	sm := make(map[hazards.Parameter]interface{})
	var sdf = gstructs.DamageFunctionFamilyStochastic{DamageFunctions: sm}

	cm := make(map[hazards.Parameter]interface{})
	var cdf = gstructs.DamageFunctionFamilyStochastic{DamageFunctions: cm}
	//Default hazard.
	sdf.DamageFunctions[hazards.Default] = structuredamagefunctionStochastic
	cdf.DamageFunctions[hazards.Default] = contentdamagefunctionStochastic

	return gstructs.OccupancyTypeStochastic{Name: "RES1-1SWB", StructureDFF: sdf, ContentDFF: cdf}
}
func res12snb() gstructs.OccupancyTypeStochastic {
	xvals := []float64{-8.0, -7.0, -6.0, -5.0, -4.0, -3.0, -2.0, -1.0, 0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0}
	structureydists := make([]statistics.ContinuousDistribution, 25)
	structureydists[0] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[1] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[2] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[3] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[4] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[5] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[6] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[7] = statistics.NormalDistribution{Mean: 2.939710557, StandardDeviation: 0.30000001192092896}
	structureydists[8] = statistics.NormalDistribution{Mean: 7.529996304, StandardDeviation: 1}
	structureydists[9] = statistics.NormalDistribution{Mean: 16.08748453, StandardDeviation: 1.5}
	structureydists[10] = statistics.NormalDistribution{Mean: 22.18031775, StandardDeviation: 2}
	structureydists[11] = statistics.NormalDistribution{Mean: 25.55483981, StandardDeviation: 2.4000000953674316}
	structureydists[12] = statistics.NormalDistribution{Mean: 28.63015146, StandardDeviation: 2.7000000476837158}
	structureydists[13] = statistics.NormalDistribution{Mean: 31.62127385, StandardDeviation: 3.0999999046325684}
	structureydists[14] = statistics.NormalDistribution{Mean: 34.58785417, StandardDeviation: 3.2999999523162842}
	structureydists[15] = statistics.NormalDistribution{Mean: 37.56694206, StandardDeviation: 3.4500000476837158}
	structureydists[16] = statistics.NormalDistribution{Mean: 40.65711043, StandardDeviation: 3.4900000095367432}
	structureydists[17] = statistics.NormalDistribution{Mean: 44.225069, StandardDeviation: 3.5099999904632568}
	structureydists[18] = statistics.NormalDistribution{Mean: 50.31651122, StandardDeviation: 3.5499999523162842}
	structureydists[19] = statistics.NormalDistribution{Mean: 59.04275233, StandardDeviation: 3.5999999046325684}
	structureydists[20] = statistics.NormalDistribution{Mean: 65.61021228, StandardDeviation: 3.6500000953674316}
	structureydists[21] = statistics.NormalDistribution{Mean: 70.16134872, StandardDeviation: 3.7000000476837158}
	structureydists[22] = statistics.NormalDistribution{Mean: 74.71248516, StandardDeviation: 3.7200000286102295}
	structureydists[23] = statistics.NormalDistribution{Mean: 78.33639264, StandardDeviation: 3.75}
	structureydists[24] = statistics.NormalDistribution{Mean: 78.96225982, StandardDeviation: 3.7999999523162842}
	contentydists := make([]statistics.ContinuousDistribution, 25)
	contentydists[0] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[1] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[2] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[3] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[4] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[5] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[6] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[7] = statistics.NormalDistribution{Mean: 4.769714108, StandardDeviation: 0.10000000149011612}
	contentydists[8] = statistics.NormalDistribution{Mean: 12.2175054, StandardDeviation: 0.550000011920929}
	contentydists[9] = statistics.NormalDistribution{Mean: 21.05706565, StandardDeviation: 1}
	contentydists[10] = statistics.NormalDistribution{Mean: 25.93632764, StandardDeviation: 1.3999999761581421}
	contentydists[11] = statistics.NormalDistribution{Mean: 29.52336958, StandardDeviation: 1.7999999523162842}
	contentydists[12] = statistics.NormalDistribution{Mean: 32.48415064, StandardDeviation: 2.0999999046325684}
	contentydists[13] = statistics.NormalDistribution{Mean: 35.0628348, StandardDeviation: 2.4000000953674316}
	contentydists[14] = statistics.NormalDistribution{Mean: 37.38839967, StandardDeviation: 2.7000000476837158}
	contentydists[15] = statistics.NormalDistribution{Mean: 39.5689533, StandardDeviation: 3}
	contentydists[16] = statistics.NormalDistribution{Mean: 41.77044572, StandardDeviation: 3.0999999046325684}
	contentydists[17] = statistics.NormalDistribution{Mean: 44.33330409, StandardDeviation: 3.2999999523162842}
	contentydists[18] = statistics.NormalDistribution{Mean: 48.29864311, StandardDeviation: 3.5}
	contentydists[19] = statistics.NormalDistribution{Mean: 53.62893042, StandardDeviation: 3.5}
	contentydists[20] = statistics.NormalDistribution{Mean: 58.38647778, StandardDeviation: 3.5}
	contentydists[21] = statistics.NormalDistribution{Mean: 62.60908024, StandardDeviation: 3.5}
	contentydists[22] = statistics.NormalDistribution{Mean: 66.83168269, StandardDeviation: 3.5999999046325684}
	contentydists[23] = statistics.NormalDistribution{Mean: 70.08486256, StandardDeviation: 3.5999999046325684}
	contentydists[24] = statistics.NormalDistribution{Mean: 70.20357606, StandardDeviation: 3.5999999046325684}
	var structuredamagefunctionStochastic = paireddata.UncertaintyPairedData{Xvals: xvals, Yvals: structureydists}
	var contentdamagefunctionStochastic = paireddata.UncertaintyPairedData{Xvals: xvals, Yvals: contentydists}
	sm := make(map[hazards.Parameter]interface{})
	var sdf = gstructs.DamageFunctionFamilyStochastic{DamageFunctions: sm}

	cm := make(map[hazards.Parameter]interface{})
	var cdf = gstructs.DamageFunctionFamilyStochastic{DamageFunctions: cm}
	//Default hazard.
	sdf.DamageFunctions[hazards.Default] = structuredamagefunctionStochastic
	cdf.DamageFunctions[hazards.Default] = contentdamagefunctionStochastic

	return gstructs.OccupancyTypeStochastic{Name: "RES1-2SNB", StructureDFF: sdf, ContentDFF: cdf}
}

func res12swb() gstructs.OccupancyTypeStochastic {
	xvals := []float64{-8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	structureydists := make([]statistics.ContinuousDistribution, 25)
	structureydists[0] = statistics.NormalDistribution{Mean: 0.09170235, StandardDeviation: 0}
	structureydists[1] = statistics.NormalDistribution{Mean: 0.54545034, StandardDeviation: 0}
	structureydists[2] = statistics.NormalDistribution{Mean: 0.999198331, StandardDeviation: 0.0099999997764825821}
	structureydists[3] = statistics.NormalDistribution{Mean: 1.452946322, StandardDeviation: 0.10000000149011612}
	structureydists[4] = statistics.NormalDistribution{Mean: 1.906694312, StandardDeviation: 0.30000001192092896}
	structureydists[5] = statistics.NormalDistribution{Mean: 2.360442303, StandardDeviation: 0.60000002384185791}
	structureydists[6] = statistics.NormalDistribution{Mean: 2.814190293, StandardDeviation: 0.89999997615814209}
	structureydists[7] = statistics.NormalDistribution{Mean: 3.267938284, StandardDeviation: 1.1000000238418579}
	structureydists[8] = statistics.NormalDistribution{Mean: 3.721686275, StandardDeviation: 1.3200000524520874}
	structureydists[9] = statistics.NormalDistribution{Mean: 8.980655072, StandardDeviation: 1.3500000238418579}
	structureydists[10] = statistics.NormalDistribution{Mean: 12.2919343, StandardDeviation: 1.5}
	structureydists[11] = statistics.NormalDistribution{Mean: 13.28200361, StandardDeviation: 1.75}
	structureydists[12] = statistics.NormalDistribution{Mean: 13.98298013, StandardDeviation: 2.0399999618530273}
	structureydists[13] = statistics.NormalDistribution{Mean: 14.60177938, StandardDeviation: 2.3399999141693115}
	structureydists[14] = statistics.NormalDistribution{Mean: 15.19125161, StandardDeviation: 2.5999999046325684}
	structureydists[15] = statistics.NormalDistribution{Mean: 15.77461125, StandardDeviation: 2.7000000476837158}
	structureydists[16] = statistics.NormalDistribution{Mean: 16.38572858, StandardDeviation: 2.75}
	structureydists[17] = statistics.NormalDistribution{Mean: 17.15123848, StandardDeviation: 2.7599999904632568}
	structureydists[18] = statistics.NormalDistribution{Mean: 19.08670375, StandardDeviation: 2.7699999809265137}
	structureydists[19] = statistics.NormalDistribution{Mean: 22.52422032, StandardDeviation: 2.7799999713897705}
	structureydists[20] = statistics.NormalDistribution{Mean: 29.51224106, StandardDeviation: 2.7899999618530273}
	structureydists[21] = statistics.NormalDistribution{Mean: 39.81646872, StandardDeviation: 2.7999999523162842}
	structureydists[22] = statistics.NormalDistribution{Mean: 50.12069638, StandardDeviation: 2.8299999237060547}
	structureydists[23] = statistics.NormalDistribution{Mean: 58.69396948, StandardDeviation: 2.8599998950958252}
	structureydists[24] = statistics.NormalDistribution{Mean: 61.67048954, StandardDeviation: 2.8599998950958252}
	contentydists := make([]statistics.ContinuousDistribution, 25)
	contentydists[0] = statistics.NormalDistribution{Mean: 0.308201347, StandardDeviation: 0}
	contentydists[1] = statistics.NormalDistribution{Mean: 1.833197625, StandardDeviation: 0.0099999997764825821}
	contentydists[2] = statistics.NormalDistribution{Mean: 3.358193903, StandardDeviation: 0.10000000149011612}
	contentydists[3] = statistics.NormalDistribution{Mean: 4.88319018, StandardDeviation: 0.20000000298023224}
	contentydists[4] = statistics.NormalDistribution{Mean: 6.408186458, StandardDeviation: 0.34999999403953552}
	contentydists[5] = statistics.NormalDistribution{Mean: 7.933182736, StandardDeviation: 0.5}
	contentydists[6] = statistics.NormalDistribution{Mean: 9.458179013, StandardDeviation: 0.699999988079071}
	contentydists[7] = statistics.NormalDistribution{Mean: 10.98317529, StandardDeviation: 0.86000001430511475}
	contentydists[8] = statistics.NormalDistribution{Mean: 12.50817157, StandardDeviation: 1}
	contentydists[9] = statistics.NormalDistribution{Mean: 17.44715762, StandardDeviation: 1.1100000143051148}
	contentydists[10] = statistics.NormalDistribution{Mean: 21.10806642, StandardDeviation: 1.2300000190734863}
	contentydists[11] = statistics.NormalDistribution{Mean: 24.19765392, StandardDeviation: 1.4299999475479126}
	contentydists[12] = statistics.NormalDistribution{Mean: 26.96661489, StandardDeviation: 1.6699999570846558}
	contentydists[13] = statistics.NormalDistribution{Mean: 29.49530306, StandardDeviation: 1.9199999570846558}
	contentydists[14] = statistics.NormalDistribution{Mean: 31.82217694, StandardDeviation: 2.1500000953674316}
	contentydists[15] = statistics.NormalDistribution{Mean: 33.97046798, StandardDeviation: 2.3599998950958252}
	contentydists[16] = statistics.NormalDistribution{Mean: 35.95676362, StandardDeviation: 2.559999942779541}
	contentydists[17] = statistics.NormalDistribution{Mean: 37.79461003, StandardDeviation: 2.7599999904632568}
	contentydists[18] = statistics.NormalDistribution{Mean: 39.97675387, StandardDeviation: 3.0399999618530273}
	contentydists[19] = statistics.NormalDistribution{Mean: 43.12683181, StandardDeviation: 3.2999999523162842}
	contentydists[20] = statistics.NormalDistribution{Mean: 48.61795127, StandardDeviation: 3.5999999046325684}
	contentydists[21] = statistics.NormalDistribution{Mean: 56.29562726, StandardDeviation: 3.9000000953674316}
	contentydists[22] = statistics.NormalDistribution{Mean: 63.97330326, StandardDeviation: 4.2600002288818359}
	contentydists[23] = statistics.NormalDistribution{Mean: 69.85646812, StandardDeviation: 4.5999999046325684}
	contentydists[24] = statistics.NormalDistribution{Mean: 69.93738033, StandardDeviation: 5}
	var structuredamagefunctionStochastic = paireddata.UncertaintyPairedData{Xvals: xvals, Yvals: structureydists}
	var contentdamagefunctionStochastic = paireddata.UncertaintyPairedData{Xvals: xvals, Yvals: contentydists}
	sm := make(map[hazards.Parameter]interface{})
	var sdf = gstructs.DamageFunctionFamilyStochastic{DamageFunctions: sm}

	cm := make(map[hazards.Parameter]interface{})
	var cdf = gstructs.DamageFunctionFamilyStochastic{DamageFunctions: cm}
	//Default hazard.
	sdf.DamageFunctions[hazards.Default] = structuredamagefunctionStochastic
	cdf.DamageFunctions[hazards.Default] = contentdamagefunctionStochastic
	return gstructs.OccupancyTypeStochastic{Name: "RES1-2SWB", StructureDFF: sdf, ContentDFF: cdf}
}

func res13snb() gstructs.OccupancyTypeStochastic {
	xvals := []float64{-8.0, -7.0, -6.0, -5.0, -4.0, -3.0, -2.0, -1.0, 0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0}
	structureydists := make([]statistics.ContinuousDistribution, 25)
	structureydists[0] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[1] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[2] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[3] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[4] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[5] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[6] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[7] = statistics.NormalDistribution{Mean: 1.514906645, StandardDeviation: 0.30000001192092896}
	structureydists[8] = statistics.NormalDistribution{Mean: 3.88039612, StandardDeviation: 1}
	structureydists[9] = statistics.NormalDistribution{Mean: 11.05754009, StandardDeviation: 1.5}
	structureydists[10] = statistics.NormalDistribution{Mean: 15.40045926, StandardDeviation: 2.0999999046325684}
	structureydists[11] = statistics.NormalDistribution{Mean: 17.63659009, StandardDeviation: 2.5999999046325684}
	structureydists[12] = statistics.NormalDistribution{Mean: 19.35562508, StandardDeviation: 3}
	structureydists[13] = statistics.NormalDistribution{Mean: 20.88880829, StandardDeviation: 3.2000000476837158}
	structureydists[14] = statistics.NormalDistribution{Mean: 22.36864664, StandardDeviation: 3.5}
	structureydists[15] = statistics.NormalDistribution{Mean: 23.87466606, StandardDeviation: 3.5499999523162842}
	structureydists[16] = statistics.NormalDistribution{Mean: 25.4772364, StandardDeviation: 3.5999999046325684}
	structureydists[17] = statistics.NormalDistribution{Mean: 27.25861467, StandardDeviation: 3.6500000953674316}
	structureydists[18] = statistics.NormalDistribution{Mean: 29.4712456, StandardDeviation: 3.7000000476837158}
	structureydists[19] = statistics.NormalDistribution{Mean: 32.10213449, StandardDeviation: 3.7300000190734863}
	structureydists[20] = statistics.NormalDistribution{Mean: 34.84277995, StandardDeviation: 3.7699999809265137}
	structureydists[21] = statistics.NormalDistribution{Mean: 37.68593917, StandardDeviation: 3.7799999713897705}
	structureydists[22] = statistics.NormalDistribution{Mean: 40.52909839, StandardDeviation: 3.7899999618530273}
	structureydists[23] = statistics.NormalDistribution{Mean: 43.07239525, StandardDeviation: 3.7999999523162842}
	structureydists[24] = statistics.NormalDistribution{Mean: 44.64613714, StandardDeviation: 3.7999999523162842}
	contentydists := make([]statistics.ContinuousDistribution, 25)
	contentydists[0] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[1] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[2] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[3] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[4] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[5] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[6] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[7] = statistics.NormalDistribution{Mean: 4.87139255, StandardDeviation: 0.05000000074505806}
	contentydists[8] = statistics.NormalDistribution{Mean: 12.47795223, StandardDeviation: 0.5}
	contentydists[9] = statistics.NormalDistribution{Mean: 21.50579094, StandardDeviation: 0.89999997615814209}
	contentydists[10] = statistics.NormalDistribution{Mean: 26.10964048, StandardDeviation: 1.2999999523162842}
	contentydists[11] = statistics.NormalDistribution{Mean: 29.13100792, StandardDeviation: 1.7000000476837158}
	contentydists[12] = statistics.NormalDistribution{Mean: 31.36660701, StandardDeviation: 2}
	contentydists[13] = statistics.NormalDistribution{Mean: 33.12125272, StandardDeviation: 2.2999999523162842}
	contentydists[14] = statistics.NormalDistribution{Mean: 34.55091112, StandardDeviation: 2.5999999046325684}
	contentydists[15] = statistics.NormalDistribution{Mean: 35.75582101, StandardDeviation: 2.9000000953674316}
	contentydists[16] = statistics.NormalDistribution{Mean: 36.83385441, StandardDeviation: 3.0999999046325684}
	contentydists[17] = statistics.NormalDistribution{Mean: 37.95169193, StandardDeviation: 3.2999999523162842}
	contentydists[18] = statistics.NormalDistribution{Mean: 39.58341398, StandardDeviation: 3.5}
	contentydists[19] = statistics.NormalDistribution{Mean: 41.74808174, StandardDeviation: 3.5999999046325684}
	contentydists[20] = statistics.NormalDistribution{Mean: 43.07696659, StandardDeviation: 3.7000000476837158}
	contentydists[21] = statistics.NormalDistribution{Mean: 43.62522172, StandardDeviation: 3.7999999523162842}
	contentydists[22] = statistics.NormalDistribution{Mean: 44.17347686, StandardDeviation: 3.9000000953674316}
	contentydists[23] = statistics.NormalDistribution{Mean: 44.62749063, StandardDeviation: 3.9000000953674316}
	contentydists[24] = statistics.NormalDistribution{Mean: 44.77679067, StandardDeviation: 3.9000000953674316}
	var structuredamagefunctionStochastic = paireddata.UncertaintyPairedData{Xvals: xvals, Yvals: structureydists}
	var contentdamagefunctionStochastic = paireddata.UncertaintyPairedData{Xvals: xvals, Yvals: contentydists}
	sm := make(map[hazards.Parameter]interface{})
	var sdf = gstructs.DamageFunctionFamilyStochastic{DamageFunctions: sm}

	cm := make(map[hazards.Parameter]interface{})
	var cdf = gstructs.DamageFunctionFamilyStochastic{DamageFunctions: cm}
	//Default hazard.
	sdf.DamageFunctions[hazards.Default] = structuredamagefunctionStochastic
	cdf.DamageFunctions[hazards.Default] = contentdamagefunctionStochastic

	return gstructs.OccupancyTypeStochastic{Name: "RES1-3SNB", StructureDFF: sdf, ContentDFF: cdf}
}
func res13swb() gstructs.OccupancyTypeStochastic {
	xvals := []float64{-8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	structureydists := make([]statistics.ContinuousDistribution, 25)
	structureydists[0] = statistics.NormalDistribution{Mean: 0.09170235, StandardDeviation: 0}
	structureydists[1] = statistics.NormalDistribution{Mean: 0.54545034, StandardDeviation: 0}
	structureydists[2] = statistics.NormalDistribution{Mean: 0.999198331, StandardDeviation: 0.0099999997764825821}
	structureydists[3] = statistics.NormalDistribution{Mean: 1.452946322, StandardDeviation: 0.10000000149011612}
	structureydists[4] = statistics.NormalDistribution{Mean: 1.906694312, StandardDeviation: 0.30000001192092896}
	structureydists[5] = statistics.NormalDistribution{Mean: 2.360442303, StandardDeviation: 0.60000002384185791}
	structureydists[6] = statistics.NormalDistribution{Mean: 2.814190293, StandardDeviation: 0.89999997615814209}
	structureydists[7] = statistics.NormalDistribution{Mean: 3.267938284, StandardDeviation: 1.1000000238418579}
	structureydists[8] = statistics.NormalDistribution{Mean: 3.721686275, StandardDeviation: 1.3200000524520874}
	structureydists[9] = statistics.NormalDistribution{Mean: 8.980655072, StandardDeviation: 1.3500000238418579}
	structureydists[10] = statistics.NormalDistribution{Mean: 12.2919343, StandardDeviation: 1.5}
	structureydists[11] = statistics.NormalDistribution{Mean: 13.28200361, StandardDeviation: 1.75}
	structureydists[12] = statistics.NormalDistribution{Mean: 13.98298013, StandardDeviation: 2.0399999618530273}
	structureydists[13] = statistics.NormalDistribution{Mean: 14.60177938, StandardDeviation: 2.3399999141693115}
	structureydists[14] = statistics.NormalDistribution{Mean: 15.19125161, StandardDeviation: 2.5999999046325684}
	structureydists[15] = statistics.NormalDistribution{Mean: 15.77461125, StandardDeviation: 2.7000000476837158}
	structureydists[16] = statistics.NormalDistribution{Mean: 16.38572858, StandardDeviation: 2.75}
	structureydists[17] = statistics.NormalDistribution{Mean: 17.15123848, StandardDeviation: 2.7599999904632568}
	structureydists[18] = statistics.NormalDistribution{Mean: 19.08670375, StandardDeviation: 2.7699999809265137}
	structureydists[19] = statistics.NormalDistribution{Mean: 22.52422032, StandardDeviation: 2.7799999713897705}
	structureydists[20] = statistics.NormalDistribution{Mean: 29.51224106, StandardDeviation: 2.7899999618530273}
	structureydists[21] = statistics.NormalDistribution{Mean: 39.81646872, StandardDeviation: 2.7999999523162842}
	structureydists[22] = statistics.NormalDistribution{Mean: 50.12069638, StandardDeviation: 2.8299999237060547}
	structureydists[23] = statistics.NormalDistribution{Mean: 58.69396948, StandardDeviation: 2.8599998950958252}
	structureydists[24] = statistics.NormalDistribution{Mean: 61.67048954, StandardDeviation: 2.8599998950958252}
	contentydists := make([]statistics.ContinuousDistribution, 25)
	contentydists[0] = statistics.NormalDistribution{Mean: 0.33712374, StandardDeviation: 0}
	contentydists[1] = statistics.NormalDistribution{Mean: 2.005229518, StandardDeviation: 0.0099999997764825821}
	contentydists[2] = statistics.NormalDistribution{Mean: 3.673335297, StandardDeviation: 0.10000000149011612}
	contentydists[3] = statistics.NormalDistribution{Mean: 5.341441075, StandardDeviation: 0.20000000298023224}
	contentydists[4] = statistics.NormalDistribution{Mean: 7.009546854, StandardDeviation: 0.34999999403953552}
	contentydists[5] = statistics.NormalDistribution{Mean: 8.677652632, StandardDeviation: 0.5}
	contentydists[6] = statistics.NormalDistribution{Mean: 10.34575841, StandardDeviation: 0.699999988079071}
	contentydists[7] = statistics.NormalDistribution{Mean: 12.01386419, StandardDeviation: 0.86000001430511475}
	contentydists[8] = statistics.NormalDistribution{Mean: 13.68196997, StandardDeviation: 1}
	contentydists[9] = statistics.NormalDistribution{Mean: 19.54269382, StandardDeviation: 1.1100000143051148}
	contentydists[10] = statistics.NormalDistribution{Mean: 23.14237988, StandardDeviation: 1.2300000190734863}
	contentydists[11] = statistics.NormalDistribution{Mean: 25.81469644, StandardDeviation: 1.4299999475479126}
	contentydists[12] = statistics.NormalDistribution{Mean: 28.01627865, StandardDeviation: 1.6699999570846558}
	contentydists[13] = statistics.NormalDistribution{Mean: 29.90890272, StandardDeviation: 1.9199999570846558}
	contentydists[14] = statistics.NormalDistribution{Mean: 31.57217278, StandardDeviation: 2.1500000953674316}
	contentydists[15] = statistics.NormalDistribution{Mean: 33.05261981, StandardDeviation: 2.3599998950958252}
	contentydists[16] = statistics.NormalDistribution{Mean: 34.38094526, StandardDeviation: 2.559999942779541}
	contentydists[17] = statistics.NormalDistribution{Mean: 35.57914047, StandardDeviation: 2.7599999904632568}
	contentydists[18] = statistics.NormalDistribution{Mean: 36.74044882, StandardDeviation: 3.0399999618530273}
	contentydists[19] = statistics.NormalDistribution{Mean: 38.02704693, StandardDeviation: 3.2999999523162842}
	contentydists[20] = statistics.NormalDistribution{Mean: 39.90541042, StandardDeviation: 3.5999999046325684}
	contentydists[21] = statistics.NormalDistribution{Mean: 42.33648877, StandardDeviation: 3.9000000953674316}
	contentydists[22] = statistics.NormalDistribution{Mean: 44.76756712, StandardDeviation: 4.2600002288818359}
	contentydists[23] = statistics.NormalDistribution{Mean: 46.67445569, StandardDeviation: 4.5999999046325684}
	contentydists[24] = statistics.NormalDistribution{Mean: 46.88646396, StandardDeviation: 5}
	var structuredamagefunctionStochastic = paireddata.UncertaintyPairedData{Xvals: xvals, Yvals: structureydists}
	var contentdamagefunctionStochastic = paireddata.UncertaintyPairedData{Xvals: xvals, Yvals: contentydists}
	sm := make(map[hazards.Parameter]interface{})
	var sdf = gstructs.DamageFunctionFamilyStochastic{DamageFunctions: sm}

	cm := make(map[hazards.Parameter]interface{})
	var cdf = gstructs.DamageFunctionFamilyStochastic{DamageFunctions: cm}
	//Default hazard.
	sdf.DamageFunctions[hazards.Default] = structuredamagefunctionStochastic
	cdf.DamageFunctions[hazards.Default] = contentdamagefunctionStochastic

	return gstructs.OccupancyTypeStochastic{Name: "RES1-3SWB", StructureDFF: sdf, ContentDFF: cdf}
}
func res1slnb() gstructs.OccupancyTypeStochastic {
	xvals := []float64{-8.0, -7.0, -6.0, -5.0, -4.0, -3.0, -2.0, -1.0, 0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0}
	structureydists := make([]statistics.ContinuousDistribution, 25)
	structureydists[0] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[1] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[2] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[3] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[4] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[5] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[6] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[7] = statistics.NormalDistribution{Mean: 1.982910128, StandardDeviation: 0}
	structureydists[8] = statistics.NormalDistribution{Mean: 5.079175534, StandardDeviation: 0.039999999105930328}
	structureydists[9] = statistics.NormalDistribution{Mean: 11.56330818, StandardDeviation: 0.20000000298023224}
	structureydists[10] = statistics.NormalDistribution{Mean: 17.29454137, StandardDeviation: 0.5}
	structureydists[11] = statistics.NormalDistribution{Mean: 20.86810816, StandardDeviation: 1}
	structureydists[12] = statistics.NormalDistribution{Mean: 23.95834494, StandardDeviation: 1.5}
	structureydists[13] = statistics.NormalDistribution{Mean: 26.86835757, StandardDeviation: 2}
	structureydists[14] = statistics.NormalDistribution{Mean: 29.6950892, StandardDeviation: 2.7000000476837158}
	structureydists[15] = statistics.NormalDistribution{Mean: 32.48186926, StandardDeviation: 3.2000000476837158}
	structureydists[16] = statistics.NormalDistribution{Mean: 35.28250612, StandardDeviation: 3.5}
	structureydists[17] = statistics.NormalDistribution{Mean: 38.23139718, StandardDeviation: 3.7999999523162842}
	structureydists[18] = statistics.NormalDistribution{Mean: 41.70627823, StandardDeviation: 4}
	structureydists[19] = statistics.NormalDistribution{Mean: 45.51872052, StandardDeviation: 3.5}
	structureydists[20] = statistics.NormalDistribution{Mean: 48.60909276, StandardDeviation: 3}
	structureydists[21] = statistics.NormalDistribution{Mean: 51.02504424, StandardDeviation: 2.5}
	structureydists[22] = statistics.NormalDistribution{Mean: 53.44099572, StandardDeviation: 2.0999999046325684}
	structureydists[23] = statistics.NormalDistribution{Mean: 55.76183376, StandardDeviation: 1.8999999761581421}
	structureydists[24] = statistics.NormalDistribution{Mean: 57.77513834, StandardDeviation: 1.8999999761581421}

	contentydists := make([]statistics.ContinuousDistribution, 25)
	contentydists[0] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[1] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[2] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[3] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[4] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[5] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[6] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[7] = statistics.NormalDistribution{Mean: 4.505130891, StandardDeviation: 0}
	contentydists[8] = statistics.NormalDistribution{Mean: 11.53978199, StandardDeviation: 0.059999998658895493}
	contentydists[9] = statistics.NormalDistribution{Mean: 18.66884561, StandardDeviation: 0.25}
	contentydists[10] = statistics.NormalDistribution{Mean: 22.8379740, StandardDeviation: 0.60000002384185791}
	contentydists[11] = statistics.NormalDistribution{Mean: 26.06119159, StandardDeviation: 1}
	contentydists[12] = statistics.NormalDistribution{Mean: 28.83287613, StandardDeviation: 1.5}
	contentydists[13] = statistics.NormalDistribution{Mean: 31.38412047, StandardDeviation: 1.6000000238418579}
	contentydists[14] = statistics.NormalDistribution{Mean: 33.93813824, StandardDeviation: 1.7999999523162842}
	contentydists[15] = statistics.NormalDistribution{Mean: 36.78403549, StandardDeviation: 2.0999999046325684}
	contentydists[16] = statistics.NormalDistribution{Mean: 40.27191538, StandardDeviation: 2.5}
	contentydists[17] = statistics.NormalDistribution{Mean: 44.70602747, StandardDeviation: 3}
	contentydists[18] = statistics.NormalDistribution{Mean: 50.21277879, StandardDeviation: 3.5}
	contentydists[19] = statistics.NormalDistribution{Mean: 56.18787891, StandardDeviation: 4}
	contentydists[20] = statistics.NormalDistribution{Mean: 61.25223889, StandardDeviation: 4.4000000953674316}
	contentydists[21] = statistics.NormalDistribution{Mean: 65.46595833, StandardDeviation: 4.75}
	contentydists[22] = statistics.NormalDistribution{Mean: 69.67967776, StandardDeviation: 4.8000001907348633}
	contentydists[23] = statistics.NormalDistribution{Mean: 72.91197317, StandardDeviation: 4.8000001907348633}
	contentydists[24] = statistics.NormalDistribution{Mean: 72.97099757, StandardDeviation: 4.8000001907348633}
	var structuredamagefunctionStochastic = paireddata.UncertaintyPairedData{Xvals: xvals, Yvals: structureydists}
	var contentdamagefunctionStochastic = paireddata.UncertaintyPairedData{Xvals: xvals, Yvals: contentydists}
	sm := make(map[hazards.Parameter]interface{})
	var sdf = gstructs.DamageFunctionFamilyStochastic{DamageFunctions: sm}

	cm := make(map[hazards.Parameter]interface{})
	var cdf = gstructs.DamageFunctionFamilyStochastic{DamageFunctions: cm}
	//Default hazard.
	sdf.DamageFunctions[hazards.Default] = structuredamagefunctionStochastic
	cdf.DamageFunctions[hazards.Default] = contentdamagefunctionStochastic

	return gstructs.OccupancyTypeStochastic{Name: "RES1-SLNB", StructureDFF: sdf, ContentDFF: cdf}
}
func res1slwb() gstructs.OccupancyTypeStochastic {
	xvals := []float64{-8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	structureydists := make([]statistics.ContinuousDistribution, 25)
	structureydists[0] = statistics.NormalDistribution{Mean: 0.187552916, StandardDeviation: 0}
	structureydists[1] = statistics.NormalDistribution{Mean: 1.115574488, StandardDeviation: 0}
	structureydists[2] = statistics.NormalDistribution{Mean: 2.04359606, StandardDeviation: 0.30000001192092896}
	structureydists[3] = statistics.NormalDistribution{Mean: 2.971617631, StandardDeviation: 0.30000001192092896}
	structureydists[4] = statistics.NormalDistribution{Mean: 3.899639203, StandardDeviation: 0.5}
	structureydists[5] = statistics.NormalDistribution{Mean: 4.827660775, StandardDeviation: 0.699999988079071}
	structureydists[6] = statistics.NormalDistribution{Mean: 5.755682347, StandardDeviation: 1}
	structureydists[7] = statistics.NormalDistribution{Mean: 6.683703918, StandardDeviation: 1.2000000476837158}
	structureydists[8] = statistics.NormalDistribution{Mean: 7.61172549, StandardDeviation: 1.6000000238418579}
	structureydists[9] = statistics.NormalDistribution{Mean: 11.2104594, StandardDeviation: 1.7000000476837158}
	structureydists[10] = statistics.NormalDistribution{Mean: 13.23905094, StandardDeviation: 1.8999999761581421}
	structureydists[11] = statistics.NormalDistribution{Mean: 14.85944755, StandardDeviation: 2.0999999046325684}
	structureydists[12] = statistics.NormalDistribution{Mean: 16.27590526, StandardDeviation: 2.4000000953674316}
	structureydists[13] = statistics.NormalDistribution{Mean: 17.57425364, StandardDeviation: 2.5999999046325684}
	structureydists[14] = statistics.NormalDistribution{Mean: 18.79539355, StandardDeviation: 2.9000000953674316}
	structureydists[15] = statistics.NormalDistribution{Mean: 19.99670165, StandardDeviation: 3.2000000476837158}
	structureydists[16] = statistics.NormalDistribution{Mean: 21.3506072, StandardDeviation: 3.2999999523162842}
	structureydists[17] = statistics.NormalDistribution{Mean: 23.46652043, StandardDeviation: 3.4000000953674316}
	structureydists[18] = statistics.NormalDistribution{Mean: 31.2370149, StandardDeviation: 3.4500000476837158}
	structureydists[19] = statistics.NormalDistribution{Mean: 42.74599823, StandardDeviation: 3.5}
	structureydists[20] = statistics.NormalDistribution{Mean: 49.74738846, StandardDeviation: 3.5699999332427979}
	structureydists[21] = statistics.NormalDistribution{Mean: 52.53864097, StandardDeviation: 3.619999885559082}
	structureydists[22] = statistics.NormalDistribution{Mean: 55.32989348, StandardDeviation: 3.619999885559082}
	structureydists[23] = statistics.NormalDistribution{Mean: 57.46423328, StandardDeviation: 3.619999885559082}
	structureydists[24] = statistics.NormalDistribution{Mean: 57.47455533, StandardDeviation: 3.619999885559082}
	contentydists := make([]statistics.ContinuousDistribution, 25)
	contentydists[0] = statistics.NormalDistribution{Mean: 0.410480696, StandardDeviation: 0}
	contentydists[1] = statistics.NormalDistribution{Mean: 2.441560501, StandardDeviation: 0}
	contentydists[2] = statistics.NormalDistribution{Mean: 4.472640307, StandardDeviation: 0.05000000074505806}
	contentydists[3] = statistics.NormalDistribution{Mean: 6.503720112, StandardDeviation: 0.15000000596046448}
	contentydists[4] = statistics.NormalDistribution{Mean: 8.53479991, StandardDeviation: 0.30000001192092896}
	contentydists[5] = statistics.NormalDistribution{Mean: 10.56587972, StandardDeviation: 0.5}
	contentydists[6] = statistics.NormalDistribution{Mean: 12.59695953, StandardDeviation: 0.699999988079071}
	contentydists[7] = statistics.NormalDistribution{Mean: 14.62803933, StandardDeviation: 0.89999997615814209}
	contentydists[8] = statistics.NormalDistribution{Mean: 16.65911914, StandardDeviation: 1.059999942779541}
	contentydists[9] = statistics.NormalDistribution{Mean: 20.80415477, StandardDeviation: 1.2000000476837158}
	contentydists[10] = statistics.NormalDistribution{Mean: 23.02387255, StandardDeviation: 1.3999999761581421}
	contentydists[11] = statistics.NormalDistribution{Mean: 24.65332945, StandardDeviation: 1.6000000238418579}
	contentydists[12] = statistics.NormalDistribution{Mean: 26.15944434, StandardDeviation: 1.7999999523162842}
	contentydists[13] = statistics.NormalDistribution{Mean: 27.83655466, StandardDeviation: 2}
	contentydists[14] = statistics.NormalDistribution{Mean: 29.94829832, StandardDeviation: 2.1800000667572021}
	contentydists[15] = statistics.NormalDistribution{Mean: 32.70971169, StandardDeviation: 2.2999999523162842}
	contentydists[16] = statistics.NormalDistribution{Mean: 36.21101358, StandardDeviation: 2.4000000953674316}
	contentydists[17] = statistics.NormalDistribution{Mean: 40.34738664, StandardDeviation: 2.440000057220459}
	contentydists[18] = statistics.NormalDistribution{Mean: 44.91796959, StandardDeviation: 2.440000057220459}
	contentydists[19] = statistics.NormalDistribution{Mean: 49.72666497, StandardDeviation: 2.440000057220459}
	contentydists[20] = statistics.NormalDistribution{Mean: 54.08442275, StandardDeviation: 2.440000057220459}
	contentydists[21] = statistics.NormalDistribution{Mean: 58.02100023, StandardDeviation: 2.440000057220459}
	contentydists[22] = statistics.NormalDistribution{Mean: 61.95757771, StandardDeviation: 2.440000057220459}
	contentydists[23] = statistics.NormalDistribution{Mean: 64.97327353, StandardDeviation: 2.440000057220459}
	contentydists[24] = statistics.NormalDistribution{Mean: 65.01145199, StandardDeviation: 2.440000057220459}
	var structuredamagefunctionStochastic = paireddata.UncertaintyPairedData{Xvals: xvals, Yvals: structureydists}
	var contentdamagefunctionStochastic = paireddata.UncertaintyPairedData{Xvals: xvals, Yvals: contentydists}
	sm := make(map[hazards.Parameter]interface{})
	var sdf = gstructs.DamageFunctionFamilyStochastic{DamageFunctions: sm}

	cm := make(map[hazards.Parameter]interface{})
	var cdf = gstructs.DamageFunctionFamilyStochastic{DamageFunctions: cm}
	//Default hazard.
	sdf.DamageFunctions[hazards.Default] = structuredamagefunctionStochastic
	cdf.DamageFunctions[hazards.Default] = contentdamagefunctionStochastic

	return gstructs.OccupancyTypeStochastic{Name: "RES1-SLWB", StructureDFF: sdf, ContentDFF: cdf}
}
func res3a() gstructs.OccupancyTypeStochastic {
	xvals := []float64{-8.0, -7.0, -6.0, -5.0, -4.0, -3.0, -2.0, -1.0, 0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0}
	structureydists := make([]statistics.ContinuousDistribution, 25)
	structureydists[0] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[1] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[2] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[3] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[4] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[5] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[6] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[7] = statistics.NormalDistribution{Mean: 2.653144212, StandardDeviation: 0.30000001192092896}
	structureydists[8] = statistics.NormalDistribution{Mean: 6.795963659, StandardDeviation: 1}
	structureydists[9] = statistics.NormalDistribution{Mean: 12.08185272, StandardDeviation: 1.5}
	structureydists[10] = statistics.NormalDistribution{Mean: 15.68959506, StandardDeviation: 2.0999999046325684}
	structureydists[11] = statistics.NormalDistribution{Mean: 17.87611447, StandardDeviation: 2.5999999046325684}
	structureydists[12] = statistics.NormalDistribution{Mean: 20.58109678, StandardDeviation: 3}
	structureydists[13] = statistics.NormalDistribution{Mean: 24.24706094, StandardDeviation: 3.2000000476837158}
	structureydists[14] = statistics.NormalDistribution{Mean: 29.21343129, StandardDeviation: 3.5}
	structureydists[15] = statistics.NormalDistribution{Mean: 35.65241119, StandardDeviation: 3.5499999523162842}
	structureydists[16] = statistics.NormalDistribution{Mean: 43.33923708, StandardDeviation: 3.5999999046325684}
	structureydists[17] = statistics.NormalDistribution{Mean: 51.50952162, StandardDeviation: 3.6500000953674316}
	structureydists[18] = statistics.NormalDistribution{Mean: 58.73172842, StandardDeviation: 3.7000000476837158}
	structureydists[19] = statistics.NormalDistribution{Mean: 64.77799354, StandardDeviation: 3.7300000190734863}
	structureydists[20] = statistics.NormalDistribution{Mean: 68.57021265, StandardDeviation: 3.7699999809265137}
	structureydists[21] = statistics.NormalDistribution{Mean: 70.2571299, StandardDeviation: 3.7799999713897705}
	structureydists[22] = statistics.NormalDistribution{Mean: 71.94404714, StandardDeviation: 3.7899999618530273}
	structureydists[23] = statistics.NormalDistribution{Mean: 73.57856024, StandardDeviation: 3.7999999523162842}
	structureydists[24] = statistics.NormalDistribution{Mean: 75.04363326, StandardDeviation: 3.7999999523162842}
	contentydists := make([]statistics.ContinuousDistribution, 25)
	contentydists[0] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[1] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[2] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[3] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[4] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[5] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[6] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[7] = statistics.NormalDistribution{Mean: 5.761828871, StandardDeviation: 0.05000000074505806}
	contentydists[8] = statistics.NormalDistribution{Mean: 14.75878297, StandardDeviation: 0.5}
	contentydists[9] = statistics.NormalDistribution{Mean: 22.9234344, StandardDeviation: 0.89999997615814209}
	contentydists[10] = statistics.NormalDistribution{Mean: 26.42803974, StandardDeviation: 1.2999999523162842}
	contentydists[11] = statistics.NormalDistribution{Mean: 29.1072136, StandardDeviation: 1.7000000476837158}
	contentydists[12] = statistics.NormalDistribution{Mean: 31.72280577, StandardDeviation: 2}
	contentydists[13] = statistics.NormalDistribution{Mean: 34.55707845, StandardDeviation: 2.2999999523162842}
	contentydists[14] = statistics.NormalDistribution{Mean: 37.73179607, StandardDeviation: 2.9000000953674316}
	contentydists[15] = statistics.NormalDistribution{Mean: 41.26328754, StandardDeviation: 3.0999999046325684}
	contentydists[16] = statistics.NormalDistribution{Mean: 45.07268587, StandardDeviation: 3.2999999523162842}
	contentydists[17] = statistics.NormalDistribution{Mean: 48.99708898, StandardDeviation: 3.5}
	contentydists[18] = statistics.NormalDistribution{Mean: 52.72125582, StandardDeviation: 3.5999999046325684}
	contentydists[19] = statistics.NormalDistribution{Mean: 56.16067757, StandardDeviation: 3.7000000476837158}
	contentydists[20] = statistics.NormalDistribution{Mean: 58.88887261, StandardDeviation: 3.7999999523162842}
	contentydists[21] = statistics.NormalDistribution{Mean: 60.95277468, StandardDeviation: 3.9000000953674316}
	contentydists[22] = statistics.NormalDistribution{Mean: 63.01667676, StandardDeviation: 3.9000000953674316}
	contentydists[23] = statistics.NormalDistribution{Mean: 64.82002884, StandardDeviation: 3.9000000953674316}
	contentydists[24] = statistics.NormalDistribution{Mean: 65.78093593, StandardDeviation: 3.9000000953674316}
	var structuredamagefunctionStochastic = paireddata.UncertaintyPairedData{Xvals: xvals, Yvals: structureydists}
	var contentdamagefunctionStochastic = paireddata.UncertaintyPairedData{Xvals: xvals, Yvals: contentydists}
	sm := make(map[hazards.Parameter]interface{})
	var sdf = gstructs.DamageFunctionFamilyStochastic{DamageFunctions: sm}

	cm := make(map[hazards.Parameter]interface{})
	var cdf = gstructs.DamageFunctionFamilyStochastic{DamageFunctions: cm}
	//Default hazard.
	sdf.DamageFunctions[hazards.Default] = structuredamagefunctionStochastic
	cdf.DamageFunctions[hazards.Default] = contentdamagefunctionStochastic

	return gstructs.OccupancyTypeStochastic{Name: "RES3A", StructureDFF: sdf, ContentDFF: cdf}
}
func res3b() gstructs.OccupancyTypeStochastic {
	xvals := []float64{-8.0, -7.0, -6.0, -5.0, -4.0, -3.0, -2.0, -1.0, 0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0}
	structureydists := make([]statistics.ContinuousDistribution, 25)
	structureydists[0] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[1] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[2] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[3] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[4] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[5] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[6] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[7] = statistics.NormalDistribution{Mean: 2.653144212, StandardDeviation: 0.30000001192092896}
	structureydists[8] = statistics.NormalDistribution{Mean: 6.795963659, StandardDeviation: 1}
	structureydists[9] = statistics.NormalDistribution{Mean: 12.08185272, StandardDeviation: 1.5}
	structureydists[10] = statistics.NormalDistribution{Mean: 15.68959506, StandardDeviation: 2.0999999046325684}
	structureydists[11] = statistics.NormalDistribution{Mean: 17.87611447, StandardDeviation: 2.5999999046325684}
	structureydists[12] = statistics.NormalDistribution{Mean: 20.58109678, StandardDeviation: 3}
	structureydists[13] = statistics.NormalDistribution{Mean: 24.24706094, StandardDeviation: 3.2000000476837158}
	structureydists[14] = statistics.NormalDistribution{Mean: 29.21343129, StandardDeviation: 3.5}
	structureydists[15] = statistics.NormalDistribution{Mean: 35.65241119, StandardDeviation: 3.5499999523162842}
	structureydists[16] = statistics.NormalDistribution{Mean: 43.33923708, StandardDeviation: 3.5999999046325684}
	structureydists[17] = statistics.NormalDistribution{Mean: 51.50952162, StandardDeviation: 3.6500000953674316}
	structureydists[18] = statistics.NormalDistribution{Mean: 58.73172842, StandardDeviation: 3.7000000476837158}
	structureydists[19] = statistics.NormalDistribution{Mean: 64.77799354, StandardDeviation: 3.7300000190734863}
	structureydists[20] = statistics.NormalDistribution{Mean: 68.57021265, StandardDeviation: 3.7699999809265137}
	structureydists[21] = statistics.NormalDistribution{Mean: 70.2571299, StandardDeviation: 3.7799999713897705}
	structureydists[22] = statistics.NormalDistribution{Mean: 71.94404714, StandardDeviation: 3.7899999618530273}
	structureydists[23] = statistics.NormalDistribution{Mean: 73.57856024, StandardDeviation: 3.7999999523162842}
	structureydists[24] = statistics.NormalDistribution{Mean: 75.04363326, StandardDeviation: 3.7999999523162842}
	contentydists := make([]statistics.ContinuousDistribution, 25)
	contentydists[0] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[1] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[2] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[3] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[4] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[5] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[6] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[7] = statistics.NormalDistribution{Mean: 5.761828871, StandardDeviation: 0.05000000074505806}
	contentydists[8] = statistics.NormalDistribution{Mean: 14.75878297, StandardDeviation: 0.5}
	contentydists[9] = statistics.NormalDistribution{Mean: 22.9234344, StandardDeviation: 0.89999997615814209}
	contentydists[10] = statistics.NormalDistribution{Mean: 26.42803974, StandardDeviation: 1.2999999523162842}
	contentydists[11] = statistics.NormalDistribution{Mean: 29.1072136, StandardDeviation: 1.7000000476837158}
	contentydists[12] = statistics.NormalDistribution{Mean: 31.72280577, StandardDeviation: 2}
	contentydists[13] = statistics.NormalDistribution{Mean: 34.55707845, StandardDeviation: 2.2999999523162842}
	contentydists[14] = statistics.NormalDistribution{Mean: 37.73179607, StandardDeviation: 2.9000000953674316}
	contentydists[15] = statistics.NormalDistribution{Mean: 41.26328754, StandardDeviation: 3.0999999046325684}
	contentydists[16] = statistics.NormalDistribution{Mean: 45.07268587, StandardDeviation: 3.2999999523162842}
	contentydists[17] = statistics.NormalDistribution{Mean: 48.99708898, StandardDeviation: 3.5}
	contentydists[18] = statistics.NormalDistribution{Mean: 52.72125582, StandardDeviation: 3.5999999046325684}
	contentydists[19] = statistics.NormalDistribution{Mean: 56.16067757, StandardDeviation: 3.7000000476837158}
	contentydists[20] = statistics.NormalDistribution{Mean: 58.88887261, StandardDeviation: 3.7999999523162842}
	contentydists[21] = statistics.NormalDistribution{Mean: 60.95277468, StandardDeviation: 3.9000000953674316}
	contentydists[22] = statistics.NormalDistribution{Mean: 63.01667676, StandardDeviation: 3.9000000953674316}
	contentydists[23] = statistics.NormalDistribution{Mean: 64.82002884, StandardDeviation: 3.9000000953674316}
	contentydists[24] = statistics.NormalDistribution{Mean: 65.78093593, StandardDeviation: 3.9000000953674316}
	var structuredamagefunctionStochastic = paireddata.UncertaintyPairedData{Xvals: xvals, Yvals: structureydists}
	var contentdamagefunctionStochastic = paireddata.UncertaintyPairedData{Xvals: xvals, Yvals: contentydists}
	sm := make(map[hazards.Parameter]interface{})
	var sdf = gstructs.DamageFunctionFamilyStochastic{DamageFunctions: sm}

	cm := make(map[hazards.Parameter]interface{})
	var cdf = gstructs.DamageFunctionFamilyStochastic{DamageFunctions: cm}
	//Default hazard.
	sdf.DamageFunctions[hazards.Default] = structuredamagefunctionStochastic
	cdf.DamageFunctions[hazards.Default] = contentdamagefunctionStochastic

	return gstructs.OccupancyTypeStochastic{Name: "RES3B", StructureDFF: sdf, ContentDFF: cdf}
}
func res3c() gstructs.OccupancyTypeStochastic {
	xvals := []float64{-8.0, -7.0, -6.0, -5.0, -4.0, -3.0, -2.0, -1.0, 0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0}
	structureydists := make([]statistics.ContinuousDistribution, 25)
	structureydists[0] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[1] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[2] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[3] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[4] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[5] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[6] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[7] = statistics.NormalDistribution{Mean: 2.653144212, StandardDeviation: 0.30000001192092896}
	structureydists[8] = statistics.NormalDistribution{Mean: 6.795963659, StandardDeviation: 1}
	structureydists[9] = statistics.NormalDistribution{Mean: 12.08185272, StandardDeviation: 1.5}
	structureydists[10] = statistics.NormalDistribution{Mean: 15.68959506, StandardDeviation: 2.0999999046325684}
	structureydists[11] = statistics.NormalDistribution{Mean: 17.87611447, StandardDeviation: 2.5999999046325684}
	structureydists[12] = statistics.NormalDistribution{Mean: 20.58109678, StandardDeviation: 3}
	structureydists[13] = statistics.NormalDistribution{Mean: 24.24706094, StandardDeviation: 3.2000000476837158}
	structureydists[14] = statistics.NormalDistribution{Mean: 29.21343129, StandardDeviation: 3.5}
	structureydists[15] = statistics.NormalDistribution{Mean: 35.65241119, StandardDeviation: 3.5499999523162842}
	structureydists[16] = statistics.NormalDistribution{Mean: 43.33923708, StandardDeviation: 3.5999999046325684}
	structureydists[17] = statistics.NormalDistribution{Mean: 51.50952162, StandardDeviation: 3.6500000953674316}
	structureydists[18] = statistics.NormalDistribution{Mean: 58.73172842, StandardDeviation: 3.7000000476837158}
	structureydists[19] = statistics.NormalDistribution{Mean: 64.77799354, StandardDeviation: 3.7300000190734863}
	structureydists[20] = statistics.NormalDistribution{Mean: 68.57021265, StandardDeviation: 3.7699999809265137}
	structureydists[21] = statistics.NormalDistribution{Mean: 70.2571299, StandardDeviation: 3.7799999713897705}
	structureydists[22] = statistics.NormalDistribution{Mean: 71.94404714, StandardDeviation: 3.7899999618530273}
	structureydists[23] = statistics.NormalDistribution{Mean: 73.57856024, StandardDeviation: 3.7999999523162842}
	structureydists[24] = statistics.NormalDistribution{Mean: 75.04363326, StandardDeviation: 3.7999999523162842}
	contentydists := make([]statistics.ContinuousDistribution, 25)
	contentydists[0] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[1] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[2] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[3] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[4] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[5] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[6] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[7] = statistics.NormalDistribution{Mean: 5.761828871, StandardDeviation: 0.05000000074505806}
	contentydists[8] = statistics.NormalDistribution{Mean: 14.75878297, StandardDeviation: 0.5}
	contentydists[9] = statistics.NormalDistribution{Mean: 22.9234344, StandardDeviation: 0.89999997615814209}
	contentydists[10] = statistics.NormalDistribution{Mean: 26.42803974, StandardDeviation: 1.2999999523162842}
	contentydists[11] = statistics.NormalDistribution{Mean: 29.1072136, StandardDeviation: 1.7000000476837158}
	contentydists[12] = statistics.NormalDistribution{Mean: 31.72280577, StandardDeviation: 2}
	contentydists[13] = statistics.NormalDistribution{Mean: 34.55707845, StandardDeviation: 2.2999999523162842}
	contentydists[14] = statistics.NormalDistribution{Mean: 37.73179607, StandardDeviation: 2.9000000953674316}
	contentydists[15] = statistics.NormalDistribution{Mean: 41.26328754, StandardDeviation: 3.0999999046325684}
	contentydists[16] = statistics.NormalDistribution{Mean: 45.07268587, StandardDeviation: 3.2999999523162842}
	contentydists[17] = statistics.NormalDistribution{Mean: 48.99708898, StandardDeviation: 3.5}
	contentydists[18] = statistics.NormalDistribution{Mean: 52.72125582, StandardDeviation: 3.5999999046325684}
	contentydists[19] = statistics.NormalDistribution{Mean: 56.16067757, StandardDeviation: 3.7000000476837158}
	contentydists[20] = statistics.NormalDistribution{Mean: 58.88887261, StandardDeviation: 3.7999999523162842}
	contentydists[21] = statistics.NormalDistribution{Mean: 60.95277468, StandardDeviation: 3.9000000953674316}
	contentydists[22] = statistics.NormalDistribution{Mean: 63.01667676, StandardDeviation: 3.9000000953674316}
	contentydists[23] = statistics.NormalDistribution{Mean: 64.82002884, StandardDeviation: 3.9000000953674316}
	contentydists[24] = statistics.NormalDistribution{Mean: 65.78093593, StandardDeviation: 3.9000000953674316}
	var structuredamagefunctionStochastic = paireddata.UncertaintyPairedData{Xvals: xvals, Yvals: structureydists}
	var contentdamagefunctionStochastic = paireddata.UncertaintyPairedData{Xvals: xvals, Yvals: contentydists}
	sm := make(map[hazards.Parameter]interface{})
	var sdf = gstructs.DamageFunctionFamilyStochastic{DamageFunctions: sm}

	cm := make(map[hazards.Parameter]interface{})
	var cdf = gstructs.DamageFunctionFamilyStochastic{DamageFunctions: cm}
	//Default hazard.
	sdf.DamageFunctions[hazards.Default] = structuredamagefunctionStochastic
	cdf.DamageFunctions[hazards.Default] = contentdamagefunctionStochastic

	return gstructs.OccupancyTypeStochastic{Name: "RES3C", StructureDFF: sdf, ContentDFF: cdf}
}
func res3d() gstructs.OccupancyTypeStochastic {
	xvals := []float64{-8.0, -7.0, -6.0, -5.0, -4.0, -3.0, -2.0, -1.0, 0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0}
	structureydists := make([]statistics.ContinuousDistribution, 25)
	structureydists[0] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[1] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[2] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[3] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[4] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[5] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[6] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[7] = statistics.NormalDistribution{Mean: 3.830024524, StandardDeviation: 0.30000001192092896}
	structureydists[8] = statistics.NormalDistribution{Mean: 9.810513636, StandardDeviation: 1}
	structureydists[9] = statistics.NormalDistribution{Mean: 11.46303702, StandardDeviation: 1.5}
	structureydists[10] = statistics.NormalDistribution{Mean: 12.09134065, StandardDeviation: 2.0999999046325684}
	structureydists[11] = statistics.NormalDistribution{Mean: 12.59532742, StandardDeviation: 2.5999999046325684}
	structureydists[12] = statistics.NormalDistribution{Mean: 13.02717253, StandardDeviation: 3}
	structureydists[13] = statistics.NormalDistribution{Mean: 13.41002329, StandardDeviation: 3.2000000476837158}
	structureydists[14] = statistics.NormalDistribution{Mean: 13.75688355, StandardDeviation: 3.5}
	structureydists[15] = statistics.NormalDistribution{Mean: 14.0756337, StandardDeviation: 3.5499999523162842}
	structureydists[16] = statistics.NormalDistribution{Mean: 14.37158235, StandardDeviation: 3.5999999046325684}
	structureydists[17] = statistics.NormalDistribution{Mean: 14.64847612, StandardDeviation: 3.6500000953674316}
	structureydists[18] = statistics.NormalDistribution{Mean: 14.9058736, StandardDeviation: 3.7000000476837158}
	structureydists[19] = statistics.NormalDistribution{Mean: 15.15145096, StandardDeviation: 3.7300000190734863}
	structureydists[20] = statistics.NormalDistribution{Mean: 15.38234764, StandardDeviation: 3.7699999809265137}
	structureydists[21] = statistics.NormalDistribution{Mean: 15.59953242, StandardDeviation: 3.7799999713897705}
	structureydists[22] = statistics.NormalDistribution{Mean: 15.8167172, StandardDeviation: 3.7899999618530273}
	structureydists[23] = statistics.NormalDistribution{Mean: 16.02737209, StandardDeviation: 3.7999999523162842}
	structureydists[24] = statistics.NormalDistribution{Mean: 16.21691368, StandardDeviation: 3.7999999523162842}
	contentydists := make([]statistics.ContinuousDistribution, 25)
	contentydists[0] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[1] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[2] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[3] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[4] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[5] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[6] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[7] = statistics.NormalDistribution{Mean: 3.620482647, StandardDeviation: 0.05000000074505806}
	contentydists[8] = statistics.NormalDistribution{Mean: 9.273777273, StandardDeviation: 0.5}
	contentydists[9] = statistics.NormalDistribution{Mean: 16.88241833, StandardDeviation: 0.89999997615814209}
	contentydists[10] = statistics.NormalDistribution{Mean: 21.3011933, StandardDeviation: 1.2999999523162842}
	contentydists[11] = statistics.NormalDistribution{Mean: 24.56464375, StandardDeviation: 1.7000000476837158}
	contentydists[12] = statistics.NormalDistribution{Mean: 27.21140912, StandardDeviation: 2}
	contentydists[13] = statistics.NormalDistribution{Mean: 29.44113737, StandardDeviation: 2.2999999523162842}
	contentydists[14] = statistics.NormalDistribution{Mean: 31.35873606, StandardDeviation: 2.9000000953674316}
	contentydists[15] = statistics.NormalDistribution{Mean: 33.02969389, StandardDeviation: 3.0999999046325684}
	contentydists[16] = statistics.NormalDistribution{Mean: 34.49982529, StandardDeviation: 3.2999999523162842}
	contentydists[17] = statistics.NormalDistribution{Mean: 35.80323759, StandardDeviation: 3.5}
	contentydists[18] = statistics.NormalDistribution{Mean: 36.93941462, StandardDeviation: 3.5999999046325684}
	contentydists[19] = statistics.NormalDistribution{Mean: 37.97522304, StandardDeviation: 3.7000000476837158}
	contentydists[20] = statistics.NormalDistribution{Mean: 38.89276994, StandardDeviation: 3.7999999523162842}
	contentydists[21] = statistics.NormalDistribution{Mean: 39.6998594, StandardDeviation: 3.9000000953674316}
	contentydists[22] = statistics.NormalDistribution{Mean: 40.50694885, StandardDeviation: 3.9000000953674316}
	contentydists[23] = statistics.NormalDistribution{Mean: 41.26605536, StandardDeviation: 3.9000000953674316}
	contentydists[24] = statistics.NormalDistribution{Mean: 41.87001703, StandardDeviation: 3.9000000953674316}
	var structuredamagefunctionStochastic = paireddata.UncertaintyPairedData{Xvals: xvals, Yvals: structureydists}
	var contentdamagefunctionStochastic = paireddata.UncertaintyPairedData{Xvals: xvals, Yvals: contentydists}
	sm := make(map[hazards.Parameter]interface{})
	var sdf = gstructs.DamageFunctionFamilyStochastic{DamageFunctions: sm}

	cm := make(map[hazards.Parameter]interface{})
	var cdf = gstructs.DamageFunctionFamilyStochastic{DamageFunctions: cm}
	//Default hazard.
	sdf.DamageFunctions[hazards.Default] = structuredamagefunctionStochastic
	cdf.DamageFunctions[hazards.Default] = contentdamagefunctionStochastic

	return gstructs.OccupancyTypeStochastic{Name: "RES3D", StructureDFF: sdf, ContentDFF: cdf}
}
func res3e() gstructs.OccupancyTypeStochastic {
	xvals := []float64{-8.0, -7.0, -6.0, -5.0, -4.0, -3.0, -2.0, -1.0, 0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0}
	structureydists := make([]statistics.ContinuousDistribution, 25)
	structureydists[0] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[1] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[2] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[3] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[4] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[5] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[6] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[7] = statistics.NormalDistribution{Mean: 3.830024524, StandardDeviation: 0.30000001192092896}
	structureydists[8] = statistics.NormalDistribution{Mean: 9.810513636, StandardDeviation: 1}
	structureydists[9] = statistics.NormalDistribution{Mean: 11.46303702, StandardDeviation: 1.5}
	structureydists[10] = statistics.NormalDistribution{Mean: 12.09134065, StandardDeviation: 2.0999999046325684}
	structureydists[11] = statistics.NormalDistribution{Mean: 12.59532742, StandardDeviation: 2.5999999046325684}
	structureydists[12] = statistics.NormalDistribution{Mean: 13.02717253, StandardDeviation: 3}
	structureydists[13] = statistics.NormalDistribution{Mean: 13.41002329, StandardDeviation: 3.2000000476837158}
	structureydists[14] = statistics.NormalDistribution{Mean: 13.75688355, StandardDeviation: 3.5}
	structureydists[15] = statistics.NormalDistribution{Mean: 14.0756337, StandardDeviation: 3.5499999523162842}
	structureydists[16] = statistics.NormalDistribution{Mean: 14.37158235, StandardDeviation: 3.5999999046325684}
	structureydists[17] = statistics.NormalDistribution{Mean: 14.64847612, StandardDeviation: 3.6500000953674316}
	structureydists[18] = statistics.NormalDistribution{Mean: 14.9058736, StandardDeviation: 3.7000000476837158}
	structureydists[19] = statistics.NormalDistribution{Mean: 15.15145096, StandardDeviation: 3.7300000190734863}
	structureydists[20] = statistics.NormalDistribution{Mean: 15.38234764, StandardDeviation: 3.7699999809265137}
	structureydists[21] = statistics.NormalDistribution{Mean: 15.59953242, StandardDeviation: 3.7799999713897705}
	structureydists[22] = statistics.NormalDistribution{Mean: 15.8167172, StandardDeviation: 3.7899999618530273}
	structureydists[23] = statistics.NormalDistribution{Mean: 16.02737209, StandardDeviation: 3.7999999523162842}
	structureydists[24] = statistics.NormalDistribution{Mean: 16.21691368, StandardDeviation: 3.7999999523162842}
	contentydists := make([]statistics.ContinuousDistribution, 25)
	contentydists[0] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[1] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[2] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[3] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[4] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[5] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[6] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[7] = statistics.NormalDistribution{Mean: 3.620482647, StandardDeviation: 0.05000000074505806}
	contentydists[8] = statistics.NormalDistribution{Mean: 9.273777273, StandardDeviation: 0.5}
	contentydists[9] = statistics.NormalDistribution{Mean: 16.88241833, StandardDeviation: 0.89999997615814209}
	contentydists[10] = statistics.NormalDistribution{Mean: 21.3011933, StandardDeviation: 1.2999999523162842}
	contentydists[11] = statistics.NormalDistribution{Mean: 24.56464375, StandardDeviation: 1.7000000476837158}
	contentydists[12] = statistics.NormalDistribution{Mean: 27.21140912, StandardDeviation: 2}
	contentydists[13] = statistics.NormalDistribution{Mean: 29.44113737, StandardDeviation: 2.2999999523162842}
	contentydists[14] = statistics.NormalDistribution{Mean: 31.35873606, StandardDeviation: 2.9000000953674316}
	contentydists[15] = statistics.NormalDistribution{Mean: 33.02969389, StandardDeviation: 3.0999999046325684}
	contentydists[16] = statistics.NormalDistribution{Mean: 34.49982529, StandardDeviation: 3.2999999523162842}
	contentydists[17] = statistics.NormalDistribution{Mean: 35.80323759, StandardDeviation: 3.5}
	contentydists[18] = statistics.NormalDistribution{Mean: 36.93941462, StandardDeviation: 3.5999999046325684}
	contentydists[19] = statistics.NormalDistribution{Mean: 37.97522304, StandardDeviation: 3.7000000476837158}
	contentydists[20] = statistics.NormalDistribution{Mean: 38.89276994, StandardDeviation: 3.7999999523162842}
	contentydists[21] = statistics.NormalDistribution{Mean: 39.6998594, StandardDeviation: 3.9000000953674316}
	contentydists[22] = statistics.NormalDistribution{Mean: 40.50694885, StandardDeviation: 3.9000000953674316}
	contentydists[23] = statistics.NormalDistribution{Mean: 41.26605536, StandardDeviation: 3.9000000953674316}
	contentydists[24] = statistics.NormalDistribution{Mean: 41.87001703, StandardDeviation: 3.9000000953674316}
	var structuredamagefunctionStochastic = paireddata.UncertaintyPairedData{Xvals: xvals, Yvals: structureydists}
	var contentdamagefunctionStochastic = paireddata.UncertaintyPairedData{Xvals: xvals, Yvals: contentydists}
	sm := make(map[hazards.Parameter]interface{})
	var sdf = gstructs.DamageFunctionFamilyStochastic{DamageFunctions: sm}

	cm := make(map[hazards.Parameter]interface{})
	var cdf = gstructs.DamageFunctionFamilyStochastic{DamageFunctions: cm}
	//Default hazard.
	sdf.DamageFunctions[hazards.Default] = structuredamagefunctionStochastic
	cdf.DamageFunctions[hazards.Default] = contentdamagefunctionStochastic
	return gstructs.OccupancyTypeStochastic{Name: "RES3E", StructureDFF: sdf, ContentDFF: cdf}
}
func res3f() gstructs.OccupancyTypeStochastic {
	xvals := []float64{-8.0, -7.0, -6.0, -5.0, -4.0, -3.0, -2.0, -1.0, 0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0}
	structureydists := make([]statistics.ContinuousDistribution, 25)
	structureydists[0] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[1] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[2] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[3] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[4] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[5] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[6] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	structureydists[7] = statistics.NormalDistribution{Mean: 3.830024524, StandardDeviation: 0.30000001192092896}
	structureydists[8] = statistics.NormalDistribution{Mean: 9.810513636, StandardDeviation: 1}
	structureydists[9] = statistics.NormalDistribution{Mean: 11.46303702, StandardDeviation: 1.5}
	structureydists[10] = statistics.NormalDistribution{Mean: 12.09134065, StandardDeviation: 2.0999999046325684}
	structureydists[11] = statistics.NormalDistribution{Mean: 12.59532742, StandardDeviation: 2.5999999046325684}
	structureydists[12] = statistics.NormalDistribution{Mean: 13.02717253, StandardDeviation: 3}
	structureydists[13] = statistics.NormalDistribution{Mean: 13.41002329, StandardDeviation: 3.2000000476837158}
	structureydists[14] = statistics.NormalDistribution{Mean: 13.75688355, StandardDeviation: 3.5}
	structureydists[15] = statistics.NormalDistribution{Mean: 14.0756337, StandardDeviation: 3.5499999523162842}
	structureydists[16] = statistics.NormalDistribution{Mean: 14.37158235, StandardDeviation: 3.5999999046325684}
	structureydists[17] = statistics.NormalDistribution{Mean: 14.64847612, StandardDeviation: 3.6500000953674316}
	structureydists[18] = statistics.NormalDistribution{Mean: 14.9058736, StandardDeviation: 3.7000000476837158}
	structureydists[19] = statistics.NormalDistribution{Mean: 15.15145096, StandardDeviation: 3.7300000190734863}
	structureydists[20] = statistics.NormalDistribution{Mean: 15.38234764, StandardDeviation: 3.7699999809265137}
	structureydists[21] = statistics.NormalDistribution{Mean: 15.59953242, StandardDeviation: 3.7799999713897705}
	structureydists[22] = statistics.NormalDistribution{Mean: 15.8167172, StandardDeviation: 3.7899999618530273}
	structureydists[23] = statistics.NormalDistribution{Mean: 16.02737209, StandardDeviation: 3.7999999523162842}
	structureydists[24] = statistics.NormalDistribution{Mean: 16.21691368, StandardDeviation: 3.7999999523162842}
	contentydists := make([]statistics.ContinuousDistribution, 25)
	contentydists[0] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[1] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[2] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[3] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[4] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[5] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[6] = statistics.NormalDistribution{Mean: 0, StandardDeviation: 0}
	contentydists[7] = statistics.NormalDistribution{Mean: 3.620482647, StandardDeviation: 0.05000000074505806}
	contentydists[8] = statistics.NormalDistribution{Mean: 9.273777273, StandardDeviation: 0.5}
	contentydists[9] = statistics.NormalDistribution{Mean: 16.88241833, StandardDeviation: 0.89999997615814209}
	contentydists[10] = statistics.NormalDistribution{Mean: 21.3011933, StandardDeviation: 1.2999999523162842}
	contentydists[11] = statistics.NormalDistribution{Mean: 24.56464375, StandardDeviation: 1.7000000476837158}
	contentydists[12] = statistics.NormalDistribution{Mean: 27.21140912, StandardDeviation: 2}
	contentydists[13] = statistics.NormalDistribution{Mean: 29.44113737, StandardDeviation: 2.2999999523162842}
	contentydists[14] = statistics.NormalDistribution{Mean: 31.35873606, StandardDeviation: 2.9000000953674316}
	contentydists[15] = statistics.NormalDistribution{Mean: 33.02969389, StandardDeviation: 3.0999999046325684}
	contentydists[16] = statistics.NormalDistribution{Mean: 34.49982529, StandardDeviation: 3.2999999523162842}
	contentydists[17] = statistics.NormalDistribution{Mean: 35.80323759, StandardDeviation: 3.5}
	contentydists[18] = statistics.NormalDistribution{Mean: 36.93941462, StandardDeviation: 3.5999999046325684}
	contentydists[19] = statistics.NormalDistribution{Mean: 37.97522304, StandardDeviation: 3.7000000476837158}
	contentydists[20] = statistics.NormalDistribution{Mean: 38.89276994, StandardDeviation: 3.7999999523162842}
	contentydists[21] = statistics.NormalDistribution{Mean: 39.6998594, StandardDeviation: 3.9000000953674316}
	contentydists[22] = statistics.NormalDistribution{Mean: 40.50694885, StandardDeviation: 3.9000000953674316}
	contentydists[23] = statistics.NormalDistribution{Mean: 41.26605536, StandardDeviation: 3.9000000953674316}
	contentydists[24] = statistics.NormalDistribution{Mean: 41.87001703, StandardDeviation: 3.9000000953674316}
	var structuredamagefunctionStochastic = paireddata.UncertaintyPairedData{Xvals: xvals, Yvals: structureydists}
	var contentdamagefunctionStochastic = paireddata.UncertaintyPairedData{Xvals: xvals, Yvals: contentydists}
	sm := make(map[hazards.Parameter]interface{})
	var sdf = gstructs.DamageFunctionFamilyStochastic{DamageFunctions: sm}

	cm := make(map[hazards.Parameter]interface{})
	var cdf = gstructs.DamageFunctionFamilyStochastic{DamageFunctions: cm}
	//Default hazard.
	sdf.DamageFunctions[hazards.Default] = structuredamagefunctionStochastic
	cdf.DamageFunctions[hazards.Default] = contentdamagefunctionStochastic

	return gstructs.OccupancyTypeStochastic{Name: "RES3F", StructureDFF: sdf, ContentDFF: cdf}
}
