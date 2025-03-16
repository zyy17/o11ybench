package templates

import (
	"github.com/zyy17/o11ybench/pkg/generator/common"
	"github.com/zyy17/o11ybench/pkg/generator/faker"
	"github.com/zyy17/o11ybench/pkg/generator/logs/types"
)

var RFC5424Log = BuiltinLogFormat{
	Template:        RFC5424LogTemplate,
	Tokens:          RFC5424LogTokens,
	TimestampFormat: common.TimestampFormatTypeRFC5424,
}

const (
	// RFC5424LogTemplate is the template for outputting the fake data in RFC5424Log format.
	// RFC5424Log : <priority>{version} {iso-timestamp} {hostname} {application} {pid} {message-id} {structured-data} {message}
	// Example: <87>3 2025-03-06T07:32:21.000Z chiefmonetize.org suscipit 5541 ID648 - We need to quantify the bluetooth RAM firewall!
	RFC5424LogTemplate string = "<{{ ." + ReservedTokenNamePriority + " }}>{{ ." + ReservedTokenNameVersion + " }} {{ ." + ReservedTokenNameTimestamp + " }} {{ ." + ReservedTokenNameHost + " }} {{ ." + ReservedTokenNameApplication + " }} {{ ." + ReservedTokenNamePid + " }} ID{{ ." + ReservedTokenNameMessageID + " }} {{ ." + ReservedTokenNameStructuredData + " }} {{ ." + ReservedTokenNameMessage + " }}"
)

var (
	// RFC5424LogTokens is the list of tokens for the RFC5424Log format.
	RFC5424LogTokens = []*types.LogToken{
		{
			Name: ReservedTokenNamePriority,
			Type: common.ElementTypeInt32,
			FakeConfig: &faker.FakeConfig{
				Kind: faker.FakeDataKindNumber,
				Options: map[string]any{
					"min": "0",
					"max": "191",
				},
			},
		},
		{
			Name: ReservedTokenNameVersion,
			Type: common.ElementTypeInt32,
			FakeConfig: &faker.FakeConfig{
				Kind: faker.FakeDataKindNumber,
				Options: map[string]any{
					"min": "1",
					"max": "3",
				},
			},
		},
		{
			Name: ReservedTokenNameHost,
			Type: common.ElementTypeString,
			FakeConfig: &faker.FakeConfig{
				Kind: faker.FakeDataKindDomainName,
			},
		},
		{
			Name: ReservedTokenNameApplication,
			Type: common.ElementTypeString,
			FakeConfig: &faker.FakeConfig{
				Kind: faker.FakeDataKindWords,
				Options: map[string]any{
					"count": 1,
				},
			},
		},
		{
			Name: ReservedTokenNamePid,
			Type: common.ElementTypeInt32,
			FakeConfig: &faker.FakeConfig{
				Kind: faker.FakeDataKindNumber,
				Options: map[string]any{
					"min": "1",
					"max": "10000",
				},
			},
		},
		{
			Name: ReservedTokenNameMessageID,
			Type: common.ElementTypeInt32,
			FakeConfig: &faker.FakeConfig{
				Kind: faker.FakeDataKindNumber,
				Options: map[string]any{
					"min": "1",
					"max": "1000",
				},
			},
		},
		{
			Name:  ReservedTokenNameStructuredData,
			Value: "-",
		},
		{
			Name: ReservedTokenNameMessage,
			Type: common.ElementTypeString,
			FakeConfig: &faker.FakeConfig{
				Kind: faker.FakeDataKindHackerPhrase,
			},
		},
	}
)
