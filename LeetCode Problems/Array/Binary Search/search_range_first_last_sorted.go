package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange(nums, target))
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	first := findOccurence(nums, target, true)
	last := findOccurence(nums, target, false)

	return []int{first, last}
}

func findOccurence(nums []int, target int, firstOccurence bool) int {
	l := 0
	r := len(nums) - 1
	first := -1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			first = m
			if firstOccurence {
				r = m - 1   // To find first occurence in the left side, we will reduce r by m - 1 whenever target is found
			} else {
				l = m + 1   // To find last occurence in the right side, we will increase l by m + 1 whenever target is found
			}
		}
	}
	return first
}
