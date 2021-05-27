package cn

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	num_loc := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		v, ok := num_loc[target-nums[i]]
		if ok {
			return []int{v, i}
		} else {
			num_loc[nums[i]] = i
		}
	}
	return []int{-1, -1}
}

//leetcode submit region end(Prohibit modification and deletion)
