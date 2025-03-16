package faker

import (
	"github.com/brianvoe/gofakeit"

	"github.com/zyy17/o11ybench/pkg/generator/common"
)

// FakeHackerPhraseOptions is the options for generating the fake hacker phrase.
type FakeHackerPhraseOptions struct {
	// TODO(zyy17)
}

// FakeHackerPhrase generates a fake hacker phrase.
func FakeHackerPhrase(_ common.ElementType, opts Options) (string, error) {
	var options FakeHackerPhraseOptions
	if err := parseOptions(opts, &options); err != nil {
		return "", err
	}

	return gofakeit.HackerPhrase(), nil
}
