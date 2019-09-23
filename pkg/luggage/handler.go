package luggage

import (
	"io/ioutil"
	"log"
	"net/http"

	"../utils"
	"github.com/Lockwarr/POSAR/pkg/models"
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
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-FL-API-User", "542M9auMnJjwbdyrw97t9EhhFze2PDpG")
	req.Header.Add("X-FL-API-Key", "Nz286rGEwwnLDuNpa6X3KaVw3eJLpTLC")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err, "get")
	}
	defer resp.Body.Close()
	utils.Response(w, resp.StatusCode, body)
}

//GetItemTypes - retrieves list of item types
func (h *Handler) GetItemTypes(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://flapi.bucklehosting.com/v2/itemtypes", nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-FL-API-User", "542M9auMnJjwbdyrw97t9EhhFze2PDpG")
	req.Header.Add("X-FL-API-Key", "Nz286rGEwwnLDuNpa6X3KaVw3eJLpTLC")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err, "get")
	}
	defer resp.Body.Close()
	utils.Response(w, resp.StatusCode, body)
}

//GlobalMinimalQuotation is first_luggage get a minimal quote endpoint handler
func (h *Handler) GlobalMinimalQuotation(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("POST", "https://flapi.bucklehosting.com/v2/quote", r.Body)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	r = req
	//_ = json.NewDecoder(r.Body).Decode(&luggageResp)
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}
	defer resp.Body.Close()

}
