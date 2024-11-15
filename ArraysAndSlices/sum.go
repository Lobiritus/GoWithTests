package ArraysAndSlices

func SumSlice(slice []int) int {
	add := func(acc, x int) int { return acc + x }

	return Reduce(slice, add, 0)
}

func SumArray(arr [5]int) int {
	var result int
	for _, v := range arr {
		result += v
	}
	return result
}

func SumAll(slice ...[]int) []int {
	var result []int
	for _, numbers := range slice {
		result = append(result, SumSlice(numbers))
	}
	return result
}

func SumAllTails(slice ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, SumSlice(tail))
		}
	}
	return Reduce(slice, sumTail, []int{})
}

func Reduce[A any](collection []A, f func(A, A) A, initialValue A) A {
	var result = initialValue
	for _, x := range collection {
		result = f(result, x)
	}
	return result
}
