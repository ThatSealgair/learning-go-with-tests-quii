package main

import "fmt"

const englishHelloPrefix = "Hello, "
const fullStop = "."

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name + fullStop
}

func main() {
	fmt.Println(Hello("world"))
}
