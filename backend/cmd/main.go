package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	webserver "github.com/Nikolat27/king_of_numbers/backend/WebServer"
	"github.com/Nikolat27/king_of_numbers/backend/database"
	"github.com/Nikolat27/king_of_numbers/backend/database/models"
	"github.com/Nikolat27/king_of_numbers/backend/handlers"
	"github.com/Nikolat27/king_of_numbers/backend/paseto"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("reading .env file: %s", err.Error()))
	}

	mongoURI, err := getMongoURI()
	if err != nil {
		panic(fmt.Errorf("getting mongoURI %s", err.Error()))
	}

	db, err := database.New(mongoURI)
	if err != nil {
		panic(fmt.Errorf("creating new database: %s ", err.Error()))
	}

	allModels := models.New(db)

	pasetoInstance, errResp := paseto.New()
	if errResp != nil {
		panic(fmt.Errorf("creating paseto: %s ", errResp.Detail))
	}

	handlerInstance := handlers.New(allModels, pasetoInstance)

	webServerInstance := webserver.New("8000", handlerInstance)
	defer webServerInstance.Close()

	if err := webServerInstance.Run(); err != nil {
		slog.Error("WebServer running", "error", err.Error())
	}
}

func getMongoURI() (string, error) {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		return "", errors.New("MONGO_URI env variable does not exist")
	}

	return uri, nil
}
