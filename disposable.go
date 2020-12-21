//go:generate go run ./cmd/parser/ -pkg=disposable -var=disposable -out=disposable_list.go -url=https://raw.githubusercontent.com/ivolo/disposable-email-domains/master/index.json
//go:generate go run ./cmd/parser/ -pkg=disposable -var=wildcards -out=disposable_list_wildcard.go -url=https://raw.githubusercontent.com/ivolo/disposable-email-domains/master/wildcard.json

// Package disposable can be used to detect if a domain or email address is
// potentially fake or temporary based on it being registered using a
// disposable email provider (identified by the domain name).
package disposable

import "strings"

var mapping map[string]struct{}

func init() {
	mapping = make(map[string]struct{}, len(disposable))
	for _, domain := range disposable {
		mapping[domain] = struct{}{}
	}
}

// IsDomainDisposable determines if a specified domain is disposable based
// on an exact match check against a list of domains known for creating
// diposable email addresses.
func IsDomainDisposable(domain string) bool {
	_, found := mapping[domain]
	return found
}

// IsDomainWildcard determines if a specified domain is disposable based on if
// it is a subdomain of a known wildcard domain used to create disposable
// email addresses.
func IsDomainWildcard(domain string) bool {
	for _, wildcard := range wildcards {
		if strings.HasSuffix(domain, wildcard) {
			return true
		}
	}
	return false
}

// IsEmailAddressDisposable parses a valid email address and checks the domain
// part against known disposable domains. If an invalid email is passed,
// correct behavior can not be guaranteed.
func IsEmailAddressDisposable(email string) bool {
	parts := strings.SplitN(email, "@", 2)
	if len(parts) != 2 {
		return false
	}
	domain := parts[1]
	return IsDomainDisposable(domain) || IsDomainWildcard(domain)
}
