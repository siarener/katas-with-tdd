package kata

import "testing"

func TestFirstProperty(t *testing.T) {
	t.Run("common end of two letters", func(t *testing.T) {
		assertEqualLcs(t,
			processLcsForSameEnding(TmpLcs{"abcdef", "xyzef", ""}),
			TmpLcs{"abcd", "xyz", "ef"})
	})
	t.Run("no common end", func(t *testing.T) {
		assertEqualLcs(t,
			processLcsForSameEnding(TmpLcs{"abcdefg", "xyzef", ""}),
			TmpLcs{"abcdefg", "xyzef", ""})
	})
	t.Run("append to already found subsequence", func(t *testing.T) {
		assertEqualLcs(t,
			processLcsForSameEnding(TmpLcs{"abcdef", "xyzef", "nml"}),
			TmpLcs{"abcd", "xyz", "efnml"})
	})
}

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

func assertEqualLcs(t *testing.T, got, want TmpLcs) {
	t.Helper()
	if got != want {
		t.Errorf("\ngot the common ending\t'%v' with the left differing starts: '%v', '%v'\nexpected the common ending\t'%v' with the left differing starts: '%v', '%v'", got.foundSubsequence, got.s1, got.s2, want.foundSubsequence, want.s1, want.s2)
	}
}
