package leetcode

func runningSum2(nums []int) []int {
	res := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		res = append(res, res[i-1]+nums[i])
	}

	return res
}

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	return nums
}
