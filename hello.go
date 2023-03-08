package main

import "fmt"

const prefix = "Hello, "

func Hello(name string) string {
	n := "world"
	if name != "" {
		n = name
	}
	return prefix + n
}

func main() {
	fmt.Println(Hello("Shikhar"))
}
