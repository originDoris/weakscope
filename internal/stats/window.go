package stats

type Window struct {
	rtts []float64
	size int
}

type WindowStats struct {
	Count int
	Min   float64
	Max   float64
	Mean  float64
	P50   float64
	P95   float64
	P99   float64
}

func NewWindow(size int) *Window {
	if size <= 0 {
		size = 1
	}
	window := Window{}
	window.rtts = make([]float64, 0, size)
	window.size = size
	return &window
}

func (w *Window) Add(x float64) {
	if len(w.rtts) == w.size {
		copy(w.rtts, w.rtts[1:])
		w.rtts[w.size-1] = x
	} else {
		w.rtts = append(w.rtts, x)
	}
}

func (w *Window) Stats() WindowStats {
	stats := WindowStats{}
	stats.Count = len(w.rtts)
	stats.Min = Min(w.rtts)
	stats.Max = Max(w.rtts)
	stats.Mean = Mean(w.rtts)
	stats.P50 = Percentile(w.rtts, 0.5)
	stats.P95 = Percentile(w.rtts, 0.95)
	stats.P99 = Percentile(w.rtts, 0.99)
	return stats
}
