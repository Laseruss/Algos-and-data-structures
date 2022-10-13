package triangularnumbers

import "testing"

func TestNthTriangular(t *testing.T) {
	got := nthTriangular(7)
	want := 28

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
