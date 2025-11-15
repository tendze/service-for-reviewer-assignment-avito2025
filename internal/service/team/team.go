package teamservice

import (
	"log/slog"

	"dang.z.v.task/internal/domain"
)

type TeamRepository interface {
	AddTeamWithUsersAtomic(team domain.Team, users []domain.User) ([]domain.User, error)
	GetTeamMembers(teamName string) ([]domain.User, error)
}

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
	const op = "teamservice.team.AddTeam"
	log := s.log.With(slog.String("op", op))

	savedUsers, err := s.teamRepo.AddTeamWithUsersAtomic(domain.Team{Name: teamName}, users)
	if err != nil {
		log.Error("failed to add team", slog.Any("err", err))

		return []domain.User{}, err
	}

	return savedUsers, nil
}

func (s *TeamService) GetTeam(teamName string) ([]domain.User, error) {
	const op = "teamservice.team.GetTeam"
	log := s.log.With(slog.String("op", op))

	teamMembers, err := s.teamRepo.GetTeamMembers(teamName)
	if err != nil {
		log.Error("failed to add team", slog.Any("err", err))

		return []domain.User{}, err
	}

	return teamMembers, err
}
