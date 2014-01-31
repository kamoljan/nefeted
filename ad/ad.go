package ad

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
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

func (a *Ad) Save() error {
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
	fmt.Println("Description lah:", result.Description)

	return nil //XXX tmp
}
