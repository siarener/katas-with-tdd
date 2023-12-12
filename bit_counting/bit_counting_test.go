package kata

import "testing"

func TestCountBits(t *testing.T) {
	t.Run("small number", func(t *testing.T) {
		assertExpectedBitCount(t, 3, CountBits(7))
	})
	t.Run("large number", func(t *testing.T) {
		assertExpectedBitCount(t, 5, CountBits(1234))
	})
	t.Run("number is '0'", func(t *testing.T) {
		assertExpectedBitCount(t, 0, CountBits(0))
	})
}

func assertExpectedBitCount(t *testing.T, want int, got int) {
	t.Helper()
	if got != want {
		t.Errorf("expected %d, but got %d", want, got)
	}
}
