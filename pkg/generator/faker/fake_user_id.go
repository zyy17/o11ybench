package faker

import (
	"github.com/brianvoe/gofakeit"
)

// FakeUserIDOptions is the options for the fake user ID.
type FakeUserIDOptions struct {
	// TODO(zyy17)
}

// FakeUserID generates a fake user ID.
func FakeUserID(opts Options) (string, error) {
	var options FakeUserIDOptions
	if err := parseOptions(opts, &options); err != nil {
		return "", err
	}

	return gofakeit.Username(), nil
}
