// generated code - do not edit
package aixmdatatypes

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
	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.CodeColourTypes)

	res = __gong__appendInstances(res, stage.CodeDayTypes)

	res = __gong__appendInstances(res, stage.CodeGroundLightingTypes)

	res = __gong__appendInstances(res, stage.CodeLightIntensityTypes)

	res = __gong__appendInstances(res, stage.CodeLightSourceTypes)

	res = __gong__appendInstances(res, stage.CodeNotePurposeTypes)

	res = __gong__appendInstances(res, stage.CodeObstacleAreaTypes)

	res = __gong__appendInstances(res, stage.CodeObstacleAssessmentSurfaceTypes)

	res = __gong__appendInstances(res, stage.CodeStatusConstructionTypes)

	res = __gong__appendInstances(res, stage.CodeStatusOperationsTypes)

	res = __gong__appendInstances(res, stage.CodeTimeEventCombinationTypes)

	res = __gong__appendInstances(res, stage.CodeTimeEventTypes)

	res = __gong__appendInstances(res, stage.CodeTimeReferenceTypes)

	res = __gong__appendInstances(res, stage.CodeVerticalDatumTypes)

	res = __gong__appendInstances(res, stage.CodeVerticalStructureMarkingTypes)

	res = __gong__appendInstances(res, stage.CodeVerticalStructureMaterialTypes)

	res = __gong__appendInstances(res, stage.CodeVerticalStructureTypes)

	res = __gong__appendInstances(res, stage.CodeYesNoTypes)

	res = __gong__appendInstances(res, stage.DateMonthDayTypes)

	res = __gong__appendInstances(res, stage.TextDesignatorTypes)

	res = __gong__appendInstances(res, stage.TextNameTypes)

	res = __gong__appendInstances(res, stage.TextNoteTypes)

	res = __gong__appendInstances(res, stage.TextPropertyNameTypes)

	res = __gong__appendInstances(res, stage.TimeTypes)

	res = __gong__appendInstances(res, stage.ValDistanceSignedTypes)

	res = __gong__appendInstances(res, stage.ValDistanceTypes)

	res = __gong__appendInstances(res, stage.ValDistanceVerticalTypes)

	res = __gong__appendInstances(res, stage.ValDurationTypes)

	res = __gong__appendInstances(res, stage.ValLightIntensityTypes)

	return
}

// insertion point per named struct
func (codecolourtype *CodeColourType) GongCopy() GongstructIF {
	newInstance := new(CodeColourType)
	codecolourtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codedaytype *CodeDayType) GongCopy() GongstructIF {
	newInstance := new(CodeDayType)
	codedaytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codegroundlightingtype *CodeGroundLightingType) GongCopy() GongstructIF {
	newInstance := new(CodeGroundLightingType)
	codegroundlightingtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codelightintensitytype *CodeLightIntensityType) GongCopy() GongstructIF {
	newInstance := new(CodeLightIntensityType)
	codelightintensitytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codelightsourcetype *CodeLightSourceType) GongCopy() GongstructIF {
	newInstance := new(CodeLightSourceType)
	codelightsourcetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codenotepurposetype *CodeNotePurposeType) GongCopy() GongstructIF {
	newInstance := new(CodeNotePurposeType)
	codenotepurposetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeobstacleareatype *CodeObstacleAreaType) GongCopy() GongstructIF {
	newInstance := new(CodeObstacleAreaType)
	codeobstacleareatype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongCopy() GongstructIF {
	newInstance := new(CodeObstacleAssessmentSurfaceType)
	codeobstacleassessmentsurfacetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codestatusconstructiontype *CodeStatusConstructionType) GongCopy() GongstructIF {
	newInstance := new(CodeStatusConstructionType)
	codestatusconstructiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codestatusoperationstype *CodeStatusOperationsType) GongCopy() GongstructIF {
	newInstance := new(CodeStatusOperationsType)
	codestatusoperationstype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongCopy() GongstructIF {
	newInstance := new(CodeTimeEventCombinationType)
	codetimeeventcombinationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codetimeeventtype *CodeTimeEventType) GongCopy() GongstructIF {
	newInstance := new(CodeTimeEventType)
	codetimeeventtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codetimereferencetype *CodeTimeReferenceType) GongCopy() GongstructIF {
	newInstance := new(CodeTimeReferenceType)
	codetimereferencetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongCopy() GongstructIF {
	newInstance := new(CodeVerticalDatumType)
	codeverticaldatumtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongCopy() GongstructIF {
	newInstance := new(CodeVerticalStructureMarkingType)
	codeverticalstructuremarkingtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongCopy() GongstructIF {
	newInstance := new(CodeVerticalStructureMaterialType)
	codeverticalstructurematerialtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeverticalstructuretype *CodeVerticalStructureType) GongCopy() GongstructIF {
	newInstance := new(CodeVerticalStructureType)
	codeverticalstructuretype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeyesnotype *CodeYesNoType) GongCopy() GongstructIF {
	newInstance := new(CodeYesNoType)
	codeyesnotype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (datemonthdaytype *DateMonthDayType) GongCopy() GongstructIF {
	newInstance := new(DateMonthDayType)
	datemonthdaytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (textdesignatortype *TextDesignatorType) GongCopy() GongstructIF {
	newInstance := new(TextDesignatorType)
	textdesignatortype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (textnametype *TextNameType) GongCopy() GongstructIF {
	newInstance := new(TextNameType)
	textnametype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (textnotetype *TextNoteType) GongCopy() GongstructIF {
	newInstance := new(TextNoteType)
	textnotetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (textpropertynametype *TextPropertyNameType) GongCopy() GongstructIF {
	newInstance := new(TextPropertyNameType)
	textpropertynametype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (timetype *TimeType) GongCopy() GongstructIF {
	newInstance := new(TimeType)
	timetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (valdistancesignedtype *ValDistanceSignedType) GongCopy() GongstructIF {
	newInstance := new(ValDistanceSignedType)
	valdistancesignedtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (valdistancetype *ValDistanceType) GongCopy() GongstructIF {
	newInstance := new(ValDistanceType)
	valdistancetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (valdistanceverticaltype *ValDistanceVerticalType) GongCopy() GongstructIF {
	newInstance := new(ValDistanceVerticalType)
	valdistanceverticaltype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (valdurationtype *ValDurationType) GongCopy() GongstructIF {
	newInstance := new(ValDurationType)
	valdurationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (vallightintensitytype *ValLightIntensityType) GongCopy() GongstructIF {
	newInstance := new(ValLightIntensityType)
	vallightintensitytype.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (codecolourtype *CodeColourType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codecolourtype)
}

func (codedaytype *CodeDayType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codedaytype)
}

func (codegroundlightingtype *CodeGroundLightingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codegroundlightingtype)
}

func (codelightintensitytype *CodeLightIntensityType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codelightintensitytype)
}

func (codelightsourcetype *CodeLightSourceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codelightsourcetype)
}

func (codenotepurposetype *CodeNotePurposeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codenotepurposetype)
}

func (codeobstacleareatype *CodeObstacleAreaType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeobstacleareatype)
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeobstacleassessmentsurfacetype)
}

func (codestatusconstructiontype *CodeStatusConstructionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codestatusconstructiontype)
}

func (codestatusoperationstype *CodeStatusOperationsType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codestatusoperationstype)
}

func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codetimeeventcombinationtype)
}

func (codetimeeventtype *CodeTimeEventType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codetimeeventtype)
}

func (codetimereferencetype *CodeTimeReferenceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codetimereferencetype)
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeverticaldatumtype)
}

func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeverticalstructuremarkingtype)
}

func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeverticalstructurematerialtype)
}

func (codeverticalstructuretype *CodeVerticalStructureType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeverticalstructuretype)
}

func (codeyesnotype *CodeYesNoType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeyesnotype)
}

func (datemonthdaytype *DateMonthDayType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, datemonthdaytype)
}

func (textdesignatortype *TextDesignatorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, textdesignatortype)
}

func (textnametype *TextNameType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, textnametype)
}

func (textnotetype *TextNoteType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, textnotetype)
}

func (textpropertynametype *TextPropertyNameType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, textpropertynametype)
}

func (timetype *TimeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, timetype)
}

func (valdistancesignedtype *ValDistanceSignedType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valdistancesignedtype)
}

func (valdistancetype *ValDistanceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valdistancetype)
}

func (valdistanceverticaltype *ValDistanceVerticalType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valdistanceverticaltype)
}

func (valdurationtype *ValDurationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valdurationtype)
}

func (vallightintensitytype *ValLightIntensityType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, vallightintensitytype)
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
		stage.CodeColourTypes,
		stage.CodeColourType_stagedOrder,
		stage.CodeColourTypes_reference,
		&stage.CodeColourTypes_referenceOrder,
		stage.CodeColourTypes_instance,
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
		stage.CodeDayTypes,
		stage.CodeDayType_stagedOrder,
		stage.CodeDayTypes_reference,
		&stage.CodeDayTypes_referenceOrder,
		stage.CodeDayTypes_instance,
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
		stage.CodeGroundLightingTypes,
		stage.CodeGroundLightingType_stagedOrder,
		stage.CodeGroundLightingTypes_reference,
		&stage.CodeGroundLightingTypes_referenceOrder,
		stage.CodeGroundLightingTypes_instance,
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
		stage.CodeLightIntensityTypes,
		stage.CodeLightIntensityType_stagedOrder,
		stage.CodeLightIntensityTypes_reference,
		&stage.CodeLightIntensityTypes_referenceOrder,
		stage.CodeLightIntensityTypes_instance,
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
		stage.CodeLightSourceTypes,
		stage.CodeLightSourceType_stagedOrder,
		stage.CodeLightSourceTypes_reference,
		&stage.CodeLightSourceTypes_referenceOrder,
		stage.CodeLightSourceTypes_instance,
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
		stage.CodeNotePurposeTypes,
		stage.CodeNotePurposeType_stagedOrder,
		stage.CodeNotePurposeTypes_reference,
		&stage.CodeNotePurposeTypes_referenceOrder,
		stage.CodeNotePurposeTypes_instance,
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
		stage.CodeObstacleAreaTypes,
		stage.CodeObstacleAreaType_stagedOrder,
		stage.CodeObstacleAreaTypes_reference,
		&stage.CodeObstacleAreaTypes_referenceOrder,
		stage.CodeObstacleAreaTypes_instance,
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
		stage.CodeObstacleAssessmentSurfaceTypes,
		stage.CodeObstacleAssessmentSurfaceType_stagedOrder,
		stage.CodeObstacleAssessmentSurfaceTypes_reference,
		&stage.CodeObstacleAssessmentSurfaceTypes_referenceOrder,
		stage.CodeObstacleAssessmentSurfaceTypes_instance,
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
		stage.CodeStatusConstructionTypes,
		stage.CodeStatusConstructionType_stagedOrder,
		stage.CodeStatusConstructionTypes_reference,
		&stage.CodeStatusConstructionTypes_referenceOrder,
		stage.CodeStatusConstructionTypes_instance,
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
		stage.CodeStatusOperationsTypes,
		stage.CodeStatusOperationsType_stagedOrder,
		stage.CodeStatusOperationsTypes_reference,
		&stage.CodeStatusOperationsTypes_referenceOrder,
		stage.CodeStatusOperationsTypes_instance,
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
		stage.CodeTimeEventCombinationTypes,
		stage.CodeTimeEventCombinationType_stagedOrder,
		stage.CodeTimeEventCombinationTypes_reference,
		&stage.CodeTimeEventCombinationTypes_referenceOrder,
		stage.CodeTimeEventCombinationTypes_instance,
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
		stage.CodeTimeEventTypes,
		stage.CodeTimeEventType_stagedOrder,
		stage.CodeTimeEventTypes_reference,
		&stage.CodeTimeEventTypes_referenceOrder,
		stage.CodeTimeEventTypes_instance,
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
		stage.CodeTimeReferenceTypes,
		stage.CodeTimeReferenceType_stagedOrder,
		stage.CodeTimeReferenceTypes_reference,
		&stage.CodeTimeReferenceTypes_referenceOrder,
		stage.CodeTimeReferenceTypes_instance,
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
		stage.CodeVerticalDatumTypes,
		stage.CodeVerticalDatumType_stagedOrder,
		stage.CodeVerticalDatumTypes_reference,
		&stage.CodeVerticalDatumTypes_referenceOrder,
		stage.CodeVerticalDatumTypes_instance,
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
		stage.CodeVerticalStructureMarkingTypes,
		stage.CodeVerticalStructureMarkingType_stagedOrder,
		stage.CodeVerticalStructureMarkingTypes_reference,
		&stage.CodeVerticalStructureMarkingTypes_referenceOrder,
		stage.CodeVerticalStructureMarkingTypes_instance,
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
		stage.CodeVerticalStructureMaterialTypes,
		stage.CodeVerticalStructureMaterialType_stagedOrder,
		stage.CodeVerticalStructureMaterialTypes_reference,
		&stage.CodeVerticalStructureMaterialTypes_referenceOrder,
		stage.CodeVerticalStructureMaterialTypes_instance,
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
		stage.CodeVerticalStructureTypes,
		stage.CodeVerticalStructureType_stagedOrder,
		stage.CodeVerticalStructureTypes_reference,
		&stage.CodeVerticalStructureTypes_referenceOrder,
		stage.CodeVerticalStructureTypes_instance,
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
		stage.CodeYesNoTypes,
		stage.CodeYesNoType_stagedOrder,
		stage.CodeYesNoTypes_reference,
		&stage.CodeYesNoTypes_referenceOrder,
		stage.CodeYesNoTypes_instance,
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
		stage.DateMonthDayTypes,
		stage.DateMonthDayType_stagedOrder,
		stage.DateMonthDayTypes_reference,
		&stage.DateMonthDayTypes_referenceOrder,
		stage.DateMonthDayTypes_instance,
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
		stage.TextDesignatorTypes,
		stage.TextDesignatorType_stagedOrder,
		stage.TextDesignatorTypes_reference,
		&stage.TextDesignatorTypes_referenceOrder,
		stage.TextDesignatorTypes_instance,
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
		stage.TextNameTypes,
		stage.TextNameType_stagedOrder,
		stage.TextNameTypes_reference,
		&stage.TextNameTypes_referenceOrder,
		stage.TextNameTypes_instance,
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
		stage.TextNoteTypes,
		stage.TextNoteType_stagedOrder,
		stage.TextNoteTypes_reference,
		&stage.TextNoteTypes_referenceOrder,
		stage.TextNoteTypes_instance,
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
		stage.TextPropertyNameTypes,
		stage.TextPropertyNameType_stagedOrder,
		stage.TextPropertyNameTypes_reference,
		&stage.TextPropertyNameTypes_referenceOrder,
		stage.TextPropertyNameTypes_instance,
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
		stage.TimeTypes,
		stage.TimeType_stagedOrder,
		stage.TimeTypes_reference,
		&stage.TimeTypes_referenceOrder,
		stage.TimeTypes_instance,
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
		stage.ValDistanceSignedTypes,
		stage.ValDistanceSignedType_stagedOrder,
		stage.ValDistanceSignedTypes_reference,
		&stage.ValDistanceSignedTypes_referenceOrder,
		stage.ValDistanceSignedTypes_instance,
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
		stage.ValDistanceTypes,
		stage.ValDistanceType_stagedOrder,
		stage.ValDistanceTypes_reference,
		&stage.ValDistanceTypes_referenceOrder,
		stage.ValDistanceTypes_instance,
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
		stage.ValDistanceVerticalTypes,
		stage.ValDistanceVerticalType_stagedOrder,
		stage.ValDistanceVerticalTypes_reference,
		&stage.ValDistanceVerticalTypes_referenceOrder,
		stage.ValDistanceVerticalTypes_instance,
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
		stage.ValDurationTypes,
		stage.ValDurationType_stagedOrder,
		stage.ValDurationTypes_reference,
		&stage.ValDurationTypes_referenceOrder,
		stage.ValDurationTypes_instance,
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
		stage.ValLightIntensityTypes,
		stage.ValLightIntensityType_stagedOrder,
		stage.ValLightIntensityTypes_reference,
		&stage.ValLightIntensityTypes_referenceOrder,
		stage.ValLightIntensityTypes_instance,
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
	__gong__computeReferencePass1(stage, stage.CodeColourTypes, &stage.CodeColourTypes_reference, &stage.CodeColourTypes_referenceOrder, &stage.CodeColourTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDayTypes, &stage.CodeDayTypes_reference, &stage.CodeDayTypes_referenceOrder, &stage.CodeDayTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeGroundLightingTypes, &stage.CodeGroundLightingTypes_reference, &stage.CodeGroundLightingTypes_referenceOrder, &stage.CodeGroundLightingTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeLightIntensityTypes, &stage.CodeLightIntensityTypes_reference, &stage.CodeLightIntensityTypes_referenceOrder, &stage.CodeLightIntensityTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeLightSourceTypes, &stage.CodeLightSourceTypes_reference, &stage.CodeLightSourceTypes_referenceOrder, &stage.CodeLightSourceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNotePurposeTypes, &stage.CodeNotePurposeTypes_reference, &stage.CodeNotePurposeTypes_referenceOrder, &stage.CodeNotePurposeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeObstacleAreaTypes, &stage.CodeObstacleAreaTypes_reference, &stage.CodeObstacleAreaTypes_referenceOrder, &stage.CodeObstacleAreaTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeObstacleAssessmentSurfaceTypes, &stage.CodeObstacleAssessmentSurfaceTypes_reference, &stage.CodeObstacleAssessmentSurfaceTypes_referenceOrder, &stage.CodeObstacleAssessmentSurfaceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeStatusConstructionTypes, &stage.CodeStatusConstructionTypes_reference, &stage.CodeStatusConstructionTypes_referenceOrder, &stage.CodeStatusConstructionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeStatusOperationsTypes, &stage.CodeStatusOperationsTypes_reference, &stage.CodeStatusOperationsTypes_referenceOrder, &stage.CodeStatusOperationsTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTimeEventCombinationTypes, &stage.CodeTimeEventCombinationTypes_reference, &stage.CodeTimeEventCombinationTypes_referenceOrder, &stage.CodeTimeEventCombinationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTimeEventTypes, &stage.CodeTimeEventTypes_reference, &stage.CodeTimeEventTypes_referenceOrder, &stage.CodeTimeEventTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTimeReferenceTypes, &stage.CodeTimeReferenceTypes_reference, &stage.CodeTimeReferenceTypes_referenceOrder, &stage.CodeTimeReferenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeVerticalDatumTypes, &stage.CodeVerticalDatumTypes_reference, &stage.CodeVerticalDatumTypes_referenceOrder, &stage.CodeVerticalDatumTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeVerticalStructureMarkingTypes, &stage.CodeVerticalStructureMarkingTypes_reference, &stage.CodeVerticalStructureMarkingTypes_referenceOrder, &stage.CodeVerticalStructureMarkingTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeVerticalStructureMaterialTypes, &stage.CodeVerticalStructureMaterialTypes_reference, &stage.CodeVerticalStructureMaterialTypes_referenceOrder, &stage.CodeVerticalStructureMaterialTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeVerticalStructureTypes, &stage.CodeVerticalStructureTypes_reference, &stage.CodeVerticalStructureTypes_referenceOrder, &stage.CodeVerticalStructureTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeYesNoTypes, &stage.CodeYesNoTypes_reference, &stage.CodeYesNoTypes_referenceOrder, &stage.CodeYesNoTypes_instance)

	__gong__computeReferencePass1(stage, stage.DateMonthDayTypes, &stage.DateMonthDayTypes_reference, &stage.DateMonthDayTypes_referenceOrder, &stage.DateMonthDayTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextDesignatorTypes, &stage.TextDesignatorTypes_reference, &stage.TextDesignatorTypes_referenceOrder, &stage.TextDesignatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextNameTypes, &stage.TextNameTypes_reference, &stage.TextNameTypes_referenceOrder, &stage.TextNameTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextNoteTypes, &stage.TextNoteTypes_reference, &stage.TextNoteTypes_referenceOrder, &stage.TextNoteTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextPropertyNameTypes, &stage.TextPropertyNameTypes_reference, &stage.TextPropertyNameTypes_referenceOrder, &stage.TextPropertyNameTypes_instance)

	__gong__computeReferencePass1(stage, stage.TimeTypes, &stage.TimeTypes_reference, &stage.TimeTypes_referenceOrder, &stage.TimeTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDistanceSignedTypes, &stage.ValDistanceSignedTypes_reference, &stage.ValDistanceSignedTypes_referenceOrder, &stage.ValDistanceSignedTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDistanceTypes, &stage.ValDistanceTypes_reference, &stage.ValDistanceTypes_referenceOrder, &stage.ValDistanceTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDistanceVerticalTypes, &stage.ValDistanceVerticalTypes_reference, &stage.ValDistanceVerticalTypes_referenceOrder, &stage.ValDistanceVerticalTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDurationTypes, &stage.ValDurationTypes_reference, &stage.ValDurationTypes_referenceOrder, &stage.ValDurationTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValLightIntensityTypes, &stage.ValLightIntensityTypes_reference, &stage.ValLightIntensityTypes_referenceOrder, &stage.ValLightIntensityTypes_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.CodeColourTypes, stage.CodeColourTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDayTypes, stage.CodeDayTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeGroundLightingTypes, stage.CodeGroundLightingTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeLightIntensityTypes, stage.CodeLightIntensityTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeLightSourceTypes, stage.CodeLightSourceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNotePurposeTypes, stage.CodeNotePurposeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeObstacleAreaTypes, stage.CodeObstacleAreaTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeObstacleAssessmentSurfaceTypes, stage.CodeObstacleAssessmentSurfaceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeStatusConstructionTypes, stage.CodeStatusConstructionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeStatusOperationsTypes, stage.CodeStatusOperationsTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTimeEventCombinationTypes, stage.CodeTimeEventCombinationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTimeEventTypes, stage.CodeTimeEventTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTimeReferenceTypes, stage.CodeTimeReferenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeVerticalDatumTypes, stage.CodeVerticalDatumTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeVerticalStructureMarkingTypes, stage.CodeVerticalStructureMarkingTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeVerticalStructureMaterialTypes, stage.CodeVerticalStructureMaterialTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeVerticalStructureTypes, stage.CodeVerticalStructureTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeYesNoTypes, stage.CodeYesNoTypes_reference, stage)

	__gong__computeReferencePass2(stage.DateMonthDayTypes, stage.DateMonthDayTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextDesignatorTypes, stage.TextDesignatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextNameTypes, stage.TextNameTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextNoteTypes, stage.TextNoteTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextPropertyNameTypes, stage.TextPropertyNameTypes_reference, stage)

	__gong__computeReferencePass2(stage.TimeTypes, stage.TimeTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDistanceSignedTypes, stage.ValDistanceSignedTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDistanceTypes, stage.ValDistanceTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDistanceVerticalTypes, stage.ValDistanceVerticalTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDurationTypes, stage.ValDurationTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValLightIntensityTypes, stage.ValLightIntensityTypes_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (codecolourtype *CodeColourType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeColourType_stagedOrder, stage.CodeColourTypes_referenceOrder, codecolourtype, "CodeColourType")
}

func (codedaytype *CodeDayType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeDayType_stagedOrder, stage.CodeDayTypes_referenceOrder, codedaytype, "CodeDayType")
}

func (codegroundlightingtype *CodeGroundLightingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeGroundLightingType_stagedOrder, stage.CodeGroundLightingTypes_referenceOrder, codegroundlightingtype, "CodeGroundLightingType")
}

func (codelightintensitytype *CodeLightIntensityType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeLightIntensityType_stagedOrder, stage.CodeLightIntensityTypes_referenceOrder, codelightintensitytype, "CodeLightIntensityType")
}

func (codelightsourcetype *CodeLightSourceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeLightSourceType_stagedOrder, stage.CodeLightSourceTypes_referenceOrder, codelightsourcetype, "CodeLightSourceType")
}

func (codenotepurposetype *CodeNotePurposeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeNotePurposeType_stagedOrder, stage.CodeNotePurposeTypes_referenceOrder, codenotepurposetype, "CodeNotePurposeType")
}

func (codeobstacleareatype *CodeObstacleAreaType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeObstacleAreaType_stagedOrder, stage.CodeObstacleAreaTypes_referenceOrder, codeobstacleareatype, "CodeObstacleAreaType")
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeObstacleAssessmentSurfaceType_stagedOrder, stage.CodeObstacleAssessmentSurfaceTypes_referenceOrder, codeobstacleassessmentsurfacetype, "CodeObstacleAssessmentSurfaceType")
}

func (codestatusconstructiontype *CodeStatusConstructionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeStatusConstructionType_stagedOrder, stage.CodeStatusConstructionTypes_referenceOrder, codestatusconstructiontype, "CodeStatusConstructionType")
}

func (codestatusoperationstype *CodeStatusOperationsType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeStatusOperationsType_stagedOrder, stage.CodeStatusOperationsTypes_referenceOrder, codestatusoperationstype, "CodeStatusOperationsType")
}

func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeTimeEventCombinationType_stagedOrder, stage.CodeTimeEventCombinationTypes_referenceOrder, codetimeeventcombinationtype, "CodeTimeEventCombinationType")
}

func (codetimeeventtype *CodeTimeEventType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeTimeEventType_stagedOrder, stage.CodeTimeEventTypes_referenceOrder, codetimeeventtype, "CodeTimeEventType")
}

func (codetimereferencetype *CodeTimeReferenceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeTimeReferenceType_stagedOrder, stage.CodeTimeReferenceTypes_referenceOrder, codetimereferencetype, "CodeTimeReferenceType")
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeVerticalDatumType_stagedOrder, stage.CodeVerticalDatumTypes_referenceOrder, codeverticaldatumtype, "CodeVerticalDatumType")
}

func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeVerticalStructureMarkingType_stagedOrder, stage.CodeVerticalStructureMarkingTypes_referenceOrder, codeverticalstructuremarkingtype, "CodeVerticalStructureMarkingType")
}

func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeVerticalStructureMaterialType_stagedOrder, stage.CodeVerticalStructureMaterialTypes_referenceOrder, codeverticalstructurematerialtype, "CodeVerticalStructureMaterialType")
}

func (codeverticalstructuretype *CodeVerticalStructureType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeVerticalStructureType_stagedOrder, stage.CodeVerticalStructureTypes_referenceOrder, codeverticalstructuretype, "CodeVerticalStructureType")
}

func (codeyesnotype *CodeYesNoType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeYesNoType_stagedOrder, stage.CodeYesNoTypes_referenceOrder, codeyesnotype, "CodeYesNoType")
}

func (datemonthdaytype *DateMonthDayType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DateMonthDayType_stagedOrder, stage.DateMonthDayTypes_referenceOrder, datemonthdaytype, "DateMonthDayType")
}

func (textdesignatortype *TextDesignatorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TextDesignatorType_stagedOrder, stage.TextDesignatorTypes_referenceOrder, textdesignatortype, "TextDesignatorType")
}

func (textnametype *TextNameType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TextNameType_stagedOrder, stage.TextNameTypes_referenceOrder, textnametype, "TextNameType")
}

func (textnotetype *TextNoteType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TextNoteType_stagedOrder, stage.TextNoteTypes_referenceOrder, textnotetype, "TextNoteType")
}

func (textpropertynametype *TextPropertyNameType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TextPropertyNameType_stagedOrder, stage.TextPropertyNameTypes_referenceOrder, textpropertynametype, "TextPropertyNameType")
}

func (timetype *TimeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TimeType_stagedOrder, stage.TimeTypes_referenceOrder, timetype, "TimeType")
}

func (valdistancesignedtype *ValDistanceSignedType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValDistanceSignedType_stagedOrder, stage.ValDistanceSignedTypes_referenceOrder, valdistancesignedtype, "ValDistanceSignedType")
}

func (valdistancetype *ValDistanceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValDistanceType_stagedOrder, stage.ValDistanceTypes_referenceOrder, valdistancetype, "ValDistanceType")
}

func (valdistanceverticaltype *ValDistanceVerticalType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValDistanceVerticalType_stagedOrder, stage.ValDistanceVerticalTypes_referenceOrder, valdistanceverticaltype, "ValDistanceVerticalType")
}

func (valdurationtype *ValDurationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValDurationType_stagedOrder, stage.ValDurationTypes_referenceOrder, valdurationtype, "ValDurationType")
}

func (vallightintensitytype *ValLightIntensityType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValLightIntensityType_stagedOrder, stage.ValLightIntensityTypes_referenceOrder, vallightintensitytype, "ValLightIntensityType")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (codecolourtype *CodeColourType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codecolourtype, codecolourtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codecolourtype *CodeColourType) GongGetReferenceIdentifier(stage *Stage) string {
	return codecolourtype.GongGetIdentifier(stage)
}

func (codedaytype *CodeDayType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codedaytype, codedaytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codedaytype *CodeDayType) GongGetReferenceIdentifier(stage *Stage) string {
	return codedaytype.GongGetIdentifier(stage)
}

func (codegroundlightingtype *CodeGroundLightingType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codegroundlightingtype, codegroundlightingtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codegroundlightingtype *CodeGroundLightingType) GongGetReferenceIdentifier(stage *Stage) string {
	return codegroundlightingtype.GongGetIdentifier(stage)
}

func (codelightintensitytype *CodeLightIntensityType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codelightintensitytype, codelightintensitytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codelightintensitytype *CodeLightIntensityType) GongGetReferenceIdentifier(stage *Stage) string {
	return codelightintensitytype.GongGetIdentifier(stage)
}

func (codelightsourcetype *CodeLightSourceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codelightsourcetype, codelightsourcetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codelightsourcetype *CodeLightSourceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codelightsourcetype.GongGetIdentifier(stage)
}

func (codenotepurposetype *CodeNotePurposeType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codenotepurposetype, codenotepurposetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codenotepurposetype *CodeNotePurposeType) GongGetReferenceIdentifier(stage *Stage) string {
	return codenotepurposetype.GongGetIdentifier(stage)
}

func (codeobstacleareatype *CodeObstacleAreaType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeobstacleareatype, codeobstacleareatype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeobstacleareatype *CodeObstacleAreaType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeobstacleareatype.GongGetIdentifier(stage)
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeobstacleassessmentsurfacetype, codeobstacleassessmentsurfacetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeobstacleassessmentsurfacetype.GongGetIdentifier(stage)
}

func (codestatusconstructiontype *CodeStatusConstructionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codestatusconstructiontype, codestatusconstructiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codestatusconstructiontype *CodeStatusConstructionType) GongGetReferenceIdentifier(stage *Stage) string {
	return codestatusconstructiontype.GongGetIdentifier(stage)
}

func (codestatusoperationstype *CodeStatusOperationsType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codestatusoperationstype, codestatusoperationstype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codestatusoperationstype *CodeStatusOperationsType) GongGetReferenceIdentifier(stage *Stage) string {
	return codestatusoperationstype.GongGetIdentifier(stage)
}

func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codetimeeventcombinationtype, codetimeeventcombinationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongGetReferenceIdentifier(stage *Stage) string {
	return codetimeeventcombinationtype.GongGetIdentifier(stage)
}

func (codetimeeventtype *CodeTimeEventType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codetimeeventtype, codetimeeventtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codetimeeventtype *CodeTimeEventType) GongGetReferenceIdentifier(stage *Stage) string {
	return codetimeeventtype.GongGetIdentifier(stage)
}

func (codetimereferencetype *CodeTimeReferenceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codetimereferencetype, codetimereferencetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codetimereferencetype *CodeTimeReferenceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codetimereferencetype.GongGetIdentifier(stage)
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeverticaldatumtype, codeverticaldatumtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeverticaldatumtype *CodeVerticalDatumType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeverticaldatumtype.GongGetIdentifier(stage)
}

func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeverticalstructuremarkingtype, codeverticalstructuremarkingtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeverticalstructuremarkingtype.GongGetIdentifier(stage)
}

func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeverticalstructurematerialtype, codeverticalstructurematerialtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeverticalstructurematerialtype.GongGetIdentifier(stage)
}

func (codeverticalstructuretype *CodeVerticalStructureType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeverticalstructuretype, codeverticalstructuretype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeverticalstructuretype *CodeVerticalStructureType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeverticalstructuretype.GongGetIdentifier(stage)
}

func (codeyesnotype *CodeYesNoType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeyesnotype, codeyesnotype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeyesnotype *CodeYesNoType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeyesnotype.GongGetIdentifier(stage)
}

func (datemonthdaytype *DateMonthDayType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(datemonthdaytype, datemonthdaytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (datemonthdaytype *DateMonthDayType) GongGetReferenceIdentifier(stage *Stage) string {
	return datemonthdaytype.GongGetIdentifier(stage)
}

func (textdesignatortype *TextDesignatorType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(textdesignatortype, textdesignatortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (textdesignatortype *TextDesignatorType) GongGetReferenceIdentifier(stage *Stage) string {
	return textdesignatortype.GongGetIdentifier(stage)
}

func (textnametype *TextNameType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(textnametype, textnametype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (textnametype *TextNameType) GongGetReferenceIdentifier(stage *Stage) string {
	return textnametype.GongGetIdentifier(stage)
}

func (textnotetype *TextNoteType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(textnotetype, textnotetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (textnotetype *TextNoteType) GongGetReferenceIdentifier(stage *Stage) string {
	return textnotetype.GongGetIdentifier(stage)
}

func (textpropertynametype *TextPropertyNameType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(textpropertynametype, textpropertynametype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (textpropertynametype *TextPropertyNameType) GongGetReferenceIdentifier(stage *Stage) string {
	return textpropertynametype.GongGetIdentifier(stage)
}

func (timetype *TimeType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(timetype, timetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (timetype *TimeType) GongGetReferenceIdentifier(stage *Stage) string {
	return timetype.GongGetIdentifier(stage)
}

func (valdistancesignedtype *ValDistanceSignedType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valdistancesignedtype, valdistancesignedtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valdistancesignedtype *ValDistanceSignedType) GongGetReferenceIdentifier(stage *Stage) string {
	return valdistancesignedtype.GongGetIdentifier(stage)
}

func (valdistancetype *ValDistanceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valdistancetype, valdistancetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valdistancetype *ValDistanceType) GongGetReferenceIdentifier(stage *Stage) string {
	return valdistancetype.GongGetIdentifier(stage)
}

func (valdistanceverticaltype *ValDistanceVerticalType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valdistanceverticaltype, valdistanceverticaltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valdistanceverticaltype *ValDistanceVerticalType) GongGetReferenceIdentifier(stage *Stage) string {
	return valdistanceverticaltype.GongGetIdentifier(stage)
}

func (valdurationtype *ValDurationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valdurationtype, valdurationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valdurationtype *ValDurationType) GongGetReferenceIdentifier(stage *Stage) string {
	return valdurationtype.GongGetIdentifier(stage)
}

func (vallightintensitytype *ValLightIntensityType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(vallightintensitytype, vallightintensitytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (vallightintensitytype *ValLightIntensityType) GongGetReferenceIdentifier(stage *Stage) string {
	return vallightintensitytype.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (codecolourtype *CodeColourType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codecolourtype.GongGetIdentifier(stage), "CodeColourType", codecolourtype.Name)
}

func (codedaytype *CodeDayType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codedaytype.GongGetIdentifier(stage), "CodeDayType", codedaytype.Name)
}

func (codegroundlightingtype *CodeGroundLightingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codegroundlightingtype.GongGetIdentifier(stage), "CodeGroundLightingType", codegroundlightingtype.Name)
}

func (codelightintensitytype *CodeLightIntensityType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codelightintensitytype.GongGetIdentifier(stage), "CodeLightIntensityType", codelightintensitytype.Name)
}

func (codelightsourcetype *CodeLightSourceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codelightsourcetype.GongGetIdentifier(stage), "CodeLightSourceType", codelightsourcetype.Name)
}

func (codenotepurposetype *CodeNotePurposeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codenotepurposetype.GongGetIdentifier(stage), "CodeNotePurposeType", codenotepurposetype.Name)
}

func (codeobstacleareatype *CodeObstacleAreaType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeobstacleareatype.GongGetIdentifier(stage), "CodeObstacleAreaType", codeobstacleareatype.Name)
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeobstacleassessmentsurfacetype.GongGetIdentifier(stage), "CodeObstacleAssessmentSurfaceType", codeobstacleassessmentsurfacetype.Name)
}

func (codestatusconstructiontype *CodeStatusConstructionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codestatusconstructiontype.GongGetIdentifier(stage), "CodeStatusConstructionType", codestatusconstructiontype.Name)
}

func (codestatusoperationstype *CodeStatusOperationsType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codestatusoperationstype.GongGetIdentifier(stage), "CodeStatusOperationsType", codestatusoperationstype.Name)
}

func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codetimeeventcombinationtype.GongGetIdentifier(stage), "CodeTimeEventCombinationType", codetimeeventcombinationtype.Name)
}

func (codetimeeventtype *CodeTimeEventType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codetimeeventtype.GongGetIdentifier(stage), "CodeTimeEventType", codetimeeventtype.Name)
}

func (codetimereferencetype *CodeTimeReferenceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codetimereferencetype.GongGetIdentifier(stage), "CodeTimeReferenceType", codetimereferencetype.Name)
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeverticaldatumtype.GongGetIdentifier(stage), "CodeVerticalDatumType", codeverticaldatumtype.Name)
}

func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeverticalstructuremarkingtype.GongGetIdentifier(stage), "CodeVerticalStructureMarkingType", codeverticalstructuremarkingtype.Name)
}

func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeverticalstructurematerialtype.GongGetIdentifier(stage), "CodeVerticalStructureMaterialType", codeverticalstructurematerialtype.Name)
}

func (codeverticalstructuretype *CodeVerticalStructureType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeverticalstructuretype.GongGetIdentifier(stage), "CodeVerticalStructureType", codeverticalstructuretype.Name)
}

func (codeyesnotype *CodeYesNoType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeyesnotype.GongGetIdentifier(stage), "CodeYesNoType", codeyesnotype.Name)
}

func (datemonthdaytype *DateMonthDayType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(datemonthdaytype.GongGetIdentifier(stage), "DateMonthDayType", datemonthdaytype.Name)
}

func (textdesignatortype *TextDesignatorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(textdesignatortype.GongGetIdentifier(stage), "TextDesignatorType", textdesignatortype.Name)
}

func (textnametype *TextNameType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(textnametype.GongGetIdentifier(stage), "TextNameType", textnametype.Name)
}

func (textnotetype *TextNoteType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(textnotetype.GongGetIdentifier(stage), "TextNoteType", textnotetype.Name)
}

func (textpropertynametype *TextPropertyNameType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(textpropertynametype.GongGetIdentifier(stage), "TextPropertyNameType", textpropertynametype.Name)
}

func (timetype *TimeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(timetype.GongGetIdentifier(stage), "TimeType", timetype.Name)
}

func (valdistancesignedtype *ValDistanceSignedType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valdistancesignedtype.GongGetIdentifier(stage), "ValDistanceSignedType", valdistancesignedtype.Name)
}

func (valdistancetype *ValDistanceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valdistancetype.GongGetIdentifier(stage), "ValDistanceType", valdistancetype.Name)
}

func (valdistanceverticaltype *ValDistanceVerticalType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valdistanceverticaltype.GongGetIdentifier(stage), "ValDistanceVerticalType", valdistanceverticaltype.Name)
}

func (valdurationtype *ValDurationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valdurationtype.GongGetIdentifier(stage), "ValDurationType", valdurationtype.Name)
}

func (vallightintensitytype *ValLightIntensityType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(vallightintensitytype.GongGetIdentifier(stage), "ValLightIntensityType", vallightintensitytype.Name)
}

// insertion point for unstaging
func (codecolourtype *CodeColourType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codecolourtype.GongGetReferenceIdentifier(stage))
}

func (codedaytype *CodeDayType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codedaytype.GongGetReferenceIdentifier(stage))
}

func (codegroundlightingtype *CodeGroundLightingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codegroundlightingtype.GongGetReferenceIdentifier(stage))
}

func (codelightintensitytype *CodeLightIntensityType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codelightintensitytype.GongGetReferenceIdentifier(stage))
}

func (codelightsourcetype *CodeLightSourceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codelightsourcetype.GongGetReferenceIdentifier(stage))
}

func (codenotepurposetype *CodeNotePurposeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codenotepurposetype.GongGetReferenceIdentifier(stage))
}

func (codeobstacleareatype *CodeObstacleAreaType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeobstacleareatype.GongGetReferenceIdentifier(stage))
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeobstacleassessmentsurfacetype.GongGetReferenceIdentifier(stage))
}

func (codestatusconstructiontype *CodeStatusConstructionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codestatusconstructiontype.GongGetReferenceIdentifier(stage))
}

func (codestatusoperationstype *CodeStatusOperationsType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codestatusoperationstype.GongGetReferenceIdentifier(stage))
}

func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codetimeeventcombinationtype.GongGetReferenceIdentifier(stage))
}

func (codetimeeventtype *CodeTimeEventType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codetimeeventtype.GongGetReferenceIdentifier(stage))
}

func (codetimereferencetype *CodeTimeReferenceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codetimereferencetype.GongGetReferenceIdentifier(stage))
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeverticaldatumtype.GongGetReferenceIdentifier(stage))
}

func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeverticalstructuremarkingtype.GongGetReferenceIdentifier(stage))
}

func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeverticalstructurematerialtype.GongGetReferenceIdentifier(stage))
}

func (codeverticalstructuretype *CodeVerticalStructureType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeverticalstructuretype.GongGetReferenceIdentifier(stage))
}

func (codeyesnotype *CodeYesNoType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeyesnotype.GongGetReferenceIdentifier(stage))
}

func (datemonthdaytype *DateMonthDayType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(datemonthdaytype.GongGetReferenceIdentifier(stage))
}

func (textdesignatortype *TextDesignatorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(textdesignatortype.GongGetReferenceIdentifier(stage))
}

func (textnametype *TextNameType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(textnametype.GongGetReferenceIdentifier(stage))
}

func (textnotetype *TextNoteType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(textnotetype.GongGetReferenceIdentifier(stage))
}

func (textpropertynametype *TextPropertyNameType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(textpropertynametype.GongGetReferenceIdentifier(stage))
}

func (timetype *TimeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(timetype.GongGetReferenceIdentifier(stage))
}

func (valdistancesignedtype *ValDistanceSignedType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valdistancesignedtype.GongGetReferenceIdentifier(stage))
}

func (valdistancetype *ValDistanceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valdistancetype.GongGetReferenceIdentifier(stage))
}

func (valdistanceverticaltype *ValDistanceVerticalType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valdistanceverticaltype.GongGetReferenceIdentifier(stage))
}

func (valdurationtype *ValDurationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valdurationtype.GongGetReferenceIdentifier(stage))
}

func (vallightintensitytype *ValLightIntensityType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(vallightintensitytype.GongGetReferenceIdentifier(stage))
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
	return fmt.Sprintf("__aixmdatatypes__%s__%08d_", s.GongGetGongstructName(), order)
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
