package database

import (
	"context"
	"errors"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(uri string) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbClient, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	dbName, err := getDatabaseName()
	if err != nil {
		return nil, err
	}

	database := dbClient.Database(dbName)

	return database, nil
}

func getDatabaseName() (string, error) {
	dbName := os.Getenv("DATABASE_NAME")
	if dbName == "" {
		return "", errors.New("DATABASE_NAME env variable does not exist")
	}

	return dbName, nil
}
