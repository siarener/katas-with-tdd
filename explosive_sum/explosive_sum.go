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
	if u == 0 {
		return uint64(1)
	}

	dp := make([][]uint64, u+1)
	for i := range dp {
		dp[i] = make([]uint64, u+1)
	}

	for i := 0; i <= int(u); i++ {
		dp[i][1] = 1
		dp[0][i] = 1

	}

	for i := 1; i <= int(u); i++ {
		dp[i][1] = 1
		for j := 2; j <= int(u); j++ {
			dp[i][j] = dp[i][j-1]
			if i >= j {
				dp[i][j] += dp[i-j][j]
			}
		}
	}
	return uint64(dp[u][u])
}
