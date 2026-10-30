// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	gongprobe "github.com/fullstack-lang/gong/pkg/runtime/probe"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/xlink"
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
		case "StandardInstrumentDeparture":
			StageSetFormDivSelectFieldToField(&inst.StandardInstrumentDeparture, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.StandardInstrumentDepartureType](), formDiv)
		case "StandardInstrumentArrival":
			StageSetFormDivSelectFieldToField(&inst.StandardInstrumentArrival, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.StandardInstrumentArrivalType](), formDiv)
		case "InstrumentApproachProcedure":
			StageSetFormDivSelectFieldToField(&inst.InstrumentApproachProcedure, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.InstrumentApproachProcedureType](), formDiv)
		case "NavigationArea":
			StageSetFormDivSelectFieldToField(&inst.NavigationArea, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavigationAreaType](), formDiv)
		case "TerminalArrivalArea":
			StageSetFormDivSelectFieldToField(&inst.TerminalArrivalArea, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TerminalArrivalAreaType](), formDiv)
		case "CirclingArea":
			StageSetFormDivSelectFieldToField(&inst.CirclingArea, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CirclingAreaType](), formDiv)
		case "DepartureLeg":
			StageSetFormDivSelectFieldToField(&inst.DepartureLeg, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DepartureLegType](), formDiv)
		case "ArrivalLeg":
			StageSetFormDivSelectFieldToField(&inst.ArrivalLeg, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ArrivalLegType](), formDiv)
		case "ArrivalFeederLeg":
			StageSetFormDivSelectFieldToField(&inst.ArrivalFeederLeg, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ArrivalFeederLegType](), formDiv)
		case "InitialLeg":
			StageSetFormDivSelectFieldToField(&inst.InitialLeg, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.InitialLegType](), formDiv)
		case "IntermediateLeg":
			StageSetFormDivSelectFieldToField(&inst.IntermediateLeg, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.IntermediateLegType](), formDiv)
		case "FinalLeg":
			StageSetFormDivSelectFieldToField(&inst.FinalLeg, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FinalLegType](), formDiv)
		case "MissedApproachLeg":
			StageSetFormDivSelectFieldToField(&inst.MissedApproachLeg, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.MissedApproachLegType](), formDiv)
		case "ProcedureDME":
			StageSetFormDivSelectFieldToField(&inst.ProcedureDME, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ProcedureDMEType](), formDiv)
		case "NavigationAreaRestriction":
			StageSetFormDivSelectFieldToField(&inst.NavigationAreaRestriction, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavigationAreaRestrictionType](), formDiv)
		case "SafeAltitudeArea":
			StageSetFormDivSelectFieldToField(&inst.SafeAltitudeArea, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SafeAltitudeAreaType](), formDiv)
		case "HoldingPattern":
			StageSetFormDivSelectFieldToField(&inst.HoldingPattern, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.HoldingPatternType](), formDiv)
		case "UnplannedHolding":
			StageSetFormDivSelectFieldToField(&inst.UnplannedHolding, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.UnplannedHoldingType](), formDiv)
		case "HoldingAssessment":
			StageSetFormDivSelectFieldToField(&inst.HoldingAssessment, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.HoldingAssessmentType](), formDiv)
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

func saveStageSet_A_extension_42_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_42,
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

func saveStageSet_A_extension_43_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_43,
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

func saveStageSet_A_extension_44_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_44,
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

func saveStageSet_A_extension_45_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_45,
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

func saveStageSet_A_extension_46_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_46,
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

func saveStageSet_A_extension_47_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_47,
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

func saveStageSet_A_extension_48_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_48,
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

func saveStageSet_A_extension_49_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_49,
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

func saveStageSet_A_extension_50_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_50,
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

func saveStageSet_A_extension_51_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_51,
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

func saveStageSet_A_extension_52_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_52,
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

func saveStageSet_A_extension_53_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_53,
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

func saveStageSet_A_extension_54_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_54,
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

func saveStageSet_A_extension_55_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_55,
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

func saveStageSet_A_extension_56_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_56,
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

func saveStageSet_A_extension_57_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_57,
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

func saveStageSet_A_extension_58_AixmfeaturesStage(
	inst *aixmfeatures.A_extension_58,
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

func saveStageSet_AbstractApproachLegType_AixmfeaturesStage(
	inst *aixmfeatures.AbstractApproachLegType,
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

func saveStageSet_AbstractProcedureType_AixmfeaturesStage(
	inst *aixmfeatures.AbstractProcedureType,
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

func saveStageSet_AbstractSegmentLegType_AixmfeaturesStage(
	inst *aixmfeatures.AbstractSegmentLegType,
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

func saveStageSet_AbstractSegmentPointType_AixmfeaturesStage(
	inst *aixmfeatures.AbstractSegmentPointType,
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

func saveStageSet_AerialRefuellingPointType_AixmfeaturesStage(
	inst *aixmfeatures.AerialRefuellingPointType,
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
		case "ReportingATC":
			StageSetFormDivSelectFieldToField(&inst.ReportingATC, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeATCReportingType](), formDiv)
		case "FlyOver":
			StageSetFormDivSelectFieldToField(&inst.FlyOver, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "Waypoint":
			StageSetFormDivSelectFieldToField(&inst.Waypoint, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "RadarGuidance":
			StageSetFormDivSelectFieldToField(&inst.RadarGuidance, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
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
		case "ExtendedServiceVolume":
			StageSetFormDivSelectFieldToField(&inst.ExtendedServiceVolume, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RadioFrequencyAreaPropertyType](), formDiv)
		case "Sequence":
			StageSetFormDivSelectFieldToField(&inst.Sequence, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoSequenceType](), formDiv)
		case "UsageType":
			StageSetFormDivSelectFieldToField(&inst.UsageType, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAerialRefuellingPointType](), formDiv)
		}
	}
}

func saveStageSet_AircraftCharacteristicPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AircraftCharacteristicPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "AircraftCharacteristic":
			StageSetFormDivSelectFieldToField(&inst.AircraftCharacteristic, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AircraftCharacteristicType](), formDiv)
		}
	}
}

func saveStageSet_AircraftCharacteristicType_AixmfeaturesStage(
	inst *aixmfeatures.AircraftCharacteristicType,
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
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAircraftType](), formDiv)
		case "Engine":
			StageSetFormDivSelectFieldToField(&inst.Engine, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAircraftEngineType](), formDiv)
		case "NumberEngine":
			StageSetFormDivSelectFieldToField(&inst.NumberEngine, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAircraftEngineNumberType](), formDiv)
		case "TypeAircraftICAO":
			StageSetFormDivSelectFieldToField(&inst.TypeAircraftICAO, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAircraftICAOType](), formDiv)
		case "AircraftLandingCategory":
			StageSetFormDivSelectFieldToField(&inst.AircraftLandingCategory, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAircraftCategoryType](), formDiv)
		case "WingSpan":
			StageSetFormDivSelectFieldToField(&inst.WingSpan, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "WingSpanInterpretation":
			StageSetFormDivSelectFieldToField(&inst.WingSpanInterpretation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeValueInterpretationType](), formDiv)
		case "ClassWingSpan":
			StageSetFormDivSelectFieldToField(&inst.ClassWingSpan, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAircraftWingspanClassType](), formDiv)
		case "Weight":
			StageSetFormDivSelectFieldToField(&inst.Weight, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValWeightType](), formDiv)
		case "WeightInterpretation":
			StageSetFormDivSelectFieldToField(&inst.WeightInterpretation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeValueInterpretationType](), formDiv)
		case "Passengers":
			StageSetFormDivSelectFieldToField(&inst.Passengers, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoNumberType](), formDiv)
		case "PassengersInterpretation":
			StageSetFormDivSelectFieldToField(&inst.PassengersInterpretation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeValueInterpretationType](), formDiv)
		case "Speed":
			StageSetFormDivSelectFieldToField(&inst.Speed, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValSpeedType](), formDiv)
		case "SpeedInterpretation":
			StageSetFormDivSelectFieldToField(&inst.SpeedInterpretation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeValueInterpretationType](), formDiv)
		case "WakeTurbulence":
			StageSetFormDivSelectFieldToField(&inst.WakeTurbulence, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeWakeTurbulenceType](), formDiv)
		case "NavigationEquipment":
			StageSetFormDivSelectFieldToField(&inst.NavigationEquipment, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavigationEquipmentType](), formDiv)
		case "NavigationSpecification":
			StageSetFormDivSelectFieldToField(&inst.NavigationSpecification, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavigationSpecificationType](), formDiv)
		case "VerticalSeparationCapability":
			StageSetFormDivSelectFieldToField(&inst.VerticalSeparationCapability, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRVSMType](), formDiv)
		case "AntiCollisionAndSeparationEquipment":
			StageSetFormDivSelectFieldToField(&inst.AntiCollisionAndSeparationEquipment, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeEquipmentAntiCollisionType](), formDiv)
		case "CommunicationEquipment":
			StageSetFormDivSelectFieldToField(&inst.CommunicationEquipment, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCommunicationModeType](), formDiv)
		case "SurveillanceEquipment":
			StageSetFormDivSelectFieldToField(&inst.SurveillanceEquipment, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTransponderType](), formDiv)
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

func saveStageSet_AltimeterSourcePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AltimeterSourcePropertyType,
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

func saveStageSet_AltitudeAdjustmentPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AltitudeAdjustmentPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "AltitudeAdjustment":
			StageSetFormDivSelectFieldToField(&inst.AltitudeAdjustment, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AltitudeAdjustmentType](), formDiv)
		}
	}
}

func saveStageSet_AltitudeAdjustmentType_AixmfeaturesStage(
	inst *aixmfeatures.AltitudeAdjustmentType,
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
		case "AltitudeAdjustmentType":
			StageSetFormDivSelectFieldToField(&inst.AltitudeAdjustmentType, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAltitudeAdjustmentType](), formDiv)
		case "PrimaryAlternateMinimum":
			StageSetFormDivSelectFieldToField(&inst.PrimaryAlternateMinimum, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "AltitudeAdjustment":
			StageSetFormDivSelectFieldToField(&inst.AltitudeAdjustment, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "LocalRemoteCode":
			StageSetFormDivSelectFieldToField(&inst.LocalRemoteCode, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		}
	}
}

func saveStageSet_AngleIndicationPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AngleIndicationPropertyType,
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

func saveStageSet_AngleUsePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AngleUsePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "AngleUse":
			StageSetFormDivSelectFieldToField(&inst.AngleUse, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AngleUseType](), formDiv)
		}
	}
}

func saveStageSet_AngleUseType_AixmfeaturesStage(
	inst *aixmfeatures.AngleUseType,
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
		case "AlongCourseGuidance":
			StageSetFormDivSelectFieldToField(&inst.AlongCourseGuidance, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "TheAngleIndication":
			StageSetFormDivSelectFieldToField(&inst.TheAngleIndication, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AngleIndicationPropertyType](), formDiv)
		}
	}
}

func saveStageSet_ApproachAltitudeTablePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.ApproachAltitudeTablePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "ApproachAltitudeTable":
			StageSetFormDivSelectFieldToField(&inst.ApproachAltitudeTable, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ApproachAltitudeTableType](), formDiv)
		}
	}
}

func saveStageSet_ApproachAltitudeTableType_AixmfeaturesStage(
	inst *aixmfeatures.ApproachAltitudeTableType,
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
		case "MeasurementPoint":
			StageSetFormDivSelectFieldToField(&inst.MeasurementPoint, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeProcedureDistanceType](), formDiv)
		case "Altitude":
			StageSetFormDivSelectFieldToField(&inst.Altitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "AltitudeReference":
			StageSetFormDivSelectFieldToField(&inst.AltitudeReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		}
	}
}

func saveStageSet_ApproachConditionPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.ApproachConditionPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "ApproachCondition":
			StageSetFormDivSelectFieldToField(&inst.ApproachCondition, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ApproachConditionType](), formDiv)
		}
	}
}

func saveStageSet_ApproachConditionType_AixmfeaturesStage(
	inst *aixmfeatures.ApproachConditionType,
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
		case "FinalApproachPath":
			StageSetFormDivSelectFieldToField(&inst.FinalApproachPath, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeMinimaFinalApproachPathType](), formDiv)
		case "RequiredNavigationPerformance":
			StageSetFormDivSelectFieldToField(&inst.RequiredNavigationPerformance, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRNPType](), formDiv)
		case "ClimbGradient":
			StageSetFormDivSelectFieldToField(&inst.ClimbGradient, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValSlopeType](), formDiv)
		case "MinimumSet":
			StageSetFormDivSelectFieldToField(&inst.MinimumSet, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.MinimaPropertyType](), formDiv)
		case "Altimeter":
			StageSetFormDivSelectFieldToField(&inst.Altimeter, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AltimeterSourcePropertyType](), formDiv)
		}
	}
}

func saveStageSet_ApproachDistanceTablePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.ApproachDistanceTablePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "ApproachDistanceTable":
			StageSetFormDivSelectFieldToField(&inst.ApproachDistanceTable, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ApproachDistanceTableType](), formDiv)
		}
	}
}

func saveStageSet_ApproachDistanceTableType_AixmfeaturesStage(
	inst *aixmfeatures.ApproachDistanceTableType,
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
		case "StartingMeasurementPoint":
			StageSetFormDivSelectFieldToField(&inst.StartingMeasurementPoint, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeProcedureDistanceType](), formDiv)
		case "ValueHAT":
			StageSetFormDivSelectFieldToField(&inst.ValueHAT, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "EndingMeasurementPoint":
			StageSetFormDivSelectFieldToField(&inst.EndingMeasurementPoint, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeProcedureDistanceType](), formDiv)
		case "Distance":
			StageSetFormDivSelectFieldToField(&inst.Distance, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		}
	}
}

func saveStageSet_ApproachTimingTablePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.ApproachTimingTablePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "ApproachTimingTable":
			StageSetFormDivSelectFieldToField(&inst.ApproachTimingTable, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ApproachTimingTableType](), formDiv)
		}
	}
}

func saveStageSet_ApproachTimingTableType_AixmfeaturesStage(
	inst *aixmfeatures.ApproachTimingTableType,
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
		case "StartingMeasurementPoint":
			StageSetFormDivSelectFieldToField(&inst.StartingMeasurementPoint, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeProcedureDistanceType](), formDiv)
		case "EndingMeasurementPoint":
			StageSetFormDivSelectFieldToField(&inst.EndingMeasurementPoint, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeProcedureDistanceType](), formDiv)
		case "Time":
			StageSetFormDivSelectFieldToField(&inst.Time, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDurationType](), formDiv)
		case "Speed":
			StageSetFormDivSelectFieldToField(&inst.Speed, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValSpeedType](), formDiv)
		}
	}
}

func saveStageSet_ArrivalFeederLegTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.ArrivalFeederLegTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "ArrivalFeederLegTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.ArrivalFeederLegTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ArrivalFeederLegTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_ArrivalFeederLegTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.ArrivalFeederLegTimeSliceType,
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
		case "EndConditionDesignator":
			StageSetFormDivSelectFieldToField(&inst.EndConditionDesignator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSegmentTerminationType](), formDiv)
		case "LegPath":
			StageSetFormDivSelectFieldToField(&inst.LegPath, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTrajectoryType](), formDiv)
		case "LegTypeARINC":
			StageSetFormDivSelectFieldToField(&inst.LegTypeARINC, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSegmentPathType](), formDiv)
		case "Course":
			StageSetFormDivSelectFieldToField(&inst.Course, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "CourseType":
			StageSetFormDivSelectFieldToField(&inst.CourseType, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCourseType](), formDiv)
		case "CourseDirection":
			StageSetFormDivSelectFieldToField(&inst.CourseDirection, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDirectionReferenceType](), formDiv)
		case "TurnDirection":
			StageSetFormDivSelectFieldToField(&inst.TurnDirection, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDirectionTurnType](), formDiv)
		case "SpeedLimit":
			StageSetFormDivSelectFieldToField(&inst.SpeedLimit, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValSpeedType](), formDiv)
		case "SpeedReference":
			StageSetFormDivSelectFieldToField(&inst.SpeedReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSpeedReferenceType](), formDiv)
		case "SpeedInterpretation":
			StageSetFormDivSelectFieldToField(&inst.SpeedInterpretation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAltitudeUseType](), formDiv)
		case "BankAngle":
			StageSetFormDivSelectFieldToField(&inst.BankAngle, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "Length":
			StageSetFormDivSelectFieldToField(&inst.Length, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "Duration":
			StageSetFormDivSelectFieldToField(&inst.Duration, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDurationType](), formDiv)
		case "ProcedureTurnRequired":
			StageSetFormDivSelectFieldToField(&inst.ProcedureTurnRequired, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "UpperLimitAltitude":
			StageSetFormDivSelectFieldToField(&inst.UpperLimitAltitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "UpperLimitReference":
			StageSetFormDivSelectFieldToField(&inst.UpperLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "LowerLimitAltitude":
			StageSetFormDivSelectFieldToField(&inst.LowerLimitAltitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "LowerLimitReference":
			StageSetFormDivSelectFieldToField(&inst.LowerLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "AltitudeInterpretation":
			StageSetFormDivSelectFieldToField(&inst.AltitudeInterpretation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAltitudeUseType](), formDiv)
		case "AltitudeOverrideATC":
			StageSetFormDivSelectFieldToField(&inst.AltitudeOverrideATC, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "AltitudeOverrideReference":
			StageSetFormDivSelectFieldToField(&inst.AltitudeOverrideReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "VerticalAngle":
			StageSetFormDivSelectFieldToField(&inst.VerticalAngle, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "StartPoint":
			StageSetFormDivSelectFieldToField(&inst.StartPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TerminalSegmentPointPropertyType](), formDiv)
		case "EndPoint":
			StageSetFormDivSelectFieldToField(&inst.EndPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TerminalSegmentPointPropertyType](), formDiv)
		case "Trajectory":
			StageSetFormDivSelectFieldToField(&inst.Trajectory, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CurvePropertyType](), formDiv)
		case "ArcCentre":
			StageSetFormDivSelectFieldToField(&inst.ArcCentre, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TerminalSegmentPointPropertyType](), formDiv)
		case "Angle":
			StageSetFormDivSelectFieldToField(&inst.Angle, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AngleIndicationPropertyType](), formDiv)
		case "Distance":
			StageSetFormDivSelectFieldToField(&inst.Distance, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DistanceIndicationPropertyType](), formDiv)
		case "Holding":
			StageSetFormDivSelectFieldToField(&inst.Holding, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.HoldingUsePropertyType](), formDiv)
		case "Approach":
			StageSetFormDivSelectFieldToField(&inst.Approach, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.InstrumentApproachProcedurePropertyType](), formDiv)
		case "RequiredNavigationPerformance":
			StageSetFormDivSelectFieldToField(&inst.RequiredNavigationPerformance, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRNPType](), formDiv)
		}
	}
}

func saveStageSet_ArrivalFeederLegType_AixmfeaturesStage(
	inst *aixmfeatures.ArrivalFeederLegType,
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

func saveStageSet_ArrivalLegTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.ArrivalLegTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "ArrivalLegTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.ArrivalLegTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ArrivalLegTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_ArrivalLegTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.ArrivalLegTimeSliceType,
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
		case "EndConditionDesignator":
			StageSetFormDivSelectFieldToField(&inst.EndConditionDesignator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSegmentTerminationType](), formDiv)
		case "LegPath":
			StageSetFormDivSelectFieldToField(&inst.LegPath, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTrajectoryType](), formDiv)
		case "LegTypeARINC":
			StageSetFormDivSelectFieldToField(&inst.LegTypeARINC, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSegmentPathType](), formDiv)
		case "Course":
			StageSetFormDivSelectFieldToField(&inst.Course, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "CourseType":
			StageSetFormDivSelectFieldToField(&inst.CourseType, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCourseType](), formDiv)
		case "CourseDirection":
			StageSetFormDivSelectFieldToField(&inst.CourseDirection, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDirectionReferenceType](), formDiv)
		case "TurnDirection":
			StageSetFormDivSelectFieldToField(&inst.TurnDirection, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDirectionTurnType](), formDiv)
		case "SpeedLimit":
			StageSetFormDivSelectFieldToField(&inst.SpeedLimit, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValSpeedType](), formDiv)
		case "SpeedReference":
			StageSetFormDivSelectFieldToField(&inst.SpeedReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSpeedReferenceType](), formDiv)
		case "SpeedInterpretation":
			StageSetFormDivSelectFieldToField(&inst.SpeedInterpretation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAltitudeUseType](), formDiv)
		case "BankAngle":
			StageSetFormDivSelectFieldToField(&inst.BankAngle, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "Length":
			StageSetFormDivSelectFieldToField(&inst.Length, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "Duration":
			StageSetFormDivSelectFieldToField(&inst.Duration, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDurationType](), formDiv)
		case "ProcedureTurnRequired":
			StageSetFormDivSelectFieldToField(&inst.ProcedureTurnRequired, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "UpperLimitAltitude":
			StageSetFormDivSelectFieldToField(&inst.UpperLimitAltitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "UpperLimitReference":
			StageSetFormDivSelectFieldToField(&inst.UpperLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "LowerLimitAltitude":
			StageSetFormDivSelectFieldToField(&inst.LowerLimitAltitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "LowerLimitReference":
			StageSetFormDivSelectFieldToField(&inst.LowerLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "AltitudeInterpretation":
			StageSetFormDivSelectFieldToField(&inst.AltitudeInterpretation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAltitudeUseType](), formDiv)
		case "AltitudeOverrideATC":
			StageSetFormDivSelectFieldToField(&inst.AltitudeOverrideATC, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "AltitudeOverrideReference":
			StageSetFormDivSelectFieldToField(&inst.AltitudeOverrideReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "VerticalAngle":
			StageSetFormDivSelectFieldToField(&inst.VerticalAngle, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "StartPoint":
			StageSetFormDivSelectFieldToField(&inst.StartPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TerminalSegmentPointPropertyType](), formDiv)
		case "EndPoint":
			StageSetFormDivSelectFieldToField(&inst.EndPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TerminalSegmentPointPropertyType](), formDiv)
		case "Trajectory":
			StageSetFormDivSelectFieldToField(&inst.Trajectory, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CurvePropertyType](), formDiv)
		case "ArcCentre":
			StageSetFormDivSelectFieldToField(&inst.ArcCentre, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TerminalSegmentPointPropertyType](), formDiv)
		case "Angle":
			StageSetFormDivSelectFieldToField(&inst.Angle, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AngleIndicationPropertyType](), formDiv)
		case "Distance":
			StageSetFormDivSelectFieldToField(&inst.Distance, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DistanceIndicationPropertyType](), formDiv)
		case "Holding":
			StageSetFormDivSelectFieldToField(&inst.Holding, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.HoldingUsePropertyType](), formDiv)
		case "RequiredNavigationPerformance":
			StageSetFormDivSelectFieldToField(&inst.RequiredNavigationPerformance, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRNPType](), formDiv)
		case "Arrival":
			StageSetFormDivSelectFieldToField(&inst.Arrival, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.StandardInstrumentArrivalPropertyType](), formDiv)
		}
	}
}

func saveStageSet_ArrivalLegType_AixmfeaturesStage(
	inst *aixmfeatures.ArrivalLegType,
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

func saveStageSet_CircleSectorPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.CircleSectorPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "CircleSector":
			StageSetFormDivSelectFieldToField(&inst.CircleSector, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CircleSectorType](), formDiv)
		}
	}
}

func saveStageSet_CircleSectorType_AixmfeaturesStage(
	inst *aixmfeatures.CircleSectorType,
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
		case "ArcDirection":
			StageSetFormDivSelectFieldToField(&inst.ArcDirection, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeArcDirectionType](), formDiv)
		case "FromAngle":
			StageSetFormDivSelectFieldToField(&inst.FromAngle, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "ToAngle":
			StageSetFormDivSelectFieldToField(&inst.ToAngle, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "AngleType":
			StageSetFormDivSelectFieldToField(&inst.AngleType, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeBearingType](), formDiv)
		case "AngleDirectionReference":
			StageSetFormDivSelectFieldToField(&inst.AngleDirectionReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDirectionReferenceType](), formDiv)
		case "InnerDistance":
			StageSetFormDivSelectFieldToField(&inst.InnerDistance, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "OuterDistance":
			StageSetFormDivSelectFieldToField(&inst.OuterDistance, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "UpperLimit":
			StageSetFormDivSelectFieldToField(&inst.UpperLimit, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "UpperLimitReference":
			StageSetFormDivSelectFieldToField(&inst.UpperLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "LowerLimit":
			StageSetFormDivSelectFieldToField(&inst.LowerLimit, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "LowerLimitReference":
			StageSetFormDivSelectFieldToField(&inst.LowerLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		}
	}
}

func saveStageSet_CirclingAreaTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.CirclingAreaTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "CirclingAreaTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.CirclingAreaTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CirclingAreaTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_CirclingAreaTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.CirclingAreaTimeSliceType,
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
		case "Extent":
			StageSetFormDivSelectFieldToField(&inst.Extent, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SurfacePropertyType](), formDiv)
		case "Approach":
			StageSetFormDivSelectFieldToField(&inst.Approach, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.InstrumentApproachProcedurePropertyType](), formDiv)
		case "AircraftCategory":
			StageSetFormDivSelectFieldToField(&inst.AircraftCategory, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AircraftCharacteristicPropertyType](), formDiv)
		}
	}
}

func saveStageSet_CirclingAreaType_AixmfeaturesStage(
	inst *aixmfeatures.CirclingAreaType,
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

func saveStageSet_CirclingRestrictionPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.CirclingRestrictionPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "CirclingRestriction":
			StageSetFormDivSelectFieldToField(&inst.CirclingRestriction, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CirclingRestrictionType](), formDiv)
		}
	}
}

func saveStageSet_CirclingRestrictionType_AixmfeaturesStage(
	inst *aixmfeatures.CirclingRestrictionType,
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
		case "SectorDescription":
			StageSetFormDivSelectFieldToField(&inst.SectorDescription, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CircleSectorPropertyType](), formDiv)
		case "RestrictionArea":
			StageSetFormDivSelectFieldToField(&inst.RestrictionArea, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SurfacePropertyType](), formDiv)
		}
	}
}

func saveStageSet_CurvePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.CurvePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "Curve":
			StageSetFormDivSelectFieldToField(&inst.Curve, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CurveType](), formDiv)
		case "ElevatedCurve":
			StageSetFormDivSelectFieldToField(&inst.ElevatedCurve, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedCurveType](), formDiv)
		}
	}
}

func saveStageSet_CurveType_AixmfeaturesStage(
	inst *aixmfeatures.CurveType,
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
		case "Segments":
			StageSetFormDivSelectFieldToField(&inst.Segments, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CurveSegmentArrayPropertyType](), formDiv)
		case "HorizontalAccuracy":
			StageSetFormDivSelectFieldToField(&inst.HorizontalAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		}
	}
}

func saveStageSet_DMEPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.DMEPropertyType,
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

func saveStageSet_DepartureArrivalConditionPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.DepartureArrivalConditionPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "DepartureArrivalCondition":
			StageSetFormDivSelectFieldToField(&inst.DepartureArrivalCondition, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DepartureArrivalConditionType](), formDiv)
		}
	}
}

func saveStageSet_DepartureArrivalConditionType_AixmfeaturesStage(
	inst *aixmfeatures.DepartureArrivalConditionType,
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
		case "MinimumEnrouteAltitude":
			StageSetFormDivSelectFieldToField(&inst.MinimumEnrouteAltitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "MinimumCrossingAtEnd":
			StageSetFormDivSelectFieldToField(&inst.MinimumCrossingAtEnd, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "MinimumCrossingAtEndReference":
			StageSetFormDivSelectFieldToField(&inst.MinimumCrossingAtEndReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "MaximumCrossingAtEnd":
			StageSetFormDivSelectFieldToField(&inst.MaximumCrossingAtEnd, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "MaximumCrossingAtEndReference":
			StageSetFormDivSelectFieldToField(&inst.MaximumCrossingAtEndReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "EngineType":
			StageSetFormDivSelectFieldToField(&inst.EngineType, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AircraftCharacteristicPropertyType](), formDiv)
		}
	}
}

func saveStageSet_DepartureLegTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.DepartureLegTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "DepartureLegTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.DepartureLegTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DepartureLegTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_DepartureLegTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.DepartureLegTimeSliceType,
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
		case "EndConditionDesignator":
			StageSetFormDivSelectFieldToField(&inst.EndConditionDesignator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSegmentTerminationType](), formDiv)
		case "LegPath":
			StageSetFormDivSelectFieldToField(&inst.LegPath, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTrajectoryType](), formDiv)
		case "LegTypeARINC":
			StageSetFormDivSelectFieldToField(&inst.LegTypeARINC, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSegmentPathType](), formDiv)
		case "Course":
			StageSetFormDivSelectFieldToField(&inst.Course, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "CourseType":
			StageSetFormDivSelectFieldToField(&inst.CourseType, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCourseType](), formDiv)
		case "CourseDirection":
			StageSetFormDivSelectFieldToField(&inst.CourseDirection, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDirectionReferenceType](), formDiv)
		case "TurnDirection":
			StageSetFormDivSelectFieldToField(&inst.TurnDirection, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDirectionTurnType](), formDiv)
		case "SpeedLimit":
			StageSetFormDivSelectFieldToField(&inst.SpeedLimit, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValSpeedType](), formDiv)
		case "SpeedReference":
			StageSetFormDivSelectFieldToField(&inst.SpeedReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSpeedReferenceType](), formDiv)
		case "SpeedInterpretation":
			StageSetFormDivSelectFieldToField(&inst.SpeedInterpretation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAltitudeUseType](), formDiv)
		case "BankAngle":
			StageSetFormDivSelectFieldToField(&inst.BankAngle, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "Length":
			StageSetFormDivSelectFieldToField(&inst.Length, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "Duration":
			StageSetFormDivSelectFieldToField(&inst.Duration, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDurationType](), formDiv)
		case "ProcedureTurnRequired":
			StageSetFormDivSelectFieldToField(&inst.ProcedureTurnRequired, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "UpperLimitAltitude":
			StageSetFormDivSelectFieldToField(&inst.UpperLimitAltitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "UpperLimitReference":
			StageSetFormDivSelectFieldToField(&inst.UpperLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "LowerLimitAltitude":
			StageSetFormDivSelectFieldToField(&inst.LowerLimitAltitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "LowerLimitReference":
			StageSetFormDivSelectFieldToField(&inst.LowerLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "AltitudeInterpretation":
			StageSetFormDivSelectFieldToField(&inst.AltitudeInterpretation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAltitudeUseType](), formDiv)
		case "AltitudeOverrideATC":
			StageSetFormDivSelectFieldToField(&inst.AltitudeOverrideATC, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "AltitudeOverrideReference":
			StageSetFormDivSelectFieldToField(&inst.AltitudeOverrideReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "VerticalAngle":
			StageSetFormDivSelectFieldToField(&inst.VerticalAngle, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "StartPoint":
			StageSetFormDivSelectFieldToField(&inst.StartPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TerminalSegmentPointPropertyType](), formDiv)
		case "EndPoint":
			StageSetFormDivSelectFieldToField(&inst.EndPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TerminalSegmentPointPropertyType](), formDiv)
		case "Trajectory":
			StageSetFormDivSelectFieldToField(&inst.Trajectory, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CurvePropertyType](), formDiv)
		case "ArcCentre":
			StageSetFormDivSelectFieldToField(&inst.ArcCentre, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TerminalSegmentPointPropertyType](), formDiv)
		case "Angle":
			StageSetFormDivSelectFieldToField(&inst.Angle, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AngleIndicationPropertyType](), formDiv)
		case "Distance":
			StageSetFormDivSelectFieldToField(&inst.Distance, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DistanceIndicationPropertyType](), formDiv)
		case "Holding":
			StageSetFormDivSelectFieldToField(&inst.Holding, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.HoldingUsePropertyType](), formDiv)
		case "RequiredNavigationPerformance":
			StageSetFormDivSelectFieldToField(&inst.RequiredNavigationPerformance, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRNPType](), formDiv)
		case "MinimumObstacleClearanceAltitude":
			StageSetFormDivSelectFieldToField(&inst.MinimumObstacleClearanceAltitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "Departure":
			StageSetFormDivSelectFieldToField(&inst.Departure, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.StandardInstrumentDeparturePropertyType](), formDiv)
		}
	}
}

func saveStageSet_DepartureLegType_AixmfeaturesStage(
	inst *aixmfeatures.DepartureLegType,
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

func saveStageSet_DistanceIndicationPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.DistanceIndicationPropertyType,
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

func saveStageSet_ElevatedCurveType_AixmfeaturesStage(
	inst *aixmfeatures.ElevatedCurveType,
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
		case "Segments":
			StageSetFormDivSelectFieldToField(&inst.Segments, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CurveSegmentArrayPropertyType](), formDiv)
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

func saveStageSet_ElevatedSurfaceType_AixmfeaturesStage(
	inst *aixmfeatures.ElevatedSurfaceType,
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
		case "Patches":
			StageSetFormDivSelectFieldToField(&inst.Patches, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.SurfacePatchArrayPropertyType](), formDiv)
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

func saveStageSet_EnRouteSegmentPointType_AixmfeaturesStage(
	inst *aixmfeatures.EnRouteSegmentPointType,
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
		case "ReportingATC":
			StageSetFormDivSelectFieldToField(&inst.ReportingATC, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeATCReportingType](), formDiv)
		case "FlyOver":
			StageSetFormDivSelectFieldToField(&inst.FlyOver, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "Waypoint":
			StageSetFormDivSelectFieldToField(&inst.Waypoint, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "RadarGuidance":
			StageSetFormDivSelectFieldToField(&inst.RadarGuidance, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
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
		case "ExtendedServiceVolume":
			StageSetFormDivSelectFieldToField(&inst.ExtendedServiceVolume, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RadioFrequencyAreaPropertyType](), formDiv)
		case "RoleFreeFlight":
			StageSetFormDivSelectFieldToField(&inst.RoleFreeFlight, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFreeFlightType](), formDiv)
		case "RoleRVSM":
			StageSetFormDivSelectFieldToField(&inst.RoleRVSM, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRVSMPointRoleType](), formDiv)
		case "TurnRadius":
			StageSetFormDivSelectFieldToField(&inst.TurnRadius, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "RoleMilitaryTraining":
			StageSetFormDivSelectFieldToField(&inst.RoleMilitaryTraining, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeMilitaryRoutePointType](), formDiv)
		}
	}
}

func saveStageSet_EquipmentUnavailableAdjustmentColumnPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.EquipmentUnavailableAdjustmentColumnPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EquipmentUnavailableAdjustmentColumn":
			StageSetFormDivSelectFieldToField(&inst.EquipmentUnavailableAdjustmentColumn, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.EquipmentUnavailableAdjustmentColumnType](), formDiv)
		}
	}
}

func saveStageSet_EquipmentUnavailableAdjustmentColumnType_AixmfeaturesStage(
	inst *aixmfeatures.EquipmentUnavailableAdjustmentColumnType,
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
		case "GuidanceEquipment":
			StageSetFormDivSelectFieldToField(&inst.GuidanceEquipment, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeApproachType](), formDiv)
		case "LandingSystemLights":
			StageSetFormDivSelectFieldToField(&inst.LandingSystemLights, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "EquipmentRVR":
			StageSetFormDivSelectFieldToField(&inst.EquipmentRVR, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "VisibilityAdjustment":
			StageSetFormDivSelectFieldToField(&inst.VisibilityAdjustment, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "ApproachLightingInoperative":
			StageSetFormDivSelectFieldToField(&inst.ApproachLightingInoperative, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		}
	}
}

func saveStageSet_EquipmentUnavailableAdjustmentPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.EquipmentUnavailableAdjustmentPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EquipmentUnavailableAdjustment":
			StageSetFormDivSelectFieldToField(&inst.EquipmentUnavailableAdjustment, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.EquipmentUnavailableAdjustmentType](), formDiv)
		}
	}
}

func saveStageSet_EquipmentUnavailableAdjustmentType_AixmfeaturesStage(
	inst *aixmfeatures.EquipmentUnavailableAdjustmentType,
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
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeEquipmentUnavailableType](), formDiv)
		case "ApproachLightingInoperative":
			StageSetFormDivSelectFieldToField(&inst.ApproachLightingInoperative, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		}
	}
}

func saveStageSet_FASDataBlockPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.FASDataBlockPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "FASDataBlock":
			StageSetFormDivSelectFieldToField(&inst.FASDataBlock, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FASDataBlockType](), formDiv)
		}
	}
}

func saveStageSet_FASDataBlockType_AixmfeaturesStage(
	inst *aixmfeatures.FASDataBlockType,
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
		case "HorizontalAlarmLimit":
			StageSetFormDivSelectFieldToField(&inst.HorizontalAlarmLimit, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAlarmLimitType](), formDiv)
		case "VerticalAlarmLimit":
			StageSetFormDivSelectFieldToField(&inst.VerticalAlarmLimit, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAlarmLimitType](), formDiv)
		case "ThresholdCourseWidth":
			StageSetFormDivSelectFieldToField(&inst.ThresholdCourseWidth, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "LengthOffset":
			StageSetFormDivSelectFieldToField(&inst.LengthOffset, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "CRCRemainder":
			StageSetFormDivSelectFieldToField(&inst.CRCRemainder, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValHexType](), formDiv)
		case "OperationType":
			StageSetFormDivSelectFieldToField(&inst.OperationType, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoSequenceType](), formDiv)
		case "ServiceProviderSBAS":
			StageSetFormDivSelectFieldToField(&inst.ServiceProviderSBAS, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoSequenceType](), formDiv)
		case "ApproachPerformanceDesignator":
			StageSetFormDivSelectFieldToField(&inst.ApproachPerformanceDesignator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoSequenceType](), formDiv)
		case "RouteIndicator":
			StageSetFormDivSelectFieldToField(&inst.RouteIndicator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRouteIndicatorType](), formDiv)
		case "ReferencePathDataSelector":
			StageSetFormDivSelectFieldToField(&inst.ReferencePathDataSelector, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoSequenceType](), formDiv)
		case "ReferencePathIdentifier":
			StageSetFormDivSelectFieldToField(&inst.ReferencePathIdentifier, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeReferencePathIdentifierType](), formDiv)
		case "CodeICAO":
			StageSetFormDivSelectFieldToField(&inst.CodeICAO, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeICAOCountryType](), formDiv)
		}
	}
}

func saveStageSet_FinalLegTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.FinalLegTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "FinalLegTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.FinalLegTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FinalLegTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_FinalLegTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.FinalLegTimeSliceType,
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
		case "EndConditionDesignator":
			StageSetFormDivSelectFieldToField(&inst.EndConditionDesignator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSegmentTerminationType](), formDiv)
		case "LegPath":
			StageSetFormDivSelectFieldToField(&inst.LegPath, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTrajectoryType](), formDiv)
		case "LegTypeARINC":
			StageSetFormDivSelectFieldToField(&inst.LegTypeARINC, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSegmentPathType](), formDiv)
		case "Course":
			StageSetFormDivSelectFieldToField(&inst.Course, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "CourseType":
			StageSetFormDivSelectFieldToField(&inst.CourseType, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCourseType](), formDiv)
		case "CourseDirection":
			StageSetFormDivSelectFieldToField(&inst.CourseDirection, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDirectionReferenceType](), formDiv)
		case "TurnDirection":
			StageSetFormDivSelectFieldToField(&inst.TurnDirection, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDirectionTurnType](), formDiv)
		case "SpeedLimit":
			StageSetFormDivSelectFieldToField(&inst.SpeedLimit, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValSpeedType](), formDiv)
		case "SpeedReference":
			StageSetFormDivSelectFieldToField(&inst.SpeedReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSpeedReferenceType](), formDiv)
		case "SpeedInterpretation":
			StageSetFormDivSelectFieldToField(&inst.SpeedInterpretation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAltitudeUseType](), formDiv)
		case "BankAngle":
			StageSetFormDivSelectFieldToField(&inst.BankAngle, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "Length":
			StageSetFormDivSelectFieldToField(&inst.Length, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "Duration":
			StageSetFormDivSelectFieldToField(&inst.Duration, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDurationType](), formDiv)
		case "ProcedureTurnRequired":
			StageSetFormDivSelectFieldToField(&inst.ProcedureTurnRequired, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "UpperLimitAltitude":
			StageSetFormDivSelectFieldToField(&inst.UpperLimitAltitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "UpperLimitReference":
			StageSetFormDivSelectFieldToField(&inst.UpperLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "LowerLimitAltitude":
			StageSetFormDivSelectFieldToField(&inst.LowerLimitAltitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "LowerLimitReference":
			StageSetFormDivSelectFieldToField(&inst.LowerLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "AltitudeInterpretation":
			StageSetFormDivSelectFieldToField(&inst.AltitudeInterpretation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAltitudeUseType](), formDiv)
		case "AltitudeOverrideATC":
			StageSetFormDivSelectFieldToField(&inst.AltitudeOverrideATC, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "AltitudeOverrideReference":
			StageSetFormDivSelectFieldToField(&inst.AltitudeOverrideReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "VerticalAngle":
			StageSetFormDivSelectFieldToField(&inst.VerticalAngle, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "StartPoint":
			StageSetFormDivSelectFieldToField(&inst.StartPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TerminalSegmentPointPropertyType](), formDiv)
		case "EndPoint":
			StageSetFormDivSelectFieldToField(&inst.EndPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TerminalSegmentPointPropertyType](), formDiv)
		case "Trajectory":
			StageSetFormDivSelectFieldToField(&inst.Trajectory, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CurvePropertyType](), formDiv)
		case "ArcCentre":
			StageSetFormDivSelectFieldToField(&inst.ArcCentre, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TerminalSegmentPointPropertyType](), formDiv)
		case "Angle":
			StageSetFormDivSelectFieldToField(&inst.Angle, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AngleIndicationPropertyType](), formDiv)
		case "Distance":
			StageSetFormDivSelectFieldToField(&inst.Distance, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DistanceIndicationPropertyType](), formDiv)
		case "Holding":
			StageSetFormDivSelectFieldToField(&inst.Holding, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.HoldingUsePropertyType](), formDiv)
		case "Approach":
			StageSetFormDivSelectFieldToField(&inst.Approach, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.InstrumentApproachProcedurePropertyType](), formDiv)
		case "GuidanceSystem":
			StageSetFormDivSelectFieldToField(&inst.GuidanceSystem, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFinalGuidanceType](), formDiv)
		case "LandingSystemCategory":
			StageSetFormDivSelectFieldToField(&inst.LandingSystemCategory, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeApproachGuidanceType](), formDiv)
		case "MinimumBaroVnavTemperature":
			StageSetFormDivSelectFieldToField(&inst.MinimumBaroVnavTemperature, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValTemperatureType](), formDiv)
		case "RnpDMEAuthorized":
			StageSetFormDivSelectFieldToField(&inst.RnpDMEAuthorized, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "CourseOffsetAngle":
			StageSetFormDivSelectFieldToField(&inst.CourseOffsetAngle, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "CourseOffsetSide":
			StageSetFormDivSelectFieldToField(&inst.CourseOffsetSide, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSideType](), formDiv)
		case "CourseCentrelineDistance":
			StageSetFormDivSelectFieldToField(&inst.CourseCentrelineDistance, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "CourseOffsetDistance":
			StageSetFormDivSelectFieldToField(&inst.CourseOffsetDistance, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "CourseCentrelineIntersect":
			StageSetFormDivSelectFieldToField(&inst.CourseCentrelineIntersect, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRelativePositionType](), formDiv)
		case "FinalPathAlignmentPoint_fixDesignatedPoint":
			StageSetFormDivSelectFieldToField(&inst.FinalPathAlignmentPoint_fixDesignatedPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), formDiv)
		case "FinalPathAlignmentPoint_navaidSystem":
			StageSetFormDivSelectFieldToField(&inst.FinalPathAlignmentPoint_navaidSystem, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), formDiv)
		case "FinalPathAlignmentPoint_position":
			StageSetFormDivSelectFieldToField(&inst.FinalPathAlignmentPoint_position, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), formDiv)
		case "FinalPathAlignmentPoint_runwayPoint":
			StageSetFormDivSelectFieldToField(&inst.FinalPathAlignmentPoint_runwayPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), formDiv)
		case "FinalPathAlignmentPoint_aimingPoint":
			StageSetFormDivSelectFieldToField(&inst.FinalPathAlignmentPoint_aimingPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), formDiv)
		case "FinalPathAlignmentPoint_airportReferencePoint":
			StageSetFormDivSelectFieldToField(&inst.FinalPathAlignmentPoint_airportReferencePoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), formDiv)
		case "VisualDescentPoint":
			StageSetFormDivSelectFieldToField(&inst.VisualDescentPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TerminalSegmentPointPropertyType](), formDiv)
		case "FASData":
			StageSetFormDivSelectFieldToField(&inst.FASData, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FASDataBlockPropertyType](), formDiv)
		}
	}
}

func saveStageSet_FinalLegType_AixmfeaturesStage(
	inst *aixmfeatures.FinalLegType,
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

func saveStageSet_FinalProfilePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.FinalProfilePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "FinalProfile":
			StageSetFormDivSelectFieldToField(&inst.FinalProfile, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FinalProfileType](), formDiv)
		}
	}
}

func saveStageSet_FinalProfileType_AixmfeaturesStage(
	inst *aixmfeatures.FinalProfileType,
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

func saveStageSet_HoldingAssessmentTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.HoldingAssessmentTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "HoldingAssessmentTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.HoldingAssessmentTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.HoldingAssessmentTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_HoldingAssessmentTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.HoldingAssessmentTimeSliceType,
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
		case "UpperLimit":
			StageSetFormDivSelectFieldToField(&inst.UpperLimit, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "UpperLimitReference":
			StageSetFormDivSelectFieldToField(&inst.UpperLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "LowerLimit":
			StageSetFormDivSelectFieldToField(&inst.LowerLimit, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "LowerLimitReference":
			StageSetFormDivSelectFieldToField(&inst.LowerLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "SpeedLimit":
			StageSetFormDivSelectFieldToField(&inst.SpeedLimit, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValSpeedType](), formDiv)
		case "PatternTemplate":
			StageSetFormDivSelectFieldToField(&inst.PatternTemplate, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextNameType](), formDiv)
		case "TurbulentAir":
			StageSetFormDivSelectFieldToField(&inst.TurbulentAir, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "LegLengthToward":
			StageSetFormDivSelectFieldToField(&inst.LegLengthToward, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "LegLengthAway":
			StageSetFormDivSelectFieldToField(&inst.LegLengthAway, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "HoldingPoint":
			StageSetFormDivSelectFieldToField(&inst.HoldingPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SegmentPointPropertyType](), formDiv)
		case "UnplannedHolding":
			StageSetFormDivSelectFieldToField(&inst.UnplannedHolding, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.UnplannedHoldingPropertyType](), formDiv)
		case "AssessedHoldingPattern":
			StageSetFormDivSelectFieldToField(&inst.AssessedHoldingPattern, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.HoldingPatternPropertyType](), formDiv)
		}
	}
}

func saveStageSet_HoldingAssessmentType_AixmfeaturesStage(
	inst *aixmfeatures.HoldingAssessmentType,
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

func saveStageSet_HoldingPatternDistancePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.HoldingPatternDistancePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "HoldingPatternDistance":
			StageSetFormDivSelectFieldToField(&inst.HoldingPatternDistance, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.HoldingPatternDistanceType](), formDiv)
		}
	}
}

func saveStageSet_HoldingPatternDistanceType_AixmfeaturesStage(
	inst *aixmfeatures.HoldingPatternDistanceType,
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
		case "Length":
			StageSetFormDivSelectFieldToField(&inst.Length, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		}
	}
}

func saveStageSet_HoldingPatternDurationPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.HoldingPatternDurationPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "HoldingPatternDuration":
			StageSetFormDivSelectFieldToField(&inst.HoldingPatternDuration, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.HoldingPatternDurationType](), formDiv)
		}
	}
}

func saveStageSet_HoldingPatternDurationType_AixmfeaturesStage(
	inst *aixmfeatures.HoldingPatternDurationType,
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
		case "Duration":
			StageSetFormDivSelectFieldToField(&inst.Duration, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDurationType](), formDiv)
		}
	}
}

func saveStageSet_HoldingPatternPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.HoldingPatternPropertyType,
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

func saveStageSet_HoldingPatternTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.HoldingPatternTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "HoldingPatternTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.HoldingPatternTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.HoldingPatternTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_HoldingPatternTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.HoldingPatternTimeSliceType,
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
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeHoldingUsageType](), formDiv)
		case "OutboundCourse":
			StageSetFormDivSelectFieldToField(&inst.OutboundCourse, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "OutboundCourseType":
			StageSetFormDivSelectFieldToField(&inst.OutboundCourseType, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCourseType](), formDiv)
		case "InboundCourse":
			StageSetFormDivSelectFieldToField(&inst.InboundCourse, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "TurnDirection":
			StageSetFormDivSelectFieldToField(&inst.TurnDirection, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDirectionTurnType](), formDiv)
		case "UpperLimit":
			StageSetFormDivSelectFieldToField(&inst.UpperLimit, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "UpperLimitReference":
			StageSetFormDivSelectFieldToField(&inst.UpperLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "LowerLimit":
			StageSetFormDivSelectFieldToField(&inst.LowerLimit, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "LowerLimitReference":
			StageSetFormDivSelectFieldToField(&inst.LowerLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "SpeedLimit":
			StageSetFormDivSelectFieldToField(&inst.SpeedLimit, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValSpeedType](), formDiv)
		case "Instruction":
			StageSetFormDivSelectFieldToField(&inst.Instruction, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextInstructionType](), formDiv)
		case "NonStandardHolding":
			StageSetFormDivSelectFieldToField(&inst.NonStandardHolding, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "OutboundLegSpan_endPoint":
			StageSetFormDivSelectFieldToField(&inst.OutboundLegSpan_endPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SegmentPointPropertyType](), formDiv)
		case "OutboundLegSpan_endDistance":
			StageSetFormDivSelectFieldToField(&inst.OutboundLegSpan_endDistance, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.HoldingPatternDistancePropertyType](), formDiv)
		case "OutboundLegSpan_endTime":
			StageSetFormDivSelectFieldToField(&inst.OutboundLegSpan_endTime, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.HoldingPatternDurationPropertyType](), formDiv)
		case "HoldingPoint":
			StageSetFormDivSelectFieldToField(&inst.HoldingPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SegmentPointPropertyType](), formDiv)
		case "Extent":
			StageSetFormDivSelectFieldToField(&inst.Extent, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CurvePropertyType](), formDiv)
		}
	}
}

func saveStageSet_HoldingPatternType_AixmfeaturesStage(
	inst *aixmfeatures.HoldingPatternType,
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

func saveStageSet_HoldingUsePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.HoldingUsePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "HoldingUse":
			StageSetFormDivSelectFieldToField(&inst.HoldingUse, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.HoldingUseType](), formDiv)
		}
	}
}

func saveStageSet_HoldingUseType_AixmfeaturesStage(
	inst *aixmfeatures.HoldingUseType,
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
		case "HoldingUse":
			StageSetFormDivSelectFieldToField(&inst.HoldingUse, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeHoldingUseType](), formDiv)
		case "Instruction":
			StageSetFormDivSelectFieldToField(&inst.Instruction, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextInstructionType](), formDiv)
		case "InstructedAltitude":
			StageSetFormDivSelectFieldToField(&inst.InstructedAltitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "InstructionAltitudeReference":
			StageSetFormDivSelectFieldToField(&inst.InstructionAltitudeReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "TheHoldingPattern":
			StageSetFormDivSelectFieldToField(&inst.TheHoldingPattern, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.HoldingPatternPropertyType](), formDiv)
		}
	}
}

func saveStageSet_InitialLegTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.InitialLegTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "InitialLegTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.InitialLegTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.InitialLegTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_InitialLegTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.InitialLegTimeSliceType,
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
		case "EndConditionDesignator":
			StageSetFormDivSelectFieldToField(&inst.EndConditionDesignator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSegmentTerminationType](), formDiv)
		case "LegPath":
			StageSetFormDivSelectFieldToField(&inst.LegPath, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTrajectoryType](), formDiv)
		case "LegTypeARINC":
			StageSetFormDivSelectFieldToField(&inst.LegTypeARINC, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSegmentPathType](), formDiv)
		case "Course":
			StageSetFormDivSelectFieldToField(&inst.Course, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "CourseType":
			StageSetFormDivSelectFieldToField(&inst.CourseType, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCourseType](), formDiv)
		case "CourseDirection":
			StageSetFormDivSelectFieldToField(&inst.CourseDirection, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDirectionReferenceType](), formDiv)
		case "TurnDirection":
			StageSetFormDivSelectFieldToField(&inst.TurnDirection, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDirectionTurnType](), formDiv)
		case "SpeedLimit":
			StageSetFormDivSelectFieldToField(&inst.SpeedLimit, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValSpeedType](), formDiv)
		case "SpeedReference":
			StageSetFormDivSelectFieldToField(&inst.SpeedReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSpeedReferenceType](), formDiv)
		case "SpeedInterpretation":
			StageSetFormDivSelectFieldToField(&inst.SpeedInterpretation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAltitudeUseType](), formDiv)
		case "BankAngle":
			StageSetFormDivSelectFieldToField(&inst.BankAngle, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "Length":
			StageSetFormDivSelectFieldToField(&inst.Length, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "Duration":
			StageSetFormDivSelectFieldToField(&inst.Duration, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDurationType](), formDiv)
		case "ProcedureTurnRequired":
			StageSetFormDivSelectFieldToField(&inst.ProcedureTurnRequired, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "UpperLimitAltitude":
			StageSetFormDivSelectFieldToField(&inst.UpperLimitAltitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "UpperLimitReference":
			StageSetFormDivSelectFieldToField(&inst.UpperLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "LowerLimitAltitude":
			StageSetFormDivSelectFieldToField(&inst.LowerLimitAltitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "LowerLimitReference":
			StageSetFormDivSelectFieldToField(&inst.LowerLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "AltitudeInterpretation":
			StageSetFormDivSelectFieldToField(&inst.AltitudeInterpretation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAltitudeUseType](), formDiv)
		case "AltitudeOverrideATC":
			StageSetFormDivSelectFieldToField(&inst.AltitudeOverrideATC, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "AltitudeOverrideReference":
			StageSetFormDivSelectFieldToField(&inst.AltitudeOverrideReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "VerticalAngle":
			StageSetFormDivSelectFieldToField(&inst.VerticalAngle, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "StartPoint":
			StageSetFormDivSelectFieldToField(&inst.StartPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TerminalSegmentPointPropertyType](), formDiv)
		case "EndPoint":
			StageSetFormDivSelectFieldToField(&inst.EndPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TerminalSegmentPointPropertyType](), formDiv)
		case "Trajectory":
			StageSetFormDivSelectFieldToField(&inst.Trajectory, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CurvePropertyType](), formDiv)
		case "ArcCentre":
			StageSetFormDivSelectFieldToField(&inst.ArcCentre, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TerminalSegmentPointPropertyType](), formDiv)
		case "Angle":
			StageSetFormDivSelectFieldToField(&inst.Angle, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AngleIndicationPropertyType](), formDiv)
		case "Distance":
			StageSetFormDivSelectFieldToField(&inst.Distance, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DistanceIndicationPropertyType](), formDiv)
		case "Holding":
			StageSetFormDivSelectFieldToField(&inst.Holding, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.HoldingUsePropertyType](), formDiv)
		case "Approach":
			StageSetFormDivSelectFieldToField(&inst.Approach, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.InstrumentApproachProcedurePropertyType](), formDiv)
		case "RequiredNavigationPerformance":
			StageSetFormDivSelectFieldToField(&inst.RequiredNavigationPerformance, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRNPType](), formDiv)
		}
	}
}

func saveStageSet_InitialLegType_AixmfeaturesStage(
	inst *aixmfeatures.InitialLegType,
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

func saveStageSet_InstrumentApproachProcedurePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.InstrumentApproachProcedurePropertyType,
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

func saveStageSet_InstrumentApproachProcedureTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.InstrumentApproachProcedureTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "InstrumentApproachProcedureTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.InstrumentApproachProcedureTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.InstrumentApproachProcedureTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_InstrumentApproachProcedureTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.InstrumentApproachProcedureTimeSliceType,
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
		case "CommunicationFailureInstruction":
			StageSetFormDivSelectFieldToField(&inst.CommunicationFailureInstruction, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextInstructionType](), formDiv)
		case "Instruction":
			StageSetFormDivSelectFieldToField(&inst.Instruction, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextInstructionType](), formDiv)
		case "DesignCriteria":
			StageSetFormDivSelectFieldToField(&inst.DesignCriteria, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDesignStandardType](), formDiv)
		case "CodingStandard":
			StageSetFormDivSelectFieldToField(&inst.CodingStandard, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeProcedureCodingStandardType](), formDiv)
		case "FlightChecked":
			StageSetFormDivSelectFieldToField(&inst.FlightChecked, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "RNAV":
			StageSetFormDivSelectFieldToField(&inst.RNAV, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "GuidanceFacility_specialNavigationSystem":
			StageSetFormDivSelectFieldToField(&inst.GuidanceFacility_specialNavigationSystem, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SpecialNavigationSystemPropertyType](), formDiv)
		case "GuidanceFacility_radar":
			StageSetFormDivSelectFieldToField(&inst.GuidanceFacility_radar, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RadarSystemPropertyType](), formDiv)
		case "GuidanceFacility_navaid":
			StageSetFormDivSelectFieldToField(&inst.GuidanceFacility_navaid, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), formDiv)
		case "SafeAltitude":
			StageSetFormDivSelectFieldToField(&inst.SafeAltitude, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SafeAltitudeAreaPropertyType](), formDiv)
		case "ApproachPrefix":
			StageSetFormDivSelectFieldToField(&inst.ApproachPrefix, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeApproachPrefixType](), formDiv)
		case "ApproachType":
			StageSetFormDivSelectFieldToField(&inst.ApproachType, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeApproachType](), formDiv)
		case "MultipleIdentification":
			StageSetFormDivSelectFieldToField(&inst.MultipleIdentification, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeUpperAlphaType](), formDiv)
		case "CopterTrack":
			StageSetFormDivSelectFieldToField(&inst.CopterTrack, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "CirclingIdentification":
			StageSetFormDivSelectFieldToField(&inst.CirclingIdentification, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeUpperAlphaType](), formDiv)
		case "CourseReversalInstruction":
			StageSetFormDivSelectFieldToField(&inst.CourseReversalInstruction, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextInstructionType](), formDiv)
		case "AdditionalEquipment":
			StageSetFormDivSelectFieldToField(&inst.AdditionalEquipment, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeApproachEquipmentAdditionalType](), formDiv)
		case "ChannelGNSS":
			StageSetFormDivSelectFieldToField(&inst.ChannelGNSS, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValChannelNumberType](), formDiv)
		case "WAASReliable":
			StageSetFormDivSelectFieldToField(&inst.WAASReliable, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "Landing":
			StageSetFormDivSelectFieldToField(&inst.Landing, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.LandingTakeoffAreaCollectionPropertyType](), formDiv)
		case "FinalProfile":
			StageSetFormDivSelectFieldToField(&inst.FinalProfile, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FinalProfilePropertyType](), formDiv)
		}
	}
}

func saveStageSet_InstrumentApproachProcedureType_AixmfeaturesStage(
	inst *aixmfeatures.InstrumentApproachProcedureType,
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

func saveStageSet_IntermediateLegTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.IntermediateLegTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "IntermediateLegTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.IntermediateLegTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.IntermediateLegTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_IntermediateLegTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.IntermediateLegTimeSliceType,
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
		case "EndConditionDesignator":
			StageSetFormDivSelectFieldToField(&inst.EndConditionDesignator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSegmentTerminationType](), formDiv)
		case "LegPath":
			StageSetFormDivSelectFieldToField(&inst.LegPath, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTrajectoryType](), formDiv)
		case "LegTypeARINC":
			StageSetFormDivSelectFieldToField(&inst.LegTypeARINC, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSegmentPathType](), formDiv)
		case "Course":
			StageSetFormDivSelectFieldToField(&inst.Course, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "CourseType":
			StageSetFormDivSelectFieldToField(&inst.CourseType, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCourseType](), formDiv)
		case "CourseDirection":
			StageSetFormDivSelectFieldToField(&inst.CourseDirection, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDirectionReferenceType](), formDiv)
		case "TurnDirection":
			StageSetFormDivSelectFieldToField(&inst.TurnDirection, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDirectionTurnType](), formDiv)
		case "SpeedLimit":
			StageSetFormDivSelectFieldToField(&inst.SpeedLimit, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValSpeedType](), formDiv)
		case "SpeedReference":
			StageSetFormDivSelectFieldToField(&inst.SpeedReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSpeedReferenceType](), formDiv)
		case "SpeedInterpretation":
			StageSetFormDivSelectFieldToField(&inst.SpeedInterpretation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAltitudeUseType](), formDiv)
		case "BankAngle":
			StageSetFormDivSelectFieldToField(&inst.BankAngle, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "Length":
			StageSetFormDivSelectFieldToField(&inst.Length, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "Duration":
			StageSetFormDivSelectFieldToField(&inst.Duration, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDurationType](), formDiv)
		case "ProcedureTurnRequired":
			StageSetFormDivSelectFieldToField(&inst.ProcedureTurnRequired, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "UpperLimitAltitude":
			StageSetFormDivSelectFieldToField(&inst.UpperLimitAltitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "UpperLimitReference":
			StageSetFormDivSelectFieldToField(&inst.UpperLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "LowerLimitAltitude":
			StageSetFormDivSelectFieldToField(&inst.LowerLimitAltitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "LowerLimitReference":
			StageSetFormDivSelectFieldToField(&inst.LowerLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "AltitudeInterpretation":
			StageSetFormDivSelectFieldToField(&inst.AltitudeInterpretation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAltitudeUseType](), formDiv)
		case "AltitudeOverrideATC":
			StageSetFormDivSelectFieldToField(&inst.AltitudeOverrideATC, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "AltitudeOverrideReference":
			StageSetFormDivSelectFieldToField(&inst.AltitudeOverrideReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "VerticalAngle":
			StageSetFormDivSelectFieldToField(&inst.VerticalAngle, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "StartPoint":
			StageSetFormDivSelectFieldToField(&inst.StartPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TerminalSegmentPointPropertyType](), formDiv)
		case "EndPoint":
			StageSetFormDivSelectFieldToField(&inst.EndPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TerminalSegmentPointPropertyType](), formDiv)
		case "Trajectory":
			StageSetFormDivSelectFieldToField(&inst.Trajectory, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CurvePropertyType](), formDiv)
		case "ArcCentre":
			StageSetFormDivSelectFieldToField(&inst.ArcCentre, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TerminalSegmentPointPropertyType](), formDiv)
		case "Angle":
			StageSetFormDivSelectFieldToField(&inst.Angle, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AngleIndicationPropertyType](), formDiv)
		case "Distance":
			StageSetFormDivSelectFieldToField(&inst.Distance, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DistanceIndicationPropertyType](), formDiv)
		case "Holding":
			StageSetFormDivSelectFieldToField(&inst.Holding, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.HoldingUsePropertyType](), formDiv)
		case "Approach":
			StageSetFormDivSelectFieldToField(&inst.Approach, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.InstrumentApproachProcedurePropertyType](), formDiv)
		case "RequiredNavigationPerformance":
			StageSetFormDivSelectFieldToField(&inst.RequiredNavigationPerformance, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRNPType](), formDiv)
		}
	}
}

func saveStageSet_IntermediateLegType_AixmfeaturesStage(
	inst *aixmfeatures.IntermediateLegType,
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

func saveStageSet_LandingTakeoffAreaCollectionPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.LandingTakeoffAreaCollectionPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "LandingTakeoffAreaCollection":
			StageSetFormDivSelectFieldToField(&inst.LandingTakeoffAreaCollection, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.LandingTakeoffAreaCollectionType](), formDiv)
		}
	}
}

func saveStageSet_LandingTakeoffAreaCollectionType_AixmfeaturesStage(
	inst *aixmfeatures.LandingTakeoffAreaCollectionType,
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

func saveStageSet_MinimaPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.MinimaPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "Minima":
			StageSetFormDivSelectFieldToField(&inst.Minima, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.MinimaType](), formDiv)
		}
	}
}

func saveStageSet_MinimaType_AixmfeaturesStage(
	inst *aixmfeatures.MinimaType,
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
		case "Altitude":
			StageSetFormDivSelectFieldToField(&inst.Altitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "AltitudeCode":
			StageSetFormDivSelectFieldToField(&inst.AltitudeCode, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeMinimumAltitudeType](), formDiv)
		case "AltitudeReference":
			StageSetFormDivSelectFieldToField(&inst.AltitudeReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "Height":
			StageSetFormDivSelectFieldToField(&inst.Height, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "MilitaryHeight":
			StageSetFormDivSelectFieldToField(&inst.MilitaryHeight, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "RadioHeight":
			StageSetFormDivSelectFieldToField(&inst.RadioHeight, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "HeightCode":
			StageSetFormDivSelectFieldToField(&inst.HeightCode, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeMinimumHeightType](), formDiv)
		case "HeightReference":
			StageSetFormDivSelectFieldToField(&inst.HeightReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeHeightReferenceType](), formDiv)
		case "Visibility":
			StageSetFormDivSelectFieldToField(&inst.Visibility, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "MilitaryVisibility":
			StageSetFormDivSelectFieldToField(&inst.MilitaryVisibility, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "MandatoryRVR":
			StageSetFormDivSelectFieldToField(&inst.MandatoryRVR, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "RemoteAltimeterMinima":
			StageSetFormDivSelectFieldToField(&inst.RemoteAltimeterMinima, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		}
	}
}

func saveStageSet_MissedApproachGroupPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.MissedApproachGroupPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "MissedApproachGroup":
			StageSetFormDivSelectFieldToField(&inst.MissedApproachGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.MissedApproachGroupType](), formDiv)
		}
	}
}

func saveStageSet_MissedApproachGroupType_AixmfeaturesStage(
	inst *aixmfeatures.MissedApproachGroupType,
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
		case "Instruction":
			StageSetFormDivSelectFieldToField(&inst.Instruction, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextInstructionType](), formDiv)
		case "AlternateClimbInstruction":
			StageSetFormDivSelectFieldToField(&inst.AlternateClimbInstruction, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextInstructionType](), formDiv)
		case "AlternateClimbAltitude":
			StageSetFormDivSelectFieldToField(&inst.AlternateClimbAltitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		}
	}
}

func saveStageSet_MissedApproachLegTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.MissedApproachLegTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "MissedApproachLegTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.MissedApproachLegTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.MissedApproachLegTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_MissedApproachLegTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.MissedApproachLegTimeSliceType,
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
		case "EndConditionDesignator":
			StageSetFormDivSelectFieldToField(&inst.EndConditionDesignator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSegmentTerminationType](), formDiv)
		case "LegPath":
			StageSetFormDivSelectFieldToField(&inst.LegPath, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTrajectoryType](), formDiv)
		case "LegTypeARINC":
			StageSetFormDivSelectFieldToField(&inst.LegTypeARINC, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSegmentPathType](), formDiv)
		case "Course":
			StageSetFormDivSelectFieldToField(&inst.Course, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "CourseType":
			StageSetFormDivSelectFieldToField(&inst.CourseType, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCourseType](), formDiv)
		case "CourseDirection":
			StageSetFormDivSelectFieldToField(&inst.CourseDirection, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDirectionReferenceType](), formDiv)
		case "TurnDirection":
			StageSetFormDivSelectFieldToField(&inst.TurnDirection, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDirectionTurnType](), formDiv)
		case "SpeedLimit":
			StageSetFormDivSelectFieldToField(&inst.SpeedLimit, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValSpeedType](), formDiv)
		case "SpeedReference":
			StageSetFormDivSelectFieldToField(&inst.SpeedReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSpeedReferenceType](), formDiv)
		case "SpeedInterpretation":
			StageSetFormDivSelectFieldToField(&inst.SpeedInterpretation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAltitudeUseType](), formDiv)
		case "BankAngle":
			StageSetFormDivSelectFieldToField(&inst.BankAngle, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "Length":
			StageSetFormDivSelectFieldToField(&inst.Length, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "Duration":
			StageSetFormDivSelectFieldToField(&inst.Duration, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDurationType](), formDiv)
		case "ProcedureTurnRequired":
			StageSetFormDivSelectFieldToField(&inst.ProcedureTurnRequired, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "UpperLimitAltitude":
			StageSetFormDivSelectFieldToField(&inst.UpperLimitAltitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "UpperLimitReference":
			StageSetFormDivSelectFieldToField(&inst.UpperLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "LowerLimitAltitude":
			StageSetFormDivSelectFieldToField(&inst.LowerLimitAltitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "LowerLimitReference":
			StageSetFormDivSelectFieldToField(&inst.LowerLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "AltitudeInterpretation":
			StageSetFormDivSelectFieldToField(&inst.AltitudeInterpretation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAltitudeUseType](), formDiv)
		case "AltitudeOverrideATC":
			StageSetFormDivSelectFieldToField(&inst.AltitudeOverrideATC, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "AltitudeOverrideReference":
			StageSetFormDivSelectFieldToField(&inst.AltitudeOverrideReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "VerticalAngle":
			StageSetFormDivSelectFieldToField(&inst.VerticalAngle, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "StartPoint":
			StageSetFormDivSelectFieldToField(&inst.StartPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TerminalSegmentPointPropertyType](), formDiv)
		case "EndPoint":
			StageSetFormDivSelectFieldToField(&inst.EndPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TerminalSegmentPointPropertyType](), formDiv)
		case "Trajectory":
			StageSetFormDivSelectFieldToField(&inst.Trajectory, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CurvePropertyType](), formDiv)
		case "ArcCentre":
			StageSetFormDivSelectFieldToField(&inst.ArcCentre, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TerminalSegmentPointPropertyType](), formDiv)
		case "Angle":
			StageSetFormDivSelectFieldToField(&inst.Angle, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AngleIndicationPropertyType](), formDiv)
		case "Distance":
			StageSetFormDivSelectFieldToField(&inst.Distance, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DistanceIndicationPropertyType](), formDiv)
		case "Holding":
			StageSetFormDivSelectFieldToField(&inst.Holding, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.HoldingUsePropertyType](), formDiv)
		case "Approach":
			StageSetFormDivSelectFieldToField(&inst.Approach, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.InstrumentApproachProcedurePropertyType](), formDiv)
		case "Type_XSD":
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeMissedApproachType](), formDiv)
		case "ThresholdAfterMAPT":
			StageSetFormDivSelectFieldToField(&inst.ThresholdAfterMAPT, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "HeightMAPT":
			StageSetFormDivSelectFieldToField(&inst.HeightMAPT, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "RequiredNavigationPerformance":
			StageSetFormDivSelectFieldToField(&inst.RequiredNavigationPerformance, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRNPType](), formDiv)
		}
	}
}

func saveStageSet_MissedApproachLegType_AixmfeaturesStage(
	inst *aixmfeatures.MissedApproachLegType,
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

func saveStageSet_NavigationAreaRestrictionTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.NavigationAreaRestrictionTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "NavigationAreaRestrictionTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.NavigationAreaRestrictionTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavigationAreaRestrictionTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_NavigationAreaRestrictionTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.NavigationAreaRestrictionTimeSliceType,
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
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavigationAreaRestrictionType](), formDiv)
		case "DesignSurface":
			StageSetFormDivSelectFieldToField(&inst.DesignSurface, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ObstacleAssessmentAreaPropertyType](), formDiv)
		case "SectorDefinition":
			StageSetFormDivSelectFieldToField(&inst.SectorDefinition, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CircleSectorPropertyType](), formDiv)
		}
	}
}

func saveStageSet_NavigationAreaRestrictionType_AixmfeaturesStage(
	inst *aixmfeatures.NavigationAreaRestrictionType,
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

func saveStageSet_NavigationAreaSectorPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.NavigationAreaSectorPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "NavigationAreaSector":
			StageSetFormDivSelectFieldToField(&inst.NavigationAreaSector, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavigationAreaSectorType](), formDiv)
		}
	}
}

func saveStageSet_NavigationAreaSectorType_AixmfeaturesStage(
	inst *aixmfeatures.NavigationAreaSectorType,
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
		case "SectorDefinition":
			StageSetFormDivSelectFieldToField(&inst.SectorDefinition, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CircleSectorPropertyType](), formDiv)
		case "Extent":
			StageSetFormDivSelectFieldToField(&inst.Extent, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SurfacePropertyType](), formDiv)
		}
	}
}

func saveStageSet_NavigationAreaTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.NavigationAreaTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "NavigationAreaTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.NavigationAreaTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavigationAreaTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_NavigationAreaTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.NavigationAreaTimeSliceType,
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
		case "NavigationAreaType":
			StageSetFormDivSelectFieldToField(&inst.NavigationAreaType, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavigationAreaType](), formDiv)
		case "MinimumCeiling":
			StageSetFormDivSelectFieldToField(&inst.MinimumCeiling, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "MinimumVisibility":
			StageSetFormDivSelectFieldToField(&inst.MinimumVisibility, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "Departure":
			StageSetFormDivSelectFieldToField(&inst.Departure, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.StandardInstrumentDeparturePropertyType](), formDiv)
		case "CentrePoint_fixDesignatedPoint":
			StageSetFormDivSelectFieldToField(&inst.CentrePoint_fixDesignatedPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), formDiv)
		case "CentrePoint_navaidSystem":
			StageSetFormDivSelectFieldToField(&inst.CentrePoint_navaidSystem, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), formDiv)
		case "CentrePoint_position":
			StageSetFormDivSelectFieldToField(&inst.CentrePoint_position, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), formDiv)
		case "CentrePoint_runwayPoint":
			StageSetFormDivSelectFieldToField(&inst.CentrePoint_runwayPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), formDiv)
		case "CentrePoint_aimingPoint":
			StageSetFormDivSelectFieldToField(&inst.CentrePoint_aimingPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), formDiv)
		case "CentrePoint_airportReferencePoint":
			StageSetFormDivSelectFieldToField(&inst.CentrePoint_airportReferencePoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), formDiv)
		}
	}
}

func saveStageSet_NavigationAreaType_AixmfeaturesStage(
	inst *aixmfeatures.NavigationAreaType,
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

func saveStageSet_ObstacleAssessmentAreaPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.ObstacleAssessmentAreaPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "ObstacleAssessmentArea":
			StageSetFormDivSelectFieldToField(&inst.ObstacleAssessmentArea, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ObstacleAssessmentAreaType](), formDiv)
		}
	}
}

func saveStageSet_ObstacleAssessmentAreaType_AixmfeaturesStage(
	inst *aixmfeatures.ObstacleAssessmentAreaType,
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
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeObstacleAssessmentSurfaceType](), formDiv)
		case "SectionNumber":
			StageSetFormDivSelectFieldToField(&inst.SectionNumber, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoNumberType](), formDiv)
		case "Slope":
			StageSetFormDivSelectFieldToField(&inst.Slope, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValSlopeType](), formDiv)
		case "AssessedAltitude":
			StageSetFormDivSelectFieldToField(&inst.AssessedAltitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "SlopeLowerAltitude":
			StageSetFormDivSelectFieldToField(&inst.SlopeLowerAltitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "GradientLowHigh":
			StageSetFormDivSelectFieldToField(&inst.GradientLowHigh, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValSlopeType](), formDiv)
		case "SurfaceZone":
			StageSetFormDivSelectFieldToField(&inst.SurfaceZone, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeObstructionIdSurfaceZoneType](), formDiv)
		case "SafetyRegulation":
			StageSetFormDivSelectFieldToField(&inst.SafetyRegulation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextNameType](), formDiv)
		case "Surface":
			StageSetFormDivSelectFieldToField(&inst.Surface, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SurfacePropertyType](), formDiv)
		case "StartingCurve":
			StageSetFormDivSelectFieldToField(&inst.StartingCurve, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CurvePropertyType](), formDiv)
		}
	}
}

func saveStageSet_ObstaclePlacementPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.ObstaclePlacementPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "ObstaclePlacement":
			StageSetFormDivSelectFieldToField(&inst.ObstaclePlacement, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ObstaclePlacementType](), formDiv)
		}
	}
}

func saveStageSet_ObstaclePlacementType_AixmfeaturesStage(
	inst *aixmfeatures.ObstaclePlacementType,
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
		case "ObstacleBearing":
			StageSetFormDivSelectFieldToField(&inst.ObstacleBearing, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "ObstacleDistance":
			StageSetFormDivSelectFieldToField(&inst.ObstacleDistance, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "PointType":
			StageSetFormDivSelectFieldToField(&inst.PointType, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextNameType](), formDiv)
		case "ObstaclePlacement":
			StageSetFormDivSelectFieldToField(&inst.ObstaclePlacement, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSideType](), formDiv)
		}
	}
}

func saveStageSet_ObstructionPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.ObstructionPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "Obstruction":
			StageSetFormDivSelectFieldToField(&inst.Obstruction, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ObstructionType](), formDiv)
		}
	}
}

func saveStageSet_ObstructionType_AixmfeaturesStage(
	inst *aixmfeatures.ObstructionType,
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
		case "RequiredClearance":
			StageSetFormDivSelectFieldToField(&inst.RequiredClearance, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "MinimumAltitude":
			StageSetFormDivSelectFieldToField(&inst.MinimumAltitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "SurfacePenetration":
			StageSetFormDivSelectFieldToField(&inst.SurfacePenetration, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "SlopePenetration":
			StageSetFormDivSelectFieldToField(&inst.SlopePenetration, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), formDiv)
		case "Controlling":
			StageSetFormDivSelectFieldToField(&inst.Controlling, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "CloseIn":
			StageSetFormDivSelectFieldToField(&inst.CloseIn, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "TheVerticalStructure":
			StageSetFormDivSelectFieldToField(&inst.TheVerticalStructure, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.VerticalStructurePropertyType](), formDiv)
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

func saveStageSet_PointReferencePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.PointReferencePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "PointReference":
			StageSetFormDivSelectFieldToField(&inst.PointReference, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointReferenceType](), formDiv)
		}
	}
}

func saveStageSet_PointReferenceType_AixmfeaturesStage(
	inst *aixmfeatures.PointReferenceType,
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
		case "Role":
			StageSetFormDivSelectFieldToField(&inst.Role, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeReferenceRoleType](), formDiv)
		case "PriorFixTolerance":
			StageSetFormDivSelectFieldToField(&inst.PriorFixTolerance, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceSignedType](), formDiv)
		case "PostFixTolerance":
			StageSetFormDivSelectFieldToField(&inst.PostFixTolerance, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceSignedType](), formDiv)
		case "Point":
			StageSetFormDivSelectFieldToField(&inst.Point, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), formDiv)
		case "FixToleranceArea":
			StageSetFormDivSelectFieldToField(&inst.FixToleranceArea, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SurfacePropertyType](), formDiv)
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

func saveStageSet_ProcedureAvailabilityPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.ProcedureAvailabilityPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "ProcedureAvailability":
			StageSetFormDivSelectFieldToField(&inst.ProcedureAvailability, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ProcedureAvailabilityType](), formDiv)
		}
	}
}

func saveStageSet_ProcedureAvailabilityType_AixmfeaturesStage(
	inst *aixmfeatures.ProcedureAvailabilityType,
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
		case "Status":
			StageSetFormDivSelectFieldToField(&inst.Status, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeProcedureAvailabilityType](), formDiv)
		}
	}
}

func saveStageSet_ProcedureDMETimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.ProcedureDMETimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "ProcedureDMETimeSlice":
			StageSetFormDivSelectFieldToField(&inst.ProcedureDMETimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ProcedureDMETimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_ProcedureDMETimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.ProcedureDMETimeSliceType,
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
		case "CriticalDME":
			StageSetFormDivSelectFieldToField(&inst.CriticalDME, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "Satisfactory":
			StageSetFormDivSelectFieldToField(&inst.Satisfactory, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "DME":
			StageSetFormDivSelectFieldToField(&inst.DME, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DMEPropertyType](), formDiv)
		case "SegmentLeg":
			StageSetFormDivSelectFieldToField(&inst.SegmentLeg, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SegmentLegPropertyType](), formDiv)
		}
	}
}

func saveStageSet_ProcedureDMEType_AixmfeaturesStage(
	inst *aixmfeatures.ProcedureDMEType,
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

func saveStageSet_ProcedurePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.ProcedurePropertyType,
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

func saveStageSet_ProcedureTransitionLegPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.ProcedureTransitionLegPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "ProcedureTransitionLeg":
			StageSetFormDivSelectFieldToField(&inst.ProcedureTransitionLeg, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ProcedureTransitionLegType](), formDiv)
		}
	}
}

func saveStageSet_ProcedureTransitionLegType_AixmfeaturesStage(
	inst *aixmfeatures.ProcedureTransitionLegType,
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
		case "SeqNumberARINC":
			StageSetFormDivSelectFieldToField(&inst.SeqNumberARINC, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoSequenceType](), formDiv)
		case "TheSegmentLeg":
			StageSetFormDivSelectFieldToField(&inst.TheSegmentLeg, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SegmentLegPropertyType](), formDiv)
		}
	}
}

func saveStageSet_ProcedureTransitionPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.ProcedureTransitionPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "ProcedureTransition":
			StageSetFormDivSelectFieldToField(&inst.ProcedureTransition, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ProcedureTransitionType](), formDiv)
		}
	}
}

func saveStageSet_ProcedureTransitionType_AixmfeaturesStage(
	inst *aixmfeatures.ProcedureTransitionType,
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
		case "TransitionId":
			StageSetFormDivSelectFieldToField(&inst.TransitionId, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDesignatedPointDesignatorType](), formDiv)
		case "Type_XSD":
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeProcedurePhaseType](), formDiv)
		case "Instruction":
			StageSetFormDivSelectFieldToField(&inst.Instruction, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextInstructionType](), formDiv)
		case "VectorHeading":
			StageSetFormDivSelectFieldToField(&inst.VectorHeading, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "DepartureRunwayTransition":
			StageSetFormDivSelectFieldToField(&inst.DepartureRunwayTransition, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.LandingTakeoffAreaCollectionPropertyType](), formDiv)
		case "Trajectory":
			StageSetFormDivSelectFieldToField(&inst.Trajectory, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CurvePropertyType](), formDiv)
		}
	}
}

func saveStageSet_RadarSystemPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.RadarSystemPropertyType,
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

func saveStageSet_RadioFrequencyAreaPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.RadioFrequencyAreaPropertyType,
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

func saveStageSet_SafeAltitudeAreaPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.SafeAltitudeAreaPropertyType,
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

func saveStageSet_SafeAltitudeAreaSectorPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.SafeAltitudeAreaSectorPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "SafeAltitudeAreaSector":
			StageSetFormDivSelectFieldToField(&inst.SafeAltitudeAreaSector, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SafeAltitudeAreaSectorType](), formDiv)
		}
	}
}

func saveStageSet_SafeAltitudeAreaSectorType_AixmfeaturesStage(
	inst *aixmfeatures.SafeAltitudeAreaSectorType,
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
		case "BufferWidth":
			StageSetFormDivSelectFieldToField(&inst.BufferWidth, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "Extent":
			StageSetFormDivSelectFieldToField(&inst.Extent, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SurfacePropertyType](), formDiv)
		case "SectorDefinition":
			StageSetFormDivSelectFieldToField(&inst.SectorDefinition, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CircleSectorPropertyType](), formDiv)
		}
	}
}

func saveStageSet_SafeAltitudeAreaTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.SafeAltitudeAreaTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "SafeAltitudeAreaTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.SafeAltitudeAreaTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SafeAltitudeAreaTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_SafeAltitudeAreaTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.SafeAltitudeAreaTimeSliceType,
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
		case "SafeAreaType":
			StageSetFormDivSelectFieldToField(&inst.SafeAreaType, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSafeAltitudeType](), formDiv)
		case "CentrePoint_fixDesignatedPoint":
			StageSetFormDivSelectFieldToField(&inst.CentrePoint_fixDesignatedPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), formDiv)
		case "CentrePoint_navaidSystem":
			StageSetFormDivSelectFieldToField(&inst.CentrePoint_navaidSystem, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), formDiv)
		case "CentrePoint_position":
			StageSetFormDivSelectFieldToField(&inst.CentrePoint_position, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), formDiv)
		case "CentrePoint_runwayPoint":
			StageSetFormDivSelectFieldToField(&inst.CentrePoint_runwayPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), formDiv)
		case "CentrePoint_aimingPoint":
			StageSetFormDivSelectFieldToField(&inst.CentrePoint_aimingPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), formDiv)
		case "CentrePoint_airportReferencePoint":
			StageSetFormDivSelectFieldToField(&inst.CentrePoint_airportReferencePoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), formDiv)
		}
	}
}

func saveStageSet_SafeAltitudeAreaType_AixmfeaturesStage(
	inst *aixmfeatures.SafeAltitudeAreaType,
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

func saveStageSet_SectorDesignPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.SectorDesignPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "SectorDesign":
			StageSetFormDivSelectFieldToField(&inst.SectorDesign, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SectorDesignType](), formDiv)
		}
	}
}

func saveStageSet_SectorDesignType_AixmfeaturesStage(
	inst *aixmfeatures.SectorDesignType,
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
		case "TurnDirection":
			StageSetFormDivSelectFieldToField(&inst.TurnDirection, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDirectionTurnType](), formDiv)
		case "DesignGradient":
			StageSetFormDivSelectFieldToField(&inst.DesignGradient, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValSlopeType](), formDiv)
		case "TerminationAltitude":
			StageSetFormDivSelectFieldToField(&inst.TerminationAltitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "TurnPermitted":
			StageSetFormDivSelectFieldToField(&inst.TurnPermitted, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		}
	}
}

func saveStageSet_SegmentLegPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.SegmentLegPropertyType,
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

func saveStageSet_SegmentPointPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.SegmentPointPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "AerialRefuellingPoint":
			StageSetFormDivSelectFieldToField(&inst.AerialRefuellingPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AerialRefuellingPointType](), formDiv)
		case "EnRouteSegmentPoint":
			StageSetFormDivSelectFieldToField(&inst.EnRouteSegmentPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.EnRouteSegmentPointType](), formDiv)
		case "TerminalSegmentPoint":
			StageSetFormDivSelectFieldToField(&inst.TerminalSegmentPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TerminalSegmentPointType](), formDiv)
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

func saveStageSet_StandardInstrumentArrivalPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.StandardInstrumentArrivalPropertyType,
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

func saveStageSet_StandardInstrumentArrivalTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.StandardInstrumentArrivalTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "StandardInstrumentArrivalTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.StandardInstrumentArrivalTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.StandardInstrumentArrivalTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_StandardInstrumentArrivalTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.StandardInstrumentArrivalTimeSliceType,
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
		case "CommunicationFailureInstruction":
			StageSetFormDivSelectFieldToField(&inst.CommunicationFailureInstruction, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextInstructionType](), formDiv)
		case "Instruction":
			StageSetFormDivSelectFieldToField(&inst.Instruction, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextInstructionType](), formDiv)
		case "DesignCriteria":
			StageSetFormDivSelectFieldToField(&inst.DesignCriteria, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDesignStandardType](), formDiv)
		case "CodingStandard":
			StageSetFormDivSelectFieldToField(&inst.CodingStandard, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeProcedureCodingStandardType](), formDiv)
		case "FlightChecked":
			StageSetFormDivSelectFieldToField(&inst.FlightChecked, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "RNAV":
			StageSetFormDivSelectFieldToField(&inst.RNAV, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "GuidanceFacility_specialNavigationSystem":
			StageSetFormDivSelectFieldToField(&inst.GuidanceFacility_specialNavigationSystem, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SpecialNavigationSystemPropertyType](), formDiv)
		case "GuidanceFacility_radar":
			StageSetFormDivSelectFieldToField(&inst.GuidanceFacility_radar, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RadarSystemPropertyType](), formDiv)
		case "GuidanceFacility_navaid":
			StageSetFormDivSelectFieldToField(&inst.GuidanceFacility_navaid, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), formDiv)
		case "SafeAltitude":
			StageSetFormDivSelectFieldToField(&inst.SafeAltitude, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SafeAltitudeAreaPropertyType](), formDiv)
		case "Designator":
			StageSetFormDivSelectFieldToField(&inst.Designator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextSIDSTARDesignatorType](), formDiv)
		case "Arrival":
			StageSetFormDivSelectFieldToField(&inst.Arrival, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.LandingTakeoffAreaCollectionPropertyType](), formDiv)
		}
	}
}

func saveStageSet_StandardInstrumentArrivalType_AixmfeaturesStage(
	inst *aixmfeatures.StandardInstrumentArrivalType,
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

func saveStageSet_StandardInstrumentDeparturePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.StandardInstrumentDeparturePropertyType,
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

func saveStageSet_StandardInstrumentDepartureTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.StandardInstrumentDepartureTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "StandardInstrumentDepartureTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.StandardInstrumentDepartureTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.StandardInstrumentDepartureTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_StandardInstrumentDepartureTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.StandardInstrumentDepartureTimeSliceType,
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
		case "CommunicationFailureInstruction":
			StageSetFormDivSelectFieldToField(&inst.CommunicationFailureInstruction, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextInstructionType](), formDiv)
		case "Instruction":
			StageSetFormDivSelectFieldToField(&inst.Instruction, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextInstructionType](), formDiv)
		case "DesignCriteria":
			StageSetFormDivSelectFieldToField(&inst.DesignCriteria, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDesignStandardType](), formDiv)
		case "CodingStandard":
			StageSetFormDivSelectFieldToField(&inst.CodingStandard, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeProcedureCodingStandardType](), formDiv)
		case "FlightChecked":
			StageSetFormDivSelectFieldToField(&inst.FlightChecked, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "RNAV":
			StageSetFormDivSelectFieldToField(&inst.RNAV, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "GuidanceFacility_specialNavigationSystem":
			StageSetFormDivSelectFieldToField(&inst.GuidanceFacility_specialNavigationSystem, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SpecialNavigationSystemPropertyType](), formDiv)
		case "GuidanceFacility_radar":
			StageSetFormDivSelectFieldToField(&inst.GuidanceFacility_radar, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RadarSystemPropertyType](), formDiv)
		case "GuidanceFacility_navaid":
			StageSetFormDivSelectFieldToField(&inst.GuidanceFacility_navaid, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), formDiv)
		case "SafeAltitude":
			StageSetFormDivSelectFieldToField(&inst.SafeAltitude, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SafeAltitudeAreaPropertyType](), formDiv)
		case "Designator":
			StageSetFormDivSelectFieldToField(&inst.Designator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextSIDSTARDesignatorType](), formDiv)
		case "ContingencyRoute":
			StageSetFormDivSelectFieldToField(&inst.ContingencyRoute, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "Takeoff":
			StageSetFormDivSelectFieldToField(&inst.Takeoff, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.LandingTakeoffAreaCollectionPropertyType](), formDiv)
		}
	}
}

func saveStageSet_StandardInstrumentDepartureType_AixmfeaturesStage(
	inst *aixmfeatures.StandardInstrumentDepartureType,
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

func saveStageSet_SurfacePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.SurfacePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "Surface":
			StageSetFormDivSelectFieldToField(&inst.Surface, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SurfaceType](), formDiv)
		case "ElevatedSurface":
			StageSetFormDivSelectFieldToField(&inst.ElevatedSurface, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedSurfaceType](), formDiv)
		}
	}
}

func saveStageSet_SurfaceType_AixmfeaturesStage(
	inst *aixmfeatures.SurfaceType,
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
		case "Patches":
			StageSetFormDivSelectFieldToField(&inst.Patches, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.SurfacePatchArrayPropertyType](), formDiv)
		case "HorizontalAccuracy":
			StageSetFormDivSelectFieldToField(&inst.HorizontalAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		}
	}
}

func saveStageSet_TerminalArrivalAreaSectorPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.TerminalArrivalAreaSectorPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "TerminalArrivalAreaSector":
			StageSetFormDivSelectFieldToField(&inst.TerminalArrivalAreaSector, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TerminalArrivalAreaSectorType](), formDiv)
		}
	}
}

func saveStageSet_TerminalArrivalAreaSectorType_AixmfeaturesStage(
	inst *aixmfeatures.TerminalArrivalAreaSectorType,
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
		case "FlyByCode":
			StageSetFormDivSelectFieldToField(&inst.FlyByCode, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "ProcedureTurnRequired":
			StageSetFormDivSelectFieldToField(&inst.ProcedureTurnRequired, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "AltitudeDescription":
			StageSetFormDivSelectFieldToField(&inst.AltitudeDescription, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAltitudeUseType](), formDiv)
		case "SectorDefinition":
			StageSetFormDivSelectFieldToField(&inst.SectorDefinition, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CircleSectorPropertyType](), formDiv)
		case "Extent":
			StageSetFormDivSelectFieldToField(&inst.Extent, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SurfacePropertyType](), formDiv)
		}
	}
}

func saveStageSet_TerminalArrivalAreaTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.TerminalArrivalAreaTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "TerminalArrivalAreaTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.TerminalArrivalAreaTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TerminalArrivalAreaTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_TerminalArrivalAreaTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.TerminalArrivalAreaTimeSliceType,
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
		case "ArrivalAreaType":
			StageSetFormDivSelectFieldToField(&inst.ArrivalAreaType, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTAAType](), formDiv)
		case "OuterBufferWidth":
			StageSetFormDivSelectFieldToField(&inst.OuterBufferWidth, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "LateralBufferWidth":
			StageSetFormDivSelectFieldToField(&inst.LateralBufferWidth, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "IF_fixDesignatedPoint":
			StageSetFormDivSelectFieldToField(&inst.IF_fixDesignatedPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), formDiv)
		case "IF_navaidSystem":
			StageSetFormDivSelectFieldToField(&inst.IF_navaidSystem, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), formDiv)
		case "IF_position":
			StageSetFormDivSelectFieldToField(&inst.IF_position, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), formDiv)
		case "IF_runwayPoint":
			StageSetFormDivSelectFieldToField(&inst.IF_runwayPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), formDiv)
		case "IF_aimingPoint":
			StageSetFormDivSelectFieldToField(&inst.IF_aimingPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), formDiv)
		case "IF_airportReferencePoint":
			StageSetFormDivSelectFieldToField(&inst.IF_airportReferencePoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), formDiv)
		case "IAF_fixDesignatedPoint":
			StageSetFormDivSelectFieldToField(&inst.IAF_fixDesignatedPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), formDiv)
		case "IAF_navaidSystem":
			StageSetFormDivSelectFieldToField(&inst.IAF_navaidSystem, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), formDiv)
		case "IAF_position":
			StageSetFormDivSelectFieldToField(&inst.IAF_position, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), formDiv)
		case "IAF_runwayPoint":
			StageSetFormDivSelectFieldToField(&inst.IAF_runwayPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), formDiv)
		case "IAF_aimingPoint":
			StageSetFormDivSelectFieldToField(&inst.IAF_aimingPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), formDiv)
		case "IAF_airportReferencePoint":
			StageSetFormDivSelectFieldToField(&inst.IAF_airportReferencePoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), formDiv)
		case "Buffer":
			StageSetFormDivSelectFieldToField(&inst.Buffer, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SurfacePropertyType](), formDiv)
		case "ApproachRNAV":
			StageSetFormDivSelectFieldToField(&inst.ApproachRNAV, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.InstrumentApproachProcedurePropertyType](), formDiv)
		}
	}
}

func saveStageSet_TerminalArrivalAreaType_AixmfeaturesStage(
	inst *aixmfeatures.TerminalArrivalAreaType,
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

func saveStageSet_TerminalSegmentPointPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.TerminalSegmentPointPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "TerminalSegmentPoint":
			StageSetFormDivSelectFieldToField(&inst.TerminalSegmentPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TerminalSegmentPointType](), formDiv)
		}
	}
}

func saveStageSet_TerminalSegmentPointType_AixmfeaturesStage(
	inst *aixmfeatures.TerminalSegmentPointType,
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
		case "ReportingATC":
			StageSetFormDivSelectFieldToField(&inst.ReportingATC, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeATCReportingType](), formDiv)
		case "FlyOver":
			StageSetFormDivSelectFieldToField(&inst.FlyOver, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "Waypoint":
			StageSetFormDivSelectFieldToField(&inst.Waypoint, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "RadarGuidance":
			StageSetFormDivSelectFieldToField(&inst.RadarGuidance, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
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
		case "ExtendedServiceVolume":
			StageSetFormDivSelectFieldToField(&inst.ExtendedServiceVolume, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RadioFrequencyAreaPropertyType](), formDiv)
		case "Role":
			StageSetFormDivSelectFieldToField(&inst.Role, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeProcedureFixRoleType](), formDiv)
		case "LeadRadial":
			StageSetFormDivSelectFieldToField(&inst.LeadRadial, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "LeadDME":
			StageSetFormDivSelectFieldToField(&inst.LeadDME, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "IndicatorFACF":
			StageSetFormDivSelectFieldToField(&inst.IndicatorFACF, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
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

func saveStageSet_UnplannedHoldingPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.UnplannedHoldingPropertyType,
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

func saveStageSet_UnplannedHoldingTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.UnplannedHoldingTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "UnplannedHoldingTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.UnplannedHoldingTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.UnplannedHoldingTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_UnplannedHoldingTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.UnplannedHoldingTimeSliceType,
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
		case "UnplannedHolding":
			StageSetFormDivSelectFieldToField(&inst.UnplannedHolding, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeApprovalType](), formDiv)
		case "AuthorizedAltitude":
			StageSetFormDivSelectFieldToField(&inst.AuthorizedAltitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "AltitudeReference":
			StageSetFormDivSelectFieldToField(&inst.AltitudeReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "ControlledAirspace":
			StageSetFormDivSelectFieldToField(&inst.ControlledAirspace, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "HoldingPoint":
			StageSetFormDivSelectFieldToField(&inst.HoldingPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SegmentPointPropertyType](), formDiv)
		}
	}
}

func saveStageSet_UnplannedHoldingType_AixmfeaturesStage(
	inst *aixmfeatures.UnplannedHoldingType,
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

func saveStageSet_VerticalStructurePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.VerticalStructurePropertyType,
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

func saveStageSet_CodeATCReportingType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeATCReportingType,
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

func saveStageSet_CodeAerialRefuellingPointType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeAerialRefuellingPointType,
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

func saveStageSet_CodeAircraftCategoryType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeAircraftCategoryType,
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

func saveStageSet_CodeAircraftEngineNumberType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeAircraftEngineNumberType,
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

func saveStageSet_CodeAircraftEngineType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeAircraftEngineType,
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

func saveStageSet_CodeAircraftICAOType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeAircraftICAOType,
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

func saveStageSet_CodeAircraftType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeAircraftType,
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

func saveStageSet_CodeAircraftWingspanClassType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeAircraftWingspanClassType,
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

func saveStageSet_CodeAltitudeAdjustmentType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeAltitudeAdjustmentType,
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

func saveStageSet_CodeApproachEquipmentAdditionalType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeApproachEquipmentAdditionalType,
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

func saveStageSet_CodeApproachGuidanceType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeApproachGuidanceType,
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

func saveStageSet_CodeApproachPrefixType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeApproachPrefixType,
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

func saveStageSet_CodeApproachType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeApproachType,
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

func saveStageSet_CodeApprovalType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeApprovalType,
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

func saveStageSet_CodeArcDirectionType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeArcDirectionType,
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

func saveStageSet_CodeCommunicationModeType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeCommunicationModeType,
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

func saveStageSet_CodeCourseType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeCourseType,
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

func saveStageSet_CodeDesignStandardType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeDesignStandardType,
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

func saveStageSet_CodeDirectionTurnType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeDirectionTurnType,
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

func saveStageSet_CodeEquipmentAntiCollisionType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeEquipmentAntiCollisionType,
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

func saveStageSet_CodeEquipmentUnavailableType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeEquipmentUnavailableType,
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

func saveStageSet_CodeFinalGuidanceType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeFinalGuidanceType,
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

func saveStageSet_CodeFreeFlightType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeFreeFlightType,
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

func saveStageSet_CodeHeightReferenceType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeHeightReferenceType,
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

func saveStageSet_CodeHoldingUsageType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeHoldingUsageType,
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

func saveStageSet_CodeHoldingUseType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeHoldingUseType,
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

func saveStageSet_CodeICAOCountryType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeICAOCountryType,
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

func saveStageSet_CodeMilitaryRoutePointType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeMilitaryRoutePointType,
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

func saveStageSet_CodeMinimaFinalApproachPathType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeMinimaFinalApproachPathType,
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

func saveStageSet_CodeMinimumAltitudeType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeMinimumAltitudeType,
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

func saveStageSet_CodeMinimumHeightType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeMinimumHeightType,
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

func saveStageSet_CodeMissedApproachType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeMissedApproachType,
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

func saveStageSet_CodeNavigationAreaRestrictionType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeNavigationAreaRestrictionType,
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

func saveStageSet_CodeNavigationAreaType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeNavigationAreaType,
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

func saveStageSet_CodeNavigationEquipmentType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeNavigationEquipmentType,
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

func saveStageSet_CodeNavigationSpecificationType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeNavigationSpecificationType,
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

func saveStageSet_CodeObstacleAssessmentSurfaceType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeObstacleAssessmentSurfaceType,
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

func saveStageSet_CodeObstructionIdSurfaceZoneType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeObstructionIdSurfaceZoneType,
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

func saveStageSet_CodeProcedureAvailabilityType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeProcedureAvailabilityType,
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

func saveStageSet_CodeProcedureCodingStandardType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeProcedureCodingStandardType,
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

func saveStageSet_CodeProcedureDistanceType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeProcedureDistanceType,
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

func saveStageSet_CodeProcedureFixRoleType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeProcedureFixRoleType,
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

func saveStageSet_CodeProcedurePhaseType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeProcedurePhaseType,
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

func saveStageSet_CodeRNPType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeRNPType,
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

func saveStageSet_CodeRVSMPointRoleType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeRVSMPointRoleType,
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

func saveStageSet_CodeRVSMType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeRVSMType,
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

func saveStageSet_CodeReferencePathIdentifierType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeReferencePathIdentifierType,
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

func saveStageSet_CodeReferenceRoleType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeReferenceRoleType,
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

func saveStageSet_CodeRelativePositionType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeRelativePositionType,
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

func saveStageSet_CodeRouteIndicatorType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeRouteIndicatorType,
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

func saveStageSet_CodeSafeAltitudeType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeSafeAltitudeType,
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

func saveStageSet_CodeSegmentPathType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeSegmentPathType,
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

func saveStageSet_CodeSegmentTerminationType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeSegmentTerminationType,
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

func saveStageSet_CodeSideType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeSideType,
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

func saveStageSet_CodeSpeedReferenceType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeSpeedReferenceType,
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

func saveStageSet_CodeTAAType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeTAAType,
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

func saveStageSet_CodeTrajectoryType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeTrajectoryType,
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

func saveStageSet_CodeUpperAlphaType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeUpperAlphaType,
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

func saveStageSet_CodeValueInterpretationType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeValueInterpretationType,
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

func saveStageSet_CodeWakeTurbulenceType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeWakeTurbulenceType,
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

func saveStageSet_NoNumberType_AixmdatatypesStage(
	inst *aixmdatatypes.NoNumberType,
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

func saveStageSet_TextInstructionType_AixmdatatypesStage(
	inst *aixmdatatypes.TextInstructionType,
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

func saveStageSet_TextSIDSTARDesignatorType_AixmdatatypesStage(
	inst *aixmdatatypes.TextSIDSTARDesignatorType,
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

func saveStageSet_ValAlarmLimitType_AixmdatatypesStage(
	inst *aixmdatatypes.ValAlarmLimitType,
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

func saveStageSet_ValChannelNumberType_AixmdatatypesStage(
	inst *aixmdatatypes.ValChannelNumberType,
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

func saveStageSet_ValHexType_AixmdatatypesStage(
	inst *aixmdatatypes.ValHexType,
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

func saveStageSet_ValSlopeType_AixmdatatypesStage(
	inst *aixmdatatypes.ValSlopeType,
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

func saveStageSet_ValSpeedType_AixmdatatypesStage(
	inst *aixmdatatypes.ValSpeedType,
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

func saveStageSet_ValTemperatureType_AixmdatatypesStage(
	inst *aixmdatatypes.ValTemperatureType,
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

func saveStageSet_ValWeightType_AixmdatatypesStage(
	inst *aixmdatatypes.ValWeightType,
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

func saveStageSet_AbstractCurveSegmentType_GmlStage(
	inst *gml.AbstractCurveSegmentType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NumDerivativesAtStart":
			FormDivBasicFieldToField(&inst.NumDerivativesAtStart, formDiv)
		case "NumDerivativesAtEnd":
			FormDivBasicFieldToField(&inst.NumDerivativesAtEnd, formDiv)
		case "NumDerivativeInterior":
			FormDivBasicFieldToField(&inst.NumDerivativeInterior, formDiv)
		}
	}
}

func saveStageSet_AbstractCurveType_GmlStage(
	inst *gml.AbstractCurveType,
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

func saveStageSet_AbstractRingPropertyType_GmlStage(
	inst *gml.AbstractRingPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Ring":
			StageSetFormDivSelectFieldToField(&inst.Ring, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.RingType](), formDiv)
		}
	}
}

func saveStageSet_AbstractRingType_GmlStage(
	inst *gml.AbstractRingType,
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

func saveStageSet_AbstractSurfacePatchType_GmlStage(
	inst *gml.AbstractSurfacePatchType,
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

func saveStageSet_AbstractSurfaceType_GmlStage(
	inst *gml.AbstractSurfaceType,
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

func saveStageSet_AngleType_GmlStage(
	inst *gml.AngleType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_ArcByCenterPointType_GmlStage(
	inst *gml.ArcByCenterPointType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NumDerivativesAtStart":
			FormDivBasicFieldToField(&inst.NumDerivativesAtStart, formDiv)
		case "NumDerivativesAtEnd":
			FormDivBasicFieldToField(&inst.NumDerivativesAtEnd, formDiv)
		case "NumDerivativeInterior":
			FormDivBasicFieldToField(&inst.NumDerivativeInterior, formDiv)
		case "Interpolation":
			StageSetFormDivEnumStringFieldToField(&inst.Interpolation, formDiv)
		case "NumArc":
			FormDivBasicFieldToField(&inst.NumArc, formDiv)
		case "Pos":
			StageSetFormDivSelectFieldToField(&inst.Pos, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionType](), formDiv)
		case "PointProperty":
			StageSetFormDivSelectFieldToField(&inst.PointProperty, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.PointPropertyType](), formDiv)
		case "PosList":
			StageSetFormDivSelectFieldToField(&inst.PosList, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionListType](), formDiv)
		case "Radius":
			StageSetFormDivSelectFieldToField(&inst.Radius, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.LengthType](), formDiv)
		case "StartAngle":
			StageSetFormDivSelectFieldToField(&inst.StartAngle, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.AngleType](), formDiv)
		case "EndAngle":
			StageSetFormDivSelectFieldToField(&inst.EndAngle, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.AngleType](), formDiv)
		}
	}
}

func saveStageSet_ArcStringType_GmlStage(
	inst *gml.ArcStringType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NumDerivativesAtStart":
			FormDivBasicFieldToField(&inst.NumDerivativesAtStart, formDiv)
		case "NumDerivativesAtEnd":
			FormDivBasicFieldToField(&inst.NumDerivativesAtEnd, formDiv)
		case "NumDerivativeInterior":
			FormDivBasicFieldToField(&inst.NumDerivativeInterior, formDiv)
		case "Interpolation":
			StageSetFormDivEnumStringFieldToField(&inst.Interpolation, formDiv)
		case "NumArc":
			FormDivBasicFieldToField(&inst.NumArc, formDiv)
		case "PosList":
			StageSetFormDivSelectFieldToField(&inst.PosList, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionListType](), formDiv)
		}
	}
}

func saveStageSet_ArcType_GmlStage(
	inst *gml.ArcType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NumDerivativesAtStart":
			FormDivBasicFieldToField(&inst.NumDerivativesAtStart, formDiv)
		case "NumDerivativesAtEnd":
			FormDivBasicFieldToField(&inst.NumDerivativesAtEnd, formDiv)
		case "NumDerivativeInterior":
			FormDivBasicFieldToField(&inst.NumDerivativeInterior, formDiv)
		case "Interpolation":
			StageSetFormDivEnumStringFieldToField(&inst.Interpolation, formDiv)
		case "NumArc":
			FormDivBasicFieldToField(&inst.NumArc, formDiv)
		case "Pos":
			StageSetFormDivSelectFieldToField(&inst.Pos, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionType](), formDiv)
		case "PointProperty":
			StageSetFormDivSelectFieldToField(&inst.PointProperty, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.PointPropertyType](), formDiv)
		case "PosList":
			StageSetFormDivSelectFieldToField(&inst.PosList, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionListType](), formDiv)
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

func saveStageSet_CircleByCenterPointType_GmlStage(
	inst *gml.CircleByCenterPointType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NumDerivativesAtStart":
			FormDivBasicFieldToField(&inst.NumDerivativesAtStart, formDiv)
		case "NumDerivativesAtEnd":
			FormDivBasicFieldToField(&inst.NumDerivativesAtEnd, formDiv)
		case "NumDerivativeInterior":
			FormDivBasicFieldToField(&inst.NumDerivativeInterior, formDiv)
		case "Interpolation":
			StageSetFormDivEnumStringFieldToField(&inst.Interpolation, formDiv)
		case "NumArc":
			FormDivBasicFieldToField(&inst.NumArc, formDiv)
		case "Pos":
			StageSetFormDivSelectFieldToField(&inst.Pos, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionType](), formDiv)
		case "PointProperty":
			StageSetFormDivSelectFieldToField(&inst.PointProperty, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.PointPropertyType](), formDiv)
		case "PosList":
			StageSetFormDivSelectFieldToField(&inst.PosList, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionListType](), formDiv)
		case "Radius":
			StageSetFormDivSelectFieldToField(&inst.Radius, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.LengthType](), formDiv)
		}
	}
}

func saveStageSet_CircleType_GmlStage(
	inst *gml.CircleType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NumDerivativesAtStart":
			FormDivBasicFieldToField(&inst.NumDerivativesAtStart, formDiv)
		case "NumDerivativesAtEnd":
			FormDivBasicFieldToField(&inst.NumDerivativesAtEnd, formDiv)
		case "NumDerivativeInterior":
			FormDivBasicFieldToField(&inst.NumDerivativeInterior, formDiv)
		case "Interpolation":
			StageSetFormDivEnumStringFieldToField(&inst.Interpolation, formDiv)
		case "NumArc":
			FormDivBasicFieldToField(&inst.NumArc, formDiv)
		case "PosList":
			StageSetFormDivSelectFieldToField(&inst.PosList, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionListType](), formDiv)
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

func saveStageSet_CompositeCurveType_GmlStage(
	inst *gml.CompositeCurveType,
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
		case "AggregationType":
			FormDivBasicFieldToField(&inst.AggregationType, formDiv)
		}
	}
}

func saveStageSet_CurvePropertyType_GmlStage(
	inst *gml.CurvePropertyType,
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
		case "Curve":
			StageSetFormDivSelectFieldToField(&inst.Curve, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CurveType](), formDiv)
		case "OrientableCurve":
			StageSetFormDivSelectFieldToField(&inst.OrientableCurve, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.OrientableCurveType](), formDiv)
		case "CompositeCurve":
			StageSetFormDivSelectFieldToField(&inst.CompositeCurve, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CompositeCurveType](), formDiv)
		}
	}
}

func saveStageSet_CurveSegmentArrayPropertyType_GmlStage(
	inst *gml.CurveSegmentArrayPropertyType,
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

func saveStageSet_CurveType_GmlStage(
	inst *gml.CurveType,
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
		case "Segments":
			StageSetFormDivSelectFieldToField(&inst.Segments, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CurveSegmentArrayPropertyType](), formDiv)
		}
	}
}

func saveStageSet_DirectPositionListType_GmlStage(
	inst *gml.DirectPositionListType,
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

func saveStageSet_GeodesicStringType_GmlStage(
	inst *gml.GeodesicStringType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NumDerivativesAtStart":
			FormDivBasicFieldToField(&inst.NumDerivativesAtStart, formDiv)
		case "NumDerivativesAtEnd":
			FormDivBasicFieldToField(&inst.NumDerivativesAtEnd, formDiv)
		case "NumDerivativeInterior":
			FormDivBasicFieldToField(&inst.NumDerivativeInterior, formDiv)
		case "Interpolation":
			StageSetFormDivEnumStringFieldToField(&inst.Interpolation, formDiv)
		case "PosList":
			StageSetFormDivSelectFieldToField(&inst.PosList, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionListType](), formDiv)
		case "Pos":
			StageSetFormDivSelectFieldToField(&inst.Pos, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionType](), formDiv)
		case "PointProperty":
			StageSetFormDivSelectFieldToField(&inst.PointProperty, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.PointPropertyType](), formDiv)
		}
	}
}

func saveStageSet_GeodesicType_GmlStage(
	inst *gml.GeodesicType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NumDerivativesAtStart":
			FormDivBasicFieldToField(&inst.NumDerivativesAtStart, formDiv)
		case "NumDerivativesAtEnd":
			FormDivBasicFieldToField(&inst.NumDerivativesAtEnd, formDiv)
		case "NumDerivativeInterior":
			FormDivBasicFieldToField(&inst.NumDerivativeInterior, formDiv)
		case "Interpolation":
			StageSetFormDivEnumStringFieldToField(&inst.Interpolation, formDiv)
		case "PosList":
			StageSetFormDivSelectFieldToField(&inst.PosList, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionListType](), formDiv)
		case "Pos":
			StageSetFormDivSelectFieldToField(&inst.Pos, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionType](), formDiv)
		case "PointProperty":
			StageSetFormDivSelectFieldToField(&inst.PointProperty, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.PointPropertyType](), formDiv)
		}
	}
}

func saveStageSet_LengthType_GmlStage(
	inst *gml.LengthType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_LineStringSegmentType_GmlStage(
	inst *gml.LineStringSegmentType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NumDerivativesAtStart":
			FormDivBasicFieldToField(&inst.NumDerivativesAtStart, formDiv)
		case "NumDerivativesAtEnd":
			FormDivBasicFieldToField(&inst.NumDerivativesAtEnd, formDiv)
		case "NumDerivativeInterior":
			FormDivBasicFieldToField(&inst.NumDerivativeInterior, formDiv)
		case "Interpolation":
			StageSetFormDivEnumStringFieldToField(&inst.Interpolation, formDiv)
		case "PosList":
			StageSetFormDivSelectFieldToField(&inst.PosList, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionListType](), formDiv)
		}
	}
}

func saveStageSet_MeasureType_GmlStage(
	inst *gml.MeasureType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Uom":
			FormDivBasicFieldToField(&inst.Uom, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_OrientableCurveType_GmlStage(
	inst *gml.OrientableCurveType,
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
		case "Orientation":
			StageSetFormDivEnumStringFieldToField(&inst.Orientation, formDiv)
		case "BaseCurve":
			StageSetFormDivSelectFieldToField(&inst.BaseCurve, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CurvePropertyType](), formDiv)
		}
	}
}

func saveStageSet_PointPropertyType_GmlStage(
	inst *gml.PointPropertyType,
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
		case "Point":
			StageSetFormDivSelectFieldToField(&inst.Point, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.PointType](), formDiv)
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

func saveStageSet_PolygonPatchType_GmlStage(
	inst *gml.PolygonPatchType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Interpolation":
			StageSetFormDivEnumStringFieldToField(&inst.Interpolation, formDiv)
		case "Exterior":
			StageSetFormDivSelectFieldToField(&inst.Exterior, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.AbstractRingPropertyType](), formDiv)
		}
	}
}

func saveStageSet_RingType_GmlStage(
	inst *gml.RingType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "AggregationType":
			FormDivBasicFieldToField(&inst.AggregationType, formDiv)
		}
	}
}

func saveStageSet_SurfacePatchArrayPropertyType_GmlStage(
	inst *gml.SurfacePatchArrayPropertyType,
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

func saveStageSet_SurfaceType_GmlStage(
	inst *gml.SurfaceType,
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
		case "Patches":
			StageSetFormDivSelectFieldToField(&inst.Patches, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.SurfacePatchArrayPropertyType](), formDiv)
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

func StageSetNewInstance_A_extension_42_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_42",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_42)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_42_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_42_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_43_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_43",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_43)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_43_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_43_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_44_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_44",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_44)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_44_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_44_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_45_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_45",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_45)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_45_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_45_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_46_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_46",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_46)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_46_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_46_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_47_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_47",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_47)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_47_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_47_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_48_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_48",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_48)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_48_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_48_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_49_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_49",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_49)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_49_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_49_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_A_extension_50_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_50",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_50)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_50_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_50_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_51_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_51",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_51)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_51_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_51_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_52_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_52",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_52)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_52_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_52_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_53_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_53",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_53)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_53_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_53_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_54_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_54",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_54)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_54_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_54_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_55_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_55",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_55)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_55_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_55_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_56_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_56",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_56)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_56_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_56_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_57_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_57",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_57)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_57_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_57_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_extension_58_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_extension_58",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.A_extension_58)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_A_extension_58_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_A_extension_58_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_AbstractApproachLegType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractApproachLegType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AbstractApproachLegType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AbstractApproachLegType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AbstractApproachLegType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractProcedureType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractProcedureType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AbstractProcedureType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AbstractProcedureType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AbstractProcedureType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_AbstractSegmentLegType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractSegmentLegType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AbstractSegmentLegType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AbstractSegmentLegType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AbstractSegmentLegType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractSegmentPointType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractSegmentPointType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AbstractSegmentPointType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AbstractSegmentPointType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AbstractSegmentPointType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AerialRefuellingPointType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AerialRefuellingPointType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AerialRefuellingPointType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AerialRefuellingPointType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AerialRefuellingPointType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AircraftCharacteristicPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AircraftCharacteristicPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AircraftCharacteristicPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AircraftCharacteristicPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AircraftCharacteristicPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AircraftCharacteristicType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AircraftCharacteristicType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AircraftCharacteristicType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AircraftCharacteristicType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AircraftCharacteristicType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_AltimeterSourcePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AltimeterSourcePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AltimeterSourcePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AltimeterSourcePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AltimeterSourcePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AltitudeAdjustmentPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AltitudeAdjustmentPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AltitudeAdjustmentPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AltitudeAdjustmentPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AltitudeAdjustmentPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AltitudeAdjustmentType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AltitudeAdjustmentType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AltitudeAdjustmentType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AltitudeAdjustmentType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AltitudeAdjustmentType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AngleIndicationPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AngleIndicationPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AngleIndicationPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AngleIndicationPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AngleIndicationPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AngleUsePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AngleUsePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AngleUsePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AngleUsePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AngleUsePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AngleUseType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AngleUseType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AngleUseType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AngleUseType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AngleUseType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ApproachAltitudeTablePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ApproachAltitudeTablePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ApproachAltitudeTablePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ApproachAltitudeTablePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ApproachAltitudeTablePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ApproachAltitudeTableType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ApproachAltitudeTableType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ApproachAltitudeTableType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ApproachAltitudeTableType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ApproachAltitudeTableType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ApproachConditionPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ApproachConditionPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ApproachConditionPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ApproachConditionPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ApproachConditionPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ApproachConditionType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ApproachConditionType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ApproachConditionType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ApproachConditionType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ApproachConditionType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ApproachDistanceTablePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ApproachDistanceTablePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ApproachDistanceTablePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ApproachDistanceTablePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ApproachDistanceTablePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ApproachDistanceTableType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ApproachDistanceTableType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ApproachDistanceTableType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ApproachDistanceTableType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ApproachDistanceTableType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ApproachTimingTablePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ApproachTimingTablePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ApproachTimingTablePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ApproachTimingTablePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ApproachTimingTablePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ApproachTimingTableType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ApproachTimingTableType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ApproachTimingTableType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ApproachTimingTableType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ApproachTimingTableType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ArrivalFeederLegTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ArrivalFeederLegTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ArrivalFeederLegTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ArrivalFeederLegTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ArrivalFeederLegTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ArrivalFeederLegTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ArrivalFeederLegTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ArrivalFeederLegTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ArrivalFeederLegTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ArrivalFeederLegTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ArrivalFeederLegType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ArrivalFeederLegType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ArrivalFeederLegType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ArrivalFeederLegType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ArrivalFeederLegType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ArrivalLegTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ArrivalLegTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ArrivalLegTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ArrivalLegTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ArrivalLegTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ArrivalLegTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ArrivalLegTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ArrivalLegTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ArrivalLegTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ArrivalLegTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ArrivalLegType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ArrivalLegType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ArrivalLegType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ArrivalLegType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ArrivalLegType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CircleSectorPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CircleSectorPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.CircleSectorPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_CircleSectorPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_CircleSectorPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CircleSectorType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CircleSectorType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.CircleSectorType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_CircleSectorType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_CircleSectorType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CirclingAreaTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CirclingAreaTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.CirclingAreaTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_CirclingAreaTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_CirclingAreaTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CirclingAreaTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CirclingAreaTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.CirclingAreaTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_CirclingAreaTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_CirclingAreaTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CirclingAreaType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CirclingAreaType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.CirclingAreaType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_CirclingAreaType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_CirclingAreaType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CirclingRestrictionPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CirclingRestrictionPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.CirclingRestrictionPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_CirclingRestrictionPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_CirclingRestrictionPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CirclingRestrictionType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CirclingRestrictionType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.CirclingRestrictionType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_CirclingRestrictionType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_CirclingRestrictionType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CurvePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CurvePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.CurvePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_CurvePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_CurvePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CurveType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CurveType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.CurveType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_CurveType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_CurveType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DMEPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DMEPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.DMEPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_DMEPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_DMEPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DepartureArrivalConditionPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DepartureArrivalConditionPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.DepartureArrivalConditionPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_DepartureArrivalConditionPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_DepartureArrivalConditionPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DepartureArrivalConditionType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DepartureArrivalConditionType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.DepartureArrivalConditionType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_DepartureArrivalConditionType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_DepartureArrivalConditionType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DepartureLegTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DepartureLegTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.DepartureLegTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_DepartureLegTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_DepartureLegTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DepartureLegTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DepartureLegTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.DepartureLegTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_DepartureLegTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_DepartureLegTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DepartureLegType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DepartureLegType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.DepartureLegType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_DepartureLegType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_DepartureLegType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_DistanceIndicationPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DistanceIndicationPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.DistanceIndicationPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_DistanceIndicationPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_DistanceIndicationPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ElevatedCurveType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ElevatedCurveType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ElevatedCurveType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ElevatedCurveType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ElevatedCurveType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_ElevatedSurfaceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ElevatedSurfaceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ElevatedSurfaceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ElevatedSurfaceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ElevatedSurfaceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_EnRouteSegmentPointType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New EnRouteSegmentPointType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.EnRouteSegmentPointType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_EnRouteSegmentPointType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_EnRouteSegmentPointType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_EquipmentUnavailableAdjustmentColumnPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New EquipmentUnavailableAdjustmentColumnPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.EquipmentUnavailableAdjustmentColumnPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_EquipmentUnavailableAdjustmentColumnPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_EquipmentUnavailableAdjustmentColumnPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_EquipmentUnavailableAdjustmentColumnType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New EquipmentUnavailableAdjustmentColumnType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.EquipmentUnavailableAdjustmentColumnType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_EquipmentUnavailableAdjustmentColumnType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_EquipmentUnavailableAdjustmentColumnType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_EquipmentUnavailableAdjustmentPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New EquipmentUnavailableAdjustmentPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.EquipmentUnavailableAdjustmentPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_EquipmentUnavailableAdjustmentPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_EquipmentUnavailableAdjustmentPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_EquipmentUnavailableAdjustmentType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New EquipmentUnavailableAdjustmentType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.EquipmentUnavailableAdjustmentType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_EquipmentUnavailableAdjustmentType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_EquipmentUnavailableAdjustmentType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FASDataBlockPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New FASDataBlockPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.FASDataBlockPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_FASDataBlockPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_FASDataBlockPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FASDataBlockType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New FASDataBlockType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.FASDataBlockType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_FASDataBlockType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_FASDataBlockType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FinalLegTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New FinalLegTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.FinalLegTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_FinalLegTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_FinalLegTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FinalLegTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New FinalLegTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.FinalLegTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_FinalLegTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_FinalLegTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FinalLegType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New FinalLegType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.FinalLegType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_FinalLegType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_FinalLegType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FinalProfilePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New FinalProfilePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.FinalProfilePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_FinalProfilePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_FinalProfilePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FinalProfileType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New FinalProfileType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.FinalProfileType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_FinalProfileType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_FinalProfileType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_HoldingAssessmentTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New HoldingAssessmentTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.HoldingAssessmentTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_HoldingAssessmentTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_HoldingAssessmentTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_HoldingAssessmentTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New HoldingAssessmentTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.HoldingAssessmentTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_HoldingAssessmentTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_HoldingAssessmentTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_HoldingAssessmentType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New HoldingAssessmentType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.HoldingAssessmentType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_HoldingAssessmentType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_HoldingAssessmentType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_HoldingPatternDistancePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New HoldingPatternDistancePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.HoldingPatternDistancePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_HoldingPatternDistancePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_HoldingPatternDistancePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_HoldingPatternDistanceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New HoldingPatternDistanceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.HoldingPatternDistanceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_HoldingPatternDistanceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_HoldingPatternDistanceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_HoldingPatternDurationPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New HoldingPatternDurationPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.HoldingPatternDurationPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_HoldingPatternDurationPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_HoldingPatternDurationPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_HoldingPatternDurationType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New HoldingPatternDurationType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.HoldingPatternDurationType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_HoldingPatternDurationType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_HoldingPatternDurationType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_HoldingPatternPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New HoldingPatternPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.HoldingPatternPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_HoldingPatternPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_HoldingPatternPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_HoldingPatternTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New HoldingPatternTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.HoldingPatternTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_HoldingPatternTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_HoldingPatternTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_HoldingPatternTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New HoldingPatternTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.HoldingPatternTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_HoldingPatternTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_HoldingPatternTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_HoldingPatternType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New HoldingPatternType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.HoldingPatternType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_HoldingPatternType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_HoldingPatternType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_HoldingUsePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New HoldingUsePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.HoldingUsePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_HoldingUsePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_HoldingUsePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_HoldingUseType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New HoldingUseType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.HoldingUseType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_HoldingUseType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_HoldingUseType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_InitialLegTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New InitialLegTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.InitialLegTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_InitialLegTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_InitialLegTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_InitialLegTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New InitialLegTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.InitialLegTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_InitialLegTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_InitialLegTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_InitialLegType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New InitialLegType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.InitialLegType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_InitialLegType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_InitialLegType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_InstrumentApproachProcedurePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New InstrumentApproachProcedurePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.InstrumentApproachProcedurePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_InstrumentApproachProcedurePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_InstrumentApproachProcedurePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_InstrumentApproachProcedureTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New InstrumentApproachProcedureTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.InstrumentApproachProcedureTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_InstrumentApproachProcedureTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_InstrumentApproachProcedureTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_InstrumentApproachProcedureTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New InstrumentApproachProcedureTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.InstrumentApproachProcedureTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_InstrumentApproachProcedureTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_InstrumentApproachProcedureTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_InstrumentApproachProcedureType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New InstrumentApproachProcedureType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.InstrumentApproachProcedureType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_InstrumentApproachProcedureType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_InstrumentApproachProcedureType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_IntermediateLegTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New IntermediateLegTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.IntermediateLegTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_IntermediateLegTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_IntermediateLegTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_IntermediateLegTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New IntermediateLegTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.IntermediateLegTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_IntermediateLegTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_IntermediateLegTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_IntermediateLegType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New IntermediateLegType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.IntermediateLegType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_IntermediateLegType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_IntermediateLegType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_LandingTakeoffAreaCollectionPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New LandingTakeoffAreaCollectionPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.LandingTakeoffAreaCollectionPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_LandingTakeoffAreaCollectionPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_LandingTakeoffAreaCollectionPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_LandingTakeoffAreaCollectionType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New LandingTakeoffAreaCollectionType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.LandingTakeoffAreaCollectionType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_LandingTakeoffAreaCollectionType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_LandingTakeoffAreaCollectionType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_MinimaPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New MinimaPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.MinimaPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_MinimaPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_MinimaPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_MinimaType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New MinimaType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.MinimaType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_MinimaType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_MinimaType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_MissedApproachGroupPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New MissedApproachGroupPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.MissedApproachGroupPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_MissedApproachGroupPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_MissedApproachGroupPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_MissedApproachGroupType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New MissedApproachGroupType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.MissedApproachGroupType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_MissedApproachGroupType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_MissedApproachGroupType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_MissedApproachLegTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New MissedApproachLegTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.MissedApproachLegTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_MissedApproachLegTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_MissedApproachLegTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_MissedApproachLegTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New MissedApproachLegTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.MissedApproachLegTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_MissedApproachLegTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_MissedApproachLegTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_MissedApproachLegType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New MissedApproachLegType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.MissedApproachLegType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_MissedApproachLegType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_MissedApproachLegType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_NavigationAreaRestrictionTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New NavigationAreaRestrictionTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.NavigationAreaRestrictionTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_NavigationAreaRestrictionTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_NavigationAreaRestrictionTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NavigationAreaRestrictionTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New NavigationAreaRestrictionTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.NavigationAreaRestrictionTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_NavigationAreaRestrictionTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_NavigationAreaRestrictionTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NavigationAreaRestrictionType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New NavigationAreaRestrictionType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.NavigationAreaRestrictionType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_NavigationAreaRestrictionType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_NavigationAreaRestrictionType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NavigationAreaSectorPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New NavigationAreaSectorPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.NavigationAreaSectorPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_NavigationAreaSectorPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_NavigationAreaSectorPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NavigationAreaSectorType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New NavigationAreaSectorType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.NavigationAreaSectorType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_NavigationAreaSectorType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_NavigationAreaSectorType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NavigationAreaTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New NavigationAreaTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.NavigationAreaTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_NavigationAreaTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_NavigationAreaTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NavigationAreaTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New NavigationAreaTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.NavigationAreaTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_NavigationAreaTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_NavigationAreaTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_NavigationAreaType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New NavigationAreaType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.NavigationAreaType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_NavigationAreaType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_NavigationAreaType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_ObstacleAssessmentAreaPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ObstacleAssessmentAreaPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ObstacleAssessmentAreaPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ObstacleAssessmentAreaPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ObstacleAssessmentAreaPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ObstacleAssessmentAreaType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ObstacleAssessmentAreaType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ObstacleAssessmentAreaType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ObstacleAssessmentAreaType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ObstacleAssessmentAreaType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ObstaclePlacementPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ObstaclePlacementPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ObstaclePlacementPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ObstaclePlacementPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ObstaclePlacementPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ObstaclePlacementType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ObstaclePlacementType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ObstaclePlacementType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ObstaclePlacementType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ObstaclePlacementType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ObstructionPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ObstructionPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ObstructionPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ObstructionPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ObstructionPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ObstructionType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ObstructionType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ObstructionType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ObstructionType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ObstructionType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_PointReferencePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New PointReferencePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.PointReferencePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_PointReferencePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_PointReferencePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_PointReferenceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New PointReferenceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.PointReferenceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_PointReferenceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_PointReferenceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_ProcedureAvailabilityPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ProcedureAvailabilityPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ProcedureAvailabilityPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ProcedureAvailabilityPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ProcedureAvailabilityPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ProcedureAvailabilityType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ProcedureAvailabilityType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ProcedureAvailabilityType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ProcedureAvailabilityType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ProcedureAvailabilityType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ProcedureDMETimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ProcedureDMETimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ProcedureDMETimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ProcedureDMETimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ProcedureDMETimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ProcedureDMETimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ProcedureDMETimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ProcedureDMETimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ProcedureDMETimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ProcedureDMETimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ProcedureDMEType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ProcedureDMEType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ProcedureDMEType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ProcedureDMEType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ProcedureDMEType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ProcedurePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ProcedurePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ProcedurePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ProcedurePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ProcedurePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ProcedureTransitionLegPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ProcedureTransitionLegPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ProcedureTransitionLegPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ProcedureTransitionLegPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ProcedureTransitionLegPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ProcedureTransitionLegType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ProcedureTransitionLegType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ProcedureTransitionLegType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ProcedureTransitionLegType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ProcedureTransitionLegType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ProcedureTransitionPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ProcedureTransitionPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ProcedureTransitionPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ProcedureTransitionPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ProcedureTransitionPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ProcedureTransitionType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ProcedureTransitionType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ProcedureTransitionType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ProcedureTransitionType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ProcedureTransitionType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RadarSystemPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RadarSystemPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.RadarSystemPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_RadarSystemPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_RadarSystemPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RadioFrequencyAreaPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RadioFrequencyAreaPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.RadioFrequencyAreaPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_RadioFrequencyAreaPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_RadioFrequencyAreaPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_SafeAltitudeAreaPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SafeAltitudeAreaPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SafeAltitudeAreaPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SafeAltitudeAreaPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SafeAltitudeAreaPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SafeAltitudeAreaSectorPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SafeAltitudeAreaSectorPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SafeAltitudeAreaSectorPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SafeAltitudeAreaSectorPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SafeAltitudeAreaSectorPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SafeAltitudeAreaSectorType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SafeAltitudeAreaSectorType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SafeAltitudeAreaSectorType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SafeAltitudeAreaSectorType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SafeAltitudeAreaSectorType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SafeAltitudeAreaTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SafeAltitudeAreaTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SafeAltitudeAreaTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SafeAltitudeAreaTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SafeAltitudeAreaTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SafeAltitudeAreaTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SafeAltitudeAreaTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SafeAltitudeAreaTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SafeAltitudeAreaTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SafeAltitudeAreaTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SafeAltitudeAreaType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SafeAltitudeAreaType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SafeAltitudeAreaType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SafeAltitudeAreaType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SafeAltitudeAreaType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SectorDesignPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SectorDesignPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SectorDesignPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SectorDesignPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SectorDesignPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SectorDesignType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SectorDesignType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SectorDesignType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SectorDesignType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SectorDesignType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SegmentLegPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SegmentLegPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SegmentLegPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SegmentLegPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SegmentLegPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SegmentPointPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SegmentPointPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SegmentPointPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SegmentPointPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SegmentPointPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_StandardInstrumentArrivalPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New StandardInstrumentArrivalPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.StandardInstrumentArrivalPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_StandardInstrumentArrivalPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_StandardInstrumentArrivalPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_StandardInstrumentArrivalTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New StandardInstrumentArrivalTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.StandardInstrumentArrivalTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_StandardInstrumentArrivalTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_StandardInstrumentArrivalTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_StandardInstrumentArrivalTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New StandardInstrumentArrivalTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.StandardInstrumentArrivalTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_StandardInstrumentArrivalTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_StandardInstrumentArrivalTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_StandardInstrumentArrivalType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New StandardInstrumentArrivalType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.StandardInstrumentArrivalType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_StandardInstrumentArrivalType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_StandardInstrumentArrivalType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_StandardInstrumentDeparturePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New StandardInstrumentDeparturePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.StandardInstrumentDeparturePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_StandardInstrumentDeparturePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_StandardInstrumentDeparturePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_StandardInstrumentDepartureTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New StandardInstrumentDepartureTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.StandardInstrumentDepartureTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_StandardInstrumentDepartureTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_StandardInstrumentDepartureTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_StandardInstrumentDepartureTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New StandardInstrumentDepartureTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.StandardInstrumentDepartureTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_StandardInstrumentDepartureTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_StandardInstrumentDepartureTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_StandardInstrumentDepartureType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New StandardInstrumentDepartureType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.StandardInstrumentDepartureType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_StandardInstrumentDepartureType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_StandardInstrumentDepartureType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SurfacePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SurfacePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SurfacePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SurfacePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SurfacePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SurfaceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SurfaceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SurfaceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SurfaceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SurfaceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TerminalArrivalAreaSectorPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TerminalArrivalAreaSectorPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.TerminalArrivalAreaSectorPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_TerminalArrivalAreaSectorPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_TerminalArrivalAreaSectorPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TerminalArrivalAreaSectorType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TerminalArrivalAreaSectorType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.TerminalArrivalAreaSectorType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_TerminalArrivalAreaSectorType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_TerminalArrivalAreaSectorType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TerminalArrivalAreaTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TerminalArrivalAreaTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.TerminalArrivalAreaTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_TerminalArrivalAreaTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_TerminalArrivalAreaTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TerminalArrivalAreaTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TerminalArrivalAreaTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.TerminalArrivalAreaTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_TerminalArrivalAreaTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_TerminalArrivalAreaTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TerminalArrivalAreaType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TerminalArrivalAreaType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.TerminalArrivalAreaType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_TerminalArrivalAreaType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_TerminalArrivalAreaType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TerminalSegmentPointPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TerminalSegmentPointPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.TerminalSegmentPointPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_TerminalSegmentPointPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_TerminalSegmentPointPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_TerminalSegmentPointType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TerminalSegmentPointType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.TerminalSegmentPointType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_TerminalSegmentPointType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_TerminalSegmentPointType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_UnplannedHoldingPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New UnplannedHoldingPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.UnplannedHoldingPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_UnplannedHoldingPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_UnplannedHoldingPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_UnplannedHoldingTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New UnplannedHoldingTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.UnplannedHoldingTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_UnplannedHoldingTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_UnplannedHoldingTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_UnplannedHoldingTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New UnplannedHoldingTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.UnplannedHoldingTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_UnplannedHoldingTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_UnplannedHoldingTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_UnplannedHoldingType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New UnplannedHoldingType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.UnplannedHoldingType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_UnplannedHoldingType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_UnplannedHoldingType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_VerticalStructurePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New VerticalStructurePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.VerticalStructurePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_VerticalStructurePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_VerticalStructurePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeATCReportingType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeATCReportingType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeATCReportingType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeATCReportingType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeATCReportingType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeAerialRefuellingPointType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeAerialRefuellingPointType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeAerialRefuellingPointType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeAerialRefuellingPointType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeAerialRefuellingPointType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeAircraftCategoryType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeAircraftCategoryType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeAircraftCategoryType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeAircraftCategoryType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeAircraftCategoryType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeAircraftEngineNumberType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeAircraftEngineNumberType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeAircraftEngineNumberType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeAircraftEngineNumberType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeAircraftEngineNumberType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeAircraftEngineType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeAircraftEngineType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeAircraftEngineType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeAircraftEngineType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeAircraftEngineType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeAircraftICAOType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeAircraftICAOType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeAircraftICAOType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeAircraftICAOType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeAircraftICAOType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeAircraftType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeAircraftType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeAircraftType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeAircraftType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeAircraftType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeAircraftWingspanClassType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeAircraftWingspanClassType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeAircraftWingspanClassType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeAircraftWingspanClassType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeAircraftWingspanClassType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeAltitudeAdjustmentType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeAltitudeAdjustmentType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeAltitudeAdjustmentType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeAltitudeAdjustmentType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeAltitudeAdjustmentType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_CodeApproachEquipmentAdditionalType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeApproachEquipmentAdditionalType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeApproachEquipmentAdditionalType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeApproachEquipmentAdditionalType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeApproachEquipmentAdditionalType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeApproachGuidanceType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeApproachGuidanceType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeApproachGuidanceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeApproachGuidanceType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeApproachGuidanceType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeApproachPrefixType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeApproachPrefixType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeApproachPrefixType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeApproachPrefixType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeApproachPrefixType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeApproachType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeApproachType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeApproachType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeApproachType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeApproachType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeApprovalType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeApprovalType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeApprovalType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeApprovalType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeApprovalType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeArcDirectionType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeArcDirectionType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeArcDirectionType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeArcDirectionType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeArcDirectionType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_CodeCommunicationModeType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeCommunicationModeType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeCommunicationModeType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeCommunicationModeType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeCommunicationModeType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeCourseType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeCourseType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeCourseType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeCourseType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeCourseType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_CodeDesignStandardType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeDesignStandardType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeDesignStandardType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeDesignStandardType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeDesignStandardType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_CodeDirectionTurnType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeDirectionTurnType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeDirectionTurnType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeDirectionTurnType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeDirectionTurnType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeEquipmentAntiCollisionType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeEquipmentAntiCollisionType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeEquipmentAntiCollisionType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeEquipmentAntiCollisionType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeEquipmentAntiCollisionType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeEquipmentUnavailableType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeEquipmentUnavailableType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeEquipmentUnavailableType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeEquipmentUnavailableType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeEquipmentUnavailableType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeFinalGuidanceType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeFinalGuidanceType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeFinalGuidanceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeFinalGuidanceType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeFinalGuidanceType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeFreeFlightType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeFreeFlightType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeFreeFlightType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeFreeFlightType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeFreeFlightType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeHeightReferenceType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeHeightReferenceType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeHeightReferenceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeHeightReferenceType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeHeightReferenceType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeHoldingUsageType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeHoldingUsageType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeHoldingUsageType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeHoldingUsageType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeHoldingUsageType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeHoldingUseType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeHoldingUseType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeHoldingUseType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeHoldingUseType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeHoldingUseType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeICAOCountryType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeICAOCountryType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeICAOCountryType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeICAOCountryType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeICAOCountryType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeMilitaryRoutePointType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeMilitaryRoutePointType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeMilitaryRoutePointType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeMilitaryRoutePointType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeMilitaryRoutePointType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeMinimaFinalApproachPathType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeMinimaFinalApproachPathType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeMinimaFinalApproachPathType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeMinimaFinalApproachPathType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeMinimaFinalApproachPathType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeMinimumAltitudeType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeMinimumAltitudeType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeMinimumAltitudeType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeMinimumAltitudeType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeMinimumAltitudeType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeMinimumHeightType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeMinimumHeightType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeMinimumHeightType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeMinimumHeightType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeMinimumHeightType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeMissedApproachType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeMissedApproachType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeMissedApproachType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeMissedApproachType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeMissedApproachType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeNavigationAreaRestrictionType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeNavigationAreaRestrictionType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeNavigationAreaRestrictionType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeNavigationAreaRestrictionType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeNavigationAreaRestrictionType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeNavigationAreaType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeNavigationAreaType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeNavigationAreaType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeNavigationAreaType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeNavigationAreaType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeNavigationEquipmentType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeNavigationEquipmentType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeNavigationEquipmentType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeNavigationEquipmentType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeNavigationEquipmentType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeNavigationSpecificationType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeNavigationSpecificationType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeNavigationSpecificationType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeNavigationSpecificationType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeNavigationSpecificationType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_CodeObstacleAssessmentSurfaceType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeObstacleAssessmentSurfaceType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeObstacleAssessmentSurfaceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeObstacleAssessmentSurfaceType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeObstacleAssessmentSurfaceType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeObstructionIdSurfaceZoneType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeObstructionIdSurfaceZoneType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeObstructionIdSurfaceZoneType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeObstructionIdSurfaceZoneType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeObstructionIdSurfaceZoneType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeProcedureAvailabilityType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeProcedureAvailabilityType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeProcedureAvailabilityType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeProcedureAvailabilityType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeProcedureAvailabilityType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeProcedureCodingStandardType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeProcedureCodingStandardType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeProcedureCodingStandardType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeProcedureCodingStandardType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeProcedureCodingStandardType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeProcedureDistanceType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeProcedureDistanceType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeProcedureDistanceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeProcedureDistanceType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeProcedureDistanceType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeProcedureFixRoleType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeProcedureFixRoleType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeProcedureFixRoleType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeProcedureFixRoleType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeProcedureFixRoleType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeProcedurePhaseType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeProcedurePhaseType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeProcedurePhaseType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeProcedurePhaseType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeProcedurePhaseType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeRNPType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeRNPType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeRNPType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeRNPType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeRNPType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeRVSMPointRoleType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeRVSMPointRoleType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeRVSMPointRoleType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeRVSMPointRoleType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeRVSMPointRoleType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeRVSMType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeRVSMType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeRVSMType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeRVSMType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeRVSMType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeReferencePathIdentifierType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeReferencePathIdentifierType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeReferencePathIdentifierType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeReferencePathIdentifierType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeReferencePathIdentifierType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeReferenceRoleType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeReferenceRoleType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeReferenceRoleType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeReferenceRoleType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeReferenceRoleType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeRelativePositionType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeRelativePositionType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeRelativePositionType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeRelativePositionType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeRelativePositionType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeRouteIndicatorType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeRouteIndicatorType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeRouteIndicatorType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeRouteIndicatorType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeRouteIndicatorType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeSafeAltitudeType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeSafeAltitudeType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeSafeAltitudeType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeSafeAltitudeType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeSafeAltitudeType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeSegmentPathType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeSegmentPathType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeSegmentPathType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeSegmentPathType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeSegmentPathType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeSegmentTerminationType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeSegmentTerminationType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeSegmentTerminationType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeSegmentTerminationType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeSegmentTerminationType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeSideType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeSideType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeSideType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeSideType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeSideType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeSpeedReferenceType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeSpeedReferenceType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeSpeedReferenceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeSpeedReferenceType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeSpeedReferenceType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeTAAType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeTAAType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeTAAType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeTAAType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeTAAType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_CodeTrajectoryType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeTrajectoryType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeTrajectoryType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeTrajectoryType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeTrajectoryType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_CodeUpperAlphaType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeUpperAlphaType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeUpperAlphaType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeUpperAlphaType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeUpperAlphaType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeValueInterpretationType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeValueInterpretationType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeValueInterpretationType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeValueInterpretationType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeValueInterpretationType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_CodeWakeTurbulenceType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeWakeTurbulenceType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeWakeTurbulenceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeWakeTurbulenceType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeWakeTurbulenceType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_NoNumberType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New NoNumberType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.NoNumberType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_NoNumberType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_NoNumberType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_TextInstructionType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TextInstructionType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.TextInstructionType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_TextInstructionType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_TextInstructionType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_TextSIDSTARDesignatorType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New TextSIDSTARDesignatorType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.TextSIDSTARDesignatorType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_TextSIDSTARDesignatorType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_TextSIDSTARDesignatorType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_ValAlarmLimitType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ValAlarmLimitType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.ValAlarmLimitType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_ValAlarmLimitType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_ValAlarmLimitType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_ValChannelNumberType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ValChannelNumberType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.ValChannelNumberType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_ValChannelNumberType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_ValChannelNumberType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_ValHexType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ValHexType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.ValHexType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_ValHexType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_ValHexType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ValSlopeType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ValSlopeType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.ValSlopeType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_ValSlopeType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_ValSlopeType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ValSpeedType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ValSpeedType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.ValSpeedType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_ValSpeedType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_ValSpeedType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ValTemperatureType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ValTemperatureType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.ValTemperatureType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_ValTemperatureType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_ValTemperatureType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ValWeightType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ValWeightType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.ValWeightType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_ValWeightType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_ValWeightType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_AbstractCurveSegmentType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractCurveSegmentType",
	}).Stage(probe.formStage)
	inst := new(gml.AbstractCurveSegmentType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_AbstractCurveSegmentType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_AbstractCurveSegmentType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractCurveType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractCurveType",
	}).Stage(probe.formStage)
	inst := new(gml.AbstractCurveType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_AbstractCurveType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_AbstractCurveType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_AbstractRingPropertyType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractRingPropertyType",
	}).Stage(probe.formStage)
	inst := new(gml.AbstractRingPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_AbstractRingPropertyType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_AbstractRingPropertyType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractRingType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractRingType",
	}).Stage(probe.formStage)
	inst := new(gml.AbstractRingType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_AbstractRingType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_AbstractRingType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractSurfacePatchType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractSurfacePatchType",
	}).Stage(probe.formStage)
	inst := new(gml.AbstractSurfacePatchType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_AbstractSurfacePatchType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_AbstractSurfacePatchType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AbstractSurfaceType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractSurfaceType",
	}).Stage(probe.formStage)
	inst := new(gml.AbstractSurfaceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_AbstractSurfaceType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_AbstractSurfaceType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_AngleType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AngleType",
	}).Stage(probe.formStage)
	inst := new(gml.AngleType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_AngleType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_AngleType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ArcByCenterPointType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ArcByCenterPointType",
	}).Stage(probe.formStage)
	inst := new(gml.ArcByCenterPointType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_ArcByCenterPointType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_ArcByCenterPointType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ArcStringType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ArcStringType",
	}).Stage(probe.formStage)
	inst := new(gml.ArcStringType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_ArcStringType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_ArcStringType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ArcType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ArcType",
	}).Stage(probe.formStage)
	inst := new(gml.ArcType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_ArcType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_ArcType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_CircleByCenterPointType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CircleByCenterPointType",
	}).Stage(probe.formStage)
	inst := new(gml.CircleByCenterPointType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_CircleByCenterPointType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_CircleByCenterPointType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CircleType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CircleType",
	}).Stage(probe.formStage)
	inst := new(gml.CircleType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_CircleType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_CircleType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_CompositeCurveType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CompositeCurveType",
	}).Stage(probe.formStage)
	inst := new(gml.CompositeCurveType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_CompositeCurveType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_CompositeCurveType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CurvePropertyType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CurvePropertyType",
	}).Stage(probe.formStage)
	inst := new(gml.CurvePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_CurvePropertyType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_CurvePropertyType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CurveSegmentArrayPropertyType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CurveSegmentArrayPropertyType",
	}).Stage(probe.formStage)
	inst := new(gml.CurveSegmentArrayPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_CurveSegmentArrayPropertyType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_CurveSegmentArrayPropertyType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CurveType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CurveType",
	}).Stage(probe.formStage)
	inst := new(gml.CurveType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_CurveType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_CurveType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DirectPositionListType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DirectPositionListType",
	}).Stage(probe.formStage)
	inst := new(gml.DirectPositionListType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_DirectPositionListType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_DirectPositionListType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_GeodesicStringType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New GeodesicStringType",
	}).Stage(probe.formStage)
	inst := new(gml.GeodesicStringType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_GeodesicStringType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_GeodesicStringType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_GeodesicType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New GeodesicType",
	}).Stage(probe.formStage)
	inst := new(gml.GeodesicType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_GeodesicType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_GeodesicType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_LengthType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New LengthType",
	}).Stage(probe.formStage)
	inst := new(gml.LengthType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_LengthType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_LengthType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_LineStringSegmentType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New LineStringSegmentType",
	}).Stage(probe.formStage)
	inst := new(gml.LineStringSegmentType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_LineStringSegmentType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_LineStringSegmentType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_MeasureType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New MeasureType",
	}).Stage(probe.formStage)
	inst := new(gml.MeasureType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_MeasureType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_MeasureType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_OrientableCurveType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New OrientableCurveType",
	}).Stage(probe.formStage)
	inst := new(gml.OrientableCurveType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_OrientableCurveType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_OrientableCurveType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_PointPropertyType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New PointPropertyType",
	}).Stage(probe.formStage)
	inst := new(gml.PointPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_PointPropertyType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_PointPropertyType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_PolygonPatchType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New PolygonPatchType",
	}).Stage(probe.formStage)
	inst := new(gml.PolygonPatchType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_PolygonPatchType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_PolygonPatchType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RingType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RingType",
	}).Stage(probe.formStage)
	inst := new(gml.RingType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_RingType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_RingType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SurfacePatchArrayPropertyType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SurfacePatchArrayPropertyType",
	}).Stage(probe.formStage)
	inst := new(gml.SurfacePatchArrayPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_SurfacePatchArrayPropertyType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_SurfacePatchArrayPropertyType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SurfaceType_GmlStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SurfaceType",
	}).Stage(probe.formStage)
	inst := new(gml.SurfaceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.GmlStage.Lock()
			defer probe.stageOfInterest.GmlStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.GmlStage)
			saveStageSet_SurfaceType_GmlStage(inst, probe, formGroup)
			probe.stageOfInterest.GmlStage.Commit()
			updateStageSetTable_SurfaceType_GmlStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

