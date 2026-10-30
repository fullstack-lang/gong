// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	gongprobe "github.com/fullstack-lang/gong/pkg/runtime/probe"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/xlink"
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
		case "Airspace":
			StageSetFormDivSelectFieldToField(&inst.Airspace, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceType](), formDiv)
		case "AuthorityForAirspace":
			StageSetFormDivSelectFieldToField(&inst.AuthorityForAirspace, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AuthorityForAirspaceType](), formDiv)
		case "GeoBorder":
			StageSetFormDivSelectFieldToField(&inst.GeoBorder, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.GeoBorderType](), formDiv)
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

func saveStageSet_AirspaceActivationPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AirspaceActivationPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "AirspaceActivation":
			StageSetFormDivSelectFieldToField(&inst.AirspaceActivation, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceActivationType](), formDiv)
		}
	}
}

func saveStageSet_AirspaceActivationType_AixmfeaturesStage(
	inst *aixmfeatures.AirspaceActivationType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "Activity":
			StageSetFormDivSelectFieldToField(&inst.Activity, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAirspaceActivityType](), formDiv)
		case "Status":
			StageSetFormDivSelectFieldToField(&inst.Status, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeStatusAirspaceType](), formDiv)
		}
	}
}

func saveStageSet_AirspaceGeometryComponentPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AirspaceGeometryComponentPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "AirspaceGeometryComponent":
			StageSetFormDivSelectFieldToField(&inst.AirspaceGeometryComponent, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceGeometryComponentType](), formDiv)
		}
	}
}

func saveStageSet_AirspaceGeometryComponentType_AixmfeaturesStage(
	inst *aixmfeatures.AirspaceGeometryComponentType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "Operation":
			StageSetFormDivSelectFieldToField(&inst.Operation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAirspaceAggregationType](), formDiv)
		case "OperationSequence":
			StageSetFormDivSelectFieldToField(&inst.OperationSequence, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoSequenceType](), formDiv)
		case "TheAirspaceVolume":
			StageSetFormDivSelectFieldToField(&inst.TheAirspaceVolume, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceVolumePropertyType](), formDiv)
		}
	}
}

func saveStageSet_AirspaceLayerClassPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AirspaceLayerClassPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "AirspaceLayerClass":
			StageSetFormDivSelectFieldToField(&inst.AirspaceLayerClass, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceLayerClassType](), formDiv)
		}
	}
}

func saveStageSet_AirspaceLayerClassType_AixmfeaturesStage(
	inst *aixmfeatures.AirspaceLayerClassType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "Classification":
			StageSetFormDivSelectFieldToField(&inst.Classification, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAirspaceClassificationType](), formDiv)
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

func saveStageSet_AirspaceTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AirspaceTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "AirspaceTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.AirspaceTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_AirspaceTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.AirspaceTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
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
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAirspaceType](), formDiv)
		case "Designator":
			StageSetFormDivSelectFieldToField(&inst.Designator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAirspaceDesignatorType](), formDiv)
		case "LocalType":
			StageSetFormDivSelectFieldToField(&inst.LocalType, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextNameType](), formDiv)
		case "DesignatorICAO":
			StageSetFormDivSelectFieldToField(&inst.DesignatorICAO, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "ControlType":
			StageSetFormDivSelectFieldToField(&inst.ControlType, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeMilitaryOperationsType](), formDiv)
		case "UpperLowerSeparation":
			StageSetFormDivSelectFieldToField(&inst.UpperLowerSeparation, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValFLType](), formDiv)
		case "ProtectedRoute":
			StageSetFormDivSelectFieldToField(&inst.ProtectedRoute, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RoutePropertyType](), formDiv)
		}
	}
}

func saveStageSet_AirspaceType_AixmfeaturesStage(
	inst *aixmfeatures.AirspaceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
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

func saveStageSet_AirspaceVolumeDependencyPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AirspaceVolumeDependencyPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "AirspaceVolumeDependency":
			StageSetFormDivSelectFieldToField(&inst.AirspaceVolumeDependency, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceVolumeDependencyType](), formDiv)
		}
	}
}

func saveStageSet_AirspaceVolumeDependencyType_AixmfeaturesStage(
	inst *aixmfeatures.AirspaceVolumeDependencyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "Dependency":
			StageSetFormDivSelectFieldToField(&inst.Dependency, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAirspaceDependencyType](), formDiv)
		case "TheAirspace":
			StageSetFormDivSelectFieldToField(&inst.TheAirspace, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspacePropertyType](), formDiv)
		}
	}
}

func saveStageSet_AirspaceVolumePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AirspaceVolumePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "AirspaceVolume":
			StageSetFormDivSelectFieldToField(&inst.AirspaceVolume, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceVolumeType](), formDiv)
		}
	}
}

func saveStageSet_AirspaceVolumeType_AixmfeaturesStage(
	inst *aixmfeatures.AirspaceVolumeType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
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
		case "MaximumLimit":
			StageSetFormDivSelectFieldToField(&inst.MaximumLimit, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "MaximumLimitReference":
			StageSetFormDivSelectFieldToField(&inst.MaximumLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "LowerLimit":
			StageSetFormDivSelectFieldToField(&inst.LowerLimit, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "LowerLimitReference":
			StageSetFormDivSelectFieldToField(&inst.LowerLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "MinimumLimit":
			StageSetFormDivSelectFieldToField(&inst.MinimumLimit, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), formDiv)
		case "MinimumLimitReference":
			StageSetFormDivSelectFieldToField(&inst.MinimumLimitReference, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), formDiv)
		case "Width":
			StageSetFormDivSelectFieldToField(&inst.Width, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "HorizontalProjection":
			StageSetFormDivSelectFieldToField(&inst.HorizontalProjection, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SurfacePropertyType](), formDiv)
		case "Centreline":
			StageSetFormDivSelectFieldToField(&inst.Centreline, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CurvePropertyType](), formDiv)
		case "ContributorAirspace":
			StageSetFormDivSelectFieldToField(&inst.ContributorAirspace, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceVolumeDependencyPropertyType](), formDiv)
		}
	}
}

func saveStageSet_AuthorityForAirspaceTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AuthorityForAirspaceTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "AuthorityForAirspaceTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.AuthorityForAirspaceTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AuthorityForAirspaceTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_AuthorityForAirspaceTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.AuthorityForAirspaceTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
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
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAuthorityType](), formDiv)
		case "ResponsibleOrganisation":
			StageSetFormDivSelectFieldToField(&inst.ResponsibleOrganisation, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.OrganisationAuthorityPropertyType](), formDiv)
		case "AssignedAirspace":
			StageSetFormDivSelectFieldToField(&inst.AssignedAirspace, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspacePropertyType](), formDiv)
		}
	}
}

func saveStageSet_AuthorityForAirspaceType_AixmfeaturesStage(
	inst *aixmfeatures.AuthorityForAirspaceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
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

func saveStageSet_GeoBorderTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.GeoBorderTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "GeoBorderTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.GeoBorderTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.GeoBorderTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_GeoBorderTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.GeoBorderTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
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
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeGeoBorderType](), formDiv)
		case "Border":
			StageSetFormDivSelectFieldToField(&inst.Border, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CurvePropertyType](), formDiv)
		}
	}
}

func saveStageSet_GeoBorderType_AixmfeaturesStage(
	inst *aixmfeatures.GeoBorderType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
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

func saveStageSet_CodeAirspaceActivityType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeAirspaceActivityType,
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

func saveStageSet_CodeAirspaceAggregationType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeAirspaceAggregationType,
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

func saveStageSet_CodeAirspaceClassificationType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeAirspaceClassificationType,
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

func saveStageSet_CodeAirspaceDependencyType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeAirspaceDependencyType,
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

func saveStageSet_CodeAirspaceDesignatorType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeAirspaceDesignatorType,
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

func saveStageSet_CodeAirspaceType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeAirspaceType,
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

func saveStageSet_CodeGeoBorderType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeGeoBorderType,
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

func saveStageSet_CodeMilitaryOperationsType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeMilitaryOperationsType,
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

func saveStageSet_CodeStatusAirspaceType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeStatusAirspaceType,
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

func saveStageSet_ValFLType_AixmdatatypesStage(
	inst *aixmdatatypes.ValFLType,
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

func StageSetNewInstance_AirspaceActivationPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AirspaceActivationPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AirspaceActivationPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AirspaceActivationPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AirspaceActivationPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AirspaceActivationType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AirspaceActivationType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AirspaceActivationType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AirspaceActivationType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AirspaceActivationType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AirspaceGeometryComponentPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AirspaceGeometryComponentPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AirspaceGeometryComponentPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AirspaceGeometryComponentPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AirspaceGeometryComponentPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AirspaceGeometryComponentType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AirspaceGeometryComponentType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AirspaceGeometryComponentType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AirspaceGeometryComponentType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AirspaceGeometryComponentType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AirspaceLayerClassPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AirspaceLayerClassPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AirspaceLayerClassPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AirspaceLayerClassPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AirspaceLayerClassPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AirspaceLayerClassType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AirspaceLayerClassType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AirspaceLayerClassType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AirspaceLayerClassType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AirspaceLayerClassType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_AirspaceTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AirspaceTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AirspaceTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AirspaceTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AirspaceTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AirspaceTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AirspaceTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AirspaceTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AirspaceTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AirspaceTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AirspaceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AirspaceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AirspaceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AirspaceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AirspaceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AirspaceVolumeDependencyPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AirspaceVolumeDependencyPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AirspaceVolumeDependencyPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AirspaceVolumeDependencyPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AirspaceVolumeDependencyPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AirspaceVolumeDependencyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AirspaceVolumeDependencyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AirspaceVolumeDependencyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AirspaceVolumeDependencyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AirspaceVolumeDependencyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AirspaceVolumePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AirspaceVolumePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AirspaceVolumePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AirspaceVolumePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AirspaceVolumePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AirspaceVolumeType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AirspaceVolumeType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AirspaceVolumeType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AirspaceVolumeType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AirspaceVolumeType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AuthorityForAirspaceTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AuthorityForAirspaceTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AuthorityForAirspaceTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AuthorityForAirspaceTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AuthorityForAirspaceTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AuthorityForAirspaceTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AuthorityForAirspaceTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AuthorityForAirspaceTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AuthorityForAirspaceTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AuthorityForAirspaceTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AuthorityForAirspaceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AuthorityForAirspaceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AuthorityForAirspaceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AuthorityForAirspaceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AuthorityForAirspaceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_GeoBorderTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New GeoBorderTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.GeoBorderTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_GeoBorderTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_GeoBorderTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_GeoBorderTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New GeoBorderTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.GeoBorderTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_GeoBorderTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_GeoBorderTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_GeoBorderType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New GeoBorderType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.GeoBorderType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_GeoBorderType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_GeoBorderType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_CodeAirspaceActivityType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeAirspaceActivityType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeAirspaceActivityType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeAirspaceActivityType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeAirspaceActivityType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeAirspaceAggregationType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeAirspaceAggregationType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeAirspaceAggregationType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeAirspaceAggregationType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeAirspaceAggregationType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeAirspaceClassificationType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeAirspaceClassificationType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeAirspaceClassificationType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeAirspaceClassificationType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeAirspaceClassificationType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeAirspaceDependencyType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeAirspaceDependencyType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeAirspaceDependencyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeAirspaceDependencyType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeAirspaceDependencyType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeAirspaceDesignatorType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeAirspaceDesignatorType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeAirspaceDesignatorType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeAirspaceDesignatorType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeAirspaceDesignatorType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeAirspaceType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeAirspaceType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeAirspaceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeAirspaceType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeAirspaceType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_CodeGeoBorderType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeGeoBorderType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeGeoBorderType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeGeoBorderType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeGeoBorderType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeMilitaryOperationsType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeMilitaryOperationsType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeMilitaryOperationsType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeMilitaryOperationsType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeMilitaryOperationsType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_CodeStatusAirspaceType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeStatusAirspaceType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeStatusAirspaceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeStatusAirspaceType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeStatusAirspaceType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_ValFLType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ValFLType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.ValFLType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_ValFLType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_ValFLType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

