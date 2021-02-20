package amanz

import "testing"

func TestGet(t *testing.T) {
	want := "test"
	got := "test"

	if got != want {
		t.Errorf("got %v want %v", want, got)
	}
} 