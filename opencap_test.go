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

func TestValidateAlias(t *testing.T) {
	username, domain, err := ValidateAlias("donate$ogdolo.com")
	assert.Equal(t, "donate", username)
	assert.Equal(t, "ogdolo.com", domain)
	assert.Nil(t, err)

	_, _, err = ValidateAlias("donateogdolo.com")
	assert.NotNil(t, err)

	_, _, err = ValidateAlias("donateogd$olocom")
	assert.NotNil(t, err)

	_, _, err = ValidateAlias("donateogdolocom")
	assert.NotNil(t, err)

	_, _, err = ValidateAlias("donateogdolo$com.")
	assert.NotNil(t, err)
}
