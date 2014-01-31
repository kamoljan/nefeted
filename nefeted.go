package main

import (
	//"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Ad struct {
	Profile     uint64 // Facebook profile ID
	Title       string
	Image       string
	Thumb       string
	Category    uint64
	Description string
	Price       uint64
	Phone       string
	Date        time.Time
	Currency    string
	Report      uint64
}

func (a *Ad) save() error {
	// filename := "data/" + p.Title + ".txt"
	// return ioutil.WriteFile(filename, p.Body, 0600)
	session, err := mgo.Dial("mongodb://admin:12345678@localhost:27017/sa")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("sa").C("ad")
	err = c.Insert(&a)
	if err != nil {
		panic(err)
	}

	result := Ad{}
	err = c.Find(bson.M{"title": "test"}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("Phone:", result.Phone)

	return nil //XXX tmp
}

func adHandler(w http.ResponseWriter, r *http.Request, fid string) {
	if r.Method == "POST" {
		fmt.Printf("r.Method = %s\n", r.Method)
		fmt.Printf("r.URL = %s\n", r.URL)
		fmt.Printf("r.profile = %s\n", r.FormValue("profile"))

		profile, err := strconv.ParseUint(r.FormValue("profile"), 10, 64)
		if err != nil {
			// panic(err)
		}
		category, err := strconv.ParseUint(r.FormValue("category"), 10, 64)
		if err != nil {
			//panic(err)
		}

		price, err := strconv.ParseUint(r.FormValue("price"), 10, 64)
		if err != nil {
			//panic(err)
		}

		ad := &Ad{
			Profile:     profile,
			Title:       r.FormValue("title"),
			Image:       r.FormValue("image"),
			Thumb:       r.FormValue("thumb"),
			Category:    category,
			Description: r.FormValue("description"),
			Price:       price,
			Phone:       r.FormValue("phone"),
			Date:        time.Now(),
			Currency:    r.FormValue("currency"),
			Report:      0,
		}

		fmt.Printf("ad = %s", ad)
		err = ad.save()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("This is GET request " + r.Method))
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
