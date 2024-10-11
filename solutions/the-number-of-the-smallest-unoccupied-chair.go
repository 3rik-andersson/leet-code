package solutions

import (
	"sort"
)

func SmallestChair(times [][]int, targetFriend int) int {
	events := [][3]int{}
	for i, time := range times {
		events = append(events, [3]int{time[0] * 2, i, 1})
		events = append(events, [3]int{time[1]*2 - 1, i, 0})
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	chairs := []int{}
	targetChair := -1
	for _, event := range events {
		if event[2] == 1 {
			sat := false
			for i, chair := range chairs {
				if chair == -1 {
					chairs[i] = event[1]
					sat = true
					if event[1] == targetFriend {
						targetChair = i
					}
					break
				}
			}
			if !sat {
				chairs = append(chairs, event[1])
				if event[1] == targetFriend {
					targetChair = len(chairs) - 1
				}
			}
		} else {
			for i := range chairs {
				if chairs[i] == event[1] {
					chairs[i] = -1
					break
				}
			}
		}
	}
	return targetChair
}
