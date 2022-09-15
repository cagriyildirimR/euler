package multiplies

func MultipliesOfXAndY(target, x, y int) int {

	list := []int{}

	for i := 1; i < target; i++ {
		if i%3 == 0 || i%5 == 0 {
			list = append(list, i)
		}
	}
	return sumList(list)
}

func sumList(list []int) (result int) {
	for _, v := range list {
		result += v
	}
	return result
}
