package response

import (
	"time"

	"dang.z.v.task/internal/domain"
)

type CreatePRResponse struct {
	PullRequestID        uint   `json:"pull_request_id"`
	PullRequestName      string `json:"pull_request_name"`
	AuthorID             uint   `json:"author_id"`
	Status               string `json:"status"`
	AssignedReviewersIDs []uint `json:"assigned_reviewers"`
}

func NewCreatePRResponse(
	pr domain.PullRequest,
	status string,
	assignedUsers *[]domain.User,
) map[string]interface{} {
	return map[string]interface{}{
		"pr": CreatePRResponse{
			PullRequestID:        pr.ID,
			PullRequestName:      pr.Title,
			AuthorID:             pr.AuthorID,
			Status:               status,
			AssignedReviewersIDs: userToIDs(assignedUsers),
		},
	}
}

type MergePRResponse struct {
	PullRequestID        uint      `json:"pull_request_id"`
	PullRequestName      string    `json:"pull_request_name"`
	AuthorID             uint      `json:"author_id"`
	Status               string    `json:"status"`
	AssignedReviewersIDs []uint    `json:"assigned_reviewers"`
	MergedAt             time.Time `json:"merged_at"`
}

func NewMergePRResponse(
	pr domain.PullRequest,
	status string,
	assignedUsers *[]domain.User,
) map[string]interface{} {
	return map[string]interface{}{
		"pr": MergePRResponse{
			PullRequestID:        pr.ID,
			PullRequestName:      pr.Title,
			AuthorID:             pr.AuthorID,
			Status:               status,
			AssignedReviewersIDs: userToIDs(assignedUsers),
			MergedAt:             *pr.MergedAt,
		},
	}
}

type ReassingReviewerResponse struct {
	PullRequestID       uint   `json:"pull_request_id"`
	PullRequestName     string `json:"pull_request_name"`
	AuthorID            uint   `json:"author_id"`
	Status              string `json:"status"`
	AssignedReviewerIDs []uint `json:"assigned_reviewers"`
}

func NewReassignReviewerResponse(
	pr domain.PullRequest,
	replacedByID uint,
	assignedUsers *[]domain.User,
) map[string]interface{} {
	return map[string]interface{}{
		"pr": ReassingReviewerResponse{
			PullRequestID:       pr.ID,
			PullRequestName:     pr.Title,
			AuthorID:            pr.AuthorID,
			Status:              pr.Status,
			AssignedReviewerIDs: userToIDs(assignedUsers),
		},
		"replaced_by": replacedByID,
	}
}

func userToIDs(assignedUsers *[]domain.User) []uint {
	if assignedUsers == nil {
		return []uint{}
	}

	res := make([]uint, 0, len(*assignedUsers))
	for _, usr := range *assignedUsers {
		res = append(res, usr.ID)
	}

	return res
}
