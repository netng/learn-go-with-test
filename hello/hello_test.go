package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("greeting with default message", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("hello people with empty language", func(t *testing.T) {
		got := Hello("Nandang", "")
		want := "Hello, Nandang"
		assertCorrectMessage(t, got, want)
	})

	t.Run("hello people with a french language", func(t *testing.T) {
		got := Hello("Nandang", "French")
		want := "Bonjour, Nandang"
		assertCorrectMessage(t, got, want)
	})

	t.Run("hello people with a spanish language", func(t *testing.T) {
		got := Hello("Nandang", "Spanish")
		want := "Hola, Nandang"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
