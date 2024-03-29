package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Luka", "")
		want := "Hello, Luka"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

  t.Run("in Serbian", func(t *testing.T){
    got := Hello("Luka", "Serbian")
    want := "Zdravo, Luka"
    assertCorrectMessage(t, got, want)
  })
  
  t.Run("in Spanish", func(t *testing.T){
    got := Hello("Luka", "Spanish")
    want := "Hola, Luka"
    assertCorrectMessage(t, got, want)
  })
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
