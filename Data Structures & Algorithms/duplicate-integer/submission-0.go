func hasDuplicate(nums []int) bool {
	var seenNumbers map[int]bool = make(map[int]bool)

	for _, val := range nums {
		if seenNumbers[val] {
			return seenNumbers[val]
		} 

		seenNumbers[val] = true
	}

	return false
}
