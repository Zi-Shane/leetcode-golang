package leetcode

// https://zhenchaogan.gitbook.io/leetcode-solution/leetcode-1570-dot-product-of-two-sparse-vectors

type SparseVector struct {
	Vector map[int]int
}

func NewSparseVector(nums []int) *SparseVector {
	vec := SparseVector{Vector: map[int]int{}}
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			vec.Vector[i] = nums[i]
		}
	}
	return &vec
}

func (vec1 *SparseVector) dotProduct(vec2 *SparseVector) int {
	res := 0
	for i, v1 := range vec1.Vector {
		v2, ok := vec2.Vector[i]
		if ok {
			res += v1 * v2
		}
	}
	return res
}
