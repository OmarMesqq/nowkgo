import socket 
import threading
from punchlines import getPunchline

HOST = "127.0.0.1"

def tostr(bytes):
    """ 
    Converte bytes para string
    """
    return bytes.decode('utf-8')


def grabPort(server, port):
    """    
    Tenta bindar o socket do servidor à porta fornecida. 
    Caso esteja ocupada, tenta a próxima disponível.
    """
    try:
        server.bind((HOST,port))
        return  
    except OSError:
        port += 1 
        grabPort(server, port) 


def main(): 
    """
    Inicializa o servidor
    """
    port = 9001
    
    server = socket.socket() 
    grabPort(server, port)

    server.listen(50)     # Avaliar backlog e conexões simultâneas
    print(f"[*] Esperando conexões em {HOST}:{port}")
    count = 0

    while True: 
        conn, addr = server.accept()    # Bloqueante 
        count += 1
        threading.Thread(target=handle_client, args=(conn, addr, count)).start() # Avaliar jogar start() para linha de baixo
        
        
def handle_client(conn, addr, count):
    """
    Lida com a conexão de um cliente
    """
    print(f"[*] Cliente novo na porta {addr[1]}")

    conn.sendall(b"\nToc Toc\n") 

    msg = conn.recv(1024)     # Bloqueante
    msg = tostr(msg)    # Fazer decode aqui 
    print(f"[*] Cliente {count} diz: {msg}") 

    conn.sendall(getPunchline())
 
    conn.close() 
    print(f"[*] Conexão com cliente {addr} terminou com sucesso")
    


if __name__ == '__main__':
    try:
        main() 
    except KeyboardInterrupt: 
        print("\n[*] Saindo...")
        exit(0)


## Considerações:
    # Sistema de filas do teatro
    # Implementar disconnect message (com while ou return)
    # Timeout p/ espera no recv do cliente
    # Sendall pode jogar exceção

# Features futuras:
    # Client to client 
    # Receber non strings (json ou pickle)