package probe

import (
	"context"
	"net"
	"time"
)

type Dialer interface {
	DialContext(ctx context.Context, network, address string) (net.Conn, error)
	ResolveTCPAddr(network, address string) (*net.TCPAddr, error)
}

type NetDialer struct {
}

func (d NetDialer) ResolveTCPAddr(network, address string) (*net.TCPAddr, error) {
	return net.ResolveTCPAddr(network, address)
}

func (d NetDialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	dialer := net.Dialer{
		Timeout: time.Duration(100 * time.Millisecond),
	}
	return dialer.DialContext(ctx, network, address)
}
