package main

import (
	"fmt"
)

const (
	french  = "French"
	spanish = "Spanish"

	englishHelloPrefix = "Hello, "
	frenchHelloPrefix  = "Bonjour, "
	spanishHelloPrefix = "Hola, "
)

func main() {
	fmt.Println(Hello("Kazu", ""))
}

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingWithPrefix(language) + name
}

func greetingWithPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return

}
