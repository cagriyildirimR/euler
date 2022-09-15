package exponentiation

import "testing"

func BenchmarkNaiveExponentiation(b *testing.B) {
	for i := 1; i < b.N; i++ {
		NaiveExponentiation(2, 100)
	}
}

func BenchmarkBinaryExponentiation(b *testing.B) {
	for i := 1; i < b.N; i++ {
		BinaryExponentiation(2, 100)
	}
}

func BenchmarkMinimumExponentiation(b *testing.B) {
	for i := 1; i < b.N; i++ {
		MinimumExponentiation(2, 100)
	}
}

func TestMinimumExponentiation(t *testing.T) {
	for _, test := range []struct {
		base, power int
		want        int
	}{
		{2, 30, 1073741824},
		{3, 20, 3486784401},
		{2, 17, 131072},
	} {
		if got := MinimumExponentiation(test.base, test.power); test.want != got {
			t.Errorf("%v**%v is %v but got: %v", test.base, test.power, test.want, got)
		}
	}
}
