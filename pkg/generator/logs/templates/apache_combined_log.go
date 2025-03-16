package templates

import (
	"github.com/zyy17/o11ybench/pkg/generator/common"
	"github.com/zyy17/o11ybench/pkg/generator/faker"
	"github.com/zyy17/o11ybench/pkg/generator/logs/types"
)

var ApacheCombinedLog = BuiltinLogFormat{
	Template:        ApacheCombinedLogTemplate,
	Tokens:          ApacheCombinedLogTokens,
	TimestampFormat: common.TimestampFormatTypeApache,
}

const (
	// ApacheCombinedLogTemplate is the template for outputting the fake data in ApacheCombinedLog format.
	// ApacheCombinedLog : {host} {user-identifier} {auth-user-id} [{datetime}] "{method} {request} {protocol}" {response-code} {bytes} "{referrer}" "{agent}"
	// Example: 45.239.98.253 - Hoppe5924 [06/Mar/2025:21:23:09 +0800] "PATCH /functionalities/reinvent/24%2f365 HTTP/2.0" 205 49431 "https://www.corporatereinvent.io/killer/vortals/mission-critical" "Mozilla/5.0 (iPad; CPU OS 9_1_3 like Mac OS X; en-US) AppleWebKit/534.10.1 (KHTML, like Gecko) Version/3.0.5 Mobile/8B118 Safari/6534.10.1"
	ApacheCombinedLogTemplate string = "{{ ." + ReservedTokenNameHost + "}} - {{ ." + ReservedTokenNameUserID + "}} [{{ ." + ReservedTokenNameTimestamp + " }}] \"{{ ." + ReservedTokenNameHTTPMethod + " }} {{ ." + ReservedTokenNameHTTPURL + " }} {{ ." + ReservedTokenNameHTTPVersion + " }}\" {{ ." + ReservedTokenNameHTTPStatusCode + " }} {{ ." + ReservedTokenNameHTTPContentLength + " }} \"{{ ." + ReservedTokenNameReferer + " }}\" \"{{ ." + ReservedTokenNameHTTPUserAgent + " }}\""
)

var (
	// ApacheCombinedLogTokens is the list of tokens for the ApacheCombinedLog format.
	ApacheCombinedLogTokens = []*types.LogToken{
		{
			Name: ReservedTokenNameHost,
			Type: common.ElementTypeString,
			FakeConfig: &faker.FakeConfig{
				Kind: faker.FakeDataKindIPv4,
			},
		},
		{
			Name: ReservedTokenNameUserID,
			Type: common.ElementTypeString,
			FakeConfig: &faker.FakeConfig{
				Kind: faker.FakeDataKindUsername,
			},
		},
		{
			Name: ReservedTokenNameHTTPMethod,
			Type: common.ElementTypeString,
			FakeConfig: &faker.FakeConfig{
				Kind: faker.FakeDataKindHTTPMethod,
			},
		},
		{
			Name: ReservedTokenNameHTTPURL,
			Type: common.ElementTypeString,
			FakeConfig: &faker.FakeConfig{
				Kind: faker.FakeDataKindURI,
				Options: map[string]any{
					"url": true,
				},
			},
		},
		{
			Name: ReservedTokenNameHTTPVersion,
			Type: common.ElementTypeString,
			FakeConfig: &faker.FakeConfig{
				Kind: faker.FakeDataKindHTTPVersion,
			},
		},
		{
			Name: ReservedTokenNameHTTPStatusCode,
			Type: common.ElementTypeInt32,
			FakeConfig: &faker.FakeConfig{
				Kind: faker.FakeDataKindHTTPStatusCode,
			},
		},
		{
			Name: ReservedTokenNameHTTPContentLength,
			Type: common.ElementTypeInt32,
			FakeConfig: &faker.FakeConfig{
				Kind: faker.FakeDataKindNumber,
				Options: map[string]any{
					"min": "100",
					"max": "100000",
				},
			},
		},
		{
			Name: ReservedTokenNameReferer,
			Type: common.ElementTypeString,
			FakeConfig: &faker.FakeConfig{
				Kind: faker.FakeDataKindURI,
				Options: map[string]any{
					"url": true,
				},
			},
		},
		{
			Name: ReservedTokenNameHTTPUserAgent,
			Type: common.ElementTypeString,
			FakeConfig: &faker.FakeConfig{
				Kind: faker.FakeDataKindHTTPUserAgent,
			},
		},
	}
)
