package templates

// ReservedTokenName is the reserved token name for the fake data. It only take effect in specific format.
const (
	// ReservedTokenNameHost is the reserved token name for the host.
	// It only take effect in `FormatApacheCommonLog` and `FormatApacheCombinedLog`.
	ReservedTokenNameHost string = "host"

	// ReservedTokenNameUserID is the reserved token name for the user ID.
	// It only take effect in `FormatApacheCommonLog` and `FormatApacheCombinedLog`.
	ReservedTokenNameUserID string = "userID"

	// ReservedTokenNameHTTPMethod is the reserved token name for the HTTP method.
	// It only take effect in `FormatApacheCommonLog` and `FormatApacheCombinedLog`.
	ReservedTokenNameHTTPMethod string = "httpMethod"

	// ReservedTokenNameHTTPVersion is the reserved token name for the HTTP version.
	// It only take effect in `FormatApacheCommonLog` and `FormatApacheCombinedLog`.
	ReservedTokenNameHTTPVersion string = "httpVersion"

	// ReservedTokenNameHTTPStatusCode is the reserved token name for the HTTP status code.
	// It only take effect in `FormatApacheCommonLog` and `FormatApacheCombinedLog`.
	ReservedTokenNameHTTPStatusCode string = "httpStatusCode"

	// ReservedTokenNameHTTPURL is the reserved token name for the HTTP URL.
	// It only take effect in `FormatApacheCommonLog` and `FormatApacheCombinedLog`.
	ReservedTokenNameHTTPURL string = "httpURL"

	// ReservedTokenNameHTTPContentLength is the reserved token name for the HTTP content length.
	// It only take effect in `FormatApacheCommonLog` and `FormatApacheCombinedLog`.
	ReservedTokenNameHTTPContentLength string = "httpContentLength"

	// ReservedTokenNameHTTPUserAgent is the reserved token name for the HTTP user agent.
	// It only take effect in `FormatApacheCombinedLog`.
	ReservedTokenNameHTTPUserAgent string = "httpUserAgent"

	// ReservedTokenNameReferer is the reserved token name for the referer.
	// It only take effect in `FormatApacheCombinedLog`.
	ReservedTokenNameReferer string = "referer"
)

// ReservedTokenNameTimestamp is the reserved token name for the timestamp.
const (
	ReservedTokenNameTimestamp string = "timestamp"
)
