package luggage

import (
	"io/ioutil"
	"net/http"

	"github.com/Lockwarr/POSAR/models"
	"github.com/Lockwarr/POSAR/utils"
)

// Handler handler for luggages
type Handler struct {
	repo models.FirstLuggage
}

// NewLuggageHandler creates new LuggageHandler
func NewLuggageHandler(repo models.FirstLuggage) *Handler {
	return &Handler{repo: repo}
}

//GetCountries - retrieves list of countries
func (h *Handler) GetCountries(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://flapi.bucklehosting.com/v2/countries", nil)
	req.Header = r.Header
	resp, err := client.Do(req)
	if err != nil {
		utils.ResponseError(w, 500, "Application error: ", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		utils.ResponseError(w, 500, "Application error: ", err)
		return
	}
	defer resp.Body.Close()
	utils.Response(w, resp.StatusCode, body)
}

//GetItemTypes - retrieves list of item types
func (h *Handler) GetItemTypes(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://flapi.bucklehosting.com/v2/itemtypes", nil)
	req.Header = r.Header
	resp, err := client.Do(req)
	if err != nil {
		utils.ResponseError(w, 500, "Application error: ", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		utils.ResponseError(w, 500, "Application error: ", err)
		return
	}
	defer resp.Body.Close()
	utils.Response(w, resp.StatusCode, body)
}

//GlobalMinimalQuotation is first_luggage get a minimal quote endpoint handler
func (h *Handler) GlobalMinimalQuotation(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://flapi.bucklehosting.com/v2/quote", r.Body)
	if err != nil {
		utils.ResponseError(w, 500, "Application error: ", err)
		return
	}
	req.Header = r.Header
	//_ = json.NewDecoder(r.Body).Decode(&luggageResp)
	resp, err := client.Do(r)
	if err != nil {
		utils.ResponseError(w, 500, "Application error: ", err)
		return
	}
	defer resp.Body.Close()

}
