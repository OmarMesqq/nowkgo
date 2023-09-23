package main

import (
	"math/rand"
)

func getJoke() (string, string) {
	setups := []string{
		"A mosca",
		"O cavalo",
		"A uva",
		"O macaco",
		"O pato",
		"A aranha",
		"O milho",
		"O elefante",
		"A abelha",
		"O canguru",
		"O cachorro",
		"O mané",
	}
	punchlines := []string{
		"O azar da mosca!",
		"O cavalo que foi fazer compras na ferragem!",
		"A uva dando tchauzinho!",
		"O macaco que foi cortar o cabelo!",
		"O pato praticando natação!",
		"A aranha batendo palmas!",
		"O milho assando na fogueira!",
		"O elefante jogando pingue-pongue!",
		"A abelha tomando banho de sol!",
		"O canguru que veio entregar uma carta!",
		"Que comeu sua branch!",
		"Que programa!",
	}

	jokes := make(map[string]string)

	for i := range setups {
		jokes[setups[i]] = punchlines[i]
	}

	size := len(setups)
	key := setups[rand.Intn(size)]
	punchline := jokes[key]

	return key, punchline
}
