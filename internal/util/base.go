package util

import "golang.org/x/exp/constraints"

// ToPtrSlice converts a slice of T to a slice of *T.
func ToPtrSlice[T any](slice []T) []*T {
	var ptrSlice []*T
	for _, item := range slice {
		ptrSlice = append(ptrSlice, &item)
	}
	return ptrSlice
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
