package kata

import "testing"

func TestIsSolved(t *testing.T) {
	t.Run("not yet finished", func(t *testing.T) {
		assertCorrectCheck(t,
			InProgress,
			IsSolved([3][3]int{
				{p_, p_, pX},
				{p_, pX, pO},
				{pO, pX, p_},
			}))
	})

	t.Run("winning row", func(t *testing.T) {
		assertCorrectCheck(t,
			WinO,
			IsSolved([3][3]int{
				{p_, pO, pX},
				{pO, pO, pO},
				{p_, p_, p_},
			}))
	})
	t.Run("winning column", func(t *testing.T) {
		assertCorrectCheck(t,
			WinO,
			IsSolved([3][3]int{
				{p_, pO, pX},
				{pO, pO, pX},
				{p_, pO, p_},
			}))
	})
	t.Run("winning diagonal O", func(t *testing.T) {
		assertCorrectCheck(t,
			WinO,
			IsSolved([3][3]int{
				{pO, pO, pX},
				{pO, pO, pX},
				{p_, pX, pO},
			}))
	})
	t.Run("winning diagonal X", func(t *testing.T) {
		assertCorrectCheck(t,
			WinX,
			IsSolved([3][3]int{
				{pO, pO, pX},
				{pO, pX, pX},
				{pX, p_, pO},
			}))
	})
	t.Run("draw", func(t *testing.T) {
		assertCorrectCheck(t,
			Draw,
			IsSolved([3][3]int{
				{pX, pO, pX},
				{pO, pO, pX},
				{pX, pX, pO},
			}))
	})
}

func TestCheckWinning(t *testing.T) {
	t.Run("not yet finished", func(t *testing.T) {
		assertCorrectCheck(t,
			InProgress,
			checkWinning([3]int{p_, p_, pX}))
	})
	t.Run("not yet finished", func(t *testing.T) {
		assertCorrectCheck(t,
			InProgress,
			checkWinning([3]int{pO, pO, p_}))
	})
	t.Run("not yet finished", func(t *testing.T) {
		assertCorrectCheck(t,
			InProgress,
			checkWinning([3]int{pO, pX, pO}))
	})
	t.Run("win for X", func(t *testing.T) {
		assertCorrectCheck(t,
			WinX,
			checkWinning([3]int{pX, pX, pX}))
	})
	t.Run("win for O", func(t *testing.T) {
		assertCorrectCheck(t,
			WinO,
			checkWinning([3]int{pO, pO, pO}))
	})
}

func assertCorrectCheck(t *testing.T, want int, got int) {
	t.Helper()
	if got != want {
		t.Errorf("expected %d, but got %d", want, got)
	}
}
