// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/xlink"
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
		StageSetAssociationFieldToForm("Navaid", inst.Navaid, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidType](), probe.formStage)
		StageSetAssociationFieldToForm("VOR", inst.VOR, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.VORType](), probe.formStage)
		StageSetAssociationFieldToForm("DME", inst.DME, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DMEType](), probe.formStage)
		StageSetAssociationFieldToForm("TACAN", inst.TACAN, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TACANType](), probe.formStage)
		StageSetAssociationFieldToForm("NDB", inst.NDB, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NDBType](), probe.formStage)
		StageSetAssociationFieldToForm("Localizer", inst.Localizer, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.LocalizerType](), probe.formStage)
		StageSetAssociationFieldToForm("Glidepath", inst.Glidepath, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.GlidepathType](), probe.formStage)
		StageSetAssociationFieldToForm("MarkerBeacon", inst.MarkerBeacon, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.MarkerBeaconType](), probe.formStage)
		StageSetAssociationFieldToForm("SDF", inst.SDF, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SDFType](), probe.formStage)
		StageSetAssociationFieldToForm("DirectionFinder", inst.DirectionFinder, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DirectionFinderType](), probe.formStage)
		StageSetAssociationFieldToForm("Azimuth", inst.Azimuth, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AzimuthType](), probe.formStage)
		StageSetAssociationFieldToForm("Elevation", inst.Elevation, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevationType](), probe.formStage)
		StageSetAssociationFieldToForm("CheckpointVOR", inst.CheckpointVOR, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CheckpointVORType](), probe.formStage)
		StageSetAssociationFieldToForm("CheckpointINS", inst.CheckpointINS, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CheckpointINSType](), probe.formStage)
		StageSetAssociationFieldToForm("SpecialNavigationSystem", inst.SpecialNavigationSystem, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SpecialNavigationSystemType](), probe.formStage)
		StageSetAssociationFieldToForm("SpecialNavigationStation", inst.SpecialNavigationStation, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SpecialNavigationStationType](), probe.formStage)
		StageSetAssociationFieldToForm("DesignatedPoint", inst.DesignatedPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointType](), probe.formStage)
		StageSetAssociationFieldToForm("SignificantPointInAirspace", inst.SignificantPointInAirspace, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SignificantPointInAirspaceType](), probe.formStage)
		StageSetAssociationFieldToForm("AngleIndication", inst.AngleIndication, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AngleIndicationType](), probe.formStage)
		StageSetAssociationFieldToForm("DistanceIndication", inst.DistanceIndication, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DistanceIndicationType](), probe.formStage)
		StageSetAssociationFieldToForm("RadarSystem", inst.RadarSystem, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RadarSystemType](), probe.formStage)
		StageSetAssociationFieldToForm("PrimarySurveillanceRadar", inst.PrimarySurveillanceRadar, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PrimarySurveillanceRadarType](), probe.formStage)
		StageSetAssociationFieldToForm("SecondarySurveillanceRadar", inst.SecondarySurveillanceRadar, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SecondarySurveillanceRadarType](), probe.formStage)
		StageSetAssociationFieldToForm("PrecisionApproachRadar", inst.PrecisionApproachRadar, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PrecisionApproachRadarType](), probe.formStage)

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
	case *aixmfeatures.A_extension_1:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractAuthorityForNavaidEquipmentExtension
			div := (&form.FormDiv{Name: "AbstractAuthorityForNavaidEquipmentExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAuthorityForNavaidEquipmentExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAuthorityForNavaidEquipmentExtension",
				Label: "AbstractAuthorityForNavaidEquipmentExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForNavaidEquipmentTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForNavaidEquipmentType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_10:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractGlidepathExtension
			div := (&form.FormDiv{Name: "AbstractGlidepathExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractGlidepathExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractGlidepathExtension",
				Label: "AbstractGlidepathExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractNavaidEquipmentExtension
			div := (&form.FormDiv{Name: "AbstractNavaidEquipmentExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractNavaidEquipmentExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractNavaidEquipmentExtension",
				Label: "AbstractNavaidEquipmentExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_11:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractLocalizerExtension
			div := (&form.FormDiv{Name: "AbstractLocalizerExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractLocalizerExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractLocalizerExtension",
				Label: "AbstractLocalizerExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractNavaidEquipmentExtension
			div := (&form.FormDiv{Name: "AbstractNavaidEquipmentExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractNavaidEquipmentExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractNavaidEquipmentExtension",
				Label: "AbstractNavaidEquipmentExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_12:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractMarkerBeaconExtension
			div := (&form.FormDiv{Name: "AbstractMarkerBeaconExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractMarkerBeaconExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractMarkerBeaconExtension",
				Label: "AbstractMarkerBeaconExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractNavaidEquipmentExtension
			div := (&form.FormDiv{Name: "AbstractNavaidEquipmentExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractNavaidEquipmentExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractNavaidEquipmentExtension",
				Label: "AbstractNavaidEquipmentExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MarkerBeaconTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_13:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractNavaidExtension
			div := (&form.FormDiv{Name: "AbstractNavaidExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractNavaidExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractNavaidExtension",
				Label: "AbstractNavaidExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_14:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractNavaidComponentExtension
			div := (&form.FormDiv{Name: "AbstractNavaidComponentExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractNavaidComponentExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractNavaidComponentExtension",
				Label: "AbstractNavaidComponentExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidComponentTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidComponentType", "Extension", refNames, formGroup, probe.formStage)
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
			// Slice of pointers: AbstractNavaidEquipmentMonitoringExtension
			div := (&form.FormDiv{Name: "AbstractNavaidEquipmentMonitoringExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractNavaidEquipmentMonitoringExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractNavaidEquipmentMonitoringExtension",
				Label: "AbstractNavaidEquipmentMonitoringExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidEquipmentMonitoringTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidEquipmentMonitoringType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_16:
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
			// Slice of pointers: AbstractNavaidOperationalStatusExtension
			div := (&form.FormDiv{Name: "AbstractNavaidOperationalStatusExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractNavaidOperationalStatusExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractNavaidOperationalStatusExtension",
				Label: "AbstractNavaidOperationalStatusExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidOperationalStatusTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidOperationalStatusType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_17:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractNDBExtension
			div := (&form.FormDiv{Name: "AbstractNDBExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractNDBExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractNDBExtension",
				Label: "AbstractNDBExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractNavaidEquipmentExtension
			div := (&form.FormDiv{Name: "AbstractNavaidEquipmentExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractNavaidEquipmentExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractNavaidEquipmentExtension",
				Label: "AbstractNavaidEquipmentExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NDBTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NDBTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_18:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractSDFExtension
			div := (&form.FormDiv{Name: "AbstractSDFExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractSDFExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractSDFExtension",
				Label: "AbstractSDFExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractNavaidEquipmentExtension
			div := (&form.FormDiv{Name: "AbstractNavaidEquipmentExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractNavaidEquipmentExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractNavaidEquipmentExtension",
				Label: "AbstractNavaidEquipmentExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SDFTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SDFTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_19:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractSpecialNavigationStationExtension
			div := (&form.FormDiv{Name: "AbstractSpecialNavigationStationExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractSpecialNavigationStationExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractSpecialNavigationStationExtension",
				Label: "AbstractSpecialNavigationStationExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationStationTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_2:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractAuthorityForSpecialNavigationStationExtension
			div := (&form.FormDiv{Name: "AbstractAuthorityForSpecialNavigationStationExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAuthorityForSpecialNavigationStationExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAuthorityForSpecialNavigationStationExtension",
				Label: "AbstractAuthorityForSpecialNavigationStationExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForSpecialNavigationStationTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForSpecialNavigationStationType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_20:
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
			// Slice of pointers: AbstractSpecialNavigationStationStatusExtension
			div := (&form.FormDiv{Name: "AbstractSpecialNavigationStationStatusExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractSpecialNavigationStationStatusExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractSpecialNavigationStationStatusExtension",
				Label: "AbstractSpecialNavigationStationStatusExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationStatusTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationStationStatusType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_21:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractSpecialNavigationSystemExtension
			div := (&form.FormDiv{Name: "AbstractSpecialNavigationSystemExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractSpecialNavigationSystemExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractSpecialNavigationSystemExtension",
				Label: "AbstractSpecialNavigationSystemExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationSystemTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationSystemTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_22:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractTACANExtension
			div := (&form.FormDiv{Name: "AbstractTACANExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractTACANExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractTACANExtension",
				Label: "AbstractTACANExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractNavaidEquipmentExtension
			div := (&form.FormDiv{Name: "AbstractNavaidEquipmentExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractNavaidEquipmentExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractNavaidEquipmentExtension",
				Label: "AbstractNavaidEquipmentExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TACANTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TACANTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_23:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractVORExtension
			div := (&form.FormDiv{Name: "AbstractVORExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractVORExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractVORExtension",
				Label: "AbstractVORExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractNavaidEquipmentExtension
			div := (&form.FormDiv{Name: "AbstractNavaidEquipmentExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractNavaidEquipmentExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractNavaidEquipmentExtension",
				Label: "AbstractNavaidEquipmentExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VORTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VORTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_24:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractAngleIndicationExtension
			div := (&form.FormDiv{Name: "AbstractAngleIndicationExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAngleIndicationExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAngleIndicationExtension",
				Label: "AbstractAngleIndicationExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleIndicationTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_25:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractDistanceIndicationExtension
			div := (&form.FormDiv{Name: "AbstractDistanceIndicationExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractDistanceIndicationExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractDistanceIndicationExtension",
				Label: "AbstractDistanceIndicationExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DistanceIndicationTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DistanceIndicationTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_26:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractDesignatedPointExtension
			div := (&form.FormDiv{Name: "AbstractDesignatedPointExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractDesignatedPointExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractDesignatedPointExtension",
				Label: "AbstractDesignatedPointExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DesignatedPointTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DesignatedPointTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_27:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractSignificantPointInAirspaceExtension
			div := (&form.FormDiv{Name: "AbstractSignificantPointInAirspaceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractSignificantPointInAirspaceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractSignificantPointInAirspaceExtension",
				Label: "AbstractSignificantPointInAirspaceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SignificantPointInAirspaceTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SignificantPointInAirspaceTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_28:
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
	case *aixmfeatures.A_extension_29:
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
	case *aixmfeatures.A_extension_3:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractAuthorityForSpecialNavigationSystemExtension
			div := (&form.FormDiv{Name: "AbstractAuthorityForSpecialNavigationSystemExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAuthorityForSpecialNavigationSystemExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAuthorityForSpecialNavigationSystemExtension",
				Label: "AbstractAuthorityForSpecialNavigationSystemExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForSpecialNavigationSystemTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForSpecialNavigationSystemType", "Extension", refNames, formGroup, probe.formStage)
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
	case *aixmfeatures.A_extension_35:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractPrecisionApproachRadarExtension
			div := (&form.FormDiv{Name: "AbstractPrecisionApproachRadarExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractPrecisionApproachRadarExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractPrecisionApproachRadarExtension",
				Label: "AbstractPrecisionApproachRadarExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractRadarEquipmentExtension
			div := (&form.FormDiv{Name: "AbstractRadarEquipmentExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractRadarEquipmentExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractRadarEquipmentExtension",
				Label: "AbstractRadarEquipmentExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_36:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractPrimarySurveillanceRadarExtension
			div := (&form.FormDiv{Name: "AbstractPrimarySurveillanceRadarExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractPrimarySurveillanceRadarExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractPrimarySurveillanceRadarExtension",
				Label: "AbstractPrimarySurveillanceRadarExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractSurveillanceRadarExtension
			div := (&form.FormDiv{Name: "AbstractSurveillanceRadarExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractSurveillanceRadarExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractSurveillanceRadarExtension",
				Label: "AbstractSurveillanceRadarExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractRadarEquipmentExtension
			div := (&form.FormDiv{Name: "AbstractRadarEquipmentExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractRadarEquipmentExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractRadarEquipmentExtension",
				Label: "AbstractRadarEquipmentExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_37:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractRadarComponentExtension
			div := (&form.FormDiv{Name: "AbstractRadarComponentExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractRadarComponentExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractRadarComponentExtension",
				Label: "AbstractRadarComponentExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarComponentTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarComponentType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_38:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractRadarSystemExtension
			div := (&form.FormDiv{Name: "AbstractRadarSystemExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractRadarSystemExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractRadarSystemExtension",
				Label: "AbstractRadarSystemExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarSystemTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarSystemTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_39:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractReflectorExtension
			div := (&form.FormDiv{Name: "AbstractReflectorExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractReflectorExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractReflectorExtension",
				Label: "AbstractReflectorExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ReflectorTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ReflectorType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_4:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractAzimuthExtension
			div := (&form.FormDiv{Name: "AbstractAzimuthExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAzimuthExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAzimuthExtension",
				Label: "AbstractAzimuthExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractNavaidEquipmentExtension
			div := (&form.FormDiv{Name: "AbstractNavaidEquipmentExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractNavaidEquipmentExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractNavaidEquipmentExtension",
				Label: "AbstractNavaidEquipmentExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_40:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractSecondarySurveillanceRadarExtension
			div := (&form.FormDiv{Name: "AbstractSecondarySurveillanceRadarExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractSecondarySurveillanceRadarExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractSecondarySurveillanceRadarExtension",
				Label: "AbstractSecondarySurveillanceRadarExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractSurveillanceRadarExtension
			div := (&form.FormDiv{Name: "AbstractSurveillanceRadarExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractSurveillanceRadarExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractSurveillanceRadarExtension",
				Label: "AbstractSurveillanceRadarExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractRadarEquipmentExtension
			div := (&form.FormDiv{Name: "AbstractRadarEquipmentExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractRadarEquipmentExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractRadarEquipmentExtension",
				Label: "AbstractRadarEquipmentExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_41:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractSurveillanceGroundStationExtension
			div := (&form.FormDiv{Name: "AbstractSurveillanceGroundStationExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractSurveillanceGroundStationExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractSurveillanceGroundStationExtension",
				Label: "AbstractSurveillanceGroundStationExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SurveillanceGroundStationTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SurveillanceGroundStationType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_5:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractCheckpointINSExtension
			div := (&form.FormDiv{Name: "AbstractCheckpointINSExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractCheckpointINSExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractCheckpointINSExtension",
				Label: "AbstractCheckpointINSExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractNavigationSystemCheckpointExtension
			div := (&form.FormDiv{Name: "AbstractNavigationSystemCheckpointExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractNavigationSystemCheckpointExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractNavigationSystemCheckpointExtension",
				Label: "AbstractNavigationSystemCheckpointExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointINSTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointINSTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_6:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractCheckpointVORExtension
			div := (&form.FormDiv{Name: "AbstractCheckpointVORExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractCheckpointVORExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractCheckpointVORExtension",
				Label: "AbstractCheckpointVORExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractNavigationSystemCheckpointExtension
			div := (&form.FormDiv{Name: "AbstractNavigationSystemCheckpointExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractNavigationSystemCheckpointExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractNavigationSystemCheckpointExtension",
				Label: "AbstractNavigationSystemCheckpointExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointVORTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointVORTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_7:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractDirectionFinderExtension
			div := (&form.FormDiv{Name: "AbstractDirectionFinderExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractDirectionFinderExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractDirectionFinderExtension",
				Label: "AbstractDirectionFinderExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractNavaidEquipmentExtension
			div := (&form.FormDiv{Name: "AbstractNavaidEquipmentExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractNavaidEquipmentExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractNavaidEquipmentExtension",
				Label: "AbstractNavaidEquipmentExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectionFinderTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_8:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractDMEExtension
			div := (&form.FormDiv{Name: "AbstractDMEExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractDMEExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractDMEExtension",
				Label: "AbstractDMEExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractNavaidEquipmentExtension
			div := (&form.FormDiv{Name: "AbstractNavaidEquipmentExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractNavaidEquipmentExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractNavaidEquipmentExtension",
				Label: "AbstractNavaidEquipmentExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DMETimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DMETimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_9:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractElevationExtension
			div := (&form.FormDiv{Name: "AbstractElevationExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractElevationExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractElevationExtension",
				Label: "AbstractElevationExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractNavaidEquipmentExtension
			div := (&form.FormDiv{Name: "AbstractNavaidEquipmentExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractNavaidEquipmentExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractNavaidEquipmentExtension",
				Label: "AbstractNavaidEquipmentExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevationTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevationTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AbstractNavaidEquipmentType:
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
	case *aixmfeatures.AbstractNavigationSystemCheckpointType:
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
	case *aixmfeatures.AbstractRadarEquipmentType:
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
	case *aixmfeatures.AbstractSurveillanceRadarType:
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTimeSliceTypes {
				if src.PointChoice_airportReferencePoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleIndicationTimeSliceType", "PointChoice_airportReferencePoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointINSTimeSliceTypes {
				if src.AirportHeliport == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointINSTimeSliceType", "AirportHeliport", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointVORTimeSliceTypes {
				if src.AirportHeliport == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointVORTimeSliceType", "AirportHeliport", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DesignatedPointTimeSliceTypes {
				if src.AirportHeliport == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DesignatedPointTimeSliceType", "AirportHeliport", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DistanceIndicationTimeSliceTypes {
				if src.PointChoice_airportReferencePoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DistanceIndicationTimeSliceType", "PointChoice_airportReferencePoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidTimeSliceTypes {
				for _, target := range src.ServedAirport {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidTimeSliceType", "ServedAirport", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarSystemTimeSliceTypes {
				if src.AirportHeliport == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarSystemTimeSliceType", "AirportHeliport", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SignificantPointInAirspaceTimeSliceTypes {
				if src.Location_airportReferencePoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SignificantPointInAirspaceTimeSliceType", "Location_airportReferencePoint", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.SignificantPointInAirspaceTimeSliceTypes {
				if src.ContainingAirspace == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SignificantPointInAirspaceTimeSliceType", "ContainingAirspace", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AngleIndicationTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AngleIndicationTimeSlice", inst.AngleIndicationTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AngleIndicationTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleIndicationType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AngleIndicationTimeSliceType:
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
		StageSetAssociationFieldToForm("Angle", inst.Angle, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), probe.formStage)
		StageSetAssociationFieldToForm("AngleType", inst.AngleType, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeBearingType](), probe.formStage)
		StageSetAssociationFieldToForm("IndicationDirection", inst.IndicationDirection, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDirectionReferenceType](), probe.formStage)
		StageSetAssociationFieldToForm("TrueAngle", inst.TrueAngle, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), probe.formStage)
		StageSetAssociationFieldToForm("CardinalDirection", inst.CardinalDirection, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCardinalDirectionType](), probe.formStage)
		StageSetAssociationFieldToForm("MinimumReceptionAltitude", inst.MinimumReceptionAltitude, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("Fix", inst.Fix, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PointChoice_fixDesignatedPoint", inst.PointChoice_fixDesignatedPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PointChoice_navaidSystem", inst.PointChoice_navaidSystem, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PointChoice_position", inst.PointChoice_position, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PointChoice_runwayPoint", inst.PointChoice_runwayPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PointChoice_aimingPoint", inst.PointChoice_aimingPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PointChoice_airportReferencePoint", inst.PointChoice_airportReferencePoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTimeSlicePropertyTypes {
				if src.AngleIndicationTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleIndicationTimeSlicePropertyType", "AngleIndicationTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AngleIndicationType:
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
				if src.AngleIndication == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "AngleIndication", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AuthorityForNavaidEquipmentPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AuthorityForNavaidEquipment", inst.AuthorityForNavaidEquipment, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AuthorityForNavaidEquipmentType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSliceTypes {
				for _, target := range src.Authority {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthTimeSliceType", "Authority", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DMETimeSliceTypes {
				for _, target := range src.Authority {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DMETimeSliceType", "Authority", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTimeSliceTypes {
				for _, target := range src.Authority {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectionFinderTimeSliceType", "Authority", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevationTimeSliceTypes {
				for _, target := range src.Authority {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevationTimeSliceType", "Authority", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTimeSliceTypes {
				for _, target := range src.Authority {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathTimeSliceType", "Authority", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSliceTypes {
				for _, target := range src.Authority {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerTimeSliceType", "Authority", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTimeSliceTypes {
				for _, target := range src.Authority {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MarkerBeaconTimeSliceType", "Authority", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NDBTimeSliceTypes {
				for _, target := range src.Authority {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NDBTimeSliceType", "Authority", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SDFTimeSliceTypes {
				for _, target := range src.Authority {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SDFTimeSliceType", "Authority", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TACANTimeSliceTypes {
				for _, target := range src.Authority {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TACANTimeSliceType", "Authority", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VORTimeSliceTypes {
				for _, target := range src.Authority {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VORTimeSliceType", "Authority", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AuthorityForNavaidEquipmentType:
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
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAuthorityRoleType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForNavaidEquipmentPropertyTypes {
				if src.AuthorityForNavaidEquipment == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForNavaidEquipmentPropertyType", "AuthorityForNavaidEquipment", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AuthorityForSpecialNavigationStationPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AuthorityForSpecialNavigationStation", inst.AuthorityForSpecialNavigationStation, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AuthorityForSpecialNavigationStationType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationTimeSliceTypes {
				if src.ResponsibleOrganisation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationStationTimeSliceType", "ResponsibleOrganisation", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AuthorityForSpecialNavigationStationType:
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
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAuthorityRoleType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForSpecialNavigationStationPropertyTypes {
				if src.AuthorityForSpecialNavigationStation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForSpecialNavigationStationPropertyType", "AuthorityForSpecialNavigationStation", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AuthorityForSpecialNavigationSystemPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AuthorityForSpecialNavigationSystem", inst.AuthorityForSpecialNavigationSystem, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AuthorityForSpecialNavigationSystemType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationSystemTimeSliceTypes {
				if src.ResponsibleOrganisation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationSystemTimeSliceType", "ResponsibleOrganisation", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AuthorityForSpecialNavigationSystemType:
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
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAuthorityRoleType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForSpecialNavigationSystemPropertyTypes {
				if src.AuthorityForSpecialNavigationSystem == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForSpecialNavigationSystemPropertyType", "AuthorityForSpecialNavigationSystem", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AzimuthTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AzimuthTimeSlice", inst.AzimuthTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AzimuthTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AzimuthTimeSliceType:
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
		StageSetAssociationFieldToForm("Designator", inst.Designator, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavaidDesignatorType](), probe.formStage)
		StageSetAssociationFieldToForm("EmissionClass", inst.EmissionClass, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadioEmissionType](), probe.formStage)
		StageSetAssociationFieldToForm("Mobile", inst.Mobile, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticVariation", inst.MagneticVariation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticVariationAccuracy", inst.MagneticVariationAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("DateMagneticVariation", inst.DateMagneticVariation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), probe.formStage)
		StageSetAssociationFieldToForm("FlightChecked", inst.FlightChecked, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Location", inst.Location, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)

		{
			// Slice of pointers: Authority
			div := (&form.FormDiv{Name: "Authority"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Authority {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Authority",
				Label: "Authority",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Monitoring
			div := (&form.FormDiv{Name: "Monitoring"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Monitoring {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Monitoring",
				Label: "Monitoring",
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
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeMLSAzimuthType](), probe.formStage)
		StageSetAssociationFieldToForm("TrueBearing", inst.TrueBearing, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), probe.formStage)
		StageSetAssociationFieldToForm("TrueBearingAccuracy", inst.TrueBearingAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticBearing", inst.MagneticBearing, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), probe.formStage)
		StageSetAssociationFieldToForm("AngleProportionalLeft", inst.AngleProportionalLeft, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("AngleProportionalRight", inst.AngleProportionalRight, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("AngleCoverLeft", inst.AngleCoverLeft, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("AngleCoverRight", inst.AngleCoverRight, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("Channel", inst.Channel, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeMLSChannelType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSlicePropertyTypes {
				if src.AzimuthTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthTimeSlicePropertyType", "AzimuthTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AzimuthType:
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
				if src.Azimuth == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "Azimuth", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.CheckpointINSTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("CheckpointINSTimeSlice", inst.CheckpointINSTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CheckpointINSTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointINSTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointINSType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.CheckpointINSTimeSliceType:
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
		StageSetAssociationFieldToForm("Category", inst.Category, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCheckpointCategoryType](), probe.formStage)
		StageSetAssociationFieldToForm("UpperLimit", inst.UpperLimit, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("UpperLimitReference", inst.UpperLimitReference, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), probe.formStage)
		StageSetAssociationFieldToForm("LowerLimit", inst.LowerLimit, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("LowerLimitReference", inst.LowerLimitReference, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), probe.formStage)
		StageSetAssociationFieldToForm("AltitudeInterpretation", inst.AltitudeInterpretation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAltitudeUseType](), probe.formStage)
		StageSetAssociationFieldToForm("Distance", inst.Distance, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)
		StageSetAssociationFieldToForm("Angle", inst.Angle, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), probe.formStage)
		StageSetAssociationFieldToForm("Position", inst.Position, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("AirportHeliport", inst.AirportHeliport, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointINSTimeSlicePropertyTypes {
				if src.CheckpointINSTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointINSTimeSlicePropertyType", "CheckpointINSTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.CheckpointINSType:
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
				if src.CheckpointINS == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "CheckpointINS", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.CheckpointVORTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("CheckpointVORTimeSlice", inst.CheckpointVORTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CheckpointVORTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointVORTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointVORType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.CheckpointVORTimeSliceType:
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
		StageSetAssociationFieldToForm("Category", inst.Category, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCheckpointCategoryType](), probe.formStage)
		StageSetAssociationFieldToForm("UpperLimit", inst.UpperLimit, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("UpperLimitReference", inst.UpperLimitReference, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), probe.formStage)
		StageSetAssociationFieldToForm("LowerLimit", inst.LowerLimit, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("LowerLimitReference", inst.LowerLimitReference, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVerticalReferenceType](), probe.formStage)
		StageSetAssociationFieldToForm("AltitudeInterpretation", inst.AltitudeInterpretation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAltitudeUseType](), probe.formStage)
		StageSetAssociationFieldToForm("Distance", inst.Distance, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)
		StageSetAssociationFieldToForm("Angle", inst.Angle, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), probe.formStage)
		StageSetAssociationFieldToForm("Position", inst.Position, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("AirportHeliport", inst.AirportHeliport, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), probe.formStage)

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
		StageSetAssociationFieldToForm("CheckPointFacility", inst.CheckPointFacility, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.VORPropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointVORTimeSlicePropertyTypes {
				if src.CheckpointVORTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointVORTimeSlicePropertyType", "CheckpointVORTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.CheckpointVORType:
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
				if src.CheckpointVOR == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "CheckpointVOR", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ContactInformationPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ContactInformation", inst.ContactInformation, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ContactInformationType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSliceTypes {
				for _, target := range src.Contact {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarTimeSliceType", "Contact", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				for _, target := range src.Contact {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "Contact", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				for _, target := range src.Contact {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "Contact", refNames, formGroup, probe.formStage)
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
	case *aixmfeatures.DMETimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("DMETimeSlice", inst.DMETimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DMETimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DMETypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DMEType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.DMETimeSliceType:
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
		StageSetAssociationFieldToForm("Designator", inst.Designator, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavaidDesignatorType](), probe.formStage)
		StageSetAssociationFieldToForm("EmissionClass", inst.EmissionClass, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadioEmissionType](), probe.formStage)
		StageSetAssociationFieldToForm("Mobile", inst.Mobile, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticVariation", inst.MagneticVariation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticVariationAccuracy", inst.MagneticVariationAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("DateMagneticVariation", inst.DateMagneticVariation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), probe.formStage)
		StageSetAssociationFieldToForm("FlightChecked", inst.FlightChecked, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Location", inst.Location, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)

		{
			// Slice of pointers: Authority
			div := (&form.FormDiv{Name: "Authority"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Authority {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Authority",
				Label: "Authority",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Monitoring
			div := (&form.FormDiv{Name: "Monitoring"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Monitoring {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Monitoring",
				Label: "Monitoring",
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
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDMEType](), probe.formStage)
		StageSetAssociationFieldToForm("Channel", inst.Channel, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDMEChannelType](), probe.formStage)
		StageSetAssociationFieldToForm("GhostFrequency", inst.GhostFrequency, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValFrequencyType](), probe.formStage)
		StageSetAssociationFieldToForm("Displace", inst.Displace, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.DMETimeSlicePropertyTypes {
				if src.DMETimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DMETimeSlicePropertyType", "DMETimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.DMEType:
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
				if src.DME == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "DME", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTimeSliceTypes {
				if src.Fix == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleIndicationTimeSliceType", "Fix", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTimeSliceTypes {
				if src.PointChoice_fixDesignatedPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleIndicationTimeSliceType", "PointChoice_fixDesignatedPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DistanceIndicationTimeSliceTypes {
				if src.Fix == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DistanceIndicationTimeSliceType", "Fix", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DistanceIndicationTimeSliceTypes {
				if src.PointChoice_fixDesignatedPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DistanceIndicationTimeSliceType", "PointChoice_fixDesignatedPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SignificantPointInAirspaceTimeSliceTypes {
				if src.Location_fixDesignatedPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SignificantPointInAirspaceTimeSliceType", "Location_fixDesignatedPoint", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.DesignatedPointTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("DesignatedPointTimeSlice", inst.DesignatedPointTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DesignatedPointTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DesignatedPointType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.DesignatedPointTimeSliceType:
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
		StageSetAssociationFieldToForm("Designator", inst.Designator, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDesignatedPointDesignatorType](), probe.formStage)
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDesignatedPointType](), probe.formStage)
		StageSetAssociationFieldToForm("Location", inst.Location, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("AimingPoint", inst.AimingPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("AirportHeliport", inst.AirportHeliport, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("RunwayPoint", inst.RunwayPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.DesignatedPointTimeSlicePropertyTypes {
				if src.DesignatedPointTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DesignatedPointTimeSlicePropertyType", "DesignatedPointTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.DesignatedPointType:
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
				if src.DesignatedPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "DesignatedPoint", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.DirectionFinderTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("DirectionFinderTimeSlice", inst.DirectionFinderTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DirectionFinderTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectionFinderType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.DirectionFinderTimeSliceType:
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
		StageSetAssociationFieldToForm("Designator", inst.Designator, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavaidDesignatorType](), probe.formStage)
		StageSetAssociationFieldToForm("EmissionClass", inst.EmissionClass, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadioEmissionType](), probe.formStage)
		StageSetAssociationFieldToForm("Mobile", inst.Mobile, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticVariation", inst.MagneticVariation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticVariationAccuracy", inst.MagneticVariationAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("DateMagneticVariation", inst.DateMagneticVariation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), probe.formStage)
		StageSetAssociationFieldToForm("FlightChecked", inst.FlightChecked, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Location", inst.Location, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)

		{
			// Slice of pointers: Authority
			div := (&form.FormDiv{Name: "Authority"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Authority {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Authority",
				Label: "Authority",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Monitoring
			div := (&form.FormDiv{Name: "Monitoring"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Monitoring {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Monitoring",
				Label: "Monitoring",
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
		StageSetAssociationFieldToForm("Doppler", inst.Doppler, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)

		{
			// Slice of pointers: InformationProvision
			div := (&form.FormDiv{Name: "InformationProvision"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.InformationProvision {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "InformationProvision",
				Label: "InformationProvision",
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTimeSlicePropertyTypes {
				if src.DirectionFinderTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectionFinderTimeSlicePropertyType", "DirectionFinderTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.DirectionFinderType:
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
				if src.DirectionFinder == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "DirectionFinder", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.DistanceIndicationTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("DistanceIndicationTimeSlice", inst.DistanceIndicationTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DistanceIndicationTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DistanceIndicationTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DistanceIndicationType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.DistanceIndicationTimeSliceType:
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
		StageSetAssociationFieldToForm("MinimumReceptionAltitude", inst.MinimumReceptionAltitude, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeDistanceIndicationType](), probe.formStage)
		StageSetAssociationFieldToForm("Fix", inst.Fix, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PointChoice_fixDesignatedPoint", inst.PointChoice_fixDesignatedPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PointChoice_navaidSystem", inst.PointChoice_navaidSystem, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PointChoice_position", inst.PointChoice_position, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PointChoice_runwayPoint", inst.PointChoice_runwayPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PointChoice_aimingPoint", inst.PointChoice_aimingPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("PointChoice_airportReferencePoint", inst.PointChoice_airportReferencePoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.DistanceIndicationTimeSlicePropertyTypes {
				if src.DistanceIndicationTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DistanceIndicationTimeSlicePropertyType", "DistanceIndicationTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.DistanceIndicationType:
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
				if src.DistanceIndication == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "DistanceIndication", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ElevatedPointPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ElevatedPoint", inst.ElevatedPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSliceTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthTimeSliceType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointINSTimeSliceTypes {
				if src.Position == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointINSTimeSliceType", "Position", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointVORTimeSliceTypes {
				if src.Position == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointVORTimeSliceType", "Position", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DMETimeSliceTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DMETimeSliceType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTimeSliceTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectionFinderTimeSliceType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevationTimeSliceTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevationTimeSliceType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTimeSliceTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathTimeSliceType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSliceTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerTimeSliceType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTimeSliceTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MarkerBeaconTimeSliceType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NDBTimeSliceTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NDBTimeSliceType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidTimeSliceTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidTimeSliceType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSliceTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarTimeSliceType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarSystemTimeSliceTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarSystemTimeSliceType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ReflectorTypes {
				if src.TouchdownReflector == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ReflectorType", "TouchdownReflector", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SDFTimeSliceTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SDFTimeSliceType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationTimeSliceTypes {
				if src.Position == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationStationTimeSliceType", "Position", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TACANTimeSliceTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TACANTimeSliceType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VORTimeSliceTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VORTimeSliceType", "Location", refNames, formGroup, probe.formStage)
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
	case *aixmfeatures.ElevationTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ElevationTimeSlice", inst.ElevationTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevationTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevationTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevationType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ElevationTimeSliceType:
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
		StageSetAssociationFieldToForm("Designator", inst.Designator, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavaidDesignatorType](), probe.formStage)
		StageSetAssociationFieldToForm("EmissionClass", inst.EmissionClass, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadioEmissionType](), probe.formStage)
		StageSetAssociationFieldToForm("Mobile", inst.Mobile, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticVariation", inst.MagneticVariation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticVariationAccuracy", inst.MagneticVariationAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("DateMagneticVariation", inst.DateMagneticVariation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), probe.formStage)
		StageSetAssociationFieldToForm("FlightChecked", inst.FlightChecked, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Location", inst.Location, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)

		{
			// Slice of pointers: Authority
			div := (&form.FormDiv{Name: "Authority"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Authority {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Authority",
				Label: "Authority",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Monitoring
			div := (&form.FormDiv{Name: "Monitoring"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Monitoring {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Monitoring",
				Label: "Monitoring",
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
		StageSetAssociationFieldToForm("AngleNominal", inst.AngleNominal, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("AngleMinimum", inst.AngleMinimum, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("AngleSpan", inst.AngleSpan, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("AngleAccuracy", inst.AngleAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevationTimeSlicePropertyTypes {
				if src.ElevationTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevationTimeSlicePropertyType", "ElevationTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ElevationType:
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
				if src.Elevation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "Elevation", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.GlidepathTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("GlidepathTimeSlice", inst.GlidepathTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.GlidepathTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.GlidepathTimeSliceType:
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
		StageSetAssociationFieldToForm("Designator", inst.Designator, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavaidDesignatorType](), probe.formStage)
		StageSetAssociationFieldToForm("EmissionClass", inst.EmissionClass, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadioEmissionType](), probe.formStage)
		StageSetAssociationFieldToForm("Mobile", inst.Mobile, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticVariation", inst.MagneticVariation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticVariationAccuracy", inst.MagneticVariationAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("DateMagneticVariation", inst.DateMagneticVariation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), probe.formStage)
		StageSetAssociationFieldToForm("FlightChecked", inst.FlightChecked, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Location", inst.Location, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)

		{
			// Slice of pointers: Authority
			div := (&form.FormDiv{Name: "Authority"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Authority {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Authority",
				Label: "Authority",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Monitoring
			div := (&form.FormDiv{Name: "Monitoring"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Monitoring {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Monitoring",
				Label: "Monitoring",
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
		StageSetAssociationFieldToForm("Frequency", inst.Frequency, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValFrequencyType](), probe.formStage)
		StageSetAssociationFieldToForm("Slope", inst.Slope, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("AngleAccuracy", inst.AngleAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("Rdh", inst.Rdh, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("RdhAccuracy", inst.RdhAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTimeSlicePropertyTypes {
				if src.GlidepathTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathTimeSlicePropertyType", "GlidepathTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.GlidepathType:
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
				if src.Glidepath == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "Glidepath", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.InformationServicePropertyType:
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTimeSliceTypes {
				for _, target := range src.InformationProvision {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectionFinderTimeSliceType", "InformationProvision", refNames, formGroup, probe.formStage)
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
	case *aixmfeatures.LocalizerTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("LocalizerTimeSlice", inst.LocalizerTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.LocalizerTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.LocalizerTimeSliceType:
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
		StageSetAssociationFieldToForm("Designator", inst.Designator, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavaidDesignatorType](), probe.formStage)
		StageSetAssociationFieldToForm("EmissionClass", inst.EmissionClass, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadioEmissionType](), probe.formStage)
		StageSetAssociationFieldToForm("Mobile", inst.Mobile, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticVariation", inst.MagneticVariation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticVariationAccuracy", inst.MagneticVariationAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("DateMagneticVariation", inst.DateMagneticVariation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), probe.formStage)
		StageSetAssociationFieldToForm("FlightChecked", inst.FlightChecked, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Location", inst.Location, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)

		{
			// Slice of pointers: Authority
			div := (&form.FormDiv{Name: "Authority"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Authority {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Authority",
				Label: "Authority",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Monitoring
			div := (&form.FormDiv{Name: "Monitoring"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Monitoring {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Monitoring",
				Label: "Monitoring",
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
		StageSetAssociationFieldToForm("Frequency", inst.Frequency, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValFrequencyType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticBearing", inst.MagneticBearing, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticBearingAccuracy", inst.MagneticBearingAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("TrueBearing", inst.TrueBearing, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), probe.formStage)
		StageSetAssociationFieldToForm("TrueBearingAccuracy", inst.TrueBearingAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("Declination", inst.Declination, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), probe.formStage)
		StageSetAssociationFieldToForm("WidthCourse", inst.WidthCourse, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("WidthCourseAccuracy", inst.WidthCourseAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("BackCourseUsable", inst.BackCourseUsable, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeILSBackCourseType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSlicePropertyTypes {
				if src.LocalizerTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerTimeSlicePropertyType", "LocalizerTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.LocalizerType:
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
				if src.Localizer == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "Localizer", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.MarkerBeaconTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("MarkerBeaconTimeSlice", inst.MarkerBeaconTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.MarkerBeaconTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MarkerBeaconType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.MarkerBeaconTimeSliceType:
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
		StageSetAssociationFieldToForm("Designator", inst.Designator, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavaidDesignatorType](), probe.formStage)
		StageSetAssociationFieldToForm("EmissionClass", inst.EmissionClass, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadioEmissionType](), probe.formStage)
		StageSetAssociationFieldToForm("Mobile", inst.Mobile, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticVariation", inst.MagneticVariation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticVariationAccuracy", inst.MagneticVariationAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("DateMagneticVariation", inst.DateMagneticVariation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), probe.formStage)
		StageSetAssociationFieldToForm("FlightChecked", inst.FlightChecked, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Location", inst.Location, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)

		{
			// Slice of pointers: Authority
			div := (&form.FormDiv{Name: "Authority"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Authority {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Authority",
				Label: "Authority",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Monitoring
			div := (&form.FormDiv{Name: "Monitoring"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Monitoring {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Monitoring",
				Label: "Monitoring",
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
		StageSetAssociationFieldToForm("Class", inst.Class, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeMarkerBeaconSignalType](), probe.formStage)
		StageSetAssociationFieldToForm("Frequency", inst.Frequency, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValFrequencyType](), probe.formStage)
		StageSetAssociationFieldToForm("AxisBearing", inst.AxisBearing, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), probe.formStage)
		StageSetAssociationFieldToForm("AuralMorseCode", inst.AuralMorseCode, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAuralMorseType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTimeSlicePropertyTypes {
				if src.MarkerBeaconTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MarkerBeaconTimeSlicePropertyType", "MarkerBeaconTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.MarkerBeaconType:
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
				if src.MarkerBeacon == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "MarkerBeacon", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.NDBTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("NDBTimeSlice", inst.NDBTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NDBTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NDBTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NDBType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.NDBTimeSliceType:
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
		StageSetAssociationFieldToForm("Designator", inst.Designator, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavaidDesignatorType](), probe.formStage)
		StageSetAssociationFieldToForm("EmissionClass", inst.EmissionClass, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadioEmissionType](), probe.formStage)
		StageSetAssociationFieldToForm("Mobile", inst.Mobile, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticVariation", inst.MagneticVariation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticVariationAccuracy", inst.MagneticVariationAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("DateMagneticVariation", inst.DateMagneticVariation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), probe.formStage)
		StageSetAssociationFieldToForm("FlightChecked", inst.FlightChecked, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Location", inst.Location, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)

		{
			// Slice of pointers: Authority
			div := (&form.FormDiv{Name: "Authority"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Authority {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Authority",
				Label: "Authority",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Monitoring
			div := (&form.FormDiv{Name: "Monitoring"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Monitoring {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Monitoring",
				Label: "Monitoring",
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
		StageSetAssociationFieldToForm("Frequency", inst.Frequency, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValFrequencyType](), probe.formStage)
		StageSetAssociationFieldToForm("Class", inst.Class, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNDBUsageType](), probe.formStage)
		StageSetAssociationFieldToForm("EmissionBand", inst.EmissionBand, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeEmissionBandType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.NDBTimeSlicePropertyTypes {
				if src.NDBTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NDBTimeSlicePropertyType", "NDBTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.NDBType:
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
				if src.NDB == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "NDB", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.NavaidComponentPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("NavaidComponent", inst.NavaidComponent, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidComponentType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidTimeSliceTypes {
				for _, target := range src.NavaidEquipment {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidTimeSliceType", "NavaidEquipment", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.NavaidComponentType:
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
		StageSetAssociationFieldToForm("CollocationGroup", inst.CollocationGroup, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoSequenceType](), probe.formStage)
		StageSetAssociationFieldToForm("MarkerPosition", inst.MarkerPosition, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodePositionInILSType](), probe.formStage)
		StageSetAssociationFieldToForm("ProvidesNavigableLocation", inst.ProvidesNavigableLocation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)

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
		StageSetAssociationFieldToForm("TheNavaidEquipment", inst.TheNavaidEquipment, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidEquipmentPropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidComponentPropertyTypes {
				if src.NavaidComponent == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidComponentPropertyType", "NavaidComponent", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.NavaidEquipmentMonitoringPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("NavaidEquipmentMonitoring", inst.NavaidEquipmentMonitoring, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidEquipmentMonitoringType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSliceTypes {
				for _, target := range src.Monitoring {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthTimeSliceType", "Monitoring", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DMETimeSliceTypes {
				for _, target := range src.Monitoring {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DMETimeSliceType", "Monitoring", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTimeSliceTypes {
				for _, target := range src.Monitoring {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectionFinderTimeSliceType", "Monitoring", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevationTimeSliceTypes {
				for _, target := range src.Monitoring {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevationTimeSliceType", "Monitoring", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTimeSliceTypes {
				for _, target := range src.Monitoring {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathTimeSliceType", "Monitoring", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSliceTypes {
				for _, target := range src.Monitoring {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerTimeSliceType", "Monitoring", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTimeSliceTypes {
				for _, target := range src.Monitoring {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MarkerBeaconTimeSliceType", "Monitoring", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NDBTimeSliceTypes {
				for _, target := range src.Monitoring {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NDBTimeSliceType", "Monitoring", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SDFTimeSliceTypes {
				for _, target := range src.Monitoring {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SDFTimeSliceType", "Monitoring", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TACANTimeSliceTypes {
				for _, target := range src.Monitoring {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TACANTimeSliceType", "Monitoring", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VORTimeSliceTypes {
				for _, target := range src.Monitoring {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VORTimeSliceType", "Monitoring", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.NavaidEquipmentMonitoringType:
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
		StageSetAssociationFieldToForm("Monitored", inst.Monitored, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidEquipmentMonitoringPropertyTypes {
				if src.NavaidEquipmentMonitoring == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidEquipmentMonitoringPropertyType", "NavaidEquipmentMonitoring", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidComponentTypes {
				if src.TheNavaidEquipment == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidComponentType", "TheNavaidEquipment", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.NavaidOperationalStatusPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("NavaidOperationalStatus", inst.NavaidOperationalStatus, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidOperationalStatusType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSliceTypes {
				for _, target := range src.Availability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthTimeSliceType", "Availability", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DMETimeSliceTypes {
				for _, target := range src.Availability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DMETimeSliceType", "Availability", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTimeSliceTypes {
				for _, target := range src.Availability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectionFinderTimeSliceType", "Availability", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevationTimeSliceTypes {
				for _, target := range src.Availability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevationTimeSliceType", "Availability", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTimeSliceTypes {
				for _, target := range src.Availability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathTimeSliceType", "Availability", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSliceTypes {
				for _, target := range src.Availability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerTimeSliceType", "Availability", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTimeSliceTypes {
				for _, target := range src.Availability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MarkerBeaconTimeSliceType", "Availability", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NDBTimeSliceTypes {
				for _, target := range src.Availability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NDBTimeSliceType", "Availability", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidTimeSliceTypes {
				for _, target := range src.Availability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidTimeSliceType", "Availability", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SDFTimeSliceTypes {
				for _, target := range src.Availability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SDFTimeSliceType", "Availability", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TACANTimeSliceTypes {
				for _, target := range src.Availability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TACANTimeSliceType", "Availability", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VORTimeSliceTypes {
				for _, target := range src.Availability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VORTimeSliceType", "Availability", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.NavaidOperationalStatusType:
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
		StageSetAssociationFieldToForm("OperationalStatus", inst.OperationalStatus, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeStatusNavaidType](), probe.formStage)
		StageSetAssociationFieldToForm("SignalType", inst.SignalType, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadioSignalType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidOperationalStatusPropertyTypes {
				if src.NavaidOperationalStatus == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidOperationalStatusPropertyType", "NavaidOperationalStatus", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTimeSliceTypes {
				if src.PointChoice_navaidSystem == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleIndicationTimeSliceType", "PointChoice_navaidSystem", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DistanceIndicationTimeSliceTypes {
				if src.PointChoice_navaidSystem == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DistanceIndicationTimeSliceType", "PointChoice_navaidSystem", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SignificantPointInAirspaceTimeSliceTypes {
				if src.Location_navaidSystem == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SignificantPointInAirspaceTimeSliceType", "Location_navaidSystem", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.NavaidTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("NavaidTimeSlice", inst.NavaidTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.NavaidTimeSliceType:
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
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavaidServiceType](), probe.formStage)
		StageSetAssociationFieldToForm("Designator", inst.Designator, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavaidDesignatorType](), probe.formStage)
		StageSetAssociationFieldToForm("FlightChecked", inst.FlightChecked, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Purpose", inst.Purpose, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavaidPurposeType](), probe.formStage)
		StageSetAssociationFieldToForm("SignalPerformance", inst.SignalPerformance, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSignalPerformanceILSType](), probe.formStage)
		StageSetAssociationFieldToForm("CourseQuality", inst.CourseQuality, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCourseQualityILSType](), probe.formStage)
		StageSetAssociationFieldToForm("IntegrityLevel", inst.IntegrityLevel, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeIntegrityLevelILSType](), probe.formStage)

		{
			// Slice of pointers: TouchDownLiftOff
			div := (&form.FormDiv{Name: "TouchDownLiftOff"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TouchDownLiftOff {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TouchDownLiftOff",
				Label: "TouchDownLiftOff",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: NavaidEquipment
			div := (&form.FormDiv{Name: "NavaidEquipment"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NavaidEquipment {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NavaidEquipment",
				Label: "NavaidEquipment",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Location", inst.Location, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)

		{
			// Slice of pointers: RunwayDirection
			div := (&form.FormDiv{Name: "RunwayDirection"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RunwayDirection {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RunwayDirection",
				Label: "RunwayDirection",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ServedAirport
			div := (&form.FormDiv{Name: "ServedAirport"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ServedAirport {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ServedAirport",
				Label: "ServedAirport",
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidTimeSlicePropertyTypes {
				if src.NavaidTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidTimeSlicePropertyType", "NavaidTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.NavaidType:
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
				if src.Navaid == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "Navaid", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.NotePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Note", inst.Note, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NoteType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleIndicationTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForNavaidEquipmentTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForNavaidEquipmentType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForSpecialNavigationStationTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForSpecialNavigationStationType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForSpecialNavigationSystemTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForSpecialNavigationSystemType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointINSTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointINSTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointVORTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointVORTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.DMETimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DMETimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DesignatedPointTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DesignatedPointTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectionFinderTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DistanceIndicationTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DistanceIndicationTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevationTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevationTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MarkerBeaconTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NDBTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NDBTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidComponentTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidComponentType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidEquipmentMonitoringTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidEquipmentMonitoringType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidOperationalStatusTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidOperationalStatusType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarComponentTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarComponentType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarSystemTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarSystemTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ReflectorTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ReflectorType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SDFTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SDFTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SignificantPointInAirspaceTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SignificantPointInAirspaceTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationStatusTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationStationStatusType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationStationTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationSystemTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationSystemTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SurveillanceGroundStationTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SurveillanceGroundStationType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TACANTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TACANTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
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

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VORTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VORTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForNavaidEquipmentTypes {
				if src.TheOrganisationAuthority == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForNavaidEquipmentType", "TheOrganisationAuthority", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForSpecialNavigationStationTypes {
				if src.TheOrganisationAuthority == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForSpecialNavigationStationType", "TheOrganisationAuthority", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForSpecialNavigationSystemTypes {
				if src.TheOrganisationAuthority == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForSpecialNavigationSystemType", "TheOrganisationAuthority", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidEquipmentMonitoringTypes {
				for _, target := range src.SpecialDateAuthority {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidEquipmentMonitoringType", "SpecialDateAuthority", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidOperationalStatusTypes {
				for _, target := range src.SpecialDateAuthority {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidOperationalStatusType", "SpecialDateAuthority", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarSystemTimeSliceTypes {
				for _, target := range src.Office {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarSystemTimeSliceType", "Office", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationStatusTypes {
				for _, target := range src.SpecialDateAuthority {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationStationStatusType", "SpecialDateAuthority", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTimeSliceTypes {
				if src.PointChoice_position == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleIndicationTimeSliceType", "PointChoice_position", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DesignatedPointTimeSliceTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DesignatedPointTimeSliceType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DistanceIndicationTimeSliceTypes {
				if src.PointChoice_position == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DistanceIndicationTimeSliceType", "PointChoice_position", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SignificantPointInAirspaceTimeSliceTypes {
				if src.Location_position == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SignificantPointInAirspaceTimeSliceType", "Location_position", refNames, formGroup, probe.formStage)
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
	case *aixmfeatures.PrecisionApproachRadarTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("PrecisionApproachRadarTimeSlice", inst.PrecisionApproachRadarTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PrecisionApproachRadarTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.PrecisionApproachRadarTimeSliceType:
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
		StageSetAssociationFieldToForm("SerialNumber", inst.SerialNumber, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextDesignatorType](), probe.formStage)
		StageSetAssociationFieldToForm("Range_XSD", inst.Range_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)
		StageSetAssociationFieldToForm("RangeAccuracy", inst.RangeAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)
		StageSetAssociationFieldToForm("DualChannel", inst.DualChannel, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("MovingTargetIndicator", inst.MovingTargetIndicator, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("StandbyPower", inst.StandbyPower, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeStandbyPowerType](), probe.formStage)
		StageSetAssociationFieldToForm("Digital", inst.Digital, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("MilitaryUseOnly", inst.MilitaryUseOnly, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("SpecialUseOnly", inst.SpecialUseOnly, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("SpecialAircraftOnly", inst.SpecialAircraftOnly, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticVariation", inst.MagneticVariation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticVariationAccuracy", inst.MagneticVariationAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("DateMagneticVariation", inst.DateMagneticVariation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), probe.formStage)

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
		StageSetAssociationFieldToForm("PrecisionApproachRadarType", inst.PrecisionApproachRadarType, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodePARType](), probe.formStage)
		StageSetAssociationFieldToForm("Slope", inst.Slope, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("SlopeAccuracy", inst.SlopeAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)

		{
			// Slice of pointers: Reflector
			div := (&form.FormDiv{Name: "Reflector"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Reflector {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Reflector",
				Label: "Reflector",
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSlicePropertyTypes {
				if src.PrecisionApproachRadarTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarTimeSlicePropertyType", "PrecisionApproachRadarTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.PrecisionApproachRadarType:
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
				if src.PrecisionApproachRadar == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "PrecisionApproachRadar", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.PrimarySurveillanceRadarTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("PrimarySurveillanceRadarTimeSlice", inst.PrimarySurveillanceRadarTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PrimarySurveillanceRadarTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.PrimarySurveillanceRadarTimeSliceType:
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
		StageSetAssociationFieldToForm("SerialNumber", inst.SerialNumber, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextDesignatorType](), probe.formStage)
		StageSetAssociationFieldToForm("Range_XSD", inst.Range_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)
		StageSetAssociationFieldToForm("RangeAccuracy", inst.RangeAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)
		StageSetAssociationFieldToForm("DualChannel", inst.DualChannel, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("MovingTargetIndicator", inst.MovingTargetIndicator, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("StandbyPower", inst.StandbyPower, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeStandbyPowerType](), probe.formStage)
		StageSetAssociationFieldToForm("Digital", inst.Digital, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("MilitaryUseOnly", inst.MilitaryUseOnly, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("SpecialUseOnly", inst.SpecialUseOnly, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("SpecialAircraftOnly", inst.SpecialAircraftOnly, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticVariation", inst.MagneticVariation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticVariationAccuracy", inst.MagneticVariationAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("DateMagneticVariation", inst.DateMagneticVariation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), probe.formStage)

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
		StageSetAssociationFieldToForm("VerticalCoverageAltitude", inst.VerticalCoverageAltitude, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("VerticalCoverageDistance", inst.VerticalCoverageDistance, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)
		StageSetAssociationFieldToForm("VerticalCoverageAzimuth", inst.VerticalCoverageAzimuth, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), probe.formStage)
		StageSetAssociationFieldToForm("AntennaTiltFixed", inst.AntennaTiltFixed, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("TiltAngle", inst.TiltAngle, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("AutomatedRadarTerminalSystem", inst.AutomatedRadarTerminalSystem, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextDesignatorType](), probe.formStage)

		{
			// Slice of pointers: GroundStation
			div := (&form.FormDiv{Name: "GroundStation"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GroundStation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GroundStation",
				Label: "GroundStation",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodePrimaryRadarType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSlicePropertyTypes {
				if src.PrimarySurveillanceRadarTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSlicePropertyType", "PrimarySurveillanceRadarTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.PrimarySurveillanceRadarType:
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
				if src.PrimarySurveillanceRadar == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "PrimarySurveillanceRadar", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RadarComponentPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("RadarComponent", inst.RadarComponent, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RadarComponentType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarSystemTimeSliceTypes {
				for _, target := range src.RadarEquipment {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarSystemTimeSliceType", "RadarEquipment", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RadarComponentType:
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
		StageSetAssociationFieldToForm("CollocationGroup", inst.CollocationGroup, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoSequenceType](), probe.formStage)

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
		StageSetAssociationFieldToForm("TheRadarEquipment", inst.TheRadarEquipment, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RadarEquipmentPropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarComponentPropertyTypes {
				if src.RadarComponent == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarComponentPropertyType", "RadarComponent", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RadarEquipmentPropertyType:
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarComponentTypes {
				if src.TheRadarEquipment == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarComponentType", "TheRadarEquipment", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RadarSystemTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("RadarSystemTimeSlice", inst.RadarSystemTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RadarSystemTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarSystemTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarSystemType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RadarSystemTimeSliceType:
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
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadarServiceType](), probe.formStage)
		StageSetAssociationFieldToForm("Model", inst.Model, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextNameType](), probe.formStage)
		StageSetAssociationFieldToForm("GeneralTerrainMonitor", inst.GeneralTerrainMonitor, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("BroadcastIdentifier", inst.BroadcastIdentifier, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextDesignatorType](), probe.formStage)

		{
			// Slice of pointers: RadarEquipment
			div := (&form.FormDiv{Name: "RadarEquipment"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RadarEquipment {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RadarEquipment",
				Label: "RadarEquipment",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Office
			div := (&form.FormDiv{Name: "Office"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Office {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Office",
				Label: "Office",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("AirportHeliport", inst.AirportHeliport, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), probe.formStage)

		{
			// Slice of pointers: PARRunway
			div := (&form.FormDiv{Name: "PARRunway"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.PARRunway {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "PARRunway",
				Label: "PARRunway",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarSystemTimeSlicePropertyTypes {
				if src.RadarSystemTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarSystemTimeSlicePropertyType", "RadarSystemTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RadarSystemType:
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
				if src.RadarSystem == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "RadarSystem", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ReflectorPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Reflector", inst.Reflector, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ReflectorType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSliceTypes {
				for _, target := range src.Reflector {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarTimeSliceType", "Reflector", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ReflectorType:
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
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeReflectorType](), probe.formStage)
		StageSetAssociationFieldToForm("TouchdownReflector", inst.TouchdownReflector, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.ReflectorPropertyTypes {
				if src.Reflector == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ReflectorPropertyType", "Reflector", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTimeSliceTypes {
				if src.PointChoice_runwayPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleIndicationTimeSliceType", "PointChoice_runwayPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DesignatedPointTimeSliceTypes {
				if src.RunwayPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DesignatedPointTimeSliceType", "RunwayPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DistanceIndicationTimeSliceTypes {
				if src.PointChoice_runwayPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DistanceIndicationTimeSliceType", "PointChoice_runwayPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SignificantPointInAirspaceTimeSliceTypes {
				if src.Location_runwayPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SignificantPointInAirspaceTimeSliceType", "Location_runwayPoint", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidTimeSliceTypes {
				for _, target := range src.RunwayDirection {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidTimeSliceType", "RunwayDirection", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RunwayPropertyType:
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarSystemTimeSliceTypes {
				for _, target := range src.PARRunway {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarSystemTimeSliceType", "PARRunway", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SDFTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("SDFTimeSlice", inst.SDFTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SDFTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SDFTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SDFType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SDFTimeSliceType:
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
		StageSetAssociationFieldToForm("Designator", inst.Designator, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavaidDesignatorType](), probe.formStage)
		StageSetAssociationFieldToForm("EmissionClass", inst.EmissionClass, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadioEmissionType](), probe.formStage)
		StageSetAssociationFieldToForm("Mobile", inst.Mobile, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticVariation", inst.MagneticVariation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticVariationAccuracy", inst.MagneticVariationAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("DateMagneticVariation", inst.DateMagneticVariation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), probe.formStage)
		StageSetAssociationFieldToForm("FlightChecked", inst.FlightChecked, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Location", inst.Location, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)

		{
			// Slice of pointers: Authority
			div := (&form.FormDiv{Name: "Authority"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Authority {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Authority",
				Label: "Authority",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Monitoring
			div := (&form.FormDiv{Name: "Monitoring"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Monitoring {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Monitoring",
				Label: "Monitoring",
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
		StageSetAssociationFieldToForm("Frequency", inst.Frequency, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValFrequencyType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticBearing", inst.MagneticBearing, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), probe.formStage)
		StageSetAssociationFieldToForm("TrueBearing", inst.TrueBearing, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.SDFTimeSlicePropertyTypes {
				if src.SDFTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SDFTimeSlicePropertyType", "SDFTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SDFType:
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
				if src.SDF == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "SDF", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SecondarySurveillanceRadarTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("SecondarySurveillanceRadarTimeSlice", inst.SecondarySurveillanceRadarTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SecondarySurveillanceRadarTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SecondarySurveillanceRadarTimeSliceType:
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
		StageSetAssociationFieldToForm("SerialNumber", inst.SerialNumber, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextDesignatorType](), probe.formStage)
		StageSetAssociationFieldToForm("Range_XSD", inst.Range_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)
		StageSetAssociationFieldToForm("RangeAccuracy", inst.RangeAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)
		StageSetAssociationFieldToForm("DualChannel", inst.DualChannel, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("MovingTargetIndicator", inst.MovingTargetIndicator, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("StandbyPower", inst.StandbyPower, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeStandbyPowerType](), probe.formStage)
		StageSetAssociationFieldToForm("Digital", inst.Digital, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("MilitaryUseOnly", inst.MilitaryUseOnly, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("SpecialUseOnly", inst.SpecialUseOnly, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("SpecialAircraftOnly", inst.SpecialAircraftOnly, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticVariation", inst.MagneticVariation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticVariationAccuracy", inst.MagneticVariationAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("DateMagneticVariation", inst.DateMagneticVariation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), probe.formStage)

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
		StageSetAssociationFieldToForm("VerticalCoverageAltitude", inst.VerticalCoverageAltitude, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)
		StageSetAssociationFieldToForm("VerticalCoverageDistance", inst.VerticalCoverageDistance, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceType](), probe.formStage)
		StageSetAssociationFieldToForm("VerticalCoverageAzimuth", inst.VerticalCoverageAzimuth, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), probe.formStage)
		StageSetAssociationFieldToForm("AntennaTiltFixed", inst.AntennaTiltFixed, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("TiltAngle", inst.TiltAngle, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("AutomatedRadarTerminalSystem", inst.AutomatedRadarTerminalSystem, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextDesignatorType](), probe.formStage)

		{
			// Slice of pointers: GroundStation
			div := (&form.FormDiv{Name: "GroundStation"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GroundStation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GroundStation",
				Label: "GroundStation",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Transponder", inst.Transponder, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTransponderType](), probe.formStage)
		StageSetAssociationFieldToForm("Autonomous", inst.Autonomous, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Monopulse", inst.Monopulse, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSlicePropertyTypes {
				if src.SecondarySurveillanceRadarTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSlicePropertyType", "SecondarySurveillanceRadarTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SecondarySurveillanceRadarType:
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
				if src.SecondarySurveillanceRadar == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "SecondarySurveillanceRadar", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SignificantPointInAirspaceTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("SignificantPointInAirspaceTimeSlice", inst.SignificantPointInAirspaceTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SignificantPointInAirspaceTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SignificantPointInAirspaceTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SignificantPointInAirspaceType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SignificantPointInAirspaceTimeSliceType:
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
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAirspacePointRoleType](), probe.formStage)
		StageSetAssociationFieldToForm("RelativeLocation", inst.RelativeLocation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAirspacePointPositionType](), probe.formStage)
		StageSetAssociationFieldToForm("ContainingAirspace", inst.ContainingAirspace, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspacePropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Location_fixDesignatedPoint", inst.Location_fixDesignatedPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Location_navaidSystem", inst.Location_navaidSystem, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Location_position", inst.Location_position, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Location_runwayPoint", inst.Location_runwayPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RunwayCentrelinePointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Location_aimingPoint", inst.Location_aimingPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TouchDownLiftOffPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Location_airportReferencePoint", inst.Location_airportReferencePoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.SignificantPointInAirspaceTimeSlicePropertyTypes {
				if src.SignificantPointInAirspaceTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SignificantPointInAirspaceTimeSlicePropertyType", "SignificantPointInAirspaceTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SignificantPointInAirspaceType:
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
				if src.SignificantPointInAirspace == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "SignificantPointInAirspace", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SpecialNavigationStationStatusPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("SpecialNavigationStationStatus", inst.SpecialNavigationStationStatus, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SpecialNavigationStationStatusType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationTimeSliceTypes {
				for _, target := range src.Availability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationStationTimeSliceType", "Availability", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SpecialNavigationStationStatusType:
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
		StageSetAssociationFieldToForm("OperationalStatus", inst.OperationalStatus, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeStatusNavaidType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationStatusPropertyTypes {
				if src.SpecialNavigationStationStatus == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationStationStatusPropertyType", "SpecialNavigationStationStatus", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SpecialNavigationStationTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("SpecialNavigationStationTimeSlice", inst.SpecialNavigationStationTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SpecialNavigationStationTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationStationType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SpecialNavigationStationTimeSliceType:
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
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSpecialNavigationStationType](), probe.formStage)
		StageSetAssociationFieldToForm("Frequency", inst.Frequency, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValFrequencyType](), probe.formStage)
		StageSetAssociationFieldToForm("Emission", inst.Emission, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadioEmissionType](), probe.formStage)
		StageSetAssociationFieldToForm("SystemChain", inst.SystemChain, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SpecialNavigationSystemPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("ResponsibleOrganisation", inst.ResponsibleOrganisation, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AuthorityForSpecialNavigationStationPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("Position", inst.Position, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationTimeSlicePropertyTypes {
				if src.SpecialNavigationStationTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationStationTimeSlicePropertyType", "SpecialNavigationStationTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SpecialNavigationStationType:
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
				if src.SpecialNavigationStation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "SpecialNavigationStation", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SpecialNavigationSystemPropertyType:
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationTimeSliceTypes {
				if src.SystemChain == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationStationTimeSliceType", "SystemChain", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SpecialNavigationSystemTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("SpecialNavigationSystemTimeSlice", inst.SpecialNavigationSystemTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SpecialNavigationSystemTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationSystemTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationSystemType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SpecialNavigationSystemTimeSliceType:
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
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSpecialNavigationSystemType](), probe.formStage)
		StageSetAssociationFieldToForm("Designator", inst.Designator, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSpecialNavigationChainDesignatorType](), probe.formStage)
		StageSetAssociationFieldToForm("ResponsibleOrganisation", inst.ResponsibleOrganisation, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AuthorityForSpecialNavigationSystemPropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationSystemTimeSlicePropertyTypes {
				if src.SpecialNavigationSystemTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationSystemTimeSlicePropertyType", "SpecialNavigationSystemTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SpecialNavigationSystemType:
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
				if src.SpecialNavigationSystem == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "SpecialNavigationSystem", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SurveillanceGroundStationPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("SurveillanceGroundStation", inst.SurveillanceGroundStation, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SurveillanceGroundStationType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				for _, target := range src.GroundStation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "GroundStation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				for _, target := range src.GroundStation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "GroundStation", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SurveillanceGroundStationType:
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
		StageSetAssociationFieldToForm("VideoMap", inst.VideoMap, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)

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
		StageSetAssociationFieldToForm("TheUnit", inst.TheUnit, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.UnitPropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.SurveillanceGroundStationPropertyTypes {
				if src.SurveillanceGroundStation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SurveillanceGroundStationPropertyType", "SurveillanceGroundStation", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.TACANTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("TACANTimeSlice", inst.TACANTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TACANTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TACANTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TACANType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.TACANTimeSliceType:
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
		StageSetAssociationFieldToForm("Designator", inst.Designator, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavaidDesignatorType](), probe.formStage)
		StageSetAssociationFieldToForm("EmissionClass", inst.EmissionClass, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadioEmissionType](), probe.formStage)
		StageSetAssociationFieldToForm("Mobile", inst.Mobile, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticVariation", inst.MagneticVariation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticVariationAccuracy", inst.MagneticVariationAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("DateMagneticVariation", inst.DateMagneticVariation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), probe.formStage)
		StageSetAssociationFieldToForm("FlightChecked", inst.FlightChecked, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Location", inst.Location, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)

		{
			// Slice of pointers: Authority
			div := (&form.FormDiv{Name: "Authority"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Authority {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Authority",
				Label: "Authority",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Monitoring
			div := (&form.FormDiv{Name: "Monitoring"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Monitoring {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Monitoring",
				Label: "Monitoring",
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
		StageSetAssociationFieldToForm("Channel", inst.Channel, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeTACANChannelType](), probe.formStage)
		StageSetAssociationFieldToForm("Declination", inst.Declination, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.TACANTimeSlicePropertyTypes {
				if src.TACANTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TACANTimeSlicePropertyType", "TACANTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.TACANType:
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
				if src.TACAN == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "TACAN", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidEquipmentMonitoringTypes {
				for _, target := range src.TimeInterval {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidEquipmentMonitoringType", "TimeInterval", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidOperationalStatusTypes {
				for _, target := range src.TimeInterval {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidOperationalStatusType", "TimeInterval", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationStatusTypes {
				for _, target := range src.TimeInterval {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationStationStatusType", "TimeInterval", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTimeSliceTypes {
				if src.PointChoice_aimingPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleIndicationTimeSliceType", "PointChoice_aimingPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DesignatedPointTimeSliceTypes {
				if src.AimingPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DesignatedPointTimeSliceType", "AimingPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DistanceIndicationTimeSliceTypes {
				if src.PointChoice_aimingPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DistanceIndicationTimeSliceType", "PointChoice_aimingPoint", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidTimeSliceTypes {
				for _, target := range src.TouchDownLiftOff {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidTimeSliceType", "TouchDownLiftOff", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SignificantPointInAirspaceTimeSliceTypes {
				if src.Location_aimingPoint == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SignificantPointInAirspaceTimeSliceType", "Location_aimingPoint", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.SurveillanceGroundStationTypes {
				if src.TheUnit == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SurveillanceGroundStationType", "TheUnit", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.VORPropertyType:
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointVORTimeSliceTypes {
				if src.CheckPointFacility == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointVORTimeSliceType", "CheckPointFacility", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.VORTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("VORTimeSlice", inst.VORTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.VORTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VORTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VORType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.VORTimeSliceType:
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
		StageSetAssociationFieldToForm("Designator", inst.Designator, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNavaidDesignatorType](), probe.formStage)
		StageSetAssociationFieldToForm("EmissionClass", inst.EmissionClass, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadioEmissionType](), probe.formStage)
		StageSetAssociationFieldToForm("Mobile", inst.Mobile, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticVariation", inst.MagneticVariation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), probe.formStage)
		StageSetAssociationFieldToForm("MagneticVariationAccuracy", inst.MagneticVariationAccuracy, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValAngleType](), probe.formStage)
		StageSetAssociationFieldToForm("DateMagneticVariation", inst.DateMagneticVariation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), probe.formStage)
		StageSetAssociationFieldToForm("FlightChecked", inst.FlightChecked, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Location", inst.Location, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)

		{
			// Slice of pointers: Authority
			div := (&form.FormDiv{Name: "Authority"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Authority {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Authority",
				Label: "Authority",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Monitoring
			div := (&form.FormDiv{Name: "Monitoring"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Monitoring {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Monitoring",
				Label: "Monitoring",
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
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeVORType](), probe.formStage)
		StageSetAssociationFieldToForm("Frequency", inst.Frequency, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValFrequencyType](), probe.formStage)
		StageSetAssociationFieldToForm("ZeroBearingDirection", inst.ZeroBearingDirection, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNorthReferenceType](), probe.formStage)
		StageSetAssociationFieldToForm("Declination", inst.Declination, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValMagneticVariationType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.VORTimeSlicePropertyTypes {
				if src.VORTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VORTimeSlicePropertyType", "VORTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.VORType:
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
				if src.VOR == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "VOR", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeAirspacePointPositionType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SignificantPointInAirspaceTimeSliceTypes {
				if src.RelativeLocation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SignificantPointInAirspaceTimeSliceType", "RelativeLocation", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeAirspacePointRoleType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SignificantPointInAirspaceTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SignificantPointInAirspaceTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeAltitudeUseType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointINSTimeSliceTypes {
				if src.AltitudeInterpretation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointINSTimeSliceType", "AltitudeInterpretation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointVORTimeSliceTypes {
				if src.AltitudeInterpretation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointVORTimeSliceType", "AltitudeInterpretation", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeAuralMorseType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTimeSliceTypes {
				if src.AuralMorseCode == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MarkerBeaconTimeSliceType", "AuralMorseCode", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeAuthorityRoleType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForNavaidEquipmentTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForNavaidEquipmentType", "Type_XSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForSpecialNavigationStationTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForSpecialNavigationStationType", "Type_XSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForSpecialNavigationSystemTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForSpecialNavigationSystemType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeBearingType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTimeSliceTypes {
				if src.AngleType == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleIndicationTimeSliceType", "AngleType", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeCardinalDirectionType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTimeSliceTypes {
				if src.CardinalDirection == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleIndicationTimeSliceType", "CardinalDirection", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeCheckpointCategoryType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointINSTimeSliceTypes {
				if src.Category == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointINSTimeSliceType", "Category", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointVORTimeSliceTypes {
				if src.Category == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointVORTimeSliceType", "Category", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeCourseQualityILSType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidTimeSliceTypes {
				if src.CourseQuality == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidTimeSliceType", "CourseQuality", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeDMEChannelType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DMETimeSliceTypes {
				if src.Channel == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DMETimeSliceType", "Channel", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeDMEType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DMETimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DMETimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
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
	case *aixmdatatypes.CodeDesignatedPointDesignatorType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DesignatedPointTimeSliceTypes {
				if src.Designator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DesignatedPointTimeSliceType", "Designator", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeDesignatedPointType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DesignatedPointTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DesignatedPointTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeDirectionReferenceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTimeSliceTypes {
				if src.IndicationDirection == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleIndicationTimeSliceType", "IndicationDirection", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeDistanceIndicationType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DistanceIndicationTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DistanceIndicationTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeEmissionBandType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NDBTimeSliceTypes {
				if src.EmissionBand == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NDBTimeSliceType", "EmissionBand", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeILSBackCourseType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSliceTypes {
				if src.BackCourseUsable == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerTimeSliceType", "BackCourseUsable", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeIntegrityLevelILSType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidTimeSliceTypes {
				if src.IntegrityLevel == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidTimeSliceType", "IntegrityLevel", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeMLSAzimuthType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeMLSChannelType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSliceTypes {
				if src.Channel == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthTimeSliceType", "Channel", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeMarkerBeaconSignalType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTimeSliceTypes {
				if src.Class == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MarkerBeaconTimeSliceType", "Class", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeNDBUsageType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NDBTimeSliceTypes {
				if src.Class == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NDBTimeSliceType", "Class", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeNavaidDesignatorType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSliceTypes {
				if src.Designator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthTimeSliceType", "Designator", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DMETimeSliceTypes {
				if src.Designator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DMETimeSliceType", "Designator", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTimeSliceTypes {
				if src.Designator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectionFinderTimeSliceType", "Designator", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevationTimeSliceTypes {
				if src.Designator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevationTimeSliceType", "Designator", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTimeSliceTypes {
				if src.Designator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathTimeSliceType", "Designator", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSliceTypes {
				if src.Designator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerTimeSliceType", "Designator", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTimeSliceTypes {
				if src.Designator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MarkerBeaconTimeSliceType", "Designator", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NDBTimeSliceTypes {
				if src.Designator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NDBTimeSliceType", "Designator", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidTimeSliceTypes {
				if src.Designator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidTimeSliceType", "Designator", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SDFTimeSliceTypes {
				if src.Designator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SDFTimeSliceType", "Designator", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TACANTimeSliceTypes {
				if src.Designator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TACANTimeSliceType", "Designator", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VORTimeSliceTypes {
				if src.Designator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VORTimeSliceType", "Designator", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeNavaidPurposeType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidTimeSliceTypes {
				if src.Purpose == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidTimeSliceType", "Purpose", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeNavaidServiceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeNorthReferenceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VORTimeSliceTypes {
				if src.ZeroBearingDirection == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VORTimeSliceType", "ZeroBearingDirection", refNames, formGroup, probe.formStage)
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
	case *aixmdatatypes.CodePARType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSliceTypes {
				if src.PrecisionApproachRadarType == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarTimeSliceType", "PrecisionApproachRadarType", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodePositionInILSType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidComponentTypes {
				if src.MarkerPosition == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidComponentType", "MarkerPosition", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodePrimaryRadarType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeRadarServiceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarSystemTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarSystemTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeRadioEmissionType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSliceTypes {
				if src.EmissionClass == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthTimeSliceType", "EmissionClass", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DMETimeSliceTypes {
				if src.EmissionClass == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DMETimeSliceType", "EmissionClass", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTimeSliceTypes {
				if src.EmissionClass == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectionFinderTimeSliceType", "EmissionClass", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevationTimeSliceTypes {
				if src.EmissionClass == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevationTimeSliceType", "EmissionClass", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTimeSliceTypes {
				if src.EmissionClass == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathTimeSliceType", "EmissionClass", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSliceTypes {
				if src.EmissionClass == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerTimeSliceType", "EmissionClass", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTimeSliceTypes {
				if src.EmissionClass == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MarkerBeaconTimeSliceType", "EmissionClass", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NDBTimeSliceTypes {
				if src.EmissionClass == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NDBTimeSliceType", "EmissionClass", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SDFTimeSliceTypes {
				if src.EmissionClass == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SDFTimeSliceType", "EmissionClass", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationTimeSliceTypes {
				if src.Emission == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationStationTimeSliceType", "Emission", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TACANTimeSliceTypes {
				if src.EmissionClass == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TACANTimeSliceType", "EmissionClass", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VORTimeSliceTypes {
				if src.EmissionClass == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VORTimeSliceType", "EmissionClass", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeRadioSignalType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidOperationalStatusTypes {
				if src.SignalType == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidOperationalStatusType", "SignalType", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeReflectorType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ReflectorTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ReflectorType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeSignalPerformanceILSType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidTimeSliceTypes {
				if src.SignalPerformance == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidTimeSliceType", "SignalPerformance", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeSpecialNavigationChainDesignatorType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationSystemTimeSliceTypes {
				if src.Designator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationSystemTimeSliceType", "Designator", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeSpecialNavigationStationType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationStationTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeSpecialNavigationSystemType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationSystemTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationSystemTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeStandbyPowerType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSliceTypes {
				if src.StandbyPower == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarTimeSliceType", "StandbyPower", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				if src.StandbyPower == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "StandbyPower", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				if src.StandbyPower == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "StandbyPower", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeStatusNavaidType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidOperationalStatusTypes {
				if src.OperationalStatus == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidOperationalStatusType", "OperationalStatus", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationStatusTypes {
				if src.OperationalStatus == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationStationStatusType", "OperationalStatus", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeTACANChannelType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TACANTimeSliceTypes {
				if src.Channel == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TACANTimeSliceType", "Channel", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				if src.Transponder == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "Transponder", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeVORType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VORTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VORTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeVerticalDatumType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

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
	case *aixmdatatypes.CodeVerticalReferenceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointINSTimeSliceTypes {
				if src.UpperLimitReference == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointINSTimeSliceType", "UpperLimitReference", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointINSTimeSliceTypes {
				if src.LowerLimitReference == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointINSTimeSliceType", "LowerLimitReference", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointVORTimeSliceTypes {
				if src.UpperLimitReference == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointVORTimeSliceType", "UpperLimitReference", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointVORTimeSliceTypes {
				if src.LowerLimitReference == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointVORTimeSliceType", "LowerLimitReference", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeYesNoType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSliceTypes {
				if src.Mobile == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthTimeSliceType", "Mobile", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSliceTypes {
				if src.FlightChecked == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthTimeSliceType", "FlightChecked", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DMETimeSliceTypes {
				if src.Mobile == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DMETimeSliceType", "Mobile", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DMETimeSliceTypes {
				if src.FlightChecked == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DMETimeSliceType", "FlightChecked", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTimeSliceTypes {
				if src.Mobile == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectionFinderTimeSliceType", "Mobile", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTimeSliceTypes {
				if src.FlightChecked == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectionFinderTimeSliceType", "FlightChecked", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTimeSliceTypes {
				if src.Doppler == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectionFinderTimeSliceType", "Doppler", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevationTimeSliceTypes {
				if src.Mobile == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevationTimeSliceType", "Mobile", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevationTimeSliceTypes {
				if src.FlightChecked == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevationTimeSliceType", "FlightChecked", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTimeSliceTypes {
				if src.Mobile == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathTimeSliceType", "Mobile", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTimeSliceTypes {
				if src.FlightChecked == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathTimeSliceType", "FlightChecked", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSliceTypes {
				if src.Mobile == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerTimeSliceType", "Mobile", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSliceTypes {
				if src.FlightChecked == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerTimeSliceType", "FlightChecked", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTimeSliceTypes {
				if src.Mobile == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MarkerBeaconTimeSliceType", "Mobile", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTimeSliceTypes {
				if src.FlightChecked == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MarkerBeaconTimeSliceType", "FlightChecked", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NDBTimeSliceTypes {
				if src.Mobile == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NDBTimeSliceType", "Mobile", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NDBTimeSliceTypes {
				if src.FlightChecked == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NDBTimeSliceType", "FlightChecked", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidComponentTypes {
				if src.ProvidesNavigableLocation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidComponentType", "ProvidesNavigableLocation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidEquipmentMonitoringTypes {
				if src.Monitored == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidEquipmentMonitoringType", "Monitored", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidTimeSliceTypes {
				if src.FlightChecked == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidTimeSliceType", "FlightChecked", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSliceTypes {
				if src.DualChannel == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarTimeSliceType", "DualChannel", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSliceTypes {
				if src.MovingTargetIndicator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarTimeSliceType", "MovingTargetIndicator", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSliceTypes {
				if src.Digital == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarTimeSliceType", "Digital", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSliceTypes {
				if src.MilitaryUseOnly == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarTimeSliceType", "MilitaryUseOnly", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSliceTypes {
				if src.SpecialUseOnly == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarTimeSliceType", "SpecialUseOnly", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSliceTypes {
				if src.SpecialAircraftOnly == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarTimeSliceType", "SpecialAircraftOnly", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				if src.DualChannel == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "DualChannel", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				if src.MovingTargetIndicator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "MovingTargetIndicator", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				if src.Digital == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "Digital", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				if src.MilitaryUseOnly == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "MilitaryUseOnly", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				if src.SpecialUseOnly == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "SpecialUseOnly", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				if src.SpecialAircraftOnly == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "SpecialAircraftOnly", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				if src.AntennaTiltFixed == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "AntennaTiltFixed", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarSystemTimeSliceTypes {
				if src.GeneralTerrainMonitor == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarSystemTimeSliceType", "GeneralTerrainMonitor", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SDFTimeSliceTypes {
				if src.Mobile == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SDFTimeSliceType", "Mobile", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SDFTimeSliceTypes {
				if src.FlightChecked == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SDFTimeSliceType", "FlightChecked", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				if src.DualChannel == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "DualChannel", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				if src.MovingTargetIndicator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "MovingTargetIndicator", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				if src.Digital == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "Digital", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				if src.MilitaryUseOnly == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "MilitaryUseOnly", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				if src.SpecialUseOnly == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "SpecialUseOnly", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				if src.SpecialAircraftOnly == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "SpecialAircraftOnly", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				if src.AntennaTiltFixed == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "AntennaTiltFixed", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				if src.Autonomous == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "Autonomous", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				if src.Monopulse == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "Monopulse", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SurveillanceGroundStationTypes {
				if src.VideoMap == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SurveillanceGroundStationType", "VideoMap", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TACANTimeSliceTypes {
				if src.Mobile == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TACANTimeSliceType", "Mobile", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TACANTimeSliceTypes {
				if src.FlightChecked == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TACANTimeSliceType", "FlightChecked", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.VORTimeSliceTypes {
				if src.Mobile == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VORTimeSliceType", "Mobile", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VORTimeSliceTypes {
				if src.FlightChecked == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VORTimeSliceType", "FlightChecked", refNames, formGroup, probe.formStage)
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
	case *aixmdatatypes.DateYearType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSliceTypes {
				if src.DateMagneticVariation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthTimeSliceType", "DateMagneticVariation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DMETimeSliceTypes {
				if src.DateMagneticVariation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DMETimeSliceType", "DateMagneticVariation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTimeSliceTypes {
				if src.DateMagneticVariation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectionFinderTimeSliceType", "DateMagneticVariation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevationTimeSliceTypes {
				if src.DateMagneticVariation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevationTimeSliceType", "DateMagneticVariation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTimeSliceTypes {
				if src.DateMagneticVariation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathTimeSliceType", "DateMagneticVariation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSliceTypes {
				if src.DateMagneticVariation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerTimeSliceType", "DateMagneticVariation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTimeSliceTypes {
				if src.DateMagneticVariation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MarkerBeaconTimeSliceType", "DateMagneticVariation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NDBTimeSliceTypes {
				if src.DateMagneticVariation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NDBTimeSliceType", "DateMagneticVariation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSliceTypes {
				if src.DateMagneticVariation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarTimeSliceType", "DateMagneticVariation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				if src.DateMagneticVariation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "DateMagneticVariation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SDFTimeSliceTypes {
				if src.DateMagneticVariation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SDFTimeSliceType", "DateMagneticVariation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				if src.DateMagneticVariation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "DateMagneticVariation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TACANTimeSliceTypes {
				if src.DateMagneticVariation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TACANTimeSliceType", "DateMagneticVariation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VORTimeSliceTypes {
				if src.DateMagneticVariation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VORTimeSliceType", "DateMagneticVariation", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.NoSequenceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidComponentTypes {
				if src.CollocationGroup == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidComponentType", "CollocationGroup", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarComponentTypes {
				if src.CollocationGroup == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarComponentType", "CollocationGroup", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSliceTypes {
				if src.SerialNumber == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarTimeSliceType", "SerialNumber", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				if src.SerialNumber == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "SerialNumber", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				if src.AutomatedRadarTerminalSystem == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "AutomatedRadarTerminalSystem", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarSystemTimeSliceTypes {
				if src.BroadcastIdentifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarSystemTimeSliceType", "BroadcastIdentifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				if src.SerialNumber == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "SerialNumber", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				if src.AutomatedRadarTerminalSystem == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "AutomatedRadarTerminalSystem", refNames, formGroup, probe.formStage)
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

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarSystemTimeSliceTypes {
				if src.Model == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarSystemTimeSliceType", "Model", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSliceTypes {
				if src.MagneticVariationAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthTimeSliceType", "MagneticVariationAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSliceTypes {
				if src.TrueBearingAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthTimeSliceType", "TrueBearingAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSliceTypes {
				if src.AngleProportionalLeft == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthTimeSliceType", "AngleProportionalLeft", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSliceTypes {
				if src.AngleProportionalRight == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthTimeSliceType", "AngleProportionalRight", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSliceTypes {
				if src.AngleCoverLeft == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthTimeSliceType", "AngleCoverLeft", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSliceTypes {
				if src.AngleCoverRight == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthTimeSliceType", "AngleCoverRight", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DMETimeSliceTypes {
				if src.MagneticVariationAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DMETimeSliceType", "MagneticVariationAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTimeSliceTypes {
				if src.MagneticVariationAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectionFinderTimeSliceType", "MagneticVariationAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevationTimeSliceTypes {
				if src.MagneticVariationAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevationTimeSliceType", "MagneticVariationAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevationTimeSliceTypes {
				if src.AngleNominal == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevationTimeSliceType", "AngleNominal", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevationTimeSliceTypes {
				if src.AngleMinimum == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevationTimeSliceType", "AngleMinimum", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevationTimeSliceTypes {
				if src.AngleSpan == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevationTimeSliceType", "AngleSpan", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevationTimeSliceTypes {
				if src.AngleAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevationTimeSliceType", "AngleAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTimeSliceTypes {
				if src.MagneticVariationAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathTimeSliceType", "MagneticVariationAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTimeSliceTypes {
				if src.Slope == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathTimeSliceType", "Slope", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTimeSliceTypes {
				if src.AngleAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathTimeSliceType", "AngleAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSliceTypes {
				if src.MagneticVariationAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerTimeSliceType", "MagneticVariationAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSliceTypes {
				if src.MagneticBearingAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerTimeSliceType", "MagneticBearingAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSliceTypes {
				if src.TrueBearingAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerTimeSliceType", "TrueBearingAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSliceTypes {
				if src.WidthCourse == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerTimeSliceType", "WidthCourse", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSliceTypes {
				if src.WidthCourseAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerTimeSliceType", "WidthCourseAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTimeSliceTypes {
				if src.MagneticVariationAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MarkerBeaconTimeSliceType", "MagneticVariationAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NDBTimeSliceTypes {
				if src.MagneticVariationAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NDBTimeSliceType", "MagneticVariationAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSliceTypes {
				if src.MagneticVariationAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarTimeSliceType", "MagneticVariationAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSliceTypes {
				if src.Slope == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarTimeSliceType", "Slope", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSliceTypes {
				if src.SlopeAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarTimeSliceType", "SlopeAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				if src.MagneticVariationAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "MagneticVariationAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				if src.TiltAngle == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "TiltAngle", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SDFTimeSliceTypes {
				if src.MagneticVariationAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SDFTimeSliceType", "MagneticVariationAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				if src.MagneticVariationAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "MagneticVariationAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				if src.TiltAngle == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "TiltAngle", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TACANTimeSliceTypes {
				if src.MagneticVariationAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TACANTimeSliceType", "MagneticVariationAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VORTimeSliceTypes {
				if src.MagneticVariationAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VORTimeSliceType", "MagneticVariationAccuracy", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.ValBearingType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTimeSliceTypes {
				if src.Angle == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleIndicationTimeSliceType", "Angle", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTimeSliceTypes {
				if src.TrueAngle == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleIndicationTimeSliceType", "TrueAngle", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSliceTypes {
				if src.TrueBearing == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthTimeSliceType", "TrueBearing", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSliceTypes {
				if src.MagneticBearing == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthTimeSliceType", "MagneticBearing", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointINSTimeSliceTypes {
				if src.Angle == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointINSTimeSliceType", "Angle", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointVORTimeSliceTypes {
				if src.Angle == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointVORTimeSliceType", "Angle", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSliceTypes {
				if src.MagneticBearing == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerTimeSliceType", "MagneticBearing", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSliceTypes {
				if src.TrueBearing == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerTimeSliceType", "TrueBearing", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTimeSliceTypes {
				if src.AxisBearing == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MarkerBeaconTimeSliceType", "AxisBearing", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				if src.VerticalCoverageAzimuth == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "VerticalCoverageAzimuth", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SDFTimeSliceTypes {
				if src.MagneticBearing == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SDFTimeSliceType", "MagneticBearing", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SDFTimeSliceTypes {
				if src.TrueBearing == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SDFTimeSliceType", "TrueBearing", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				if src.VerticalCoverageAzimuth == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "VerticalCoverageAzimuth", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.ValDistanceSignedType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Uom", inst.Uom, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

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
	case *aixmdatatypes.ValDistanceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Uom", inst.Uom, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointINSTimeSliceTypes {
				if src.Distance == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointINSTimeSliceType", "Distance", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointVORTimeSliceTypes {
				if src.Distance == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointVORTimeSliceType", "Distance", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DMETimeSliceTypes {
				if src.Displace == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DMETimeSliceType", "Displace", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DistanceIndicationTimeSliceTypes {
				if src.Distance == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DistanceIndicationTimeSliceType", "Distance", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSliceTypes {
				if src.Range_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarTimeSliceType", "Range_XSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSliceTypes {
				if src.RangeAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarTimeSliceType", "RangeAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				if src.Range_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "Range_XSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				if src.RangeAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "RangeAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				if src.VerticalCoverageDistance == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "VerticalCoverageDistance", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				if src.Range_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "Range_XSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				if src.RangeAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "RangeAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				if src.VerticalCoverageDistance == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "VerticalCoverageDistance", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.ValDistanceVerticalType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Uom", inst.Uom, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTimeSliceTypes {
				if src.MinimumReceptionAltitude == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleIndicationTimeSliceType", "MinimumReceptionAltitude", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointINSTimeSliceTypes {
				if src.UpperLimit == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointINSTimeSliceType", "UpperLimit", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointINSTimeSliceTypes {
				if src.LowerLimit == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointINSTimeSliceType", "LowerLimit", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointVORTimeSliceTypes {
				if src.UpperLimit == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointVORTimeSliceType", "UpperLimit", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointVORTimeSliceTypes {
				if src.LowerLimit == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointVORTimeSliceType", "LowerLimit", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DistanceIndicationTimeSliceTypes {
				if src.MinimumReceptionAltitude == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DistanceIndicationTimeSliceType", "MinimumReceptionAltitude", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTimeSliceTypes {
				if src.Rdh == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathTimeSliceType", "Rdh", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTimeSliceTypes {
				if src.RdhAccuracy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathTimeSliceType", "RdhAccuracy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				if src.VerticalCoverageAltitude == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "VerticalCoverageAltitude", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				if src.VerticalCoverageAltitude == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "VerticalCoverageAltitude", refNames, formGroup, probe.formStage)
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
	case *aixmdatatypes.ValFrequencyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Uom", inst.Uom, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DMETimeSliceTypes {
				if src.GhostFrequency == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DMETimeSliceType", "GhostFrequency", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTimeSliceTypes {
				if src.Frequency == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathTimeSliceType", "Frequency", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSliceTypes {
				if src.Frequency == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerTimeSliceType", "Frequency", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTimeSliceTypes {
				if src.Frequency == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MarkerBeaconTimeSliceType", "Frequency", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NDBTimeSliceTypes {
				if src.Frequency == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NDBTimeSliceType", "Frequency", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SDFTimeSliceTypes {
				if src.Frequency == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SDFTimeSliceType", "Frequency", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationTimeSliceTypes {
				if src.Frequency == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationStationTimeSliceType", "Frequency", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VORTimeSliceTypes {
				if src.Frequency == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VORTimeSliceType", "Frequency", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.ValMagneticVariationType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSliceTypes {
				if src.MagneticVariation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthTimeSliceType", "MagneticVariation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DMETimeSliceTypes {
				if src.MagneticVariation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DMETimeSliceType", "MagneticVariation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTimeSliceTypes {
				if src.MagneticVariation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectionFinderTimeSliceType", "MagneticVariation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevationTimeSliceTypes {
				if src.MagneticVariation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevationTimeSliceType", "MagneticVariation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTimeSliceTypes {
				if src.MagneticVariation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathTimeSliceType", "MagneticVariation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSliceTypes {
				if src.MagneticVariation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerTimeSliceType", "MagneticVariation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSliceTypes {
				if src.Declination == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerTimeSliceType", "Declination", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTimeSliceTypes {
				if src.MagneticVariation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MarkerBeaconTimeSliceType", "MagneticVariation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NDBTimeSliceTypes {
				if src.MagneticVariation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NDBTimeSliceType", "MagneticVariation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSliceTypes {
				if src.MagneticVariation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarTimeSliceType", "MagneticVariation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				if src.MagneticVariation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "MagneticVariation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SDFTimeSliceTypes {
				if src.MagneticVariation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SDFTimeSliceType", "MagneticVariation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				if src.MagneticVariation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "MagneticVariation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TACANTimeSliceTypes {
				if src.MagneticVariation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TACANTimeSliceType", "MagneticVariation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TACANTimeSliceTypes {
				if src.Declination == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TACANTimeSliceType", "Declination", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VORTimeSliceTypes {
				if src.MagneticVariation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VORTimeSliceType", "MagneticVariation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VORTimeSliceTypes {
				if src.Declination == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VORTimeSliceType", "Declination", refNames, formGroup, probe.formStage)
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
				for _, target := range src.AbstractElevatedPointExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension", "AbstractElevatedPointExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_1s {
				for _, target := range src.AbstractAuthorityForNavaidEquipmentExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_1", "AbstractAuthorityForNavaidEquipmentExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_10s {
				for _, target := range src.AbstractGlidepathExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_10", "AbstractGlidepathExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_10s {
				for _, target := range src.AbstractNavaidEquipmentExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_10", "AbstractNavaidEquipmentExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_11s {
				for _, target := range src.AbstractLocalizerExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_11", "AbstractLocalizerExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_11s {
				for _, target := range src.AbstractNavaidEquipmentExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_11", "AbstractNavaidEquipmentExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_12s {
				for _, target := range src.AbstractMarkerBeaconExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_12", "AbstractMarkerBeaconExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_12s {
				for _, target := range src.AbstractNavaidEquipmentExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_12", "AbstractNavaidEquipmentExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_13s {
				for _, target := range src.AbstractNavaidExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_13", "AbstractNavaidExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_14s {
				for _, target := range src.AbstractNavaidComponentExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_14", "AbstractNavaidComponentExtension", refNames, formGroup, probe.formStage)
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
				for _, target := range src.AbstractNavaidEquipmentMonitoringExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_15", "AbstractNavaidEquipmentMonitoringExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_16s {
				for _, target := range src.AbstractPropertiesWithScheduleExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_16", "AbstractPropertiesWithScheduleExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_16s {
				for _, target := range src.AbstractNavaidOperationalStatusExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_16", "AbstractNavaidOperationalStatusExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_17s {
				for _, target := range src.AbstractNDBExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_17", "AbstractNDBExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_17s {
				for _, target := range src.AbstractNavaidEquipmentExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_17", "AbstractNavaidEquipmentExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_18s {
				for _, target := range src.AbstractSDFExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_18", "AbstractSDFExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_18s {
				for _, target := range src.AbstractNavaidEquipmentExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_18", "AbstractNavaidEquipmentExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_19s {
				for _, target := range src.AbstractSpecialNavigationStationExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_19", "AbstractSpecialNavigationStationExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_2s {
				for _, target := range src.AbstractAuthorityForSpecialNavigationStationExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_2", "AbstractAuthorityForSpecialNavigationStationExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_20s {
				for _, target := range src.AbstractPropertiesWithScheduleExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_20", "AbstractPropertiesWithScheduleExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_20s {
				for _, target := range src.AbstractSpecialNavigationStationStatusExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_20", "AbstractSpecialNavigationStationStatusExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_21s {
				for _, target := range src.AbstractSpecialNavigationSystemExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_21", "AbstractSpecialNavigationSystemExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_22s {
				for _, target := range src.AbstractTACANExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_22", "AbstractTACANExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_22s {
				for _, target := range src.AbstractNavaidEquipmentExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_22", "AbstractNavaidEquipmentExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_23s {
				for _, target := range src.AbstractVORExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_23", "AbstractVORExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_23s {
				for _, target := range src.AbstractNavaidEquipmentExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_23", "AbstractNavaidEquipmentExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_24s {
				for _, target := range src.AbstractAngleIndicationExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_24", "AbstractAngleIndicationExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_25s {
				for _, target := range src.AbstractDistanceIndicationExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_25", "AbstractDistanceIndicationExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_26s {
				for _, target := range src.AbstractDesignatedPointExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_26", "AbstractDesignatedPointExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_27s {
				for _, target := range src.AbstractSignificantPointInAirspaceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_27", "AbstractSignificantPointInAirspaceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_28s {
				for _, target := range src.AbstractLinguisticNoteExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_28", "AbstractLinguisticNoteExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_29s {
				for _, target := range src.AbstractNoteExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_29", "AbstractNoteExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_3s {
				for _, target := range src.AbstractAuthorityForSpecialNavigationSystemExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_3", "AbstractAuthorityForSpecialNavigationSystemExtension", refNames, formGroup, probe.formStage)
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
				for _, target := range src.AbstractTimesheetExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_34", "AbstractTimesheetExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_35s {
				for _, target := range src.AbstractPrecisionApproachRadarExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_35", "AbstractPrecisionApproachRadarExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_35s {
				for _, target := range src.AbstractRadarEquipmentExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_35", "AbstractRadarEquipmentExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_36s {
				for _, target := range src.AbstractPrimarySurveillanceRadarExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_36", "AbstractPrimarySurveillanceRadarExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_36s {
				for _, target := range src.AbstractSurveillanceRadarExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_36", "AbstractSurveillanceRadarExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_36s {
				for _, target := range src.AbstractRadarEquipmentExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_36", "AbstractRadarEquipmentExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_37s {
				for _, target := range src.AbstractRadarComponentExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_37", "AbstractRadarComponentExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_38s {
				for _, target := range src.AbstractRadarSystemExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_38", "AbstractRadarSystemExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_39s {
				for _, target := range src.AbstractReflectorExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_39", "AbstractReflectorExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_4s {
				for _, target := range src.AbstractAzimuthExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_4", "AbstractAzimuthExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_4s {
				for _, target := range src.AbstractNavaidEquipmentExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_4", "AbstractNavaidEquipmentExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_40s {
				for _, target := range src.AbstractSecondarySurveillanceRadarExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_40", "AbstractSecondarySurveillanceRadarExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_40s {
				for _, target := range src.AbstractSurveillanceRadarExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_40", "AbstractSurveillanceRadarExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_40s {
				for _, target := range src.AbstractRadarEquipmentExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_40", "AbstractRadarEquipmentExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_41s {
				for _, target := range src.AbstractSurveillanceGroundStationExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_41", "AbstractSurveillanceGroundStationExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_5s {
				for _, target := range src.AbstractCheckpointINSExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_5", "AbstractCheckpointINSExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_5s {
				for _, target := range src.AbstractNavigationSystemCheckpointExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_5", "AbstractNavigationSystemCheckpointExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_6s {
				for _, target := range src.AbstractCheckpointVORExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_6", "AbstractCheckpointVORExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_6s {
				for _, target := range src.AbstractNavigationSystemCheckpointExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_6", "AbstractNavigationSystemCheckpointExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_7s {
				for _, target := range src.AbstractDirectionFinderExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_7", "AbstractDirectionFinderExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_7s {
				for _, target := range src.AbstractNavaidEquipmentExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_7", "AbstractNavaidEquipmentExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_8s {
				for _, target := range src.AbstractDMEExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_8", "AbstractDMEExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_8s {
				for _, target := range src.AbstractNavaidEquipmentExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_8", "AbstractNavaidEquipmentExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_9s {
				for _, target := range src.AbstractElevationExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_9", "AbstractElevationExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_9s {
				for _, target := range src.AbstractNavaidEquipmentExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_9", "AbstractNavaidEquipmentExtension", refNames, formGroup, probe.formStage)
		}
	case *aixmabstractgmlobjecttypes.FeatureMetadataPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractNavaidEquipmentTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractNavaidEquipmentType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractNavigationSystemCheckpointTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractNavigationSystemCheckpointType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractRadarEquipmentTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractRadarEquipmentType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractSurveillanceRadarTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractSurveillanceRadarType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleIndicationType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointINSTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointINSType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointVORTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointVORType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DMETypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DMEType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DesignatedPointTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DesignatedPointType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectionFinderType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DistanceIndicationTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DistanceIndicationType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevationTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevationType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MarkerBeaconType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NDBTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NDBType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarSystemTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarSystemType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SDFTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SDFType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SignificantPointInAirspaceTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SignificantPointInAirspaceType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationStationType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationSystemTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationSystemType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TACANTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TACANType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VORTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VORType", "FeatureMetadata", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleIndicationTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointINSTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointINSTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointVORTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointVORTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DMETimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DMETimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DesignatedPointTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DesignatedPointTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectionFinderTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DistanceIndicationTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DistanceIndicationTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevationTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevationTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MarkerBeaconTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NDBTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NDBTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarSystemTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarSystemTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SDFTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SDFTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SignificantPointInAirspaceTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SignificantPointInAirspaceTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationStationTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationSystemTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationSystemTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TACANTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TACANTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VORTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VORTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractNavaidEquipmentTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractNavaidEquipmentType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractNavigationSystemCheckpointTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractNavigationSystemCheckpointType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractRadarEquipmentTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractRadarEquipmentType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractSurveillanceRadarTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractSurveillanceRadarType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleIndicationType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointINSTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointINSType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointVORTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointVORType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DMETypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DMEType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DesignatedPointTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DesignatedPointType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectionFinderType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DistanceIndicationTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DistanceIndicationType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevationTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevationType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MarkerBeaconType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NDBTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NDBType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarSystemTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarSystemType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SDFTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SDFType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SignificantPointInAirspaceTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SignificantPointInAirspaceType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationStationType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationSystemTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationSystemType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TACANTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TACANType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VORTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VORType", "BoundedBy", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractNavaidEquipmentTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractNavaidEquipmentType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractNavigationSystemCheckpointTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractNavigationSystemCheckpointType", "NameXSD", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractRadarEquipmentTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractRadarEquipmentType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractSurveillanceRadarTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractSurveillanceRadarType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleIndicationTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleIndicationType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForNavaidEquipmentTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForNavaidEquipmentType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForSpecialNavigationStationTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForSpecialNavigationStationType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForSpecialNavigationSystemTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForSpecialNavigationSystemType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointINSTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointINSTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointINSTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointINSType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointVORTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointVORTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointVORTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointVORType", "NameXSD", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.DMETimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DMETimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DMETypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DMEType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DesignatedPointTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DesignatedPointTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DesignatedPointTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DesignatedPointType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectionFinderTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectionFinderType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DistanceIndicationTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DistanceIndicationTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DistanceIndicationTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DistanceIndicationType", "NameXSD", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevationTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevationTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevationTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevationType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathType", "NameXSD", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MarkerBeaconTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MarkerBeaconType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NDBTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NDBTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NDBTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NDBType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidComponentTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidComponentType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidEquipmentMonitoringTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidEquipmentMonitoringType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidOperationalStatusTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidOperationalStatusType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidType", "NameXSD", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarComponentTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarComponentType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarSystemTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarSystemTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarSystemTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarSystemType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ReflectorTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ReflectorType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SDFTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SDFTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SDFTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SDFType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SignificantPointInAirspaceTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SignificantPointInAirspaceTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SignificantPointInAirspaceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SignificantPointInAirspaceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationStatusTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationStationStatusType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationStationTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationStationType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationSystemTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationSystemTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationSystemTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationSystemType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SurveillanceGroundStationTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SurveillanceGroundStationType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TACANTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TACANTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TACANTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TACANType", "NameXSD", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.VORTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VORTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VORTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VORType", "NameXSD", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractNavaidEquipmentTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractNavaidEquipmentType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractNavigationSystemCheckpointTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractNavigationSystemCheckpointType", "Identifier", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractRadarEquipmentTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractRadarEquipmentType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractSurveillanceRadarTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractSurveillanceRadarType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleIndicationTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleIndicationType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForNavaidEquipmentTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForNavaidEquipmentType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForSpecialNavigationStationTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForSpecialNavigationStationType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AuthorityForSpecialNavigationSystemTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AuthorityForSpecialNavigationSystemType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointINSTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointINSTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointINSTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointINSType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointVORTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointVORTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointVORTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointVORType", "Identifier", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.DMETimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DMETimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DMETypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DMEType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DesignatedPointTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DesignatedPointTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DesignatedPointTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DesignatedPointType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectionFinderTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectionFinderType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DistanceIndicationTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DistanceIndicationTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DistanceIndicationTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DistanceIndicationType", "Identifier", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevationTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevationTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevationTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevationType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathType", "Identifier", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MarkerBeaconTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MarkerBeaconType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NDBTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NDBTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NDBTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NDBType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidComponentTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidComponentType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidEquipmentMonitoringTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidEquipmentMonitoringType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidOperationalStatusTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidOperationalStatusType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidType", "Identifier", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarComponentTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarComponentType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarSystemTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarSystemTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarSystemTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarSystemType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ReflectorTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ReflectorType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SDFTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SDFTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SDFTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SDFType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SignificantPointInAirspaceTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SignificantPointInAirspaceTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SignificantPointInAirspaceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SignificantPointInAirspaceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationStatusTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationStationStatusType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationStationTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationStationType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationSystemTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationSystemTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationSystemTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationSystemType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SurveillanceGroundStationTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SurveillanceGroundStationType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TACANTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TACANTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TACANTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TACANType", "Identifier", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.VORTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VORTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VORTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VORType", "Identifier", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleIndicationTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AngleIndicationTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AzimuthTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointINSTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointINSTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointINSTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointINSTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointVORTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointVORTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CheckpointVORTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CheckpointVORTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DMETimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DMETimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DMETimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DMETimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DesignatedPointTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DesignatedPointTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DesignatedPointTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DesignatedPointTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectionFinderTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DirectionFinderTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DistanceIndicationTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DistanceIndicationTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.DistanceIndicationTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.DistanceIndicationTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevationTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevationTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ElevationTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ElevationTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GlidepathTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GlidepathTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LocalizerTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MarkerBeaconTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.MarkerBeaconTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NDBTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NDBTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NDBTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NDBTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NavaidTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NavaidTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrecisionApproachRadarTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PrimarySurveillanceRadarTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarSystemTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarSystemTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadarSystemTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadarSystemTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SDFTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SDFTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SDFTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SDFTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SecondarySurveillanceRadarTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SignificantPointInAirspaceTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SignificantPointInAirspaceTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SignificantPointInAirspaceTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SignificantPointInAirspaceTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationStationTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationStationTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationSystemTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationSystemTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationSystemTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialNavigationSystemTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TACANTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TACANTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.TACANTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.TACANTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VORTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VORTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.VORTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.VORTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
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
