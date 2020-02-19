package main

import "testing"

/*test funcs must begin with Test
and they only take one arg t *testing.T
*/
func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := hello("Chris", "English")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World', when an empty string is supplied", func(t *testing.T) {
		got := hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := hello("Marcel", "French")
		want := "Bonjour, Marcel"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in korean", func(t *testing.T) {
		got := hello("Shinmyeong", "Korean")
		want := "Annyeong, Shinmyeong"
		assertCorrectMessage(t, got, want)
	})
}
