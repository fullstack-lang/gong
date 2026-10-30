// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	gongprobe "github.com/fullstack-lang/gong/pkg/runtime/probe"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/xlink"
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
		case "Route":
			StageSetFormDivSelectFieldToField(&inst.Route, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RouteType](), formDiv)
		case "RouteSegment":
			StageSetFormDivSelectFieldToField(&inst.RouteSegment, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RouteSegmentType](), formDiv)
		case "RouteDME":
			StageSetFormDivSelectFieldToField(&inst.RouteDME, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RouteDMEType](), formDiv)
		case "ChangeOverPoint":
			StageSetFormDivSelectFieldToField(&inst.ChangeOverPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ChangeOverPointType](), formDiv)
		case "FlightRestriction":
			StageSetFormDivSelectFieldToField(&inst.FlightRestriction, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FlightRestrictionType](), formDiv)
		case "AirspaceBorderCrossing":
			StageSetFormDivSelectFieldToField(&inst.AirspaceBorderCrossing, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceBorderCrossingType](), formDiv)
		case "RulesProcedures":
			StageSetFormDivSelectFieldToField(&inst.RulesProcedures, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RulesProceduresType](), formDiv)
		case "AerialRefuelling":
			StageSetFormDivSelectFieldToField(&inst.AerialRefuelling, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AerialRefuellingType](), formDiv)
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

func saveStageSet_AbstractDirectFlightType_AixmfeaturesStage(
	inst *aixmfeatures.AbstractDirectFlightType,
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

func saveStageSet_AerialRefuellingAnchorPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AerialRefuellingAnchorPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "AerialRefuellingAnchor":
			StageSetFormDivSelectFieldToField(&inst.AerialRefuellingAnchor, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AerialRefuellingAnchorType](), formDiv)
		}
	}
}

func saveStageSet_AerialRefuellingAnchorType_AixmfeaturesStage(
	inst *aixmfeatures.AerialRefuellingAnchorType,
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
		case "OutboundCourse":
			StageSetFormDivSelectFieldToField(&inst.OutboundCourse, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "OutboundCourseType":
			StageSetFormDivSelectFieldToField(&inst.OutboundCourseType, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCourseType](), formDiv)
		case "InboundCourse":
			StageSetFormDivSelectFieldToField(&inst.InboundCourse, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "TurnDirection":
			StageSetFormDivSelectFieldToField(&inst.TurnDirection, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDirectionTurnType](), formDiv)
		case "SpeedLimit":
			StageSetFormDivSelectFieldToField(&inst.SpeedLimit, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValSpeedType](), formDiv)
		case "LegSeparation":
			StageSetFormDivSelectFieldToField(&inst.LegSeparation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "LegLength":
			StageSetFormDivSelectFieldToField(&inst.LegLength, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "RefuellingBaseLevel":
			StageSetFormDivSelectFieldToField(&inst.RefuellingBaseLevel, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "RefuellingBaseLevelReference":
			StageSetFormDivSelectFieldToField(&inst.RefuellingBaseLevelReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "Extent":
			StageSetFormDivSelectFieldToField(&inst.Extent, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SurfacePropertyType](), formDiv)
		}
	}
}

func saveStageSet_AerialRefuellingPointPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AerialRefuellingPointPropertyType,
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

func saveStageSet_AerialRefuellingPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AerialRefuellingPropertyType,
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

func saveStageSet_AerialRefuellingTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AerialRefuellingTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "AerialRefuellingTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.AerialRefuellingTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AerialRefuellingTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_AerialRefuellingTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.AerialRefuellingTimeSliceType,
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
		case "DesignatorPrefix":
			StageSetFormDivSelectFieldToField(&inst.DesignatorPrefix, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAerialRefuellingPrefixType](), formDiv)
		case "DesignatorNumber":
			StageSetFormDivSelectFieldToField(&inst.DesignatorNumber, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoNumberType](), formDiv)
		case "DesignatorSuffix":
			StageSetFormDivSelectFieldToField(&inst.DesignatorSuffix, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextDesignatorType](), formDiv)
		case "DesignatorDirection":
			StageSetFormDivSelectFieldToField(&inst.DesignatorDirection, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCardinalDirectionType](), formDiv)
		case "Type_XSD":
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAerialRefuellingType](), formDiv)
		case "RadarBeaconSetting":
			StageSetFormDivSelectFieldToField(&inst.RadarBeaconSetting, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoNumberType](), formDiv)
		case "XbandRadarSetting":
			StageSetFormDivSelectFieldToField(&inst.XbandRadarSetting, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoNumberType](), formDiv)
		case "TankerChannel":
			StageSetFormDivSelectFieldToField(&inst.TankerChannel, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTACANChannelType](), formDiv)
		case "ReceiverChannel":
			StageSetFormDivSelectFieldToField(&inst.ReceiverChannel, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTACANChannelType](), formDiv)
		case "HelicopterRoute":
			StageSetFormDivSelectFieldToField(&inst.HelicopterRoute, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "SpecialRefuelling":
			StageSetFormDivSelectFieldToField(&inst.SpecialRefuelling, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "BidirectionalUse":
			StageSetFormDivSelectFieldToField(&inst.BidirectionalUse, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "ReverseDirectionTurn":
			StageSetFormDivSelectFieldToField(&inst.ReverseDirectionTurn, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDirectionTurnType](), formDiv)
		case "OppositeTrack":
			StageSetFormDivSelectFieldToField(&inst.OppositeTrack, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AerialRefuellingPropertyType](), formDiv)
		}
	}
}

func saveStageSet_AerialRefuellingTrackPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AerialRefuellingTrackPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "AerialRefuellingTrack":
			StageSetFormDivSelectFieldToField(&inst.AerialRefuellingTrack, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AerialRefuellingTrackType](), formDiv)
		}
	}
}

func saveStageSet_AerialRefuellingTrackType_AixmfeaturesStage(
	inst *aixmfeatures.AerialRefuellingTrackType,
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
		case "Extent":
			StageSetFormDivSelectFieldToField(&inst.Extent, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CurvePropertyType](), formDiv)
		}
	}
}

func saveStageSet_AerialRefuellingType_AixmfeaturesStage(
	inst *aixmfeatures.AerialRefuellingType,
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

func saveStageSet_AirspaceBorderCrossingPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AirspaceBorderCrossingPropertyType,
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

func saveStageSet_AirspaceBorderCrossingTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AirspaceBorderCrossingTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "AirspaceBorderCrossingTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.AirspaceBorderCrossingTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceBorderCrossingTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_AirspaceBorderCrossingTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.AirspaceBorderCrossingTimeSliceType,
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
		case "ExitedAirspace":
			StageSetFormDivSelectFieldToField(&inst.ExitedAirspace, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspacePropertyType](), formDiv)
		case "EnteredAirspace":
			StageSetFormDivSelectFieldToField(&inst.EnteredAirspace, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspacePropertyType](), formDiv)
		}
	}
}

func saveStageSet_AirspaceBorderCrossingType_AixmfeaturesStage(
	inst *aixmfeatures.AirspaceBorderCrossingType,
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

func saveStageSet_AirspaceLayerPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AirspaceLayerPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "AirspaceLayer":
			StageSetFormDivSelectFieldToField(&inst.AirspaceLayer, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceLayerType](), formDiv)
		}
	}
}

func saveStageSet_AirspaceLayerType_AixmfeaturesStage(
	inst *aixmfeatures.AirspaceLayerType,
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
		case "DiscreteLevelSeries":
			StageSetFormDivSelectFieldToField(&inst.DiscreteLevelSeries, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.StandardLevelColumnPropertyType](), formDiv)
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

func saveStageSet_AuthorityForAerialRefuellingPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AuthorityForAerialRefuellingPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "AuthorityForAerialRefuelling":
			StageSetFormDivSelectFieldToField(&inst.AuthorityForAerialRefuelling, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AuthorityForAerialRefuellingType](), formDiv)
		}
	}
}

func saveStageSet_AuthorityForAerialRefuellingType_AixmfeaturesStage(
	inst *aixmfeatures.AuthorityForAerialRefuellingType,
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
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAuthorityType](), formDiv)
		case "TheOrganisationAuthority":
			StageSetFormDivSelectFieldToField(&inst.TheOrganisationAuthority, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.OrganisationAuthorityPropertyType](), formDiv)
		}
	}
}

func saveStageSet_ChangeOverPointTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.ChangeOverPointTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "ChangeOverPointTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.ChangeOverPointTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ChangeOverPointTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_ChangeOverPointTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.ChangeOverPointTimeSliceType,
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
		case "ApplicableRoutePortion":
			StageSetFormDivSelectFieldToField(&inst.ApplicableRoutePortion, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RoutePortionPropertyType](), formDiv)
		}
	}
}

func saveStageSet_ChangeOverPointType_AixmfeaturesStage(
	inst *aixmfeatures.ChangeOverPointType,
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

func saveStageSet_DirectFlightClassType_AixmfeaturesStage(
	inst *aixmfeatures.DirectFlightClassType,
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
		case "ExceedLength":
			StageSetFormDivSelectFieldToField(&inst.ExceedLength, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		}
	}
}

func saveStageSet_DirectFlightPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.DirectFlightPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "DirectFlightClass":
			StageSetFormDivSelectFieldToField(&inst.DirectFlightClass, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DirectFlightClassType](), formDiv)
		case "DirectFlightSegment":
			StageSetFormDivSelectFieldToField(&inst.DirectFlightSegment, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DirectFlightSegmentType](), formDiv)
		}
	}
}

func saveStageSet_DirectFlightSegmentPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.DirectFlightSegmentPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "DirectFlightSegment":
			StageSetFormDivSelectFieldToField(&inst.DirectFlightSegment, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DirectFlightSegmentType](), formDiv)
		}
	}
}

func saveStageSet_DirectFlightSegmentType_AixmfeaturesStage(
	inst *aixmfeatures.DirectFlightSegmentType,
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
		case "End_fixDesignatedPoint":
			StageSetFormDivSelectFieldToField(&inst.End_fixDesignatedPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), formDiv)
		case "End_navaidSystem":
			StageSetFormDivSelectFieldToField(&inst.End_navaidSystem, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), formDiv)
		case "End_position":
			StageSetFormDivSelectFieldToField(&inst.End_position, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), formDiv)
		case "End_runwayPoint":
			StageSetFormDivSelectFieldToField(&inst.End_runwayPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), formDiv)
		case "End_aimingPoint":
			StageSetFormDivSelectFieldToField(&inst.End_aimingPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), formDiv)
		case "End_airportReferencePoint":
			StageSetFormDivSelectFieldToField(&inst.End_airportReferencePoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), formDiv)
		case "Start_fixDesignatedPoint":
			StageSetFormDivSelectFieldToField(&inst.Start_fixDesignatedPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), formDiv)
		case "Start_navaidSystem":
			StageSetFormDivSelectFieldToField(&inst.Start_navaidSystem, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), formDiv)
		case "Start_position":
			StageSetFormDivSelectFieldToField(&inst.Start_position, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), formDiv)
		case "Start_runwayPoint":
			StageSetFormDivSelectFieldToField(&inst.Start_runwayPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), formDiv)
		case "Start_aimingPoint":
			StageSetFormDivSelectFieldToField(&inst.Start_aimingPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), formDiv)
		case "Start_airportReferencePoint":
			StageSetFormDivSelectFieldToField(&inst.Start_airportReferencePoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), formDiv)
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

func saveStageSet_EnRouteSegmentPointPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.EnRouteSegmentPointPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnRouteSegmentPoint":
			StageSetFormDivSelectFieldToField(&inst.EnRouteSegmentPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.EnRouteSegmentPointType](), formDiv)
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

func saveStageSet_FlightCharacteristicPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.FlightCharacteristicPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "FlightCharacteristic":
			StageSetFormDivSelectFieldToField(&inst.FlightCharacteristic, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FlightCharacteristicType](), formDiv)
		}
	}
}

func saveStageSet_FlightCharacteristicType_AixmfeaturesStage(
	inst *aixmfeatures.FlightCharacteristicType,
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
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlightType](), formDiv)
		case "Rule":
			StageSetFormDivSelectFieldToField(&inst.Rule, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlightRuleType](), formDiv)
		case "Status":
			StageSetFormDivSelectFieldToField(&inst.Status, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlightStatusType](), formDiv)
		case "Military":
			StageSetFormDivSelectFieldToField(&inst.Military, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeMilitaryStatusType](), formDiv)
		case "Origin":
			StageSetFormDivSelectFieldToField(&inst.Origin, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlightOriginType](), formDiv)
		case "Purpose":
			StageSetFormDivSelectFieldToField(&inst.Purpose, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlightPurposeType](), formDiv)
		}
	}
}

func saveStageSet_FlightConditionCircumstancePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.FlightConditionCircumstancePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "FlightConditionCircumstance":
			StageSetFormDivSelectFieldToField(&inst.FlightConditionCircumstance, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FlightConditionCircumstanceType](), formDiv)
		}
	}
}

func saveStageSet_FlightConditionCircumstanceType_AixmfeaturesStage(
	inst *aixmfeatures.FlightConditionCircumstanceType,
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
		case "ReferenceLocation":
			StageSetFormDivSelectFieldToField(&inst.ReferenceLocation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "RelationWithLocation":
			StageSetFormDivSelectFieldToField(&inst.RelationWithLocation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeLocationQualifierType](), formDiv)
		}
	}
}

func saveStageSet_FlightConditionCombinationPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.FlightConditionCombinationPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "FlightConditionCombination":
			StageSetFormDivSelectFieldToField(&inst.FlightConditionCombination, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FlightConditionCombinationType](), formDiv)
		}
	}
}

func saveStageSet_FlightConditionCombinationType_AixmfeaturesStage(
	inst *aixmfeatures.FlightConditionCombinationType,
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
		case "LogicalOperator":
			StageSetFormDivSelectFieldToField(&inst.LogicalOperator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlowConditionOperationType](), formDiv)
		}
	}
}

func saveStageSet_FlightConditionElementPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.FlightConditionElementPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "FlightConditionElement":
			StageSetFormDivSelectFieldToField(&inst.FlightConditionElement, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FlightConditionElementType](), formDiv)
		}
	}
}

func saveStageSet_FlightConditionElementType_AixmfeaturesStage(
	inst *aixmfeatures.FlightConditionElementType,
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
		case "Index":
			StageSetFormDivSelectFieldToField(&inst.Index, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoSequenceType](), formDiv)
		case "FlightCondition_weather":
			StageSetFormDivSelectFieldToField(&inst.FlightCondition_weather, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.MeteorologyPropertyType](), formDiv)
		case "FlightCondition_aircraft":
			StageSetFormDivSelectFieldToField(&inst.FlightCondition_aircraft, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AircraftCharacteristicPropertyType](), formDiv)
		case "FlightCondition_flight":
			StageSetFormDivSelectFieldToField(&inst.FlightCondition_flight, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FlightCharacteristicPropertyType](), formDiv)
		case "FlightCondition_directFlightCondition":
			StageSetFormDivSelectFieldToField(&inst.FlightCondition_directFlightCondition, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DirectFlightPropertyType](), formDiv)
		case "FlightCondition_borderCrossingCondition":
			StageSetFormDivSelectFieldToField(&inst.FlightCondition_borderCrossingCondition, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceBorderCrossingPropertyType](), formDiv)
		case "FlightCondition_operand":
			StageSetFormDivSelectFieldToField(&inst.FlightCondition_operand, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FlightConditionCombinationPropertyType](), formDiv)
		case "FlightCondition_routePortionCondition":
			StageSetFormDivSelectFieldToField(&inst.FlightCondition_routePortionCondition, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RoutePortionPropertyType](), formDiv)
		case "FlightCondition_standardInstrumentDepartureCondition":
			StageSetFormDivSelectFieldToField(&inst.FlightCondition_standardInstrumentDepartureCondition, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.StandardInstrumentDeparturePropertyType](), formDiv)
		case "FlightCondition_standardInstrumentArrivalCondition":
			StageSetFormDivSelectFieldToField(&inst.FlightCondition_standardInstrumentArrivalCondition, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.StandardInstrumentArrivalPropertyType](), formDiv)
		case "FlightCondition_organisationCondition":
			StageSetFormDivSelectFieldToField(&inst.FlightCondition_organisationCondition, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.OrganisationAuthorityPropertyType](), formDiv)
		case "SignificantPointCondition_fixDesignatedPoint":
			StageSetFormDivSelectFieldToField(&inst.SignificantPointCondition_fixDesignatedPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), formDiv)
		case "SignificantPointCondition_navaidSystem":
			StageSetFormDivSelectFieldToField(&inst.SignificantPointCondition_navaidSystem, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), formDiv)
		case "SignificantPointCondition_position":
			StageSetFormDivSelectFieldToField(&inst.SignificantPointCondition_position, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), formDiv)
		case "SignificantPointCondition_runwayPoint":
			StageSetFormDivSelectFieldToField(&inst.SignificantPointCondition_runwayPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), formDiv)
		case "SignificantPointCondition_aimingPoint":
			StageSetFormDivSelectFieldToField(&inst.SignificantPointCondition_aimingPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), formDiv)
		case "SignificantPointCondition_airportReferencePoint":
			StageSetFormDivSelectFieldToField(&inst.SignificantPointCondition_airportReferencePoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), formDiv)
		case "FlightCondition_airspaceCondition":
			StageSetFormDivSelectFieldToField(&inst.FlightCondition_airspaceCondition, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspacePropertyType](), formDiv)
		case "FlightCondition_airportHeliportCondition":
			StageSetFormDivSelectFieldToField(&inst.FlightCondition_airportHeliportCondition, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), formDiv)
		case "FlightCondition_aerialRefuellingCondition":
			StageSetFormDivSelectFieldToField(&inst.FlightCondition_aerialRefuellingCondition, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AerialRefuellingPropertyType](), formDiv)
		case "OperationalCondition":
			StageSetFormDivSelectFieldToField(&inst.OperationalCondition, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FlightConditionCircumstancePropertyType](), formDiv)
		}
	}
}

func saveStageSet_FlightRestrictionLevelPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.FlightRestrictionLevelPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "FlightRestrictionLevel":
			StageSetFormDivSelectFieldToField(&inst.FlightRestrictionLevel, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FlightRestrictionLevelType](), formDiv)
		}
	}
}

func saveStageSet_FlightRestrictionLevelType_AixmfeaturesStage(
	inst *aixmfeatures.FlightRestrictionLevelType,
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
		case "UpperLevel":
			StageSetFormDivSelectFieldToField(&inst.UpperLevel, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "UpperLevelReference":
			StageSetFormDivSelectFieldToField(&inst.UpperLevelReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "LowerLevel":
			StageSetFormDivSelectFieldToField(&inst.LowerLevel, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "LowerLevelReference":
			StageSetFormDivSelectFieldToField(&inst.LowerLevelReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		}
	}
}

func saveStageSet_FlightRestrictionRoutePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.FlightRestrictionRoutePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "FlightRestrictionRoute":
			StageSetFormDivSelectFieldToField(&inst.FlightRestrictionRoute, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FlightRestrictionRouteType](), formDiv)
		}
	}
}

func saveStageSet_FlightRestrictionRouteType_AixmfeaturesStage(
	inst *aixmfeatures.FlightRestrictionRouteType,
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
		case "PriorPermission":
			StageSetFormDivSelectFieldToField(&inst.PriorPermission, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		}
	}
}

func saveStageSet_FlightRestrictionTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.FlightRestrictionTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "FlightRestrictionTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.FlightRestrictionTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FlightRestrictionTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_FlightRestrictionTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.FlightRestrictionTimeSliceType,
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
			StageSetFormDivSelectFieldToField(&inst.Designator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlightRestrictionDesignatorType](), formDiv)
		case "Type_XSD":
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlightRestrictionType](), formDiv)
		case "Instruction":
			StageSetFormDivSelectFieldToField(&inst.Instruction, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextInstructionType](), formDiv)
		case "Flight":
			StageSetFormDivSelectFieldToField(&inst.Flight, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FlightConditionCombinationPropertyType](), formDiv)
		}
	}
}

func saveStageSet_FlightRestrictionType_AixmfeaturesStage(
	inst *aixmfeatures.FlightRestrictionType,
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

func saveStageSet_FlightRoutingElementPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.FlightRoutingElementPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "FlightRoutingElement":
			StageSetFormDivSelectFieldToField(&inst.FlightRoutingElement, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FlightRoutingElementType](), formDiv)
		}
	}
}

func saveStageSet_FlightRoutingElementType_AixmfeaturesStage(
	inst *aixmfeatures.FlightRoutingElementType,
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
		case "OrderNumber":
			StageSetFormDivSelectFieldToField(&inst.OrderNumber, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoSequenceType](), formDiv)
		case "Speed":
			StageSetFormDivSelectFieldToField(&inst.Speed, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValSpeedType](), formDiv)
		case "SpeedReference":
			StageSetFormDivSelectFieldToField(&inst.SpeedReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSpeedReferenceType](), formDiv)
		case "SpeedCriteria":
			StageSetFormDivSelectFieldToField(&inst.SpeedCriteria, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeComparisonType](), formDiv)
		case "Element_directFlightElement":
			StageSetFormDivSelectFieldToField(&inst.Element_directFlightElement, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DirectFlightSegmentPropertyType](), formDiv)
		case "Element_routePortionElement":
			StageSetFormDivSelectFieldToField(&inst.Element_routePortionElement, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RoutePortionPropertyType](), formDiv)
		case "Element_standardInstrumentDepartureElement":
			StageSetFormDivSelectFieldToField(&inst.Element_standardInstrumentDepartureElement, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.StandardInstrumentDeparturePropertyType](), formDiv)
		case "Element_standardInstrumentArrivalElement":
			StageSetFormDivSelectFieldToField(&inst.Element_standardInstrumentArrivalElement, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.StandardInstrumentArrivalPropertyType](), formDiv)
		case "PointElement_fixDesignatedPoint":
			StageSetFormDivSelectFieldToField(&inst.PointElement_fixDesignatedPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), formDiv)
		case "PointElement_navaidSystem":
			StageSetFormDivSelectFieldToField(&inst.PointElement_navaidSystem, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), formDiv)
		case "PointElement_position":
			StageSetFormDivSelectFieldToField(&inst.PointElement_position, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), formDiv)
		case "PointElement_runwayPoint":
			StageSetFormDivSelectFieldToField(&inst.PointElement_runwayPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), formDiv)
		case "PointElement_aimingPoint":
			StageSetFormDivSelectFieldToField(&inst.PointElement_aimingPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), formDiv)
		case "PointElement_airportReferencePoint":
			StageSetFormDivSelectFieldToField(&inst.PointElement_airportReferencePoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), formDiv)
		case "Element_airspaceElement":
			StageSetFormDivSelectFieldToField(&inst.Element_airspaceElement, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspacePropertyType](), formDiv)
		case "Element_airportHeliportElement":
			StageSetFormDivSelectFieldToField(&inst.Element_airportHeliportElement, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), formDiv)
		case "Element_aerialRefuellingElement":
			StageSetFormDivSelectFieldToField(&inst.Element_aerialRefuellingElement, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AerialRefuellingPropertyType](), formDiv)
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

func saveStageSet_MeteorologyPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.MeteorologyPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "Meteorology":
			StageSetFormDivSelectFieldToField(&inst.Meteorology, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.MeteorologyType](), formDiv)
		}
	}
}

func saveStageSet_MeteorologyType_AixmfeaturesStage(
	inst *aixmfeatures.MeteorologyType,
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
		case "FlightConditions":
			StageSetFormDivSelectFieldToField(&inst.FlightConditions, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeMeteoConditionsType](), formDiv)
		case "Visibility":
			StageSetFormDivSelectFieldToField(&inst.Visibility, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "VisibilityInterpretation":
			StageSetFormDivSelectFieldToField(&inst.VisibilityInterpretation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeValueInterpretationType](), formDiv)
		case "RunwayVisualRange":
			StageSetFormDivSelectFieldToField(&inst.RunwayVisualRange, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "RunwayVisualRangeInterpretation":
			StageSetFormDivSelectFieldToField(&inst.RunwayVisualRangeInterpretation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeValueInterpretationType](), formDiv)
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

func saveStageSet_RouteAvailabilityPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.RouteAvailabilityPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "RouteAvailability":
			StageSetFormDivSelectFieldToField(&inst.RouteAvailability, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RouteAvailabilityType](), formDiv)
		}
	}
}

func saveStageSet_RouteAvailabilityType_AixmfeaturesStage(
	inst *aixmfeatures.RouteAvailabilityType,
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
		case "Direction":
			StageSetFormDivSelectFieldToField(&inst.Direction, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDirectionType](), formDiv)
		case "CardinalDirection":
			StageSetFormDivSelectFieldToField(&inst.CardinalDirection, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCardinalDirectionType](), formDiv)
		case "Status":
			StageSetFormDivSelectFieldToField(&inst.Status, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRouteAvailabilityType](), formDiv)
		}
	}
}

func saveStageSet_RouteDMETimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.RouteDMETimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "RouteDMETimeSlice":
			StageSetFormDivSelectFieldToField(&inst.RouteDMETimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RouteDMETimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_RouteDMETimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.RouteDMETimeSliceType,
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
		case "ReferencedDME":
			StageSetFormDivSelectFieldToField(&inst.ReferencedDME, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DMEPropertyType](), formDiv)
		case "ApplicableRoutePortion":
			StageSetFormDivSelectFieldToField(&inst.ApplicableRoutePortion, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RoutePortionPropertyType](), formDiv)
		}
	}
}

func saveStageSet_RouteDMEType_AixmfeaturesStage(
	inst *aixmfeatures.RouteDMEType,
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

func saveStageSet_RoutePortionPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.RoutePortionPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "RoutePortion":
			StageSetFormDivSelectFieldToField(&inst.RoutePortion, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RoutePortionType](), formDiv)
		}
	}
}

func saveStageSet_RoutePortionType_AixmfeaturesStage(
	inst *aixmfeatures.RoutePortionType,
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
		case "Start_fixDesignatedPoint":
			StageSetFormDivSelectFieldToField(&inst.Start_fixDesignatedPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), formDiv)
		case "Start_navaidSystem":
			StageSetFormDivSelectFieldToField(&inst.Start_navaidSystem, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), formDiv)
		case "Start_position":
			StageSetFormDivSelectFieldToField(&inst.Start_position, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), formDiv)
		case "Start_runwayPoint":
			StageSetFormDivSelectFieldToField(&inst.Start_runwayPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), formDiv)
		case "Start_aimingPoint":
			StageSetFormDivSelectFieldToField(&inst.Start_aimingPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), formDiv)
		case "Start_airportReferencePoint":
			StageSetFormDivSelectFieldToField(&inst.Start_airportReferencePoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), formDiv)
		case "IntermediatePoint_fixDesignatedPoint":
			StageSetFormDivSelectFieldToField(&inst.IntermediatePoint_fixDesignatedPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), formDiv)
		case "IntermediatePoint_navaidSystem":
			StageSetFormDivSelectFieldToField(&inst.IntermediatePoint_navaidSystem, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), formDiv)
		case "IntermediatePoint_position":
			StageSetFormDivSelectFieldToField(&inst.IntermediatePoint_position, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), formDiv)
		case "IntermediatePoint_runwayPoint":
			StageSetFormDivSelectFieldToField(&inst.IntermediatePoint_runwayPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), formDiv)
		case "IntermediatePoint_aimingPoint":
			StageSetFormDivSelectFieldToField(&inst.IntermediatePoint_aimingPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), formDiv)
		case "IntermediatePoint_airportReferencePoint":
			StageSetFormDivSelectFieldToField(&inst.IntermediatePoint_airportReferencePoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), formDiv)
		case "ReferencedRoute":
			StageSetFormDivSelectFieldToField(&inst.ReferencedRoute, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RoutePropertyType](), formDiv)
		case "End_fixDesignatedPoint":
			StageSetFormDivSelectFieldToField(&inst.End_fixDesignatedPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), formDiv)
		case "End_navaidSystem":
			StageSetFormDivSelectFieldToField(&inst.End_navaidSystem, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), formDiv)
		case "End_position":
			StageSetFormDivSelectFieldToField(&inst.End_position, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), formDiv)
		case "End_runwayPoint":
			StageSetFormDivSelectFieldToField(&inst.End_runwayPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), formDiv)
		case "End_aimingPoint":
			StageSetFormDivSelectFieldToField(&inst.End_aimingPoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), formDiv)
		case "End_airportReferencePoint":
			StageSetFormDivSelectFieldToField(&inst.End_airportReferencePoint, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), formDiv)
		}
	}
}

func saveStageSet_RoutePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.RoutePropertyType,
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

func saveStageSet_RouteSegmentTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.RouteSegmentTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "RouteSegmentTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.RouteSegmentTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RouteSegmentTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_RouteSegmentTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.RouteSegmentTimeSliceType,
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
		case "Level":
			StageSetFormDivSelectFieldToField(&inst.Level, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeLevelType](), formDiv)
		case "UpperLimit":
			StageSetFormDivSelectFieldToField(&inst.UpperLimit, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "UpperLimitReference":
			StageSetFormDivSelectFieldToField(&inst.UpperLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "LowerLimit":
			StageSetFormDivSelectFieldToField(&inst.LowerLimit, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "LowerLimitReference":
			StageSetFormDivSelectFieldToField(&inst.LowerLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "MinimumObstacleClearanceAltitude":
			StageSetFormDivSelectFieldToField(&inst.MinimumObstacleClearanceAltitude, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "PathType":
			StageSetFormDivSelectFieldToField(&inst.PathType, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRouteSegmentPathType](), formDiv)
		case "TrueTrack":
			StageSetFormDivSelectFieldToField(&inst.TrueTrack, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "MagneticTrack":
			StageSetFormDivSelectFieldToField(&inst.MagneticTrack, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "ReverseTrueTrack":
			StageSetFormDivSelectFieldToField(&inst.ReverseTrueTrack, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "ReverseMagneticTrack":
			StageSetFormDivSelectFieldToField(&inst.ReverseMagneticTrack, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), formDiv)
		case "Length":
			StageSetFormDivSelectFieldToField(&inst.Length, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "WidthLeft":
			StageSetFormDivSelectFieldToField(&inst.WidthLeft, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "WidthRight":
			StageSetFormDivSelectFieldToField(&inst.WidthRight, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "TurnDirection":
			StageSetFormDivSelectFieldToField(&inst.TurnDirection, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDirectionTurnType](), formDiv)
		case "SignalGap":
			StageSetFormDivSelectFieldToField(&inst.SignalGap, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
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
		case "NavigationType":
			StageSetFormDivSelectFieldToField(&inst.NavigationType, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRouteNavigationType](), formDiv)
		case "RequiredNavigationPerformance":
			StageSetFormDivSelectFieldToField(&inst.RequiredNavigationPerformance, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRNPType](), formDiv)
		case "DesignatorSuffix":
			StageSetFormDivSelectFieldToField(&inst.DesignatorSuffix, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRouteDesignatorSuffixType](), formDiv)
		case "Start":
			StageSetFormDivSelectFieldToField(&inst.Start, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.EnRouteSegmentPointPropertyType](), formDiv)
		case "RouteFormed":
			StageSetFormDivSelectFieldToField(&inst.RouteFormed, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RoutePropertyType](), formDiv)
		case "EvaluationArea":
			StageSetFormDivSelectFieldToField(&inst.EvaluationArea, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ObstacleAssessmentAreaPropertyType](), formDiv)
		case "CurveExtent":
			StageSetFormDivSelectFieldToField(&inst.CurveExtent, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CurvePropertyType](), formDiv)
		case "End":
			StageSetFormDivSelectFieldToField(&inst.End, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.EnRouteSegmentPointPropertyType](), formDiv)
		}
	}
}

func saveStageSet_RouteSegmentType_AixmfeaturesStage(
	inst *aixmfeatures.RouteSegmentType,
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

func saveStageSet_RouteTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.RouteTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "RouteTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.RouteTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RouteTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_RouteTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.RouteTimeSliceType,
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
		case "DesignatorPrefix":
			StageSetFormDivSelectFieldToField(&inst.DesignatorPrefix, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRouteDesignatorPrefixType](), formDiv)
		case "DesignatorSecondLetter":
			StageSetFormDivSelectFieldToField(&inst.DesignatorSecondLetter, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRouteDesignatorLetterType](), formDiv)
		case "DesignatorNumber":
			StageSetFormDivSelectFieldToField(&inst.DesignatorNumber, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoNumberType](), formDiv)
		case "MultipleIdentifier":
			StageSetFormDivSelectFieldToField(&inst.MultipleIdentifier, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeUpperAlphaType](), formDiv)
		case "LocationDesignator":
			StageSetFormDivSelectFieldToField(&inst.LocationDesignator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextDesignatorType](), formDiv)
		case "Type_XSD":
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRouteType](), formDiv)
		case "FlightRule":
			StageSetFormDivSelectFieldToField(&inst.FlightRule, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlightRuleType](), formDiv)
		case "InternationalUse":
			StageSetFormDivSelectFieldToField(&inst.InternationalUse, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRouteOriginType](), formDiv)
		case "MilitaryUse":
			StageSetFormDivSelectFieldToField(&inst.MilitaryUse, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeMilitaryStatusType](), formDiv)
		case "MilitaryTrainingType":
			StageSetFormDivSelectFieldToField(&inst.MilitaryTrainingType, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeMilitaryTrainingType](), formDiv)
		case "UserOrganisation":
			StageSetFormDivSelectFieldToField(&inst.UserOrganisation, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.OrganisationAuthorityPropertyType](), formDiv)
		}
	}
}

func saveStageSet_RouteType_AixmfeaturesStage(
	inst *aixmfeatures.RouteType,
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

func saveStageSet_RulesProceduresTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.RulesProceduresTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "RulesProceduresTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.RulesProceduresTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RulesProceduresTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_RulesProceduresTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.RulesProceduresTimeSliceType,
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
			StageSetFormDivSelectFieldToField(&inst.Category, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRuleProcedureType](), formDiv)
		case "Title":
			StageSetFormDivSelectFieldToField(&inst.Title, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRuleProcedureTitleType](), formDiv)
		case "Content":
			StageSetFormDivSelectFieldToField(&inst.Content, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.XHTMLType](), formDiv)
		}
	}
}

func saveStageSet_RulesProceduresType_AixmfeaturesStage(
	inst *aixmfeatures.RulesProceduresType,
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

func saveStageSet_StandardLevelColumnPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.StandardLevelColumnPropertyType,
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

func saveStageSet_CodeAerialRefuellingPrefixType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeAerialRefuellingPrefixType,
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

func saveStageSet_CodeAerialRefuellingType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeAerialRefuellingType,
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

func saveStageSet_CodeAuthorityType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeAuthorityType,
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

func saveStageSet_CodeComparisonType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeComparisonType,
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

func saveStageSet_CodeDirectionType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeDirectionType,
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

func saveStageSet_CodeFlightOriginType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeFlightOriginType,
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

func saveStageSet_CodeFlightPurposeType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeFlightPurposeType,
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

func saveStageSet_CodeFlightRestrictionDesignatorType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeFlightRestrictionDesignatorType,
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

func saveStageSet_CodeFlightRestrictionType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeFlightRestrictionType,
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

func saveStageSet_CodeFlightRuleType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeFlightRuleType,
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

func saveStageSet_CodeFlightStatusType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeFlightStatusType,
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

func saveStageSet_CodeFlightType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeFlightType,
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

func saveStageSet_CodeFlowConditionOperationType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeFlowConditionOperationType,
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

func saveStageSet_CodeLevelType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeLevelType,
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

func saveStageSet_CodeLocationQualifierType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeLocationQualifierType,
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

func saveStageSet_CodeMeteoConditionsType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeMeteoConditionsType,
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

func saveStageSet_CodeMilitaryStatusType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeMilitaryStatusType,
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

func saveStageSet_CodeMilitaryTrainingType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeMilitaryTrainingType,
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

func saveStageSet_CodeRouteAvailabilityType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeRouteAvailabilityType,
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

func saveStageSet_CodeRouteDesignatorLetterType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeRouteDesignatorLetterType,
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

func saveStageSet_CodeRouteDesignatorPrefixType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeRouteDesignatorPrefixType,
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

func saveStageSet_CodeRouteDesignatorSuffixType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeRouteDesignatorSuffixType,
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

func saveStageSet_CodeRouteNavigationType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeRouteNavigationType,
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

func saveStageSet_CodeRouteOriginType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeRouteOriginType,
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

func saveStageSet_CodeRouteSegmentPathType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeRouteSegmentPathType,
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

func saveStageSet_CodeRouteType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeRouteType,
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

func saveStageSet_CodeRuleProcedureTitleType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeRuleProcedureTitleType,
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

func saveStageSet_CodeRuleProcedureType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeRuleProcedureType,
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

func saveStageSet_XHTMLType_AixmdatatypesStage(
	inst *aixmdatatypes.XHTMLType,
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

func StageSetNewInstance_AbstractDirectFlightType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AbstractDirectFlightType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AbstractDirectFlightType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AbstractDirectFlightType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AbstractDirectFlightType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_AerialRefuellingAnchorPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AerialRefuellingAnchorPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AerialRefuellingAnchorPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AerialRefuellingAnchorPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AerialRefuellingAnchorPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AerialRefuellingAnchorType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AerialRefuellingAnchorType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AerialRefuellingAnchorType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AerialRefuellingAnchorType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AerialRefuellingAnchorType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AerialRefuellingPointPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AerialRefuellingPointPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AerialRefuellingPointPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AerialRefuellingPointPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AerialRefuellingPointPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_AerialRefuellingPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AerialRefuellingPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AerialRefuellingPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AerialRefuellingPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AerialRefuellingPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AerialRefuellingTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AerialRefuellingTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AerialRefuellingTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AerialRefuellingTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AerialRefuellingTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AerialRefuellingTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AerialRefuellingTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AerialRefuellingTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AerialRefuellingTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AerialRefuellingTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AerialRefuellingTrackPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AerialRefuellingTrackPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AerialRefuellingTrackPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AerialRefuellingTrackPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AerialRefuellingTrackPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AerialRefuellingTrackType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AerialRefuellingTrackType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AerialRefuellingTrackType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AerialRefuellingTrackType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AerialRefuellingTrackType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AerialRefuellingType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AerialRefuellingType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AerialRefuellingType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AerialRefuellingType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AerialRefuellingType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_AirspaceBorderCrossingPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AirspaceBorderCrossingPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AirspaceBorderCrossingPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AirspaceBorderCrossingPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AirspaceBorderCrossingPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AirspaceBorderCrossingTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AirspaceBorderCrossingTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AirspaceBorderCrossingTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AirspaceBorderCrossingTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AirspaceBorderCrossingTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AirspaceBorderCrossingTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AirspaceBorderCrossingTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AirspaceBorderCrossingTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AirspaceBorderCrossingTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AirspaceBorderCrossingTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AirspaceBorderCrossingType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AirspaceBorderCrossingType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AirspaceBorderCrossingType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AirspaceBorderCrossingType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AirspaceBorderCrossingType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AirspaceLayerPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AirspaceLayerPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AirspaceLayerPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AirspaceLayerPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AirspaceLayerPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AirspaceLayerType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AirspaceLayerType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AirspaceLayerType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AirspaceLayerType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AirspaceLayerType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_AuthorityForAerialRefuellingPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AuthorityForAerialRefuellingPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AuthorityForAerialRefuellingPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AuthorityForAerialRefuellingPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AuthorityForAerialRefuellingPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AuthorityForAerialRefuellingType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AuthorityForAerialRefuellingType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AuthorityForAerialRefuellingType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AuthorityForAerialRefuellingType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AuthorityForAerialRefuellingType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ChangeOverPointTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ChangeOverPointTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ChangeOverPointTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ChangeOverPointTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ChangeOverPointTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ChangeOverPointTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ChangeOverPointTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ChangeOverPointTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ChangeOverPointTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ChangeOverPointTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ChangeOverPointType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ChangeOverPointType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ChangeOverPointType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ChangeOverPointType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ChangeOverPointType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_DirectFlightClassType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DirectFlightClassType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.DirectFlightClassType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_DirectFlightClassType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_DirectFlightClassType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DirectFlightPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DirectFlightPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.DirectFlightPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_DirectFlightPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_DirectFlightPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DirectFlightSegmentPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DirectFlightSegmentPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.DirectFlightSegmentPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_DirectFlightSegmentPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_DirectFlightSegmentPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DirectFlightSegmentType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DirectFlightSegmentType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.DirectFlightSegmentType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_DirectFlightSegmentType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_DirectFlightSegmentType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_EnRouteSegmentPointPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New EnRouteSegmentPointPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.EnRouteSegmentPointPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_EnRouteSegmentPointPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_EnRouteSegmentPointPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_FlightCharacteristicPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New FlightCharacteristicPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.FlightCharacteristicPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_FlightCharacteristicPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_FlightCharacteristicPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FlightCharacteristicType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New FlightCharacteristicType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.FlightCharacteristicType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_FlightCharacteristicType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_FlightCharacteristicType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FlightConditionCircumstancePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New FlightConditionCircumstancePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.FlightConditionCircumstancePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_FlightConditionCircumstancePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_FlightConditionCircumstancePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FlightConditionCircumstanceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New FlightConditionCircumstanceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.FlightConditionCircumstanceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_FlightConditionCircumstanceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_FlightConditionCircumstanceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FlightConditionCombinationPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New FlightConditionCombinationPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.FlightConditionCombinationPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_FlightConditionCombinationPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_FlightConditionCombinationPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FlightConditionCombinationType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New FlightConditionCombinationType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.FlightConditionCombinationType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_FlightConditionCombinationType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_FlightConditionCombinationType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FlightConditionElementPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New FlightConditionElementPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.FlightConditionElementPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_FlightConditionElementPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_FlightConditionElementPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FlightConditionElementType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New FlightConditionElementType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.FlightConditionElementType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_FlightConditionElementType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_FlightConditionElementType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FlightRestrictionLevelPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New FlightRestrictionLevelPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.FlightRestrictionLevelPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_FlightRestrictionLevelPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_FlightRestrictionLevelPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FlightRestrictionLevelType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New FlightRestrictionLevelType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.FlightRestrictionLevelType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_FlightRestrictionLevelType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_FlightRestrictionLevelType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FlightRestrictionRoutePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New FlightRestrictionRoutePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.FlightRestrictionRoutePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_FlightRestrictionRoutePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_FlightRestrictionRoutePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FlightRestrictionRouteType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New FlightRestrictionRouteType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.FlightRestrictionRouteType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_FlightRestrictionRouteType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_FlightRestrictionRouteType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FlightRestrictionTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New FlightRestrictionTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.FlightRestrictionTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_FlightRestrictionTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_FlightRestrictionTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FlightRestrictionTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New FlightRestrictionTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.FlightRestrictionTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_FlightRestrictionTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_FlightRestrictionTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FlightRestrictionType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New FlightRestrictionType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.FlightRestrictionType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_FlightRestrictionType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_FlightRestrictionType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FlightRoutingElementPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New FlightRoutingElementPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.FlightRoutingElementPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_FlightRoutingElementPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_FlightRoutingElementPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_FlightRoutingElementType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New FlightRoutingElementType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.FlightRoutingElementType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_FlightRoutingElementType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_FlightRoutingElementType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_MeteorologyPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New MeteorologyPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.MeteorologyPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_MeteorologyPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_MeteorologyPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_MeteorologyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New MeteorologyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.MeteorologyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_MeteorologyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_MeteorologyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_RouteAvailabilityPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RouteAvailabilityPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.RouteAvailabilityPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_RouteAvailabilityPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_RouteAvailabilityPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RouteAvailabilityType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RouteAvailabilityType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.RouteAvailabilityType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_RouteAvailabilityType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_RouteAvailabilityType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RouteDMETimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RouteDMETimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.RouteDMETimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_RouteDMETimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_RouteDMETimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RouteDMETimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RouteDMETimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.RouteDMETimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_RouteDMETimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_RouteDMETimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RouteDMEType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RouteDMEType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.RouteDMEType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_RouteDMEType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_RouteDMEType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RoutePortionPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RoutePortionPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.RoutePortionPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_RoutePortionPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_RoutePortionPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RoutePortionType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RoutePortionType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.RoutePortionType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_RoutePortionType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_RoutePortionType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RoutePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RoutePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.RoutePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_RoutePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_RoutePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RouteSegmentTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RouteSegmentTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.RouteSegmentTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_RouteSegmentTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_RouteSegmentTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RouteSegmentTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RouteSegmentTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.RouteSegmentTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_RouteSegmentTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_RouteSegmentTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RouteSegmentType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RouteSegmentType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.RouteSegmentType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_RouteSegmentType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_RouteSegmentType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RouteTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RouteTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.RouteTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_RouteTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_RouteTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RouteTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RouteTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.RouteTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_RouteTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_RouteTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RouteType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RouteType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.RouteType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_RouteType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_RouteType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RulesProceduresTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RulesProceduresTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.RulesProceduresTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_RulesProceduresTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_RulesProceduresTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RulesProceduresTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RulesProceduresTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.RulesProceduresTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_RulesProceduresTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_RulesProceduresTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RulesProceduresType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RulesProceduresType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.RulesProceduresType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_RulesProceduresType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_RulesProceduresType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_StandardLevelColumnPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New StandardLevelColumnPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.StandardLevelColumnPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_StandardLevelColumnPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_StandardLevelColumnPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_CodeAerialRefuellingPrefixType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeAerialRefuellingPrefixType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeAerialRefuellingPrefixType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeAerialRefuellingPrefixType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeAerialRefuellingPrefixType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeAerialRefuellingType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeAerialRefuellingType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeAerialRefuellingType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeAerialRefuellingType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeAerialRefuellingType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_CodeAuthorityType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeAuthorityType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeAuthorityType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeAuthorityType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeAuthorityType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_CodeComparisonType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeComparisonType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeComparisonType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeComparisonType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeComparisonType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_CodeDirectionType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeDirectionType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeDirectionType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeDirectionType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeDirectionType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_CodeFlightOriginType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeFlightOriginType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeFlightOriginType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeFlightOriginType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeFlightOriginType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeFlightPurposeType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeFlightPurposeType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeFlightPurposeType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeFlightPurposeType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeFlightPurposeType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeFlightRestrictionDesignatorType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeFlightRestrictionDesignatorType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeFlightRestrictionDesignatorType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeFlightRestrictionDesignatorType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeFlightRestrictionDesignatorType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeFlightRestrictionType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeFlightRestrictionType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeFlightRestrictionType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeFlightRestrictionType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeFlightRestrictionType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeFlightRuleType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeFlightRuleType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeFlightRuleType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeFlightRuleType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeFlightRuleType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeFlightStatusType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeFlightStatusType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeFlightStatusType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeFlightStatusType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeFlightStatusType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeFlightType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeFlightType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeFlightType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeFlightType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeFlightType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeFlowConditionOperationType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeFlowConditionOperationType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeFlowConditionOperationType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeFlowConditionOperationType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeFlowConditionOperationType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_CodeLevelType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeLevelType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeLevelType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeLevelType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeLevelType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeLocationQualifierType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeLocationQualifierType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeLocationQualifierType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeLocationQualifierType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeLocationQualifierType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeMeteoConditionsType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeMeteoConditionsType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeMeteoConditionsType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeMeteoConditionsType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeMeteoConditionsType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_CodeMilitaryStatusType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeMilitaryStatusType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeMilitaryStatusType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeMilitaryStatusType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeMilitaryStatusType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeMilitaryTrainingType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeMilitaryTrainingType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeMilitaryTrainingType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeMilitaryTrainingType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeMilitaryTrainingType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_CodeRouteAvailabilityType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeRouteAvailabilityType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeRouteAvailabilityType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeRouteAvailabilityType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeRouteAvailabilityType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeRouteDesignatorLetterType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeRouteDesignatorLetterType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeRouteDesignatorLetterType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeRouteDesignatorLetterType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeRouteDesignatorLetterType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeRouteDesignatorPrefixType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeRouteDesignatorPrefixType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeRouteDesignatorPrefixType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeRouteDesignatorPrefixType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeRouteDesignatorPrefixType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeRouteDesignatorSuffixType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeRouteDesignatorSuffixType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeRouteDesignatorSuffixType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeRouteDesignatorSuffixType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeRouteDesignatorSuffixType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeRouteNavigationType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeRouteNavigationType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeRouteNavigationType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeRouteNavigationType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeRouteNavigationType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeRouteOriginType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeRouteOriginType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeRouteOriginType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeRouteOriginType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeRouteOriginType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeRouteSegmentPathType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeRouteSegmentPathType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeRouteSegmentPathType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeRouteSegmentPathType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeRouteSegmentPathType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeRouteType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeRouteType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeRouteType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeRouteType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeRouteType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeRuleProcedureTitleType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeRuleProcedureTitleType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeRuleProcedureTitleType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeRuleProcedureTitleType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeRuleProcedureTitleType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeRuleProcedureType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeRuleProcedureType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeRuleProcedureType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeRuleProcedureType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeRuleProcedureType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_XHTMLType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New XHTMLType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.XHTMLType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_XHTMLType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_XHTMLType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

