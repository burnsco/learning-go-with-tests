package main

import "testing"

func TestHello(q *testing.T) {
	q.Run("says hello when given name", func(t *testing.T) {
		got := Hello("Corey")
		want := "Hello, Corey"

		assertCorrectMessage(t, got, want)

	})
	q.Run("says hello world when no name given", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(q testing.TB, got, want string) {
	q.Helper()
	if got != want {
		q.Errorf("got %s want %s", got, want)
	}

}
