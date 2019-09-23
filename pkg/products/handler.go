package products

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Lockwarr/POSAR/models"
	"github.com/Lockwarr/POSAR/utils"
)

// Handler handler for luggages
type Handler struct {
	repo models.Products
}

// NewProductsHandler creates new ProductsHandler
func NewProductsHandler(repo models.Products) *Handler {
	return &Handler{repo: repo}
}

//SelectProducts - selects products in our API
func (h *Handler) SelectProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	products := models.Products{}
	decoder := json.NewDecoder(r.Body)
	log.Println(r.Body)
	log.Println(decoder)
	err := decoder.Decode(&products)
	if err != nil {
		utils.ResponseError(w, 500, "Application error: ", err)
		return
	}

	_ = h.repo.Create()

	utils.Response(w, 200, []byte("200"))
}
