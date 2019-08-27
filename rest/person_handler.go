package rest

import (
	"encoding/json"
	"fmt"
	"golang-rest-crud/structs"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/go-chi/render"
)

// CreatePerson creates a person in the database
func (h *Handler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	var p structs.Person
	err := json.NewDecoder(r.Body).Decode(&p)
	result, err := h.db.AddPerson(p)
	w.Header().Set("content-type", "application/json")
	if err != nil {
		render.JSON(w, r, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, result)
}

// GetPersonByID returns a record by id from the collection
func (h *Handler) GetPersonByID(w http.ResponseWriter, r *http.Request) {
	personID := chi.URLParam(r, "personID")
	result, err := h.db.FindPerson(personID)
	w.Header().Set("content-type", "application/json")
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, result)
}
