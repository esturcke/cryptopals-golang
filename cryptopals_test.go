package cryptopals

import (
	"fmt"
	"testing"
)

func TestSolutions(t *testing.T) {
	for _, challenge := range Challenges {
		t.Run(fmt.Sprintf("Challenge %d", challenge.number), func(t *testing.T) {
			if got, want := challenge.solver(), challenge.solution; got != want {
				t.Errorf("Failed to solve challenge %d: Got %s, expected %s", challenge.number, got, want)
			}
		})
	}
}

func BenchmarkSolution(b *testing.B) {
	for _, challenge := range Challenges {
		b.Run(fmt.Sprintf("Challenge %d", challenge.number), func(b *testing.B) {
			challenge.solver()
		})
	}
}
