package fibonacci

func fibonacciEvenSumTailRec(limit, first, second, sum int) (int, int, int, int) {
	if limit < second {
		return limit, first, second, sum
	}
	first, second = second, first+second
	if isEven(second) {
		sum += second
	}
	return fibonacciEvenSumTailRec(limit, first, second, sum)
}

func isEven(v int) bool {
	return v%2 == 0
}

func FibonacciEvenSum(limit int) (sum int) {
	_, _, _, sum = fibonacciEvenSumTailRec(limit, 0, 1, 0)
	return sum
}
