package onlyevens

import (
	"reflect"
	"testing"
)

func TestOnlyEvens(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{2, 4}},
		{[]int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run("great test", func(t *testing.T) {
			ans := onlyEvens(tt.nums, []int{})
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %q, want %q", ans, tt.want)
			}
		})
	}
}
