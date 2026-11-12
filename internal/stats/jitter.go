package stats

import "math"

type Jitter struct {
	s       float64
	r       float64
	j       float64
	Alpha   float64
	hasPrev bool
}

func NewJitter(alpha float64) *Jitter {
	e := Jitter{}
	e.Alpha = alpha
	return &e
}

func (jitter *Jitter) Update(r float64, s float64) float64 {
	d := 0.0
	result := 0.0
	if jitter.hasPrev {
		d = (r - jitter.r) - (s - jitter.s)
		result = jitter.j + (math.Abs(d)-jitter.j)*jitter.Alpha
	}
	jitter.r = r
	jitter.s = s
	jitter.j = result
	jitter.hasPrev = true
	return result
}
