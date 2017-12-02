package mongodb

import (
	"gopkg.in/mgo.v2"

	"github.com/AidHamza/gomvc-demo/models/logger"
	"github.com/AidHamza/gomvc-demo/setting"
)

// DB .
var DB *mgo.Database

func init() {
	logger.Log("mongodb#init")

	session, err := mgo.Dial(setting.Option.MongoURL)
	if err != nil {
		logger.Log("ERR mgo.Dial:", err)
		panic(err)
	}

	logger.Log("mongodb#init: success")
	DB = session.DB(setting.Option.MongoDBName)
}
