package main

import "fmt"

const englishPrefix = "Hello, "
const spanish = "Spanish"
const spanishPrefix = "Hola, "
const french = "French"
const frenchPrefix = "Bonjour, "

func Hello(input, language string) string {
	prefix := englishPrefix
	if input == "" {
		input = "world"
	}
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	}
	return prefix + input
}

func main() {
	fmt.Println(Hello("banana", ""))
}
