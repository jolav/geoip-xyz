package lib

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"math"
	"net"
	"net/http"
	"strings"
)

// SendJSONToClient ...
func SendJSONToClient(w http.ResponseWriter, d interface{}) {
	w.Header().Set("Content-Type", "application/json")
	var dataJSON = []byte(`{}`)
	dataJSON, err := json.MarshalIndent(d, "", " ")
	if err != nil {
		log.Printf("ERROR Marshaling %s\n", err)
		w.Write([]byte(`{}`))
	}
	w.Write(dataJSON)
}

// SendXMLToClient ...
func SendXMLToClient(w http.ResponseWriter, d interface{}) {
	w.Header().Set("Content-Type", "application/xml")
	var dataXML = []byte(`<output></output>`)
	dataXML, err := xml.Marshal(&d)
	if err != nil {
		log.Printf("ERROR Parsing into XML %s\n", err)
		w.Write([]byte(`{}`))
	}
	w.Write(dataXML)
}

// SendErrorToClient ...
func SendErrorToClient(w http.ResponseWriter, d interface{}) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	var dataJSON = []byte(`{}`)
	dataJSON, err := json.MarshalIndent(d, "", " ")
	if err != nil {
		log.Printf("ERROR Marshaling %s\n", err)
		w.Write([]byte(`{}`))
	}
	w.Write(dataJSON)
}

// ToFixedFloat64 (untruncated, num) -> untruncated.toFixed(num)
func ToFixedFloat64(untruncated float64, precision int) float64 {
	coef := math.Pow10(precision)
	truncated := float64(int(untruncated*coef)) / coef
	return truncated
}

// LoadConfig ...
func LoadConfig(configjson []byte, c interface{}) {
	err := json.Unmarshal(configjson, &c)
	if err != nil {
		log.Printf("ERROR LoadConfig %s\n", err)
	}
}

// GetIP ...
func GetIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	ip = strings.Split(ip, ",")[0]
	if len(ip) > 0 {
		return ip
	}
	ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	ip = strings.Split(ip, ",")[0]
	return ip
}

// SliceContainsString ... returns true/false
func SliceContainsString(str string, slice []string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
