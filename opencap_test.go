package opencap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHost(t *testing.T) {
	url, err := GetHost("ogdolo.com")
	assert.Nil(t, err)
	assert.Equal(t, "cap.ogdolo.com", url)
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

	alias := "lane$ogdolo.com"
	username, domain, err = ValidateAlias(alias)
	assert.Nil(t, err)
	assert.Equal(t, "lane", username)
	assert.Equal(t, "ogdolo.com", domain)

	alias = "ogdolo.com"
	domain, username, err = ValidateAlias(alias)
	assert.NotNil(t, err)

	alias = "lane@ogdolo@.com"
	domain, username, err = ValidateAlias(alias)
	assert.NotNil(t, err)
}

func TestValidateDomain(t *testing.T) {
	domain := "ogdolo.com"
	assert.True(t, ValidateDomain(domain))

	domain = "com"
	assert.False(t, ValidateDomain(domain))

	domain = "ogdolo.co"
	assert.True(t, ValidateDomain(domain))

	domain = "ogdolo.bump"
	assert.True(t, ValidateDomain(domain))

	domain = "ogdolo."
	assert.False(t, ValidateDomain(domain))

	domain = "ogdolo.b"
	assert.False(t, ValidateDomain(domain))
}

func TestValidateUsername(t *testing.T) {
	username := "lane.c.Wagner"
	newName, valid := ValidateUsername(username)
	assert.True(t, valid)
	assert.Equal(t, "lane.c.wagner", newName)

	username = "lane-c-wail"
	newName, valid = ValidateUsername(username)
	assert.True(t, valid)

	username = "ner@gail.com"
	newName, valid = ValidateUsername(username)
	assert.False(t, valid)

	username = "kvothe435245"
	newName, valid = ValidateUsername(username)
	assert.True(t, valid)

	username = "fdasdilfsudfgshghgjghfjjhghjghjgjhgjghjhjidfsbfibkfjk"
	newName, valid = ValidateUsername(username)
	assert.False(t, valid)

	username = ""
	newName, valid = ValidateUsername(username)
	assert.False(t, valid)
}
