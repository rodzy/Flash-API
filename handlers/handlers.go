package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rodzy/flash/middlewares"
	"github.com/rodzy/flash/routers"
	"github.com/rs/cors"
)

//DirectDrivers func allows to connect API w/ routing using mux
func DirectDrivers() {
	//Instancing a new router w mux
	router := mux.NewRouter()
	//Register
	router.HandleFunc("/register", middlewares.CheckDataBase(routers.Register)).Methods("POST")

	//Setting port env var
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	//Setting up the cors for the cloud API
	handler := cors.AllowAll().Handler(router)
	//Just in case for http error listening
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
