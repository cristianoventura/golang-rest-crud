package rest

import (
	"golang-rest-crud/database"
	"net/http"
)

// HandlerInterface contains the implementation for each route
type HandlerInterface interface {
	CreatePerson(w http.ResponseWriter, r *http.Request)
	GetPersonByID(w http.ResponseWriter, r *http.Request)
}

// Handler gives access to request, response and database
type Handler struct {
	r  *http.Request
	w  http.ResponseWriter
	db database.ServiceInterface
}

// NewHadler connects to the database and returns a new HandlerInterface
func NewHadler() (HandlerInterface, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	return &Handler{
		db: db,
	}, nil
}
