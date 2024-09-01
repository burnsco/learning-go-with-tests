package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("says hello when given name", func(t *testing.T) {
		got := Hello("Corey", "")
		want := "Hello, Corey"

		assertCorrectMessage(t, got, want)

	})
	t.Run("says hello world when no name given", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("says hello in Spanish", func(t *testing.T) {
		got := Hello("Frank", "Spanish")
		want := "Hola, Frank"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}
