package leetcode

func maxSubArray(nums []int) int {
	cur, res := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		cur = cur + nums[i]
		if cur < 0 || cur < nums[i] {
			cur = nums[i]
		}
		if cur > res {
			res = cur
		}
	}

	return res
}

// O(nlogn)
func maxSubArray2(nums []int) int {
	return maxSubArrayHelperFunc(nums, 0, len(nums)-1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArrayHelperFunc(nums []int, left int, right int) int {
	if left == right {
		return nums[left]
	}
	mid := left + (right-left)/2
	// mid := (left + right) / 2
	leftRet := maxSubArrayHelperFunc(nums, left, mid)
	rightRet := maxSubArrayHelperFunc(nums, mid+1, right)
	// cross mid
	tmp := 0
	leftMax := nums[mid]
	for i := mid; i >= left; i-- {
		tmp += nums[i]
		if tmp > leftMax {
			leftMax = tmp
		}
	}
	tmp = 0
	rightMax := nums[mid+1]
	for i := mid + 1; i <= right; i++ {
		tmp += nums[i]
		if tmp > rightMax {
			rightMax = tmp
		}
	}
	return max(max(leftRet, rightRet), leftMax+rightMax)
}

// O(n^2)
func maxSubArray3(nums []int) int {
	max := nums[0]

	for i := 0; i < len(nums); i++ {
		sumNums := map[int]int{i: nums[i]}
		for j := i; j < len(nums); j++ {
			sumNums[j] = sumNums[j-1] + nums[j]
			if sumNums[j] > max {
				max = sumNums[j]
			}
		}
	}

	return max
}

func maxSubArray4(nums []int) int {
	maxSum := 0
	// size = # of elements need to sum
	for size := 1; size <= len(nums); size++ {
		// start = start index of sum
		for start := 0; start < len(nums)-(size-1); start++ {
			s := 0
			// Do sum
			for i := 0; i < size; i++ {
				s += nums[start+i]
			}
			if s > maxSum {
				maxSum = s
			}
		}
	}
	return maxSum
}
