package faker

import (
	"github.com/brianvoe/gofakeit"

	"github.com/zyy17/o11ybench/pkg/generator/common"
)

// FakeDomainNameOptions is the options for generating the fake domain name.
type FakeDomainNameOptions struct {
	// TODO(zyy17)
}

// FakeDomainName generates a fake domain name.
func FakeDomainName(_ common.ElementType, opts Options) (string, error) {
	var options FakeDomainNameOptions
	if err := parseOptions(opts, &options); err != nil {
		return "", err
	}

	return gofakeit.DomainName(), nil
}
