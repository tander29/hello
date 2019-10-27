package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Travis", "")
		want := "Hello, Travis"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'hello world' when no name supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hola in spanish", func(t *testing.T) {
		got := Hello("Travis", "Spanish")
		want := "Hola, Travis"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying Bonjour in French", func(t *testing.T) {
		got := Hello("Travis", "French")
		want := "Bonjour, Travis"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in german", func(t *testing.T) {
		got := Hello("Travis", "German")
		want := "Guttentag, Travis"

		assertCorrectMessage(t, got, want)
	})
}
