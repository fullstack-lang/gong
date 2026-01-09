// generated code - do not edit
package models

import (
	"fmt"
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

	var newInstancesStmt string
	_ = newInstancesStmt
	var fieldsEditStmt string
	_ = fieldsEditStmt
	var deletedInstancesStmt string
	_ = deletedInstancesStmt

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
			newInstancesStmt += button.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := button.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := button.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", button.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for button := range stage.Buttons_reference {
		if _, ok := stage.Buttons[button]; !ok {
			buttons_deletedInstances = append(buttons_deletedInstances, button)
			deletedInstancesStmt += button.GongMarshallUnstaging(stage)
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
			newInstancesStmt += node.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := node.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := node.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", node.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for node := range stage.Nodes_reference {
		if _, ok := stage.Nodes[node]; !ok {
			nodes_deletedInstances = append(nodes_deletedInstances, node)
			deletedInstancesStmt += node.GongMarshallUnstaging(stage)
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
			newInstancesStmt += svgicon.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := svgicon.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := svgicon.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", svgicon.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for svgicon := range stage.SVGIcons_reference {
		if _, ok := stage.SVGIcons[svgicon]; !ok {
			svgicons_deletedInstances = append(svgicons_deletedInstances, svgicon)
			deletedInstancesStmt += svgicon.GongMarshallUnstaging(stage)
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
			newInstancesStmt += tree.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := tree.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := tree.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", tree.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for tree := range stage.Trees_reference {
		if _, ok := stage.Trees[tree]; !ok {
			trees_deletedInstances = append(trees_deletedInstances, tree)
			deletedInstancesStmt += tree.GongMarshallUnstaging(stage)
		}
	}

	lenNewInstances += len(trees_newInstances)
	lenDeletedInstances += len(trees_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		notif := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		notif += fmt.Sprintf("\n\t// %s", time.Now().Format(time.RFC3339Nano))
		notif += "\n\tstage.Commit()"
		if stage.GetProbeIF() != nil {
			stage.GetProbeIF().AddNotification(
				time.Now(),
				notif,
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
	return stage.ButtonMap_Staged_Order[button]
}

func (button *Button) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Buttons_referenceOrder[button]
}

func (node *Node) GongGetOrder(stage *Stage) uint {
	return stage.NodeMap_Staged_Order[node]
}

func (node *Node) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Nodes_referenceOrder[node]
}

func (svgicon *SVGIcon) GongGetOrder(stage *Stage) uint {
	return stage.SVGIconMap_Staged_Order[svgicon]
}

func (svgicon *SVGIcon) GongGetReferenceOrder(stage *Stage) uint {
	return stage.SVGIcons_referenceOrder[svgicon]
}

func (tree *Tree) GongGetOrder(stage *Stage) uint {
	return stage.TreeMap_Staged_Order[tree]
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
