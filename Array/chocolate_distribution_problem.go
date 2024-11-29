package main

import (
	"fmt"
	"sort"
	"math"
)

//arr = [3, 4, 1, 9, 56, 7, 9, 12], m = 5
//sort the array
// loop till m

//1,3,4,5,7

func main() {
	arr := []int{3, 4, 1, 9, 56, 7, 9, 12}
	//1,3,4,7,9,9,12
	//arr := []int{7, 3, 2, 4, 9, 12, 56}
	m := 5
	sort.Ints(arr)
	fmt.Println(arr)
	n := len(arr)

	// min := arr[0]
	// max := arr[m-1]
	// mindiff := max - min
	

	//Second approach
	mindiff :=  math.MaxInt
	if n > m {
		
		for i := 0; i < n -m ; i++ {
			diff := arr[i+m-1] - arr[i]
			if diff < mindiff {
				mindiff = diff
			}
		}
		
	}
	fmt.Println(mindiff)
	

}
