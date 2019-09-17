package controllers

import (
	"net/http"
	u "../utils"
	"log"
)

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	u.Respond(w, map[string]interface{}{"test":"test"})
}