package main

import (
	//這邊引入的地方要看module名字 不同package就分資料夾放

	CRUD "mongo/CRUD"
	util "mongo/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	var (
		client     = util.GetMgoCli()
		Collection *mongo.Collection
	)
	Collection = client.Database("user").Collection("user")
	CRUD.Delete(Collection)
	// CRUD.Create(Collection)
	CRUD.Read(Collection)
	// CRUD.Update(Collection)
}
