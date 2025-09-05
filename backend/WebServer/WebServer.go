package webserver

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Nikolat27/king_of_numbers/backend/handlers"
)

type WebServer struct {
	Server *http.Server
	Port   string
}

// New -> Construction
func New(port string, handler *handlers.Handler) *WebServer {
	var srv = &WebServer{
		Port: port,
	}

	routerInstance := newRouter(handler)

	srv.setupServer(routerInstance)

	return srv
}

func (srv *WebServer) setupServer(router *Router) {
	srv.Server = &http.Server{
		Addr:    fmt.Sprintf(":%s", srv.Port),
		Handler: router.CoreRouter,
	}
}

// Run -> Http
func (srv *WebServer) Run() error {
	log.Println("backend started successfully")
	return srv.Server.ListenAndServe()
}

// RunHttps -> Https
func (srv *WebServer) RunHttps(certFile, keyFile string) error {
	return srv.Server.ListenAndServeTLS(certFile, keyFile)
}

func (srv *WebServer) Close() error {
	if srv.Server == nil {
		return errors.New("'Server' field is nil")
	}

	return srv.Server.Close()
}
