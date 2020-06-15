package utils

import (
	"testing"
)

type Query struct {
	X        int64 `form:"x" validate:"required"`
	Y        int64 `form:"y" validate:"required"`
	Distance int64 `form:"distance" validate:"required"`
}

func TestAbs(t *testing.T) {
	res := Abs((-5) + (-5))
	if res != 10 {
		t.Errorf("Expected to return a positive number")
	}
}

func TestValidation(t *testing.T) {
	var query Query
	err := Validation(query)
	if err == nil {
		t.Errorf("Expected to return an error if struct fields are empty")
	}
}
