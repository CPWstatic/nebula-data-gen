package generator

import (
	"fmt"
	"math/rand"
	"strconv"
)

func GeneratePoiGeo(fromIdx int64, nodes []Node) ([]Poi, []Geo) {
	pois := make([]Poi, len(nodes))
	geos := make([]Geo, len(nodes))

	for i := range pois {
		lat, _ := strconv.ParseFloat(nodes[i].Lat, 64)
		lng, _ := strconv.ParseFloat(nodes[i].Lon, 64)
		vid := int64(i) + fromIdx
		pois[i] = Poi{
			VID:                   vid,
			Mid:                   fmt.Sprint("mid:", vid),
			DpId:                  rand.Int(),
			DpMainId:              rand.Int(),
			MtId:                  rand.Int(),
			MtMainId:              rand.Int(),
			PoiId:                 rand.Int(),
			PoiMainId:             rand.Int(),
			DpPoiName:             fmt.Sprint("poi_name_", vid),
			DpPoiCityName:         "上海",
			DpPoiBusinessDistrict: "华东",
			DpPoiLatitude:         lat,
			DpPoiLongitude:        lng,
			DpPoiStatus:           fmt.Sprint("poi_stat_", vid),
			DpPoiFirstCateName:    fmt.Sprint("poi_fc_", vid),
			DpPoiSecondCateName:   fmt.Sprint("poi_sc_", vid),
			DpPoiScore:            0,
			DpPoiStar:             fmt.Sprint("poi_star_", vid),
		}

		geos[i] = Geo {
			Lat: lat,
			Lng: lng,
			Dst: int64(i),
		}
	}

	return pois, geos
}

func GenerateUsers(fromIdx, cnt int64) []User {
	users := make([]User, cnt)
	for i := range users {
		vid := int64(i) + fromIdx
		users[i] = User{
			VID:                    vid,
			Mid:                    fmt.Sprint("mid:", vid),
			MtUserId:               rand.Int(),
			MtUserName:             fmt.Sprint("mun:", vid),
			DpUserId:               rand.Int(),
			DpUserName:             fmt.Sprint("dun:", vid),
			IsEmployee:             fmt.Sprint("empe:", vid),
			Mobile:                 fmt.Sprint("mobile:", vid),
			Gender:                 fmt.Sprint("gender:", vid),
			Age:                    fmt.Sprint("age:", vid),
			Residence:              fmt.Sprint("rs:", vid),
			MtUserFirstPayDate:     fmt.Sprint("fpd:", vid),
			MtUserXmdCampaignMoney: rand.Float64(),
			MtUserPoiCampaignMoney: rand.Float64(),
			DpUserFirstPayDate:     fmt.Sprint("dpfpd:", vid),
			DpUserXmdCampaignMoney: rand.Float64(),
			DpUserPoiCampaignMoney: rand.Float64(),
		}
	}

	return users
}

func GenerateConsumePoi(users []User, pois []Poi) []ConsumePoi {
	var consumes []ConsumePoi

	for _, u :=range users {
		numConsume := rand.Intn(10);
		for i := 0; i < numConsume; i++ {
			pId := rand.Intn(len(pois))
			consume := ConsumePoi{
				Src:       u.VID,
				Dst:       int64(pId),
				Ranking:   0,
				SourceMid: u.Mid,
				TargetMid: pois[pId].Mid,
				Weight:    rand.Float64(),
			}

			consumes = append(consumes, consume)
		}
	}

	return consumes
}
