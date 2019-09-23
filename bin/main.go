package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Lockwarr/POSAR/models"
	"github.com/Lockwarr/POSAR/pkg/luggage"
	"github.com/Lockwarr/POSAR/pkg/products"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	var test models.FirstLuggage
	var test1 models.Products
	luggageHandler := luggage.NewLuggageHandler(test)
	productsHandler := products.NewProductsHandler(test1)
	r.Route("/v1", func(r chi.Router) {
		r.Post("/products/select", productsHandler.SelectProducts)
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
