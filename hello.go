package main

import "fmt"

const enHello = "Hello" // English
const esHello = "Hola"  // Spanish

var helloMultilingualMap = map[string]string{
	"en": enHello,
	"es": esHello,
}

// Accepts any string for the name and an ISO 639-1 language code (as string)
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
