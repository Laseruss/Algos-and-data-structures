package fibonacci

import "testing"

func TestNthFibonacci(t *testing.T) {
	tests := []struct {
		testname string
		n        uint
		want     uint
		m        map[uint]uint
	}{
		{"testing for 6", 6, 8, map[uint]uint{}},
		{"testing for 0", 0, 0, map[uint]uint{}},
	}

	for _, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
			got := bottomUpFibonacci(tt.n)
			want := tt.want

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
