package utils

import (
	"log"
	"net/http"
	"strings"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Response(w http.ResponseWriter, code int, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	w.Write(body)
}

func ResponseError(w http.ResponseWriter, code int, prependText string, err ...error) {
	for _, v := range err {
		prependText = strings.Join([]string{prependText, v.Error()}, " ")
	}
	prependText = strings.Join([]string{`{ "message":"`, prependText, `"}`}, ``)
	log.Println(prependText)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	w.Write([]byte(prependText))
}
