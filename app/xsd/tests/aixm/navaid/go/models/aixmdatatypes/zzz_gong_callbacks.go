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
func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAirspacePointPositionTypeCreateCallback != nil {
		stage.OnAfterCodeAirspacePointPositionTypeCreateCallback.OnAfterCreate(stage, codeairspacepointpositiontype)
	}
}

func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAirspacePointPositionTypeUpdateCallback != nil {
		var frontCodeAirspacePointPositionType *CodeAirspacePointPositionType
		if front != nil {
			frontCodeAirspacePointPositionType, _ = front.(*CodeAirspacePointPositionType)
		}
		stage.OnAfterCodeAirspacePointPositionTypeUpdateCallback.OnAfterUpdate(stage, codeairspacepointpositiontype, frontCodeAirspacePointPositionType)
	}
}

func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAirspacePointPositionTypeDeleteCallback != nil {
		var frontCodeAirspacePointPositionType *CodeAirspacePointPositionType
		if front != nil {
			frontCodeAirspacePointPositionType, _ = front.(*CodeAirspacePointPositionType)
		}
		stage.OnAfterCodeAirspacePointPositionTypeDeleteCallback.OnAfterDelete(stage, codeairspacepointpositiontype, frontCodeAirspacePointPositionType)
	}
}

func (codeairspacepointroletype *CodeAirspacePointRoleType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAirspacePointRoleTypeCreateCallback != nil {
		stage.OnAfterCodeAirspacePointRoleTypeCreateCallback.OnAfterCreate(stage, codeairspacepointroletype)
	}
}

func (codeairspacepointroletype *CodeAirspacePointRoleType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAirspacePointRoleTypeUpdateCallback != nil {
		var frontCodeAirspacePointRoleType *CodeAirspacePointRoleType
		if front != nil {
			frontCodeAirspacePointRoleType, _ = front.(*CodeAirspacePointRoleType)
		}
		stage.OnAfterCodeAirspacePointRoleTypeUpdateCallback.OnAfterUpdate(stage, codeairspacepointroletype, frontCodeAirspacePointRoleType)
	}
}

func (codeairspacepointroletype *CodeAirspacePointRoleType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAirspacePointRoleTypeDeleteCallback != nil {
		var frontCodeAirspacePointRoleType *CodeAirspacePointRoleType
		if front != nil {
			frontCodeAirspacePointRoleType, _ = front.(*CodeAirspacePointRoleType)
		}
		stage.OnAfterCodeAirspacePointRoleTypeDeleteCallback.OnAfterDelete(stage, codeairspacepointroletype, frontCodeAirspacePointRoleType)
	}
}

func (codealtitudeusetype *CodeAltitudeUseType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAltitudeUseTypeCreateCallback != nil {
		stage.OnAfterCodeAltitudeUseTypeCreateCallback.OnAfterCreate(stage, codealtitudeusetype)
	}
}

func (codealtitudeusetype *CodeAltitudeUseType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAltitudeUseTypeUpdateCallback != nil {
		var frontCodeAltitudeUseType *CodeAltitudeUseType
		if front != nil {
			frontCodeAltitudeUseType, _ = front.(*CodeAltitudeUseType)
		}
		stage.OnAfterCodeAltitudeUseTypeUpdateCallback.OnAfterUpdate(stage, codealtitudeusetype, frontCodeAltitudeUseType)
	}
}

func (codealtitudeusetype *CodeAltitudeUseType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAltitudeUseTypeDeleteCallback != nil {
		var frontCodeAltitudeUseType *CodeAltitudeUseType
		if front != nil {
			frontCodeAltitudeUseType, _ = front.(*CodeAltitudeUseType)
		}
		stage.OnAfterCodeAltitudeUseTypeDeleteCallback.OnAfterDelete(stage, codealtitudeusetype, frontCodeAltitudeUseType)
	}
}

func (codeauralmorsetype *CodeAuralMorseType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAuralMorseTypeCreateCallback != nil {
		stage.OnAfterCodeAuralMorseTypeCreateCallback.OnAfterCreate(stage, codeauralmorsetype)
	}
}

func (codeauralmorsetype *CodeAuralMorseType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAuralMorseTypeUpdateCallback != nil {
		var frontCodeAuralMorseType *CodeAuralMorseType
		if front != nil {
			frontCodeAuralMorseType, _ = front.(*CodeAuralMorseType)
		}
		stage.OnAfterCodeAuralMorseTypeUpdateCallback.OnAfterUpdate(stage, codeauralmorsetype, frontCodeAuralMorseType)
	}
}

func (codeauralmorsetype *CodeAuralMorseType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAuralMorseTypeDeleteCallback != nil {
		var frontCodeAuralMorseType *CodeAuralMorseType
		if front != nil {
			frontCodeAuralMorseType, _ = front.(*CodeAuralMorseType)
		}
		stage.OnAfterCodeAuralMorseTypeDeleteCallback.OnAfterDelete(stage, codeauralmorsetype, frontCodeAuralMorseType)
	}
}

func (codeauthorityroletype *CodeAuthorityRoleType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAuthorityRoleTypeCreateCallback != nil {
		stage.OnAfterCodeAuthorityRoleTypeCreateCallback.OnAfterCreate(stage, codeauthorityroletype)
	}
}

func (codeauthorityroletype *CodeAuthorityRoleType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAuthorityRoleTypeUpdateCallback != nil {
		var frontCodeAuthorityRoleType *CodeAuthorityRoleType
		if front != nil {
			frontCodeAuthorityRoleType, _ = front.(*CodeAuthorityRoleType)
		}
		stage.OnAfterCodeAuthorityRoleTypeUpdateCallback.OnAfterUpdate(stage, codeauthorityroletype, frontCodeAuthorityRoleType)
	}
}

func (codeauthorityroletype *CodeAuthorityRoleType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAuthorityRoleTypeDeleteCallback != nil {
		var frontCodeAuthorityRoleType *CodeAuthorityRoleType
		if front != nil {
			frontCodeAuthorityRoleType, _ = front.(*CodeAuthorityRoleType)
		}
		stage.OnAfterCodeAuthorityRoleTypeDeleteCallback.OnAfterDelete(stage, codeauthorityroletype, frontCodeAuthorityRoleType)
	}
}

func (codebearingtype *CodeBearingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeBearingTypeCreateCallback != nil {
		stage.OnAfterCodeBearingTypeCreateCallback.OnAfterCreate(stage, codebearingtype)
	}
}

func (codebearingtype *CodeBearingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeBearingTypeUpdateCallback != nil {
		var frontCodeBearingType *CodeBearingType
		if front != nil {
			frontCodeBearingType, _ = front.(*CodeBearingType)
		}
		stage.OnAfterCodeBearingTypeUpdateCallback.OnAfterUpdate(stage, codebearingtype, frontCodeBearingType)
	}
}

func (codebearingtype *CodeBearingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeBearingTypeDeleteCallback != nil {
		var frontCodeBearingType *CodeBearingType
		if front != nil {
			frontCodeBearingType, _ = front.(*CodeBearingType)
		}
		stage.OnAfterCodeBearingTypeDeleteCallback.OnAfterDelete(stage, codebearingtype, frontCodeBearingType)
	}
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeCardinalDirectionTypeCreateCallback != nil {
		stage.OnAfterCodeCardinalDirectionTypeCreateCallback.OnAfterCreate(stage, codecardinaldirectiontype)
	}
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeCardinalDirectionTypeUpdateCallback != nil {
		var frontCodeCardinalDirectionType *CodeCardinalDirectionType
		if front != nil {
			frontCodeCardinalDirectionType, _ = front.(*CodeCardinalDirectionType)
		}
		stage.OnAfterCodeCardinalDirectionTypeUpdateCallback.OnAfterUpdate(stage, codecardinaldirectiontype, frontCodeCardinalDirectionType)
	}
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeCardinalDirectionTypeDeleteCallback != nil {
		var frontCodeCardinalDirectionType *CodeCardinalDirectionType
		if front != nil {
			frontCodeCardinalDirectionType, _ = front.(*CodeCardinalDirectionType)
		}
		stage.OnAfterCodeCardinalDirectionTypeDeleteCallback.OnAfterDelete(stage, codecardinaldirectiontype, frontCodeCardinalDirectionType)
	}
}

func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeCheckpointCategoryTypeCreateCallback != nil {
		stage.OnAfterCodeCheckpointCategoryTypeCreateCallback.OnAfterCreate(stage, codecheckpointcategorytype)
	}
}

func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeCheckpointCategoryTypeUpdateCallback != nil {
		var frontCodeCheckpointCategoryType *CodeCheckpointCategoryType
		if front != nil {
			frontCodeCheckpointCategoryType, _ = front.(*CodeCheckpointCategoryType)
		}
		stage.OnAfterCodeCheckpointCategoryTypeUpdateCallback.OnAfterUpdate(stage, codecheckpointcategorytype, frontCodeCheckpointCategoryType)
	}
}

func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeCheckpointCategoryTypeDeleteCallback != nil {
		var frontCodeCheckpointCategoryType *CodeCheckpointCategoryType
		if front != nil {
			frontCodeCheckpointCategoryType, _ = front.(*CodeCheckpointCategoryType)
		}
		stage.OnAfterCodeCheckpointCategoryTypeDeleteCallback.OnAfterDelete(stage, codecheckpointcategorytype, frontCodeCheckpointCategoryType)
	}
}

func (codecoursequalityilstype *CodeCourseQualityILSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeCourseQualityILSTypeCreateCallback != nil {
		stage.OnAfterCodeCourseQualityILSTypeCreateCallback.OnAfterCreate(stage, codecoursequalityilstype)
	}
}

func (codecoursequalityilstype *CodeCourseQualityILSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeCourseQualityILSTypeUpdateCallback != nil {
		var frontCodeCourseQualityILSType *CodeCourseQualityILSType
		if front != nil {
			frontCodeCourseQualityILSType, _ = front.(*CodeCourseQualityILSType)
		}
		stage.OnAfterCodeCourseQualityILSTypeUpdateCallback.OnAfterUpdate(stage, codecoursequalityilstype, frontCodeCourseQualityILSType)
	}
}

func (codecoursequalityilstype *CodeCourseQualityILSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeCourseQualityILSTypeDeleteCallback != nil {
		var frontCodeCourseQualityILSType *CodeCourseQualityILSType
		if front != nil {
			frontCodeCourseQualityILSType, _ = front.(*CodeCourseQualityILSType)
		}
		stage.OnAfterCodeCourseQualityILSTypeDeleteCallback.OnAfterDelete(stage, codecoursequalityilstype, frontCodeCourseQualityILSType)
	}
}

func (codedmechanneltype *CodeDMEChannelType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeDMEChannelTypeCreateCallback != nil {
		stage.OnAfterCodeDMEChannelTypeCreateCallback.OnAfterCreate(stage, codedmechanneltype)
	}
}

func (codedmechanneltype *CodeDMEChannelType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDMEChannelTypeUpdateCallback != nil {
		var frontCodeDMEChannelType *CodeDMEChannelType
		if front != nil {
			frontCodeDMEChannelType, _ = front.(*CodeDMEChannelType)
		}
		stage.OnAfterCodeDMEChannelTypeUpdateCallback.OnAfterUpdate(stage, codedmechanneltype, frontCodeDMEChannelType)
	}
}

func (codedmechanneltype *CodeDMEChannelType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDMEChannelTypeDeleteCallback != nil {
		var frontCodeDMEChannelType *CodeDMEChannelType
		if front != nil {
			frontCodeDMEChannelType, _ = front.(*CodeDMEChannelType)
		}
		stage.OnAfterCodeDMEChannelTypeDeleteCallback.OnAfterDelete(stage, codedmechanneltype, frontCodeDMEChannelType)
	}
}

func (codedmetype *CodeDMEType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeDMETypeCreateCallback != nil {
		stage.OnAfterCodeDMETypeCreateCallback.OnAfterCreate(stage, codedmetype)
	}
}

func (codedmetype *CodeDMEType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDMETypeUpdateCallback != nil {
		var frontCodeDMEType *CodeDMEType
		if front != nil {
			frontCodeDMEType, _ = front.(*CodeDMEType)
		}
		stage.OnAfterCodeDMETypeUpdateCallback.OnAfterUpdate(stage, codedmetype, frontCodeDMEType)
	}
}

func (codedmetype *CodeDMEType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDMETypeDeleteCallback != nil {
		var frontCodeDMEType *CodeDMEType
		if front != nil {
			frontCodeDMEType, _ = front.(*CodeDMEType)
		}
		stage.OnAfterCodeDMETypeDeleteCallback.OnAfterDelete(stage, codedmetype, frontCodeDMEType)
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

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeDesignatedPointDesignatorTypeCreateCallback != nil {
		stage.OnAfterCodeDesignatedPointDesignatorTypeCreateCallback.OnAfterCreate(stage, codedesignatedpointdesignatortype)
	}
}

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDesignatedPointDesignatorTypeUpdateCallback != nil {
		var frontCodeDesignatedPointDesignatorType *CodeDesignatedPointDesignatorType
		if front != nil {
			frontCodeDesignatedPointDesignatorType, _ = front.(*CodeDesignatedPointDesignatorType)
		}
		stage.OnAfterCodeDesignatedPointDesignatorTypeUpdateCallback.OnAfterUpdate(stage, codedesignatedpointdesignatortype, frontCodeDesignatedPointDesignatorType)
	}
}

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDesignatedPointDesignatorTypeDeleteCallback != nil {
		var frontCodeDesignatedPointDesignatorType *CodeDesignatedPointDesignatorType
		if front != nil {
			frontCodeDesignatedPointDesignatorType, _ = front.(*CodeDesignatedPointDesignatorType)
		}
		stage.OnAfterCodeDesignatedPointDesignatorTypeDeleteCallback.OnAfterDelete(stage, codedesignatedpointdesignatortype, frontCodeDesignatedPointDesignatorType)
	}
}

func (codedesignatedpointtype *CodeDesignatedPointType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeDesignatedPointTypeCreateCallback != nil {
		stage.OnAfterCodeDesignatedPointTypeCreateCallback.OnAfterCreate(stage, codedesignatedpointtype)
	}
}

func (codedesignatedpointtype *CodeDesignatedPointType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDesignatedPointTypeUpdateCallback != nil {
		var frontCodeDesignatedPointType *CodeDesignatedPointType
		if front != nil {
			frontCodeDesignatedPointType, _ = front.(*CodeDesignatedPointType)
		}
		stage.OnAfterCodeDesignatedPointTypeUpdateCallback.OnAfterUpdate(stage, codedesignatedpointtype, frontCodeDesignatedPointType)
	}
}

func (codedesignatedpointtype *CodeDesignatedPointType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDesignatedPointTypeDeleteCallback != nil {
		var frontCodeDesignatedPointType *CodeDesignatedPointType
		if front != nil {
			frontCodeDesignatedPointType, _ = front.(*CodeDesignatedPointType)
		}
		stage.OnAfterCodeDesignatedPointTypeDeleteCallback.OnAfterDelete(stage, codedesignatedpointtype, frontCodeDesignatedPointType)
	}
}

func (codedirectionreferencetype *CodeDirectionReferenceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeDirectionReferenceTypeCreateCallback != nil {
		stage.OnAfterCodeDirectionReferenceTypeCreateCallback.OnAfterCreate(stage, codedirectionreferencetype)
	}
}

func (codedirectionreferencetype *CodeDirectionReferenceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDirectionReferenceTypeUpdateCallback != nil {
		var frontCodeDirectionReferenceType *CodeDirectionReferenceType
		if front != nil {
			frontCodeDirectionReferenceType, _ = front.(*CodeDirectionReferenceType)
		}
		stage.OnAfterCodeDirectionReferenceTypeUpdateCallback.OnAfterUpdate(stage, codedirectionreferencetype, frontCodeDirectionReferenceType)
	}
}

func (codedirectionreferencetype *CodeDirectionReferenceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDirectionReferenceTypeDeleteCallback != nil {
		var frontCodeDirectionReferenceType *CodeDirectionReferenceType
		if front != nil {
			frontCodeDirectionReferenceType, _ = front.(*CodeDirectionReferenceType)
		}
		stage.OnAfterCodeDirectionReferenceTypeDeleteCallback.OnAfterDelete(stage, codedirectionreferencetype, frontCodeDirectionReferenceType)
	}
}

func (codedistanceindicationtype *CodeDistanceIndicationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeDistanceIndicationTypeCreateCallback != nil {
		stage.OnAfterCodeDistanceIndicationTypeCreateCallback.OnAfterCreate(stage, codedistanceindicationtype)
	}
}

func (codedistanceindicationtype *CodeDistanceIndicationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDistanceIndicationTypeUpdateCallback != nil {
		var frontCodeDistanceIndicationType *CodeDistanceIndicationType
		if front != nil {
			frontCodeDistanceIndicationType, _ = front.(*CodeDistanceIndicationType)
		}
		stage.OnAfterCodeDistanceIndicationTypeUpdateCallback.OnAfterUpdate(stage, codedistanceindicationtype, frontCodeDistanceIndicationType)
	}
}

func (codedistanceindicationtype *CodeDistanceIndicationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDistanceIndicationTypeDeleteCallback != nil {
		var frontCodeDistanceIndicationType *CodeDistanceIndicationType
		if front != nil {
			frontCodeDistanceIndicationType, _ = front.(*CodeDistanceIndicationType)
		}
		stage.OnAfterCodeDistanceIndicationTypeDeleteCallback.OnAfterDelete(stage, codedistanceindicationtype, frontCodeDistanceIndicationType)
	}
}

func (codeemissionbandtype *CodeEmissionBandType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeEmissionBandTypeCreateCallback != nil {
		stage.OnAfterCodeEmissionBandTypeCreateCallback.OnAfterCreate(stage, codeemissionbandtype)
	}
}

func (codeemissionbandtype *CodeEmissionBandType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeEmissionBandTypeUpdateCallback != nil {
		var frontCodeEmissionBandType *CodeEmissionBandType
		if front != nil {
			frontCodeEmissionBandType, _ = front.(*CodeEmissionBandType)
		}
		stage.OnAfterCodeEmissionBandTypeUpdateCallback.OnAfterUpdate(stage, codeemissionbandtype, frontCodeEmissionBandType)
	}
}

func (codeemissionbandtype *CodeEmissionBandType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeEmissionBandTypeDeleteCallback != nil {
		var frontCodeEmissionBandType *CodeEmissionBandType
		if front != nil {
			frontCodeEmissionBandType, _ = front.(*CodeEmissionBandType)
		}
		stage.OnAfterCodeEmissionBandTypeDeleteCallback.OnAfterDelete(stage, codeemissionbandtype, frontCodeEmissionBandType)
	}
}

func (codeilsbackcoursetype *CodeILSBackCourseType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeILSBackCourseTypeCreateCallback != nil {
		stage.OnAfterCodeILSBackCourseTypeCreateCallback.OnAfterCreate(stage, codeilsbackcoursetype)
	}
}

func (codeilsbackcoursetype *CodeILSBackCourseType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeILSBackCourseTypeUpdateCallback != nil {
		var frontCodeILSBackCourseType *CodeILSBackCourseType
		if front != nil {
			frontCodeILSBackCourseType, _ = front.(*CodeILSBackCourseType)
		}
		stage.OnAfterCodeILSBackCourseTypeUpdateCallback.OnAfterUpdate(stage, codeilsbackcoursetype, frontCodeILSBackCourseType)
	}
}

func (codeilsbackcoursetype *CodeILSBackCourseType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeILSBackCourseTypeDeleteCallback != nil {
		var frontCodeILSBackCourseType *CodeILSBackCourseType
		if front != nil {
			frontCodeILSBackCourseType, _ = front.(*CodeILSBackCourseType)
		}
		stage.OnAfterCodeILSBackCourseTypeDeleteCallback.OnAfterDelete(stage, codeilsbackcoursetype, frontCodeILSBackCourseType)
	}
}

func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeIntegrityLevelILSTypeCreateCallback != nil {
		stage.OnAfterCodeIntegrityLevelILSTypeCreateCallback.OnAfterCreate(stage, codeintegritylevelilstype)
	}
}

func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeIntegrityLevelILSTypeUpdateCallback != nil {
		var frontCodeIntegrityLevelILSType *CodeIntegrityLevelILSType
		if front != nil {
			frontCodeIntegrityLevelILSType, _ = front.(*CodeIntegrityLevelILSType)
		}
		stage.OnAfterCodeIntegrityLevelILSTypeUpdateCallback.OnAfterUpdate(stage, codeintegritylevelilstype, frontCodeIntegrityLevelILSType)
	}
}

func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeIntegrityLevelILSTypeDeleteCallback != nil {
		var frontCodeIntegrityLevelILSType *CodeIntegrityLevelILSType
		if front != nil {
			frontCodeIntegrityLevelILSType, _ = front.(*CodeIntegrityLevelILSType)
		}
		stage.OnAfterCodeIntegrityLevelILSTypeDeleteCallback.OnAfterDelete(stage, codeintegritylevelilstype, frontCodeIntegrityLevelILSType)
	}
}

func (codemlsazimuthtype *CodeMLSAzimuthType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeMLSAzimuthTypeCreateCallback != nil {
		stage.OnAfterCodeMLSAzimuthTypeCreateCallback.OnAfterCreate(stage, codemlsazimuthtype)
	}
}

func (codemlsazimuthtype *CodeMLSAzimuthType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMLSAzimuthTypeUpdateCallback != nil {
		var frontCodeMLSAzimuthType *CodeMLSAzimuthType
		if front != nil {
			frontCodeMLSAzimuthType, _ = front.(*CodeMLSAzimuthType)
		}
		stage.OnAfterCodeMLSAzimuthTypeUpdateCallback.OnAfterUpdate(stage, codemlsazimuthtype, frontCodeMLSAzimuthType)
	}
}

func (codemlsazimuthtype *CodeMLSAzimuthType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMLSAzimuthTypeDeleteCallback != nil {
		var frontCodeMLSAzimuthType *CodeMLSAzimuthType
		if front != nil {
			frontCodeMLSAzimuthType, _ = front.(*CodeMLSAzimuthType)
		}
		stage.OnAfterCodeMLSAzimuthTypeDeleteCallback.OnAfterDelete(stage, codemlsazimuthtype, frontCodeMLSAzimuthType)
	}
}

func (codemlschanneltype *CodeMLSChannelType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeMLSChannelTypeCreateCallback != nil {
		stage.OnAfterCodeMLSChannelTypeCreateCallback.OnAfterCreate(stage, codemlschanneltype)
	}
}

func (codemlschanneltype *CodeMLSChannelType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMLSChannelTypeUpdateCallback != nil {
		var frontCodeMLSChannelType *CodeMLSChannelType
		if front != nil {
			frontCodeMLSChannelType, _ = front.(*CodeMLSChannelType)
		}
		stage.OnAfterCodeMLSChannelTypeUpdateCallback.OnAfterUpdate(stage, codemlschanneltype, frontCodeMLSChannelType)
	}
}

func (codemlschanneltype *CodeMLSChannelType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMLSChannelTypeDeleteCallback != nil {
		var frontCodeMLSChannelType *CodeMLSChannelType
		if front != nil {
			frontCodeMLSChannelType, _ = front.(*CodeMLSChannelType)
		}
		stage.OnAfterCodeMLSChannelTypeDeleteCallback.OnAfterDelete(stage, codemlschanneltype, frontCodeMLSChannelType)
	}
}

func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeMarkerBeaconSignalTypeCreateCallback != nil {
		stage.OnAfterCodeMarkerBeaconSignalTypeCreateCallback.OnAfterCreate(stage, codemarkerbeaconsignaltype)
	}
}

func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMarkerBeaconSignalTypeUpdateCallback != nil {
		var frontCodeMarkerBeaconSignalType *CodeMarkerBeaconSignalType
		if front != nil {
			frontCodeMarkerBeaconSignalType, _ = front.(*CodeMarkerBeaconSignalType)
		}
		stage.OnAfterCodeMarkerBeaconSignalTypeUpdateCallback.OnAfterUpdate(stage, codemarkerbeaconsignaltype, frontCodeMarkerBeaconSignalType)
	}
}

func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMarkerBeaconSignalTypeDeleteCallback != nil {
		var frontCodeMarkerBeaconSignalType *CodeMarkerBeaconSignalType
		if front != nil {
			frontCodeMarkerBeaconSignalType, _ = front.(*CodeMarkerBeaconSignalType)
		}
		stage.OnAfterCodeMarkerBeaconSignalTypeDeleteCallback.OnAfterDelete(stage, codemarkerbeaconsignaltype, frontCodeMarkerBeaconSignalType)
	}
}

func (codendbusagetype *CodeNDBUsageType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeNDBUsageTypeCreateCallback != nil {
		stage.OnAfterCodeNDBUsageTypeCreateCallback.OnAfterCreate(stage, codendbusagetype)
	}
}

func (codendbusagetype *CodeNDBUsageType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNDBUsageTypeUpdateCallback != nil {
		var frontCodeNDBUsageType *CodeNDBUsageType
		if front != nil {
			frontCodeNDBUsageType, _ = front.(*CodeNDBUsageType)
		}
		stage.OnAfterCodeNDBUsageTypeUpdateCallback.OnAfterUpdate(stage, codendbusagetype, frontCodeNDBUsageType)
	}
}

func (codendbusagetype *CodeNDBUsageType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNDBUsageTypeDeleteCallback != nil {
		var frontCodeNDBUsageType *CodeNDBUsageType
		if front != nil {
			frontCodeNDBUsageType, _ = front.(*CodeNDBUsageType)
		}
		stage.OnAfterCodeNDBUsageTypeDeleteCallback.OnAfterDelete(stage, codendbusagetype, frontCodeNDBUsageType)
	}
}

func (codenavaiddesignatortype *CodeNavaidDesignatorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeNavaidDesignatorTypeCreateCallback != nil {
		stage.OnAfterCodeNavaidDesignatorTypeCreateCallback.OnAfterCreate(stage, codenavaiddesignatortype)
	}
}

func (codenavaiddesignatortype *CodeNavaidDesignatorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNavaidDesignatorTypeUpdateCallback != nil {
		var frontCodeNavaidDesignatorType *CodeNavaidDesignatorType
		if front != nil {
			frontCodeNavaidDesignatorType, _ = front.(*CodeNavaidDesignatorType)
		}
		stage.OnAfterCodeNavaidDesignatorTypeUpdateCallback.OnAfterUpdate(stage, codenavaiddesignatortype, frontCodeNavaidDesignatorType)
	}
}

func (codenavaiddesignatortype *CodeNavaidDesignatorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNavaidDesignatorTypeDeleteCallback != nil {
		var frontCodeNavaidDesignatorType *CodeNavaidDesignatorType
		if front != nil {
			frontCodeNavaidDesignatorType, _ = front.(*CodeNavaidDesignatorType)
		}
		stage.OnAfterCodeNavaidDesignatorTypeDeleteCallback.OnAfterDelete(stage, codenavaiddesignatortype, frontCodeNavaidDesignatorType)
	}
}

func (codenavaidpurposetype *CodeNavaidPurposeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeNavaidPurposeTypeCreateCallback != nil {
		stage.OnAfterCodeNavaidPurposeTypeCreateCallback.OnAfterCreate(stage, codenavaidpurposetype)
	}
}

func (codenavaidpurposetype *CodeNavaidPurposeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNavaidPurposeTypeUpdateCallback != nil {
		var frontCodeNavaidPurposeType *CodeNavaidPurposeType
		if front != nil {
			frontCodeNavaidPurposeType, _ = front.(*CodeNavaidPurposeType)
		}
		stage.OnAfterCodeNavaidPurposeTypeUpdateCallback.OnAfterUpdate(stage, codenavaidpurposetype, frontCodeNavaidPurposeType)
	}
}

func (codenavaidpurposetype *CodeNavaidPurposeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNavaidPurposeTypeDeleteCallback != nil {
		var frontCodeNavaidPurposeType *CodeNavaidPurposeType
		if front != nil {
			frontCodeNavaidPurposeType, _ = front.(*CodeNavaidPurposeType)
		}
		stage.OnAfterCodeNavaidPurposeTypeDeleteCallback.OnAfterDelete(stage, codenavaidpurposetype, frontCodeNavaidPurposeType)
	}
}

func (codenavaidservicetype *CodeNavaidServiceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeNavaidServiceTypeCreateCallback != nil {
		stage.OnAfterCodeNavaidServiceTypeCreateCallback.OnAfterCreate(stage, codenavaidservicetype)
	}
}

func (codenavaidservicetype *CodeNavaidServiceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNavaidServiceTypeUpdateCallback != nil {
		var frontCodeNavaidServiceType *CodeNavaidServiceType
		if front != nil {
			frontCodeNavaidServiceType, _ = front.(*CodeNavaidServiceType)
		}
		stage.OnAfterCodeNavaidServiceTypeUpdateCallback.OnAfterUpdate(stage, codenavaidservicetype, frontCodeNavaidServiceType)
	}
}

func (codenavaidservicetype *CodeNavaidServiceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNavaidServiceTypeDeleteCallback != nil {
		var frontCodeNavaidServiceType *CodeNavaidServiceType
		if front != nil {
			frontCodeNavaidServiceType, _ = front.(*CodeNavaidServiceType)
		}
		stage.OnAfterCodeNavaidServiceTypeDeleteCallback.OnAfterDelete(stage, codenavaidservicetype, frontCodeNavaidServiceType)
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

func (codepartype *CodePARType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodePARTypeCreateCallback != nil {
		stage.OnAfterCodePARTypeCreateCallback.OnAfterCreate(stage, codepartype)
	}
}

func (codepartype *CodePARType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodePARTypeUpdateCallback != nil {
		var frontCodePARType *CodePARType
		if front != nil {
			frontCodePARType, _ = front.(*CodePARType)
		}
		stage.OnAfterCodePARTypeUpdateCallback.OnAfterUpdate(stage, codepartype, frontCodePARType)
	}
}

func (codepartype *CodePARType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodePARTypeDeleteCallback != nil {
		var frontCodePARType *CodePARType
		if front != nil {
			frontCodePARType, _ = front.(*CodePARType)
		}
		stage.OnAfterCodePARTypeDeleteCallback.OnAfterDelete(stage, codepartype, frontCodePARType)
	}
}

func (codepositioninilstype *CodePositionInILSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodePositionInILSTypeCreateCallback != nil {
		stage.OnAfterCodePositionInILSTypeCreateCallback.OnAfterCreate(stage, codepositioninilstype)
	}
}

func (codepositioninilstype *CodePositionInILSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodePositionInILSTypeUpdateCallback != nil {
		var frontCodePositionInILSType *CodePositionInILSType
		if front != nil {
			frontCodePositionInILSType, _ = front.(*CodePositionInILSType)
		}
		stage.OnAfterCodePositionInILSTypeUpdateCallback.OnAfterUpdate(stage, codepositioninilstype, frontCodePositionInILSType)
	}
}

func (codepositioninilstype *CodePositionInILSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodePositionInILSTypeDeleteCallback != nil {
		var frontCodePositionInILSType *CodePositionInILSType
		if front != nil {
			frontCodePositionInILSType, _ = front.(*CodePositionInILSType)
		}
		stage.OnAfterCodePositionInILSTypeDeleteCallback.OnAfterDelete(stage, codepositioninilstype, frontCodePositionInILSType)
	}
}

func (codeprimaryradartype *CodePrimaryRadarType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodePrimaryRadarTypeCreateCallback != nil {
		stage.OnAfterCodePrimaryRadarTypeCreateCallback.OnAfterCreate(stage, codeprimaryradartype)
	}
}

func (codeprimaryradartype *CodePrimaryRadarType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodePrimaryRadarTypeUpdateCallback != nil {
		var frontCodePrimaryRadarType *CodePrimaryRadarType
		if front != nil {
			frontCodePrimaryRadarType, _ = front.(*CodePrimaryRadarType)
		}
		stage.OnAfterCodePrimaryRadarTypeUpdateCallback.OnAfterUpdate(stage, codeprimaryradartype, frontCodePrimaryRadarType)
	}
}

func (codeprimaryradartype *CodePrimaryRadarType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodePrimaryRadarTypeDeleteCallback != nil {
		var frontCodePrimaryRadarType *CodePrimaryRadarType
		if front != nil {
			frontCodePrimaryRadarType, _ = front.(*CodePrimaryRadarType)
		}
		stage.OnAfterCodePrimaryRadarTypeDeleteCallback.OnAfterDelete(stage, codeprimaryradartype, frontCodePrimaryRadarType)
	}
}

func (coderadarservicetype *CodeRadarServiceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRadarServiceTypeCreateCallback != nil {
		stage.OnAfterCodeRadarServiceTypeCreateCallback.OnAfterCreate(stage, coderadarservicetype)
	}
}

func (coderadarservicetype *CodeRadarServiceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRadarServiceTypeUpdateCallback != nil {
		var frontCodeRadarServiceType *CodeRadarServiceType
		if front != nil {
			frontCodeRadarServiceType, _ = front.(*CodeRadarServiceType)
		}
		stage.OnAfterCodeRadarServiceTypeUpdateCallback.OnAfterUpdate(stage, coderadarservicetype, frontCodeRadarServiceType)
	}
}

func (coderadarservicetype *CodeRadarServiceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRadarServiceTypeDeleteCallback != nil {
		var frontCodeRadarServiceType *CodeRadarServiceType
		if front != nil {
			frontCodeRadarServiceType, _ = front.(*CodeRadarServiceType)
		}
		stage.OnAfterCodeRadarServiceTypeDeleteCallback.OnAfterDelete(stage, coderadarservicetype, frontCodeRadarServiceType)
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

func (coderadiosignaltype *CodeRadioSignalType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRadioSignalTypeCreateCallback != nil {
		stage.OnAfterCodeRadioSignalTypeCreateCallback.OnAfterCreate(stage, coderadiosignaltype)
	}
}

func (coderadiosignaltype *CodeRadioSignalType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRadioSignalTypeUpdateCallback != nil {
		var frontCodeRadioSignalType *CodeRadioSignalType
		if front != nil {
			frontCodeRadioSignalType, _ = front.(*CodeRadioSignalType)
		}
		stage.OnAfterCodeRadioSignalTypeUpdateCallback.OnAfterUpdate(stage, coderadiosignaltype, frontCodeRadioSignalType)
	}
}

func (coderadiosignaltype *CodeRadioSignalType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRadioSignalTypeDeleteCallback != nil {
		var frontCodeRadioSignalType *CodeRadioSignalType
		if front != nil {
			frontCodeRadioSignalType, _ = front.(*CodeRadioSignalType)
		}
		stage.OnAfterCodeRadioSignalTypeDeleteCallback.OnAfterDelete(stage, coderadiosignaltype, frontCodeRadioSignalType)
	}
}

func (codereflectortype *CodeReflectorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeReflectorTypeCreateCallback != nil {
		stage.OnAfterCodeReflectorTypeCreateCallback.OnAfterCreate(stage, codereflectortype)
	}
}

func (codereflectortype *CodeReflectorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeReflectorTypeUpdateCallback != nil {
		var frontCodeReflectorType *CodeReflectorType
		if front != nil {
			frontCodeReflectorType, _ = front.(*CodeReflectorType)
		}
		stage.OnAfterCodeReflectorTypeUpdateCallback.OnAfterUpdate(stage, codereflectortype, frontCodeReflectorType)
	}
}

func (codereflectortype *CodeReflectorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeReflectorTypeDeleteCallback != nil {
		var frontCodeReflectorType *CodeReflectorType
		if front != nil {
			frontCodeReflectorType, _ = front.(*CodeReflectorType)
		}
		stage.OnAfterCodeReflectorTypeDeleteCallback.OnAfterDelete(stage, codereflectortype, frontCodeReflectorType)
	}
}

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeSignalPerformanceILSTypeCreateCallback != nil {
		stage.OnAfterCodeSignalPerformanceILSTypeCreateCallback.OnAfterCreate(stage, codesignalperformanceilstype)
	}
}

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSignalPerformanceILSTypeUpdateCallback != nil {
		var frontCodeSignalPerformanceILSType *CodeSignalPerformanceILSType
		if front != nil {
			frontCodeSignalPerformanceILSType, _ = front.(*CodeSignalPerformanceILSType)
		}
		stage.OnAfterCodeSignalPerformanceILSTypeUpdateCallback.OnAfterUpdate(stage, codesignalperformanceilstype, frontCodeSignalPerformanceILSType)
	}
}

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSignalPerformanceILSTypeDeleteCallback != nil {
		var frontCodeSignalPerformanceILSType *CodeSignalPerformanceILSType
		if front != nil {
			frontCodeSignalPerformanceILSType, _ = front.(*CodeSignalPerformanceILSType)
		}
		stage.OnAfterCodeSignalPerformanceILSTypeDeleteCallback.OnAfterDelete(stage, codesignalperformanceilstype, frontCodeSignalPerformanceILSType)
	}
}

func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeSpecialNavigationChainDesignatorTypeCreateCallback != nil {
		stage.OnAfterCodeSpecialNavigationChainDesignatorTypeCreateCallback.OnAfterCreate(stage, codespecialnavigationchaindesignatortype)
	}
}

func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSpecialNavigationChainDesignatorTypeUpdateCallback != nil {
		var frontCodeSpecialNavigationChainDesignatorType *CodeSpecialNavigationChainDesignatorType
		if front != nil {
			frontCodeSpecialNavigationChainDesignatorType, _ = front.(*CodeSpecialNavigationChainDesignatorType)
		}
		stage.OnAfterCodeSpecialNavigationChainDesignatorTypeUpdateCallback.OnAfterUpdate(stage, codespecialnavigationchaindesignatortype, frontCodeSpecialNavigationChainDesignatorType)
	}
}

func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSpecialNavigationChainDesignatorTypeDeleteCallback != nil {
		var frontCodeSpecialNavigationChainDesignatorType *CodeSpecialNavigationChainDesignatorType
		if front != nil {
			frontCodeSpecialNavigationChainDesignatorType, _ = front.(*CodeSpecialNavigationChainDesignatorType)
		}
		stage.OnAfterCodeSpecialNavigationChainDesignatorTypeDeleteCallback.OnAfterDelete(stage, codespecialnavigationchaindesignatortype, frontCodeSpecialNavigationChainDesignatorType)
	}
}

func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeSpecialNavigationStationTypeCreateCallback != nil {
		stage.OnAfterCodeSpecialNavigationStationTypeCreateCallback.OnAfterCreate(stage, codespecialnavigationstationtype)
	}
}

func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSpecialNavigationStationTypeUpdateCallback != nil {
		var frontCodeSpecialNavigationStationType *CodeSpecialNavigationStationType
		if front != nil {
			frontCodeSpecialNavigationStationType, _ = front.(*CodeSpecialNavigationStationType)
		}
		stage.OnAfterCodeSpecialNavigationStationTypeUpdateCallback.OnAfterUpdate(stage, codespecialnavigationstationtype, frontCodeSpecialNavigationStationType)
	}
}

func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSpecialNavigationStationTypeDeleteCallback != nil {
		var frontCodeSpecialNavigationStationType *CodeSpecialNavigationStationType
		if front != nil {
			frontCodeSpecialNavigationStationType, _ = front.(*CodeSpecialNavigationStationType)
		}
		stage.OnAfterCodeSpecialNavigationStationTypeDeleteCallback.OnAfterDelete(stage, codespecialnavigationstationtype, frontCodeSpecialNavigationStationType)
	}
}

func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeSpecialNavigationSystemTypeCreateCallback != nil {
		stage.OnAfterCodeSpecialNavigationSystemTypeCreateCallback.OnAfterCreate(stage, codespecialnavigationsystemtype)
	}
}

func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSpecialNavigationSystemTypeUpdateCallback != nil {
		var frontCodeSpecialNavigationSystemType *CodeSpecialNavigationSystemType
		if front != nil {
			frontCodeSpecialNavigationSystemType, _ = front.(*CodeSpecialNavigationSystemType)
		}
		stage.OnAfterCodeSpecialNavigationSystemTypeUpdateCallback.OnAfterUpdate(stage, codespecialnavigationsystemtype, frontCodeSpecialNavigationSystemType)
	}
}

func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSpecialNavigationSystemTypeDeleteCallback != nil {
		var frontCodeSpecialNavigationSystemType *CodeSpecialNavigationSystemType
		if front != nil {
			frontCodeSpecialNavigationSystemType, _ = front.(*CodeSpecialNavigationSystemType)
		}
		stage.OnAfterCodeSpecialNavigationSystemTypeDeleteCallback.OnAfterDelete(stage, codespecialnavigationsystemtype, frontCodeSpecialNavigationSystemType)
	}
}

func (codestandbypowertype *CodeStandbyPowerType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeStandbyPowerTypeCreateCallback != nil {
		stage.OnAfterCodeStandbyPowerTypeCreateCallback.OnAfterCreate(stage, codestandbypowertype)
	}
}

func (codestandbypowertype *CodeStandbyPowerType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeStandbyPowerTypeUpdateCallback != nil {
		var frontCodeStandbyPowerType *CodeStandbyPowerType
		if front != nil {
			frontCodeStandbyPowerType, _ = front.(*CodeStandbyPowerType)
		}
		stage.OnAfterCodeStandbyPowerTypeUpdateCallback.OnAfterUpdate(stage, codestandbypowertype, frontCodeStandbyPowerType)
	}
}

func (codestandbypowertype *CodeStandbyPowerType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeStandbyPowerTypeDeleteCallback != nil {
		var frontCodeStandbyPowerType *CodeStandbyPowerType
		if front != nil {
			frontCodeStandbyPowerType, _ = front.(*CodeStandbyPowerType)
		}
		stage.OnAfterCodeStandbyPowerTypeDeleteCallback.OnAfterDelete(stage, codestandbypowertype, frontCodeStandbyPowerType)
	}
}

func (codestatusnavaidtype *CodeStatusNavaidType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeStatusNavaidTypeCreateCallback != nil {
		stage.OnAfterCodeStatusNavaidTypeCreateCallback.OnAfterCreate(stage, codestatusnavaidtype)
	}
}

func (codestatusnavaidtype *CodeStatusNavaidType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeStatusNavaidTypeUpdateCallback != nil {
		var frontCodeStatusNavaidType *CodeStatusNavaidType
		if front != nil {
			frontCodeStatusNavaidType, _ = front.(*CodeStatusNavaidType)
		}
		stage.OnAfterCodeStatusNavaidTypeUpdateCallback.OnAfterUpdate(stage, codestatusnavaidtype, frontCodeStatusNavaidType)
	}
}

func (codestatusnavaidtype *CodeStatusNavaidType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeStatusNavaidTypeDeleteCallback != nil {
		var frontCodeStatusNavaidType *CodeStatusNavaidType
		if front != nil {
			frontCodeStatusNavaidType, _ = front.(*CodeStatusNavaidType)
		}
		stage.OnAfterCodeStatusNavaidTypeDeleteCallback.OnAfterDelete(stage, codestatusnavaidtype, frontCodeStatusNavaidType)
	}
}

func (codetacanchanneltype *CodeTACANChannelType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeTACANChannelTypeCreateCallback != nil {
		stage.OnAfterCodeTACANChannelTypeCreateCallback.OnAfterCreate(stage, codetacanchanneltype)
	}
}

func (codetacanchanneltype *CodeTACANChannelType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTACANChannelTypeUpdateCallback != nil {
		var frontCodeTACANChannelType *CodeTACANChannelType
		if front != nil {
			frontCodeTACANChannelType, _ = front.(*CodeTACANChannelType)
		}
		stage.OnAfterCodeTACANChannelTypeUpdateCallback.OnAfterUpdate(stage, codetacanchanneltype, frontCodeTACANChannelType)
	}
}

func (codetacanchanneltype *CodeTACANChannelType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTACANChannelTypeDeleteCallback != nil {
		var frontCodeTACANChannelType *CodeTACANChannelType
		if front != nil {
			frontCodeTACANChannelType, _ = front.(*CodeTACANChannelType)
		}
		stage.OnAfterCodeTACANChannelTypeDeleteCallback.OnAfterDelete(stage, codetacanchanneltype, frontCodeTACANChannelType)
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

func (codetranspondertype *CodeTransponderType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeTransponderTypeCreateCallback != nil {
		stage.OnAfterCodeTransponderTypeCreateCallback.OnAfterCreate(stage, codetranspondertype)
	}
}

func (codetranspondertype *CodeTransponderType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTransponderTypeUpdateCallback != nil {
		var frontCodeTransponderType *CodeTransponderType
		if front != nil {
			frontCodeTransponderType, _ = front.(*CodeTransponderType)
		}
		stage.OnAfterCodeTransponderTypeUpdateCallback.OnAfterUpdate(stage, codetranspondertype, frontCodeTransponderType)
	}
}

func (codetranspondertype *CodeTransponderType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTransponderTypeDeleteCallback != nil {
		var frontCodeTransponderType *CodeTransponderType
		if front != nil {
			frontCodeTransponderType, _ = front.(*CodeTransponderType)
		}
		stage.OnAfterCodeTransponderTypeDeleteCallback.OnAfterDelete(stage, codetranspondertype, frontCodeTransponderType)
	}
}

func (codevortype *CodeVORType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeVORTypeCreateCallback != nil {
		stage.OnAfterCodeVORTypeCreateCallback.OnAfterCreate(stage, codevortype)
	}
}

func (codevortype *CodeVORType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeVORTypeUpdateCallback != nil {
		var frontCodeVORType *CodeVORType
		if front != nil {
			frontCodeVORType, _ = front.(*CodeVORType)
		}
		stage.OnAfterCodeVORTypeUpdateCallback.OnAfterUpdate(stage, codevortype, frontCodeVORType)
	}
}

func (codevortype *CodeVORType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeVORTypeDeleteCallback != nil {
		var frontCodeVORType *CodeVORType
		if front != nil {
			frontCodeVORType, _ = front.(*CodeVORType)
		}
		stage.OnAfterCodeVORTypeDeleteCallback.OnAfterDelete(stage, codevortype, frontCodeVORType)
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

func (codeverticalreferencetype *CodeVerticalReferenceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeVerticalReferenceTypeCreateCallback != nil {
		stage.OnAfterCodeVerticalReferenceTypeCreateCallback.OnAfterCreate(stage, codeverticalreferencetype)
	}
}

func (codeverticalreferencetype *CodeVerticalReferenceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeVerticalReferenceTypeUpdateCallback != nil {
		var frontCodeVerticalReferenceType *CodeVerticalReferenceType
		if front != nil {
			frontCodeVerticalReferenceType, _ = front.(*CodeVerticalReferenceType)
		}
		stage.OnAfterCodeVerticalReferenceTypeUpdateCallback.OnAfterUpdate(stage, codeverticalreferencetype, frontCodeVerticalReferenceType)
	}
}

func (codeverticalreferencetype *CodeVerticalReferenceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeVerticalReferenceTypeDeleteCallback != nil {
		var frontCodeVerticalReferenceType *CodeVerticalReferenceType
		if front != nil {
			frontCodeVerticalReferenceType, _ = front.(*CodeVerticalReferenceType)
		}
		stage.OnAfterCodeVerticalReferenceTypeDeleteCallback.OnAfterDelete(stage, codeverticalreferencetype, frontCodeVerticalReferenceType)
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

func (nosequencetype *NoSequenceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoSequenceTypeCreateCallback != nil {
		stage.OnAfterNoSequenceTypeCreateCallback.OnAfterCreate(stage, nosequencetype)
	}
}

func (nosequencetype *NoSequenceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoSequenceTypeUpdateCallback != nil {
		var frontNoSequenceType *NoSequenceType
		if front != nil {
			frontNoSequenceType, _ = front.(*NoSequenceType)
		}
		stage.OnAfterNoSequenceTypeUpdateCallback.OnAfterUpdate(stage, nosequencetype, frontNoSequenceType)
	}
}

func (nosequencetype *NoSequenceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoSequenceTypeDeleteCallback != nil {
		var frontNoSequenceType *NoSequenceType
		if front != nil {
			frontNoSequenceType, _ = front.(*NoSequenceType)
		}
		stage.OnAfterNoSequenceTypeDeleteCallback.OnAfterDelete(stage, nosequencetype, frontNoSequenceType)
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

func (valangletype *ValAngleType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValAngleTypeCreateCallback != nil {
		stage.OnAfterValAngleTypeCreateCallback.OnAfterCreate(stage, valangletype)
	}
}

func (valangletype *ValAngleType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValAngleTypeUpdateCallback != nil {
		var frontValAngleType *ValAngleType
		if front != nil {
			frontValAngleType, _ = front.(*ValAngleType)
		}
		stage.OnAfterValAngleTypeUpdateCallback.OnAfterUpdate(stage, valangletype, frontValAngleType)
	}
}

func (valangletype *ValAngleType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValAngleTypeDeleteCallback != nil {
		var frontValAngleType *ValAngleType
		if front != nil {
			frontValAngleType, _ = front.(*ValAngleType)
		}
		stage.OnAfterValAngleTypeDeleteCallback.OnAfterDelete(stage, valangletype, frontValAngleType)
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

func (valmagneticvariationtype *ValMagneticVariationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValMagneticVariationTypeCreateCallback != nil {
		stage.OnAfterValMagneticVariationTypeCreateCallback.OnAfterCreate(stage, valmagneticvariationtype)
	}
}

func (valmagneticvariationtype *ValMagneticVariationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValMagneticVariationTypeUpdateCallback != nil {
		var frontValMagneticVariationType *ValMagneticVariationType
		if front != nil {
			frontValMagneticVariationType, _ = front.(*ValMagneticVariationType)
		}
		stage.OnAfterValMagneticVariationTypeUpdateCallback.OnAfterUpdate(stage, valmagneticvariationtype, frontValMagneticVariationType)
	}
}

func (valmagneticvariationtype *ValMagneticVariationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValMagneticVariationTypeDeleteCallback != nil {
		var frontValMagneticVariationType *ValMagneticVariationType
		if front != nil {
			frontValMagneticVariationType, _ = front.(*ValMagneticVariationType)
		}
		stage.OnAfterValMagneticVariationTypeDeleteCallback.OnAfterDelete(stage, valmagneticvariationtype, frontValMagneticVariationType)
	}
}

