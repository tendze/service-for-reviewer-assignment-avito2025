package postgresql

import (
	"fmt"

	"dang.z.v.task/internal/domain"
	"dang.z.v.task/internal/storage/postgresql/mapper"
	"dang.z.v.task/internal/storage/postgresql/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Storage struct {
	db *gorm.DB
}

func New(dsn string) (*Storage, error) {
	const op = "postgres.New"

	// TODO: turn on silent mode
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	storage := &Storage{db: db}

	return storage, nil
}

func (s *Storage) SaveUser(usr domain.User) error {
	const op = "postgres.SaveUser"

	pgmodel := mapper.UserDomainToModel(usr)

	result := s.db.Create(&pgmodel)
	if err := result.Error; err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) SaveTeam(team domain.Team) error {
	const op = "postgres.SaveTeam"

	pgmodel := mapper.TeamDomainToModel(team)

	result := s.db.Create(&pgmodel)
	if err := result.Error; err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) AddTeamWithUsersAtomic(team domain.Team, users []domain.User) ([]domain.User, error) {
	const op = "postgres.SaveTeamWithUsersAtomic"

	var savedUsers []domain.User

	return savedUsers,
		s.db.Transaction(func(tx *gorm.DB) error {
			teamModel := mapper.TeamDomainToModel(team)
			if err := tx.Create(&teamModel).Error; err != nil {
				return fmt.Errorf("%s: save team: %w", op, err)
			}

			userModels := make([]models.User, 0, len(users))
			for _, user := range users {
				user.TeamID = teamModel.ID
				userModels = append(userModels, mapper.UserDomainToModel(user))
			}

			if err := tx.Create(&userModels).Error; err != nil {
				return fmt.Errorf("%s: save users: %w", op, err)
			}

			savedUsers = mapper.UserModelsToDomains(userModels)

			return nil
		})
}

func (s *Storage) GetTeamMembers(teamName string) ([]domain.User, error) {
	const op = "postgres.GetTeamMembers"

	var team models.Team
	if err := s.db.Where("name = ?", teamName).First(&team).Error; err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var users []models.User
	if err := s.db.Where("team_id = ?", team.ID).Find(&users).Error; err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return mapper.UserModelsToDomains(users), nil
}
