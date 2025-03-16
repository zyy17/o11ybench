package common

import (
	"fmt"
	"time"
)

const (
	// TimezoneLocal is a special value that means use the current system timezone.
	TimezoneLocal = "Local"
)

// TimeConfig is the configuration for the time of the data to be generated.
type TimeConfig struct {
	// Range is the configuration for the time range of the data to be generated.
	Range *TimeRange `yaml:"range,omitempty"`

	// TimestampFormat is the configuration for the timestamp format of the data to be generated.
	TimestampFormat *TimestampFormat `yaml:"timestamp,omitempty"`
}

// Validate validates the time configuration.
func (t *TimeConfig) Validate() error {
	if t.Range != nil {
		if err := t.Range.validate(); err != nil {
			return err
		}
	}

	if t.TimestampFormat != nil {
		if err := t.TimestampFormat.validate(); err != nil {
			return err
		}
	}

	return nil
}

// Defaults returns the default time configuration.
func (t TimeConfig) Defaults() *TimeConfig {
	return &TimeConfig{
		TimestampFormat: &TimestampFormat{
			Zone: TimezoneLocal,
		},
	}
}

// TimeRange is the configuration for the time range of the data to be generated.
// If not set, the data will be generated for the current time.
type TimeRange struct {
	// Start is the start time of the time range.
	Start time.Time `yaml:"start"`

	// End is the end time of the time range.
	End time.Time `yaml:"end"`
}

func (t *TimeRange) validate() error {
	if t.Start.IsZero() {
		return fmt.Errorf("start is required")
	}

	if t.End.IsZero() {
		return fmt.Errorf("end is required")
	}

	if t.Start.After(t.End) {
		return fmt.Errorf("start must be before end")
	}

	return nil
}

// TimestampFormat is the output timestamp format of the log.
type TimestampFormat struct {
	// Type is the type of the timestamp format.
	Type TimestampFormatType `yaml:"type,omitempty"`

	// Custom is the custom format of the timestamp.
	Custom string `yaml:"custom,omitempty"`

	// Zone is the timezone of the timestamp. If not set, it will use current system timezone.
	Zone string `yaml:"zone,omitempty"`
}

func (t *TimestampFormat) validate() error {
	if t.Zone != "" && t.Zone != TimezoneLocal {
		_, err := time.LoadLocation(t.Zone)
		if err != nil {
			return fmt.Errorf("invalid timezone '%s': %w", t.Zone, err)
		}
	}

	return nil
}

// TimestampFormatType is the type of the timestamp format.
type TimestampFormatType string

const (
	// TimestampFormatTypeApache is the format of the apache timestamp.
	TimestampFormatTypeApache TimestampFormatType = "apache"

	// TimestampFormatTypeApacheError is the format of the apache error timestamp.
	TimestampFormatTypeApacheError TimestampFormatType = "apache_error"

	// TimestampFormatTypeRFC3164 is the format of the rfc3164 timestamp.
	TimestampFormatTypeRFC3164 TimestampFormatType = "rfc3164"

	// TimestampFormatTypeRFC5424 is the format of the rfc5424 timestamp.
	TimestampFormatTypeRFC5424 TimestampFormatType = "rfc5424"

	// TimestampFormatTypeRFC3339 is the format of the rfc3339 timestamp.
	TimestampFormatTypeRFC3339 TimestampFormatType = "rfc3339"

	// TimestampFormatTypeUnix is the format of the unix timestamp in seconds.
	TimestampFormatTypeUnix TimestampFormatType = "unix_seconds"
)

// OutputTimestamp outputs the timestamp in the given format.
func OutputTimestamp(input time.Time, cfg *TimestampFormat) string {
	var loc *time.Location
	if cfg.Zone == TimezoneLocal {
		loc = time.Local
	} else {
		// We can safely ignore the error here since we've already validated the timezone in the configuration
		loc, _ = time.LoadLocation(cfg.Zone)
	}

	// It's high priority to use the custom format.
	if cfg.Custom != "" {
		return input.In(loc).Format(cfg.Custom)
	}

	// If the timestamp format is one of the built-in formats.
	format, ok := timestampFormats[cfg.Type]
	if ok {
		return input.In(loc).Format(format)
	}

	if cfg.Type == TimestampFormatTypeUnix {
		return fmt.Sprintf("%d", input.Unix())
	}

	// Use RFC3339 as the default format.
	return input.Format(time.RFC3339)
}

// TimestampFormatType -> Specific timestamp format string.
var timestampFormats = map[TimestampFormatType]string{
	TimestampFormatTypeApache:      Apache,
	TimestampFormatTypeApacheError: ApacheError,
	TimestampFormatTypeRFC3164:     RFC3164,
	TimestampFormatTypeRFC5424:     RFC5424,
}

const (
	Apache      = "02/Jan/2006:15:04:05 -0700"
	ApacheError = "Mon Jan 02 15:04:05 2006"
	RFC3164     = "Jan 02 15:04:05"
	RFC5424     = "2006-01-02T15:04:05.000Z"
)
