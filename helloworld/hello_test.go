package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("world!", "")
		want := "Hello, world!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World! when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)

	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Mbappe", "French")
		want := "Bonjour, Mbappe"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Azerbaijan", func(t *testing.T) {
		got := Hello("Ali", "Azerbaijan")
		want := "Salam, Ali"

		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
