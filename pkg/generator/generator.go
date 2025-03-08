package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"sort"
	"strings"
	"sync"
	"text/template"
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

	// Count is the total number of logs to generate.
	Count int64 `yaml:"count,omitempty" json:"count,omitempty"`

	// IntervalCount is the number of logs to generate per interval. It's mutually exclusive with `Count`.
	IntervalCount int64 `yaml:"intervalCount,omitempty" json:"intervalCount,omitempty"`

	// Parallel is the number of parallel workers that used to generate the fake data.
	Parallel int `yaml:"parallel,omitempty" json:"parallel,omitempty"`

	// DisableTimestamp is used to disable the timestamp of the fake data.
	DisableTimestamp *bool `yaml:"disableTimestamp,omitempty" json:"disableTimestamp,omitempty"`
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

	// LogFormatApacheErrorLog is the format of apache error log.
	LogFormatApacheErrorLog LogFormat = "ApacheErrorLog"

	// LogFormatRFC3164 is the format of rfc3164 log.
	LogFormatRFC3164 LogFormat = "RFC3164"

	// LogFormatRFC5424 is the format of rfc5424 log.
	LogFormatRFC5424 LogFormat = "RFC5424"

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
	TimestampFormatUnix        TimestampFormat = "UnixSeconds"
)

const (
	Apache      = "02/Jan/2006:15:04:05 -0700"
	ApacheError = "Mon Jan 02 15:04:05 2006"
	RFC3164     = "Jan 02 15:04:05"
	RFC5424     = "2006-01-02T15:04:05.000Z"
)

const (
	// DefaultCount is the default number of logs to generate.
	DefaultCount = 100

	// DefaultTimezone is the default timezone of the time range.
	DefaultTimezone = "UTC"

	// DefaultInterval is the default interval for each log.
	DefaultInterval = "5s"

	// DefaultTimeFormat is the default time format of the time range.
	DefaultTimeFormat = string(TimestampFormatRFC3339)
)

// Generator is the interface for load data generator.
type Generator interface {
	Generate() ([]byte, error)
}

var _ Generator = &generator{}

type generator struct {
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

func NewGenerator(config *FakeDataConfig) (Generator, error) {
	// Set the default values for the config.
	config.setDefault()

	timeRange, err := parseTimeRange(&config.TimeRange)
	if err != nil {
		return nil, err
	}

	if !timeRange.end.IsZero() {
		// Calculate the total number of logs to generate.
		totalCount := (timeRange.end.UnixNano()-timeRange.start.UnixNano())/int64(timeRange.interval) + 1

		if config.Output.Count > 0 && config.Output.Count < totalCount {
			timeRange.end = timeRange.start.Add(timeRange.interval * time.Duration(config.Output.Count))
		} else {
			config.Output.Count = totalCount
		}
	} else {
		if config.Output.Count == 0 {
			// Default count is 100.
			config.Output.Count = 100
		}

		timeRange.end = timeRange.start.Add(timeRange.interval * time.Duration(config.Output.Count))
	}

	return &generator{
		tokens:    config.Tokens,
		output:    &config.Output,
		timeRange: timeRange,
	}, nil
}

func NewGeneratorFromFile(configFile string) (Generator, error) {
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

func (c *FakeDataConfig) setDefault() {
	if c.Output.Count == 0 && c.TimeRange.End == "" {
		// Limit the default count to 100 to avoid generating too many logs.
		c.Output.Count = DefaultCount
	}

	if c.TimeRange.Timezone == "" {
		c.TimeRange.Timezone = DefaultTimezone
	}

	if c.TimeRange.Interval == "" {
		c.TimeRange.Interval = DefaultInterval
	}

	if c.TimeRange.Format == "" {
		if c.Output.LogFormat == LogFormatApacheCommonLog || c.Output.LogFormat == LogFormatApacheCombinedLog {
			c.TimeRange.Format = string(TimestampFormatApache)
		} else if c.Output.LogFormat == LogFormatApacheErrorLog {
			c.TimeRange.Format = string(TimestampFormatApacheError)
		} else if c.Output.LogFormat == LogFormatRFC3164 {
			c.TimeRange.Format = string(TimestampFormatRFC3164)
		} else if c.Output.LogFormat == LogFormatRFC5424 {
			c.TimeRange.Format = string(TimestampFormatRFC5424)
		} else {
			c.TimeRange.Format = DefaultTimeFormat
		}
	}

	if c.TimeRange.Start == "" {
		// 6 hours ago
		c.TimeRange.Start = time.Now().Add(-6 * time.Hour).Format(time.RFC3339)
	}
}

func (g *generator) Generate() ([]byte, error) {
	var (
		chunks = make([]*chunk, 0)
	)

	if g.output.Parallel > 0 {
		wg := sync.WaitGroup{}
		chunks = make([]*chunk, g.output.Parallel)
		count := g.output.Count / int64(g.output.Parallel)
		for i := 0; i < g.output.Parallel; i++ {
			if i == g.output.Parallel-1 {
				count += g.output.Count % int64(g.output.Parallel)
			}
			start := g.timeRange.start.Add(time.Duration(i) * g.timeRange.interval * time.Duration(count))
			end := start.Add(g.timeRange.interval * time.Duration(count))
			wg.Add(1)
			go func() {
				defer wg.Done()
				chunk, err := g.doGenerate(start, end)
				if err != nil {
					fmt.Printf("Generate [%d] chunk failed: %v\n", i, err)
					return
				}

				// Write the chunk by worker index to avoid using mutex.
				chunks[i] = chunk
			}()
		}
		wg.Wait()
	} else {
		chunk, err := g.doGenerate(g.timeRange.start, g.timeRange.end)
		if err != nil {
			return nil, err
		}
		chunks = append(chunks, chunk)
	}

	var output []string
	if len(chunks) == 1 {
		output = chunks[0].logs
	} else {
		logs := make([]string, 0)
		// Sort the chunks by start time and merge them.
		sort.Slice(chunks, func(i, j int) bool {
			return chunks[i].start.Before(chunks[j].start)
		})

		for _, chunk := range chunks {
			logs = append(logs, chunk.logs...)
		}

		output = logs
	}

	return []byte(strings.Join(output, "\n")), nil
}

func parseTimeRange(cfg *TimeRange) (*timeRange, error) {
	var tr timeRange

	location, err := time.LoadLocation(cfg.Timezone)
	if err != nil {
		return nil, err
	}
	tr.location = location

	start, err := time.Parse(time.RFC3339, cfg.Start)
	if err != nil {
		return nil, err
	}
	tr.start = start

	if cfg.End != "" {
		end, err := time.Parse(time.RFC3339, cfg.End)
		if err != nil {
			return nil, err
		}

		if tr.start.After(end) {
			return nil, fmt.Errorf("start time '%s' is after end time '%s'", tr.start, end)
		}

		tr.end = end
	}

	if cfg.Interval != "" {
		interval, err := time.ParseDuration(cfg.Interval)
		if err != nil {
			return nil, err
		}
		tr.interval = interval
	}

	tr.timeFormat = cfg.Format

	return &tr, nil
}

func (g *generator) generateByTemplate(overrideData map[string]any, tokens []*faker.Token, templateString string) (string, error) {
	data := make(map[string]any)

	for _, token := range tokens {
		var (
			value any
			err   error
		)

		if token.Value != nil {
			value = token.Value
		} else {
			value, err = faker.Fake(&token.FakeConfig)
			if err != nil {
				return "", err
			}
		}

		data[token.Name] = value
	}

	maps.Copy(data, overrideData)

	return g.templateOutput(templateString, data)
}

func (g *generator) templateOutput(templateString string, data map[string]any) (string, error) {
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

func (g *generator) outputJSON(data map[string]any) (string, error) {
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

func (g *generator) timestamp(timestamp time.Time) string {
	switch g.timeRange.timeFormat {
	case string(TimestampFormatApache):
		return timestamp.In(g.timeRange.location).Format(Apache)
	case string(TimestampFormatApacheError):
		return timestamp.In(g.timeRange.location).Format(ApacheError)
	case string(TimestampFormatRFC3164):
		return timestamp.In(g.timeRange.location).Format(RFC3164)
	case string(TimestampFormatRFC5424):
		return timestamp.In(g.timeRange.location).Format(RFC5424)
	case string(TimestampFormatRFC3339):
		return timestamp.In(g.timeRange.location).Format(time.RFC3339)
	case string(TimestampFormatUnix):
		return fmt.Sprintf("%d", timestamp.Unix())
	default:
		return timestamp.In(g.timeRange.location).Format(g.timeRange.timeFormat)
	}
}

type chunk struct {
	start time.Time
	end   time.Time
	logs  []string
}

func (g *generator) doGenerate(start time.Time, end time.Time) (*chunk, error) {
	var (
		logs    = make([]string, 0)
		current = start
	)

	for {
		if current.Equal(end) {
			break
		}

		var (
			generatedData = make(map[string]any)

			log string
			err error
		)

		current = current.Add(g.timeRange.interval)
		for _, token := range g.tokens {
			var value any
			if token.Value != nil {
				value = token.Value
			} else if token.Name != templates.ReservedTokenNameTimestamp {
				value, err = faker.Fake(&token.FakeConfig)
				if err != nil {
					return nil, err
				}
			}

			generatedData[token.Name] = value
		}

		// Set the timestamp.
		if g.output.DisableTimestamp == nil || !*g.output.DisableTimestamp {
			generatedData[templates.ReservedTokenNameTimestamp] = g.timestamp(current)
		}

		if g.output.Custom != "" {
			log, err = g.templateOutput(g.output.Custom, generatedData)
			if err != nil {
				return nil, err
			}
			logs = append(logs, log)
			continue
		}

		if g.output.IntervalCount > 0 {
			for i := 0; i < int(g.output.IntervalCount); i++ {
				log, err := g.generateLogs(generatedData)
				if err != nil {
					return nil, err
				}
				logs = append(logs, log)
			}
		} else {
			log, err := g.generateLogs(generatedData)
			if err != nil {
				return nil, err
			}
			logs = append(logs, log)
		}
	}

	return &chunk{
		logs:  logs,
		start: start,
		end:   end,
	}, nil
}

func (g *generator) generateLogs(generatedData map[string]any) (string, error) {
	switch g.output.LogFormat {
	case LogFormatApacheCommonLog:
		log, err := g.generateByTemplate(generatedData, templates.ApacheCommonLogTokens, templates.ApacheCommonLogTemplate)
		if err != nil {
			return "", err
		}
		return log, nil
	case LogFormatApacheCombinedLog:
		log, err := g.generateByTemplate(generatedData, templates.ApacheCombinedLogTokens, templates.ApacheCombinedLogTemplate)
		if err != nil {
			return "", err
		}
		return log, nil
	case LogFormatApacheErrorLog:
		log, err := g.generateByTemplate(generatedData, templates.ApacheErrorLogTokens, templates.ApacheErrorLogTemplate)
		if err != nil {
			return "", err
		}
		return log, nil
	case LogFormatRFC3164:
		log, err := g.generateByTemplate(generatedData, templates.RFC3164LogTokens, templates.RFC3164LogTemplate)
		if err != nil {
			return "", err
		}
		return log, nil
	case LogFormatRFC5424:
		log, err := g.generateByTemplate(generatedData, templates.RFC5424LogTokens, templates.RFC5424LogTemplate)
		if err != nil {
			return "", err
		}
		return log, nil
	case LogFormatJSON:
		log, err := g.outputJSON(generatedData)
		if err != nil {
			return "", err
		}
		return log, nil
	}

	return "", nil
}
