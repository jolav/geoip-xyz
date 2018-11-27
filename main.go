package main

/*
GOOS=linux GOARCH=amd64 go build -o geoip
*/
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	db "./_db"
	lib "./_lib"
)

func init() {
	lib.LoadConfig(configjson, &c)
	if c.App.Mode != "production" {
		c.App.Port = 3000
	}
}

func main() {
	////////////// SEND LOGS TO FILE //////////////////
	if c.App.Mode == "production" {
		var f = c.App.ErrLog
		mylog, err := os.OpenFile(f, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Printf("ERROR opening log file %s\n", err)
		}
		defer mylog.Close() // defer must be in main
		log.SetOutput(mylog)
	}
	///////////////////////////////////////////////////
	db.MYDB.InitDB()

	http.DefaultClient.Timeout = 10 * time.Second
	mux := http.NewServeMux()
	mux.HandleFunc("/", router)

	server := http.Server{
		Addr:           fmt.Sprintf("localhost:%d", c.App.Port),
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("Server listening ...%s", server.Addr)
	server.ListenAndServe()
}

func router(w http.ResponseWriter, r *http.Request) {
	e = myError{}
	params := strings.Split(strings.ToLower(r.URL.Path), "/")
	path := params[1:len(params)]
	if path[len(path)-1] == "" { // remove last empty slot after /
		path = path[:len(path)-1]
	}
	//log.Printf("Going ....%s %s %d", path, r.Method, len(path))
	if len(path) != 2 {
		msg := fmt.Sprintf("%s", "Bad Request")
		badRequest(w, r, msg)
		return
	}

	ip := lib.GetIP(r)
	var origin = r.Header.Get("Origin")
	/*for k, v := range r.Header {
		log.Printf("Header field %q, Value %q\n", k, v)
	}*/
	saveHit(ip, origin, r.URL.RequestURI())

	format := strings.ToLower(path[1])
	if !lib.SliceContainsString(format, c.Geoip.Formats) {
		msg := fmt.Sprintf("Bad Request, unknown format '%s'", path[1])
		badRequest(w, r, msg)
		return
	}
	r.ParseForm()
	target := strings.ToLower(r.Form.Get("q"))
	if target == "" {
		target = lib.GetIP(r)
	}
	doGeoipRequest(w, format, target)

}

func badRequest(w http.ResponseWriter, r *http.Request, msg string) {
	e.Error = msg
	if c.App.Mode != "test" {
		log.Printf("ERROR Bad Request -> %s\n", r.URL.RequestURI())
	}
	lib.SendErrorToClient(w, e)
}
