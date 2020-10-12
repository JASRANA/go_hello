package hello

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		// 当测试失败时所报告的行号将在函数调用中而不是在辅助函数内部
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to Spanish people", func(t *testing.T) {
		got := Hello("Carols", "Spanish")
		want := "Hola, Carols"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to French people", func(t *testing.T) {
		got := Hello("Carols", "French")
		want := "Bonjour, Carols"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})
}