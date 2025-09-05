package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
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

func (user *UserModel) Create(username, password string) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var newUser = &User{
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
	}

	result, err := user.collection.InsertOne(ctx, newUser)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}
