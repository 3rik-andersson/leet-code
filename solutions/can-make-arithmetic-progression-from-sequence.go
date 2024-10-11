package solutions

import "sort"

func CanMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	n := len(arr)
	diff := 0
	for i := 1; i < n; i++ {
		if i == 1 {
			diff = arr[i] - arr[i-1]
		} else if arr[i]-arr[i-1] != diff {
			return false
		}
	}
	return true
}
