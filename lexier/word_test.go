package lexier

import (
	"github.com/stretchr/testify/require"

	"testing"
)

func TestWordToken(t *testing.T) {
	word := NewWordToken("variable", ID)
	require.Equal(t, "variable", word.ToString())
	word_tag := word.Tag
	require.Equal(t, word_tag.ToString(), "identifier")
}
