package main

const GREETING = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World!"
	}
	return GREETING + name
}

func main() {
	println(Hello("Corey"))
}
