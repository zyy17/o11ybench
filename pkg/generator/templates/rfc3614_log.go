package templates

import (
	"github.com/zyy17/o11ybench/pkg/generator/faker"
)

const (
	// RFC3164LogTemplate is the template for outputting the fake data in RFC3164Log format.
	// RFC3164Log: <priority>{timestamp} {hostname} {application}[{pid}]: {message}
	// Example: <57>Mar 06 07:30:21 Lockman1496 veniam[7031]: Copying the program won't do anything, we need to connect the mobile SMTP protocol!
	RFC3164LogTemplate string = "<{{ ." + ReservedTokenNamePriority + " }}>{{ ." + ReservedTokenNameTimestamp + " }} {{ ." + ReservedTokenNameHost + " }} {{ ." + ReservedTokenNameApplication + " }}[{{ ." + ReservedTokenNamePid + " }}]: {{ ." + ReservedTokenNameMessage + " }}"
)

var (
	// RFC3164LogTokens is the list of tokens for the RFC3164Log format.
	RFC3164LogTokens = []*faker.Token{
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
					"max":  "10000",
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
