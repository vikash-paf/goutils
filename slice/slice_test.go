package slice

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		iteratee func(int) string
		want     []string
	}{
		{
			name:     "nil slice",
			input:    nil,
			iteratee: strconv.Itoa,
			want:     nil,
		},
		{
			name:     "empty slice",
			input:    []int{},
			iteratee: strconv.Itoa,
			want:     []string{},
		},
		{
			name:     "valid slice",
			input:    []int{1, 2, 3},
			iteratee: strconv.Itoa,
			want:     []string{"1", "2", "3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Map(tt.input, tt.iteratee)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleMap() {
	numbers := []int{1, 2, 3, 4}
	strings := Map(numbers, func(n int) string {
		return strconv.Itoa(n * 2)
	})
	fmt.Println(strings)
	// Output: [2 4 6 8]
}
