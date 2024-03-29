package hazard_providers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/USACE/go-consequences/hazardproviders"
	"github.com/USACE/go-consequences/hazards"
)

//"C:\Users\Q0HECWPL\Documents\NSI\NSI_Fathom_depths\NSI_Fathom_depths.csv" - old data
//new data in state folders.

//FrequencyData is a structure to describe the variables in the csv files.
type FrequencyData struct {
	fluvial bool      //false is pluvial //new data has false is coastal and fluvial (max of pluvial and fluvial really)
	year    int       //2020, 2050
	Values  []float64 //2yr, 5yr,20yr,100yr,250yr,500yr // new data adds 2 year
}
type DataSet struct {
	Data map[string]Record
}
type Record struct {
	Fd_id          string
	FutureFluvial  FrequencyData
	FuturePluvial  FrequencyData //coastal?
	CurrentFluvial FrequencyData
	CurrentPluvial FrequencyData //coastal?
}
type FathomEvent struct {
	Year      int
	Fluvial   bool //false is coastal?
	Frequency int  //2,5,20,100,250,500
}
type FathomQuery struct {
	Fd_id string
	FathomEvent
	NewData bool
}

func (ds DataSet) ProvideHazard(args interface{}) (hazards.HazardEvent, error) {
	fd_id, ok := args.(FathomQuery)
	if ok {
		r, found := ds.Data[fd_id.Fd_id]
		if found {
			if fd_id.Fluvial {
				if fd_id.Year == 2020 {
					return generateDepthEvent(fd_id.Frequency, r.CurrentFluvial, fd_id.NewData)
				} else if fd_id.Year == 2050 {
					return generateDepthEvent(fd_id.Frequency, r.FutureFluvial, fd_id.NewData)
				} else {
					//throw error?
					return nil, hazardproviders.HazardError{Input: "Bad Year Argument"}
				}

			} else {
				if fd_id.Year == 2020 {
					return generateDepthEvent(fd_id.Frequency, r.CurrentPluvial, fd_id.NewData) //under new data it is coastal
				} else if fd_id.Year == 2050 {
					return generateDepthEvent(fd_id.Frequency, r.FuturePluvial, fd_id.NewData) //under new data it is coastal
				} else {
					//throw error?
					return nil, hazardproviders.HazardError{Input: "Bad Year Argument"}
				}
			}
		} else {
			return nil, hazardproviders.NoHazardFoundError{Input: fd_id.Fd_id}
		}
	} else {
		return nil, hazardproviders.HazardError{Input: "Unable to parse args to hazard_providers.FathomQuery"}
	}
}
func generateDepthEvent(frequency int, data FrequencyData, newData bool) (hazards.DepthEvent, error) {
	switch frequency {
	case 2:
		if newData {
			h := hazards.DepthEvent{}
			h.SetDepth(data.Values[0])
			return h, nil
		}
		return hazards.DepthEvent{}, hazardproviders.NoFrequencyFoundError{Input: fmt.Sprintf("%v", frequency)} //bad frequency
	case 5:
		if newData {
			h := hazards.DepthEvent{}
			h.SetDepth(data.Values[1])
			return h, nil
		}
		h := hazards.DepthEvent{}
		h.SetDepth(data.Values[0])
		return h, nil
	case 20:
		if newData {
			h := hazards.DepthEvent{}
			h.SetDepth(data.Values[2])
			return h, nil
		}
		h := hazards.DepthEvent{}
		h.SetDepth(data.Values[1])
		return h, nil
	case 100:
		if newData {
			h := hazards.DepthEvent{}
			h.SetDepth(data.Values[3])
			return h, nil
		}
		h := hazards.DepthEvent{}
		h.SetDepth(data.Values[2])
		return h, nil
	case 250:
		if newData {
			h := hazards.DepthEvent{}
			h.SetDepth(data.Values[4])
			return h, nil
		}
		h := hazards.DepthEvent{}
		h.SetDepth(data.Values[3])
		return h, nil
	case 500:
		if newData {
			h := hazards.DepthEvent{}
			h.SetDepth(data.Values[5])
			return h, nil
		}
		h := hazards.DepthEvent{}
		h.SetDepth(data.Values[4])
		return h, nil
	default:
		return hazards.DepthEvent{}, hazardproviders.NoFrequencyFoundError{Input: fmt.Sprintf("%v", frequency)} //bad frequency
	}
}
func ConvertFile(file string) DataSet {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return DataSet{}
	}
	scanner := bufio.NewScanner(f)
	if err != nil {
		return DataSet{}
	}
	scanner.Scan()
	//fmt.Println(scanner.Text()) //header row
	m := make(map[string]Record)
	newData := strings.Contains(scanner.Text(), "cstl") //not present in old data.
	count := 0
	for scanner.Scan() {
		lines := strings.Split(scanner.Text(), ",")
		fd_id := lines[0]
		//fmt.Println(fd_id)
		//fluv_2020_5yr,pluv_2020_5yr,fluv_2020_20yr,pluv_2020_20yr,fluv_2020_100yr,pluv_2020_100yr,fluv_2020_250yr,pluv_2020_250yr,fluv_2020_500yr,pluv_2020_500yr,fluv_2050_5yr,pluv_2050_5yr,fluv_2050_20yr,pluv_2050_20yr,fluv_2050_100yr,pluv_2050_100yr,fluv_2050_250yr,pluv_2050_250yr,fluv_2050_500yr,pluv_2050_500yr
		//,fluv_2020_2yr,cstl_2020_2yr,fluv_2020_5yr,cstl_2020_5yr,fluv_2020_20yr,cstl_2020_20yr,fluv_2020_100yr,cstl_2020_100yr,fluv_2020_250yr,cstl_2020_250yr,fluv_2020_500yr,cstl_2020_500yr,fluv_2050_2yr,cstl_2050_2yr,fluv_2050_5yr,cstl_2050_5yr,fluv_2050_20yr,cstl_2050_20yr,fluv_2050_100yr,cstl_2050_100yr,fluv_2050_250yr,cstl_2050_250yr,fluv_2050_500yr,cstl_2050_500yr
		elements := 10
		size := 5
		if newData {
			elements = 10
			size = 5
		}
		fluvial := true
		cfvals := make([]float64, size)
		cpvals := make([]float64, size)
		ffvals := make([]float64, size)
		fpvals := make([]float64, size)
		twentyTwenty := 0
		fpidx := 0
		ffidx := 0
		cpidx := 0
		cfidx := 0
		for i := 1; i < len(lines); i++ {
			if twentyTwenty >= elements { //new data requires this to be 12, old data 10... (to accomidate the 2 year hazard.)
				//2050
				if fluvial {
					ffvals[ffidx], err = strconv.ParseFloat(lines[i], 64)
					if ffvals[ffidx] != 9999.0 {
						ffvals[ffidx] = ffvals[ffidx] / 30.48 //centimeters to feet
					}
					ffidx++
				} else {
					fpvals[fpidx], err = strconv.ParseFloat(lines[i], 64)
					if fpvals[fpidx] != 9999.0 {
						fpvals[fpidx] = fpvals[fpidx] / 30.48 //centimeters to feet
					}
					fpidx++ //new data coastal...
				}
			} else {
				//2020
				if fluvial {
					cfvals[cfidx], err = strconv.ParseFloat(lines[i], 64)
					if cfvals[cfidx] != 9999.0 {
						cfvals[cfidx] = cfvals[cfidx] / 30.48 //centimeters to feet
					}
					cfidx++
				} else {
					cpvals[cpidx], err = strconv.ParseFloat(lines[i], 64)
					if cpvals[cpidx] != 9999.0 {
						cpvals[cpidx] = cpvals[cpidx] / 30.48 //centimeters to feet
					}
					//fmt.Println("current pluvial") //new data coastal...
					cpidx++
				}
			}
			fluvial = !fluvial
			twentyTwenty++
		}
		futurefluvial := FrequencyData{fluvial: true, year: 2050, Values: ffvals}
		futurepluvial := FrequencyData{fluvial: false, year: 2050, Values: fpvals}
		currentfluvial := FrequencyData{fluvial: true, year: 2020, Values: cfvals}
		currentpluvial := FrequencyData{fluvial: false, year: 2020, Values: cpvals}
		if hasNonZeroValues(ffvals, fpvals, cfvals, cpvals, newData) {
			r := Record{Fd_id: fd_id, FutureFluvial: futurefluvial, FuturePluvial: futurepluvial, CurrentFluvial: currentfluvial, CurrentPluvial: currentpluvial}
			m[fd_id] = r
			count++
		} else {
			//skipping.
			fmt.Println("skipping " + fd_id)
		}

	}
	fmt.Println(count)
	ds := DataSet{Data: m}
	return ds
}
func ReadFeetFile(file string) DataSet {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return DataSet{}
	}
	scanner := bufio.NewScanner(f)
	if err != nil {
		return DataSet{}
	}
	scanner.Scan()
	fmt.Println(scanner.Text()) //header row
	m := make(map[string]Record)
	count := 0
	newData := strings.Contains(scanner.Text(), "cstl") //not present in old data.
	elements := 10
	size := 5
	if newData {
		elements = 10
		size = 5
	}
	for scanner.Scan() {
		lines := strings.Split(scanner.Text(), ",")
		fd_id := lines[0]
		//fmt.Println(fd_id)
		//fluv_2020_5yr,pluv_2020_5yr,fluv_2020_20yr,pluv_2020_20yr,fluv_2020_100yr,pluv_2020_100yr,fluv_2020_250yr,pluv_2020_250yr,fluv_2020_500yr,pluv_2020_500yr,fluv_2050_5yr,pluv_2050_5yr,fluv_2050_20yr,pluv_2050_20yr,fluv_2050_100yr,pluv_2050_100yr,fluv_2050_250yr,pluv_2050_250yr,fluv_2050_500yr,pluv_2050_500yr
		fluvial := true
		cfvals := make([]float64, size)
		cpvals := make([]float64, size)
		ffvals := make([]float64, size)
		fpvals := make([]float64, size)
		twentyTwenty := 0
		fpidx := 0
		ffidx := 0
		cpidx := 0
		cfidx := 0
		for i := 1; i < len(lines); i++ {
			if twentyTwenty >= elements {
				//2050
				if fluvial {
					ffvals[ffidx], err = strconv.ParseFloat(lines[i], 64)
					ffidx++
				} else {
					fpvals[fpidx], err = strconv.ParseFloat(lines[i], 64)
					fpidx++
				}
			} else {
				//2020
				if fluvial {
					cfvals[cfidx], err = strconv.ParseFloat(lines[i], 64)
					cfidx++
				} else {
					cpvals[cpidx], err = strconv.ParseFloat(lines[i], 64)
					cpidx++
				}
			}
			fluvial = !fluvial
			twentyTwenty++
		}
		futurefluvial := FrequencyData{fluvial: true, year: 2050, Values: ffvals}
		futurepluvial := FrequencyData{fluvial: false, year: 2050, Values: fpvals}
		currentfluvial := FrequencyData{fluvial: true, year: 2020, Values: cfvals}
		currentpluvial := FrequencyData{fluvial: false, year: 2020, Values: cpvals}
		if hasNonZeroValues(ffvals, fpvals, cfvals, cpvals, newData) {
			if hasValidData(fd_id, ffvals, fpvals, cfvals, cpvals, newData) {
				r := Record{Fd_id: fd_id, FutureFluvial: futurefluvial, FuturePluvial: futurepluvial, CurrentFluvial: currentfluvial, CurrentPluvial: currentpluvial}
				m[fd_id] = r
				count++
			}
		} else {
			//skipping.
		}

	}
	fmt.Println(count)
	ds := DataSet{Data: m}
	return ds
}
func hasNonZeroValues(ffvals []float64, fpvals []float64, cfvals []float64, cpvals []float64, newData bool) bool {
	records := len(ffvals)
	missingDataValue := 9999.0
	for i := 0; i < records; i++ {
		if ffvals[i] != missingDataValue {
			return true
		}
		if fpvals[i] != missingDataValue {
			return true
		}
		if cfvals[i] != missingDataValue {
			return true
		}
		if cpvals[i] != missingDataValue {
			return true
		}
	}
	return false
}
func hasValidData(fd_id string, ffvals []float64, fpvals []float64, cfvals []float64, cpvals []float64, newData bool) bool {
	//ff
	records := 5
	if newData {
		records = 5
	}
	ffvalid := true
	fpvalid := true
	cfvalid := true
	cpvalid := true
	datasetvalid := true
	for i := 1; i < records; i++ {
		if ffvals[i] < ffvals[i-1] {
			if ffvals[i-1] != 9999 {
				ffvalid = false
				datasetvalid = false
			}
		}
		if fpvals[i] < fpvals[i-1] {
			if fpvals[i-1] != 9999 {
				fpvalid = false
				datasetvalid = false
			}
		}
		if cfvals[i] < cfvals[i-1] {
			if cfvals[i-1] != 9999 {
				cfvalid = false
				datasetvalid = false
			}
		}
		if cpvals[i] < cpvals[i-1] {
			if cpvals[i-1] != 9999 {
				cpvalid = false
				datasetvalid = false
			}
		}
	}
	if !datasetvalid {
		s := fd_id + " is not valid for: "
		if !ffvalid {
			s += "future fluvial,"
		}
		if !fpvalid {
			s += "future pluvial,"
		}
		if !cfvalid {
			s += "current fluvial,"
		}
		if !cpvalid {
			s += "current pluvial,"
		}
		s = strings.Trim(s, ",")
		fmt.Println(s)
		return false
	}
	return true
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func WriteBackToDisk(ds DataSet, outputPath string, newData bool) {
	f, err := os.Create(outputPath)
	check(err)
	defer f.Close()
	//write header.
	//FD_ID,fluv_2020_5yr,pluv_2020_5yr,fluv_2020_20yr,pluv_2020_20yr,fluv_2020_100yr,pluv_2020_100yr,fluv_2020_250yr,pluv_2020_250yr,fluv_2020_500yr,pluv_2020_500yr,fluv_2050_5yr,pluv_2050_5yr,fluv_2050_20yr,pluv_2050_20yr,fluv_2050_100yr,pluv_2050_100yr,fluv_2050_250yr,pluv_2050_250yr,fluv_2050_500yr,pluv_2050_500yr
	w := bufio.NewWriter(f)
	if newData {
		w.WriteString("FD_ID,fluv_2020_5yr,cstl_2020_5yr,fluv_2020_20yr,cstl_2020_20yr,fluv_2020_100yr,cstl_2020_100yr,fluv_2020_250yr,cstl_2020_250yr,fluv_2020_500yr,cstl_2020_500yr,fluv_2050_5yr,cstl_2050_5yr,fluv_2050_20yr,cstl_2050_20yr,fluv_2050_100yr,cstl_2050_100yr,fluv_2050_250yr,cstl_2050_250yr,fluv_2050_500yr,cstl_2050_500yr\n")
	} else {
		w.WriteString("FD_ID,fluv_2020_5yr,pluv_2020_5yr,fluv_2020_20yr,pluv_2020_20yr,fluv_2020_100yr,pluv_2020_100yr,fluv_2020_250yr,pluv_2020_250yr,fluv_2020_500yr,pluv_2020_500yr,fluv_2050_5yr,pluv_2050_5yr,fluv_2050_20yr,pluv_2050_20yr,fluv_2050_100yr,pluv_2050_100yr,fluv_2050_250yr,pluv_2050_250yr,fluv_2050_500yr,pluv_2050_500yr\n")
	}
	w.Flush()
	size := len(ds.Data)
	count := 0
	records := 5
	if newData {
		records = 5
	}
	for _, r := range ds.Data {
		s := r.Fd_id + ","
		for i := 0; i < records; i++ {
			s += fmt.Sprintf("%f", r.CurrentFluvial.Values[i]) + ","
			s += fmt.Sprintf("%f", r.CurrentPluvial.Values[i]) + ","
		}
		for i := 0; i < records; i++ {
			s += fmt.Sprintf("%f", r.FutureFluvial.Values[i]) + ","
			s += fmt.Sprintf("%f", r.FuturePluvial.Values[i]) + ","
		}
		s = strings.Trim(s, ",")
		if count <= size-1 {
			s += "\n"
		}
		count++
		w.WriteString(s)
		w.Flush()
	}

}
func (h DataSet) Close() {

}
