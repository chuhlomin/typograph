package main

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
	}

	s := Space{}

	for _, test := range tests {
		actual := s.Process([]byte(test.in))
		require.Equal(t, test.out, string(actual))
	}
}