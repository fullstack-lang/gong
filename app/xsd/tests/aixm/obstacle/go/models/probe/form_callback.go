// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	gongprobe "github.com/fullstack-lang/gong/pkg/runtime/probe"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/xlink"
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
		case "VerticalStructure":
			StageSetFormDivSelectFieldToField(&inst.VerticalStructure, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.VerticalStructureType](), formDiv)
		case "ObstacleArea":
			StageSetFormDivSelectFieldToField(&inst.ObstacleArea, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ObstacleAreaType](), formDiv)
		case "AeronauticalGroundLight":
			StageSetFormDivSelectFieldToField(&inst.AeronauticalGroundLight, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AeronauticalGroundLightType](), formDiv)
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

func saveStageSet_AeronauticalGroundLightTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.AeronauticalGroundLightTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "AeronauticalGroundLightTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.AeronauticalGroundLightTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AeronauticalGroundLightTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_AeronauticalGroundLightTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.AeronauticalGroundLightTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
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
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeGroundLightingType](), formDiv)
		case "Colour":
			StageSetFormDivSelectFieldToField(&inst.Colour, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeColourType](), formDiv)
		case "Flashing":
			StageSetFormDivSelectFieldToField(&inst.Flashing, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "StructureBeacon":
			StageSetFormDivSelectFieldToField(&inst.StructureBeacon, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.VerticalStructurePropertyType](), formDiv)
		case "AerodromeBeacon":
			StageSetFormDivSelectFieldToField(&inst.AerodromeBeacon, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), formDiv)
		case "Location":
			StageSetFormDivSelectFieldToField(&inst.Location, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), formDiv)
		}
	}
}

func saveStageSet_AeronauticalGroundLightType_AixmfeaturesStage(
	inst *aixmfeatures.AeronauticalGroundLightType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
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

func saveStageSet_ElevatedCurvePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.ElevatedCurvePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "ElevatedCurve":
			StageSetFormDivSelectFieldToField(&inst.ElevatedCurve, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedCurveType](), formDiv)
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

func saveStageSet_ElevatedSurfacePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.ElevatedSurfacePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "ElevatedSurface":
			StageSetFormDivSelectFieldToField(&inst.ElevatedSurface, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedSurfaceType](), formDiv)
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

func saveStageSet_GroundLightSystemPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.GroundLightSystemPropertyType,
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

func saveStageSet_LightElementPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.LightElementPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "LightElement":
			StageSetFormDivSelectFieldToField(&inst.LightElement, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.LightElementType](), formDiv)
		}
	}
}

func saveStageSet_LightElementStatusPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.LightElementStatusPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "LightElementStatus":
			StageSetFormDivSelectFieldToField(&inst.LightElementStatus, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.LightElementStatusType](), formDiv)
		}
	}
}

func saveStageSet_LightElementStatusType_AixmfeaturesStage(
	inst *aixmfeatures.LightElementStatusType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
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
			StageSetFormDivSelectFieldToField(&inst.Status, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeStatusOperationsType](), formDiv)
		}
	}
}

func saveStageSet_LightElementType_AixmfeaturesStage(
	inst *aixmfeatures.LightElementType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "Colour":
			StageSetFormDivSelectFieldToField(&inst.Colour, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeColourType](), formDiv)
		case "IntensityLevel":
			StageSetFormDivSelectFieldToField(&inst.IntensityLevel, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeLightIntensityType](), formDiv)
		case "Intensity":
			StageSetFormDivSelectFieldToField(&inst.Intensity, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValLightIntensityType](), formDiv)
		case "Type_XSD":
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeLightSourceType](), formDiv)
		case "Location":
			StageSetFormDivSelectFieldToField(&inst.Location, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), formDiv)
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

func saveStageSet_MarkerBeaconPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.MarkerBeaconPropertyType,
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

func saveStageSet_ObstacleAreaTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.ObstacleAreaTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "ObstacleAreaTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.ObstacleAreaTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ObstacleAreaTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_ObstacleAreaTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.ObstacleAreaTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
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
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeObstacleAreaType](), formDiv)
		case "ObstructionIdSurfaceCondition":
			StageSetFormDivSelectFieldToField(&inst.ObstructionIdSurfaceCondition, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeObstacleAssessmentSurfaceType](), formDiv)
		case "Reference_ownerOrganisation":
			StageSetFormDivSelectFieldToField(&inst.Reference_ownerOrganisation, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.OrganisationAuthorityPropertyType](), formDiv)
		case "Reference_ownerRunway":
			StageSetFormDivSelectFieldToField(&inst.Reference_ownerRunway, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayDirectionPropertyType](), formDiv)
		case "Reference_ownerAirport":
			StageSetFormDivSelectFieldToField(&inst.Reference_ownerAirport, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), formDiv)
		case "SurfaceExtent":
			StageSetFormDivSelectFieldToField(&inst.SurfaceExtent, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SurfacePropertyType](), formDiv)
		}
	}
}

func saveStageSet_ObstacleAreaType_AixmfeaturesStage(
	inst *aixmfeatures.ObstacleAreaType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
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

func saveStageSet_PassengerServicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.PassengerServicePropertyType,
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

func saveStageSet_ServicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.ServicePropertyType,
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

func saveStageSet_SpecialNavigationStationPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.SpecialNavigationStationPropertyType,
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

func saveStageSet_VerticalStructureLightingStatusPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.VerticalStructureLightingStatusPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "VerticalStructureLightingStatus":
			StageSetFormDivSelectFieldToField(&inst.VerticalStructureLightingStatus, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.VerticalStructureLightingStatusType](), formDiv)
		}
	}
}

func saveStageSet_VerticalStructureLightingStatusType_AixmfeaturesStage(
	inst *aixmfeatures.VerticalStructureLightingStatusType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
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
			StageSetFormDivSelectFieldToField(&inst.Status, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeStatusOperationsType](), formDiv)
		}
	}
}

func saveStageSet_VerticalStructurePartPropertyType_AixmfeaturesStage(
	inst *aixmfeatures.VerticalStructurePartPropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "VerticalStructurePart":
			StageSetFormDivSelectFieldToField(&inst.VerticalStructurePart, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.VerticalStructurePartType](), formDiv)
		}
	}
}

func saveStageSet_VerticalStructurePartType_AixmfeaturesStage(
	inst *aixmfeatures.VerticalStructurePartType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Id":
			FormDivBasicFieldToField(&inst.Id, formDiv)
		case "Description":
			FormDivBasicFieldToField(&inst.Description, formDiv)
		case "Identifier":
			StageSetFormDivSelectFieldToField(&inst.Identifier, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), formDiv)
		case "VerticalExtent":
			StageSetFormDivSelectFieldToField(&inst.VerticalExtent, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "VerticalExtentAccuracy":
			StageSetFormDivSelectFieldToField(&inst.VerticalExtentAccuracy, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "Type_XSD":
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalStructureType](), formDiv)
		case "ConstructionStatus":
			StageSetFormDivSelectFieldToField(&inst.ConstructionStatus, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeStatusConstructionType](), formDiv)
		case "MarkingPattern":
			StageSetFormDivSelectFieldToField(&inst.MarkingPattern, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalStructureMarkingType](), formDiv)
		case "MarkingFirstColour":
			StageSetFormDivSelectFieldToField(&inst.MarkingFirstColour, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeColourType](), formDiv)
		case "MarkingSecondColour":
			StageSetFormDivSelectFieldToField(&inst.MarkingSecondColour, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeColourType](), formDiv)
		case "Mobile":
			StageSetFormDivSelectFieldToField(&inst.Mobile, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "Frangible":
			StageSetFormDivSelectFieldToField(&inst.Frangible, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "VisibleMaterial":
			StageSetFormDivSelectFieldToField(&inst.VisibleMaterial, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalStructureMaterialType](), formDiv)
		case "Designator":
			StageSetFormDivSelectFieldToField(&inst.Designator, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextDesignatorType](), formDiv)
		case "HorizontalProjection_surfaceExtent":
			StageSetFormDivSelectFieldToField(&inst.HorizontalProjection_surfaceExtent, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedSurfacePropertyType](), formDiv)
		case "HorizontalProjection_linearExtent":
			StageSetFormDivSelectFieldToField(&inst.HorizontalProjection_linearExtent, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedCurvePropertyType](), formDiv)
		case "HorizontalProjection_location":
			StageSetFormDivSelectFieldToField(&inst.HorizontalProjection_location, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), formDiv)
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

func saveStageSet_VerticalStructureTimeSlicePropertyType_AixmfeaturesStage(
	inst *aixmfeatures.VerticalStructureTimeSlicePropertyType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Owns":
			FormDivBasicFieldToField(&inst.Owns, formDiv)
		case "VerticalStructureTimeSlice":
			StageSetFormDivSelectFieldToField(&inst.VerticalStructureTimeSlice, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.VerticalStructureTimeSliceType](), formDiv)
		}
	}
}

func saveStageSet_VerticalStructureTimeSliceType_AixmfeaturesStage(
	inst *aixmfeatures.VerticalStructureTimeSliceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
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
			StageSetFormDivSelectFieldToField(&inst.Type_XSD, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalStructureType](), formDiv)
		case "Lighted":
			StageSetFormDivSelectFieldToField(&inst.Lighted, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "MarkingICAOStandard":
			StageSetFormDivSelectFieldToField(&inst.MarkingICAOStandard, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "Group":
			StageSetFormDivSelectFieldToField(&inst.Group, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "Length":
			StageSetFormDivSelectFieldToField(&inst.Length, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "Width":
			StageSetFormDivSelectFieldToField(&inst.Width, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "Radius":
			StageSetFormDivSelectFieldToField(&inst.Radius, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), formDiv)
		case "LightingICAOStandard":
			StageSetFormDivSelectFieldToField(&inst.LightingICAOStandard, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "SynchronisedLighting":
			StageSetFormDivSelectFieldToField(&inst.SynchronisedLighting, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), formDiv)
		case "Marker":
			StageSetFormDivSelectFieldToField(&inst.Marker, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.MarkerBeaconPropertyType](), formDiv)
		}
	}
}

func saveStageSet_VerticalStructureType_AixmfeaturesStage(
	inst *aixmfeatures.VerticalStructureType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
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

func saveStageSet_CodeColourType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeColourType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
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

func saveStageSet_CodeGroundLightingType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeGroundLightingType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeLightIntensityType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeLightIntensityType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeLightSourceType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeLightSourceType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
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

func saveStageSet_CodeObstacleAreaType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeObstacleAreaType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
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

func saveStageSet_CodeStatusConstructionType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeStatusConstructionType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeStatusOperationsType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeStatusOperationsType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
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

func saveStageSet_CodeVerticalStructureMarkingType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeVerticalStructureMarkingType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeVerticalStructureMaterialType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeVerticalStructureMaterialType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "NilReason":
			FormDivBasicFieldToField(&inst.NilReason, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_CodeVerticalStructureType_AixmdatatypesStage(
	inst *aixmdatatypes.CodeVerticalStructureType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
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

func saveStageSet_ValLightIntensityType_AixmdatatypesStage(
	inst *aixmdatatypes.ValLightIntensityType,
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

func StageSetNewInstance_AeronauticalGroundLightTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AeronauticalGroundLightTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AeronauticalGroundLightTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AeronauticalGroundLightTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AeronauticalGroundLightTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AeronauticalGroundLightTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AeronauticalGroundLightTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AeronauticalGroundLightTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AeronauticalGroundLightTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AeronauticalGroundLightTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_AeronauticalGroundLightType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New AeronauticalGroundLightType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.AeronauticalGroundLightType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_AeronauticalGroundLightType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_AeronauticalGroundLightType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_ElevatedCurvePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ElevatedCurvePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ElevatedCurvePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ElevatedCurvePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ElevatedCurvePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_ElevatedSurfacePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ElevatedSurfacePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ElevatedSurfacePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ElevatedSurfacePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ElevatedSurfacePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_GroundLightSystemPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New GroundLightSystemPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.GroundLightSystemPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_GroundLightSystemPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_GroundLightSystemPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_LightElementPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New LightElementPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.LightElementPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_LightElementPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_LightElementPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_LightElementStatusPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New LightElementStatusPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.LightElementStatusPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_LightElementStatusPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_LightElementStatusPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_LightElementStatusType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New LightElementStatusType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.LightElementStatusType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_LightElementStatusType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_LightElementStatusType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_LightElementType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New LightElementType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.LightElementType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_LightElementType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_LightElementType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_MarkerBeaconPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New MarkerBeaconPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.MarkerBeaconPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_MarkerBeaconPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_MarkerBeaconPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_ObstacleAreaTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ObstacleAreaTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ObstacleAreaTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ObstacleAreaTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ObstacleAreaTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ObstacleAreaTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ObstacleAreaTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ObstacleAreaTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ObstacleAreaTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ObstacleAreaTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ObstacleAreaType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ObstacleAreaType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ObstacleAreaType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ObstacleAreaType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ObstacleAreaType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_PassengerServicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New PassengerServicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.PassengerServicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_PassengerServicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_PassengerServicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_ServicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ServicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.ServicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_ServicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_ServicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SpecialNavigationStationPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SpecialNavigationStationPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.SpecialNavigationStationPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_SpecialNavigationStationPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_SpecialNavigationStationPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_VerticalStructureLightingStatusPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New VerticalStructureLightingStatusPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.VerticalStructureLightingStatusPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_VerticalStructureLightingStatusPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_VerticalStructureLightingStatusPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_VerticalStructureLightingStatusType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New VerticalStructureLightingStatusType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.VerticalStructureLightingStatusType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_VerticalStructureLightingStatusType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_VerticalStructureLightingStatusType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_VerticalStructurePartPropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New VerticalStructurePartPropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.VerticalStructurePartPropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_VerticalStructurePartPropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_VerticalStructurePartPropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_VerticalStructurePartType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New VerticalStructurePartType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.VerticalStructurePartType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_VerticalStructurePartType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_VerticalStructurePartType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_VerticalStructureTimeSlicePropertyType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New VerticalStructureTimeSlicePropertyType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.VerticalStructureTimeSlicePropertyType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_VerticalStructureTimeSlicePropertyType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_VerticalStructureTimeSlicePropertyType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_VerticalStructureTimeSliceType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New VerticalStructureTimeSliceType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.VerticalStructureTimeSliceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_VerticalStructureTimeSliceType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_VerticalStructureTimeSliceType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_VerticalStructureType_AixmfeaturesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New VerticalStructureType",
	}).Stage(probe.formStage)
	inst := new(aixmfeatures.VerticalStructureType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmfeaturesStage.Lock()
			defer probe.stageOfInterest.AixmfeaturesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmfeaturesStage)
			saveStageSet_VerticalStructureType_AixmfeaturesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmfeaturesStage.Commit()
			updateStageSetTable_VerticalStructureType_AixmfeaturesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeColourType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeColourType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeColourType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeColourType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeColourType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_CodeGroundLightingType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeGroundLightingType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeGroundLightingType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeGroundLightingType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeGroundLightingType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeLightIntensityType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeLightIntensityType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeLightIntensityType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeLightIntensityType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeLightIntensityType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeLightSourceType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeLightSourceType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeLightSourceType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeLightSourceType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeLightSourceType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_CodeObstacleAreaType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeObstacleAreaType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeObstacleAreaType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeObstacleAreaType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeObstacleAreaType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_CodeStatusConstructionType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeStatusConstructionType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeStatusConstructionType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeStatusConstructionType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeStatusConstructionType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeStatusOperationsType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeStatusOperationsType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeStatusOperationsType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeStatusOperationsType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeStatusOperationsType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_CodeVerticalStructureMarkingType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeVerticalStructureMarkingType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeVerticalStructureMarkingType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeVerticalStructureMarkingType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeVerticalStructureMarkingType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeVerticalStructureMaterialType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeVerticalStructureMaterialType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeVerticalStructureMaterialType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeVerticalStructureMaterialType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeVerticalStructureMaterialType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_CodeVerticalStructureType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New CodeVerticalStructureType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.CodeVerticalStructureType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_CodeVerticalStructureType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_CodeVerticalStructureType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

func StageSetNewInstance_ValLightIntensityType_AixmdatatypesStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ValLightIntensityType",
	}).Stage(probe.formStage)
	inst := new(aixmdatatypes.ValLightIntensityType)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.AixmdatatypesStage.Lock()
			defer probe.stageOfInterest.AixmdatatypesStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.AixmdatatypesStage)
			saveStageSet_ValLightIntensityType_AixmdatatypesStage(inst, probe, formGroup)
			probe.stageOfInterest.AixmdatatypesStage.Commit()
			updateStageSetTable_ValLightIntensityType_AixmdatatypesStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
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

