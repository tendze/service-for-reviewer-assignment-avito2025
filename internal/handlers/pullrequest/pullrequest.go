package pullrequest

import (
	"net/http"

	"dang.z.v.task/internal/domain"
	"dang.z.v.task/internal/handlers/request"
	"dang.z.v.task/internal/handlers/response"
	"github.com/go-chi/chi/v5"
)

type PRService interface {
	CreatePullRequest(domain.PullRequest) (*[]domain.User, error)
	MergePullRequest(uint) (domain.PullRequest, *[]domain.User, error)
	ReassignReviewer(uint, uint) (domain.PullRequest, *[]domain.User, uint, error)
}

type PRHandler struct {
	prService PRService
}

func NewHandler(prService PRService) http.Handler {
	handler := &PRHandler{
		prService: prService,
	}

	r := chi.NewRouter()

	r.Post("/create", handler.createPullRequest)
	r.Post("/merge", handler.mergePullRequest)
	r.Post("/reassign", handler.reassingPullRequest)

	return r
}

func (h *PRHandler) createPullRequest(w http.ResponseWriter, r *http.Request) {
	var req request.CreatePRRequest
	if err := req.Bind(r); err != nil {
		response.JSONError(w, http.StatusBadRequest, response.BAD_REQUEST, err.Error())

		return
	}

	prDomain := req.Domain()

	assignedUsers, err := h.prService.CreatePullRequest(prDomain)
	if err != nil {
		response.JSONError(
			w,
			http.StatusInternalServerError,
			response.INTERNAL_SERVER_ERRROR,
			err.Error(),
		)

		return
	}

	response.JSONSuccess(
		w,
		http.StatusCreated,
		response.NewCreatePRResponse(
			prDomain,
			domain.StatusOpen,
			assignedUsers,
		),
	)
}

func (h *PRHandler) mergePullRequest(w http.ResponseWriter, r *http.Request) {
	var req request.MergePRRequest
	if err := req.Bind(r); err != nil {
		response.JSONError(w, http.StatusBadRequest, response.BAD_REQUEST, err.Error())

		return
	}

	pr, assignedUsers, err := h.prService.MergePullRequest(req.PullRequestID)
	if err != nil {
		response.JSONError(
			w,
			http.StatusInternalServerError,
			response.INTERNAL_SERVER_ERRROR,
			err.Error(),
		)

		return
	}

	response.JSONSuccess(
		w,
		http.StatusOK,
		response.NewMergePRResponse(
			pr,
			domain.StatusMerged,
			assignedUsers,
		),
	)
}

func (h *PRHandler) reassingPullRequest(w http.ResponseWriter, r *http.Request) {
	var req request.ReassignRequest
	if err := req.Bind(r); err != nil {
		response.JSONError(
			w,
			http.StatusInternalServerError,
			response.INTERNAL_SERVER_ERRROR,
			err.Error(),
		)

		return
	}

	pr, assignedUsers, replacedByID, err := h.prService.ReassignReviewer(req.PullRequestID, req.OldReviewerID)
	if err != nil {
		response.JSONError(
			w,
			http.StatusInternalServerError,
			response.INTERNAL_SERVER_ERRROR,
			err.Error(),
		)

		return
	}

	response.JSONSuccess(
		w,
		http.StatusCreated,
		response.NewReassignReviewerResponse(
			pr,
			replacedByID,
			assignedUsers,
		),
	)
}
