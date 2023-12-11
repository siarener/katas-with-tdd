package kata

import "testing"

func TestReverseOrRotate(t *testing.T) {
	t.Run("first example", func(t *testing.T) {
		got := Revrot("123456987654", 6)
		want := "234561876549"

		if got != want {
			t.Errorf("expected %s, but got %s", want, got)
		}
	})
}
