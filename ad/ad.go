package ad

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
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

type Msg struct {
	Status, Message string
}

type MsgOk struct {
	Status  string
	Message Ad
}

//********************** POST { **********************
func (ad *Ad) saveAd() error {
	session, err := mgo.Dial("mongodb://admin:12345678@localhost:27017/sa")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("sa").C("ad")
	err = c.Insert(&ad)
	if err != nil {
		panic(err)
	}
	return err
}

func Message(status string, message string) []byte {
	m := Msg{
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

func PostAd(w http.ResponseWriter, r *http.Request) {
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
		w.Write(Message("ERROR", "Could not save your ad, please, try again later"))
		return
	}
	w.Write(Message("OK", "Saved!"))
}

//********************** } POST *********************

//********************** GET { **********************
func getAdById(fid string) (Ad, error) {
	session, err := mgo.Dial("mongodb://admin:12345678@localhost:27017/sa")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	result := Ad{}
	c := session.DB("sa").C("ad")
	err = c.FindId(bson.ObjectIdHex(fid)).One(&result)
	return result, err
}

func GetAd(w http.ResponseWriter, r *http.Request, fid string) {
	result, err := getAdById(fid)
	var m interface{}
	if err != nil {
		m = Msg{
			Status:  "ERROR",
			Message: "Ad not found",
		}
	} else {
		m = MsgOk{
			Status:  "OK",
			Message: result,
		}
	}

	fmt.Println("result = %+v", result)

	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	w.Write(b)
}

//********************** } GET **********************
