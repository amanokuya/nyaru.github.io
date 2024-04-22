package lexier

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTokenName(t *testing.T) {
	id_token := NewToken(ID)
	require.Equal(t, "identifier", id_token.ToString())
}
