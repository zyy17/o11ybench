package faker

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strings"

	"github.com/brianvoe/gofakeit"
	"github.com/zyy17/o11ybench/pkg/utils"
)

const (
	// DefaultSeparator is the default separator between fake words.
	DefaultSeparator = " "
)

// FakeWordsOptions is the options to generate fake words.
type FakeWordsOptions struct {
	// Count is the number of fake words to generate.
	// It's exclusive with Size, SizeRange and SizeRangeByPossibility options.
	Count int `yaml:"count,omitempty" json:"count,omitempty"`

	// Size is the approximate total size of all generated fake words. It can be in format "100bytes", "1kib", "1mib", etc.
	// The size should be a positive integer.
	// It's exclusive with Count, SizeRange and SizeRangeByPossibility options.
	Size string `yaml:"size,omitempty" json:"size,omitempty"`

	// SizeRange specifies a approximate total size range of all generated fake words.
	// It can be in format of <min-size>-<max-size>, for example: "100bytes-200bytes", "1kib-2kib", "1mib-2mib", etc.
	// The size should be a positive integer.
	// It's exclusive with Count, Size and SizeRangeByPossibility options.
	SizeRange string `yaml:"sizeRange,omitempty" json:"sizeRange,omitempty"`

	// SizeRangeWithPossibility specifies multiple approximate total size range of all generated fake words with their probabilities.
	// Each entry should be in format "<probability>%:<min-size>-<max-size>" (e.g. "10%:30bytes-40bytes").
	// The size should be a positive integer and each range should not overlap.
	// The sum of all probabilities must equal 100%.
	// It's exclusive with Count, Size and SizeRange options.
	SizeRangeWithPossibility []string `yaml:"sizeRangeWithPossibility,omitempty" json:"sizeRangeWithPossibility,omitempty"`

	// Separator is the string used to separate fake words. Defaults to a single space if not set.
	Separator string `yaml:"separator,omitempty" json:"separator,omitempty"`

	// FixedWords is the list of fixed words to be used.
	FixedWords []string `yaml:"fixedWords,omitempty" json:"fixedWords,omitempty"`
}

func (o *FakeWordsOptions) validate() error {
	sizeMethods := 0

	if o.Count < 0 {
		return fmt.Errorf("count should be a positive integer")
	}

	if o.Count > 0 {
		sizeMethods++
	}

	if o.Size != "" {
		sizeMethods++
		if sizeMethods > 1 {
			return fmt.Errorf("only one of count, size, sizeRange or sizeRangeWithPossibility should be set")
		}

		sizeInBytes, err := utils.ParseSize(o.Size)
		if err != nil {
			return fmt.Errorf("invalid size: %s", o.Size)
		}

		if sizeInBytes <= 0 {
			return fmt.Errorf("size should be a positive integer")
		}
	}

	if o.SizeRange != "" {
		sizeMethods++
		if sizeMethods > 1 {
			return fmt.Errorf("only one of count, size, sizeRange or sizeRangeWithPossibility should be set")
		}

		min, max, err := parseSizeRange(o.SizeRange)
		if err != nil {
			return fmt.Errorf("invalid size range: %s", o.SizeRange)
		}

		if min <= 0 || max <= 0 {
			return fmt.Errorf("size range should be a positive integer")
		}

		if min >= max {
			return fmt.Errorf("min should be less than max")
		}
	}

	if len(o.SizeRangeWithPossibility) > 0 {
		sizeMethods++
		if sizeMethods > 1 {
			return fmt.Errorf("only one of count, size, sizeRange or sizeRangeWithPossibility should be set")
		}

		sizeRangeWithProbabilities, err := parseSizeRangeWithProbability(o.SizeRangeWithPossibility)
		if err != nil {
			return fmt.Errorf("invalid size range with probability: %s", o.SizeRangeWithPossibility)
		}

		if err := validateSizeRangeWithProbability(sizeRangeWithProbabilities); err != nil {
			return fmt.Errorf("invalid size range with probability: %s", o.SizeRangeWithPossibility)
		}
	}

	if len(o.FixedWords) > 0 {
		sizeMethods++
		if sizeMethods > 1 {
			return fmt.Errorf("only one of count, size, sizeRange, sizeRangeWithPossibility or fixedWords should be set")
		}
	}

	if sizeMethods == 0 {
		return fmt.Errorf("should set at least one of count, size, sizeRange or sizeRangeWithPossibility")
	}

	return nil
}

// FakeWords generates fake words with the given options.
func FakeWords(opts Options) (string, error) {
	var options FakeWordsOptions
	if err := parseOptions(opts, &options); err != nil {
		return "", err
	}

	if err := options.validate(); err != nil {
		return "", err
	}

	var (
		words     []string
		separator = DefaultSeparator
	)

	if options.Separator != "" {
		separator = options.Separator
	}

	if options.Count > 0 {
		words = fakeNWords(options.Count)
	} else if options.Size != "" {
		words = fakeWordsWithSize(options.Size, separator)
	} else if options.SizeRange != "" {
		words = fakeWordsWithSizeRange(options.SizeRange, separator)
	} else if len(options.SizeRangeWithPossibility) > 0 {
		words = fakeWordsWithSizeRangeWithProbability(options.SizeRangeWithPossibility, separator)
	} else if len(options.FixedWords) > 0 {
		choose := options.FixedWords[rand.Intn(len(options.FixedWords))]
		words = []string{choose}
	} else {
		return "", fmt.Errorf("no valid options provided")
	}

	if len(words) == 1 {
		return words[0], nil
	}

	return strings.Join(words, separator), nil
}

func fakeNWords(n int) []string {
	words := make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = fakeWord()
	}
	return words
}

func fakeWordsWithSize(size, separator string) []string {
	// Already validated before calling this function.
	bytes, _ := utils.ParseSize(size)
	return doFakeWordsWithSize(bytes, separator)
}

func fakeWordsWithSizeRange(sizeRange string, separator string) []string {
	// Already validated before calling this function.
	min, max, _ := parseSizeRange(sizeRange)
	return doFakeWordsWithSize(randomNumber(min, max), separator)
}

func fakeWordsWithSizeRangeWithProbability(sizeRangeWithProbability []string, separator string) []string {
	// Already validated before calling this function.
	sizeRangeWithProbabilities, _ := parseSizeRangeWithProbability(sizeRangeWithProbability)
	return doFakeWordsWithSize(generateSizeByProbability(sizeRangeWithProbabilities), separator)
}

func doFakeWordsWithSize(sizeInBytes int64, separator string) []string {
	var (
		words       []string
		currentSize int64
	)

	for currentSize < sizeInBytes {
		word := fakeWord()
		if len(words) > 0 {
			currentSize += int64(len(separator))
		}
		currentSize += int64(len(word))
		words = append(words, word)
	}

	return words
}

func fakeWord() string {
	return gofakeit.Word()
}

func parseSizeRange(sizeRange string) (int64, int64, error) {
	parts := strings.Split(sizeRange, "-")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid size range: %s", sizeRange)
	}

	min, err := utils.ParseSize(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, 0, fmt.Errorf("invalid size range: %s", sizeRange)
	}

	max, err := utils.ParseSize(strings.TrimSpace(parts[1]))
	if err != nil {
		return 0, 0, fmt.Errorf("invalid size range: %s", sizeRange)
	}

	return min, max, nil
}

func generateSizeByProbability(sizeRangeWithProbabilities []*sizeRangeWithProbability) int64 {
	sort.Slice(sizeRangeWithProbabilities, func(i, j int) bool {
		return sizeRangeWithProbabilities[i].min < sizeRangeWithProbabilities[j].min &&
			sizeRangeWithProbabilities[i].max < sizeRangeWithProbabilities[j].max
	})

	// Calculate the cumulative probability that the generated size will fall into the range.
	for i, sr := range sizeRangeWithProbabilities {
		if i > 0 {
			sr.probability += sizeRangeWithProbabilities[i-1].probability
		}
	}

	// Generate a random number in [0.0, 1.0).
	num := rand.Float64()

	for _, sizeRangeWithProbability := range sizeRangeWithProbabilities {
		if num < sizeRangeWithProbability.probability {
			return randomNumber(sizeRangeWithProbability.min, sizeRangeWithProbability.max)
		}
	}

	// It's impossible to reach here because the sum of all probabilities is 100%. It's already validated before calling this function.
	return 0
}

type sizeRangeWithProbability struct {
	min         int64
	max         int64
	probability float64
}

func parseSizeRangeWithProbability(inputs []string) ([]*sizeRangeWithProbability, error) {
	if len(inputs) == 0 {
		return nil, fmt.Errorf("empty size range with probability")
	}

	var sizeRangeWithProbabilities []*sizeRangeWithProbability
	for _, input := range inputs {
		parts := strings.Split(input, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid size range with probability: %s", input)
		}

		probability, err := utils.ParsePercentage(strings.TrimSpace(parts[0]))
		if err != nil {
			return nil, fmt.Errorf("invalid size range with probability: %s", input)
		}

		min, max, err := parseSizeRange(strings.TrimSpace(parts[1]))
		if err != nil {
			return nil, fmt.Errorf("invalid size range with probability: %s", input)
		}

		sizeRangeWithProbabilities = append(sizeRangeWithProbabilities, &sizeRangeWithProbability{
			min:         min,
			max:         max,
			probability: probability,
		})
	}

	return sizeRangeWithProbabilities, nil
}

func validateSizeRangeWithProbability(sizeRangeWithProbabilities []*sizeRangeWithProbability) error {
	var sum float64
	for i, sr := range sizeRangeWithProbabilities {
		// The min should be less than the max.
		if sr.min >= sr.max {
			return fmt.Errorf("invalid size range with probability: %v", sizeRangeWithProbabilities[i])
		}

		// The probability should be a positive number in (0, 1).
		if sr.probability < 0 || sr.probability > 1 {
			return fmt.Errorf("invalid size range with probability: %v", sizeRangeWithProbabilities[i])
		}

		// Check if the range overlaps with any other range.
		for j, other := range sizeRangeWithProbabilities {
			if i == j {
				continue
			}

			if sr.max > other.min && sr.min < other.max {
				return fmt.Errorf("the range [%d, %d) overlaps with the range [%d, %d)", sr.min, sr.max, other.min, other.max)
			}
		}

		sum += sr.probability
	}

	// The sum of all probabilities should be 1.
	if math.Abs(sum-1) > 0.000001 {
		return fmt.Errorf("invalid size range with probability: %v", sizeRangeWithProbabilities)
	}

	return nil
}

// Generate a random number in [min, max).
func randomNumber(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}
