package leetcode

func removeDuplicates(nums []int) int {
	k, j := 0, 1
	for j < len(nums) {
		if nums[k] != nums[j] {
			k++
			nums[k] = nums[j]
		} else {
			j++
		}
	}

	return k + 1
}

// 	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

/*
i  j
0, 0, 1, 1, 1, 2, 2, 3, 3, 4

i     j
0, 0, 1, 1, 1, 2, 2, 3, 3, 4

   i  j
0, 1, 1, 1, 1, 2, 2, 3, 3, 4

   i     j
0, 1, 1, 1, 1, 2, 2, 3, 3, 4

   i        j
0, 1, 1, 1, 1, 2, 2, 3, 3, 4

      i        j
0, 1, 1, 1, 1, 2, 2, 3, 3, 4

      i        j
0, 1, 2, 1, 1, 2, 2, 3, 3, 4


*/
