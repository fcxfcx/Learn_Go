package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(got, want string, t *testing.T) {
		t.Helper()
		if want != got {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("in English", func(t *testing.T) {
		want := "Hello, Chris"
		got := Hello("Chris", "English")
		assertCorrectMessage(got, want, t)
	})

	t.Run("in English", func(t *testing.T) {
		want := "Hola, Elodie"
		got := Hello("Elodie", "Spanish")
		assertCorrectMessage(got, want, t)
	})

	t.Run("in French", func(t *testing.T) {
		want := "Bonjour, Alberta"
		got := Hello("Alberta", "French")
		assertCorrectMessage(got, want, t)
	})

	t.Run("empty Name", func(t *testing.T) {
		want := "Hello, World"
		got := Hello("", "English")
		assertCorrectMessage(got, want, t)
	})
}
