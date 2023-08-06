import random
FORMAT = 'utf-8'

def getJoke():
    """
    Retorna uma tupla em bytes com a chave e o valor da piada
    """
    key =  random.choice(animals)
    bytes_key = key.encode(FORMAT)
    bytes_punchline = jokes[key].encode(FORMAT)
    return (bytes_key, bytes_punchline)


animals = [
    "A mosca!",
    "O cavalo!",
    "A uva!",
    "O macaco!",
    "O pato!",
    "A aranha!",
    "O milho!",
    "O elefante!",
    "A abelha!",
    "O canguru!"
]

punchlines = [
    "O azar da mosca!",
    "O cavalo que foi fazer compras na ferragem!",
    "A uva dando tchauzinho!",
    "O macaco que foi cortar o cabelo!",
    "O pato praticando natação!",
    "A aranha batendo palmas!",
    "O milho assando na fogueira!",
    "O elefante jogando pingue-pongue!",
    "A abelha tomando banho de sol!",
    "O canguru que veio entregar uma carta!"
]


# Dicionário com compreensão para as piadas
# Um jeito pythonico de declarar um dicionário
jokes = {animal: punchline for animal, punchline in zip(animals, punchlines)}


# Piadas internas da empresa
# calindra = ["Cachorro", "Mané"]
# calindra_punchlines = ["Que comeu sua branch", "que programa"]