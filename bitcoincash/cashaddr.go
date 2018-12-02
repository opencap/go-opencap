package bitcoincash

import (
	"errors"

	"github.com/cpacia/bchutil"
)

// Validate returns nil if the address is a valid cashaddr
func Validate(address string) error {
	prefix, _, err := bchutil.DecodeCashAddress(address)
	if err != nil {
		return err
	}
	if prefix != "bitcoincash" {
		return errors.New("Invalid prefix")
	}
	return nil
}
