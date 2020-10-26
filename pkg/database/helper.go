package database

import (
	"errors"
	"strings"

	"github.com/lib/pq"
)

// UniqueViolation return true if the column violates the unique constraint.
func UniqueViolation(err error, column string) bool {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		isUniqueViolation := pqErr.Code.Name() == "unique_violation"
		return isUniqueViolation && strings.Contains(pqErr.Detail, column)
	}
	return false
}
