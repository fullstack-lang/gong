// generated code - do not edit
package models

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

var __GongSliceTemplate_time__dummyDeclaration time.Duration
var _ = __GongSliceTemplate_time__dummyDeclaration

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Button
	// insertion point per field

	// Compute reverse map for named struct Node
	// insertion point per field
	stage.Node_Children_reverseMap = make(map[*Node]*Node)
	for node := range stage.Nodes {
		_ = node
		for _, _node := range node.Children {
			stage.Node_Children_reverseMap[_node] = node
		}
	}
	stage.Node_Buttons_reverseMap = make(map[*Button]*Node)
	for node := range stage.Nodes {
		_ = node
		for _, _button := range node.Buttons {
			stage.Node_Buttons_reverseMap[_button] = node
		}
	}

	// Compute reverse map for named struct SVGIcon
	// insertion point per field

	// Compute reverse map for named struct Tree
	// insertion point per field
	stage.Tree_RootNodes_reverseMap = make(map[*Node]*Tree)
	for tree := range stage.Trees {
		_ = tree
		for _, _node := range tree.RootNodes {
			stage.Tree_RootNodes_reverseMap[_node] = tree
		}
	}

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.Buttons {
		res = append(res, instance)
	}

	for instance := range stage.Nodes {
		res = append(res, instance)
	}

	for instance := range stage.SVGIcons {
		res = append(res, instance)
	}

	for instance := range stage.Trees {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (button *Button) GongCopy() GongstructIF {
	newInstance := *button
	return &newInstance
}

func (node *Node) GongCopy() GongstructIF {
	newInstance := *node
	return &newInstance
}

func (svgicon *SVGIcon) GongCopy() GongstructIF {
	newInstance := *svgicon
	return &newInstance
}

func (tree *Tree) GongCopy() GongstructIF {
	newInstance := *tree
	return &newInstance
}

func (stage *Stage) ComputeDifference() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	var buttons_newInstances []*Button
	var buttons_deletedInstances []*Button

	// parse all staged instances and check if they have a reference
	for button := range stage.Buttons {
		if ref, ok := stage.Buttons_reference[button]; !ok {
			buttons_newInstances = append(buttons_newInstances, button)
			newInstancesSlice = append(newInstancesSlice, button.GongMarshallIdentifier(stage))
			if stage.Buttons_referenceOrder == nil {
				stage.Buttons_referenceOrder = make(map[*Button]uint)
			}
			stage.Buttons_referenceOrder[button] = stage.ButtonMap_Staged_Order[button]
			newInstancesReverseSlice = append(newInstancesReverseSlice, button.GongMarshallUnstaging(stage))
			delete(stage.Buttons_referenceOrder, button)
			fieldInitializers, pointersInitializations := button.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ButtonMap_Staged_Order[ref] = stage.ButtonMap_Staged_Order[button]
			diffs := button.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, button)
			delete(stage.ButtonMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", button.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Buttons_reference {
		if _, ok := stage.Buttons[ref]; !ok {
			buttons_deletedInstances = append(buttons_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(buttons_newInstances)
	lenDeletedInstances += len(buttons_deletedInstances)
	var nodes_newInstances []*Node
	var nodes_deletedInstances []*Node

	// parse all staged instances and check if they have a reference
	for node := range stage.Nodes {
		if ref, ok := stage.Nodes_reference[node]; !ok {
			nodes_newInstances = append(nodes_newInstances, node)
			newInstancesSlice = append(newInstancesSlice, node.GongMarshallIdentifier(stage))
			if stage.Nodes_referenceOrder == nil {
				stage.Nodes_referenceOrder = make(map[*Node]uint)
			}
			stage.Nodes_referenceOrder[node] = stage.NodeMap_Staged_Order[node]
			newInstancesReverseSlice = append(newInstancesReverseSlice, node.GongMarshallUnstaging(stage))
			delete(stage.Nodes_referenceOrder, node)
			fieldInitializers, pointersInitializations := node.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.NodeMap_Staged_Order[ref] = stage.NodeMap_Staged_Order[node]
			diffs := node.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, node)
			delete(stage.NodeMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", node.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Nodes_reference {
		if _, ok := stage.Nodes[ref]; !ok {
			nodes_deletedInstances = append(nodes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(nodes_newInstances)
	lenDeletedInstances += len(nodes_deletedInstances)
	var svgicons_newInstances []*SVGIcon
	var svgicons_deletedInstances []*SVGIcon

	// parse all staged instances and check if they have a reference
	for svgicon := range stage.SVGIcons {
		if ref, ok := stage.SVGIcons_reference[svgicon]; !ok {
			svgicons_newInstances = append(svgicons_newInstances, svgicon)
			newInstancesSlice = append(newInstancesSlice, svgicon.GongMarshallIdentifier(stage))
			if stage.SVGIcons_referenceOrder == nil {
				stage.SVGIcons_referenceOrder = make(map[*SVGIcon]uint)
			}
			stage.SVGIcons_referenceOrder[svgicon] = stage.SVGIconMap_Staged_Order[svgicon]
			newInstancesReverseSlice = append(newInstancesReverseSlice, svgicon.GongMarshallUnstaging(stage))
			delete(stage.SVGIcons_referenceOrder, svgicon)
			fieldInitializers, pointersInitializations := svgicon.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SVGIconMap_Staged_Order[ref] = stage.SVGIconMap_Staged_Order[svgicon]
			diffs := svgicon.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, svgicon)
			delete(stage.SVGIconMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", svgicon.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.SVGIcons_reference {
		if _, ok := stage.SVGIcons[ref]; !ok {
			svgicons_deletedInstances = append(svgicons_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(svgicons_newInstances)
	lenDeletedInstances += len(svgicons_deletedInstances)
	var trees_newInstances []*Tree
	var trees_deletedInstances []*Tree

	// parse all staged instances and check if they have a reference
	for tree := range stage.Trees {
		if ref, ok := stage.Trees_reference[tree]; !ok {
			trees_newInstances = append(trees_newInstances, tree)
			newInstancesSlice = append(newInstancesSlice, tree.GongMarshallIdentifier(stage))
			if stage.Trees_referenceOrder == nil {
				stage.Trees_referenceOrder = make(map[*Tree]uint)
			}
			stage.Trees_referenceOrder[tree] = stage.TreeMap_Staged_Order[tree]
			newInstancesReverseSlice = append(newInstancesReverseSlice, tree.GongMarshallUnstaging(stage))
			delete(stage.Trees_referenceOrder, tree)
			fieldInitializers, pointersInitializations := tree.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TreeMap_Staged_Order[ref] = stage.TreeMap_Staged_Order[tree]
			diffs := tree.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, tree)
			delete(stage.TreeMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", tree.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Trees_reference {
		if _, ok := stage.Trees[ref]; !ok {
			trees_deletedInstances = append(trees_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(trees_newInstances)
	lenDeletedInstances += len(trees_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.forwardCommits = append(stage.forwardCommits, forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.backwardCommits = append(stage.backwardCommits, backwardCommit)

		if stage.GetProbeIF() != nil {
			var mergedCommits string
			for _, commit := range stage.forwardCommits {
				mergedCommits += commit
			}
			stage.GetProbeIF().AddNotification(
				time.Now(),
				"	// Forward commits:\n"+
					mergedCommits,
			)

			var reverseMergedCommits string
			for _, reverserCommit := range stage.backwardCommits {
				reverseMergedCommits += reverserCommit
			}
			stage.GetProbeIF().AddNotification(
				time.Now(),
				"	// Backward commits:\n"+
					reverseMergedCommits,
			)

			stage.GetProbeIF().CommitNotificationTable()
		}
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.Buttons_reference = make(map[*Button]*Button)
	stage.Buttons_referenceOrder = make(map[*Button]uint) // diff Unstage needs the reference order
	for instance := range stage.Buttons {
		stage.Buttons_reference[instance] = instance.GongCopy().(*Button)
		stage.Buttons_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Nodes_reference = make(map[*Node]*Node)
	stage.Nodes_referenceOrder = make(map[*Node]uint) // diff Unstage needs the reference order
	for instance := range stage.Nodes {
		stage.Nodes_reference[instance] = instance.GongCopy().(*Node)
		stage.Nodes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.SVGIcons_reference = make(map[*SVGIcon]*SVGIcon)
	stage.SVGIcons_referenceOrder = make(map[*SVGIcon]uint) // diff Unstage needs the reference order
	for instance := range stage.SVGIcons {
		stage.SVGIcons_reference[instance] = instance.GongCopy().(*SVGIcon)
		stage.SVGIcons_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Trees_reference = make(map[*Tree]*Tree)
	stage.Trees_referenceOrder = make(map[*Tree]uint) // diff Unstage needs the reference order
	for instance := range stage.Trees {
		stage.Trees_reference[instance] = instance.GongCopy().(*Tree)
		stage.Trees_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (button *Button) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ButtonMap_Staged_Order[button]; ok {
		return order
	}
	return stage.Buttons_referenceOrder[button]
}

func (button *Button) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Buttons_referenceOrder[button]
}

func (node *Node) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.NodeMap_Staged_Order[node]; ok {
		return order
	}
	return stage.Nodes_referenceOrder[node]
}

func (node *Node) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Nodes_referenceOrder[node]
}

func (svgicon *SVGIcon) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SVGIconMap_Staged_Order[svgicon]; ok {
		return order
	}
	return stage.SVGIcons_referenceOrder[svgicon]
}

func (svgicon *SVGIcon) GongGetReferenceOrder(stage *Stage) uint {
	return stage.SVGIcons_referenceOrder[svgicon]
}

func (tree *Tree) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TreeMap_Staged_Order[tree]; ok {
		return order
	}
	return stage.Trees_referenceOrder[tree]
}

func (tree *Tree) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Trees_referenceOrder[tree]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (button *Button) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", button.GongGetGongstructName(), button.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (button *Button) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", button.GongGetGongstructName(), button.GongGetReferenceOrder(stage))
}

func (node *Node) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", node.GongGetGongstructName(), node.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (node *Node) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", node.GongGetGongstructName(), node.GongGetReferenceOrder(stage))
}

func (svgicon *SVGIcon) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", svgicon.GongGetGongstructName(), svgicon.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (svgicon *SVGIcon) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", svgicon.GongGetGongstructName(), svgicon.GongGetReferenceOrder(stage))
}

func (tree *Tree) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tree.GongGetGongstructName(), tree.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tree *Tree) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tree.GongGetGongstructName(), tree.GongGetReferenceOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (button *Button) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", button.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Button")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", button.Name)
	return
}
func (node *Node) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", node.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Node")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", node.Name)
	return
}
func (svgicon *SVGIcon) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", svgicon.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SVGIcon")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", svgicon.Name)
	return
}
func (tree *Tree) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tree.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Tree")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", tree.Name)
	return
}

// insertion point for unstaging
func (button *Button) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", button.GongGetReferenceIdentifier(stage))
	return
}
func (node *Node) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", node.GongGetReferenceIdentifier(stage))
	return
}
func (svgicon *SVGIcon) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", svgicon.GongGetReferenceIdentifier(stage))
	return
}
func (tree *Tree) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tree.GongGetReferenceIdentifier(stage))
	return
}
