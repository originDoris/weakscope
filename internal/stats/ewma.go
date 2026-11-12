package stats

type Ewma struct {
	Alpha       float64
	srtt        float64
	initialized bool
}

func NewEwma(alpha float64) *Ewma {
	e := Ewma{}
	e.Alpha = alpha
	return &e
}

func (e *Ewma) Update(rtt float64) float64 {
	if rtt <= 0 {
		return e.srtt
	}
	if e.initialized {
		e.srtt = (1-e.Alpha)*e.srtt + e.Alpha*rtt
	} else {
		e.srtt = rtt
	}
	e.initialized = true
	return e.srtt
}
