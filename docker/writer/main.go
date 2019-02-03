package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/globalsign/mgo/bson"

	"github.com/globalsign/mgo"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

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
			_, err := c.Find(bson.M{
				"Name": "Mehdi",
			}).Apply(mgo.Change{
				Upsert: true,
				Update: bson.M{
					"Name": "Mehdi",
					"Age":  20 + rand.Intn(10),
				},
			}, nil)

			if err != nil {
				panic(err)
			}
		}
	}
}
