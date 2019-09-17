package main

import (
	"fmt"
	"net/http"
	"os"

	"./controllers"
	"github.com/gorilla/mux"
	"github.com/goji/httpauth"
	"log"
)

func main() {
	router := mux.NewRouter()

	router.Use(httpauth.SimpleBasicAuth("dave", "somepassword"))
	router.HandleFunc("/endpoint", controllers.Authenticate).Methods("GET")

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
