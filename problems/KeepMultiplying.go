package problems

func FindFinalValue(nums []int, original int) int {
	numMap := make(map[int]bool)

	for _, num := range nums {
		numMap[num] = true
	}

	isFound := numMap[original]
	currentValue := original

	for isFound {
		currentValue = currentValue * 2
		isFound = numMap[currentValue]
	}

	return currentValue
}
