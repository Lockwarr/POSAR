package controllers

import (
	"net/http"
	u "../utils"
)

var Test = func(w http.ResponseWriter, r *http.Request) {
	u.Respond(w, map[string]interface{}{"test":"test"})
}