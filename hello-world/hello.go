package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const korean = "Korean"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const koreanHelloPrefix = "Annyeong, "

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case korean:
		prefix = koreanHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return

}

//hello takes in a string arg and returns a string
func hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name

}

func main() {
	fmt.Println(hello("world", "English"))

}
