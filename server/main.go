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

	fmt.Printf("[*] Cliente novo (número %d) no endereço: %s\n", clientNumber, clientSocket.RemoteAddr())

	clientSocket.Write(bytes("Toc Toc", "\n")) // começa a interação

	response, err := clientSocket.Read(clientBuffer) // espera "quem é?"
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			clientSocket.Write(bytes("\nVOCE EXPLODIU!\n"))
			fmt.Printf("[*] Cliente %d desconectado por inatividade\n", clientNumber)
			return
		}
		fmt.Printf("[*] Cliente %d encerrou a conexão\n", clientNumber)
		return
	}
	fmt.Printf("[*] Cliente %d diz: %s\n", clientNumber, string(clientBuffer[:response]))

	clientSocket.Write(bytes(intro, "\n")) // primeira parte da piada

	response, err = clientSocket.Read(clientBuffer) // espera "fulano quem?"
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			clientSocket.Write(bytes("\nVOCE EXPLODIU!\n"))
			fmt.Printf("[*] Cliente %d desconectado por inatividade\n", clientNumber)
			return
		}
		fmt.Printf("[*] Cliente %d encerrou a conexão\n", clientNumber)
		return
	}
	fmt.Printf("[*] Cliente %d diz: %s\n", clientNumber, string(clientBuffer[:response]))

	clientSocket.Write(bytes(punchline, "\n"))

	theater.Leave()

	fmt.Printf("[*] Conexão com cliente %d terminou com sucesso\n", clientNumber)
}

func getInQueue(clientSocket net.Conn, theater *Theater) {
	<-theater.ready
	nextInLine := theater.queue.Dequeue()
	go handleClient(nextInLine, -1, theater)
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
	signal.Notify(sigCh, syscall.SIGINT) // Avisa quando o usuário apertar Ctrl+C
	go func() {
		<-sigCh
		fmt.Println("\n[*] Saindo...")
		serverSocket.Close()
		// fechar sockets dos clientes?
		os.Exit(0)
	}()

	fmt.Printf("[*] Servidor local ligado na porta %d\n", port)

	// Espera por novas conexões
	for {
		clientSocket, err := serverSocket.Accept()
		if err != nil {
			fmt.Printf("[!] Erro ao aceitar nova conexão de %s: %v\n", clientSocket.RemoteAddr(), err)
			continue
		}

		if theater.IsThereRoom() {
			go handleClient(clientSocket, clientNumber, theater)
			clientNumber++
		} else {
			clientSocket.Write(bytes("O teatro está cheio! Você está na fila e já entra", "\n"))
			queue.Enqueue(clientSocket)
			fmt.Printf("Cliente %s entrou na fila", clientSocket.RemoteAddr())
			go getInQueue(clientSocket, theater)
		}
	}
}
