import random

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

random_punchline = f"{random.choice(punchlines)}" + "\n\n"

def getPunchline() -> bytes:
    return random_punchline.encode('utf-8')