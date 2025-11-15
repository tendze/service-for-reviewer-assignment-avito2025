package teamservice

import (
	"log/slog"

	"dang.z.v.task/internal/domain"
)

type TeamRepository interface{}

type TeamService struct {
	teamRepo TeamRepository
	log      *slog.Logger
}

func NewTeamService(
	teamRepository TeamRepository,
	log *slog.Logger,
) *TeamService {
	return &TeamService{
		teamRepo: teamRepository,
		log:      log,
	}
}

func (s *TeamService) AddTeam(teamName string, users []domain.User) ([]domain.User, error) {
	return []domain.User{}, nil
}

func (s *TeamService) GetTeam(teamName string) ([]domain.User, error) {
	return []domain.User{}, nil
}
