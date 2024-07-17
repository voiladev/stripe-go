package connectiontoken

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/voiladev/stripe-go/v79"
	_ "github.com/voiladev/stripe-go/v79/testing"
)

func TestTerminalConnectionTokenNew(t *testing.T) {
	connectiontoken, err := New(&stripe.TerminalConnectionTokenParams{})
	assert.Nil(t, err)
	assert.NotNil(t, connectiontoken)
}
