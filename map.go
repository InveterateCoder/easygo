package easygo

func Map[InT, OutT any](slice []InT, mapper func(item *InT) OutT) []OutT {
	ret := []OutT{}
	for _, item := range slice {
		ret = append(ret, mapper(&item))
	}
	return ret
}
