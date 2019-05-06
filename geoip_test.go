package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGeoipApi(t *testing.T) {
	c.App.Mode = "test"

	for _, test := range geoipTests {
		var to = geoipTestOutput{}
		pass := true
		//fmt.Println(`Test...`, test.endpoint, "...", test.it)
		req, err := http.NewRequest("GET", test.endpoint, nil)
		if err != nil {
			//fmt.Println(`------------------------------`)
			//fmt.Println(err.Error())
			//fmt.Println(test.errorText)
			//fmt.Println(`------------------------------`)
			if err.Error() != test.errorText {
				t.Fatalf("Error Request %s\n", err.Error())
			} else {
				pass = false
			}
		}
		if pass {
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(router)
			handler.ServeHTTP(rr, req)
			if rr.Code != test.statusCode {
				t.Errorf("%s got %v want %v\n", test.endpoint, rr.Code, test.statusCode)
			}

			class := strings.Split(test.it, " ")[0]

			if class == "JSON" {
				_ = json.Unmarshal(rr.Body.Bytes(), &to)
				if to.Error != test.errorText {
					t.Errorf("%s got %v want %v\n", test.endpoint, to.Error, test.errorText)
				}
				var expected geoipOutput
				_ = json.Unmarshal([]byte(test.expected), &expected)
				if to.geoipOutput != expected {
					var fail string
					fail = fmt.Sprintf("\n")
					fail += fmt.Sprintf("---------- got ---------- \n")
					fail += fmt.Sprintf("%v \n", to.geoipOutput)
					fail += fmt.Sprintf("---------- want ---------- \n")
					fail += fmt.Sprintf("%v \n", expected)
					fail += fmt.Sprintf("------------------------- \n")
					t.Error(fail)
				}
			}

			if class == "XML" {
				if string(rr.Body.Bytes()) != test.expected {
					var fail string
					fail = fmt.Sprintf("\n")
					fail += fmt.Sprintf("---------- got ---------- \n")
					fail += fmt.Sprintf("%s \n", string(rr.Body.Bytes()))
					fail += fmt.Sprintf("---------- want ---------- \n")
					fail += fmt.Sprintf("%s \n", test.expected)
					fail += fmt.Sprintf("------------------------- \n")
					t.Error(fail)
				}
			}
		}
		fmt.Printf("Test OK...%s\n", test.endpoint)
	}
}

type geoipTestOutput struct {
	StatusCode int    // `json:"statusCode"`
	Error      string `json:"Error,omitempty"`
	geoipOutput
}

var geoipTests = []struct {
	it         string
	endpoint   string
	errorText  string
	statusCode int
	expected   string
}{
	{
		"JSON /v1/json no parameters",
		"/v1/json",
		" is a unknown host, not a valid IP or hostname",
		400,
		``,
	},
	{
		"JSON valid IP /v1/json?q=208.67.222.222",
		"/v1/json?q=208.67.222.222",
		"",
		200,
		`{
			"ip": "208.67.222.222",
			"country_code": "US",
			"country_name": "United States",
			"region_code": "",
			"region_name": "",
			"city": "",
			"zip_code": "",
			"time_zone": "America/Chicago",
			"latitude": 37.751,
			"longitude": -97.822
		 }`,
	},
	{
		"JSON valid IP /v1/json?q=8.8.8.8",
		"/v1/json?q=8.8.8.8",
		"",
		200,
		`{
			"ip": "8.8.8.8",
			"country_code": "US",
			"country_name": "United States",
			"region_code": "",
			"region_name": "",
			"city": "",
			"zip_code": "",
			"time_zone": "America/Chicago",
			"latitude": 37.751,
			"longitude": -97.822
		 }`,
	},
	{
		"JSON valid IPV6 /v1/json?q=2a00:1450:4006:803::200e",
		"/v1/json?q=2a00:1450:4006:803::200e",
		"",
		200,
		`{
			"ip": "2a00:1450:4006:803::200e",
			"country_code": "IE",
			"country_name": "Ireland",
			"region_code": "",
			"region_name": "",
			"city": "",
			"zip_code": "",
			"time_zone": "Europe/Dublin",
			"latitude": 53,
			"longitude": -8
		 }`,
	},
	{
		"JSON invalid IP /v1/json?q=260.50.50.50",
		"/v1/json?q=260.50.50.50",
		"260.50.50.50 is a unknown host, not a valid IP or hostname",
		400,
		``,
	},
	{
		"JSON invalid IP /v1/json?q=20.20.-5.20",
		"/v1/json?q=20.20.-5.20",
		"20.20.-5.20 is a unknown host, not a valid IP or hostname",
		400,
		``,
	},
	{
		"JSON valid Hostname /v1/json?q=riojandogs.com",
		"/v1/json?q=riojandogs.com",
		"",
		200,
		`{
			"ip": "89.38.144.25",
			"country_code": "GB",
			"country_name": "United Kingdom",
			"region_code": "ENG",
			"region_name": "England",
			"city": "Slough",
			"zip_code": "SL1",
			"time_zone": "Europe/London",
			"latitude": 51.5353,
			"longitude": -0.6657
		 }`,
	},
	{
		"JSON non existent Hostname /v1/json?q=code123456tabs.com",
		"/v1/json?q=code123456tabs.com",
		"code123456tabs.com is a unknown host, not a valid IP or hostname",
		400,
		``,
	},
	{
		"XML /v1/xml no parameters",
		"/v1/xml",
		"",
		400,
		`{
 "Error": " is a unknown host, not a valid IP or hostname"
}`,
	},
	{
		"XML valid IP /v1/xml?q=208.67.222.222",
		"/v1/xml?q=208.67.222.222",
		"",
		200,
		`<geoipOutput><ip>208.67.222.222</ip><country_code>US</country_code><country_name>United States</country_name><time_zone>America/Chicago</time_zone><latitude>37.751</latitude><longitude>-97.822</longitude></geoipOutput>`,
	},
	{
		"XML valid IP /v1/xml?q=8.8.8.8",
		"/v1/xml?q=8.8.8.8",
		``,
		200,
		"<geoipOutput><ip>8.8.8.8</ip><country_code>US</country_code><country_name>United States</country_name><time_zone>America/Chicago</time_zone><latitude>37.751</latitude><longitude>-97.822</longitude></geoipOutput>",
	},
	{
		"XML valid IPV6 /v1/xml?q=2a00:1450:4006:803::200e",
		"/v1/xml?q=2a00:1450:4006:803::200e",
		``,
		200,
		"<geoipOutput><ip>2a00:1450:4006:803::200e</ip><country_code>IE</country_code><country_name>Ireland</country_name><time_zone>Europe/Dublin</time_zone><latitude>53</latitude><longitude>-8</longitude></geoipOutput>",
	},
	{
		"XML invalid IP /v1/xml?q=260.50.50.50",
		"/v1/xml?q=260.50.50.50",
		"260.50.50.50 is a unknown host, not a valid IP or hostname",
		400,
		`{
 "Error": "260.50.50.50 is a unknown host, not a valid IP or hostname"
}`,
	},
	{
		"XML invalid IP /v1/xml?q=20.20.-5.20",
		"/v1/xml?q=20.20.-5.20",
		"20.20.-5.20 is a unknown host, not a valid IP or hostname",
		400,
		`{
 "Error": "20.20.-5.20 is a unknown host, not a valid IP or hostname"
}`,
	},
	{
		"XML valid Hostname /v1/xml?q=riojandogs.com",
		"/v1/xml?q=riojandogs.com",
		"",
		200,
		"<geoipOutput><ip>89.38.144.25</ip><country_code>GB</country_code><country_name>United Kingdom</country_name><region_code>ENG</region_code><region_name>England</region_name><city>Slough</city><zip_code>SL1</zip_code><time_zone>Europe/London</time_zone><latitude>51.5353</latitude><longitude>-0.6657</longitude></geoipOutput>",
	},
	{
		"XML non existent Hostname /v1/xml?q=code123456tabs.com",
		"/v1/xml?q=code123456tabs.com",
		"",
		400,
		`{
 "Error": "code123456tabs.com is a unknown host, not a valid IP or hostname"
}`,
	},
}
