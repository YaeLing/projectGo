package userService

import (
	"amenBack/dbService/userDBService"
	"amenBack/model/apiModel"
	"amenBack/model/dbModel"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	accountExist = iota
	createAccountFailed
	createUserInfoFailed
	queryUserInfosFailed
	querySelfAccountFailed
	updateSelfInfoFailed
	updateSelfAccountFailed
	deleteSelfProfileFailed
)

var errorMsgs = map[int]string{
	accountExist:            "Register new user failed. Account already exists.",
	createAccountFailed:     "Register new user failed. Create account failed.",
	createUserInfoFailed:    "Register new user failed. Create user info failed.",
	queryUserInfosFailed:    "Query user infos failed.",
	querySelfAccountFailed:  "Query self account failed.",
	updateSelfInfoFailed:    "Update self info failed.",
	updateSelfAccountFailed: "Update self account failed.",
	deleteSelfProfileFailed: "Delete self profile failed.",
}

func registerNewUser(user apiModel.RequestRegisterUser) error {
	if result, err := userDBService.GetUserAccounts("account", user.Account); err == nil { //if err == nil means account already exists
		if len(result) > 0 {
			errorMsg := errorMsgs[accountExist]
			log.Println(errorMsg)
			return errors.New(errorMsg)
		}
	} else {
		errorMsg := errorMsgs[createAccountFailed]
		log.Println(errorMsg)
		return errors.New(errorMsg)
	}
	id := primitive.NewObjectID()
	userAccount := dbModel.UserAccount{Account: user.Account, Password: user.Password, Role: "user"}
	userInfo := dbModel.UserInfo{Name: user.Name, Phone: user.Phone}
	userProfile := dbModel.UserProfile{ID: id, Info: userInfo, Account: userAccount}

	if err := userDBService.CreateUserProfile(userProfile); err != nil {
		errorMsg := errorMsgs[createAccountFailed]
		log.Println(errorMsg)
		return errors.New(errorMsg)
	}
	return nil
}

func queryUserInfos(key string, value string) (apiModel.ResponseUserInfos, error) {
	var response apiModel.ResponseUserInfos
	if results, err := userDBService.GetUserInfos(key, value); err != nil {
		errorMsg := errorMsgs[queryUserInfosFailed]
		log.Println(errorMsg)
		return response, errors.New(errorMsg)
	} else {
		for _, result := range results {
			responseUserInfo := apiModel.ResponseUserInfo{Name: result.Name, Phone: result.Phone}
			response.UserInfos = append(response.UserInfos, responseUserInfo)
		}
		return response, nil
	}
}

func querySelfAccount(id string) (apiModel.ResponseUserAccount, error) {
	var response apiModel.ResponseUserAccount
	if results, err := userDBService.GetUserAccounts("_id", id); err != nil {
		errorMsg := errorMsgs[querySelfAccountFailed]
		log.Println(errorMsg)
		return response, errors.New(errorMsg)
	} else {
		responseSelfAccount := apiModel.ResponseUserAccount{Account: results[0].Account, Password: results[0].Password, Role: results[0].Role}
		response = responseSelfAccount
		return response, nil
	}
}

func updateSelfInfo(id string, selfInfo apiModel.RequestUpdateUserInfo) error {
	info := dbModel.UserInfo{Name: selfInfo.Name, Phone: selfInfo.Phone}
	if err := userDBService.UpdateUserInfo(id, info); err != nil {
		errorMsg := errorMsgs[updateSelfInfoFailed]
		log.Println(errorMsg)
		return errors.New(errorMsg)
	} else {
		return nil
	}
}

func updateSelfAccount(id string, selfAccount apiModel.RequestUpdateUserAccount) error {
	if result, err := userDBService.GetUserAccounts("_id", id); err != nil {
		errorMsg := errorMsgs[updateSelfAccountFailed]
		log.Println(errorMsg)
		return errors.New(errorMsg)
	} else {
		account := dbModel.UserAccount{Account: selfAccount.Account, Password: selfAccount.Password, Role: result[0].Role}
		if err := userDBService.UpdateUserAccount(id, account); err != nil {
			errorMsg := errorMsgs[updateSelfAccountFailed]
			log.Println(errorMsg)
			return errors.New(errorMsg)
		} else {
			return nil
		}
	}
}

func deleteSelfProfile(id string) error {
	if err := userDBService.DeleteUserProfile(id); err != nil {
		errorMsg := errorMsgs[deleteSelfProfileFailed]
		log.Println(errorMsg)
		return errors.New(errorMsg)
	} else {
		return nil
	}
}
