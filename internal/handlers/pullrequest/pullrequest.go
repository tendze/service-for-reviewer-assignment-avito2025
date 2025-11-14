package pullrequest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewHandler() http.Handler {
	r := chi.NewRouter()

	r.Post("/create", createPullRequest)
	r.Post("/merge", mergePullRequest)
	r.Post("/reassign", reassingPullRequest)

	return r
}

func createPullRequest(w http.ResponseWriter, r *http.Request) {
	// TODO:
}

func mergePullRequest(w http.ResponseWriter, r *http.Request) {
	// TODO:
}

func reassingPullRequest(w http.ResponseWriter, r *http.Request) {
	// TODO:
}
