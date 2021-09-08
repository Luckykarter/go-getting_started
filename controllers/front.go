package controllers

import (
	"encoding/json"
	"net/http"
)

func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)

}

func encodeResponseAsJSON(data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
