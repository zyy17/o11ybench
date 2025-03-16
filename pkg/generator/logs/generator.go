package logs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"maps"
	"text/template"
	"time"

	"github.com/zyy17/o11ybench/pkg/generator/common"
	"github.com/zyy17/o11ybench/pkg/generator/faker"
	"github.com/zyy17/o11ybench/pkg/generator/logs/templates"
	"github.com/zyy17/o11ybench/pkg/generator/logs/types"
)

// LogsGenerator is the generator for the logs.
type LogsGenerator struct {
	cfg     *types.LogsGeneratorConfig
	timeCfg *common.TimeConfig
}

// NewLogsGenerator creates a new LogsGenerator.
func NewLogsGenerator(cfg *types.LogsGeneratorConfig, timeCfg *common.TimeConfig) (*LogsGenerator, error) {
	return &LogsGenerator{cfg: cfg, timeCfg: timeCfg}, nil
}

func (g *LogsGenerator) Generate(opts *types.GeneratorOptions) ([]byte, error) {
	if opts != nil && opts.LogsCount > 0 {
		if opts.Timestamp.IsZero() {
			return nil, fmt.Errorf("timestamp is required")
		}

		logs, err := g.generateMultipleLogs(opts.LogsCount, opts.Timestamp, g.timeCfg)
		if err != nil {
			return nil, err
		}

		return logs, nil
	}

	if g.cfg.Output != nil {
		if g.timeCfg == nil || g.timeCfg.Range == nil {
			logs, err := g.generateMultipleLogs(g.cfg.Output.Count, time.Now(), g.timeCfg)
			if err != nil {
				return nil, err
			}

			return logs, nil
		}

		var (
			logs  = make([]byte, 0)
			start = g.timeCfg.Range.Start
			end   = g.timeCfg.Range.End
		)
		if g.cfg.Output.Count > 0 && g.cfg.Output.Interval > 0 {
			expectedEnd := start.Add(g.cfg.Output.Interval * time.Duration(g.cfg.Output.Count))
			if !expectedEnd.After(g.timeCfg.Range.End) {
				end = expectedEnd
			}
		}

		current := start
		for current.Before(end) {
			log, err := g.generateOneLineLog(current, g.timeCfg)
			if err != nil {
				return nil, err
			}

			// Add the newline to the log.
			log = append(log, '\n')

			logs = append(logs, log...)

			current = current.Add(g.cfg.Output.Interval)
		}

		return logs, nil
	}

	return nil, nil
}

func (g *LogsGenerator) generateMultipleLogs(count int, timestamp time.Time, timeCfg *common.TimeConfig) ([]byte, error) {
	logs := make([]byte, 0)

	for i := 0; i < count; i++ {
		log, err := g.generateOneLineLog(timestamp, timeCfg)
		if err != nil {
			return nil, err
		}

		// Add the newline to the log.
		log = append(log, '\n')

		logs = append(logs, log...)
	}

	return logs, nil
}

func (g *LogsGenerator) generateOneLineLog(timestamp time.Time, timeCfg *common.TimeConfig) ([]byte, error) {
	// The logs tokens that are from the config.
	generatedData, err := generateTokenValues(g.cfg.Tokens)
	if err != nil {
		return nil, err
	}

	// Set the timestamp.
	generatedData[templates.ReservedTokenNameTimestamp] = common.OutputTimestamp(timestamp, timeCfg.TimestampFormat)

	if g.cfg.Format.Type == types.LogFormatTypeJSON {
		return g.jsonOutput(generatedData)
	}

	if g.cfg.Format.Custom != "" {
		return g.templateOutput(g.cfg.Format.Custom, generatedData)
	}

	// The logs tokens that are from the builtin templates.
	if g.cfg.Format.Type != "" {
		builtinTemplate, ok := templates.LogFormats[g.cfg.Format.Type]
		if !ok {
			return nil, fmt.Errorf("invalid format: '%s'", g.cfg.Format.Type)
		}

		builtinGeneratedData, err := generateTokenValues(builtinTemplate.Tokens)
		if err != nil {
			return nil, err
		}

		// Set the timestamp.
		if timeCfg.TimestampFormat.Type == "" && timeCfg.TimestampFormat.Custom == "" && builtinTemplate.TimestampFormat != "" {
			timeCfg.TimestampFormat.Type = builtinTemplate.TimestampFormat
		}
		builtinGeneratedData[templates.ReservedTokenNameTimestamp] = common.OutputTimestamp(timestamp, timeCfg.TimestampFormat)

		// Merge the generated data with the builtin generated data.
		maps.Copy(generatedData, builtinGeneratedData)

		return g.templateOutput(builtinTemplate.Template, generatedData)
	}

	return nil, fmt.Errorf("can't find a valid log format")
}

func (g *LogsGenerator) templateOutput(templateString string, data map[string]any) ([]byte, error) {
	tmpl, err := template.New("output").Parse(templateString)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (g *LogsGenerator) jsonOutput(data map[string]any) ([]byte, error) {
	for k, v := range data {
		for _, token := range g.cfg.Tokens {
			// If the token has a display name, use the display name as the key.
			if token.Name == k && token.Display != "" {
				data[token.Display] = v
				delete(data, k)
			}
		}
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func generateTokenValues(tokens []*types.LogToken) (map[string]any, error) {
	generatedData := make(map[string]any)

	for _, token := range tokens {
		var (
			value any
			err   error
		)

		if token.Value != nil {
			value = token.Value
		} else {
			value, err = faker.Fake(token.Type, token.FakeConfig)
			if err != nil {
				return nil, err
			}
		}

		generatedData[token.Name] = value
	}

	return generatedData, nil
}
