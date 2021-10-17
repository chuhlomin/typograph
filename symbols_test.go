package typograph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSymbols(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"->", "→"},
		{"<-", "←"},
		{"20 + 10 -> 30", "20 + 10 → 30"},
		{"20 + 10 <- 30", "20 + 10 ← 30"},
		// copy
		{"(c)", "©"},
		{"(с)", "©"},
		{"Copyright (с)", "©"},
		{"copyright (с)", "©"},
		{"(r)", "®"},
		// tm
		{"(tm)", "™"},
		// cf
		{" 200 C", " 200 °C"},
		{"200 C", "200 °C"},
		{"&minus;20 C", "&minus;20 °C"},
		{"-20 C", "-20 °C"},
		{"+10 C", "+10 °C"},
		{"±2,4 C", "±2,4 °C"},
		{"+1.5 C", "+1.5 °C"},
		{"1,5 C", "1,5 °C"},
		{"≈99 C", "≈99 °C"},
		{"B2C", "B2C"},
		{" 200 C.", " 200 °C."},
		{" 20d C", " 20d C"},
		{" 20 C1", " 20 C1"},
		{" 200 F", " 200 °F"},
		{"200 F", "200 °F"},
		{"200 C\n 300 F", "200 °C\n 300 °F"},
		{"200 C\n300 F", "200 °C\n300 °F"},
	}

	s := Symbols{}

	for _, test := range tests {
		actual := s.Process([]byte(test.in))
		require.Equal(t, test.out, string(actual))
	}
}
