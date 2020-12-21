package disposable_test

import (
	"testing"

	"github.com/hsson/go-disposable"
)

func TestDisposableDomain(t *testing.T) {
	table := map[string]bool{
		"0-180.com":     true,
		"gmail":         false,
		"hotmail":       false,
		"19292.monster": true,
		"quant2new.pro": true,
	}

	for domain, isDisposable := range table {
		res := disposable.IsDomainDisposable(domain)
		if res != isDisposable {
			t.Errorf("%s: got %v want %v", domain, res, isDisposable)
		}
	}
}

func TestDisposableWildcard(t *testing.T) {
	table := map[string]bool{
		"A.33mail.com":       true,
		"A.B.C.33mail.com":   true,
		"33mail.com":         true,
		"not.disposable.com": false,
	}

	for domain, isDisposable := range table {
		res := disposable.IsDomainWildcard(domain)
		if res != isDisposable {
			t.Errorf("%s: got %v want %v", domain, res, isDisposable)
		}
	}
}

func TestDisposableEmail(t *testing.T) {
	table := map[string]bool{
		"test@gmail.com":                   false,
		"disposable@quant2new.pro":         true,
		"foo@33mail.com":                   true,
		"33mail.com@some.stuff.33mail.com": true,
		"legit@hotmail.com":                false,
	}
	for domain, isDisposable := range table {
		res := disposable.IsEmailAddressDisposable(domain)
		if res != isDisposable {
			t.Errorf("%s: got %v want %v", domain, res, isDisposable)
		}
	}
}

func TestDisposableEmailShouldNotPanicOnInvalid(t *testing.T) {
	cases := []string{
		"33mail.com",
		"   ",
		"@ @@ asd@@__",
		"",
	}
	for _, invalid := range cases {
		disposable.IsEmailAddressDisposable(invalid)
	}
}
