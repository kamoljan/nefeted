package main

import (
	"fmt"
	"log"
	"net/http"
	// "regexp"

	"github.com/kamoljan/ikura/conf"
	"github.com/kamoljan/nefeted/ad"
)

// var validPath = regexp.MustCompile("^/(ad|search)/([a-zA-Z0-9]+)$")
// var validPath = regexp.MustCompile("^/(ad)/([a-zA-Z0-9]+)?$")

// func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		m := validPath.FindStringSubmatch(r.URL.Path)
// 		if m == nil {
// 			http.NotFound(w, r)
// 			return
// 		}
// 		fn(w, r)
// 	}
// }

// /ad/
func adHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("r.Method = %s\n", r.Method)
	fmt.Printf("r.URL = %s\n", r.URL)

	if r.Method == "POST" {
		ad.PostAd(w, r)
	} else if r.Method == "GET" {
		ad.GetAd(w, r, r.URL.Path[4:]) // /ad/52ef6ae7f12eb2aa1635f66b > 52ef6ae7f12eb2aa1635f66b
	} else {
		http.NotFound(w, r)
		return
	}
}

func logHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		log.Println(req.URL)
		h.ServeHTTP(rw, req)
	})
}

func main() {
	http.HandleFunc("/ad/", adHandler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", conf.IkuradPort), logHandler(http.DefaultServeMux))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
