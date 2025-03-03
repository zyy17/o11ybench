package faker

import (
	"math/rand"
)

// FakeHTTPVersionOptions is the options for the fake HTTP version.
type FakeHTTPVersionOptions struct {
	// TODO(zyy17)
}

// FakeHTTPVersion generates a fake HTTP version.
func FakeHTTPVersion(opts Options) (string, error) {
	var options FakeHTTPVersionOptions
	if err := parseOptions(opts, &options); err != nil {
		return "", err
	}

	return randomHTTPVersion(), nil
}

func randomHTTPVersion() string {
	versions := []string{"HTTP/1.0", "HTTP/1.1", "HTTP/2.0"}
	return versions[rand.Intn(3)]
}
