package main

import "fmt"

const (
	englishGreetingPrefix = "Hello "
	spanishGreetingPrefix = "Hola "
	frenchGreetingPrefix  = "Bonjour "
	Spanish               = "spanish"
	English               = "english"
	French                = "french"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return getLanguageGreeting(language) + name
}

func getLanguageGreeting(language string) (greeting string) {
	switch language {
	case Spanish:
		greeting = spanishGreetingPrefix
	case French:
		greeting = frenchGreetingPrefix
	default:
		greeting = englishGreetingPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Rodrigo", English))
}
