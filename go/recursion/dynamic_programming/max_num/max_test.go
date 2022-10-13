package max

import "testing"

func TestMaxNum(t *testing.T) {
	tests := []struct {
		Name string
		Nums []int
		Max  int
	}{
		{"testing with numbers", []int{1, 2, 3, 4}, 4},
		{"testing with numbers in reverse order", []int{4, 3, 2, 1}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := maxNum(tt.Nums)
			want := tt.Max

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}

}
