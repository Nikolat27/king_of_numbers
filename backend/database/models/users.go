package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserModel struct {
	collection *mongo.Collection
}

type User struct {
	Id        primitive.ObjectID
	Username  string
	Password  string
	CreatedAt time.Time
}

func NewUserModel(db *mongo.Database) *UserModel {
	return &UserModel{
		collection: db.Collection("users"),
	}
}
