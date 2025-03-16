package templates

import (
	"github.com/zyy17/o11ybench/pkg/generator/common"
	"github.com/zyy17/o11ybench/pkg/generator/faker"
	"github.com/zyy17/o11ybench/pkg/generator/logs/types"
)

var RFC3164Log = BuiltinLogFormat{
	Template:        RFC3164LogTemplate,
	Tokens:          RFC3164LogTokens,
	TimestampFormat: common.TimestampFormatTypeRFC3164,
}

const (
	// RFC3164LogTemplate is the template for outputting the fake data in RFC3164Log format.
	// RFC3164Log: <priority>{timestamp} {hostname} {application}[{pid}]: {message}
	// Example: <57>Mar 06 07:30:21 Lockman1496 veniam[7031]: Copying the program won't do anything, we need to connect the mobile SMTP protocol!
	RFC3164LogTemplate string = "<{{ ." + ReservedTokenNamePriority + " }}>{{ ." + ReservedTokenNameTimestamp + " }} {{ ." + ReservedTokenNameHost + " }} {{ ." + ReservedTokenNameApplication + " }}[{{ ." + ReservedTokenNamePid + " }}]: {{ ." + ReservedTokenNameMessage + " }}"
)

var (
	// RFC3164LogTokens is the list of tokens for the RFC3164Log format.
	RFC3164LogTokens = []*types.LogToken{
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
			Name: ReservedTokenNameHost,
			Type: common.ElementTypeString,
			FakeConfig: &faker.FakeConfig{
				Kind: faker.FakeDataKindUsername,
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
			Name: ReservedTokenNameMessage,
			Type: common.ElementTypeString,
			FakeConfig: &faker.FakeConfig{
				Kind: faker.FakeDataKindHackerPhrase,
			},
		},
	}
)
