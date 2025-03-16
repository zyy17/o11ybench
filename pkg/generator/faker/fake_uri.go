package faker

import (
	"net/url"
	"strings"

	"github.com/brianvoe/gofakeit"

	"github.com/zyy17/o11ybench/pkg/generator/common"
)

// FakeURIOptions is the options for the fake URI.
type FakeURIOptions struct {
	// URL indicates whether to generate a full URL.
	URL bool `yaml:"url"`
}

// FakeURI generates a fake URI.
func FakeURI(_ common.ElementType, opts Options) (string, error) {
	var options FakeURIOptions
	if err := parseOptions(opts, &options); err != nil {
		return "", err
	}

	if options.URL {
		return gofakeit.URL(), nil
	}

	return randomResourceURI(), nil
}

func randomResourceURI() string {
	var uri string
	num := gofakeit.Number(1, 4)
	for i := 0; i < num; i++ {
		uri += "/" + url.QueryEscape(gofakeit.BS())
	}
	uri = strings.ToLower(uri)
	return uri
}
