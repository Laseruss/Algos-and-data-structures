package add_until

import "testing"

func TestAddUntilX(t *testing.T) {
	got := addUntilX([]int{3, 4, 7}, 5)
	want := 4

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
