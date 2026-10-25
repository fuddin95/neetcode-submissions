
func hasDuplicate(nums []int) bool {
    duplicate := map[int]int{}

	for _, value := range nums {
		_, exist := duplicate[value]
		if exist{
			return true
		}
		duplicate[value] = 1
	}
	return false
}

