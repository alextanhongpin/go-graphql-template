package model_test

import (
	"testing"

	"github.com/alextanhongpin/go-graphql-template/model"
)

func TestNullStringFactory(t *testing.T) {
	ns := model.NewNullString("")
	if ns.Valid == true {
		t.Fatal("expected false, got true")
	}
	if ns.String != "" {
		t.Fatalf(`expected "", got %s`, ns.String)
	}
}
