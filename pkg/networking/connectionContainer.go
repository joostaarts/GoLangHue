package networking

import (
	"net"
	"sync"
)

// ConnectionContainer is a simple list of connections which is thread safe to add to
type ConnectionContainer struct {
	mu          sync.Mutex
	connections []net.UDPConn
}

// AddConnection adds a connection to the list, may be called concurrently
func (container *ConnectionContainer) AddConnection(con net.UDPConn) {
	container.mu.Lock()
	container.connections = append(container.connections, con)
	container.mu.Unlock()
}

// Dispose closes all connections
func (container *ConnectionContainer) Dispose() {
	container.mu.Lock()

	for _, con := range container.connections {
		con.Close()
	}

	container.connections = nil

	container.mu.Unlock()
}
