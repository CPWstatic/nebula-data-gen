package generator

type Poi struct {
	VID                   int64
	Mid                   string
	DpId                  int
	DpMainId              int
	MtId                  int
	MtMainId              int
	PoiId                 int
	PoiMainId             int
	DpPoiName             string
	DpPoiCityName         string
	DpPoiBusinessDistrict string
	DpPoiLatitude         float64
	DpPoiLongitude        float64
	DpPoiStatus           string
	DpPoiFirstCateName    string
	DpPoiSecondCateName   string
	DpPoiScore            int
	DpPoiStar             string
}

type Geo struct {
	Lat float64
	Lng float64
	Dst int64
}

type User struct {
	VID					   int64
	Mid                    string
	MtUserId               int
	MtUserName             string
	DpUserId               int
	DpUserName             string
	IsEmployee             string
	Mobile                 string
	Gender                 string
	Age                    string
	Residence              string
	MtUserFirstPayDate     string
	MtUserXmdCampaignMoney float64
	MtUserPoiCampaignMoney float64
	DpUserFirstPayDate     string
	DpUserXmdCampaignMoney float64
	DpUserPoiCampaignMoney float64
}

type ConsumePoi struct {
	Src       int64
	Dst       int64
	Ranking   int64
	SourceMid string
	TargetMid string
	Weight    float64
}
