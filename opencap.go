package opencap

import (
	"errors"
	"net"
)

// GetHost returns the highest priority opencap host URL at a given
// domain name.
func GetHost(domain string) (string, error) {
	_, addresses, err := net.LookupSRV("opencap", "tcp", domain)
	if err != nil {
		return "", err
	}
	if len(addresses) < 1 {
		return "", errors.New("No addresses found")
	}

	target := addresses[0].Target

	// strip trailing period for convenience
	if len(target) > 1 {
		if string(target[len(target)-1]) == "." {
			target = target[:len(target)-1]
		}
	}

	return target, nil
}
