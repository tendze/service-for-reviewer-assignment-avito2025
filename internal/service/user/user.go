package userservice

import (
	"fmt"
	"log/slog"

	"dang.z.v.task/internal/domain"
	"dang.z.v.task/internal/service/mapper"
)

type UserRepository interface {
	UpdateUserActiveStatus(userID uint, isActive bool) (*domain.User, error)
	GetUserTeamName(userID uint) (string, error)
	GetPRsByReviewer(userID uint) (*[]domain.PullRequest, error)
	GetPRsByStatus(userID uint, status string) (*[]domain.PullRequest, error)
}

type UserService struct {
	userRepo UserRepository
	log      *slog.Logger
}

func NewUserService(
	usrRepository UserRepository,
	log *slog.Logger,
) *UserService {
	return &UserService{
		userRepo: usrRepository,
		log:      log,
	}
}

func (s *UserService) GetPRsByStatus(userID uint, status string) (*[]domain.PullRequest, error) {
	const op = "userservice.user.GetPRsByStatus"
	log := s.log.With(slog.String("op", op))

	usr, err := s.userRepo.GetPRsByStatus(userID, status)
	if err != nil {
		log.Error("failed to get pr's by status", slog.Any("err", err))

		return nil, fmt.Errorf("%s: %w", op, mapper.MapStorageError(err))
	}

	return usr, nil
}

// / Returns User, User's team name and error
func (s *UserService) SetIsActive(userID uint, isActive bool) (*domain.User, string, error) {
	const op = "userservice.user.SetIsActive"
	log := s.log.With(slog.String("op", op))

	usr, err := s.userRepo.UpdateUserActiveStatus(userID, isActive)
	if err != nil {
		log.Error("failed to update user's active status", slog.Any("err", err))

		return nil, "", fmt.Errorf("%s: %w", op, mapper.MapStorageError(err))
	}

	teamName, err := s.userRepo.GetUserTeamName(userID)
	if err != nil {
		log.Error("failed to get user's team name", slog.Any("err", err))

		return nil, "", fmt.Errorf("%s: %w", op, mapper.MapStorageError(err))
	}

	return usr, teamName, nil
}

func (s *UserService) GetReview(userID uint) (*[]domain.PullRequest, error) {
	const op = "userservice.user.GetReview"
	log := s.log.With(slog.String("op", op))

	prs, err := s.userRepo.GetPRsByReviewer(userID)
	if err != nil {
		log.Error("failed to update user's active status", slog.Any("err", err))

		return nil, fmt.Errorf("%s: %w", op, mapper.MapStorageError(err))
	}

	return prs, nil
}
