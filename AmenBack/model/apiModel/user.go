package apiModel

// RequestRegisterUser model info
// @Description User register information
// @Description with acoount and user info
type RequestRegisterUser struct {
	Account  string `bson:"account,required"`  //this is account
	Password string `bson:"password,required"` //this is user password
	Name     string `bson:"name,required"`     //this is user name
	Phone    string `bson:"phone,required"`    //this is user phone
}

// RequestUpdateUserInfo model info
// @Description Update user information
type RequestUpdateUserInfo struct {
	Name  string `bson:"name,required"`  //this is user name
	Phone string `bson:"phone,required"` //this is user phone
}

// RequestUpdateUserAccount model info
// @Description Update user account
type RequestUpdateUserAccount struct {
	Account  string `bson:"account,required"`  //this is account
	Password string `bson:"password,required"` //this is user password
}

type RequestUpdateUserRole struct {
	ID   string `bson:"id,required"`
	Role string `bson:"role,required"`
}

// ResponseUserInfo model info
// @Description Response of user information
type ResponseUserInfo struct {
	Name  string `bson:"name,required"`  //this is user name
	Phone string `bson:"phone,required"` //this is user phone
}

// ResponseUserInfos model info
// @Description Response of multiple user informations
type ResponseUserInfos struct {
	UserInfos []ResponseUserInfo `bson:"UserInfos,required"`
}

// ResponseUserAccount model info
// @Description Response of user account
type ResponseUserAccount struct {
	Account  string `bson:"account,required"`  //this is account
	Password string `bson:"password,required"` //this is user password
	Role     string `bson:"role,required"`     //this is user role
}

// ResponseUserAccount model info
// @Description Response of user account without password
type ResponseUserAccountNoPass struct {
	Account string `bson:"account,required"` //this is account
	Role    string `bson:"role,required"`    //this is user role
}

// ResponseUserProfiles model info
// @Description Response of user profile
type ResponseUserProfile struct {
	ID      string                    `bson:"id,required"`
	Info    ResponseUserInfo          `bson:"Info,required"`
	Account ResponseUserAccountNoPass `bson:"Account,required"`
}

// ResponseUserProfiles model info
// @Description Response of user profiles
type ResponseUserProfiles struct {
	UserProfiles []ResponseUserProfile `bson:"UserProfiles,required"`
}
