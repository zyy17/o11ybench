package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"maps"
	"os"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/zyy17/o11ybench/pkg/generator/faker"
	"github.com/zyy17/o11ybench/pkg/generator/templates"
)

// FakeDataConfig is the configuration for the fake data generator.
type FakeDataConfig struct {
	// Tokens is the list of tokens that are part of the fake data.
	Tokens []*faker.Token `yaml:"tokens,omitempty" json:"tokens,omitempty"`

	// Output is used to configure the output of the fake data.
	Output Output `yaml:"output" json:"output"`

	// TimeRange is the time range of the fake data.
	TimeRange TimeRange `yaml:"timeRange,omitempty" json:"timeRange,omitempty"`
}

// Output is used to configure the output of the fake data.
type Output struct {
	// LogFormat is the output log format of the fake data, for example, `ApacheCommonLog`, `JSON`, etc.
	LogFormat LogFormat `yaml:"logFormat,omitempty" json:"logFormat,omitempty"`

	// Custom is the custom format of the fake data. You can use `{{ .<token_name> }}` to refer to the token.
	Custom string `yaml:"custom,omitempty" json:"custom,omitempty"`

	// Count is the number of logs to generate.
	Count int `yaml:"count,omitempty" json:"count,omitempty"`
}

// TimeRange is the time range of the fake data.
type TimeRange struct {
	// Start is the start time of the time range.
	Start string `yaml:"start,omitempty" json:"start,omitempty"`

	// End is the end time of the time range.
	End string `yaml:"end,omitempty" json:"end,omitempty"`

	// Format is the format of the time range.
	Format string `yaml:"format,omitempty" json:"format,omitempty"`

	// Timezone is the timezone of the time range.
	Timezone string `yaml:"timezone,omitempty" json:"timezone,omitempty"`

	// Interval is the interval for each log.
	Interval string `yaml:"interval,omitempty" json:"interval,omitempty"`
}

// LogFormat is the output format of the fake data.
type LogFormat string

const (
	// LogFormatApacheCommonLog is the format of apache common log.
	LogFormatApacheCommonLog LogFormat = "ApacheCommonLog"

	// LogFormatApacheCombinedLog is the format of apache combined log.
	LogFormatApacheCombinedLog LogFormat = "ApacheCombinedLog"

	// LogFormatJSON is the format of JSON.
	LogFormatJSON LogFormat = "JSON"
)

type TimestampFormat string

const (
	TimestampFormatApache      TimestampFormat = "Apache"
	TimestampFormatApacheError TimestampFormat = "ApacheError"
	TimestampFormatRFC3164     TimestampFormat = "RFC3164"
	TimestampFormatRFC5424     TimestampFormat = "RFC5424"
	TimestampFormatRFC3339     TimestampFormat = "RFC3339"
	TimestampFormatCommonLog   TimestampFormat = "CommonLog"
	TimestampFormatClickHouse  TimestampFormat = "ClickHouse"
	TimestampFormatUnix        TimestampFormat = "UnixSeconds"
)

const (
	Apache      = "02/Jan/2006:15:04:05 -0700"
	ApacheError = "Mon Jan 02 15:04:05 2006"
	RFC3164     = "Jan 02 15:04:05"
	RFC5424     = "2006-01-02T15:04:05.000Z"
	CommonLog   = "02/Jan/2006:15:04:05 -0700"
	ClickHouse  = "02/Jan/2006:15:04:05 -0700"
)

type Generator struct {
	tokens    []*faker.Token
	output    *Output
	timeRange *timeRange
}

type timeRange struct {
	start      time.Time
	end        time.Time
	location   *time.Location
	interval   time.Duration
	timeFormat string
}

func NewGenerator(config *FakeDataConfig) (*Generator, error) {
	timeRange, err := parseTimeRange(&config.TimeRange)
	if err != nil {
		return nil, err
	}

	return &Generator{
		tokens:    config.Tokens,
		output:    &config.Output,
		timeRange: timeRange,
	}, nil
}

func NewGeneratorFromFile(configFile string) (*Generator, error) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	cfg := &FakeDataConfig{}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	return NewGenerator(cfg)
}

func (g *Generator) Generate() ([]string, error) {
	var (
		logs    = make([]string, 0)
		count   = 0
		current = g.timeRange.start
	)

	for {
		if g.output.Count > 0 && count >= g.output.Count {
			break
		}

		if current.After(g.timeRange.end) {
			break
		}

		var (
			generatedData = make(map[string]any)

			log string
			err error
		)

		current = current.Add(g.timeRange.interval)
		for _, token := range g.tokens {
			value, err := faker.Fake(&token.FakeConfig)
			if err != nil {
				return nil, err
			}

			generatedData[token.Name] = value
		}

		// Set the timestamp to the current time.
		generatedData[templates.ReservedTokenNameTimestamp] = g.timestamp(current.UnixNano())

		if g.output.Custom != "" {
			log, err = g.templateOutput(g.output.Custom, generatedData)
			if err != nil {
				return nil, err
			}
			logs = append(logs, log)
			count++
			continue
		}

		switch g.output.LogFormat {
		case LogFormatApacheCommonLog:
			log, err = g.doGenerate(generatedData, templates.ApacheCommonLogTokens, templates.ApacheCommonLogTemplate)
			if err != nil {
				return nil, err
			}
		case LogFormatJSON:
			log, err = g.outputJSON(generatedData)
			if err != nil {
				return nil, err
			}
		}

		logs = append(logs, log)
		count++
	}

	return logs, nil
}

func parseTimeRange(cfg *TimeRange) (*timeRange, error) {
	location, err := time.LoadLocation(cfg.Timezone)
	if err != nil {
		return nil, err
	}

	start, err := time.Parse(time.RFC3339, cfg.Start)
	if err != nil {
		return nil, err
	}

	end, err := time.Parse(time.RFC3339, cfg.End)
	if err != nil {
		return nil, err
	}

	if start.After(end) {
		return nil, fmt.Errorf("start time '%s' is after end time '%s'", start, end)
	}

	interval, err := time.ParseDuration(cfg.Interval)
	if err != nil {
		return nil, err
	}

	if start.Add(interval).After(end) {
		return nil, fmt.Errorf("interval '%s' is too long, it will exceed the end time '%s'", interval, end)
	}

	return &timeRange{
		start:      start,
		end:        end,
		location:   location,
		interval:   interval,
		timeFormat: cfg.Format,
	}, nil
}

func (g *Generator) doGenerate(overrideData map[string]any, tokens []*faker.Token, templateString string) (string, error) {
	data := make(map[string]any)

	for _, token := range tokens {
		value, err := faker.Fake(&token.FakeConfig)
		if err != nil {
			return "", err
		}

		data[token.Name] = value
	}

	maps.Copy(data, overrideData)

	return g.templateOutput(templateString, data)
}

func (g *Generator) templateOutput(templateString string, data map[string]any) (string, error) {
	tmpl, err := template.New("output").Parse(templateString)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g *Generator) outputJSON(data map[string]any) (string, error) {
	for k, v := range data {
		for _, token := range g.tokens {
			// If the token has a display name, use the display name as the key.
			if token.Name == k && token.Display != "" {
				data[token.Display] = v
				delete(data, k)
			}
		}
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func (g *Generator) timestamp(timestampNanoseconds int64) string {
	timestamp := time.Unix(0, timestampNanoseconds)
	switch g.timeRange.timeFormat {
	case string(TimestampFormatApache):
		return timestamp.In(g.timeRange.location).Format(Apache)
	case string(TimestampFormatApacheError):
		return timestamp.In(g.timeRange.location).Format(ApacheError)
	case string(TimestampFormatRFC3164):
		return timestamp.In(g.timeRange.location).Format(RFC3164)
	case string(TimestampFormatRFC5424):
		return timestamp.In(g.timeRange.location).Format(RFC5424)
	case string(TimestampFormatCommonLog):
		return timestamp.In(g.timeRange.location).Format(CommonLog)
	case string(TimestampFormatClickHouse):
		return timestamp.In(g.timeRange.location).Format(ClickHouse)
	case string(TimestampFormatRFC3339):
		return timestamp.In(g.timeRange.location).Format(time.RFC3339)
	case string(TimestampFormatUnix):
		return fmt.Sprintf("%d", timestamp.Unix())
	default:
		return timestamp.In(g.timeRange.location).Format(g.timeRange.timeFormat)
	}
}
