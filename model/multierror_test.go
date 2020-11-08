package model_test

import (
	"errors"
	"testing"

	"github.com/alextanhongpin/go-graphql-template/model"
)

func TestMultiError_InitEmpty(t *testing.T) {
	merr := model.NewMultiError()
	if exp, got := "", merr.Error(); exp != got {
		t.Fatalf("expected %s, got %s", exp, got)
	}
}

func TestMultiError_InitWithValue(t *testing.T) {
	merr := model.NewMultiError(errors.New("bad request"))
	if exp, got := "bad request", merr.Error(); exp != got {
		t.Fatalf("expected %s, got %s", exp, got)
	}
}

func TestMultiError_Add(t *testing.T) {
	merr := model.NewMultiError()
	merr.Add(nil)
	merr.Add(errors.New("hello"))
	merr.Add(errors.New("world"))
	if exp, got := "hello\nworld", merr.Error(); exp != got {
		t.Fatalf("expected %s, got %s", exp, got)
	}
}

func TestMultiError_HasError(t *testing.T) {
	merr := model.NewMultiError()
	if exp, got := false, merr.HasError(); exp != got {
		t.Fatalf("expected %t, got %t", exp, got)
	}

	merr.Add(errors.New("hello"))
	if exp, got := true, merr.HasError(); exp != got {
		t.Fatalf("expected %t, got %t", exp, got)
	}
}

func TestMultiError_Extensions(t *testing.T) {
	merr := model.NewMultiError()
	if exp, got := false, merr.HasError(); exp != got {
		t.Fatalf("expected %t, got %t", exp, got)
	}

	ext := merr.Extensions()
	if exp, got := model.CodeBadUserInput, ext["code"]; exp != got {
		t.Fatalf("expected %s, got %s", exp, got)
	}
}
