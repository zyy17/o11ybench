package common

import (
	"testing"
	"time"
)

func TestOutputTimestamp(t *testing.T) {
	testStartTime, _ := time.Parse(time.RFC3339, "2025-03-23T00:00:00Z")
	timestampFormat := &TimestampFormat{
		Zone: "UTC",
	}

	tests := []struct {
		name     string
		format   TimestampFormatType
		custom   string
		expected string
	}{
		{
			name:     "Test apache format",
			format:   TimestampFormatTypeApache,
			expected: "23/Mar/2025:00:00:00 +0000",
		},
		{
			name:     "Test apache_error format",
			format:   TimestampFormatTypeApacheError,
			expected: "Sun Mar 23 00:00:00 2025",
		},
		{
			name:     "Test rfc3164 format",
			format:   TimestampFormatTypeRFC3164,
			expected: "Mar 23 00:00:00",
		},
		{
			name:     "Test rfc5424 format",
			format:   TimestampFormatTypeRFC5424,
			expected: "2025-03-23T00:00:00.000Z",
		},
		{
			name:     "Test rfc3339 format",
			format:   TimestampFormatTypeRFC3339,
			expected: "2025-03-23T00:00:00Z",
		},
		{
			name:     "Test unix_seconds format",
			format:   TimestampFormatTypeUnix,
			expected: "1742688000",
		},
		{
			name:     "Test custom format",
			custom:   "2006/01/02|15:04:05",
			expected: "2025/03/23|00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timestampFormat.Type = tt.format
			timestampFormat.Custom = tt.custom
			actual := OutputTimestamp(testStartTime, timestampFormat)
			if actual != tt.expected {
				t.Errorf("expected '%s', but got '%s'", tt.expected, actual)
			}
		})
	}
}
