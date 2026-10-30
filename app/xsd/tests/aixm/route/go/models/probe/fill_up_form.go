// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/xlink"
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
		StageSetAssociationFieldToForm("Route", inst.Route, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RouteType](), probe.formStage)
		StageSetAssociationFieldToForm("RouteSegment", inst.RouteSegment, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RouteSegmentType](), probe.formStage)
		StageSetAssociationFieldToForm("RouteDME", inst.RouteDME, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RouteDMEType](), probe.formStage)
		StageSetAssociationFieldToForm("ChangeOverPoint", inst.ChangeOverPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ChangeOverPointType](), probe.formStage)
		StageSetAssociationFieldToForm("FlightRestriction", inst.FlightRestriction, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FlightRestrictionType](), probe.formStage)
		StageSetAssociationFieldToForm("AirspaceBorderCrossing", inst.AirspaceBorderCrossing, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceBorderCrossingType](), probe.formStage)
		StageSetAssociationFieldToForm("RulesProcedures", inst.RulesProcedures, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RulesProceduresType](), probe.formStage)
		StageSetAssociationFieldToForm("AerialRefuelling", inst.AerialRefuelling, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AerialRefuellingType](), probe.formStage)

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
			// Slice of pointers: AbstractAerialRefuellingExtension
			div := (&form.FormDiv{Name: "AbstractAerialRefuellingExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAerialRefuellingExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAerialRefuellingExtension",
				Label: "AbstractAerialRefuellingExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_1:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractAerialRefuellingAnchorExtension
			div := (&form.FormDiv{Name: "AbstractAerialRefuellingAnchorExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAerialRefuellingAnchorExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAerialRefuellingAnchorExtension",
				Label: "AbstractAerialRefuellingAnchorExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingAnchorTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingAnchorType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_10:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractPointReferenceExtension
			div := (&form.FormDiv{Name: "AbstractPointReferenceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractPointReferenceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractPointReferenceExtension",
				Label: "AbstractPointReferenceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PointReferenceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PointReferenceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_11:
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
	case *aixmfeatures.A_extension_12:
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
	case *aixmfeatures.A_extension_13:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractChangeOverPointExtension
			div := (&form.FormDiv{Name: "AbstractChangeOverPointExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractChangeOverPointExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractChangeOverPointExtension",
				Label: "AbstractChangeOverPointExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ChangeOverPointTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ChangeOverPointTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_14:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractRouteExtension
			div := (&form.FormDiv{Name: "AbstractRouteExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractRouteExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractRouteExtension",
				Label: "AbstractRouteExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_15:
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
			// Slice of pointers: AbstractRouteAvailabilityExtension
			div := (&form.FormDiv{Name: "AbstractRouteAvailabilityExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractRouteAvailabilityExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractRouteAvailabilityExtension",
				Label: "AbstractRouteAvailabilityExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteAvailabilityTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteAvailabilityType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_16:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractRouteDMEExtension
			div := (&form.FormDiv{Name: "AbstractRouteDMEExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractRouteDMEExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractRouteDMEExtension",
				Label: "AbstractRouteDMEExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteDMETimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteDMETimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_17:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractRoutePortionExtension
			div := (&form.FormDiv{Name: "AbstractRoutePortionExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractRoutePortionExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractRoutePortionExtension",
				Label: "AbstractRoutePortionExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RoutePortionTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RoutePortionType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_18:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractRouteSegmentExtension
			div := (&form.FormDiv{Name: "AbstractRouteSegmentExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractRouteSegmentExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractRouteSegmentExtension",
				Label: "AbstractRouteSegmentExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_19:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractAirspaceBorderCrossingExtension
			div := (&form.FormDiv{Name: "AbstractAirspaceBorderCrossingExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAirspaceBorderCrossingExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAirspaceBorderCrossingExtension",
				Label: "AbstractAirspaceBorderCrossingExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceBorderCrossingTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceBorderCrossingTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_2:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractSegmentPointExtension
			div := (&form.FormDiv{Name: "AbstractSegmentPointExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractSegmentPointExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractSegmentPointExtension",
				Label: "AbstractSegmentPointExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractAerialRefuellingPointExtension
			div := (&form.FormDiv{Name: "AbstractAerialRefuellingPointExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAerialRefuellingPointExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAerialRefuellingPointExtension",
				Label: "AbstractAerialRefuellingPointExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingPointTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingPointType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_20:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractDirectFlightExtension
			div := (&form.FormDiv{Name: "AbstractDirectFlightExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractDirectFlightExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractDirectFlightExtension",
				Label: "AbstractDirectFlightExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractDirectFlightClassExtension
			div := (&form.FormDiv{Name: "AbstractDirectFlightClassExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractDirectFlightClassExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractDirectFlightClassExtension",
				Label: "AbstractDirectFlightClassExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectFlightClassTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectFlightClassType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_21:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractDirectFlightExtension
			div := (&form.FormDiv{Name: "AbstractDirectFlightExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractDirectFlightExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractDirectFlightExtension",
				Label: "AbstractDirectFlightExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractDirectFlightSegmentExtension
			div := (&form.FormDiv{Name: "AbstractDirectFlightSegmentExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractDirectFlightSegmentExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractDirectFlightSegmentExtension",
				Label: "AbstractDirectFlightSegmentExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectFlightSegmentTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectFlightSegmentType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_22:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractFlightConditionCircumstanceExtension
			div := (&form.FormDiv{Name: "AbstractFlightConditionCircumstanceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractFlightConditionCircumstanceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractFlightConditionCircumstanceExtension",
				Label: "AbstractFlightConditionCircumstanceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionCircumstanceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionCircumstanceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_23:
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
			// Slice of pointers: AbstractFlightConditionCombinationExtension
			div := (&form.FormDiv{Name: "AbstractFlightConditionCombinationExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractFlightConditionCombinationExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractFlightConditionCombinationExtension",
				Label: "AbstractFlightConditionCombinationExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionCombinationTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionCombinationType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_24:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractFlightConditionElementExtension
			div := (&form.FormDiv{Name: "AbstractFlightConditionElementExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractFlightConditionElementExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractFlightConditionElementExtension",
				Label: "AbstractFlightConditionElementExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionElementType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_25:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractFlightRestrictionExtension
			div := (&form.FormDiv{Name: "AbstractFlightRestrictionExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractFlightRestrictionExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractFlightRestrictionExtension",
				Label: "AbstractFlightRestrictionExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_26:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractFlightRestrictionLevelExtension
			div := (&form.FormDiv{Name: "AbstractFlightRestrictionLevelExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractFlightRestrictionLevelExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractFlightRestrictionLevelExtension",
				Label: "AbstractFlightRestrictionLevelExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionLevelTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionLevelType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_27:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractFlightRestrictionRouteExtension
			div := (&form.FormDiv{Name: "AbstractFlightRestrictionRouteExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractFlightRestrictionRouteExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractFlightRestrictionRouteExtension",
				Label: "AbstractFlightRestrictionRouteExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionRouteTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionRouteType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_28:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractFlightRoutingElementExtension
			div := (&form.FormDiv{Name: "AbstractFlightRoutingElementExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractFlightRoutingElementExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractFlightRoutingElementExtension",
				Label: "AbstractFlightRoutingElementExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRoutingElementTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRoutingElementType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_29:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractRulesProceduresExtension
			div := (&form.FormDiv{Name: "AbstractRulesProceduresExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractRulesProceduresExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractRulesProceduresExtension",
				Label: "AbstractRulesProceduresExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RulesProceduresTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RulesProceduresTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_3:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractAerialRefuellingTrackExtension
			div := (&form.FormDiv{Name: "AbstractAerialRefuellingTrackExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAerialRefuellingTrackExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAerialRefuellingTrackExtension",
				Label: "AbstractAerialRefuellingTrackExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTrackTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTrackType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_30:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractContactInformationExtension
			div := (&form.FormDiv{Name: "AbstractContactInformationExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractContactInformationExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractContactInformationExtension",
				Label: "AbstractContactInformationExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ContactInformationTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ContactInformationType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_31:
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
			// Slice of pointers: AbstractOnlineContactExtension
			div := (&form.FormDiv{Name: "AbstractOnlineContactExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractOnlineContactExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractOnlineContactExtension",
				Label: "AbstractOnlineContactExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OnlineContactTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OnlineContactType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_32:
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
			// Slice of pointers: AbstractPostalAddressExtension
			div := (&form.FormDiv{Name: "AbstractPostalAddressExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractPostalAddressExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractPostalAddressExtension",
				Label: "AbstractPostalAddressExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PostalAddressTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PostalAddressType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_33:
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
			// Slice of pointers: AbstractTelephoneContactExtension
			div := (&form.FormDiv{Name: "AbstractTelephoneContactExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractTelephoneContactExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractTelephoneContactExtension",
				Label: "AbstractTelephoneContactExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TelephoneContactTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TelephoneContactType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_34:
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
	case *aixmfeatures.A_extension_35:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractFlightCharacteristicExtension
			div := (&form.FormDiv{Name: "AbstractFlightCharacteristicExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractFlightCharacteristicExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractFlightCharacteristicExtension",
				Label: "AbstractFlightCharacteristicExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightCharacteristicTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightCharacteristicType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_36:
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
	case *aixmfeatures.A_extension_37:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractMeteorologyExtension
			div := (&form.FormDiv{Name: "AbstractMeteorologyExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractMeteorologyExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractMeteorologyExtension",
				Label: "AbstractMeteorologyExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MeteorologyTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MeteorologyType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_38:
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
	case *aixmfeatures.A_extension_39:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractAltitudeAdjustmentExtension
			div := (&form.FormDiv{Name: "AbstractAltitudeAdjustmentExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAltitudeAdjustmentExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAltitudeAdjustmentExtension",
				Label: "AbstractAltitudeAdjustmentExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AltitudeAdjustmentTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AltitudeAdjustmentType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_4:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractAuthorityForAerialRefuellingExtension
			div := (&form.FormDiv{Name: "AbstractAuthorityForAerialRefuellingExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAuthorityForAerialRefuellingExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAuthorityForAerialRefuellingExtension",
				Label: "AbstractAuthorityForAerialRefuellingExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForAerialRefuellingTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForAerialRefuellingType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_40:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractObstacleAssessmentAreaExtension
			div := (&form.FormDiv{Name: "AbstractObstacleAssessmentAreaExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractObstacleAssessmentAreaExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractObstacleAssessmentAreaExtension",
				Label: "AbstractObstacleAssessmentAreaExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAssessmentAreaTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAssessmentAreaType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_41:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractObstaclePlacementExtension
			div := (&form.FormDiv{Name: "AbstractObstaclePlacementExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractObstaclePlacementExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractObstaclePlacementExtension",
				Label: "AbstractObstaclePlacementExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstaclePlacementTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstaclePlacementType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_42:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractObstructionExtension
			div := (&form.FormDiv{Name: "AbstractObstructionExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractObstructionExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractObstructionExtension",
				Label: "AbstractObstructionExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstructionTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstructionType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_5:
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
	case *aixmfeatures.A_extension_6:
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
	case *aixmfeatures.A_extension_7:
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
	case *aixmfeatures.A_extension_8:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractAngleUseExtension
			div := (&form.FormDiv{Name: "AbstractAngleUseExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAngleUseExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAngleUseExtension",
				Label: "AbstractAngleUseExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleUseTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleUseType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_9:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractSegmentPointExtension
			div := (&form.FormDiv{Name: "AbstractSegmentPointExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractSegmentPointExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractSegmentPointExtension",
				Label: "AbstractSegmentPointExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractEnRouteSegmentPointExtension
			div := (&form.FormDiv{Name: "AbstractEnRouteSegmentPointExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractEnRouteSegmentPointExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractEnRouteSegmentPointExtension",
				Label: "AbstractEnRouteSegmentPointExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.EnRouteSegmentPointTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.EnRouteSegmentPointType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AbstractDirectFlightType:
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
	case *aixmfeatures.AbstractSegmentPointType:
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
	case *aixmfeatures.AerialRefuellingAnchorPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AerialRefuellingAnchor", inst.AerialRefuellingAnchor, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AerialRefuellingAnchorType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSliceTypes {
				for _, target := range src.Anchor {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTimeSliceType", "Anchor", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AerialRefuellingAnchorType:
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
		StageSetAssociationFieldToForm("OutboundCourse", inst.OutboundCourse, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), probe.formStage)
		StageSetAssociationFieldToForm("OutboundCourseType", inst.OutboundCourseType, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCourseType](), probe.formStage)
		StageSetAssociationFieldToForm("InboundCourse", inst.InboundCourse, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), probe.formStage)
		StageSetAssociationFieldToForm("TurnDirection", inst.TurnDirection, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDirectionTurnType](), probe.formStage)
		StageSetAssociationFieldToForm("SpeedLimit", inst.SpeedLimit, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValSpeedType](), probe.formStage)
		StageSetAssociationFieldToForm("LegSeparation", inst.LegSeparation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)
		StageSetAssociationFieldToForm("LegLength", inst.LegLength, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)
		StageSetAssociationFieldToForm("RefuellingBaseLevel", inst.RefuellingBaseLevel, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("RefuellingBaseLevelReference", inst.RefuellingBaseLevelReference, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), probe.formStage)
		StageSetAssociationFieldToForm("Extent", inst.Extent, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SurfacePropertyType](), probe.formStage)

		{
			// Slice of pointers: VerticalExtent
			div := (&form.FormDiv{Name: "VerticalExtent"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.VerticalExtent {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "VerticalExtent",
				Label: "VerticalExtent",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Point
			div := (&form.FormDiv{Name: "Point"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Point {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Point",
				Label: "Point",
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingAnchorPropertyTypes {
				if src.AerialRefuellingAnchor == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingAnchorPropertyType", "AerialRefuellingAnchor", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AerialRefuellingPointPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AerialRefuellingPoint", inst.AerialRefuellingPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AerialRefuellingPointType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingAnchorTypes {
				for _, target := range src.Point {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingAnchorType", "Point", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTrackTypes {
				for _, target := range src.Point {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTrackType", "Point", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AerialRefuellingPointType:
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
		StageSetAssociationFieldToForm("ReportingATC", inst.ReportingATC, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeATCReportingType](), probe.formStage)
		StageSetAssociationFieldToForm("FlyOver", inst.FlyOver, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Waypoint", inst.Waypoint, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("RadarGuidance", inst.RadarGuidance, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)

		{
			// Slice of pointers: FacilityMakeup
			div := (&form.FormDiv{Name: "FacilityMakeup"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.FacilityMakeup {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "FacilityMakeup",
				Label: "FacilityMakeup",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("PointChoice_fixDesignatedPoint", inst.PointChoice_fixDesignatedPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PointChoice_navaidSystem", inst.PointChoice_navaidSystem, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PointChoice_position", inst.PointChoice_position, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PointChoice_runwayPoint", inst.PointChoice_runwayPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PointChoice_aimingPoint", inst.PointChoice_aimingPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PointChoice_airportReferencePoint", inst.PointChoice_airportReferencePoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("ExtendedServiceVolume", inst.ExtendedServiceVolume, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RadioFrequencyAreaPropertyType](), probe.formStage)

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
		StageSetAssociationFieldToForm("Sequence", inst.Sequence, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoSequenceType](), probe.formStage)
		StageSetAssociationFieldToForm("UsageType", inst.UsageType, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAerialRefuellingPointType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingPointPropertyTypes {
				if src.AerialRefuellingPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingPointPropertyType", "AerialRefuellingPoint", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AerialRefuellingPropertyType:
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSliceTypes {
				if src.OppositeTrack == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTimeSliceType", "OppositeTrack", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementTypes {
				if src.FlightCondition_aerialRefuellingCondition == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionElementType", "FlightCondition_aerialRefuellingCondition", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRoutingElementTypes {
				if src.Element_aerialRefuellingElement == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRoutingElementType", "Element_aerialRefuellingElement", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AerialRefuellingTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AerialRefuellingTimeSlice", inst.AerialRefuellingTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AerialRefuellingTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AerialRefuellingTimeSliceType:
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
		StageSetAssociationFieldToForm("DesignatorPrefix", inst.DesignatorPrefix, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAerialRefuellingPrefixType](), probe.formStage)
		StageSetAssociationFieldToForm("DesignatorNumber", inst.DesignatorNumber, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoNumberType](), probe.formStage)
		StageSetAssociationFieldToForm("DesignatorSuffix", inst.DesignatorSuffix, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextDesignatorType](), probe.formStage)
		StageSetAssociationFieldToForm("DesignatorDirection", inst.DesignatorDirection, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCardinalDirectionType](), probe.formStage)
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAerialRefuellingType](), probe.formStage)
		StageSetAssociationFieldToForm("RadarBeaconSetting", inst.RadarBeaconSetting, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoNumberType](), probe.formStage)
		StageSetAssociationFieldToForm("XbandRadarSetting", inst.XbandRadarSetting, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoNumberType](), probe.formStage)
		StageSetAssociationFieldToForm("TankerChannel", inst.TankerChannel, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTACANChannelType](), probe.formStage)
		StageSetAssociationFieldToForm("ReceiverChannel", inst.ReceiverChannel, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTACANChannelType](), probe.formStage)
		StageSetAssociationFieldToForm("HelicopterRoute", inst.HelicopterRoute, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("SpecialRefuelling", inst.SpecialRefuelling, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("BidirectionalUse", inst.BidirectionalUse, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("ReverseDirectionTurn", inst.ReverseDirectionTurn, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDirectionTurnType](), probe.formStage)

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
			// Slice of pointers: ProtectingAirspace
			div := (&form.FormDiv{Name: "ProtectingAirspace"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ProtectingAirspace {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ProtectingAirspace",
				Label: "ProtectingAirspace",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Track
			div := (&form.FormDiv{Name: "Track"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Track {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Track",
				Label: "Track",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Anchor
			div := (&form.FormDiv{Name: "Anchor"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Anchor {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Anchor",
				Label: "Anchor",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("OppositeTrack", inst.OppositeTrack, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AerialRefuellingPropertyType](), probe.formStage)

		{
			// Slice of pointers: ManagingOrganisation
			div := (&form.FormDiv{Name: "ManagingOrganisation"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ManagingOrganisation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ManagingOrganisation",
				Label: "ManagingOrganisation",
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSlicePropertyTypes {
				if src.AerialRefuellingTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTimeSlicePropertyType", "AerialRefuellingTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AerialRefuellingTrackPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AerialRefuellingTrack", inst.AerialRefuellingTrack, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AerialRefuellingTrackType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSliceTypes {
				for _, target := range src.Track {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTimeSliceType", "Track", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AerialRefuellingTrackType:
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
		StageSetAssociationFieldToForm("Extent", inst.Extent, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CurvePropertyType](), probe.formStage)

		{
			// Slice of pointers: Point
			div := (&form.FormDiv{Name: "Point"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Point {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Point",
				Label: "Point",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: VerticalExtent
			div := (&form.FormDiv{Name: "VerticalExtent"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.VerticalExtent {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "VerticalExtent",
				Label: "VerticalExtent",
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTrackPropertyTypes {
				if src.AerialRefuellingTrack == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTrackPropertyType", "AerialRefuellingTrack", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AerialRefuellingType:
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
				if src.AerialRefuelling == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "AerialRefuelling", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AircraftCharacteristicPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AircraftCharacteristic", inst.AircraftCharacteristic, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AircraftCharacteristicType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementTypes {
				if src.FlightCondition_aircraft == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionElementType", "FlightCondition_aircraft", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAssessmentAreaTypes {
				for _, target := range src.AircraftCategory {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAssessmentAreaType", "AircraftCategory", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingPointTypes {
				if src.PointChoice_airportReferencePoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingPointType", "PointChoice_airportReferencePoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ChangeOverPointTimeSliceTypes {
				if src.Location_airportReferencePoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ChangeOverPointTimeSliceType", "Location_airportReferencePoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectFlightSegmentTypes {
				if src.End_airportReferencePoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectFlightSegmentType", "End_airportReferencePoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectFlightSegmentTypes {
				if src.Start_airportReferencePoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectFlightSegmentType", "Start_airportReferencePoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.EnRouteSegmentPointTypes {
				if src.PointChoice_airportReferencePoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.EnRouteSegmentPointType", "PointChoice_airportReferencePoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementTypes {
				if src.SignificantPointCondition_airportReferencePoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionElementType", "SignificantPointCondition_airportReferencePoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementTypes {
				if src.FlightCondition_airportHeliportCondition == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionElementType", "FlightCondition_airportHeliportCondition", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRoutingElementTypes {
				if src.PointElement_airportReferencePoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRoutingElementType", "PointElement_airportReferencePoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRoutingElementTypes {
				if src.Element_airportHeliportElement == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRoutingElementType", "Element_airportHeliportElement", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RoutePortionTypes {
				if src.Start_airportReferencePoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RoutePortionType", "Start_airportReferencePoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RoutePortionTypes {
				if src.IntermediatePoint_airportReferencePoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RoutePortionType", "IntermediatePoint_airportReferencePoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RoutePortionTypes {
				if src.End_airportReferencePoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RoutePortionType", "End_airportReferencePoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RulesProceduresTimeSliceTypes {
				for _, target := range src.AffectedLocation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RulesProceduresTimeSliceType", "AffectedLocation", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirspaceBorderCrossingPropertyType:
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementTypes {
				if src.FlightCondition_borderCrossingCondition == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionElementType", "FlightCondition_borderCrossingCondition", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirspaceBorderCrossingTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AirspaceBorderCrossingTimeSlice", inst.AirspaceBorderCrossingTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceBorderCrossingTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceBorderCrossingTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceBorderCrossingType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirspaceBorderCrossingTimeSliceType:
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
		StageSetAssociationFieldToForm("ExitedAirspace", inst.ExitedAirspace, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspacePropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("EnteredAirspace", inst.EnteredAirspace, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspacePropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceBorderCrossingTimeSlicePropertyTypes {
				if src.AirspaceBorderCrossingTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceBorderCrossingTimeSlicePropertyType", "AirspaceBorderCrossingTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirspaceBorderCrossingType:
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
				if src.AirspaceBorderCrossing == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "AirspaceBorderCrossing", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirspaceLayerPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AirspaceLayer", inst.AirspaceLayer, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceLayerType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingAnchorTypes {
				for _, target := range src.VerticalExtent {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingAnchorType", "VerticalExtent", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTrackTypes {
				for _, target := range src.VerticalExtent {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTrackType", "VerticalExtent", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteAvailabilityTypes {
				for _, target := range src.Levels {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteAvailabilityType", "Levels", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSliceTypes {
				for _, target := range src.ProtectingAirspace {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTimeSliceType", "ProtectingAirspace", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceBorderCrossingTimeSliceTypes {
				if src.ExitedAirspace == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceBorderCrossingTimeSliceType", "ExitedAirspace", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceBorderCrossingTimeSliceTypes {
				if src.EnteredAirspace == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceBorderCrossingTimeSliceType", "EnteredAirspace", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementTypes {
				if src.FlightCondition_airspaceCondition == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionElementType", "FlightCondition_airspaceCondition", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRoutingElementTypes {
				if src.Element_airspaceElement == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRoutingElementType", "Element_airspaceElement", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RulesProceduresTimeSliceTypes {
				for _, target := range src.AffectedArea {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RulesProceduresTimeSliceType", "AffectedArea", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AltitudeAdjustmentPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AltitudeAdjustment", inst.AltitudeAdjustment, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AltitudeAdjustmentType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstructionTypes {
				for _, target := range src.Adjustment {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstructionType", "Adjustment", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AltitudeAdjustmentType:
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
		StageSetAssociationFieldToForm("AltitudeAdjustmentType", inst.AltitudeAdjustmentType, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAltitudeAdjustmentType](), probe.formStage)
		StageSetAssociationFieldToForm("PrimaryAlternateMinimum", inst.PrimaryAlternateMinimum, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("AltitudeAdjustment", inst.AltitudeAdjustment, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("LocalRemoteCode", inst.LocalRemoteCode, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AltitudeAdjustmentPropertyTypes {
				if src.AltitudeAdjustment == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AltitudeAdjustmentPropertyType", "AltitudeAdjustment", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AngleIndicationPropertyType:
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleUseTypes {
				if src.TheAngleIndication == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleUseType", "TheAngleIndication", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AngleUsePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AngleUse", inst.AngleUse, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AngleUseType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PointReferenceTypes {
				for _, target := range src.FacilityAngle {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PointReferenceType", "FacilityAngle", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AngleUseType:
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
		StageSetAssociationFieldToForm("AlongCourseGuidance", inst.AlongCourseGuidance, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)

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
		StageSetAssociationFieldToForm("TheAngleIndication", inst.TheAngleIndication, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AngleIndicationPropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleUsePropertyTypes {
				if src.AngleUse == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleUsePropertyType", "AngleUse", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AuthorityForAerialRefuellingPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AuthorityForAerialRefuelling", inst.AuthorityForAerialRefuelling, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AuthorityForAerialRefuellingType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSliceTypes {
				for _, target := range src.ManagingOrganisation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTimeSliceType", "ManagingOrganisation", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AuthorityForAerialRefuellingType:
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
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAuthorityType](), probe.formStage)

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
		StageSetAssociationFieldToForm("TheOrganisationAuthority", inst.TheOrganisationAuthority, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.OrganisationAuthorityPropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForAerialRefuellingPropertyTypes {
				if src.AuthorityForAerialRefuelling == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForAerialRefuellingPropertyType", "AuthorityForAerialRefuelling", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ChangeOverPointTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ChangeOverPointTimeSlice", inst.ChangeOverPointTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ChangeOverPointTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ChangeOverPointTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ChangeOverPointType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ChangeOverPointTimeSliceType:
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
		StageSetAssociationFieldToForm("Distance", inst.Distance, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)
		StageSetAssociationFieldToForm("Location_fixDesignatedPoint", inst.Location_fixDesignatedPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Location_navaidSystem", inst.Location_navaidSystem, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Location_position", inst.Location_position, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Location_runwayPoint", inst.Location_runwayPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Location_aimingPoint", inst.Location_aimingPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Location_airportReferencePoint", inst.Location_airportReferencePoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("ApplicableRoutePortion", inst.ApplicableRoutePortion, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RoutePortionPropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.ChangeOverPointTimeSlicePropertyTypes {
				if src.ChangeOverPointTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ChangeOverPointTimeSlicePropertyType", "ChangeOverPointTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ChangeOverPointType:
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
				if src.ChangeOverPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "ChangeOverPoint", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ContactInformationPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ContactInformation", inst.ContactInformation, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ContactInformationType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionRouteTypes {
				for _, target := range src.Contact {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionRouteType", "Contact", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ContactInformationType:
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
		StageSetAssociationFieldToForm("Title", inst.Title, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextNameType](), probe.formStage)

		{
			// Slice of pointers: Address
			div := (&form.FormDiv{Name: "Address"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Address {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Address",
				Label: "Address",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: NetworkNode
			div := (&form.FormDiv{Name: "NetworkNode"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NetworkNode {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NetworkNode",
				Label: "NetworkNode",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: PhoneFax
			div := (&form.FormDiv{Name: "PhoneFax"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.PhoneFax {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "PhoneFax",
				Label: "PhoneFax",
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.ContactInformationPropertyTypes {
				if src.ContactInformation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ContactInformationPropertyType", "ContactInformation", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.CurvePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Curve", inst.Curve, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CurveType](), probe.formStage)
		StageSetAssociationFieldToForm("ElevatedCurve", inst.ElevatedCurve, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedCurveType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTrackTypes {
				if src.Extent == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTrackType", "Extent", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAssessmentAreaTypes {
				if src.StartingCurve == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAssessmentAreaType", "StartingCurve", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.CurveExtent == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "CurveExtent", refNames, formGroup, probe.formStage)
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
	case *aixmfeatures.DMEPropertyType:
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteDMETimeSliceTypes {
				if src.ReferencedDME == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteDMETimeSliceType", "ReferencedDME", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.DesignatedPointPropertyType:
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingPointTypes {
				if src.PointChoice_fixDesignatedPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingPointType", "PointChoice_fixDesignatedPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ChangeOverPointTimeSliceTypes {
				if src.Location_fixDesignatedPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ChangeOverPointTimeSliceType", "Location_fixDesignatedPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectFlightSegmentTypes {
				if src.End_fixDesignatedPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectFlightSegmentType", "End_fixDesignatedPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectFlightSegmentTypes {
				if src.Start_fixDesignatedPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectFlightSegmentType", "Start_fixDesignatedPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.EnRouteSegmentPointTypes {
				if src.PointChoice_fixDesignatedPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.EnRouteSegmentPointType", "PointChoice_fixDesignatedPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementTypes {
				if src.SignificantPointCondition_fixDesignatedPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionElementType", "SignificantPointCondition_fixDesignatedPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRoutingElementTypes {
				if src.PointElement_fixDesignatedPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRoutingElementType", "PointElement_fixDesignatedPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PointReferenceTypes {
				if src.Point == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PointReferenceType", "Point", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RoutePortionTypes {
				if src.Start_fixDesignatedPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RoutePortionType", "Start_fixDesignatedPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RoutePortionTypes {
				if src.IntermediatePoint_fixDesignatedPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RoutePortionType", "IntermediatePoint_fixDesignatedPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RoutePortionTypes {
				if src.End_fixDesignatedPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RoutePortionType", "End_fixDesignatedPoint", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.DirectFlightClassType:
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
		StageSetAssociationFieldToForm("ExceedLength", inst.ExceedLength, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectFlightPropertyTypes {
				if src.DirectFlightClass == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectFlightPropertyType", "DirectFlightClass", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.DirectFlightPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("DirectFlightClass", inst.DirectFlightClass, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DirectFlightClassType](), probe.formStage)
		StageSetAssociationFieldToForm("DirectFlightSegment", inst.DirectFlightSegment, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DirectFlightSegmentType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementTypes {
				if src.FlightCondition_directFlightCondition == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionElementType", "FlightCondition_directFlightCondition", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.DirectFlightSegmentPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("DirectFlightSegment", inst.DirectFlightSegment, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DirectFlightSegmentType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRoutingElementTypes {
				if src.Element_directFlightElement == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRoutingElementType", "Element_directFlightElement", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.DirectFlightSegmentType:
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
		StageSetAssociationFieldToForm("End_fixDesignatedPoint", inst.End_fixDesignatedPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("End_navaidSystem", inst.End_navaidSystem, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("End_position", inst.End_position, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("End_runwayPoint", inst.End_runwayPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("End_aimingPoint", inst.End_aimingPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("End_airportReferencePoint", inst.End_airportReferencePoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Start_fixDesignatedPoint", inst.Start_fixDesignatedPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Start_navaidSystem", inst.Start_navaidSystem, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Start_position", inst.Start_position, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Start_runwayPoint", inst.Start_runwayPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Start_aimingPoint", inst.Start_aimingPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Start_airportReferencePoint", inst.Start_airportReferencePoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectFlightPropertyTypes {
				if src.DirectFlightSegment == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectFlightPropertyType", "DirectFlightSegment", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectFlightSegmentPropertyTypes {
				if src.DirectFlightSegment == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectFlightSegmentPropertyType", "DirectFlightSegment", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.DistanceIndicationPropertyType:
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.PointReferenceTypes {
				for _, target := range src.FacilityDistance {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PointReferenceType", "FacilityDistance", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.PointPropertyTypes {
				if src.ElevatedPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PointPropertyType", "ElevatedPoint", refNames, formGroup, probe.formStage)
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
	case *aixmfeatures.EnRouteSegmentPointPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("EnRouteSegmentPoint", inst.EnRouteSegmentPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.EnRouteSegmentPointType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.Start == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "Start", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.End == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "End", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.EnRouteSegmentPointType:
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
		StageSetAssociationFieldToForm("ReportingATC", inst.ReportingATC, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeATCReportingType](), probe.formStage)
		StageSetAssociationFieldToForm("FlyOver", inst.FlyOver, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Waypoint", inst.Waypoint, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("RadarGuidance", inst.RadarGuidance, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)

		{
			// Slice of pointers: FacilityMakeup
			div := (&form.FormDiv{Name: "FacilityMakeup"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.FacilityMakeup {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "FacilityMakeup",
				Label: "FacilityMakeup",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("PointChoice_fixDesignatedPoint", inst.PointChoice_fixDesignatedPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PointChoice_navaidSystem", inst.PointChoice_navaidSystem, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PointChoice_position", inst.PointChoice_position, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PointChoice_runwayPoint", inst.PointChoice_runwayPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PointChoice_aimingPoint", inst.PointChoice_aimingPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PointChoice_airportReferencePoint", inst.PointChoice_airportReferencePoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("ExtendedServiceVolume", inst.ExtendedServiceVolume, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RadioFrequencyAreaPropertyType](), probe.formStage)

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
		StageSetAssociationFieldToForm("RoleFreeFlight", inst.RoleFreeFlight, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFreeFlightType](), probe.formStage)
		StageSetAssociationFieldToForm("RoleRVSM", inst.RoleRVSM, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRVSMPointRoleType](), probe.formStage)
		StageSetAssociationFieldToForm("TurnRadius", inst.TurnRadius, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)
		StageSetAssociationFieldToForm("RoleMilitaryTraining", inst.RoleMilitaryTraining, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeMilitaryRoutePointType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.EnRouteSegmentPointPropertyTypes {
				if src.EnRouteSegmentPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.EnRouteSegmentPointPropertyType", "EnRouteSegmentPoint", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.FlightCharacteristicPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("FlightCharacteristic", inst.FlightCharacteristic, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FlightCharacteristicType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementTypes {
				if src.FlightCondition_flight == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionElementType", "FlightCondition_flight", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.FlightCharacteristicType:
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
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlightType](), probe.formStage)
		StageSetAssociationFieldToForm("Rule", inst.Rule, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlightRuleType](), probe.formStage)
		StageSetAssociationFieldToForm("Status", inst.Status, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlightStatusType](), probe.formStage)
		StageSetAssociationFieldToForm("Military", inst.Military, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeMilitaryStatusType](), probe.formStage)
		StageSetAssociationFieldToForm("Origin", inst.Origin, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlightOriginType](), probe.formStage)
		StageSetAssociationFieldToForm("Purpose", inst.Purpose, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlightPurposeType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightCharacteristicPropertyTypes {
				if src.FlightCharacteristic == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightCharacteristicPropertyType", "FlightCharacteristic", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.FlightConditionCircumstancePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("FlightConditionCircumstance", inst.FlightConditionCircumstance, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FlightConditionCircumstanceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementTypes {
				if src.OperationalCondition == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionElementType", "OperationalCondition", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.FlightConditionCircumstanceType:
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
		StageSetAssociationFieldToForm("ReferenceLocation", inst.ReferenceLocation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("RelationWithLocation", inst.RelationWithLocation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeLocationQualifierType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionCircumstancePropertyTypes {
				if src.FlightConditionCircumstance == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionCircumstancePropertyType", "FlightConditionCircumstance", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.FlightConditionCombinationPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("FlightConditionCombination", inst.FlightConditionCombination, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FlightConditionCombinationType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementTypes {
				if src.FlightCondition_operand == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionElementType", "FlightCondition_operand", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionTimeSliceTypes {
				if src.Flight == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionTimeSliceType", "Flight", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.FlightConditionCombinationType:
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
		StageSetAssociationFieldToForm("LogicalOperator", inst.LogicalOperator, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlowConditionOperationType](), probe.formStage)

		{
			// Slice of pointers: Element
			div := (&form.FormDiv{Name: "Element"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Element {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Element",
				Label: "Element",
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionCombinationPropertyTypes {
				if src.FlightConditionCombination == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionCombinationPropertyType", "FlightConditionCombination", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.FlightConditionElementPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("FlightConditionElement", inst.FlightConditionElement, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FlightConditionElementType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionCombinationTypes {
				for _, target := range src.Element {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionCombinationType", "Element", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.FlightConditionElementType:
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
		StageSetAssociationFieldToForm("Index", inst.Index, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoSequenceType](), probe.formStage)
		StageSetAssociationFieldToForm("FlightCondition_weather", inst.FlightCondition_weather, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.MeteorologyPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("FlightCondition_aircraft", inst.FlightCondition_aircraft, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AircraftCharacteristicPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("FlightCondition_flight", inst.FlightCondition_flight, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FlightCharacteristicPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("FlightCondition_directFlightCondition", inst.FlightCondition_directFlightCondition, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DirectFlightPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("FlightCondition_borderCrossingCondition", inst.FlightCondition_borderCrossingCondition, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceBorderCrossingPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("FlightCondition_operand", inst.FlightCondition_operand, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FlightConditionCombinationPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("FlightCondition_routePortionCondition", inst.FlightCondition_routePortionCondition, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RoutePortionPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("FlightCondition_standardInstrumentDepartureCondition", inst.FlightCondition_standardInstrumentDepartureCondition, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.StandardInstrumentDeparturePropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("FlightCondition_standardInstrumentArrivalCondition", inst.FlightCondition_standardInstrumentArrivalCondition, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.StandardInstrumentArrivalPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("FlightCondition_organisationCondition", inst.FlightCondition_organisationCondition, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.OrganisationAuthorityPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("SignificantPointCondition_fixDesignatedPoint", inst.SignificantPointCondition_fixDesignatedPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("SignificantPointCondition_navaidSystem", inst.SignificantPointCondition_navaidSystem, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("SignificantPointCondition_position", inst.SignificantPointCondition_position, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("SignificantPointCondition_runwayPoint", inst.SignificantPointCondition_runwayPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("SignificantPointCondition_aimingPoint", inst.SignificantPointCondition_aimingPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("SignificantPointCondition_airportReferencePoint", inst.SignificantPointCondition_airportReferencePoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("FlightCondition_airspaceCondition", inst.FlightCondition_airspaceCondition, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspacePropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("FlightCondition_airportHeliportCondition", inst.FlightCondition_airportHeliportCondition, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("FlightCondition_aerialRefuellingCondition", inst.FlightCondition_aerialRefuellingCondition, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AerialRefuellingPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("OperationalCondition", inst.OperationalCondition, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FlightConditionCircumstancePropertyType](), probe.formStage)

		{
			// Slice of pointers: FlightLevel
			div := (&form.FormDiv{Name: "FlightLevel"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.FlightLevel {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "FlightLevel",
				Label: "FlightLevel",
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementPropertyTypes {
				if src.FlightConditionElement == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionElementPropertyType", "FlightConditionElement", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.FlightRestrictionLevelPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("FlightRestrictionLevel", inst.FlightRestrictionLevel, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FlightRestrictionLevelType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementTypes {
				for _, target := range src.FlightLevel {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionElementType", "FlightLevel", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRoutingElementTypes {
				for _, target := range src.FlightLevel {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRoutingElementType", "FlightLevel", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.FlightRestrictionLevelType:
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
		StageSetAssociationFieldToForm("UpperLevel", inst.UpperLevel, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("UpperLevelReference", inst.UpperLevelReference, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), probe.formStage)
		StageSetAssociationFieldToForm("LowerLevel", inst.LowerLevel, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("LowerLevelReference", inst.LowerLevelReference, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionLevelPropertyTypes {
				if src.FlightRestrictionLevel == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionLevelPropertyType", "FlightRestrictionLevel", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.FlightRestrictionRoutePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("FlightRestrictionRoute", inst.FlightRestrictionRoute, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FlightRestrictionRouteType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionTimeSliceTypes {
				for _, target := range src.RegulatedRoute {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionTimeSliceType", "RegulatedRoute", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.FlightRestrictionRouteType:
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
		StageSetAssociationFieldToForm("PriorPermission", inst.PriorPermission, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)

		{
			// Slice of pointers: RouteElement
			div := (&form.FormDiv{Name: "RouteElement"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RouteElement {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RouteElement",
				Label: "RouteElement",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Contact
			div := (&form.FormDiv{Name: "Contact"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Contact {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Contact",
				Label: "Contact",
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionRoutePropertyTypes {
				if src.FlightRestrictionRoute == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionRoutePropertyType", "FlightRestrictionRoute", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.FlightRestrictionTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("FlightRestrictionTimeSlice", inst.FlightRestrictionTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FlightRestrictionTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.FlightRestrictionTimeSliceType:
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
		StageSetAssociationFieldToForm("Designator", inst.Designator, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlightRestrictionDesignatorType](), probe.formStage)
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlightRestrictionType](), probe.formStage)
		StageSetAssociationFieldToForm("Instruction", inst.Instruction, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextInstructionType](), probe.formStage)
		StageSetAssociationFieldToForm("Flight", inst.Flight, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FlightConditionCombinationPropertyType](), probe.formStage)

		{
			// Slice of pointers: RegulatedRoute
			div := (&form.FormDiv{Name: "RegulatedRoute"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RegulatedRoute {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RegulatedRoute",
				Label: "RegulatedRoute",
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionTimeSlicePropertyTypes {
				if src.FlightRestrictionTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionTimeSlicePropertyType", "FlightRestrictionTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.FlightRestrictionType:
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
				if src.FlightRestriction == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "FlightRestriction", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.FlightRoutingElementPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("FlightRoutingElement", inst.FlightRoutingElement, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FlightRoutingElementType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionRouteTypes {
				for _, target := range src.RouteElement {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionRouteType", "RouteElement", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.FlightRoutingElementType:
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
		StageSetAssociationFieldToForm("OrderNumber", inst.OrderNumber, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoSequenceType](), probe.formStage)
		StageSetAssociationFieldToForm("Speed", inst.Speed, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValSpeedType](), probe.formStage)
		StageSetAssociationFieldToForm("SpeedReference", inst.SpeedReference, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSpeedReferenceType](), probe.formStage)
		StageSetAssociationFieldToForm("SpeedCriteria", inst.SpeedCriteria, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeComparisonType](), probe.formStage)

		{
			// Slice of pointers: FlightLevel
			div := (&form.FormDiv{Name: "FlightLevel"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.FlightLevel {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "FlightLevel",
				Label: "FlightLevel",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Element_directFlightElement", inst.Element_directFlightElement, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DirectFlightSegmentPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Element_routePortionElement", inst.Element_routePortionElement, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RoutePortionPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Element_standardInstrumentDepartureElement", inst.Element_standardInstrumentDepartureElement, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.StandardInstrumentDeparturePropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Element_standardInstrumentArrivalElement", inst.Element_standardInstrumentArrivalElement, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.StandardInstrumentArrivalPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PointElement_fixDesignatedPoint", inst.PointElement_fixDesignatedPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PointElement_navaidSystem", inst.PointElement_navaidSystem, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PointElement_position", inst.PointElement_position, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PointElement_runwayPoint", inst.PointElement_runwayPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PointElement_aimingPoint", inst.PointElement_aimingPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PointElement_airportReferencePoint", inst.PointElement_airportReferencePoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Element_airspaceElement", inst.Element_airspaceElement, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspacePropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Element_airportHeliportElement", inst.Element_airportHeliportElement, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Element_aerialRefuellingElement", inst.Element_aerialRefuellingElement, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AerialRefuellingPropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRoutingElementPropertyTypes {
				if src.FlightRoutingElement == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRoutingElementPropertyType", "FlightRoutingElement", refNames, formGroup, probe.formStage)
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
	case *aixmfeatures.MeteorologyPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Meteorology", inst.Meteorology, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.MeteorologyType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementTypes {
				if src.FlightCondition_weather == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionElementType", "FlightCondition_weather", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.MeteorologyType:
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
		StageSetAssociationFieldToForm("FlightConditions", inst.FlightConditions, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeMeteoConditionsType](), probe.formStage)
		StageSetAssociationFieldToForm("Visibility", inst.Visibility, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)
		StageSetAssociationFieldToForm("VisibilityInterpretation", inst.VisibilityInterpretation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeValueInterpretationType](), probe.formStage)
		StageSetAssociationFieldToForm("RunwayVisualRange", inst.RunwayVisualRange, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)
		StageSetAssociationFieldToForm("RunwayVisualRangeInterpretation", inst.RunwayVisualRangeInterpretation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeValueInterpretationType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.MeteorologyPropertyTypes {
				if src.Meteorology == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MeteorologyPropertyType", "Meteorology", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.NavaidPropertyType:
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingPointTypes {
				if src.PointChoice_navaidSystem == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingPointType", "PointChoice_navaidSystem", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ChangeOverPointTimeSliceTypes {
				if src.Location_navaidSystem == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ChangeOverPointTimeSliceType", "Location_navaidSystem", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectFlightSegmentTypes {
				if src.End_navaidSystem == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectFlightSegmentType", "End_navaidSystem", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectFlightSegmentTypes {
				if src.Start_navaidSystem == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectFlightSegmentType", "Start_navaidSystem", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.EnRouteSegmentPointTypes {
				if src.PointChoice_navaidSystem == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.EnRouteSegmentPointType", "PointChoice_navaidSystem", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementTypes {
				if src.SignificantPointCondition_navaidSystem == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionElementType", "SignificantPointCondition_navaidSystem", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRoutingElementTypes {
				if src.PointElement_navaidSystem == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRoutingElementType", "PointElement_navaidSystem", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RoutePortionTypes {
				if src.Start_navaidSystem == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RoutePortionType", "Start_navaidSystem", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RoutePortionTypes {
				if src.IntermediatePoint_navaidSystem == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RoutePortionType", "IntermediatePoint_navaidSystem", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RoutePortionTypes {
				if src.End_navaidSystem == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RoutePortionType", "End_navaidSystem", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.NotePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Note", inst.Note, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NoteType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingAnchorTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingAnchorType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingPointTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingPointType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTrackTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTrackType", "Annotation", refNames, formGroup, probe.formStage)
		}

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceBorderCrossingTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceBorderCrossingTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AltitudeAdjustmentTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AltitudeAdjustmentType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleUseTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleUseType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForAerialRefuellingTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForAerialRefuellingType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ChangeOverPointTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ChangeOverPointTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ContactInformationTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ContactInformationType", "Annotation", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectFlightClassTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectFlightClassType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectFlightSegmentTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectFlightSegmentType", "Annotation", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.EnRouteSegmentPointTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.EnRouteSegmentPointType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightCharacteristicTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightCharacteristicType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionCircumstanceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionCircumstanceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionCombinationTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionCombinationType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionElementType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionLevelTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionLevelType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionRouteTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionRouteType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRoutingElementTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRoutingElementType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MeteorologyTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MeteorologyType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAssessmentAreaTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAssessmentAreaType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstaclePlacementTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstaclePlacementType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstructionTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstructionType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OnlineContactTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OnlineContactType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PointReferenceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PointReferenceType", "Annotation", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.PostalAddressTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PostalAddressType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteAvailabilityTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteAvailabilityType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteDMETimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteDMETimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RoutePortionTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RoutePortionType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RulesProceduresTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RulesProceduresTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.TelephoneContactTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TelephoneContactType", "Annotation", refNames, formGroup, probe.formStage)
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
	case *aixmfeatures.ObstacleAssessmentAreaPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ObstacleAssessmentArea", inst.ObstacleAssessmentArea, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ObstacleAssessmentAreaType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.EvaluationArea == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "EvaluationArea", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ObstacleAssessmentAreaType:
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
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeObstacleAssessmentSurfaceType](), probe.formStage)
		StageSetAssociationFieldToForm("SectionNumber", inst.SectionNumber, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoNumberType](), probe.formStage)
		StageSetAssociationFieldToForm("Slope", inst.Slope, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValSlopeType](), probe.formStage)
		StageSetAssociationFieldToForm("AssessedAltitude", inst.AssessedAltitude, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("SlopeLowerAltitude", inst.SlopeLowerAltitude, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("GradientLowHigh", inst.GradientLowHigh, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValSlopeType](), probe.formStage)
		StageSetAssociationFieldToForm("SurfaceZone", inst.SurfaceZone, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeObstructionIdSurfaceZoneType](), probe.formStage)
		StageSetAssociationFieldToForm("SafetyRegulation", inst.SafetyRegulation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextNameType](), probe.formStage)

		{
			// Slice of pointers: AircraftCategory
			div := (&form.FormDiv{Name: "AircraftCategory"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AircraftCategory {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AircraftCategory",
				Label: "AircraftCategory",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: SignificantObstacle
			div := (&form.FormDiv{Name: "SignificantObstacle"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SignificantObstacle {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SignificantObstacle",
				Label: "SignificantObstacle",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Surface", inst.Surface, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SurfacePropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("StartingCurve", inst.StartingCurve, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CurvePropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAssessmentAreaPropertyTypes {
				if src.ObstacleAssessmentArea == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAssessmentAreaPropertyType", "ObstacleAssessmentArea", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ObstaclePlacementPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ObstaclePlacement", inst.ObstaclePlacement, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ObstaclePlacementType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstructionTypes {
				for _, target := range src.ObstaclePlacement {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstructionType", "ObstaclePlacement", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ObstaclePlacementType:
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
		StageSetAssociationFieldToForm("ObstacleBearing", inst.ObstacleBearing, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), probe.formStage)
		StageSetAssociationFieldToForm("ObstacleDistance", inst.ObstacleDistance, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)
		StageSetAssociationFieldToForm("PointType", inst.PointType, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextNameType](), probe.formStage)
		StageSetAssociationFieldToForm("ObstaclePlacement", inst.ObstaclePlacement, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSideType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstaclePlacementPropertyTypes {
				if src.ObstaclePlacement == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstaclePlacementPropertyType", "ObstaclePlacement", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ObstructionPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Obstruction", inst.Obstruction, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ObstructionType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAssessmentAreaTypes {
				for _, target := range src.SignificantObstacle {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAssessmentAreaType", "SignificantObstacle", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ObstructionType:
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
		StageSetAssociationFieldToForm("RequiredClearance", inst.RequiredClearance, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)
		StageSetAssociationFieldToForm("MinimumAltitude", inst.MinimumAltitude, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("SurfacePenetration", inst.SurfacePenetration, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("SlopePenetration", inst.SlopePenetration, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("Controlling", inst.Controlling, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("CloseIn", inst.CloseIn, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("TheVerticalStructure", inst.TheVerticalStructure, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.VerticalStructurePropertyType](), probe.formStage)

		{
			// Slice of pointers: Adjustment
			div := (&form.FormDiv{Name: "Adjustment"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Adjustment {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Adjustment",
				Label: "Adjustment",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ObstaclePlacement
			div := (&form.FormDiv{Name: "ObstaclePlacement"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ObstaclePlacement {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ObstaclePlacement",
				Label: "ObstaclePlacement",
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstructionPropertyTypes {
				if src.Obstruction == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstructionPropertyType", "Obstruction", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.OnlineContactPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("OnlineContact", inst.OnlineContact, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.OnlineContactType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ContactInformationTypes {
				for _, target := range src.NetworkNode {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ContactInformationType", "NetworkNode", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.OnlineContactType:
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
		StageSetAssociationFieldToForm("Network", inst.Network, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTelecomNetworkType](), probe.formStage)
		StageSetAssociationFieldToForm("Linkage", inst.Linkage, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextAddressType](), probe.formStage)
		StageSetAssociationFieldToForm("Protocol", inst.Protocol, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextNameType](), probe.formStage)
		StageSetAssociationFieldToForm("EMail", inst.EMail, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextAddressType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.OnlineContactPropertyTypes {
				if src.OnlineContact == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OnlineContactPropertyType", "OnlineContact", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForAerialRefuellingTypes {
				if src.TheOrganisationAuthority == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForAerialRefuellingType", "TheOrganisationAuthority", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionCombinationTypes {
				for _, target := range src.SpecialDateAuthority {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionCombinationType", "SpecialDateAuthority", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementTypes {
				if src.FlightCondition_organisationCondition == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionElementType", "FlightCondition_organisationCondition", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OnlineContactTypes {
				for _, target := range src.SpecialDateAuthority {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OnlineContactType", "SpecialDateAuthority", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PostalAddressTypes {
				for _, target := range src.SpecialDateAuthority {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PostalAddressType", "SpecialDateAuthority", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteAvailabilityTypes {
				for _, target := range src.SpecialDateAuthority {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteAvailabilityType", "SpecialDateAuthority", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteTimeSliceTypes {
				if src.UserOrganisation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteTimeSliceType", "UserOrganisation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TelephoneContactTypes {
				for _, target := range src.SpecialDateAuthority {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TelephoneContactType", "SpecialDateAuthority", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.PointPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Point", inst.Point, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointType](), probe.formStage)
		StageSetAssociationFieldToForm("ElevatedPoint", inst.ElevatedPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingPointTypes {
				if src.PointChoice_position == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingPointType", "PointChoice_position", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ChangeOverPointTimeSliceTypes {
				if src.Location_position == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ChangeOverPointTimeSliceType", "Location_position", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectFlightSegmentTypes {
				if src.End_position == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectFlightSegmentType", "End_position", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectFlightSegmentTypes {
				if src.Start_position == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectFlightSegmentType", "Start_position", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.EnRouteSegmentPointTypes {
				if src.PointChoice_position == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.EnRouteSegmentPointType", "PointChoice_position", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementTypes {
				if src.SignificantPointCondition_position == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionElementType", "SignificantPointCondition_position", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRoutingElementTypes {
				if src.PointElement_position == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRoutingElementType", "PointElement_position", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RoutePortionTypes {
				if src.Start_position == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RoutePortionType", "Start_position", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RoutePortionTypes {
				if src.IntermediatePoint_position == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RoutePortionType", "IntermediatePoint_position", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RoutePortionTypes {
				if src.End_position == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RoutePortionType", "End_position", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.PointReferencePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("PointReference", inst.PointReference, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointReferenceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingPointTypes {
				for _, target := range src.FacilityMakeup {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingPointType", "FacilityMakeup", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.EnRouteSegmentPointTypes {
				for _, target := range src.FacilityMakeup {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.EnRouteSegmentPointType", "FacilityMakeup", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.PointReferenceType:
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
		StageSetAssociationFieldToForm("Role", inst.Role, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeReferenceRoleType](), probe.formStage)
		StageSetAssociationFieldToForm("PriorFixTolerance", inst.PriorFixTolerance, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceSignedType](), probe.formStage)
		StageSetAssociationFieldToForm("PostFixTolerance", inst.PostFixTolerance, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceSignedType](), probe.formStage)
		StageSetAssociationFieldToForm("Point", inst.Point, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), probe.formStage)

		{
			// Slice of pointers: FacilityAngle
			div := (&form.FormDiv{Name: "FacilityAngle"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.FacilityAngle {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "FacilityAngle",
				Label: "FacilityAngle",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: FacilityDistance
			div := (&form.FormDiv{Name: "FacilityDistance"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.FacilityDistance {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "FacilityDistance",
				Label: "FacilityDistance",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("FixToleranceArea", inst.FixToleranceArea, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SurfacePropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.PointReferencePropertyTypes {
				if src.PointReference == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PointReferencePropertyType", "PointReference", refNames, formGroup, probe.formStage)
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

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PointPropertyTypes {
				if src.Point == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PointPropertyType", "Point", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.PostalAddressPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("PostalAddress", inst.PostalAddress, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PostalAddressType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ContactInformationTypes {
				for _, target := range src.Address {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ContactInformationType", "Address", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.PostalAddressType:
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
		StageSetAssociationFieldToForm("DeliveryPoint", inst.DeliveryPoint, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextAddressType](), probe.formStage)
		StageSetAssociationFieldToForm("City", inst.City, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextNameType](), probe.formStage)
		StageSetAssociationFieldToForm("AdministrativeArea", inst.AdministrativeArea, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextNameType](), probe.formStage)
		StageSetAssociationFieldToForm("PostalCode", inst.PostalCode, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextNameType](), probe.formStage)
		StageSetAssociationFieldToForm("Country", inst.Country, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextNameType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.PostalAddressPropertyTypes {
				if src.PostalAddress == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PostalAddressPropertyType", "PostalAddress", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RadioFrequencyAreaPropertyType:
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingPointTypes {
				if src.ExtendedServiceVolume == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingPointType", "ExtendedServiceVolume", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.EnRouteSegmentPointTypes {
				if src.ExtendedServiceVolume == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.EnRouteSegmentPointType", "ExtendedServiceVolume", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RouteAvailabilityPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("RouteAvailability", inst.RouteAvailability, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RouteAvailabilityType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSliceTypes {
				for _, target := range src.Availability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTimeSliceType", "Availability", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				for _, target := range src.Availability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "Availability", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RouteAvailabilityType:
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
		StageSetAssociationFieldToForm("Direction", inst.Direction, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDirectionType](), probe.formStage)
		StageSetAssociationFieldToForm("CardinalDirection", inst.CardinalDirection, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCardinalDirectionType](), probe.formStage)
		StageSetAssociationFieldToForm("Status", inst.Status, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRouteAvailabilityType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteAvailabilityPropertyTypes {
				if src.RouteAvailability == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteAvailabilityPropertyType", "RouteAvailability", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RouteDMETimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("RouteDMETimeSlice", inst.RouteDMETimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RouteDMETimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteDMETypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteDMEType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RouteDMETimeSliceType:
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
		StageSetAssociationFieldToForm("CriticalDME", inst.CriticalDME, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Satisfactory", inst.Satisfactory, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("ReferencedDME", inst.ReferencedDME, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DMEPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("ApplicableRoutePortion", inst.ApplicableRoutePortion, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RoutePortionPropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteDMETimeSlicePropertyTypes {
				if src.RouteDMETimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteDMETimeSlicePropertyType", "RouteDMETimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RouteDMEType:
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
				if src.RouteDME == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "RouteDME", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RoutePortionPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("RoutePortion", inst.RoutePortion, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RoutePortionType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ChangeOverPointTimeSliceTypes {
				if src.ApplicableRoutePortion == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ChangeOverPointTimeSliceType", "ApplicableRoutePortion", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementTypes {
				if src.FlightCondition_routePortionCondition == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionElementType", "FlightCondition_routePortionCondition", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRoutingElementTypes {
				if src.Element_routePortionElement == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRoutingElementType", "Element_routePortionElement", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteDMETimeSliceTypes {
				if src.ApplicableRoutePortion == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteDMETimeSliceType", "ApplicableRoutePortion", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RoutePortionType:
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
		StageSetAssociationFieldToForm("Start_fixDesignatedPoint", inst.Start_fixDesignatedPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Start_navaidSystem", inst.Start_navaidSystem, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Start_position", inst.Start_position, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Start_runwayPoint", inst.Start_runwayPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Start_aimingPoint", inst.Start_aimingPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Start_airportReferencePoint", inst.Start_airportReferencePoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("IntermediatePoint_fixDesignatedPoint", inst.IntermediatePoint_fixDesignatedPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("IntermediatePoint_navaidSystem", inst.IntermediatePoint_navaidSystem, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("IntermediatePoint_position", inst.IntermediatePoint_position, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("IntermediatePoint_runwayPoint", inst.IntermediatePoint_runwayPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("IntermediatePoint_aimingPoint", inst.IntermediatePoint_aimingPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("IntermediatePoint_airportReferencePoint", inst.IntermediatePoint_airportReferencePoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("ReferencedRoute", inst.ReferencedRoute, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RoutePropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("End_fixDesignatedPoint", inst.End_fixDesignatedPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("End_navaidSystem", inst.End_navaidSystem, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("End_position", inst.End_position, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("End_runwayPoint", inst.End_runwayPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("End_aimingPoint", inst.End_aimingPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("End_airportReferencePoint", inst.End_airportReferencePoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.RoutePortionPropertyTypes {
				if src.RoutePortion == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RoutePortionPropertyType", "RoutePortion", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.RoutePortionTypes {
				if src.ReferencedRoute == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RoutePortionType", "ReferencedRoute", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.RouteFormed == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "RouteFormed", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RouteSegmentTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("RouteSegmentTimeSlice", inst.RouteSegmentTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RouteSegmentTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RouteSegmentTimeSliceType:
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
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeLevelType](), probe.formStage)
		StageSetAssociationFieldToForm("UpperLimit", inst.UpperLimit, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("UpperLimitReference", inst.UpperLimitReference, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), probe.formStage)
		StageSetAssociationFieldToForm("LowerLimit", inst.LowerLimit, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("LowerLimitReference", inst.LowerLimitReference, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), probe.formStage)
		StageSetAssociationFieldToForm("MinimumObstacleClearanceAltitude", inst.MinimumObstacleClearanceAltitude, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("PathType", inst.PathType, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRouteSegmentPathType](), probe.formStage)
		StageSetAssociationFieldToForm("TrueTrack", inst.TrueTrack, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticTrack", inst.MagneticTrack, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), probe.formStage)
		StageSetAssociationFieldToForm("ReverseTrueTrack", inst.ReverseTrueTrack, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), probe.formStage)
		StageSetAssociationFieldToForm("ReverseMagneticTrack", inst.ReverseMagneticTrack, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), probe.formStage)
		StageSetAssociationFieldToForm("Length", inst.Length, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)
		StageSetAssociationFieldToForm("WidthLeft", inst.WidthLeft, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)
		StageSetAssociationFieldToForm("WidthRight", inst.WidthRight, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)
		StageSetAssociationFieldToForm("TurnDirection", inst.TurnDirection, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDirectionTurnType](), probe.formStage)
		StageSetAssociationFieldToForm("SignalGap", inst.SignalGap, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("MinimumEnrouteAltitude", inst.MinimumEnrouteAltitude, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("MinimumCrossingAtEnd", inst.MinimumCrossingAtEnd, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("MinimumCrossingAtEndReference", inst.MinimumCrossingAtEndReference, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), probe.formStage)
		StageSetAssociationFieldToForm("MaximumCrossingAtEnd", inst.MaximumCrossingAtEnd, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("MaximumCrossingAtEndReference", inst.MaximumCrossingAtEndReference, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), probe.formStage)
		StageSetAssociationFieldToForm("NavigationType", inst.NavigationType, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRouteNavigationType](), probe.formStage)
		StageSetAssociationFieldToForm("RequiredNavigationPerformance", inst.RequiredNavigationPerformance, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRNPType](), probe.formStage)
		StageSetAssociationFieldToForm("DesignatorSuffix", inst.DesignatorSuffix, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRouteDesignatorSuffixType](), probe.formStage)
		StageSetAssociationFieldToForm("Start", inst.Start, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.EnRouteSegmentPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("RouteFormed", inst.RouteFormed, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RoutePropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("EvaluationArea", inst.EvaluationArea, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ObstacleAssessmentAreaPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("CurveExtent", inst.CurveExtent, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CurvePropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("End", inst.End, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.EnRouteSegmentPointPropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSlicePropertyTypes {
				if src.RouteSegmentTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSlicePropertyType", "RouteSegmentTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RouteSegmentType:
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
				if src.RouteSegment == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "RouteSegment", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RouteTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("RouteTimeSlice", inst.RouteTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RouteTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RouteTimeSliceType:
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
		StageSetAssociationFieldToForm("DesignatorPrefix", inst.DesignatorPrefix, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRouteDesignatorPrefixType](), probe.formStage)
		StageSetAssociationFieldToForm("DesignatorSecondLetter", inst.DesignatorSecondLetter, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRouteDesignatorLetterType](), probe.formStage)
		StageSetAssociationFieldToForm("DesignatorNumber", inst.DesignatorNumber, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoNumberType](), probe.formStage)
		StageSetAssociationFieldToForm("MultipleIdentifier", inst.MultipleIdentifier, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeUpperAlphaType](), probe.formStage)
		StageSetAssociationFieldToForm("LocationDesignator", inst.LocationDesignator, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextDesignatorType](), probe.formStage)
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRouteType](), probe.formStage)
		StageSetAssociationFieldToForm("FlightRule", inst.FlightRule, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlightRuleType](), probe.formStage)
		StageSetAssociationFieldToForm("InternationalUse", inst.InternationalUse, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRouteOriginType](), probe.formStage)
		StageSetAssociationFieldToForm("MilitaryUse", inst.MilitaryUse, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeMilitaryStatusType](), probe.formStage)
		StageSetAssociationFieldToForm("MilitaryTrainingType", inst.MilitaryTrainingType, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeMilitaryTrainingType](), probe.formStage)
		StageSetAssociationFieldToForm("UserOrganisation", inst.UserOrganisation, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.OrganisationAuthorityPropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteTimeSlicePropertyTypes {
				if src.RouteTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteTimeSlicePropertyType", "RouteTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RouteType:
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
				if src.Route == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "Route", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RulesProceduresTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("RulesProceduresTimeSlice", inst.RulesProceduresTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RulesProceduresTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RulesProceduresTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RulesProceduresType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RulesProceduresTimeSliceType:
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
		StageSetAssociationFieldToForm("Category", inst.Category, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRuleProcedureType](), probe.formStage)
		StageSetAssociationFieldToForm("Title", inst.Title, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRuleProcedureTitleType](), probe.formStage)
		StageSetAssociationFieldToForm("Content", inst.Content, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.XHTMLType](), probe.formStage)

		{
			// Slice of pointers: AffectedLocation
			div := (&form.FormDiv{Name: "AffectedLocation"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AffectedLocation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AffectedLocation",
				Label: "AffectedLocation",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AffectedArea
			div := (&form.FormDiv{Name: "AffectedArea"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AffectedArea {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AffectedArea",
				Label: "AffectedArea",
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.RulesProceduresTimeSlicePropertyTypes {
				if src.RulesProceduresTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RulesProceduresTimeSlicePropertyType", "RulesProceduresTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RulesProceduresType:
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
				if src.RulesProcedures == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "RulesProcedures", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RunwayCentrelinePointPropertyType:
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingPointTypes {
				if src.PointChoice_runwayPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingPointType", "PointChoice_runwayPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ChangeOverPointTimeSliceTypes {
				if src.Location_runwayPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ChangeOverPointTimeSliceType", "Location_runwayPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectFlightSegmentTypes {
				if src.End_runwayPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectFlightSegmentType", "End_runwayPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectFlightSegmentTypes {
				if src.Start_runwayPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectFlightSegmentType", "Start_runwayPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.EnRouteSegmentPointTypes {
				if src.PointChoice_runwayPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.EnRouteSegmentPointType", "PointChoice_runwayPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementTypes {
				if src.SignificantPointCondition_runwayPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionElementType", "SignificantPointCondition_runwayPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRoutingElementTypes {
				if src.PointElement_runwayPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRoutingElementType", "PointElement_runwayPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RoutePortionTypes {
				if src.Start_runwayPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RoutePortionType", "Start_runwayPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RoutePortionTypes {
				if src.IntermediatePoint_runwayPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RoutePortionType", "IntermediatePoint_runwayPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RoutePortionTypes {
				if src.End_runwayPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RoutePortionType", "End_runwayPoint", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.StandardInstrumentArrivalPropertyType:
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementTypes {
				if src.FlightCondition_standardInstrumentArrivalCondition == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionElementType", "FlightCondition_standardInstrumentArrivalCondition", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRoutingElementTypes {
				if src.Element_standardInstrumentArrivalElement == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRoutingElementType", "Element_standardInstrumentArrivalElement", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.StandardInstrumentDeparturePropertyType:
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementTypes {
				if src.FlightCondition_standardInstrumentDepartureCondition == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionElementType", "FlightCondition_standardInstrumentDepartureCondition", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRoutingElementTypes {
				if src.Element_standardInstrumentDepartureElement == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRoutingElementType", "Element_standardInstrumentDepartureElement", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingAnchorTypes {
				if src.Extent == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingAnchorType", "Extent", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAssessmentAreaTypes {
				if src.Surface == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAssessmentAreaType", "Surface", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PointReferenceTypes {
				if src.FixToleranceArea == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PointReferenceType", "FixToleranceArea", refNames, formGroup, probe.formStage)
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
	case *aixmfeatures.TelephoneContactPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("TelephoneContact", inst.TelephoneContact, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TelephoneContactType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ContactInformationTypes {
				for _, target := range src.PhoneFax {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ContactInformationType", "PhoneFax", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.TelephoneContactType:
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
		StageSetAssociationFieldToForm("Voice", inst.Voice, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextPhoneType](), probe.formStage)
		StageSetAssociationFieldToForm("Facsimile", inst.Facsimile, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextPhoneType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.TelephoneContactPropertyTypes {
				if src.TelephoneContact == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TelephoneContactPropertyType", "TelephoneContact", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.TimesheetPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Timesheet", inst.Timesheet, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TimesheetType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionCombinationTypes {
				for _, target := range src.TimeInterval {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionCombinationType", "TimeInterval", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OnlineContactTypes {
				for _, target := range src.TimeInterval {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OnlineContactType", "TimeInterval", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PostalAddressTypes {
				for _, target := range src.TimeInterval {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PostalAddressType", "TimeInterval", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteAvailabilityTypes {
				for _, target := range src.TimeInterval {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteAvailabilityType", "TimeInterval", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TelephoneContactTypes {
				for _, target := range src.TimeInterval {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TelephoneContactType", "TimeInterval", refNames, formGroup, probe.formStage)
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
	case *aixmfeatures.TouchDownLiftOffPropertyType:
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingPointTypes {
				if src.PointChoice_aimingPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingPointType", "PointChoice_aimingPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ChangeOverPointTimeSliceTypes {
				if src.Location_aimingPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ChangeOverPointTimeSliceType", "Location_aimingPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectFlightSegmentTypes {
				if src.End_aimingPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectFlightSegmentType", "End_aimingPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectFlightSegmentTypes {
				if src.Start_aimingPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectFlightSegmentType", "Start_aimingPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.EnRouteSegmentPointTypes {
				if src.PointChoice_aimingPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.EnRouteSegmentPointType", "PointChoice_aimingPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementTypes {
				if src.SignificantPointCondition_aimingPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionElementType", "SignificantPointCondition_aimingPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRoutingElementTypes {
				if src.PointElement_aimingPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRoutingElementType", "PointElement_aimingPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RoutePortionTypes {
				if src.Start_aimingPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RoutePortionType", "Start_aimingPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RoutePortionTypes {
				if src.IntermediatePoint_aimingPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RoutePortionType", "IntermediatePoint_aimingPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RoutePortionTypes {
				if src.End_aimingPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RoutePortionType", "End_aimingPoint", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstructionTypes {
				if src.TheVerticalStructure == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstructionType", "TheVerticalStructure", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeATCReportingType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingPointTypes {
				if src.ReportingATC == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingPointType", "ReportingATC", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.EnRouteSegmentPointTypes {
				if src.ReportingATC == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.EnRouteSegmentPointType", "ReportingATC", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeAerialRefuellingPointType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingPointTypes {
				if src.UsageType == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingPointType", "UsageType", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeAerialRefuellingPrefixType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSliceTypes {
				if src.DesignatorPrefix == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTimeSliceType", "DesignatorPrefix", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeAerialRefuellingType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
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
	case *aixmdatatypes.CodeAltitudeAdjustmentType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AltitudeAdjustmentTypes {
				if src.AltitudeAdjustmentType == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AltitudeAdjustmentType", "AltitudeAdjustmentType", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForAerialRefuellingTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForAerialRefuellingType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeCardinalDirectionType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSliceTypes {
				if src.DesignatorDirection == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTimeSliceType", "DesignatorDirection", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteAvailabilityTypes {
				if src.CardinalDirection == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteAvailabilityType", "CardinalDirection", refNames, formGroup, probe.formStage)
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
	case *aixmdatatypes.CodeComparisonType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRoutingElementTypes {
				if src.SpeedCriteria == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRoutingElementType", "SpeedCriteria", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeCourseType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingAnchorTypes {
				if src.OutboundCourseType == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingAnchorType", "OutboundCourseType", refNames, formGroup, probe.formStage)
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
	case *aixmdatatypes.CodeDirectionTurnType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingAnchorTypes {
				if src.TurnDirection == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingAnchorType", "TurnDirection", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSliceTypes {
				if src.ReverseDirectionTurn == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTimeSliceType", "ReverseDirectionTurn", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.TurnDirection == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "TurnDirection", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeDirectionType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteAvailabilityTypes {
				if src.Direction == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteAvailabilityType", "Direction", refNames, formGroup, probe.formStage)
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
	case *aixmdatatypes.CodeFlightOriginType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightCharacteristicTypes {
				if src.Origin == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightCharacteristicType", "Origin", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeFlightPurposeType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightCharacteristicTypes {
				if src.Purpose == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightCharacteristicType", "Purpose", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeFlightRestrictionDesignatorType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionTimeSliceTypes {
				if src.Designator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionTimeSliceType", "Designator", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeFlightRestrictionType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeFlightRuleType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightCharacteristicTypes {
				if src.Rule == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightCharacteristicType", "Rule", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteTimeSliceTypes {
				if src.FlightRule == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteTimeSliceType", "FlightRule", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeFlightStatusType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightCharacteristicTypes {
				if src.Status == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightCharacteristicType", "Status", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeFlightType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightCharacteristicTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightCharacteristicType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeFlowConditionOperationType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionCombinationTypes {
				if src.LogicalOperator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionCombinationType", "LogicalOperator", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeFreeFlightType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.EnRouteSegmentPointTypes {
				if src.RoleFreeFlight == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.EnRouteSegmentPointType", "RoleFreeFlight", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeLevelType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "Level", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeLocationQualifierType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionCircumstanceTypes {
				if src.RelationWithLocation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionCircumstanceType", "RelationWithLocation", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeMeteoConditionsType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MeteorologyTypes {
				if src.FlightConditions == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MeteorologyType", "FlightConditions", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeMilitaryRoutePointType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.EnRouteSegmentPointTypes {
				if src.RoleMilitaryTraining == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.EnRouteSegmentPointType", "RoleMilitaryTraining", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeMilitaryStatusType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightCharacteristicTypes {
				if src.Military == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightCharacteristicType", "Military", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteTimeSliceTypes {
				if src.MilitaryUse == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteTimeSliceType", "MilitaryUse", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeMilitaryTrainingType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteTimeSliceTypes {
				if src.MilitaryTrainingType == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteTimeSliceType", "MilitaryTrainingType", refNames, formGroup, probe.formStage)
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
	case *aixmdatatypes.CodeObstacleAssessmentSurfaceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAssessmentAreaTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAssessmentAreaType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeObstructionIdSurfaceZoneType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAssessmentAreaTypes {
				if src.SurfaceZone == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAssessmentAreaType", "SurfaceZone", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeRNPType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.RequiredNavigationPerformance == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "RequiredNavigationPerformance", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeRVSMPointRoleType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.EnRouteSegmentPointTypes {
				if src.RoleRVSM == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.EnRouteSegmentPointType", "RoleRVSM", refNames, formGroup, probe.formStage)
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
	case *aixmdatatypes.CodeReferenceRoleType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PointReferenceTypes {
				if src.Role == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PointReferenceType", "Role", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeRouteAvailabilityType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteAvailabilityTypes {
				if src.Status == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteAvailabilityType", "Status", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeRouteDesignatorLetterType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteTimeSliceTypes {
				if src.DesignatorSecondLetter == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteTimeSliceType", "DesignatorSecondLetter", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeRouteDesignatorPrefixType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteTimeSliceTypes {
				if src.DesignatorPrefix == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteTimeSliceType", "DesignatorPrefix", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeRouteDesignatorSuffixType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.DesignatorSuffix == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "DesignatorSuffix", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeRouteNavigationType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.NavigationType == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "NavigationType", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeRouteOriginType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteTimeSliceTypes {
				if src.InternationalUse == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteTimeSliceType", "InternationalUse", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeRouteSegmentPathType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.PathType == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "PathType", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeRouteType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeRuleProcedureTitleType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RulesProceduresTimeSliceTypes {
				if src.Title == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RulesProceduresTimeSliceType", "Title", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeRuleProcedureType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RulesProceduresTimeSliceTypes {
				if src.Category == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RulesProceduresTimeSliceType", "Category", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeSideType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstaclePlacementTypes {
				if src.ObstaclePlacement == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstaclePlacementType", "ObstaclePlacement", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeSpeedReferenceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRoutingElementTypes {
				if src.SpeedReference == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRoutingElementType", "SpeedReference", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeTACANChannelType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSliceTypes {
				if src.TankerChannel == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTimeSliceType", "TankerChannel", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSliceTypes {
				if src.ReceiverChannel == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTimeSliceType", "ReceiverChannel", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeTelecomNetworkType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OnlineContactTypes {
				if src.Network == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OnlineContactType", "Network", refNames, formGroup, probe.formStage)
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
	case *aixmdatatypes.CodeUpperAlphaType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteTimeSliceTypes {
				if src.MultipleIdentifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteTimeSliceType", "MultipleIdentifier", refNames, formGroup, probe.formStage)
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

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MeteorologyTypes {
				if src.VisibilityInterpretation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MeteorologyType", "VisibilityInterpretation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MeteorologyTypes {
				if src.RunwayVisualRangeInterpretation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MeteorologyType", "RunwayVisualRangeInterpretation", refNames, formGroup, probe.formStage)
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
	case *aixmdatatypes.CodeVerticalReferenceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingAnchorTypes {
				if src.RefuellingBaseLevelReference == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingAnchorType", "RefuellingBaseLevelReference", refNames, formGroup, probe.formStage)
		}

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionLevelTypes {
				if src.UpperLevelReference == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionLevelType", "UpperLevelReference", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionLevelTypes {
				if src.LowerLevelReference == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionLevelType", "LowerLevelReference", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.UpperLimitReference == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "UpperLimitReference", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.LowerLimitReference == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "LowerLimitReference", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.MinimumCrossingAtEndReference == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "MinimumCrossingAtEndReference", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.MaximumCrossingAtEndReference == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "MaximumCrossingAtEndReference", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingPointTypes {
				if src.FlyOver == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingPointType", "FlyOver", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingPointTypes {
				if src.Waypoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingPointType", "Waypoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingPointTypes {
				if src.RadarGuidance == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingPointType", "RadarGuidance", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSliceTypes {
				if src.HelicopterRoute == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTimeSliceType", "HelicopterRoute", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSliceTypes {
				if src.SpecialRefuelling == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTimeSliceType", "SpecialRefuelling", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSliceTypes {
				if src.BidirectionalUse == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTimeSliceType", "BidirectionalUse", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AltitudeAdjustmentTypes {
				if src.PrimaryAlternateMinimum == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AltitudeAdjustmentType", "PrimaryAlternateMinimum", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AltitudeAdjustmentTypes {
				if src.LocalRemoteCode == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AltitudeAdjustmentType", "LocalRemoteCode", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleUseTypes {
				if src.AlongCourseGuidance == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleUseType", "AlongCourseGuidance", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.EnRouteSegmentPointTypes {
				if src.FlyOver == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.EnRouteSegmentPointType", "FlyOver", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.EnRouteSegmentPointTypes {
				if src.Waypoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.EnRouteSegmentPointType", "Waypoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.EnRouteSegmentPointTypes {
				if src.RadarGuidance == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.EnRouteSegmentPointType", "RadarGuidance", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionCircumstanceTypes {
				if src.ReferenceLocation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionCircumstanceType", "ReferenceLocation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionRouteTypes {
				if src.PriorPermission == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionRouteType", "PriorPermission", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstructionTypes {
				if src.SurfacePenetration == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstructionType", "SurfacePenetration", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstructionTypes {
				if src.Controlling == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstructionType", "Controlling", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstructionTypes {
				if src.CloseIn == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstructionType", "CloseIn", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteDMETimeSliceTypes {
				if src.CriticalDME == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteDMETimeSliceType", "CriticalDME", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteDMETimeSliceTypes {
				if src.Satisfactory == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteDMETimeSliceType", "Satisfactory", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.SignalGap == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "SignalGap", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSliceTypes {
				if src.DesignatorNumber == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTimeSliceType", "DesignatorNumber", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSliceTypes {
				if src.RadarBeaconSetting == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTimeSliceType", "RadarBeaconSetting", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSliceTypes {
				if src.XbandRadarSetting == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTimeSliceType", "XbandRadarSetting", refNames, formGroup, probe.formStage)
		}

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

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAssessmentAreaTypes {
				if src.SectionNumber == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAssessmentAreaType", "SectionNumber", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteTimeSliceTypes {
				if src.DesignatorNumber == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteTimeSliceType", "DesignatorNumber", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.NoSequenceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingPointTypes {
				if src.Sequence == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingPointType", "Sequence", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementTypes {
				if src.Index == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionElementType", "Index", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRoutingElementTypes {
				if src.OrderNumber == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRoutingElementType", "OrderNumber", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.TextAddressType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OnlineContactTypes {
				if src.Linkage == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OnlineContactType", "Linkage", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OnlineContactTypes {
				if src.EMail == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OnlineContactType", "EMail", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PostalAddressTypes {
				if src.DeliveryPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PostalAddressType", "DeliveryPoint", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.TextDesignatorType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSliceTypes {
				if src.DesignatorSuffix == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTimeSliceType", "DesignatorSuffix", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteTimeSliceTypes {
				if src.LocationDesignator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteTimeSliceType", "LocationDesignator", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.TextInstructionType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionTimeSliceTypes {
				if src.Instruction == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionTimeSliceType", "Instruction", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.TextNameType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ContactInformationTypes {
				if src.Title == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ContactInformationType", "Title", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAssessmentAreaTypes {
				if src.SafetyRegulation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAssessmentAreaType", "SafetyRegulation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstaclePlacementTypes {
				if src.PointType == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstaclePlacementType", "PointType", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OnlineContactTypes {
				if src.Protocol == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OnlineContactType", "Protocol", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PostalAddressTypes {
				if src.City == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PostalAddressType", "City", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PostalAddressTypes {
				if src.AdministrativeArea == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PostalAddressType", "AdministrativeArea", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PostalAddressTypes {
				if src.PostalCode == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PostalAddressType", "PostalCode", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PostalAddressTypes {
				if src.Country == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PostalAddressType", "Country", refNames, formGroup, probe.formStage)
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
	case *aixmdatatypes.TextPhoneType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TelephoneContactTypes {
				if src.Voice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TelephoneContactType", "Voice", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TelephoneContactTypes {
				if src.Facsimile == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TelephoneContactType", "Facsimile", refNames, formGroup, probe.formStage)
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
	case *aixmdatatypes.ValAngleType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstructionTypes {
				if src.SlopePenetration == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstructionType", "SlopePenetration", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.ValBearingType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingAnchorTypes {
				if src.OutboundCourse == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingAnchorType", "OutboundCourse", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingAnchorTypes {
				if src.InboundCourse == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingAnchorType", "InboundCourse", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstaclePlacementTypes {
				if src.ObstacleBearing == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstaclePlacementType", "ObstacleBearing", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.TrueTrack == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "TrueTrack", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.MagneticTrack == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "MagneticTrack", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.ReverseTrueTrack == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "ReverseTrueTrack", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.ReverseMagneticTrack == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "ReverseMagneticTrack", refNames, formGroup, probe.formStage)
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

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PointReferenceTypes {
				if src.PriorFixTolerance == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PointReferenceType", "PriorFixTolerance", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PointReferenceTypes {
				if src.PostFixTolerance == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PointReferenceType", "PostFixTolerance", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.ValDistanceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Uom", inst.Uom, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingAnchorTypes {
				if src.LegSeparation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingAnchorType", "LegSeparation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingAnchorTypes {
				if src.LegLength == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingAnchorType", "LegLength", refNames, formGroup, probe.formStage)
		}

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.ChangeOverPointTimeSliceTypes {
				if src.Distance == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ChangeOverPointTimeSliceType", "Distance", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectFlightClassTypes {
				if src.ExceedLength == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectFlightClassType", "ExceedLength", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.EnRouteSegmentPointTypes {
				if src.TurnRadius == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.EnRouteSegmentPointType", "TurnRadius", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MeteorologyTypes {
				if src.Visibility == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MeteorologyType", "Visibility", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MeteorologyTypes {
				if src.RunwayVisualRange == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MeteorologyType", "RunwayVisualRange", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstaclePlacementTypes {
				if src.ObstacleDistance == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstaclePlacementType", "ObstacleDistance", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstructionTypes {
				if src.RequiredClearance == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstructionType", "RequiredClearance", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.Length == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "Length", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.WidthLeft == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "WidthLeft", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.WidthRight == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "WidthRight", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingAnchorTypes {
				if src.RefuellingBaseLevel == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingAnchorType", "RefuellingBaseLevel", refNames, formGroup, probe.formStage)
		}

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AltitudeAdjustmentTypes {
				if src.AltitudeAdjustment == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AltitudeAdjustmentType", "AltitudeAdjustment", refNames, formGroup, probe.formStage)
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

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionLevelTypes {
				if src.UpperLevel == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionLevelType", "UpperLevel", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionLevelTypes {
				if src.LowerLevel == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionLevelType", "LowerLevel", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAssessmentAreaTypes {
				if src.AssessedAltitude == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAssessmentAreaType", "AssessedAltitude", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAssessmentAreaTypes {
				if src.SlopeLowerAltitude == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAssessmentAreaType", "SlopeLowerAltitude", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstructionTypes {
				if src.MinimumAltitude == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstructionType", "MinimumAltitude", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.UpperLimit == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "UpperLimit", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.LowerLimit == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "LowerLimit", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.MinimumObstacleClearanceAltitude == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "MinimumObstacleClearanceAltitude", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.MinimumEnrouteAltitude == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "MinimumEnrouteAltitude", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.MinimumCrossingAtEnd == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "MinimumCrossingAtEnd", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.MaximumCrossingAtEnd == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "MaximumCrossingAtEnd", refNames, formGroup, probe.formStage)
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
	case *aixmdatatypes.ValSlopeType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAssessmentAreaTypes {
				if src.Slope == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAssessmentAreaType", "Slope", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAssessmentAreaTypes {
				if src.GradientLowHigh == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAssessmentAreaType", "GradientLowHigh", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.ValSpeedType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Uom", inst.Uom, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingAnchorTypes {
				if src.SpeedLimit == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingAnchorType", "SpeedLimit", refNames, formGroup, probe.formStage)
		}

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

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRoutingElementTypes {
				if src.Speed == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRoutingElementType", "Speed", refNames, formGroup, probe.formStage)
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
	case *aixmdatatypes.XHTMLType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RulesProceduresTimeSliceTypes {
				if src.Content == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RulesProceduresTimeSliceType", "Content", refNames, formGroup, probe.formStage)
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
				for _, target := range src.AbstractAerialRefuellingExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension", "AbstractAerialRefuellingExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_1s {
				for _, target := range src.AbstractAerialRefuellingAnchorExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_1", "AbstractAerialRefuellingAnchorExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_10s {
				for _, target := range src.AbstractPointReferenceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_10", "AbstractPointReferenceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_11s {
				for _, target := range src.AbstractLinguisticNoteExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_11", "AbstractLinguisticNoteExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_12s {
				for _, target := range src.AbstractNoteExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_12", "AbstractNoteExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_13s {
				for _, target := range src.AbstractChangeOverPointExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_13", "AbstractChangeOverPointExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_14s {
				for _, target := range src.AbstractRouteExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_14", "AbstractRouteExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_15s {
				for _, target := range src.AbstractPropertiesWithScheduleExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_15", "AbstractPropertiesWithScheduleExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_15s {
				for _, target := range src.AbstractRouteAvailabilityExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_15", "AbstractRouteAvailabilityExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_16s {
				for _, target := range src.AbstractRouteDMEExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_16", "AbstractRouteDMEExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_17s {
				for _, target := range src.AbstractRoutePortionExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_17", "AbstractRoutePortionExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_18s {
				for _, target := range src.AbstractRouteSegmentExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_18", "AbstractRouteSegmentExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_19s {
				for _, target := range src.AbstractAirspaceBorderCrossingExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_19", "AbstractAirspaceBorderCrossingExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_2s {
				for _, target := range src.AbstractSegmentPointExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_2", "AbstractSegmentPointExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_2s {
				for _, target := range src.AbstractAerialRefuellingPointExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_2", "AbstractAerialRefuellingPointExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_20s {
				for _, target := range src.AbstractDirectFlightExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_20", "AbstractDirectFlightExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_20s {
				for _, target := range src.AbstractDirectFlightClassExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_20", "AbstractDirectFlightClassExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_21s {
				for _, target := range src.AbstractDirectFlightExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_21", "AbstractDirectFlightExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_21s {
				for _, target := range src.AbstractDirectFlightSegmentExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_21", "AbstractDirectFlightSegmentExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_22s {
				for _, target := range src.AbstractFlightConditionCircumstanceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_22", "AbstractFlightConditionCircumstanceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_23s {
				for _, target := range src.AbstractPropertiesWithScheduleExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_23", "AbstractPropertiesWithScheduleExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_23s {
				for _, target := range src.AbstractFlightConditionCombinationExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_23", "AbstractFlightConditionCombinationExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_24s {
				for _, target := range src.AbstractFlightConditionElementExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_24", "AbstractFlightConditionElementExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_25s {
				for _, target := range src.AbstractFlightRestrictionExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_25", "AbstractFlightRestrictionExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_26s {
				for _, target := range src.AbstractFlightRestrictionLevelExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_26", "AbstractFlightRestrictionLevelExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_27s {
				for _, target := range src.AbstractFlightRestrictionRouteExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_27", "AbstractFlightRestrictionRouteExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_28s {
				for _, target := range src.AbstractFlightRoutingElementExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_28", "AbstractFlightRoutingElementExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_29s {
				for _, target := range src.AbstractRulesProceduresExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_29", "AbstractRulesProceduresExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_3s {
				for _, target := range src.AbstractAerialRefuellingTrackExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_3", "AbstractAerialRefuellingTrackExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_30s {
				for _, target := range src.AbstractContactInformationExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_30", "AbstractContactInformationExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_31s {
				for _, target := range src.AbstractPropertiesWithScheduleExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_31", "AbstractPropertiesWithScheduleExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_31s {
				for _, target := range src.AbstractOnlineContactExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_31", "AbstractOnlineContactExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_32s {
				for _, target := range src.AbstractPropertiesWithScheduleExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_32", "AbstractPropertiesWithScheduleExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_32s {
				for _, target := range src.AbstractPostalAddressExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_32", "AbstractPostalAddressExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_33s {
				for _, target := range src.AbstractPropertiesWithScheduleExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_33", "AbstractPropertiesWithScheduleExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_33s {
				for _, target := range src.AbstractTelephoneContactExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_33", "AbstractTelephoneContactExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_34s {
				for _, target := range src.AbstractAircraftCharacteristicExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_34", "AbstractAircraftCharacteristicExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_35s {
				for _, target := range src.AbstractFlightCharacteristicExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_35", "AbstractFlightCharacteristicExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_36s {
				for _, target := range src.AbstractAirspaceLayerExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_36", "AbstractAirspaceLayerExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_37s {
				for _, target := range src.AbstractMeteorologyExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_37", "AbstractMeteorologyExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_38s {
				for _, target := range src.AbstractTimesheetExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_38", "AbstractTimesheetExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_39s {
				for _, target := range src.AbstractAltitudeAdjustmentExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_39", "AbstractAltitudeAdjustmentExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_4s {
				for _, target := range src.AbstractAuthorityForAerialRefuellingExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_4", "AbstractAuthorityForAerialRefuellingExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_40s {
				for _, target := range src.AbstractObstacleAssessmentAreaExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_40", "AbstractObstacleAssessmentAreaExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_41s {
				for _, target := range src.AbstractObstaclePlacementExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_41", "AbstractObstaclePlacementExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_42s {
				for _, target := range src.AbstractObstructionExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_42", "AbstractObstructionExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_5s {
				for _, target := range src.AbstractElevatedCurveExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_5", "AbstractElevatedCurveExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_6s {
				for _, target := range src.AbstractElevatedPointExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_6", "AbstractElevatedPointExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_7s {
				for _, target := range src.AbstractElevatedSurfaceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_7", "AbstractElevatedSurfaceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_8s {
				for _, target := range src.AbstractAngleUseExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_8", "AbstractAngleUseExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_9s {
				for _, target := range src.AbstractSegmentPointExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_9", "AbstractSegmentPointExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_9s {
				for _, target := range src.AbstractEnRouteSegmentPointExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_9", "AbstractEnRouteSegmentPointExtension", refNames, formGroup, probe.formStage)
		}
	case *aixmabstractgmlobjecttypes.FeatureMetadataPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceBorderCrossingTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceBorderCrossingType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ChangeOverPointTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ChangeOverPointType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteDMETypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteDMEType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RulesProceduresTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RulesProceduresType", "FeatureMetadata", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceBorderCrossingTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceBorderCrossingTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ChangeOverPointTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ChangeOverPointTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteDMETimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteDMETimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RulesProceduresTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RulesProceduresTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceBorderCrossingTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceBorderCrossingType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ChangeOverPointTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ChangeOverPointType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteDMETypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteDMEType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RulesProceduresTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RulesProceduresType", "BoundedBy", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractDirectFlightTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractDirectFlightType", "NameXSD", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractSegmentPointTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractSegmentPointType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingAnchorTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingAnchorType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingPointTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingPointType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTrackTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTrackType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingType", "NameXSD", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceBorderCrossingTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceBorderCrossingTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceBorderCrossingTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceBorderCrossingType", "NameXSD", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AltitudeAdjustmentTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AltitudeAdjustmentType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleUseTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleUseType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForAerialRefuellingTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForAerialRefuellingType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ChangeOverPointTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ChangeOverPointTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ChangeOverPointTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ChangeOverPointType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ContactInformationTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ContactInformationType", "NameXSD", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectFlightClassTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectFlightClassType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectFlightSegmentTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectFlightSegmentType", "NameXSD", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.EnRouteSegmentPointTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.EnRouteSegmentPointType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightCharacteristicTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightCharacteristicType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionCircumstanceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionCircumstanceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionCombinationTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionCombinationType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionElementType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionLevelTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionLevelType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionRouteTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionRouteType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRoutingElementTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRoutingElementType", "NameXSD", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.MeteorologyTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MeteorologyType", "NameXSD", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAssessmentAreaTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAssessmentAreaType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstaclePlacementTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstaclePlacementType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstructionTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstructionType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OnlineContactTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OnlineContactType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PointReferenceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PointReferenceType", "NameXSD", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.PostalAddressTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PostalAddressType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteAvailabilityTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteAvailabilityType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteDMETimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteDMETimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteDMETypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteDMEType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RoutePortionTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RoutePortionType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RulesProceduresTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RulesProceduresTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RulesProceduresTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RulesProceduresType", "NameXSD", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.TelephoneContactTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TelephoneContactType", "NameXSD", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractDirectFlightTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractDirectFlightType", "Identifier", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractSegmentPointTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractSegmentPointType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingAnchorTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingAnchorType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingPointTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingPointType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTrackTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTrackType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingType", "Identifier", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceBorderCrossingTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceBorderCrossingTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceBorderCrossingTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceBorderCrossingType", "Identifier", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AltitudeAdjustmentTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AltitudeAdjustmentType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleUseTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleUseType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForAerialRefuellingTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForAerialRefuellingType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ChangeOverPointTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ChangeOverPointTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ChangeOverPointTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ChangeOverPointType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ContactInformationTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ContactInformationType", "Identifier", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectFlightClassTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectFlightClassType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectFlightSegmentTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectFlightSegmentType", "Identifier", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.EnRouteSegmentPointTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.EnRouteSegmentPointType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightCharacteristicTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightCharacteristicType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionCircumstanceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionCircumstanceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionCombinationTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionCombinationType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightConditionElementType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionLevelTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionLevelType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionRouteTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionRouteType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRoutingElementTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRoutingElementType", "Identifier", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.MeteorologyTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MeteorologyType", "Identifier", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstacleAssessmentAreaTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstacleAssessmentAreaType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstaclePlacementTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstaclePlacementType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ObstructionTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ObstructionType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OnlineContactTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OnlineContactType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PointReferenceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PointReferenceType", "Identifier", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.PostalAddressTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PostalAddressType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteAvailabilityTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteAvailabilityType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteDMETimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteDMETimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteDMETypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteDMEType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RoutePortionTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RoutePortionType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RulesProceduresTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RulesProceduresTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RulesProceduresTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RulesProceduresType", "Identifier", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.TelephoneContactTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TelephoneContactType", "Identifier", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AerialRefuellingTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceBorderCrossingTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceBorderCrossingTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirspaceBorderCrossingTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirspaceBorderCrossingTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ChangeOverPointTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ChangeOverPointTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ChangeOverPointTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ChangeOverPointTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FlightRestrictionTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteDMETimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteDMETimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteDMETimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteDMETimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteSegmentTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RouteTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RouteTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RulesProceduresTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RulesProceduresTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RulesProceduresTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RulesProceduresTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
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
