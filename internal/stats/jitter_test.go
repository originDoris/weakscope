package stats

import "testing"

type packet struct {
	s    float64
	r    float64
	want float64
}

func TestJitter_Update(t *testing.T) {
	tests := []struct {
		name    string
		jitter  *Jitter
		dresult []packet
	}{
		{name: "rfc3550", jitter: NewJitter(0.0625), dresult: []packet{
			{10, 10, 0},
			{20, 20, 0},
			{35, 40, 0.3125},
			{45, 60, 0.91796875},
		}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, p := range test.dresult {
				got := test.jitter.Update(p.r, p.s)
				if got != p.want {
					t.Fatalf("got %v want %v", got, p.want)
				}
			}
		})
	}
}
