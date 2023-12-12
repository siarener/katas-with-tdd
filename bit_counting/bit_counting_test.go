package kata

import "testing"

func TestCountBits(t *testing.T) {
	t.Run("small number", func(t *testing.T) {
		got, err := CountBits(7)
		assertExpectedBitCount(t, 3, got, err)
	})
	t.Run("large number", func(t *testing.T) {
		got, err := CountBits(1234)
		assertExpectedBitCount(t, 5, got, err)
	})
	t.Run("number is '0'", func(t *testing.T) {
		got, err := CountBits(0)
		assertExpectedBitCount(t, 0, got, err)
	})
}

func assertExpectedBitCount(t *testing.T, want int, got int, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("expected %d, but got error: '%v'", want, err)
	}
	if got != want {
		t.Errorf("expected %d, but got %d", want, got)
	}
}
