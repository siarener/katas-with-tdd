/* https://www.codewars.com/kata/534d2f5b5371ecf8d2000a08/train/go

Your task, is to create N×N multiplication table, of size provided in parameter.

For example, when given size is 3:

1 2 3
2 4 6
3 6 9

For the given example, the return value should be:

[[1,2,3],[2,4,6],[3,6,9]] */

package kata

func MultiplicationTable(size int) [][]int {
	matrix := make([][]int, size)

	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			matrix[i-1] = append(matrix[i-1], i*j)
		}
	}

	return matrix
}
