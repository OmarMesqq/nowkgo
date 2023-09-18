package main

import (
	"fmt"
	"net"
	"time"
)

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
	port := 9001
	listener, boundPort := bindPort(port)
	fmt.Printf("[*] Servidor ligado em 127.0.0.1:%d\n", boundPort)

	count := 0

	// Espera por novas conexões
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("[!] Erro ao aceitar nova conexão:", err)
			continue
		}
		count++
		go handleClient(conn, count)
	}
}

func handleClient(conn net.Conn, count int) {
	defer conn.Close()

	// 1min entre cada interação
	conn.SetDeadline(time.Now().Add(1 * time.Minute))

	fmt.Printf("[*] Cliente novo (número %d) na porta %s\n", count, conn.RemoteAddr())

	intro, punchline := getJoke()

	// Começa a interação
	conn.Write([]byte("Toc Toc\n"))

	// Espera "quem é?"
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

	// Primeira parte da piada
	conn.Write([]byte(intro + "\n"))

	// Espera "fulano quem?"
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Printf("[*] Cliente %d encerrou a conexão\n", count)
		return
	}

	// Punchline
	conn.Write([]byte(punchline + "\n\n"))

	fmt.Printf("[*] Conexão com cliente %d terminou com sucesso\n", count)
}
