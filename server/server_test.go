package main

import (
	"net"
	"testing"
)

// Toma um ponteiro para o struct T do testing
func TestBindPort(reporter *testing.T) {
	testListener, err := net.Listen("tcp", "127.0.0.1:9001")
	if err != nil {
		reporter.Fatalf("Failed to create a testListener: %v", err)
	}
	defer testListener.Close()

	port := 9001
	serverSocket, boundPort := bindPort(port)
	// defer serverSocket.Close()

	if serverSocket == nil {
		reporter.Error("Expected a non-nil server socket")
	}

	if boundPort != port {
		reporter.Errorf("Expected bound port to be %d, but got %d", port, boundPort)
	}
}
