package user

import (
	"net/http"

	"dang.z.v.task/internal/domain"
	"dang.z.v.task/internal/handlers/mapper"
	"dang.z.v.task/internal/handlers/request"
	"dang.z.v.task/internal/handlers/response"
	"github.com/go-chi/chi/v5"
)

type UserService interface {
	SetIsActive(userID uint, isActive bool) (*domain.User, string, error)
	GetReview(userID uint) (*[]domain.PullRequest, error)
	GetPRsByStatus(userID uint, status string) (*[]domain.PullRequest, error)
}

type UserHandler struct {
	userService UserService
}

func NewHandler(userService UserService) http.Handler {
	handler := &UserHandler{
		userService: userService,
	}

	r := chi.NewRouter()

	r.Post("/setIsActive", handler.setIsActive)
	r.Get("/getReview", handler.getReview)
	r.Get("/getPR", handler.getPRsByStatus)

	return r
}

func (h *UserHandler) getPRsByStatus(w http.ResponseWriter, r *http.Request) {
	var req request.GetPRsByStatusRequest
	if err := req.Bind(r); err != nil {
		response.JSONError(w, http.StatusBadRequest, response.BAD_REQUEST, err.Error())

		return
	}

	prs, err := h.userService.GetPRsByStatus(req.UserID, req.Status)
	if err != nil {
		var (
			code         = mapper.HTTPStatusFromError(err)
			errorMsg     = mapper.ErrorMessageFromError(err)
			errorCodeMsg = mapper.ErrorCodeMessageFromError(err)
		)

		response.JSONError(
			w,
			code,
			errorCodeMsg,
			errorMsg,
		)

		return
	}

	response.JSONSuccess(w, http.StatusOK, response.NewGetPRsByStatusResponse(prs))
}

func (h *UserHandler) setIsActive(w http.ResponseWriter, r *http.Request) {
	var req request.SetIsActiveRequest
	if err := req.Bind(r); err != nil {
		response.JSONError(w, http.StatusBadRequest, response.BAD_REQUEST, err.Error())

		return
	}

	user, teamName, err := h.userService.SetIsActive(req.UserID, *req.IsActive)
	if err != nil {
		var (
			code         = mapper.HTTPStatusFromError(err)
			errorMsg     = mapper.ErrorMessageFromError(err)
			errorCodeMsg = mapper.ErrorCodeMessageFromError(err)
		)

		response.JSONError(
			w,
			code,
			errorCodeMsg,
			errorMsg,
		)

		return
	}

	response.JSONSuccess(w, http.StatusOK, response.NewSetIsActiveReponse(*user, teamName))
}

func (h *UserHandler) getReview(w http.ResponseWriter, r *http.Request) {
	var req request.GetReviewRequest
	if err := req.Bind(r); err != nil {
		response.JSONError(w, http.StatusBadRequest, response.BAD_REQUEST, err.Error())

		return
	}

	prs, err := h.userService.GetReview(req.UserID)
	if err != nil {
		var (
			code         = mapper.HTTPStatusFromError(err)
			errorMsg     = mapper.ErrorMessageFromError(err)
			errorCodeMsg = mapper.ErrorCodeMessageFromError(err)
		)

		response.JSONError(
			w,
			code,
			errorCodeMsg,
			errorMsg,
		)

		return
	}

	response.JSONSuccess(w, http.StatusOK, response.NewGetPRResponse(req.UserID, prs))
}
