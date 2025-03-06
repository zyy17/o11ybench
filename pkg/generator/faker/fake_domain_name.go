package faker

import (
	"github.com/brianvoe/gofakeit"
)

// FakeDomainNameOptions is the options for generating the fake domain name.
type FakeDomainNameOptions struct {
	// TODO(zyy17)
}

// FakeDomainName generates a fake domain name.
func FakeDomainName(opts Options) (string, error) {
	var options FakeDomainNameOptions
	if err := parseOptions(opts, &options); err != nil {
		return "", err
	}

	return gofakeit.DomainName(), nil
}
