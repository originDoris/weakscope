package stats

import (
	"cmp"
	"math"
	"slices"
)

func Mean(xs []float64) float64 {
	size := len(xs)
	if size == 0 {
		return 0
	}
	var total float64
	for _, x := range xs {
		total += x
	}
	return total / float64(size)
}

func Min(xs []float64) (min float64) {
	if len(xs) == 0 {
		return 0
	}
	min = xs[0]
	for _, x := range xs {
		if x < min {
			min = x
		}
	}
	return min
}

func Max(xs []float64) (max float64) {
	if len(xs) == 0 {
		return 0
	}
	max = xs[0]
	for _, x := range xs {
		if x > max {
			max = x
		}
	}
	return max
}

func Percentile(xs []float64, p float64) (result float64) {
	clone := slices.Clone(xs)
	slices.SortFunc(clone, func(e, e2 float64) int {
		return cmp.Compare(e, e2)
	})
	size := len(clone)
	if size == 0 {
		return 0
	}
	if p <= 0 {
		return clone[0]
	}
	if p >= 1 {
		return clone[size-1]
	}

	f := float64(size) * p
	if f < 0 {
		return 0
	}
	i := int(math.Ceil(f))
	return clone[i-1]
}
