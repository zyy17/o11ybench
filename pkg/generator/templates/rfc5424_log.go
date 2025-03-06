package templates

import (
	"github.com/zyy17/o11ybench/pkg/generator/faker"
)

const (
	// RFC5424LogTemplate is the template for outputting the fake data in RFC5424Log format.
	// RFC5424Log : <priority>{version} {iso-timestamp} {hostname} {application} {pid} {message-id} {structured-data} {message}
	// Example: <87>3 2025-03-06T07:32:21.000Z chiefmonetize.org suscipit 5541 ID648 - We need to quantify the bluetooth RAM firewall!
	RFC5424LogTemplate string = "<{{ ." + ReservedTokenNamePriority + " }}>{{ ." + ReservedTokenNameVersion + " }} {{ ." + ReservedTokenNameTimestamp + " }} {{ ." + ReservedTokenNameHost + " }} {{ ." + ReservedTokenNameApplication + " }} {{ ." + ReservedTokenNamePid + " }} ID{{ ." + ReservedTokenNameMessageID + " }} {{ ." + ReservedTokenNameStructuredData + " }} {{ ." + ReservedTokenNameMessage + " }}"
)

var (
	// RFC5424LogTokens is the list of tokens for the RFC5424Log format.
	RFC5424LogTokens = []*faker.Token{
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
			Name: ReservedTokenNameVersion,
			FakeConfig: faker.FakeConfig{
				Type: faker.FakeTypeNumber,
				Options: map[string]any{
					"type": faker.NumberTypeInt32,
					"min":  "1",
					"max":  "3",
				},
			},
		},
		{
			Name: ReservedTokenNameHost,
			FakeConfig: faker.FakeConfig{
				Type: faker.FakeTypeDomainName,
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
			Name: ReservedTokenNameMessageID,
			FakeConfig: faker.FakeConfig{
				Type: faker.FakeTypeNumber,
				Options: map[string]any{
					"type": faker.NumberTypeInt32,
					"min":  "1",
					"max":  "1000",
				},
			},
		},
		{
			Name:  ReservedTokenNameStructuredData,
			Value: "-",
		},
		{
			Name: ReservedTokenNameMessage,
			FakeConfig: faker.FakeConfig{
				Type: faker.FakeTypeHackerPhrase,
			},
		},
	}
)
