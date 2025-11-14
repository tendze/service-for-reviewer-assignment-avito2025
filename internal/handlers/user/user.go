package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewHandler() http.Handler {
	r := chi.NewRouter()

	r.Post("/add", saveUser)
	r.Post("/setIsActive", setIsActive)
	r.Get("/getReview", getReview)

	return r
}

func saveUser(w http.ResponseWriter, r *http.Request) {
	// TODO: 
}

func setIsActive(w http.ResponseWriter, r *http.Request) {
	// TODO:
}

func getReview(w http.ResponseWriter, r *http.Request) {
	// TODO:
}
