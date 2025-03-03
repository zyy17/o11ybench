package templates

import (
	"github.com/zyy17/o11ybench/pkg/generator/faker"
)

// ReservedTokenName is the reserved token name for the fake data. It only take effect in specific format.
const (
	// ReservedTokenNameHost is the reserved token name for the host.
	// It only take effect in `FormatApacheCommonLog` and `FormatApacheCombinedLog`.
	ReservedTokenNameHost string = "host"

	// ReservedTokenNameUserID is the reserved token name for the user ID.
	// It only take effect in `FormatApacheCommonLog` and `FormatApacheCombinedLog`.
	ReservedTokenNameUserID string = "userID"

	// ReservedTokenNameTimestamp is the reserved token name for the timestamp.
	// It only take effect in `FormatApacheCommonLog` and `FormatApacheCombinedLog`.
	ReservedTokenNameTimestamp string = "timestamp"

	// ReservedTokenNameHTTPMethod is the reserved token name for the HTTP method.
	// It only take effect in `FormatApacheCommonLog` and `FormatApacheCombinedLog`.
	ReservedTokenNameHTTPMethod string = "httpMethod"

	// ReservedTokenNameHTTPVersion is the reserved token name for the HTTP version.
	// It only take effect in `FormatApacheCommonLog` and `FormatApacheCombinedLog`.
	ReservedTokenNameHTTPVersion string = "httpVersion"

	// ReservedTokenNameHTTPStatusCode is the reserved token name for the HTTP status code.
	// It only take effect in `FormatApacheCommonLog` and `FormatApacheCombinedLog`.
	ReservedTokenNameHTTPStatusCode string = "httpStatusCode"

	// ReservedTokenNameHTTPURL is the reserved token name for the HTTP URL.
	// It only take effect in `FormatApacheCommonLog` and `FormatApacheCombinedLog`.
	ReservedTokenNameHTTPURL string = "httpURL"

	// ReservedTokenNameHTTPContentLength is the reserved token name for the HTTP content length.
	// It only take effect in `FormatApacheCommonLog` and `FormatApacheCombinedLog`.
	ReservedTokenNameHTTPContentLength string = "httpContentLength"
)

const (
	// ApacheCommonLogTemplate is the template for outputting the fake data in `FormatApacheCommonLog` format.
	ApacheCommonLogTemplate string = "{{ ." + ReservedTokenNameHost + "}} - {{ ." + ReservedTokenNameUserID + "}} [{{ ." + ReservedTokenNameTimestamp + " }}] \"{{ ." + ReservedTokenNameHTTPMethod + " }} {{ ." + ReservedTokenNameHTTPURL + " }} {{ ." + ReservedTokenNameHTTPVersion + " }}\" {{ ." + ReservedTokenNameHTTPStatusCode + " }} {{ ." + ReservedTokenNameHTTPContentLength + " }}"
)

var (
	// ApacheCommonLogTokens is the list of tokens for the ApacheCommonLog format.
	ApacheCommonLogTokens = []*faker.Token{
		{
			Name: ReservedTokenNameHost,
			FakeConfig: faker.FakeConfig{
				Type: faker.FakeTypeIPv4,
			},
		},
		{
			Name: ReservedTokenNameUserID,
			FakeConfig: faker.FakeConfig{
				Type: faker.FakeTypeUserID,
			},
		},
		{
			Name: ReservedTokenNameHTTPMethod,
			FakeConfig: faker.FakeConfig{
				Type: faker.FakeTypeHTTPMethod,
			},
		},
		{
			Name: ReservedTokenNameHTTPURL,
			FakeConfig: faker.FakeConfig{
				Type: faker.FakeTypeURI,
				Options: map[string]any{
					"url": true,
				},
			},
		},
		{
			Name: ReservedTokenNameHTTPVersion,
			FakeConfig: faker.FakeConfig{
				Type: faker.FakeTypeHTTPVersion,
			},
		},
		{
			Name: ReservedTokenNameHTTPStatusCode,
			FakeConfig: faker.FakeConfig{
				Type: faker.FakeTypeHTTPStatusCode,
			},
		},
		{
			Name: ReservedTokenNameHTTPContentLength,
			FakeConfig: faker.FakeConfig{
				Type: faker.FakeTypeNumber,
				Options: map[string]any{
					"min":  "100",
					"max":  "100000",
					"type": "int32",
				},
			},
		},
	}
)
