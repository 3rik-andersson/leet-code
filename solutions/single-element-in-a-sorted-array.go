package solutions

import "fmt"

func SingleNonDuplicate(nums []int) int {
	n := len(nums)
	a, b := 0, n-1
	for a < b {
		c := (a + b) / 2
		d := c
		if c > 0 && c < n-1 && nums[c+1] != nums[c] && nums[c-1] != nums[c] {
			return nums[c]
		}
		if c > 0 && nums[c-1] == nums[c] {
			d = c - 2
		} else if c < n-1 && nums[c+1] == nums[c] {
			d = c - 1
		}
		fmt.Println(nums)
		fmt.Println(a, b, c, d)
		if d%2 == 1 {
			a = c
		} else {
			b = c
		}
	}
	return 0
}
