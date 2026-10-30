// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	gongprobe "github.com/fullstack-lang/gong/pkg/runtime/probe"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/xlink"
)

type FormCallbackIF = gongprobe.FormCallbackIF

type functionalStageSetFormCallback struct {
	onSave func()
}

func (cb *functionalStageSetFormCallback) OnSave() {
	if cb.onSave != nil {
		cb.onSave()
	}
}

func saveStageSet_AIXMBasicMessage_Stage(
	inst *models.AIXMBasicMessage,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "MessageMetadata":
			StageSetFormDivSelectFieldToField(&inst.MessageMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.MessageMetadataPropertyType](), formDiv)
		case "AggregationType":
			FormDivBasicFieldToField(&inst.AggregationType, formDiv)
		case "HasMember":
			FormDivSliceOfPointersToField(inst, "HasMember", &inst.HasMember, formDiv, probe)
		}
	}
}

func saveStageSet_AIXMBasicMessageType_Stage(
	inst *models.AIXMBasicMessageType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "MessageMetadata":
			StageSetFormDivSelectFieldToField(&inst.MessageMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.MessageMetadataPropertyType](), formDiv)
		case "AggregationType":
			FormDivBasicFieldToField(&inst.AggregationType, formDiv)
		case "HasMember":
			FormDivSliceOfPointersToField(inst, "HasMember", &inst.HasMember, formDiv, probe)
		}
	}
}

func saveStageSet_BasicMessageMemberAIXMPropertyType_Stage(
	inst *models.BasicMessageMemberAIXMPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "Href":
			FormDivBasicFieldToField(&inst.Href, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Arcrole":
			FormDivBasicFieldToField(&inst.Arcrole, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Show":
			FormDivBasicFieldToField(&inst.Show, formDiv)
		case "Actuate":
			FormDivBasicFieldToField(&inst.Actuate, formDiv)
		case "Navaid":
			StageSetFormDivSelectFieldToField(&inst.Navaid, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidType](), formDiv)
		case "VOR":
			StageSetFormDivSelectFieldToField(&inst.VOR, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.VORType](), formDiv)
		case "DME":
			StageSetFormDivSelectFieldToField(&inst.DME, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DMEType](), formDiv)
		case "TACAN":
			StageSetFormDivSelectFieldToField(&inst.TACAN, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TACANType](), formDiv)
		case "NDB":
			StageSetFormDivSelectFieldToField(&inst.NDB, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NDBType](), formDiv)
		case "Localizer":
			StageSetFormDivSelectFieldToField(&inst.Localizer, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.LocalizerType](), formDiv)
		case "Glidepath":
			StageSetFormDivSelectFieldToField(&inst.Glidepath, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.GlidepathType](), formDiv)
		case "MarkerBeacon":
			StageSetFormDivSelectFieldToField(&inst.MarkerBeacon, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.MarkerBeaconType](), formDiv)
		case "SDF":
			StageSetFormDivSelectFieldToField(&inst.SDF, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SDFType](), formDiv)
		case "DirectionFinder":
			StageSetFormDivSelectFieldToField(&inst.DirectionFinder, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DirectionFinderType](), formDiv)
		case "Azimuth":
			StageSetFormDivSelectFieldToField(&inst.Azimuth, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AzimuthType](), formDiv)
		case "Elevation":
			StageSetFormDivSelectFieldToField(&inst.Elevation, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevationType](), formDiv)
		case "CheckpointVOR":
			StageSetFormDivSelectFieldToField(&inst.CheckpointVOR, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CheckpointVORType](), formDiv)
		case "CheckpointINS":
			StageSetFormDivSelectFieldToField(&inst.CheckpointINS, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CheckpointINSType](), formDiv)
		case "SpecialNavigationSystem":
			StageSetFormDivSelectFieldToField(&inst.SpecialNavigationSystem, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SpecialNavigationSystemType](), formDiv)
		case "SpecialNavigationStation":
			StageSetFormDivSelectFieldToField(&inst.SpecialNavigationStation, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SpecialNavigationStationType](), formDiv)
		case "DesignatedPoint":
			StageSetFormDivSelectFieldToField(&inst.DesignatedPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointType](), formDiv)
		case "SignificantPointInAirspace":
			StageSetFormDivSelectFieldToField(&inst.SignificantPointInAirspace, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SignificantPointInAirspaceType](), formDiv)
		case "AngleIndication":
			StageSetFormDivSelectFieldToField(&inst.AngleIndication, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AngleIndicationType](), formDiv)
		case "DistanceIndication":
			StageSetFormDivSelectFieldToField(&inst.DistanceIndication, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DistanceIndicationType](), formDiv)
		case "RadarSystem":
			StageSetFormDivSelectFieldToField(&inst.RadarSystem, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RadarSystemType](), formDiv)
		case "PrimarySurveillanceRadar":
			StageSetFormDivSelectFieldToField(&inst.PrimarySurveillanceRadar, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PrimarySurveillanceRadarType](), formDiv)
		case "SecondarySurveillanceRadar":
			StageSetFormDivSelectFieldToField(&inst.SecondarySurveillanceRadar, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SecondarySurveillanceRadarType](), formDiv)
		case "PrecisionApproachRadar":
			StageSetFormDivSelectFieldToField(&inst.PrecisionApproachRadar, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PrecisionApproachRadarType](), formDiv)
		}
	}
}

func saveStageSet_A_extension_AixmfeaturesStage(
	inst *aixmfeatures.A_extension,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_1_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_1,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_10_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_10,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_11_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_11,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_12_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_12,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_13_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_13,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_14_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_14,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_15_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_15,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_16_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_16,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_17_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_17,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_18_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_18,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_19_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_19,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_2_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_2,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_20_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_20,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_21_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_21,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_22_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_22,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_23_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_23,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_24_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_24,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_25_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_25,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_26_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_26,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_27_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_27,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_28_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_28,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_29_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_29,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_3_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_3,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_30_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_30,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_31_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_31,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_32_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_32,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_33_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_33,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_34_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_34,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_35_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_35,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_36_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_36,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_37_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_37,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_38_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_38,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_39_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_39,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_4_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_4,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_40_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_40,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_41_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_41,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_5_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_5,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_6_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_6,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_7_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_7,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_8_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_8,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_A_extension_9_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_9,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_AbstractNavaidEquipmentType_AixmfeaturesStage(
	inst *aixmfeatures.AbstractNavaidEquipmentType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_AbstractNavigationSystemCheckpointType_AixmfeaturesStage(
	inst *aixmfeatures.AbstractNavigationSystemCheckpointType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_AbstractPropertiesWithScheduleType_AixmfeaturesStage(
	inst *aixmfeatures.AbstractPropertiesWithScheduleType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		}
	}
}

func saveStageSet_AbstractRadarEquipmentType_AixmfeaturesStage(
	inst *aixmfeatures.AbstractRadarEquipmentType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_AbstractSurveillanceRadarType_AixmfeaturesStage(
	inst *aixmfeatures.AbstractSurveillanceRadarType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_AirportHeliportPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AirportHeliportPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "Href":
			FormDivBasicFieldToField(&inst.Href, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Arcrole":
			FormDivBasicFieldToField(&inst.Arcrole, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Show":
			FormDivBasicFieldToField(&inst.Show, formDiv)
		case "Actuate":
			FormDivBasicFieldToField(&inst.Actuate, formDiv)
		}
	}
}

func saveStageSet_AirspacePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AirspacePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "Href":
			FormDivBasicFieldToField(&inst.Href, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Arcrole":
			FormDivBasicFieldToField(&inst.Arcrole, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Show":
			FormDivBasicFieldToField(&inst.Show, formDiv)
		case "Actuate":
			FormDivBasicFieldToField(&inst.Actuate, formDiv)
		}
	}
}

func saveStageSet_AngleIndicationTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AngleIndicationTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "AngleIndicationTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.AngleIndicationTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AngleIndicationTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_AngleIndicationTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.AngleIndicationTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "ValidTime":
			StageSetFormDivSelectFieldToField(&inst.ValidTime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Interpretation":
			FormDivBasicFieldToField(&inst.Interpretation, formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "CorrectionNumber":
			FormDivBasicFieldToField(&inst.CorrectionNumber, formDiv)
		case "TimeSliceMetadata":
			StageSetFormDivSelectFieldToField(&inst.TimeSliceMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), formDiv)
		case "FeatureLifetime":
			StageSetFormDivSelectFieldToField(&inst.FeatureLifetime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Angle":
			StageSetFormDivSelectFieldToField(&inst.Angle, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "AngleType":
			StageSetFormDivSelectFieldToField(&inst.AngleType, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeBearingType](), formDiv)
		case "IndicationDirection":
			StageSetFormDivSelectFieldToField(&inst.IndicationDirection, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDirectionReferenceType](), formDiv)
		case "TrueAngle":
			StageSetFormDivSelectFieldToField(&inst.TrueAngle, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "CardinalDirection":
			StageSetFormDivSelectFieldToField(&inst.CardinalDirection, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCardinalDirectionType](), formDiv)
		case "MinimumReceptionAltitude":
			StageSetFormDivSelectFieldToField(&inst.MinimumReceptionAltitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "Fix":
			StageSetFormDivSelectFieldToField(&inst.Fix, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), formDiv)
		case "PointChoice_fixDesignatedPoint":
			StageSetFormDivSelectFieldToField(&inst.PointChoice_fixDesignatedPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), formDiv)
		case "PointChoice_navaidSystem":
			StageSetFormDivSelectFieldToField(&inst.PointChoice_navaidSystem, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), formDiv)
		case "PointChoice_position":
			StageSetFormDivSelectFieldToField(&inst.PointChoice_position, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), formDiv)
		case "PointChoice_runwayPoint":
			StageSetFormDivSelectFieldToField(&inst.PointChoice_runwayPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), formDiv)
		case "PointChoice_aimingPoint":
			StageSetFormDivSelectFieldToField(&inst.PointChoice_aimingPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), formDiv)
		case "PointChoice_airportReferencePoint":
			StageSetFormDivSelectFieldToField(&inst.PointChoice_airportReferencePoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), formDiv)
		}
	}
}

func saveStageSet_AngleIndicationType_AixmfeaturesStage(
	inst *aixmfeatures.AngleIndicationType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_AuthorityForNavaidEquipmentPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AuthorityForNavaidEquipmentPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "AuthorityForNavaidEquipment":
			StageSetFormDivSelectFieldToField(&inst.AuthorityForNavaidEquipment, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AuthorityForNavaidEquipmentType](), formDiv)
		}
	}
}

func saveStageSet_AuthorityForNavaidEquipmentType_AixmfeaturesStage(
	inst *aixmfeatures.AuthorityForNavaidEquipmentType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "Type_XSD":
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAuthorityRoleType](), formDiv)
		case "TheOrganisationAuthority":
			StageSetFormDivSelectFieldToField(&inst.TheOrganisationAuthority, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.OrganisationAuthorityPropertyType](), formDiv)
		}
	}
}

func saveStageSet_AuthorityForSpecialNavigationStationPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AuthorityForSpecialNavigationStationPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "AuthorityForSpecialNavigationStation":
			StageSetFormDivSelectFieldToField(&inst.AuthorityForSpecialNavigationStation, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AuthorityForSpecialNavigationStationType](), formDiv)
		}
	}
}

func saveStageSet_AuthorityForSpecialNavigationStationType_AixmfeaturesStage(
	inst *aixmfeatures.AuthorityForSpecialNavigationStationType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "Type_XSD":
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAuthorityRoleType](), formDiv)
		case "TheOrganisationAuthority":
			StageSetFormDivSelectFieldToField(&inst.TheOrganisationAuthority, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.OrganisationAuthorityPropertyType](), formDiv)
		}
	}
}

func saveStageSet_AuthorityForSpecialNavigationSystemPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AuthorityForSpecialNavigationSystemPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "AuthorityForSpecialNavigationSystem":
			StageSetFormDivSelectFieldToField(&inst.AuthorityForSpecialNavigationSystem, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AuthorityForSpecialNavigationSystemType](), formDiv)
		}
	}
}

func saveStageSet_AuthorityForSpecialNavigationSystemType_AixmfeaturesStage(
	inst *aixmfeatures.AuthorityForSpecialNavigationSystemType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "Type_XSD":
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAuthorityRoleType](), formDiv)
		case "TheOrganisationAuthority":
			StageSetFormDivSelectFieldToField(&inst.TheOrganisationAuthority, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.OrganisationAuthorityPropertyType](), formDiv)
		}
	}
}

func saveStageSet_AzimuthTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AzimuthTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "AzimuthTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.AzimuthTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AzimuthTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_AzimuthTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.AzimuthTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "ValidTime":
			StageSetFormDivSelectFieldToField(&inst.ValidTime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Interpretation":
			FormDivBasicFieldToField(&inst.Interpretation, formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "CorrectionNumber":
			FormDivBasicFieldToField(&inst.CorrectionNumber, formDiv)
		case "TimeSliceMetadata":
			StageSetFormDivSelectFieldToField(&inst.TimeSliceMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), formDiv)
		case "FeatureLifetime":
			StageSetFormDivSelectFieldToField(&inst.FeatureLifetime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Designator":
			StageSetFormDivSelectFieldToField(&inst.Designator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavaidDesignatorType](), formDiv)
		case "EmissionClass":
			StageSetFormDivSelectFieldToField(&inst.EmissionClass, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadioEmissionType](), formDiv)
		case "Mobile":
			StageSetFormDivSelectFieldToField(&inst.Mobile, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "MagneticVariation":
			StageSetFormDivSelectFieldToField(&inst.MagneticVariation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), formDiv)
		case "MagneticVariationAccuracy":
			StageSetFormDivSelectFieldToField(&inst.MagneticVariationAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "DateMagneticVariation":
			StageSetFormDivSelectFieldToField(&inst.DateMagneticVariation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), formDiv)
		case "FlightChecked":
			StageSetFormDivSelectFieldToField(&inst.FlightChecked, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "Location":
			StageSetFormDivSelectFieldToField(&inst.Location, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), formDiv)
		case "Type_XSD":
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeMLSAzimuthType](), formDiv)
		case "TrueBearing":
			StageSetFormDivSelectFieldToField(&inst.TrueBearing, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "TrueBearingAccuracy":
			StageSetFormDivSelectFieldToField(&inst.TrueBearingAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "MagneticBearing":
			StageSetFormDivSelectFieldToField(&inst.MagneticBearing, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "AngleProportionalLeft":
			StageSetFormDivSelectFieldToField(&inst.AngleProportionalLeft, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "AngleProportionalRight":
			StageSetFormDivSelectFieldToField(&inst.AngleProportionalRight, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "AngleCoverLeft":
			StageSetFormDivSelectFieldToField(&inst.AngleCoverLeft, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "AngleCoverRight":
			StageSetFormDivSelectFieldToField(&inst.AngleCoverRight, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "Channel":
			StageSetFormDivSelectFieldToField(&inst.Channel, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeMLSChannelType](), formDiv)
		}
	}
}

func saveStageSet_AzimuthType_AixmfeaturesStage(
	inst *aixmfeatures.AzimuthType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_CheckpointINSTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.CheckpointINSTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "CheckpointINSTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.CheckpointINSTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CheckpointINSTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_CheckpointINSTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.CheckpointINSTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "ValidTime":
			StageSetFormDivSelectFieldToField(&inst.ValidTime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Interpretation":
			FormDivBasicFieldToField(&inst.Interpretation, formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "CorrectionNumber":
			FormDivBasicFieldToField(&inst.CorrectionNumber, formDiv)
		case "TimeSliceMetadata":
			StageSetFormDivSelectFieldToField(&inst.TimeSliceMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), formDiv)
		case "FeatureLifetime":
			StageSetFormDivSelectFieldToField(&inst.FeatureLifetime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Category":
			StageSetFormDivSelectFieldToField(&inst.Category, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCheckpointCategoryType](), formDiv)
		case "UpperLimit":
			StageSetFormDivSelectFieldToField(&inst.UpperLimit, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "UpperLimitReference":
			StageSetFormDivSelectFieldToField(&inst.UpperLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "LowerLimit":
			StageSetFormDivSelectFieldToField(&inst.LowerLimit, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "LowerLimitReference":
			StageSetFormDivSelectFieldToField(&inst.LowerLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "AltitudeInterpretation":
			StageSetFormDivSelectFieldToField(&inst.AltitudeInterpretation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAltitudeUseType](), formDiv)
		case "Distance":
			StageSetFormDivSelectFieldToField(&inst.Distance, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "Angle":
			StageSetFormDivSelectFieldToField(&inst.Angle, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "Position":
			StageSetFormDivSelectFieldToField(&inst.Position, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), formDiv)
		case "AirportHeliport":
			StageSetFormDivSelectFieldToField(&inst.AirportHeliport, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), formDiv)
		}
	}
}

func saveStageSet_CheckpointINSType_AixmfeaturesStage(
	inst *aixmfeatures.CheckpointINSType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_CheckpointVORTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.CheckpointVORTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "CheckpointVORTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.CheckpointVORTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CheckpointVORTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_CheckpointVORTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.CheckpointVORTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "ValidTime":
			StageSetFormDivSelectFieldToField(&inst.ValidTime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Interpretation":
			FormDivBasicFieldToField(&inst.Interpretation, formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "CorrectionNumber":
			FormDivBasicFieldToField(&inst.CorrectionNumber, formDiv)
		case "TimeSliceMetadata":
			StageSetFormDivSelectFieldToField(&inst.TimeSliceMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), formDiv)
		case "FeatureLifetime":
			StageSetFormDivSelectFieldToField(&inst.FeatureLifetime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Category":
			StageSetFormDivSelectFieldToField(&inst.Category, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCheckpointCategoryType](), formDiv)
		case "UpperLimit":
			StageSetFormDivSelectFieldToField(&inst.UpperLimit, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "UpperLimitReference":
			StageSetFormDivSelectFieldToField(&inst.UpperLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "LowerLimit":
			StageSetFormDivSelectFieldToField(&inst.LowerLimit, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "LowerLimitReference":
			StageSetFormDivSelectFieldToField(&inst.LowerLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "AltitudeInterpretation":
			StageSetFormDivSelectFieldToField(&inst.AltitudeInterpretation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAltitudeUseType](), formDiv)
		case "Distance":
			StageSetFormDivSelectFieldToField(&inst.Distance, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "Angle":
			StageSetFormDivSelectFieldToField(&inst.Angle, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "Position":
			StageSetFormDivSelectFieldToField(&inst.Position, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), formDiv)
		case "AirportHeliport":
			StageSetFormDivSelectFieldToField(&inst.AirportHeliport, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), formDiv)
		case "CheckPointFacility":
			StageSetFormDivSelectFieldToField(&inst.CheckPointFacility, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.VORPropertyType](), formDiv)
		}
	}
}

func saveStageSet_CheckpointVORType_AixmfeaturesStage(
	inst *aixmfeatures.CheckpointVORType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_ContactInformationPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.ContactInformationPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "ContactInformation":
			StageSetFormDivSelectFieldToField(&inst.ContactInformation, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ContactInformationType](), formDiv)
		}
	}
}

func saveStageSet_ContactInformationType_AixmfeaturesStage(
	inst *aixmfeatures.ContactInformationType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "Title":
			StageSetFormDivSelectFieldToField(&inst.Title, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextNameType](), formDiv)
		}
	}
}

func saveStageSet_DMETimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.DMETimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "DMETimeSlice":
			StageSetFormDivSelectFieldToField(&inst.DMETimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DMETimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_DMETimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.DMETimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "ValidTime":
			StageSetFormDivSelectFieldToField(&inst.ValidTime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Interpretation":
			FormDivBasicFieldToField(&inst.Interpretation, formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "CorrectionNumber":
			FormDivBasicFieldToField(&inst.CorrectionNumber, formDiv)
		case "TimeSliceMetadata":
			StageSetFormDivSelectFieldToField(&inst.TimeSliceMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), formDiv)
		case "FeatureLifetime":
			StageSetFormDivSelectFieldToField(&inst.FeatureLifetime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Designator":
			StageSetFormDivSelectFieldToField(&inst.Designator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavaidDesignatorType](), formDiv)
		case "EmissionClass":
			StageSetFormDivSelectFieldToField(&inst.EmissionClass, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadioEmissionType](), formDiv)
		case "Mobile":
			StageSetFormDivSelectFieldToField(&inst.Mobile, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "MagneticVariation":
			StageSetFormDivSelectFieldToField(&inst.MagneticVariation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), formDiv)
		case "MagneticVariationAccuracy":
			StageSetFormDivSelectFieldToField(&inst.MagneticVariationAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "DateMagneticVariation":
			StageSetFormDivSelectFieldToField(&inst.DateMagneticVariation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), formDiv)
		case "FlightChecked":
			StageSetFormDivSelectFieldToField(&inst.FlightChecked, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "Location":
			StageSetFormDivSelectFieldToField(&inst.Location, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), formDiv)
		case "Type_XSD":
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDMEType](), formDiv)
		case "Channel":
			StageSetFormDivSelectFieldToField(&inst.Channel, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDMEChannelType](), formDiv)
		case "GhostFrequency":
			StageSetFormDivSelectFieldToField(&inst.GhostFrequency, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValFrequencyType](), formDiv)
		case "Displace":
			StageSetFormDivSelectFieldToField(&inst.Displace, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		}
	}
}

func saveStageSet_DMEType_AixmfeaturesStage(
	inst *aixmfeatures.DMEType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_DesignatedPointPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.DesignatedPointPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "Href":
			FormDivBasicFieldToField(&inst.Href, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Arcrole":
			FormDivBasicFieldToField(&inst.Arcrole, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Show":
			FormDivBasicFieldToField(&inst.Show, formDiv)
		case "Actuate":
			FormDivBasicFieldToField(&inst.Actuate, formDiv)
		}
	}
}

func saveStageSet_DesignatedPointTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.DesignatedPointTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "DesignatedPointTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.DesignatedPointTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_DesignatedPointTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.DesignatedPointTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "ValidTime":
			StageSetFormDivSelectFieldToField(&inst.ValidTime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Interpretation":
			FormDivBasicFieldToField(&inst.Interpretation, formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "CorrectionNumber":
			FormDivBasicFieldToField(&inst.CorrectionNumber, formDiv)
		case "TimeSliceMetadata":
			StageSetFormDivSelectFieldToField(&inst.TimeSliceMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), formDiv)
		case "FeatureLifetime":
			StageSetFormDivSelectFieldToField(&inst.FeatureLifetime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Designator":
			StageSetFormDivSelectFieldToField(&inst.Designator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDesignatedPointDesignatorType](), formDiv)
		case "Type_XSD":
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDesignatedPointType](), formDiv)
		case "Location":
			StageSetFormDivSelectFieldToField(&inst.Location, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), formDiv)
		case "AimingPoint":
			StageSetFormDivSelectFieldToField(&inst.AimingPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), formDiv)
		case "AirportHeliport":
			StageSetFormDivSelectFieldToField(&inst.AirportHeliport, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), formDiv)
		case "RunwayPoint":
			StageSetFormDivSelectFieldToField(&inst.RunwayPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), formDiv)
		}
	}
}

func saveStageSet_DesignatedPointType_AixmfeaturesStage(
	inst *aixmfeatures.DesignatedPointType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_DirectionFinderTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.DirectionFinderTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "DirectionFinderTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.DirectionFinderTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DirectionFinderTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_DirectionFinderTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.DirectionFinderTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "ValidTime":
			StageSetFormDivSelectFieldToField(&inst.ValidTime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Interpretation":
			FormDivBasicFieldToField(&inst.Interpretation, formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "CorrectionNumber":
			FormDivBasicFieldToField(&inst.CorrectionNumber, formDiv)
		case "TimeSliceMetadata":
			StageSetFormDivSelectFieldToField(&inst.TimeSliceMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), formDiv)
		case "FeatureLifetime":
			StageSetFormDivSelectFieldToField(&inst.FeatureLifetime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Designator":
			StageSetFormDivSelectFieldToField(&inst.Designator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavaidDesignatorType](), formDiv)
		case "EmissionClass":
			StageSetFormDivSelectFieldToField(&inst.EmissionClass, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadioEmissionType](), formDiv)
		case "Mobile":
			StageSetFormDivSelectFieldToField(&inst.Mobile, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "MagneticVariation":
			StageSetFormDivSelectFieldToField(&inst.MagneticVariation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), formDiv)
		case "MagneticVariationAccuracy":
			StageSetFormDivSelectFieldToField(&inst.MagneticVariationAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "DateMagneticVariation":
			StageSetFormDivSelectFieldToField(&inst.DateMagneticVariation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), formDiv)
		case "FlightChecked":
			StageSetFormDivSelectFieldToField(&inst.FlightChecked, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "Location":
			StageSetFormDivSelectFieldToField(&inst.Location, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), formDiv)
		case "Doppler":
			StageSetFormDivSelectFieldToField(&inst.Doppler, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		}
	}
}

func saveStageSet_DirectionFinderType_AixmfeaturesStage(
	inst *aixmfeatures.DirectionFinderType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_DistanceIndicationTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.DistanceIndicationTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "DistanceIndicationTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.DistanceIndicationTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DistanceIndicationTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_DistanceIndicationTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.DistanceIndicationTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "ValidTime":
			StageSetFormDivSelectFieldToField(&inst.ValidTime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Interpretation":
			FormDivBasicFieldToField(&inst.Interpretation, formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "CorrectionNumber":
			FormDivBasicFieldToField(&inst.CorrectionNumber, formDiv)
		case "TimeSliceMetadata":
			StageSetFormDivSelectFieldToField(&inst.TimeSliceMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), formDiv)
		case "FeatureLifetime":
			StageSetFormDivSelectFieldToField(&inst.FeatureLifetime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Distance":
			StageSetFormDivSelectFieldToField(&inst.Distance, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "MinimumReceptionAltitude":
			StageSetFormDivSelectFieldToField(&inst.MinimumReceptionAltitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "Type_XSD":
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDistanceIndicationType](), formDiv)
		case "Fix":
			StageSetFormDivSelectFieldToField(&inst.Fix, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), formDiv)
		case "PointChoice_fixDesignatedPoint":
			StageSetFormDivSelectFieldToField(&inst.PointChoice_fixDesignatedPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), formDiv)
		case "PointChoice_navaidSystem":
			StageSetFormDivSelectFieldToField(&inst.PointChoice_navaidSystem, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), formDiv)
		case "PointChoice_position":
			StageSetFormDivSelectFieldToField(&inst.PointChoice_position, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), formDiv)
		case "PointChoice_runwayPoint":
			StageSetFormDivSelectFieldToField(&inst.PointChoice_runwayPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), formDiv)
		case "PointChoice_aimingPoint":
			StageSetFormDivSelectFieldToField(&inst.PointChoice_aimingPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), formDiv)
		case "PointChoice_airportReferencePoint":
			StageSetFormDivSelectFieldToField(&inst.PointChoice_airportReferencePoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), formDiv)
		}
	}
}

func saveStageSet_DistanceIndicationType_AixmfeaturesStage(
	inst *aixmfeatures.DistanceIndicationType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_ElevatedPointPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.ElevatedPointPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "ElevatedPoint":
			StageSetFormDivSelectFieldToField(&inst.ElevatedPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointType](), formDiv)
		}
	}
}

func saveStageSet_ElevatedPointType_AixmfeaturesStage(
	inst *aixmfeatures.ElevatedPointType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "SrsName":
			FormDivBasicFieldToField(&inst.SrsName, formDiv)
		case "Pos":
			StageSetFormDivSelectFieldToField(&inst.Pos, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionType](), formDiv)
		case "HorizontalAccuracy":
			StageSetFormDivSelectFieldToField(&inst.HorizontalAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "Elevation":
			StageSetFormDivSelectFieldToField(&inst.Elevation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "GeoidUndulation":
			StageSetFormDivSelectFieldToField(&inst.GeoidUndulation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceSignedType](), formDiv)
		case "VerticalDatum":
			StageSetFormDivSelectFieldToField(&inst.VerticalDatum, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalDatumType](), formDiv)
		case "VerticalAccuracy":
			StageSetFormDivSelectFieldToField(&inst.VerticalAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		}
	}
}

func saveStageSet_ElevationTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.ElevationTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "ElevationTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.ElevationTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevationTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_ElevationTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.ElevationTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "ValidTime":
			StageSetFormDivSelectFieldToField(&inst.ValidTime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Interpretation":
			FormDivBasicFieldToField(&inst.Interpretation, formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "CorrectionNumber":
			FormDivBasicFieldToField(&inst.CorrectionNumber, formDiv)
		case "TimeSliceMetadata":
			StageSetFormDivSelectFieldToField(&inst.TimeSliceMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), formDiv)
		case "FeatureLifetime":
			StageSetFormDivSelectFieldToField(&inst.FeatureLifetime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Designator":
			StageSetFormDivSelectFieldToField(&inst.Designator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavaidDesignatorType](), formDiv)
		case "EmissionClass":
			StageSetFormDivSelectFieldToField(&inst.EmissionClass, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadioEmissionType](), formDiv)
		case "Mobile":
			StageSetFormDivSelectFieldToField(&inst.Mobile, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "MagneticVariation":
			StageSetFormDivSelectFieldToField(&inst.MagneticVariation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), formDiv)
		case "MagneticVariationAccuracy":
			StageSetFormDivSelectFieldToField(&inst.MagneticVariationAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "DateMagneticVariation":
			StageSetFormDivSelectFieldToField(&inst.DateMagneticVariation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), formDiv)
		case "FlightChecked":
			StageSetFormDivSelectFieldToField(&inst.FlightChecked, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "Location":
			StageSetFormDivSelectFieldToField(&inst.Location, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), formDiv)
		case "AngleNominal":
			StageSetFormDivSelectFieldToField(&inst.AngleNominal, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "AngleMinimum":
			StageSetFormDivSelectFieldToField(&inst.AngleMinimum, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "AngleSpan":
			StageSetFormDivSelectFieldToField(&inst.AngleSpan, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "AngleAccuracy":
			StageSetFormDivSelectFieldToField(&inst.AngleAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		}
	}
}

func saveStageSet_ElevationType_AixmfeaturesStage(
	inst *aixmfeatures.ElevationType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_GlidepathTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.GlidepathTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "GlidepathTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.GlidepathTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.GlidepathTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_GlidepathTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.GlidepathTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "ValidTime":
			StageSetFormDivSelectFieldToField(&inst.ValidTime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Interpretation":
			FormDivBasicFieldToField(&inst.Interpretation, formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "CorrectionNumber":
			FormDivBasicFieldToField(&inst.CorrectionNumber, formDiv)
		case "TimeSliceMetadata":
			StageSetFormDivSelectFieldToField(&inst.TimeSliceMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), formDiv)
		case "FeatureLifetime":
			StageSetFormDivSelectFieldToField(&inst.FeatureLifetime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Designator":
			StageSetFormDivSelectFieldToField(&inst.Designator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavaidDesignatorType](), formDiv)
		case "EmissionClass":
			StageSetFormDivSelectFieldToField(&inst.EmissionClass, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadioEmissionType](), formDiv)
		case "Mobile":
			StageSetFormDivSelectFieldToField(&inst.Mobile, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "MagneticVariation":
			StageSetFormDivSelectFieldToField(&inst.MagneticVariation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), formDiv)
		case "MagneticVariationAccuracy":
			StageSetFormDivSelectFieldToField(&inst.MagneticVariationAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "DateMagneticVariation":
			StageSetFormDivSelectFieldToField(&inst.DateMagneticVariation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), formDiv)
		case "FlightChecked":
			StageSetFormDivSelectFieldToField(&inst.FlightChecked, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "Location":
			StageSetFormDivSelectFieldToField(&inst.Location, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), formDiv)
		case "Frequency":
			StageSetFormDivSelectFieldToField(&inst.Frequency, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValFrequencyType](), formDiv)
		case "Slope":
			StageSetFormDivSelectFieldToField(&inst.Slope, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "AngleAccuracy":
			StageSetFormDivSelectFieldToField(&inst.AngleAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "Rdh":
			StageSetFormDivSelectFieldToField(&inst.Rdh, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "RdhAccuracy":
			StageSetFormDivSelectFieldToField(&inst.RdhAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		}
	}
}

func saveStageSet_GlidepathType_AixmfeaturesStage(
	inst *aixmfeatures.GlidepathType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_InformationServicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.InformationServicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "Href":
			FormDivBasicFieldToField(&inst.Href, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Arcrole":
			FormDivBasicFieldToField(&inst.Arcrole, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Show":
			FormDivBasicFieldToField(&inst.Show, formDiv)
		case "Actuate":
			FormDivBasicFieldToField(&inst.Actuate, formDiv)
		}
	}
}

func saveStageSet_LinguisticNotePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.LinguisticNotePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "LinguisticNote":
			StageSetFormDivSelectFieldToField(&inst.LinguisticNote, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.LinguisticNoteType](), formDiv)
		}
	}
}

func saveStageSet_LinguisticNoteType_AixmfeaturesStage(
	inst *aixmfeatures.LinguisticNoteType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "Note":
			StageSetFormDivSelectFieldToField(&inst.Note, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextNoteType](), formDiv)
		}
	}
}

func saveStageSet_LocalizerTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.LocalizerTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "LocalizerTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.LocalizerTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.LocalizerTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_LocalizerTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.LocalizerTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "ValidTime":
			StageSetFormDivSelectFieldToField(&inst.ValidTime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Interpretation":
			FormDivBasicFieldToField(&inst.Interpretation, formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "CorrectionNumber":
			FormDivBasicFieldToField(&inst.CorrectionNumber, formDiv)
		case "TimeSliceMetadata":
			StageSetFormDivSelectFieldToField(&inst.TimeSliceMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), formDiv)
		case "FeatureLifetime":
			StageSetFormDivSelectFieldToField(&inst.FeatureLifetime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Designator":
			StageSetFormDivSelectFieldToField(&inst.Designator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavaidDesignatorType](), formDiv)
		case "EmissionClass":
			StageSetFormDivSelectFieldToField(&inst.EmissionClass, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadioEmissionType](), formDiv)
		case "Mobile":
			StageSetFormDivSelectFieldToField(&inst.Mobile, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "MagneticVariation":
			StageSetFormDivSelectFieldToField(&inst.MagneticVariation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), formDiv)
		case "MagneticVariationAccuracy":
			StageSetFormDivSelectFieldToField(&inst.MagneticVariationAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "DateMagneticVariation":
			StageSetFormDivSelectFieldToField(&inst.DateMagneticVariation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), formDiv)
		case "FlightChecked":
			StageSetFormDivSelectFieldToField(&inst.FlightChecked, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "Location":
			StageSetFormDivSelectFieldToField(&inst.Location, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), formDiv)
		case "Frequency":
			StageSetFormDivSelectFieldToField(&inst.Frequency, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValFrequencyType](), formDiv)
		case "MagneticBearing":
			StageSetFormDivSelectFieldToField(&inst.MagneticBearing, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "MagneticBearingAccuracy":
			StageSetFormDivSelectFieldToField(&inst.MagneticBearingAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "TrueBearing":
			StageSetFormDivSelectFieldToField(&inst.TrueBearing, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "TrueBearingAccuracy":
			StageSetFormDivSelectFieldToField(&inst.TrueBearingAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "Declination":
			StageSetFormDivSelectFieldToField(&inst.Declination, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), formDiv)
		case "WidthCourse":
			StageSetFormDivSelectFieldToField(&inst.WidthCourse, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "WidthCourseAccuracy":
			StageSetFormDivSelectFieldToField(&inst.WidthCourseAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "BackCourseUsable":
			StageSetFormDivSelectFieldToField(&inst.BackCourseUsable, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeILSBackCourseType](), formDiv)
		}
	}
}

func saveStageSet_LocalizerType_AixmfeaturesStage(
	inst *aixmfeatures.LocalizerType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_MarkerBeaconTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.MarkerBeaconTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "MarkerBeaconTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.MarkerBeaconTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.MarkerBeaconTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_MarkerBeaconTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.MarkerBeaconTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "ValidTime":
			StageSetFormDivSelectFieldToField(&inst.ValidTime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Interpretation":
			FormDivBasicFieldToField(&inst.Interpretation, formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "CorrectionNumber":
			FormDivBasicFieldToField(&inst.CorrectionNumber, formDiv)
		case "TimeSliceMetadata":
			StageSetFormDivSelectFieldToField(&inst.TimeSliceMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), formDiv)
		case "FeatureLifetime":
			StageSetFormDivSelectFieldToField(&inst.FeatureLifetime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Designator":
			StageSetFormDivSelectFieldToField(&inst.Designator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavaidDesignatorType](), formDiv)
		case "EmissionClass":
			StageSetFormDivSelectFieldToField(&inst.EmissionClass, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadioEmissionType](), formDiv)
		case "Mobile":
			StageSetFormDivSelectFieldToField(&inst.Mobile, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "MagneticVariation":
			StageSetFormDivSelectFieldToField(&inst.MagneticVariation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), formDiv)
		case "MagneticVariationAccuracy":
			StageSetFormDivSelectFieldToField(&inst.MagneticVariationAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "DateMagneticVariation":
			StageSetFormDivSelectFieldToField(&inst.DateMagneticVariation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), formDiv)
		case "FlightChecked":
			StageSetFormDivSelectFieldToField(&inst.FlightChecked, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "Location":
			StageSetFormDivSelectFieldToField(&inst.Location, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), formDiv)
		case "Class":
			StageSetFormDivSelectFieldToField(&inst.Class, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeMarkerBeaconSignalType](), formDiv)
		case "Frequency":
			StageSetFormDivSelectFieldToField(&inst.Frequency, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValFrequencyType](), formDiv)
		case "AxisBearing":
			StageSetFormDivSelectFieldToField(&inst.AxisBearing, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "AuralMorseCode":
			StageSetFormDivSelectFieldToField(&inst.AuralMorseCode, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAuralMorseType](), formDiv)
		}
	}
}

func saveStageSet_MarkerBeaconType_AixmfeaturesStage(
	inst *aixmfeatures.MarkerBeaconType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_NDBTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.NDBTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "NDBTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.NDBTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NDBTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_NDBTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.NDBTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "ValidTime":
			StageSetFormDivSelectFieldToField(&inst.ValidTime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Interpretation":
			FormDivBasicFieldToField(&inst.Interpretation, formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "CorrectionNumber":
			FormDivBasicFieldToField(&inst.CorrectionNumber, formDiv)
		case "TimeSliceMetadata":
			StageSetFormDivSelectFieldToField(&inst.TimeSliceMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), formDiv)
		case "FeatureLifetime":
			StageSetFormDivSelectFieldToField(&inst.FeatureLifetime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Designator":
			StageSetFormDivSelectFieldToField(&inst.Designator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavaidDesignatorType](), formDiv)
		case "EmissionClass":
			StageSetFormDivSelectFieldToField(&inst.EmissionClass, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadioEmissionType](), formDiv)
		case "Mobile":
			StageSetFormDivSelectFieldToField(&inst.Mobile, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "MagneticVariation":
			StageSetFormDivSelectFieldToField(&inst.MagneticVariation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), formDiv)
		case "MagneticVariationAccuracy":
			StageSetFormDivSelectFieldToField(&inst.MagneticVariationAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "DateMagneticVariation":
			StageSetFormDivSelectFieldToField(&inst.DateMagneticVariation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), formDiv)
		case "FlightChecked":
			StageSetFormDivSelectFieldToField(&inst.FlightChecked, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "Location":
			StageSetFormDivSelectFieldToField(&inst.Location, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), formDiv)
		case "Frequency":
			StageSetFormDivSelectFieldToField(&inst.Frequency, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValFrequencyType](), formDiv)
		case "Class":
			StageSetFormDivSelectFieldToField(&inst.Class, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNDBUsageType](), formDiv)
		case "EmissionBand":
			StageSetFormDivSelectFieldToField(&inst.EmissionBand, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeEmissionBandType](), formDiv)
		}
	}
}

func saveStageSet_NDBType_AixmfeaturesStage(
	inst *aixmfeatures.NDBType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_NavaidComponentPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.NavaidComponentPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "NavaidComponent":
			StageSetFormDivSelectFieldToField(&inst.NavaidComponent, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidComponentType](), formDiv)
		}
	}
}

func saveStageSet_NavaidComponentType_AixmfeaturesStage(
	inst *aixmfeatures.NavaidComponentType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "CollocationGroup":
			StageSetFormDivSelectFieldToField(&inst.CollocationGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoSequenceType](), formDiv)
		case "MarkerPosition":
			StageSetFormDivSelectFieldToField(&inst.MarkerPosition, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodePositionInILSType](), formDiv)
		case "ProvidesNavigableLocation":
			StageSetFormDivSelectFieldToField(&inst.ProvidesNavigableLocation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "TheNavaidEquipment":
			StageSetFormDivSelectFieldToField(&inst.TheNavaidEquipment, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidEquipmentPropertyType](), formDiv)
		}
	}
}

func saveStageSet_NavaidEquipmentMonitoringPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.NavaidEquipmentMonitoringPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "NavaidEquipmentMonitoring":
			StageSetFormDivSelectFieldToField(&inst.NavaidEquipmentMonitoring, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidEquipmentMonitoringType](), formDiv)
		}
	}
}

func saveStageSet_NavaidEquipmentMonitoringType_AixmfeaturesStage(
	inst *aixmfeatures.NavaidEquipmentMonitoringType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "Monitored":
			StageSetFormDivSelectFieldToField(&inst.Monitored, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		}
	}
}

func saveStageSet_NavaidEquipmentPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.NavaidEquipmentPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "Href":
			FormDivBasicFieldToField(&inst.Href, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Arcrole":
			FormDivBasicFieldToField(&inst.Arcrole, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Show":
			FormDivBasicFieldToField(&inst.Show, formDiv)
		case "Actuate":
			FormDivBasicFieldToField(&inst.Actuate, formDiv)
		}
	}
}

func saveStageSet_NavaidOperationalStatusPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.NavaidOperationalStatusPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "NavaidOperationalStatus":
			StageSetFormDivSelectFieldToField(&inst.NavaidOperationalStatus, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidOperationalStatusType](), formDiv)
		}
	}
}

func saveStageSet_NavaidOperationalStatusType_AixmfeaturesStage(
	inst *aixmfeatures.NavaidOperationalStatusType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "OperationalStatus":
			StageSetFormDivSelectFieldToField(&inst.OperationalStatus, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeStatusNavaidType](), formDiv)
		case "SignalType":
			StageSetFormDivSelectFieldToField(&inst.SignalType, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadioSignalType](), formDiv)
		}
	}
}

func saveStageSet_NavaidPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.NavaidPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "Href":
			FormDivBasicFieldToField(&inst.Href, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Arcrole":
			FormDivBasicFieldToField(&inst.Arcrole, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Show":
			FormDivBasicFieldToField(&inst.Show, formDiv)
		case "Actuate":
			FormDivBasicFieldToField(&inst.Actuate, formDiv)
		}
	}
}

func saveStageSet_NavaidTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.NavaidTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "NavaidTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.NavaidTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_NavaidTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.NavaidTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "ValidTime":
			StageSetFormDivSelectFieldToField(&inst.ValidTime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Interpretation":
			FormDivBasicFieldToField(&inst.Interpretation, formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "CorrectionNumber":
			FormDivBasicFieldToField(&inst.CorrectionNumber, formDiv)
		case "TimeSliceMetadata":
			StageSetFormDivSelectFieldToField(&inst.TimeSliceMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), formDiv)
		case "FeatureLifetime":
			StageSetFormDivSelectFieldToField(&inst.FeatureLifetime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Type_XSD":
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavaidServiceType](), formDiv)
		case "Designator":
			StageSetFormDivSelectFieldToField(&inst.Designator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavaidDesignatorType](), formDiv)
		case "FlightChecked":
			StageSetFormDivSelectFieldToField(&inst.FlightChecked, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "Purpose":
			StageSetFormDivSelectFieldToField(&inst.Purpose, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavaidPurposeType](), formDiv)
		case "SignalPerformance":
			StageSetFormDivSelectFieldToField(&inst.SignalPerformance, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSignalPerformanceILSType](), formDiv)
		case "CourseQuality":
			StageSetFormDivSelectFieldToField(&inst.CourseQuality, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCourseQualityILSType](), formDiv)
		case "IntegrityLevel":
			StageSetFormDivSelectFieldToField(&inst.IntegrityLevel, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeIntegrityLevelILSType](), formDiv)
		case "Location":
			StageSetFormDivSelectFieldToField(&inst.Location, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), formDiv)
		}
	}
}

func saveStageSet_NavaidType_AixmfeaturesStage(
	inst *aixmfeatures.NavaidType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_NotePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.NotePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "Note":
			StageSetFormDivSelectFieldToField(&inst.Note, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NoteType](), formDiv)
		}
	}
}

func saveStageSet_NoteType_AixmfeaturesStage(
	inst *aixmfeatures.NoteType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "PropertyName":
			StageSetFormDivSelectFieldToField(&inst.PropertyName, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextPropertyNameType](), formDiv)
		case "Purpose":
			StageSetFormDivSelectFieldToField(&inst.Purpose, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNotePurposeType](), formDiv)
		}
	}
}

func saveStageSet_OnlineContactPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.OnlineContactPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "OnlineContact":
			StageSetFormDivSelectFieldToField(&inst.OnlineContact, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.OnlineContactType](), formDiv)
		}
	}
}

func saveStageSet_OnlineContactType_AixmfeaturesStage(
	inst *aixmfeatures.OnlineContactType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "Network":
			StageSetFormDivSelectFieldToField(&inst.Network, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTelecomNetworkType](), formDiv)
		case "Linkage":
			StageSetFormDivSelectFieldToField(&inst.Linkage, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextAddressType](), formDiv)
		case "Protocol":
			StageSetFormDivSelectFieldToField(&inst.Protocol, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextNameType](), formDiv)
		case "EMail":
			StageSetFormDivSelectFieldToField(&inst.EMail, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextAddressType](), formDiv)
		}
	}
}

func saveStageSet_OrganisationAuthorityPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.OrganisationAuthorityPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "Href":
			FormDivBasicFieldToField(&inst.Href, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Arcrole":
			FormDivBasicFieldToField(&inst.Arcrole, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Show":
			FormDivBasicFieldToField(&inst.Show, formDiv)
		case "Actuate":
			FormDivBasicFieldToField(&inst.Actuate, formDiv)
		}
	}
}

func saveStageSet_PointPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.PointPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "Point":
			StageSetFormDivSelectFieldToField(&inst.Point, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointType](), formDiv)
		case "ElevatedPoint":
			StageSetFormDivSelectFieldToField(&inst.ElevatedPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointType](), formDiv)
		}
	}
}

func saveStageSet_PointType_AixmfeaturesStage(
	inst *aixmfeatures.PointType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "SrsName":
			FormDivBasicFieldToField(&inst.SrsName, formDiv)
		case "Pos":
			StageSetFormDivSelectFieldToField(&inst.Pos, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionType](), formDiv)
		case "HorizontalAccuracy":
			StageSetFormDivSelectFieldToField(&inst.HorizontalAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		}
	}
}

func saveStageSet_PostalAddressPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.PostalAddressPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "PostalAddress":
			StageSetFormDivSelectFieldToField(&inst.PostalAddress, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PostalAddressType](), formDiv)
		}
	}
}

func saveStageSet_PostalAddressType_AixmfeaturesStage(
	inst *aixmfeatures.PostalAddressType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "DeliveryPoint":
			StageSetFormDivSelectFieldToField(&inst.DeliveryPoint, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextAddressType](), formDiv)
		case "City":
			StageSetFormDivSelectFieldToField(&inst.City, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextNameType](), formDiv)
		case "AdministrativeArea":
			StageSetFormDivSelectFieldToField(&inst.AdministrativeArea, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextNameType](), formDiv)
		case "PostalCode":
			StageSetFormDivSelectFieldToField(&inst.PostalCode, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextNameType](), formDiv)
		case "Country":
			StageSetFormDivSelectFieldToField(&inst.Country, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextNameType](), formDiv)
		}
	}
}

func saveStageSet_PrecisionApproachRadarTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.PrecisionApproachRadarTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "PrecisionApproachRadarTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.PrecisionApproachRadarTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PrecisionApproachRadarTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_PrecisionApproachRadarTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.PrecisionApproachRadarTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "ValidTime":
			StageSetFormDivSelectFieldToField(&inst.ValidTime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Interpretation":
			FormDivBasicFieldToField(&inst.Interpretation, formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "CorrectionNumber":
			FormDivBasicFieldToField(&inst.CorrectionNumber, formDiv)
		case "TimeSliceMetadata":
			StageSetFormDivSelectFieldToField(&inst.TimeSliceMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), formDiv)
		case "FeatureLifetime":
			StageSetFormDivSelectFieldToField(&inst.FeatureLifetime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "SerialNumber":
			StageSetFormDivSelectFieldToField(&inst.SerialNumber, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextDesignatorType](), formDiv)
		case "Range_XSD":
			StageSetFormDivSelectFieldToField(&inst.Range_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "RangeAccuracy":
			StageSetFormDivSelectFieldToField(&inst.RangeAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "DualChannel":
			StageSetFormDivSelectFieldToField(&inst.DualChannel, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "MovingTargetIndicator":
			StageSetFormDivSelectFieldToField(&inst.MovingTargetIndicator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "StandbyPower":
			StageSetFormDivSelectFieldToField(&inst.StandbyPower, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeStandbyPowerType](), formDiv)
		case "Digital":
			StageSetFormDivSelectFieldToField(&inst.Digital, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "MilitaryUseOnly":
			StageSetFormDivSelectFieldToField(&inst.MilitaryUseOnly, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "SpecialUseOnly":
			StageSetFormDivSelectFieldToField(&inst.SpecialUseOnly, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "SpecialAircraftOnly":
			StageSetFormDivSelectFieldToField(&inst.SpecialAircraftOnly, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "MagneticVariation":
			StageSetFormDivSelectFieldToField(&inst.MagneticVariation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), formDiv)
		case "MagneticVariationAccuracy":
			StageSetFormDivSelectFieldToField(&inst.MagneticVariationAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "DateMagneticVariation":
			StageSetFormDivSelectFieldToField(&inst.DateMagneticVariation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), formDiv)
		case "Location":
			StageSetFormDivSelectFieldToField(&inst.Location, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), formDiv)
		case "PrecisionApproachRadarType":
			StageSetFormDivSelectFieldToField(&inst.PrecisionApproachRadarType, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodePARType](), formDiv)
		case "Slope":
			StageSetFormDivSelectFieldToField(&inst.Slope, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "SlopeAccuracy":
			StageSetFormDivSelectFieldToField(&inst.SlopeAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		}
	}
}

func saveStageSet_PrecisionApproachRadarType_AixmfeaturesStage(
	inst *aixmfeatures.PrecisionApproachRadarType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_PrimarySurveillanceRadarTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.PrimarySurveillanceRadarTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "PrimarySurveillanceRadarTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.PrimarySurveillanceRadarTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PrimarySurveillanceRadarTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_PrimarySurveillanceRadarTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.PrimarySurveillanceRadarTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "ValidTime":
			StageSetFormDivSelectFieldToField(&inst.ValidTime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Interpretation":
			FormDivBasicFieldToField(&inst.Interpretation, formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "CorrectionNumber":
			FormDivBasicFieldToField(&inst.CorrectionNumber, formDiv)
		case "TimeSliceMetadata":
			StageSetFormDivSelectFieldToField(&inst.TimeSliceMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), formDiv)
		case "FeatureLifetime":
			StageSetFormDivSelectFieldToField(&inst.FeatureLifetime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "SerialNumber":
			StageSetFormDivSelectFieldToField(&inst.SerialNumber, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextDesignatorType](), formDiv)
		case "Range_XSD":
			StageSetFormDivSelectFieldToField(&inst.Range_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "RangeAccuracy":
			StageSetFormDivSelectFieldToField(&inst.RangeAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "DualChannel":
			StageSetFormDivSelectFieldToField(&inst.DualChannel, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "MovingTargetIndicator":
			StageSetFormDivSelectFieldToField(&inst.MovingTargetIndicator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "StandbyPower":
			StageSetFormDivSelectFieldToField(&inst.StandbyPower, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeStandbyPowerType](), formDiv)
		case "Digital":
			StageSetFormDivSelectFieldToField(&inst.Digital, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "MilitaryUseOnly":
			StageSetFormDivSelectFieldToField(&inst.MilitaryUseOnly, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "SpecialUseOnly":
			StageSetFormDivSelectFieldToField(&inst.SpecialUseOnly, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "SpecialAircraftOnly":
			StageSetFormDivSelectFieldToField(&inst.SpecialAircraftOnly, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "MagneticVariation":
			StageSetFormDivSelectFieldToField(&inst.MagneticVariation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), formDiv)
		case "MagneticVariationAccuracy":
			StageSetFormDivSelectFieldToField(&inst.MagneticVariationAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "DateMagneticVariation":
			StageSetFormDivSelectFieldToField(&inst.DateMagneticVariation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), formDiv)
		case "Location":
			StageSetFormDivSelectFieldToField(&inst.Location, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), formDiv)
		case "VerticalCoverageAltitude":
			StageSetFormDivSelectFieldToField(&inst.VerticalCoverageAltitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "VerticalCoverageDistance":
			StageSetFormDivSelectFieldToField(&inst.VerticalCoverageDistance, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "VerticalCoverageAzimuth":
			StageSetFormDivSelectFieldToField(&inst.VerticalCoverageAzimuth, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "AntennaTiltFixed":
			StageSetFormDivSelectFieldToField(&inst.AntennaTiltFixed, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "TiltAngle":
			StageSetFormDivSelectFieldToField(&inst.TiltAngle, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "AutomatedRadarTerminalSystem":
			StageSetFormDivSelectFieldToField(&inst.AutomatedRadarTerminalSystem, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextDesignatorType](), formDiv)
		case "Type_XSD":
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodePrimaryRadarType](), formDiv)
		}
	}
}

func saveStageSet_PrimarySurveillanceRadarType_AixmfeaturesStage(
	inst *aixmfeatures.PrimarySurveillanceRadarType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_RadarComponentPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.RadarComponentPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "RadarComponent":
			StageSetFormDivSelectFieldToField(&inst.RadarComponent, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RadarComponentType](), formDiv)
		}
	}
}

func saveStageSet_RadarComponentType_AixmfeaturesStage(
	inst *aixmfeatures.RadarComponentType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "CollocationGroup":
			StageSetFormDivSelectFieldToField(&inst.CollocationGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoSequenceType](), formDiv)
		case "TheRadarEquipment":
			StageSetFormDivSelectFieldToField(&inst.TheRadarEquipment, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RadarEquipmentPropertyType](), formDiv)
		}
	}
}

func saveStageSet_RadarEquipmentPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.RadarEquipmentPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "Href":
			FormDivBasicFieldToField(&inst.Href, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Arcrole":
			FormDivBasicFieldToField(&inst.Arcrole, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Show":
			FormDivBasicFieldToField(&inst.Show, formDiv)
		case "Actuate":
			FormDivBasicFieldToField(&inst.Actuate, formDiv)
		}
	}
}

func saveStageSet_RadarSystemTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.RadarSystemTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "RadarSystemTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.RadarSystemTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RadarSystemTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_RadarSystemTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.RadarSystemTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "ValidTime":
			StageSetFormDivSelectFieldToField(&inst.ValidTime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Interpretation":
			FormDivBasicFieldToField(&inst.Interpretation, formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "CorrectionNumber":
			FormDivBasicFieldToField(&inst.CorrectionNumber, formDiv)
		case "TimeSliceMetadata":
			StageSetFormDivSelectFieldToField(&inst.TimeSliceMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), formDiv)
		case "FeatureLifetime":
			StageSetFormDivSelectFieldToField(&inst.FeatureLifetime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Type_XSD":
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadarServiceType](), formDiv)
		case "Model":
			StageSetFormDivSelectFieldToField(&inst.Model, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextNameType](), formDiv)
		case "GeneralTerrainMonitor":
			StageSetFormDivSelectFieldToField(&inst.GeneralTerrainMonitor, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "BroadcastIdentifier":
			StageSetFormDivSelectFieldToField(&inst.BroadcastIdentifier, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextDesignatorType](), formDiv)
		case "AirportHeliport":
			StageSetFormDivSelectFieldToField(&inst.AirportHeliport, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), formDiv)
		case "Location":
			StageSetFormDivSelectFieldToField(&inst.Location, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), formDiv)
		}
	}
}

func saveStageSet_RadarSystemType_AixmfeaturesStage(
	inst *aixmfeatures.RadarSystemType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_ReflectorPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.ReflectorPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "Reflector":
			StageSetFormDivSelectFieldToField(&inst.Reflector, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ReflectorType](), formDiv)
		}
	}
}

func saveStageSet_ReflectorType_AixmfeaturesStage(
	inst *aixmfeatures.ReflectorType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "Type_XSD":
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeReflectorType](), formDiv)
		case "TouchdownReflector":
			StageSetFormDivSelectFieldToField(&inst.TouchdownReflector, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), formDiv)
		}
	}
}

func saveStageSet_RunwayCentrelinePointPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.RunwayCentrelinePointPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "Href":
			FormDivBasicFieldToField(&inst.Href, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Arcrole":
			FormDivBasicFieldToField(&inst.Arcrole, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Show":
			FormDivBasicFieldToField(&inst.Show, formDiv)
		case "Actuate":
			FormDivBasicFieldToField(&inst.Actuate, formDiv)
		}
	}
}

func saveStageSet_RunwayDirectionPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.RunwayDirectionPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "Href":
			FormDivBasicFieldToField(&inst.Href, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Arcrole":
			FormDivBasicFieldToField(&inst.Arcrole, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Show":
			FormDivBasicFieldToField(&inst.Show, formDiv)
		case "Actuate":
			FormDivBasicFieldToField(&inst.Actuate, formDiv)
		}
	}
}

func saveStageSet_RunwayPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.RunwayPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "Href":
			FormDivBasicFieldToField(&inst.Href, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Arcrole":
			FormDivBasicFieldToField(&inst.Arcrole, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Show":
			FormDivBasicFieldToField(&inst.Show, formDiv)
		case "Actuate":
			FormDivBasicFieldToField(&inst.Actuate, formDiv)
		}
	}
}

func saveStageSet_SDFTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.SDFTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "SDFTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.SDFTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SDFTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_SDFTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.SDFTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "ValidTime":
			StageSetFormDivSelectFieldToField(&inst.ValidTime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Interpretation":
			FormDivBasicFieldToField(&inst.Interpretation, formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "CorrectionNumber":
			FormDivBasicFieldToField(&inst.CorrectionNumber, formDiv)
		case "TimeSliceMetadata":
			StageSetFormDivSelectFieldToField(&inst.TimeSliceMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), formDiv)
		case "FeatureLifetime":
			StageSetFormDivSelectFieldToField(&inst.FeatureLifetime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Designator":
			StageSetFormDivSelectFieldToField(&inst.Designator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavaidDesignatorType](), formDiv)
		case "EmissionClass":
			StageSetFormDivSelectFieldToField(&inst.EmissionClass, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadioEmissionType](), formDiv)
		case "Mobile":
			StageSetFormDivSelectFieldToField(&inst.Mobile, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "MagneticVariation":
			StageSetFormDivSelectFieldToField(&inst.MagneticVariation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), formDiv)
		case "MagneticVariationAccuracy":
			StageSetFormDivSelectFieldToField(&inst.MagneticVariationAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "DateMagneticVariation":
			StageSetFormDivSelectFieldToField(&inst.DateMagneticVariation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), formDiv)
		case "FlightChecked":
			StageSetFormDivSelectFieldToField(&inst.FlightChecked, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "Location":
			StageSetFormDivSelectFieldToField(&inst.Location, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), formDiv)
		case "Frequency":
			StageSetFormDivSelectFieldToField(&inst.Frequency, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValFrequencyType](), formDiv)
		case "MagneticBearing":
			StageSetFormDivSelectFieldToField(&inst.MagneticBearing, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "TrueBearing":
			StageSetFormDivSelectFieldToField(&inst.TrueBearing, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		}
	}
}

func saveStageSet_SDFType_AixmfeaturesStage(
	inst *aixmfeatures.SDFType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_SecondarySurveillanceRadarTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.SecondarySurveillanceRadarTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "SecondarySurveillanceRadarTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.SecondarySurveillanceRadarTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SecondarySurveillanceRadarTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_SecondarySurveillanceRadarTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.SecondarySurveillanceRadarTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "ValidTime":
			StageSetFormDivSelectFieldToField(&inst.ValidTime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Interpretation":
			FormDivBasicFieldToField(&inst.Interpretation, formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "CorrectionNumber":
			FormDivBasicFieldToField(&inst.CorrectionNumber, formDiv)
		case "TimeSliceMetadata":
			StageSetFormDivSelectFieldToField(&inst.TimeSliceMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), formDiv)
		case "FeatureLifetime":
			StageSetFormDivSelectFieldToField(&inst.FeatureLifetime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "SerialNumber":
			StageSetFormDivSelectFieldToField(&inst.SerialNumber, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextDesignatorType](), formDiv)
		case "Range_XSD":
			StageSetFormDivSelectFieldToField(&inst.Range_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "RangeAccuracy":
			StageSetFormDivSelectFieldToField(&inst.RangeAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "DualChannel":
			StageSetFormDivSelectFieldToField(&inst.DualChannel, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "MovingTargetIndicator":
			StageSetFormDivSelectFieldToField(&inst.MovingTargetIndicator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "StandbyPower":
			StageSetFormDivSelectFieldToField(&inst.StandbyPower, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeStandbyPowerType](), formDiv)
		case "Digital":
			StageSetFormDivSelectFieldToField(&inst.Digital, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "MilitaryUseOnly":
			StageSetFormDivSelectFieldToField(&inst.MilitaryUseOnly, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "SpecialUseOnly":
			StageSetFormDivSelectFieldToField(&inst.SpecialUseOnly, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "SpecialAircraftOnly":
			StageSetFormDivSelectFieldToField(&inst.SpecialAircraftOnly, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "MagneticVariation":
			StageSetFormDivSelectFieldToField(&inst.MagneticVariation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), formDiv)
		case "MagneticVariationAccuracy":
			StageSetFormDivSelectFieldToField(&inst.MagneticVariationAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "DateMagneticVariation":
			StageSetFormDivSelectFieldToField(&inst.DateMagneticVariation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), formDiv)
		case "Location":
			StageSetFormDivSelectFieldToField(&inst.Location, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), formDiv)
		case "VerticalCoverageAltitude":
			StageSetFormDivSelectFieldToField(&inst.VerticalCoverageAltitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "VerticalCoverageDistance":
			StageSetFormDivSelectFieldToField(&inst.VerticalCoverageDistance, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "VerticalCoverageAzimuth":
			StageSetFormDivSelectFieldToField(&inst.VerticalCoverageAzimuth, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "AntennaTiltFixed":
			StageSetFormDivSelectFieldToField(&inst.AntennaTiltFixed, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "TiltAngle":
			StageSetFormDivSelectFieldToField(&inst.TiltAngle, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "AutomatedRadarTerminalSystem":
			StageSetFormDivSelectFieldToField(&inst.AutomatedRadarTerminalSystem, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextDesignatorType](), formDiv)
		case "Transponder":
			StageSetFormDivSelectFieldToField(&inst.Transponder, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTransponderType](), formDiv)
		case "Autonomous":
			StageSetFormDivSelectFieldToField(&inst.Autonomous, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "Monopulse":
			StageSetFormDivSelectFieldToField(&inst.Monopulse, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		}
	}
}

func saveStageSet_SecondarySurveillanceRadarType_AixmfeaturesStage(
	inst *aixmfeatures.SecondarySurveillanceRadarType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_SignificantPointInAirspaceTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.SignificantPointInAirspaceTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "SignificantPointInAirspaceTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.SignificantPointInAirspaceTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SignificantPointInAirspaceTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_SignificantPointInAirspaceTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.SignificantPointInAirspaceTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "ValidTime":
			StageSetFormDivSelectFieldToField(&inst.ValidTime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Interpretation":
			FormDivBasicFieldToField(&inst.Interpretation, formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "CorrectionNumber":
			FormDivBasicFieldToField(&inst.CorrectionNumber, formDiv)
		case "TimeSliceMetadata":
			StageSetFormDivSelectFieldToField(&inst.TimeSliceMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), formDiv)
		case "FeatureLifetime":
			StageSetFormDivSelectFieldToField(&inst.FeatureLifetime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Type_XSD":
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAirspacePointRoleType](), formDiv)
		case "RelativeLocation":
			StageSetFormDivSelectFieldToField(&inst.RelativeLocation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAirspacePointPositionType](), formDiv)
		case "ContainingAirspace":
			StageSetFormDivSelectFieldToField(&inst.ContainingAirspace, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspacePropertyType](), formDiv)
		case "Location_fixDesignatedPoint":
			StageSetFormDivSelectFieldToField(&inst.Location_fixDesignatedPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), formDiv)
		case "Location_navaidSystem":
			StageSetFormDivSelectFieldToField(&inst.Location_navaidSystem, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), formDiv)
		case "Location_position":
			StageSetFormDivSelectFieldToField(&inst.Location_position, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), formDiv)
		case "Location_runwayPoint":
			StageSetFormDivSelectFieldToField(&inst.Location_runwayPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), formDiv)
		case "Location_aimingPoint":
			StageSetFormDivSelectFieldToField(&inst.Location_aimingPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), formDiv)
		case "Location_airportReferencePoint":
			StageSetFormDivSelectFieldToField(&inst.Location_airportReferencePoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), formDiv)
		}
	}
}

func saveStageSet_SignificantPointInAirspaceType_AixmfeaturesStage(
	inst *aixmfeatures.SignificantPointInAirspaceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_SpecialNavigationStationStatusPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.SpecialNavigationStationStatusPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "SpecialNavigationStationStatus":
			StageSetFormDivSelectFieldToField(&inst.SpecialNavigationStationStatus, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SpecialNavigationStationStatusType](), formDiv)
		}
	}
}

func saveStageSet_SpecialNavigationStationStatusType_AixmfeaturesStage(
	inst *aixmfeatures.SpecialNavigationStationStatusType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "OperationalStatus":
			StageSetFormDivSelectFieldToField(&inst.OperationalStatus, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeStatusNavaidType](), formDiv)
		}
	}
}

func saveStageSet_SpecialNavigationStationTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.SpecialNavigationStationTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "SpecialNavigationStationTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.SpecialNavigationStationTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SpecialNavigationStationTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_SpecialNavigationStationTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.SpecialNavigationStationTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "ValidTime":
			StageSetFormDivSelectFieldToField(&inst.ValidTime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Interpretation":
			FormDivBasicFieldToField(&inst.Interpretation, formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "CorrectionNumber":
			FormDivBasicFieldToField(&inst.CorrectionNumber, formDiv)
		case "TimeSliceMetadata":
			StageSetFormDivSelectFieldToField(&inst.TimeSliceMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), formDiv)
		case "FeatureLifetime":
			StageSetFormDivSelectFieldToField(&inst.FeatureLifetime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Type_XSD":
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSpecialNavigationStationType](), formDiv)
		case "Frequency":
			StageSetFormDivSelectFieldToField(&inst.Frequency, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValFrequencyType](), formDiv)
		case "Emission":
			StageSetFormDivSelectFieldToField(&inst.Emission, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadioEmissionType](), formDiv)
		case "SystemChain":
			StageSetFormDivSelectFieldToField(&inst.SystemChain, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SpecialNavigationSystemPropertyType](), formDiv)
		case "ResponsibleOrganisation":
			StageSetFormDivSelectFieldToField(&inst.ResponsibleOrganisation, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AuthorityForSpecialNavigationStationPropertyType](), formDiv)
		case "Position":
			StageSetFormDivSelectFieldToField(&inst.Position, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), formDiv)
		}
	}
}

func saveStageSet_SpecialNavigationStationType_AixmfeaturesStage(
	inst *aixmfeatures.SpecialNavigationStationType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_SpecialNavigationSystemPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.SpecialNavigationSystemPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "Href":
			FormDivBasicFieldToField(&inst.Href, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Arcrole":
			FormDivBasicFieldToField(&inst.Arcrole, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Show":
			FormDivBasicFieldToField(&inst.Show, formDiv)
		case "Actuate":
			FormDivBasicFieldToField(&inst.Actuate, formDiv)
		}
	}
}

func saveStageSet_SpecialNavigationSystemTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.SpecialNavigationSystemTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "SpecialNavigationSystemTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.SpecialNavigationSystemTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SpecialNavigationSystemTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_SpecialNavigationSystemTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.SpecialNavigationSystemTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "ValidTime":
			StageSetFormDivSelectFieldToField(&inst.ValidTime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Interpretation":
			FormDivBasicFieldToField(&inst.Interpretation, formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "CorrectionNumber":
			FormDivBasicFieldToField(&inst.CorrectionNumber, formDiv)
		case "TimeSliceMetadata":
			StageSetFormDivSelectFieldToField(&inst.TimeSliceMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), formDiv)
		case "FeatureLifetime":
			StageSetFormDivSelectFieldToField(&inst.FeatureLifetime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Type_XSD":
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSpecialNavigationSystemType](), formDiv)
		case "Designator":
			StageSetFormDivSelectFieldToField(&inst.Designator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSpecialNavigationChainDesignatorType](), formDiv)
		case "ResponsibleOrganisation":
			StageSetFormDivSelectFieldToField(&inst.ResponsibleOrganisation, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AuthorityForSpecialNavigationSystemPropertyType](), formDiv)
		}
	}
}

func saveStageSet_SpecialNavigationSystemType_AixmfeaturesStage(
	inst *aixmfeatures.SpecialNavigationSystemType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_SurveillanceGroundStationPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.SurveillanceGroundStationPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "SurveillanceGroundStation":
			StageSetFormDivSelectFieldToField(&inst.SurveillanceGroundStation, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SurveillanceGroundStationType](), formDiv)
		}
	}
}

func saveStageSet_SurveillanceGroundStationType_AixmfeaturesStage(
	inst *aixmfeatures.SurveillanceGroundStationType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "VideoMap":
			StageSetFormDivSelectFieldToField(&inst.VideoMap, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "TheUnit":
			StageSetFormDivSelectFieldToField(&inst.TheUnit, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.UnitPropertyType](), formDiv)
		}
	}
}

func saveStageSet_TACANTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.TACANTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "TACANTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.TACANTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TACANTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_TACANTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.TACANTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "ValidTime":
			StageSetFormDivSelectFieldToField(&inst.ValidTime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Interpretation":
			FormDivBasicFieldToField(&inst.Interpretation, formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "CorrectionNumber":
			FormDivBasicFieldToField(&inst.CorrectionNumber, formDiv)
		case "TimeSliceMetadata":
			StageSetFormDivSelectFieldToField(&inst.TimeSliceMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), formDiv)
		case "FeatureLifetime":
			StageSetFormDivSelectFieldToField(&inst.FeatureLifetime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Designator":
			StageSetFormDivSelectFieldToField(&inst.Designator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavaidDesignatorType](), formDiv)
		case "EmissionClass":
			StageSetFormDivSelectFieldToField(&inst.EmissionClass, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadioEmissionType](), formDiv)
		case "Mobile":
			StageSetFormDivSelectFieldToField(&inst.Mobile, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "MagneticVariation":
			StageSetFormDivSelectFieldToField(&inst.MagneticVariation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), formDiv)
		case "MagneticVariationAccuracy":
			StageSetFormDivSelectFieldToField(&inst.MagneticVariationAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "DateMagneticVariation":
			StageSetFormDivSelectFieldToField(&inst.DateMagneticVariation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), formDiv)
		case "FlightChecked":
			StageSetFormDivSelectFieldToField(&inst.FlightChecked, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "Location":
			StageSetFormDivSelectFieldToField(&inst.Location, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), formDiv)
		case "Channel":
			StageSetFormDivSelectFieldToField(&inst.Channel, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTACANChannelType](), formDiv)
		case "Declination":
			StageSetFormDivSelectFieldToField(&inst.Declination, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), formDiv)
		}
	}
}

func saveStageSet_TACANType_AixmfeaturesStage(
	inst *aixmfeatures.TACANType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_TelephoneContactPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.TelephoneContactPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "TelephoneContact":
			StageSetFormDivSelectFieldToField(&inst.TelephoneContact, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TelephoneContactType](), formDiv)
		}
	}
}

func saveStageSet_TelephoneContactType_AixmfeaturesStage(
	inst *aixmfeatures.TelephoneContactType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "Voice":
			StageSetFormDivSelectFieldToField(&inst.Voice, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextPhoneType](), formDiv)
		case "Facsimile":
			StageSetFormDivSelectFieldToField(&inst.Facsimile, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextPhoneType](), formDiv)
		}
	}
}

func saveStageSet_TimesheetPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.TimesheetPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "Timesheet":
			StageSetFormDivSelectFieldToField(&inst.Timesheet, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TimesheetType](), formDiv)
		}
	}
}

func saveStageSet_TimesheetType_AixmfeaturesStage(
	inst *aixmfeatures.TimesheetType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "TimeReference":
			StageSetFormDivSelectFieldToField(&inst.TimeReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTimeReferenceType](), formDiv)
		case "StartDate":
			StageSetFormDivSelectFieldToField(&inst.StartDate, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateMonthDayType](), formDiv)
		case "EndDate":
			StageSetFormDivSelectFieldToField(&inst.EndDate, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateMonthDayType](), formDiv)
		case "Day":
			StageSetFormDivSelectFieldToField(&inst.Day, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDayType](), formDiv)
		case "DayTil":
			StageSetFormDivSelectFieldToField(&inst.DayTil, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDayType](), formDiv)
		case "StartTime":
			StageSetFormDivSelectFieldToField(&inst.StartTime, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TimeType](), formDiv)
		case "StartEvent":
			StageSetFormDivSelectFieldToField(&inst.StartEvent, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTimeEventType](), formDiv)
		case "StartTimeRelativeEvent":
			StageSetFormDivSelectFieldToField(&inst.StartTimeRelativeEvent, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDurationType](), formDiv)
		case "StartEventInterpretation":
			StageSetFormDivSelectFieldToField(&inst.StartEventInterpretation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTimeEventCombinationType](), formDiv)
		case "EndTime":
			StageSetFormDivSelectFieldToField(&inst.EndTime, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TimeType](), formDiv)
		case "EndEvent":
			StageSetFormDivSelectFieldToField(&inst.EndEvent, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTimeEventType](), formDiv)
		case "EndTimeRelativeEvent":
			StageSetFormDivSelectFieldToField(&inst.EndTimeRelativeEvent, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDurationType](), formDiv)
		case "EndEventInterpretation":
			StageSetFormDivSelectFieldToField(&inst.EndEventInterpretation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTimeEventCombinationType](), formDiv)
		case "DaylightSavingAdjust":
			StageSetFormDivSelectFieldToField(&inst.DaylightSavingAdjust, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "Excluded":
			StageSetFormDivSelectFieldToField(&inst.Excluded, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		}
	}
}

func saveStageSet_TouchDownLiftOffPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.TouchDownLiftOffPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "Href":
			FormDivBasicFieldToField(&inst.Href, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Arcrole":
			FormDivBasicFieldToField(&inst.Arcrole, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Show":
			FormDivBasicFieldToField(&inst.Show, formDiv)
		case "Actuate":
			FormDivBasicFieldToField(&inst.Actuate, formDiv)
		}
	}
}

func saveStageSet_UnitPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.UnitPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "Href":
			FormDivBasicFieldToField(&inst.Href, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Arcrole":
			FormDivBasicFieldToField(&inst.Arcrole, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Show":
			FormDivBasicFieldToField(&inst.Show, formDiv)
		case "Actuate":
			FormDivBasicFieldToField(&inst.Actuate, formDiv)
		}
	}
}

func saveStageSet_VORPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.VORPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "Href":
			FormDivBasicFieldToField(&inst.Href, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Arcrole":
			FormDivBasicFieldToField(&inst.Arcrole, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Show":
			FormDivBasicFieldToField(&inst.Show, formDiv)
		case "Actuate":
			FormDivBasicFieldToField(&inst.Actuate, formDiv)
		}
	}
}

func saveStageSet_VORTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.VORTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "VORTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.VORTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.VORTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_VORTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.VORTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "ValidTime":
			StageSetFormDivSelectFieldToField(&inst.ValidTime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Interpretation":
			FormDivBasicFieldToField(&inst.Interpretation, formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "CorrectionNumber":
			FormDivBasicFieldToField(&inst.CorrectionNumber, formDiv)
		case "TimeSliceMetadata":
			StageSetFormDivSelectFieldToField(&inst.TimeSliceMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), formDiv)
		case "FeatureLifetime":
			StageSetFormDivSelectFieldToField(&inst.FeatureLifetime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Designator":
			StageSetFormDivSelectFieldToField(&inst.Designator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavaidDesignatorType](), formDiv)
		case "EmissionClass":
			StageSetFormDivSelectFieldToField(&inst.EmissionClass, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadioEmissionType](), formDiv)
		case "Mobile":
			StageSetFormDivSelectFieldToField(&inst.Mobile, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "MagneticVariation":
			StageSetFormDivSelectFieldToField(&inst.MagneticVariation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), formDiv)
		case "MagneticVariationAccuracy":
			StageSetFormDivSelectFieldToField(&inst.MagneticVariationAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "DateMagneticVariation":
			StageSetFormDivSelectFieldToField(&inst.DateMagneticVariation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), formDiv)
		case "FlightChecked":
			StageSetFormDivSelectFieldToField(&inst.FlightChecked, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "Location":
			StageSetFormDivSelectFieldToField(&inst.Location, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), formDiv)
		case "Type_XSD":
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVORType](), formDiv)
		case "Frequency":
			StageSetFormDivSelectFieldToField(&inst.Frequency, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValFrequencyType](), formDiv)
		case "ZeroBearingDirection":
			StageSetFormDivSelectFieldToField(&inst.ZeroBearingDirection, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNorthReferenceType](), formDiv)
		case "Declination":
			StageSetFormDivSelectFieldToField(&inst.Declination, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), formDiv)
		}
	}
}

func saveStageSet_VORType_AixmfeaturesStage(
	inst *aixmfeatures.VORType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_CodeAirspacePointPositionType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeAirspacePointPositionType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeAirspacePointRoleType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeAirspacePointRoleType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeAltitudeUseType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeAltitudeUseType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeAuralMorseType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeAuralMorseType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeAuthorityRoleType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeAuthorityRoleType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeBearingType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeBearingType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeCardinalDirectionType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeCardinalDirectionType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeCheckpointCategoryType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeCheckpointCategoryType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeCourseQualityILSType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeCourseQualityILSType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeDMEChannelType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeDMEChannelType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeDMEType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeDMEType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeDayType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeDayType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeDesignatedPointDesignatorType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeDesignatedPointDesignatorType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeDesignatedPointType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeDesignatedPointType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeDirectionReferenceType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeDirectionReferenceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeDistanceIndicationType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeDistanceIndicationType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeEmissionBandType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeEmissionBandType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeILSBackCourseType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeILSBackCourseType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeIntegrityLevelILSType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeIntegrityLevelILSType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeMLSAzimuthType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeMLSAzimuthType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeMLSChannelType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeMLSChannelType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeMarkerBeaconSignalType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeMarkerBeaconSignalType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeNDBUsageType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeNDBUsageType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeNavaidDesignatorType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeNavaidDesignatorType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeNavaidPurposeType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeNavaidPurposeType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeNavaidServiceType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeNavaidServiceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeNorthReferenceType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeNorthReferenceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeNotePurposeType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeNotePurposeType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodePARType_AixmdatatypesStage(
	inst *aixmdatatypes.CodePARType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodePositionInILSType_AixmdatatypesStage(
	inst *aixmdatatypes.CodePositionInILSType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodePrimaryRadarType_AixmdatatypesStage(
	inst *aixmdatatypes.CodePrimaryRadarType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeRadarServiceType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeRadarServiceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeRadioEmissionType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeRadioEmissionType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeRadioSignalType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeRadioSignalType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeReflectorType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeReflectorType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeSignalPerformanceILSType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeSignalPerformanceILSType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeSpecialNavigationChainDesignatorType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeSpecialNavigationChainDesignatorType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeSpecialNavigationStationType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeSpecialNavigationStationType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeSpecialNavigationSystemType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeSpecialNavigationSystemType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeStandbyPowerType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeStandbyPowerType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeStatusNavaidType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeStatusNavaidType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeTACANChannelType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeTACANChannelType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeTelecomNetworkType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeTelecomNetworkType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeTimeEventCombinationType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeTimeEventCombinationType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeTimeEventType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeTimeEventType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeTimeReferenceType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeTimeReferenceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeTransponderType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeTransponderType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeVORType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeVORType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeVerticalDatumType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeVerticalDatumType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeVerticalReferenceType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeVerticalReferenceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeYesNoType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeYesNoType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_DateMonthDayType_AixmdatatypesStage(
	inst *aixmdatatypes.DateMonthDayType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_DateYearType_AixmdatatypesStage(
	inst *aixmdatatypes.DateYearType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_NoSequenceType_AixmdatatypesStage(
	inst *aixmdatatypes.NoSequenceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_TextAddressType_AixmdatatypesStage(
	inst *aixmdatatypes.TextAddressType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_TextDesignatorType_AixmdatatypesStage(
	inst *aixmdatatypes.TextDesignatorType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_TextNameType_AixmdatatypesStage(
	inst *aixmdatatypes.TextNameType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_TextNoteType_AixmdatatypesStage(
	inst *aixmdatatypes.TextNoteType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "Lang":
			FormDivBasicFieldToField(&inst.Lang, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_TextPhoneType_AixmdatatypesStage(
	inst *aixmdatatypes.TextPhoneType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_TextPropertyNameType_AixmdatatypesStage(
	inst *aixmdatatypes.TextPropertyNameType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_TimeType_AixmdatatypesStage(
	inst *aixmdatatypes.TimeType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_ValAngleType_AixmdatatypesStage(
	inst *aixmdatatypes.ValAngleType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_ValBearingType_AixmdatatypesStage(
	inst *aixmdatatypes.ValBearingType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_ValDistanceSignedType_AixmdatatypesStage(
	inst *aixmdatatypes.ValDistanceSignedType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Uom":
			FormDivBasicFieldToField(&inst.Uom, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_ValDistanceType_AixmdatatypesStage(
	inst *aixmdatatypes.ValDistanceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Uom":
			FormDivBasicFieldToField(&inst.Uom, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_ValDistanceVerticalType_AixmdatatypesStage(
	inst *aixmdatatypes.ValDistanceVerticalType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Uom":
			FormDivBasicFieldToField(&inst.Uom, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_ValDurationType_AixmdatatypesStage(
	inst *aixmdatatypes.ValDurationType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Uom":
			FormDivBasicFieldToField(&inst.Uom, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_ValFrequencyType_AixmdatatypesStage(
	inst *aixmdatatypes.ValFrequencyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Uom":
			FormDivBasicFieldToField(&inst.Uom, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_ValMagneticVariationType_AixmdatatypesStage(
	inst *aixmdatatypes.ValMagneticVariationType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_AbstractAIXMFeatureBaseType_AixmabstractgmlobjecttypesStage(
	inst *aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		}
	}
}

func saveStageSet_AbstractAIXMFeatureType_AixmabstractgmlobjecttypesStage(
	inst *aixmabstractgmlobjecttypes.AbstractAIXMFeatureType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "FeatureMetadata":
			StageSetFormDivSelectFieldToField(&inst.FeatureMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_AbstractAIXMMessageBaseType_AixmabstractgmlobjecttypesStage(
	inst *aixmabstractgmlobjecttypes.AbstractAIXMMessageBaseType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "MessageMetadata":
			StageSetFormDivSelectFieldToField(&inst.MessageMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.MessageMetadataPropertyType](), formDiv)
		}
	}
}

func saveStageSet_AbstractAIXMMessageType_AixmabstractgmlobjecttypesStage(
	inst *aixmabstractgmlobjecttypes.AbstractAIXMMessageType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "MessageMetadata":
			StageSetFormDivSelectFieldToField(&inst.MessageMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.MessageMetadataPropertyType](), formDiv)
		case "AggregationType":
			FormDivBasicFieldToField(&inst.AggregationType, formDiv)
		}
	}
}

func saveStageSet_AbstractAIXMObjectType_AixmabstractgmlobjecttypesStage(
	inst *aixmabstractgmlobjecttypes.AbstractAIXMObjectType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		}
	}
}

func saveStageSet_AbstractAIXMPropertyType_AixmabstractgmlobjecttypesStage(
	inst *aixmabstractgmlobjecttypes.AbstractAIXMPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		}
	}
}

func saveStageSet_AbstractAIXMTimeSliceBaseType_AixmabstractgmlobjecttypesStage(
	inst *aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceBaseType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "ValidTime":
			StageSetFormDivSelectFieldToField(&inst.ValidTime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		}
	}
}

func saveStageSet_AbstractAIXMTimeSliceType_AixmabstractgmlobjecttypesStage(
	inst *aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "ValidTime":
			StageSetFormDivSelectFieldToField(&inst.ValidTime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		case "Interpretation":
			StageSetFormDivEnumStringFieldToField(&inst.Interpretation, formDiv)
		case "SequenceNumber":
			FormDivBasicFieldToField(&inst.SequenceNumber, formDiv)
		case "CorrectionNumber":
			FormDivBasicFieldToField(&inst.CorrectionNumber, formDiv)
		case "TimeSliceMetadata":
			StageSetFormDivSelectFieldToField(&inst.TimeSliceMetadata, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), formDiv)
		case "FeatureLifetime":
			StageSetFormDivSelectFieldToField(&inst.FeatureLifetime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		}
	}
}

func saveStageSet_AbstractExtensionType_AixmabstractgmlobjecttypesStage(
	inst *aixmabstractgmlobjecttypes.AbstractExtensionType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		}
	}
}

func saveStageSet_FeatureMetadataPropertyType_AixmabstractgmlobjecttypesStage(
	inst *aixmabstractgmlobjecttypes.FeatureMetadataPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_FeatureTimeSliceMetadataPropertyType_AixmabstractgmlobjecttypesStage(
	inst *aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_MessageMetadataPropertyType_AixmabstractgmlobjecttypesStage(
	inst *aixmabstractgmlobjecttypes.MessageMetadataPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_AbstractFeatureMemberType_GmlStage(
	inst *gml.AbstractFeatureMemberType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		}
	}
}

func saveStageSet_AbstractFeatureType_GmlStage(
	inst *gml.AbstractFeatureType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		}
	}
}

func saveStageSet_AbstractGMLType_GmlStage(
	inst *gml.AbstractGMLType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		}
	}
}

func saveStageSet_AbstractGeometricPrimitiveType_GmlStage(
	inst *gml.AbstractGeometricPrimitiveType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "SrsName":
			FormDivBasicFieldToField(&inst.SrsName, formDiv)
		}
	}
}

func saveStageSet_AbstractGeometryType_GmlStage(
	inst *gml.AbstractGeometryType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "SrsName":
			FormDivBasicFieldToField(&inst.SrsName, formDiv)
		}
	}
}

func saveStageSet_AbstractMetadataPropertyType_GmlStage(
	inst *gml.AbstractMetadataPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		}
	}
}

func saveStageSet_AbstractTimeGeometricPrimitiveType_GmlStage(
	inst *gml.AbstractTimeGeometricPrimitiveType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "Frame":
			FormDivBasicFieldToField(&inst.Frame, formDiv)
		}
	}
}

func saveStageSet_AbstractTimeObjectType_GmlStage(
	inst *gml.AbstractTimeObjectType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		}
	}
}

func saveStageSet_AbstractTimePrimitiveType_GmlStage(
	inst *gml.AbstractTimePrimitiveType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		}
	}
}

func saveStageSet_AbstractTimeSliceType_GmlStage(
	inst *gml.AbstractTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "ValidTime":
			StageSetFormDivSelectFieldToField(&inst.ValidTime, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), formDiv)
		}
	}
}

func saveStageSet_BoundingShapeType_GmlStage(
	inst *gml.BoundingShapeType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "Envelope":
			StageSetFormDivSelectFieldToField(&inst.Envelope, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.EnvelopeType](), formDiv)
		}
	}
}

func saveStageSet_CodeType_GmlStage(
	inst *gml.CodeType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "CodeSpace":
			FormDivBasicFieldToField(&inst.CodeSpace, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeWithAuthorityType_GmlStage(
	inst *gml.CodeWithAuthorityType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "CodeSpace":
			FormDivBasicFieldToField(&inst.CodeSpace, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_DirectPositionType_GmlStage(
	inst *gml.DirectPositionType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "SrsName":
			FormDivBasicFieldToField(&inst.SrsName, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_DynamicFeatureType_GmlStage(
	inst *gml.DynamicFeatureType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "BoundedBy":
			StageSetFormDivSelectFieldToField(&inst.BoundedBy, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), formDiv)
		}
	}
}

func saveStageSet_EnvelopeType_GmlStage(
	inst *gml.EnvelopeType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "SrsName":
			FormDivBasicFieldToField(&inst.SrsName, formDiv)
		case "LowerCorner":
			StageSetFormDivSelectFieldToField(&inst.LowerCorner, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionType](), formDiv)
		case "UpperCorner":
			StageSetFormDivSelectFieldToField(&inst.UpperCorner, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionType](), formDiv)
		}
	}
}

func saveStageSet_PointType_GmlStage(
	inst *gml.PointType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "SrsName":
			FormDivBasicFieldToField(&inst.SrsName, formDiv)
		case "Pos":
			StageSetFormDivSelectFieldToField(&inst.Pos, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionType](), formDiv)
		}
	}
}

func saveStageSet_TimeInstantType_GmlStage(
	inst *gml.TimeInstantType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "Frame":
			FormDivBasicFieldToField(&inst.Frame, formDiv)
		case "TimePosition":
			StageSetFormDivSelectFieldToField(&inst.TimePosition, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePositionType](), formDiv)
		}
	}
}

func saveStageSet_TimePeriodType_GmlStage(
	inst *gml.TimePeriodType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "Frame":
			FormDivBasicFieldToField(&inst.Frame, formDiv)
		case "BeginPosition":
			StageSetFormDivSelectFieldToField(&inst.BeginPosition, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePositionType](), formDiv)
		case "EndPosition":
			StageSetFormDivSelectFieldToField(&inst.EndPosition, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePositionType](), formDiv)
		}
	}
}

func saveStageSet_TimePositionType_GmlStage(
	inst *gml.TimePositionType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Frame":
			FormDivBasicFieldToField(&inst.Frame, formDiv)
		case "IndeterminatePosition":
			StageSetFormDivEnumStringFieldToField(&inst.IndeterminatePosition, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_TimePrimitivePropertyType_GmlStage(
	inst *gml.TimePrimitivePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "Type_XSD":
			FormDivBasicFieldToField(&inst.Type_XSD, formDiv)
		case "Href":
			FormDivBasicFieldToField(&inst.Href, formDiv)
		case "Role":
			FormDivBasicFieldToField(&inst.Role, formDiv)
		case "Arcrole":
			FormDivBasicFieldToField(&inst.Arcrole, formDiv)
		case "Title":
			FormDivBasicFieldToField(&inst.Title, formDiv)
		case "Show":
			FormDivBasicFieldToField(&inst.Show, formDiv)
		case "Actuate":
			FormDivBasicFieldToField(&inst.Actuate, formDiv)
		case "TimeInstant":
			StageSetFormDivSelectFieldToField(&inst.TimeInstant, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimeInstantType](), formDiv)
		case "TimePeriod":
			StageSetFormDivSelectFieldToField(&inst.TimePeriod, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePeriodType](), formDiv)
		}
	}
}

func saveStageSet_Xlink_SchemaElement_XlinkStage(
	inst *xlink.Xlink_SchemaElement,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		}
	}
}


func StageSetNewInstance_AIXMBasicMessage_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AIXMBasicMessage",
	}).Stage(probe.formStage)
	inst := new(models.AIXMBasicMessage)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_AIXMBasicMessage_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_AIXMBasicMessage_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AIXMBasicMessageType_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AIXMBasicMessageType",
	}).Stage(probe.formStage)
	inst := new(models.AIXMBasicMessageType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_AIXMBasicMessageType_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_AIXMBasicMessageType_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_BasicMessageMemberAIXMPropertyType_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New BasicMessageMemberAIXMPropertyType",
	}).Stage(probe.formStage)
	inst := new(models.BasicMessageMemberAIXMPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_BasicMessageMemberAIXMPropertyType_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_BasicMessageMemberAIXMPropertyType_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_1_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_1",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_1)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_1_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_1_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_10_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_10",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_10)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_10_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_10_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_11_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_11",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_11)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_11_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_11_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_12_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_12",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_12)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_12_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_12_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_13_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_13",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_13)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_13_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_13_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_14_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_14",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_14)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_14_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_14_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_15_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_15",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_15)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_15_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_15_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_16_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_16",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_16)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_16_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_16_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_17_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_17",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_17)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_17_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_17_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_18_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_18",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_18)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_18_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_18_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_19_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_19",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_19)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_19_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_19_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_2_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_2",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_2)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_2_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_2_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_20_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_20",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_20)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_20_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_20_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_21_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_21",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_21)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_21_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_21_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_22_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_22",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_22)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_22_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_22_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_23_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_23",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_23)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_23_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_23_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_24_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_24",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_24)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_24_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_24_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_25_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_25",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_25)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_25_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_25_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_26_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_26",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_26)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_26_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_26_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_27_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_27",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_27)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_27_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_27_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_28_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_28",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_28)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_28_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_28_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_29_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_29",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_29)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_29_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_29_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_3_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_3",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_3)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_3_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_3_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_30_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_30",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_30)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_30_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_30_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_31_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_31",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_31)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_31_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_31_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_32_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_32",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_32)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_32_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_32_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_33_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_33",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_33)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_33_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_33_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_34_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_34",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_34)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_34_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_34_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_35_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_35",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_35)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_35_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_35_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_36_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_36",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_36)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_36_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_36_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_37_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_37",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_37)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_37_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_37_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_38_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_38",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_38)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_38_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_38_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_39_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_39",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_39)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_39_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_39_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_4_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_4",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_4)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_4_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_4_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_40_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_40",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_40)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_40_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_40_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_41_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_41",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_41)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_41_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_41_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_5_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_5",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_5)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_5_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_5_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_6_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_6",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_6)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_6_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_6_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_7_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_7",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_7)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_7_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_7_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_8_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_8",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_8)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_8_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_8_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_9_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_9",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_9)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_9_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_9_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractNavaidEquipmentType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractNavaidEquipmentType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AbstractNavaidEquipmentType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AbstractNavaidEquipmentType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AbstractNavaidEquipmentType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractNavigationSystemCheckpointType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractNavigationSystemCheckpointType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AbstractNavigationSystemCheckpointType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AbstractNavigationSystemCheckpointType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AbstractNavigationSystemCheckpointType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractPropertiesWithScheduleType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractPropertiesWithScheduleType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AbstractPropertiesWithScheduleType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AbstractPropertiesWithScheduleType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AbstractPropertiesWithScheduleType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractRadarEquipmentType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractRadarEquipmentType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AbstractRadarEquipmentType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AbstractRadarEquipmentType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AbstractRadarEquipmentType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractSurveillanceRadarType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractSurveillanceRadarType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AbstractSurveillanceRadarType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AbstractSurveillanceRadarType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AbstractSurveillanceRadarType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AirportHeliportPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AirportHeliportPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AirportHeliportPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AirportHeliportPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AirportHeliportPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AirspacePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AirspacePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AirspacePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AirspacePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AirspacePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AngleIndicationTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AngleIndicationTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AngleIndicationTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AngleIndicationTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AngleIndicationTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AngleIndicationTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AngleIndicationTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AngleIndicationTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AngleIndicationTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AngleIndicationTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AngleIndicationType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AngleIndicationType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AngleIndicationType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AngleIndicationType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AngleIndicationType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AuthorityForNavaidEquipmentPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AuthorityForNavaidEquipmentPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AuthorityForNavaidEquipmentPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AuthorityForNavaidEquipmentPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AuthorityForNavaidEquipmentPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AuthorityForNavaidEquipmentType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AuthorityForNavaidEquipmentType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AuthorityForNavaidEquipmentType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AuthorityForNavaidEquipmentType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AuthorityForNavaidEquipmentType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AuthorityForSpecialNavigationStationPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AuthorityForSpecialNavigationStationPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AuthorityForSpecialNavigationStationPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AuthorityForSpecialNavigationStationPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AuthorityForSpecialNavigationStationPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AuthorityForSpecialNavigationStationType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AuthorityForSpecialNavigationStationType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AuthorityForSpecialNavigationStationType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AuthorityForSpecialNavigationStationType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AuthorityForSpecialNavigationStationType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AuthorityForSpecialNavigationSystemPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AuthorityForSpecialNavigationSystemPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AuthorityForSpecialNavigationSystemPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AuthorityForSpecialNavigationSystemPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AuthorityForSpecialNavigationSystemPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AuthorityForSpecialNavigationSystemType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AuthorityForSpecialNavigationSystemType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AuthorityForSpecialNavigationSystemType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AuthorityForSpecialNavigationSystemType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AuthorityForSpecialNavigationSystemType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AzimuthTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AzimuthTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AzimuthTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AzimuthTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AzimuthTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AzimuthTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AzimuthTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AzimuthTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AzimuthTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AzimuthTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AzimuthType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AzimuthType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AzimuthType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AzimuthType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AzimuthType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CheckpointINSTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CheckpointINSTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.CheckpointINSTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_CheckpointINSTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_CheckpointINSTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CheckpointINSTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CheckpointINSTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.CheckpointINSTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_CheckpointINSTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_CheckpointINSTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CheckpointINSType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CheckpointINSType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.CheckpointINSType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_CheckpointINSType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_CheckpointINSType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CheckpointVORTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CheckpointVORTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.CheckpointVORTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_CheckpointVORTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_CheckpointVORTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CheckpointVORTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CheckpointVORTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.CheckpointVORTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_CheckpointVORTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_CheckpointVORTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CheckpointVORType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CheckpointVORType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.CheckpointVORType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_CheckpointVORType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_CheckpointVORType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ContactInformationPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ContactInformationPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ContactInformationPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ContactInformationPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ContactInformationPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ContactInformationType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ContactInformationType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ContactInformationType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ContactInformationType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ContactInformationType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DMETimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DMETimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.DMETimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_DMETimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_DMETimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DMETimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DMETimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.DMETimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_DMETimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_DMETimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DMEType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DMEType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.DMEType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_DMEType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_DMEType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DesignatedPointPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DesignatedPointPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.DesignatedPointPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_DesignatedPointPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_DesignatedPointPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DesignatedPointTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DesignatedPointTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.DesignatedPointTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_DesignatedPointTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_DesignatedPointTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DesignatedPointTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DesignatedPointTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.DesignatedPointTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_DesignatedPointTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_DesignatedPointTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DesignatedPointType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DesignatedPointType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.DesignatedPointType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_DesignatedPointType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_DesignatedPointType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DirectionFinderTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DirectionFinderTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.DirectionFinderTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_DirectionFinderTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_DirectionFinderTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DirectionFinderTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DirectionFinderTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.DirectionFinderTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_DirectionFinderTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_DirectionFinderTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DirectionFinderType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DirectionFinderType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.DirectionFinderType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_DirectionFinderType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_DirectionFinderType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DistanceIndicationTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DistanceIndicationTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.DistanceIndicationTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_DistanceIndicationTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_DistanceIndicationTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DistanceIndicationTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DistanceIndicationTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.DistanceIndicationTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_DistanceIndicationTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_DistanceIndicationTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DistanceIndicationType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DistanceIndicationType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.DistanceIndicationType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_DistanceIndicationType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_DistanceIndicationType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ElevatedPointPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ElevatedPointPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ElevatedPointPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ElevatedPointPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ElevatedPointPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ElevatedPointType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ElevatedPointType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ElevatedPointType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ElevatedPointType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ElevatedPointType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ElevationTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ElevationTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ElevationTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ElevationTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ElevationTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ElevationTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ElevationTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ElevationTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ElevationTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ElevationTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ElevationType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ElevationType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ElevationType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ElevationType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ElevationType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_GlidepathTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New GlidepathTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.GlidepathTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_GlidepathTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_GlidepathTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_GlidepathTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New GlidepathTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.GlidepathTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_GlidepathTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_GlidepathTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_GlidepathType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New GlidepathType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.GlidepathType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_GlidepathType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_GlidepathType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_InformationServicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New InformationServicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.InformationServicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_InformationServicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_InformationServicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_LinguisticNotePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New LinguisticNotePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.LinguisticNotePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_LinguisticNotePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_LinguisticNotePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_LinguisticNoteType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New LinguisticNoteType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.LinguisticNoteType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_LinguisticNoteType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_LinguisticNoteType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_LocalizerTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New LocalizerTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.LocalizerTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_LocalizerTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_LocalizerTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_LocalizerTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New LocalizerTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.LocalizerTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_LocalizerTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_LocalizerTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_LocalizerType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New LocalizerType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.LocalizerType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_LocalizerType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_LocalizerType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_MarkerBeaconTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New MarkerBeaconTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.MarkerBeaconTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_MarkerBeaconTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_MarkerBeaconTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_MarkerBeaconTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New MarkerBeaconTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.MarkerBeaconTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_MarkerBeaconTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_MarkerBeaconTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_MarkerBeaconType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New MarkerBeaconType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.MarkerBeaconType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_MarkerBeaconType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_MarkerBeaconType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NDBTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New NDBTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.NDBTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_NDBTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_NDBTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NDBTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New NDBTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.NDBTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_NDBTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_NDBTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NDBType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New NDBType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.NDBType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_NDBType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_NDBType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NavaidComponentPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New NavaidComponentPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.NavaidComponentPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_NavaidComponentPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_NavaidComponentPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NavaidComponentType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New NavaidComponentType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.NavaidComponentType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_NavaidComponentType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_NavaidComponentType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NavaidEquipmentMonitoringPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New NavaidEquipmentMonitoringPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.NavaidEquipmentMonitoringPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_NavaidEquipmentMonitoringPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_NavaidEquipmentMonitoringPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NavaidEquipmentMonitoringType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New NavaidEquipmentMonitoringType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.NavaidEquipmentMonitoringType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_NavaidEquipmentMonitoringType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_NavaidEquipmentMonitoringType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NavaidEquipmentPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New NavaidEquipmentPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.NavaidEquipmentPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_NavaidEquipmentPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_NavaidEquipmentPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NavaidOperationalStatusPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New NavaidOperationalStatusPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.NavaidOperationalStatusPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_NavaidOperationalStatusPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_NavaidOperationalStatusPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NavaidOperationalStatusType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New NavaidOperationalStatusType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.NavaidOperationalStatusType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_NavaidOperationalStatusType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_NavaidOperationalStatusType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NavaidPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New NavaidPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.NavaidPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_NavaidPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_NavaidPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NavaidTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New NavaidTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.NavaidTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_NavaidTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_NavaidTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NavaidTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New NavaidTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.NavaidTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_NavaidTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_NavaidTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NavaidType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New NavaidType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.NavaidType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_NavaidType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_NavaidType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NotePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New NotePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.NotePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_NotePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_NotePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NoteType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New NoteType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.NoteType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_NoteType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_NoteType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_OnlineContactPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New OnlineContactPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.OnlineContactPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_OnlineContactPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_OnlineContactPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_OnlineContactType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New OnlineContactType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.OnlineContactType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_OnlineContactType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_OnlineContactType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_OrganisationAuthorityPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New OrganisationAuthorityPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.OrganisationAuthorityPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_OrganisationAuthorityPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_OrganisationAuthorityPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_PointPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New PointPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.PointPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_PointPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_PointPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_PointType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New PointType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.PointType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_PointType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_PointType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_PostalAddressPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New PostalAddressPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.PostalAddressPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_PostalAddressPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_PostalAddressPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_PostalAddressType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New PostalAddressType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.PostalAddressType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_PostalAddressType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_PostalAddressType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_PrecisionApproachRadarTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New PrecisionApproachRadarTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.PrecisionApproachRadarTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_PrecisionApproachRadarTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_PrecisionApproachRadarTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_PrecisionApproachRadarTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New PrecisionApproachRadarTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.PrecisionApproachRadarTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_PrecisionApproachRadarTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_PrecisionApproachRadarTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_PrecisionApproachRadarType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New PrecisionApproachRadarType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.PrecisionApproachRadarType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_PrecisionApproachRadarType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_PrecisionApproachRadarType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_PrimarySurveillanceRadarTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New PrimarySurveillanceRadarTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.PrimarySurveillanceRadarTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_PrimarySurveillanceRadarTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_PrimarySurveillanceRadarTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_PrimarySurveillanceRadarTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New PrimarySurveillanceRadarTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.PrimarySurveillanceRadarTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_PrimarySurveillanceRadarTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_PrimarySurveillanceRadarTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_PrimarySurveillanceRadarType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New PrimarySurveillanceRadarType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.PrimarySurveillanceRadarType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_PrimarySurveillanceRadarType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_PrimarySurveillanceRadarType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RadarComponentPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RadarComponentPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.RadarComponentPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_RadarComponentPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_RadarComponentPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RadarComponentType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RadarComponentType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.RadarComponentType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_RadarComponentType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_RadarComponentType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RadarEquipmentPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RadarEquipmentPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.RadarEquipmentPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_RadarEquipmentPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_RadarEquipmentPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RadarSystemTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RadarSystemTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.RadarSystemTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_RadarSystemTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_RadarSystemTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RadarSystemTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RadarSystemTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.RadarSystemTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_RadarSystemTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_RadarSystemTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RadarSystemType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RadarSystemType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.RadarSystemType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_RadarSystemType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_RadarSystemType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ReflectorPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ReflectorPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ReflectorPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ReflectorPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ReflectorPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ReflectorType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ReflectorType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ReflectorType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ReflectorType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ReflectorType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RunwayCentrelinePointPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RunwayCentrelinePointPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.RunwayCentrelinePointPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_RunwayCentrelinePointPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_RunwayCentrelinePointPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RunwayDirectionPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RunwayDirectionPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.RunwayDirectionPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_RunwayDirectionPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_RunwayDirectionPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RunwayPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RunwayPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.RunwayPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_RunwayPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_RunwayPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SDFTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SDFTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SDFTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SDFTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SDFTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SDFTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SDFTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SDFTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SDFTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SDFTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SDFType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SDFType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SDFType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SDFType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SDFType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SecondarySurveillanceRadarTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SecondarySurveillanceRadarTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SecondarySurveillanceRadarTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SecondarySurveillanceRadarTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SecondarySurveillanceRadarTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SecondarySurveillanceRadarTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SecondarySurveillanceRadarTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SecondarySurveillanceRadarTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SecondarySurveillanceRadarTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SecondarySurveillanceRadarTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SecondarySurveillanceRadarType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SecondarySurveillanceRadarType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SecondarySurveillanceRadarType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SecondarySurveillanceRadarType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SecondarySurveillanceRadarType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SignificantPointInAirspaceTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SignificantPointInAirspaceTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SignificantPointInAirspaceTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SignificantPointInAirspaceTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SignificantPointInAirspaceTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SignificantPointInAirspaceTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SignificantPointInAirspaceTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SignificantPointInAirspaceTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SignificantPointInAirspaceTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SignificantPointInAirspaceTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SignificantPointInAirspaceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SignificantPointInAirspaceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SignificantPointInAirspaceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SignificantPointInAirspaceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SignificantPointInAirspaceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SpecialNavigationStationStatusPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SpecialNavigationStationStatusPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SpecialNavigationStationStatusPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SpecialNavigationStationStatusPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SpecialNavigationStationStatusPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SpecialNavigationStationStatusType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SpecialNavigationStationStatusType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SpecialNavigationStationStatusType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SpecialNavigationStationStatusType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SpecialNavigationStationStatusType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SpecialNavigationStationTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SpecialNavigationStationTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SpecialNavigationStationTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SpecialNavigationStationTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SpecialNavigationStationTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SpecialNavigationStationTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SpecialNavigationStationTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SpecialNavigationStationTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SpecialNavigationStationTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SpecialNavigationStationTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SpecialNavigationStationType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SpecialNavigationStationType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SpecialNavigationStationType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SpecialNavigationStationType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SpecialNavigationStationType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SpecialNavigationSystemPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SpecialNavigationSystemPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SpecialNavigationSystemPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SpecialNavigationSystemPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SpecialNavigationSystemPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SpecialNavigationSystemTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SpecialNavigationSystemTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SpecialNavigationSystemTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SpecialNavigationSystemTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SpecialNavigationSystemTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SpecialNavigationSystemTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SpecialNavigationSystemTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SpecialNavigationSystemTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SpecialNavigationSystemTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SpecialNavigationSystemTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SpecialNavigationSystemType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SpecialNavigationSystemType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SpecialNavigationSystemType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SpecialNavigationSystemType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SpecialNavigationSystemType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SurveillanceGroundStationPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SurveillanceGroundStationPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SurveillanceGroundStationPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SurveillanceGroundStationPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SurveillanceGroundStationPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SurveillanceGroundStationType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SurveillanceGroundStationType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SurveillanceGroundStationType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SurveillanceGroundStationType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SurveillanceGroundStationType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TACANTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TACANTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.TACANTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_TACANTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_TACANTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TACANTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TACANTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.TACANTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_TACANTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_TACANTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TACANType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TACANType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.TACANType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_TACANType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_TACANType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TelephoneContactPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TelephoneContactPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.TelephoneContactPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_TelephoneContactPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_TelephoneContactPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TelephoneContactType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TelephoneContactType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.TelephoneContactType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_TelephoneContactType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_TelephoneContactType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TimesheetPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TimesheetPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.TimesheetPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_TimesheetPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_TimesheetPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TimesheetType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TimesheetType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.TimesheetType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_TimesheetType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_TimesheetType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TouchDownLiftOffPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TouchDownLiftOffPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.TouchDownLiftOffPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_TouchDownLiftOffPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_TouchDownLiftOffPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_UnitPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New UnitPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.UnitPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_UnitPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_UnitPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_VORPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New VORPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.VORPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_VORPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_VORPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_VORTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New VORTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.VORTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_VORTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_VORTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_VORTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New VORTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.VORTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_VORTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_VORTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_VORType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New VORType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.VORType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_VORType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_VORType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeAirspacePointPositionType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeAirspacePointPositionType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeAirspacePointPositionType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeAirspacePointPositionType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeAirspacePointPositionType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeAirspacePointRoleType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeAirspacePointRoleType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeAirspacePointRoleType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeAirspacePointRoleType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeAirspacePointRoleType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeAltitudeUseType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeAltitudeUseType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeAltitudeUseType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeAltitudeUseType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeAltitudeUseType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeAuralMorseType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeAuralMorseType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeAuralMorseType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeAuralMorseType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeAuralMorseType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeAuthorityRoleType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeAuthorityRoleType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeAuthorityRoleType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeAuthorityRoleType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeAuthorityRoleType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeBearingType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeBearingType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeBearingType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeBearingType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeBearingType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeCardinalDirectionType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeCardinalDirectionType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeCardinalDirectionType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeCardinalDirectionType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeCardinalDirectionType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeCheckpointCategoryType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeCheckpointCategoryType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeCheckpointCategoryType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeCheckpointCategoryType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeCheckpointCategoryType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeCourseQualityILSType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeCourseQualityILSType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeCourseQualityILSType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeCourseQualityILSType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeCourseQualityILSType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeDMEChannelType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeDMEChannelType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeDMEChannelType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeDMEChannelType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeDMEChannelType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeDMEType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeDMEType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeDMEType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeDMEType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeDMEType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeDayType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeDayType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeDayType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeDayType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeDayType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeDesignatedPointDesignatorType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeDesignatedPointDesignatorType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeDesignatedPointDesignatorType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeDesignatedPointDesignatorType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeDesignatedPointDesignatorType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeDesignatedPointType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeDesignatedPointType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeDesignatedPointType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeDesignatedPointType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeDesignatedPointType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeDirectionReferenceType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeDirectionReferenceType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeDirectionReferenceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeDirectionReferenceType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeDirectionReferenceType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeDistanceIndicationType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeDistanceIndicationType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeDistanceIndicationType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeDistanceIndicationType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeDistanceIndicationType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeEmissionBandType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeEmissionBandType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeEmissionBandType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeEmissionBandType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeEmissionBandType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeILSBackCourseType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeILSBackCourseType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeILSBackCourseType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeILSBackCourseType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeILSBackCourseType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeIntegrityLevelILSType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeIntegrityLevelILSType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeIntegrityLevelILSType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeIntegrityLevelILSType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeIntegrityLevelILSType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeMLSAzimuthType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeMLSAzimuthType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeMLSAzimuthType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeMLSAzimuthType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeMLSAzimuthType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeMLSChannelType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeMLSChannelType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeMLSChannelType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeMLSChannelType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeMLSChannelType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeMarkerBeaconSignalType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeMarkerBeaconSignalType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeMarkerBeaconSignalType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeMarkerBeaconSignalType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeMarkerBeaconSignalType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeNDBUsageType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeNDBUsageType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeNDBUsageType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeNDBUsageType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeNDBUsageType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeNavaidDesignatorType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeNavaidDesignatorType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeNavaidDesignatorType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeNavaidDesignatorType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeNavaidDesignatorType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeNavaidPurposeType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeNavaidPurposeType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeNavaidPurposeType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeNavaidPurposeType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeNavaidPurposeType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeNavaidServiceType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeNavaidServiceType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeNavaidServiceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeNavaidServiceType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeNavaidServiceType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeNorthReferenceType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeNorthReferenceType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeNorthReferenceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeNorthReferenceType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeNorthReferenceType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeNotePurposeType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeNotePurposeType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeNotePurposeType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeNotePurposeType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeNotePurposeType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodePARType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodePARType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodePARType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodePARType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodePARType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodePositionInILSType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodePositionInILSType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodePositionInILSType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodePositionInILSType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodePositionInILSType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodePrimaryRadarType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodePrimaryRadarType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodePrimaryRadarType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodePrimaryRadarType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodePrimaryRadarType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeRadarServiceType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeRadarServiceType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeRadarServiceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeRadarServiceType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeRadarServiceType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeRadioEmissionType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeRadioEmissionType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeRadioEmissionType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeRadioEmissionType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeRadioEmissionType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeRadioSignalType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeRadioSignalType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeRadioSignalType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeRadioSignalType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeRadioSignalType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeReflectorType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeReflectorType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeReflectorType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeReflectorType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeReflectorType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeSignalPerformanceILSType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeSignalPerformanceILSType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeSignalPerformanceILSType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeSignalPerformanceILSType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeSignalPerformanceILSType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeSpecialNavigationChainDesignatorType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeSpecialNavigationChainDesignatorType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeSpecialNavigationChainDesignatorType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeSpecialNavigationChainDesignatorType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeSpecialNavigationChainDesignatorType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeSpecialNavigationStationType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeSpecialNavigationStationType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeSpecialNavigationStationType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeSpecialNavigationStationType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeSpecialNavigationStationType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeSpecialNavigationSystemType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeSpecialNavigationSystemType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeSpecialNavigationSystemType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeSpecialNavigationSystemType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeSpecialNavigationSystemType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeStandbyPowerType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeStandbyPowerType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeStandbyPowerType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeStandbyPowerType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeStandbyPowerType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeStatusNavaidType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeStatusNavaidType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeStatusNavaidType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeStatusNavaidType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeStatusNavaidType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeTACANChannelType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeTACANChannelType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeTACANChannelType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeTACANChannelType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeTACANChannelType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeTelecomNetworkType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeTelecomNetworkType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeTelecomNetworkType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeTelecomNetworkType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeTelecomNetworkType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeTimeEventCombinationType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeTimeEventCombinationType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeTimeEventCombinationType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeTimeEventCombinationType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeTimeEventCombinationType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeTimeEventType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeTimeEventType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeTimeEventType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeTimeEventType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeTimeEventType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeTimeReferenceType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeTimeReferenceType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeTimeReferenceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeTimeReferenceType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeTimeReferenceType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeTransponderType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeTransponderType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeTransponderType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeTransponderType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeTransponderType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeVORType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeVORType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeVORType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeVORType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeVORType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeVerticalDatumType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeVerticalDatumType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeVerticalDatumType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeVerticalDatumType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeVerticalDatumType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeVerticalReferenceType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeVerticalReferenceType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeVerticalReferenceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeVerticalReferenceType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeVerticalReferenceType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeYesNoType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeYesNoType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeYesNoType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeYesNoType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeYesNoType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DateMonthDayType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DateMonthDayType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.DateMonthDayType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_DateMonthDayType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_DateMonthDayType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DateYearType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DateYearType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.DateYearType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_DateYearType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_DateYearType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NoSequenceType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New NoSequenceType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.NoSequenceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_NoSequenceType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_NoSequenceType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TextAddressType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TextAddressType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.TextAddressType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_TextAddressType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_TextAddressType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TextDesignatorType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TextDesignatorType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.TextDesignatorType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_TextDesignatorType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_TextDesignatorType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TextNameType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TextNameType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.TextNameType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_TextNameType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_TextNameType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TextNoteType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TextNoteType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.TextNoteType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_TextNoteType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_TextNoteType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TextPhoneType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TextPhoneType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.TextPhoneType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_TextPhoneType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_TextPhoneType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TextPropertyNameType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TextPropertyNameType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.TextPropertyNameType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_TextPropertyNameType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_TextPropertyNameType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TimeType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TimeType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.TimeType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_TimeType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_TimeType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ValAngleType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ValAngleType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.ValAngleType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_ValAngleType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_ValAngleType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ValBearingType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ValBearingType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.ValBearingType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_ValBearingType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_ValBearingType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ValDistanceSignedType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ValDistanceSignedType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.ValDistanceSignedType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_ValDistanceSignedType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_ValDistanceSignedType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ValDistanceType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ValDistanceType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.ValDistanceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_ValDistanceType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_ValDistanceType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ValDistanceVerticalType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ValDistanceVerticalType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.ValDistanceVerticalType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_ValDistanceVerticalType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_ValDistanceVerticalType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ValDurationType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ValDurationType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.ValDurationType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_ValDurationType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_ValDurationType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ValFrequencyType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ValFrequencyType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.ValFrequencyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_ValFrequencyType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_ValFrequencyType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ValMagneticVariationType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ValMagneticVariationType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.ValMagneticVariationType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_ValMagneticVariationType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_ValMagneticVariationType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractAIXMFeatureBaseType_AixmabstractgmlobjecttypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractAIXMFeatureBaseType",
	}).Stage(probe.formStage)
	inst := new(aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Lock()
			defer probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
			saveStageSet_AbstractAIXMFeatureBaseType_AixmabstractgmlobjecttypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Commit()
			updateStageSetTable_AbstractAIXMFeatureBaseType_AixmabstractgmlobjecttypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractAIXMFeatureType_AixmabstractgmlobjecttypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractAIXMFeatureType",
	}).Stage(probe.formStage)
	inst := new(aixmabstractgmlobjecttypes.AbstractAIXMFeatureType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Lock()
			defer probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
			saveStageSet_AbstractAIXMFeatureType_AixmabstractgmlobjecttypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Commit()
			updateStageSetTable_AbstractAIXMFeatureType_AixmabstractgmlobjecttypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractAIXMMessageBaseType_AixmabstractgmlobjecttypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractAIXMMessageBaseType",
	}).Stage(probe.formStage)
	inst := new(aixmabstractgmlobjecttypes.AbstractAIXMMessageBaseType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Lock()
			defer probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
			saveStageSet_AbstractAIXMMessageBaseType_AixmabstractgmlobjecttypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Commit()
			updateStageSetTable_AbstractAIXMMessageBaseType_AixmabstractgmlobjecttypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractAIXMMessageType_AixmabstractgmlobjecttypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractAIXMMessageType",
	}).Stage(probe.formStage)
	inst := new(aixmabstractgmlobjecttypes.AbstractAIXMMessageType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Lock()
			defer probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
			saveStageSet_AbstractAIXMMessageType_AixmabstractgmlobjecttypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Commit()
			updateStageSetTable_AbstractAIXMMessageType_AixmabstractgmlobjecttypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractAIXMObjectType_AixmabstractgmlobjecttypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractAIXMObjectType",
	}).Stage(probe.formStage)
	inst := new(aixmabstractgmlobjecttypes.AbstractAIXMObjectType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Lock()
			defer probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
			saveStageSet_AbstractAIXMObjectType_AixmabstractgmlobjecttypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Commit()
			updateStageSetTable_AbstractAIXMObjectType_AixmabstractgmlobjecttypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractAIXMPropertyType_AixmabstractgmlobjecttypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractAIXMPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmabstractgmlobjecttypes.AbstractAIXMPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Lock()
			defer probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
			saveStageSet_AbstractAIXMPropertyType_AixmabstractgmlobjecttypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Commit()
			updateStageSetTable_AbstractAIXMPropertyType_AixmabstractgmlobjecttypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractAIXMTimeSliceBaseType_AixmabstractgmlobjecttypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractAIXMTimeSliceBaseType",
	}).Stage(probe.formStage)
	inst := new(aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceBaseType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Lock()
			defer probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
			saveStageSet_AbstractAIXMTimeSliceBaseType_AixmabstractgmlobjecttypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Commit()
			updateStageSetTable_AbstractAIXMTimeSliceBaseType_AixmabstractgmlobjecttypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractAIXMTimeSliceType_AixmabstractgmlobjecttypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractAIXMTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Lock()
			defer probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
			saveStageSet_AbstractAIXMTimeSliceType_AixmabstractgmlobjecttypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Commit()
			updateStageSetTable_AbstractAIXMTimeSliceType_AixmabstractgmlobjecttypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractExtensionType_AixmabstractgmlobjecttypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractExtensionType",
	}).Stage(probe.formStage)
	inst := new(aixmabstractgmlobjecttypes.AbstractExtensionType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Lock()
			defer probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
			saveStageSet_AbstractExtensionType_AixmabstractgmlobjecttypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Commit()
			updateStageSetTable_AbstractExtensionType_AixmabstractgmlobjecttypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FeatureMetadataPropertyType_AixmabstractgmlobjecttypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New FeatureMetadataPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmabstractgmlobjecttypes.FeatureMetadataPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Lock()
			defer probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
			saveStageSet_FeatureMetadataPropertyType_AixmabstractgmlobjecttypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Commit()
			updateStageSetTable_FeatureMetadataPropertyType_AixmabstractgmlobjecttypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FeatureTimeSliceMetadataPropertyType_AixmabstractgmlobjecttypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New FeatureTimeSliceMetadataPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Lock()
			defer probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
			saveStageSet_FeatureTimeSliceMetadataPropertyType_AixmabstractgmlobjecttypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Commit()
			updateStageSetTable_FeatureTimeSliceMetadataPropertyType_AixmabstractgmlobjecttypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_MessageMetadataPropertyType_AixmabstractgmlobjecttypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New MessageMetadataPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmabstractgmlobjecttypes.MessageMetadataPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Lock()
			defer probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmabstractgmlobjecttypesStage)
			saveStageSet_MessageMetadataPropertyType_AixmabstractgmlobjecttypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmabstractgmlobjecttypesStage.Commit()
			updateStageSetTable_MessageMetadataPropertyType_AixmabstractgmlobjecttypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractFeatureMemberType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractFeatureMemberType",
	}).Stage(probe.formStage)
	inst := new(gml.AbstractFeatureMemberType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_AbstractFeatureMemberType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_AbstractFeatureMemberType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractFeatureType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractFeatureType",
	}).Stage(probe.formStage)
	inst := new(gml.AbstractFeatureType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_AbstractFeatureType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_AbstractFeatureType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractGMLType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractGMLType",
	}).Stage(probe.formStage)
	inst := new(gml.AbstractGMLType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_AbstractGMLType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_AbstractGMLType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractGeometricPrimitiveType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractGeometricPrimitiveType",
	}).Stage(probe.formStage)
	inst := new(gml.AbstractGeometricPrimitiveType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_AbstractGeometricPrimitiveType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_AbstractGeometricPrimitiveType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractGeometryType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractGeometryType",
	}).Stage(probe.formStage)
	inst := new(gml.AbstractGeometryType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_AbstractGeometryType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_AbstractGeometryType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractMetadataPropertyType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractMetadataPropertyType",
	}).Stage(probe.formStage)
	inst := new(gml.AbstractMetadataPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_AbstractMetadataPropertyType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_AbstractMetadataPropertyType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractTimeGeometricPrimitiveType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractTimeGeometricPrimitiveType",
	}).Stage(probe.formStage)
	inst := new(gml.AbstractTimeGeometricPrimitiveType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_AbstractTimeGeometricPrimitiveType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_AbstractTimeGeometricPrimitiveType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractTimeObjectType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractTimeObjectType",
	}).Stage(probe.formStage)
	inst := new(gml.AbstractTimeObjectType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_AbstractTimeObjectType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_AbstractTimeObjectType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractTimePrimitiveType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractTimePrimitiveType",
	}).Stage(probe.formStage)
	inst := new(gml.AbstractTimePrimitiveType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_AbstractTimePrimitiveType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_AbstractTimePrimitiveType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractTimeSliceType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(gml.AbstractTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_AbstractTimeSliceType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_AbstractTimeSliceType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_BoundingShapeType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New BoundingShapeType",
	}).Stage(probe.formStage)
	inst := new(gml.BoundingShapeType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_BoundingShapeType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_BoundingShapeType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeType",
	}).Stage(probe.formStage)
	inst := new(gml.CodeType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_CodeType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_CodeType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeWithAuthorityType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeWithAuthorityType",
	}).Stage(probe.formStage)
	inst := new(gml.CodeWithAuthorityType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_CodeWithAuthorityType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_CodeWithAuthorityType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DirectPositionType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DirectPositionType",
	}).Stage(probe.formStage)
	inst := new(gml.DirectPositionType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_DirectPositionType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_DirectPositionType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DynamicFeatureType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DynamicFeatureType",
	}).Stage(probe.formStage)
	inst := new(gml.DynamicFeatureType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_DynamicFeatureType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_DynamicFeatureType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_EnvelopeType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New EnvelopeType",
	}).Stage(probe.formStage)
	inst := new(gml.EnvelopeType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_EnvelopeType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_EnvelopeType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_PointType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New PointType",
	}).Stage(probe.formStage)
	inst := new(gml.PointType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_PointType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_PointType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TimeInstantType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TimeInstantType",
	}).Stage(probe.formStage)
	inst := new(gml.TimeInstantType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_TimeInstantType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_TimeInstantType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TimePeriodType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TimePeriodType",
	}).Stage(probe.formStage)
	inst := new(gml.TimePeriodType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_TimePeriodType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_TimePeriodType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TimePositionType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TimePositionType",
	}).Stage(probe.formStage)
	inst := new(gml.TimePositionType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_TimePositionType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_TimePositionType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TimePrimitivePropertyType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TimePrimitivePropertyType",
	}).Stage(probe.formStage)
	inst := new(gml.TimePrimitivePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_TimePrimitivePropertyType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_TimePrimitivePropertyType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Xlink_SchemaElement_XlinkStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Xlink_SchemaElement",
	}).Stage(probe.formStage)
	inst := new(xlink.Xlink_SchemaElement)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.XlinkStage.Lock()
			defer probe.stageOfInterest.XlinkStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.XlinkStage)
			saveStageSet_Xlink_SchemaElement_XlinkStage(inst, probe, formGroup)
			probe.stageOfInterest.XlinkStage.Commit()
			updateStageSetTable_Xlink_SchemaElement_XlinkStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

