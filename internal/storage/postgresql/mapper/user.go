package mapper

import (
	"dang.z.v.task/internal/domain"
	"dang.z.v.task/internal/storage/postgresql/models"
)

func UserDomainToModel(usr domain.User) models.User {
	return models.User{
		ID:        usr.ID,
		Name:      usr.Name,
		IsActive:  usr.IsActive,
		TeamID:    usr.TeamID,
		CreatedAt: usr.CreatedAt,
	}
}

func UserModelToDomain(user models.User) domain.User {
	return domain.User{
		ID:        user.ID,
		Name:      user.Name,
		IsActive:  user.IsActive,
		TeamID:    user.TeamID,
		CreatedAt: user.CreatedAt,
	}
}
