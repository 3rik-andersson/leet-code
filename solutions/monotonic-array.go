package solutions

func IsMonotonic(nums []int) bool {
	n := len(nums)
	sign := 0
	for i := 1; i < n; i++ {
		if sign == 0 {
			if nums[i] > nums[i-1] {
				sign = 1
			} else if nums[i] < nums[i-1] {
				sign = -1
			}
		} else {
			if nums[i] > nums[i-1] && sign == -1 {
				return false
			} else if nums[i] < nums[i-1] && sign == 1 {
				return false
			}
		}
	}
	return true
}
