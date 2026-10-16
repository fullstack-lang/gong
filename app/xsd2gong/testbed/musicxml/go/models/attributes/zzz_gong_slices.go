// generated code - do not edit
package attributes

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Attributes
	// insertion point per field
	stage.Attributes_Key_reverseMap = make(map[*Key]*Attributes)
	for attributes := range stage.Attributess {
		_ = attributes
		for _, _key := range attributes.Key {
			stage.Attributes_Key_reverseMap[_key] = attributes
		}
	}
	stage.Attributes_Time_reverseMap = make(map[*Time]*Attributes)
	for attributes := range stage.Attributess {
		_ = attributes
		for _, _time := range attributes.Time {
			stage.Attributes_Time_reverseMap[_time] = attributes
		}
	}
	stage.Attributes_Clef_reverseMap = make(map[*Clef]*Attributes)
	for attributes := range stage.Attributess {
		_ = attributes
		for _, _clef := range attributes.Clef {
			stage.Attributes_Clef_reverseMap[_clef] = attributes
		}
	}
	stage.Attributes_StaffDetails_reverseMap = make(map[*StaffDetails]*Attributes)
	for attributes := range stage.Attributess {
		_ = attributes
		for _, _staffdetails := range attributes.StaffDetails {
			stage.Attributes_StaffDetails_reverseMap[_staffdetails] = attributes
		}
	}
	stage.Attributes_Transpose_reverseMap = make(map[*Transpose]*Attributes)
	for attributes := range stage.Attributess {
		_ = attributes
		for _, _transpose := range attributes.Transpose {
			stage.Attributes_Transpose_reverseMap[_transpose] = attributes
		}
	}
	stage.Attributes_ForPart_reverseMap = make(map[*ForPart]*Attributes)
	for attributes := range stage.Attributess {
		_ = attributes
		for _, _forpart := range attributes.ForPart {
			stage.Attributes_ForPart_reverseMap[_forpart] = attributes
		}
	}
	stage.Attributes_Directive_reverseMap = make(map[*AttributesDirective]*Attributes)
	for attributes := range stage.Attributess {
		_ = attributes
		for _, _attributesdirective := range attributes.Directive {
			stage.Attributes_Directive_reverseMap[_attributesdirective] = attributes
		}
	}
	stage.Attributes_MeasureStyle_reverseMap = make(map[*MeasureStyle]*Attributes)
	for attributes := range stage.Attributess {
		_ = attributes
		for _, _measurestyle := range attributes.MeasureStyle {
			stage.Attributes_MeasureStyle_reverseMap[_measurestyle] = attributes
		}
	}

	// Compute reverse map for named struct Key
	// insertion point per field
	stage.Key_KeyAccidental_reverseMap = make(map[*KeyAccidental]*Key)
	for key := range stage.Keys {
		_ = key
		for _, _keyaccidental := range key.KeyAccidental {
			stage.Key_KeyAccidental_reverseMap[_keyaccidental] = key
		}
	}
	stage.Key_KeyOctave_reverseMap = make(map[*KeyOctave]*Key)
	for key := range stage.Keys {
		_ = key
		for _, _keyoctave := range key.KeyOctave {
			stage.Key_KeyOctave_reverseMap[_keyoctave] = key
		}
	}

	// Compute reverse map for named struct StaffDetails
	// insertion point per field
	stage.StaffDetails_LineDetail_reverseMap = make(map[*LineDetail]*StaffDetails)
	for staffdetails := range stage.StaffDetailss {
		_ = staffdetails
		for _, _linedetail := range staffdetails.LineDetail {
			stage.StaffDetails_LineDetail_reverseMap[_linedetail] = staffdetails
		}
	}
	stage.StaffDetails_StaffTuning_reverseMap = make(map[*StaffTuning]*StaffDetails)
	for staffdetails := range stage.StaffDetailss {
		_ = staffdetails
		for _, _stafftuning := range staffdetails.StaffTuning {
			stage.StaffDetails_StaffTuning_reverseMap[_stafftuning] = staffdetails
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.Attributess)

	res = __gong__appendInstances(res, stage.AttributesDirectives)

	res = __gong__appendInstances(res, stage.BeatRepeats)

	res = __gong__appendInstances(res, stage.Cancels)

	res = __gong__appendInstances(res, stage.Clefs)

	res = __gong__appendInstances(res, stage.Doubles)

	res = __gong__appendInstances(res, stage.ForParts)

	res = __gong__appendInstances(res, stage.Interchangeables)

	res = __gong__appendInstances(res, stage.Keys)

	res = __gong__appendInstances(res, stage.KeyAccidentals)

	res = __gong__appendInstances(res, stage.KeyOctaves)

	res = __gong__appendInstances(res, stage.LineDetails)

	res = __gong__appendInstances(res, stage.MeasureRepeats)

	res = __gong__appendInstances(res, stage.MeasureStyles)

	res = __gong__appendInstances(res, stage.MultipleRests)

	res = __gong__appendInstances(res, stage.PartClefs)

	res = __gong__appendInstances(res, stage.PartSymbols)

	res = __gong__appendInstances(res, stage.PartTransposes)

	res = __gong__appendInstances(res, stage.Slashs)

	res = __gong__appendInstances(res, stage.StaffDetailss)

	res = __gong__appendInstances(res, stage.StaffSizes)

	res = __gong__appendInstances(res, stage.StaffTunings)

	res = __gong__appendInstances(res, stage.Times)

	res = __gong__appendInstances(res, stage.Transposes)

	return
}

// insertion point per named struct
func (attributes *Attributes) GongCopy() GongstructIF {
	newInstance := new(Attributes)
	attributes.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attributesdirective *AttributesDirective) GongCopy() GongstructIF {
	newInstance := new(AttributesDirective)
	attributesdirective.GongCopyBasicFields(newInstance)
	return newInstance
}

func (beatrepeat *BeatRepeat) GongCopy() GongstructIF {
	newInstance := new(BeatRepeat)
	beatrepeat.GongCopyBasicFields(newInstance)
	return newInstance
}

func (cancel *Cancel) GongCopy() GongstructIF {
	newInstance := new(Cancel)
	cancel.GongCopyBasicFields(newInstance)
	return newInstance
}

func (clef *Clef) GongCopy() GongstructIF {
	newInstance := new(Clef)
	clef.GongCopyBasicFields(newInstance)
	return newInstance
}

func (double *Double) GongCopy() GongstructIF {
	newInstance := new(Double)
	double.GongCopyBasicFields(newInstance)
	return newInstance
}

func (forpart *ForPart) GongCopy() GongstructIF {
	newInstance := new(ForPart)
	forpart.GongCopyBasicFields(newInstance)
	return newInstance
}

func (interchangeable *Interchangeable) GongCopy() GongstructIF {
	newInstance := new(Interchangeable)
	interchangeable.GongCopyBasicFields(newInstance)
	return newInstance
}

func (key *Key) GongCopy() GongstructIF {
	newInstance := new(Key)
	key.GongCopyBasicFields(newInstance)
	return newInstance
}

func (keyaccidental *KeyAccidental) GongCopy() GongstructIF {
	newInstance := new(KeyAccidental)
	keyaccidental.GongCopyBasicFields(newInstance)
	return newInstance
}

func (keyoctave *KeyOctave) GongCopy() GongstructIF {
	newInstance := new(KeyOctave)
	keyoctave.GongCopyBasicFields(newInstance)
	return newInstance
}

func (linedetail *LineDetail) GongCopy() GongstructIF {
	newInstance := new(LineDetail)
	linedetail.GongCopyBasicFields(newInstance)
	return newInstance
}

func (measurerepeat *MeasureRepeat) GongCopy() GongstructIF {
	newInstance := new(MeasureRepeat)
	measurerepeat.GongCopyBasicFields(newInstance)
	return newInstance
}

func (measurestyle *MeasureStyle) GongCopy() GongstructIF {
	newInstance := new(MeasureStyle)
	measurestyle.GongCopyBasicFields(newInstance)
	return newInstance
}

func (multiplerest *MultipleRest) GongCopy() GongstructIF {
	newInstance := new(MultipleRest)
	multiplerest.GongCopyBasicFields(newInstance)
	return newInstance
}

func (partclef *PartClef) GongCopy() GongstructIF {
	newInstance := new(PartClef)
	partclef.GongCopyBasicFields(newInstance)
	return newInstance
}

func (partsymbol *PartSymbol) GongCopy() GongstructIF {
	newInstance := new(PartSymbol)
	partsymbol.GongCopyBasicFields(newInstance)
	return newInstance
}

func (parttranspose *PartTranspose) GongCopy() GongstructIF {
	newInstance := new(PartTranspose)
	parttranspose.GongCopyBasicFields(newInstance)
	return newInstance
}

func (slash *Slash) GongCopy() GongstructIF {
	newInstance := new(Slash)
	slash.GongCopyBasicFields(newInstance)
	return newInstance
}

func (staffdetails *StaffDetails) GongCopy() GongstructIF {
	newInstance := new(StaffDetails)
	staffdetails.GongCopyBasicFields(newInstance)
	return newInstance
}

func (staffsize *StaffSize) GongCopy() GongstructIF {
	newInstance := new(StaffSize)
	staffsize.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stafftuning *StaffTuning) GongCopy() GongstructIF {
	newInstance := new(StaffTuning)
	stafftuning.GongCopyBasicFields(newInstance)
	return newInstance
}

func (time *Time) GongCopy() GongstructIF {
	newInstance := new(Time)
	time.GongCopyBasicFields(newInstance)
	return newInstance
}

func (transpose *Transpose) GongCopy() GongstructIF {
	newInstance := new(Transpose)
	transpose.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (attributes *Attributes) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, attributes)
}

func (attributesdirective *AttributesDirective) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, attributesdirective)
}

func (beatrepeat *BeatRepeat) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, beatrepeat)
}

func (cancel *Cancel) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, cancel)
}

func (clef *Clef) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, clef)
}

func (double *Double) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, double)
}

func (forpart *ForPart) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, forpart)
}

func (interchangeable *Interchangeable) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, interchangeable)
}

func (key *Key) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, key)
}

func (keyaccidental *KeyAccidental) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, keyaccidental)
}

func (keyoctave *KeyOctave) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, keyoctave)
}

func (linedetail *LineDetail) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, linedetail)
}

func (measurerepeat *MeasureRepeat) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, measurerepeat)
}

func (measurestyle *MeasureStyle) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, measurestyle)
}

func (multiplerest *MultipleRest) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, multiplerest)
}

func (partclef *PartClef) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, partclef)
}

func (partsymbol *PartSymbol) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, partsymbol)
}

func (parttranspose *PartTranspose) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, parttranspose)
}

func (slash *Slash) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, slash)
}

func (staffdetails *StaffDetails) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, staffdetails)
}

func (staffsize *StaffSize) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, staffsize)
}

func (stafftuning *StaffTuning) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stafftuning)
}

func (time *Time) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, time)
}

func (transpose *Transpose) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, transpose)
}


type GongstructDiffable[T any] interface {
	GongstructPtr
	GongMarshallIdentifier(stage *Stage) string
	GongMarshallUnstaging(stage *Stage) string
	GongMarshallAllFields(stage *Stage) (string, string)
	GongReconstructPointersFromInstances(stage *Stage)
	GongDiff(stage *Stage, other T) []string
}

func computeCommitsForType[T GongstructDiffable[T]](
	stage *Stage,
	stagedInstances map[T]struct{},
	stagedOrder map[T]uint,
	referenceInstances map[T]T,
	referenceOrder *map[T]uint,
	instancesMap map[T]T,
	newInstancesSlice *[]string,
	fieldsEditSlice *[]string,
	deletedInstancesSlice *[]string,
	newInstancesReverseSlice *[]string,
	fieldsEditReverseSlice *[]string,
	deletedInstancesReverseSlice *[]string,
	lenNewInstances *int,
	lenDeletedInstances *int,
	lenModifiedInstances *int,
) {
	var newInstances []T
	var deletedInstances []T

	// parse all staged instances and check if they have a reference
	for instance := range stagedInstances {
		if ref, ok := referenceInstances[instance]; !ok {
			newInstances = append(newInstances, instance)
			*newInstancesSlice = append(*newInstancesSlice, instance.GongMarshallIdentifier(stage))
			if *referenceOrder == nil {
				*referenceOrder = make(map[T]uint)
			}
			(*referenceOrder)[instance] = stagedOrder[instance]
			*newInstancesReverseSlice = append(*newInstancesReverseSlice, instance.GongMarshallUnstaging(stage))
			fieldInitializers, pointersInitializations := instance.GongMarshallAllFields(stage)
			*fieldsEditSlice = append(*fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stagedOrder[ref] = stagedOrder[instance]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := instance.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, instance)
			if len(diffs) > 0 {
				var fieldsEdit string
				if instance.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", instance.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				*fieldsEditSlice = append(*fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, reverseDiff)
				}
				*lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range referenceInstances {
		instance := instancesMap[ref] // get the instance corresponding to the reference
		if _, ok := stagedInstances[instance]; !ok { // if the instance is not staged anymore, it means it has been unstaged
			deletedInstances = append(deletedInstances, ref)
			*deletedInstancesSlice = append(*deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			*deletedInstancesReverseSlice = append(*deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	*lenNewInstances += len(newInstances)
	*lenDeletedInstances += len(deletedInstances)
}

func (stage *Stage) ComputeForwardAndBackwardCommits() {
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
	computeCommitsForType(
		stage,
		stage.Attributess,
		stage.Attributes_stagedOrder,
		stage.Attributess_reference,
		&stage.Attributess_referenceOrder,
		stage.Attributess_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AttributesDirectives,
		stage.AttributesDirective_stagedOrder,
		stage.AttributesDirectives_reference,
		&stage.AttributesDirectives_referenceOrder,
		stage.AttributesDirectives_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.BeatRepeats,
		stage.BeatRepeat_stagedOrder,
		stage.BeatRepeats_reference,
		&stage.BeatRepeats_referenceOrder,
		stage.BeatRepeats_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Cancels,
		stage.Cancel_stagedOrder,
		stage.Cancels_reference,
		&stage.Cancels_referenceOrder,
		stage.Cancels_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Clefs,
		stage.Clef_stagedOrder,
		stage.Clefs_reference,
		&stage.Clefs_referenceOrder,
		stage.Clefs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Doubles,
		stage.Double_stagedOrder,
		stage.Doubles_reference,
		&stage.Doubles_referenceOrder,
		stage.Doubles_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ForParts,
		stage.ForPart_stagedOrder,
		stage.ForParts_reference,
		&stage.ForParts_referenceOrder,
		stage.ForParts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Interchangeables,
		stage.Interchangeable_stagedOrder,
		stage.Interchangeables_reference,
		&stage.Interchangeables_referenceOrder,
		stage.Interchangeables_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Keys,
		stage.Key_stagedOrder,
		stage.Keys_reference,
		&stage.Keys_referenceOrder,
		stage.Keys_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.KeyAccidentals,
		stage.KeyAccidental_stagedOrder,
		stage.KeyAccidentals_reference,
		&stage.KeyAccidentals_referenceOrder,
		stage.KeyAccidentals_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.KeyOctaves,
		stage.KeyOctave_stagedOrder,
		stage.KeyOctaves_reference,
		&stage.KeyOctaves_referenceOrder,
		stage.KeyOctaves_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.LineDetails,
		stage.LineDetail_stagedOrder,
		stage.LineDetails_reference,
		&stage.LineDetails_referenceOrder,
		stage.LineDetails_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MeasureRepeats,
		stage.MeasureRepeat_stagedOrder,
		stage.MeasureRepeats_reference,
		&stage.MeasureRepeats_referenceOrder,
		stage.MeasureRepeats_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MeasureStyles,
		stage.MeasureStyle_stagedOrder,
		stage.MeasureStyles_reference,
		&stage.MeasureStyles_referenceOrder,
		stage.MeasureStyles_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MultipleRests,
		stage.MultipleRest_stagedOrder,
		stage.MultipleRests_reference,
		&stage.MultipleRests_referenceOrder,
		stage.MultipleRests_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.PartClefs,
		stage.PartClef_stagedOrder,
		stage.PartClefs_reference,
		&stage.PartClefs_referenceOrder,
		stage.PartClefs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.PartSymbols,
		stage.PartSymbol_stagedOrder,
		stage.PartSymbols_reference,
		&stage.PartSymbols_referenceOrder,
		stage.PartSymbols_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.PartTransposes,
		stage.PartTranspose_stagedOrder,
		stage.PartTransposes_reference,
		&stage.PartTransposes_referenceOrder,
		stage.PartTransposes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Slashs,
		stage.Slash_stagedOrder,
		stage.Slashs_reference,
		&stage.Slashs_referenceOrder,
		stage.Slashs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StaffDetailss,
		stage.StaffDetails_stagedOrder,
		stage.StaffDetailss_reference,
		&stage.StaffDetailss_referenceOrder,
		stage.StaffDetailss_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StaffSizes,
		stage.StaffSize_stagedOrder,
		stage.StaffSizes_reference,
		&stage.StaffSizes_referenceOrder,
		stage.StaffSizes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StaffTunings,
		stage.StaffTuning_stagedOrder,
		stage.StaffTunings_reference,
		&stage.StaffTunings_referenceOrder,
		stage.StaffTunings_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Times,
		stage.Time_stagedOrder,
		stage.Times_reference,
		&stage.Times_referenceOrder,
		stage.Times_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Transposes,
		stage.Transpose_stagedOrder,
		stage.Transposes_reference,
		&stage.Transposes_referenceOrder,
		stage.Transposes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)

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
		stage.AppendForwardCommit(forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.AppendBackwardCommit(backwardCommit)
		stage.SetModified(true)
	} else {
		stage.SetModified(false)
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	__gong__computeReferencePass1(stage, stage.Attributess, &stage.Attributess_reference, &stage.Attributess_referenceOrder, &stage.Attributess_instance)

	__gong__computeReferencePass1(stage, stage.AttributesDirectives, &stage.AttributesDirectives_reference, &stage.AttributesDirectives_referenceOrder, &stage.AttributesDirectives_instance)

	__gong__computeReferencePass1(stage, stage.BeatRepeats, &stage.BeatRepeats_reference, &stage.BeatRepeats_referenceOrder, &stage.BeatRepeats_instance)

	__gong__computeReferencePass1(stage, stage.Cancels, &stage.Cancels_reference, &stage.Cancels_referenceOrder, &stage.Cancels_instance)

	__gong__computeReferencePass1(stage, stage.Clefs, &stage.Clefs_reference, &stage.Clefs_referenceOrder, &stage.Clefs_instance)

	__gong__computeReferencePass1(stage, stage.Doubles, &stage.Doubles_reference, &stage.Doubles_referenceOrder, &stage.Doubles_instance)

	__gong__computeReferencePass1(stage, stage.ForParts, &stage.ForParts_reference, &stage.ForParts_referenceOrder, &stage.ForParts_instance)

	__gong__computeReferencePass1(stage, stage.Interchangeables, &stage.Interchangeables_reference, &stage.Interchangeables_referenceOrder, &stage.Interchangeables_instance)

	__gong__computeReferencePass1(stage, stage.Keys, &stage.Keys_reference, &stage.Keys_referenceOrder, &stage.Keys_instance)

	__gong__computeReferencePass1(stage, stage.KeyAccidentals, &stage.KeyAccidentals_reference, &stage.KeyAccidentals_referenceOrder, &stage.KeyAccidentals_instance)

	__gong__computeReferencePass1(stage, stage.KeyOctaves, &stage.KeyOctaves_reference, &stage.KeyOctaves_referenceOrder, &stage.KeyOctaves_instance)

	__gong__computeReferencePass1(stage, stage.LineDetails, &stage.LineDetails_reference, &stage.LineDetails_referenceOrder, &stage.LineDetails_instance)

	__gong__computeReferencePass1(stage, stage.MeasureRepeats, &stage.MeasureRepeats_reference, &stage.MeasureRepeats_referenceOrder, &stage.MeasureRepeats_instance)

	__gong__computeReferencePass1(stage, stage.MeasureStyles, &stage.MeasureStyles_reference, &stage.MeasureStyles_referenceOrder, &stage.MeasureStyles_instance)

	__gong__computeReferencePass1(stage, stage.MultipleRests, &stage.MultipleRests_reference, &stage.MultipleRests_referenceOrder, &stage.MultipleRests_instance)

	__gong__computeReferencePass1(stage, stage.PartClefs, &stage.PartClefs_reference, &stage.PartClefs_referenceOrder, &stage.PartClefs_instance)

	__gong__computeReferencePass1(stage, stage.PartSymbols, &stage.PartSymbols_reference, &stage.PartSymbols_referenceOrder, &stage.PartSymbols_instance)

	__gong__computeReferencePass1(stage, stage.PartTransposes, &stage.PartTransposes_reference, &stage.PartTransposes_referenceOrder, &stage.PartTransposes_instance)

	__gong__computeReferencePass1(stage, stage.Slashs, &stage.Slashs_reference, &stage.Slashs_referenceOrder, &stage.Slashs_instance)

	__gong__computeReferencePass1(stage, stage.StaffDetailss, &stage.StaffDetailss_reference, &stage.StaffDetailss_referenceOrder, &stage.StaffDetailss_instance)

	__gong__computeReferencePass1(stage, stage.StaffSizes, &stage.StaffSizes_reference, &stage.StaffSizes_referenceOrder, &stage.StaffSizes_instance)

	__gong__computeReferencePass1(stage, stage.StaffTunings, &stage.StaffTunings_reference, &stage.StaffTunings_referenceOrder, &stage.StaffTunings_instance)

	__gong__computeReferencePass1(stage, stage.Times, &stage.Times_reference, &stage.Times_referenceOrder, &stage.Times_instance)

	__gong__computeReferencePass1(stage, stage.Transposes, &stage.Transposes_reference, &stage.Transposes_referenceOrder, &stage.Transposes_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.Attributess, stage.Attributess_reference, stage)

	__gong__computeReferencePass2(stage.AttributesDirectives, stage.AttributesDirectives_reference, stage)

	__gong__computeReferencePass2(stage.BeatRepeats, stage.BeatRepeats_reference, stage)

	__gong__computeReferencePass2(stage.Cancels, stage.Cancels_reference, stage)

	__gong__computeReferencePass2(stage.Clefs, stage.Clefs_reference, stage)

	__gong__computeReferencePass2(stage.Doubles, stage.Doubles_reference, stage)

	__gong__computeReferencePass2(stage.ForParts, stage.ForParts_reference, stage)

	__gong__computeReferencePass2(stage.Interchangeables, stage.Interchangeables_reference, stage)

	__gong__computeReferencePass2(stage.Keys, stage.Keys_reference, stage)

	__gong__computeReferencePass2(stage.KeyAccidentals, stage.KeyAccidentals_reference, stage)

	__gong__computeReferencePass2(stage.KeyOctaves, stage.KeyOctaves_reference, stage)

	__gong__computeReferencePass2(stage.LineDetails, stage.LineDetails_reference, stage)

	__gong__computeReferencePass2(stage.MeasureRepeats, stage.MeasureRepeats_reference, stage)

	__gong__computeReferencePass2(stage.MeasureStyles, stage.MeasureStyles_reference, stage)

	__gong__computeReferencePass2(stage.MultipleRests, stage.MultipleRests_reference, stage)

	__gong__computeReferencePass2(stage.PartClefs, stage.PartClefs_reference, stage)

	__gong__computeReferencePass2(stage.PartSymbols, stage.PartSymbols_reference, stage)

	__gong__computeReferencePass2(stage.PartTransposes, stage.PartTransposes_reference, stage)

	__gong__computeReferencePass2(stage.Slashs, stage.Slashs_reference, stage)

	__gong__computeReferencePass2(stage.StaffDetailss, stage.StaffDetailss_reference, stage)

	__gong__computeReferencePass2(stage.StaffSizes, stage.StaffSizes_reference, stage)

	__gong__computeReferencePass2(stage.StaffTunings, stage.StaffTunings_reference, stage)

	__gong__computeReferencePass2(stage.Times, stage.Times_reference, stage)

	__gong__computeReferencePass2(stage.Transposes, stage.Transposes_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (attributes *Attributes) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Attributes_stagedOrder, stage.Attributess_referenceOrder, attributes, "Attributes")
}

func (attributesdirective *AttributesDirective) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AttributesDirective_stagedOrder, stage.AttributesDirectives_referenceOrder, attributesdirective, "AttributesDirective")
}

func (beatrepeat *BeatRepeat) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.BeatRepeat_stagedOrder, stage.BeatRepeats_referenceOrder, beatrepeat, "BeatRepeat")
}

func (cancel *Cancel) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Cancel_stagedOrder, stage.Cancels_referenceOrder, cancel, "Cancel")
}

func (clef *Clef) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Clef_stagedOrder, stage.Clefs_referenceOrder, clef, "Clef")
}

func (double *Double) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Double_stagedOrder, stage.Doubles_referenceOrder, double, "Double")
}

func (forpart *ForPart) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ForPart_stagedOrder, stage.ForParts_referenceOrder, forpart, "ForPart")
}

func (interchangeable *Interchangeable) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Interchangeable_stagedOrder, stage.Interchangeables_referenceOrder, interchangeable, "Interchangeable")
}

func (key *Key) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Key_stagedOrder, stage.Keys_referenceOrder, key, "Key")
}

func (keyaccidental *KeyAccidental) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.KeyAccidental_stagedOrder, stage.KeyAccidentals_referenceOrder, keyaccidental, "KeyAccidental")
}

func (keyoctave *KeyOctave) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.KeyOctave_stagedOrder, stage.KeyOctaves_referenceOrder, keyoctave, "KeyOctave")
}

func (linedetail *LineDetail) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LineDetail_stagedOrder, stage.LineDetails_referenceOrder, linedetail, "LineDetail")
}

func (measurerepeat *MeasureRepeat) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MeasureRepeat_stagedOrder, stage.MeasureRepeats_referenceOrder, measurerepeat, "MeasureRepeat")
}

func (measurestyle *MeasureStyle) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MeasureStyle_stagedOrder, stage.MeasureStyles_referenceOrder, measurestyle, "MeasureStyle")
}

func (multiplerest *MultipleRest) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MultipleRest_stagedOrder, stage.MultipleRests_referenceOrder, multiplerest, "MultipleRest")
}

func (partclef *PartClef) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PartClef_stagedOrder, stage.PartClefs_referenceOrder, partclef, "PartClef")
}

func (partsymbol *PartSymbol) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PartSymbol_stagedOrder, stage.PartSymbols_referenceOrder, partsymbol, "PartSymbol")
}

func (parttranspose *PartTranspose) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PartTranspose_stagedOrder, stage.PartTransposes_referenceOrder, parttranspose, "PartTranspose")
}

func (slash *Slash) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Slash_stagedOrder, stage.Slashs_referenceOrder, slash, "Slash")
}

func (staffdetails *StaffDetails) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StaffDetails_stagedOrder, stage.StaffDetailss_referenceOrder, staffdetails, "StaffDetails")
}

func (staffsize *StaffSize) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StaffSize_stagedOrder, stage.StaffSizes_referenceOrder, staffsize, "StaffSize")
}

func (stafftuning *StaffTuning) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StaffTuning_stagedOrder, stage.StaffTunings_referenceOrder, stafftuning, "StaffTuning")
}

func (time *Time) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Time_stagedOrder, stage.Times_referenceOrder, time, "Time")
}

func (transpose *Transpose) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Transpose_stagedOrder, stage.Transposes_referenceOrder, transpose, "Transpose")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (attributes *Attributes) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(attributes, attributes.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attributes *Attributes) GongGetReferenceIdentifier(stage *Stage) string {
	return attributes.GongGetIdentifier(stage)
}

func (attributesdirective *AttributesDirective) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(attributesdirective, attributesdirective.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attributesdirective *AttributesDirective) GongGetReferenceIdentifier(stage *Stage) string {
	return attributesdirective.GongGetIdentifier(stage)
}

func (beatrepeat *BeatRepeat) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(beatrepeat, beatrepeat.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (beatrepeat *BeatRepeat) GongGetReferenceIdentifier(stage *Stage) string {
	return beatrepeat.GongGetIdentifier(stage)
}

func (cancel *Cancel) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(cancel, cancel.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (cancel *Cancel) GongGetReferenceIdentifier(stage *Stage) string {
	return cancel.GongGetIdentifier(stage)
}

func (clef *Clef) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(clef, clef.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (clef *Clef) GongGetReferenceIdentifier(stage *Stage) string {
	return clef.GongGetIdentifier(stage)
}

func (double *Double) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(double, double.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (double *Double) GongGetReferenceIdentifier(stage *Stage) string {
	return double.GongGetIdentifier(stage)
}

func (forpart *ForPart) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(forpart, forpart.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (forpart *ForPart) GongGetReferenceIdentifier(stage *Stage) string {
	return forpart.GongGetIdentifier(stage)
}

func (interchangeable *Interchangeable) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(interchangeable, interchangeable.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (interchangeable *Interchangeable) GongGetReferenceIdentifier(stage *Stage) string {
	return interchangeable.GongGetIdentifier(stage)
}

func (key *Key) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(key, key.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (key *Key) GongGetReferenceIdentifier(stage *Stage) string {
	return key.GongGetIdentifier(stage)
}

func (keyaccidental *KeyAccidental) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(keyaccidental, keyaccidental.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (keyaccidental *KeyAccidental) GongGetReferenceIdentifier(stage *Stage) string {
	return keyaccidental.GongGetIdentifier(stage)
}

func (keyoctave *KeyOctave) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(keyoctave, keyoctave.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (keyoctave *KeyOctave) GongGetReferenceIdentifier(stage *Stage) string {
	return keyoctave.GongGetIdentifier(stage)
}

func (linedetail *LineDetail) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(linedetail, linedetail.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (linedetail *LineDetail) GongGetReferenceIdentifier(stage *Stage) string {
	return linedetail.GongGetIdentifier(stage)
}

func (measurerepeat *MeasureRepeat) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(measurerepeat, measurerepeat.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (measurerepeat *MeasureRepeat) GongGetReferenceIdentifier(stage *Stage) string {
	return measurerepeat.GongGetIdentifier(stage)
}

func (measurestyle *MeasureStyle) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(measurestyle, measurestyle.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (measurestyle *MeasureStyle) GongGetReferenceIdentifier(stage *Stage) string {
	return measurestyle.GongGetIdentifier(stage)
}

func (multiplerest *MultipleRest) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(multiplerest, multiplerest.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (multiplerest *MultipleRest) GongGetReferenceIdentifier(stage *Stage) string {
	return multiplerest.GongGetIdentifier(stage)
}

func (partclef *PartClef) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(partclef, partclef.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partclef *PartClef) GongGetReferenceIdentifier(stage *Stage) string {
	return partclef.GongGetIdentifier(stage)
}

func (partsymbol *PartSymbol) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(partsymbol, partsymbol.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partsymbol *PartSymbol) GongGetReferenceIdentifier(stage *Stage) string {
	return partsymbol.GongGetIdentifier(stage)
}

func (parttranspose *PartTranspose) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(parttranspose, parttranspose.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (parttranspose *PartTranspose) GongGetReferenceIdentifier(stage *Stage) string {
	return parttranspose.GongGetIdentifier(stage)
}

func (slash *Slash) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(slash, slash.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (slash *Slash) GongGetReferenceIdentifier(stage *Stage) string {
	return slash.GongGetIdentifier(stage)
}

func (staffdetails *StaffDetails) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(staffdetails, staffdetails.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (staffdetails *StaffDetails) GongGetReferenceIdentifier(stage *Stage) string {
	return staffdetails.GongGetIdentifier(stage)
}

func (staffsize *StaffSize) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(staffsize, staffsize.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (staffsize *StaffSize) GongGetReferenceIdentifier(stage *Stage) string {
	return staffsize.GongGetIdentifier(stage)
}

func (stafftuning *StaffTuning) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stafftuning, stafftuning.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stafftuning *StaffTuning) GongGetReferenceIdentifier(stage *Stage) string {
	return stafftuning.GongGetIdentifier(stage)
}

func (time *Time) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(time, time.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (time *Time) GongGetReferenceIdentifier(stage *Stage) string {
	return time.GongGetIdentifier(stage)
}

func (transpose *Transpose) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(transpose, transpose.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (transpose *Transpose) GongGetReferenceIdentifier(stage *Stage) string {
	return transpose.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (attributes *Attributes) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(attributes.GongGetIdentifier(stage), "Attributes", attributes.Name)
}

func (attributesdirective *AttributesDirective) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(attributesdirective.GongGetIdentifier(stage), "AttributesDirective", attributesdirective.Name)
}

func (beatrepeat *BeatRepeat) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(beatrepeat.GongGetIdentifier(stage), "BeatRepeat", beatrepeat.Name)
}

func (cancel *Cancel) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(cancel.GongGetIdentifier(stage), "Cancel", cancel.Name)
}

func (clef *Clef) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(clef.GongGetIdentifier(stage), "Clef", clef.Name)
}

func (double *Double) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(double.GongGetIdentifier(stage), "Double", double.Name)
}

func (forpart *ForPart) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(forpart.GongGetIdentifier(stage), "ForPart", forpart.Name)
}

func (interchangeable *Interchangeable) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(interchangeable.GongGetIdentifier(stage), "Interchangeable", interchangeable.Name)
}

func (key *Key) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(key.GongGetIdentifier(stage), "Key", key.Name)
}

func (keyaccidental *KeyAccidental) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(keyaccidental.GongGetIdentifier(stage), "KeyAccidental", keyaccidental.Name)
}

func (keyoctave *KeyOctave) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(keyoctave.GongGetIdentifier(stage), "KeyOctave", keyoctave.Name)
}

func (linedetail *LineDetail) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(linedetail.GongGetIdentifier(stage), "LineDetail", linedetail.Name)
}

func (measurerepeat *MeasureRepeat) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(measurerepeat.GongGetIdentifier(stage), "MeasureRepeat", measurerepeat.Name)
}

func (measurestyle *MeasureStyle) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(measurestyle.GongGetIdentifier(stage), "MeasureStyle", measurestyle.Name)
}

func (multiplerest *MultipleRest) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(multiplerest.GongGetIdentifier(stage), "MultipleRest", multiplerest.Name)
}

func (partclef *PartClef) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(partclef.GongGetIdentifier(stage), "PartClef", partclef.Name)
}

func (partsymbol *PartSymbol) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(partsymbol.GongGetIdentifier(stage), "PartSymbol", partsymbol.Name)
}

func (parttranspose *PartTranspose) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(parttranspose.GongGetIdentifier(stage), "PartTranspose", parttranspose.Name)
}

func (slash *Slash) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(slash.GongGetIdentifier(stage), "Slash", slash.Name)
}

func (staffdetails *StaffDetails) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(staffdetails.GongGetIdentifier(stage), "StaffDetails", staffdetails.Name)
}

func (staffsize *StaffSize) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(staffsize.GongGetIdentifier(stage), "StaffSize", staffsize.Name)
}

func (stafftuning *StaffTuning) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stafftuning.GongGetIdentifier(stage), "StaffTuning", stafftuning.Name)
}

func (time *Time) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(time.GongGetIdentifier(stage), "Time", time.Name)
}

func (transpose *Transpose) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(transpose.GongGetIdentifier(stage), "Transpose", transpose.Name)
}

// insertion point for unstaging
func (attributes *Attributes) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(attributes.GongGetReferenceIdentifier(stage))
}

func (attributesdirective *AttributesDirective) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(attributesdirective.GongGetReferenceIdentifier(stage))
}

func (beatrepeat *BeatRepeat) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(beatrepeat.GongGetReferenceIdentifier(stage))
}

func (cancel *Cancel) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(cancel.GongGetReferenceIdentifier(stage))
}

func (clef *Clef) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(clef.GongGetReferenceIdentifier(stage))
}

func (double *Double) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(double.GongGetReferenceIdentifier(stage))
}

func (forpart *ForPart) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(forpart.GongGetReferenceIdentifier(stage))
}

func (interchangeable *Interchangeable) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(interchangeable.GongGetReferenceIdentifier(stage))
}

func (key *Key) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(key.GongGetReferenceIdentifier(stage))
}

func (keyaccidental *KeyAccidental) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(keyaccidental.GongGetReferenceIdentifier(stage))
}

func (keyoctave *KeyOctave) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(keyoctave.GongGetReferenceIdentifier(stage))
}

func (linedetail *LineDetail) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(linedetail.GongGetReferenceIdentifier(stage))
}

func (measurerepeat *MeasureRepeat) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(measurerepeat.GongGetReferenceIdentifier(stage))
}

func (measurestyle *MeasureStyle) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(measurestyle.GongGetReferenceIdentifier(stage))
}

func (multiplerest *MultipleRest) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(multiplerest.GongGetReferenceIdentifier(stage))
}

func (partclef *PartClef) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(partclef.GongGetReferenceIdentifier(stage))
}

func (partsymbol *PartSymbol) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(partsymbol.GongGetReferenceIdentifier(stage))
}

func (parttranspose *PartTranspose) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(parttranspose.GongGetReferenceIdentifier(stage))
}

func (slash *Slash) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(slash.GongGetReferenceIdentifier(stage))
}

func (staffdetails *StaffDetails) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(staffdetails.GongGetReferenceIdentifier(stage))
}

func (staffsize *StaffSize) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(staffsize.GongGetReferenceIdentifier(stage))
}

func (stafftuning *StaffTuning) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stafftuning.GongGetReferenceIdentifier(stage))
}

func (time *Time) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(time.GongGetReferenceIdentifier(stage))
}

func (transpose *Transpose) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(transpose.GongGetReferenceIdentifier(stage))
}

func GongIntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += GongIntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GongGenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GongGenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
	// 1. Create a deterministic hash from the inputs using SHA-256
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(seedStr))

	// Write the integer to the hash (using BigEndian to ensure consistency across architectures)
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, seedInt)
	h.Write(intBytes)

	// 2. Extract the first 16 bytes from our resulting hash
	hashBytes := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, hashBytes[:16])

	// 3. Set the Version to 4 (0100 in binary)
	// We take the 7th byte, clear the top 4 bits with & 0x0f, and set the top bits to 0100 with | 0x40
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// 4. Set the Variant to RFC4122 (10 in binary)
	// We take the 9th byte, clear the top 2 bits with & 0x3f, and set the top bits to 10 with | 0x80
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format and return the byte array as a standard UUID string
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

func __gong__appendInstances[T interface {
	comparable
	GongstructIF
}](res []GongstructIF, m map[T]struct{}) []GongstructIF {
	for instance := range m {
		res = append(res, instance)
	}
	return res
}

func __gong__getUUID(stage *Stage, instance GongstructIF) string {
	if __gong__, ok := any(instance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}
	return GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instance), uint64(stage.GetOrder(instance)))
}

func __gong__computeReferencePass1[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	staged map[T]struct{},
	ref *map[T]T,
	refOrder *map[T]uint,
	inst *map[T]T,
) {
	*ref = make(map[T]T, len(staged))
	*refOrder = make(map[T]uint, len(staged))
	*inst = make(map[T]T, len(staged))
	for instance := range staged {
		_copy := instance.GongCopy().(T)
		(*ref)[instance] = _copy
		(*inst)[_copy] = instance
		(*refOrder)[_copy] = instance.GongGetOrder(stage)
	}
}

func __gong__computeReferencePass2[T interface {
	comparable
	GongstructIF
	GongReconstructPointersFromReferences(*Stage, T)
}](staged map[T]struct{}, reference map[T]T, stage *Stage) {
	for instance := range staged {
		reference[instance].GongReconstructPointersFromReferences(stage, instance)
	}
}

func __gong__getOrder[T comparable](stagedOrder, refOrder map[T]uint, instance T, typeName string) uint {
	if order, ok := stagedOrder[instance]; ok {
		return order
	}
	if order, ok := refOrder[instance]; ok {
		return order
	}
	log.Printf("instance %p of type %s was not staged and does not have a reference order", any(instance), typeName)
	return 0
}

func __gong__formatIdentifier(s GongstructIF, order uint) string {
	return fmt.Sprintf("__%s__%08d_", s.GongGetGongstructName(), order)
}

func __gong__marshallIdentifier(identifier, structName, name string) string {
	decl := strings.ReplaceAll(GongIdentifiersDecls, "{{Identifier}}", identifier)
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", structName)
	return strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(name))
}

func __gong__marshallUnstaging(identifier string) string {
	return strings.ReplaceAll(GongUnstageStmt, "{{Identifier}}", identifier)
}

// end of template
