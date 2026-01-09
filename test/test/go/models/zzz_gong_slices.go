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
	// Compute reverse map for named struct Astruct
	// insertion point per field
	stage.Astruct_Anarrayofb_reverseMap = make(map[*Bstruct]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _bstruct := range astruct.Anarrayofb {
			stage.Astruct_Anarrayofb_reverseMap[_bstruct] = astruct
		}
	}
	stage.Astruct_Dstruct4s_reverseMap = make(map[*Dstruct]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _dstruct := range astruct.Dstruct4s {
			stage.Astruct_Dstruct4s_reverseMap[_dstruct] = astruct
		}
	}
	stage.Astruct_Anarrayofa_reverseMap = make(map[*Astruct]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _astruct := range astruct.Anarrayofa {
			stage.Astruct_Anarrayofa_reverseMap[_astruct] = astruct
		}
	}
	stage.Astruct_Anotherarrayofb_reverseMap = make(map[*Bstruct]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _bstruct := range astruct.Anotherarrayofb {
			stage.Astruct_Anotherarrayofb_reverseMap[_bstruct] = astruct
		}
	}
	stage.Astruct_AnarrayofbUse_reverseMap = make(map[*AstructBstructUse]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _astructbstructuse := range astruct.AnarrayofbUse {
			stage.Astruct_AnarrayofbUse_reverseMap[_astructbstructuse] = astruct
		}
	}
	stage.Astruct_Anarrayofb2Use_reverseMap = make(map[*AstructBstruct2Use]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _astructbstruct2use := range astruct.Anarrayofb2Use {
			stage.Astruct_Anarrayofb2Use_reverseMap[_astructbstruct2use] = astruct
		}
	}

	// Compute reverse map for named struct AstructBstruct2Use
	// insertion point per field

	// Compute reverse map for named struct AstructBstructUse
	// insertion point per field

	// Compute reverse map for named struct Bstruct
	// insertion point per field

	// Compute reverse map for named struct Dstruct
	// insertion point per field
	stage.Dstruct_Anarrayofb_reverseMap = make(map[*Bstruct]*Dstruct)
	for dstruct := range stage.Dstructs {
		_ = dstruct
		for _, _bstruct := range dstruct.Anarrayofb {
			stage.Dstruct_Anarrayofb_reverseMap[_bstruct] = dstruct
		}
	}
	stage.Dstruct_Gstructs_reverseMap = make(map[*Gstruct]*Dstruct)
	for dstruct := range stage.Dstructs {
		_ = dstruct
		for _, _gstruct := range dstruct.Gstructs {
			stage.Dstruct_Gstructs_reverseMap[_gstruct] = dstruct
		}
	}

	// Compute reverse map for named struct F0123456789012345678901234567890
	// insertion point per field

	// Compute reverse map for named struct Gstruct
	// insertion point per field

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.Astructs {
		res = append(res, instance)
	}

	for instance := range stage.AstructBstruct2Uses {
		res = append(res, instance)
	}

	for instance := range stage.AstructBstructUses {
		res = append(res, instance)
	}

	for instance := range stage.Bstructs {
		res = append(res, instance)
	}

	for instance := range stage.Dstructs {
		res = append(res, instance)
	}

	for instance := range stage.F0123456789012345678901234567890s {
		res = append(res, instance)
	}

	for instance := range stage.Gstructs {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (astruct *Astruct) GongCopy() GongstructIF {
	newInstance := *astruct
	return &newInstance
}

func (astructbstruct2use *AstructBstruct2Use) GongCopy() GongstructIF {
	newInstance := *astructbstruct2use
	return &newInstance
}

func (astructbstructuse *AstructBstructUse) GongCopy() GongstructIF {
	newInstance := *astructbstructuse
	return &newInstance
}

func (bstruct *Bstruct) GongCopy() GongstructIF {
	newInstance := *bstruct
	return &newInstance
}

func (dstruct *Dstruct) GongCopy() GongstructIF {
	newInstance := *dstruct
	return &newInstance
}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongCopy() GongstructIF {
	newInstance := *f0123456789012345678901234567890
	return &newInstance
}

func (gstruct *Gstruct) GongCopy() GongstructIF {
	newInstance := *gstruct
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
	var astructs_newInstances []*Astruct
	var astructs_deletedInstances []*Astruct

	// parse all staged instances and check if they have a reference
	for astruct := range stage.Astructs {
		if ref, ok := stage.Astructs_reference[astruct]; !ok {
			astructs_newInstances = append(astructs_newInstances, astruct)
			newInstancesStmt += astruct.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := astruct.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := astruct.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", astruct.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for astruct := range stage.Astructs_reference {
		if _, ok := stage.Astructs[astruct]; !ok {
			astructs_deletedInstances = append(astructs_deletedInstances, astruct)
			deletedInstancesStmt += astruct.GongMarshallUnstaging(stage)
		}
	}

	lenNewInstances += len(astructs_newInstances)
	lenDeletedInstances += len(astructs_deletedInstances)
	var astructbstruct2uses_newInstances []*AstructBstruct2Use
	var astructbstruct2uses_deletedInstances []*AstructBstruct2Use

	// parse all staged instances and check if they have a reference
	for astructbstruct2use := range stage.AstructBstruct2Uses {
		if ref, ok := stage.AstructBstruct2Uses_reference[astructbstruct2use]; !ok {
			astructbstruct2uses_newInstances = append(astructbstruct2uses_newInstances, astructbstruct2use)
			newInstancesStmt += astructbstruct2use.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := astructbstruct2use.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := astructbstruct2use.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", astructbstruct2use.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for astructbstruct2use := range stage.AstructBstruct2Uses_reference {
		if _, ok := stage.AstructBstruct2Uses[astructbstruct2use]; !ok {
			astructbstruct2uses_deletedInstances = append(astructbstruct2uses_deletedInstances, astructbstruct2use)
			deletedInstancesStmt += astructbstruct2use.GongMarshallUnstaging(stage)
		}
	}

	lenNewInstances += len(astructbstruct2uses_newInstances)
	lenDeletedInstances += len(astructbstruct2uses_deletedInstances)
	var astructbstructuses_newInstances []*AstructBstructUse
	var astructbstructuses_deletedInstances []*AstructBstructUse

	// parse all staged instances and check if they have a reference
	for astructbstructuse := range stage.AstructBstructUses {
		if ref, ok := stage.AstructBstructUses_reference[astructbstructuse]; !ok {
			astructbstructuses_newInstances = append(astructbstructuses_newInstances, astructbstructuse)
			newInstancesStmt += astructbstructuse.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := astructbstructuse.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := astructbstructuse.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", astructbstructuse.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for astructbstructuse := range stage.AstructBstructUses_reference {
		if _, ok := stage.AstructBstructUses[astructbstructuse]; !ok {
			astructbstructuses_deletedInstances = append(astructbstructuses_deletedInstances, astructbstructuse)
			deletedInstancesStmt += astructbstructuse.GongMarshallUnstaging(stage)
		}
	}

	lenNewInstances += len(astructbstructuses_newInstances)
	lenDeletedInstances += len(astructbstructuses_deletedInstances)
	var bstructs_newInstances []*Bstruct
	var bstructs_deletedInstances []*Bstruct

	// parse all staged instances and check if they have a reference
	for bstruct := range stage.Bstructs {
		if ref, ok := stage.Bstructs_reference[bstruct]; !ok {
			bstructs_newInstances = append(bstructs_newInstances, bstruct)
			newInstancesStmt += bstruct.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := bstruct.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := bstruct.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", bstruct.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for bstruct := range stage.Bstructs_reference {
		if _, ok := stage.Bstructs[bstruct]; !ok {
			bstructs_deletedInstances = append(bstructs_deletedInstances, bstruct)
			deletedInstancesStmt += bstruct.GongMarshallUnstaging(stage)
		}
	}

	lenNewInstances += len(bstructs_newInstances)
	lenDeletedInstances += len(bstructs_deletedInstances)
	var dstructs_newInstances []*Dstruct
	var dstructs_deletedInstances []*Dstruct

	// parse all staged instances and check if they have a reference
	for dstruct := range stage.Dstructs {
		if ref, ok := stage.Dstructs_reference[dstruct]; !ok {
			dstructs_newInstances = append(dstructs_newInstances, dstruct)
			newInstancesStmt += dstruct.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := dstruct.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := dstruct.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", dstruct.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for dstruct := range stage.Dstructs_reference {
		if _, ok := stage.Dstructs[dstruct]; !ok {
			dstructs_deletedInstances = append(dstructs_deletedInstances, dstruct)
			deletedInstancesStmt += dstruct.GongMarshallUnstaging(stage)
		}
	}

	lenNewInstances += len(dstructs_newInstances)
	lenDeletedInstances += len(dstructs_deletedInstances)
	var f0123456789012345678901234567890s_newInstances []*F0123456789012345678901234567890
	var f0123456789012345678901234567890s_deletedInstances []*F0123456789012345678901234567890

	// parse all staged instances and check if they have a reference
	for f0123456789012345678901234567890 := range stage.F0123456789012345678901234567890s {
		if ref, ok := stage.F0123456789012345678901234567890s_reference[f0123456789012345678901234567890]; !ok {
			f0123456789012345678901234567890s_newInstances = append(f0123456789012345678901234567890s_newInstances, f0123456789012345678901234567890)
			newInstancesStmt += f0123456789012345678901234567890.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := f0123456789012345678901234567890.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := f0123456789012345678901234567890.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", f0123456789012345678901234567890.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for f0123456789012345678901234567890 := range stage.F0123456789012345678901234567890s_reference {
		if _, ok := stage.F0123456789012345678901234567890s[f0123456789012345678901234567890]; !ok {
			f0123456789012345678901234567890s_deletedInstances = append(f0123456789012345678901234567890s_deletedInstances, f0123456789012345678901234567890)
			deletedInstancesStmt += f0123456789012345678901234567890.GongMarshallUnstaging(stage)
		}
	}

	lenNewInstances += len(f0123456789012345678901234567890s_newInstances)
	lenDeletedInstances += len(f0123456789012345678901234567890s_deletedInstances)
	var gstructs_newInstances []*Gstruct
	var gstructs_deletedInstances []*Gstruct

	// parse all staged instances and check if they have a reference
	for gstruct := range stage.Gstructs {
		if ref, ok := stage.Gstructs_reference[gstruct]; !ok {
			gstructs_newInstances = append(gstructs_newInstances, gstruct)
			newInstancesStmt += gstruct.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := gstruct.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := gstruct.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", gstruct.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for gstruct := range stage.Gstructs_reference {
		if _, ok := stage.Gstructs[gstruct]; !ok {
			gstructs_deletedInstances = append(gstructs_deletedInstances, gstruct)
			deletedInstancesStmt += gstruct.GongMarshallUnstaging(stage)
		}
	}

	lenNewInstances += len(gstructs_newInstances)
	lenDeletedInstances += len(gstructs_deletedInstances)

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
	stage.Astructs_reference = make(map[*Astruct]*Astruct)
	stage.Astructs_referenceOrder = make(map[*Astruct]uint) // diff Unstage needs the reference order
	for instance := range stage.Astructs {
		stage.Astructs_reference[instance] = instance.GongCopy().(*Astruct)
		stage.Astructs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.AstructBstruct2Uses_reference = make(map[*AstructBstruct2Use]*AstructBstruct2Use)
	stage.AstructBstruct2Uses_referenceOrder = make(map[*AstructBstruct2Use]uint) // diff Unstage needs the reference order
	for instance := range stage.AstructBstruct2Uses {
		stage.AstructBstruct2Uses_reference[instance] = instance.GongCopy().(*AstructBstruct2Use)
		stage.AstructBstruct2Uses_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.AstructBstructUses_reference = make(map[*AstructBstructUse]*AstructBstructUse)
	stage.AstructBstructUses_referenceOrder = make(map[*AstructBstructUse]uint) // diff Unstage needs the reference order
	for instance := range stage.AstructBstructUses {
		stage.AstructBstructUses_reference[instance] = instance.GongCopy().(*AstructBstructUse)
		stage.AstructBstructUses_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Bstructs_reference = make(map[*Bstruct]*Bstruct)
	stage.Bstructs_referenceOrder = make(map[*Bstruct]uint) // diff Unstage needs the reference order
	for instance := range stage.Bstructs {
		stage.Bstructs_reference[instance] = instance.GongCopy().(*Bstruct)
		stage.Bstructs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Dstructs_reference = make(map[*Dstruct]*Dstruct)
	stage.Dstructs_referenceOrder = make(map[*Dstruct]uint) // diff Unstage needs the reference order
	for instance := range stage.Dstructs {
		stage.Dstructs_reference[instance] = instance.GongCopy().(*Dstruct)
		stage.Dstructs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.F0123456789012345678901234567890s_reference = make(map[*F0123456789012345678901234567890]*F0123456789012345678901234567890)
	stage.F0123456789012345678901234567890s_referenceOrder = make(map[*F0123456789012345678901234567890]uint) // diff Unstage needs the reference order
	for instance := range stage.F0123456789012345678901234567890s {
		stage.F0123456789012345678901234567890s_reference[instance] = instance.GongCopy().(*F0123456789012345678901234567890)
		stage.F0123456789012345678901234567890s_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Gstructs_reference = make(map[*Gstruct]*Gstruct)
	stage.Gstructs_referenceOrder = make(map[*Gstruct]uint) // diff Unstage needs the reference order
	for instance := range stage.Gstructs {
		stage.Gstructs_reference[instance] = instance.GongCopy().(*Gstruct)
		stage.Gstructs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (astruct *Astruct) GongGetOrder(stage *Stage) uint {
	return stage.AstructMap_Staged_Order[astruct]
}

func (astruct *Astruct) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Astructs_referenceOrder[astruct]
}

func (astructbstruct2use *AstructBstruct2Use) GongGetOrder(stage *Stage) uint {
	return stage.AstructBstruct2UseMap_Staged_Order[astructbstruct2use]
}

func (astructbstruct2use *AstructBstruct2Use) GongGetReferenceOrder(stage *Stage) uint {
	return stage.AstructBstruct2Uses_referenceOrder[astructbstruct2use]
}

func (astructbstructuse *AstructBstructUse) GongGetOrder(stage *Stage) uint {
	return stage.AstructBstructUseMap_Staged_Order[astructbstructuse]
}

func (astructbstructuse *AstructBstructUse) GongGetReferenceOrder(stage *Stage) uint {
	return stage.AstructBstructUses_referenceOrder[astructbstructuse]
}

func (bstruct *Bstruct) GongGetOrder(stage *Stage) uint {
	return stage.BstructMap_Staged_Order[bstruct]
}

func (bstruct *Bstruct) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Bstructs_referenceOrder[bstruct]
}

func (dstruct *Dstruct) GongGetOrder(stage *Stage) uint {
	return stage.DstructMap_Staged_Order[dstruct]
}

func (dstruct *Dstruct) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Dstructs_referenceOrder[dstruct]
}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongGetOrder(stage *Stage) uint {
	return stage.F0123456789012345678901234567890Map_Staged_Order[f0123456789012345678901234567890]
}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongGetReferenceOrder(stage *Stage) uint {
	return stage.F0123456789012345678901234567890s_referenceOrder[f0123456789012345678901234567890]
}

func (gstruct *Gstruct) GongGetOrder(stage *Stage) uint {
	return stage.GstructMap_Staged_Order[gstruct]
}

func (gstruct *Gstruct) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Gstructs_referenceOrder[gstruct]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (astruct *Astruct) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", astruct.GongGetGongstructName(), astruct.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (astruct *Astruct) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", astruct.GongGetGongstructName(), astruct.GongGetReferenceOrder(stage))
}

func (astructbstruct2use *AstructBstruct2Use) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", astructbstruct2use.GongGetGongstructName(), astructbstruct2use.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (astructbstruct2use *AstructBstruct2Use) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", astructbstruct2use.GongGetGongstructName(), astructbstruct2use.GongGetReferenceOrder(stage))
}

func (astructbstructuse *AstructBstructUse) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", astructbstructuse.GongGetGongstructName(), astructbstructuse.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (astructbstructuse *AstructBstructUse) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", astructbstructuse.GongGetGongstructName(), astructbstructuse.GongGetReferenceOrder(stage))
}

func (bstruct *Bstruct) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bstruct.GongGetGongstructName(), bstruct.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bstruct *Bstruct) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bstruct.GongGetGongstructName(), bstruct.GongGetReferenceOrder(stage))
}

func (dstruct *Dstruct) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", dstruct.GongGetGongstructName(), dstruct.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (dstruct *Dstruct) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", dstruct.GongGetGongstructName(), dstruct.GongGetReferenceOrder(stage))
}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", f0123456789012345678901234567890.GongGetGongstructName(), f0123456789012345678901234567890.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", f0123456789012345678901234567890.GongGetGongstructName(), f0123456789012345678901234567890.GongGetReferenceOrder(stage))
}

func (gstruct *Gstruct) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gstruct.GongGetGongstructName(), gstruct.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gstruct *Gstruct) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gstruct.GongGetGongstructName(), gstruct.GongGetReferenceOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (astruct *Astruct) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", astruct.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Astruct")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", astruct.Name)
	return
}
func (astructbstruct2use *AstructBstruct2Use) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", astructbstruct2use.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AstructBstruct2Use")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", astructbstruct2use.Name)
	return
}
func (astructbstructuse *AstructBstructUse) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", astructbstructuse.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AstructBstructUse")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", astructbstructuse.Name)
	return
}
func (bstruct *Bstruct) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bstruct.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Bstruct")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", bstruct.Name)
	return
}
func (dstruct *Dstruct) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", dstruct.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Dstruct")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", dstruct.Name)
	return
}
func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", f0123456789012345678901234567890.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "F0123456789012345678901234567890")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", f0123456789012345678901234567890.Name)
	return
}
func (gstruct *Gstruct) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = IdentifiersDeclsWithoutNameInit
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gstruct.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Gstruct")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", gstruct.Name)
	return
}

// insertion point for unstaging
func (astruct *Astruct) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", astruct.GongGetReferenceIdentifier(stage))
	return
}
func (astructbstruct2use *AstructBstruct2Use) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", astructbstruct2use.GongGetReferenceIdentifier(stage))
	return
}
func (astructbstructuse *AstructBstructUse) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", astructbstructuse.GongGetReferenceIdentifier(stage))
	return
}
func (bstruct *Bstruct) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bstruct.GongGetReferenceIdentifier(stage))
	return
}
func (dstruct *Dstruct) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", dstruct.GongGetReferenceIdentifier(stage))
	return
}
func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", f0123456789012345678901234567890.GongGetReferenceIdentifier(stage))
	return
}
func (gstruct *Gstruct) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gstruct.GongGetReferenceIdentifier(stage))
	return
}
