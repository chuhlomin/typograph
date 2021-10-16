package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSymbols(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// arrows
		{input: "->", expected: "→"},
		{input: "<-", expected: "←"},
		{input: "20 + 10 -> 30", expected: "20 + 10 → 30"},
		{input: "20 + 10 <- 30", expected: "20 + 10 ← 30"},
		// copy
		{input: "(c)", expected: "©"},
		{input: "(с)", expected: "©"},
		{input: "Copyright (с)", expected: "©"},
		{input: "copyright (с)", expected: "©"},
		{input: "(r)", expected: "®"},
		// tm
		{input: "(tm)", expected: "™"},
		// cf
		{input: " 200 C", expected: " 200 °C"},
		{input: "200 C", expected: "200 °C"},
		{input: "&minus;20 C", expected: "&minus;20 °C"},
		{input: "-20 C", expected: "-20 °C"},
		{input: "+10 C", expected: "+10 °C"},
		{input: "±2,4 C", expected: "±2,4 °C"},
		{input: "+1.5 C", expected: "+1.5 °C"},
		{input: "1,5 C", expected: "1,5 °C"},
		{input: "≈99 C", expected: "≈99 °C"},
		{input: "B2C", expected: "B2C"},
		{input: " 200 C.", expected: " 200 °C."},
		{input: " 20d C", expected: " 20d C"},
		{input: " 20 C1", expected: " 20 C1"},
		{input: " 200 F", expected: " 200 °F"},
		{input: "200 F", expected: "200 °F"},
		{input: "200 C\n 300 F", expected: "200 °C\n 300 °F"},
		{input: "200 C\n300 F", expected: "200 °C\n300 °F"},
	}

	for _, test := range tests {
		actual := Symbols(test.input)
		require.Equal(t, test.expected, actual)
	}
}
