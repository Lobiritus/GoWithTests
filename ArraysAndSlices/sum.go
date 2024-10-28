package ArraysAndSlices

func SumSlice(slice []int) int {
	var result int
	for _, v := range slice {
		result += v
	}
	return result
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
	var result []int
	for _, numbers := range slice {

		if len(numbers) == 0 {
			result = append(result, 0)
		} else {
			result = append(result, SumSlice(numbers[1:]))
		}
	}
	return result
}
