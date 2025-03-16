package faker

import (
	"github.com/brianvoe/gofakeit"

	"github.com/zyy17/o11ybench/pkg/generator/common"
)

// FakeHTTPUserAgentOptions is the options for the fake HTTP user agent.
type FakeHTTPUserAgentOptions struct {
	// TODO(zyy17)
}

// FakeHTTPUserAgent generates a fake HTTP user agent.
func FakeHTTPUserAgent(_ common.ElementType, opts Options) (string, error) {
	var options FakeHTTPUserAgentOptions
	if err := parseOptions(opts, &options); err != nil {
		return "", err
	}

	return gofakeit.UserAgent(), nil
}
