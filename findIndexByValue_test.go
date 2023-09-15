package easygo

import "testing"

func TestFindIndexByValue(t *testing.T) {
	param := []int{1, 2, 3, 4, 5, 6, 7}
	want := 3
	got := FindIndexByValue(param, 4)
	if got != want {
		t.Errorf("got %d, wanted: %d\n", got, want)
	}
}
