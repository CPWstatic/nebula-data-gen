package generator

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"
	"reflect"
	"strconv"
)

func ExportPoiToCsv(file string, pois []Poi)  {
	ifaces := make([]interface{}, len(pois))
	for i := range pois {
		ifaces[i] = pois[i]
	}

	exportToCSVFile(file, ifaces)
}

func ExportGeoCsv(file string, geos []Geo)  {
	ifaces := make([]interface{}, len(geos))
	for i := range geos {
		ifaces[i] = geos[i]
	}

	exportToCSVFile(file, ifaces)
}

func ExportUserCsv(file string, users []User)  {
	ifaces := make([]interface{}, len(users))
	for i := range users {
		ifaces[i] = users[i]
	}

	exportToCSVFile(file, ifaces)
}

func ExportConsumeCsv(file string, consumes []ConsumePoi)  {
	ifaces := make([]interface{}, len(consumes))
	for i := range consumes {
		ifaces[i] = consumes[i]
	}

	exportToCSVFile(file, ifaces)
}

func exportToCSVFile(filename string, ifaces []interface{}) {
	if err := os.MkdirAll(path.Dir(filename), 0755); err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	for _, iface := range ifaces {
		writer.Write(Record(iface))
	}

	writer.Flush()
}

func Record(t interface{}) []string {
	numFields := reflect.ValueOf(t).NumField()
	record := make([]string, numFields)
	for i := range record {
		f := reflect.ValueOf(t).Field(i)
		if f.Type().Name() == "string" {
			record[i] = fmt.Sprintf("\"%s\"", f.Interface())
		} else if f.Type().Name() == "float" {
			record[i] = strconv.FormatFloat(f.Float(), 'f', 8, 64)
		} else {
			record[i] = fmt.Sprintf("%v", f.Interface())
		}
	}
	return record
}