package mapper

import (
	"errors"
	"fmt"

	"dang.z.v.task/internal/service"
	"dang.z.v.task/internal/storage"
)

// Returns business errors from storage errors
func MapStorageError(err error) error {
	if err == nil {
		return nil
	}

	// base errors
	switch {
	case errors.Is(err, storage.ErrNotFound):
		return fmt.Errorf("%w: %v", service.ErrNotFound, err)
	case errors.Is(err, storage.ErrAlreadyExists):
		return fmt.Errorf("%w: %v", service.ErrAlreadyExists, err)
	}

	// logic errors
	switch {
	case errors.Is(err, storage.ErrPullRequestMerged):
		return fmt.Errorf("%w: %v", service.ErrPullRequestMerged, err)
	case errors.Is(err, storage.ErrNoReviewersAvailable):
		return fmt.Errorf("%w: %v", service.ErrNoReviewersAvailable, err)
	case errors.Is(err, storage.ErrReviewerNotAssigned):
		return fmt.Errorf("%w: %v", service.ErrReviewerNotAssigned, err)
	case errors.Is(err, storage.ErrPRNotFound):
		return fmt.Errorf("%w: %v", service.ErrPRNotFound, err)
	case errors.Is(err, storage.ErrAuthorNotFound):
		return fmt.Errorf("%w: %v", service.ErrAuthorNotFound, err)
	case errors.Is(err, storage.ErrTeamExists):
		return fmt.Errorf("%w: %v", service.ErrTeamExists, err)
	}

	// db constraints
	switch {
	case errors.Is(err, storage.ErrDuplicateKey):
		return fmt.Errorf("%w: %v", service.ErrDuplicateKey, err)
	case errors.Is(err, storage.ErrForeignKeyViolation):
		return fmt.Errorf("%w: %v", service.ErrForeignKeyViolation, err)
	case errors.Is(err, storage.ErrCheckConstraint):
		return fmt.Errorf("%w: %v", service.ErrCheckConstraint, err)
	case errors.Is(err, storage.ErrNotNullConstraint):
		return fmt.Errorf("%w: %v", service.ErrNotNullConstraint, err)
	case errors.Is(err, storage.ErrUniqueConstraint):
		return fmt.Errorf("%w: %v", service.ErrUniqueConstraint, err)
	}

	// fallback
	return fmt.Errorf("%w: %v", service.ErrInternalServer, err)
}
