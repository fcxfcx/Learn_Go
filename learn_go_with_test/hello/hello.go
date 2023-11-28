package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "English":
		prefix = englishHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	}
	return
}

func main() {
	fmt.Print(Hello("Chris", "English"))
}
