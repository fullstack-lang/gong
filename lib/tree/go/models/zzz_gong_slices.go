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

	var pointersInitializesStatements string

	// insertion point per named struct
	var buttons_newInstances []*Button
	var buttons_deletedInstances []*Button

	// parse all staged instances and check if they have a reference
	for button := range stage.Buttons {
		if ref, ok := stage.Buttons_reference[button]; !ok {
			buttons_newInstances = append(buttons_newInstances, button)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Button "+button.Name,
				)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					button.GongMarshallIdentifier(stage),
				)
				basicFieldInitializers, pointersInitializations := button.GongMarshallAllFields(stage)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					basicFieldInitializers,
				)
				pointersInitializesStatements += pointersInitializations
			}
		} else {
			diffs := button.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Button \""+button.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
					for _, diff := range diffs {
						stage.GetProbeIF().AddNotification(
							time.Now(),
							button.GongMarshallField(stage, diff),
						)
					}
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for button := range stage.Buttons_reference {
		if _, ok := stage.Buttons[button]; !ok {
			buttons_deletedInstances = append(buttons_deletedInstances, button)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					button.GongMarshallUnstaging(stage),
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Node "+node.Name,
				)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					node.GongMarshallIdentifier(stage),
				)
				basicFieldInitializers, pointersInitializations := node.GongMarshallAllFields(stage)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					basicFieldInitializers,
				)
				pointersInitializesStatements += pointersInitializations
			}
		} else {
			diffs := node.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Node \""+node.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
					for _, diff := range diffs {
						stage.GetProbeIF().AddNotification(
							time.Now(),
							node.GongMarshallField(stage, diff),
						)
					}
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for node := range stage.Nodes_reference {
		if _, ok := stage.Nodes[node]; !ok {
			nodes_deletedInstances = append(nodes_deletedInstances, node)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					node.GongMarshallUnstaging(stage),
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of SVGIcon "+svgicon.Name,
				)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					svgicon.GongMarshallIdentifier(stage),
				)
				basicFieldInitializers, pointersInitializations := svgicon.GongMarshallAllFields(stage)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					basicFieldInitializers,
				)
				pointersInitializesStatements += pointersInitializations
			}
		} else {
			diffs := svgicon.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of SVGIcon \""+svgicon.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
					for _, diff := range diffs {
						stage.GetProbeIF().AddNotification(
							time.Now(),
							svgicon.GongMarshallField(stage, diff),
						)
					}
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for svgicon := range stage.SVGIcons_reference {
		if _, ok := stage.SVGIcons[svgicon]; !ok {
			svgicons_deletedInstances = append(svgicons_deletedInstances, svgicon)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					svgicon.GongMarshallUnstaging(stage),
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Tree "+tree.Name,
				)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					tree.GongMarshallIdentifier(stage),
				)
				basicFieldInitializers, pointersInitializations := tree.GongMarshallAllFields(stage)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					basicFieldInitializers,
				)
				pointersInitializesStatements += pointersInitializations
			}
		} else {
			diffs := tree.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Tree \""+tree.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
					for _, diff := range diffs {
						stage.GetProbeIF().AddNotification(
							time.Now(),
							tree.GongMarshallField(stage, diff),
						)
					}
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for tree := range stage.Trees_reference {
		if _, ok := stage.Trees[tree]; !ok {
			trees_deletedInstances = append(trees_deletedInstances, tree)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					tree.GongMarshallUnstaging(stage),
				)
			}
		}
	}

	lenNewInstances += len(trees_newInstances)
	lenDeletedInstances += len(trees_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		// if stage.GetProbeIF() != nil {
		// 	stage.GetProbeIF().CommitNotificationTable()
		// }
	}

	if pointersInitializesStatements != "" {
		if stage.GetProbeIF() != nil {
			stage.GetProbeIF().AddNotification(
				time.Now(),
				pointersInitializesStatements,
			)
		}
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.Buttons_reference = make(map[*Button]*Button)
	for instance := range stage.Buttons {
		stage.Buttons_reference[instance] = instance.GongCopy().(*Button)
	}

	stage.Nodes_reference = make(map[*Node]*Node)
	for instance := range stage.Nodes {
		stage.Nodes_reference[instance] = instance.GongCopy().(*Node)
	}

	stage.SVGIcons_reference = make(map[*SVGIcon]*SVGIcon)
	for instance := range stage.SVGIcons {
		stage.SVGIcons_reference[instance] = instance.GongCopy().(*SVGIcon)
	}

	stage.Trees_reference = make(map[*Tree]*Tree)
	for instance := range stage.Trees {
		stage.Trees_reference[instance] = instance.GongCopy().(*Tree)
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

func (node *Node) GongGetOrder(stage *Stage) uint {
	return stage.NodeMap_Staged_Order[node]
}

func (svgicon *SVGIcon) GongGetOrder(stage *Stage) uint {
	return stage.SVGIconMap_Staged_Order[svgicon]
}

func (tree *Tree) GongGetOrder(stage *Stage) uint {
	return stage.TreeMap_Staged_Order[tree]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (button *Button) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", button.GongGetGongstructName(), button.GongGetOrder(stage))
}

func (node *Node) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", node.GongGetGongstructName(), node.GongGetOrder(stage))
}

func (svgicon *SVGIcon) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", svgicon.GongGetGongstructName(), svgicon.GongGetOrder(stage))
}

func (tree *Tree) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tree.GongGetGongstructName(), tree.GongGetOrder(stage))
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
	decl = strings.ReplaceAll(decl, "{{Identifier}}", button.GongGetIdentifier(stage))
	return
}
func (node *Node) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", node.GongGetIdentifier(stage))
	return
}
func (svgicon *SVGIcon) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", svgicon.GongGetIdentifier(stage))
	return
}
func (tree *Tree) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tree.GongGetIdentifier(stage))
	return
}
