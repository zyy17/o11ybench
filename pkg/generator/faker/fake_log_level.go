package faker

import (
	"math/rand"
	"strings"

	"github.com/brianvoe/gofakeit"

	"github.com/zyy17/o11ybench/pkg/generator/common"
)

// FakeLogLevelOptions is the options for the fake log level.
type FakeLogLevelOptions struct {
	// Type is the type of the log level. It can be "apache", "syslog" and "general".
	Type string `yaml:"type"`

	// Levels is the levels of the log level.
	// If set the levels, the log level will be selected from the levels.
	Levels []string `yaml:"levels"`

	// Uppercase is whether the log level is uppercase.
	Uppercase bool `yaml:"uppercase"`
}

// FakeLogLevel generates a fake log level.
func FakeLogLevel(_ common.ElementType, opts Options) (string, error) {
	var options FakeLogLevelOptions
	if err := parseOptions(opts, &options); err != nil {
		return "", err
	}

	var logLevel string
	if options.Type == "" && len(options.Levels) > 0 {
		logLevel = options.Levels[rand.Intn(len(options.Levels))]
	} else {
		logLevel = gofakeit.LogLevel(options.Type)
	}

	if options.Uppercase {
		logLevel = strings.ToUpper(logLevel)
	}

	return logLevel, nil
}
