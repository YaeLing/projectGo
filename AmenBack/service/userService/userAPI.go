package userService

import (
	"amenBack/model/apiModel"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register
// @Summary      Register
// @Description  Register user
// @Tags         User
// @Accept		 json
// @Produce      json
// @Param		 Request body apiModel.RequestRegisterUser  true "User info"
// @Success      200  {string}  string   "OK"
// @Failure      400  string  "Register failed"
// @Router       /user/register [POST]
func RegisterAPI(ctx *gin.Context) {
	newUser := apiModel.RequestRegisterUser{}
	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	if err := registerNewUser(newUser); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	} else {
		ctx.String(http.StatusAccepted, "OK")
		return
	}
}

// Query user info api
// @Summary      Query user info
// @Description  Query user informations by key
// @Tags         User
// @Produce      json
// @Param		 key   path  string  true	"Query key"
// @Param		 value   path  string  true	"Query value"
// @Success      200  {object}   apiModel.ResponseUserInfos   "User informations"
// @Failure      404  string  "User informations not found"
// @Router       /user/info/{key}/{value} [GET]
func QueryUserInfosAPI(ctx *gin.Context) {
	key := ctx.Param("key")
	value := ctx.Param("value")
	if infos, err := queryUserInfos(key, value); err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	} else {
		ctx.JSON(http.StatusAccepted, infos)
		return
	}
}

// Query user self account
// @Summary      Query self account
// @Description  Query user self account
// @Tags         User
// @Produce      json
// @Success      200  {object}   apiModel.ResponseUserAccount   "User account"
// @Failure      404  string  "User account not found"
// @Router       /user/account [GET]
func QuerySelfAccountAPI(ctx *gin.Context) {
	if userID := ctx.GetString("userID"); userID == "" {
		ctx.Status(http.StatusInternalServerError)
	} else {
		if account, err := querySelfAccount(userID); err != nil {
			ctx.String(http.StatusNotFound, err.Error())
			return
		} else {
			ctx.JSON(http.StatusAccepted, account)
			return
		}
	}
}

// Update user self information
// @Summary      Update user self info
// @Description  Update user self information
// @Tags         User
// @Accept		 json
// @Produce      json
// @Param		 Request body apiModel.RequestUpdateUserInfo  true "User info"
// @Success      200  string  "OK"
// @Failure      400  string  "Update user self information failed"
// @Failure      500  string  "Internal server error"
// @Router       /user/info [PUT]
func UpdateSelfInfoAPI(ctx *gin.Context) {
	userInfo := apiModel.RequestUpdateUserInfo{}
	if err := ctx.BindJSON(&userInfo); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	if userID := ctx.GetString("userID"); userID == "" {
		ctx.Status(http.StatusInternalServerError)
	} else {
		if err := updateSelfInfo(userID, userInfo); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		} else {
			ctx.JSON(http.StatusAccepted, "OK")
			return
		}
	}
}

// Update user self account
// @Summary      Update user self account
// @Description  Update user self account
// @Tags         User
// @Accept		 json
// @Produce      json
// @Param		 Request body apiModel.RequestUpdateUserAccount  true "User account"
// @Success      200  string  "OK"
// @Failure      400  string  "Update user self account failed"
// @Failure      500  string  "Internal server error"
// @Router       /user/account [PUT]
func UpdateSelfAccountAPI(ctx *gin.Context) {
	userAccount := apiModel.RequestUpdateUserAccount{}
	if err := ctx.BindJSON(&userAccount); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	if userID := ctx.GetString("userID"); userID == "" {
		ctx.Status(http.StatusInternalServerError)
	} else {
		if err := updateSelfAccount(userID, userAccount); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		} else {
			ctx.JSON(http.StatusAccepted, "OK")
			return
		}
	}
}

// Delete user
// @Summary      Delete user
// @Description  Delete user
// @Tags         User
// @Produce      json
// @Success      200  string  "OK"
// @Failure      400  string  "Delete user self profile failed"
// @Failure      500  string  "Internal server error"
// @Router       /user [DELETE]
func DeleteSelfProfile(ctx *gin.Context) {
	if userID := ctx.GetString("userID"); userID == "" {
		ctx.Status(http.StatusInternalServerError)
	} else {
		if err := deleteSelfProfile(userID); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		} else {
			ctx.JSON(http.StatusAccepted, "OK")
			return
		}
	}
}
