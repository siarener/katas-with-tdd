package kata

import "testing"

func TestWeightForWeight(t *testing.T) {
	t.Run("order weights with unique digit sums", func(t *testing.T) {
		got := OrderWeight("56 65 74 100 99 68 86 180 90")
		want := "100 180 90 56 65 74 68 86 99"

		if got != want {
			t.Errorf("expected %q but got %q", want, got)
		}
	})
	t.Run("order weights with same digit sums", func(t *testing.T) {
		got := OrderWeight("90 180 900")
		want := "180 90 900"

		if got != want {
			t.Errorf("expected %q but got %q", want, got)
		}
	})
}
