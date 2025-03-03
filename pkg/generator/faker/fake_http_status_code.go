package faker

import (
	"github.com/brianvoe/gofakeit"
)

// FakeUserIDOptions is the options for the fake user ID.
type FakeHTTPStatusCodeOptions struct {
	// TODO(zyy17)
}

// FakeHTTPStatusCode generates a fake HTTP status code.
func FakeHTTPStatusCode(opts Options) (int, error) {
	var options FakeHTTPStatusCodeOptions
	if err := parseOptions(opts, &options); err != nil {
		return 0, err
	}

	return gofakeit.StatusCode(), nil
}
