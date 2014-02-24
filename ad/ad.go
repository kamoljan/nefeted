package ad

import (
	"net/http"
	"strconv"
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"github.com/kamoljan/nefeted/json"
)

type Ad struct {
	Profile     uint64 // Facebook profile ID
	Image       string
	Thumb       string
	Title       string
	Category    uint64
	Description string
	Price       uint64
	Currency    string
	Report      uint64
	Date        time.Time
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

//********************** POST { **********************
func (ad *Ad) saveAd() error {
	session, err := mgo.Dial("mongodb://admin:12345678@localhost:27017/sa")
	check(err)
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("sa").C("ad")
	err = c.Insert(&ad)
	check(err)
	return err
}

func PostAd(w http.ResponseWriter, r *http.Request) {
	// TODO: refactor it!
	profile, err := strconv.ParseUint(r.FormValue("profile"), 10, 64)
	if err != nil {
		w.Write(json.Message("Error", "Profile is missing"))
		return
	}
	category, err := strconv.ParseUint(r.FormValue("category"), 10, 64)
	if err != nil {
		w.Write(json.Message("Error", "Category is missing"))
		return
	}
	price, err := strconv.ParseUint(r.FormValue("price"), 10, 64)
	if err != nil {
		w.Write(json.Message("Error", "Price is missing"))
		return
	}
	title := r.FormValue("title")
	if title == "" {
		w.Write(json.Message("Error", "Title is missing"))
		return
	}
	image := r.FormValue("image")
	if image == "" {
		w.Write(json.Message("Error", "Image is missing"))
		return
	}
	thumb := r.FormValue("thumb")
	if thumb == "" {
		w.Write(json.Message("Error", "Thumb is missing"))
		return
	}
	description := r.FormValue("description")
	if description == "" {
		w.Write(json.Message("Error", "Description is missing"))
		return
	}
	currency := r.FormValue("currency")
	if currency == "" {
		w.Write(json.Message("Error", "Currency is missing"))
		return
	}
	ad := Ad{
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
	err = ad.saveAd()
	if err != nil {
		w.Write(json.Message("ERROR", "Could not save your ad, please, try again later"))
		return
	}
	w.Write(json.Message("OK", "Saved!"))
}

//********************** } POST *********************

//********************** GET { **********************
func getAdById(id string) (Ad, error) {
	session, err := mgo.Dial("mongodb://admin:12345678@localhost:27017/sa")
	check(err)
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	result := Ad{}
	c := session.DB("sa").C("ad")
	err = c.FindId(bson.ObjectIdHex(id)).One(&result)
	return result, err
}

// TODO: read it http://stackoverflow.com/questions/17998943/golang-library-package-that-returns-json-string-from-http-request
/*
{
	Status: "OK",
	Message: {
	Profile: 123412341134123,
	Image: "12412341234",
	Thumb: "asdfasfasdfaf",
	Title: "test",
	Category: 323,
	Description: "dasfasdfas asdfadsf adsfadfadsfadsf qwerqwerqwer adfasdfdf",
	Price: 1241234123,
	Currency: "qwerqwer",
	Report: 0,
	Date: "2014-02-03T18:09:43.309+08:00"
	}
}
*/
func GetAd(w http.ResponseWriter, r *http.Request, id string) {
	result, err := getAdById(id)
	if err != nil {
		w.Write(json.Message("ERROR", "Ad not found"))
	} else {
		w.Write(json.Message("OK", result))
	}
}

//********************** } GET **********************
