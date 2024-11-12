package eroticgo

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	require.Error(t, errors.New("wa"))
}

func TestCOLOR_Sprint(t *testing.T) {
	args := []any{"a", "b", "c", "\n", "x", "y", "z", "\n", 1, 2, 3}
	t.Log(BLACK.Sprint(args...))
	t.Log(RED.Sprint(args...))
	t.Log(GREEN.Sprint(args...))
	t.Log(YELLOW.Sprint(args...))
	t.Log(BLUE.Sprint(args...))
	t.Log(MAGENTA.Sprint(args...))
	t.Log(CYAN.Sprint(args...))
	t.Log(WHITE.Sprint(args...))
}
