package easygo

import "testing"

func TestFindIndex(t *testing.T) {
	param := []int{1, 5, 10, 15, 20}
	want := 2
	got := FindIndex(param, func(item *int) bool {
		return *item == 10
	})
	if got != want {
		t.Errorf("got %d, wanted: %d\n", got, want)
	}
	want = -1
	got = FindIndex(param, func(item *int) bool {
		return *item == 11
	})
	if got != want {
		t.Errorf("got %d, wanted: %d\n", got, want)
	}
}
