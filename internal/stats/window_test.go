package stats

import "testing"

type content struct {
	adds []float64
	want WindowStats
}

func TestWindow(t *testing.T) {
	tests := []struct {
		name    string
		window  Window
		content []content
	}{
		{name: "test1", window: *NewWindow(3), content: []content{
			{adds: []float64{1, 2, 3}, want: WindowStats{Count: 3, Min: 1, Max: 3, Mean: 2, P50: 2, P95: 3, P99: 3}},
			{adds: []float64{10}, want: WindowStats{Count: 3, Mean: 5, Min: 2, Max: 10, P50: 3, P95: 10, P99: 10}},
			{adds: []float64{}, want: WindowStats{Count: 3, Mean: 5, Min: 2, Max: 10, P50: 3, P95: 10, P99: 10}},
		}},
		{name: "empty", window: *NewWindow(2), content: []content{
			{want: WindowStats{Count: 0, Min: 0, Max: 0, Mean: 0, P50: 0, P95: 0, P99: 0}},
		}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			contents := test.content
			window := test.window
			for _, c := range contents {
				for _, add := range c.adds {
					window.Add(add)
				}
				got := window.Stats()
				if got != c.want {
					t.Fatalf("got %v want %v", got, c.want)
				}
			}
		})
	}
}
