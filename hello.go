package main

import "fmt"

const enHello = "Hello"   // English
const frHello = "Bonjour" // French
const itHello = "Ciao"    // Italian
const deHello = "Hallo"   // German
const esHello = "Hola"    // Spanish

// Coder's note: performance-wise, slices are best, then switches, maps last
var helloMultilingualMap = map[string]string{
	"en": enHello,
	"fr": frHello,
	"it": itHello,
	"de": deHello,
	"es": esHello,
}

// Accepts any string for the name and an ISO 639-1 language code (as string)
func Hello(name string, language string) string {
	greeting := "Hello" // default greeting
	if name == "" {
		name = "World" // default name
	}
	for k, v := range helloMultilingualMap {
		if language == k {
			greeting = v
		}
	}
	return greeting + ", " + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
