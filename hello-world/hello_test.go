//It needs to be in a file with a name like xxx_test.go
package main

// In order to use the *testing.T type, you need to import "testing", like we did with fmt in the other file
import "testing"

//The test function must start with the word "Test".
func TestHello(t *testing.T) { //The test function takes one argument only t *testing.T
	got := Hello("Yasin", "English")
	want := "Hello, Yasin"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	t.Run("Hello to people", func(t *testing.T) {
		got := Hello("Yasin", "English")
		want := "Hello, Yasin"
		assertCorrectMessage(t, got, want)

	})

	t.Run("If name empty write 'world' being default", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
