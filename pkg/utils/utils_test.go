package utils

import (
	"math"
	"testing"
)

func TestParseSize(t *testing.T) {
	tests := []struct {
		name string
		size string
		want int64
	}{
		{
			size: "100bytes",
			want: 100,
		},
		{
			size: "1kib",
			want: 1024,
		},
		{
			size: "5kb",
			want: 5 * 1000,
		},
		{
			size: "2mib",
			want: 2 * 1024 * 1024,
		},
		{
			size: "9mb",
			want: 9 * 1000 * 1000,
		},
		{
			size: "3gib",
			want: 3 * 1024 * 1024 * 1024,
		},
		{
			size: "3gb",
			want: 3 * 1000 * 1000 * 1000,
		},
		{
			size: "4TiB",
			want: 4 * 1024 * 1024 * 1024 * 1024,
		},
		{
			size: "4TB",
			want: 4 * 1000 * 1000 * 1000 * 1000,
		},
		{
			size: "5PiB",
			want: 5 * 1024 * 1024 * 1024 * 1024 * 1024,
		},
		{
			size: "5PB",
			want: 5 * 1000 * 1000 * 1000 * 1000 * 1000,
		},
	}

	for i, test := range tests {
		got, err := ParseSize(test.size)
		if err != nil {
			t.Errorf("Run test [%d]: %v", i, err)
		}

		if got != test.want {
			t.Errorf("Run test [%d]: got '%d', want '%d'", i, got, test.want)
		}
	}
}

func TestParsePercentage(t *testing.T) {
	tests := []struct {
		percentage string
		want       float64
	}{
		{
			percentage: "100%",
			want:       1,
		},
		{
			percentage: "5.5%",
			want:       0.055,
		},
		{
			percentage: "7.895%",
			want:       0.07895,
		},
	}

	for i, test := range tests {
		got, err := ParsePercentage(test.percentage)
		if err != nil {
			t.Errorf("Run test [%d]: %v", i, err)
		}

		if math.Abs(got-test.want) > 0.000001 {
			t.Errorf("Run test [%d]: got '%f', want '%f'", i, got, test.want)
		}
	}
}

func TestRandomNumber(t *testing.T) {
	tests := []struct {
		min int64
		max int64
	}{
		{
			min: 100,
			max: 200,
		},
		{
			min: 477,
			max: 789,
		},
		{
			min: 1000000,
			max: 2000000,
		},
	}

	for i, test := range tests {
		got := RandomNumber(test.min, test.max)
		if got < test.min || got > test.max {
			t.Errorf("Run test [%d]: the random number is out of range [%d, %d), got '%d'", i, test.min, test.max, got)
		}
	}
}
