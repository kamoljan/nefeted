package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/kamoljan/nefeted/ad"
)

var validPath = regexp.MustCompile("^/(fid)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func adHandler(w http.ResponseWriter, r *http.Request, fid string) {
	fmt.Printf("r.Method = %s\n", r.Method)
	fmt.Printf("r.URL = %s\n", r.URL)

	if r.Method == "POST" {
		ad.PostAd(w, r, fid)
	} else if r.Method == "GET" {
		ad.GetAd(w, r, fid)
	} else {
		http.NotFound(w, r)
		return
	}
}

func main() {
	http.HandleFunc("/", makeHandler(adHandler))
	http.ListenAndServe(":8080", nil)
}
