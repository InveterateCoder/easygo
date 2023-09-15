package easygo

func FindIndices[T any](slice []T, filter func(item *T) bool) []int {
	ret := []int{}
	for i := 0; i < len(slice); i++ {
		if filter(&slice[i]) {
			ret = append(ret, i)
		}
	}
	return ret
}
