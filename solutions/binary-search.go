package solutions

func Search(nums []int, target int) int {
	a, b := 0, len(nums)-1
	for a < b {
		c := (a + b) / 2
		if nums[c] > target {
			if b == c {
				return -1
			}
			b = c
		} else if nums[c] < target {
			if a == c {
				return -1
			}
			a = c
		} else {
			return c
		}
	}
	return -1
}
