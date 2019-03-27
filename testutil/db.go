package testutil

import (
	"git.thoughtworks.net/mahadeva/sample-golang/pkg/config"
	"gopkg.in/mgo.v2"
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
