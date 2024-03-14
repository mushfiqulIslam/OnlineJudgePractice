package main

func numSubarraysWithSum(nums []int, goal int) int {
	prefixSum, numOfSubarray := 0, 0
	frequencyOfPrefix := map[int]int{0: 1}

	for _, num := range nums {
		prefixSum += num
		if val, ok := frequencyOfPrefix[prefixSum-goal]; ok {
			numOfSubarray += val
		}
		frequencyOfPrefix[prefixSum]++
	}

	return numOfSubarray
}
