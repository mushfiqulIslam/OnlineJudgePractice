import "math"

func countGoodTriplets(arr []int, a int, b int, c int) int {
	count := 0
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			if math.Abs(float64(arr[i]-arr[j])) > float64(a) {
				continue
			}
			for k := j + 1; k < len(arr); k++ {
				if math.Abs(float64(arr[j]-arr[k])) <= float64(b) &&
					math.Abs(float64(arr[i]-arr[k])) <= float64(c) {
					count++
				}
			}
		}
	}

	return count
}
