package solutions

func JudgeCircle(moves string) bool {
	u, d, l, r := 0, 0, 0, 0
	for _, move := range moves {
		switch move {
		case 'U':
			u++
		case 'D':
			d++
		case 'L':
			l++
		case 'R':
			r++
		}
	}
	if u == d && l == r {
		return true
	}
	return false
}
