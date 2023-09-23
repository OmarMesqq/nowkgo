package main

import (
	"fmt"
	"net"
	"time"
)

func bindPort(port int) (net.Listener, int) {
	// Loop infinito para para criar um socket TCP conectado a uma porta
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

	// Começa a interação
	clientSocket.Write(toBytesSlice("Toc Toc"))

	// Espera "quem é?"
	response, err := clientSocket.Read(clientBuffer)
	if err != nil {
		// Checa se erro é erro de rede. Se for, "ok" é verdade e atribui o erro a netErr.
		// Depois, verifica se "ok" é verdade. Se for, verifica se o erro é de timeout.
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			clientSocket.Write(toBytesSlice("\nVOCE EXPLODIU!\n"))
			fmt.Printf("[*] Cliente %d desconectado por inatividade\n", count)
			return
		}
		fmt.Printf("[*] Cliente %d encerrou a conexão\n", count)
		return
	}

	fmt.Printf("[*] Cliente %d diz: %s\n", count, string(clientBuffer[:response]))

	// Primeira parte da piada
	clientSocket.Write(toBytesSlice(intro))

	// Espera "fulano quem?"
	response, err = clientSocket.Read(clientBuffer)
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

	// Punchline
	clientSocket.Write(toBytesSlice(punchline))

	fmt.Printf("[*] Conexão com cliente %d terminou com sucesso\n", count)
}

func toBytesSlice(str string) []byte {
	return []byte(str)
}

/*
Obs.: A linguagem Go lida com erros de uma forma diferente da utilizada em outras linguagens.
Em Go, erros são valores de retorno de funções. Se uma função retorna um erro,
é responsabilidade de quem a chamou verificar se o erro é nil ou não.
*/
