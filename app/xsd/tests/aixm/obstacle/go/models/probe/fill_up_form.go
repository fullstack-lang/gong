// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/xlink"
)

const FormName = "Form"

var (
	_ = sort.Strings
	_ = strings.Join
)

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {
	switch inst := instance.(type) {
	case *models.AIXMBasicMessage:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("BoundedBy", inst.BoundedBy, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), probe.formStage)
		StageSetBasicFieldtoForm("SequenceNumber", inst.SequenceNumber, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("MessageMetadata", inst.MessageMetadata, formGroup, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.MessageMetadataPropertyType](), probe.formStage)
		StageSetBasicFieldtoForm("AggregationType", inst.AggregationType, probe.formStage, formGroup)
		AssociationSliceToForm("HasMember", inst, &inst.HasMember, formGroup, probe)
	case *models.AIXMBasicMessageType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("BoundedBy", inst.BoundedBy, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), probe.formStage)
		StageSetBasicFieldtoForm("SequenceNumber", inst.SequenceNumber, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("MessageMetadata", inst.MessageMetadata, formGroup, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.MessageMetadataPropertyType](), probe.formStage)
		StageSetBasicFieldtoForm("AggregationType", inst.AggregationType, probe.formStage, formGroup)
		AssociationSliceToForm("HasMember", inst, &inst.HasMember, formGroup, probe)
	case *models.BasicMessageMemberAIXMPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Href", inst.Href, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Arcrole", inst.Arcrole, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Show", inst.Show, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Actuate", inst.Actuate, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("VerticalStructure", inst.VerticalStructure, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.VerticalStructureType](), probe.formStage)
		StageSetAssociationFieldToForm("ObstacleArea", inst.ObstacleArea, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ObstacleAreaType](), probe.formStage)
		StageSetAssociationFieldToForm("AeronauticalGroundLight", inst.AeronauticalGroundLight, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AeronauticalGroundLightType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AIXMBasicMessages {
				for _, target := range src.HasMember {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AIXMBasicMessage", "HasMember", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AIXMBasicMessageTypes {
				for _, target := range src.HasMember {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AIXMBasicMessageType", "HasMember", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractElevatedCurveExtension
			div := (&form.FormDiv{Name: "AbstractElevatedCurveExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractElevatedCurveExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractElevatedCurveExtension",
				Label: "AbstractElevatedCurveExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedCurveTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedCurveType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_1:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractElevatedPointExtension
			div := (&form.FormDiv{Name: "AbstractElevatedPointExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractElevatedPointExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractElevatedPointExtension",
				Label: "AbstractElevatedPointExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedPointTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedPointType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_10:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractLightElementExtension
			div := (&form.FormDiv{Name: "AbstractLightElementExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractLightElementExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractLightElementExtension",
				Label: "AbstractLightElementExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LightElementType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_11:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractPropertiesWithScheduleExtension
			div := (&form.FormDiv{Name: "AbstractPropertiesWithScheduleExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractPropertiesWithScheduleExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractPropertiesWithScheduleExtension",
				Label: "AbstractPropertiesWithScheduleExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractLightElementStatusExtension
			div := (&form.FormDiv{Name: "AbstractLightElementStatusExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractLightElementStatusExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractLightElementStatusExtension",
				Label: "AbstractLightElementStatusExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementStatusTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LightElementStatusType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_12:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractTimesheetExtension
			div := (&form.FormDiv{Name: "AbstractTimesheetExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractTimesheetExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractTimesheetExtension",
				Label: "AbstractTimesheetExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TimesheetType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_2:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractElevatedSurfaceExtension
			div := (&form.FormDiv{Name: "AbstractElevatedSurfaceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractElevatedSurfaceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractElevatedSurfaceExtension",
				Label: "AbstractElevatedSurfaceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedSurfaceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedSurfaceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_3:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractAeronauticalGroundLightExtension
			div := (&form.FormDiv{Name: "AbstractAeronauticalGroundLightExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAeronauticalGroundLightExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAeronauticalGroundLightExtension",
				Label: "AbstractAeronauticalGroundLightExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AeronauticalGroundLightTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_4:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractLinguisticNoteExtension
			div := (&form.FormDiv{Name: "AbstractLinguisticNoteExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractLinguisticNoteExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractLinguisticNoteExtension",
				Label: "AbstractLinguisticNoteExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LinguisticNoteTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LinguisticNoteType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_5:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractNoteExtension
			div := (&form.FormDiv{Name: "AbstractNoteExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractNoteExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractNoteExtension",
				Label: "AbstractNoteExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NoteTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NoteType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_6:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractObstacleAreaExtension
			div := (&form.FormDiv{Name: "AbstractObstacleAreaExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractObstacleAreaExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractObstacleAreaExtension",
				Label: "AbstractObstacleAreaExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAreaTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_7:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractVerticalStructureExtension
			div := (&form.FormDiv{Name: "AbstractVerticalStructureExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractVerticalStructureExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractVerticalStructureExtension",
				Label: "AbstractVerticalStructureExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_8:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractPropertiesWithScheduleExtension
			div := (&form.FormDiv{Name: "AbstractPropertiesWithScheduleExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractPropertiesWithScheduleExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractPropertiesWithScheduleExtension",
				Label: "AbstractPropertiesWithScheduleExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractVerticalStructureLightingStatusExtension
			div := (&form.FormDiv{Name: "AbstractVerticalStructureLightingStatusExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractVerticalStructureLightingStatusExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractVerticalStructureLightingStatusExtension",
				Label: "AbstractVerticalStructureLightingStatusExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureLightingStatusTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureLightingStatusType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_9:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractPropertiesWithScheduleExtension
			div := (&form.FormDiv{Name: "AbstractPropertiesWithScheduleExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractPropertiesWithScheduleExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractPropertiesWithScheduleExtension",
				Label: "AbstractPropertiesWithScheduleExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractVerticalStructurePartExtension
			div := (&form.FormDiv{Name: "AbstractVerticalStructurePartExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractVerticalStructurePartExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractVerticalStructurePartExtension",
				Label: "AbstractVerticalStructurePartExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructurePartType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AbstractPropertiesWithScheduleType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
	case *aixmfeatures.AeronauticalGroundLightTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AeronauticalGroundLightTimeSlice", inst.AeronauticalGroundLightTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AeronauticalGroundLightTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AeronauticalGroundLightType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AeronauticalGroundLightTimeSliceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("ValidTime", inst.ValidTime, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), probe.formStage)
		StageSetBasicFieldtoForm("Interpretation", inst.Interpretation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SequenceNumber", inst.SequenceNumber, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CorrectionNumber", inst.CorrectionNumber, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("TimeSliceMetadata", inst.TimeSliceMetadata, formGroup, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("FeatureLifetime", inst.FeatureLifetime, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeGroundLightingType](), probe.formStage)
		StageSetAssociationFieldToForm("Colour", inst.Colour, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeColourType](), probe.formStage)
		StageSetAssociationFieldToForm("Flashing", inst.Flashing, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("StructureBeacon", inst.StructureBeacon, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.VerticalStructurePropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("AerodromeBeacon", inst.AerodromeBeacon, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Location", inst.Location, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)

		{
			// Slice of pointers: Annotation
			div := (&form.FormDiv{Name: "Annotation"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Annotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Annotation",
				Label: "Annotation",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Extension
			div := (&form.FormDiv{Name: "Extension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Extension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Extension",
				Label: "Extension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTimeSlicePropertyTypes {
				if src.AeronauticalGroundLightTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AeronauticalGroundLightTimeSlicePropertyType", "AeronauticalGroundLightTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AeronauticalGroundLightType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("BoundedBy", inst.BoundedBy, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), probe.formStage)
		StageSetAssociationFieldToForm("FeatureMetadata", inst.FeatureMetadata, formGroup, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), probe.formStage)

		{
			// Slice of pointers: TimeSlice
			div := (&form.FormDiv{Name: "TimeSlice"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TimeSlice {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TimeSlice",
				Label: "TimeSlice",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.BasicMessageMemberAIXMPropertyTypes {
				if src.AeronauticalGroundLight == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "AeronauticalGroundLight", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirportHeliportPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Href", inst.Href, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Arcrole", inst.Arcrole, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Show", inst.Show, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Actuate", inst.Actuate, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTimeSliceTypes {
				if src.AerodromeBeacon == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AeronauticalGroundLightTimeSliceType", "AerodromeBeacon", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSliceTypes {
				if src.Reference_ownerAirport == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAreaTimeSliceType", "Reference_ownerAirport", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.CurveType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("SrsName", inst.SrsName, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Segments", inst.Segments, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CurveSegmentArrayPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("HorizontalAccuracy", inst.HorizontalAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)

		{
			// Slice of pointers: Annotation
			div := (&form.FormDiv{Name: "Annotation"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Annotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Annotation",
				Label: "Annotation",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
	case *aixmfeatures.ElevatedCurvePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ElevatedCurve", inst.ElevatedCurve, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedCurveType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.HorizontalProjection_linearExtent == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructurePartType", "HorizontalProjection_linearExtent", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ElevatedCurveType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("SrsName", inst.SrsName, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Segments", inst.Segments, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CurveSegmentArrayPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("HorizontalAccuracy", inst.HorizontalAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)

		{
			// Slice of pointers: Annotation
			div := (&form.FormDiv{Name: "Annotation"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Annotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Annotation",
				Label: "Annotation",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Elevation", inst.Elevation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("GeoidUndulation", inst.GeoidUndulation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceSignedType](), probe.formStage)
		StageSetAssociationFieldToForm("VerticalDatum", inst.VerticalDatum, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalDatumType](), probe.formStage)
		StageSetAssociationFieldToForm("VerticalAccuracy", inst.VerticalAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)

		{
			// Slice of pointers: Extension
			div := (&form.FormDiv{Name: "Extension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Extension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Extension",
				Label: "Extension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedCurvePropertyTypes {
				if src.ElevatedCurve == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedCurvePropertyType", "ElevatedCurve", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ElevatedPointPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ElevatedPoint", inst.ElevatedPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTimeSliceTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AeronauticalGroundLightTimeSliceType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LightElementType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.HorizontalProjection_location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructurePartType", "HorizontalProjection_location", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ElevatedPointType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("SrsName", inst.SrsName, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Pos", inst.Pos, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionType](), probe.formStage)
		StageSetAssociationFieldToForm("HorizontalAccuracy", inst.HorizontalAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)

		{
			// Slice of pointers: Annotation
			div := (&form.FormDiv{Name: "Annotation"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Annotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Annotation",
				Label: "Annotation",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Elevation", inst.Elevation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("GeoidUndulation", inst.GeoidUndulation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceSignedType](), probe.formStage)
		StageSetAssociationFieldToForm("VerticalDatum", inst.VerticalDatum, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalDatumType](), probe.formStage)
		StageSetAssociationFieldToForm("VerticalAccuracy", inst.VerticalAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)

		{
			// Slice of pointers: Extension
			div := (&form.FormDiv{Name: "Extension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Extension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Extension",
				Label: "Extension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedPointPropertyTypes {
				if src.ElevatedPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedPointPropertyType", "ElevatedPoint", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ElevatedSurfacePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ElevatedSurface", inst.ElevatedSurface, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedSurfaceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.HorizontalProjection_surfaceExtent == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructurePartType", "HorizontalProjection_surfaceExtent", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ElevatedSurfaceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("SrsName", inst.SrsName, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Patches", inst.Patches, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.SurfacePatchArrayPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("HorizontalAccuracy", inst.HorizontalAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)

		{
			// Slice of pointers: Annotation
			div := (&form.FormDiv{Name: "Annotation"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Annotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Annotation",
				Label: "Annotation",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Elevation", inst.Elevation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("GeoidUndulation", inst.GeoidUndulation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceSignedType](), probe.formStage)
		StageSetAssociationFieldToForm("VerticalDatum", inst.VerticalDatum, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalDatumType](), probe.formStage)
		StageSetAssociationFieldToForm("VerticalAccuracy", inst.VerticalAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)

		{
			// Slice of pointers: Extension
			div := (&form.FormDiv{Name: "Extension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Extension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Extension",
				Label: "Extension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedSurfacePropertyTypes {
				if src.ElevatedSurface == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedSurfacePropertyType", "ElevatedSurface", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SurfacePropertyTypes {
				if src.ElevatedSurface == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SurfacePropertyType", "ElevatedSurface", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.GroundLightSystemPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Href", inst.Href, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Arcrole", inst.Arcrole, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Show", inst.Show, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Actuate", inst.Actuate, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				for _, target := range src.SupportedGroundLight {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureTimeSliceType", "SupportedGroundLight", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.LightElementPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("LightElement", inst.LightElement, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.LightElementType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				for _, target := range src.Lighting {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructurePartType", "Lighting", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.LightElementStatusPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("LightElementStatus", inst.LightElementStatus, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.LightElementStatusType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementTypes {
				for _, target := range src.Availability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LightElementType", "Availability", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.LightElementStatusType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: TimeInterval
			div := (&form.FormDiv{Name: "TimeInterval"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TimeInterval {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TimeInterval",
				Label: "TimeInterval",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Annotation
			div := (&form.FormDiv{Name: "Annotation"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Annotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Annotation",
				Label: "Annotation",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: SpecialDateAuthority
			div := (&form.FormDiv{Name: "SpecialDateAuthority"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SpecialDateAuthority {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SpecialDateAuthority",
				Label: "SpecialDateAuthority",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Status", inst.Status, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeStatusOperationsType](), probe.formStage)

		{
			// Slice of pointers: Extension
			div := (&form.FormDiv{Name: "Extension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Extension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Extension",
				Label: "Extension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementStatusPropertyTypes {
				if src.LightElementStatus == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LightElementStatusPropertyType", "LightElementStatus", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.LightElementType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Colour", inst.Colour, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeColourType](), probe.formStage)
		StageSetAssociationFieldToForm("IntensityLevel", inst.IntensityLevel, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeLightIntensityType](), probe.formStage)
		StageSetAssociationFieldToForm("Intensity", inst.Intensity, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValLightIntensityType](), probe.formStage)
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeLightSourceType](), probe.formStage)
		StageSetAssociationFieldToForm("Location", inst.Location, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)

		{
			// Slice of pointers: Annotation
			div := (&form.FormDiv{Name: "Annotation"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Annotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Annotation",
				Label: "Annotation",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Availability
			div := (&form.FormDiv{Name: "Availability"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Availability {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Availability",
				Label: "Availability",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Extension
			div := (&form.FormDiv{Name: "Extension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Extension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Extension",
				Label: "Extension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementPropertyTypes {
				if src.LightElement == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LightElementPropertyType", "LightElement", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.LinguisticNotePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("LinguisticNote", inst.LinguisticNote, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.LinguisticNoteType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NoteTypes {
				for _, target := range src.TranslatedNote {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NoteType", "TranslatedNote", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.LinguisticNoteType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Note", inst.Note, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextNoteType](), probe.formStage)

		{
			// Slice of pointers: Extension
			div := (&form.FormDiv{Name: "Extension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Extension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Extension",
				Label: "Extension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LinguisticNotePropertyTypes {
				if src.LinguisticNote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LinguisticNotePropertyType", "LinguisticNote", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.MarkerBeaconPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Href", inst.Href, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Arcrole", inst.Arcrole, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Show", inst.Show, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Actuate", inst.Actuate, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				if src.Marker == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureTimeSliceType", "Marker", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.NavaidEquipmentPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Href", inst.Href, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Arcrole", inst.Arcrole, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Show", inst.Show, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Actuate", inst.Actuate, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				for _, target := range src.HostedNavaidEquipment {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureTimeSliceType", "HostedNavaidEquipment", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.NotePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Note", inst.Note, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NoteType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AeronauticalGroundLightTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CurveTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CurveType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedCurveTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedCurveType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedPointTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedPointType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedSurfaceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedSurfaceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementStatusTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LightElementStatusType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LightElementType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAreaTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PointTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PointType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SurfaceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SurfaceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TimesheetType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureLightingStatusTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureLightingStatusType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructurePartType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.NoteType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("PropertyName", inst.PropertyName, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextPropertyNameType](), probe.formStage)
		StageSetAssociationFieldToForm("Purpose", inst.Purpose, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNotePurposeType](), probe.formStage)

		{
			// Slice of pointers: TranslatedNote
			div := (&form.FormDiv{Name: "TranslatedNote"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TranslatedNote {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TranslatedNote",
				Label: "TranslatedNote",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Extension
			div := (&form.FormDiv{Name: "Extension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Extension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Extension",
				Label: "Extension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NotePropertyTypes {
				if src.Note == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NotePropertyType", "Note", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ObstacleAreaTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ObstacleAreaTimeSlice", inst.ObstacleAreaTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ObstacleAreaTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAreaType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ObstacleAreaTimeSliceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("ValidTime", inst.ValidTime, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), probe.formStage)
		StageSetBasicFieldtoForm("Interpretation", inst.Interpretation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SequenceNumber", inst.SequenceNumber, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CorrectionNumber", inst.CorrectionNumber, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("TimeSliceMetadata", inst.TimeSliceMetadata, formGroup, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("FeatureLifetime", inst.FeatureLifetime, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeObstacleAreaType](), probe.formStage)
		StageSetAssociationFieldToForm("ObstructionIdSurfaceCondition", inst.ObstructionIdSurfaceCondition, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeObstacleAssessmentSurfaceType](), probe.formStage)
		StageSetAssociationFieldToForm("Reference_ownerOrganisation", inst.Reference_ownerOrganisation, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.OrganisationAuthorityPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Reference_ownerRunway", inst.Reference_ownerRunway, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayDirectionPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Reference_ownerAirport", inst.Reference_ownerAirport, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("SurfaceExtent", inst.SurfaceExtent, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SurfacePropertyType](), probe.formStage)

		{
			// Slice of pointers: Obstacle
			div := (&form.FormDiv{Name: "Obstacle"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Obstacle {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Obstacle",
				Label: "Obstacle",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Annotation
			div := (&form.FormDiv{Name: "Annotation"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Annotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Annotation",
				Label: "Annotation",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Extension
			div := (&form.FormDiv{Name: "Extension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Extension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Extension",
				Label: "Extension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSlicePropertyTypes {
				if src.ObstacleAreaTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAreaTimeSlicePropertyType", "ObstacleAreaTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ObstacleAreaType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("BoundedBy", inst.BoundedBy, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), probe.formStage)
		StageSetAssociationFieldToForm("FeatureMetadata", inst.FeatureMetadata, formGroup, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), probe.formStage)

		{
			// Slice of pointers: TimeSlice
			div := (&form.FormDiv{Name: "TimeSlice"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TimeSlice {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TimeSlice",
				Label: "TimeSlice",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.BasicMessageMemberAIXMPropertyTypes {
				if src.ObstacleArea == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "ObstacleArea", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.OrganisationAuthorityPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Href", inst.Href, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Arcrole", inst.Arcrole, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Show", inst.Show, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Actuate", inst.Actuate, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementStatusTypes {
				for _, target := range src.SpecialDateAuthority {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LightElementStatusType", "SpecialDateAuthority", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSliceTypes {
				if src.Reference_ownerOrganisation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAreaTimeSliceType", "Reference_ownerOrganisation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureLightingStatusTypes {
				for _, target := range src.SpecialDateAuthority {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureLightingStatusType", "SpecialDateAuthority", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				for _, target := range src.SpecialDateAuthority {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructurePartType", "SpecialDateAuthority", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				for _, target := range src.HostedOrganisation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureTimeSliceType", "HostedOrganisation", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.PassengerServicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Href", inst.Href, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Arcrole", inst.Arcrole, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Show", inst.Show, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Actuate", inst.Actuate, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				for _, target := range src.HostedPassengerService {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureTimeSliceType", "HostedPassengerService", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.PointType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("SrsName", inst.SrsName, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Pos", inst.Pos, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionType](), probe.formStage)
		StageSetAssociationFieldToForm("HorizontalAccuracy", inst.HorizontalAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)

		{
			// Slice of pointers: Annotation
			div := (&form.FormDiv{Name: "Annotation"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Annotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Annotation",
				Label: "Annotation",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
	case *aixmfeatures.RunwayDirectionPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Href", inst.Href, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Arcrole", inst.Arcrole, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Show", inst.Show, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Actuate", inst.Actuate, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSliceTypes {
				if src.Reference_ownerRunway == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAreaTimeSliceType", "Reference_ownerRunway", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ServicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Href", inst.Href, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Arcrole", inst.Arcrole, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Show", inst.Show, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Actuate", inst.Actuate, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				for _, target := range src.SupportedService {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureTimeSliceType", "SupportedService", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SpecialNavigationStationPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Href", inst.Href, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Arcrole", inst.Arcrole, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Show", inst.Show, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Actuate", inst.Actuate, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				for _, target := range src.HostedSpecialNavStation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureTimeSliceType", "HostedSpecialNavStation", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SurfacePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Surface", inst.Surface, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SurfaceType](), probe.formStage)
		StageSetAssociationFieldToForm("ElevatedSurface", inst.ElevatedSurface, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedSurfaceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSliceTypes {
				if src.SurfaceExtent == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAreaTimeSliceType", "SurfaceExtent", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SurfaceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("SrsName", inst.SrsName, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Patches", inst.Patches, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.SurfacePatchArrayPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("HorizontalAccuracy", inst.HorizontalAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)

		{
			// Slice of pointers: Annotation
			div := (&form.FormDiv{Name: "Annotation"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Annotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Annotation",
				Label: "Annotation",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SurfacePropertyTypes {
				if src.Surface == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SurfacePropertyType", "Surface", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.TimesheetPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Timesheet", inst.Timesheet, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TimesheetType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementStatusTypes {
				for _, target := range src.TimeInterval {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LightElementStatusType", "TimeInterval", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureLightingStatusTypes {
				for _, target := range src.TimeInterval {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureLightingStatusType", "TimeInterval", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				for _, target := range src.TimeInterval {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructurePartType", "TimeInterval", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.TimesheetType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("TimeReference", inst.TimeReference, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTimeReferenceType](), probe.formStage)
		StageSetAssociationFieldToForm("StartDate", inst.StartDate, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateMonthDayType](), probe.formStage)
		StageSetAssociationFieldToForm("EndDate", inst.EndDate, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateMonthDayType](), probe.formStage)
		StageSetAssociationFieldToForm("Day", inst.Day, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDayType](), probe.formStage)
		StageSetAssociationFieldToForm("DayTil", inst.DayTil, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDayType](), probe.formStage)
		StageSetAssociationFieldToForm("StartTime", inst.StartTime, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TimeType](), probe.formStage)
		StageSetAssociationFieldToForm("StartEvent", inst.StartEvent, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTimeEventType](), probe.formStage)
		StageSetAssociationFieldToForm("StartTimeRelativeEvent", inst.StartTimeRelativeEvent, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDurationType](), probe.formStage)
		StageSetAssociationFieldToForm("StartEventInterpretation", inst.StartEventInterpretation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTimeEventCombinationType](), probe.formStage)
		StageSetAssociationFieldToForm("EndTime", inst.EndTime, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TimeType](), probe.formStage)
		StageSetAssociationFieldToForm("EndEvent", inst.EndEvent, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTimeEventType](), probe.formStage)
		StageSetAssociationFieldToForm("EndTimeRelativeEvent", inst.EndTimeRelativeEvent, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDurationType](), probe.formStage)
		StageSetAssociationFieldToForm("EndEventInterpretation", inst.EndEventInterpretation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTimeEventCombinationType](), probe.formStage)
		StageSetAssociationFieldToForm("DaylightSavingAdjust", inst.DaylightSavingAdjust, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Excluded", inst.Excluded, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)

		{
			// Slice of pointers: Annotation
			div := (&form.FormDiv{Name: "Annotation"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Annotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Annotation",
				Label: "Annotation",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Extension
			div := (&form.FormDiv{Name: "Extension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Extension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Extension",
				Label: "Extension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetPropertyTypes {
				if src.Timesheet == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TimesheetPropertyType", "Timesheet", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.UnitPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Href", inst.Href, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Arcrole", inst.Arcrole, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Show", inst.Show, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Actuate", inst.Actuate, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				for _, target := range src.HostedUnit {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureTimeSliceType", "HostedUnit", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.VerticalStructureLightingStatusPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("VerticalStructureLightingStatus", inst.VerticalStructureLightingStatus, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.VerticalStructureLightingStatusType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				for _, target := range src.LightingAvailability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureTimeSliceType", "LightingAvailability", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.VerticalStructureLightingStatusType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: TimeInterval
			div := (&form.FormDiv{Name: "TimeInterval"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TimeInterval {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TimeInterval",
				Label: "TimeInterval",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Annotation
			div := (&form.FormDiv{Name: "Annotation"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Annotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Annotation",
				Label: "Annotation",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: SpecialDateAuthority
			div := (&form.FormDiv{Name: "SpecialDateAuthority"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SpecialDateAuthority {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SpecialDateAuthority",
				Label: "SpecialDateAuthority",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Status", inst.Status, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeStatusOperationsType](), probe.formStage)

		{
			// Slice of pointers: Extension
			div := (&form.FormDiv{Name: "Extension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Extension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Extension",
				Label: "Extension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureLightingStatusPropertyTypes {
				if src.VerticalStructureLightingStatus == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureLightingStatusPropertyType", "VerticalStructureLightingStatus", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.VerticalStructurePartPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("VerticalStructurePart", inst.VerticalStructurePart, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.VerticalStructurePartType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				for _, target := range src.Part {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureTimeSliceType", "Part", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.VerticalStructurePartType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: TimeInterval
			div := (&form.FormDiv{Name: "TimeInterval"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TimeInterval {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TimeInterval",
				Label: "TimeInterval",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Annotation
			div := (&form.FormDiv{Name: "Annotation"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Annotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Annotation",
				Label: "Annotation",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: SpecialDateAuthority
			div := (&form.FormDiv{Name: "SpecialDateAuthority"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SpecialDateAuthority {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SpecialDateAuthority",
				Label: "SpecialDateAuthority",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("VerticalExtent", inst.VerticalExtent, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)
		StageSetAssociationFieldToForm("VerticalExtentAccuracy", inst.VerticalExtentAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalStructureType](), probe.formStage)
		StageSetAssociationFieldToForm("ConstructionStatus", inst.ConstructionStatus, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeStatusConstructionType](), probe.formStage)
		StageSetAssociationFieldToForm("MarkingPattern", inst.MarkingPattern, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalStructureMarkingType](), probe.formStage)
		StageSetAssociationFieldToForm("MarkingFirstColour", inst.MarkingFirstColour, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeColourType](), probe.formStage)
		StageSetAssociationFieldToForm("MarkingSecondColour", inst.MarkingSecondColour, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeColourType](), probe.formStage)
		StageSetAssociationFieldToForm("Mobile", inst.Mobile, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Frangible", inst.Frangible, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("VisibleMaterial", inst.VisibleMaterial, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalStructureMaterialType](), probe.formStage)
		StageSetAssociationFieldToForm("Designator", inst.Designator, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextDesignatorType](), probe.formStage)
		StageSetAssociationFieldToForm("HorizontalProjection_surfaceExtent", inst.HorizontalProjection_surfaceExtent, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedSurfacePropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("HorizontalProjection_linearExtent", inst.HorizontalProjection_linearExtent, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedCurvePropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("HorizontalProjection_location", inst.HorizontalProjection_location, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)

		{
			// Slice of pointers: Lighting
			div := (&form.FormDiv{Name: "Lighting"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Lighting {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Lighting",
				Label: "Lighting",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Extension
			div := (&form.FormDiv{Name: "Extension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Extension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Extension",
				Label: "Extension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartPropertyTypes {
				if src.VerticalStructurePart == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructurePartPropertyType", "VerticalStructurePart", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.VerticalStructurePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Href", inst.Href, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Arcrole", inst.Arcrole, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Show", inst.Show, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Actuate", inst.Actuate, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTimeSliceTypes {
				if src.StructureBeacon == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AeronauticalGroundLightTimeSliceType", "StructureBeacon", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSliceTypes {
				for _, target := range src.Obstacle {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAreaTimeSliceType", "Obstacle", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.VerticalStructureTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("VerticalStructureTimeSlice", inst.VerticalStructureTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.VerticalStructureTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.VerticalStructureTimeSliceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("ValidTime", inst.ValidTime, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), probe.formStage)
		StageSetBasicFieldtoForm("Interpretation", inst.Interpretation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SequenceNumber", inst.SequenceNumber, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CorrectionNumber", inst.CorrectionNumber, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("TimeSliceMetadata", inst.TimeSliceMetadata, formGroup, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("FeatureLifetime", inst.FeatureLifetime, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalStructureType](), probe.formStage)
		StageSetAssociationFieldToForm("Lighted", inst.Lighted, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("MarkingICAOStandard", inst.MarkingICAOStandard, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Group", inst.Group, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Length", inst.Length, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)
		StageSetAssociationFieldToForm("Width", inst.Width, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)
		StageSetAssociationFieldToForm("Radius", inst.Radius, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)
		StageSetAssociationFieldToForm("LightingICAOStandard", inst.LightingICAOStandard, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("SynchronisedLighting", inst.SynchronisedLighting, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Marker", inst.Marker, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.MarkerBeaconPropertyType](), probe.formStage)

		{
			// Slice of pointers: Part
			div := (&form.FormDiv{Name: "Part"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Part {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Part",
				Label: "Part",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: HostedPassengerService
			div := (&form.FormDiv{Name: "HostedPassengerService"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.HostedPassengerService {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "HostedPassengerService",
				Label: "HostedPassengerService",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: SupportedGroundLight
			div := (&form.FormDiv{Name: "SupportedGroundLight"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SupportedGroundLight {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SupportedGroundLight",
				Label: "SupportedGroundLight",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: HostedNavaidEquipment
			div := (&form.FormDiv{Name: "HostedNavaidEquipment"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.HostedNavaidEquipment {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "HostedNavaidEquipment",
				Label: "HostedNavaidEquipment",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: HostedSpecialNavStation
			div := (&form.FormDiv{Name: "HostedSpecialNavStation"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.HostedSpecialNavStation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "HostedSpecialNavStation",
				Label: "HostedSpecialNavStation",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: HostedUnit
			div := (&form.FormDiv{Name: "HostedUnit"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.HostedUnit {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "HostedUnit",
				Label: "HostedUnit",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: HostedOrganisation
			div := (&form.FormDiv{Name: "HostedOrganisation"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.HostedOrganisation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "HostedOrganisation",
				Label: "HostedOrganisation",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: SupportedService
			div := (&form.FormDiv{Name: "SupportedService"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SupportedService {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SupportedService",
				Label: "SupportedService",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Annotation
			div := (&form.FormDiv{Name: "Annotation"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Annotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Annotation",
				Label: "Annotation",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: LightingAvailability
			div := (&form.FormDiv{Name: "LightingAvailability"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.LightingAvailability {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "LightingAvailability",
				Label: "LightingAvailability",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Extension
			div := (&form.FormDiv{Name: "Extension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Extension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Extension",
				Label: "Extension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSlicePropertyTypes {
				if src.VerticalStructureTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureTimeSlicePropertyType", "VerticalStructureTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.VerticalStructureType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("BoundedBy", inst.BoundedBy, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), probe.formStage)
		StageSetAssociationFieldToForm("FeatureMetadata", inst.FeatureMetadata, formGroup, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), probe.formStage)

		{
			// Slice of pointers: TimeSlice
			div := (&form.FormDiv{Name: "TimeSlice"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TimeSlice {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TimeSlice",
				Label: "TimeSlice",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.BasicMessageMemberAIXMPropertyTypes {
				if src.VerticalStructure == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "VerticalStructure", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeColourType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTimeSliceTypes {
				if src.Colour == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AeronauticalGroundLightTimeSliceType", "Colour", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementTypes {
				if src.Colour == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LightElementType", "Colour", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.MarkingFirstColour == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructurePartType", "MarkingFirstColour", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.MarkingSecondColour == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructurePartType", "MarkingSecondColour", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeDayType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.Day == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TimesheetType", "Day", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.DayTil == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TimesheetType", "DayTil", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeGroundLightingType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AeronauticalGroundLightTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeLightIntensityType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementTypes {
				if src.IntensityLevel == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LightElementType", "IntensityLevel", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeLightSourceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LightElementType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeNotePurposeType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NoteTypes {
				if src.Purpose == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NoteType", "Purpose", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeObstacleAreaType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAreaTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeObstacleAssessmentSurfaceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSliceTypes {
				if src.ObstructionIdSurfaceCondition == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAreaTimeSliceType", "ObstructionIdSurfaceCondition", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeStatusConstructionType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.ConstructionStatus == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructurePartType", "ConstructionStatus", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeStatusOperationsType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementStatusTypes {
				if src.Status == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LightElementStatusType", "Status", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureLightingStatusTypes {
				if src.Status == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureLightingStatusType", "Status", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeTimeEventCombinationType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.StartEventInterpretation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TimesheetType", "StartEventInterpretation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.EndEventInterpretation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TimesheetType", "EndEventInterpretation", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeTimeEventType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.StartEvent == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TimesheetType", "StartEvent", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.EndEvent == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TimesheetType", "EndEvent", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeTimeReferenceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.TimeReference == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TimesheetType", "TimeReference", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeVerticalDatumType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedCurveTypes {
				if src.VerticalDatum == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedCurveType", "VerticalDatum", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedPointTypes {
				if src.VerticalDatum == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedPointType", "VerticalDatum", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedSurfaceTypes {
				if src.VerticalDatum == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedSurfaceType", "VerticalDatum", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeVerticalStructureMarkingType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.MarkingPattern == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructurePartType", "MarkingPattern", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeVerticalStructureMaterialType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.VisibleMaterial == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructurePartType", "VisibleMaterial", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeVerticalStructureType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructurePartType", "Type_XSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeYesNoType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTimeSliceTypes {
				if src.Flashing == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AeronauticalGroundLightTimeSliceType", "Flashing", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.DaylightSavingAdjust == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TimesheetType", "DaylightSavingAdjust", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.Excluded == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TimesheetType", "Excluded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.Mobile == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructurePartType", "Mobile", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.Frangible == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructurePartType", "Frangible", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				if src.Lighted == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureTimeSliceType", "Lighted", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				if src.MarkingICAOStandard == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureTimeSliceType", "MarkingICAOStandard", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				if src.Group == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureTimeSliceType", "Group", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				if src.LightingICAOStandard == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureTimeSliceType", "LightingICAOStandard", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				if src.SynchronisedLighting == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureTimeSliceType", "SynchronisedLighting", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.DateMonthDayType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.StartDate == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TimesheetType", "StartDate", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.EndDate == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TimesheetType", "EndDate", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.TextDesignatorType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.Designator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructurePartType", "Designator", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.TextNameType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)
	case *aixmdatatypes.TextNoteType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Lang", inst.Lang, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LinguisticNoteTypes {
				if src.Note == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LinguisticNoteType", "Note", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.TextPropertyNameType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NoteTypes {
				if src.PropertyName == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NoteType", "PropertyName", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.TimeType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.StartTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TimesheetType", "StartTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.EndTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TimesheetType", "EndTime", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.ValDistanceSignedType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Uom", inst.Uom, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedCurveTypes {
				if src.GeoidUndulation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedCurveType", "GeoidUndulation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedPointTypes {
				if src.GeoidUndulation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedPointType", "GeoidUndulation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedSurfaceTypes {
				if src.GeoidUndulation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedSurfaceType", "GeoidUndulation", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.ValDistanceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Uom", inst.Uom, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CurveTypes {
				if src.HorizontalAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CurveType", "HorizontalAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedCurveTypes {
				if src.HorizontalAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedCurveType", "HorizontalAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedCurveTypes {
				if src.VerticalAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedCurveType", "VerticalAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedPointTypes {
				if src.HorizontalAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedPointType", "HorizontalAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedPointTypes {
				if src.VerticalAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedPointType", "VerticalAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedSurfaceTypes {
				if src.HorizontalAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedSurfaceType", "HorizontalAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedSurfaceTypes {
				if src.VerticalAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedSurfaceType", "VerticalAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PointTypes {
				if src.HorizontalAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PointType", "HorizontalAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SurfaceTypes {
				if src.HorizontalAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SurfaceType", "HorizontalAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.VerticalExtent == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructurePartType", "VerticalExtent", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.VerticalExtentAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructurePartType", "VerticalExtentAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				if src.Length == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureTimeSliceType", "Length", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				if src.Width == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureTimeSliceType", "Width", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				if src.Radius == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureTimeSliceType", "Radius", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.ValDistanceVerticalType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Uom", inst.Uom, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedCurveTypes {
				if src.Elevation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedCurveType", "Elevation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedPointTypes {
				if src.Elevation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedPointType", "Elevation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedSurfaceTypes {
				if src.Elevation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedSurfaceType", "Elevation", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.ValDurationType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Uom", inst.Uom, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.StartTimeRelativeEvent == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TimesheetType", "StartTimeRelativeEvent", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.EndTimeRelativeEvent == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TimesheetType", "EndTimeRelativeEvent", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.ValLightIntensityType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Uom", inst.Uom, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementTypes {
				if src.Intensity == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LightElementType", "Intensity", refNames, formGroup, probe.formStage)
		}
	case *aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("BoundedBy", inst.BoundedBy, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), probe.formStage)
	case *aixmabstractgmlobjecttypes.AbstractAIXMFeatureType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("BoundedBy", inst.BoundedBy, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), probe.formStage)
		StageSetAssociationFieldToForm("FeatureMetadata", inst.FeatureMetadata, formGroup, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType](), probe.formStage)
	case *aixmabstractgmlobjecttypes.AbstractAIXMMessageBaseType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("BoundedBy", inst.BoundedBy, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), probe.formStage)
		StageSetBasicFieldtoForm("SequenceNumber", inst.SequenceNumber, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("MessageMetadata", inst.MessageMetadata, formGroup, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.MessageMetadataPropertyType](), probe.formStage)
	case *aixmabstractgmlobjecttypes.AbstractAIXMMessageType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("BoundedBy", inst.BoundedBy, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), probe.formStage)
		StageSetBasicFieldtoForm("SequenceNumber", inst.SequenceNumber, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("MessageMetadata", inst.MessageMetadata, formGroup, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.MessageMetadataPropertyType](), probe.formStage)
		StageSetBasicFieldtoForm("AggregationType", inst.AggregationType, probe.formStage, formGroup)
	case *aixmabstractgmlobjecttypes.AbstractAIXMObjectType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
	case *aixmabstractgmlobjecttypes.AbstractAIXMPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
	case *aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceBaseType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ValidTime", inst.ValidTime, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), probe.formStage)
	case *aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("ValidTime", inst.ValidTime, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), probe.formStage)
		StageSetEnumStringFieldToForm("Interpretation", inst.Interpretation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("SequenceNumber", inst.SequenceNumber, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CorrectionNumber", inst.CorrectionNumber, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("TimeSliceMetadata", inst.TimeSliceMetadata, formGroup, probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("FeatureLifetime", inst.FeatureLifetime, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), probe.formStage)
	case *aixmabstractgmlobjecttypes.AbstractExtensionType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extensions {
				for _, target := range src.AbstractElevatedCurveExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension", "AbstractElevatedCurveExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_1s {
				for _, target := range src.AbstractElevatedPointExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_1", "AbstractElevatedPointExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_10s {
				for _, target := range src.AbstractLightElementExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_10", "AbstractLightElementExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_11s {
				for _, target := range src.AbstractPropertiesWithScheduleExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_11", "AbstractPropertiesWithScheduleExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_11s {
				for _, target := range src.AbstractLightElementStatusExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_11", "AbstractLightElementStatusExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_12s {
				for _, target := range src.AbstractTimesheetExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_12", "AbstractTimesheetExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_2s {
				for _, target := range src.AbstractElevatedSurfaceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_2", "AbstractElevatedSurfaceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_3s {
				for _, target := range src.AbstractAeronauticalGroundLightExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_3", "AbstractAeronauticalGroundLightExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_4s {
				for _, target := range src.AbstractLinguisticNoteExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_4", "AbstractLinguisticNoteExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_5s {
				for _, target := range src.AbstractNoteExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_5", "AbstractNoteExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_6s {
				for _, target := range src.AbstractObstacleAreaExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_6", "AbstractObstacleAreaExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_7s {
				for _, target := range src.AbstractVerticalStructureExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_7", "AbstractVerticalStructureExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_8s {
				for _, target := range src.AbstractPropertiesWithScheduleExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_8", "AbstractPropertiesWithScheduleExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_8s {
				for _, target := range src.AbstractVerticalStructureLightingStatusExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_8", "AbstractVerticalStructureLightingStatusExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_9s {
				for _, target := range src.AbstractPropertiesWithScheduleExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_9", "AbstractPropertiesWithScheduleExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_9s {
				for _, target := range src.AbstractVerticalStructurePartExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_9", "AbstractVerticalStructurePartExtension", refNames, formGroup, probe.formStage)
		}
	case *aixmabstractgmlobjecttypes.FeatureMetadataPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AeronauticalGroundLightType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAreaType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMFeatureTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmabstractgmlobjecttypes.AbstractAIXMFeatureType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}
	case *aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AeronauticalGroundLightTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAreaTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}
	case *aixmabstractgmlobjecttypes.MessageMetadataPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AIXMBasicMessages {
				if src.MessageMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AIXMBasicMessage", "MessageMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AIXMBasicMessageTypes {
				if src.MessageMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AIXMBasicMessageType", "MessageMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMMessageBaseTypes {
				if src.MessageMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmabstractgmlobjecttypes.AbstractAIXMMessageBaseType", "MessageMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMMessageTypes {
				if src.MessageMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmabstractgmlobjecttypes.AbstractAIXMMessageType", "MessageMetadata", refNames, formGroup, probe.formStage)
		}
	case *gml.AbstractCurveSegmentType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumDerivativesAtStart", inst.NumDerivativesAtStart, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumDerivativesAtEnd", inst.NumDerivativesAtEnd, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumDerivativeInterior", inst.NumDerivativeInterior, probe.formStage, formGroup)
	case *gml.AbstractCurveType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("SrsName", inst.SrsName, probe.formStage, formGroup)
	case *gml.AbstractFeatureMemberType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
	case *gml.AbstractFeatureType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("BoundedBy", inst.BoundedBy, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), probe.formStage)
	case *gml.AbstractGMLType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
	case *gml.AbstractGeometricPrimitiveType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("SrsName", inst.SrsName, probe.formStage, formGroup)
	case *gml.AbstractGeometryType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("SrsName", inst.SrsName, probe.formStage, formGroup)
	case *gml.AbstractMetadataPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
	case *gml.AbstractRingPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Ring", inst.Ring, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.RingType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.PolygonPatchTypes {
				if src.Exterior == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.PolygonPatchType", "Exterior", refNames, formGroup, probe.formStage)
		}
	case *gml.AbstractRingType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
	case *gml.AbstractSurfacePatchType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
	case *gml.AbstractSurfaceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("SrsName", inst.SrsName, probe.formStage, formGroup)
	case *gml.AbstractTimeGeometricPrimitiveType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("Frame", inst.Frame, probe.formStage, formGroup)
	case *gml.AbstractTimeObjectType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
	case *gml.AbstractTimePrimitiveType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
	case *gml.AbstractTimeSliceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("ValidTime", inst.ValidTime, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType](), probe.formStage)
	case *gml.AngleType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.ArcByCenterPointTypes {
				if src.StartAngle == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.ArcByCenterPointType", "StartAngle", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.ArcByCenterPointTypes {
				if src.EndAngle == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.ArcByCenterPointType", "EndAngle", refNames, formGroup, probe.formStage)
		}
	case *gml.ArcByCenterPointType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumDerivativesAtStart", inst.NumDerivativesAtStart, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumDerivativesAtEnd", inst.NumDerivativesAtEnd, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumDerivativeInterior", inst.NumDerivativeInterior, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Interpolation", inst.Interpolation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("NumArc", inst.NumArc, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Pos", inst.Pos, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionType](), probe.formStage)
		StageSetAssociationFieldToForm("PointProperty", inst.PointProperty, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.PointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PosList", inst.PosList, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionListType](), probe.formStage)
		StageSetAssociationFieldToForm("Radius", inst.Radius, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.LengthType](), probe.formStage)
		StageSetAssociationFieldToForm("StartAngle", inst.StartAngle, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.AngleType](), probe.formStage)
		StageSetAssociationFieldToForm("EndAngle", inst.EndAngle, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.AngleType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CurveSegmentArrayPropertyTypes {
				for _, target := range src.ArcByCenterPoint {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.CurveSegmentArrayPropertyType", "ArcByCenterPoint", refNames, formGroup, probe.formStage)
		}
	case *gml.ArcStringType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumDerivativesAtStart", inst.NumDerivativesAtStart, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumDerivativesAtEnd", inst.NumDerivativesAtEnd, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumDerivativeInterior", inst.NumDerivativeInterior, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Interpolation", inst.Interpolation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("NumArc", inst.NumArc, probe.formStage, formGroup)

		{
			// Slice of pointers: Pos
			div := (&form.FormDiv{Name: "Pos"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Pos {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Pos",
				Label: "Pos",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: PointProperty
			div := (&form.FormDiv{Name: "PointProperty"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.PointProperty {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "PointProperty",
				Label: "PointProperty",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("PosList", inst.PosList, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionListType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CurveSegmentArrayPropertyTypes {
				for _, target := range src.ArcString {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.CurveSegmentArrayPropertyType", "ArcString", refNames, formGroup, probe.formStage)
		}
	case *gml.ArcType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumDerivativesAtStart", inst.NumDerivativesAtStart, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumDerivativesAtEnd", inst.NumDerivativesAtEnd, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumDerivativeInterior", inst.NumDerivativeInterior, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Interpolation", inst.Interpolation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("NumArc", inst.NumArc, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Pos", inst.Pos, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionType](), probe.formStage)
		StageSetAssociationFieldToForm("PointProperty", inst.PointProperty, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.PointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PosList", inst.PosList, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionListType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CurveSegmentArrayPropertyTypes {
				for _, target := range src.Arc {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.CurveSegmentArrayPropertyType", "Arc", refNames, formGroup, probe.formStage)
		}
	case *gml.BoundingShapeType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Envelope", inst.Envelope, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.EnvelopeType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AIXMBasicMessages {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AIXMBasicMessage", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AIXMBasicMessageTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AIXMBasicMessageType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AeronauticalGroundLightType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAreaType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMFeatureBaseTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMFeatureTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmabstractgmlobjecttypes.AbstractAIXMFeatureType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMMessageBaseTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmabstractgmlobjecttypes.AbstractAIXMMessageBaseType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMMessageTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmabstractgmlobjecttypes.AbstractAIXMMessageType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractFeatureTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.AbstractFeatureType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.DynamicFeatureTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.DynamicFeatureType", "BoundedBy", refNames, formGroup, probe.formStage)
		}
	case *gml.CircleByCenterPointType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumDerivativesAtStart", inst.NumDerivativesAtStart, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumDerivativesAtEnd", inst.NumDerivativesAtEnd, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumDerivativeInterior", inst.NumDerivativeInterior, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Interpolation", inst.Interpolation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("NumArc", inst.NumArc, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Pos", inst.Pos, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionType](), probe.formStage)
		StageSetAssociationFieldToForm("PointProperty", inst.PointProperty, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.PointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PosList", inst.PosList, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionListType](), probe.formStage)
		StageSetAssociationFieldToForm("Radius", inst.Radius, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.LengthType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CurveSegmentArrayPropertyTypes {
				for _, target := range src.CircleByCenterPoint {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.CurveSegmentArrayPropertyType", "CircleByCenterPoint", refNames, formGroup, probe.formStage)
		}
	case *gml.CircleType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumDerivativesAtStart", inst.NumDerivativesAtStart, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumDerivativesAtEnd", inst.NumDerivativesAtEnd, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumDerivativeInterior", inst.NumDerivativeInterior, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Interpolation", inst.Interpolation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("NumArc", inst.NumArc, probe.formStage, formGroup)

		{
			// Slice of pointers: Pos
			div := (&form.FormDiv{Name: "Pos"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Pos {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Pos",
				Label: "Pos",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: PointProperty
			div := (&form.FormDiv{Name: "PointProperty"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.PointProperty {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "PointProperty",
				Label: "PointProperty",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("PosList", inst.PosList, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionListType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CurveSegmentArrayPropertyTypes {
				for _, target := range src.Circle {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.CurveSegmentArrayPropertyType", "Circle", refNames, formGroup, probe.formStage)
		}
	case *gml.CodeType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CodeSpace", inst.CodeSpace, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AIXMBasicMessages {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AIXMBasicMessage", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AIXMBasicMessageTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AIXMBasicMessageType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractPropertiesWithScheduleTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractPropertiesWithScheduleType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AeronauticalGroundLightTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AeronauticalGroundLightType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CurveTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CurveType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedCurveTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedCurveType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedPointTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedPointType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedSurfaceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedSurfaceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementStatusTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LightElementStatusType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LightElementType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LinguisticNoteTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LinguisticNoteType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NoteTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NoteType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAreaTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAreaType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PointTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PointType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SurfaceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SurfaceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TimesheetType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureLightingStatusTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureLightingStatusType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructurePartType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMFeatureBaseTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMFeatureTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmabstractgmlobjecttypes.AbstractAIXMFeatureType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMMessageBaseTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmabstractgmlobjecttypes.AbstractAIXMMessageBaseType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMMessageTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmabstractgmlobjecttypes.AbstractAIXMMessageType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractExtensionTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmabstractgmlobjecttypes.AbstractExtensionType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractCurveTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.AbstractCurveType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractFeatureTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.AbstractFeatureType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractGMLTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.AbstractGMLType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractGeometricPrimitiveTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.AbstractGeometricPrimitiveType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractGeometryTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.AbstractGeometryType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractSurfaceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.AbstractSurfaceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractTimeGeometricPrimitiveTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.AbstractTimeGeometricPrimitiveType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractTimeObjectTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.AbstractTimeObjectType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractTimePrimitiveTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.AbstractTimePrimitiveType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.AbstractTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CompositeCurveTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.CompositeCurveType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CurveTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.CurveType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.DynamicFeatureTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.DynamicFeatureType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.OrientableCurveTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.OrientableCurveType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.PointTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.PointType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.SurfaceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.SurfaceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.TimeInstantTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.TimeInstantType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.TimePeriodTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.TimePeriodType", "NameXSD", refNames, formGroup, probe.formStage)
		}
	case *gml.CodeWithAuthorityType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CodeSpace", inst.CodeSpace, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AIXMBasicMessages {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AIXMBasicMessage", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AIXMBasicMessageTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AIXMBasicMessageType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractPropertiesWithScheduleTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractPropertiesWithScheduleType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AeronauticalGroundLightTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AeronauticalGroundLightType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CurveTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CurveType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedCurveTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedCurveType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedPointTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedPointType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedSurfaceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedSurfaceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementStatusTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LightElementStatusType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightElementTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LightElementType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LinguisticNoteTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LinguisticNoteType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NoteTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NoteType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAreaTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAreaType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PointTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PointType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SurfaceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SurfaceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TimesheetType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureLightingStatusTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureLightingStatusType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePartTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructurePartType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMFeatureBaseTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMFeatureTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmabstractgmlobjecttypes.AbstractAIXMFeatureType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMMessageBaseTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmabstractgmlobjecttypes.AbstractAIXMMessageBaseType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMMessageTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmabstractgmlobjecttypes.AbstractAIXMMessageType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractExtensionTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmabstractgmlobjecttypes.AbstractExtensionType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractCurveTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.AbstractCurveType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractFeatureTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.AbstractFeatureType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractGMLTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.AbstractGMLType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractGeometricPrimitiveTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.AbstractGeometricPrimitiveType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractGeometryTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.AbstractGeometryType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractSurfaceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.AbstractSurfaceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractTimeGeometricPrimitiveTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.AbstractTimeGeometricPrimitiveType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractTimeObjectTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.AbstractTimeObjectType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractTimePrimitiveTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.AbstractTimePrimitiveType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.AbstractTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CompositeCurveTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.CompositeCurveType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CurveTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.CurveType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.DynamicFeatureTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.DynamicFeatureType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.OrientableCurveTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.OrientableCurveType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.PointTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.PointType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.SurfaceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.SurfaceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.TimeInstantTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.TimeInstantType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.TimePeriodTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.TimePeriodType", "Identifier", refNames, formGroup, probe.formStage)
		}
	case *gml.CompositeCurveType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("SrsName", inst.SrsName, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("AggregationType", inst.AggregationType, probe.formStage, formGroup)

		{
			// Slice of pointers: CurveMember
			div := (&form.FormDiv{Name: "CurveMember"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.CurveMember {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "CurveMember",
				Label: "CurveMember",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CurvePropertyTypes {
				if src.CompositeCurve == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.CurvePropertyType", "CompositeCurve", refNames, formGroup, probe.formStage)
		}
	case *gml.CurvePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Href", inst.Href, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Arcrole", inst.Arcrole, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Show", inst.Show, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Actuate", inst.Actuate, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Curve", inst.Curve, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CurveType](), probe.formStage)
		StageSetAssociationFieldToForm("OrientableCurve", inst.OrientableCurve, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.OrientableCurveType](), probe.formStage)
		StageSetAssociationFieldToForm("CompositeCurve", inst.CompositeCurve, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CompositeCurveType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CompositeCurveTypes {
				for _, target := range src.CurveMember {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.CompositeCurveType", "CurveMember", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.OrientableCurveTypes {
				if src.BaseCurve == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.OrientableCurveType", "BaseCurve", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.RingTypes {
				for _, target := range src.CurveMember {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.RingType", "CurveMember", refNames, formGroup, probe.formStage)
		}
	case *gml.CurveSegmentArrayPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: LineStringSegment
			div := (&form.FormDiv{Name: "LineStringSegment"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.LineStringSegment {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "LineStringSegment",
				Label: "LineStringSegment",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ArcString
			div := (&form.FormDiv{Name: "ArcString"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ArcString {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ArcString",
				Label: "ArcString",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Arc
			div := (&form.FormDiv{Name: "Arc"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Arc {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Arc",
				Label: "Arc",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Circle
			div := (&form.FormDiv{Name: "Circle"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Circle {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Circle",
				Label: "Circle",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ArcByCenterPoint
			div := (&form.FormDiv{Name: "ArcByCenterPoint"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ArcByCenterPoint {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ArcByCenterPoint",
				Label: "ArcByCenterPoint",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: CircleByCenterPoint
			div := (&form.FormDiv{Name: "CircleByCenterPoint"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.CircleByCenterPoint {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "CircleByCenterPoint",
				Label: "CircleByCenterPoint",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: GeodesicString
			div := (&form.FormDiv{Name: "GeodesicString"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GeodesicString {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GeodesicString",
				Label: "GeodesicString",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Geodesic
			div := (&form.FormDiv{Name: "Geodesic"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Geodesic {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Geodesic",
				Label: "Geodesic",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CurveTypes {
				if src.Segments == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CurveType", "Segments", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedCurveTypes {
				if src.Segments == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedCurveType", "Segments", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CurveTypes {
				if src.Segments == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.CurveType", "Segments", refNames, formGroup, probe.formStage)
		}
	case *gml.CurveType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("SrsName", inst.SrsName, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Segments", inst.Segments, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CurveSegmentArrayPropertyType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CurvePropertyTypes {
				if src.Curve == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.CurvePropertyType", "Curve", refNames, formGroup, probe.formStage)
		}
	case *gml.DirectPositionListType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SrsName", inst.SrsName, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.ArcByCenterPointTypes {
				if src.PosList == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.ArcByCenterPointType", "PosList", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.ArcStringTypes {
				if src.PosList == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.ArcStringType", "PosList", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.ArcTypes {
				if src.PosList == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.ArcType", "PosList", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CircleByCenterPointTypes {
				if src.PosList == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.CircleByCenterPointType", "PosList", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CircleTypes {
				if src.PosList == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.CircleType", "PosList", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.GeodesicStringTypes {
				if src.PosList == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.GeodesicStringType", "PosList", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.GeodesicTypes {
				if src.PosList == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.GeodesicType", "PosList", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.LineStringSegmentTypes {
				if src.PosList == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.LineStringSegmentType", "PosList", refNames, formGroup, probe.formStage)
		}
	case *gml.DirectPositionType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SrsName", inst.SrsName, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedPointTypes {
				if src.Pos == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedPointType", "Pos", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PointTypes {
				if src.Pos == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PointType", "Pos", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.ArcByCenterPointTypes {
				if src.Pos == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.ArcByCenterPointType", "Pos", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.ArcStringTypes {
				for _, target := range src.Pos {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.ArcStringType", "Pos", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.ArcTypes {
				if src.Pos == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.ArcType", "Pos", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CircleByCenterPointTypes {
				if src.Pos == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.CircleByCenterPointType", "Pos", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CircleTypes {
				for _, target := range src.Pos {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.CircleType", "Pos", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.EnvelopeTypes {
				if src.LowerCorner == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.EnvelopeType", "LowerCorner", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.EnvelopeTypes {
				if src.UpperCorner == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.EnvelopeType", "UpperCorner", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.GeodesicStringTypes {
				if src.Pos == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.GeodesicStringType", "Pos", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.GeodesicTypes {
				if src.Pos == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.GeodesicType", "Pos", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.LineStringSegmentTypes {
				for _, target := range src.Pos {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.LineStringSegmentType", "Pos", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.PointTypes {
				if src.Pos == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.PointType", "Pos", refNames, formGroup, probe.formStage)
		}
	case *gml.DynamicFeatureType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("BoundedBy", inst.BoundedBy, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.BoundingShapeType](), probe.formStage)
	case *gml.EnvelopeType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SrsName", inst.SrsName, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("LowerCorner", inst.LowerCorner, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionType](), probe.formStage)
		StageSetAssociationFieldToForm("UpperCorner", inst.UpperCorner, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.BoundingShapeTypes {
				if src.Envelope == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.BoundingShapeType", "Envelope", refNames, formGroup, probe.formStage)
		}
	case *gml.GeodesicStringType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumDerivativesAtStart", inst.NumDerivativesAtStart, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumDerivativesAtEnd", inst.NumDerivativesAtEnd, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumDerivativeInterior", inst.NumDerivativeInterior, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Interpolation", inst.Interpolation, formGroup, probe.formStage)
		StageSetAssociationFieldToForm("PosList", inst.PosList, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionListType](), probe.formStage)
		StageSetAssociationFieldToForm("Pos", inst.Pos, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionType](), probe.formStage)
		StageSetAssociationFieldToForm("PointProperty", inst.PointProperty, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.PointPropertyType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CurveSegmentArrayPropertyTypes {
				for _, target := range src.GeodesicString {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.CurveSegmentArrayPropertyType", "GeodesicString", refNames, formGroup, probe.formStage)
		}
	case *gml.GeodesicType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumDerivativesAtStart", inst.NumDerivativesAtStart, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumDerivativesAtEnd", inst.NumDerivativesAtEnd, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumDerivativeInterior", inst.NumDerivativeInterior, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Interpolation", inst.Interpolation, formGroup, probe.formStage)
		StageSetAssociationFieldToForm("PosList", inst.PosList, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionListType](), probe.formStage)
		StageSetAssociationFieldToForm("Pos", inst.Pos, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionType](), probe.formStage)
		StageSetAssociationFieldToForm("PointProperty", inst.PointProperty, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.PointPropertyType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CurveSegmentArrayPropertyTypes {
				for _, target := range src.Geodesic {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.CurveSegmentArrayPropertyType", "Geodesic", refNames, formGroup, probe.formStage)
		}
	case *gml.LengthType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.ArcByCenterPointTypes {
				if src.Radius == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.ArcByCenterPointType", "Radius", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CircleByCenterPointTypes {
				if src.Radius == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.CircleByCenterPointType", "Radius", refNames, formGroup, probe.formStage)
		}
	case *gml.LineStringSegmentType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumDerivativesAtStart", inst.NumDerivativesAtStart, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumDerivativesAtEnd", inst.NumDerivativesAtEnd, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumDerivativeInterior", inst.NumDerivativeInterior, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Interpolation", inst.Interpolation, formGroup, probe.formStage)

		{
			// Slice of pointers: Pos
			div := (&form.FormDiv{Name: "Pos"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Pos {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Pos",
				Label: "Pos",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: PointProperty
			div := (&form.FormDiv{Name: "PointProperty"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.PointProperty {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "PointProperty",
				Label: "PointProperty",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("PosList", inst.PosList, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionListType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CurveSegmentArrayPropertyTypes {
				for _, target := range src.LineStringSegment {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.CurveSegmentArrayPropertyType", "LineStringSegment", refNames, formGroup, probe.formStage)
		}
	case *gml.MeasureType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Uom", inst.Uom, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)
	case *gml.OrientableCurveType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("SrsName", inst.SrsName, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Orientation", inst.Orientation, formGroup, probe.formStage)
		StageSetAssociationFieldToForm("BaseCurve", inst.BaseCurve, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CurvePropertyType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CurvePropertyTypes {
				if src.OrientableCurve == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.CurvePropertyType", "OrientableCurve", refNames, formGroup, probe.formStage)
		}
	case *gml.PointPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Href", inst.Href, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Arcrole", inst.Arcrole, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Show", inst.Show, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Actuate", inst.Actuate, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Point", inst.Point, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.PointType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.ArcByCenterPointTypes {
				if src.PointProperty == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.ArcByCenterPointType", "PointProperty", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.ArcStringTypes {
				for _, target := range src.PointProperty {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.ArcStringType", "PointProperty", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.ArcTypes {
				if src.PointProperty == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.ArcType", "PointProperty", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CircleByCenterPointTypes {
				if src.PointProperty == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.CircleByCenterPointType", "PointProperty", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.CircleTypes {
				for _, target := range src.PointProperty {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.CircleType", "PointProperty", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.GeodesicStringTypes {
				if src.PointProperty == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.GeodesicStringType", "PointProperty", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.GeodesicTypes {
				if src.PointProperty == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.GeodesicType", "PointProperty", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.LineStringSegmentTypes {
				for _, target := range src.PointProperty {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.LineStringSegmentType", "PointProperty", refNames, formGroup, probe.formStage)
		}
	case *gml.PointType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("SrsName", inst.SrsName, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Pos", inst.Pos, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.DirectPositionType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.PointPropertyTypes {
				if src.Point == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.PointPropertyType", "Point", refNames, formGroup, probe.formStage)
		}
	case *gml.PolygonPatchType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Interpolation", inst.Interpolation, formGroup, probe.formStage)
		StageSetAssociationFieldToForm("Exterior", inst.Exterior, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.AbstractRingPropertyType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.SurfacePatchArrayPropertyTypes {
				for _, target := range src.PolygonPatch {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.SurfacePatchArrayPropertyType", "PolygonPatch", refNames, formGroup, probe.formStage)
		}
	case *gml.RingType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("AggregationType", inst.AggregationType, probe.formStage, formGroup)

		{
			// Slice of pointers: CurveMember
			div := (&form.FormDiv{Name: "CurveMember"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.CurveMember {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "CurveMember",
				Label: "CurveMember",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractRingPropertyTypes {
				if src.Ring == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.AbstractRingPropertyType", "Ring", refNames, formGroup, probe.formStage)
		}
	case *gml.SurfacePatchArrayPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: PolygonPatch
			div := (&form.FormDiv{Name: "PolygonPatch"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.PolygonPatch {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "PolygonPatch",
				Label: "PolygonPatch",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedSurfaceTypes {
				if src.Patches == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedSurfaceType", "Patches", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SurfaceTypes {
				if src.Patches == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SurfaceType", "Patches", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.SurfaceTypes {
				if src.Patches == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.SurfaceType", "Patches", refNames, formGroup, probe.formStage)
		}
	case *gml.SurfaceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("SrsName", inst.SrsName, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Patches", inst.Patches, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.SurfacePatchArrayPropertyType](), probe.formStage)
	case *gml.TimeInstantType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("Frame", inst.Frame, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("TimePosition", inst.TimePosition, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePositionType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.TimePrimitivePropertyTypes {
				if src.TimeInstant == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.TimePrimitivePropertyType", "TimeInstant", refNames, formGroup, probe.formStage)
		}
	case *gml.TimePeriodType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("Frame", inst.Frame, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("BeginPosition", inst.BeginPosition, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePositionType](), probe.formStage)
		StageSetAssociationFieldToForm("EndPosition", inst.EndPosition, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePositionType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.TimePrimitivePropertyTypes {
				if src.TimePeriod == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.TimePrimitivePropertyType", "TimePeriod", refNames, formGroup, probe.formStage)
		}
	case *gml.TimePositionType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Frame", inst.Frame, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("IndeterminatePosition", inst.IndeterminatePosition, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.TimeInstantTypes {
				if src.TimePosition == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.TimeInstantType", "TimePosition", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.TimePeriodTypes {
				if src.BeginPosition == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.TimePeriodType", "BeginPosition", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.TimePeriodTypes {
				if src.EndPosition == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.TimePeriodType", "EndPosition", refNames, formGroup, probe.formStage)
		}
	case *gml.TimePrimitivePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type_XSD", inst.Type_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Href", inst.Href, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Arcrole", inst.Arcrole, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Show", inst.Show, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Actuate", inst.Actuate, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("TimeInstant", inst.TimeInstant, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimeInstantType](), probe.formStage)
		StageSetAssociationFieldToForm("TimePeriod", inst.TimePeriod, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.TimePeriodType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AeronauticalGroundLightTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AeronauticalGroundLightTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AeronauticalGroundLightTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAreaTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAreaTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAreaTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VerticalStructureTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VerticalStructureTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMTimeSliceBaseTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceBaseType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.GmlStage.AbstractTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("gml.AbstractTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}
	case *xlink.Xlink_SchemaElement:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
	default:
		_ = inst
	}
}
