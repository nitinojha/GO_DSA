// 'Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.
//  
// Example 1:
// Input: nums = [1,2,3]
// Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// Example 2:
// Input: nums = [0,1]
// Output: [[0,1],[1,0]]
// Example 3:
// Input: nums = [1]
// Output: [[1]]
//  
// Constraints:
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// All the integers of nums are unique.'

//[][]int
// fixed number is i and set it to true
// bacxktrack (path, used)
// i as false

package main

import "fmt"

func main() {
	nums := []int{1,2,3}
	possible_permutations :=  combination(nums)
	fmt.Println(possible_permutations)
	//fmt.Println("Possible Permutations is of type %T with value %v",possible_permutations,possible_permutations)

}

func combination(nums []int) [][]int {
	var result [][]int
	var traversal  func(path []int , used []bool)
	traversal =  func(path []int, used []bool) {
		if len(path) == len(nums) {
			temp := make([]int, len(nums))
			//fmt.Println(temp,path)
			copy(temp, path)
			result = append(result, temp)
		
			return
		}



		for i := 0; i < len(nums); i++ {
			if (used[i]) {
				continue
			}

			used[i] = true
			path = append(path, nums[i])
			traversal(path,used)

			path = path[:len(path)-1]
			used[i] = false

		}
	}

	used := make([]bool, len(nums))
	traversal([]int{},used)
	return result

}