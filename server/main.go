package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func createServer(port int) (net.Listener, int) {
	for {
		serverSocket, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
		if err == nil {
			return serverSocket, port
		}
		port++
	}
}

func main() {
	port := 9001
	clientNumber := 0

	theater := &Theater{
		peopleInRoom:       0,
		maxTheaterCapacity: 10,
	}

	queue := &Queue{}

	serverSocket, port := createServer(port)
	defer serverSocket.Close()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT) // Avisa quando o usuário apertar Ctrl+C
	go func() {
		<-sigCh
		fmt.Println("\n[*] Saindo...")
		serverSocket.Close()
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
			clientNumber++
			go handleClient(clientSocket, clientNumber, theater)
		} else {
			go sendToQueue(clientSocket, queue)
		}

	}
}

func sendToQueue(clientSocket net.Conn, queue *Queue) {
	defer clientSocket.Close()
	clientSocket.Write(toBytesSlice("O teatro está cheio! Fique na fila e você já entra"))

	fmt.Printf("Cliente %s entrou na fila", clientSocket.RemoteAddr())

	queue.Enqueue(clientSocket.RemoteAddr())
}

func handleClient(clientSocket net.Conn, clientNumber int, theater *Theater) {
	defer clientSocket.Close()
	clientSocket.SetDeadline(time.Now().Add(5 * time.Minute))

	intro, punchline := getJoke()
	clientBuffer := make([]byte, 1024)

	fmt.Printf("[*] Cliente novo (número %d) no endereço: %s\n", clientNumber, clientSocket.RemoteAddr())

	clientSocket.Write(toBytesSlice("Toc Toc")) // começa a interação

	response, err := clientSocket.Read(clientBuffer) // espera "quem é?"
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			clientSocket.Write(toBytesSlice("\nVOCE EXPLODIU!\n"))
			fmt.Printf("[*] Cliente %d desconectado por inatividade\n", clientNumber)
			return
		}
		fmt.Printf("[*] Cliente %d encerrou a conexão\n", clientNumber)
		return
	}
	fmt.Printf("[*] Cliente %d diz: %s\n", clientNumber, string(clientBuffer[:response]))

	clientSocket.Write(toBytesSlice(intro)) // primeira parte da piada

	response, err = clientSocket.Read(clientBuffer) // espera "fulano quem?"
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			clientSocket.Write(toBytesSlice("\nVOCE EXPLODIU!\n"))
			fmt.Printf("[*] Cliente %d desconectado por inatividade\n", clientNumber)
			return
		}
		fmt.Printf("[*] Cliente %d encerrou a conexão\n", clientNumber)
		return
	}
	fmt.Printf("[*] Cliente %d diz: %s\n", clientNumber, string(clientBuffer[:response]))

	clientSocket.Write(toBytesSlice(punchline))

	fmt.Printf("[*] Conexão com cliente %d terminou com sucesso\n", clientNumber)

	theater.peopleInRoom -= 1
}

func toBytesSlice(str string) []byte {
	return []byte(str)
}
