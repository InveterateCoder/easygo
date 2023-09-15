package easygo

import (
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	param := []int{4, 7, 9}
	want := []string{"8", "14", "18"}
	lambda := func(item *int) string {
		return strconv.Itoa(*item * 2)
	}
	got := Map(param, lambda)
	if len(got) != len(want) {
		t.Errorf("length got: %d, wanted: %d\n", len(got), len(want))
	}
	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("got %q, wanted: %q\n", got, want)
		}
	}
}
