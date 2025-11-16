package postgresql

import (
	"fmt"
	"math/rand"
	"time"

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

func (s *Storage) GetUserTeamName(userID uint) (string, error) {
	var user models.User
	if err := s.db.Preload("Team").First(&user, userID).Error; err != nil {
		return "", err
	}
	return user.Team.Name, nil
}

func (s *Storage) UpdateUserActiveStatus(userID uint, isActive bool) (*domain.User, error) {
	const op = "poostgres.UpdateUserActiveStatus"

	var user models.User

	if err := s.db.First(&user, userID).Error; err != nil {
		return nil, fmt.Errorf("%s: user not found: %w", op, err)
	}

	if err := s.db.Model(&user).Update("is_active", isActive).Error; err != nil {
		return nil, fmt.Errorf("%s: update failed: %w", op, err)
	}

	domainUser := mapper.UserModelToDomain(user)
	domainUser.IsActive = isActive

	return &domainUser, nil
}

func (s *Storage) GetPRsByReviewer(userID uint) (*[]domain.PullRequest, error) {
	const op = "postgres.GetPRsByReviewer"

	var prs []models.PullRequest

	err := s.db.
		Joins("JOIN pr_reviewer ON pr_reviewer.pr_id = pull_request.id").
		Where("pr_reviewer.reviewer_id = ?", userID).
		Preload("Author").
		Find(&prs).Error
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get pr's %w", op, err)
	}

	domainPRs := mapper.PullRequestModelsToDomains(prs)

	return &domainPRs, nil
}

func (s *Storage) SetMergedAt(prID uint, time time.Time) (domain.PullRequest, error) {
	const op = "postgres.SetMergedAt"

	var pr models.PullRequest

	if err := s.db.First(&pr, prID).Error; err != nil {
		return domain.PullRequest{}, fmt.Errorf("%s: pull request not found: %w", op, err)
	}

	if pr.MergedAt != nil {
		return domain.PullRequest{}, fmt.Errorf("%s: PR is already merged", op)
	}

	if err := s.db.Model(&pr).Update("merged_at", time).Error; err != nil {
		return domain.PullRequest{}, fmt.Errorf("%s: failed to update merged at: %w", op, err)
	}

	domainPR := mapper.PullRequestModelToDomain(pr)

	return domainPR, nil
}

func (s *Storage) GetUserReviewersByPRID(prID uint) (*[]domain.User, error) {
	const op = "postgres.GetUserReviewersByPRID"

	var users []models.User

	err := s.db.
		Table(`"user" AS u`).
		Joins(`JOIN pr_reviewer r ON r.reviewer_id = u.id`).
		Where("r.pr_id = ?", prID).
		Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	domainUsers := mapper.UserModelsToDomains(users)

	return &domainUsers, nil
}

func (s *Storage) CreatePullRequest(pr domain.PullRequest) (uint, *[]domain.User, error) {
	const op = "postgres.CreatePullRequest"

	var assignedUsers []domain.User
	var createdPRID uint

	err := s.db.Transaction(func(tx *gorm.DB) error {
		prModel := mapper.PullRequestDomainToModel(pr)
		if err := tx.Create(&prModel).Error; err != nil {
			return fmt.Errorf("%s: create PR: %w", op, err)
		}

		var author models.User
		if err := tx.First(&author, pr.AuthorID).Error; err != nil {
			return fmt.Errorf("%s: get author: %w", op, err)
		}

		// getting every team member except author
		var teamMembers []models.User
		if err := tx.
			Where("team_id = ? AND id <> ? AND is_active = true", author.TeamID, pr.AuthorID).
			Find(&teamMembers).Error; err != nil {
			return fmt.Errorf("%s: get team members: %w", op, err)
		}

		// shuffling
		shuffleSlice(teamMembers)

		var selected []models.User
		for i := 0; i < len(teamMembers) && i < 2; i++ {
			selected = append(selected, teamMembers[i])
		}

		for _, u := range selected {
			link := models.PRReviewer{
				PRID:       prModel.ID,
				ReviewerID: u.ID,
			}
			if err := tx.Create(&link).Error; err != nil {
				return fmt.Errorf("%s: link reviewer: %w", op, err)
			}
		}

		assignedUsers = mapper.UserModelsToDomains(selected)
		createdPRID = prModel.ID

		return nil
	})
	if err != nil {
		return 0, nil, err
	}

	return createdPRID, &assignedUsers, nil
}

func (s *Storage) ReassignReviewer(prID uint, oldReviewerID uint) (domain.PullRequest, *[]domain.User, uint, error) {
	const op = "postgres.ReassignReviewer"

	var pr models.PullRequest
	var currentReviewer models.PRReviewer

	err := s.db.Transaction(func(tx *gorm.DB) error {
		// check if pr exists
		if err := tx.First(&pr, prID).Error; err != nil {
			return fmt.Errorf("%s: PR not found: %w", op, err)
		}

		// check if reviewerID is actually reviewer
		if err := tx.
			Where("pr_id = ? AND reviewer_id = ?", prID, oldReviewerID).
			First(&currentReviewer).Error; err != nil {
			return fmt.Errorf("%s: reviewer is not assigned to this PR: %w", op, err)
		}

		// lookup for pr's author
		var author models.User
		if err := tx.First(&author, pr.AuthorID).Error; err != nil {
			return fmt.Errorf("%s: author not found: %w", op, err)
		}

		// lookup for active team members except author
		var teamMembers []models.User
		if err := tx.
			Where(
				"team_id = ? AND id <> ? AND is_active = true",
				author.TeamID, author.ID,
			).
			Find(&teamMembers).Error; err != nil {
			return fmt.Errorf("%s: failed to get team members: %w", op, err)
		}

		// lookup for pr's reviewers
		var currentReviewers []models.PRReviewer
		if err := tx.Where("pr_id = ?", prID).Find(&currentReviewers).Error; err != nil {
			return fmt.Errorf("%s: failed to get reviewers: %w", op, err)
		}

		// set of current reviewers id's
		excluded := map[uint]bool{}
		for _, r := range currentReviewers {
			excluded[r.ReviewerID] = true
		}

		// choosing all candidates
		var candidates []models.User
		for _, u := range teamMembers {
			if !excluded[u.ID] {
				candidates = append(candidates, u)
			}
		}

		if len(candidates) == 0 {
			return fmt.Errorf("%s: no available reviewers for reassignment", op)
		}

		// new random reviewer
		newReviewer := candidates[rand.Intn(len(candidates))]

		// update
		if err := tx.Model(&currentReviewer).Update("reviewer_id", newReviewer.ID).Error; err != nil {
			return fmt.Errorf("%s: failed to update reviewer: %w", op, err)
		}

		return nil
	})
	if err != nil {
		return domain.PullRequest{}, nil, 0, err
	}

	domainPR := mapper.PullRequestModelToDomain(pr)

	reviewers, err := s.GetUserReviewersByPRID(prID)
	if err != nil {
		return domain.PullRequest{}, nil, 0, err
	}

	newID := currentReviewer.ReviewerID

	return domainPR, reviewers, newID, nil
}

func shuffleSlice[T any](slice []T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := len(slice) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}
