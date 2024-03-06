package main

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	sort.Ints(nums)
	var solutions [][]int
	var j, k, jkSum, target int

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j = i + 1
		k = len(nums) - 1
		for j < k {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if k < len(nums)-1 && nums[k] == nums[k+1] {
				k--
				continue
			}
			target = 0 - nums[i]
			jkSum = nums[j] + nums[k]

			if jkSum == target {
				solutions = append(solutions, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			} else if jkSum > target {
				k--
			} else {
				j++
			}
		}
	}
	return solutions
}
