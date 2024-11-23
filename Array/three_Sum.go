package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, -1, -2, 1, 1, 0, 2}

	result := threeSum(nums)
	fmt.Println(result)
	result1 := threeSumAlternate(nums)
	fmt.Println(result1)
}

func threeSum(nums []int) [][]int {

	//to be fixed
	var result [][]int
	//sum := 0
	sort.Ints(nums)

	j := 1
	k := len(nums) - 1
	for i := 0; i < k && j < k; {
		if nums[i]+nums[j]+nums[k] == 0 {
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
		} else if nums[i]+nums[j]+nums[k] < 0 {
			i++
			j++
		} else {
			k--
		}
	}
	return result
}

// My previous logic submission with O(n3)
func threeSumAlternate(nums []int) [][]int {
	var arr [][]int
	var sum, zero_temp int
	size := len(nums)
	if size < 3 {
		return arr
	}
	for k, _ := range nums {
		for k1 := k + 1; k1 < size; k1 += 1 {
			for k2 := k1 + 1; k2 < size; k2 += 1 {
				//fmt.Printf("%d,%d,%d", k, k1, k2)
				//fmt.Println("\n")
				//fmt.Printf("%d,%d,%d", nums[k], nums[k1], nums[k2])
				//fmt.Println("\n")
				if (nums[k] + nums[k1] + nums[k2]) == 0 {
					// fmt.Println("\nsum is 0 for ", k, k1, k2)
					// fmt.Printf("\n%d,%d,%d", k, k1, k2)
					// fmt.Printf("\n%d,%d,%d", nums[k], nums[k1], nums[k2])
					if len(arr) > 0 {
						for _, row := range arr {
							sum = 0
							temp_arr := []int{nums[k], nums[k1], nums[k2]}
							for _, val := range row {
								// if row[j] == int(0) && row[j+1] == int(0) && row[j+2] == int(0) {
								// 	//fmt.Println("\n row all values equals to zero")
								// 	break
								// }
								for _, values := range temp_arr {
									if values == val {
										sum += 1
										break
									}
								}
							}
							if sum == 3 {
								// fmt.Println("one matched \n")
								// fmt.Println("\nsum:", sum)
								// fmt.Println("\n row:", row)
								// fmt.Println("\ntemp_arr:", temp_arr)
								break
							}
						}
					}
					if sum == 3 {
						// fmt.Println("repated \n")
						continue
					}
					if nums[k] == int(0) && nums[k1] == int(0) && nums[k2] == int(0) {
						//fmt.Println("\n test all are values 0")
						zero_temp = 1
						continue
					}
					arr = append(arr, []int{nums[k], nums[k1], nums[k2]})
					//temp_arr = append(temp_arr, []int{nums[k], nums[k1], nums[k2]})
				}
			}

		}
	}
	if zero_temp == 1 {
		arr = append(arr, []int{0, 0, 0})
	}
	return arr
}
