package kata

import "testing"

func TestSumOfCubes(t *testing.T) {
	t.Run("only '1's", func(t *testing.T) {
		assertEqualInt(t, 4, sumOfCubes("1111"))
	})
	t.Run("different numbers", func(t *testing.T) {
		assertEqualInt(t, 352, sumOfCubes("172"))
	})
}

func TestReverseStrings(t *testing.T) {
	t.Run("string with even length", func(t *testing.T) {
		assertEqualStrings(t, "654321", reverseString("123456"))
	})
	t.Run("string with odd length", func(t *testing.T) {
		assertEqualStrings(t, "42", reverseString("24"))
	})
	t.Run("only one digit", func(t *testing.T) {
		assertEqualStrings(t, "7", reverseString("7"))
	})
}

func TestRotateStringToLeft(t *testing.T) {
	t.Run("string with several digits", func(t *testing.T) {
		assertEqualStrings(t, "234561", rotateStringToLeft("123456"))
	})
	t.Run("string with one digit", func(t *testing.T) {
		assertEqualStrings(t, "2", rotateStringToLeft("2"))
	})
}

func TestReverseOrRotate(t *testing.T) {
	t.Run("string with rotated chunks of length '6'", func(t *testing.T) {
		assertEqualStrings(t, "234561876549", Revrot("123456987654", 6))
	})
	t.Run("string with reversed chunks of length '6'", func(t *testing.T) {
		assertEqualStrings(t, "234561356789", Revrot("123456987653", 6))
	})
	t.Run("string with chunks of length '4'", func(t *testing.T) {
		assertEqualStrings(t, "44668753", Revrot("66443875", 4))
	})
	t.Run("ignore last chunk if smaller than given chunks size", func(t *testing.T) {
		assertEqualStrings(t, "67834466", Revrot("664438769", 8))
	})
}

func assertEqualInt(t *testing.T, want, got int) {
	t.Helper()
	if got != want {
		t.Errorf("expected %d, but got %d", want, got)
	}
}

func assertEqualStrings(t *testing.T, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("expected %s, but got %s", want, got)
	}
}
