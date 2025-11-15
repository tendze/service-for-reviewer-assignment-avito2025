package request

import (
	"encoding/json"
	"fmt"
	"net/http"

	"dang.z.v.task/internal/domain"
)

type CreatePRRequest struct {
	PullRequestID   uint   `json:"pull_request_id"`
	PullRequestName string `json:"pull_request_name"`
	AuthorID        uint   `json:"author_id"`
}

func (req *CreatePRRequest) Bind(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	return req.validate()
}

func (req *CreatePRRequest) validate() error {
	if req.AuthorID <= 0 {
		return fmt.Errorf("author_id is required field")
	}

	if req.PullRequestID <= 0 {
		return fmt.Errorf("pull_request_id is required field")
	}

	if req.PullRequestName == "" {
		return fmt.Errorf("pull_request_name is required field")
	}

	return nil
}

func (req *CreatePRRequest) Domain() domain.PullRequest {
	return domain.PullRequest{
		Title:    req.PullRequestName,
		AuthorID: req.AuthorID,
		Status:   domain.StatusOpen,
	}
}

type MergePRRequest struct {
	PullRequestID uint `json:"pull_request_id"`
}

func (req *MergePRRequest) Bind(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	return req.validate()
}

func (req *MergePRRequest) validate() error {
	if req.PullRequestID <= 0 {
		return fmt.Errorf("pull_request_id is required field")
	}

	return nil
}

type ReassignRequest struct {
	PullRequestID uint `json:"pull_request_id"`
	OldReviewerID uint `json:"old_reviewer_id"`
}

func (req *ReassignRequest) Bind(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	return req.validate()
}

func (req *ReassignRequest) validate() error {
	if req.OldReviewerID <= 0 {
		return fmt.Errorf("old_reviewer_id is required field")
	}

	if req.PullRequestID <= 0 {
		return fmt.Errorf("pull_request_id is required field")
	}

	return nil
}
