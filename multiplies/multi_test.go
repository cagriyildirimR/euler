package multiplies

import "testing"

func TestMultipliesOfXAndY(t *testing.T) {

	for _, test := range []struct {
		target, x, y int
		want         int
	}{
		{0, 1, 2, 0},
		{10, 3, 5, 23},
		{1000, 3, 5, 233168},
	} {
		if got := MultipliesOfXAndY(test.target, test.x, test.y); got != test.want {
			t.Errorf("Wrong result: want: '%v' but function MultipliesOfXAndY(%v, %v, %v) returns %v", test.want, test.target, test.x, test.y, got)
		}
	}
}

func TestSumList(t *testing.T) {
	for _, test := range []struct {
		list []int
		want int
	}{
		{[]int{}, 0}, // Empty list
		{[]int{1, 2, 3, 4, 5, 6}, 21},
	} {
		if got := sumList(test.list); got != test.want {
			t.Errorf("List values are not add up to %v", test.want)
		}
	}
}
