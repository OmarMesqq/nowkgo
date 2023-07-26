import socket 

def tostr(bytes):
    return bytes.decode('utf-8')

# Cria socket do cliente
client = socket.socket()

# Conecta ao servidor
client.connect(("127.0.0.1", 9001))
