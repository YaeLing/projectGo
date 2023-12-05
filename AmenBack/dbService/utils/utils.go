package dbUtils

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mgoCli *mongo.Client

func getMongoURI() string {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return "mongodb://root:3345678@localhost:27017"
	} else {
		return uri
	}
}

func initEngine() {
	var err error
	mongoURI := getMongoURI()
	clientOptions := options.Client().ApplyURI(mongoURI)

	// 連接mongoDB
	mgoCli, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 檢查連線
	err = mgoCli.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func GetMgoCli() *mongo.Client {
	if mgoCli == nil {
		initEngine()
	}
	return mgoCli
}
