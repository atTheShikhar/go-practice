package main

import (
	"fmt"
)

const enPrefix = "Hello, "
const esPrefix = "Hola, "
const frPrefix = "Bonjour, "
const es = "Spanish"
const fr = "French"

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case fr:
		prefix = frPrefix
	case es:
		prefix = esPrefix
	default:
		prefix = enPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Shikhar", ""))
}
