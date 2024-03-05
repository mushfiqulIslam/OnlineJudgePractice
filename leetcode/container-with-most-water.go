func maxArea(height []int) int {
	left, area := 0, 0
	right := len(height) - 1
	for left < right {
		area = max(area, min(height[left], height[right]) * (right - left))
		if height[left] > height[right] {
			right -= 1
		} else {
			left += 1
		}
	}
	return area
}
