package solutions

func PartitionString(s string) int {
	cnt := 0
	flag := make(map[rune]bool)
	for _, c := range s {
		if flag[c] {
			cnt++
			flag = make(map[rune]bool)
		}
		flag[c] = true
	}
	return cnt + 1
}
