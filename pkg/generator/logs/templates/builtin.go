package templates

import (
	"github.com/zyy17/o11ybench/pkg/generator/common"
	"github.com/zyy17/o11ybench/pkg/generator/logs/types"
)

// BuiltinLogFormat is the internal representation of the built-in log format, for example, ApacheCombinedLog, RFC3164, etc.
type BuiltinLogFormat struct {
	Template        string
	Tokens          []*types.LogToken
	TimestampFormat common.TimestampFormatType
}

// LogFormatType -> BuiltinLogFormat.
var LogFormats = map[types.LogFormatType]BuiltinLogFormat{
	types.LogFormatTypeApacheCombinedLog: ApacheCombinedLog,
	types.LogFormatTypeApacheCommonLog:   ApacheCommonLog,
	types.LogFormatTypeApacheErrorLog:    ApacheErrorLog,
	types.LogFormatTypeRFC3164:           RFC3164Log,
	types.LogFormatTypeRFC5424:           RFC5424Log,
}
