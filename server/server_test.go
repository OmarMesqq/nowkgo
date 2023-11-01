package main

import (
	"fmt"
	"net"
	"strings"
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

func TestClientQueue(reporter *testing.T) {
	listener, _ := createServer(9005)
	defer listener.Close()

	done := make(chan bool)        // Channel to signal when goroutine is done
	errors := make(chan error, 10) // Channel to communicate errors

	go func() {
		clientCount := 0
		for {
			_, err := listener.Accept()
			if err != nil {
				// If listener is closed, just close the done channel without pushing an error
				if strings.Contains(err.Error(), "use of closed network connection") {
					close(done)
					return
				}
				errors <- err
				close(done)
				return
			}
			clientCount++
			if clientCount > 10 {
				errors <- fmt.Errorf("Too many clients connected.")
				close(done)
				return
			}
		}
	}()

	for i := 0; i < 11; i++ {
		_, err := net.Dial("tcp", "127.0.0.1:9005")
		if err != nil {
			reporter.Fatalf("Failed to create a client instance: %v", err)
		}
	}

	// Signal to the goroutine that clients are done connecting
	listener.Close()
	// Wait for the goroutine to finish processing
	<-done

	// Check for errors sent from the goroutine
	if len(errors) > 0 {
		reporter.Fatalf("Error in goroutine: %v", <-errors)
	}
}
