package mapper

import (
	"errors"
	"net/http"

	"dang.z.v.task/internal/handlers/response"
	"dang.z.v.task/internal/service"
)

// Returns HTTP code by business layers error message
func HTTPStatusFromError(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch {
	// not Found
	case errors.Is(err, service.ErrNotFound),
		errors.Is(err, service.ErrPRNotFound),
		errors.Is(err, service.ErrAuthorNotFound),
		errors.Is(err, service.ErrReviewerNotAssigned):
		return http.StatusNotFound

	// already exists / conflict
	case errors.Is(err, service.ErrAlreadyExists),
		errors.Is(err, service.ErrTeamExists),
		errors.Is(err, service.ErrPullRequestExists),
		errors.Is(err, service.ErrDuplicateKey),
		errors.Is(err, service.ErrUniqueConstraint):
		return http.StatusConflict // 409

	// bad Request / domain violation
	case errors.Is(err, service.ErrForeignKeyViolation),
		errors.Is(err, service.ErrCheckConstraint),
		errors.Is(err, service.ErrNotNullConstraint),
		errors.Is(err, service.ErrReassignOnMergedPR),
		errors.Is(err, service.ErrPullRequestMerged),
		errors.Is(err, service.ErrNoReviewersAvailable):
		return http.StatusBadRequest // 400

	// internal
	case errors.Is(err, service.ErrInternalServer):
		return http.StatusInternalServerError // 500
	}

	return http.StatusInternalServerError
}

// ErrorMessageFromError returns error message by business layer error
func ErrorMessageFromError(err error) string {
	if err == nil {
		return ""
	}

	switch {
	// not found
	case errors.Is(err, service.ErrNotFound):
		return "resource not found"
	case errors.Is(err, service.ErrPRNotFound):
		return "pull request not found"
	case errors.Is(err, service.ErrAuthorNotFound):
		return "author not found"
	case errors.Is(err, service.ErrReviewerNotAssigned):
		return "reviewer is not assigned to this PR"

	// already exists / conflict
	case errors.Is(err, service.ErrAlreadyExists):
		return "resource already exists"
	case errors.Is(err, service.ErrTeamExists):
		return "team_name already exists"
	case errors.Is(err, service.ErrPullRequestExists):
		return "pull request already exists"
	case errors.Is(err, service.ErrDuplicateKey),
		errors.Is(err, service.ErrUniqueConstraint):
		return "duplicate key violation"

	// bad Request / domain violation
	case errors.Is(err, service.ErrForeignKeyViolation):
		return "foreign key violation"
	case errors.Is(err, service.ErrCheckConstraint):
		return "check constraint violation"
	case errors.Is(err, service.ErrNotNullConstraint):
		return "not null constraint violation"
	case errors.Is(err, service.ErrReassignOnMergedPR):
		return "cannot reassign on merged PR"
	case errors.Is(err, service.ErrPullRequestMerged):
		return "pull request is already merged"
	case errors.Is(err, service.ErrNoReviewersAvailable):
		return "no reviewers available for reassignment"
	}

	return "internal server error"
}

func ErrorCodeMessageFromError(err error) string {
	if err == nil {
		return response.SUCCESS
	}

	switch {
	// not Found
	case errors.Is(err, service.ErrNotFound),
		errors.Is(err, service.ErrPRNotFound),
		errors.Is(err, service.ErrAuthorNotFound):
		return response.NOT_FOUND

	// already exists / conflict
	case errors.Is(err, service.ErrAlreadyExists),
		errors.Is(err, service.ErrTeamExists):
		return response.TEAM_EXISTS
	case errors.Is(err, service.ErrPullRequestExists):
		return response.PR_EXISTS
	case errors.Is(err, service.ErrPullRequestMerged):
		return response.PR_MERGED

	// reassign errors
	case errors.Is(err, service.ErrReviewerNotAssigned):
		return response.NOT_ASSIGNED
	case errors.Is(err, service.ErrNoReviewersAvailable):
		return response.NO_CANDIDATE
	case errors.Is(err, service.ErrReassignOnMergedPR):
		return response.PR_MERGED

	// bad request / constraints
	case errors.Is(err, service.ErrDuplicateKey),
		errors.Is(err, service.ErrUniqueConstraint),
		errors.Is(err, service.ErrForeignKeyViolation),
		errors.Is(err, service.ErrCheckConstraint),
		errors.Is(err, service.ErrNotNullConstraint):
		return response.BAD_REQUEST

	case errors.Is(err, service.ErrInternalServer):
		return response.INTERNAL_SERVER_ERRROR
	}

	return response.INTERNAL_SERVER_ERRROR
}
