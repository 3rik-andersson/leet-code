package solutions

func CheckStraightLine(coordinates [][]int) bool {
	n := len(coordinates)
	if n < 3 {
		return true
	}
	dx, dy := coordinates[1][0]-coordinates[0][0], coordinates[1][1]-coordinates[0][1]
	for i := 2; i < n; i++ {
		tx, ty := coordinates[i][0]-coordinates[0][0], coordinates[i][1]-coordinates[0][1]
		if tx%dx != 0 || ty%dy != 0 || tx/dx != ty/dy {
			return false
		}
	}
	return true
}
