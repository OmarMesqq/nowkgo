## Um servidor de piadas estilo *toc toc* 

### O código cria um socket que se liga a porta 9001 ou acima e age como servidor local. Os clientes - também locais<sup>[1]</sup> - se conectam a ela por [`telnet`](https://en.wikipedia.org/wiki/Telnet). 

### Como utilizar:
1 - `git clone https://github.com/OmarMesqq/socketcom` 

2 - `cd socketcom` 

3 - `python3 server.py` 

4 - Em outro terminal, faça `telnet localhost 9001` ou a porta que lhe for informada pelo servidor.

5 - Converse com o servidor e se divirta!

## Funcionamento: 
### São utilizados os módulos `socket` e `threading` do Python. O primeiro para a comunicação entre cliente e servidor e o segundo para a criação de threads que permitem que o servidor atenda a mais de um cliente simultaneamente. Dessa forma, o programa pode ser executado sem módulos externos ou ambientes virtuais.


### Em tese, o número de clientes simultâneos é limitado apenas pela capacidade da máquina. Para alterar a capacidade máxima de clientes simultâneos, altere a seguinte linha de código: 
<br>

```python 
server.listen(X) # onde X é o limite de clientes simultâneos
```

### [1] Caso deseje receber conexões que não em `localhost`, altere a seguinte linha de código:
<br>

```python
HOST = 0.0.0.0 
```
o que irá permitir que o servidor escute conexões em todas as interfaces de rede.

No entanto, esteja ciente de que esse processo abre uma porta em sua máquina e pode torná-la vulnerável a ataques. Além disso, o protocolo [`telnet` não é seguro para acesso remoto](https://www.makeuseof.com/why-you-should-not-use-telnet/).
