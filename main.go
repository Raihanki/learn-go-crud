package main

import (
	"crud_app/config"
	"crud_app/routes"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()

	router := mux.NewRouter()
	routes.Router(router)

	server := http.Server{
		Addr:    "localhost:" + config.ENV.APP_PORT,
		Handler: router,
	}

	log.Println("Server running on port", config.ENV.APP_PORT)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
