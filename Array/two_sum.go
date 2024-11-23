package main

import(
	"fmt"
)

func main() {
	nums := []int{3,2,3}
	sum := 6
	result := twoSum(nums,sum)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
    arr := []int{}
    keymap  := make(map[int]int)
    for  i := 0; i < len(nums) ; i++ {
		
        if value,ok := keymap[target - nums[i]]; ok {
            arr = append(arr, value, i)
			return arr
        } else {
            keymap[nums[i]] = i
        }
    }
	
    return arr

}