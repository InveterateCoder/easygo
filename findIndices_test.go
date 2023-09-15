package easygo

import "testing"

func TestFindIndices(t *testing.T) {
	param := []int{1, 2, 3, 4, 5, 6}
	want := []int{1, 3, 5}
	lambda := func(item *int) bool {
		return *item%2 == 0
	}
	got := FindIndices(param, lambda)
	if len(got) != len(want) {
		t.Errorf("length got: %d, wanted: %d\n", len(got), len(want))
	}
	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("got %q, wanted: %q\n", got, want)
		}
	}
}
