package main

import (
	"fmt"
	"strings"
)

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func greetingPrefix(language string) (prefix string) {
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

func Hello(name string, language string) string {
	if strings.TrimSpace(name) == "" {
		name = "World"
	}

	prefix := greetingPrefix(language)

	return prefix + name

}

func main() {
	fmt.Println(Hello("Chris", ""))
}
