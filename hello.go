package main

import "fmt"

func main() {
	fmt.Println(Hello("world", ""))
}

const englishHelloPrefix = "Hello, "
const spanishHelloPerefix = "Hola, "
const frenchHelloPerefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	if language == "Spanish" {
		return spanishHelloPerefix + name
	}
	if language == "French" {
		return frenchHelloPerefix + name
	}
	return englishHelloPrefix + name
}
