package faker

import (
	"fmt"
	"strings"

	"github.com/zyy17/o11ybench/pkg/generator/common"
	"github.com/zyy17/o11ybench/pkg/generator/faker/dataset"
	"github.com/zyy17/o11ybench/pkg/generator/faker/distribution"
	"github.com/zyy17/o11ybench/pkg/utils"
)

// FakeLogsOptions is the options for generating fake logs.
type FakeLogsOptions struct {
	// Dataset is the dataset of the fake logs.
	Dataset string `yaml:"dataset"`
	// Size is the approximate total size of all generated fake words. It can be in format "100bytes", "1kib", "1mib", etc.
	// The size should be a positive integer.
	// It's exclusive with Count, SizeRange and SizeRangeByPossibility options.
	Size string `yaml:"size,omitempty"`

	// SizeRangeWithPossibility specifies multiple approximate total size range of all generated fake words with their probabilities.
	// Each entry should be in format "<probability>%:<min-size>-<max-size>" (e.g. "10%:30bytes-40bytes").
	// The size should be a positive integer and each range should not overlap.
	// The sum of all probabilities must equal 100%.
	// It's exclusive with Count, Size and SizeRange options.
	SizeRangeWithPossibility []string `yaml:"sizeRangeWithPossibility,omitempty"`
}

// FakeLogs generates a fake log.
func FakeLogs(_ common.ElementType, opts Options) (string, error) {
	var options FakeLogsOptions
	if err := parseOptions(opts, &options); err != nil {
		return "", err
	}

	if options.Dataset == "" {
		return "", fmt.Errorf("dataset is required")
	}
	var ds *dataset.Dataset
	if _, ok := dataset.Datasets[dataset.DatasetType(options.Dataset)]; !ok {
		return "", fmt.Errorf("dataset %s not found", options.Dataset)
	}
	ds = dataset.Datasets[dataset.DatasetType(options.Dataset)]

	if options.Size != "" {
		size, err := utils.ParseSize(options.Size)
		if err != nil {
			return "", err
		}

		return generateLogs(ds, size), nil
	}

	if options.SizeRangeWithPossibility != nil {
		d, err := distribution.NewDistribution(options.SizeRangeWithPossibility)
		if err != nil {
			return "", err
		}

		return generateLogs(ds, d.RandomNumber()), nil
	}

	return "", nil
}

func generateLogs(dataset *dataset.Dataset, size int64) string {
	logs := make([]string, 0, size/int64(dataset.AverageSize))
	for i := 0; i < int(size/int64(dataset.AverageSize)); i++ {
		randomIndex := utils.RandomNumber(0, int64(len(dataset.Logs)))
		logs = append(logs, dataset.Logs[randomIndex])
	}

	return strings.Join(logs, "\n")
}
