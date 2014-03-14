package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kamoljan/nefeted/ad"
	"github.com/kamoljan/nefeted/conf"
)

const (
	Mongodb     = "mongodb://admin:12345678@localhost:27017/sa"
	NefetedPort = 8080
)

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

func searchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("r.Method = %s\n", r.Method)
	fmt.Printf("r.URL = %s\n", r.URL)

	if r.Method == "GET" {
		ad.GetSearch(w, r)
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
	http.HandleFunc("/search/", searchHandler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", conf.NefetedPort), logHandler(http.DefaultServeMux))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
