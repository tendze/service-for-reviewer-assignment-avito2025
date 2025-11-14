package team

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewHandler() http.Handler {
	r := chi.NewRouter()

	r.Post("/add", saveTeam)
	r.Get("/get", getTeam)

	return r
}

func saveTeam(w http.ResponseWriter, r *http.Request) {
	// TODO: 
}

func getTeam(w http.ResponseWriter, r *http.Request) {
	// TODO: 
}
