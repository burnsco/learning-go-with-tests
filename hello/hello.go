package main

const ENGLISH_GREETING = "Hello, "
const SPANISH_GREETING = "Hola, "
const FRENCH_GREETING = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World!"
	}

	greetingType := ENGLISH_GREETING

	switch language {
	case "Spanish":
		greetingType = SPANISH_GREETING
	case "French":
		greetingType = FRENCH_GREETING
	}
	return greetingType + name

}

func main() {
	println(Hello("Corey", "English"))
}
