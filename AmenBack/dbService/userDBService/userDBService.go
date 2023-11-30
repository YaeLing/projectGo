package userDBService

import (
	utils "amenBack/dbService/utils"
	model "amenBack/model/dbModel"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	client = utils.GetMgoCli()                                 //mongodb client
	coll   = client.Database("Amen").Collection("UserProfile") //user account collection
)

func CreateUserProfile(profile model.UserProfile) error {
	if result, err := coll.InsertOne(context.TODO(), profile); err != nil {
		log.Println(err)
		return err
	} else {
		id := result.InsertedID.(primitive.ObjectID)
		log.Println("New User Account ID", id.Hex())
		return nil
	}
}

func GetUserProfiles(key string, value string) ([]model.UserProfile, error) {
	var (
		err      error
		cursor   *mongo.Cursor
		profiles []model.UserProfile
		filter   bson.M
	)
	if key == "id" {
		ObjectID, err := primitive.ObjectIDFromHex(value)
		if err != nil {
			log.Println("Get user account through id failed")
			return profiles, err
		}
		filter = bson.M{"_id": ObjectID}
	} else {
		filter = bson.M{key: value}
	}
	if cursor, err = coll.Find(context.TODO(), filter); err != nil {
		log.Println(err)
		return profiles, err
	}
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			log.Println(err)
		}
	}()
	if err = cursor.All(context.TODO(), &profiles); err != nil {
		log.Println(err)
		return profiles, err
	}
	return profiles, nil
}

func GetUserAccounts(key string, value string) ([]model.UserAccount, error) {
	var (
		err          error
		cursor       *mongo.Cursor
		results      []model.UserProfile
		userAccounts []model.UserAccount
		filter       bson.M
	)
	if key == "id" {
		ObjectID, err := primitive.ObjectIDFromHex(value)
		if err != nil {
			log.Println("Get user account through id failed")
			return userAccounts, err
		}
		filter = bson.M{"_id": ObjectID}
	} else {
		key = "Account." + key
		filter = bson.M{key: value}
	}
	if cursor, err = coll.Find(context.TODO(), filter); err != nil {
		log.Println(err)
		return userAccounts, err
	}
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			log.Println(err)
		}
	}()
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Println(err)
		return userAccounts, err
	}
	for _, result := range results {
		userAccounts = append(userAccounts, result.Account)
	}
	return userAccounts, nil
}

func GetUserInfos(key string, value string) ([]model.UserInfo, error) {
	var (
		err       error
		cursor    *mongo.Cursor
		results   []model.UserProfile
		userInfos []model.UserInfo
		filter    bson.M
	)
	if key == "id" {
		ObjectID, err := primitive.ObjectIDFromHex(value)
		if err != nil {
			log.Println("Get user account through id failed")
			return userInfos, err
		}
		filter = bson.M{"_id": ObjectID}
	} else {
		key = "Info." + key
		filter = bson.M{key: value}
	}
	if cursor, err = coll.Find(context.TODO(), filter); err != nil {
		log.Println(err)
		return userInfos, err
	}
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			log.Println(err)
		}
	}()
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Println(err)
		return userInfos, err
	}
	for _, result := range results {
		userInfos = append(userInfos, result.Info)
	}
	return userInfos, nil
}

func UpdateUserAccount(id string, account model.UserAccount) error {
	update := bson.D{{"$set", bson.M{"Account": account}}}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = coll.UpdateByID(context.TODO(), objectId, update)
	if err != nil {
		log.Println(err)
	}
	return err
}

func UpdateUserInfo(id string, info model.UserInfo) error {
	update := bson.D{{"$set", bson.M{"Info": info}}}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = coll.UpdateByID(context.TODO(), objectId, update)
	if err != nil {
		log.Println(err)
	}
	return err
}

func DeleteUserProfile(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return err
	}
	filter := bson.M{"_id": objectId}
	_, err = coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println(err)
	}
	return err
}
