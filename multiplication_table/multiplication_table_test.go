package kata

import (
	"reflect"
	"testing"
)

func TestMultiplicationTable(t *testing.T) {
	t.Run("matrix table of length '3'", func(t *testing.T) {
		assertExpectedTable(t,
			MultiplicationTable(3),
			[][]int{
				{1, 2, 3},
				{2, 4, 6},
				{3, 6, 9},
			})
	})
	t.Run("matrix table of length '0'", func(t *testing.T) {
		assertExpectedTable(t,
			MultiplicationTable(0),
			[][]int{})
	})
	t.Run("matrix table of length '1'", func(t *testing.T) {
		assertExpectedTable(t,
			MultiplicationTable(1),
			[][]int{
				{1},
			})
	})
	t.Run("matrix table of length '1'", func(t *testing.T) {
		assertExpectedTable(t,
			MultiplicationTable(10),
			[][]int{
				{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				{2, 4, 6, 8, 10, 12, 14, 16, 18, 20},
				{3, 6, 9, 12, 15, 18, 21, 24, 27, 30},
				{4, 8, 12, 16, 20, 24, 28, 32, 36, 40},
				{5, 10, 15, 20, 25, 30, 35, 40, 45, 50},
				{6, 12, 18, 24, 30, 36, 42, 48, 54, 60},
				{7, 14, 21, 28, 35, 42, 49, 56, 63, 70},
				{8, 16, 24, 32, 40, 48, 56, 64, 72, 80},
				{9, 18, 27, 36, 45, 54, 63, 72, 81, 90},
				{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			})
	})
}

func assertExpectedTable(t *testing.T, got, want [][]int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but expected %v", got, want)
	}
}
