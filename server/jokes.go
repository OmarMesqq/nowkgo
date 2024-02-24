package main

import (
	"math/rand"
)

func getJoke() (string, string) {
	setups := []string{
		"The fly",
		"The horse",
		"The grape",
		"The monkey",
		"The duck",
		"The spider",
		"The corn",
		"The elephant",
		"The bee",
		"The kangaroo",
		"The dog",
		"The fool",
	}
	punchlines := []string{
		"The fly's bad luck!",
		"The horse went shopping at the hardware store!",
		"The grape saying goodbye!",
		"The monkey went to get a haircut!",
		"The duck practicing swimming!",
		"The spider clapping hands!",
		"The corn roasting on the fire!",
		"The elephant playing ping-pong!",
		"The bee sunbathing!",
		"The kangaroo that came to deliver a letter!",
		"Who ate your branch!",
		"That's programming!",
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
