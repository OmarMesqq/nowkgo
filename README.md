## Ideia de produto: 
Um teatro de piadas estilo *toc toc* que pode ser acessado por um cliente local. 

## O que aprendi:
- O que são sockets e boas práticas (conexões TCP)
- Abertura e controle de múltiplas threads (ambientes multithread)

## Como utilizar:
1)  `git clone https://github.com/OmarMesqq/socketcom` 

2) `cd socketcom` 

3) `python3 server.py` 

4) Em outro terminal: 
    1) `cd socketcom` 
    2) Compile o cliente: `gcc -o client client.c`
    3) Rode o cliente: `./client`

5) Você se conectou ao servidor!

## Referências:
- [Documentação Python - socket](https://docs.python.org/3/library/socket.html) 
- [Documentação Python - threading](https://docs.python.org/3/library/threading.html)
