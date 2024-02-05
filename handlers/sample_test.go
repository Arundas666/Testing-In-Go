package handlers

import (
	"testing"
) 

func TestHello(t *testing.T) {
	t.Run("say hello", func(t *testing.T) {

		got := Hello("Arun")
		want := "Hey, Myself Arun"

		assertCorrectMessage(t, got, want)

	})

	t.Run("say hello , passing empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hey, Myself "
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func TestAdd(t *testing.T) {
	got := Add(2, 2)
	want := 4
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {

		Repeat("a")
	}

}
