package main

import (
	"log/slog"

	webserver "github.com/Nikolat27/king_of_numbers/backend/WebServer"
	"github.com/Nikolat27/king_of_numbers/backend/handlers"
)

func main() {
	handlerInstance := handlers.New()

	webServerInstance := webserver.New("8000", handlerInstance)
	defer webServerInstance.Close()

	if err := webServerInstance.Run(); err != nil {
		slog.Error("WebServer running", "error", err.Error())
	}
}
