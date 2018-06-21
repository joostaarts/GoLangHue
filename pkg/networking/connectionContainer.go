package networking

import (
	"net"
	"sync"
)

// ConnectionContainer is a simple list of connections which is thread safe to add to
type ConnectionContainer struct {
	mu          sync.Mutex
	connections []*net.UDPConn
}

// AddConnection adds a connection to the list, may be called concurrently
func (container *ConnectionContainer) AddConnection(con net.UDPConn) {
	container.mu.Lock()
	container.connections = append(container.connections, &con)
	container.mu.Unlock()
}

// CloseConnection closes and removes a connection from the list, may be called concurrently
func (container *ConnectionContainer) CloseConnection(con *net.UDPConn) {
	container.mu.Lock()
	con.Close()

	for i, item := range container.connections {
		if *item == *con {
			if len(container.connections) == 1 {
				container.connections[i] = nil
			} else {
				container.connections[i] = container.connections[len(container.connections)-1]
				container.connections[len(container.connections)-1] = nil
			}

			container.connections = container.connections[:len(container.connections)-1]
			break
		}
	}

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
