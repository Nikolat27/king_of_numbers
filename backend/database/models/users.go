package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserModel struct {
	collection *mongo.Collection
}

type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username  string             `bson:"username" json:"username"`
	Password  string             `bson:"password" json:"password"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
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

func (user *UserModel) Get(filter bson.M) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var userInstance User

	if err := user.collection.FindOne(ctx, filter).Decode(&userInstance); err != nil {
		return nil, err
	}

	return &userInstance, nil
}

func (user *UserModel) GetAll(filter, projection bson.M, page, limit int64) ([]User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	options := &options.FindOptions{}

	options.SetLimit(limit)
	options.SetSkip((page - 1) * limit)
	options.SetProjection(projection)

	cursor, err := user.collection.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}

	var users []User

	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (user *UserModel) DeleteOne(filter bson.M) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return user.collection.DeleteOne(ctx, filter)
}

func (user *UserModel) DeleteAll(filter bson.M) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return user.collection.DeleteMany(ctx, filter)
}

func (user *UserModel) UpdateOne(filter, updates bson.M) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return user.collection.UpdateOne(ctx, filter, updates)
}

func (user *UserModel) UpdateMany(filter, updates bson.M) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return user.collection.UpdateMany(ctx, filter, updates)
}
