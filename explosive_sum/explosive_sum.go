/* https://www.codewars.com/kata/52ec24228a515e620b0005ef/train/go

How many ways can you make the sum of a number?

From wikipedia: https://en.wikipedia.org/wiki/Partition_(number_theory)

    In number theory and combinatorics, a partition of a positive integer n, also called an integer partition, is a way of writing n
	as a sum of positive integers. Two sums that differ only in the order of their summands are considered the same partition.
	If order matters, the sum becomes a composition. For example, 4 can be partitioned in five distinct ways:

4
3 + 1
2 + 2
2 + 1 + 1
1 + 1 + 1 + 1

Examples
Basic

ExpSum(1) // 1
ExpSum(2) // 2 -> 1+1 , 2
ExpSum(3) // 3 -> 1+1+1, 1+2, 3
ExpSum(4) // 5 -> 1+1+1+1, 1+1+2, 1+3, 2+2, 4
ExpSum(5) // 7 -> 1+1+1+1+1, 1+1+1+2, 1+1+3, 1+2+2, 1+4, 5, 2+3

ExpSum(10) // 42

Explosive

ExpSum(50) // 204226
ExpSum(80) // 15796476
ExpSum(100) // 190569292

See here for more examples. */

package kata

func ExpSum(u uint64) uint64 {

	sumOptions := make([]uint64, u+1)
	sumOptions[0] = 1

	/* My former solution is replaced by the [solution from Noob's blog](https://yurukute.github.io/Blog/en/post/explosivesum/),
	    as it couldn't be beat in it's simplicity. I rewrote the
		explanation in the following with more in-between steps,
		to grasp better what happens in the program.

		For the calculation for the options we build a table, until
		we get to the row and cell where our result is located. As
		we only need the last row to compute the following row, we
		don't need to save the rows before the last row.

		Example: Finding the options to find the partitions of '4'

			Starting row:
						  | 0  1  2  3  4
						------------------
						0 | 1  0  0  0  0

			1. row calculation:
					The values from the starting row are used to calculate the new row.

					'sumOptions[cellIdx] += sumOptions[cellIdx-rowIdx]'
					->
					row[1] = row[0] + row[1]
							where
							-> row[0] holds the value from the starting row, i.e. '1'
							-> row[1] holds the value from the starting row, i.e. '0'
						-> cell row[1] of the currently calculated row now holds the value '1'
					...

						  | 0  1  2  3  4
						------------------
						1 | /  1  1  1  1


			2. row calculation:

					'sumOptions[cellIdx] += sumOptions[cellIdx-rowIdx]'
					->
					row[2] = row[2] + row[2-2]
							where
							-> row[2] holds the value from the former row, i.e. '1'
							-> row[0] holds the value from the starting row, i.e. '1'
						-> cell row[2] of the currently calculated row now holds the value '2'
					...
					row[4] = row[4] + row[4-2]
							where
							-> row[4] holds the value from the former row, i.e. '1'
							-> row[2] holds the value from the current row, i.e. '2'
						-> cell row[4] of the currently calculated row now holds the value '3'
					...
						  | 0  1  2  3  4
						------------------
						2 | /  /  2  2  3

			3. row calculation:
						  | 0  1  2  3  4
						------------------
						3 | /  /  /  3  4

			4. row calculation:
						  | 0  1  2  3  4
						------------------
						4 | /  /  /  /  5


			If we would add what we calculate in each loop, we have:
						  | 0  1  2  3  4
						------------------
						0 | 1  0  0  0  0
						1 | /  1  1  1  1
						2 | /  /  2  2  3
						3 | /  /  /  3  4
						4 | /  /  /  /  5

			We can skip calculating one more cell value each loop while progressing each loop,
			as we use the value still stored from the last row for the calculation.
			The '/' shows that the respective value that was not calculated. As we
			are always using the same row, this means that this cell still holds the last
			value assigned to it, i.e. in the last loop, the value in row[0] is still from
			the starting row, the value in row[1] is from the 1st loop, the value in row[2]
			from the 2nd loop, etc., so our final row looks like this:
						  | 0  1  2  3  4
						------------------
						4 | 1  1  2  3  5
	*/
	for rowIdx := uint64(1); rowIdx <= u; rowIdx++ {
		for cellIdx := rowIdx; cellIdx <= u; cellIdx++ {
			sumOptions[cellIdx] += sumOptions[cellIdx-rowIdx]
		}
	}

	return sumOptions[u]
}
