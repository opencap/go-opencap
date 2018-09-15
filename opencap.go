package opencap

import (
	"errors"
	"net"
	"strings"
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

// ValidateAlias returns the username, domain, and error
func ValidateAlias(alias string) (string, string, error) {
	split := strings.Split(alias, "$")
	if len(split) != 2 {
		return "", "", errors.New("Invalid seperator use in alias")
	}
	username := split[0]
	domain := split[1]
	if len(username) < 1 {
		return "", "", errors.New("No username in alias")
	}
	if len(domain) < 1 {
		return "", "", errors.New("No domain in alias")
	}

	domainParts := strings.Split(domain, ".")
	if len(domainParts) != 2 {
		return "", "", errors.New("Invalid domain in alias")
	}
	domainName := domainParts[0]
	ext := domainParts[1]
	if len(domainName) < 1 {
		return "", "", errors.New("Invalid domain in alias")
	}
	if len(ext) < 1 {
		return "", "", errors.New("Invalid domain in alias")
	}
	return username, domain, nil
}
