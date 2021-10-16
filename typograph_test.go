package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTypograph(t *testing.T) {
	tests := []struct {
		name, in, out string
	}{
		{"symbols", "10(tm) -> 30 C", "10™ → 30 °C"},
	}

	typograph := NewTypograph()

	for _, test := range tests {
		actual := typograph.Process(test.in)
		require.Equal(t, test.out, actual, test.name)
	}
}
