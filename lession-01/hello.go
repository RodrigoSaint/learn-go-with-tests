package main

import "fmt"

const englishGreetingPrefix = "Hello "
const spanishGreetingPrefix = "Hola "
const frenchGreetingPrefix = "Bonjour "
const spanish = "spanish"
const english = "english"
const french = "french"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	greeting := englishGreetingPrefix

	switch language {
	case spanish:
		greeting = spanishGreetingPrefix
	case french:
		greeting = frenchGreetingPrefix
	}

	return greeting + name
}

func main() {
	fmt.Println(Hello("Rodrigo", ""))
}
