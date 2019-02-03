package main

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"

	"github.com/globalsign/mgo"
)

func main() {
	var session *mgo.Session

	for {
		var err error

		session, err = mgo.Dial("mongo")
		if err == nil {
			break
		}

		fmt.Errorf("%s\n", err.Error())
	}
	defer session.Close()

	session.SetMode(mgo.PrimaryPreferred, true)
	session.SetSafe(&mgo.Safe{W: 1, J: true, WTimeout: 20000})

	c := session.DB("distributed").C("distributed")

	timer := time.NewTicker(1000 * time.Millisecond)

	for {
		select {
		case <-timer.C:
			var result bson.M

			err := c.Find(bson.M{
				"Name": "Mehdi",
			}).One(&result)

			switch err {
			case mgo.ErrNotFound:
				fmt.Println("<No Data Yet>")

			case nil:
				fmt.Printf("Mehdi's Age is %v\n", result["Age"])

			default:
				panic(err)
			}
		}
	}
}
