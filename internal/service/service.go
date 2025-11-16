package service

import "errors"

var (
	ErrNotFound       = errors.New("record not found")
	ErrAlreadyExists  = errors.New("already exists")
	ErrInternalServer = errors.New("internal server error")
)

var (
	ErrDuplicateKey        = errors.New("duplicate key violation")
	ErrForeignKeyViolation = errors.New("foreign key violation")
	ErrCheckConstraint     = errors.New("check constraint violation")
	ErrNotNullConstraint   = errors.New("not null constraint violation")
	ErrUniqueConstraint    = errors.New("unique constraint violation")
)

var (
	ErrPullRequestMerged    = errors.New("pull request is already merged")
	ErrTeamExists           = errors.New("team with this name already exists")
	ErrPullRequestExists    = errors.New("pull request already exists")
	ErrReassignOnMergedPR   = errors.New("cannot reassign on merged PR")
	ErrNoReviewersAvailable = errors.New("no reviewers available for reassignment")
	ErrReviewerNotAssigned  = errors.New("reviewer is not assigned to this PR")
	ErrPRNotFound           = errors.New("pull request not found")
	ErrAuthorNotFound       = errors.New("author not found")
)
