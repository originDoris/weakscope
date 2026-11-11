package probe

import (
	"context"
	"errors"
	"net"
	"strings"
	"testing"
	"time"
)

type fakeDialer struct {
	resolveDelay time.Duration
	resolveErr   error
	connectDelay time.Duration
	connectErr   error
}

type fakeConn struct{ net.Conn }

func (fakeConn) Close() error {
	return nil
}

func (f fakeDialer) ResolveTCPAddr(network, address string) (*net.TCPAddr, error) {
	if f.resolveDelay > 0 {
		time.Sleep(f.resolveDelay)
	}
	if f.resolveErr != nil {
		return nil, f.resolveErr
	}
	return &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 80}, nil
}

func (f fakeDialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	if f.connectDelay > 0 {
		time.Sleep(f.connectDelay)
	}
	if f.connectErr != nil {
		return nil, f.connectErr
	}
	return fakeConn{}, nil
}

func TestProbe(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name      string
		dialer    fakeDialer
		target    string
		wantLen   int
		wantPhase []Phase
		wantOk    []bool
		wantErr   []string
	}{
		{name: "dns_slow", dialer: fakeDialer{resolveDelay: 50 * time.Millisecond}, target: "127.0.0.1:80", wantLen: 2, wantPhase: []Phase{PhaseDns, PhaseConnect}, wantOk: []bool{true, true}, wantErr: []string{"", ""}},
		{name: "dns_err", dialer: fakeDialer{resolveErr: errors.New("dns无法访问")}, target: "127.0.0.1:80", wantLen: 1, wantPhase: []Phase{PhaseDns}, wantOk: []bool{false}, wantErr: []string{"dns无法访问"}},
		{name: "connect_slow", dialer: fakeDialer{connectDelay: 100 * time.Millisecond}, target: "127.0.0.1:80", wantLen: 2, wantPhase: []Phase{PhaseDns, PhaseConnect}, wantOk: []bool{true, true}, wantErr: []string{"", ""}},
		{name: "connect_error", dialer: fakeDialer{connectErr: errors.New("地址无法访问")}, target: "127.0.0.1:80", wantLen: 2, wantPhase: []Phase{PhaseDns, PhaseConnect}, wantOk: []bool{true, false}, wantErr: []string{"", "地址无法访问"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Probe(ctx, test.dialer, test.target)
			if len(got) != test.wantLen {
				t.Fatalf("len(got) = %d, want %d", len(got), test.wantLen)
			}

			for i, r := range got {
				if test.wantErr[i] != "" && !strings.Contains(r.Err, test.wantErr[i]) {
					t.Errorf("got '%s' want %s", r.Err, test.wantErr[i])
				}
				if r.Phase != test.wantPhase[i] {
					t.Errorf("got %v want %s", r.Phase, test.wantPhase[i])
				}
				if r.Ok != test.wantOk[i] {
					t.Errorf("got %v want %v", r.Ok, test.wantOk[i])
				}

			}
		})
	}
}
