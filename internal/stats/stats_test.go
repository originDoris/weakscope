package stats

import (
	"math"
	"testing"
)

func TestMean(t *testing.T) {
	tests := []struct {
		name string
		data []float64
		want float64
	}{
		{name: "nil", data: make([]float64, 0), want: 0},
		{name: "normal", data: []float64{1, 2, 3, 4}, want: 2.5},
		{name: "single", data: []float64{42}, want: 42},
		{name: "float_case", data: []float64{0.1, 0.2, 0.3}, want: 0.2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mean := Mean(test.data)
			if math.IsNaN(mean) || math.Abs(mean-test.want) > 1e-9 {
				t.Errorf("got %v want %v", mean, test.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name string
		data []float64
		want float64
	}{
		{name: "nil", data: make([]float64, 0), want: 0},
		{name: "normal", data: []float64{1, 2, 3, 4}, want: 1},
		{name: "single", data: []float64{42}, want: 42},
		{name: "float_case", data: []float64{0.1, 0.2, 0.3}, want: 0.1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			min := Min(test.data)
			if math.IsNaN(min) || math.Abs(min-test.want) > 1e-9 {
				t.Errorf("got %v want %v", min, test.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name string
		data []float64
		want float64
	}{
		{name: "nil", data: make([]float64, 0), want: 0},
		{name: "normal", data: []float64{1, 2, 3, 4}, want: 4},
		{name: "single", data: []float64{42}, want: 42},
		{name: "float_case", data: []float64{0.1, 0.2, 0.3}, want: 0.3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			max := Max(test.data)
			if math.IsNaN(max) || math.Abs(max-test.want) > 1e-9 {
				t.Errorf("got %v want %v", max, test.want)
			}
		})
	}
}

func TestPercentile(t *testing.T) {

	s := make([]float64, 50)
	for i := range 50 {
		s[i] = float64(i + 1)
	}
	tests := []struct {
		name string
		data []float64
		p    float64
		want float64
	}{
		{name: "nil", data: make([]float64, 0), p: 0.95, want: 0},
		{name: "normal", data: []float64{1, 2, 3, 4}, p: 0.95, want: 4},
		{name: "unsorted", data: []float64{5, 6, 1, 2, 3, 4}, p: 0.95, want: 6},
		{name: "single", data: []float64{42}, p: 0.95, want: 42},
		{name: "p_zero", data: []float64{42}, p: 0, want: 42},
		{name: "p_over", data: []float64{42}, p: 1.5, want: 42},
		{name: "off-by-one", data: s, p: 0.95, want: 48},
		{name: "float_case", data: []float64{0.1, 0.2, 0.3}, p: 0.99, want: 0.3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			max := Percentile(test.data, test.p)
			if math.IsNaN(max) || math.Abs(max-test.want) > 1e-9 {
				t.Errorf("got %v want %v", max, test.want)
			}
		})
	}
}
