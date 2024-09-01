package main

const (
	spanish = "Spanish"
	french  = "French"

	ENGLISH_GREETING = "Hello, "
	SPANISH_GREETING = "Hola, "
	FRENCH_GREETING  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World!"
	}

	return greetingType(language) + name

}

func greetingType(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = SPANISH_GREETING
	case french:
		prefix = FRENCH_GREETING
	default:
		prefix = ENGLISH_GREETING
	}
	return
}

func main() {
	println(Hello("Corey", "English"))
}
