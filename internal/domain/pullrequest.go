package domain

import "time"

const (
	StatusOpen   = "OPEN"
	StatusMerged = "MERGED"
)

type PullRequest struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	AuthorID  uint       `json:"author_id"`
	Status    string     `json:"status"` // OPEN, MERGED
	CreatedAt time.Time  `json:"created_at"`
	MergedAt  *time.Time `json:"merged_at,omitempty"`
}
