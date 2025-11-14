package response

import (
	"encoding/json"
	"net/http"

	"dang.z.v.task/internal/domain"
)

type SetIsActiveResponse struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	TeamName string `json:"team_name"`
	IsActive bool   `json:"is_active"`
}

func NewSetIsActiveReponse(user domain.User, teamName string) map[string]interface{} {
	return map[string]interface{}{
		"user": SetIsActiveResponse{
			UserID:   user.ID,
			Username: user.Name,
			IsActive: user.IsActive,
			TeamName: teamName,
		},
	}
}

type GetPRResponse struct {
	UserID       uint           `json:"user_id"`
	PullRequests []PullRequests `json:"pull_requests"`
}

type PullRequests struct {
	PullRequestID   uint   `json:"pull_request_id"`
	PullRequestName string `json:"pull_request_name"`
	AuthorID        uint   `json:"author_id"`
	Status          string `json:"open"`
}

func NewGetPRResponse(userID uint, prs *[]domain.PullRequest) map[string]interface{} {
	return map[string]interface{}{
		"user_id":       userID,
		"pull_requests": prDomainsToResponse(prs),
	}
}

func JSONSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)
}

func prDomainsToResponse(prs *[]domain.PullRequest) []PullRequests {
	if prs == nil {
		return []PullRequests{}
	}

	res := make([]PullRequests, 0, len(*prs))
	for _, pr := range *prs {
		res = append(
			res,
			PullRequests{
				PullRequestID:   pr.ID,
				PullRequestName: pr.Title,
				AuthorID:        pr.AuthorID,
				Status:          pr.Status,
			},
		)
	}

	return res
}
