package ad

import (
	// "fmt"
	"labix.org/v2/mgo"
	// "labix.org/v2/mgo/bson"
	//"time"
)

func Save() error {
	session, err := mgo.Dial("mongodb://admin:12345678@localhost:27017/sa")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	// c := session.DB("sa").C("ad")
	// err := c.Insert(&Ad{r.})
	return nil //XXX tmp
}
