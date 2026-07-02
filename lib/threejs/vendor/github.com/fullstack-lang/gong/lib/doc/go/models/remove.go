package models

// remove node from slice
func remove[T Gongstruct](slice []*T, t *T) []*T {

	// get the index of the t
	rank := -1
	for i, t_ := range slice {
		if t_ == t {
			rank = i
		}
	}
	return append(slice[:rank], slice[rank+1:]...)
}
