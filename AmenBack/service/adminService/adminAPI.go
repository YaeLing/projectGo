package adminService

import (
	"amenBack/model/apiModel"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Query user profiles
// @Summary      Admin query user profiles
// @Description  Admin query user profiles
// @Tags         Admin
// @Produce      json
// @Param		 key   path  string  true	"Query key"
// @Param		 value   path  string  true	"Query value"
// @Success      200  string  "OK"
// @Failure      400  string  "Query user profile failed"
// @Router       /admin/profile/{key}/{value} [GET]
func QueryUserProfilesAPI(ctx *gin.Context) {
	key := ctx.Param("key")
	value := ctx.Param("value")
	if result, err := queryUserProfiles(key, value); err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	} else {
		ctx.JSON(http.StatusAccepted, result)
		return
	}
}

// Update user role
// @Summary      Admin update user role
// @Description  Admin update user role
// @Tags         Admin
// @Accept		 json
// @Produce      json
// @Param		 Request body apiModel.RequestUpdateUserRole  true "User role"
// @Success      200  string  "OK"
// @Failure      400  string  "Update user role failed"
// @Router       /admin/role [PUT]
func UpdateUserRoleAPI(ctx *gin.Context) {
	request := apiModel.RequestUpdateUserRole{}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	if err := updateUserRole(request.ID, request.Role); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	} else {
		ctx.String(http.StatusAccepted, "OK")
		return
	}
}

// Delete user
// @Summary      Admin delete user
// @Description  Admin delete user
// @Tags         Admin
// @Produce      json
// @Success      200  string  "OK"
// @Failure      400  string  "Delete user profile failed"
// @Router       /profile/{userID} [DELETE]
func DeleteUserProfileAPI(ctx *gin.Context) {
	id := ctx.Param("userID")
	if err := deleteUserProfile(id); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	} else {
		ctx.String(http.StatusAccepted, "OK")
		return
	}
}
