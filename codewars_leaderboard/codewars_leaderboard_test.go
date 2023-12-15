package kata

import (
	"testing"
)

func TestGetLeaderboardHonor(t *testing.T) {
	got := GetLeaderboardHonor()

	gotLen := len(got)
	expectedLen := 500
	if expectedLen != gotLen {
		t.Errorf("got list with %v items, expected list with %d items", expectedLen, gotLen)
	}

	gotHonorOfWinner := got[0].Honor
	expectedHonorExceeds := 400000
	if expectedHonorExceeds > gotHonorOfWinner {
		t.Errorf("got %d honor for user at first position, expected at least %d", gotHonorOfWinner, expectedHonorExceeds)
	}

}
