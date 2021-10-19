package typograph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpace(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			"Солнце садилось за горизонт,и поднялся ветер. Вот.",
			"Солнце садилось за горизонт, и поднялся ветер. Вот.",
		},
		{
			"Солнце садилось за горизонт,и поднялся ветер!Вот.",
			"Солнце садилось за горизонт, и поднялся ветер! Вот.",
		},
		{
			"Солнце садилось за горизонт,и поднялся ветер?Вот.",
			"Солнце садилось за горизонт, и поднялся ветер? Вот.",
		},
		{
			"Солнце садилось за горизонт,и поднялся ветер??Вот.",
			"Солнце садилось за горизонт, и поднялся ветер?? Вот.",
		},
		{
			"Солнце садилось за горизонт,?",
			"Солнце садилось за горизонт,?",
		},
		{
			"Солнце садилось за горизонт1,и поднялся ветер?",
			"Солнце садилось за горизонт1, и поднялся ветер?",
		},
		{
			"«Я!»",
			"«Я!»",
		},
		{
			"‹I!›",
			"‹I!›",
		},
		{
			"https://tools.ietf.org/html/rfc7208",
			"https://tools.ietf.org/html/rfc7208",
		},
		{
			"![Title](path.png \"Title\")",
			"![Title](path.png \"Title\")",
		},
	}

	s := Space{}

	for _, test := range tests {
		actual := s.Process([]byte(test.in))
		require.Equal(t, test.out, string(actual))
	}
}
