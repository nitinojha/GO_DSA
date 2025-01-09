package main

import (
	"fmt"
	//"strings"
)

func main() {
	fmt.Println(removeDuplicates("abbaca"))
	fmt.Println(removeadjacentDuplicates("abbaca"))
	fmt.Println(removeadjacentDuplicatesWithStack("abbaca"))
}

func removeDuplicates(S string) string {
	var stack []rune
	for _, c := range S {
		if len(stack) > 0 && stack[len(stack)-1] == c {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}
	return string(stack)
}

func removeadjacentDuplicates(s string) string { // Without using stack(extra space for other datatype)
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			s = s[:i-1] + s[i+1:]
			i = 0
		}
	}
	return s
}

func removeadjacentDuplicatesWithStack(s string) string {
	stack := []rune{}
	for _, c := range s {
		if len(stack) > 0 && stack[len(stack)-1] == c {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}

	// stack1 :=  []byte{}
	// stack1 = append(stack1,s[0])
	// for i := 1 ; i < len(s); i++ {
	//     stack_top := len(stack1) - 1
	//     if len(stack1) == 0 ||  s[i] != stack1[stack_top] {
	//         stack1 = append(stack1,s[i])
	//     } else {
	//         stack1 = stack1[:len(stack1) - 1]
	//     }
	// }

	return string(stack)
	//return string(stack1)
}
