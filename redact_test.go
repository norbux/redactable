package redactable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedacted(t *testing.T) {
		var rs RedactableString = "This is a :secret:, friends."
		redacted := rs.Redact()

		assert.Equal(t, "This is a *****, friends.", redacted, "Wrong redacted message")
}

func TestUnredacted(t *testing.T) {
	var rs RedactableString = "This is a :secret:, friends."
	unredacted := rs.Sprint()

	assert.Equal(t, "This is a secret, friends.", unredacted, "Wrong unredacted message")
}