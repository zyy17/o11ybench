package templates

import (
	"github.com/zyy17/o11ybench/pkg/generator/faker"
)

const (
	// ApacheCombinedLogTemplate is the template for outputting the fake data in ApacheCombinedLog format.
	ApacheCombinedLogTemplate string = "{{ ." + ReservedTokenNameHost + "}} - {{ ." + ReservedTokenNameUserID + "}} [{{ ." + ReservedTokenNameTimestamp + " }}] \"{{ ." + ReservedTokenNameHTTPMethod + " }} {{ ." + ReservedTokenNameHTTPURL + " }} {{ ." + ReservedTokenNameHTTPVersion + " }}\" {{ ." + ReservedTokenNameHTTPStatusCode + " }} {{ ." + ReservedTokenNameHTTPContentLength + " }} \"{{ ." + ReservedTokenNameReferer + " }}\" \"{{ ." + ReservedTokenNameHTTPUserAgent + " }}\""
)

var (
	// ApacheCombinedLogTokens is the list of tokens for the ApacheCombinedLog format.
	ApacheCombinedLogTokens = []*faker.Token{
		{
			Name: ReservedTokenNameHost,
			FakeConfig: faker.FakeConfig{
				Type: faker.FakeTypeIPv4,
			},
		},
		{
			Name: ReservedTokenNameUserID,
			FakeConfig: faker.FakeConfig{
				Type: faker.FakeTypeUsername,
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
					"type": faker.NumberTypeInt32,
					"min":  "100",
					"max":  "100000",
				},
			},
		},
		{
			Name: ReservedTokenNameReferer,
			FakeConfig: faker.FakeConfig{
				Type: faker.FakeTypeURI,
				Options: map[string]any{
					"url": true,
				},
			},
		},
		{
			Name: ReservedTokenNameHTTPUserAgent,
			FakeConfig: faker.FakeConfig{
				Type: faker.FakeTypeHTTPUserAgent,
			},
		},
	}
)
