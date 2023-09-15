package easygo

func FindIndexByValue[T comparable](slice []T, value T) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == value {
			return i
		}
	}
	return -1
}
