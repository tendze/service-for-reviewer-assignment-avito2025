package response

import (
	"time"

	"dang.z.v.task/internal/domain"
)

type GetPRsByStatusResponse struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	AuthorID  uint       `json:"author_id"`
	Status    string     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	MergedAt  *time.Time `json:"merged_at,omitempty"`
}

func NewGetPRsByStatusResponse(prs *[]domain.PullRequest) map[string]interface{} {
	res := map[string]interface{}{
		"pull_requests": []interface{}{},
	}

	if prs == nil {
		return res
	}

	resp := make([]GetPRsByStatusResponse, 0, len(*prs))
	for _, pr := range *prs {
		resp = append(
			resp,
			GetPRsByStatusResponse{
				ID:        pr.ID,
				Title:     pr.Title,
				AuthorID:  pr.AuthorID,
				Status:    pr.Status,
				CreatedAt: pr.CreatedAt,
				MergedAt:  pr.MergedAt,
			},
		)
	}

	res["pull_requests"] = resp

	return res
}

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
