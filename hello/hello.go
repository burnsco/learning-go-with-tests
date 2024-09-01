package main

const ENGLISH_GREETING = "Hello, "
const SPANISH_GREETING = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World!"
	}
	if language == "Spanish" {
		return SPANISH_GREETING + name
	}
	return ENGLISH_GREETING + name
}

func main() {
	println(Hello("Corey", "Spanish"))
}
