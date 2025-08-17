package helloworld

import "fmt"

// language
const spanish = "Spanish"
const french = "French"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := greetingPrefix(language)
	return prefix + name
}

func greetingPrefix(language string) string {
	switch language {
	case spanish:
		return spanishHelloPrefix
	case french:
		return frenchHelloPrefix
	}

	return englishHelloPrefix
}

func main() {
	fmt.Println(Hello("Chris", french))
}
