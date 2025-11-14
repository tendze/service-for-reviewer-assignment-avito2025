package mapper

import (
	"dang.z.v.task/internal/domain"
	"dang.z.v.task/internal/storage/postgresql/models"
)

func PullRequestDomainToModel(pr domain.PullRequest) models.PullRequest {
	return models.PullRequest{
		ID:        pr.ID,
		Title:     pr.Title,
		AuthorID:  pr.AuthorID,
		Status:    pr.Status,
		CreatedAt: pr.CreatedAt,
		MergedAt:  pr.MergedAt,
	}
}

func PullRequestModelToDomain(pr models.PullRequest) domain.PullRequest {
	return domain.PullRequest{
		ID:        pr.ID,
		Title:     pr.Title,
		AuthorID:  pr.AuthorID,
		Status:    pr.Status,
		CreatedAt: pr.CreatedAt,
		MergedAt:  pr.MergedAt,
	}
}
