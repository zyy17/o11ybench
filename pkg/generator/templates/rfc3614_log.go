package templates

import (
	"github.com/zyy17/o11ybench/pkg/generator/faker"
)

const (
	// RFC3614LogTemplate is the template for outputting the fake data in RFC3614Log format.
	// RFC3164Log : <priority>{timestamp} {hostname} {application}[{pid}]: {message}
	RFC3614LogTemplate string = "<{{ ." + ReservedTokenNamePriority + " }}>{{ ." + ReservedTokenNameTimestamp + " }} {{ ." + ReservedTokenNameHost + " }} {{ ." + ReservedTokenNameApplication + " }}[{{ ." + ReservedTokenNamePid + " }}]: {{ ." + ReservedTokenNameMessage + " }}"
)

var (
	// RFC3614LogTokens is the list of tokens for the RFC3614Log format.
	RFC3614LogTokens = []*faker.Token{
		{
			Name: ReservedTokenNamePriority,
			FakeConfig: faker.FakeConfig{
				Type: faker.FakeTypeNumber,
				Options: map[string]any{
					"type": faker.NumberTypeInt32,
					"min":  "0",
					"max":  "191",
				},
			},
		},
		{
			Name: ReservedTokenNameHost,
			FakeConfig: faker.FakeConfig{
				Type: faker.FakeTypeUsername,
			},
		},
		{
			Name: ReservedTokenNameApplication,
			FakeConfig: faker.FakeConfig{
				Type: faker.FakeTypeWords,
				Options: map[string]any{
					"count": 1,
				},
			},
		},
		{
			Name: ReservedTokenNamePid,
			FakeConfig: faker.FakeConfig{
				Type: faker.FakeTypeNumber,
				Options: map[string]any{
					"type": faker.NumberTypeInt32,
					"min":  "1",
					"max":  "100000",
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
