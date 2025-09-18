package models

import "go.mongodb.org/mongo-driver/mongo"

type Models struct {
	User     *UserModel
	GameRoom *GameRoomModel
}

func New(db *mongo.Database) *Models {
	return &Models{
		User:     NewUserModel(db),
		GameRoom: NewGameRoomModel(db),
	}
}
