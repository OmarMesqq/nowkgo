package main

import (
	"fmt"
	"net"
	"testing"
)

func TestBindPort(reporter *testing.T) {
	busyPort := 9001
	expectedPort := 9002
	testListener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", busyPort))
	if err != nil {
		reporter.Fatalf("Failed to create a testListener: %v", err)
	}
	defer testListener.Close()

	serverSocket, boundPort := bindPort(busyPort)
	// defer serverSocket.Close()

	if serverSocket == nil {
		reporter.Error("Expected a non-nil server socket")
	}

	if boundPort != expectedPort {
		reporter.Errorf("Expected bound port to be %d, but got %d", busyPort, boundPort)
	}
}
