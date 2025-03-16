package faker

import (
	"fmt"

	"gopkg.in/yaml.v3"

	"github.com/zyy17/o11ybench/pkg/generator/common"
)

// Options is specific options for each fake type.
type Options map[string]any

// FakeConfig is the configuration for the fake data.
type FakeConfig struct {
	// Kind is the kind of the fake data.
	Kind FakeDataKind `yaml:"kind"`

	// Options is the specific options for the fake type.
	Options Options `yaml:"options,omitempty"`
}

func Fake(typ common.ElementType, cfg *FakeConfig) (any, error) {
	switch cfg.Kind {
	case FakeDataKindWords:
		return FakeWords(typ, cfg.Options)
	case FakeDataKindNumber:
		return FakeNumber(typ, cfg.Options)
	case FakeDataKindIPv4:
		return FakeIPv4(typ, cfg.Options)
	case FakeDataKindUsername:
		return FakeUsername(typ, cfg.Options)
	case FakeDataKindURI:
		return FakeURI(typ, cfg.Options)
	case FakeDataKindHTTPVersion:
		return FakeHTTPVersion(typ, cfg.Options)
	case FakeDataKindHTTPMethod:
		return FakeHTTPMethod(typ, cfg.Options)
	case FakeDataKindHTTPStatusCode:
		return FakeHTTPStatusCode(typ, cfg.Options)
	case FakeDataKindHTTPUserAgent:
		return FakeHTTPUserAgent(typ, cfg.Options)
	case FakeDataKindLogLevel:
		return FakeLogLevel(typ, cfg.Options)
	case FakeDataKindHackerPhrase:
		return FakeHackerPhrase(typ, cfg.Options)
	case FakeDataKindUUID:
		return FakeUUID(typ, cfg.Options)
	case FakeDataKindDomainName:
		return FakeDomainName(typ, cfg.Options)
	case FakeDataKindLogs:
		return FakeLogs(typ, cfg.Options)
	}

	return nil, fmt.Errorf("unknown fake data kind: %s", cfg.Kind)
}

// FakeDataKind is the kind of the fake data.
type FakeDataKind string

const (
	// FakeDataKindWords is used to generate fake words.
	FakeDataKindWords FakeDataKind = "words"

	// FakeDataKindNumber is used to generate a fake number.
	FakeDataKindNumber FakeDataKind = "number"

	// FakeDataKindIPv4 is used to generate a fake IPv4 address.
	FakeDataKindIPv4 FakeDataKind = "ipv4"

	// FakeDataKindUsername is used to generate a fake username.
	FakeDataKindUsername FakeDataKind = "username"

	// FakeDataKindURI is used to generate a fake URI.
	FakeDataKindURI FakeDataKind = "uri"

	// FakeDataKindHTTPVersion is used to generate a fake HTTP version.
	FakeDataKindHTTPVersion FakeDataKind = "httpVersion"

	// FakeDataKindHTTPMethod is used to generate a fake HTTP method.
	FakeDataKindHTTPMethod FakeDataKind = "httpMethod"

	// FakeDataKindHTTPStatusCode is used to generate a fake HTTP status code.
	FakeDataKindHTTPStatusCode FakeDataKind = "httpStatusCode"

	// FakeDataKindHTTPUserAgent is used to generate a fake HTTP user agent.
	FakeDataKindHTTPUserAgent FakeDataKind = "httpUserAgent"

	// FakeDataKindLogLevel is used to generate a fake log level.
	FakeDataKindLogLevel FakeDataKind = "logLevel"

	// FakeDataKindHackerPhrase is used to generate a fake hacker phrase.
	FakeDataKindHackerPhrase FakeDataKind = "hackerPhrase"

	// FakeDataKindUUID is used to generate a fake uuid.
	FakeDataKindUUID FakeDataKind = "uuid"

	// FakeDataKindDomainName is used to generate a fake domain name.
	FakeDataKindDomainName FakeDataKind = "domainName"

	// FakeDataKindLogs is used to generate a fake log.
	FakeDataKindLogs FakeDataKind = "logs"
)

// parseOptions parse the options to the target config type.
// target must be a pointer to the target config type.
func parseOptions(options Options, target any) error {
	// If the options is nil, return nil.
	if options == nil {
		return nil
	}

	yamlBytes, err := yaml.Marshal(options)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(yamlBytes, target)
}
