package solutions

func LemonadeChange(bills []int) bool {
	c5, c10, c20 := 0, 0, 0
	for _, bill := range bills {
		switch bill {
		case 5:
			c5++
		case 10:
			if c5 == 0 {
				return false
			}
			c5, c10 = c5-1, c10+1
		case 20:
			if c10 > 0 && c5 > 0 {
				c5, c10, c20 = c5-1, c10-1, c20+1
			} else if c5 > 2 {
				c5, c20 = c5-3, c20+1
			} else {
				return false
			}
		}
	}
	return true
}
