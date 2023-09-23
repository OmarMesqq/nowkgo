### Uma nota sobre checagem de erros em Go

A linguagem Go lida com erros de uma forma diferente da utilizada em outras linguagens.
Em Go, erros são valores de retorno de funções. Se uma função retorna um erro,
é responsabilidade de quem a chamou verificá-lo e, se aplicável, tratá-lo.

Essa pode ser uma forma estranha de lidar com erros, mas é uma forma muito eficiente. 
Em Go, não é necessário utilizar exceções para lidar com erros, o que torna o código mais simples e legível. 