package main

import "fmt"

func main() {
	var matrix [][]int
	// matrix = [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }

	matrix = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	fmt.Printf("Type of Matrix is %T with value is %v \n", matrix, matrix)
	fmt.Println("Original matrix  is :")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}

	// rotated matrix with extra rotated array mattrix

	rotated := rotateMatrix(matrix)

	fmt.Println("\n Rotated matrix with extra matrix space is :\n")
	for i := 0; i < len(rotated); i++ {
		for j := 0; j < len(rotated[0]); j++ {
			fmt.Printf("%d ", rotated[i][j])
		}
		fmt.Println()
	}

	rotateMatrixInPlace(matrix)

	// reversing row elements of a tranposed matrix is it's 90 degree clockwise rotation
	fmt.Println("\n Rotated matrix without extra matrix space is  or Reversal of Tranpose matrix  is :\n")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}

func rotateMatrix(matrix [][]int) [][]int {
	rotatedMatrix := make([][]int, len(matrix[0]))

	for i := range rotatedMatrix {
		rows := make([]int, len(matrix[i]))
		rotatedMatrix[i] = rows
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			rotatedMatrix[j][len(matrix)-i-1] = matrix[i][j]
		}
	}
	return rotatedMatrix
}

// Function to rotate the matrix by 90 degrees clockwise in place
func rotateMatrixInPlace(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	fmt.Println("\n Tranpose matrix  is :\n")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}

	// reversal of transposed matrix is clockwise rotation of matrix by 90 degrees
	for i := 0; i < n; i++ {
		for j, k := 0, n-1; j < k; j, k = j+1, k-1 { // Logic to remember , swap leftmost and rightmost and element till j < k  and increase and decrease k where i used to be hold on each iteration
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}

}
