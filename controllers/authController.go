package controllers

import (
	"net/http"
	u "../utils"
	"../models"
	"encoding/json"
)

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	client := &models.Client{}
	err := json.NewDecoder(r.Body).Decode(client) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	u.Respond(w, map[string]interface{}{"test":"test"})
}