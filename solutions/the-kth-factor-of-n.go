package solutions

func KthFactor(n int, k int) int {
	nums := []int{}
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			nums = append(nums, i)
		}
	}
	m := len(nums)
	if k <= m {
		return nums[k-1]
	}
	if n/nums[m-1] == nums[m-1] {
		if k >= 2*m {
			return -1
		}
		return n / nums[2*m-k-1]
	}
	if k > 2*m {
		return -1
	}
	return n / nums[2*m-k]
}
