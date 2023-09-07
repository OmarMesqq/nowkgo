import socket 
import threading
from jokes import getJoke

HOST = "127.0.0.1"
FORMAT = 'utf-8'
DISCLAIMER = '''
Olá!
O servidor de piadas te recebe de braços abertos! 
Você tem um minuto entre conversas antes de explodir por inatividade :)
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
    
    server.listen(5)  
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
        conn.sendall(b"> " + b"Toc Toc") 

        msg = conn.recv(1024)     # Bloqueante
        msg = msg.decode(FORMAT)
        print(f"[*] Cliente {count} diz: {msg}") 

        conn.sendall(b"> " + intro)

        msg = conn.recv(1024)     # Bloqueante
        msg = msg.decode(FORMAT) 
        print(f"[*] Cliente {count} diz: {msg}")

        conn.sendall(b"> " + punchline)
 
        print(f"[*] Conexão com cliente {count} terminou com sucesso")
    except socket.timeout:
        conn.sendall(b"> " + b"VOCE EXPLODIU!")
        print(f"[*] Cliente {count} desconectado por inatividade")
    except (BrokenPipeError, UnicodeDecodeError):
        print(f"[*] Cliente {count} encerrou a conexão") 
    except Exception as e:
        print(f"[!] Erro desconhecido entre cliente e servidor:\n", e)
    finally:
        conn.close()
    

if __name__ == '__main__':
    try:
        main() 
    except KeyboardInterrupt: 
        print("\n[*] Saindo...")
        exit(0)


## TO-DO:
    # Sistema de filas do teatro
    # Avaliar backlog/conexões simultâneas e rate limit atrelado a um numero baixo no server.listen()
