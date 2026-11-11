package probe

import (
	"context"
	"net"
)

type Dialer interface {
	DialContext(ctx context.Context, network, address string) (net.Conn, error)
	ResolveTCPAddr(network, address string) (*net.TCPAddr, error)
}
