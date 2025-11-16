package prservice

import (
	"fmt"
	"log/slog"
	"time"

	"dang.z.v.task/internal/domain"
	"dang.z.v.task/internal/service"
	"dang.z.v.task/internal/service/mapper"
)

type PRRepository interface {
	SetMergedAt(prID uint, time time.Time) (domain.PullRequest, error)
	GetUserReviewersByPRID(prID uint) (*[]domain.User, error)
	CreatePullRequest(pr domain.PullRequest) (uint, *[]domain.User, error)
	ReassignReviewer(prID uint, oldReviewerID uint) (domain.PullRequest, *[]domain.User, uint, error)
	GetPRStatus(prID uint) (string, error)
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

		return 0, nil, fmt.Errorf("%s: %w", op, mapper.MapStorageError(err))
	}

	return prID, assigned, nil
}

func (s *PRService) MergePullRequest(prID uint) (domain.PullRequest, *[]domain.User, error) {
	const op = "prservice.MergePullRequest"
	log := s.log.With(slog.String("op", op))

	updatedPR, err := s.prRepo.SetMergedAt(prID, time.Now())
	if err != nil {
		log.Error("failed to set merged at", slog.Any("err", err))

		return domain.PullRequest{}, nil, fmt.Errorf("%s: %w", op, mapper.MapStorageError(err))
	}

	prReviewers, err := s.prRepo.GetUserReviewersByPRID(prID)
	if err != nil {
		log.Error("failed to get pr's reviewers", slog.Any("err", err))

		return domain.PullRequest{}, nil, fmt.Errorf("%s: %w", op, mapper.MapStorageError(err))
	}

	return updatedPR, prReviewers, nil
}

// / Returns pull request, all reviewers, id of new reviewer and an error
func (s *PRService) ReassignReviewer(prID uint, oldReviewerID uint) (domain.PullRequest, *[]domain.User, uint, error) {
	const op = "prservice.ReassignReviewer"
	log := s.log.With(slog.String("op", op))

	prStatus, err := s.prRepo.GetPRStatus(prID)
	if err != nil {
		log.Error("failed to get pr's status", slog.Any("err", err))

		return domain.PullRequest{}, nil, 0, fmt.Errorf("%s: %w", op, mapper.MapStorageError(err))
	}

	if prStatus == "MERGED" {
		log.Error("cannot reassign on merged pr", slog.Any("err", err))

		return domain.PullRequest{}, nil, 0, fmt.Errorf("%s: %w", op, service.ErrPullRequestMerged)
	}

	updatedPR, reviewers, newReviewerID, err := s.prRepo.ReassignReviewer(prID, oldReviewerID)
	if err != nil {
		log.Error("failed to reassign reviewer", slog.Any("err", err))

		return domain.PullRequest{}, nil, 0, fmt.Errorf("%s: %w", op, mapper.MapStorageError(err))
	}

	return updatedPR, reviewers, newReviewerID, nil
}
