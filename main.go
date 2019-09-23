package main

import (
	"fmt"
	"net/http"
	"os"

	"./models"
	"./pkg/luggage"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	var test models.FirstLuggage
	luggageHandler := luggage.NewLuggageHandler(test)
	r.Route("/v1", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Route("/luggage", func(r chi.Router) {
				r.Post("/quote", luggageHandler.GlobalMinimalQuotation)
				r.Get("/countries", luggageHandler.GetCountries)
				r.Get("/itemtypes", luggageHandler.GetItemTypes)
			})
		})
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println("http://localhost:" + port)

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		fmt.Print(err)
	}
}
