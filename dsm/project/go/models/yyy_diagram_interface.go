// generated code (do not edit)
package models

type DiagramIF interface {
	GetName() string
	GetIsChecked() bool
	SetIsChecked(bool)
	SetIsShowPrefix(bool)
	GetIsShowPrefix() bool

	GetDefaultBoxWidth() float64
	GetDefaultBoxHeigth() float64
	IsEditable() bool

	// within the tree branch of one diagram, when an element is present in more than one diagram,
	// it is possible to access it via a list. Only one element have a list that is available per diagram.
	GetDiagramListElement() AbstractType
	SetDiagramListElement(AbstractType)
}
