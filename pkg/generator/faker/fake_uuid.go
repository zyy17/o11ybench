package faker

import (
	"github.com/brianvoe/gofakeit"
)

// FakeUUIDOptions is the options for generating the fake uuid.
type FakeUUIDOptions struct {
	// TODO(zyy17)
}

// FakeUUID generates a fake uuid.
func FakeUUID(opts Options) (string, error) {
	var options FakeUUIDOptions
	if err := parseOptions(opts, &options); err != nil {
		return "", err
	}

	return gofakeit.UUID(), nil
}
