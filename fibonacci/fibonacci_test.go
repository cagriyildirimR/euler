package fibonacci

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {

	for _, test := range []struct {
		limit int
		want  int
	}{
		{4000000, 4613732},
	} {
		if got := FibonacciEvenSum(test.limit); test.want != got {
			t.Errorf("Sum value of all fibonacci numbers smaller than limit:%v should be %v but got :%v", test.limit, test.want, got)
		}
	}
}

func ExampleIsEven() {
	fmt.Println(IsEven(1))
	fmt.Println(IsEven(2))
	fmt.Println(IsEven(1088347120))
	// Output:
	// false
	// true
	// true
}

// go test -bench=. flag will run only benchmarks or run specific benchmark -bench=BenchmarkName
// -benchmem flag will show memory allocation stats
func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciEvenSum(i)
	}
}

// for coverage use
