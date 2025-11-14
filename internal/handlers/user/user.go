package user

import (
	"net/http"

	"dang.z.v.task/internal/domain"
	"dang.z.v.task/internal/handlers/request"
	"dang.z.v.task/internal/handlers/response"
	"github.com/go-chi/chi/v5"
)

type UserService interface {
	SaveUser(domain.User) error
	SetIsActive(userID uint, isActive bool) (*domain.User, string, error)
	GetReview(userID uint) (*[]domain.PullRequest, error)
}

type UserHandler struct {
	userService UserService
}

func NewHandler(userService UserService) http.Handler {
	handler := &UserHandler{
		userService: userService,
	}

	r := chi.NewRouter()

	r.Post("/add", handler.saveUser)
	r.Post("/setIsActive", handler.setIsActive)
	r.Get("/getReview", handler.getReview)

	return r
}

func (h *UserHandler) saveUser(w http.ResponseWriter, r *http.Request) {
	var req request.SaveUserRequest
	if err := req.Bind(r); err != nil {
		response.JSONError(w, http.StatusBadRequest, response.BAD_REQUEST, err.Error())

		return
	}

	userDomain := req.Domain()

	err := h.userService.SaveUser(userDomain)
	if err != nil {
		response.JSONError(
			w,
			http.StatusInternalServerError,
			response.INTERNAL_SERVER_ERRROR,
			err.Error(),
		)

		return
	}
}

func (h *UserHandler) setIsActive(w http.ResponseWriter, r *http.Request) {
	var req request.SetIsActiveRequest
	if err := req.Bind(r); err != nil {
		response.JSONError(w, http.StatusBadRequest, response.BAD_REQUEST, err.Error())

		return
	}

	user, teamName, err := h.userService.SetIsActive(req.UserID, *req.IsActive)
	if err != nil {
		response.JSONError(
			w,
			http.StatusInternalServerError,
			response.INTERNAL_SERVER_ERRROR,
			err.Error(),
		)

		return
	}

	response.JSONSuccess(w, response.NewSetIsActiveReponse(*user, teamName))
}

func (h *UserHandler) getReview(w http.ResponseWriter, r *http.Request) {
	var req request.GetReviewRequest
	if err := req.Bind(r); err != nil {
		response.JSONError(w, http.StatusBadRequest, response.BAD_REQUEST, err.Error())

		return
	}

	prs, err := h.userService.GetReview(req.UserID)
	if err != nil {
		response.JSONError(
			w,
			http.StatusInternalServerError,
			response.INTERNAL_SERVER_ERRROR,
			err.Error(),
		)

		return
	}

	response.JSONSuccess(w, response.NewGetPRResponse(req.UserID, prs))
}
