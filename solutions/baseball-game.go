package solutions

import "strconv"

func CalPoints(operations []string) int {
	nums := []int{}
	for _, op := range operations {
		n := len(nums)
		if op == "+" {
			nums = append(nums, nums[n-1]+nums[n-2])
		} else if op == "D" {
			nums = append(nums, nums[n-1]*2)
		} else if op == "C" {
			nums = nums[:n-1]
		} else {
			num, _ := strconv.Atoi(op)
			nums = append(nums, num)
		}
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
