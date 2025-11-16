package prservice

import (
	"log/slog"
	"time"

	"dang.z.v.task/internal/domain"
)

type PRRepository interface {
	SetMergedAt(prID uint, time time.Time) (domain.PullRequest, error)
	GetUserReviewersByPRID(prID uint) (*[]domain.User, error)
	CreatePullRequest(pr domain.PullRequest) (uint, *[]domain.User, error)
	ReassignReviewer(prID uint, oldReviewerID uint) (domain.PullRequest, *[]domain.User, uint, error)
}

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

func (s *PRService) CreatePullRequest(pr domain.PullRequest) (uint, *[]domain.User, error) {
	const op = "prservice.CreatePullRequest"
	log := s.log.With(slog.String("op", op))

	prID, assigned, err := s.prRepo.CreatePullRequest(pr)
	if err != nil {
		log.Error("failed to create PR", slog.Any("err", err))
		return 0, nil, err
	}

	return prID, assigned, nil
}

func (s *PRService) MergePullRequest(prID uint) (domain.PullRequest, *[]domain.User, error) {
	const op = "prservice.MergePullRequest"
	log := s.log.With(slog.String("op", op))

	updatedPR, err := s.prRepo.SetMergedAt(prID, time.Now())
	if err != nil {
		log.Error("failed to set merged at", slog.Any("err", err))

		return domain.PullRequest{}, nil, err
	}

	prReviewers, err := s.prRepo.GetUserReviewersByPRID(prID)
	if err != nil {
		log.Error("failed to get pr's reviewers", slog.Any("err", err))

		return domain.PullRequest{}, nil, err
	}

	return updatedPR, prReviewers, nil
}

// / Returns pull request, all reviewers, id of new reviewer and an error
func (s *PRService) ReassignReviewer(prID uint, oldReviewerID uint) (domain.PullRequest, *[]domain.User, uint, error) {
	const op = "prservice.ReassignReviewer"
	log := s.log.With(slog.String("op", op))

	updatedPR, reviewers, newReviewerID, err := s.prRepo.ReassignReviewer(prID, oldReviewerID)
	if err != nil {
		log.Error("failed to reassign reviewer", slog.Any("err", err))
		return domain.PullRequest{}, nil, 0, err
	}

	return updatedPR, reviewers, newReviewerID, nil
}
