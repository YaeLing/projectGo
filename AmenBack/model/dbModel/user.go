package dbModel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserInfo struct {
	Name  string `bson:"name,required"`
	Phone string `bson:"phone,required"`
}

type UserAccount struct {
	Account  string `bson:"account,required"`
	Password string `bson:"password,required"`
	Role     string `bson:"role,required"`
}

type UserProfile struct {
	ID      primitive.ObjectID `bson:"_id,required"`
	Info    UserInfo           `bson:"Info,required"`
	Account UserAccount        `bson:"Account,required"`
}
