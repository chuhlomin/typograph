package typograph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNBSP(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			"Повторять, пока процесс не свернётся в навык.",
			"Повторять, пока процесс не\u00A0свернётся в\u00A0навык.",
		},
	}

	n := NBSP{}

	for _, test := range tests {
		actual := n.Process([]byte(test.in))
		require.Equal(t, test.out, string(actual))
	}
}
