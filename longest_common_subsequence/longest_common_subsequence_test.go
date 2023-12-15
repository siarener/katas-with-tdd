package kata

import "testing"

func TestLcs(t *testing.T) {
	t.Run("common start", func(t *testing.T) {
		assertEqualString(t,
			LCS("abcdef", "abc"),
			"abc")
	})
	t.Run("subsequence consists of non-consecutive letters", func(t *testing.T) {
		assertEqualString(t,
			LCS("abcdef", "acf"),
			"acf")
	})

	t.Run("string consists of digits", func(t *testing.T) {
		assertEqualString(t,
			LCS("132535365", "123456789"),
			"12356")
	})

}

func assertEqualString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' but expected '%s'", got, want)
	}
}
