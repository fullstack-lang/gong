package node2gongdoc

import (
	"cmp"

	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

func CompareNodeNames(a, b *gongtree_models.Node) int {
	return cmp.Compare(a.Name, b.Name)
}
