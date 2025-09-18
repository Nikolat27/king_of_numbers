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

type GameRoomModel struct {
	collection *mongo.Collection
}

type Player struct {
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username     string             `bson:"username" json:"username"`
	IsBot        bool               `bson:"is_bot" json:"is_bot"`
	Score        int64              `bson:"score" json:"score"`
	ChosenNumber int64              `bson:"chosen_number" json:"chosen_number"`
	JoinedAt     time.Time          `bson:"joined_at" json:"joined_at"`
}

type GameRoom struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	OwnerId   primitive.ObjectID `bson:"owner_id" json:"owner_id"`
	Players   []Player           `bson:"players" json:"players"`
	Round     int64              `bson:"round" bson:"round"`
	Status    string             `bson:"status" json:"status"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

func NewGameRoomModel(db *mongo.Database) *GameRoomModel {
	return &GameRoomModel{
		collection: db.Collection("game_rooms"),
	}
}

func (room *GameRoomModel) Create(ownerId primitive.ObjectID,
	players []Player) (*mongo.InsertOneResult, *types.ErrorResponse) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newRoomInstance := &GameRoom{
		OwnerId:   ownerId,
		Players:   players,
		Round:     1,
		Status:    "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := room.collection.InsertOne(ctx, newRoomInstance)
	if err != nil {
		return nil, types.NewErrorResponse("createRoom", err)
	}

	return result, nil
}

func (room *GameRoomModel) Get(filter, projection bson.M) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	findOptions := options.FindOne()
	findOptions.SetProjection(projection)

	return room.collection.FindOne(ctx, filter, findOptions)
}

func (room *GameRoomModel) Update(filter, updates bson.M) (*mongo.UpdateResult, *types.ErrorResponse) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": updates,
	}

	result, err := room.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, types.NewErrorResponse("updateRoom", err)
	}

	return result, nil
}
