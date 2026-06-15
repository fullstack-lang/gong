// generated code (do not edit)
package models

type AbstractType interface {
	GongstructIF
	GetIsExpanded() bool
	SetIsExpanded(bool)
	GetComputedPrefix() string
	SetComputedPrefix(string)
	GetComputedWidth() int
	SetComputedWidth(int)
	SetComputedPrefixInt([]int)
	GetIsInRenameMode() bool
	SetIsInRenameMode(bool)

	GetLayoutDirection() LayoutDirection
	SetLayoutDirection(LayoutDirection)
}

type AbstractTypeFields struct {
	// ComputedPrefix is automaticaly computed by the semantic enforcing mechanism
	ComputedPrefix string
	computedPrefix []int

	// nodes can be edited
	isInRenameMode bool
	IsExpanded     bool // to be made private once in production (no need to persist)

	// When the full PBS is displayed, the computedWidth is the number of node
	// aligned below. A leaf node has a computedWidth of 1
	computedWidth int

	// Directive for display in the concrete diagram
	LayoutDirection LayoutDirection
}

func (r *AbstractTypeFields) GetLayoutDirection() LayoutDirection {
	return r.LayoutDirection
}

func (r *AbstractTypeFields) SetLayoutDirection(d LayoutDirection) {
	r.LayoutDirection = d
}

func (r *AbstractTypeFields) GetComputedWidth() int {
	return r.computedWidth
}

func (r *AbstractTypeFields) SetComputedWidth(w int) {
	r.computedWidth = w
}

func (r *AbstractTypeFields) SetComputedPrefixInt(p []int) {
	r.computedPrefix = p
}

func (r *AbstractTypeFields) GetIsExpanded() bool {
	return r.IsExpanded
}

func (r *AbstractTypeFields) SetIsExpanded(isExpanded bool) {
	r.IsExpanded = isExpanded
}

func (r *AbstractTypeFields) GetComputedPrefix() string {
	return r.ComputedPrefix
}

func (r *AbstractTypeFields) SetComputedPrefix(ComputedPrefix string) {
	r.ComputedPrefix = ComputedPrefix
}

func (r *AbstractTypeFields) GetIsInRenameMode() bool {
	return r.isInRenameMode
}

func (r *AbstractTypeFields) SetIsInRenameMode(isInRenameMode bool) {
	r.isInRenameMode = isInRenameMode
}
