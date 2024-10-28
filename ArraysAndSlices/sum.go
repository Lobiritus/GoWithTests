package ArraysAndSlices

func Sum(arr []int) int {
	var result int
	for _, v := range arr {
		result += v
	}
	return result
}
