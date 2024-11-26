package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 0, 0, 3, 0, 0, 12, 0, 1}
	moveZeroesAnother(nums)
	fmt.Println(nums)
}

func moveZeroesAnother(nums []int) {
	i := 0
	j := 0
	//{0, 0, 0, 3,0, 0, 12, 0, 1}
	for j < len(nums) {
		if nums[j] != 0 {
			temp := nums[j]
			nums[j] = nums[i]
			nums[i] = temp
			i++
		}
		j++
	}
}
