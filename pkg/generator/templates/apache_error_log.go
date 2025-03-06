package templates

import (
	"github.com/zyy17/o11ybench/pkg/generator/faker"
)

const (
	// ApacheErrorLogTemplate is the template for outputting the fake data in ApacheErrorLog format.
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
