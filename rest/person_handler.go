package rest

import (
	"encoding/json"
	"golang-rest-crud/structs"
	"net/http"

	"github.com/go-chi/render"
)

// CreatePerson creates a person in the database
func (h *Handler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	var p structs.Person
	err := json.NewDecoder(r.Body).Decode(&p)
	result, err := h.db.AddPerson(p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, result)
}
