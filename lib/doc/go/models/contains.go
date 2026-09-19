package models

import "slices"

func contains[T comparable](elems []T, v T) bool {
	return slices.Contains(elems, v)
}
