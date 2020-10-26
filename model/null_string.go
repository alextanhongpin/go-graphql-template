package model

import "database/sql"

// NewNullString returns a new sql.NullString.
func NewNullString(s string) sql.NullString {
	return sql.NullString{
		Valid:  len(s) > 0,
		String: s,
	}
}
