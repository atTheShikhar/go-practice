package main

import "fmt"

const enPrefix = "Hello, "
const esPrefix = "Hola, "
const frPrefix = "Bonjour, "
const es = "Spanish"
const fr = "French"

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == es {
		return esPrefix + name
	}
	if language == fr {
		return frPrefix + name
	}
	return enPrefix + name
}

func main() {
	fmt.Println(Hello("Shikhar", ""))
}
