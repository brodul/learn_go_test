package main

import "fmt"

func Hello(input string) string {
	return "Hello, " + input
}

func main() {
	fmt.Println(Hello("banana"))
}
