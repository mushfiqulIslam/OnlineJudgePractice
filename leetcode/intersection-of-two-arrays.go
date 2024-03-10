package main

func intersection(nums1 []int, nums2 []int) []int {
	numsNotVisited := make(map[int]bool)
	var intersections []int

	for _, num := range nums1 {
		if _, exists := numsNotVisited[num]; exists {
			continue
		}
		numsNotVisited[num] = true
	}

	for _, num := range nums2 {
		if numsNotVisited[num] {
			intersections = append(intersections, num)
			numsNotVisited[num] = false
		}
	}

	return intersections
}
