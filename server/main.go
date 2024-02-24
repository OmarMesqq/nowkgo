package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func bytes(strs ...string) []byte {
	var result string
	for _, str := range strs {
		result += str
	}
	return []byte(result)
}

func createServer(port int) (net.Listener, int) {
	for {
		serverSocket, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
		if err == nil {
			return serverSocket, port
		}
		port++
	}
}

func handleClient(clientSocket net.Conn, clientNumber int, theater *Theater) {
	defer clientSocket.Close()
	clientSocket.SetDeadline(time.Now().Add(5 * time.Minute))
	theater.Enter()

	intro, punchline := getJoke()
	clientBuffer := make([]byte, 1024)

	fmt.Printf("[T] Client %d is in the theater\n", clientNumber)

	clientSocket.Write(bytes("Knock knock", "\n")) // starts interaction

	response, err := clientSocket.Read(clientBuffer) // waits for "who's there?"
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			clientSocket.Write(bytes("\nKABOOM!\n"))
			fmt.Printf("[*] Client %d disconnected due to inactivity\n", clientNumber)
			return
		}
		fmt.Printf("[*] Cliente %d closed the connection\n", clientNumber)
		return
	}
	fmt.Printf("[T] Cliente %d says: %s\n", clientNumber, string(clientBuffer[:response]))

	clientSocket.Write(bytes(intro, "\n")) // joke setup

	response, err = clientSocket.Read(clientBuffer) // waits for "who?"
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			clientSocket.Write(bytes("\nKABOOM!\n"))
			fmt.Printf("[*] Client %d disconnected due to inactivity\n", clientNumber)
			return
		}
		fmt.Printf("[*] Cliente %d closed the connection\n", clientNumber)
		return
	}
	fmt.Printf("[T] Cliente %d says: %s\n", clientNumber, string(clientBuffer[:response]))

	clientSocket.Write(bytes(punchline, "\n"))

	theater.Leave()

	fmt.Printf("[T] Connection with client %d is successfully finished\n", clientNumber)
}

func getInQueue(clientSocket net.Conn, theater *Theater, clientNumber int) {
	<-theater.ready
	nextInLine := theater.queue.Dequeue()
	go handleClient(nextInLine, clientNumber, theater)
}

func main() {
	port := 9001
	clientNumber := 0

	queue := &Queue{}
	theater := &Theater{
		peopleInRoom:       0,
		maxTheaterCapacity: 1,
		ready:              make(chan bool),
		queue:              queue,
	}

	serverSocket, port := createServer(port)
	defer serverSocket.Close()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT) // Handles Ctrl+C
	go func() {
		<-sigCh
		fmt.Println("\n[*] Saindo...")
		serverSocket.Close()
		// TODO: close client sockets?
		os.Exit(0)
	}()

	fmt.Printf("[*] Local server on port %d\n", port)

	// Waits indefinitely  for new connections
	for {
		clientSocket, err := serverSocket.Accept()
		if err != nil {
			fmt.Printf("[!] Error accepting connnection from %s: %v\n", clientSocket.RemoteAddr(), err)
			continue
		}

		if theater.IsThereRoom() {
			go handleClient(clientSocket, clientNumber, theater)
			clientNumber++
		} else {
			clientSocket.Write(bytes("The theater is full! You are in line and will be in a moment.", "\n"))
			queue.Enqueue(clientSocket)
			fmt.Printf("[F] Client %d in line\n", clientNumber)
			go getInQueue(clientSocket, theater, clientNumber)
		}
	}
}
