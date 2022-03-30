package problems

func nextPermutation(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			// find minimun number which is greater than [i-1] in [i:]
			for j := len(nums) - 1; j >= i; j-- {
				if nums[j] > nums[i-1] {
					nums[j], nums[i-1] = nums[i-1], nums[j]
				}
				break
			}
			// reverse
			start := i
			end := len(nums) - 1
			for start < end {
				nums[start], nums[end] = nums[end], nums[start]
				start++
				end--
			}

			return
		}
	}

	// last permutation
	start := 0
	end := len(nums) - 1
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
	return

}
