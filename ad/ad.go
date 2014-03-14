package ad

import (
	// "fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"github.com/kamoljan/ikura/api"
	"github.com/kamoljan/ikura/json"
	"github.com/kamoljan/nefeted/conf"
)

type Ad struct {
	Id          bson.ObjectId `json:"id"                bson:"_id"`
	Profile     uint64        `json:"profile"           bson:"profile"`
	Title       string        `json:"title"             bson:"title"`
	Category    uint64        `json:"category"          bson:"category"`
	Description string        `json:"description"       bson:"description"`
	Price       uint64        `json:"price"             bson:"price"`
	Currency    string        `json:"currency"          bson:"currency"`
	Report      uint64        `json:"report,omitempty"  bson:"report,omitempty"`
	Date        time.Time     `json:"date"              bson:"date"`
	Image1      json.Egg      `json:"image1"            bson:"image1"`
	Image2      json.Egg      `json:"image2,omitempty"  bson:"image2,omitempty"`
	Image3      json.Egg      `json:"image3,omitempty"  bson:"image2,omitempty"`
}

//********************** POST { **********************
func (ad *Ad) saveAd() error {
	session, err := mgo.Dial(conf.Mongodb)
	if err != nil {
		log.Fatal("Unable to connect to DB ", err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true) // Optional. Switch the session to a monotonic behavior.
	c := session.DB("sa").C("ad")
	err = c.Insert(&ad)
	if err != nil {
		log.Fatal("Unable to save to DB ", err)
	}
	return err
}

// newborn1, newborn2, newborn3 ...
func PostAd(w http.ResponseWriter, r *http.Request) {
	// TODO: refactor it!
	profile, err := strconv.ParseUint(r.FormValue("profile"), 10, 64)
	if err != nil {
		w.Write(json.Message("ERROR", "Profile is missing"))
		return
	}
	log.Println(r.FormValue("category"))
	category, err := strconv.ParseUint(r.FormValue("category"), 10, 64)
	if err != nil {
		w.Write(json.Message("ERROR", "Category is missing"))
		return
	}
	price, err := strconv.ParseUint(r.FormValue("price"), 10, 64)
	if err != nil {
		w.Write(json.Message("ERROR", "Price is missing"))
		return
	}
	title := r.FormValue("title")
	if title == "" {
		w.Write(json.Message("ERROR", "Title is missing"))
		return
	}
	description := r.FormValue("description")
	if description == "" {
		w.Write(json.Message("ERROR", "Description is missing"))
		return
	}
	currency := r.FormValue("currency")
	if currency == "" {
		w.Write(json.Message("ERROR", "Currency is missing"))
		return
	}
	ad := Ad{
		Profile:     profile,
		Title:       title,
		Category:    category,
		Description: description,
		Price:       price,
		Currency:    currency,
		Report:      0,
		Date:        time.Now(),
	}
	newborn1 := r.FormValue("newborn1") // Newborn image1
	if newborn1 == "" {
		w.Write(json.Message("ERROR", "At least one image should be uploaded"))
	} else {
		image1, err := api.GetEggBySize("newborn", newborn1)
		if err == nil {
			ad.Image1 = image1
		}
	}
	newborn2 := r.FormValue("newborn2") // Newborn image2
	if newborn2 != "" {
		image2, err := api.GetEggBySize("newborn", newborn2)
		if err == nil {
			ad.Image2 = image2
		}
	}
	newborn3 := r.FormValue("newborn3") // Newborn image3
	if newborn3 != "" {
		image3, err := api.GetEggBySize("newborn", newborn3)
		if err == nil {
			ad.Image3 = image3
		}
	}
	err = ad.saveAd()
	if err != nil {
		w.Write(json.Message("ERROR", "Could not save your ad, please, try again later"))
		return
	}
	w.Write(json.Message("OK", "Saved!"))
}

//********************** GET { **********************
func getAdById(id string) (Ad, error) {
	session, err := mgo.Dial(conf.Mongodb)
	if err != nil {
		log.Fatal("Unable to connect to DB ", err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true) // Optional. Switch the session to a monotonic behavior.
	result := Ad{}
	c := session.DB("sa").C("ad")
	err = c.FindId(bson.ObjectIdHex(id)).One(&result)
	return result, err
}

/*
size=infant(default)
{
	status: "OK",
	result: {
	    profile: 123412341134123,
	    title: "test",
	    category: 323,
	    description: "dasfasdfas asdfadsf adsfadfadsfadsf qwerqwerqwer adfasdfdf",
	    price: 1241234123,
	    currency: "qwerqwer",
	    report: 0,
	    date: "2014-02-03T18:09:43.309+08:00"
   	    image1: [
	        "newborn" : "0001_040db0bc2fc49ab41fd81294c7d195c7d1de358b_ACA0AC_100_160"
	        "infant" : "0001_ff41e42b0134e219bc09eddda87687822460afcf_ACA0AC_200_319"
	        "baby" : "0001_6881db255b21c864c9d1e28db50dc3b71dab5b78_ACA0AC_400_637"
	    ],
   	    image2: [
        	"newborn" : "0001_040db0bc2fc49ab41fd81294c7d195c7d1de358b_ACA0AC_100_160"
	        "infant" : "0001_ff41e42b0134e219bc09eddda87687822460afcf_ACA0AC_200_319"
	        "baby" : "0001_6881db255b21c864c9d1e28db50dc3b71dab5b78_ACA0AC_400_637"
	    ],
   	    image3: [],
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

/*
{
	status: "OK",
	result: {
	    language: "english",
	    ok: 1,
	    queryDebugString: "hsjsjd||||||",
	    results: [
	        {
            	obj: {
                	_id: "53202376058424b87c9d9368",
                	image1: {
                	    baby: "0001_68415c85528ccf9e763eb48d9dc0fca8a540f701_655B4C_400_300",
                	    egg: "0001_c0448ef44bf7fd00476fadf11805d94fe94a5820_655B4C_816_612",
                	    infant: "0001_9b0e36f28be91dae81a02863fadce2bc2f196312_655B4C_200_150",
                	    newborn: "0001_72a53f664db6f415e9e862c607d9c0ba177c20af_655B4C_100_75"
                    },
            	    price: 6468
                },
            	score: 1.1
            },
	        ...
        ],
        stats: {
        	n: 2,
        	nfound: 2,
        	nscanned: 3,
        	nscannedObjects: 0,
        	timeMicros: 69
        }
	}
}
*/
func GetSearch(w http.ResponseWriter, r *http.Request) {
	session, err := mgo.Dial(conf.Mongodb)
	if err != nil {
		log.Fatal("Unable to connect to DB ", err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true) // Optional. Switch the session to a monotonic behavior.
	db := session.DB("sa")
	var result interface{}
	//db.ad.runCommand("text", { search: "hsjsjd", limit: 20, project: { "price" : 1, "image1": 1 }})
	// q := bson.D{{"text", "ad"}, {"search", "hsjsjd"}} // working one
	q := bson.D{
		{"text", "ad"},
		{"search", "hsjsjd"},
		{"limit", 2},
		{"project",
			bson.D{
				{"price", 1},
				{"image1", 1},
			},
		},
	}
	err = db.Run(q, &result)
	log.Printf("err = %s\n", err)
	if err != nil {
		w.Write(json.Message("ERROR", "Ads not found"))
	} else {
		w.Write(json.Message("OK", result))
	}
}
