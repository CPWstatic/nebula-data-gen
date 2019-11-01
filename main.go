package main

import (
	gen "./generator"
	"sync"
)

var (
	poiCnt int64 = 100
	userCnt int64 = 1000
)

func main()  {
	var pois []gen.Poi
	var geos []gen.Geo
	var users []gen.User
	var consumes []gen.ConsumePoi


	/* gen vertex*/
	var genWg sync.WaitGroup

	genWg.Add(1)
	go func() {
		defer genWg.Done()
		if nodes, err := gen.ReadNode(poiCnt); err == nil {
			pois,geos = gen.GeneratePoiGeo(0, nodes)
		}
	}()

	genWg.Add(1)
	go func() {
		defer genWg.Done()
		users = gen.GenerateUsers(poiCnt, userCnt)
	}()

	genWg.Wait()

	/*gen edges*/
	consumes = gen.GenerateConsumePoi(users, pois)

	/*export*/
	var expWg sync.WaitGroup

	expWg.Add(1)
	go func() {
		defer expWg.Done()
		gen.ExportPoiToCsv("./poi.csv", pois)
	}()

	expWg.Add(1)
	go func() {
		defer expWg.Done()
		gen.ExportGeoCsv("./geo.csv", geos)
	}()

	expWg.Add(1)
	go func() {
		defer expWg.Done()
		gen.ExportUserCsv("./user.csv", users)
	}()

	expWg.Add(1)
	go func() {
		defer expWg.Done()
		gen.ExportConsumeCsv("./consume.csv", consumes)
	}()

	expWg.Wait()
}

