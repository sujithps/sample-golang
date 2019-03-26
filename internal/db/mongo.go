package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"spikes/sample-golang/pkg/logger"
)

type MongoDB struct {
	session *mgo.Session
	User    UserDbClient
}

func NewMongoClient(url, dbName string) *MongoDB {
	logger.NonContext.Info("MongoClient", fmt.Sprintf("Connecting to DB: %s%s", url, dbName), nil)
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	db := session.DB(dbName)
	return &MongoDB{
		User:    NewUser(db),
		session: session,
	}
}

func (client *MongoDB) Close() {
	client.session.Close()
}

func isNotFoundErr(err error) bool {
	return err == mgo.ErrNotFound
}
