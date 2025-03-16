package faker

import (
	"github.com/brianvoe/gofakeit"

	"github.com/zyy17/o11ybench/pkg/generator/common"
)

// FakeIPv4Options is the options for the fake IPv4 address.
type FakeIPv4Options struct {
	// TODO(zyy17)
}

// FakeIPv4 generates a fake IPv4 address.
func FakeIPv4(_ common.ElementType, opts Options) (string, error) {
	var options FakeIPv4Options
	if err := parseOptions(opts, &options); err != nil {
		return "", err
	}

	return gofakeit.IPv4Address(), nil
}
