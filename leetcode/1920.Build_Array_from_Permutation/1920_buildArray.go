package leetcode

func buildArray(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		nums[i] = nums[i] + n*(nums[nums[i]]%n)
	}
	for i := 0; i < n; i++ {
		nums[i] /= n
	}

	return nums
}

func buildArray2(nums []int) []int {
	ans := []int{}

	for i := 0; i < len(nums); i++ {
		ans = append(ans, nums[nums[i]])
	}

	return ans
}
