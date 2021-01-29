package model_test

import (
	"testing"
	"time"

	"github.com/alextanhongpin/go-graphql-template/domain/model"
	"github.com/google/uuid"
)

func TestNewCursor(t *testing.T) {
	cursor := model.NewCursor()
	if exp, got := model.MaxUUID, cursor.ID; exp != got {
		t.Fatalf("expected %s, got %s", exp, got)
	}
}

func TestCursorToAndFromBase64(t *testing.T) {
	cursor := model.Cursor{
		CreatedAt: time.Now(),
		ID:        uuid.New(),
	}
	b64, err := cursor.ToBase64()
	if err != nil {
		t.Fatal(err)
	}
	newCursor := model.NewCursor()
	if err != nil {
		t.Fatal(err)
	}
	err = newCursor.FromBase64(b64)
	if err != nil {
		t.Fatal(err)
	}
	if exp, got := cursor.ID, newCursor.ID; exp != got {
		t.Fatalf("expected %s, got %s", exp, got)
	}
}
