package models

import (
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID            primitive.ObjectID   `json:"_id" bson:"_id"`
	Username      string               `json:"username" validate:"required,min=5,max=20"`
	AccountStatus string               `json:"accountStatus" bson:"accountStatus" validate:"required"`
	Email         string               `json:"email" validate:"required"`
	Address       string               `json:"address" validate:"required"`
	PhoneNumber   string               `json:"phoneNumber" bson:"phoneNumber" validate:"required"`
	RoleID        int                  `json:"roleId" bson:"roleId" validate:"required"`
	Password      string               `json:"password" validate:"required"`
	TmpPassword   string               `json:"tmpPassword" bson:"tmpPassword"`
	CreatedBy     primitive.ObjectID   `json:"createdBy"  bson:"createdBy" validate:"required"`
	CreatedAt     time.Time            `json:"createdAt" bson:"createdAt"`
	UpdatedAt     time.Time            `json:"updatedAt" bson:"updatedAt"`
	Permissions   []Permission         `json:"permissions" bson:"permissions"`
	Shops         []primitive.ObjectID `json:"shops" bson:"shops"`
	Token         string               `json:"token" bson:"token"`
}

type Permission struct {
	ID   any      `json:"id" bson:"id"` //any to allow backward compatibility with existing data but it should be either Int | string
	Name string   `json:"name" bson:"name"`
	List []string `json:"list" bson:"list"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserToken struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	Email  string             `json:"email"`
	RoleID int                `json:"roleId" bson:"roleId"`
	jwt.RegisteredClaims
}
