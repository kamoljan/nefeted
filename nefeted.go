package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kamoljan/nefeted/ad"
	"github.com/kamoljan/nefeted/auth"
	"github.com/kamoljan/nefeted/conf"
)

func adHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("r.Method = %s\n", r.Method)
	fmt.Printf("r.URL = %s\n", r.URL)

	if r.Method == "POST" {
		ad.PostAd(w, r)
	} else if r.Method == "GET" {
		ad.GetAd(w, r, r.URL.Path[4:]) // /ad/52ef6ae7f12eb2aa1635f66b > 52ef6ae7f12eb2aa1635f66b
	} else if r.Method == "DELETE" {
		auth.DeleteAd(w, r)
	} else {
		http.NotFound(w, r)
		return
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("r.Method = %s\n", r.Method)
	fmt.Printf("r.URL = %s\n", r.URL)

	if r.Method == "POST" {
		ad.Search(w, r)
	} else {
		http.NotFound(w, r)
		return
	}
}

func listingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("r.Method = %s\n", r.Method)
	fmt.Printf("r.URL = %s\n", r.URL)

	if r.Method == "GET" {
		ad.Listing(w, r)
	} else {
		http.NotFound(w, r)
		return
	}
}

func myadsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("r.Method = %s\n", r.Method)
	fmt.Printf("r.URL = %s\n", r.URL)

	if r.Method == "GET" {
		ad.Myads(w, r)
	} else {
		http.NotFound(w, r)
		return
	}
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("r.Method = %s\n", r.Method)
	fmt.Printf("r.URL = %s\n", r.URL)

	if r.Method == "PUT" {
		ad.Chat(w, r)
	} else {
		http.NotFound(w, r)
		return
	}
}

func logHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		h.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/ad/", adHandler)
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/listing/", listingHandler)
	http.HandleFunc("/myads/", myadsHandler)
	http.HandleFunc("/chat/", chatHandler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", conf.NefetedPort), logHandler(http.DefaultServeMux))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
