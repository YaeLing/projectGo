package adminService

import (
	"amenBack/dbService/userDBService"
	"amenBack/model/apiModel"
	"amenBack/model/authModel"
	"errors"
	"log"
)

const (
	queryUserProfilesFailed = iota
	userNotExist
	roleNotAccepted
	updateUserRoleFailed
	deleteUserProfileFailed
)

var errorMsgs = map[int]string{
	queryUserProfilesFailed: "Admin query user profiles failed.",
	userNotExist:            "Admin query user not exists.",
	roleNotAccepted:         "This role not accepted.",
	updateUserRoleFailed:    "Admin update user profile failed.",
	deleteUserProfileFailed: "Admin delete user profile failed.",
}

func queryUserProfiles(key string, value string) (apiModel.ResponseUserProfiles, error) {
	var response = apiModel.ResponseUserProfiles{}
	if userProfiles, err := userDBService.GetUserProfiles(key, value); err != nil {
		errorMsg := errorMsgs[queryUserProfilesFailed]
		log.Println(errorMsg)
		return response, errors.New(errorMsg)
	} else {
		for index, userProfile := range userProfiles {
			response.UserProfiles[index].ID = userProfile.ID.Hex()
			response.UserProfiles[index].Info = apiModel.ResponseUserInfo(userProfile.Info)
			response.UserProfiles[index].Account = apiModel.ResponseUserAccountNoPass{Account: userProfile.Account.Account, Role: userProfile.Account.Role}
		}
		return response, nil
	}
}

func updateUserRole(id string, role string) error {
	if role != authModel.RoleAdmin && role != authModel.RoleUser {
		errorMsg := errorMsgs[roleNotAccepted]
		log.Println(errorMsg)
		return errors.New(errorMsg)
	}
	if userProfiles, err := userDBService.GetUserProfiles("_id", id); err != nil {
		errorMsg := errorMsgs[queryUserProfilesFailed]
		log.Println(errorMsg)
		return errors.New(errorMsg)
	} else if len(userProfiles) == 0 {
		errorMsg := errorMsgs[userNotExist]
		log.Println(errorMsg)
		return errors.New(errorMsg)
	} else {
		userProfile := userProfiles[0]
		userProfile.Account.Role = role
		if err := userDBService.UpdateUserAccount(id, userProfile.Account); err != nil {
			errorMsg := errorMsgs[updateUserRoleFailed]
			log.Println(errorMsg)
			return errors.New(errorMsg)
		} else {
			return nil
		}
	}
}

func deleteUserProfile(id string) error {
	if err := userDBService.DeleteUserProfile(id); err != nil {
		errorMsg := errorMsgs[deleteUserProfileFailed]
		log.Println(errorMsg)
		return errors.New(errorMsg)
	} else {
		return nil
	}
}
