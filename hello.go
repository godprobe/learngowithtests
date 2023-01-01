package main

import "fmt"

const enHello = "Hello" // English
const esHello = "Hola"  // Spanish

var helloMultilingualMap = map[string]string{
	"en": enHello,
	"es": esHello,
}

func Hello(name string, language string) string {
	greeting := "Hello"
	if name == "" {
		name = "World"
	}
	for k, v := range helloMultilingualMap {
		if language == k {
			greeting = v
		}
	}
	return greeting + ", " + name
}

func main() {
	fmt.Println(Hello("world", "es"))
}
