package main

import "testing"

/*test funcs must begin with Test
and they only take one arg t *testing.T
*/
func TestHello(t *testing.T) {
	got := hello("Chris")
	want := "Hello, Chris"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
