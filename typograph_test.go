package typograph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTypograph(t *testing.T) {
	tests := []struct {
		name, in, out string
	}{
		{"symbols", "10(tm) -> 30 C", "10™ → 30 °C"},
		{"space", "One,two.Three", "One, two. Three"},
		{"nbsp", "Me, me again and Irene", "Me, me\u00a0again and Irene"},
	}

	typograph := NewTypograph()

	for _, test := range tests {
		actual := typograph.Process([]byte(test.in))
		require.Equal(t, test.out, string(actual), test.name)
	}
}
