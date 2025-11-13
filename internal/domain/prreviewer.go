package domain

import "time"

type PRReviewer struct {
	ID         uint      `json:"id"`
	PRID       uint      `json:"pr_id"`
	ReviewerID uint      `json:"reviewer_id"`
	CreatedAt  time.Time `json:"created_at"`
}
