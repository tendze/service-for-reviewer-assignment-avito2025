package mapper

import (
	"dang.z.v.task/internal/domain"
	"dang.z.v.task/internal/storage/postgresql/models"
)

func PRReviewerModelToDomain(reviewer models.PRReviewer) domain.PRReviewer {
	return domain.PRReviewer{
		ID:         reviewer.ID,
		PRID:       reviewer.PRID,
		ReviewerID: reviewer.ReviewerID,
		CreatedAt:  reviewer.CreatedAt,
	}
}

func PRReviewerDomainToModel(d domain.PRReviewer) models.PRReviewer {
	return models.PRReviewer{
		ID:         d.ID,
		PRID:       d.PRID,
		ReviewerID: d.ReviewerID,
		CreatedAt:  d.CreatedAt,
	}
}
