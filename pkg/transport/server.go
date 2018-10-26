package transport

import "net"

// Server TODO issue#173
type Server interface {
	// Serve TODO issue#173
	Serve(net.Listener) error
}
