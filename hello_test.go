package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHello(t *testing.T) {
	mes, err := Hello("maoz")
	require.NoError(t, err)
	expected := "Hello, maoz"
	require.Equal(t, expected, mes)
}
