package hazard_providers

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/USACE/go-consequences/geography"
	"github.com/USACE/go-consequences/hazardproviders"
	"github.com/USACE/go-consequences/hazards"
	_ "github.com/mattn/go-sqlite3"
)

//SQLDataSet is a simple struct to store a sql dataset
type SQLDataSet struct {
	db *sql.DB
}

//OpenSQLDepthDataSet opens a sqldataset
func OpenSQLDepthDataSet(depthFilePath string) SQLDataSet {
	// have to change working directory one level
	// os.Chdir("..")
	// path, err := os.Getwd()
	// fmt.Println(path, err)

	// open both databases
	//dbNsi, _ := sql.Open("sqlite3", nsiFilePath)
	dbDepth, _ := sql.Open("sqlite3", depthFilePath)
	// _, err = dbDepth.Exec("attach database '" + nsiFilePath + "'as nsi;")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// results, err := dbDepth.Query("SELECT a.*, b.fd_id, b.x, b.y FROM fathom_depths as a LEFT JOIN nsi as b on a.fd_id=b.fd_id")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//sds.db.Query("SELECT a.*, b.fd_id, b.x, b.y FROM fathom_depths as a LEFT JOIN ")
	return SQLDataSet{db: dbDepth}
}

//ProvideHazardBoundary fulfils the HazardProvider interface from go-consequences
func (sds SQLDataSet) ProvideHazardBoundary() (geography.BBox, error) {
	fmt.Println("not set up for BBox yet")
	return geography.BBox{Bbox: nil}, nil
}

//ProvideHazard fulfils the HazardProvider interface from go-consequences
func (sds SQLDataSet) ProvideHazard(args interface{}) (hazards.HazardEvent, error) {
	fd_id, ok := args.(FathomQuery)

	if ok {
		r, found := sds.getRecord(fd_id.Location) //make a query for a row...
		if fd_id.NewData {
			return nil, hazardproviders.HazardError{Input: "Sqlite data does not support new data specification."}
		}
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
					return generateDepthEvent(fd_id.Frequency, r.CurrentPluvial, fd_id.NewData)
				} else if fd_id.Year == 2050 {
					return generateDepthEvent(fd_id.Frequency, r.FuturePluvial, fd_id.NewData)
				} else {
					//throw error?
					return nil, hazardproviders.HazardError{Input: "Bad Year Argument"}
				}
			}
		} else {
			return nil, hazardproviders.NoHazardFoundError{Input: fd_id.Location.SRID}
		}
	} else {
		return nil, hazardproviders.HazardError{Input: "Unable to parse args to hazard_providers.FathomQuery"}
	}
}

var depthFilePath = "/workspaces/go-fathom/data/fathom-depths.db"
var nsiFilePath = "/workspaces/go-fathom/data/nsiv2_29.gpkg"

func (sds SQLDataSet) getRecord(loc geography.Location) (Record, bool) {
	//db2, _ := sql.Open("sqlite3", nsiFilePath)
	s := "ATTACH DATABASE '/workspaces/go-fathom/data/fathom-depths.db' as depths;"
	statement, err := sds.db.Prepare(s)
	statement.Exec()

	if err != nil {
		log.Fatal(err)
	}
	//query := "SELECT a.fd_id, b.x, b.y, a.fluv_2020_5yr, a.pluv_2020_5yr, a.fluv_2020_20yr, a.pluv_2020_20yr, a.fluv_2020_100yr, a.pluv_2020_100yr, a.fluv_2020_250yr, a.pluv_2020_250yr, a.fluv_2020_500yr, a.pluv_2020_500yr, a.fluv_2050_5yr, a.pluv_2050_5yr, a.fluv_2050_20yr, a.pluv_2050_20yr, a.fluv_2050_100yr, a.pluv_2050_100yr, a.fluv_2050_250yr, a.pluv_2050_250yr, a.fluv_2050_500yr, a.pluv_2050_500yr FROM depths.fathom_depths as a LEFT JOIN nsi as b ON a.fd_id=b.fd_id WHERE x ='" + loc.X + "'and y ='" + loc.Y + "';"
	row, err := sds.db.Query(fmt.Sprint("SELECT a.fd_id, b.x, b.y, a.fluv_2020_5yr, a.pluv_2020_5yr, a.fluv_2020_20yr, a.pluv_2020_20yr, a.fluv_2020_100yr, a.pluv_2020_100yr, a.fluv_2020_250yr, a.pluv_2020_250yr, a.fluv_2020_500yr, a.pluv_2020_500yr, a.fluv_2050_5yr, a.pluv_2050_5yr, a.fluv_2050_20yr, a.pluv_2050_20yr, a.fluv_2050_100yr, a.pluv_2050_100yr, a.fluv_2050_250yr, a.pluv_2050_250yr, a.fluv_2050_500yr, a.pluv_2050_500yr FROM depths.fathom_depths as a LEFT JOIN nsi as b ON a.fd_id=b.fd_id WHERE x=", loc.X, " and y=", loc.Y, ";"))
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var fid string
		var x float64
		var y float64
		var fluv_2020_5yr float64
		var pluv_2020_5yr float64
		var fluv_2020_20yr float64
		var pluv_2020_20yr float64
		var fluv_2020_100yr float64
		var pluv_2020_100yr float64
		var fluv_2020_250yr float64
		var pluv_2020_250yr float64
		var fluv_2020_500yr float64
		var pluv_2020_500yr float64
		var fluv_2050_5yr float64
		var pluv_2050_5yr float64
		var fluv_2050_20yr float64
		var pluv_2050_20yr float64
		var fluv_2050_100yr float64
		var pluv_2050_100yr float64
		var fluv_2050_250yr float64
		var pluv_2050_250yr float64
		var fluv_2050_500yr float64
		var pluv_2050_500yr float64
		row.Scan(&fid, &x, &y, &fluv_2020_5yr, &pluv_2020_5yr, &fluv_2020_20yr, &pluv_2020_20yr, &fluv_2020_100yr, &pluv_2020_100yr, &fluv_2020_250yr, &pluv_2020_250yr, &fluv_2020_500yr, &pluv_2020_500yr, &fluv_2050_5yr, &pluv_2050_5yr, &fluv_2050_20yr, &pluv_2050_20yr, &fluv_2050_100yr, &pluv_2050_100yr, &fluv_2050_250yr, &pluv_2050_250yr, &fluv_2050_500yr, &pluv_2050_500yr)
		location := geography.Location{X: x, Y: y, SRID: fid}
		cfvals := []float64{fluv_2020_5yr, fluv_2020_20yr, fluv_2020_100yr, fluv_2020_250yr, fluv_2020_500yr}
		cpvals := []float64{pluv_2020_5yr, pluv_2020_20yr, pluv_2020_100yr, pluv_2020_250yr, pluv_2020_500yr}
		ffvals := []float64{fluv_2050_5yr, fluv_2050_20yr, fluv_2050_100yr, fluv_2050_250yr, fluv_2050_500yr}
		fpvals := []float64{pluv_2050_5yr, pluv_2050_20yr, pluv_2050_100yr, pluv_2050_250yr, pluv_2050_500yr}
		futurefluvial := FrequencyData{fluvial: true, year: 2050, Values: ffvals}
		futurepluvial := FrequencyData{fluvial: false, year: 2050, Values: fpvals}
		currentfluvial := FrequencyData{fluvial: true, year: 2020, Values: cfvals}
		currentpluvial := FrequencyData{fluvial: false, year: 2020, Values: cpvals}
		if hasNonZeroValues(ffvals, fpvals, cfvals, cpvals, false) {
			r := Record{Location: location, FutureFluvial: futurefluvial, FuturePluvial: futurepluvial, CurrentFluvial: currentfluvial, CurrentPluvial: currentpluvial}
			return r, true
		}
	}

	return Record{}, false
}
func (ds DataSet) WriteToSqlite() {
	db := CreateDepthDatabase()
	index := 0
	maxTransaction := 500
	transaction := make([]interface{}, 500)
	for _, val := range ds.Data {
		transaction[index] = val
		index++
		if index >= maxTransaction {
			WriteArrayToDepthDatabase(db, transaction)
			index = 0
		}
	}
	fmt.Println("Done.")
}
func CreateDepthDatabase() *sql.DB {
	os.Remove("fathom-depths.db")
	fmt.Println("Creating fathom-depths.db...")
	file, err := os.Create("fathom-depths.db")
	if err != nil {
		fmt.Println("error")
	}
	file.Close()
	fmt.Println("fathom-depths.db created")

	db, _ := sql.Open("sqlite3", "./fathom-depths.db")
	//defer db.Close()
	createDepthTable(db)
	return db
}
func CreateDepthWALDatabase() *sql.DB {
	os.Remove("fathom-depths.db")
	fmt.Println("Creating fathom-depths.db...")
	file, err := os.Create("fathom-depths.db")
	if err != nil {
		fmt.Println("error")
	}
	file.Close()
	fmt.Println("fathom-depths.db created")

	db, _ := sql.Open("sqlite3", "./fathom-depths.db?_journal_mode=WAL") //https://stackoverflow.com/questions/35804884/sqlite-concurrent-writing-performance/35805826
	db.SetMaxOpenConns(1)
	//defer db.Close()
	createDepthTable(db)
	return db
}
func createDepthTable(db *sql.DB) {
	createfathom := `CREATE TABLE fathom_depths (
		"fd_id" string NOT NULL PRIMARY KEY,
		"x" float,
		"y" float,
		"fluv_2020_5yr" float,
		"pluv_2020_5yr" float,
		"fluv_2020_20yr" float,
		"pluv_2020_20yr" float,
		"fluv_2020_100yr" float,
		"pluv_2020_100yr" float,
		"fluv_2020_250yr" float,
		"pluv_2020_250yr" float,
		"fluv_2020_500yr" float,
		"pluv_2020_500yr" float,
		"fluv_2050_5yr" float,
		"pluv_2050_5yr" float,
		"fluv_2050_20yr" float,
		"pluv_2050_20yr" float,
		"fluv_2050_100yr" float,
		"pluv_2050_100yr" float,
		"fluv_2050_250yr" float,
		"pluv_2050_250yr" float,
		"fluv_2050_500yr" float,
		"pluv_2050_500yr" float
	  );`

	statement, err := db.Prepare(createfathom) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("fathom_depths table created")
}
func WriteArrayToDepthDatabase(db *sql.DB, dset []interface{}) {
	insertresult := `INSERT INTO fathom_depths(fd_id, x, y, fluv_2020_5yr, pluv_2020_5yr, fluv_2020_20yr, pluv_2020_20yr, fluv_2020_100yr, pluv_2020_100yr, fluv_2020_250yr, pluv_2020_250yr, fluv_2020_500yr, pluv_2020_500yr, fluv_2050_5yr, pluv_2050_5yr, fluv_2050_20yr, pluv_2050_20yr, fluv_2050_100yr, pluv_2050_100yr, fluv_2050_250yr, pluv_2050_250yr, fluv_2050_500yr, pluv_2050_500yr) VALUES `
	var inserts []string
	var params []interface{}
	somethingtoadd := false
	for _, data := range dset {
		record, ok := data.(Record)
		if ok {
			somethingtoadd = true
			inserts = append(inserts, "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
			params = append(params, record.Location, record.CurrentFluvial.Values[0], record.CurrentPluvial.Values[0], record.CurrentFluvial.Values[1], record.CurrentPluvial.Values[1], record.CurrentFluvial.Values[2], record.CurrentPluvial.Values[2], record.CurrentFluvial.Values[3], record.CurrentPluvial.Values[3], record.CurrentFluvial.Values[4], record.CurrentPluvial.Values[4], record.FutureFluvial.Values[0], record.FuturePluvial.Values[0], record.FutureFluvial.Values[1], record.FuturePluvial.Values[1], record.FutureFluvial.Values[2], record.FuturePluvial.Values[2], record.FutureFluvial.Values[3], record.FuturePluvial.Values[3], record.FutureFluvial.Values[4], record.FuturePluvial.Values[4])
		}

	}
	if somethingtoadd {
		queryVals := strings.Join(inserts, ",")
		insertresult += queryVals
		statement, err := db.Prepare(insertresult)
		if err != nil {
			fmt.Println(insertresult)
			log.Fatalln("ERROR WITH DB PREPARE " + err.Error())
		}
		_, err = statement.Exec(params...)
		if err != nil {
			fmt.Println(params)
			log.Fatalln("ERROR WITH EXECUTION " + err.Error())
		}
	}

}
