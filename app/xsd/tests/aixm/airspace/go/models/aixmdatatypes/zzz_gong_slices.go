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
	res = __gong__appendInstances(res, stage.CodeAircraftCategoryTypes)

	res = __gong__appendInstances(res, stage.CodeAircraftEngineNumberTypes)

	res = __gong__appendInstances(res, stage.CodeAircraftEngineTypes)

	res = __gong__appendInstances(res, stage.CodeAircraftICAOTypes)

	res = __gong__appendInstances(res, stage.CodeAircraftTypes)

	res = __gong__appendInstances(res, stage.CodeAircraftWingspanClassTypes)

	res = __gong__appendInstances(res, stage.CodeAirspaceActivityTypes)

	res = __gong__appendInstances(res, stage.CodeAirspaceAggregationTypes)

	res = __gong__appendInstances(res, stage.CodeAirspaceClassificationTypes)

	res = __gong__appendInstances(res, stage.CodeAirspaceDependencyTypes)

	res = __gong__appendInstances(res, stage.CodeAirspaceDesignatorTypes)

	res = __gong__appendInstances(res, stage.CodeAirspaceTypes)

	res = __gong__appendInstances(res, stage.CodeAltitudeUseTypes)

	res = __gong__appendInstances(res, stage.CodeAuthorityTypes)

	res = __gong__appendInstances(res, stage.CodeCommunicationModeTypes)

	res = __gong__appendInstances(res, stage.CodeDayTypes)

	res = __gong__appendInstances(res, stage.CodeEquipmentAntiCollisionTypes)

	res = __gong__appendInstances(res, stage.CodeGeoBorderTypes)

	res = __gong__appendInstances(res, stage.CodeMilitaryOperationsTypes)

	res = __gong__appendInstances(res, stage.CodeNavigationEquipmentTypes)

	res = __gong__appendInstances(res, stage.CodeNavigationSpecificationTypes)

	res = __gong__appendInstances(res, stage.CodeNotePurposeTypes)

	res = __gong__appendInstances(res, stage.CodeRVSMTypes)

	res = __gong__appendInstances(res, stage.CodeStatusAirspaceTypes)

	res = __gong__appendInstances(res, stage.CodeTimeEventCombinationTypes)

	res = __gong__appendInstances(res, stage.CodeTimeEventTypes)

	res = __gong__appendInstances(res, stage.CodeTimeReferenceTypes)

	res = __gong__appendInstances(res, stage.CodeTransponderTypes)

	res = __gong__appendInstances(res, stage.CodeValueInterpretationTypes)

	res = __gong__appendInstances(res, stage.CodeVerticalDatumTypes)

	res = __gong__appendInstances(res, stage.CodeVerticalReferenceTypes)

	res = __gong__appendInstances(res, stage.CodeWakeTurbulenceTypes)

	res = __gong__appendInstances(res, stage.CodeYesNoTypes)

	res = __gong__appendInstances(res, stage.DateMonthDayTypes)

	res = __gong__appendInstances(res, stage.NoNumberTypes)

	res = __gong__appendInstances(res, stage.NoSequenceTypes)

	res = __gong__appendInstances(res, stage.TextNameTypes)

	res = __gong__appendInstances(res, stage.TextNoteTypes)

	res = __gong__appendInstances(res, stage.TextPropertyNameTypes)

	res = __gong__appendInstances(res, stage.TimeTypes)

	res = __gong__appendInstances(res, stage.ValDistanceSignedTypes)

	res = __gong__appendInstances(res, stage.ValDistanceTypes)

	res = __gong__appendInstances(res, stage.ValDistanceVerticalTypes)

	res = __gong__appendInstances(res, stage.ValDurationTypes)

	res = __gong__appendInstances(res, stage.ValFLTypes)

	res = __gong__appendInstances(res, stage.ValSpeedTypes)

	res = __gong__appendInstances(res, stage.ValWeightTypes)

	return
}

// insertion point per named struct
func (codeaircraftcategorytype *CodeAircraftCategoryType) GongCopy() GongstructIF {
	newInstance := new(CodeAircraftCategoryType)
	codeaircraftcategorytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongCopy() GongstructIF {
	newInstance := new(CodeAircraftEngineNumberType)
	codeaircraftenginenumbertype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeaircraftenginetype *CodeAircraftEngineType) GongCopy() GongstructIF {
	newInstance := new(CodeAircraftEngineType)
	codeaircraftenginetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeaircrafticaotype *CodeAircraftICAOType) GongCopy() GongstructIF {
	newInstance := new(CodeAircraftICAOType)
	codeaircrafticaotype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeaircrafttype *CodeAircraftType) GongCopy() GongstructIF {
	newInstance := new(CodeAircraftType)
	codeaircrafttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongCopy() GongstructIF {
	newInstance := new(CodeAircraftWingspanClassType)
	codeaircraftwingspanclasstype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeairspaceactivitytype *CodeAirspaceActivityType) GongCopy() GongstructIF {
	newInstance := new(CodeAirspaceActivityType)
	codeairspaceactivitytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongCopy() GongstructIF {
	newInstance := new(CodeAirspaceAggregationType)
	codeairspaceaggregationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongCopy() GongstructIF {
	newInstance := new(CodeAirspaceClassificationType)
	codeairspaceclassificationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeairspacedependencytype *CodeAirspaceDependencyType) GongCopy() GongstructIF {
	newInstance := new(CodeAirspaceDependencyType)
	codeairspacedependencytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongCopy() GongstructIF {
	newInstance := new(CodeAirspaceDesignatorType)
	codeairspacedesignatortype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeairspacetype *CodeAirspaceType) GongCopy() GongstructIF {
	newInstance := new(CodeAirspaceType)
	codeairspacetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codealtitudeusetype *CodeAltitudeUseType) GongCopy() GongstructIF {
	newInstance := new(CodeAltitudeUseType)
	codealtitudeusetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeauthoritytype *CodeAuthorityType) GongCopy() GongstructIF {
	newInstance := new(CodeAuthorityType)
	codeauthoritytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongCopy() GongstructIF {
	newInstance := new(CodeCommunicationModeType)
	codecommunicationmodetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codedaytype *CodeDayType) GongCopy() GongstructIF {
	newInstance := new(CodeDayType)
	codedaytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongCopy() GongstructIF {
	newInstance := new(CodeEquipmentAntiCollisionType)
	codeequipmentanticollisiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codegeobordertype *CodeGeoBorderType) GongCopy() GongstructIF {
	newInstance := new(CodeGeoBorderType)
	codegeobordertype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongCopy() GongstructIF {
	newInstance := new(CodeMilitaryOperationsType)
	codemilitaryoperationstype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongCopy() GongstructIF {
	newInstance := new(CodeNavigationEquipmentType)
	codenavigationequipmenttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongCopy() GongstructIF {
	newInstance := new(CodeNavigationSpecificationType)
	codenavigationspecificationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codenotepurposetype *CodeNotePurposeType) GongCopy() GongstructIF {
	newInstance := new(CodeNotePurposeType)
	codenotepurposetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codervsmtype *CodeRVSMType) GongCopy() GongstructIF {
	newInstance := new(CodeRVSMType)
	codervsmtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codestatusairspacetype *CodeStatusAirspaceType) GongCopy() GongstructIF {
	newInstance := new(CodeStatusAirspaceType)
	codestatusairspacetype.GongCopyBasicFields(newInstance)
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

func (codetranspondertype *CodeTransponderType) GongCopy() GongstructIF {
	newInstance := new(CodeTransponderType)
	codetranspondertype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codevalueinterpretationtype *CodeValueInterpretationType) GongCopy() GongstructIF {
	newInstance := new(CodeValueInterpretationType)
	codevalueinterpretationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongCopy() GongstructIF {
	newInstance := new(CodeVerticalDatumType)
	codeverticaldatumtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeverticalreferencetype *CodeVerticalReferenceType) GongCopy() GongstructIF {
	newInstance := new(CodeVerticalReferenceType)
	codeverticalreferencetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongCopy() GongstructIF {
	newInstance := new(CodeWakeTurbulenceType)
	codewaketurbulencetype.GongCopyBasicFields(newInstance)
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

func (nonumbertype *NoNumberType) GongCopy() GongstructIF {
	newInstance := new(NoNumberType)
	nonumbertype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (nosequencetype *NoSequenceType) GongCopy() GongstructIF {
	newInstance := new(NoSequenceType)
	nosequencetype.GongCopyBasicFields(newInstance)
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

func (valfltype *ValFLType) GongCopy() GongstructIF {
	newInstance := new(ValFLType)
	valfltype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (valspeedtype *ValSpeedType) GongCopy() GongstructIF {
	newInstance := new(ValSpeedType)
	valspeedtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (valweighttype *ValWeightType) GongCopy() GongstructIF {
	newInstance := new(ValWeightType)
	valweighttype.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (codeaircraftcategorytype *CodeAircraftCategoryType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeaircraftcategorytype)
}

func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeaircraftenginenumbertype)
}

func (codeaircraftenginetype *CodeAircraftEngineType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeaircraftenginetype)
}

func (codeaircrafticaotype *CodeAircraftICAOType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeaircrafticaotype)
}

func (codeaircrafttype *CodeAircraftType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeaircrafttype)
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeaircraftwingspanclasstype)
}

func (codeairspaceactivitytype *CodeAirspaceActivityType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeairspaceactivitytype)
}

func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeairspaceaggregationtype)
}

func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeairspaceclassificationtype)
}

func (codeairspacedependencytype *CodeAirspaceDependencyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeairspacedependencytype)
}

func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeairspacedesignatortype)
}

func (codeairspacetype *CodeAirspaceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeairspacetype)
}

func (codealtitudeusetype *CodeAltitudeUseType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codealtitudeusetype)
}

func (codeauthoritytype *CodeAuthorityType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeauthoritytype)
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codecommunicationmodetype)
}

func (codedaytype *CodeDayType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codedaytype)
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeequipmentanticollisiontype)
}

func (codegeobordertype *CodeGeoBorderType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codegeobordertype)
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codemilitaryoperationstype)
}

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codenavigationequipmenttype)
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codenavigationspecificationtype)
}

func (codenotepurposetype *CodeNotePurposeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codenotepurposetype)
}

func (codervsmtype *CodeRVSMType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codervsmtype)
}

func (codestatusairspacetype *CodeStatusAirspaceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codestatusairspacetype)
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

func (codetranspondertype *CodeTransponderType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codetranspondertype)
}

func (codevalueinterpretationtype *CodeValueInterpretationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codevalueinterpretationtype)
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeverticaldatumtype)
}

func (codeverticalreferencetype *CodeVerticalReferenceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeverticalreferencetype)
}

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codewaketurbulencetype)
}

func (codeyesnotype *CodeYesNoType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeyesnotype)
}

func (datemonthdaytype *DateMonthDayType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, datemonthdaytype)
}

func (nonumbertype *NoNumberType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, nonumbertype)
}

func (nosequencetype *NoSequenceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, nosequencetype)
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

func (valfltype *ValFLType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valfltype)
}

func (valspeedtype *ValSpeedType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valspeedtype)
}

func (valweighttype *ValWeightType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valweighttype)
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
		stage.CodeAircraftCategoryTypes,
		stage.CodeAircraftCategoryType_stagedOrder,
		stage.CodeAircraftCategoryTypes_reference,
		&stage.CodeAircraftCategoryTypes_referenceOrder,
		stage.CodeAircraftCategoryTypes_instance,
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
		stage.CodeAircraftEngineNumberTypes,
		stage.CodeAircraftEngineNumberType_stagedOrder,
		stage.CodeAircraftEngineNumberTypes_reference,
		&stage.CodeAircraftEngineNumberTypes_referenceOrder,
		stage.CodeAircraftEngineNumberTypes_instance,
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
		stage.CodeAircraftEngineTypes,
		stage.CodeAircraftEngineType_stagedOrder,
		stage.CodeAircraftEngineTypes_reference,
		&stage.CodeAircraftEngineTypes_referenceOrder,
		stage.CodeAircraftEngineTypes_instance,
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
		stage.CodeAircraftICAOTypes,
		stage.CodeAircraftICAOType_stagedOrder,
		stage.CodeAircraftICAOTypes_reference,
		&stage.CodeAircraftICAOTypes_referenceOrder,
		stage.CodeAircraftICAOTypes_instance,
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
		stage.CodeAircraftTypes,
		stage.CodeAircraftType_stagedOrder,
		stage.CodeAircraftTypes_reference,
		&stage.CodeAircraftTypes_referenceOrder,
		stage.CodeAircraftTypes_instance,
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
		stage.CodeAircraftWingspanClassTypes,
		stage.CodeAircraftWingspanClassType_stagedOrder,
		stage.CodeAircraftWingspanClassTypes_reference,
		&stage.CodeAircraftWingspanClassTypes_referenceOrder,
		stage.CodeAircraftWingspanClassTypes_instance,
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
		stage.CodeAirspaceActivityTypes,
		stage.CodeAirspaceActivityType_stagedOrder,
		stage.CodeAirspaceActivityTypes_reference,
		&stage.CodeAirspaceActivityTypes_referenceOrder,
		stage.CodeAirspaceActivityTypes_instance,
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
		stage.CodeAirspaceAggregationTypes,
		stage.CodeAirspaceAggregationType_stagedOrder,
		stage.CodeAirspaceAggregationTypes_reference,
		&stage.CodeAirspaceAggregationTypes_referenceOrder,
		stage.CodeAirspaceAggregationTypes_instance,
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
		stage.CodeAirspaceClassificationTypes,
		stage.CodeAirspaceClassificationType_stagedOrder,
		stage.CodeAirspaceClassificationTypes_reference,
		&stage.CodeAirspaceClassificationTypes_referenceOrder,
		stage.CodeAirspaceClassificationTypes_instance,
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
		stage.CodeAirspaceDependencyTypes,
		stage.CodeAirspaceDependencyType_stagedOrder,
		stage.CodeAirspaceDependencyTypes_reference,
		&stage.CodeAirspaceDependencyTypes_referenceOrder,
		stage.CodeAirspaceDependencyTypes_instance,
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
		stage.CodeAirspaceDesignatorTypes,
		stage.CodeAirspaceDesignatorType_stagedOrder,
		stage.CodeAirspaceDesignatorTypes_reference,
		&stage.CodeAirspaceDesignatorTypes_referenceOrder,
		stage.CodeAirspaceDesignatorTypes_instance,
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
		stage.CodeAirspaceTypes,
		stage.CodeAirspaceType_stagedOrder,
		stage.CodeAirspaceTypes_reference,
		&stage.CodeAirspaceTypes_referenceOrder,
		stage.CodeAirspaceTypes_instance,
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
		stage.CodeAltitudeUseTypes,
		stage.CodeAltitudeUseType_stagedOrder,
		stage.CodeAltitudeUseTypes_reference,
		&stage.CodeAltitudeUseTypes_referenceOrder,
		stage.CodeAltitudeUseTypes_instance,
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
		stage.CodeAuthorityTypes,
		stage.CodeAuthorityType_stagedOrder,
		stage.CodeAuthorityTypes_reference,
		&stage.CodeAuthorityTypes_referenceOrder,
		stage.CodeAuthorityTypes_instance,
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
		stage.CodeCommunicationModeTypes,
		stage.CodeCommunicationModeType_stagedOrder,
		stage.CodeCommunicationModeTypes_reference,
		&stage.CodeCommunicationModeTypes_referenceOrder,
		stage.CodeCommunicationModeTypes_instance,
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
		stage.CodeEquipmentAntiCollisionTypes,
		stage.CodeEquipmentAntiCollisionType_stagedOrder,
		stage.CodeEquipmentAntiCollisionTypes_reference,
		&stage.CodeEquipmentAntiCollisionTypes_referenceOrder,
		stage.CodeEquipmentAntiCollisionTypes_instance,
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
		stage.CodeGeoBorderTypes,
		stage.CodeGeoBorderType_stagedOrder,
		stage.CodeGeoBorderTypes_reference,
		&stage.CodeGeoBorderTypes_referenceOrder,
		stage.CodeGeoBorderTypes_instance,
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
		stage.CodeMilitaryOperationsTypes,
		stage.CodeMilitaryOperationsType_stagedOrder,
		stage.CodeMilitaryOperationsTypes_reference,
		&stage.CodeMilitaryOperationsTypes_referenceOrder,
		stage.CodeMilitaryOperationsTypes_instance,
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
		stage.CodeNavigationEquipmentTypes,
		stage.CodeNavigationEquipmentType_stagedOrder,
		stage.CodeNavigationEquipmentTypes_reference,
		&stage.CodeNavigationEquipmentTypes_referenceOrder,
		stage.CodeNavigationEquipmentTypes_instance,
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
		stage.CodeNavigationSpecificationTypes,
		stage.CodeNavigationSpecificationType_stagedOrder,
		stage.CodeNavigationSpecificationTypes_reference,
		&stage.CodeNavigationSpecificationTypes_referenceOrder,
		stage.CodeNavigationSpecificationTypes_instance,
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
		stage.CodeRVSMTypes,
		stage.CodeRVSMType_stagedOrder,
		stage.CodeRVSMTypes_reference,
		&stage.CodeRVSMTypes_referenceOrder,
		stage.CodeRVSMTypes_instance,
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
		stage.CodeStatusAirspaceTypes,
		stage.CodeStatusAirspaceType_stagedOrder,
		stage.CodeStatusAirspaceTypes_reference,
		&stage.CodeStatusAirspaceTypes_referenceOrder,
		stage.CodeStatusAirspaceTypes_instance,
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
		stage.CodeTransponderTypes,
		stage.CodeTransponderType_stagedOrder,
		stage.CodeTransponderTypes_reference,
		&stage.CodeTransponderTypes_referenceOrder,
		stage.CodeTransponderTypes_instance,
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
		stage.CodeValueInterpretationTypes,
		stage.CodeValueInterpretationType_stagedOrder,
		stage.CodeValueInterpretationTypes_reference,
		&stage.CodeValueInterpretationTypes_referenceOrder,
		stage.CodeValueInterpretationTypes_instance,
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
		stage.CodeVerticalReferenceTypes,
		stage.CodeVerticalReferenceType_stagedOrder,
		stage.CodeVerticalReferenceTypes_reference,
		&stage.CodeVerticalReferenceTypes_referenceOrder,
		stage.CodeVerticalReferenceTypes_instance,
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
		stage.CodeWakeTurbulenceTypes,
		stage.CodeWakeTurbulenceType_stagedOrder,
		stage.CodeWakeTurbulenceTypes_reference,
		&stage.CodeWakeTurbulenceTypes_referenceOrder,
		stage.CodeWakeTurbulenceTypes_instance,
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
		stage.NoNumberTypes,
		stage.NoNumberType_stagedOrder,
		stage.NoNumberTypes_reference,
		&stage.NoNumberTypes_referenceOrder,
		stage.NoNumberTypes_instance,
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
		stage.NoSequenceTypes,
		stage.NoSequenceType_stagedOrder,
		stage.NoSequenceTypes_reference,
		&stage.NoSequenceTypes_referenceOrder,
		stage.NoSequenceTypes_instance,
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
		stage.ValFLTypes,
		stage.ValFLType_stagedOrder,
		stage.ValFLTypes_reference,
		&stage.ValFLTypes_referenceOrder,
		stage.ValFLTypes_instance,
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
		stage.ValSpeedTypes,
		stage.ValSpeedType_stagedOrder,
		stage.ValSpeedTypes_reference,
		&stage.ValSpeedTypes_referenceOrder,
		stage.ValSpeedTypes_instance,
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
		stage.ValWeightTypes,
		stage.ValWeightType_stagedOrder,
		stage.ValWeightTypes_reference,
		&stage.ValWeightTypes_referenceOrder,
		stage.ValWeightTypes_instance,
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
	__gong__computeReferencePass1(stage, stage.CodeAircraftCategoryTypes, &stage.CodeAircraftCategoryTypes_reference, &stage.CodeAircraftCategoryTypes_referenceOrder, &stage.CodeAircraftCategoryTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAircraftEngineNumberTypes, &stage.CodeAircraftEngineNumberTypes_reference, &stage.CodeAircraftEngineNumberTypes_referenceOrder, &stage.CodeAircraftEngineNumberTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAircraftEngineTypes, &stage.CodeAircraftEngineTypes_reference, &stage.CodeAircraftEngineTypes_referenceOrder, &stage.CodeAircraftEngineTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAircraftICAOTypes, &stage.CodeAircraftICAOTypes_reference, &stage.CodeAircraftICAOTypes_referenceOrder, &stage.CodeAircraftICAOTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAircraftTypes, &stage.CodeAircraftTypes_reference, &stage.CodeAircraftTypes_referenceOrder, &stage.CodeAircraftTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAircraftWingspanClassTypes, &stage.CodeAircraftWingspanClassTypes_reference, &stage.CodeAircraftWingspanClassTypes_referenceOrder, &stage.CodeAircraftWingspanClassTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAirspaceActivityTypes, &stage.CodeAirspaceActivityTypes_reference, &stage.CodeAirspaceActivityTypes_referenceOrder, &stage.CodeAirspaceActivityTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAirspaceAggregationTypes, &stage.CodeAirspaceAggregationTypes_reference, &stage.CodeAirspaceAggregationTypes_referenceOrder, &stage.CodeAirspaceAggregationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAirspaceClassificationTypes, &stage.CodeAirspaceClassificationTypes_reference, &stage.CodeAirspaceClassificationTypes_referenceOrder, &stage.CodeAirspaceClassificationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAirspaceDependencyTypes, &stage.CodeAirspaceDependencyTypes_reference, &stage.CodeAirspaceDependencyTypes_referenceOrder, &stage.CodeAirspaceDependencyTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAirspaceDesignatorTypes, &stage.CodeAirspaceDesignatorTypes_reference, &stage.CodeAirspaceDesignatorTypes_referenceOrder, &stage.CodeAirspaceDesignatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAirspaceTypes, &stage.CodeAirspaceTypes_reference, &stage.CodeAirspaceTypes_referenceOrder, &stage.CodeAirspaceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAltitudeUseTypes, &stage.CodeAltitudeUseTypes_reference, &stage.CodeAltitudeUseTypes_referenceOrder, &stage.CodeAltitudeUseTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAuthorityTypes, &stage.CodeAuthorityTypes_reference, &stage.CodeAuthorityTypes_referenceOrder, &stage.CodeAuthorityTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeCommunicationModeTypes, &stage.CodeCommunicationModeTypes_reference, &stage.CodeCommunicationModeTypes_referenceOrder, &stage.CodeCommunicationModeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDayTypes, &stage.CodeDayTypes_reference, &stage.CodeDayTypes_referenceOrder, &stage.CodeDayTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeEquipmentAntiCollisionTypes, &stage.CodeEquipmentAntiCollisionTypes_reference, &stage.CodeEquipmentAntiCollisionTypes_referenceOrder, &stage.CodeEquipmentAntiCollisionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeGeoBorderTypes, &stage.CodeGeoBorderTypes_reference, &stage.CodeGeoBorderTypes_referenceOrder, &stage.CodeGeoBorderTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMilitaryOperationsTypes, &stage.CodeMilitaryOperationsTypes_reference, &stage.CodeMilitaryOperationsTypes_referenceOrder, &stage.CodeMilitaryOperationsTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNavigationEquipmentTypes, &stage.CodeNavigationEquipmentTypes_reference, &stage.CodeNavigationEquipmentTypes_referenceOrder, &stage.CodeNavigationEquipmentTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNavigationSpecificationTypes, &stage.CodeNavigationSpecificationTypes_reference, &stage.CodeNavigationSpecificationTypes_referenceOrder, &stage.CodeNavigationSpecificationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNotePurposeTypes, &stage.CodeNotePurposeTypes_reference, &stage.CodeNotePurposeTypes_referenceOrder, &stage.CodeNotePurposeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRVSMTypes, &stage.CodeRVSMTypes_reference, &stage.CodeRVSMTypes_referenceOrder, &stage.CodeRVSMTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeStatusAirspaceTypes, &stage.CodeStatusAirspaceTypes_reference, &stage.CodeStatusAirspaceTypes_referenceOrder, &stage.CodeStatusAirspaceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTimeEventCombinationTypes, &stage.CodeTimeEventCombinationTypes_reference, &stage.CodeTimeEventCombinationTypes_referenceOrder, &stage.CodeTimeEventCombinationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTimeEventTypes, &stage.CodeTimeEventTypes_reference, &stage.CodeTimeEventTypes_referenceOrder, &stage.CodeTimeEventTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTimeReferenceTypes, &stage.CodeTimeReferenceTypes_reference, &stage.CodeTimeReferenceTypes_referenceOrder, &stage.CodeTimeReferenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTransponderTypes, &stage.CodeTransponderTypes_reference, &stage.CodeTransponderTypes_referenceOrder, &stage.CodeTransponderTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeValueInterpretationTypes, &stage.CodeValueInterpretationTypes_reference, &stage.CodeValueInterpretationTypes_referenceOrder, &stage.CodeValueInterpretationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeVerticalDatumTypes, &stage.CodeVerticalDatumTypes_reference, &stage.CodeVerticalDatumTypes_referenceOrder, &stage.CodeVerticalDatumTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeVerticalReferenceTypes, &stage.CodeVerticalReferenceTypes_reference, &stage.CodeVerticalReferenceTypes_referenceOrder, &stage.CodeVerticalReferenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeWakeTurbulenceTypes, &stage.CodeWakeTurbulenceTypes_reference, &stage.CodeWakeTurbulenceTypes_referenceOrder, &stage.CodeWakeTurbulenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeYesNoTypes, &stage.CodeYesNoTypes_reference, &stage.CodeYesNoTypes_referenceOrder, &stage.CodeYesNoTypes_instance)

	__gong__computeReferencePass1(stage, stage.DateMonthDayTypes, &stage.DateMonthDayTypes_reference, &stage.DateMonthDayTypes_referenceOrder, &stage.DateMonthDayTypes_instance)

	__gong__computeReferencePass1(stage, stage.NoNumberTypes, &stage.NoNumberTypes_reference, &stage.NoNumberTypes_referenceOrder, &stage.NoNumberTypes_instance)

	__gong__computeReferencePass1(stage, stage.NoSequenceTypes, &stage.NoSequenceTypes_reference, &stage.NoSequenceTypes_referenceOrder, &stage.NoSequenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextNameTypes, &stage.TextNameTypes_reference, &stage.TextNameTypes_referenceOrder, &stage.TextNameTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextNoteTypes, &stage.TextNoteTypes_reference, &stage.TextNoteTypes_referenceOrder, &stage.TextNoteTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextPropertyNameTypes, &stage.TextPropertyNameTypes_reference, &stage.TextPropertyNameTypes_referenceOrder, &stage.TextPropertyNameTypes_instance)

	__gong__computeReferencePass1(stage, stage.TimeTypes, &stage.TimeTypes_reference, &stage.TimeTypes_referenceOrder, &stage.TimeTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDistanceSignedTypes, &stage.ValDistanceSignedTypes_reference, &stage.ValDistanceSignedTypes_referenceOrder, &stage.ValDistanceSignedTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDistanceTypes, &stage.ValDistanceTypes_reference, &stage.ValDistanceTypes_referenceOrder, &stage.ValDistanceTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDistanceVerticalTypes, &stage.ValDistanceVerticalTypes_reference, &stage.ValDistanceVerticalTypes_referenceOrder, &stage.ValDistanceVerticalTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDurationTypes, &stage.ValDurationTypes_reference, &stage.ValDurationTypes_referenceOrder, &stage.ValDurationTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValFLTypes, &stage.ValFLTypes_reference, &stage.ValFLTypes_referenceOrder, &stage.ValFLTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValSpeedTypes, &stage.ValSpeedTypes_reference, &stage.ValSpeedTypes_referenceOrder, &stage.ValSpeedTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValWeightTypes, &stage.ValWeightTypes_reference, &stage.ValWeightTypes_referenceOrder, &stage.ValWeightTypes_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.CodeAircraftCategoryTypes, stage.CodeAircraftCategoryTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAircraftEngineNumberTypes, stage.CodeAircraftEngineNumberTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAircraftEngineTypes, stage.CodeAircraftEngineTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAircraftICAOTypes, stage.CodeAircraftICAOTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAircraftTypes, stage.CodeAircraftTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAircraftWingspanClassTypes, stage.CodeAircraftWingspanClassTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAirspaceActivityTypes, stage.CodeAirspaceActivityTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAirspaceAggregationTypes, stage.CodeAirspaceAggregationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAirspaceClassificationTypes, stage.CodeAirspaceClassificationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAirspaceDependencyTypes, stage.CodeAirspaceDependencyTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAirspaceDesignatorTypes, stage.CodeAirspaceDesignatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAirspaceTypes, stage.CodeAirspaceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAltitudeUseTypes, stage.CodeAltitudeUseTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAuthorityTypes, stage.CodeAuthorityTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeCommunicationModeTypes, stage.CodeCommunicationModeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDayTypes, stage.CodeDayTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeEquipmentAntiCollisionTypes, stage.CodeEquipmentAntiCollisionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeGeoBorderTypes, stage.CodeGeoBorderTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMilitaryOperationsTypes, stage.CodeMilitaryOperationsTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNavigationEquipmentTypes, stage.CodeNavigationEquipmentTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNavigationSpecificationTypes, stage.CodeNavigationSpecificationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNotePurposeTypes, stage.CodeNotePurposeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRVSMTypes, stage.CodeRVSMTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeStatusAirspaceTypes, stage.CodeStatusAirspaceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTimeEventCombinationTypes, stage.CodeTimeEventCombinationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTimeEventTypes, stage.CodeTimeEventTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTimeReferenceTypes, stage.CodeTimeReferenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTransponderTypes, stage.CodeTransponderTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeValueInterpretationTypes, stage.CodeValueInterpretationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeVerticalDatumTypes, stage.CodeVerticalDatumTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeVerticalReferenceTypes, stage.CodeVerticalReferenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeWakeTurbulenceTypes, stage.CodeWakeTurbulenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeYesNoTypes, stage.CodeYesNoTypes_reference, stage)

	__gong__computeReferencePass2(stage.DateMonthDayTypes, stage.DateMonthDayTypes_reference, stage)

	__gong__computeReferencePass2(stage.NoNumberTypes, stage.NoNumberTypes_reference, stage)

	__gong__computeReferencePass2(stage.NoSequenceTypes, stage.NoSequenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextNameTypes, stage.TextNameTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextNoteTypes, stage.TextNoteTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextPropertyNameTypes, stage.TextPropertyNameTypes_reference, stage)

	__gong__computeReferencePass2(stage.TimeTypes, stage.TimeTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDistanceSignedTypes, stage.ValDistanceSignedTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDistanceTypes, stage.ValDistanceTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDistanceVerticalTypes, stage.ValDistanceVerticalTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDurationTypes, stage.ValDurationTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValFLTypes, stage.ValFLTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValSpeedTypes, stage.ValSpeedTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValWeightTypes, stage.ValWeightTypes_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (codeaircraftcategorytype *CodeAircraftCategoryType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAircraftCategoryType_stagedOrder, stage.CodeAircraftCategoryTypes_referenceOrder, codeaircraftcategorytype, "CodeAircraftCategoryType")
}

func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAircraftEngineNumberType_stagedOrder, stage.CodeAircraftEngineNumberTypes_referenceOrder, codeaircraftenginenumbertype, "CodeAircraftEngineNumberType")
}

func (codeaircraftenginetype *CodeAircraftEngineType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAircraftEngineType_stagedOrder, stage.CodeAircraftEngineTypes_referenceOrder, codeaircraftenginetype, "CodeAircraftEngineType")
}

func (codeaircrafticaotype *CodeAircraftICAOType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAircraftICAOType_stagedOrder, stage.CodeAircraftICAOTypes_referenceOrder, codeaircrafticaotype, "CodeAircraftICAOType")
}

func (codeaircrafttype *CodeAircraftType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAircraftType_stagedOrder, stage.CodeAircraftTypes_referenceOrder, codeaircrafttype, "CodeAircraftType")
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAircraftWingspanClassType_stagedOrder, stage.CodeAircraftWingspanClassTypes_referenceOrder, codeaircraftwingspanclasstype, "CodeAircraftWingspanClassType")
}

func (codeairspaceactivitytype *CodeAirspaceActivityType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAirspaceActivityType_stagedOrder, stage.CodeAirspaceActivityTypes_referenceOrder, codeairspaceactivitytype, "CodeAirspaceActivityType")
}

func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAirspaceAggregationType_stagedOrder, stage.CodeAirspaceAggregationTypes_referenceOrder, codeairspaceaggregationtype, "CodeAirspaceAggregationType")
}

func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAirspaceClassificationType_stagedOrder, stage.CodeAirspaceClassificationTypes_referenceOrder, codeairspaceclassificationtype, "CodeAirspaceClassificationType")
}

func (codeairspacedependencytype *CodeAirspaceDependencyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAirspaceDependencyType_stagedOrder, stage.CodeAirspaceDependencyTypes_referenceOrder, codeairspacedependencytype, "CodeAirspaceDependencyType")
}

func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAirspaceDesignatorType_stagedOrder, stage.CodeAirspaceDesignatorTypes_referenceOrder, codeairspacedesignatortype, "CodeAirspaceDesignatorType")
}

func (codeairspacetype *CodeAirspaceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAirspaceType_stagedOrder, stage.CodeAirspaceTypes_referenceOrder, codeairspacetype, "CodeAirspaceType")
}

func (codealtitudeusetype *CodeAltitudeUseType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAltitudeUseType_stagedOrder, stage.CodeAltitudeUseTypes_referenceOrder, codealtitudeusetype, "CodeAltitudeUseType")
}

func (codeauthoritytype *CodeAuthorityType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAuthorityType_stagedOrder, stage.CodeAuthorityTypes_referenceOrder, codeauthoritytype, "CodeAuthorityType")
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeCommunicationModeType_stagedOrder, stage.CodeCommunicationModeTypes_referenceOrder, codecommunicationmodetype, "CodeCommunicationModeType")
}

func (codedaytype *CodeDayType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeDayType_stagedOrder, stage.CodeDayTypes_referenceOrder, codedaytype, "CodeDayType")
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeEquipmentAntiCollisionType_stagedOrder, stage.CodeEquipmentAntiCollisionTypes_referenceOrder, codeequipmentanticollisiontype, "CodeEquipmentAntiCollisionType")
}

func (codegeobordertype *CodeGeoBorderType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeGeoBorderType_stagedOrder, stage.CodeGeoBorderTypes_referenceOrder, codegeobordertype, "CodeGeoBorderType")
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeMilitaryOperationsType_stagedOrder, stage.CodeMilitaryOperationsTypes_referenceOrder, codemilitaryoperationstype, "CodeMilitaryOperationsType")
}

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeNavigationEquipmentType_stagedOrder, stage.CodeNavigationEquipmentTypes_referenceOrder, codenavigationequipmenttype, "CodeNavigationEquipmentType")
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeNavigationSpecificationType_stagedOrder, stage.CodeNavigationSpecificationTypes_referenceOrder, codenavigationspecificationtype, "CodeNavigationSpecificationType")
}

func (codenotepurposetype *CodeNotePurposeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeNotePurposeType_stagedOrder, stage.CodeNotePurposeTypes_referenceOrder, codenotepurposetype, "CodeNotePurposeType")
}

func (codervsmtype *CodeRVSMType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRVSMType_stagedOrder, stage.CodeRVSMTypes_referenceOrder, codervsmtype, "CodeRVSMType")
}

func (codestatusairspacetype *CodeStatusAirspaceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeStatusAirspaceType_stagedOrder, stage.CodeStatusAirspaceTypes_referenceOrder, codestatusairspacetype, "CodeStatusAirspaceType")
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

func (codetranspondertype *CodeTransponderType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeTransponderType_stagedOrder, stage.CodeTransponderTypes_referenceOrder, codetranspondertype, "CodeTransponderType")
}

func (codevalueinterpretationtype *CodeValueInterpretationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeValueInterpretationType_stagedOrder, stage.CodeValueInterpretationTypes_referenceOrder, codevalueinterpretationtype, "CodeValueInterpretationType")
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeVerticalDatumType_stagedOrder, stage.CodeVerticalDatumTypes_referenceOrder, codeverticaldatumtype, "CodeVerticalDatumType")
}

func (codeverticalreferencetype *CodeVerticalReferenceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeVerticalReferenceType_stagedOrder, stage.CodeVerticalReferenceTypes_referenceOrder, codeverticalreferencetype, "CodeVerticalReferenceType")
}

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeWakeTurbulenceType_stagedOrder, stage.CodeWakeTurbulenceTypes_referenceOrder, codewaketurbulencetype, "CodeWakeTurbulenceType")
}

func (codeyesnotype *CodeYesNoType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeYesNoType_stagedOrder, stage.CodeYesNoTypes_referenceOrder, codeyesnotype, "CodeYesNoType")
}

func (datemonthdaytype *DateMonthDayType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DateMonthDayType_stagedOrder, stage.DateMonthDayTypes_referenceOrder, datemonthdaytype, "DateMonthDayType")
}

func (nonumbertype *NoNumberType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoNumberType_stagedOrder, stage.NoNumberTypes_referenceOrder, nonumbertype, "NoNumberType")
}

func (nosequencetype *NoSequenceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoSequenceType_stagedOrder, stage.NoSequenceTypes_referenceOrder, nosequencetype, "NoSequenceType")
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

func (valfltype *ValFLType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValFLType_stagedOrder, stage.ValFLTypes_referenceOrder, valfltype, "ValFLType")
}

func (valspeedtype *ValSpeedType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValSpeedType_stagedOrder, stage.ValSpeedTypes_referenceOrder, valspeedtype, "ValSpeedType")
}

func (valweighttype *ValWeightType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValWeightType_stagedOrder, stage.ValWeightTypes_referenceOrder, valweighttype, "ValWeightType")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (codeaircraftcategorytype *CodeAircraftCategoryType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeaircraftcategorytype, codeaircraftcategorytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeaircraftcategorytype *CodeAircraftCategoryType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeaircraftcategorytype.GongGetIdentifier(stage)
}

func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeaircraftenginenumbertype, codeaircraftenginenumbertype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeaircraftenginenumbertype.GongGetIdentifier(stage)
}

func (codeaircraftenginetype *CodeAircraftEngineType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeaircraftenginetype, codeaircraftenginetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeaircraftenginetype *CodeAircraftEngineType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeaircraftenginetype.GongGetIdentifier(stage)
}

func (codeaircrafticaotype *CodeAircraftICAOType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeaircrafticaotype, codeaircrafticaotype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeaircrafticaotype *CodeAircraftICAOType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeaircrafticaotype.GongGetIdentifier(stage)
}

func (codeaircrafttype *CodeAircraftType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeaircrafttype, codeaircrafttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeaircrafttype *CodeAircraftType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeaircrafttype.GongGetIdentifier(stage)
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeaircraftwingspanclasstype, codeaircraftwingspanclasstype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeaircraftwingspanclasstype.GongGetIdentifier(stage)
}

func (codeairspaceactivitytype *CodeAirspaceActivityType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeairspaceactivitytype, codeairspaceactivitytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeairspaceactivitytype *CodeAirspaceActivityType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeairspaceactivitytype.GongGetIdentifier(stage)
}

func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeairspaceaggregationtype, codeairspaceaggregationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeairspaceaggregationtype.GongGetIdentifier(stage)
}

func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeairspaceclassificationtype, codeairspaceclassificationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeairspaceclassificationtype.GongGetIdentifier(stage)
}

func (codeairspacedependencytype *CodeAirspaceDependencyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeairspacedependencytype, codeairspacedependencytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeairspacedependencytype *CodeAirspaceDependencyType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeairspacedependencytype.GongGetIdentifier(stage)
}

func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeairspacedesignatortype, codeairspacedesignatortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeairspacedesignatortype.GongGetIdentifier(stage)
}

func (codeairspacetype *CodeAirspaceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeairspacetype, codeairspacetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeairspacetype *CodeAirspaceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeairspacetype.GongGetIdentifier(stage)
}

func (codealtitudeusetype *CodeAltitudeUseType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codealtitudeusetype, codealtitudeusetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codealtitudeusetype *CodeAltitudeUseType) GongGetReferenceIdentifier(stage *Stage) string {
	return codealtitudeusetype.GongGetIdentifier(stage)
}

func (codeauthoritytype *CodeAuthorityType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeauthoritytype, codeauthoritytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeauthoritytype *CodeAuthorityType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeauthoritytype.GongGetIdentifier(stage)
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codecommunicationmodetype, codecommunicationmodetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codecommunicationmodetype *CodeCommunicationModeType) GongGetReferenceIdentifier(stage *Stage) string {
	return codecommunicationmodetype.GongGetIdentifier(stage)
}

func (codedaytype *CodeDayType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codedaytype, codedaytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codedaytype *CodeDayType) GongGetReferenceIdentifier(stage *Stage) string {
	return codedaytype.GongGetIdentifier(stage)
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeequipmentanticollisiontype, codeequipmentanticollisiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeequipmentanticollisiontype.GongGetIdentifier(stage)
}

func (codegeobordertype *CodeGeoBorderType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codegeobordertype, codegeobordertype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codegeobordertype *CodeGeoBorderType) GongGetReferenceIdentifier(stage *Stage) string {
	return codegeobordertype.GongGetIdentifier(stage)
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codemilitaryoperationstype, codemilitaryoperationstype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongGetReferenceIdentifier(stage *Stage) string {
	return codemilitaryoperationstype.GongGetIdentifier(stage)
}

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codenavigationequipmenttype, codenavigationequipmenttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongGetReferenceIdentifier(stage *Stage) string {
	return codenavigationequipmenttype.GongGetIdentifier(stage)
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codenavigationspecificationtype, codenavigationspecificationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongGetReferenceIdentifier(stage *Stage) string {
	return codenavigationspecificationtype.GongGetIdentifier(stage)
}

func (codenotepurposetype *CodeNotePurposeType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codenotepurposetype, codenotepurposetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codenotepurposetype *CodeNotePurposeType) GongGetReferenceIdentifier(stage *Stage) string {
	return codenotepurposetype.GongGetIdentifier(stage)
}

func (codervsmtype *CodeRVSMType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codervsmtype, codervsmtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codervsmtype *CodeRVSMType) GongGetReferenceIdentifier(stage *Stage) string {
	return codervsmtype.GongGetIdentifier(stage)
}

func (codestatusairspacetype *CodeStatusAirspaceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codestatusairspacetype, codestatusairspacetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codestatusairspacetype *CodeStatusAirspaceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codestatusairspacetype.GongGetIdentifier(stage)
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

func (codetranspondertype *CodeTransponderType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codetranspondertype, codetranspondertype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codetranspondertype *CodeTransponderType) GongGetReferenceIdentifier(stage *Stage) string {
	return codetranspondertype.GongGetIdentifier(stage)
}

func (codevalueinterpretationtype *CodeValueInterpretationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codevalueinterpretationtype, codevalueinterpretationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codevalueinterpretationtype *CodeValueInterpretationType) GongGetReferenceIdentifier(stage *Stage) string {
	return codevalueinterpretationtype.GongGetIdentifier(stage)
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeverticaldatumtype, codeverticaldatumtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeverticaldatumtype *CodeVerticalDatumType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeverticaldatumtype.GongGetIdentifier(stage)
}

func (codeverticalreferencetype *CodeVerticalReferenceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeverticalreferencetype, codeverticalreferencetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeverticalreferencetype *CodeVerticalReferenceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeverticalreferencetype.GongGetIdentifier(stage)
}

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codewaketurbulencetype, codewaketurbulencetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codewaketurbulencetype *CodeWakeTurbulenceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codewaketurbulencetype.GongGetIdentifier(stage)
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

func (nonumbertype *NoNumberType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(nonumbertype, nonumbertype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (nonumbertype *NoNumberType) GongGetReferenceIdentifier(stage *Stage) string {
	return nonumbertype.GongGetIdentifier(stage)
}

func (nosequencetype *NoSequenceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(nosequencetype, nosequencetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (nosequencetype *NoSequenceType) GongGetReferenceIdentifier(stage *Stage) string {
	return nosequencetype.GongGetIdentifier(stage)
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

func (valfltype *ValFLType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valfltype, valfltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valfltype *ValFLType) GongGetReferenceIdentifier(stage *Stage) string {
	return valfltype.GongGetIdentifier(stage)
}

func (valspeedtype *ValSpeedType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valspeedtype, valspeedtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valspeedtype *ValSpeedType) GongGetReferenceIdentifier(stage *Stage) string {
	return valspeedtype.GongGetIdentifier(stage)
}

func (valweighttype *ValWeightType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valweighttype, valweighttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valweighttype *ValWeightType) GongGetReferenceIdentifier(stage *Stage) string {
	return valweighttype.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (codeaircraftcategorytype *CodeAircraftCategoryType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeaircraftcategorytype.GongGetIdentifier(stage), "CodeAircraftCategoryType", codeaircraftcategorytype.Name)
}

func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeaircraftenginenumbertype.GongGetIdentifier(stage), "CodeAircraftEngineNumberType", codeaircraftenginenumbertype.Name)
}

func (codeaircraftenginetype *CodeAircraftEngineType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeaircraftenginetype.GongGetIdentifier(stage), "CodeAircraftEngineType", codeaircraftenginetype.Name)
}

func (codeaircrafticaotype *CodeAircraftICAOType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeaircrafticaotype.GongGetIdentifier(stage), "CodeAircraftICAOType", codeaircrafticaotype.Name)
}

func (codeaircrafttype *CodeAircraftType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeaircrafttype.GongGetIdentifier(stage), "CodeAircraftType", codeaircrafttype.Name)
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeaircraftwingspanclasstype.GongGetIdentifier(stage), "CodeAircraftWingspanClassType", codeaircraftwingspanclasstype.Name)
}

func (codeairspaceactivitytype *CodeAirspaceActivityType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeairspaceactivitytype.GongGetIdentifier(stage), "CodeAirspaceActivityType", codeairspaceactivitytype.Name)
}

func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeairspaceaggregationtype.GongGetIdentifier(stage), "CodeAirspaceAggregationType", codeairspaceaggregationtype.Name)
}

func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeairspaceclassificationtype.GongGetIdentifier(stage), "CodeAirspaceClassificationType", codeairspaceclassificationtype.Name)
}

func (codeairspacedependencytype *CodeAirspaceDependencyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeairspacedependencytype.GongGetIdentifier(stage), "CodeAirspaceDependencyType", codeairspacedependencytype.Name)
}

func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeairspacedesignatortype.GongGetIdentifier(stage), "CodeAirspaceDesignatorType", codeairspacedesignatortype.Name)
}

func (codeairspacetype *CodeAirspaceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeairspacetype.GongGetIdentifier(stage), "CodeAirspaceType", codeairspacetype.Name)
}

func (codealtitudeusetype *CodeAltitudeUseType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codealtitudeusetype.GongGetIdentifier(stage), "CodeAltitudeUseType", codealtitudeusetype.Name)
}

func (codeauthoritytype *CodeAuthorityType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeauthoritytype.GongGetIdentifier(stage), "CodeAuthorityType", codeauthoritytype.Name)
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codecommunicationmodetype.GongGetIdentifier(stage), "CodeCommunicationModeType", codecommunicationmodetype.Name)
}

func (codedaytype *CodeDayType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codedaytype.GongGetIdentifier(stage), "CodeDayType", codedaytype.Name)
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeequipmentanticollisiontype.GongGetIdentifier(stage), "CodeEquipmentAntiCollisionType", codeequipmentanticollisiontype.Name)
}

func (codegeobordertype *CodeGeoBorderType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codegeobordertype.GongGetIdentifier(stage), "CodeGeoBorderType", codegeobordertype.Name)
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codemilitaryoperationstype.GongGetIdentifier(stage), "CodeMilitaryOperationsType", codemilitaryoperationstype.Name)
}

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codenavigationequipmenttype.GongGetIdentifier(stage), "CodeNavigationEquipmentType", codenavigationequipmenttype.Name)
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codenavigationspecificationtype.GongGetIdentifier(stage), "CodeNavigationSpecificationType", codenavigationspecificationtype.Name)
}

func (codenotepurposetype *CodeNotePurposeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codenotepurposetype.GongGetIdentifier(stage), "CodeNotePurposeType", codenotepurposetype.Name)
}

func (codervsmtype *CodeRVSMType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codervsmtype.GongGetIdentifier(stage), "CodeRVSMType", codervsmtype.Name)
}

func (codestatusairspacetype *CodeStatusAirspaceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codestatusairspacetype.GongGetIdentifier(stage), "CodeStatusAirspaceType", codestatusairspacetype.Name)
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

func (codetranspondertype *CodeTransponderType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codetranspondertype.GongGetIdentifier(stage), "CodeTransponderType", codetranspondertype.Name)
}

func (codevalueinterpretationtype *CodeValueInterpretationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codevalueinterpretationtype.GongGetIdentifier(stage), "CodeValueInterpretationType", codevalueinterpretationtype.Name)
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeverticaldatumtype.GongGetIdentifier(stage), "CodeVerticalDatumType", codeverticaldatumtype.Name)
}

func (codeverticalreferencetype *CodeVerticalReferenceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeverticalreferencetype.GongGetIdentifier(stage), "CodeVerticalReferenceType", codeverticalreferencetype.Name)
}

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codewaketurbulencetype.GongGetIdentifier(stage), "CodeWakeTurbulenceType", codewaketurbulencetype.Name)
}

func (codeyesnotype *CodeYesNoType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeyesnotype.GongGetIdentifier(stage), "CodeYesNoType", codeyesnotype.Name)
}

func (datemonthdaytype *DateMonthDayType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(datemonthdaytype.GongGetIdentifier(stage), "DateMonthDayType", datemonthdaytype.Name)
}

func (nonumbertype *NoNumberType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(nonumbertype.GongGetIdentifier(stage), "NoNumberType", nonumbertype.Name)
}

func (nosequencetype *NoSequenceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(nosequencetype.GongGetIdentifier(stage), "NoSequenceType", nosequencetype.Name)
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

func (valfltype *ValFLType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valfltype.GongGetIdentifier(stage), "ValFLType", valfltype.Name)
}

func (valspeedtype *ValSpeedType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valspeedtype.GongGetIdentifier(stage), "ValSpeedType", valspeedtype.Name)
}

func (valweighttype *ValWeightType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valweighttype.GongGetIdentifier(stage), "ValWeightType", valweighttype.Name)
}

// insertion point for unstaging
func (codeaircraftcategorytype *CodeAircraftCategoryType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeaircraftcategorytype.GongGetReferenceIdentifier(stage))
}

func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeaircraftenginenumbertype.GongGetReferenceIdentifier(stage))
}

func (codeaircraftenginetype *CodeAircraftEngineType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeaircraftenginetype.GongGetReferenceIdentifier(stage))
}

func (codeaircrafticaotype *CodeAircraftICAOType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeaircrafticaotype.GongGetReferenceIdentifier(stage))
}

func (codeaircrafttype *CodeAircraftType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeaircrafttype.GongGetReferenceIdentifier(stage))
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeaircraftwingspanclasstype.GongGetReferenceIdentifier(stage))
}

func (codeairspaceactivitytype *CodeAirspaceActivityType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeairspaceactivitytype.GongGetReferenceIdentifier(stage))
}

func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeairspaceaggregationtype.GongGetReferenceIdentifier(stage))
}

func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeairspaceclassificationtype.GongGetReferenceIdentifier(stage))
}

func (codeairspacedependencytype *CodeAirspaceDependencyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeairspacedependencytype.GongGetReferenceIdentifier(stage))
}

func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeairspacedesignatortype.GongGetReferenceIdentifier(stage))
}

func (codeairspacetype *CodeAirspaceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeairspacetype.GongGetReferenceIdentifier(stage))
}

func (codealtitudeusetype *CodeAltitudeUseType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codealtitudeusetype.GongGetReferenceIdentifier(stage))
}

func (codeauthoritytype *CodeAuthorityType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeauthoritytype.GongGetReferenceIdentifier(stage))
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codecommunicationmodetype.GongGetReferenceIdentifier(stage))
}

func (codedaytype *CodeDayType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codedaytype.GongGetReferenceIdentifier(stage))
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeequipmentanticollisiontype.GongGetReferenceIdentifier(stage))
}

func (codegeobordertype *CodeGeoBorderType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codegeobordertype.GongGetReferenceIdentifier(stage))
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codemilitaryoperationstype.GongGetReferenceIdentifier(stage))
}

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codenavigationequipmenttype.GongGetReferenceIdentifier(stage))
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codenavigationspecificationtype.GongGetReferenceIdentifier(stage))
}

func (codenotepurposetype *CodeNotePurposeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codenotepurposetype.GongGetReferenceIdentifier(stage))
}

func (codervsmtype *CodeRVSMType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codervsmtype.GongGetReferenceIdentifier(stage))
}

func (codestatusairspacetype *CodeStatusAirspaceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codestatusairspacetype.GongGetReferenceIdentifier(stage))
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

func (codetranspondertype *CodeTransponderType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codetranspondertype.GongGetReferenceIdentifier(stage))
}

func (codevalueinterpretationtype *CodeValueInterpretationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codevalueinterpretationtype.GongGetReferenceIdentifier(stage))
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeverticaldatumtype.GongGetReferenceIdentifier(stage))
}

func (codeverticalreferencetype *CodeVerticalReferenceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeverticalreferencetype.GongGetReferenceIdentifier(stage))
}

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codewaketurbulencetype.GongGetReferenceIdentifier(stage))
}

func (codeyesnotype *CodeYesNoType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeyesnotype.GongGetReferenceIdentifier(stage))
}

func (datemonthdaytype *DateMonthDayType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(datemonthdaytype.GongGetReferenceIdentifier(stage))
}

func (nonumbertype *NoNumberType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(nonumbertype.GongGetReferenceIdentifier(stage))
}

func (nosequencetype *NoSequenceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(nosequencetype.GongGetReferenceIdentifier(stage))
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

func (valfltype *ValFLType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valfltype.GongGetReferenceIdentifier(stage))
}

func (valspeedtype *ValSpeedType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valspeedtype.GongGetReferenceIdentifier(stage))
}

func (valweighttype *ValWeightType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valweighttype.GongGetReferenceIdentifier(stage))
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
