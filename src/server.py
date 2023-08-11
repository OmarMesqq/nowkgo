import socket 
import threading
from punchlines import getJoke

HOST = "127.0.0.1"
FORMAT = 'utf-8'


def bindPort(server, port):
    """    
    Tenta bindar o socket do servidor à porta fornecida. 
    Caso esteja ocupada, tenta a próxima disponível.
    """
    try:
        server.bind((HOST,port))
        return  
    except OSError:
        port += 1 
        bindPort(server, port) 


def main(): 
    """
    Inicializa o servidor
    """
    port = 9001
    
    server = socket.socket()    # Cria socket IPv4 TCP
    bindPort(server, port)

    server.listen(50)  
    print(f"[*] Esperando conexões em {HOST}:{port}")
    count = 0

    while True: 
        conn, addr = server.accept()    # Bloqueante 
        count += 1
        clientThread = threading.Thread(target=handle_client, args=(conn, addr, count))
        clientThread.start()
        
        
def handle_client(conn, addr, count):
    """
    Lida com a conexão de um cliente
    """
    print(f"[*] Cliente novo na porta {addr[1]}")
    intro, punchline = getJoke()

    conn.sendall(b"\nToc Toc\n") 

    msg = conn.recv(1024)     # Bloqueante
    msg = msg.decode(FORMAT)
    print(f"[*] Cliente {count} diz: {msg}") 

    conn.sendall(intro + b"\n")

    msg = conn.recv(1024)     # Bloqueante
    msg = msg.decode(FORMAT) 
    print(f"[*] Cliente {count} diz: {msg}")

    conn.sendall(punchline + b"\n\n")
 
    conn.close() 
    print(f"[*] Conexão com cliente {addr} terminou com sucesso")
    

if __name__ == '__main__':
    try:
        main() 
    except KeyboardInterrupt: 
        print("\n[*] Saindo...")
        exit(0)


## Considerações:
    # Estudar rate limit
    # Cor no terminal
    # Sistema de filas do teatro
    # Melhorar saída do cliente no meio da conexão/implementar disconnect message (com break ou return) p/ finalizar thread
    # Timeout p/ espera no recv do cliente
    # Sendall pode jogar exceção
    # Avaliar backlog/conexões simultâneas e rate limit atrelado a um numero baixo no server.listen()

## Features futuras:
    # Client to client 
    # Receber non strings (json ou pickle)