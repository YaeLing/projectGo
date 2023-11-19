package util

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mgoCli *mongo.Client

func initEngine() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

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
