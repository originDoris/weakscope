package stats

import "testing"

func TestEwma(t *testing.T) {
	tests := []struct {
		name string
		ewma *Ewma
		rtts []float64
		want []float64
	}{
		{name: "rfc6298", ewma: NewEwma(0.125), rtts: []float64{1, 1.5, 0.5}, want: []float64{1, 1.0625, 0.9921875}},
		{name: "alpha_half", ewma: NewEwma(0.5), rtts: []float64{1, 1.5, 0.5}, want: []float64{1, 1.25, 0.875}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for i, rtt := range test.rtts {
				got := test.ewma.Update(rtt)
				if got != test.want[i] {
					t.Fatalf("got %v want %v ", got, test.want[i])
				}
			}
		})
	}
}
