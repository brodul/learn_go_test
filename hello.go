package main

import "fmt"

const englishPrefix = "Hello, "

func Hello(input string) string {
	if input == "" {
		input = "world"
	}
	return englishPrefix + input
}

func main() {
	fmt.Println(Hello("banana"))
}
