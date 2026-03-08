package models

// Append is a generic helper that appends an item to a slice via a pointer
func Append[T any](slice *[]T, item T) {
	*slice = append(*slice, item)
}
