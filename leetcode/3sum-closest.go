package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var j, k, jkSum int
	closeSum := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j = i + 1
		k = len(nums) - 1
		for j < k {
			jkSum = nums[i] + nums[j] + nums[k]

			if jkSum == target {
				return jkSum
			} else if abs(closeSum-target) > abs(jkSum-target) {
				closeSum = jkSum
			}
			if jkSum > target {
				k--
			} else {
				j++
			}
		}
	}
	return closeSum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
