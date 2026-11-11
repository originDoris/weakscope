package probe

import (
	"context"
	"time"
)

func Probe(ctx context.Context, d Dialer, target string) []Result {
	dnsStart := time.Now()
	addr, err := d.ResolveTCPAddr("tcp", target)
	dnsResult := Result{Ts: time.Now(), Target: target, Phase: PhaseDns}
	if err != nil {
		dnsResult.Err = err.Error()
		dnsResult.Ok = false
		return []Result{dnsResult}
	}

	dnsResult.RttMs = float64(time.Since(dnsStart).Nanoseconds()) / 1e6
	dnsResult.Ok = true
	connectStart := time.Now()
	conn, err := d.DialContext(ctx, "tcp", addr.String())
	connectResult := Result{Ts: time.Now(), Target: target, Phase: PhaseConnect}
	if err != nil {
		connectResult.Err = err.Error()
		connectResult.Ok = false
		return []Result{dnsResult, connectResult}
	}
	connectResult.RttMs = float64(time.Since(connectStart).Nanoseconds()) / 1e6
	connectResult.Ok = true
	defer conn.Close()
	return []Result{dnsResult, connectResult}
}
