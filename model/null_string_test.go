package model_test

import (
	"testing"

	"github.com/alextanhongpin/graphql-server-starter/model"
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
