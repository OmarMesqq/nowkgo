import socket 

def tostr(bytes):
    """ 
    Converte bytes para string
    """
    return bytes.decode('utf-8')

# Cria socket do servidor 
server = socket.socket() 

# Binda (trava) o socket a uma porta
server.bind(("127.0.0.1", 9001)) 

# Fica escutando por conexões 
server.listen() 

# Pega o objeto de conexão e o endereço do cliente -> bloqueante 
conn, addr = server.accept() 

print(f"Novo cliente conectado: {addr}\n") # Isso fica aqui? 

while True: 
    # Recebe a mensagem do cliente -> Bloqueante
    msg = conn.recv(1024)  

    # Se o cliente fechar a conexão, o recv retorna uma string vazia 
    if not msg: 
        break 

    # Converte a mensagem de bytes para string e a mostra no console server side
    msg = tostr(msg) 
    print(f"Cliente diz: {msg}") 

    # Envia uma mensagem de volta
    conn.sendall(b"Toc toc\n") 

    # Espera e processa a próxima mensagem 
    msg = conn.recv(1024) 
    msg = tostr(msg)
    print(f"Cliente diz: {msg}") 

    # Envia uma mensagem de volta 
    conn.sendall(b"A galinha! KKKKKKKKKKKKK trolei\n")

    # Fecha a conexão 
    conn.close() 