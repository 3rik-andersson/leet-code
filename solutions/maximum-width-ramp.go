package solutions

func MaxWidthRamp(nums []int) int {
	n := len(nums)
	for i := n - 1; i > 0; i-- {
		for j := 0; j < n-i; j++ {
			for nums[i-j] >= nums[j] {
				return i
			}
		}
	}
	return 0
}
