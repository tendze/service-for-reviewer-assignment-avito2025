package prservice

import (
	"log/slog"

	"dang.z.v.task/internal/domain"
)

type PRRepository interface{}

type PRService struct {
	prRepo PRRepository
	log    *slog.Logger
}

func NewPullRequestService(
	prRepository PRRepository,
	log *slog.Logger,
) *PRService {
	return &PRService{
		prRepo: prRepository,
		log:    log,
	}
}

func (s *PRService) CreatePullRequest(domain.PullRequest) (*[]domain.User, error) {
	return &[]domain.User{}, nil
}

func (s *PRService) MergePullRequest(prID uint) (domain.PullRequest, *[]domain.User, error) {
	return domain.PullRequest{}, nil, nil
}

func (s *PRService) ReassignReviewer(uint, uint) (domain.PullRequest, *[]domain.User, uint, error) {
	return domain.PullRequest{}, nil, 0, nil
}
