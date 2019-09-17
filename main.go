package main

import (
	"fmt"
	"net/http"
	"os"

	"./controllers"
	"github.com/gorilla/mux"
	"github.com/goji/httpauth"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/test", controllers.Test).Methods("GET")
	http.Handle("/", httpauth.SimpleBasicAuth("test", "test")(router))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println("http://localhost:" + port)

	err := http.ListenAndServe(":"+port, router) 
	if err != nil {
		fmt.Print(err)
	}
}
