package main

import (
	"fmt"
	"net"
	"time"
)

func bindPort(port int) (net.Listener, int) {
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
	count := 0

	serverSocket, port := bindPort(port)
	defer serverSocket.Close()

	fmt.Printf("[*] Servidor local ligado na porta %d\n", port)

	// Espera por novas conexões até o programa ser encerrado
	for {
		clientSocket, err := serverSocket.Accept()
		if err != nil {
			fmt.Printf("[!] Erro ao aceitar nova conexão de %s: %v\n", clientSocket.RemoteAddr(), err)
			continue
		}
		count++
		go handleClient(clientSocket, count) //TO-DO: implementar limitador de conexões
	}
}

func handleClient(clientSocket net.Conn, count int) {
	defer clientSocket.Close()

	intro, punchline := getJoke()
	clientBuffer := make([]byte, 1024)

	// Se interação com cliente demorar mais que 5 minutos, encerra a conexão
	clientSocket.SetDeadline(time.Now().Add(5 * time.Minute))

	fmt.Printf("[*] Cliente novo (número %d) no endereço: %s\n", count, clientSocket.RemoteAddr())

	clientSocket.Write(toBytesSlice("Toc Toc")) // começa a interação

	response, err := clientSocket.Read(clientBuffer) // espera "quem é?"
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			clientSocket.Write(toBytesSlice("\nVOCE EXPLODIU!\n"))
			fmt.Printf("[*] Cliente %d desconectado por inatividade\n", count)
			return
		}
		fmt.Printf("[*] Cliente %d encerrou a conexão\n", count)
		return
	}
	fmt.Printf("[*] Cliente %d diz: %s\n", count, string(clientBuffer[:response]))

	clientSocket.Write(toBytesSlice(intro)) // primeira parte da piada

	response, err = clientSocket.Read(clientBuffer) // espera "fulano quem?"
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			clientSocket.Write(toBytesSlice("\nVOCE EXPLODIU!\n"))
			fmt.Printf("[*] Cliente %d desconectado por inatividade\n", count)
			return
		}
		fmt.Printf("[*] Cliente %d encerrou a conexão\n", count)
		return
	}
	fmt.Printf("[*] Cliente %d diz: %s\n", count, string(clientBuffer[:response]))

	clientSocket.Write(toBytesSlice(punchline))

	fmt.Printf("[*] Conexão com cliente %d terminou com sucesso\n", count)
}

func toBytesSlice(str string) []byte {
	return []byte(str)
}
