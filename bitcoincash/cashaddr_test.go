package bitcoincash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeValid(t *testing.T) {
	err := Validate("bitcoincash:qpm2qsznhks23z7629mms6s4cwef74vcwvy22gdx6a")
	assert.Nil(t, err)

	err = Validate("bitcoincash:qr95sy3j9xwd2ap32xkykttr4cvcu7as4y0qverfuy")
	assert.Nil(t, err)

	err = Validate("bitcoincash:qqq3728yw0y47sqn6l2na30mcw6zm78dzqre909m2r")
	assert.Nil(t, err)

	err = Validate("bitcoincash:ppm2qsznhks23z7629mms6s4cwef74vcwvn0h829pq")
	assert.Nil(t, err)

	err = Validate("bitcoincash:pr95sy3j9xwd2ap32xkykttr4cvcu7as4yc93ky28e")
	assert.Nil(t, err)

	err = Validate("bitcoincash:pqq3728yw0y47sqn6l2na30mcw6zm78dzq5ucqzc37")
	assert.Nil(t, err)
}

func TestDecodeInvalid(t *testing.T) {
	err := Validate("bitcoincsh:qpm2qsznhks23z7629mms6s4cwef74vcwvy22gdx6a")
	assert.NotNil(t, err)

	err = Validate("bitcoincash:3CWFddi6m4ndiGyKqzYvsFYagqDLPVMTzC")
	assert.NotNil(t, err)

	err = Validate("bitcoincash:1KXrWXciRDZUpQwQmuM1DbwsKDLYAYsVLR")
	assert.NotNil(t, err)

	err = Validate("bitcoincash:")
	assert.NotNil(t, err)

	err = Validate(":pqq3728yw0y47sqn6l2na30mcw6zm78dzq5ucqzc37")
	assert.NotNil(t, err)
}
