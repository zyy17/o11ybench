package faker

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

// FakeNumberOptions is the options for generating the fake number.
type FakeNumberOptions struct {
	// Type is the type of the number.
	Type NumberType `yaml:"type,omitempty" json:"type,omitempty"`

	// Min is the minimum value of the number.
	Min string `yaml:"min,omitempty" json:"min,omitempty"`

	// Max is the maximum value of the number.
	Max string `yaml:"max,omitempty" json:"max,omitempty"`

	// Precision is only used for float types.
	Precision *int `yaml:"precision,omitempty" json:"precision,omitempty"`

	// Prefix is the output prefix of the number.
	Prefix string `yaml:"prefix,omitempty" json:"prefix,omitempty"`

	// Suffix is the output suffix of the number.
	Suffix string `yaml:"suffix,omitempty" json:"suffix,omitempty"`
}

func (o *FakeNumberOptions) Validate() error {
	switch o.Type {
	case NumberTypeInt8:
		if err := o.validateIntRange(o.Min, o.Max, 8, int64(math.MinInt8), int64(math.MaxInt8)); err != nil {
			return err
		}
	case NumberTypeInt16:
		if err := o.validateIntRange(o.Min, o.Max, 16, int64(math.MinInt16), int64(math.MaxInt16)); err != nil {
			return err
		}
	case NumberTypeInt32:
		if err := o.validateIntRange(o.Min, o.Max, 32, int64(math.MinInt32), int64(math.MaxInt32)); err != nil {
			return err
		}
	case NumberTypeInt64:
		if err := o.validateIntRange(o.Min, o.Max, 64, int64(math.MinInt64), int64(math.MaxInt64)); err != nil {
			return err
		}
	case NumberTypeUint8:
		if err := o.validateUintRange(o.Min, o.Max, 8, uint64(math.MaxUint8)); err != nil {
			return err
		}
	case NumberTypeUint16:
		if err := o.validateUintRange(o.Min, o.Max, 16, uint64(math.MaxUint16)); err != nil {
			return err
		}
	case NumberTypeUint32:
		if err := o.validateUintRange(o.Min, o.Max, 32, uint64(math.MaxUint32)); err != nil {
			return err
		}
	case NumberTypeUint64:
		if err := o.validateUintRange(o.Min, o.Max, 64, uint64(math.MaxUint64)); err != nil {
			return err
		}
	case NumberTypeFloat32:
		if err := o.validateFloatRange(o.Min, o.Max, 32, math.SmallestNonzeroFloat32, math.MaxFloat32); err != nil {
			return err
		}
	case NumberTypeFloat64:
		if err := o.validateFloatRange(o.Min, o.Max, 64, math.SmallestNonzeroFloat64, math.MaxFloat64); err != nil {
			return err
		}
	default:
		return fmt.Errorf("invalid number type: %s", o.Type)
	}

	if o.Precision != nil {
		if o.Type != NumberTypeFloat32 && o.Type != NumberTypeFloat64 {
			return fmt.Errorf("precision is only used for float types")
		}

		if *o.Precision < 0 {
			return fmt.Errorf("precision must be greater than 0")
		}
	}

	return nil
}

func (o *FakeNumberOptions) validateIntRange(min, max string, bitSize int, allowedMin, allowedMax int64) error {
	var (
		parsedMin int64
		parsedMax int64

		err error
	)

	if min != "" {
		parsedMin, err = strconv.ParseInt(min, 10, bitSize)
		if err != nil {
			return err
		}
		if parsedMin < allowedMin {
			return fmt.Errorf("min value is less than the minimum allowed value for '%s' '%d'", o.Type, allowedMin)
		}
	}

	if max != "" {
		parsedMax, err = strconv.ParseInt(max, 10, bitSize)
		if err != nil {
			return err
		}
		if parsedMax > allowedMax {
			return fmt.Errorf("max value is greater than the maximum allowed value for '%s' '%d'", o.Type, allowedMax)
		}
	}

	if parsedMin > parsedMax {
		return fmt.Errorf("min value '%d' is greater than the max value '%d'", parsedMin, parsedMax)
	}

	return nil
}

func (o *FakeNumberOptions) validateUintRange(min, max string, bitSize int, allowedMax uint64) error {
	var (
		parsedMin uint64
		parsedMax uint64

		err error
	)

	if min != "" {
		parsedMin, err = strconv.ParseUint(min, 10, bitSize)
		if err != nil {
			return err
		}
	}

	if max != "" {
		parsedMax, err = strconv.ParseUint(max, 10, bitSize)
		if err != nil {
			return err
		}
		if parsedMax > allowedMax {
			return fmt.Errorf("max value is greater than the maximum allowed value for '%s' '%d'", o.Type, allowedMax)
		}
	}

	if parsedMin > parsedMax {
		return fmt.Errorf("min value '%d' is greater than the max value '%d'", parsedMin, parsedMax)
	}

	return nil
}

func (o *FakeNumberOptions) validateFloatRange(min, max string, bitSize int, allowedMin, allowedMax float64) error {
	var (
		parsedMin float64
		parsedMax float64

		err error
	)

	if min != "" {
		parsedMin, err = strconv.ParseFloat(min, bitSize)
		if err != nil {
			return err
		}
		if parsedMin < allowedMin {
			return fmt.Errorf("min value is less than the minimum allowed value for '%s' '%f'", o.Type, allowedMin)
		}
	}

	if max != "" {
		parsedMax, err = strconv.ParseFloat(max, bitSize)
		if err != nil {
			return err
		}
		if parsedMax > allowedMax {
			return fmt.Errorf("max value is greater than the maximum allowed value for '%s' '%f'", o.Type, allowedMax)
		}
	}

	if parsedMin > parsedMax {
		return fmt.Errorf("min value '%f' is greater than the max value '%f'", parsedMin, parsedMax)
	}

	return nil
}

// NumberType is the type of the number.
type NumberType string

const (
	// NumberTypeInt8 represents an 8-bit signed integer type.
	NumberTypeInt8 = "int8"

	// NumberTypeInt16 represents a 16-bit signed integer type.
	NumberTypeInt16 = "int16"

	// NumberTypeInt32 represents a 32-bit signed integer type.
	NumberTypeInt32 = "int32"

	// NumberTypeInt64 represents a 64-bit signed integer type.
	NumberTypeInt64 = "int64"

	// NumberTypeUint8 represents an 8-bit unsigned integer type.
	NumberTypeUint8 = "uint8"

	// NumberTypeUint16 represents a 16-bit unsigned integer type.
	NumberTypeUint16 = "uint16"

	// NumberTypeUint32 represents a 32-bit unsigned integer type.
	NumberTypeUint32 = "uint32"

	// NumberTypeUint64 represents a 64-bit unsigned integer type.
	NumberTypeUint64 = "uint64"

	// NumberTypeFloat32 represents a 32-bit floating point type.
	NumberTypeFloat32 = "float32"

	// NumberTypeFloat64 represents a 64-bit floating point type.
	NumberTypeFloat64 = "float64"
)

// FakeIPv4 generates a random IPv4 address.
func FakeNumber(opts Options) (string, error) {
	var options FakeNumberOptions
	if err := parseOptions(opts, &options); err != nil {
		return "", err
	}

	var (
		number string
		err    error
	)

	switch options.Type {
	case NumberTypeInt8:
		number, err = generateInt(options.Min, options.Max, 8, int64(math.MinInt8), int64(math.MaxInt8))
		if err != nil {
			return "", err
		}
	case NumberTypeInt16:
		number, err = generateInt(options.Min, options.Max, 16, int64(math.MinInt16), int64(math.MaxInt16))
		if err != nil {
			return "", err
		}
	case NumberTypeInt32:
		number, err = generateInt(options.Min, options.Max, 32, int64(math.MinInt32), int64(math.MaxInt32))
		if err != nil {
			return "", err
		}
	case NumberTypeInt64:
		number, err = generateInt(options.Min, options.Max, 64, int64(math.MinInt64), int64(math.MaxInt64))
		if err != nil {
			return "", err
		}
	case NumberTypeUint8:
		number, err = generateUint(options.Min, options.Max, 8, uint64(math.MaxUint8))
		if err != nil {
			return "", err
		}
	case NumberTypeUint16:
		number, err = generateUint(options.Min, options.Max, 16, uint64(math.MaxUint16))
		if err != nil {
			return "", err
		}
	case NumberTypeUint32:
		number, err = generateUint(options.Min, options.Max, 32, uint64(math.MaxUint32))
		if err != nil {
			return "", err
		}
	case NumberTypeUint64:
		number, err = generateUint(options.Min, options.Max, 64, uint64(math.MaxUint64))
		if err != nil {
			return "", err
		}
	case NumberTypeFloat32:
		number, err = generateFloat(options.Min, options.Max, 32, math.SmallestNonzeroFloat32, math.MaxFloat32, options.Precision)
		if err != nil {
			return "", err
		}
	case NumberTypeFloat64:
		number, err = generateFloat(options.Min, options.Max, 64, math.SmallestNonzeroFloat64, math.MaxFloat64, options.Precision)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("invalid number type: %s", options.Type)
	}

	if options.Prefix != "" {
		number = options.Prefix + number
	}

	if options.Suffix != "" {
		number = number + options.Suffix
	}

	return number, nil
}

func generateInt(min, max string, bitSize int, allowedMin, allowedMax int64) (string, error) {
	var (
		minValue = int64(allowedMin)
		maxValue = int64(allowedMax)
	)

	if min != "" {
		parsedMin, err := strconv.ParseInt(min, 10, bitSize)
		if err != nil {
			return "", err
		}
		minValue = parsedMin
	}

	if max != "" {
		parsedMax, err := strconv.ParseInt(max, 10, bitSize)
		if err != nil {
			return "", err
		}
		maxValue = parsedMax
	}

	return fmt.Sprintf("%d", rand.Int63n(maxValue-minValue+1)+minValue), nil
}

func generateUint(min, max string, bitSize int, allowedMax uint64) (string, error) {
	minValue := uint64(0)
	maxValue := uint64(allowedMax)

	if min != "" {
		parsedMin, err := strconv.ParseUint(min, 10, bitSize)
		if err != nil {
			return "", err
		}
		minValue = parsedMin
	}

	if max != "" {
		parsedMax, err := strconv.ParseUint(max, 10, bitSize)
		if err != nil {
			return "", err
		}
		maxValue = parsedMax
	}

	return fmt.Sprintf("%d", rand.Uint64()%(maxValue-minValue+1)+minValue), nil
}

func generateFloat(min, max string, bitSize int, allowedMin, allowedMax float64, precision *int) (string, error) {
	minValue := float64(allowedMin)
	maxValue := float64(allowedMax)

	if min != "" {
		parsedMin, err := strconv.ParseFloat(min, bitSize)
		if err != nil {
			return "", err
		}
		minValue = parsedMin
	}

	if max != "" {
		parsedMax, err := strconv.ParseFloat(max, bitSize)
		if err != nil {
			return "", err
		}
		maxValue = parsedMax
	}

	value := rand.Float64()*(maxValue-minValue) + minValue
	if precision != nil {
		return strconv.FormatFloat(value, 'f', *precision, bitSize), nil
	}

	return fmt.Sprintf("%f", value), nil
}
