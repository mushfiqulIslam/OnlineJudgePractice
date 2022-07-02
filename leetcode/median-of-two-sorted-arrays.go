package main

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	middleFloat := float64(len(nums1)) / 2.00
	middleInt := int(middleFloat)
	if middleFloat > float64(middleInt) {
		return float64(nums1[len(nums1)/2])
	}
	return float64(nums1[middleInt-1]+nums1[middleInt]) / 2.00
}
