package nano

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateAddressSuccess(t *testing.T) {
	addr := "xrb_1ipdmx5ck6fwc1hkt3fg631z6twjo9e5m6e9dbenm65sbb3e4hgmd1ygxnq3"
	err := ValidateAddress(addr)
	assert.Nil(t, err)

	addr = "xrb_1qckwc5o3obkrwbet4amnkya113xq77qpaknsmiq9hwq31tmd5bpyo7sepsw"
	err = ValidateAddress(addr)
	assert.Nil(t, err)

	addr = "nano_1os1zf6ne5ydfmugdnqhwqbgze3p6f6x4k7fsz7epadknw7fae7qmihkaj8i"
	err = ValidateAddress(addr)
	assert.Nil(t, err)
}

func TestValidateAddressFailure(t *testing.T) {
	addr := "xrb_1ipdmx5ck6fwc1hkt3fg631z6twjo9e5m6e9dbe65sbb3e4hgmd1ygxnq3"
	err := ValidateAddress(addr)
	assert.NotNil(t, err)

	addr = "xrb_1qckwc5o3obkrwbet4amnkya113xq77qpaknsmiq9hwq31tmd5bpyo7sepswdf"
	err = ValidateAddress(addr)
	assert.NotNil(t, err)

	addr = "nan_1os1zf6ne5ydfmugdnqhwqbgze3p6f6x4k7fsz7epadknw7fae7qmihkaj8i"
	err = ValidateAddress(addr)
	assert.NotNil(t, err)
}
