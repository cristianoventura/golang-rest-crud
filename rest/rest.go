package rest

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Router declares the api routes
func Router(h HandlerInterface) *chi.Mux {
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/api", func(r chi.Router) {
		r.Post("/person", h.CreatePerson)
		r.Get("/person/{personID}", h.GetPersonByID)
	})

	return r
}

// Start boots the rest api
func Start() {
	h, err := NewHadler()
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":8080", Router(h))
}
