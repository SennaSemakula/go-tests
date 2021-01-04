package main

import "testing"

func TestGreeting(t *testing.T) {

	assertCorrectMsg := func(t *testing.T, got, want string)  {
		t.Helper() // tells the test suite that this is a helper message so will fail on function that called this
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	// First subtest
	t.Run("Test name greeting", func(t *testing.T) {
		got := Greeting("Chris", "")
		expected := "Hello Chris"
		assertCorrectMsg(t, got, expected)
	})

	t.Run("Test empty name", func(t *testing.T) {
		got := Greeting("", "")
		expected := "Hello World"
		assertCorrectMsg(t, got, expected)
	})

	t.Run("Test Arabic Language", func(t *testing.T) {
		got := Greeting("Senna", "Arabic")
		expected := "Asalaam alaikum Senna"
		assertCorrectMsg(t, got, expected)
	})
	
	t.Run("Test French Language", func(t *testing.T) {
		got := Greeting("Moet", "French")
		expected := "Bonjour Moet"
		assertCorrectMsg(t, got, expected)
	})

}