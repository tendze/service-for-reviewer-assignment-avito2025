package mapper

import (
	"errors"
	"fmt"
	"strings"

	"dang.z.v.task/internal/storage"
	"gorm.io/gorm"
)

func MapPostgresError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("%w: %v", storage.ErrNotFound, err)
	}

	msg := err.Error()

	switch {
	case strings.Contains(msg, "SQLSTATE 23505"):
		return fmt.Errorf("%w: %v", storage.ErrUniqueConstraint, err)
	case strings.Contains(msg, "SQLSTATE 23503"):
		return fmt.Errorf("%w: %v", storage.ErrForeignKeyViolation, err)
	case strings.Contains(msg, "SQLSTATE 23502"):
		return fmt.Errorf("%w: %v", storage.ErrNotNullConstraint, err)
	case strings.Contains(msg, "SQLSTATE 23514"):
		return fmt.Errorf("%w: %v", storage.ErrCheckConstraint, err)
	}

	return fmt.Errorf("%w: %v", storage.ErrInternalServer, err)
}
