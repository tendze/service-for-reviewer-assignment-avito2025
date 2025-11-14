package userservice

import (
	"log/slog"

	"dang.z.v.task/internal/domain"
)

type UserRepository interface{}

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

func (s *UserService) SaveUser(user domain.User) error {
	return nil
}

// / Returns User, User's team name and error
func (s *UserService) SetIsActive(userID uint, isActive bool) (*domain.User, string, error) {
	return nil, "", nil
}

func (s *UserService) GetReview(userID uint) (*[]domain.PullRequest, error) {
	return nil, nil
}
