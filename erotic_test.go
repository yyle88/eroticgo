package eroticgo_test

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/eroticgo"
)

func TestExample(t *testing.T) {
	require.Error(t, errors.New("wrong"))
}

var colors = []eroticgo.COLOR{
	eroticgo.BLACK,
	eroticgo.RED,
	eroticgo.GREEN,
	eroticgo.YELLOW,
	eroticgo.BLUE,
	eroticgo.MAGENTA,
	eroticgo.CYAN,
	eroticgo.WHITE,
	eroticgo.GRAY,
	eroticgo.SCARLET,
	eroticgo.LIME,
	eroticgo.AMBER,
	eroticgo.AZURE,
	eroticgo.PINK,
	eroticgo.AQUA,
	eroticgo.IVORY,
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

func TestNewCOLOR(t *testing.T) {
	t.Run("case-1", func(t *testing.T) {
		co := eroticgo.NewCOLOR("\033[102m%s\033[0m")
		t.Log(co.Sprint("a b c"))
	})
	t.Run("case-2", func(t *testing.T) {
		co := eroticgo.NewCOLOR("\033[42m%s\033[0m")
		t.Log(co.Sprint("a b c"))
	})
	t.Run("case-2", func(t *testing.T) {
		co := eroticgo.NewCOLOR("\033[32m%s\033[0m")
		t.Log(co.Sprint("a b c"))
	})
}
