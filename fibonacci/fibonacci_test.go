package fibonacci

import "testing"

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
