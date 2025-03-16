package faker

import (
	"github.com/brianvoe/gofakeit"

	"github.com/zyy17/o11ybench/pkg/generator/common"
)

// FakeUsernameOptions is the options for the fake username.
type FakeUsernameOptions struct {
	// TODO(zyy17)
}

// FakeUsername generates a fake user ID.
func FakeUsername(_ common.ElementType, opts Options) (string, error) {
	var options FakeUsernameOptions
	if err := parseOptions(opts, &options); err != nil {
		return "", err
	}

	return gofakeit.Username(), nil
}
