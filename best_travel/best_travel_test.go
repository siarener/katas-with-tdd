package kata

import (
	"reflect"
	"testing"
)

/* func TestChooseBestSum(t *testing.T) {
	got := ChooseBestSum(163, 3, []int{50, 55, 56, 57, 58})
	want := 163
	if got != want {
		t.Errorf("got %d but expected %d", got, want)
	}
} */

func TestSumUpSlices(t *testing.T) {
	t.Run("slices with numbers", func(t *testing.T) {
		assertExpectedSum(
			t,
			sumUpDistances([][]int{
				{50, 55, 57},
				{50, 55, 58},
				{50, 55, 60},
				{50, 57, 58},
				{50, 57, 60},
				{50, 58, 60},
				{55, 57, 58},
				{55, 57, 60},
				{55, 58, 60},
				{57, 58, 60},
			}),
			[]int{162, 163, 165, 165, 167, 168, 170, 172, 173, 175},
		)
	})
	t.Run("empty input", func(t *testing.T) {
		assertExpectedSum(
			t,
			sumUpDistances([][]int{}),
			[]int{},
		)
	})
}

func TestGetCombinations(t *testing.T) {
	t.Run("list of four numbers, combinations of two numbers", func(t *testing.T) {
		assertExpectedSlices(t,
			getCombinations([]int{1, 2, 3, 4}, 2),
			[][]int{
				{1, 2},
				{1, 3},
				{1, 4},
				{2, 3},
				{2, 4},
				{3, 4},
			})
	})
	t.Run("list of four numbers, combinations of zero numbers", func(t *testing.T) {
		assertExpectedSlices(t,
			getCombinations([]int{1, 2, 3, 4}, 0),
			[][]int{})
	})
}

func bothSlicesAreEmpty(a, b [][]int) bool {
	return len(a) == 0 && len(b) == 0
}

func assertExpectedSlices(t *testing.T, got, want [][]int) {
	t.Helper()

	/* 	comparing empty and non-empty slice of slices are handled differently:
	- non-empty slices can be compared with `reflect.DeepEqual` as this function recursively compares the elements of the slices
	  -> this works with slices with simple types
	- empty slices cannot be compared with `reflect.DeepEqual` as the underlying arrays are different; hence we check their equality by
	  asserting that both are of length `0` */

	if !reflect.DeepEqual(got, want) && !bothSlicesAreEmpty(got, want) {
		t.Errorf("got %v but expected %v", got, want)
	}
}

func assertExpectedSum(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but expected %v", got, want)
	}
}
