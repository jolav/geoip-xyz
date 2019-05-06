package main

var configjson = []byte(`
{
  "app": {
    "version": "0.1.1",
    "mode": "production",
    "port": 3700,
		"hitslog":"./logs/hits.log",
		"errlog":"./logs/errors.log"
	},
	"geoip": {
		"geoDataBaseFile": "./_data/GeoLite2-City.mmdb",
		"formats": ["json","xml"]
	}
}
`)

var c configuration
var e myError

type configuration struct {
	App struct {
		Mode     string //`json:"mode"`
		Port     int    //`json:"port"`
		Service  string //`json:"service"`
		Version  string //`json:"version"`
		HitsLog  string
		ErrLog   string
		Services []string
	} //`json:"app"`
	Geoip struct {
		GeoDataBaseFile string
		Formats         []string
	}
}

type myError struct {
	Error string `json:"Error,omitempty"`
}

// GEOIP

type geoipOutput struct {
	IP          string  `json:"ip" xml:"ip,omitempty"`
	CountryCode string  `json:"country_code" xml:"country_code,omitempty"`
	CountryName string  `json:"country_name" xml:"country_name,omitempty"`
	RegionCode  string  `json:"region_code" xml:"region_code,omitempty"`
	RegionName  string  `json:"region_name" xml:"region_name,omitempty"`
	City        string  `json:"city" xml:"city,omitempty"`
	ZipCode     string  `json:"zip_code" xml:"zip_code,omitempty"`
	TimeZone    string  `json:"time_zone" xml:"time_zone,omitempty"`
	Latitude    float64 `json:"latitude" xml:"latitude,omitempty"`
	Longitude   float64 `json:"longitude" xml:"longitude,omitempty"`
}
