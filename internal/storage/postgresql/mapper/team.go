package mapper

import (
	"dang.z.v.task/internal/domain"
	"dang.z.v.task/internal/storage/postgresql/models"
)

func TeamDomainToModel(team domain.Team) models.Team {
	return models.Team{
		Name:      team.Name,
		CreatedAt: team.CreatedAt,
	}
}

func TeamModelToDomain(team models.Team) domain.Team {
	return domain.Team{
		ID:        team.ID,
		Name:      team.Name,
		CreatedAt: team.CreatedAt,
	}
}
