package postgresql

import (
	"fmt"

	"dang.z.v.task/internal/domain"
	"dang.z.v.task/internal/storage/postgresql/mapper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Storage struct {
	db *gorm.DB
}

func New(dsn string) (*Storage, error) {
	const op = "postgres.New"

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
