## Ideia de produto: 
Um teatro de piadas estilo *toc toc* que pode ser acessado por um cliente local. 

## O que aprendi:
- O que são sockets e boas práticas (conexões TCP)
- Abertura e controle de múltiplas threads (ambientes multithread)

## Como utilizar:
1)  `git clone https://github.com/OmarMesqq/socketcom` 

2) `cd socketcom` 

3) `cd server`

4) Compile o servidor: `go build -o server` 

5) Rode o servidor: `./server`

6) Em outro terminal: 
    1) `cd client` 
    2) Compile o cliente: `gcc -o client client.c`
    3) Rode o cliente: `./client`

7) Você se conectou ao servidor!

## Referências:
- [Documentação Python - socket](https://docs.python.org/3/library/socket.html) 
- [Documentação Python - threading](https://docs.python.org/3/library/threading.html)
- [Go - net ](https://pkg.go.dev/net)
- [Go - bytes](https://pkg.go.dev/bytes) 
- [Go - slices](https://go.dev/blog/slices-intro) 
- [Go - defer](https://www.digitalocean.com/community/tutorials/understanding-defer-in-go)
- [Go - testing](https://pkg.go.dev/testing#T)