// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/xlink"
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
		StageSetAssociationFieldToForm("Airspace", inst.Airspace, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceType](), probe.formStage)
		StageSetAssociationFieldToForm("AuthorityForAirspace", inst.AuthorityForAirspace, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AuthorityForAirspaceType](), probe.formStage)
		StageSetAssociationFieldToForm("GeoBorder", inst.GeoBorder, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.GeoBorderType](), probe.formStage)

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
			// Slice of pointers: AbstractAirspaceExtension
			div := (&form.FormDiv{Name: "AbstractAirspaceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAirspaceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAirspaceExtension",
				Label: "AbstractAirspaceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_1:
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
			// Slice of pointers: AbstractAirspaceActivationExtension
			div := (&form.FormDiv{Name: "AbstractAirspaceActivationExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAirspaceActivationExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAirspaceActivationExtension",
				Label: "AbstractAirspaceActivationExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceActivationTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceActivationType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_10:
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
	case *aixmfeatures.A_extension_11:
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
	case *aixmfeatures.A_extension_12:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractAircraftCharacteristicExtension
			div := (&form.FormDiv{Name: "AbstractAircraftCharacteristicExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAircraftCharacteristicExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAircraftCharacteristicExtension",
				Label: "AbstractAircraftCharacteristicExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftCharacteristicTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftCharacteristicType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_13:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractAirspaceLayerExtension
			div := (&form.FormDiv{Name: "AbstractAirspaceLayerExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAirspaceLayerExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAirspaceLayerExtension",
				Label: "AbstractAirspaceLayerExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceLayerTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceLayerType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_14:
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
			// Slice of pointers: AbstractAirspaceGeometryComponentExtension
			div := (&form.FormDiv{Name: "AbstractAirspaceGeometryComponentExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAirspaceGeometryComponentExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAirspaceGeometryComponentExtension",
				Label: "AbstractAirspaceGeometryComponentExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceGeometryComponentTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceGeometryComponentType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_3:
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
			// Slice of pointers: AbstractAirspaceLayerClassExtension
			div := (&form.FormDiv{Name: "AbstractAirspaceLayerClassExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAirspaceLayerClassExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAirspaceLayerClassExtension",
				Label: "AbstractAirspaceLayerClassExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceLayerClassTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceLayerClassType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_4:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractAirspaceVolumeExtension
			div := (&form.FormDiv{Name: "AbstractAirspaceVolumeExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAirspaceVolumeExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAirspaceVolumeExtension",
				Label: "AbstractAirspaceVolumeExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceVolumeTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceVolumeType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_5:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractAirspaceVolumeDependencyExtension
			div := (&form.FormDiv{Name: "AbstractAirspaceVolumeDependencyExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAirspaceVolumeDependencyExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAirspaceVolumeDependencyExtension",
				Label: "AbstractAirspaceVolumeDependencyExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceVolumeDependencyTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceVolumeDependencyType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_6:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractAuthorityForAirspaceExtension
			div := (&form.FormDiv{Name: "AbstractAuthorityForAirspaceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAuthorityForAirspaceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAuthorityForAirspaceExtension",
				Label: "AbstractAuthorityForAirspaceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForAirspaceTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForAirspaceTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_7:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractGeoBorderExtension
			div := (&form.FormDiv{Name: "AbstractGeoBorderExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractGeoBorderExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractGeoBorderExtension",
				Label: "AbstractGeoBorderExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GeoBorderTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GeoBorderTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_8:
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
	case *aixmfeatures.A_extension_9:
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
	case *aixmfeatures.AircraftCharacteristicPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AircraftCharacteristic", inst.AircraftCharacteristic, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AircraftCharacteristicType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceActivationTypes {
				for _, target := range src.Aircraft {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceActivationType", "Aircraft", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AircraftCharacteristicType:
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
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAircraftType](), probe.formStage)
		StageSetAssociationFieldToForm("Engine", inst.Engine, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAircraftEngineType](), probe.formStage)
		StageSetAssociationFieldToForm("NumberEngine", inst.NumberEngine, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAircraftEngineNumberType](), probe.formStage)
		StageSetAssociationFieldToForm("TypeAircraftICAO", inst.TypeAircraftICAO, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAircraftICAOType](), probe.formStage)
		StageSetAssociationFieldToForm("AircraftLandingCategory", inst.AircraftLandingCategory, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAircraftCategoryType](), probe.formStage)
		StageSetAssociationFieldToForm("WingSpan", inst.WingSpan, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)
		StageSetAssociationFieldToForm("WingSpanInterpretation", inst.WingSpanInterpretation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeValueInterpretationType](), probe.formStage)
		StageSetAssociationFieldToForm("ClassWingSpan", inst.ClassWingSpan, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAircraftWingspanClassType](), probe.formStage)
		StageSetAssociationFieldToForm("Weight", inst.Weight, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValWeightType](), probe.formStage)
		StageSetAssociationFieldToForm("WeightInterpretation", inst.WeightInterpretation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeValueInterpretationType](), probe.formStage)
		StageSetAssociationFieldToForm("Passengers", inst.Passengers, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoNumberType](), probe.formStage)
		StageSetAssociationFieldToForm("PassengersInterpretation", inst.PassengersInterpretation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeValueInterpretationType](), probe.formStage)
		StageSetAssociationFieldToForm("Speed", inst.Speed, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValSpeedType](), probe.formStage)
		StageSetAssociationFieldToForm("SpeedInterpretation", inst.SpeedInterpretation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeValueInterpretationType](), probe.formStage)
		StageSetAssociationFieldToForm("WakeTurbulence", inst.WakeTurbulence, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeWakeTurbulenceType](), probe.formStage)
		StageSetAssociationFieldToForm("NavigationEquipment", inst.NavigationEquipment, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavigationEquipmentType](), probe.formStage)
		StageSetAssociationFieldToForm("NavigationSpecification", inst.NavigationSpecification, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavigationSpecificationType](), probe.formStage)
		StageSetAssociationFieldToForm("VerticalSeparationCapability", inst.VerticalSeparationCapability, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRVSMType](), probe.formStage)
		StageSetAssociationFieldToForm("AntiCollisionAndSeparationEquipment", inst.AntiCollisionAndSeparationEquipment, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeEquipmentAntiCollisionType](), probe.formStage)
		StageSetAssociationFieldToForm("CommunicationEquipment", inst.CommunicationEquipment, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCommunicationModeType](), probe.formStage)
		StageSetAssociationFieldToForm("SurveillanceEquipment", inst.SurveillanceEquipment, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTransponderType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftCharacteristicPropertyTypes {
				if src.AircraftCharacteristic == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftCharacteristicPropertyType", "AircraftCharacteristic", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirspaceActivationPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AirspaceActivation", inst.AirspaceActivation, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceActivationType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceTimeSliceTypes {
				for _, target := range src.Activation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceTimeSliceType", "Activation", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirspaceActivationType:
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
		StageSetAssociationFieldToForm("Activity", inst.Activity, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAirspaceActivityType](), probe.formStage)
		StageSetAssociationFieldToForm("Status", inst.Status, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeStatusAirspaceType](), probe.formStage)

		{
			// Slice of pointers: Levels
			div := (&form.FormDiv{Name: "Levels"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Levels {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Levels",
				Label: "Levels",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: User
			div := (&form.FormDiv{Name: "User"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.User {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "User",
				Label: "User",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Aircraft
			div := (&form.FormDiv{Name: "Aircraft"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Aircraft {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Aircraft",
				Label: "Aircraft",
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceActivationPropertyTypes {
				if src.AirspaceActivation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceActivationPropertyType", "AirspaceActivation", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirspaceGeometryComponentPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AirspaceGeometryComponent", inst.AirspaceGeometryComponent, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceGeometryComponentType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceTimeSliceTypes {
				for _, target := range src.GeometryComponent {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceTimeSliceType", "GeometryComponent", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirspaceGeometryComponentType:
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
		StageSetAssociationFieldToForm("Operation", inst.Operation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAirspaceAggregationType](), probe.formStage)
		StageSetAssociationFieldToForm("OperationSequence", inst.OperationSequence, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoSequenceType](), probe.formStage)

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
		StageSetAssociationFieldToForm("TheAirspaceVolume", inst.TheAirspaceVolume, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceVolumePropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceGeometryComponentPropertyTypes {
				if src.AirspaceGeometryComponent == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceGeometryComponentPropertyType", "AirspaceGeometryComponent", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirspaceLayerClassPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AirspaceLayerClass", inst.AirspaceLayerClass, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceLayerClassType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceTimeSliceTypes {
				for _, target := range src.Class {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceTimeSliceType", "Class", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirspaceLayerClassType:
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
		StageSetAssociationFieldToForm("Classification", inst.Classification, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAirspaceClassificationType](), probe.formStage)

		{
			// Slice of pointers: AssociatedLevels
			div := (&form.FormDiv{Name: "AssociatedLevels"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AssociatedLevels {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AssociatedLevels",
				Label: "AssociatedLevels",
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceLayerClassPropertyTypes {
				if src.AirspaceLayerClass == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceLayerClassPropertyType", "AirspaceLayerClass", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirspaceLayerPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AirspaceLayer", inst.AirspaceLayer, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceLayerType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceActivationTypes {
				for _, target := range src.Levels {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceActivationType", "Levels", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceLayerClassTypes {
				for _, target := range src.AssociatedLevels {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceLayerClassType", "AssociatedLevels", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirspaceLayerType:
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
		StageSetAssociationFieldToForm("UpperLimit", inst.UpperLimit, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("UpperLimitReference", inst.UpperLimitReference, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), probe.formStage)
		StageSetAssociationFieldToForm("LowerLimit", inst.LowerLimit, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("LowerLimitReference", inst.LowerLimitReference, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), probe.formStage)
		StageSetAssociationFieldToForm("AltitudeInterpretation", inst.AltitudeInterpretation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAltitudeUseType](), probe.formStage)
		StageSetAssociationFieldToForm("DiscreteLevelSeries", inst.DiscreteLevelSeries, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.StandardLevelColumnPropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceLayerPropertyTypes {
				if src.AirspaceLayer == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceLayerPropertyType", "AirspaceLayer", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirspacePropertyType:
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceVolumeDependencyTypes {
				if src.TheAirspace == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceVolumeDependencyType", "TheAirspace", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForAirspaceTimeSliceTypes {
				if src.AssignedAirspace == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForAirspaceTimeSliceType", "AssignedAirspace", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirspaceTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AirspaceTimeSlice", inst.AirspaceTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirspaceTimeSliceType:
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
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAirspaceType](), probe.formStage)
		StageSetAssociationFieldToForm("Designator", inst.Designator, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAirspaceDesignatorType](), probe.formStage)
		StageSetAssociationFieldToForm("LocalType", inst.LocalType, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextNameType](), probe.formStage)
		StageSetAssociationFieldToForm("DesignatorICAO", inst.DesignatorICAO, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("ControlType", inst.ControlType, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeMilitaryOperationsType](), probe.formStage)
		StageSetAssociationFieldToForm("UpperLowerSeparation", inst.UpperLowerSeparation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValFLType](), probe.formStage)

		{
			// Slice of pointers: Class
			div := (&form.FormDiv{Name: "Class"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Class {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Class",
				Label: "Class",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("ProtectedRoute", inst.ProtectedRoute, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RoutePropertyType](), probe.formStage)

		{
			// Slice of pointers: GeometryComponent
			div := (&form.FormDiv{Name: "GeometryComponent"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GeometryComponent {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GeometryComponent",
				Label: "GeometryComponent",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Activation
			div := (&form.FormDiv{Name: "Activation"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Activation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Activation",
				Label: "Activation",
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceTimeSlicePropertyTypes {
				if src.AirspaceTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceTimeSlicePropertyType", "AirspaceTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirspaceType:
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
				if src.Airspace == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "Airspace", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirspaceVolumeDependencyPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AirspaceVolumeDependency", inst.AirspaceVolumeDependency, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceVolumeDependencyType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceVolumeTypes {
				if src.ContributorAirspace == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceVolumeType", "ContributorAirspace", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirspaceVolumeDependencyType:
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
		StageSetAssociationFieldToForm("Dependency", inst.Dependency, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAirspaceDependencyType](), probe.formStage)

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
		StageSetAssociationFieldToForm("TheAirspace", inst.TheAirspace, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspacePropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceVolumeDependencyPropertyTypes {
				if src.AirspaceVolumeDependency == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceVolumeDependencyPropertyType", "AirspaceVolumeDependency", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirspaceVolumePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AirspaceVolume", inst.AirspaceVolume, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceVolumeType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceGeometryComponentTypes {
				if src.TheAirspaceVolume == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceGeometryComponentType", "TheAirspaceVolume", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirspaceVolumeType:
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
		StageSetAssociationFieldToForm("UpperLimit", inst.UpperLimit, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("UpperLimitReference", inst.UpperLimitReference, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), probe.formStage)
		StageSetAssociationFieldToForm("MaximumLimit", inst.MaximumLimit, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("MaximumLimitReference", inst.MaximumLimitReference, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), probe.formStage)
		StageSetAssociationFieldToForm("LowerLimit", inst.LowerLimit, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("LowerLimitReference", inst.LowerLimitReference, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), probe.formStage)
		StageSetAssociationFieldToForm("MinimumLimit", inst.MinimumLimit, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("MinimumLimitReference", inst.MinimumLimitReference, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), probe.formStage)
		StageSetAssociationFieldToForm("Width", inst.Width, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)
		StageSetAssociationFieldToForm("HorizontalProjection", inst.HorizontalProjection, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SurfacePropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Centreline", inst.Centreline, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CurvePropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("ContributorAirspace", inst.ContributorAirspace, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceVolumeDependencyPropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceVolumePropertyTypes {
				if src.AirspaceVolume == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceVolumePropertyType", "AirspaceVolume", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AuthorityForAirspaceTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AuthorityForAirspaceTimeSlice", inst.AuthorityForAirspaceTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AuthorityForAirspaceTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForAirspaceTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForAirspaceType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AuthorityForAirspaceTimeSliceType:
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
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAuthorityType](), probe.formStage)
		StageSetAssociationFieldToForm("ResponsibleOrganisation", inst.ResponsibleOrganisation, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.OrganisationAuthorityPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("AssignedAirspace", inst.AssignedAirspace, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspacePropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForAirspaceTimeSlicePropertyTypes {
				if src.AuthorityForAirspaceTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForAirspaceTimeSlicePropertyType", "AuthorityForAirspaceTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AuthorityForAirspaceType:
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
				if src.AuthorityForAirspace == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "AuthorityForAirspace", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.CurvePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Curve", inst.Curve, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CurveType](), probe.formStage)
		StageSetAssociationFieldToForm("ElevatedCurve", inst.ElevatedCurve, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedCurveType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceVolumeTypes {
				if src.Centreline == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceVolumeType", "Centreline", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GeoBorderTimeSliceTypes {
				if src.Border == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GeoBorderTimeSliceType", "Border", refNames, formGroup, probe.formStage)
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

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CurvePropertyTypes {
				if src.Curve == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CurvePropertyType", "Curve", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.CurvePropertyTypes {
				if src.ElevatedCurve == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CurvePropertyType", "ElevatedCurve", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.SurfacePropertyTypes {
				if src.ElevatedSurface == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SurfacePropertyType", "ElevatedSurface", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.GeoBorderTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("GeoBorderTimeSlice", inst.GeoBorderTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.GeoBorderTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GeoBorderTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GeoBorderType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.GeoBorderTimeSliceType:
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
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeGeoBorderType](), probe.formStage)
		StageSetAssociationFieldToForm("Border", inst.Border, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CurvePropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.GeoBorderTimeSlicePropertyTypes {
				if src.GeoBorderTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GeoBorderTimeSlicePropertyType", "GeoBorderTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.GeoBorderType:
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
				if src.GeoBorder == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "GeoBorder", refNames, formGroup, probe.formStage)
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
	case *aixmfeatures.NotePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Note", inst.Note, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NoteType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftCharacteristicTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftCharacteristicType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceActivationTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceActivationType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceGeometryComponentTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceGeometryComponentType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceLayerClassTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceLayerClassType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceLayerTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceLayerType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceVolumeDependencyTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceVolumeDependencyType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceVolumeTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceVolumeType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForAirspaceTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForAirspaceTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.GeoBorderTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GeoBorderTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceActivationTypes {
				for _, target := range src.SpecialDateAuthority {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceActivationType", "SpecialDateAuthority", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceActivationTypes {
				for _, target := range src.User {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceActivationType", "User", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceLayerClassTypes {
				for _, target := range src.SpecialDateAuthority {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceLayerClassType", "SpecialDateAuthority", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForAirspaceTimeSliceTypes {
				if src.ResponsibleOrganisation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForAirspaceTimeSliceType", "ResponsibleOrganisation", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RoutePropertyType:
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceTimeSliceTypes {
				if src.ProtectedRoute == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceTimeSliceType", "ProtectedRoute", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.StandardLevelColumnPropertyType:
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceLayerTypes {
				if src.DiscreteLevelSeries == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceLayerType", "DiscreteLevelSeries", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SurfacePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Surface", inst.Surface, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SurfaceType](), probe.formStage)
		StageSetAssociationFieldToForm("ElevatedSurface", inst.ElevatedSurface, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedSurfaceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceVolumeTypes {
				if src.HorizontalProjection == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceVolumeType", "HorizontalProjection", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceActivationTypes {
				for _, target := range src.TimeInterval {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceActivationType", "TimeInterval", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceLayerClassTypes {
				for _, target := range src.TimeInterval {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceLayerClassType", "TimeInterval", refNames, formGroup, probe.formStage)
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
	case *aixmdatatypes.CodeAircraftCategoryType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftCharacteristicTypes {
				if src.AircraftLandingCategory == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftCharacteristicType", "AircraftLandingCategory", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeAircraftEngineNumberType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftCharacteristicTypes {
				if src.NumberEngine == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftCharacteristicType", "NumberEngine", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeAircraftEngineType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftCharacteristicTypes {
				if src.Engine == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftCharacteristicType", "Engine", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeAircraftICAOType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftCharacteristicTypes {
				if src.TypeAircraftICAO == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftCharacteristicType", "TypeAircraftICAO", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeAircraftType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftCharacteristicTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftCharacteristicType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeAircraftWingspanClassType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftCharacteristicTypes {
				if src.ClassWingSpan == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftCharacteristicType", "ClassWingSpan", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeAirspaceActivityType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceActivationTypes {
				if src.Activity == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceActivationType", "Activity", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeAirspaceAggregationType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceGeometryComponentTypes {
				if src.Operation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceGeometryComponentType", "Operation", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeAirspaceClassificationType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceLayerClassTypes {
				if src.Classification == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceLayerClassType", "Classification", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeAirspaceDependencyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceVolumeDependencyTypes {
				if src.Dependency == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceVolumeDependencyType", "Dependency", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeAirspaceDesignatorType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceTimeSliceTypes {
				if src.Designator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceTimeSliceType", "Designator", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeAirspaceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeAltitudeUseType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceLayerTypes {
				if src.AltitudeInterpretation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceLayerType", "AltitudeInterpretation", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeAuthorityType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForAirspaceTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForAirspaceTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeCommunicationModeType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftCharacteristicTypes {
				if src.CommunicationEquipment == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftCharacteristicType", "CommunicationEquipment", refNames, formGroup, probe.formStage)
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
	case *aixmdatatypes.CodeEquipmentAntiCollisionType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftCharacteristicTypes {
				if src.AntiCollisionAndSeparationEquipment == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftCharacteristicType", "AntiCollisionAndSeparationEquipment", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeGeoBorderType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GeoBorderTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GeoBorderTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeMilitaryOperationsType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceTimeSliceTypes {
				if src.ControlType == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceTimeSliceType", "ControlType", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeNavigationEquipmentType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftCharacteristicTypes {
				if src.NavigationEquipment == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftCharacteristicType", "NavigationEquipment", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeNavigationSpecificationType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftCharacteristicTypes {
				if src.NavigationSpecification == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftCharacteristicType", "NavigationSpecification", refNames, formGroup, probe.formStage)
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
	case *aixmdatatypes.CodeRVSMType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftCharacteristicTypes {
				if src.VerticalSeparationCapability == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftCharacteristicType", "VerticalSeparationCapability", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeStatusAirspaceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceActivationTypes {
				if src.Status == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceActivationType", "Status", refNames, formGroup, probe.formStage)
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
	case *aixmdatatypes.CodeTransponderType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftCharacteristicTypes {
				if src.SurveillanceEquipment == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftCharacteristicType", "SurveillanceEquipment", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeValueInterpretationType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftCharacteristicTypes {
				if src.WingSpanInterpretation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftCharacteristicType", "WingSpanInterpretation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftCharacteristicTypes {
				if src.WeightInterpretation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftCharacteristicType", "WeightInterpretation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftCharacteristicTypes {
				if src.PassengersInterpretation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftCharacteristicType", "PassengersInterpretation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftCharacteristicTypes {
				if src.SpeedInterpretation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftCharacteristicType", "SpeedInterpretation", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevatedSurfaceTypes {
				if src.VerticalDatum == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevatedSurfaceType", "VerticalDatum", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeVerticalReferenceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceLayerTypes {
				if src.UpperLimitReference == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceLayerType", "UpperLimitReference", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceLayerTypes {
				if src.LowerLimitReference == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceLayerType", "LowerLimitReference", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceVolumeTypes {
				if src.UpperLimitReference == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceVolumeType", "UpperLimitReference", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceVolumeTypes {
				if src.MaximumLimitReference == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceVolumeType", "MaximumLimitReference", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceVolumeTypes {
				if src.LowerLimitReference == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceVolumeType", "LowerLimitReference", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceVolumeTypes {
				if src.MinimumLimitReference == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceVolumeType", "MinimumLimitReference", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeWakeTurbulenceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftCharacteristicTypes {
				if src.WakeTurbulence == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftCharacteristicType", "WakeTurbulence", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeYesNoType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceTimeSliceTypes {
				if src.DesignatorICAO == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceTimeSliceType", "DesignatorICAO", refNames, formGroup, probe.formStage)
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
	case *aixmdatatypes.NoNumberType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftCharacteristicTypes {
				if src.Passengers == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftCharacteristicType", "Passengers", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.NoSequenceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceGeometryComponentTypes {
				if src.OperationSequence == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceGeometryComponentType", "OperationSequence", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.TextNameType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceTimeSliceTypes {
				if src.LocalType == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceTimeSliceType", "LocalType", refNames, formGroup, probe.formStage)
		}
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftCharacteristicTypes {
				if src.WingSpan == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftCharacteristicType", "WingSpan", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceVolumeTypes {
				if src.Width == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceVolumeType", "Width", refNames, formGroup, probe.formStage)
		}

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.SurfaceTypes {
				if src.HorizontalAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SurfaceType", "HorizontalAccuracy", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.ValDistanceVerticalType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Uom", inst.Uom, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceLayerTypes {
				if src.UpperLimit == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceLayerType", "UpperLimit", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceLayerTypes {
				if src.LowerLimit == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceLayerType", "LowerLimit", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceVolumeTypes {
				if src.UpperLimit == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceVolumeType", "UpperLimit", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceVolumeTypes {
				if src.MaximumLimit == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceVolumeType", "MaximumLimit", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceVolumeTypes {
				if src.LowerLimit == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceVolumeType", "LowerLimit", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceVolumeTypes {
				if src.MinimumLimit == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceVolumeType", "MinimumLimit", refNames, formGroup, probe.formStage)
		}

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
	case *aixmdatatypes.ValFLType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Uom", inst.Uom, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceTimeSliceTypes {
				if src.UpperLowerSeparation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceTimeSliceType", "UpperLowerSeparation", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.ValSpeedType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Uom", inst.Uom, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftCharacteristicTypes {
				if src.Speed == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftCharacteristicType", "Speed", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.ValWeightType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Uom", inst.Uom, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftCharacteristicTypes {
				if src.Weight == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftCharacteristicType", "Weight", refNames, formGroup, probe.formStage)
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
				for _, target := range src.AbstractAirspaceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension", "AbstractAirspaceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_1s {
				for _, target := range src.AbstractPropertiesWithScheduleExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_1", "AbstractPropertiesWithScheduleExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_1s {
				for _, target := range src.AbstractAirspaceActivationExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_1", "AbstractAirspaceActivationExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_10s {
				for _, target := range src.AbstractLinguisticNoteExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_10", "AbstractLinguisticNoteExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_11s {
				for _, target := range src.AbstractNoteExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_11", "AbstractNoteExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_12s {
				for _, target := range src.AbstractAircraftCharacteristicExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_12", "AbstractAircraftCharacteristicExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_13s {
				for _, target := range src.AbstractAirspaceLayerExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_13", "AbstractAirspaceLayerExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_14s {
				for _, target := range src.AbstractTimesheetExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_14", "AbstractTimesheetExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_2s {
				for _, target := range src.AbstractAirspaceGeometryComponentExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_2", "AbstractAirspaceGeometryComponentExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_3s {
				for _, target := range src.AbstractPropertiesWithScheduleExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_3", "AbstractPropertiesWithScheduleExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_3s {
				for _, target := range src.AbstractAirspaceLayerClassExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_3", "AbstractAirspaceLayerClassExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_4s {
				for _, target := range src.AbstractAirspaceVolumeExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_4", "AbstractAirspaceVolumeExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_5s {
				for _, target := range src.AbstractAirspaceVolumeDependencyExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_5", "AbstractAirspaceVolumeDependencyExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_6s {
				for _, target := range src.AbstractAuthorityForAirspaceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_6", "AbstractAuthorityForAirspaceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_7s {
				for _, target := range src.AbstractGeoBorderExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_7", "AbstractGeoBorderExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_8s {
				for _, target := range src.AbstractElevatedCurveExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_8", "AbstractElevatedCurveExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_9s {
				for _, target := range src.AbstractElevatedSurfaceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_9", "AbstractElevatedSurfaceExtension", refNames, formGroup, probe.formStage)
		}
	case *aixmabstractgmlobjecttypes.FeatureMetadataPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForAirspaceTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForAirspaceType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GeoBorderTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GeoBorderType", "FeatureMetadata", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForAirspaceTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForAirspaceTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GeoBorderTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GeoBorderTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForAirspaceTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForAirspaceType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GeoBorderTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GeoBorderType", "BoundedBy", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftCharacteristicTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftCharacteristicType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceActivationTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceActivationType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceGeometryComponentTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceGeometryComponentType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceLayerClassTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceLayerClassType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceLayerTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceLayerType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceVolumeDependencyTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceVolumeDependencyType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceVolumeTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceVolumeType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForAirspaceTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForAirspaceTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForAirspaceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForAirspaceType", "NameXSD", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.GeoBorderTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GeoBorderTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GeoBorderTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GeoBorderType", "NameXSD", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftCharacteristicTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftCharacteristicType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceActivationTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceActivationType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceGeometryComponentTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceGeometryComponentType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceLayerClassTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceLayerClassType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceLayerTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceLayerType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceVolumeDependencyTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceVolumeDependencyType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceVolumeTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceVolumeType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForAirspaceTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForAirspaceTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForAirspaceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForAirspaceType", "Identifier", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.GeoBorderTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GeoBorderTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GeoBorderTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GeoBorderType", "Identifier", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForAirspaceTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForAirspaceTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForAirspaceTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForAirspaceTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GeoBorderTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GeoBorderTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GeoBorderTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GeoBorderTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
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
