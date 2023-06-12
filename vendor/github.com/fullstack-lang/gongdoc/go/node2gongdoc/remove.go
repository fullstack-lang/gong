package node2gongdoc

import gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

func remove[T gongdoc_models.Gongstruct](slice []*T, t *T) []*T {

	// get the index of the t
	rank := -1
	for i, t_ := range slice {
		if t_ == t {
			rank = i
		}
	}
	return append(slice[:rank], slice[rank+1:]...)
}
