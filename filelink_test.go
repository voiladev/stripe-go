package stripe

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/voiladev/stripe-go/v79/form"
)

func TestFileLinkParams_AppendTo(t *testing.T) {
	{
		params := &FileLinkParams{ExpiresAtNow: Bool(true)}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"now"}, body.Get("expires_at"))
	}
}
