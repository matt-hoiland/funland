package main

import "fmt"

var codes = map[rune]string{
	'a': ".-",
	'b': "-...",
	'c': "-.-.",
	'd': "-..",
	'e': ".",
	'f': "..-.",
	'g': "--.",
	'h': "....",
	'j': ".---",
	'k': "-.-",
	'l': ".-..",
	'm': "--",
	'n': "-.",
	'o': "---",
	'p': ".--.",
	'q': "--.-",
	'r': ".-.",
	's': "...",
	't': "-",
	'u': "..-",
	'v': "...-",
	'w': ".--",
	'x': "-..-",
	'y': "-.--",
	'z': "--..",
}

var reverseCodes = make(map[string]rune)

func setup() {
	for k, v := range codes {
		reverseCodes[v] = k
	}
}

func main() {
	setup()

	for k, v := range codes {
		fmt.Printf("%q: %q\n", k, v)
	}

	for k, v := range reverseCodes {
		fmt.Printf("%q: %q\n", k, v)
	}
}
