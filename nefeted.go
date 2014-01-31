package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/kamoljan/nefeted/ad"
)

func Message(status string, message string) []byte {
	type Message struct {
		Status  string
		Message string
	}
	m := Message{
		Status:  status,
		Message: message,
	}
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error:", err)
		panic(err)
	}
	return b
}

func adHandler(w http.ResponseWriter, r *http.Request, fid string) {
	fmt.Printf("r.Method = %s\n", r.Method)
	fmt.Printf("r.URL = %s\n", r.URL)

	if r.Method == "POST" {
		// TODO: refactor it!
		profile, err := strconv.ParseUint(r.FormValue("profile"), 10, 64)
		if err != nil {
			w.Write(Message("Error", "Profile is missing"))
			return
		}
		category, err := strconv.ParseUint(r.FormValue("category"), 10, 64)
		if err != nil {
			w.Write(Message("Error", "Category is missing"))
			return
		}
		price, err := strconv.ParseUint(r.FormValue("price"), 10, 64)
		if err != nil {
			w.Write(Message("Error", "Price is missing"))
			return
		}
		title := r.FormValue("title")
		if title == "" {
			w.Write(Message("Error", "Title is missing"))
			return
		}
		image := r.FormValue("image")
		if image == "" {
			w.Write(Message("Error", "Image is missing"))
			return
		}
		thumb := r.FormValue("thumb")
		if thumb == "" {
			w.Write(Message("Error", "Thumb is missing"))
			return
		}
		description := r.FormValue("description")
		if description == "" {
			w.Write(Message("Error", "Description is missing"))
			return
		}
		currency := r.FormValue("currency")
		if currency == "" {
			w.Write(Message("Error", "Currency is missing"))
			return
		}
		ad := &ad.Ad{
			Profile:     profile,
			Image:       image,
			Thumb:       thumb,
			Title:       title,
			Category:    category,
			Description: description,
			Price:       price,
			Currency:    currency,
			Report:      0,
			Date:        time.Now(),
		}
		err = ad.Save()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(Message("OK", "Saved!"))
	} else {
		http.NotFound(w, r)
		return
	}
}

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

func main() {
	http.HandleFunc("/", makeHandler(adHandler))
	http.ListenAndServe(":8080", nil)
}
