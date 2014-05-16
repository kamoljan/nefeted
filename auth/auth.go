package auth

import (
	// "fmt"
	"log"
	"net/http"
	"strings"

	fb "github.com/huandu/facebook"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"github.com/kamoljan/ikura/json"
	"github.com/kamoljan/nefeted/conf"
)

/*
DELETE: http://localhost:8080/ad/{id}/{token}
{
	status: "OK"
    "data":"Deleted"
}
*/
func DeleteAd(w http.ResponseWriter, r *http.Request) {
	var ad, token string
	s := strings.Split(r.URL.Path, "/")
	if len(s) >= 4 {
		ad, token = s[2], s[3]
	} else {
		w.Write(json.Message3("ERROR", nil, "Wrong URL"))
		return
	}

	// create a global App var to hold your app id and secret.
	var globalApp = fb.New("1454662381420609", "9a9d5fd67edc5245f366a1a0fb9bb9bf")

	// facebook asks for a valid redirect uri when parsing signed request.
	// it's a new enforced policy starting in late 2013.
	// it can be omitted in a mobile app server.
	//globalApp.RedirectUri = ""

	// here comes a client with a facebook signed request string in query string.
	// creates a new session with signed request.
	//session, _ := globalApp.SessionFromSignedRequest(signedRequest)

	// or, you just get a valid access token in other way.
	// creates a session directly.
	session := globalApp.Session(token)

	// use session to send api request with your access token.
	//res, _ := session.Get("/me/feed", nil)

	// validate access token. err is nil if token is valid.
	err := session.Validate()

	if err == nil {
		session, err := mgo.Dial(conf.Mongodb)
		if err != nil {
			log.Fatal("Unable to connect to DB ", err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true) // Optional. Switch the session to a monotonic behavior.
		db := session.DB("sa")

		err = db.C("ad").Remove(bson.M{"_id": bson.ObjectIdHex(ad)})
		if err != nil {
			w.Write(json.Message("ERROR", "Could not DELETE"))
			log.Printf("err = %s\n", err)
		} else {
			w.Write(json.Message("OK", "Deleted"))
		}
	} else {
		w.Write(json.Message("ERROR", "Not valid session"))
	}
}
