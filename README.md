## Um servidor de piadas estilo *toc toc* 

### Como utilizar:
1)  `git clone https://github.com/OmarMesqq/socketcom` 

2) `cd socketcom` 

3) `python3 server.py` 

4) Em outro terminal, faça `telnet localhost 9001` ou a porta que lhe for informada pelo servidor.

5) Converse com o servidor e se divirta!

## Funcionamento: 
### O código cria um socket que se liga a porta 9001 ou acima e age como servidor local. Os clientes - também locais - se conectam a ela por [`telnet`](https://en.wikipedia.org/wiki/Telnet).