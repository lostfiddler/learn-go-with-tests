package main

import "testing"

func TestHello(t *testing.T) {
    t.Run("saying hello to people", func(t *testing.T){
        got := Hello("Kat", "")
        want := "Hello, Kat"
        assertCorrectMessage(t, got, want)
    })
    t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T){
        got := Hello("", "")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
    })
    t.Run("in Spanish", func(t *testing.T){
        got := Hello("Marina", "Spanish")
        want := "Hola, Marina"
        assertCorrectMessage(t, got, want)
    })
    t.Run("in French", func(t *testing.T){
        got := Hello("Ian", "French")
        want := "Bonjour, Ian"
        assertCorrectMessage(t, got, want)
    })
}

func assertCorrectMessage(t testing.TB, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
