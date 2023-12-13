/* https://www.codewars.com/kata/55e7280b40e1c4a06d0000aa/go

John and Mary want to travel between a few towns A, B, C ... Mary has on a sheet of paper a list of distances between these towns. ls = [50, 55, 57, 58, 60]. John is tired of driving and he says to Mary that he doesn't want to drive more than t = 174 miles and he will visit only 3 towns.

Which distances, hence which towns, they will choose so that the sum of the distances is the biggest possible to please Mary and John?
Example:

With list ls and 3 towns to visit they can make a choice between: [50,55,57],[50,55,58],[50,55,60],[50,57,58],[50,57,60],[50,58,60],[55,57,58],[55,57,60],[55,58,60],[57,58,60].

The sums of distances are then: 162, 163, 165, 165, 167, 168, 170, 172, 173, 175.

The biggest possible sum taking a limit of 174 into account is then 173 and the distances of the 3 corresponding towns is [55, 58, 60].

The function chooseBestSum (or choose_best_sum or ... depending on the language) will take as parameters
	t (maximum sum of distances, integer >= 0),
	k (number of towns to visit, k >= 1) and
	ls (list of distances, all distances are positive or zero integers and this list has at least one element).

The function returns the "best" sum ie the biggest possible sum of k distances less than or equal to the given limit t, if that sum exists,
or otherwise nil, null, None, Nothing, depending on the language.

In that case with C, C++, D, Dart, Fortran, F#, Go, Julia, Kotlin, Nim, OCaml, Pascal, Perl, PowerShell, Reason, Rust, Scala, Shell, Swift return -1.

Examples:

ts = [50, 55, 56, 57, 58] choose_best_sum(163, 3, ts) -> 163

xs = [50] choose_best_sum(163, 3, xs) -> nil (or null or ... or -1 (C++, C, D, Rust, Swift, Go, ...)

ys = [91, 74, 73, 85, 73, 81, 87] choose_best_sum(230, 3, ys) -> 228
Notes:

    try not to modify the input list of distances ls
    in some languages this "list" is in fact a string (see the Sample Tests). */

package kata

func sumUpDistances(slices [][]int) []int {
	sumsOfSlices := []int{}
	for _, slice := range slices {
		sum := 0
		for _, num := range slice {
			sum += num
		}
		sumsOfSlices = append(sumsOfSlices, sum)
	}
	return sumsOfSlices
}

func getHighestNumUntilLimit(numbers []int, limit int) int {
	currentMaxNum := -1
	for _, num := range numbers {
		if num > currentMaxNum && num <= limit {
			currentMaxNum = num
		}
	}
	return currentMaxNum
}

func deepCopySlice(orig []int) []int {
	sliceCopy := make([]int, len(orig))
	copy(sliceCopy, orig)
	return sliceCopy
}

func getCombinations(ls []int, numElem int) [][]int {
	var result [][]int
	if numElem == 0 {
		return result
	}

	var generateCombinations func(start int, current []int)
	generateCombinations = func(start int, current []int) {
		if len(current) == numElem {
			result = append(result, deepCopySlice(current))
			return
		}
		for newCombItem := start; newCombItem < len(ls); newCombItem++ {
			// add the new combi item to the current list,
			// and continue to recursively add the remaining possible combi items
			current = append(current, ls[newCombItem])
			generateCombinations(newCombItem+1, current)

			// remove the new combi item again, to continue the for-loop with the
			// next list item instead of the current one
			current = current[:len(current)-1]
		}
	}
	generateCombinations(0, []int{})
	return result
}

func ChooseBestSum(maxDist, numTowns int, ls []int) int {
	routeCombinations := getCombinations(ls, numTowns)
	summedUpDist := sumUpDistances(routeCombinations)
	return getHighestNumUntilLimit(summedUpDist, maxDist)
}
