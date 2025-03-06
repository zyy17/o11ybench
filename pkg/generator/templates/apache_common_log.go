package templates

import (
	"github.com/zyy17/o11ybench/pkg/generator/faker"
)

const (
	// ApacheCommonLogTemplate is the template for outputting the fake data in ApacheCommonLog format.
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
					"min":  "100",
					"max":  "100000",
					"type": "int32",
				},
			},
		},
	}
)
