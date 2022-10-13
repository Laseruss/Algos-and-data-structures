package golomb

import "testing"

func TestGolomb(t *testing.T) {
	got := golomb(5, map[int]int{})
	want := 3

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
