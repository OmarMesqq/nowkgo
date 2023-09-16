package main

import (
	"fmt"
	"net"
	"time"
)

func getJoke() (string, string) {
	// TO-DO: implement
	return "not", "implemented"
}

// Bind the socket to the given port, or find the next available port if the given port is busy
func bindPort(port int) (net.Listener, int) {
	for {
		listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
		if err == nil {
			return listener, port
		}
		port++
	}
}

func main() {
	// Initialize port and listener
	port := 9001
	listener, boundPort := bindPort(port)
	fmt.Printf("[*] Servidor ligado em 127.0.0.1:%d\n", boundPort)

	// Initialize counter
	count := 0

	// Wait for incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("[!] Erro ao aceitar nova conexão:", err)
			continue
		}

		// Increment counter
		count++

		go handleClient(conn, count) // Handle client in a new goroutine
	}
}

// handleClient handles incoming client connections
func handleClient(conn net.Conn, count int) {
	defer conn.Close()

	// Set a timeout for the connection
	conn.SetDeadline(time.Now().Add(1 * time.Minute))

	fmt.Printf("[*] Cliente novo (número %d) na porta %s\n", count, conn.RemoteAddr())

	intro, punchline := getJoke()

	// Send knock knock
	conn.Write([]byte("Toc Toc\n"))

	// Read client's response
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			conn.Write([]byte("\nVOCE EXPLODIU!\n"))
			fmt.Printf("[*] Cliente %d desconectado por inatividade\n", count)
			return
		}
		fmt.Printf("[*] Cliente %d encerrou a conexão\n", count)
		return
	}

	// Send the joke intro
	conn.Write([]byte(intro + "\n"))

	// Read client's response again
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Printf("[*] Cliente %d encerrou a conexão\n", count)
		return
	}

	// Send the joke punchline
	conn.Write([]byte(punchline + "\n\n"))

	fmt.Printf("[*] Conexão com cliente %d terminou com sucesso\n", count)
}
