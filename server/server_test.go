package main

import (
	"fmt"
	"net"
	"testing"
)

func TestCreateServer(reporter *testing.T) {
	busyPort := 9001
	expectedPort := 9002
	firstServerInstance, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", busyPort))
	if err != nil {
		reporter.Fatalf("Failed to create a firstServerInstance: %v", err)
	}
	defer firstServerInstance.Close()

	secondServerInstance, boundPort := createServer(busyPort)
	// defer secondServerInstance.Close()

	if secondServerInstance == nil {
		reporter.Error("Expected a non-nil server socket")
	}

	if boundPort != expectedPort {
		reporter.Errorf("Expected bound port to be %d, but got %d", busyPort, boundPort)
	}
}
