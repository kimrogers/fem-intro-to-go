package utils

import "testing"

func Test_Add(t *testing.T) {
	got := Add(1, 4, 19, 6)
	want := 30
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
