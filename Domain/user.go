package Domain

import (
	// "time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" bson:"_id,omitempty"`
	Name           string             `json:"name"`
	UserName       string             `json:"username"`
	Email          string             `json:"email" validate:"required"`
	Password       string             `json:"password,omitempty" validate:"required"`
	Role           string             `json:"role"`
	ProfilePicture string             `json:"profile_picture"`
	EmailVerified  bool               `bson:"email_verified"`
}

// this could have been handled in a better way but i was too lazy to do it
type OmitedUser struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" bson:"_id,omitempty"`
	UserName       string             `json:"username"`
	Email          string             `json:"email" validate:"required"`
	Password       string             `json:"-"`
	Role           string             `json:"role"`
	ProfilePicture string             `json:"profile_picture"`
	EmailVerified  bool               `bson:"email_verified"`
}

type UpdateUser struct {
	Name     string `json:"name"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
