package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		sum := Repeat("a", 5)
		expected := "aaaaa"

		if sum != expected {
			t.Errorf("expected '%q' but got '%q'", expected, sum)
		}
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("a", 5))
	// Output: "aaaaa"
}
