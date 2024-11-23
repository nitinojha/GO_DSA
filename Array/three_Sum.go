package main

import(
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1,0,-1,-2,1,1,0,2}
	
	result := threeSum(nums)
	fmt.Println(result)
}

func threeSum(nums []int) [][]int {

	//to be fixed
    var result [][]int
	//sum := 0
    sort.Ints(nums)


    j:= 1; k := len(nums) - 1
    for i:=0; (i<k && j<k); {
        if nums[i] + nums[j] + nums[k] == 0  {
            flag := true
            for l, _ := range result {
                if (result[l][0] == nums[i]) && (result[l][1] == nums[j]) && (result[l][2] == nums[k]) {
                    flag = false
                    break
                }
            }
            if flag {
                result = append(result, []int{nums[i], nums[j], nums[k]})
            }
            i++
            j++
        } else if nums[i] + nums[j] + nums[k] < 0 {
            i++
            j++
        } else {
            k--
        }
	}
	return result
}