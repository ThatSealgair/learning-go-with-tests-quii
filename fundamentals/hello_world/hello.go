package main

import "fmt"

const english = "English"
const french = "French"
const spanish = "Spanish"

const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "

const fullStop = "."

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name + fullStop
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", spanish))
}
