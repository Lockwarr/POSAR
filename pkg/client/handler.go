package products

import (
	"encoding/json"
	"net/http"

	"github.com/Lockwarr/POSAR/models"
	"github.com/Lockwarr/POSAR/utils"
)

// Handler handler for luggages
type Handler struct {
	client models.Client
}

// NewClientsHandler creates new ClientsHandler
func NewclientsHandler(client models.Client) *Handler {
	return &Handler{client: client}
}

//SelectClients - selects clients in our API
func (h *Handler) SelectClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&h.client)
	if err != nil {
		utils.ResponseError(w, 500, "Application error: ", err)
		return
	}
	h.client.ID = 0
	h.client.Create()

	utils.Response(w, 200, []byte("200"))
}
