// generated code - do not edit
package aixmdatatypes

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront(instance GongstructIF) {
	if instance != nil {
		instance.GongAfterCreateFromFront(stage)
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront(old, new GongstructIF) {
	if old != nil {
		old.GongOnAfterUpdateFromFront(stage, new)
	}
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront(staged, front GongstructIF) {
	if staged != nil {
		staged.GongAfterDeleteFromFront(stage, front)
	}
}

// insertion point
func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAircraftGroundServiceTypeCreateCallback != nil {
		stage.OnAfterCodeAircraftGroundServiceTypeCreateCallback.OnAfterCreate(stage, codeaircraftgroundservicetype)
	}
}

func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAircraftGroundServiceTypeUpdateCallback != nil {
		var frontCodeAircraftGroundServiceType *CodeAircraftGroundServiceType
		if front != nil {
			frontCodeAircraftGroundServiceType, _ = front.(*CodeAircraftGroundServiceType)
		}
		stage.OnAfterCodeAircraftGroundServiceTypeUpdateCallback.OnAfterUpdate(stage, codeaircraftgroundservicetype, frontCodeAircraftGroundServiceType)
	}
}

func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAircraftGroundServiceTypeDeleteCallback != nil {
		var frontCodeAircraftGroundServiceType *CodeAircraftGroundServiceType
		if front != nil {
			frontCodeAircraftGroundServiceType, _ = front.(*CodeAircraftGroundServiceType)
		}
		stage.OnAfterCodeAircraftGroundServiceTypeDeleteCallback.OnAfterDelete(stage, codeaircraftgroundservicetype, frontCodeAircraftGroundServiceType)
	}
}

func (codeaviationstandardstype *CodeAviationStandardsType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAviationStandardsTypeCreateCallback != nil {
		stage.OnAfterCodeAviationStandardsTypeCreateCallback.OnAfterCreate(stage, codeaviationstandardstype)
	}
}

func (codeaviationstandardstype *CodeAviationStandardsType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAviationStandardsTypeUpdateCallback != nil {
		var frontCodeAviationStandardsType *CodeAviationStandardsType
		if front != nil {
			frontCodeAviationStandardsType, _ = front.(*CodeAviationStandardsType)
		}
		stage.OnAfterCodeAviationStandardsTypeUpdateCallback.OnAfterUpdate(stage, codeaviationstandardstype, frontCodeAviationStandardsType)
	}
}

func (codeaviationstandardstype *CodeAviationStandardsType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAviationStandardsTypeDeleteCallback != nil {
		var frontCodeAviationStandardsType *CodeAviationStandardsType
		if front != nil {
			frontCodeAviationStandardsType, _ = front.(*CodeAviationStandardsType)
		}
		stage.OnAfterCodeAviationStandardsTypeDeleteCallback.OnAfterDelete(stage, codeaviationstandardstype, frontCodeAviationStandardsType)
	}
}

func (codecommunicationchanneltype *CodeCommunicationChannelType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeCommunicationChannelTypeCreateCallback != nil {
		stage.OnAfterCodeCommunicationChannelTypeCreateCallback.OnAfterCreate(stage, codecommunicationchanneltype)
	}
}

func (codecommunicationchanneltype *CodeCommunicationChannelType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeCommunicationChannelTypeUpdateCallback != nil {
		var frontCodeCommunicationChannelType *CodeCommunicationChannelType
		if front != nil {
			frontCodeCommunicationChannelType, _ = front.(*CodeCommunicationChannelType)
		}
		stage.OnAfterCodeCommunicationChannelTypeUpdateCallback.OnAfterUpdate(stage, codecommunicationchanneltype, frontCodeCommunicationChannelType)
	}
}

func (codecommunicationchanneltype *CodeCommunicationChannelType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeCommunicationChannelTypeDeleteCallback != nil {
		var frontCodeCommunicationChannelType *CodeCommunicationChannelType
		if front != nil {
			frontCodeCommunicationChannelType, _ = front.(*CodeCommunicationChannelType)
		}
		stage.OnAfterCodeCommunicationChannelTypeDeleteCallback.OnAfterDelete(stage, codecommunicationchanneltype, frontCodeCommunicationChannelType)
	}
}

func (codecommunicationdirectiontype *CodeCommunicationDirectionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeCommunicationDirectionTypeCreateCallback != nil {
		stage.OnAfterCodeCommunicationDirectionTypeCreateCallback.OnAfterCreate(stage, codecommunicationdirectiontype)
	}
}

func (codecommunicationdirectiontype *CodeCommunicationDirectionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeCommunicationDirectionTypeUpdateCallback != nil {
		var frontCodeCommunicationDirectionType *CodeCommunicationDirectionType
		if front != nil {
			frontCodeCommunicationDirectionType, _ = front.(*CodeCommunicationDirectionType)
		}
		stage.OnAfterCodeCommunicationDirectionTypeUpdateCallback.OnAfterUpdate(stage, codecommunicationdirectiontype, frontCodeCommunicationDirectionType)
	}
}

func (codecommunicationdirectiontype *CodeCommunicationDirectionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeCommunicationDirectionTypeDeleteCallback != nil {
		var frontCodeCommunicationDirectionType *CodeCommunicationDirectionType
		if front != nil {
			frontCodeCommunicationDirectionType, _ = front.(*CodeCommunicationDirectionType)
		}
		stage.OnAfterCodeCommunicationDirectionTypeDeleteCallback.OnAfterDelete(stage, codecommunicationdirectiontype, frontCodeCommunicationDirectionType)
	}
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeCommunicationModeTypeCreateCallback != nil {
		stage.OnAfterCodeCommunicationModeTypeCreateCallback.OnAfterCreate(stage, codecommunicationmodetype)
	}
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeCommunicationModeTypeUpdateCallback != nil {
		var frontCodeCommunicationModeType *CodeCommunicationModeType
		if front != nil {
			frontCodeCommunicationModeType, _ = front.(*CodeCommunicationModeType)
		}
		stage.OnAfterCodeCommunicationModeTypeUpdateCallback.OnAfterUpdate(stage, codecommunicationmodetype, frontCodeCommunicationModeType)
	}
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeCommunicationModeTypeDeleteCallback != nil {
		var frontCodeCommunicationModeType *CodeCommunicationModeType
		if front != nil {
			frontCodeCommunicationModeType, _ = front.(*CodeCommunicationModeType)
		}
		stage.OnAfterCodeCommunicationModeTypeDeleteCallback.OnAfterDelete(stage, codecommunicationmodetype, frontCodeCommunicationModeType)
	}
}

func (codedaytype *CodeDayType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeDayTypeCreateCallback != nil {
		stage.OnAfterCodeDayTypeCreateCallback.OnAfterCreate(stage, codedaytype)
	}
}

func (codedaytype *CodeDayType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDayTypeUpdateCallback != nil {
		var frontCodeDayType *CodeDayType
		if front != nil {
			frontCodeDayType, _ = front.(*CodeDayType)
		}
		stage.OnAfterCodeDayTypeUpdateCallback.OnAfterUpdate(stage, codedaytype, frontCodeDayType)
	}
}

func (codedaytype *CodeDayType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDayTypeDeleteCallback != nil {
		var frontCodeDayType *CodeDayType
		if front != nil {
			frontCodeDayType, _ = front.(*CodeDayType)
		}
		stage.OnAfterCodeDayTypeDeleteCallback.OnAfterDelete(stage, codedaytype, frontCodeDayType)
	}
}

func (codefacilityrankingtype *CodeFacilityRankingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeFacilityRankingTypeCreateCallback != nil {
		stage.OnAfterCodeFacilityRankingTypeCreateCallback.OnAfterCreate(stage, codefacilityrankingtype)
	}
}

func (codefacilityrankingtype *CodeFacilityRankingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFacilityRankingTypeUpdateCallback != nil {
		var frontCodeFacilityRankingType *CodeFacilityRankingType
		if front != nil {
			frontCodeFacilityRankingType, _ = front.(*CodeFacilityRankingType)
		}
		stage.OnAfterCodeFacilityRankingTypeUpdateCallback.OnAfterUpdate(stage, codefacilityrankingtype, frontCodeFacilityRankingType)
	}
}

func (codefacilityrankingtype *CodeFacilityRankingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFacilityRankingTypeDeleteCallback != nil {
		var frontCodeFacilityRankingType *CodeFacilityRankingType
		if front != nil {
			frontCodeFacilityRankingType, _ = front.(*CodeFacilityRankingType)
		}
		stage.OnAfterCodeFacilityRankingTypeDeleteCallback.OnAfterDelete(stage, codefacilityrankingtype, frontCodeFacilityRankingType)
	}
}

func (codefirefightingtype *CodeFireFightingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeFireFightingTypeCreateCallback != nil {
		stage.OnAfterCodeFireFightingTypeCreateCallback.OnAfterCreate(stage, codefirefightingtype)
	}
}

func (codefirefightingtype *CodeFireFightingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFireFightingTypeUpdateCallback != nil {
		var frontCodeFireFightingType *CodeFireFightingType
		if front != nil {
			frontCodeFireFightingType, _ = front.(*CodeFireFightingType)
		}
		stage.OnAfterCodeFireFightingTypeUpdateCallback.OnAfterUpdate(stage, codefirefightingtype, frontCodeFireFightingType)
	}
}

func (codefirefightingtype *CodeFireFightingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFireFightingTypeDeleteCallback != nil {
		var frontCodeFireFightingType *CodeFireFightingType
		if front != nil {
			frontCodeFireFightingType, _ = front.(*CodeFireFightingType)
		}
		stage.OnAfterCodeFireFightingTypeDeleteCallback.OnAfterDelete(stage, codefirefightingtype, frontCodeFireFightingType)
	}
}

func (codeflightdestinationtype *CodeFlightDestinationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeFlightDestinationTypeCreateCallback != nil {
		stage.OnAfterCodeFlightDestinationTypeCreateCallback.OnAfterCreate(stage, codeflightdestinationtype)
	}
}

func (codeflightdestinationtype *CodeFlightDestinationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFlightDestinationTypeUpdateCallback != nil {
		var frontCodeFlightDestinationType *CodeFlightDestinationType
		if front != nil {
			frontCodeFlightDestinationType, _ = front.(*CodeFlightDestinationType)
		}
		stage.OnAfterCodeFlightDestinationTypeUpdateCallback.OnAfterUpdate(stage, codeflightdestinationtype, frontCodeFlightDestinationType)
	}
}

func (codeflightdestinationtype *CodeFlightDestinationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFlightDestinationTypeDeleteCallback != nil {
		var frontCodeFlightDestinationType *CodeFlightDestinationType
		if front != nil {
			frontCodeFlightDestinationType, _ = front.(*CodeFlightDestinationType)
		}
		stage.OnAfterCodeFlightDestinationTypeDeleteCallback.OnAfterDelete(stage, codeflightdestinationtype, frontCodeFlightDestinationType)
	}
}

func (codeflightruletype *CodeFlightRuleType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeFlightRuleTypeCreateCallback != nil {
		stage.OnAfterCodeFlightRuleTypeCreateCallback.OnAfterCreate(stage, codeflightruletype)
	}
}

func (codeflightruletype *CodeFlightRuleType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFlightRuleTypeUpdateCallback != nil {
		var frontCodeFlightRuleType *CodeFlightRuleType
		if front != nil {
			frontCodeFlightRuleType, _ = front.(*CodeFlightRuleType)
		}
		stage.OnAfterCodeFlightRuleTypeUpdateCallback.OnAfterUpdate(stage, codeflightruletype, frontCodeFlightRuleType)
	}
}

func (codeflightruletype *CodeFlightRuleType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFlightRuleTypeDeleteCallback != nil {
		var frontCodeFlightRuleType *CodeFlightRuleType
		if front != nil {
			frontCodeFlightRuleType, _ = front.(*CodeFlightRuleType)
		}
		stage.OnAfterCodeFlightRuleTypeDeleteCallback.OnAfterDelete(stage, codeflightruletype, frontCodeFlightRuleType)
	}
}

func (codefueltype *CodeFuelType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeFuelTypeCreateCallback != nil {
		stage.OnAfterCodeFuelTypeCreateCallback.OnAfterCreate(stage, codefueltype)
	}
}

func (codefueltype *CodeFuelType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFuelTypeUpdateCallback != nil {
		var frontCodeFuelType *CodeFuelType
		if front != nil {
			frontCodeFuelType, _ = front.(*CodeFuelType)
		}
		stage.OnAfterCodeFuelTypeUpdateCallback.OnAfterUpdate(stage, codefueltype, frontCodeFuelType)
	}
}

func (codefueltype *CodeFuelType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFuelTypeDeleteCallback != nil {
		var frontCodeFuelType *CodeFuelType
		if front != nil {
			frontCodeFuelType, _ = front.(*CodeFuelType)
		}
		stage.OnAfterCodeFuelTypeDeleteCallback.OnAfterDelete(stage, codefueltype, frontCodeFuelType)
	}
}

func (codeintensitystandbytype *CodeIntensityStandByType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeIntensityStandByTypeCreateCallback != nil {
		stage.OnAfterCodeIntensityStandByTypeCreateCallback.OnAfterCreate(stage, codeintensitystandbytype)
	}
}

func (codeintensitystandbytype *CodeIntensityStandByType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeIntensityStandByTypeUpdateCallback != nil {
		var frontCodeIntensityStandByType *CodeIntensityStandByType
		if front != nil {
			frontCodeIntensityStandByType, _ = front.(*CodeIntensityStandByType)
		}
		stage.OnAfterCodeIntensityStandByTypeUpdateCallback.OnAfterUpdate(stage, codeintensitystandbytype, frontCodeIntensityStandByType)
	}
}

func (codeintensitystandbytype *CodeIntensityStandByType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeIntensityStandByTypeDeleteCallback != nil {
		var frontCodeIntensityStandByType *CodeIntensityStandByType
		if front != nil {
			frontCodeIntensityStandByType, _ = front.(*CodeIntensityStandByType)
		}
		stage.OnAfterCodeIntensityStandByTypeDeleteCallback.OnAfterDelete(stage, codeintensitystandbytype, frontCodeIntensityStandByType)
	}
}

func (codelanguagetype *CodeLanguageType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeLanguageTypeCreateCallback != nil {
		stage.OnAfterCodeLanguageTypeCreateCallback.OnAfterCreate(stage, codelanguagetype)
	}
}

func (codelanguagetype *CodeLanguageType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLanguageTypeUpdateCallback != nil {
		var frontCodeLanguageType *CodeLanguageType
		if front != nil {
			frontCodeLanguageType, _ = front.(*CodeLanguageType)
		}
		stage.OnAfterCodeLanguageTypeUpdateCallback.OnAfterUpdate(stage, codelanguagetype, frontCodeLanguageType)
	}
}

func (codelanguagetype *CodeLanguageType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLanguageTypeDeleteCallback != nil {
		var frontCodeLanguageType *CodeLanguageType
		if front != nil {
			frontCodeLanguageType, _ = front.(*CodeLanguageType)
		}
		stage.OnAfterCodeLanguageTypeDeleteCallback.OnAfterDelete(stage, codelanguagetype, frontCodeLanguageType)
	}
}

func (codelevelseriestype *CodeLevelSeriesType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeLevelSeriesTypeCreateCallback != nil {
		stage.OnAfterCodeLevelSeriesTypeCreateCallback.OnAfterCreate(stage, codelevelseriestype)
	}
}

func (codelevelseriestype *CodeLevelSeriesType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLevelSeriesTypeUpdateCallback != nil {
		var frontCodeLevelSeriesType *CodeLevelSeriesType
		if front != nil {
			frontCodeLevelSeriesType, _ = front.(*CodeLevelSeriesType)
		}
		stage.OnAfterCodeLevelSeriesTypeUpdateCallback.OnAfterUpdate(stage, codelevelseriestype, frontCodeLevelSeriesType)
	}
}

func (codelevelseriestype *CodeLevelSeriesType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLevelSeriesTypeDeleteCallback != nil {
		var frontCodeLevelSeriesType *CodeLevelSeriesType
		if front != nil {
			frontCodeLevelSeriesType, _ = front.(*CodeLevelSeriesType)
		}
		stage.OnAfterCodeLevelSeriesTypeDeleteCallback.OnAfterDelete(stage, codelevelseriestype, frontCodeLevelSeriesType)
	}
}

func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeLevelTableDesignatorTypeCreateCallback != nil {
		stage.OnAfterCodeLevelTableDesignatorTypeCreateCallback.OnAfterCreate(stage, codeleveltabledesignatortype)
	}
}

func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLevelTableDesignatorTypeUpdateCallback != nil {
		var frontCodeLevelTableDesignatorType *CodeLevelTableDesignatorType
		if front != nil {
			frontCodeLevelTableDesignatorType, _ = front.(*CodeLevelTableDesignatorType)
		}
		stage.OnAfterCodeLevelTableDesignatorTypeUpdateCallback.OnAfterUpdate(stage, codeleveltabledesignatortype, frontCodeLevelTableDesignatorType)
	}
}

func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLevelTableDesignatorTypeDeleteCallback != nil {
		var frontCodeLevelTableDesignatorType *CodeLevelTableDesignatorType
		if front != nil {
			frontCodeLevelTableDesignatorType, _ = front.(*CodeLevelTableDesignatorType)
		}
		stage.OnAfterCodeLevelTableDesignatorTypeDeleteCallback.OnAfterDelete(stage, codeleveltabledesignatortype, frontCodeLevelTableDesignatorType)
	}
}

func (codelightintensitytype *CodeLightIntensityType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeLightIntensityTypeCreateCallback != nil {
		stage.OnAfterCodeLightIntensityTypeCreateCallback.OnAfterCreate(stage, codelightintensitytype)
	}
}

func (codelightintensitytype *CodeLightIntensityType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLightIntensityTypeUpdateCallback != nil {
		var frontCodeLightIntensityType *CodeLightIntensityType
		if front != nil {
			frontCodeLightIntensityType, _ = front.(*CodeLightIntensityType)
		}
		stage.OnAfterCodeLightIntensityTypeUpdateCallback.OnAfterUpdate(stage, codelightintensitytype, frontCodeLightIntensityType)
	}
}

func (codelightintensitytype *CodeLightIntensityType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLightIntensityTypeDeleteCallback != nil {
		var frontCodeLightIntensityType *CodeLightIntensityType
		if front != nil {
			frontCodeLightIntensityType, _ = front.(*CodeLightIntensityType)
		}
		stage.OnAfterCodeLightIntensityTypeDeleteCallback.OnAfterDelete(stage, codelightintensitytype, frontCodeLightIntensityType)
	}
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeMilitaryOperationsTypeCreateCallback != nil {
		stage.OnAfterCodeMilitaryOperationsTypeCreateCallback.OnAfterCreate(stage, codemilitaryoperationstype)
	}
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMilitaryOperationsTypeUpdateCallback != nil {
		var frontCodeMilitaryOperationsType *CodeMilitaryOperationsType
		if front != nil {
			frontCodeMilitaryOperationsType, _ = front.(*CodeMilitaryOperationsType)
		}
		stage.OnAfterCodeMilitaryOperationsTypeUpdateCallback.OnAfterUpdate(stage, codemilitaryoperationstype, frontCodeMilitaryOperationsType)
	}
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMilitaryOperationsTypeDeleteCallback != nil {
		var frontCodeMilitaryOperationsType *CodeMilitaryOperationsType
		if front != nil {
			frontCodeMilitaryOperationsType, _ = front.(*CodeMilitaryOperationsType)
		}
		stage.OnAfterCodeMilitaryOperationsTypeDeleteCallback.OnAfterDelete(stage, codemilitaryoperationstype, frontCodeMilitaryOperationsType)
	}
}

func (codenitrogentype *CodeNitrogenType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeNitrogenTypeCreateCallback != nil {
		stage.OnAfterCodeNitrogenTypeCreateCallback.OnAfterCreate(stage, codenitrogentype)
	}
}

func (codenitrogentype *CodeNitrogenType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNitrogenTypeUpdateCallback != nil {
		var frontCodeNitrogenType *CodeNitrogenType
		if front != nil {
			frontCodeNitrogenType, _ = front.(*CodeNitrogenType)
		}
		stage.OnAfterCodeNitrogenTypeUpdateCallback.OnAfterUpdate(stage, codenitrogentype, frontCodeNitrogenType)
	}
}

func (codenitrogentype *CodeNitrogenType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNitrogenTypeDeleteCallback != nil {
		var frontCodeNitrogenType *CodeNitrogenType
		if front != nil {
			frontCodeNitrogenType, _ = front.(*CodeNitrogenType)
		}
		stage.OnAfterCodeNitrogenTypeDeleteCallback.OnAfterDelete(stage, codenitrogentype, frontCodeNitrogenType)
	}
}

func (codenorthreferencetype *CodeNorthReferenceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeNorthReferenceTypeCreateCallback != nil {
		stage.OnAfterCodeNorthReferenceTypeCreateCallback.OnAfterCreate(stage, codenorthreferencetype)
	}
}

func (codenorthreferencetype *CodeNorthReferenceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNorthReferenceTypeUpdateCallback != nil {
		var frontCodeNorthReferenceType *CodeNorthReferenceType
		if front != nil {
			frontCodeNorthReferenceType, _ = front.(*CodeNorthReferenceType)
		}
		stage.OnAfterCodeNorthReferenceTypeUpdateCallback.OnAfterUpdate(stage, codenorthreferencetype, frontCodeNorthReferenceType)
	}
}

func (codenorthreferencetype *CodeNorthReferenceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNorthReferenceTypeDeleteCallback != nil {
		var frontCodeNorthReferenceType *CodeNorthReferenceType
		if front != nil {
			frontCodeNorthReferenceType, _ = front.(*CodeNorthReferenceType)
		}
		stage.OnAfterCodeNorthReferenceTypeDeleteCallback.OnAfterDelete(stage, codenorthreferencetype, frontCodeNorthReferenceType)
	}
}

func (codenotepurposetype *CodeNotePurposeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeNotePurposeTypeCreateCallback != nil {
		stage.OnAfterCodeNotePurposeTypeCreateCallback.OnAfterCreate(stage, codenotepurposetype)
	}
}

func (codenotepurposetype *CodeNotePurposeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNotePurposeTypeUpdateCallback != nil {
		var frontCodeNotePurposeType *CodeNotePurposeType
		if front != nil {
			frontCodeNotePurposeType, _ = front.(*CodeNotePurposeType)
		}
		stage.OnAfterCodeNotePurposeTypeUpdateCallback.OnAfterUpdate(stage, codenotepurposetype, frontCodeNotePurposeType)
	}
}

func (codenotepurposetype *CodeNotePurposeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNotePurposeTypeDeleteCallback != nil {
		var frontCodeNotePurposeType *CodeNotePurposeType
		if front != nil {
			frontCodeNotePurposeType, _ = front.(*CodeNotePurposeType)
		}
		stage.OnAfterCodeNotePurposeTypeDeleteCallback.OnAfterDelete(stage, codenotepurposetype, frontCodeNotePurposeType)
	}
}

func (codeoiltype *CodeOilType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeOilTypeCreateCallback != nil {
		stage.OnAfterCodeOilTypeCreateCallback.OnAfterCreate(stage, codeoiltype)
	}
}

func (codeoiltype *CodeOilType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeOilTypeUpdateCallback != nil {
		var frontCodeOilType *CodeOilType
		if front != nil {
			frontCodeOilType, _ = front.(*CodeOilType)
		}
		stage.OnAfterCodeOilTypeUpdateCallback.OnAfterUpdate(stage, codeoiltype, frontCodeOilType)
	}
}

func (codeoiltype *CodeOilType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeOilTypeDeleteCallback != nil {
		var frontCodeOilType *CodeOilType
		if front != nil {
			frontCodeOilType, _ = front.(*CodeOilType)
		}
		stage.OnAfterCodeOilTypeDeleteCallback.OnAfterDelete(stage, codeoiltype, frontCodeOilType)
	}
}

func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeOrganisationDesignatorTypeCreateCallback != nil {
		stage.OnAfterCodeOrganisationDesignatorTypeCreateCallback.OnAfterCreate(stage, codeorganisationdesignatortype)
	}
}

func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeOrganisationDesignatorTypeUpdateCallback != nil {
		var frontCodeOrganisationDesignatorType *CodeOrganisationDesignatorType
		if front != nil {
			frontCodeOrganisationDesignatorType, _ = front.(*CodeOrganisationDesignatorType)
		}
		stage.OnAfterCodeOrganisationDesignatorTypeUpdateCallback.OnAfterUpdate(stage, codeorganisationdesignatortype, frontCodeOrganisationDesignatorType)
	}
}

func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeOrganisationDesignatorTypeDeleteCallback != nil {
		var frontCodeOrganisationDesignatorType *CodeOrganisationDesignatorType
		if front != nil {
			frontCodeOrganisationDesignatorType, _ = front.(*CodeOrganisationDesignatorType)
		}
		stage.OnAfterCodeOrganisationDesignatorTypeDeleteCallback.OnAfterDelete(stage, codeorganisationdesignatortype, frontCodeOrganisationDesignatorType)
	}
}

func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeOrganisationHierarchyTypeCreateCallback != nil {
		stage.OnAfterCodeOrganisationHierarchyTypeCreateCallback.OnAfterCreate(stage, codeorganisationhierarchytype)
	}
}

func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeOrganisationHierarchyTypeUpdateCallback != nil {
		var frontCodeOrganisationHierarchyType *CodeOrganisationHierarchyType
		if front != nil {
			frontCodeOrganisationHierarchyType, _ = front.(*CodeOrganisationHierarchyType)
		}
		stage.OnAfterCodeOrganisationHierarchyTypeUpdateCallback.OnAfterUpdate(stage, codeorganisationhierarchytype, frontCodeOrganisationHierarchyType)
	}
}

func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeOrganisationHierarchyTypeDeleteCallback != nil {
		var frontCodeOrganisationHierarchyType *CodeOrganisationHierarchyType
		if front != nil {
			frontCodeOrganisationHierarchyType, _ = front.(*CodeOrganisationHierarchyType)
		}
		stage.OnAfterCodeOrganisationHierarchyTypeDeleteCallback.OnAfterDelete(stage, codeorganisationhierarchytype, frontCodeOrganisationHierarchyType)
	}
}

func (codeorganisationtype *CodeOrganisationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeOrganisationTypeCreateCallback != nil {
		stage.OnAfterCodeOrganisationTypeCreateCallback.OnAfterCreate(stage, codeorganisationtype)
	}
}

func (codeorganisationtype *CodeOrganisationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeOrganisationTypeUpdateCallback != nil {
		var frontCodeOrganisationType *CodeOrganisationType
		if front != nil {
			frontCodeOrganisationType, _ = front.(*CodeOrganisationType)
		}
		stage.OnAfterCodeOrganisationTypeUpdateCallback.OnAfterUpdate(stage, codeorganisationtype, frontCodeOrganisationType)
	}
}

func (codeorganisationtype *CodeOrganisationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeOrganisationTypeDeleteCallback != nil {
		var frontCodeOrganisationType *CodeOrganisationType
		if front != nil {
			frontCodeOrganisationType, _ = front.(*CodeOrganisationType)
		}
		stage.OnAfterCodeOrganisationTypeDeleteCallback.OnAfterDelete(stage, codeorganisationtype, frontCodeOrganisationType)
	}
}

func (codeoxygentype *CodeOxygenType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeOxygenTypeCreateCallback != nil {
		stage.OnAfterCodeOxygenTypeCreateCallback.OnAfterCreate(stage, codeoxygentype)
	}
}

func (codeoxygentype *CodeOxygenType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeOxygenTypeUpdateCallback != nil {
		var frontCodeOxygenType *CodeOxygenType
		if front != nil {
			frontCodeOxygenType, _ = front.(*CodeOxygenType)
		}
		stage.OnAfterCodeOxygenTypeUpdateCallback.OnAfterUpdate(stage, codeoxygentype, frontCodeOxygenType)
	}
}

func (codeoxygentype *CodeOxygenType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeOxygenTypeDeleteCallback != nil {
		var frontCodeOxygenType *CodeOxygenType
		if front != nil {
			frontCodeOxygenType, _ = front.(*CodeOxygenType)
		}
		stage.OnAfterCodeOxygenTypeDeleteCallback.OnAfterDelete(stage, codeoxygentype, frontCodeOxygenType)
	}
}

func (codepassengerservicetype *CodePassengerServiceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodePassengerServiceTypeCreateCallback != nil {
		stage.OnAfterCodePassengerServiceTypeCreateCallback.OnAfterCreate(stage, codepassengerservicetype)
	}
}

func (codepassengerservicetype *CodePassengerServiceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodePassengerServiceTypeUpdateCallback != nil {
		var frontCodePassengerServiceType *CodePassengerServiceType
		if front != nil {
			frontCodePassengerServiceType, _ = front.(*CodePassengerServiceType)
		}
		stage.OnAfterCodePassengerServiceTypeUpdateCallback.OnAfterUpdate(stage, codepassengerservicetype, frontCodePassengerServiceType)
	}
}

func (codepassengerservicetype *CodePassengerServiceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodePassengerServiceTypeDeleteCallback != nil {
		var frontCodePassengerServiceType *CodePassengerServiceType
		if front != nil {
			frontCodePassengerServiceType, _ = front.(*CodePassengerServiceType)
		}
		stage.OnAfterCodePassengerServiceTypeDeleteCallback.OnAfterDelete(stage, codepassengerservicetype, frontCodePassengerServiceType)
	}
}

func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodePilotControlledLightingTypeCreateCallback != nil {
		stage.OnAfterCodePilotControlledLightingTypeCreateCallback.OnAfterCreate(stage, codepilotcontrolledlightingtype)
	}
}

func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodePilotControlledLightingTypeUpdateCallback != nil {
		var frontCodePilotControlledLightingType *CodePilotControlledLightingType
		if front != nil {
			frontCodePilotControlledLightingType, _ = front.(*CodePilotControlledLightingType)
		}
		stage.OnAfterCodePilotControlledLightingTypeUpdateCallback.OnAfterUpdate(stage, codepilotcontrolledlightingtype, frontCodePilotControlledLightingType)
	}
}

func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodePilotControlledLightingTypeDeleteCallback != nil {
		var frontCodePilotControlledLightingType *CodePilotControlledLightingType
		if front != nil {
			frontCodePilotControlledLightingType, _ = front.(*CodePilotControlledLightingType)
		}
		stage.OnAfterCodePilotControlledLightingTypeDeleteCallback.OnAfterDelete(stage, codepilotcontrolledlightingtype, frontCodePilotControlledLightingType)
	}
}

func (codervsmtype *CodeRVSMType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRVSMTypeCreateCallback != nil {
		stage.OnAfterCodeRVSMTypeCreateCallback.OnAfterCreate(stage, codervsmtype)
	}
}

func (codervsmtype *CodeRVSMType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRVSMTypeUpdateCallback != nil {
		var frontCodeRVSMType *CodeRVSMType
		if front != nil {
			frontCodeRVSMType, _ = front.(*CodeRVSMType)
		}
		stage.OnAfterCodeRVSMTypeUpdateCallback.OnAfterUpdate(stage, codervsmtype, frontCodeRVSMType)
	}
}

func (codervsmtype *CodeRVSMType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRVSMTypeDeleteCallback != nil {
		var frontCodeRVSMType *CodeRVSMType
		if front != nil {
			frontCodeRVSMType, _ = front.(*CodeRVSMType)
		}
		stage.OnAfterCodeRVSMTypeDeleteCallback.OnAfterDelete(stage, codervsmtype, frontCodeRVSMType)
	}
}

func (coderadioemissiontype *CodeRadioEmissionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRadioEmissionTypeCreateCallback != nil {
		stage.OnAfterCodeRadioEmissionTypeCreateCallback.OnAfterCreate(stage, coderadioemissiontype)
	}
}

func (coderadioemissiontype *CodeRadioEmissionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRadioEmissionTypeUpdateCallback != nil {
		var frontCodeRadioEmissionType *CodeRadioEmissionType
		if front != nil {
			frontCodeRadioEmissionType, _ = front.(*CodeRadioEmissionType)
		}
		stage.OnAfterCodeRadioEmissionTypeUpdateCallback.OnAfterUpdate(stage, coderadioemissiontype, frontCodeRadioEmissionType)
	}
}

func (coderadioemissiontype *CodeRadioEmissionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRadioEmissionTypeDeleteCallback != nil {
		var frontCodeRadioEmissionType *CodeRadioEmissionType
		if front != nil {
			frontCodeRadioEmissionType, _ = front.(*CodeRadioEmissionType)
		}
		stage.OnAfterCodeRadioEmissionTypeDeleteCallback.OnAfterDelete(stage, coderadioemissiontype, frontCodeRadioEmissionType)
	}
}

func (codeserviceatctype *CodeServiceATCType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeServiceATCTypeCreateCallback != nil {
		stage.OnAfterCodeServiceATCTypeCreateCallback.OnAfterCreate(stage, codeserviceatctype)
	}
}

func (codeserviceatctype *CodeServiceATCType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeServiceATCTypeUpdateCallback != nil {
		var frontCodeServiceATCType *CodeServiceATCType
		if front != nil {
			frontCodeServiceATCType, _ = front.(*CodeServiceATCType)
		}
		stage.OnAfterCodeServiceATCTypeUpdateCallback.OnAfterUpdate(stage, codeserviceatctype, frontCodeServiceATCType)
	}
}

func (codeserviceatctype *CodeServiceATCType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeServiceATCTypeDeleteCallback != nil {
		var frontCodeServiceATCType *CodeServiceATCType
		if front != nil {
			frontCodeServiceATCType, _ = front.(*CodeServiceATCType)
		}
		stage.OnAfterCodeServiceATCTypeDeleteCallback.OnAfterDelete(stage, codeserviceatctype, frontCodeServiceATCType)
	}
}

func (codeserviceatfmtype *CodeServiceATFMType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeServiceATFMTypeCreateCallback != nil {
		stage.OnAfterCodeServiceATFMTypeCreateCallback.OnAfterCreate(stage, codeserviceatfmtype)
	}
}

func (codeserviceatfmtype *CodeServiceATFMType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeServiceATFMTypeUpdateCallback != nil {
		var frontCodeServiceATFMType *CodeServiceATFMType
		if front != nil {
			frontCodeServiceATFMType, _ = front.(*CodeServiceATFMType)
		}
		stage.OnAfterCodeServiceATFMTypeUpdateCallback.OnAfterUpdate(stage, codeserviceatfmtype, frontCodeServiceATFMType)
	}
}

func (codeserviceatfmtype *CodeServiceATFMType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeServiceATFMTypeDeleteCallback != nil {
		var frontCodeServiceATFMType *CodeServiceATFMType
		if front != nil {
			frontCodeServiceATFMType, _ = front.(*CodeServiceATFMType)
		}
		stage.OnAfterCodeServiceATFMTypeDeleteCallback.OnAfterDelete(stage, codeserviceatfmtype, frontCodeServiceATFMType)
	}
}

func (codeservicegroundcontroltype *CodeServiceGroundControlType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeServiceGroundControlTypeCreateCallback != nil {
		stage.OnAfterCodeServiceGroundControlTypeCreateCallback.OnAfterCreate(stage, codeservicegroundcontroltype)
	}
}

func (codeservicegroundcontroltype *CodeServiceGroundControlType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeServiceGroundControlTypeUpdateCallback != nil {
		var frontCodeServiceGroundControlType *CodeServiceGroundControlType
		if front != nil {
			frontCodeServiceGroundControlType, _ = front.(*CodeServiceGroundControlType)
		}
		stage.OnAfterCodeServiceGroundControlTypeUpdateCallback.OnAfterUpdate(stage, codeservicegroundcontroltype, frontCodeServiceGroundControlType)
	}
}

func (codeservicegroundcontroltype *CodeServiceGroundControlType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeServiceGroundControlTypeDeleteCallback != nil {
		var frontCodeServiceGroundControlType *CodeServiceGroundControlType
		if front != nil {
			frontCodeServiceGroundControlType, _ = front.(*CodeServiceGroundControlType)
		}
		stage.OnAfterCodeServiceGroundControlTypeDeleteCallback.OnAfterDelete(stage, codeservicegroundcontroltype, frontCodeServiceGroundControlType)
	}
}

func (codeserviceinformationtype *CodeServiceInformationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeServiceInformationTypeCreateCallback != nil {
		stage.OnAfterCodeServiceInformationTypeCreateCallback.OnAfterCreate(stage, codeserviceinformationtype)
	}
}

func (codeserviceinformationtype *CodeServiceInformationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeServiceInformationTypeUpdateCallback != nil {
		var frontCodeServiceInformationType *CodeServiceInformationType
		if front != nil {
			frontCodeServiceInformationType, _ = front.(*CodeServiceInformationType)
		}
		stage.OnAfterCodeServiceInformationTypeUpdateCallback.OnAfterUpdate(stage, codeserviceinformationtype, frontCodeServiceInformationType)
	}
}

func (codeserviceinformationtype *CodeServiceInformationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeServiceInformationTypeDeleteCallback != nil {
		var frontCodeServiceInformationType *CodeServiceInformationType
		if front != nil {
			frontCodeServiceInformationType, _ = front.(*CodeServiceInformationType)
		}
		stage.OnAfterCodeServiceInformationTypeDeleteCallback.OnAfterDelete(stage, codeserviceinformationtype, frontCodeServiceInformationType)
	}
}

func (codeservicesartype *CodeServiceSARType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeServiceSARTypeCreateCallback != nil {
		stage.OnAfterCodeServiceSARTypeCreateCallback.OnAfterCreate(stage, codeservicesartype)
	}
}

func (codeservicesartype *CodeServiceSARType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeServiceSARTypeUpdateCallback != nil {
		var frontCodeServiceSARType *CodeServiceSARType
		if front != nil {
			frontCodeServiceSARType, _ = front.(*CodeServiceSARType)
		}
		stage.OnAfterCodeServiceSARTypeUpdateCallback.OnAfterUpdate(stage, codeservicesartype, frontCodeServiceSARType)
	}
}

func (codeservicesartype *CodeServiceSARType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeServiceSARTypeDeleteCallback != nil {
		var frontCodeServiceSARType *CodeServiceSARType
		if front != nil {
			frontCodeServiceSARType, _ = front.(*CodeServiceSARType)
		}
		stage.OnAfterCodeServiceSARTypeDeleteCallback.OnAfterDelete(stage, codeservicesartype, frontCodeServiceSARType)
	}
}

func (codespecialdatetype *CodeSpecialDateType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeSpecialDateTypeCreateCallback != nil {
		stage.OnAfterCodeSpecialDateTypeCreateCallback.OnAfterCreate(stage, codespecialdatetype)
	}
}

func (codespecialdatetype *CodeSpecialDateType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSpecialDateTypeUpdateCallback != nil {
		var frontCodeSpecialDateType *CodeSpecialDateType
		if front != nil {
			frontCodeSpecialDateType, _ = front.(*CodeSpecialDateType)
		}
		stage.OnAfterCodeSpecialDateTypeUpdateCallback.OnAfterUpdate(stage, codespecialdatetype, frontCodeSpecialDateType)
	}
}

func (codespecialdatetype *CodeSpecialDateType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSpecialDateTypeDeleteCallback != nil {
		var frontCodeSpecialDateType *CodeSpecialDateType
		if front != nil {
			frontCodeSpecialDateType, _ = front.(*CodeSpecialDateType)
		}
		stage.OnAfterCodeSpecialDateTypeDeleteCallback.OnAfterDelete(stage, codespecialdatetype, frontCodeSpecialDateType)
	}
}

func (codestatusoperationstype *CodeStatusOperationsType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeStatusOperationsTypeCreateCallback != nil {
		stage.OnAfterCodeStatusOperationsTypeCreateCallback.OnAfterCreate(stage, codestatusoperationstype)
	}
}

func (codestatusoperationstype *CodeStatusOperationsType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeStatusOperationsTypeUpdateCallback != nil {
		var frontCodeStatusOperationsType *CodeStatusOperationsType
		if front != nil {
			frontCodeStatusOperationsType, _ = front.(*CodeStatusOperationsType)
		}
		stage.OnAfterCodeStatusOperationsTypeUpdateCallback.OnAfterUpdate(stage, codestatusoperationstype, frontCodeStatusOperationsType)
	}
}

func (codestatusoperationstype *CodeStatusOperationsType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeStatusOperationsTypeDeleteCallback != nil {
		var frontCodeStatusOperationsType *CodeStatusOperationsType
		if front != nil {
			frontCodeStatusOperationsType, _ = front.(*CodeStatusOperationsType)
		}
		stage.OnAfterCodeStatusOperationsTypeDeleteCallback.OnAfterDelete(stage, codestatusoperationstype, frontCodeStatusOperationsType)
	}
}

func (codestatusservicetype *CodeStatusServiceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeStatusServiceTypeCreateCallback != nil {
		stage.OnAfterCodeStatusServiceTypeCreateCallback.OnAfterCreate(stage, codestatusservicetype)
	}
}

func (codestatusservicetype *CodeStatusServiceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeStatusServiceTypeUpdateCallback != nil {
		var frontCodeStatusServiceType *CodeStatusServiceType
		if front != nil {
			frontCodeStatusServiceType, _ = front.(*CodeStatusServiceType)
		}
		stage.OnAfterCodeStatusServiceTypeUpdateCallback.OnAfterUpdate(stage, codestatusservicetype, frontCodeStatusServiceType)
	}
}

func (codestatusservicetype *CodeStatusServiceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeStatusServiceTypeDeleteCallback != nil {
		var frontCodeStatusServiceType *CodeStatusServiceType
		if front != nil {
			frontCodeStatusServiceType, _ = front.(*CodeStatusServiceType)
		}
		stage.OnAfterCodeStatusServiceTypeDeleteCallback.OnAfterDelete(stage, codestatusservicetype, frontCodeStatusServiceType)
	}
}

func (codesystemactivationtype *CodeSystemActivationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeSystemActivationTypeCreateCallback != nil {
		stage.OnAfterCodeSystemActivationTypeCreateCallback.OnAfterCreate(stage, codesystemactivationtype)
	}
}

func (codesystemactivationtype *CodeSystemActivationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSystemActivationTypeUpdateCallback != nil {
		var frontCodeSystemActivationType *CodeSystemActivationType
		if front != nil {
			frontCodeSystemActivationType, _ = front.(*CodeSystemActivationType)
		}
		stage.OnAfterCodeSystemActivationTypeUpdateCallback.OnAfterUpdate(stage, codesystemactivationtype, frontCodeSystemActivationType)
	}
}

func (codesystemactivationtype *CodeSystemActivationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSystemActivationTypeDeleteCallback != nil {
		var frontCodeSystemActivationType *CodeSystemActivationType
		if front != nil {
			frontCodeSystemActivationType, _ = front.(*CodeSystemActivationType)
		}
		stage.OnAfterCodeSystemActivationTypeDeleteCallback.OnAfterDelete(stage, codesystemactivationtype, frontCodeSystemActivationType)
	}
}

func (codetelecomnetworktype *CodeTelecomNetworkType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeTelecomNetworkTypeCreateCallback != nil {
		stage.OnAfterCodeTelecomNetworkTypeCreateCallback.OnAfterCreate(stage, codetelecomnetworktype)
	}
}

func (codetelecomnetworktype *CodeTelecomNetworkType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTelecomNetworkTypeUpdateCallback != nil {
		var frontCodeTelecomNetworkType *CodeTelecomNetworkType
		if front != nil {
			frontCodeTelecomNetworkType, _ = front.(*CodeTelecomNetworkType)
		}
		stage.OnAfterCodeTelecomNetworkTypeUpdateCallback.OnAfterUpdate(stage, codetelecomnetworktype, frontCodeTelecomNetworkType)
	}
}

func (codetelecomnetworktype *CodeTelecomNetworkType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTelecomNetworkTypeDeleteCallback != nil {
		var frontCodeTelecomNetworkType *CodeTelecomNetworkType
		if front != nil {
			frontCodeTelecomNetworkType, _ = front.(*CodeTelecomNetworkType)
		}
		stage.OnAfterCodeTelecomNetworkTypeDeleteCallback.OnAfterDelete(stage, codetelecomnetworktype, frontCodeTelecomNetworkType)
	}
}

func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeTimeEventCombinationTypeCreateCallback != nil {
		stage.OnAfterCodeTimeEventCombinationTypeCreateCallback.OnAfterCreate(stage, codetimeeventcombinationtype)
	}
}

func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTimeEventCombinationTypeUpdateCallback != nil {
		var frontCodeTimeEventCombinationType *CodeTimeEventCombinationType
		if front != nil {
			frontCodeTimeEventCombinationType, _ = front.(*CodeTimeEventCombinationType)
		}
		stage.OnAfterCodeTimeEventCombinationTypeUpdateCallback.OnAfterUpdate(stage, codetimeeventcombinationtype, frontCodeTimeEventCombinationType)
	}
}

func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTimeEventCombinationTypeDeleteCallback != nil {
		var frontCodeTimeEventCombinationType *CodeTimeEventCombinationType
		if front != nil {
			frontCodeTimeEventCombinationType, _ = front.(*CodeTimeEventCombinationType)
		}
		stage.OnAfterCodeTimeEventCombinationTypeDeleteCallback.OnAfterDelete(stage, codetimeeventcombinationtype, frontCodeTimeEventCombinationType)
	}
}

func (codetimeeventtype *CodeTimeEventType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeTimeEventTypeCreateCallback != nil {
		stage.OnAfterCodeTimeEventTypeCreateCallback.OnAfterCreate(stage, codetimeeventtype)
	}
}

func (codetimeeventtype *CodeTimeEventType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTimeEventTypeUpdateCallback != nil {
		var frontCodeTimeEventType *CodeTimeEventType
		if front != nil {
			frontCodeTimeEventType, _ = front.(*CodeTimeEventType)
		}
		stage.OnAfterCodeTimeEventTypeUpdateCallback.OnAfterUpdate(stage, codetimeeventtype, frontCodeTimeEventType)
	}
}

func (codetimeeventtype *CodeTimeEventType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTimeEventTypeDeleteCallback != nil {
		var frontCodeTimeEventType *CodeTimeEventType
		if front != nil {
			frontCodeTimeEventType, _ = front.(*CodeTimeEventType)
		}
		stage.OnAfterCodeTimeEventTypeDeleteCallback.OnAfterDelete(stage, codetimeeventtype, frontCodeTimeEventType)
	}
}

func (codetimereferencetype *CodeTimeReferenceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeTimeReferenceTypeCreateCallback != nil {
		stage.OnAfterCodeTimeReferenceTypeCreateCallback.OnAfterCreate(stage, codetimereferencetype)
	}
}

func (codetimereferencetype *CodeTimeReferenceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTimeReferenceTypeUpdateCallback != nil {
		var frontCodeTimeReferenceType *CodeTimeReferenceType
		if front != nil {
			frontCodeTimeReferenceType, _ = front.(*CodeTimeReferenceType)
		}
		stage.OnAfterCodeTimeReferenceTypeUpdateCallback.OnAfterUpdate(stage, codetimereferencetype, frontCodeTimeReferenceType)
	}
}

func (codetimereferencetype *CodeTimeReferenceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTimeReferenceTypeDeleteCallback != nil {
		var frontCodeTimeReferenceType *CodeTimeReferenceType
		if front != nil {
			frontCodeTimeReferenceType, _ = front.(*CodeTimeReferenceType)
		}
		stage.OnAfterCodeTimeReferenceTypeDeleteCallback.OnAfterDelete(stage, codetimereferencetype, frontCodeTimeReferenceType)
	}
}

func (codeunitdependencytype *CodeUnitDependencyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeUnitDependencyTypeCreateCallback != nil {
		stage.OnAfterCodeUnitDependencyTypeCreateCallback.OnAfterCreate(stage, codeunitdependencytype)
	}
}

func (codeunitdependencytype *CodeUnitDependencyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeUnitDependencyTypeUpdateCallback != nil {
		var frontCodeUnitDependencyType *CodeUnitDependencyType
		if front != nil {
			frontCodeUnitDependencyType, _ = front.(*CodeUnitDependencyType)
		}
		stage.OnAfterCodeUnitDependencyTypeUpdateCallback.OnAfterUpdate(stage, codeunitdependencytype, frontCodeUnitDependencyType)
	}
}

func (codeunitdependencytype *CodeUnitDependencyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeUnitDependencyTypeDeleteCallback != nil {
		var frontCodeUnitDependencyType *CodeUnitDependencyType
		if front != nil {
			frontCodeUnitDependencyType, _ = front.(*CodeUnitDependencyType)
		}
		stage.OnAfterCodeUnitDependencyTypeDeleteCallback.OnAfterDelete(stage, codeunitdependencytype, frontCodeUnitDependencyType)
	}
}

func (codeunittype *CodeUnitType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeUnitTypeCreateCallback != nil {
		stage.OnAfterCodeUnitTypeCreateCallback.OnAfterCreate(stage, codeunittype)
	}
}

func (codeunittype *CodeUnitType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeUnitTypeUpdateCallback != nil {
		var frontCodeUnitType *CodeUnitType
		if front != nil {
			frontCodeUnitType, _ = front.(*CodeUnitType)
		}
		stage.OnAfterCodeUnitTypeUpdateCallback.OnAfterUpdate(stage, codeunittype, frontCodeUnitType)
	}
}

func (codeunittype *CodeUnitType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeUnitTypeDeleteCallback != nil {
		var frontCodeUnitType *CodeUnitType
		if front != nil {
			frontCodeUnitType, _ = front.(*CodeUnitType)
		}
		stage.OnAfterCodeUnitTypeDeleteCallback.OnAfterDelete(stage, codeunittype, frontCodeUnitType)
	}
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeVerticalDatumTypeCreateCallback != nil {
		stage.OnAfterCodeVerticalDatumTypeCreateCallback.OnAfterCreate(stage, codeverticaldatumtype)
	}
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeVerticalDatumTypeUpdateCallback != nil {
		var frontCodeVerticalDatumType *CodeVerticalDatumType
		if front != nil {
			frontCodeVerticalDatumType, _ = front.(*CodeVerticalDatumType)
		}
		stage.OnAfterCodeVerticalDatumTypeUpdateCallback.OnAfterUpdate(stage, codeverticaldatumtype, frontCodeVerticalDatumType)
	}
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeVerticalDatumTypeDeleteCallback != nil {
		var frontCodeVerticalDatumType *CodeVerticalDatumType
		if front != nil {
			frontCodeVerticalDatumType, _ = front.(*CodeVerticalDatumType)
		}
		stage.OnAfterCodeVerticalDatumTypeDeleteCallback.OnAfterDelete(stage, codeverticaldatumtype, frontCodeVerticalDatumType)
	}
}

func (codeyesnotype *CodeYesNoType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeYesNoTypeCreateCallback != nil {
		stage.OnAfterCodeYesNoTypeCreateCallback.OnAfterCreate(stage, codeyesnotype)
	}
}

func (codeyesnotype *CodeYesNoType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeYesNoTypeUpdateCallback != nil {
		var frontCodeYesNoType *CodeYesNoType
		if front != nil {
			frontCodeYesNoType, _ = front.(*CodeYesNoType)
		}
		stage.OnAfterCodeYesNoTypeUpdateCallback.OnAfterUpdate(stage, codeyesnotype, frontCodeYesNoType)
	}
}

func (codeyesnotype *CodeYesNoType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeYesNoTypeDeleteCallback != nil {
		var frontCodeYesNoType *CodeYesNoType
		if front != nil {
			frontCodeYesNoType, _ = front.(*CodeYesNoType)
		}
		stage.OnAfterCodeYesNoTypeDeleteCallback.OnAfterDelete(stage, codeyesnotype, frontCodeYesNoType)
	}
}

func (datemonthdaytype *DateMonthDayType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDateMonthDayTypeCreateCallback != nil {
		stage.OnAfterDateMonthDayTypeCreateCallback.OnAfterCreate(stage, datemonthdaytype)
	}
}

func (datemonthdaytype *DateMonthDayType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDateMonthDayTypeUpdateCallback != nil {
		var frontDateMonthDayType *DateMonthDayType
		if front != nil {
			frontDateMonthDayType, _ = front.(*DateMonthDayType)
		}
		stage.OnAfterDateMonthDayTypeUpdateCallback.OnAfterUpdate(stage, datemonthdaytype, frontDateMonthDayType)
	}
}

func (datemonthdaytype *DateMonthDayType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDateMonthDayTypeDeleteCallback != nil {
		var frontDateMonthDayType *DateMonthDayType
		if front != nil {
			frontDateMonthDayType, _ = front.(*DateMonthDayType)
		}
		stage.OnAfterDateMonthDayTypeDeleteCallback.OnAfterDelete(stage, datemonthdaytype, frontDateMonthDayType)
	}
}

func (dateyeartype *DateYearType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDateYearTypeCreateCallback != nil {
		stage.OnAfterDateYearTypeCreateCallback.OnAfterCreate(stage, dateyeartype)
	}
}

func (dateyeartype *DateYearType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDateYearTypeUpdateCallback != nil {
		var frontDateYearType *DateYearType
		if front != nil {
			frontDateYearType, _ = front.(*DateYearType)
		}
		stage.OnAfterDateYearTypeUpdateCallback.OnAfterUpdate(stage, dateyeartype, frontDateYearType)
	}
}

func (dateyeartype *DateYearType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDateYearTypeDeleteCallback != nil {
		var frontDateYearType *DateYearType
		if front != nil {
			frontDateYearType, _ = front.(*DateYearType)
		}
		stage.OnAfterDateYearTypeDeleteCallback.OnAfterDelete(stage, dateyeartype, frontDateYearType)
	}
}

func (nonumbertype *NoNumberType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoNumberTypeCreateCallback != nil {
		stage.OnAfterNoNumberTypeCreateCallback.OnAfterCreate(stage, nonumbertype)
	}
}

func (nonumbertype *NoNumberType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoNumberTypeUpdateCallback != nil {
		var frontNoNumberType *NoNumberType
		if front != nil {
			frontNoNumberType, _ = front.(*NoNumberType)
		}
		stage.OnAfterNoNumberTypeUpdateCallback.OnAfterUpdate(stage, nonumbertype, frontNoNumberType)
	}
}

func (nonumbertype *NoNumberType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoNumberTypeDeleteCallback != nil {
		var frontNoNumberType *NoNumberType
		if front != nil {
			frontNoNumberType, _ = front.(*NoNumberType)
		}
		stage.OnAfterNoNumberTypeDeleteCallback.OnAfterDelete(stage, nonumbertype, frontNoNumberType)
	}
}

func (textaddresstype *TextAddressType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTextAddressTypeCreateCallback != nil {
		stage.OnAfterTextAddressTypeCreateCallback.OnAfterCreate(stage, textaddresstype)
	}
}

func (textaddresstype *TextAddressType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextAddressTypeUpdateCallback != nil {
		var frontTextAddressType *TextAddressType
		if front != nil {
			frontTextAddressType, _ = front.(*TextAddressType)
		}
		stage.OnAfterTextAddressTypeUpdateCallback.OnAfterUpdate(stage, textaddresstype, frontTextAddressType)
	}
}

func (textaddresstype *TextAddressType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextAddressTypeDeleteCallback != nil {
		var frontTextAddressType *TextAddressType
		if front != nil {
			frontTextAddressType, _ = front.(*TextAddressType)
		}
		stage.OnAfterTextAddressTypeDeleteCallback.OnAfterDelete(stage, textaddresstype, frontTextAddressType)
	}
}

func (textdesignatortype *TextDesignatorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTextDesignatorTypeCreateCallback != nil {
		stage.OnAfterTextDesignatorTypeCreateCallback.OnAfterCreate(stage, textdesignatortype)
	}
}

func (textdesignatortype *TextDesignatorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextDesignatorTypeUpdateCallback != nil {
		var frontTextDesignatorType *TextDesignatorType
		if front != nil {
			frontTextDesignatorType, _ = front.(*TextDesignatorType)
		}
		stage.OnAfterTextDesignatorTypeUpdateCallback.OnAfterUpdate(stage, textdesignatortype, frontTextDesignatorType)
	}
}

func (textdesignatortype *TextDesignatorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextDesignatorTypeDeleteCallback != nil {
		var frontTextDesignatorType *TextDesignatorType
		if front != nil {
			frontTextDesignatorType, _ = front.(*TextDesignatorType)
		}
		stage.OnAfterTextDesignatorTypeDeleteCallback.OnAfterDelete(stage, textdesignatortype, frontTextDesignatorType)
	}
}

func (textinstructiontype *TextInstructionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTextInstructionTypeCreateCallback != nil {
		stage.OnAfterTextInstructionTypeCreateCallback.OnAfterCreate(stage, textinstructiontype)
	}
}

func (textinstructiontype *TextInstructionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextInstructionTypeUpdateCallback != nil {
		var frontTextInstructionType *TextInstructionType
		if front != nil {
			frontTextInstructionType, _ = front.(*TextInstructionType)
		}
		stage.OnAfterTextInstructionTypeUpdateCallback.OnAfterUpdate(stage, textinstructiontype, frontTextInstructionType)
	}
}

func (textinstructiontype *TextInstructionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextInstructionTypeDeleteCallback != nil {
		var frontTextInstructionType *TextInstructionType
		if front != nil {
			frontTextInstructionType, _ = front.(*TextInstructionType)
		}
		stage.OnAfterTextInstructionTypeDeleteCallback.OnAfterDelete(stage, textinstructiontype, frontTextInstructionType)
	}
}

func (textnametype *TextNameType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTextNameTypeCreateCallback != nil {
		stage.OnAfterTextNameTypeCreateCallback.OnAfterCreate(stage, textnametype)
	}
}

func (textnametype *TextNameType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextNameTypeUpdateCallback != nil {
		var frontTextNameType *TextNameType
		if front != nil {
			frontTextNameType, _ = front.(*TextNameType)
		}
		stage.OnAfterTextNameTypeUpdateCallback.OnAfterUpdate(stage, textnametype, frontTextNameType)
	}
}

func (textnametype *TextNameType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextNameTypeDeleteCallback != nil {
		var frontTextNameType *TextNameType
		if front != nil {
			frontTextNameType, _ = front.(*TextNameType)
		}
		stage.OnAfterTextNameTypeDeleteCallback.OnAfterDelete(stage, textnametype, frontTextNameType)
	}
}

func (textnotetype *TextNoteType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTextNoteTypeCreateCallback != nil {
		stage.OnAfterTextNoteTypeCreateCallback.OnAfterCreate(stage, textnotetype)
	}
}

func (textnotetype *TextNoteType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextNoteTypeUpdateCallback != nil {
		var frontTextNoteType *TextNoteType
		if front != nil {
			frontTextNoteType, _ = front.(*TextNoteType)
		}
		stage.OnAfterTextNoteTypeUpdateCallback.OnAfterUpdate(stage, textnotetype, frontTextNoteType)
	}
}

func (textnotetype *TextNoteType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextNoteTypeDeleteCallback != nil {
		var frontTextNoteType *TextNoteType
		if front != nil {
			frontTextNoteType, _ = front.(*TextNoteType)
		}
		stage.OnAfterTextNoteTypeDeleteCallback.OnAfterDelete(stage, textnotetype, frontTextNoteType)
	}
}

func (textphonetype *TextPhoneType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTextPhoneTypeCreateCallback != nil {
		stage.OnAfterTextPhoneTypeCreateCallback.OnAfterCreate(stage, textphonetype)
	}
}

func (textphonetype *TextPhoneType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextPhoneTypeUpdateCallback != nil {
		var frontTextPhoneType *TextPhoneType
		if front != nil {
			frontTextPhoneType, _ = front.(*TextPhoneType)
		}
		stage.OnAfterTextPhoneTypeUpdateCallback.OnAfterUpdate(stage, textphonetype, frontTextPhoneType)
	}
}

func (textphonetype *TextPhoneType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextPhoneTypeDeleteCallback != nil {
		var frontTextPhoneType *TextPhoneType
		if front != nil {
			frontTextPhoneType, _ = front.(*TextPhoneType)
		}
		stage.OnAfterTextPhoneTypeDeleteCallback.OnAfterDelete(stage, textphonetype, frontTextPhoneType)
	}
}

func (textpropertynametype *TextPropertyNameType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTextPropertyNameTypeCreateCallback != nil {
		stage.OnAfterTextPropertyNameTypeCreateCallback.OnAfterCreate(stage, textpropertynametype)
	}
}

func (textpropertynametype *TextPropertyNameType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextPropertyNameTypeUpdateCallback != nil {
		var frontTextPropertyNameType *TextPropertyNameType
		if front != nil {
			frontTextPropertyNameType, _ = front.(*TextPropertyNameType)
		}
		stage.OnAfterTextPropertyNameTypeUpdateCallback.OnAfterUpdate(stage, textpropertynametype, frontTextPropertyNameType)
	}
}

func (textpropertynametype *TextPropertyNameType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextPropertyNameTypeDeleteCallback != nil {
		var frontTextPropertyNameType *TextPropertyNameType
		if front != nil {
			frontTextPropertyNameType, _ = front.(*TextPropertyNameType)
		}
		stage.OnAfterTextPropertyNameTypeDeleteCallback.OnAfterDelete(stage, textpropertynametype, frontTextPropertyNameType)
	}
}

func (timetype *TimeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimeTypeCreateCallback != nil {
		stage.OnAfterTimeTypeCreateCallback.OnAfterCreate(stage, timetype)
	}
}

func (timetype *TimeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeTypeUpdateCallback != nil {
		var frontTimeType *TimeType
		if front != nil {
			frontTimeType, _ = front.(*TimeType)
		}
		stage.OnAfterTimeTypeUpdateCallback.OnAfterUpdate(stage, timetype, frontTimeType)
	}
}

func (timetype *TimeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeTypeDeleteCallback != nil {
		var frontTimeType *TimeType
		if front != nil {
			frontTimeType, _ = front.(*TimeType)
		}
		stage.OnAfterTimeTypeDeleteCallback.OnAfterDelete(stage, timetype, frontTimeType)
	}
}

func (valbearingtype *ValBearingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValBearingTypeCreateCallback != nil {
		stage.OnAfterValBearingTypeCreateCallback.OnAfterCreate(stage, valbearingtype)
	}
}

func (valbearingtype *ValBearingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValBearingTypeUpdateCallback != nil {
		var frontValBearingType *ValBearingType
		if front != nil {
			frontValBearingType, _ = front.(*ValBearingType)
		}
		stage.OnAfterValBearingTypeUpdateCallback.OnAfterUpdate(stage, valbearingtype, frontValBearingType)
	}
}

func (valbearingtype *ValBearingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValBearingTypeDeleteCallback != nil {
		var frontValBearingType *ValBearingType
		if front != nil {
			frontValBearingType, _ = front.(*ValBearingType)
		}
		stage.OnAfterValBearingTypeDeleteCallback.OnAfterDelete(stage, valbearingtype, frontValBearingType)
	}
}

func (valdistancesignedtype *ValDistanceSignedType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValDistanceSignedTypeCreateCallback != nil {
		stage.OnAfterValDistanceSignedTypeCreateCallback.OnAfterCreate(stage, valdistancesignedtype)
	}
}

func (valdistancesignedtype *ValDistanceSignedType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValDistanceSignedTypeUpdateCallback != nil {
		var frontValDistanceSignedType *ValDistanceSignedType
		if front != nil {
			frontValDistanceSignedType, _ = front.(*ValDistanceSignedType)
		}
		stage.OnAfterValDistanceSignedTypeUpdateCallback.OnAfterUpdate(stage, valdistancesignedtype, frontValDistanceSignedType)
	}
}

func (valdistancesignedtype *ValDistanceSignedType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValDistanceSignedTypeDeleteCallback != nil {
		var frontValDistanceSignedType *ValDistanceSignedType
		if front != nil {
			frontValDistanceSignedType, _ = front.(*ValDistanceSignedType)
		}
		stage.OnAfterValDistanceSignedTypeDeleteCallback.OnAfterDelete(stage, valdistancesignedtype, frontValDistanceSignedType)
	}
}

func (valdistancetype *ValDistanceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValDistanceTypeCreateCallback != nil {
		stage.OnAfterValDistanceTypeCreateCallback.OnAfterCreate(stage, valdistancetype)
	}
}

func (valdistancetype *ValDistanceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValDistanceTypeUpdateCallback != nil {
		var frontValDistanceType *ValDistanceType
		if front != nil {
			frontValDistanceType, _ = front.(*ValDistanceType)
		}
		stage.OnAfterValDistanceTypeUpdateCallback.OnAfterUpdate(stage, valdistancetype, frontValDistanceType)
	}
}

func (valdistancetype *ValDistanceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValDistanceTypeDeleteCallback != nil {
		var frontValDistanceType *ValDistanceType
		if front != nil {
			frontValDistanceType, _ = front.(*ValDistanceType)
		}
		stage.OnAfterValDistanceTypeDeleteCallback.OnAfterDelete(stage, valdistancetype, frontValDistanceType)
	}
}

func (valdistanceverticaltype *ValDistanceVerticalType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValDistanceVerticalTypeCreateCallback != nil {
		stage.OnAfterValDistanceVerticalTypeCreateCallback.OnAfterCreate(stage, valdistanceverticaltype)
	}
}

func (valdistanceverticaltype *ValDistanceVerticalType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValDistanceVerticalTypeUpdateCallback != nil {
		var frontValDistanceVerticalType *ValDistanceVerticalType
		if front != nil {
			frontValDistanceVerticalType, _ = front.(*ValDistanceVerticalType)
		}
		stage.OnAfterValDistanceVerticalTypeUpdateCallback.OnAfterUpdate(stage, valdistanceverticaltype, frontValDistanceVerticalType)
	}
}

func (valdistanceverticaltype *ValDistanceVerticalType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValDistanceVerticalTypeDeleteCallback != nil {
		var frontValDistanceVerticalType *ValDistanceVerticalType
		if front != nil {
			frontValDistanceVerticalType, _ = front.(*ValDistanceVerticalType)
		}
		stage.OnAfterValDistanceVerticalTypeDeleteCallback.OnAfterDelete(stage, valdistanceverticaltype, frontValDistanceVerticalType)
	}
}

func (valdurationtype *ValDurationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValDurationTypeCreateCallback != nil {
		stage.OnAfterValDurationTypeCreateCallback.OnAfterCreate(stage, valdurationtype)
	}
}

func (valdurationtype *ValDurationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValDurationTypeUpdateCallback != nil {
		var frontValDurationType *ValDurationType
		if front != nil {
			frontValDurationType, _ = front.(*ValDurationType)
		}
		stage.OnAfterValDurationTypeUpdateCallback.OnAfterUpdate(stage, valdurationtype, frontValDurationType)
	}
}

func (valdurationtype *ValDurationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValDurationTypeDeleteCallback != nil {
		var frontValDurationType *ValDurationType
		if front != nil {
			frontValDurationType, _ = front.(*ValDurationType)
		}
		stage.OnAfterValDurationTypeDeleteCallback.OnAfterDelete(stage, valdurationtype, frontValDurationType)
	}
}

func (valfrequencytype *ValFrequencyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValFrequencyTypeCreateCallback != nil {
		stage.OnAfterValFrequencyTypeCreateCallback.OnAfterCreate(stage, valfrequencytype)
	}
}

func (valfrequencytype *ValFrequencyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValFrequencyTypeUpdateCallback != nil {
		var frontValFrequencyType *ValFrequencyType
		if front != nil {
			frontValFrequencyType, _ = front.(*ValFrequencyType)
		}
		stage.OnAfterValFrequencyTypeUpdateCallback.OnAfterUpdate(stage, valfrequencytype, frontValFrequencyType)
	}
}

func (valfrequencytype *ValFrequencyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValFrequencyTypeDeleteCallback != nil {
		var frontValFrequencyType *ValFrequencyType
		if front != nil {
			frontValFrequencyType, _ = front.(*ValFrequencyType)
		}
		stage.OnAfterValFrequencyTypeDeleteCallback.OnAfterDelete(stage, valfrequencytype, frontValFrequencyType)
	}
}

