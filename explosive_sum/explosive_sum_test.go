package kata

import "testing"

func TestExpSum(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		assertExpectedCount(
			t,
			ExpSum(1),
			1,
		)
	})
	t.Run("2", func(t *testing.T) {
		assertExpectedCount(
			t,
			ExpSum(2),
			2,
		)
	})
	t.Run("3", func(t *testing.T) {
		assertExpectedCount(
			t,
			ExpSum(3),
			3,
		)
	})
	t.Run("4", func(t *testing.T) {
		assertExpectedCount(
			t,
			ExpSum(4),
			5,
		)
	})
	t.Run("5", func(t *testing.T) {
		assertExpectedCount(
			t,
			ExpSum(5),
			7,
		)
	})
	t.Run("10", func(t *testing.T) {
		assertExpectedCount(
			t,
			ExpSum(10),
			42,
		)
	})

	t.Run("0", func(t *testing.T) {
		assertExpectedCount(
			t,
			ExpSum(0),
			1,
		)
	})

}

func assertExpectedCount(t *testing.T, got, want uint64) {
	t.Helper()
	if got != want {
		t.Errorf("got %d but expected %d", got, want)
	}
}
