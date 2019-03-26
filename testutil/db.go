package testutil

import (
	"gopkg.in/mgo.v2"
	"spikes/sample-golang/pkg/config"
)

func CleanDb() {
	session, err := mgo.Dial(config.MongoURL())
	defer session.Close()

	if err != nil {
		panic(err)
	}

	db := session.DB(config.MongoDBName())
	_ = db.C("users").DropCollection()
}

