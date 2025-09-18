package models

import (
	"context"
	"time"

	"github.com/Nikolat27/king_of_numbers/backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserModel struct {
	collection *mongo.Collection
}

type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username  string             `bson:"username" json:"username"`
	Password  string             `bson:"password" json:"password"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

func NewUserModel(db *mongo.Database) *UserModel {
	return &UserModel{
		collection: db.Collection("users"),
	}
}

func (user *UserModel) Create(username, password string) (primitive.ObjectID, *types.ErrorResponse) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var newUser = &User{
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
	}

	result, err := user.collection.InsertOne(ctx, newUser)
	if err != nil {
		return primitive.NilObjectID, types.NewErrorResponse("create user", err)
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (user *UserModel) Get(filter, projection bson.M) (*User, *types.ErrorResponse) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result User

	findOptions := options.FindOne()
	if projection != nil {
		findOptions.SetProjection(projection)
	}

	if err := user.collection.FindOne(ctx, filter, findOptions).Decode(&result); err != nil {
		return nil, types.NewErrorResponse("getUser", err)
	}

	return &result, nil
}
