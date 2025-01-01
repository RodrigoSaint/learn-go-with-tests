package main

import "fmt"

const (
	englishGreetingPrefix = "Hello "
	spanishGreetingPrefix = "Hola "
	frenchGreetingPrefix  = "Bonjour "
	spanish               = "spanish"
	english               = "english"
	french                = "french"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return getLanguageGreeting(language) + name
}

func getLanguageGreeting(language string) (greeting string) {
	switch language {
	case spanish:
		greeting = spanishGreetingPrefix
	case french:
		greeting = frenchGreetingPrefix
	default:
		greeting = englishGreetingPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Rodrigo", ""))
}
