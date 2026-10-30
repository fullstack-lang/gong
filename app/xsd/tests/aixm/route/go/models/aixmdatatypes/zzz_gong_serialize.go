// generated code - do not edit
package aixmdatatypes

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/xuri/excelize/v2"
)

func (stage *Stage) SerializeStage(filename string) {
	stage.SerializeStage2(filename, false)
}

func (stage *Stage) SerializeStage2(filename string, addIDs bool) {
	f := stage.__gong__buildExcelizeFile(addIDs)
	if err := f.SaveAs(filename); err != nil {
		fmt.Println("cannot write xl file : ", err)
	}
}

func (stage *Stage) __gong__buildExcelizeFile(addIDs bool) *excelize.File {
	f := excelize.NewFile()
	{
		// insertion point
		{
			var instances []GongstructIF
			for instance := range stage.CodeATCReportingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeATCReportingType", instances, (*CodeATCReportingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAerialRefuellingPointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAerialRefuellingPointType", instances, (*CodeAerialRefuellingPointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAerialRefuellingPrefixTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAerialRefuellingPrefixType", instances, (*CodeAerialRefuellingPrefixType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAerialRefuellingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAerialRefuellingType", instances, (*CodeAerialRefuellingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAircraftCategoryTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAircraftCategoryType", instances, (*CodeAircraftCategoryType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAircraftEngineNumberTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAircraftEngineNumberType", instances, (*CodeAircraftEngineNumberType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAircraftEngineTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAircraftEngineType", instances, (*CodeAircraftEngineType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAircraftICAOTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAircraftICAOType", instances, (*CodeAircraftICAOType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAircraftTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAircraftType", instances, (*CodeAircraftType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAircraftWingspanClassTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAircraftWingspanClassType", instances, (*CodeAircraftWingspanClassType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAltitudeAdjustmentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAltitudeAdjustmentType", instances, (*CodeAltitudeAdjustmentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAltitudeUseTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAltitudeUseType", instances, (*CodeAltitudeUseType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAuthorityTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAuthorityType", instances, (*CodeAuthorityType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeCardinalDirectionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeCardinalDirectionType", instances, (*CodeCardinalDirectionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeCommunicationModeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeCommunicationModeType", instances, (*CodeCommunicationModeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeComparisonTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeComparisonType", instances, (*CodeComparisonType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeCourseTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeCourseType", instances, (*CodeCourseType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeDayTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeDayType", instances, (*CodeDayType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeDirectionTurnTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeDirectionTurnType", instances, (*CodeDirectionTurnType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeDirectionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeDirectionType", instances, (*CodeDirectionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeEquipmentAntiCollisionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeEquipmentAntiCollisionType", instances, (*CodeEquipmentAntiCollisionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeFlightOriginTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeFlightOriginType", instances, (*CodeFlightOriginType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeFlightPurposeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeFlightPurposeType", instances, (*CodeFlightPurposeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeFlightRestrictionDesignatorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeFlightRestrictionDesignatorType", instances, (*CodeFlightRestrictionDesignatorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeFlightRestrictionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeFlightRestrictionType", instances, (*CodeFlightRestrictionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeFlightRuleTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeFlightRuleType", instances, (*CodeFlightRuleType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeFlightStatusTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeFlightStatusType", instances, (*CodeFlightStatusType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeFlightTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeFlightType", instances, (*CodeFlightType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeFlowConditionOperationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeFlowConditionOperationType", instances, (*CodeFlowConditionOperationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeFreeFlightTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeFreeFlightType", instances, (*CodeFreeFlightType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeLevelTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeLevelType", instances, (*CodeLevelType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeLocationQualifierTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeLocationQualifierType", instances, (*CodeLocationQualifierType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeMeteoConditionsTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeMeteoConditionsType", instances, (*CodeMeteoConditionsType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeMilitaryRoutePointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeMilitaryRoutePointType", instances, (*CodeMilitaryRoutePointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeMilitaryStatusTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeMilitaryStatusType", instances, (*CodeMilitaryStatusType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeMilitaryTrainingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeMilitaryTrainingType", instances, (*CodeMilitaryTrainingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeNavigationEquipmentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeNavigationEquipmentType", instances, (*CodeNavigationEquipmentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeNavigationSpecificationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeNavigationSpecificationType", instances, (*CodeNavigationSpecificationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeNotePurposeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeNotePurposeType", instances, (*CodeNotePurposeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeObstacleAssessmentSurfaceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeObstacleAssessmentSurfaceType", instances, (*CodeObstacleAssessmentSurfaceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeObstructionIdSurfaceZoneTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeObstructionIdSurfaceZoneType", instances, (*CodeObstructionIdSurfaceZoneType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRNPTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRNPType", instances, (*CodeRNPType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRVSMPointRoleTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRVSMPointRoleType", instances, (*CodeRVSMPointRoleType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRVSMTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRVSMType", instances, (*CodeRVSMType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeReferenceRoleTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeReferenceRoleType", instances, (*CodeReferenceRoleType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRouteAvailabilityTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRouteAvailabilityType", instances, (*CodeRouteAvailabilityType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRouteDesignatorLetterTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRouteDesignatorLetterType", instances, (*CodeRouteDesignatorLetterType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRouteDesignatorPrefixTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRouteDesignatorPrefixType", instances, (*CodeRouteDesignatorPrefixType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRouteDesignatorSuffixTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRouteDesignatorSuffixType", instances, (*CodeRouteDesignatorSuffixType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRouteNavigationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRouteNavigationType", instances, (*CodeRouteNavigationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRouteOriginTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRouteOriginType", instances, (*CodeRouteOriginType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRouteSegmentPathTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRouteSegmentPathType", instances, (*CodeRouteSegmentPathType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRouteTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRouteType", instances, (*CodeRouteType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRuleProcedureTitleTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRuleProcedureTitleType", instances, (*CodeRuleProcedureTitleType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRuleProcedureTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRuleProcedureType", instances, (*CodeRuleProcedureType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeSideTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeSideType", instances, (*CodeSideType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeSpeedReferenceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeSpeedReferenceType", instances, (*CodeSpeedReferenceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeTACANChannelTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeTACANChannelType", instances, (*CodeTACANChannelType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeTelecomNetworkTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeTelecomNetworkType", instances, (*CodeTelecomNetworkType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeTimeEventCombinationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeTimeEventCombinationType", instances, (*CodeTimeEventCombinationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeTimeEventTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeTimeEventType", instances, (*CodeTimeEventType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeTimeReferenceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeTimeReferenceType", instances, (*CodeTimeReferenceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeTransponderTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeTransponderType", instances, (*CodeTransponderType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeUpperAlphaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeUpperAlphaType", instances, (*CodeUpperAlphaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeValueInterpretationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeValueInterpretationType", instances, (*CodeValueInterpretationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeVerticalDatumTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeVerticalDatumType", instances, (*CodeVerticalDatumType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeVerticalReferenceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeVerticalReferenceType", instances, (*CodeVerticalReferenceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeWakeTurbulenceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeWakeTurbulenceType", instances, (*CodeWakeTurbulenceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeYesNoTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeYesNoType", instances, (*CodeYesNoType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DateMonthDayTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DateMonthDayType", instances, (*DateMonthDayType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NoNumberTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NoNumberType", instances, (*NoNumberType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NoSequenceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NoSequenceType", instances, (*NoSequenceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TextAddressTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TextAddressType", instances, (*TextAddressType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TextDesignatorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TextDesignatorType", instances, (*TextDesignatorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TextInstructionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TextInstructionType", instances, (*TextInstructionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TextNameTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TextNameType", instances, (*TextNameType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TextNoteTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TextNoteType", instances, (*TextNoteType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TextPhoneTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TextPhoneType", instances, (*TextPhoneType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TextPropertyNameTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TextPropertyNameType", instances, (*TextPropertyNameType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimeType", instances, (*TimeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValAngleTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValAngleType", instances, (*ValAngleType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValBearingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValBearingType", instances, (*ValBearingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValDistanceSignedTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValDistanceSignedType", instances, (*ValDistanceSignedType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValDistanceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValDistanceType", instances, (*ValDistanceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValDistanceVerticalTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValDistanceVerticalType", instances, (*ValDistanceVerticalType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValDurationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValDurationType", instances, (*ValDurationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValSlopeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValSlopeType", instances, (*ValSlopeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValSpeedTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValSpeedType", instances, (*ValSpeedType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValWeightTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValWeightType", instances, (*ValWeightType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.XHTMLTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "XHTMLType", instances, (*XHTMLType)(nil).GongGetFieldHeaders(), addIDs)
		}
	}

	// Create a style with wrap text enabled
	wrapStyle, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			WrapText: true,
		},
	})
	_ = wrapStyle
	if err != nil {
		fmt.Println("failed to create style:", err)
		return f
	}

	// Create a style with bold text
	boldStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
	})
	_ = boldStyle
	if err != nil {
		fmt.Println("failed to create bold style:", err)
		return f
	}

	// Get all sheet names
	sheetList := f.GetSheetList()

	for _, sheet := range sheetList {
		// Use a lazy iterator instead of loading all rows into memory
		rows, err := f.Rows(sheet)
		if err != nil {
			fmt.Printf("failed to get rows iterator for sheet %q: %v\n", sheet, err)
			continue
		}

		// Check if there is at least one row, and move the iterator to it
		if !rows.Next() {
			rows.Close() // Always close iterators
			continue
		}

		// Read ONLY the first row
		firstRow, err := rows.Columns()

		// Close the iterator immediately since we don't need the rest of the sheet
		rows.Close()

		if err != nil {
			fmt.Printf("failed to get columns for sheet %q: %v\n", sheet, err)
			continue
		}

		// If the first row is completely empty, skip
		if len(firstRow) == 0 {
			continue
		}

		// Track the first and last “used” column in the first row,
		// so we can later apply an AutoFilter from the first to last used col
		var firstUsedColIdx, lastUsedColIdx int

		for colIdx, cellValue := range firstRow {
			if cellValue == "" {
				// Skip columns with empty first-row cells
				continue
			}

			// Convert zero-based colIdx to 1-based for Excelize,
			// then get the column name (A, B, C, etc.)
			colName, err := excelize.ColumnNumberToName(colIdx + 1)
			if err != nil {
				fmt.Printf("failed to convert column number: %v\n", err)
				continue
			}

			// Apply wrap-text style to this entire column
			colRange := colName + ":" + colName
			if err := f.SetColStyle(sheet, colRange, wrapStyle); err != nil {
				fmt.Printf("failed to set col style on %s: %v\n", colRange, err)
				continue
			}

			// Make the first row (cell in row 1) bold in this column
			cellRef := fmt.Sprintf("%s1", colName)
			if err := f.SetCellStyle(sheet, cellRef, cellRef, boldStyle); err != nil {
				fmt.Printf("failed to set cell style on %s: %v\n", cellRef, err)
				continue
			}

			// Update our “first used” and “last used” column indices
			if firstUsedColIdx == 0 {
				firstUsedColIdx = colIdx + 1
			}
			if colIdx+1 > lastUsedColIdx {
				lastUsedColIdx = colIdx + 1
			}
		}

		// If we found at least one non-empty column in row 1, enable AutoFilter
		if firstUsedColIdx != 0 && lastUsedColIdx >= firstUsedColIdx {
			startCol, _ := excelize.ColumnNumberToName(firstUsedColIdx)
			endCol, _ := excelize.ColumnNumberToName(lastUsedColIdx)
			styleRange := fmt.Sprintf("%s:%s", startCol, endCol)
			autoFilterRange := fmt.Sprintf("%s1:%s1", startCol, endCol)
			startCellString := fmt.Sprintf("%s1", startCol)
			endCellString := fmt.Sprintf("%s1", endCol)

			if err := f.SetColStyle(sheet, styleRange, wrapStyle); err != nil {
				fmt.Println("failed to set column style:", err)
				return f
			}

			// Apply the bold style to the first row (A1:XFD1)
			if err := f.SetCellStyle(sheet, startCellString, endCellString, boldStyle); err != nil {
				fmt.Println("failed to set bold style:", err)
				return f
			}

			var opts []excelize.AutoFilterOptions
			if err := f.AutoFilter(sheet, autoFilterRange, opts); err != nil {
				fmt.Printf("failed to enable auto filter on range %s: %v\n", autoFilterRange, err)
			}
		}
	}

	var tab GongExcelizeTabulator
	tab.SetExcelizeFile(f)
	{
		f.DeleteSheet("Sheet1")
	}
	return f
}

// SerializeStageAsBytes serializes the stage to a pure in-memory Excel file and returns the bytes.
func (stage *Stage) SerializeStageAsBytes(addIDs bool) ([]byte, error) {
	f := stage.__gong__buildExcelizeFile(addIDs)
	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func __gong__shortenString(s string) string {
	if len(s) > 31 {
		return s[:31]
	}
	return s
}

// GongTabulator is an interface for writing to a table strings
type GongTabulator interface {
	AddSheet(sheetName string)
	AddRow(sheetName string) int
	AddCell(sheetName string, rowId, columnIndex int, value string)
}

type Tabulator = GongTabulator

type GongExcelizeTabulator struct {
	f *excelize.File
}

type ExcelizeTabulator = GongExcelizeTabulator

func (tab *GongExcelizeTabulator) SetExcelizeFile(f *excelize.File) {
	tab.f = f
}

func (tab *GongExcelizeTabulator) AddSheet(sheetName string) {

}

func (tab *GongExcelizeTabulator) AddRow(sheetName string) (rowId int) {
	return
}

func (tab *GongExcelizeTabulator) AddCell(sheetName string, rowId, columnIndex int, value string) {

}

// SerializeExcelize is the Stage method for Excel serialization with optional IDs.
func (stage *Stage) SerializeExcelize(f *excelize.File, name string, instances []GongstructIF, fields []GongFieldHeader, addIDs bool) {
	sheetName := __gong__shortenString(name)

	// Create a new sheet.
	f.NewSheet(sheetName)

	sortedSlice := make([]GongstructIF, len(instances))
	copy(sortedSlice, instances)
	slices.SortFunc(sortedSlice, func(a, b GongstructIF) int {
		return cmp.Compare(a.GetName(), b.GetName())
	})

	line := 1

	for index, fieldHeader := range fields {
		if !addIDs {
			f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(index+1)), line), fieldHeader.Name)
		} else {
			f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+1)), line), fieldHeader.Name)
			switch fieldHeader.GongFieldValueType {
			case GongFieldValueTypePointer:
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line),
					fieldHeader.Name+":"+fieldHeader.TargetGongstructName+":ID")
			case GongFieldValueTypeSliceOfPointers:
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line),
					fieldHeader.Name+":"+fieldHeader.TargetGongstructName+":IDs")
			default:
				// if index is 0, this is the ID of the instance
				if index == 0 {
					f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), "ID")
				} else {
					// one have to put the type of the cell
					header := fieldHeader.Name
					switch fieldHeader.GongFieldValueType {
					case GongFieldValueTypeInt:
						header += ":int"
					case GongFieldValueTypeIntDuration:
						header += ":duration"
					case GongFieldValueTypeFloat:
						header += ":float"
					case GongFieldValueTypeBool:
						header += ":bool"
					case GongFieldValueTypeString:
						header += ":string"
					case GongFieldValueTypeDate:
						header += ":date"
					default:
						header += ":basicType"
					}
					header += ":noID"
					f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), header)
				}
			}
		}
	}

	// AutoFilter starting from A1
	f.AutoFilter(sheetName,
		fmt.Sprintf("%s%d", GongIntToLetters(int32(1)), line),
		[]excelize.AutoFilterOptions{})

	for _, instance := range sortedSlice {
		line = line + 1

		// 3. Add the ID value in column B

		for index, fieldName := range fields {
			fieldStringValue := stage.GetFieldStringValueFromPointer(instance, fieldName.Name)
			if !addIDs {
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(index+1)), line), fieldStringValue.GetValueString())
			} else {
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+1)), line), fieldStringValue.GetValueString())
				if index == 0 {
					f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), instance.GongGetUUID(stage))
				} else {
					switch fieldStringValue.GongFieldValueType {
					case GongFieldValueTypePointer, GongFieldValueTypeSliceOfPointers:
						f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), fieldStringValue.ids)
					}
				}

			}
		}
	}
}

// SerializeExcelizePointer is the Stage method for Excel serialization.
func (stage *Stage) SerializeExcelizePointer[Type GongstructPtr](f *excelize.File) {
	stage.SerializeExcelizePointer2[Type](f, false)
}

// SerializeExcelizePointer2 is the Stage method for Excel serialization with optional IDs.
func (stage *Stage) SerializeExcelizePointer2[Type GongstructPtr](f *excelize.File, addIDs bool) {
	var ret Type
	set := *stage.GetInstancesSet[Type]()
	var instances []GongstructIF
	for key := range set {
		instances = append(instances, key)
	}
	stage.SerializeExcelize(f, ret.GongGetGongstructName(), instances, ret.GongGetFieldHeaders(), addIDs)
}
