package templates

import (
	"github.com/zyy17/o11ybench/pkg/generator/common"
	"github.com/zyy17/o11ybench/pkg/generator/faker"
	"github.com/zyy17/o11ybench/pkg/generator/logs/types"
)

var ApacheErrorLog = BuiltinLogFormat{
	Template:        ApacheErrorLogTemplate,
	Tokens:          ApacheErrorLogTokens,
	TimestampFormat: common.TimestampFormatTypeApacheError,
}

const (
	// ApacheErrorLogTemplate is the template for outputting the fake data in ApacheErrorLog format.
	// ApacheErrorLog: [{timestamp}] [{module}:{severity}] [pid {pid}:tid {thread-id}] [client %{client}:{port}] %{message}
	// Example: [Thu Mar 06 07:26:33 2025] [dolore:emerg] [pid 19366:tid 95232] [client 135.199.249.25:39294] We need to back up the cross-platform PNG alarm!
	ApacheErrorLogTemplate string = "[{{ ." + ReservedTokenNameTimestamp + " }}] [{{ ." + ReservedTokenNameModule + " }}:{{ ." + ReservedTokenNameLogLevel + " }}] [pid {{ ." + ReservedTokenNamePid + " }}:tid {{ ." + ReservedTokenNameTid + " }}] [client {{ ." + ReservedTokenNameHost + " }}:{{ ." + ReservedTokenNamePort + " }}] {{ ." + ReservedTokenNameMessage + " }}"
)

var (
	// ApacheErrorLogTokens is the list of tokens for the ApacheErrorLog format.
	ApacheErrorLogTokens = []*types.LogToken{
		{
			Name: ReservedTokenNameModule,
			Type: common.ElementTypeString,
			FakeConfig: &faker.FakeConfig{
				Kind: faker.FakeDataKindWords,
				Options: faker.Options{
					"count": 1,
				},
			},
		},
		{
			Name: ReservedTokenNameLogLevel,
			Type: common.ElementTypeString,
			FakeConfig: &faker.FakeConfig{
				Kind: faker.FakeDataKindLogLevel,
				Options: faker.Options{
					"type": "apache",
				},
			},
		},
		{
			Name: ReservedTokenNamePid,
			Type: common.ElementTypeInt32,
			FakeConfig: &faker.FakeConfig{
				Kind: faker.FakeDataKindNumber,
				Options: faker.Options{
					"min": "1",
					"max": "100000",
				},
			},
		},
		{
			Name: ReservedTokenNameTid,
			Type: common.ElementTypeInt32,
			FakeConfig: &faker.FakeConfig{
				Kind: faker.FakeDataKindNumber,
				Options: faker.Options{
					"min": "1",
					"max": "100000",
				},
			},
		},
		{
			Name: ReservedTokenNameHost,
			Type: common.ElementTypeString,
			FakeConfig: &faker.FakeConfig{
				Kind: faker.FakeDataKindIPv4,
			},
		},
		{
			Name: ReservedTokenNamePort,
			Type: common.ElementTypeInt32,
			FakeConfig: &faker.FakeConfig{
				Kind: faker.FakeDataKindNumber,
				Options: faker.Options{
					"min": "1",
					"max": "65535",
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
