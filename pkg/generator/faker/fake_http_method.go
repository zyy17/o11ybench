package faker

import (
	"github.com/brianvoe/gofakeit"
)

// FakeHTTPMethodOptions is the options for generating the fake http method.
type FakeHTTPMethodOptions struct {
	// TODO(zyy17)
}

// FakeHTTPMethod generates a fake HTTP method.
func FakeHTTPMethod(opts Options) (string, error) {
	var options FakeHTTPMethodOptions
	if err := parseOptions(opts, &options); err != nil {
		return "", err
	}

	return gofakeit.HTTPMethod(), nil
}
