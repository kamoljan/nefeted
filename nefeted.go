package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/kamoljan/nefeted/ad"
)

// FIXME: ad/save POST, ad/objectid GET
var validPath = regexp.MustCompile("^/(ad|search)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2]) //FIXME: ad/save POST m[2] is not needed
	}
}

func adHandler(w http.ResponseWriter, r *http.Request, fid string) {
	fmt.Printf("r.Method = %s\n", r.Method)
	fmt.Printf("r.URL = %s\n", r.URL)

	if r.Method == "POST" {
		ad.PostAd(w, r)
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
