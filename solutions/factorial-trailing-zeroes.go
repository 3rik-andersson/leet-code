package solutions

func TrailingZeroes(n int) int {
	cnt := 0
	for n > 0 {
		cnt += n / 5
		n /= 5
	}
	return cnt
}
