package model

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

var MaxUUID uuid.UUID

func init() {
	MaxUUID = uuid.MustParse("ffffffff-ffff-ffff-ffff-ffffffffffff")
}

// Cursor represents a cursor by id and created at timestamp.
type Cursor struct {
	CreatedAt time.Time `json:"ts"`
	ID        uuid.UUID `json:"id"`
}

// FromBase64 unmarshals a base64 string to Cursor.
func (c *Cursor) FromBase64(b64 string) (err error) {
	cursor, err := base64.URLEncoding.DecodeString(b64)
	if err != nil {
		return err
	}
	return json.Unmarshal(cursor, c)
}

// ToBase64 returns a base64 cursor string.
func (c *Cursor) ToBase64() (string, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// NewCursor returns a pointer to Cursor.
func NewCursor() *Cursor {
	return &Cursor{
		CreatedAt: time.Now(),
		ID:        MaxUUID,
	}
}
