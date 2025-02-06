package main

import "fmt"

// group constants in a block rather than in their own line
const (
	spanish = "Spanish"
	french = "French"
	
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

// (return prefix) can return to whatever we set to, in this case it's prefix
// private function starts with a lowercase letter
// public function starts with a capital letter
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Chris", ""))
}