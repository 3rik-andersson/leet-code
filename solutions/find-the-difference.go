package solutions

func FindTheDifference(s string, t string) byte {
	letters := map[rune]int{}
	for _, c := range t {
		letters[c]++
	}
	for _, c := range s {
		letters[c]--
	}
	for k, c := range letters {
		if c != 0 {
			return byte(k)
		}
	}
	return 'a'
}
