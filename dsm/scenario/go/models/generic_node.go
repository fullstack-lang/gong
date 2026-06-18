package models

// GenericNode is a Gongstruct that is also used as a node
// in the tree
type GenericNode[T Gongstruct] interface {
	*T
	GetIsNodeExpanded() bool
	SetIsNodeExpanded(isNodeExpanded bool)
}
