package templates

import (
	"github.com/zyy17/o11ybench/pkg/generator/common"
	"github.com/zyy17/o11ybench/pkg/generator/faker"
	"github.com/zyy17/o11ybench/pkg/generator/logs/types"
)

var ApacheCommonLog = BuiltinLogFormat{
	Template:        ApacheCommonLogTemplate,
	Tokens:          ApacheCommonLogTokens,
	TimestampFormat: common.TimestampFormatTypeApache,
}

const (
	// ApacheCommonLogTemplate is the template for outputting the fake data in ApacheCommonLog format.
	// ApacheCommonLog : {host} {user-identifier} {auth-user-id} [{datetime}] "{method} {request} {protocol}" {response-code} {bytes}
	// Example: 249.153.155.15 - Bechtelar5851 [06/Mar/2025:07:20:34 +0000] "GET http://www.nationalimpactful.org/technologies/transition HTTP/1.1" 501 8134
	ApacheCommonLogTemplate string = "{{ ." + ReservedTokenNameHost + "}} - {{ ." + ReservedTokenNameUserID + "}} [{{ ." + ReservedTokenNameTimestamp + " }}] \"{{ ." + ReservedTokenNameHTTPMethod + " }} {{ ." + ReservedTokenNameHTTPURL + " }} {{ ." + ReservedTokenNameHTTPVersion + " }}\" {{ ." + ReservedTokenNameHTTPStatusCode + " }} {{ ." + ReservedTokenNameHTTPContentLength + " }}"
)

var (
	// ApacheCommonLogTokens is the list of tokens for the ApacheCommonLog format.
	ApacheCommonLogTokens = []*types.LogToken{
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
	}
)
