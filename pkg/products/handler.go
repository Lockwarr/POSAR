package products

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Lockwarr/POSAR/models"
	"github.com/Lockwarr/POSAR/utils"
	"github.com/go-chi/chi"
)

// Handler handler for luggages
type Handler struct {
	products models.Products
}

// NewProductsHandler creates new ProductsHandler
func NewProductsHandler(products models.Products) *Handler {
	return &Handler{products: products}
}

//SelectProducts - selects products in our API
func (h *Handler) SelectProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&h.products)
	if err != nil {
		utils.ResponseError(w, 500, "Application error: ", err)
		return
	}
	h.products.ID = 0
	h.products.Update()

	utils.Response(w, 200, []byte("200"))
}

//SelectedProducts - shows selected products for id in our API
func (h *Handler) SelectedProducts(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		utils.ResponseError(w, 500, "Application error: ", err)
		return
	}
	products := models.GetProducts(uint(id))
	if err != nil {
		log.Println(err, "get")
		utils.ResponseError(w, 500, "Application error: ", err)
		return
	}
	b, err := json.Marshal(products)
	if err != nil {
		utils.ResponseError(w, 500, "Application error: ", err)
		return
	}
	utils.Response(w, 200, b)
}
