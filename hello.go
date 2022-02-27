package main

const englishHelloPrefix = "Hello, "
const spanishHelloPerefix = "Hola, "
const frenchHelloPerefix = "Bonjour, "
const germanHelloPerefix = "Hallo, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = spanishHelloPerefix
	case "French":
		prefix = frenchHelloPerefix
	case "German":
		prefix = germanHelloPerefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
