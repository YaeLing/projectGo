package CRUD

import (
	"context"
	"fmt"
	"log"
	"mongo/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	iResult  *mongo.InsertOneResult
	iResults *mongo.InsertManyResult
	id       primitive.ObjectID
	err      error
)

func Create(collection *mongo.Collection) {
	//insert once
	user := model.User{Name: "芝山小學生", Phone: "4484824", Role: "user"}

	if iResult, err = collection.InsertOne(context.TODO(), user); err != nil {
		fmt.Println(err)
		return
	}
	//_id:默认生成一个全局唯一ID
	id = iResult.InsertedID.(primitive.ObjectID)
	fmt.Println("自增ID", id.Hex())
	//insert many
	users := []interface{}{model.User{Name: "張學友", Phone: "3345678", Role: "user"}, model.User{Name: "郝龍冰", Phone: "8874874", Role: "user"}}
	if iResults, err = collection.InsertMany(context.TODO(), users); err != nil {
		fmt.Println(err)
		return
	}
	if iResults == nil {
		log.Fatal("result nil")
	}
	for _, v := range iResults.InsertedIDs {
		id = v.(primitive.ObjectID)
		fmt.Println("自增ID", id.Hex())
	}
}

/*
bson.D{}: 對Document的有序描述，key-value以逗號分隔；        bson.D{{"foo", "bar"}, {"hello", "world"}, {"pi", 3.14159}}
bson.M{}: Map 結構，key-value 以冒號分隔，無順序，使用最方便；bson.M{"foo": "bar", "hello": "world", "pi": 3.14159}
bson.A{}: 有排序的bson array                                bson.A{"bar", "world", 3.14159, bson.D{{"qux", 12345}}}
*/

func Read(collection *mongo.Collection) {
	var (
		err    error
		cursor *mongo.Cursor
	)
	var result model.User
	filter := bson.M{"name": "董翰瑋"}
	if err = collection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	filter = bson.M{"name": "郝龍冰"}
	if cursor, err = collection.Find(context.TODO(), filter); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	var results []model.User
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}
}

func Update(collection *mongo.Collection) {
	id, _ := primitive.ObjectIDFromHex("655a0a514e93d4842681d24d")
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"name", "董家安"}}}}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Documents updated: %v\n", result.ModifiedCount)

	filter = bson.D{{"name", "郝龍冰"}}
	update = bson.D{{"$set", bson.D{{"phone", "55688"}}}}
	result, err = collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Documents updated: %v\n", result.ModifiedCount)
}

func Delete(collection *mongo.Collection) {
	filter := bson.M{"name": "董家安"}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Documents deleted: %v\n", result.DeletedCount)

	filter = bson.M{"name": "郝龍冰"}
	result, err = collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Documents deleted: %v\n", result.DeletedCount)
}
