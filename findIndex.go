package easygo

func FindIndex[T any](slice []T, filter func(item *T) bool) int {
	for i := 0; i < len(slice); i++ {
		if filter(&slice[i]) {
			return i
		}
	}
	return -1
}
