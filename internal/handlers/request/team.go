package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"dang.z.v.task/internal/domain"
)

type AddTeamRequest struct {
	TeamName string        `json:"team_name"`
	Members  *[]TeamMember `json:"members"`
}

type TeamMember struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}

func (req *AddTeamRequest) Bind(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	return req.validate()
}

func (req *AddTeamRequest) validate() error {
	if req.TeamName == "" {
		return fmt.Errorf("team_name is required field")
	}

	if req.Members == nil {
		return fmt.Errorf("members is required field")
	}

	return nil
}

type GetTeamRequest struct {
	TeamName string
}

func (req *GetTeamRequest) Bind(r *http.Request) error {
	teamName := r.URL.Query().Get("team_name")
	if teamName == "" {
		return fmt.Errorf("team_name is required query parameter")
	}

	req.TeamName = teamName

	return nil
}

func (req *AddTeamRequest) TeamMembersToUsersDomain() []domain.User {
	if req.Members == nil {
		return []domain.User{}
	}

	res := make([]domain.User, 0, len(*req.Members))
	for _, member := range *req.Members {
		res = append(
			res,
			domain.User{
				ID:        member.UserID,
				Name:      member.Username,
				IsActive:  member.IsActive,
				TeamID:    0,
				CreatedAt: time.Time{},
			},
		)
	}

	return res
}
