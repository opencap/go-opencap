package opencap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHost(t *testing.T) {
	url, err := GetHost("ogdolo.com")
	assert.Nil(t, err)
	assert.Equal(t, "opencap.ogdolo.com", url)
}
