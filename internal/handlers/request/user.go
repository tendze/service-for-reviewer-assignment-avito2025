package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"dang.z.v.task/internal/domain"
)

type SaveUserRequest struct {
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
	TeamId   uint   `json:"team_id"`
}

func (req *SaveUserRequest) Bind(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	return req.validate()
}

func (req *SaveUserRequest) validate() error {
	if req.Name == "" {
		return fmt.Errorf("name is required field")
	}

	if req.TeamId <= 0 {
		return fmt.Errorf("team_id is required field")
	}

	return nil
}

func (sur *SaveUserRequest) Domain() domain.User {
	return domain.User{
		ID:        0,
		Name:      sur.Name,
		IsActive:  sur.IsActive,
		TeamID:    sur.TeamId,
		CreatedAt: time.Time{},
	}
}

type SetIsActiveRequest struct {
	UserID   uint  `json:"user_id"`
	IsActive *bool `json:"is_active"`
}

func (req *SetIsActiveRequest) Bind(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	return req.validate()
}

func (req *SetIsActiveRequest) validate() error {
	if req.UserID <= 0 {
		return fmt.Errorf("user_id is required field")
	}

	if req.IsActive == nil {
		return fmt.Errorf("is_active is required field")
	}

	return nil
}

type GetReviewRequest struct {
	UserID uint
}

func (req *GetReviewRequest) Bind(r *http.Request) error {
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		return fmt.Errorf("user_id is required query parameter")
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		return fmt.Errorf("user_id must be a valid number")
	}

	req.UserID = uint(userID)
	
	return req.validate()
}

func (req *GetReviewRequest) validate() error {
	if req.UserID <= 0 {
		return fmt.Errorf("user_id must be greater than 0")
	}

	return nil
}