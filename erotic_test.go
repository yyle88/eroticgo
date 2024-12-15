package eroticgo

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	require.Error(t, errors.New("wrong"))
}

var colors = []COLOR{
	BLACK, RED, GREEN, YELLOW, BLUE, MAGENTA, CYAN, WHITE,
}

func TestCOLOR_Sprint(t *testing.T) {
	args := []any{"a", "b", "c", "\n", "x", "y", "z", "\n", 1, 2, 3}
	for _, co := range colors {
		t.Log(co.Sprint(args...))
	}
}

func TestCOLOR_ShowMessage(t *testing.T) {
	args := []any{"a", "b", "c", "\n", "x", "y", "z", "\n", 1, 2, 3}
	for _, co := range colors {
		co.ShowMessage(args...)
	}
}
