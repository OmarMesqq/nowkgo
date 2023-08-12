import socket 
import threading
from punchlines import getJoke

HOST = "127.0.0.1"
FORMAT = 'utf-8'
DISCLAIMER = '''
OLÁ!
O SERVIDOR DE PIADAS TE RECEBE DE BRAÇOS ABERTOS.
DIVIRTA-SE!
Você tem um minuto entre conversas para interagir com o servidor. 
-------------------------------------------------------------------------------------
'''

def bindPort(server, port):
    """    
    Tenta bindar o socket do servidor à porta fornecida. 
    Caso esteja ocupada, tenta a próxima disponível.
    """
    try:
        server.bind((HOST,port))
        print(f"[*] Servidor ligado em {HOST}:{port}")
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
    count = 0
    print("[*] Esperando conexões...")

    while True: 
        conn, addr = server.accept()    # Bloqueante 
        conn.settimeout(60)
        count += 1
        clientThread = threading.Thread(target=handle_client, args=(conn, addr, count))
        clientThread.start()
        
        
def handle_client(conn, addr, count):
    """
    Lida com a conexão de um cliente
    """
    print(f"[*] Cliente novo (número {count}) na porta {addr[1]}")
    intro, punchline = getJoke()

    try: 
        conn.sendall(DISCLAIMER.encode(FORMAT))
        conn.sendall(b"Toc Toc\n") 

        msg = conn.recv(1024)     # Bloqueante
        msg = msg.decode(FORMAT)
        print(f"[*] Cliente {count} diz: {msg}") 

        conn.sendall(intro + b"\n")

        msg = conn.recv(1024)     # Bloqueante
        msg = msg.decode(FORMAT) 
        print(f"[*] Cliente {count} diz: {msg}")

        conn.sendall(punchline + b"\n\n")
 
        print(f"[*] Conexão com cliente {count} terminou com sucesso")
    except socket.timeout:
        conn.sendall(b"\nVOCE FOI DESCONECTADA(O) POR INATIVIDADE\n")
        print(f"[*] Cliente {count} desconectado por inatividade")
    except (BrokenPipeError, UnicodeDecodeError):
        print(f"[*] Cliente {count} encerrou a conexão") 
    except Exception as e:
        print(f"[*] Erro desconhecido entre cliente e servidor:\n", e)
    finally:
        conn.close()
    

if __name__ == '__main__':
    try:
        main() 
    except KeyboardInterrupt: 
        print("\n[*] Saindo...")
        exit(0)


## Considerações:
    # Cor no terminal
    # Sistema de filas do teatro
    # Melhorar saída do cliente no meio da conexão/implementar disconnect message (com break ou return) p/ finalizar thread
    # Avaliar backlog/conexões simultâneas e rate limit atrelado a um numero baixo no server.listen()

## Features futuras:
    # Client to client 
    # Receber non strings (json ou pickle)