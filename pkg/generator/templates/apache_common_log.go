package templates

import (
	"github.com/zyy17/o11ybench/pkg/generator/faker"
)

const (
	// ApacheCommonLogTemplate is the template for outputting the fake data in ApacheCommonLog format.
	// ApacheCommonLog : {host} {user-identifier} {auth-user-id} [{datetime}] "{method} {request} {protocol}" {response-code} {bytes}
	// Example: 249.153.155.15 - Bechtelar5851 [06/Mar/2025:07:20:34 +0000] "GET http://www.nationalimpactful.org/technologies/transition HTTP/1.1" 501 8134
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
					"type": faker.NumberTypeInt32,
					"min":  "100",
					"max":  "100000",
				},
			},
		},
	}
)
