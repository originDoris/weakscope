package stats

type Ewma struct {
	Alpha float64
	srtt  float64
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
	if e.srtt == 0 {
		e.srtt = rtt
	}
	e.srtt = (1-e.Alpha)*e.srtt + e.Alpha*rtt
	return e.srtt
}
