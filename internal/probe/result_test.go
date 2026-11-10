package probe

import (
	"encoding/json"
	"testing"
	"time"
)

func TestResult(t *testing.T) {
	var date = time.Date(2026, 11, 10, 12, 0, 0, 0, time.UTC)
	tests := []struct {
		name string
		in   Result
		want string
	}{
		{name: "success", in: Result{Ts: date, Target: "127.0.0.1:80", Phase: PhaseDns, RttMs: 0.1, Ok: true, Err: ""}, want: `{"ts":"2026-11-10T12:00:00Z","target":"127.0.0.1:80","phase":"dns","rtt_ms":0.1,"ok":true}`},
		{name: "fail", in: Result{Ts: date, Target: "127.0.0.1:80", Phase: PhaseDns, RttMs: 0.1, Ok: false, Err: "无法访问"}, want: `{"ts":"2026-11-10T12:00:00Z","target":"127.0.0.1:80","phase":"dns","rtt_ms":0.1,"ok":false,"err":"无法访问"}`},
		{name: "rtt", in: Result{Ts: date, Target: "127.0.0.1:80", Phase: PhaseDns, RttMs: 0.2, Ok: true, Err: ""}, want: `{"ts":"2026-11-10T12:00:00Z","target":"127.0.0.1:80","phase":"dns","rtt_ms":0.2,"ok":true}`},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			marshal, err := json.Marshal(test.in)
			if err != nil {
				t.Fatalf("发生错误 %s", err)
			}
			if got := string(marshal); got != test.want {
				t.Errorf("got '%s' want '%s'", got, test.want)
			}
		})
	}
}
