package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	nums := []int{1, -2, -5, -13, -10, -11, 0, -12, -11, 13, -4, 9, 8, 10, -7, 3, -9, -12, -7, 8, -2, -12, 1, -10, -15, -8, 5, 14, -7, -8, -8, 9, -3, -6, 3, -5, -1, -11, -10, 3, -13, 1, -10, 3, -12, -10, -9, -13, -7, -1, 10, 6, -6, -12, 12, -13, -13, -6, -14, -13, -7, -7, 4, 6, -6, -8, 8, 8, -4, 13, -11, -1, -8, -14, 9, -5, -9, 7, -3, -1, 14, 14, 13, -7, 9, 2, -5, 12, 11, -12, 14, -11, -12, 3, 2, -2, 3, -5, -9, 14, -14, -13, -10, -7, -12, 14, 3, -6, -1, 8, 1, -2, -1, -1, 6, -6}

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
	var zero_temp int
	for i := 0; i < len(nums)-2; i++ {
		l := i + 1
		r := len(nums) - 1
		//fmt.Println(i, l, r)
		for l < r {
			if nums[i] == int(0) && nums[l] == int(0) && nums[r] == int(0) {
				zero_temp = 1
				break
			}
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {

				new_triplet := []int{nums[i], nums[l], nums[r]}
				// response := VerifyUniqueTriplets(result, new_triplet)
				// if response {
				// 	result = append(result, new_triplet)
				// }
				result = append(result, new_triplet)

				l++
				r--
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}

	}
	if zero_temp == 1 {
		result = append(result, []int{0, 0, 0})
	}
	result = findUniqueTriplets(result)
	return result
}

func findUniqueTriplets(arr [][]int) [][]int {
	uniqueMap := make(map[string]struct{})
	var result [][]int

	for _, triplet := range arr {
		// Sort the triplet to handle permutations
		sort.Ints(triplet)

		// Convert triplet to a string key
		key := strings.Join(strings.Fields(fmt.Sprint(triplet)), ",")

		// Check if triplet is already seen
		if _, exists := uniqueMap[key]; !exists {
			uniqueMap[key] = struct{}{}
			result = append(result, triplet)
		}
	}
	return result
}

func VerifyUniqueTriplets(result [][]int, new_triplet []int) bool {
	//uniqueMap := make(map[string]struct{})
	sort.Ints(new_triplet)
	new_triplet_key := strings.Join(strings.Fields(fmt.Sprint(new_triplet)), ",")

	for _, triplet := range result {
		// Sort the triplet to handle permutations
		sort.Ints(triplet)

		// Convert triplet to a string key
		key := strings.Join(strings.Fields(fmt.Sprint(triplet)), ",")
		if key == new_triplet_key {
			return false
		}

		// Check if triplet is already seen
		// if _, exists := uniqueMap[key]; !exists {
		// 	uniqueMap[key] = struct{}{}
		// 	result = append(result, triplet)
		// }
	}
	return true
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
