package templates

import (
	"github.com/zyy17/o11ybench/pkg/generator/faker"
)

const (
	// ApacheErrorLogTemplate is the template for outputting the fake data in ApacheErrorLog format.
	// ApacheErrorLog: [{timestamp}] [{module}:{severity}] [pid {pid}:tid {thread-id}] [client %{client}:{port}] %{message}
	// Example: [Thu Mar 06 07:26:33 2025] [dolore:emerg] [pid 19366:tid 95232] [client 135.199.249.25:39294] We need to back up the cross-platform PNG alarm!
	ApacheErrorLogTemplate string = "[{{ ." + ReservedTokenNameTimestamp + " }}] [{{ ." + ReservedTokenNameModule + " }}:{{ ." + ReservedTokenNameLogLevel + " }}] [pid {{ ." + ReservedTokenNamePid + " }}:tid {{ ." + ReservedTokenNameTid + " }}] [client {{ ." + ReservedTokenNameHost + " }}:{{ ." + ReservedTokenNamePort + " }}] {{ ." + ReservedTokenNameMessage + " }}"
)

var (
	// ApacheErrorLogTokens is the list of tokens for the ApacheErrorLog format.
	ApacheErrorLogTokens = []*faker.Token{
		{
			Name: ReservedTokenNameModule,
			FakeConfig: faker.FakeConfig{
				Type: faker.FakeTypeWords,
				Options: faker.Options{
					"count": 1,
				},
			},
		},
		{
			Name: ReservedTokenNameLogLevel,
			FakeConfig: faker.FakeConfig{
				Type: faker.FakeTypeLogLevel,
				Options: faker.Options{
					"type": "apache",
				},
			},
		},
		{
			Name: ReservedTokenNamePid,
			FakeConfig: faker.FakeConfig{
				Type: faker.FakeTypeNumber,
				Options: faker.Options{
					"type": faker.NumberTypeInt32,
					"min":  "1",
					"max":  "100000",
				},
			},
		},
		{
			Name: ReservedTokenNameTid,
			FakeConfig: faker.FakeConfig{
				Type: faker.FakeTypeNumber,
				Options: faker.Options{
					"type": faker.NumberTypeInt32,
					"min":  "1",
					"max":  "100000",
				},
			},
		},
		{
			Name: ReservedTokenNameHost,
			FakeConfig: faker.FakeConfig{
				Type: faker.FakeTypeIPv4,
			},
		},
		{
			Name: ReservedTokenNamePort,
			FakeConfig: faker.FakeConfig{
				Type: faker.FakeTypeNumber,
				Options: faker.Options{
					"type": faker.NumberTypeInt32,
					"min":  "1",
					"max":  "65535",
				},
			},
		},
		{
			Name: ReservedTokenNameMessage,
			FakeConfig: faker.FakeConfig{
				Type: faker.FakeTypeHackerPhrase,
			},
		},
	}
)
