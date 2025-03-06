package templates

import (
	"github.com/zyy17/o11ybench/pkg/generator/faker"
)

const (
	// ApacheCombinedLogTemplate is the template for outputting the fake data in ApacheCombinedLog format.
	// ApacheCombinedLog : {host} {user-identifier} {auth-user-id} [{datetime}] "{method} {request} {protocol}" {response-code} {bytes} "{referrer}" "{agent}"
	// Example: 45.239.98.253 - Hoppe5924 [06/Mar/2025:21:23:09 +0800] "PATCH /functionalities/reinvent/24%2f365 HTTP/2.0" 205 49431 "https://www.corporatereinvent.io/killer/vortals/mission-critical" "Mozilla/5.0 (iPad; CPU OS 9_1_3 like Mac OS X; en-US) AppleWebKit/534.10.1 (KHTML, like Gecko) Version/3.0.5 Mobile/8B118 Safari/6534.10.1"
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
