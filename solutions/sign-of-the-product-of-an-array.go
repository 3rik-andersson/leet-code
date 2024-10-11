package solutions

func ArraySign(nums []int) int {
	sign := 1
	for _, num := range nums {
		if num == 0 {
			return 0
		} else if num < 0 {
			sign = -sign
		}
	}
	return sign
}
