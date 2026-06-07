// generated code - do not edit
package models

import (
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

// onNameChange provides a reusable callback for tree.Node.OnNameChange
func (stager *Stager) onNameChange(element AbstractType) func(newName string) {
	return func(newName string) {
		element.SetName(newName)
		element.SetIsInRenameMode(false)
		stager.stage.Commit()
	}
}

func onNodeClicked[T AbstractType](stager *Stager, element T) func(frontNode *tree.Node) {
	return func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(element, GetPointerToGongstructName[T]())
		stager.stage.Commit()
	}
}

// onIsExpandedChangeSlice provides a reusable callback for tree.Node.OnIsExpandedChange backed by a slice
func onIsExpandedChangeSlice[T comparable](stager *Stager, element T, expandedSlice *[]T) func(isExpanded bool) {
	return func(isExpanded bool) {
		if isExpanded {
			if slices.Index(*expandedSlice, element) == -1 {
				*expandedSlice = append(*expandedSlice, element)
			}
		} else {
			if idx := slices.Index(*expandedSlice, element); idx != -1 {
				*expandedSlice = slices.Delete(*expandedSlice, idx, idx+1)
			}
		}
		stager.stage.Commit()
	}
}

// onIsExpandedChangeBool provides a reusable callback for tree.Node.OnIsExpandedChange backed by a boolean
func (stager *Stager) onIsExpandedChangeBool(isExpandedPtr *bool) func(isExpanded bool) {
	return func(isExpanded bool) {
		*isExpandedPtr = isExpanded
		stager.stage.Commit()
	}
}
