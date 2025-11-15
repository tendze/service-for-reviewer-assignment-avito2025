package team

import (
	"net/http"

	"dang.z.v.task/internal/domain"
	"dang.z.v.task/internal/handlers/request"
	"dang.z.v.task/internal/handlers/response"
	"github.com/go-chi/chi/v5"
)

type TeamService interface {
	AddTeam(string, []domain.User) ([]domain.User, error)
	GetTeam(string) ([]domain.User, error)
}

type TeamHandler struct {
	teamService TeamService
}

func NewHandler(teamService TeamService) http.Handler {
	handler := &TeamHandler{
		teamService: teamService,
	}

	r := chi.NewRouter()

	r.Post("/add", handler.saveTeam)
	r.Get("/get", handler.getTeam)

	return r
}

func (h *TeamHandler) saveTeam(w http.ResponseWriter, r *http.Request) {
	var req request.AddTeamRequest
	if err := req.Bind(r); err != nil {
		response.JSONError(w, http.StatusBadRequest, response.BAD_REQUEST, err.Error())

		return
	}

	domainUsers := req.TeamMembersToUsersDomain()

	teamMembers, err := h.teamService.AddTeam(req.TeamName, domainUsers)
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
		response.NewAddTeamResponse(req.TeamName, &teamMembers),
	)
}

func (h *TeamHandler) getTeam(w http.ResponseWriter, r *http.Request) {
	var req request.GetTeamRequest
	if err := req.Bind(r); err != nil {
		response.JSONError(w, http.StatusBadRequest, response.BAD_REQUEST, err.Error())

		return
	}

	teamMembers, err := h.teamService.GetTeam(req.TeamName)
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
		response.NewGetTeamResponse(req.TeamName, &teamMembers),
	)
}
