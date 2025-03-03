package faker

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// FakeType is the type of the fake data.
type FakeType string

const (
	// FakeTypeWords is used to generate fake words.
	FakeTypeWords FakeType = "words"

	// FakeTypeNumber is used to generate a fake number.
	FakeTypeNumber FakeType = "number"

	// FakeTypeIPv4 is used to generate a fake IPv4 address.
	FakeTypeIPv4 FakeType = "ipv4"

	// FakeTypeUserID is used to generate a fake user ID.
	FakeTypeUserID FakeType = "userID"

	// FakeTypeURI is used to generate a fake URI.
	FakeTypeURI FakeType = "uri"

	// FakeTypeHTTPVersion is used to generate a fake HTTP version.
	FakeTypeHTTPVersion FakeType = "httpVersion"

	// FakeTypeHTTPMethod is used to generate a fake HTTP method.
	FakeTypeHTTPMethod FakeType = "httpMethod"

	// FakeTypeHTTPStatusCode is used to generate a fake HTTP status code.
	FakeTypeHTTPStatusCode FakeType = "httpStatusCode"

	// FakeTypeHTTPUserAgent is used to generate a fake HTTP user agent.
	FakeTypeHTTPUserAgent FakeType = "httpUserAgent"
)

// Token is a token that is part of the fake data.
type Token struct {
	// Name is the internal name of the token. You can use this name in custom format to refer to the token.
	Name string `yaml:"name" json:"name"`

	// FakeConfig is the configuration for how to generate the fake data.
	FakeConfig FakeConfig `yaml:"fake" json:"fake"`

	// Display is the final name of the token. It will be used in the fake data output.
	Display string `yaml:"display,omitempty" json:"display,omitempty"`
}

// Options is specific options for each fake type.
type Options map[string]any

// FakeConfig is the configuration for the fake data.
type FakeConfig struct {
	// Type is the type of the fake data.
	Type FakeType `json:"type" yaml:"type"`

	// Options is the specific options for the fake type.
	Options Options `json:"options,omitempty" yaml:"options,omitempty"`
}

func Fake(cfg *FakeConfig) (any, error) {
	switch cfg.Type {
	case FakeTypeWords:
		return FakeWords(cfg.Options)
	case FakeTypeNumber:
		return FakeNumber(cfg.Options)
	case FakeTypeIPv4:
		return FakeIPv4(cfg.Options)
	case FakeTypeUserID:
		return FakeUserID(cfg.Options)
	case FakeTypeURI:
		return FakeURI(cfg.Options)
	case FakeTypeHTTPVersion:
		return FakeHTTPVersion(cfg.Options)
	case FakeTypeHTTPMethod:
		return FakeHTTPMethod(cfg.Options)
	case FakeTypeHTTPStatusCode:
		return FakeHTTPStatusCode(cfg.Options)
	case FakeTypeHTTPUserAgent:
		return FakeHTTPUserAgent(cfg.Options)
	}

	return nil, fmt.Errorf("unknown fake type: %s", cfg.Type)
}

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
