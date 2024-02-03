package main

func maxSumAfterPartitioning(arr []int, k int) int {
	lenArr := len(arr)
	maxArr := make([]int, lenArr+1)
	var maximum int

	for i := 1; i < lenArr+1; i++ {
		maximum = 0
		for j := 1; j < min(i, k)+1; j++ {
			maximum = max(maximum, arr[i-j])
			maxArr[i] = max(maxArr[i], maxArr[i-j]+maximum*j)
		}

	}

	return maxArr[lenArr]
}
