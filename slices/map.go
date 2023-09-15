package slices

func Map[Input, Output any](slice []Input, lambda func(item *Input) Output) []Output {
	ret := []Output{}
	for _, item := range slice {
		ret = append(ret, lambda(&item))
	}
	return ret
}
