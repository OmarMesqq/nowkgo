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
        print(f"[*] Esperando conexões em {HOST}:{port}") 
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
    count = 0

    while True: 
        conn, addr = server.accept()
        count += 1
        threading.Thread(target=handle_client, args=(conn, addr, count)).start()
        
        
def handle_client(conn, addr, count):
    """
    Lida com a conexão de um cliente
    """
    print(f"[*] Cliente local novo na porta {addr[1]}")

    while True:    # não precisa desse laço
        # Inicia a conversa
        conn.sendall(b"\nToc Toc\n") 

        # Espera e processa a próxima mensagem (bloqueante)
        msg = conn.recv(1024)     # Timeout p/ espera
        msg = tostr(msg)
        print(f"[*] Cliente {count} diz: {msg}") 

        # Envia a punchline de volta
        conn.sendall(getPunchline())

        # Fecha a conexão 
        conn.close() 
        print(f"[*] Conexão com cliente {addr} terminou com sucesso")
        break


if __name__ == '__main__':
    try:
        main() 
    except KeyboardInterrupt: 
        print("\n[*] Saindo...")
        exit(0)

# Sistema de filas do teatro
