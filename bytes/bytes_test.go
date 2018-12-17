package bytes

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEditDistance(t *testing.T) {
	var tests = []struct {
		a        []byte
		b        []byte
		distance int
	}{
		{[]byte("this is a test"), []byte("wokka wokka!!!"), 37},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Edit distance %d", i), func(t *testing.T) {
			if got, want := EditDistance(test.a, test.b), test.distance; got != want {
				t.Errorf("Expected edit distance %d, but got %d", want, got)
			}
		})
	}
}

func TestCycledSplit(t *testing.T) {
	var tests = []struct {
		data []byte
		n    int
		want [][]byte
	}{
		{[]byte{1, 2, 3, 4, 5}, 1, [][]byte{{1, 2, 3, 4, 5}}},
		{[]byte{1, 2, 3, 4, 5}, 2, [][]byte{{1, 3, 5}, {2, 4}}},
		{[]byte{1, 2, 3, 4, 5}, 5, [][]byte{{1}, {2}, {3}, {4}, {5}}},
		{[]byte{1, 2, 3, 4, 5}, 6, [][]byte{{1}, {2}, {3}, {4}, {5}, {}}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("CycledSplit of %d for %v", test.n, test.data), func(t *testing.T) {
			if got, want := CycledSplit(test.data, test.n), test.want; !reflect.DeepEqual(got, want) {
				t.Errorf("Expected %v, but got %v", want, got)
			}
		})
	}
}
