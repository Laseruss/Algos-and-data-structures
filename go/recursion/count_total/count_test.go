package count

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	tests := []struct {
		l    []string
		want int
	}{
		{[]string{"hello", "my", "name", "is"}, 13},
		{[]string{}, 0},
		{[]string{"hello"}, 5},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%q", tt.l)
		t.Run(testname, func(t *testing.T) {
			ans := countTotal(tt.l)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
