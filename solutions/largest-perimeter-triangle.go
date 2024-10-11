package solutions

import "sort"

func LargestPerimeter(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	for i := n - 1; i > 1; i-- {
		for nums[i] < nums[i-1]+nums[i-2] {
			return nums[i] + nums[i-1] + nums[i-2]
		}
	}
	return 0
}
