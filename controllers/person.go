package controllers

import (
	"encoding/json"
	"net/http"

	"../services"
	"go.mongodb.org/mongo-driver/mongo"
)

type server int

var dbInstance *mongo.Database

// PersonController handles the /person requests
func PersonController(db *mongo.Database) *http.ServeMux {
	dbInstance = db
	var s server
	mux := http.NewServeMux()
	mux.Handle("/person", s)
	return mux
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Internal Server Error"))
		}

		id := r.FormValue("id")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
		}

		res, err := services.FindPerson(id, dbInstance)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Internal Server Error"))
		}

		content, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Internal Server Error"))
		}
		w.Write(content)
	}
}
