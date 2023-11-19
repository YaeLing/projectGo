package model

type User struct {
	Name  string `bson:"name"`  //使用者名
	Phone string `bson:"phone"` //使用者電話
	Role  string `bson:"role"`  //使用者角色
}
