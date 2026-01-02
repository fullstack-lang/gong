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
	// Compute reverse map for named struct A
	// insertion point per field
	stage.A_Bs_reverseMap = make(map[*B]*A)
	for a := range stage.As {
		_ = a
		for _, _b := range a.Bs {
			stage.A_Bs_reverseMap[_b] = a
		}
	}

	// Compute reverse map for named struct B
	// insertion point per field

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.As {
		res = append(res, instance)
	}

	for instance := range stage.Bs {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (a *A) GongCopy() GongstructIF {
	newInstance := *a
	return &newInstance
}

func (b *B) GongCopy() GongstructIF {
	newInstance := *b
	return &newInstance
}

func (stage *Stage) ComputeDifference() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var pointersInitializesStatements string

	// insertion point per named struct
	var as_newInstances []*A
	var as_deletedInstances []*A

	// parse all staged instances and check if they have a reference
	for a := range stage.As {
		if ref, ok := stage.As_reference[a]; !ok {
			as_newInstances = append(as_newInstances, a)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of A "+a.Name,
				)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					a.GongMarshallIdentifier(stage),
				)
				basicFieldInitializers, pointersInitializations := a.GongMarshallAllFields(stage)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					basicFieldInitializers,
				)
				pointersInitializesStatements += pointersInitializations
			}
		} else {
			diffs := a.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of A \""+a.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
					for _, diff := range diffs {
						stage.GetProbeIF().AddNotification(
							time.Now(),
							a.GongMarshallField(stage, diff),
						)
					}
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for a := range stage.As_reference {
		if _, ok := stage.As[a]; !ok {
			as_deletedInstances = append(as_deletedInstances, a)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of A "+a.Name,
				)
			}
		}
	}

	lenNewInstances += len(as_newInstances)
	lenDeletedInstances += len(as_deletedInstances)
	var bs_newInstances []*B
	var bs_deletedInstances []*B

	// parse all staged instances and check if they have a reference
	for b := range stage.Bs {
		if ref, ok := stage.Bs_reference[b]; !ok {
			bs_newInstances = append(bs_newInstances, b)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of B "+b.Name,
				)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					b.GongMarshallIdentifier(stage),
				)
				basicFieldInitializers, pointersInitializations := b.GongMarshallAllFields(stage)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					basicFieldInitializers,
				)
				pointersInitializesStatements += pointersInitializations
			}
		} else {
			diffs := b.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of B \""+b.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
					for _, diff := range diffs {
						stage.GetProbeIF().AddNotification(
							time.Now(),
							b.GongMarshallField(stage, diff),
						)
					}
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for b := range stage.Bs_reference {
		if _, ok := stage.Bs[b]; !ok {
			bs_deletedInstances = append(bs_deletedInstances, b)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of B "+b.Name,
				)
			}
		}
	}

	lenNewInstances += len(bs_newInstances)
	lenDeletedInstances += len(bs_deletedInstances)

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
	stage.As_reference = make(map[*A]*A)
	for instance := range stage.As {
		stage.As_reference[instance] = instance.GongCopy().(*A)
	}

	stage.Bs_reference = make(map[*B]*B)
	for instance := range stage.Bs {
		stage.Bs_reference[instance] = instance.GongCopy().(*B)
	}

}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (a *A) GongGetOrder(stage *Stage) uint {
	return stage.AMap_Staged_Order[a]
}

func (b *B) GongGetOrder(stage *Stage) uint {
	return stage.BMap_Staged_Order[b]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (a *A) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a.GongGetGongstructName(), a.GongGetOrder(stage))
}

func (b *B) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", b.GongGetGongstructName(), b.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (a *A) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = IdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a.Name)
	return
}
func (b *B) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = IdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", b.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "B")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", b.Name)
	return
}
