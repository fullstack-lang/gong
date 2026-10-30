// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/xlink"
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
		StageSetAssociationFieldToForm("OrganisationAuthority", inst.OrganisationAuthority, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.OrganisationAuthorityType](), probe.formStage)
		StageSetAssociationFieldToForm("Unit", inst.Unit, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.UnitType](), probe.formStage)
		StageSetAssociationFieldToForm("AirTrafficControlService", inst.AirTrafficControlService, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirTrafficControlServiceType](), probe.formStage)
		StageSetAssociationFieldToForm("AirTrafficManagementService", inst.AirTrafficManagementService, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirTrafficManagementServiceType](), probe.formStage)
		StageSetAssociationFieldToForm("GroundTrafficControlService", inst.GroundTrafficControlService, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.GroundTrafficControlServiceType](), probe.formStage)
		StageSetAssociationFieldToForm("InformationService", inst.InformationService, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.InformationServiceType](), probe.formStage)
		StageSetAssociationFieldToForm("SearchRescueService", inst.SearchRescueService, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SearchRescueServiceType](), probe.formStage)
		StageSetAssociationFieldToForm("AircraftGroundService", inst.AircraftGroundService, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AircraftGroundServiceType](), probe.formStage)
		StageSetAssociationFieldToForm("AirportClearanceService", inst.AirportClearanceService, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportClearanceServiceType](), probe.formStage)
		StageSetAssociationFieldToForm("AirportSuppliesService", inst.AirportSuppliesService, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportSuppliesServiceType](), probe.formStage)
		StageSetAssociationFieldToForm("FireFightingService", inst.FireFightingService, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FireFightingServiceType](), probe.formStage)
		StageSetAssociationFieldToForm("PassengerService", inst.PassengerService, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PassengerServiceType](), probe.formStage)
		StageSetAssociationFieldToForm("PilotControlledLighting", inst.PilotControlledLighting, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PilotControlledLightingType](), probe.formStage)
		StageSetAssociationFieldToForm("RadioCommunicationChannel", inst.RadioCommunicationChannel, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RadioCommunicationChannelType](), probe.formStage)
		StageSetAssociationFieldToForm("SpecialDate", inst.SpecialDate, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SpecialDateType](), probe.formStage)
		StageSetAssociationFieldToForm("StandardLevelTable", inst.StandardLevelTable, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.StandardLevelTableType](), probe.formStage)
		StageSetAssociationFieldToForm("StandardLevelColumn", inst.StandardLevelColumn, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.StandardLevelColumnType](), probe.formStage)
		StageSetAssociationFieldToForm("StandardLevelSector", inst.StandardLevelSector, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.StandardLevelSectorType](), probe.formStage)
		StageSetAssociationFieldToForm("StandardLevel", inst.StandardLevel, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.StandardLevelType](), probe.formStage)

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
			// Slice of pointers: AbstractLightActivationExtension
			div := (&form.FormDiv{Name: "AbstractLightActivationExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractLightActivationExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractLightActivationExtension",
				Label: "AbstractLightActivationExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightActivationTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LightActivationType", "Extension", refNames, formGroup, probe.formStage)
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
			// Slice of pointers: AbstractAircraftGroundServiceExtension
			div := (&form.FormDiv{Name: "AbstractAircraftGroundServiceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAircraftGroundServiceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAircraftGroundServiceExtension",
				Label: "AbstractAircraftGroundServiceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractAirportGroundServiceExtension
			div := (&form.FormDiv{Name: "AbstractAirportGroundServiceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAirportGroundServiceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAirportGroundServiceExtension",
				Label: "AbstractAirportGroundServiceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractServiceExtension
			div := (&form.FormDiv{Name: "AbstractServiceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractServiceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractServiceExtension",
				Label: "AbstractServiceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftGroundServiceTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftGroundServiceTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_11:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractAirportClearanceServiceExtension
			div := (&form.FormDiv{Name: "AbstractAirportClearanceServiceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAirportClearanceServiceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAirportClearanceServiceExtension",
				Label: "AbstractAirportClearanceServiceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractAirportGroundServiceExtension
			div := (&form.FormDiv{Name: "AbstractAirportGroundServiceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAirportGroundServiceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAirportGroundServiceExtension",
				Label: "AbstractAirportGroundServiceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractServiceExtension
			div := (&form.FormDiv{Name: "AbstractServiceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractServiceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractServiceExtension",
				Label: "AbstractServiceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportClearanceServiceTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportClearanceServiceTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_12:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractAirportSuppliesServiceExtension
			div := (&form.FormDiv{Name: "AbstractAirportSuppliesServiceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAirportSuppliesServiceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAirportSuppliesServiceExtension",
				Label: "AbstractAirportSuppliesServiceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractAirportGroundServiceExtension
			div := (&form.FormDiv{Name: "AbstractAirportGroundServiceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAirportGroundServiceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAirportGroundServiceExtension",
				Label: "AbstractAirportGroundServiceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractServiceExtension
			div := (&form.FormDiv{Name: "AbstractServiceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractServiceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractServiceExtension",
				Label: "AbstractServiceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportSuppliesServiceTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportSuppliesServiceTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_13:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractAirTrafficControlServiceExtension
			div := (&form.FormDiv{Name: "AbstractAirTrafficControlServiceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAirTrafficControlServiceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAirTrafficControlServiceExtension",
				Label: "AbstractAirTrafficControlServiceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractTrafficSeparationServiceExtension
			div := (&form.FormDiv{Name: "AbstractTrafficSeparationServiceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractTrafficSeparationServiceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractTrafficSeparationServiceExtension",
				Label: "AbstractTrafficSeparationServiceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractServiceExtension
			div := (&form.FormDiv{Name: "AbstractServiceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractServiceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractServiceExtension",
				Label: "AbstractServiceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_14:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractAirTrafficManagementServiceExtension
			div := (&form.FormDiv{Name: "AbstractAirTrafficManagementServiceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAirTrafficManagementServiceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAirTrafficManagementServiceExtension",
				Label: "AbstractAirTrafficManagementServiceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractServiceExtension
			div := (&form.FormDiv{Name: "AbstractServiceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractServiceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractServiceExtension",
				Label: "AbstractServiceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficManagementServiceTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficManagementServiceTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_15:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractCallsignDetailExtension
			div := (&form.FormDiv{Name: "AbstractCallsignDetailExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractCallsignDetailExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractCallsignDetailExtension",
				Label: "AbstractCallsignDetailExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CallsignDetailTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CallsignDetailType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_16:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractFireFightingServiceExtension
			div := (&form.FormDiv{Name: "AbstractFireFightingServiceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractFireFightingServiceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractFireFightingServiceExtension",
				Label: "AbstractFireFightingServiceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractAirportGroundServiceExtension
			div := (&form.FormDiv{Name: "AbstractAirportGroundServiceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAirportGroundServiceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAirportGroundServiceExtension",
				Label: "AbstractAirportGroundServiceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractServiceExtension
			div := (&form.FormDiv{Name: "AbstractServiceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractServiceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractServiceExtension",
				Label: "AbstractServiceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FireFightingServiceTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FireFightingServiceTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_17:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractFuelExtension
			div := (&form.FormDiv{Name: "AbstractFuelExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractFuelExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractFuelExtension",
				Label: "AbstractFuelExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FuelTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FuelType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_18:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractGroundTrafficControlServiceExtension
			div := (&form.FormDiv{Name: "AbstractGroundTrafficControlServiceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractGroundTrafficControlServiceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractGroundTrafficControlServiceExtension",
				Label: "AbstractGroundTrafficControlServiceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractTrafficSeparationServiceExtension
			div := (&form.FormDiv{Name: "AbstractTrafficSeparationServiceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractTrafficSeparationServiceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractTrafficSeparationServiceExtension",
				Label: "AbstractTrafficSeparationServiceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractServiceExtension
			div := (&form.FormDiv{Name: "AbstractServiceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractServiceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractServiceExtension",
				Label: "AbstractServiceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GroundTrafficControlServiceTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GroundTrafficControlServiceTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_19:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractInformationServiceExtension
			div := (&form.FormDiv{Name: "AbstractInformationServiceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractInformationServiceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractInformationServiceExtension",
				Label: "AbstractInformationServiceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractServiceExtension
			div := (&form.FormDiv{Name: "AbstractServiceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractServiceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractServiceExtension",
				Label: "AbstractServiceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_2:
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
	case *aixmfeatures.A_extension_20:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractNitrogenExtension
			div := (&form.FormDiv{Name: "AbstractNitrogenExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractNitrogenExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractNitrogenExtension",
				Label: "AbstractNitrogenExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NitrogenTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NitrogenType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_21:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractOilExtension
			div := (&form.FormDiv{Name: "AbstractOilExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractOilExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractOilExtension",
				Label: "AbstractOilExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OilTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OilType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_22:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractOxygenExtension
			div := (&form.FormDiv{Name: "AbstractOxygenExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractOxygenExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractOxygenExtension",
				Label: "AbstractOxygenExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OxygenTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OxygenType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_23:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractPassengerServiceExtension
			div := (&form.FormDiv{Name: "AbstractPassengerServiceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractPassengerServiceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractPassengerServiceExtension",
				Label: "AbstractPassengerServiceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractAirportGroundServiceExtension
			div := (&form.FormDiv{Name: "AbstractAirportGroundServiceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractAirportGroundServiceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractAirportGroundServiceExtension",
				Label: "AbstractAirportGroundServiceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractServiceExtension
			div := (&form.FormDiv{Name: "AbstractServiceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractServiceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractServiceExtension",
				Label: "AbstractServiceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PassengerServiceTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PassengerServiceTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_24:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractPilotControlledLightingExtension
			div := (&form.FormDiv{Name: "AbstractPilotControlledLightingExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractPilotControlledLightingExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractPilotControlledLightingExtension",
				Label: "AbstractPilotControlledLightingExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PilotControlledLightingTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PilotControlledLightingTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_25:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractRadioCommunicationChannelExtension
			div := (&form.FormDiv{Name: "AbstractRadioCommunicationChannelExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractRadioCommunicationChannelExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractRadioCommunicationChannelExtension",
				Label: "AbstractRadioCommunicationChannelExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationChannelTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationChannelTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_26:
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
			// Slice of pointers: AbstractRadioCommunicationOperationalStatusExtension
			div := (&form.FormDiv{Name: "AbstractRadioCommunicationOperationalStatusExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractRadioCommunicationOperationalStatusExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractRadioCommunicationOperationalStatusExtension",
				Label: "AbstractRadioCommunicationOperationalStatusExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationOperationalStatusTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationOperationalStatusType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_27:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractSearchRescueServiceExtension
			div := (&form.FormDiv{Name: "AbstractSearchRescueServiceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractSearchRescueServiceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractSearchRescueServiceExtension",
				Label: "AbstractSearchRescueServiceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AbstractServiceExtension
			div := (&form.FormDiv{Name: "AbstractServiceExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractServiceExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractServiceExtension",
				Label: "AbstractServiceExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SearchRescueServiceTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SearchRescueServiceTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_28:
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
			// Slice of pointers: AbstractServiceOperationalStatusExtension
			div := (&form.FormDiv{Name: "AbstractServiceOperationalStatusExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractServiceOperationalStatusExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractServiceOperationalStatusExtension",
				Label: "AbstractServiceOperationalStatusExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ServiceOperationalStatusTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ServiceOperationalStatusType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_29:
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
	case *aixmfeatures.A_extension_3:
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
	case *aixmfeatures.A_extension_30:
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
	case *aixmfeatures.A_extension_33:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractSpecialDateExtension
			div := (&form.FormDiv{Name: "AbstractSpecialDateExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractSpecialDateExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractSpecialDateExtension",
				Label: "AbstractSpecialDateExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialDateTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialDateTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
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
			// Slice of pointers: AbstractStandardLevelExtension
			div := (&form.FormDiv{Name: "AbstractStandardLevelExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractStandardLevelExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractStandardLevelExtension",
				Label: "AbstractStandardLevelExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_36:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractStandardLevelColumnExtension
			div := (&form.FormDiv{Name: "AbstractStandardLevelColumnExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractStandardLevelColumnExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractStandardLevelColumnExtension",
				Label: "AbstractStandardLevelColumnExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelColumnTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelColumnTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_37:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractStandardLevelSectorExtension
			div := (&form.FormDiv{Name: "AbstractStandardLevelSectorExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractStandardLevelSectorExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractStandardLevelSectorExtension",
				Label: "AbstractStandardLevelSectorExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelSectorTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelSectorTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_38:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractStandardLevelTableExtension
			div := (&form.FormDiv{Name: "AbstractStandardLevelTableExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractStandardLevelTableExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractStandardLevelTableExtension",
				Label: "AbstractStandardLevelTableExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelTableTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelTableTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_4:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractOrganisationAuthorityExtension
			div := (&form.FormDiv{Name: "AbstractOrganisationAuthorityExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractOrganisationAuthorityExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractOrganisationAuthorityExtension",
				Label: "AbstractOrganisationAuthorityExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OrganisationAuthorityTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OrganisationAuthorityTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_5:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractOrganisationAuthorityAssociationExtension
			div := (&form.FormDiv{Name: "AbstractOrganisationAuthorityAssociationExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractOrganisationAuthorityAssociationExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractOrganisationAuthorityAssociationExtension",
				Label: "AbstractOrganisationAuthorityAssociationExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OrganisationAuthorityAssociationTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OrganisationAuthorityAssociationType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_6:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractUnitExtension
			div := (&form.FormDiv{Name: "AbstractUnitExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractUnitExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractUnitExtension",
				Label: "AbstractUnitExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitTimeSliceTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitTimeSliceType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_7:
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
			// Slice of pointers: AbstractUnitAvailabilityExtension
			div := (&form.FormDiv{Name: "AbstractUnitAvailabilityExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractUnitAvailabilityExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractUnitAvailabilityExtension",
				Label: "AbstractUnitAvailabilityExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitAvailabilityTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitAvailabilityType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_8:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			// Slice of pointers: AbstractUnitDependencyExtension
			div := (&form.FormDiv{Name: "AbstractUnitDependencyExtension"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AbstractUnitDependencyExtension {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AbstractUnitDependencyExtension",
				Label: "AbstractUnitDependencyExtension",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitDependencyTypes {
				for _, target := range src.Extension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitDependencyType", "Extension", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.A_extension_9:
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
	case *aixmfeatures.AbstractAirportGroundServiceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
	case *aixmfeatures.AbstractServiceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
	case *aixmfeatures.AbstractTrafficSeparationServiceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTimeSliceTypes {
				for _, target := range src.ClientAerialRefuelling {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceTimeSliceType", "ClientAerialRefuelling", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficManagementServiceTimeSliceTypes {
				for _, target := range src.ClientAerialRefuelling {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficManagementServiceTimeSliceType", "ClientAerialRefuelling", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTimeSliceTypes {
				for _, target := range src.ClientAerialRefuelling {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceTimeSliceType", "ClientAerialRefuelling", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirTrafficControlServiceTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AirTrafficControlServiceTimeSlice", inst.AirTrafficControlServiceTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirTrafficControlServiceTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirTrafficControlServiceTimeSliceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
		StageSetAssociationFieldToForm("FlightOperations", inst.FlightOperations, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlightDestinationType](), probe.formStage)
		StageSetAssociationFieldToForm("Rank", inst.Rank, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFacilityRankingType](), probe.formStage)
		StageSetAssociationFieldToForm("CompliantICAO", inst.CompliantICAO, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Location", inst.Location, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("ServiceProvider", inst.ServiceProvider, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.UnitPropertyType](), probe.formStage)

		{
			// Slice of pointers: Call_sign
			div := (&form.FormDiv{Name: "Call_sign"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Call_sign {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Call_sign",
				Label: "Call_sign",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: RadioCommunication
			div := (&form.FormDiv{Name: "RadioCommunication"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RadioCommunication {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RadioCommunication",
				Label: "RadioCommunication",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: GroundCommunication
			div := (&form.FormDiv{Name: "GroundCommunication"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GroundCommunication {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GroundCommunication",
				Label: "GroundCommunication",
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
		StageSetAssociationFieldToForm("RadarAssisted", inst.RadarAssisted, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("DataLinkEnabled", inst.DataLinkEnabled, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("DataLinkChannel", inst.DataLinkChannel, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCommunicationChannelType](), probe.formStage)
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeServiceATCType](), probe.formStage)

		{
			// Slice of pointers: ClientAirport
			div := (&form.FormDiv{Name: "ClientAirport"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ClientAirport {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ClientAirport",
				Label: "ClientAirport",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ClientAirspace
			div := (&form.FormDiv{Name: "ClientAirspace"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ClientAirspace {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ClientAirspace",
				Label: "ClientAirspace",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ClientRoute
			div := (&form.FormDiv{Name: "ClientRoute"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ClientRoute {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ClientRoute",
				Label: "ClientRoute",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ClientProcedure
			div := (&form.FormDiv{Name: "ClientProcedure"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ClientProcedure {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ClientProcedure",
				Label: "ClientProcedure",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ClientHolding
			div := (&form.FormDiv{Name: "ClientHolding"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ClientHolding {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ClientHolding",
				Label: "ClientHolding",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ClientAerialRefuelling
			div := (&form.FormDiv{Name: "ClientAerialRefuelling"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ClientAerialRefuelling {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ClientAerialRefuelling",
				Label: "ClientAerialRefuelling",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("AircraftLocator", inst.AircraftLocator, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DirectionFinderPropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTimeSlicePropertyTypes {
				if src.AirTrafficControlServiceTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceTimeSlicePropertyType", "AirTrafficControlServiceTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirTrafficControlServiceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
				if src.AirTrafficControlService == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "AirTrafficControlService", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirTrafficManagementServiceTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AirTrafficManagementServiceTimeSlice", inst.AirTrafficManagementServiceTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirTrafficManagementServiceTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficManagementServiceTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficManagementServiceType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirTrafficManagementServiceTimeSliceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
		StageSetAssociationFieldToForm("FlightOperations", inst.FlightOperations, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlightDestinationType](), probe.formStage)
		StageSetAssociationFieldToForm("Rank", inst.Rank, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFacilityRankingType](), probe.formStage)
		StageSetAssociationFieldToForm("CompliantICAO", inst.CompliantICAO, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Location", inst.Location, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("ServiceProvider", inst.ServiceProvider, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.UnitPropertyType](), probe.formStage)

		{
			// Slice of pointers: Call_sign
			div := (&form.FormDiv{Name: "Call_sign"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Call_sign {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Call_sign",
				Label: "Call_sign",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: RadioCommunication
			div := (&form.FormDiv{Name: "RadioCommunication"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RadioCommunication {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RadioCommunication",
				Label: "RadioCommunication",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: GroundCommunication
			div := (&form.FormDiv{Name: "GroundCommunication"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GroundCommunication {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GroundCommunication",
				Label: "GroundCommunication",
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
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeServiceATFMType](), probe.formStage)

		{
			// Slice of pointers: ClientAirspace
			div := (&form.FormDiv{Name: "ClientAirspace"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ClientAirspace {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ClientAirspace",
				Label: "ClientAirspace",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ClientAerialRefuelling
			div := (&form.FormDiv{Name: "ClientAerialRefuelling"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ClientAerialRefuelling {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ClientAerialRefuelling",
				Label: "ClientAerialRefuelling",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ClientRoute
			div := (&form.FormDiv{Name: "ClientRoute"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ClientRoute {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ClientRoute",
				Label: "ClientRoute",
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficManagementServiceTimeSlicePropertyTypes {
				if src.AirTrafficManagementServiceTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficManagementServiceTimeSlicePropertyType", "AirTrafficManagementServiceTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirTrafficManagementServiceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
				if src.AirTrafficManagementService == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "AirTrafficManagementService", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AircraftGroundServiceTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AircraftGroundServiceTimeSlice", inst.AircraftGroundServiceTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AircraftGroundServiceTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftGroundServiceTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftGroundServiceType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AircraftGroundServiceTimeSliceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
		StageSetAssociationFieldToForm("FlightOperations", inst.FlightOperations, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlightDestinationType](), probe.formStage)
		StageSetAssociationFieldToForm("Rank", inst.Rank, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFacilityRankingType](), probe.formStage)
		StageSetAssociationFieldToForm("CompliantICAO", inst.CompliantICAO, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Location", inst.Location, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("ServiceProvider", inst.ServiceProvider, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.UnitPropertyType](), probe.formStage)

		{
			// Slice of pointers: Call_sign
			div := (&form.FormDiv{Name: "Call_sign"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Call_sign {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Call_sign",
				Label: "Call_sign",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: RadioCommunication
			div := (&form.FormDiv{Name: "RadioCommunication"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RadioCommunication {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RadioCommunication",
				Label: "RadioCommunication",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: GroundCommunication
			div := (&form.FormDiv{Name: "GroundCommunication"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GroundCommunication {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GroundCommunication",
				Label: "GroundCommunication",
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
			// Slice of pointers: AirportHeliport
			div := (&form.FormDiv{Name: "AirportHeliport"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AirportHeliport {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AirportHeliport",
				Label: "AirportHeliport",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAircraftGroundServiceType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftGroundServiceTimeSlicePropertyTypes {
				if src.AircraftGroundServiceTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftGroundServiceTimeSlicePropertyType", "AircraftGroundServiceTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AircraftGroundServiceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
				if src.AircraftGroundService == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "AircraftGroundService", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirportClearanceServiceTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AirportClearanceServiceTimeSlice", inst.AirportClearanceServiceTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportClearanceServiceTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportClearanceServiceTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportClearanceServiceType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirportClearanceServiceTimeSliceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
		StageSetAssociationFieldToForm("FlightOperations", inst.FlightOperations, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlightDestinationType](), probe.formStage)
		StageSetAssociationFieldToForm("Rank", inst.Rank, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFacilityRankingType](), probe.formStage)
		StageSetAssociationFieldToForm("CompliantICAO", inst.CompliantICAO, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Location", inst.Location, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("ServiceProvider", inst.ServiceProvider, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.UnitPropertyType](), probe.formStage)

		{
			// Slice of pointers: Call_sign
			div := (&form.FormDiv{Name: "Call_sign"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Call_sign {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Call_sign",
				Label: "Call_sign",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: RadioCommunication
			div := (&form.FormDiv{Name: "RadioCommunication"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RadioCommunication {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RadioCommunication",
				Label: "RadioCommunication",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: GroundCommunication
			div := (&form.FormDiv{Name: "GroundCommunication"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GroundCommunication {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GroundCommunication",
				Label: "GroundCommunication",
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
			// Slice of pointers: AirportHeliport
			div := (&form.FormDiv{Name: "AirportHeliport"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AirportHeliport {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AirportHeliport",
				Label: "AirportHeliport",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("SnowPlan", inst.SnowPlan, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextInstructionType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportClearanceServiceTimeSlicePropertyTypes {
				if src.AirportClearanceServiceTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportClearanceServiceTimeSlicePropertyType", "AirportClearanceServiceTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirportClearanceServiceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
				if src.AirportClearanceService == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "AirportClearanceService", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTimeSliceTypes {
				for _, target := range src.ClientAirport {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceTimeSliceType", "ClientAirport", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftGroundServiceTimeSliceTypes {
				for _, target := range src.AirportHeliport {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftGroundServiceTimeSliceType", "AirportHeliport", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportClearanceServiceTimeSliceTypes {
				for _, target := range src.AirportHeliport {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportClearanceServiceTimeSliceType", "AirportHeliport", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportSuppliesServiceTimeSliceTypes {
				for _, target := range src.AirportHeliport {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportSuppliesServiceTimeSliceType", "AirportHeliport", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FireFightingServiceTimeSliceTypes {
				for _, target := range src.AirportHeliport {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FireFightingServiceTimeSliceType", "AirportHeliport", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GroundTrafficControlServiceTimeSliceTypes {
				if src.ClientAirport == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GroundTrafficControlServiceTimeSliceType", "ClientAirport", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTimeSliceTypes {
				for _, target := range src.ClientAirport {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceTimeSliceType", "ClientAirport", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PassengerServiceTimeSliceTypes {
				for _, target := range src.AirportHeliport {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PassengerServiceTimeSliceType", "AirportHeliport", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitTimeSliceTypes {
				if src.AirportLocation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitTimeSliceType", "AirportLocation", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirportSuppliesServiceTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AirportSuppliesServiceTimeSlice", inst.AirportSuppliesServiceTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportSuppliesServiceTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportSuppliesServiceTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportSuppliesServiceType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirportSuppliesServiceTimeSliceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
		StageSetAssociationFieldToForm("FlightOperations", inst.FlightOperations, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlightDestinationType](), probe.formStage)
		StageSetAssociationFieldToForm("Rank", inst.Rank, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFacilityRankingType](), probe.formStage)
		StageSetAssociationFieldToForm("CompliantICAO", inst.CompliantICAO, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Location", inst.Location, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("ServiceProvider", inst.ServiceProvider, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.UnitPropertyType](), probe.formStage)

		{
			// Slice of pointers: Call_sign
			div := (&form.FormDiv{Name: "Call_sign"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Call_sign {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Call_sign",
				Label: "Call_sign",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: RadioCommunication
			div := (&form.FormDiv{Name: "RadioCommunication"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RadioCommunication {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RadioCommunication",
				Label: "RadioCommunication",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: GroundCommunication
			div := (&form.FormDiv{Name: "GroundCommunication"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GroundCommunication {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GroundCommunication",
				Label: "GroundCommunication",
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
			// Slice of pointers: AirportHeliport
			div := (&form.FormDiv{Name: "AirportHeliport"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AirportHeliport {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AirportHeliport",
				Label: "AirportHeliport",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: FuelSupply
			div := (&form.FormDiv{Name: "FuelSupply"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.FuelSupply {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "FuelSupply",
				Label: "FuelSupply",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: OilSupply
			div := (&form.FormDiv{Name: "OilSupply"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.OilSupply {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "OilSupply",
				Label: "OilSupply",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: NitrogenSupply
			div := (&form.FormDiv{Name: "NitrogenSupply"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NitrogenSupply {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NitrogenSupply",
				Label: "NitrogenSupply",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: OxygenSupply
			div := (&form.FormDiv{Name: "OxygenSupply"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.OxygenSupply {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "OxygenSupply",
				Label: "OxygenSupply",
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportSuppliesServiceTimeSlicePropertyTypes {
				if src.AirportSuppliesServiceTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportSuppliesServiceTimeSlicePropertyType", "AirportSuppliesServiceTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.AirportSuppliesServiceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
				if src.AirportSuppliesService == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "AirportSuppliesService", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTimeSliceTypes {
				for _, target := range src.ClientAirspace {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceTimeSliceType", "ClientAirspace", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficManagementServiceTimeSliceTypes {
				for _, target := range src.ClientAirspace {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficManagementServiceTimeSliceType", "ClientAirspace", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTimeSliceTypes {
				for _, target := range src.ClientAirspace {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceTimeSliceType", "ClientAirspace", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SearchRescueServiceTimeSliceTypes {
				for _, target := range src.ClientAirspace {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SearchRescueServiceTimeSliceType", "ClientAirspace", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelSectorTimeSliceTypes {
				for _, target := range src.ApplicableAirspace {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelSectorTimeSliceType", "ApplicableAirspace", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.CallsignDetailPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("CallsignDetail", inst.CallsignDetail, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CallsignDetailType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTimeSliceTypes {
				for _, target := range src.Call_sign {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceTimeSliceType", "Call_sign", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficManagementServiceTimeSliceTypes {
				for _, target := range src.Call_sign {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficManagementServiceTimeSliceType", "Call_sign", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftGroundServiceTimeSliceTypes {
				for _, target := range src.Call_sign {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftGroundServiceTimeSliceType", "Call_sign", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportClearanceServiceTimeSliceTypes {
				for _, target := range src.Call_sign {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportClearanceServiceTimeSliceType", "Call_sign", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportSuppliesServiceTimeSliceTypes {
				for _, target := range src.Call_sign {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportSuppliesServiceTimeSliceType", "Call_sign", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FireFightingServiceTimeSliceTypes {
				for _, target := range src.Call_sign {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FireFightingServiceTimeSliceType", "Call_sign", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GroundTrafficControlServiceTimeSliceTypes {
				for _, target := range src.Call_sign {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GroundTrafficControlServiceTimeSliceType", "Call_sign", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTimeSliceTypes {
				for _, target := range src.Call_sign {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceTimeSliceType", "Call_sign", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PassengerServiceTimeSliceTypes {
				for _, target := range src.Call_sign {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PassengerServiceTimeSliceType", "Call_sign", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SearchRescueServiceTimeSliceTypes {
				for _, target := range src.Call_sign {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SearchRescueServiceTimeSliceType", "Call_sign", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.CallsignDetailType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("CallSign", inst.CallSign, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextNameType](), probe.formStage)
		StageSetAssociationFieldToForm("Language", inst.Language, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeLanguageType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.CallsignDetailPropertyTypes {
				if src.CallsignDetail == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CallsignDetailPropertyType", "CallsignDetail", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ContactInformationPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ContactInformation", inst.ContactInformation, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ContactInformationType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTimeSliceTypes {
				for _, target := range src.GroundCommunication {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceTimeSliceType", "GroundCommunication", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficManagementServiceTimeSliceTypes {
				for _, target := range src.GroundCommunication {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficManagementServiceTimeSliceType", "GroundCommunication", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftGroundServiceTimeSliceTypes {
				for _, target := range src.GroundCommunication {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftGroundServiceTimeSliceType", "GroundCommunication", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportClearanceServiceTimeSliceTypes {
				for _, target := range src.GroundCommunication {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportClearanceServiceTimeSliceType", "GroundCommunication", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportSuppliesServiceTimeSliceTypes {
				for _, target := range src.GroundCommunication {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportSuppliesServiceTimeSliceType", "GroundCommunication", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FireFightingServiceTimeSliceTypes {
				for _, target := range src.GroundCommunication {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FireFightingServiceTimeSliceType", "GroundCommunication", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GroundTrafficControlServiceTimeSliceTypes {
				for _, target := range src.GroundCommunication {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GroundTrafficControlServiceTimeSliceType", "GroundCommunication", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTimeSliceTypes {
				for _, target := range src.GroundCommunication {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceTimeSliceType", "GroundCommunication", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OrganisationAuthorityTimeSliceTypes {
				for _, target := range src.Contact {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OrganisationAuthorityTimeSliceType", "Contact", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PassengerServiceTimeSliceTypes {
				for _, target := range src.GroundCommunication {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PassengerServiceTimeSliceType", "GroundCommunication", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SearchRescueServiceTimeSliceTypes {
				for _, target := range src.GroundCommunication {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SearchRescueServiceTimeSliceType", "GroundCommunication", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitTimeSliceTypes {
				for _, target := range src.Contact {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitTimeSliceType", "Contact", refNames, formGroup, probe.formStage)
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
	case *aixmfeatures.DirectionFinderPropertyType:
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTimeSliceTypes {
				if src.AircraftLocator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceTimeSliceType", "AircraftLocator", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ElevatedPointPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ElevatedPoint", inst.ElevatedPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTimeSliceTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceTimeSliceType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficManagementServiceTimeSliceTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficManagementServiceTimeSliceType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftGroundServiceTimeSliceTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftGroundServiceTimeSliceType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportClearanceServiceTimeSliceTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportClearanceServiceTimeSliceType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportSuppliesServiceTimeSliceTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportSuppliesServiceTimeSliceType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FireFightingServiceTimeSliceTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FireFightingServiceTimeSliceType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GroundTrafficControlServiceTimeSliceTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GroundTrafficControlServiceTimeSliceType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTimeSliceTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceTimeSliceType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PassengerServiceTimeSliceTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PassengerServiceTimeSliceType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationChannelTimeSliceTypes {
				for _, target := range src.Location {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationChannelTimeSliceType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SearchRescueServiceTimeSliceTypes {
				if src.Location == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SearchRescueServiceTimeSliceType", "Location", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitTimeSliceTypes {
				if src.Position == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitTimeSliceType", "Position", refNames, formGroup, probe.formStage)
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
	case *aixmfeatures.FireFightingServiceTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("FireFightingServiceTimeSlice", inst.FireFightingServiceTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FireFightingServiceTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FireFightingServiceTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FireFightingServiceType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.FireFightingServiceTimeSliceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
		StageSetAssociationFieldToForm("FlightOperations", inst.FlightOperations, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlightDestinationType](), probe.formStage)
		StageSetAssociationFieldToForm("Rank", inst.Rank, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFacilityRankingType](), probe.formStage)
		StageSetAssociationFieldToForm("CompliantICAO", inst.CompliantICAO, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Location", inst.Location, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("ServiceProvider", inst.ServiceProvider, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.UnitPropertyType](), probe.formStage)

		{
			// Slice of pointers: Call_sign
			div := (&form.FormDiv{Name: "Call_sign"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Call_sign {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Call_sign",
				Label: "Call_sign",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: RadioCommunication
			div := (&form.FormDiv{Name: "RadioCommunication"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RadioCommunication {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RadioCommunication",
				Label: "RadioCommunication",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: GroundCommunication
			div := (&form.FormDiv{Name: "GroundCommunication"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GroundCommunication {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GroundCommunication",
				Label: "GroundCommunication",
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
			// Slice of pointers: AirportHeliport
			div := (&form.FormDiv{Name: "AirportHeliport"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AirportHeliport {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AirportHeliport",
				Label: "AirportHeliport",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Category", inst.Category, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFireFightingType](), probe.formStage)
		StageSetAssociationFieldToForm("Standard", inst.Standard, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeAviationStandardsType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.FireFightingServiceTimeSlicePropertyTypes {
				if src.FireFightingServiceTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FireFightingServiceTimeSlicePropertyType", "FireFightingServiceTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.FireFightingServiceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
				if src.FireFightingService == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "FireFightingService", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.FuelPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Fuel", inst.Fuel, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.FuelType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportSuppliesServiceTimeSliceTypes {
				for _, target := range src.FuelSupply {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportSuppliesServiceTimeSliceType", "FuelSupply", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.FuelType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Category", inst.Category, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFuelType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.FuelPropertyTypes {
				if src.Fuel == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FuelPropertyType", "Fuel", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.PilotControlledLightingTimeSliceTypes {
				for _, target := range src.ActivatedGroundLighting {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PilotControlledLightingTimeSliceType", "ActivatedGroundLighting", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.GroundTrafficControlServiceTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("GroundTrafficControlServiceTimeSlice", inst.GroundTrafficControlServiceTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.GroundTrafficControlServiceTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GroundTrafficControlServiceTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GroundTrafficControlServiceType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.GroundTrafficControlServiceTimeSliceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
		StageSetAssociationFieldToForm("FlightOperations", inst.FlightOperations, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlightDestinationType](), probe.formStage)
		StageSetAssociationFieldToForm("Rank", inst.Rank, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFacilityRankingType](), probe.formStage)
		StageSetAssociationFieldToForm("CompliantICAO", inst.CompliantICAO, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Location", inst.Location, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("ServiceProvider", inst.ServiceProvider, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.UnitPropertyType](), probe.formStage)

		{
			// Slice of pointers: Call_sign
			div := (&form.FormDiv{Name: "Call_sign"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Call_sign {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Call_sign",
				Label: "Call_sign",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: RadioCommunication
			div := (&form.FormDiv{Name: "RadioCommunication"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RadioCommunication {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RadioCommunication",
				Label: "RadioCommunication",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: GroundCommunication
			div := (&form.FormDiv{Name: "GroundCommunication"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GroundCommunication {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GroundCommunication",
				Label: "GroundCommunication",
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
		StageSetAssociationFieldToForm("RadarAssisted", inst.RadarAssisted, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("DataLinkEnabled", inst.DataLinkEnabled, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("DataLinkChannel", inst.DataLinkChannel, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCommunicationChannelType](), probe.formStage)
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeServiceGroundControlType](), probe.formStage)
		StageSetAssociationFieldToForm("ClientAirport", inst.ClientAirport, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.GroundTrafficControlServiceTimeSlicePropertyTypes {
				if src.GroundTrafficControlServiceTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GroundTrafficControlServiceTimeSlicePropertyType", "GroundTrafficControlServiceTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.GroundTrafficControlServiceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
				if src.GroundTrafficControlService == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "GroundTrafficControlService", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.HoldingPatternPropertyType:
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTimeSliceTypes {
				for _, target := range src.ClientHolding {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceTimeSliceType", "ClientHolding", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTimeSliceTypes {
				for _, target := range src.ClientHolding {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceTimeSliceType", "ClientHolding", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.InformationServiceTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("InformationServiceTimeSlice", inst.InformationServiceTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.InformationServiceTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.InformationServiceTimeSliceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
		StageSetAssociationFieldToForm("FlightOperations", inst.FlightOperations, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlightDestinationType](), probe.formStage)
		StageSetAssociationFieldToForm("Rank", inst.Rank, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFacilityRankingType](), probe.formStage)
		StageSetAssociationFieldToForm("CompliantICAO", inst.CompliantICAO, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Location", inst.Location, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("ServiceProvider", inst.ServiceProvider, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.UnitPropertyType](), probe.formStage)

		{
			// Slice of pointers: Call_sign
			div := (&form.FormDiv{Name: "Call_sign"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Call_sign {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Call_sign",
				Label: "Call_sign",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: RadioCommunication
			div := (&form.FormDiv{Name: "RadioCommunication"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RadioCommunication {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RadioCommunication",
				Label: "RadioCommunication",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: GroundCommunication
			div := (&form.FormDiv{Name: "GroundCommunication"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GroundCommunication {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GroundCommunication",
				Label: "GroundCommunication",
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
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeServiceInformationType](), probe.formStage)
		StageSetAssociationFieldToForm("Voice", inst.Voice, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("DataLink", inst.DataLink, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Recorded", inst.Recorded, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)

		{
			// Slice of pointers: NavaidBroadcast
			div := (&form.FormDiv{Name: "NavaidBroadcast"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NavaidBroadcast {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NavaidBroadcast",
				Label: "NavaidBroadcast",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ClientAirspace
			div := (&form.FormDiv{Name: "ClientAirspace"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ClientAirspace {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ClientAirspace",
				Label: "ClientAirspace",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ClientAirport
			div := (&form.FormDiv{Name: "ClientAirport"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ClientAirport {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ClientAirport",
				Label: "ClientAirport",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ClientRoute
			div := (&form.FormDiv{Name: "ClientRoute"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ClientRoute {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ClientRoute",
				Label: "ClientRoute",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ClientProcedure
			div := (&form.FormDiv{Name: "ClientProcedure"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ClientProcedure {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ClientProcedure",
				Label: "ClientProcedure",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ClientHolding
			div := (&form.FormDiv{Name: "ClientHolding"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ClientHolding {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ClientHolding",
				Label: "ClientHolding",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ClientAerialRefuelling
			div := (&form.FormDiv{Name: "ClientAerialRefuelling"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ClientAerialRefuelling {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ClientAerialRefuelling",
				Label: "ClientAerialRefuelling",
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTimeSlicePropertyTypes {
				if src.InformationServiceTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceTimeSlicePropertyType", "InformationServiceTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.InformationServiceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
				if src.InformationService == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "InformationService", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.LightActivationPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("LightActivation", inst.LightActivation, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.LightActivationType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PilotControlledLightingTimeSliceTypes {
				for _, target := range src.ControlledLightIntensity {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PilotControlledLightingTimeSliceType", "ControlledLightIntensity", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.LightActivationType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Clicks", inst.Clicks, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoNumberType](), probe.formStage)
		StageSetAssociationFieldToForm("IntensityLevel", inst.IntensityLevel, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeLightIntensityType](), probe.formStage)
		StageSetAssociationFieldToForm("Activation", inst.Activation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSystemActivationType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightActivationPropertyTypes {
				if src.LightActivation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LightActivationPropertyType", "LightActivation", refNames, formGroup, probe.formStage)
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
	case *aixmfeatures.NitrogenPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Nitrogen", inst.Nitrogen, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NitrogenType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportSuppliesServiceTimeSliceTypes {
				for _, target := range src.NitrogenSupply {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportSuppliesServiceTimeSliceType", "NitrogenSupply", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.NitrogenType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNitrogenType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.NitrogenPropertyTypes {
				if src.Nitrogen == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NitrogenPropertyType", "Nitrogen", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.NotePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Note", inst.Note, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NoteType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficManagementServiceTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficManagementServiceTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftGroundServiceTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftGroundServiceTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportClearanceServiceTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportClearanceServiceTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportSuppliesServiceTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportSuppliesServiceTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CallsignDetailTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CallsignDetailType", "Annotation", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.FireFightingServiceTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FireFightingServiceTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FuelTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FuelType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GroundTrafficControlServiceTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GroundTrafficControlServiceTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightActivationTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LightActivationType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NitrogenTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NitrogenType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OilTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OilType", "Annotation", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.OrganisationAuthorityAssociationTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OrganisationAuthorityAssociationType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OrganisationAuthorityTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OrganisationAuthorityTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OxygenTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OxygenType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PassengerServiceTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PassengerServiceTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PilotControlledLightingTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PilotControlledLightingTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationChannelTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationChannelTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationOperationalStatusTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationOperationalStatusType", "Annotation", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.SearchRescueServiceTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SearchRescueServiceTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ServiceOperationalStatusTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ServiceOperationalStatusType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialDateTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialDateTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelColumnTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelColumnTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelSectorTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelSectorTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelTableTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelTableTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelType", "Annotation", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitAvailabilityTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitAvailabilityType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitDependencyTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitDependencyType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitTimeSliceTypes {
				for _, target := range src.Annotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitTimeSliceType", "Annotation", refNames, formGroup, probe.formStage)
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
	case *aixmfeatures.OilPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Oil", inst.Oil, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.OilType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportSuppliesServiceTimeSliceTypes {
				for _, target := range src.OilSupply {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportSuppliesServiceTimeSliceType", "OilSupply", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.OilType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Category", inst.Category, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeOilType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.OilPropertyTypes {
				if src.Oil == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OilPropertyType", "Oil", refNames, formGroup, probe.formStage)
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
	case *aixmfeatures.OrganisationAuthorityAssociationPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("OrganisationAuthorityAssociation", inst.OrganisationAuthorityAssociation, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.OrganisationAuthorityAssociationType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OrganisationAuthorityTimeSliceTypes {
				for _, target := range src.RelatedOrganisationAuthority {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OrganisationAuthorityTimeSliceType", "RelatedOrganisationAuthority", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.OrganisationAuthorityAssociationType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeOrganisationHierarchyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.OrganisationAuthorityAssociationPropertyTypes {
				if src.OrganisationAuthorityAssociation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OrganisationAuthorityAssociationPropertyType", "OrganisationAuthorityAssociation", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.OrganisationAuthorityAssociationTypes {
				if src.TheOrganisationAuthority == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OrganisationAuthorityAssociationType", "TheOrganisationAuthority", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationOperationalStatusTypes {
				for _, target := range src.SpecialDateAuthority {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationOperationalStatusType", "SpecialDateAuthority", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ServiceOperationalStatusTypes {
				for _, target := range src.SpecialDateAuthority {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ServiceOperationalStatusType", "SpecialDateAuthority", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialDateTimeSliceTypes {
				if src.Authority == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialDateTimeSliceType", "Authority", refNames, formGroup, probe.formStage)
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

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitAvailabilityTypes {
				for _, target := range src.SpecialDateAuthority {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitAvailabilityType", "SpecialDateAuthority", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitTimeSliceTypes {
				if src.OwnerOrganisation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitTimeSliceType", "OwnerOrganisation", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.OrganisationAuthorityTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("OrganisationAuthorityTimeSlice", inst.OrganisationAuthorityTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.OrganisationAuthorityTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OrganisationAuthorityTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OrganisationAuthorityType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.OrganisationAuthorityTimeSliceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
		StageSetAssociationFieldToForm("Designator", inst.Designator, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeOrganisationDesignatorType](), probe.formStage)
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeOrganisationType](), probe.formStage)
		StageSetAssociationFieldToForm("Military", inst.Military, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeMilitaryOperationsType](), probe.formStage)

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
			// Slice of pointers: RelatedOrganisationAuthority
			div := (&form.FormDiv{Name: "RelatedOrganisationAuthority"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RelatedOrganisationAuthority {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RelatedOrganisationAuthority",
				Label: "RelatedOrganisationAuthority",
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.OrganisationAuthorityTimeSlicePropertyTypes {
				if src.OrganisationAuthorityTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OrganisationAuthorityTimeSlicePropertyType", "OrganisationAuthorityTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.OrganisationAuthorityType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
				if src.OrganisationAuthority == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "OrganisationAuthority", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.OxygenPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Oxygen", inst.Oxygen, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.OxygenType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportSuppliesServiceTimeSliceTypes {
				for _, target := range src.OxygenSupply {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportSuppliesServiceTimeSliceType", "OxygenSupply", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.OxygenType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeOxygenType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.OxygenPropertyTypes {
				if src.Oxygen == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OxygenPropertyType", "Oxygen", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.PassengerServiceTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("PassengerServiceTimeSlice", inst.PassengerServiceTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PassengerServiceTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PassengerServiceTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PassengerServiceType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.PassengerServiceTimeSliceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
		StageSetAssociationFieldToForm("FlightOperations", inst.FlightOperations, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlightDestinationType](), probe.formStage)
		StageSetAssociationFieldToForm("Rank", inst.Rank, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFacilityRankingType](), probe.formStage)
		StageSetAssociationFieldToForm("CompliantICAO", inst.CompliantICAO, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Location", inst.Location, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("ServiceProvider", inst.ServiceProvider, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.UnitPropertyType](), probe.formStage)

		{
			// Slice of pointers: Call_sign
			div := (&form.FormDiv{Name: "Call_sign"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Call_sign {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Call_sign",
				Label: "Call_sign",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: RadioCommunication
			div := (&form.FormDiv{Name: "RadioCommunication"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RadioCommunication {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RadioCommunication",
				Label: "RadioCommunication",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: GroundCommunication
			div := (&form.FormDiv{Name: "GroundCommunication"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GroundCommunication {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GroundCommunication",
				Label: "GroundCommunication",
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
			// Slice of pointers: AirportHeliport
			div := (&form.FormDiv{Name: "AirportHeliport"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AirportHeliport {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AirportHeliport",
				Label: "AirportHeliport",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodePassengerServiceType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.PassengerServiceTimeSlicePropertyTypes {
				if src.PassengerServiceTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PassengerServiceTimeSlicePropertyType", "PassengerServiceTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.PassengerServiceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
				if src.PassengerService == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "PassengerService", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.PilotControlledLightingTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("PilotControlledLightingTimeSlice", inst.PilotControlledLightingTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PilotControlledLightingTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PilotControlledLightingTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PilotControlledLightingType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.PilotControlledLightingTimeSliceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodePilotControlledLightingType](), probe.formStage)
		StageSetAssociationFieldToForm("Duration", inst.Duration, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDurationType](), probe.formStage)
		StageSetAssociationFieldToForm("IntensitySteps", inst.IntensitySteps, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.NoNumberType](), probe.formStage)
		StageSetAssociationFieldToForm("StandByIntensity", inst.StandByIntensity, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeIntensityStandByType](), probe.formStage)
		StageSetAssociationFieldToForm("RadioFrequency", inst.RadioFrequency, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValFrequencyType](), probe.formStage)
		StageSetAssociationFieldToForm("ActivationInstruction", inst.ActivationInstruction, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextInstructionType](), probe.formStage)

		{
			// Slice of pointers: ControlledLightIntensity
			div := (&form.FormDiv{Name: "ControlledLightIntensity"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ControlledLightIntensity {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ControlledLightIntensity",
				Label: "ControlledLightIntensity",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ActivatedGroundLighting
			div := (&form.FormDiv{Name: "ActivatedGroundLighting"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ActivatedGroundLighting {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ActivatedGroundLighting",
				Label: "ActivatedGroundLighting",
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.PilotControlledLightingTimeSlicePropertyTypes {
				if src.PilotControlledLightingTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PilotControlledLightingTimeSlicePropertyType", "PilotControlledLightingTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.PilotControlledLightingType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
				if src.PilotControlledLighting == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "PilotControlledLighting", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.PointPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Point", inst.Point, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.PointType](), probe.formStage)
		StageSetAssociationFieldToForm("ElevatedPoint", inst.ElevatedPoint, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointType](), probe.formStage)

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
	case *aixmfeatures.ProcedurePropertyType:
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTimeSliceTypes {
				for _, target := range src.ClientProcedure {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceTimeSliceType", "ClientProcedure", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTimeSliceTypes {
				for _, target := range src.ClientProcedure {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceTimeSliceType", "ClientProcedure", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RadioCommunicationChannelPropertyType:
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTimeSliceTypes {
				for _, target := range src.RadioCommunication {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceTimeSliceType", "RadioCommunication", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficManagementServiceTimeSliceTypes {
				for _, target := range src.RadioCommunication {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficManagementServiceTimeSliceType", "RadioCommunication", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftGroundServiceTimeSliceTypes {
				for _, target := range src.RadioCommunication {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftGroundServiceTimeSliceType", "RadioCommunication", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportClearanceServiceTimeSliceTypes {
				for _, target := range src.RadioCommunication {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportClearanceServiceTimeSliceType", "RadioCommunication", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportSuppliesServiceTimeSliceTypes {
				for _, target := range src.RadioCommunication {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportSuppliesServiceTimeSliceType", "RadioCommunication", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FireFightingServiceTimeSliceTypes {
				for _, target := range src.RadioCommunication {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FireFightingServiceTimeSliceType", "RadioCommunication", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GroundTrafficControlServiceTimeSliceTypes {
				for _, target := range src.RadioCommunication {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GroundTrafficControlServiceTimeSliceType", "RadioCommunication", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTimeSliceTypes {
				for _, target := range src.RadioCommunication {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceTimeSliceType", "RadioCommunication", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PassengerServiceTimeSliceTypes {
				for _, target := range src.RadioCommunication {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PassengerServiceTimeSliceType", "RadioCommunication", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SearchRescueServiceTimeSliceTypes {
				for _, target := range src.RadioCommunication {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SearchRescueServiceTimeSliceType", "RadioCommunication", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RadioCommunicationChannelTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("RadioCommunicationChannelTimeSlice", inst.RadioCommunicationChannelTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RadioCommunicationChannelTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationChannelTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationChannelType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RadioCommunicationChannelTimeSliceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
		StageSetAssociationFieldToForm("Mode", inst.Mode, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCommunicationModeType](), probe.formStage)
		StageSetAssociationFieldToForm("Rank", inst.Rank, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFacilityRankingType](), probe.formStage)
		StageSetAssociationFieldToForm("FrequencyTransmission", inst.FrequencyTransmission, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValFrequencyType](), probe.formStage)
		StageSetAssociationFieldToForm("FrequencyReception", inst.FrequencyReception, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValFrequencyType](), probe.formStage)
		StageSetAssociationFieldToForm("Channel", inst.Channel, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCommunicationChannelType](), probe.formStage)
		StageSetAssociationFieldToForm("Logon", inst.Logon, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.TextDesignatorType](), probe.formStage)
		StageSetAssociationFieldToForm("EmissionType", inst.EmissionType, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRadioEmissionType](), probe.formStage)
		StageSetAssociationFieldToForm("SelectiveCall", inst.SelectiveCall, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("FlightChecked", inst.FlightChecked, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("TrafficDirection", inst.TrafficDirection, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeCommunicationDirectionType](), probe.formStage)

		{
			// Slice of pointers: Location
			div := (&form.FormDiv{Name: "Location"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Location {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Location",
				Label: "Location",
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationChannelTimeSlicePropertyTypes {
				if src.RadioCommunicationChannelTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationChannelTimeSlicePropertyType", "RadioCommunicationChannelTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RadioCommunicationChannelType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
				if src.RadioCommunicationChannel == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "RadioCommunicationChannel", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RadioCommunicationOperationalStatusPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("RadioCommunicationOperationalStatus", inst.RadioCommunicationOperationalStatus, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RadioCommunicationOperationalStatusType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationChannelTimeSliceTypes {
				for _, target := range src.Availability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationChannelTimeSliceType", "Availability", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RadioCommunicationOperationalStatusType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
		StageSetAssociationFieldToForm("OperationalStatus", inst.OperationalStatus, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeStatusServiceType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationOperationalStatusPropertyTypes {
				if src.RadioCommunicationOperationalStatus == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationOperationalStatusPropertyType", "RadioCommunicationOperationalStatus", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.RoutePortionPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("RoutePortion", inst.RoutePortion, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.RoutePortionType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTimeSliceTypes {
				for _, target := range src.ClientRoute {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceTimeSliceType", "ClientRoute", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficManagementServiceTimeSliceTypes {
				for _, target := range src.ClientRoute {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficManagementServiceTimeSliceType", "ClientRoute", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTimeSliceTypes {
				for _, target := range src.ClientRoute {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceTimeSliceType", "ClientRoute", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SearchRescueServiceTimeSliceTypes {
				for _, target := range src.ClientRoute {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SearchRescueServiceTimeSliceType", "ClientRoute", refNames, formGroup, probe.formStage)
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
	case *aixmfeatures.SearchRescueServiceTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("SearchRescueServiceTimeSlice", inst.SearchRescueServiceTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SearchRescueServiceTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SearchRescueServiceTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SearchRescueServiceType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SearchRescueServiceTimeSliceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
		StageSetAssociationFieldToForm("FlightOperations", inst.FlightOperations, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlightDestinationType](), probe.formStage)
		StageSetAssociationFieldToForm("Rank", inst.Rank, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFacilityRankingType](), probe.formStage)
		StageSetAssociationFieldToForm("CompliantICAO", inst.CompliantICAO, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Location", inst.Location, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("ServiceProvider", inst.ServiceProvider, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.UnitPropertyType](), probe.formStage)

		{
			// Slice of pointers: Call_sign
			div := (&form.FormDiv{Name: "Call_sign"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Call_sign {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Call_sign",
				Label: "Call_sign",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: RadioCommunication
			div := (&form.FormDiv{Name: "RadioCommunication"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RadioCommunication {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RadioCommunication",
				Label: "RadioCommunication",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: GroundCommunication
			div := (&form.FormDiv{Name: "GroundCommunication"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GroundCommunication {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GroundCommunication",
				Label: "GroundCommunication",
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
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeServiceSARType](), probe.formStage)

		{
			// Slice of pointers: ClientAirspace
			div := (&form.FormDiv{Name: "ClientAirspace"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ClientAirspace {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ClientAirspace",
				Label: "ClientAirspace",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ClientRoute
			div := (&form.FormDiv{Name: "ClientRoute"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ClientRoute {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ClientRoute",
				Label: "ClientRoute",
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.SearchRescueServiceTimeSlicePropertyTypes {
				if src.SearchRescueServiceTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SearchRescueServiceTimeSlicePropertyType", "SearchRescueServiceTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SearchRescueServiceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
				if src.SearchRescueService == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "SearchRescueService", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ServiceOperationalStatusPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ServiceOperationalStatus", inst.ServiceOperationalStatus, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ServiceOperationalStatusType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTimeSliceTypes {
				for _, target := range src.Availability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceTimeSliceType", "Availability", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficManagementServiceTimeSliceTypes {
				for _, target := range src.Availability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficManagementServiceTimeSliceType", "Availability", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftGroundServiceTimeSliceTypes {
				for _, target := range src.Availability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftGroundServiceTimeSliceType", "Availability", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportClearanceServiceTimeSliceTypes {
				for _, target := range src.Availability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportClearanceServiceTimeSliceType", "Availability", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportSuppliesServiceTimeSliceTypes {
				for _, target := range src.Availability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportSuppliesServiceTimeSliceType", "Availability", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FireFightingServiceTimeSliceTypes {
				for _, target := range src.Availability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FireFightingServiceTimeSliceType", "Availability", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GroundTrafficControlServiceTimeSliceTypes {
				for _, target := range src.Availability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GroundTrafficControlServiceTimeSliceType", "Availability", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTimeSliceTypes {
				for _, target := range src.Availability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceTimeSliceType", "Availability", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PassengerServiceTimeSliceTypes {
				for _, target := range src.Availability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PassengerServiceTimeSliceType", "Availability", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SearchRescueServiceTimeSliceTypes {
				for _, target := range src.Availability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SearchRescueServiceTimeSliceType", "Availability", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.ServiceOperationalStatusType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
		StageSetAssociationFieldToForm("OperationalStatus", inst.OperationalStatus, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeStatusServiceType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.ServiceOperationalStatusPropertyTypes {
				if src.ServiceOperationalStatus == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ServiceOperationalStatusPropertyType", "ServiceOperationalStatus", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SpecialDateTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("SpecialDateTimeSlice", inst.SpecialDateTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SpecialDateTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialDateTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialDateType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SpecialDateTimeSliceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeSpecialDateType](), probe.formStage)
		StageSetAssociationFieldToForm("DateDay", inst.DateDay, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateMonthDayType](), probe.formStage)
		StageSetAssociationFieldToForm("DateYear", inst.DateYear, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.DateYearType](), probe.formStage)

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
		StageSetAssociationFieldToForm("Authority", inst.Authority, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.OrganisationAuthorityPropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialDateTimeSlicePropertyTypes {
				if src.SpecialDateTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialDateTimeSlicePropertyType", "SpecialDateTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.SpecialDateType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
				if src.SpecialDate == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "SpecialDate", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelSectorTimeSliceTypes {
				if src.ApplicableLevelColumn == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelSectorTimeSliceType", "ApplicableLevelColumn", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.StandardLevelColumnTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("StandardLevelColumnTimeSlice", inst.StandardLevelColumnTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.StandardLevelColumnTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelColumnTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelColumnType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.StandardLevelColumnTimeSliceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
		StageSetAssociationFieldToForm("Series", inst.Series, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeLevelSeriesType](), probe.formStage)
		StageSetBasicFieldtoForm("UnitOfMeasurement", inst.UnitOfMeasurement, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Separation", inst.Separation, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeRVSMType](), probe.formStage)

		{
			// Slice of pointers: Level
			div := (&form.FormDiv{Name: "Level"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Level {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Level",
				Label: "Level",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("LevelTable", inst.LevelTable, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.StandardLevelTablePropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelColumnTimeSlicePropertyTypes {
				if src.StandardLevelColumnTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelColumnTimeSlicePropertyType", "StandardLevelColumnTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.StandardLevelColumnType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
				if src.StandardLevelColumn == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "StandardLevelColumn", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.StandardLevelPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("StandardLevel", inst.StandardLevel, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.StandardLevelType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelColumnTimeSliceTypes {
				for _, target := range src.Level {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelColumnTimeSliceType", "Level", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.StandardLevelSectorTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("StandardLevelSectorTimeSlice", inst.StandardLevelSectorTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.StandardLevelSectorTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelSectorTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelSectorType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.StandardLevelSectorTimeSliceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
		StageSetAssociationFieldToForm("FlightRule", inst.FlightRule, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeFlightRuleType](), probe.formStage)
		StageSetAssociationFieldToForm("FromTrack", inst.FromTrack, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), probe.formStage)
		StageSetAssociationFieldToForm("ToTrack", inst.ToTrack, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValBearingType](), probe.formStage)
		StageSetAssociationFieldToForm("AngleType", inst.AngleType, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeNorthReferenceType](), probe.formStage)

		{
			// Slice of pointers: ApplicableAirspace
			div := (&form.FormDiv{Name: "ApplicableAirspace"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ApplicableAirspace {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ApplicableAirspace",
				Label: "ApplicableAirspace",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("ApplicableLevelColumn", inst.ApplicableLevelColumn, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.StandardLevelColumnPropertyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelSectorTimeSlicePropertyTypes {
				if src.StandardLevelSectorTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelSectorTimeSlicePropertyType", "StandardLevelSectorTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.StandardLevelSectorType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
				if src.StandardLevelSector == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "StandardLevelSector", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.StandardLevelTablePropertyType:
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelColumnTimeSliceTypes {
				if src.LevelTable == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelColumnTimeSliceType", "LevelTable", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.StandardLevelTableTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("StandardLevelTableTimeSlice", inst.StandardLevelTableTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.StandardLevelTableTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelTableTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelTableType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.StandardLevelTableTimeSliceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
		StageSetAssociationFieldToForm("StandardICAO", inst.StandardICAO, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelTableTimeSlicePropertyTypes {
				if src.StandardLevelTableTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelTableTimeSlicePropertyType", "StandardLevelTableTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.StandardLevelTableType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
				if src.StandardLevelTable == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "StandardLevelTable", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.StandardLevelType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("VerticalDistance", inst.VerticalDistance, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.ValDistanceVerticalType](), probe.formStage)

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
			for src := range probe.stageOfInterest.BasicMessageMemberAIXMPropertyTypes {
				if src.StandardLevel == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "StandardLevel", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelPropertyTypes {
				if src.StandardLevel == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelPropertyType", "StandardLevel", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationOperationalStatusTypes {
				for _, target := range src.TimeInterval {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationOperationalStatusType", "TimeInterval", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ServiceOperationalStatusTypes {
				for _, target := range src.TimeInterval {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ServiceOperationalStatusType", "TimeInterval", refNames, formGroup, probe.formStage)
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

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitAvailabilityTypes {
				for _, target := range src.TimeInterval {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitAvailabilityType", "TimeInterval", refNames, formGroup, probe.formStage)
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
	case *aixmfeatures.UnitAvailabilityPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("UnitAvailability", inst.UnitAvailability, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.UnitAvailabilityType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitTimeSliceTypes {
				for _, target := range src.Availability {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitTimeSliceType", "Availability", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.UnitAvailabilityType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
		StageSetAssociationFieldToForm("OperationalStatus", inst.OperationalStatus, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeStatusOperationsType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitAvailabilityPropertyTypes {
				if src.UnitAvailability == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitAvailabilityPropertyType", "UnitAvailability", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.UnitDependencyPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("UnitDependency", inst.UnitDependency, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.UnitDependencyType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitTimeSliceTypes {
				for _, target := range src.RelatedUnit {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitTimeSliceType", "RelatedUnit", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.UnitDependencyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NameXSD",
				Label: "NameXSD",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeUnitDependencyType](), probe.formStage)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitDependencyPropertyTypes {
				if src.UnitDependency == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitDependencyPropertyType", "UnitDependency", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTimeSliceTypes {
				if src.ServiceProvider == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceTimeSliceType", "ServiceProvider", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficManagementServiceTimeSliceTypes {
				if src.ServiceProvider == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficManagementServiceTimeSliceType", "ServiceProvider", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftGroundServiceTimeSliceTypes {
				if src.ServiceProvider == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftGroundServiceTimeSliceType", "ServiceProvider", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportClearanceServiceTimeSliceTypes {
				if src.ServiceProvider == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportClearanceServiceTimeSliceType", "ServiceProvider", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportSuppliesServiceTimeSliceTypes {
				if src.ServiceProvider == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportSuppliesServiceTimeSliceType", "ServiceProvider", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FireFightingServiceTimeSliceTypes {
				if src.ServiceProvider == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FireFightingServiceTimeSliceType", "ServiceProvider", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GroundTrafficControlServiceTimeSliceTypes {
				if src.ServiceProvider == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GroundTrafficControlServiceTimeSliceType", "ServiceProvider", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTimeSliceTypes {
				if src.ServiceProvider == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceTimeSliceType", "ServiceProvider", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PassengerServiceTimeSliceTypes {
				if src.ServiceProvider == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PassengerServiceTimeSliceType", "ServiceProvider", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SearchRescueServiceTimeSliceTypes {
				if src.ServiceProvider == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SearchRescueServiceTimeSliceType", "ServiceProvider", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitDependencyTypes {
				if src.TheUnit == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitDependencyType", "TheUnit", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.UnitTimeSlicePropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("UnitTimeSlice", inst.UnitTimeSlice, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.UnitTimeSliceType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitTypes {
				for _, target := range src.TimeSlice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitType", "TimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.UnitTimeSliceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
		StageSetAssociationFieldToForm("Type_XSD", inst.Type_XSD, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeUnitType](), probe.formStage)
		StageSetAssociationFieldToForm("CompliantICAO", inst.CompliantICAO, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeYesNoType](), probe.formStage)
		StageSetAssociationFieldToForm("Designator", inst.Designator, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeOrganisationDesignatorType](), probe.formStage)
		StageSetAssociationFieldToForm("Military", inst.Military, formGroup, probe.stageOfInterest.AixmdatatypesStage.GetInstancesSet[*aixmdatatypes.CodeMilitaryOperationsType](), probe.formStage)
		StageSetAssociationFieldToForm("Position", inst.Position, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("AirportLocation", inst.AirportLocation, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportPropertyType](), probe.formStage)
		StageSetAssociationFieldToForm("OwnerOrganisation", inst.OwnerOrganisation, formGroup, probe.stageOfInterest.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.OrganisationAuthorityPropertyType](), probe.formStage)

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
			// Slice of pointers: RelatedUnit
			div := (&form.FormDiv{Name: "RelatedUnit"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RelatedUnit {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RelatedUnit",
				Label: "RelatedUnit",
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitTimeSlicePropertyTypes {
				if src.UnitTimeSlice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitTimeSlicePropertyType", "UnitTimeSlice", refNames, formGroup, probe.formStage)
		}
	case *aixmfeatures.UnitType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identifier", inst.Identifier, formGroup, probe.stageOfInterest.GmlStage.GetInstancesSet[*gml.CodeWithAuthorityType](), probe.formStage)

		{
			// Slice of pointers: NameXSD
			div := (&form.FormDiv{Name: "NameXSD"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NameXSD {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
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
				if src.Unit == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BasicMessageMemberAIXMPropertyType", "Unit", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTimeSliceTypes {
				for _, target := range src.NavaidBroadcast {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceTimeSliceType", "NavaidBroadcast", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeAircraftGroundServiceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftGroundServiceTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftGroundServiceTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeAviationStandardsType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FireFightingServiceTimeSliceTypes {
				if src.Standard == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FireFightingServiceTimeSliceType", "Standard", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeCommunicationChannelType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTimeSliceTypes {
				if src.DataLinkChannel == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceTimeSliceType", "DataLinkChannel", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GroundTrafficControlServiceTimeSliceTypes {
				if src.DataLinkChannel == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GroundTrafficControlServiceTimeSliceType", "DataLinkChannel", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationChannelTimeSliceTypes {
				if src.Channel == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationChannelTimeSliceType", "Channel", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeCommunicationDirectionType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationChannelTimeSliceTypes {
				if src.TrafficDirection == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationChannelTimeSliceType", "TrafficDirection", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeCommunicationModeType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationChannelTimeSliceTypes {
				if src.Mode == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationChannelTimeSliceType", "Mode", refNames, formGroup, probe.formStage)
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
	case *aixmdatatypes.CodeFacilityRankingType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTimeSliceTypes {
				if src.Rank == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceTimeSliceType", "Rank", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficManagementServiceTimeSliceTypes {
				if src.Rank == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficManagementServiceTimeSliceType", "Rank", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftGroundServiceTimeSliceTypes {
				if src.Rank == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftGroundServiceTimeSliceType", "Rank", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportClearanceServiceTimeSliceTypes {
				if src.Rank == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportClearanceServiceTimeSliceType", "Rank", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportSuppliesServiceTimeSliceTypes {
				if src.Rank == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportSuppliesServiceTimeSliceType", "Rank", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FireFightingServiceTimeSliceTypes {
				if src.Rank == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FireFightingServiceTimeSliceType", "Rank", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GroundTrafficControlServiceTimeSliceTypes {
				if src.Rank == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GroundTrafficControlServiceTimeSliceType", "Rank", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTimeSliceTypes {
				if src.Rank == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceTimeSliceType", "Rank", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PassengerServiceTimeSliceTypes {
				if src.Rank == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PassengerServiceTimeSliceType", "Rank", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationChannelTimeSliceTypes {
				if src.Rank == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationChannelTimeSliceType", "Rank", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SearchRescueServiceTimeSliceTypes {
				if src.Rank == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SearchRescueServiceTimeSliceType", "Rank", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeFireFightingType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FireFightingServiceTimeSliceTypes {
				if src.Category == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FireFightingServiceTimeSliceType", "Category", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeFlightDestinationType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTimeSliceTypes {
				if src.FlightOperations == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceTimeSliceType", "FlightOperations", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficManagementServiceTimeSliceTypes {
				if src.FlightOperations == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficManagementServiceTimeSliceType", "FlightOperations", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftGroundServiceTimeSliceTypes {
				if src.FlightOperations == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftGroundServiceTimeSliceType", "FlightOperations", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportClearanceServiceTimeSliceTypes {
				if src.FlightOperations == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportClearanceServiceTimeSliceType", "FlightOperations", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportSuppliesServiceTimeSliceTypes {
				if src.FlightOperations == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportSuppliesServiceTimeSliceType", "FlightOperations", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FireFightingServiceTimeSliceTypes {
				if src.FlightOperations == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FireFightingServiceTimeSliceType", "FlightOperations", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GroundTrafficControlServiceTimeSliceTypes {
				if src.FlightOperations == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GroundTrafficControlServiceTimeSliceType", "FlightOperations", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTimeSliceTypes {
				if src.FlightOperations == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceTimeSliceType", "FlightOperations", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PassengerServiceTimeSliceTypes {
				if src.FlightOperations == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PassengerServiceTimeSliceType", "FlightOperations", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SearchRescueServiceTimeSliceTypes {
				if src.FlightOperations == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SearchRescueServiceTimeSliceType", "FlightOperations", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeFlightRuleType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelSectorTimeSliceTypes {
				if src.FlightRule == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelSectorTimeSliceType", "FlightRule", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeFuelType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FuelTypes {
				if src.Category == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FuelType", "Category", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeIntensityStandByType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PilotControlledLightingTimeSliceTypes {
				if src.StandByIntensity == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PilotControlledLightingTimeSliceType", "StandByIntensity", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeLanguageType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CallsignDetailTypes {
				if src.Language == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CallsignDetailType", "Language", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeLevelSeriesType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelColumnTimeSliceTypes {
				if src.Series == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelColumnTimeSliceType", "Series", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeLevelTableDesignatorType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)
	case *aixmdatatypes.CodeLightIntensityType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightActivationTypes {
				if src.IntensityLevel == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LightActivationType", "IntensityLevel", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeMilitaryOperationsType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OrganisationAuthorityTimeSliceTypes {
				if src.Military == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OrganisationAuthorityTimeSliceType", "Military", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitTimeSliceTypes {
				if src.Military == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitTimeSliceType", "Military", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeNitrogenType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.NitrogenTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NitrogenType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeNorthReferenceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelSectorTimeSliceTypes {
				if src.AngleType == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelSectorTimeSliceType", "AngleType", refNames, formGroup, probe.formStage)
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
	case *aixmdatatypes.CodeOilType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OilTypes {
				if src.Category == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OilType", "Category", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeOrganisationDesignatorType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OrganisationAuthorityTimeSliceTypes {
				if src.Designator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OrganisationAuthorityTimeSliceType", "Designator", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitTimeSliceTypes {
				if src.Designator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitTimeSliceType", "Designator", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeOrganisationHierarchyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OrganisationAuthorityAssociationTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OrganisationAuthorityAssociationType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeOrganisationType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OrganisationAuthorityTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OrganisationAuthorityTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeOxygenType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OxygenTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OxygenType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodePassengerServiceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PassengerServiceTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PassengerServiceTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodePilotControlledLightingType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PilotControlledLightingTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PilotControlledLightingTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeRVSMType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelColumnTimeSliceTypes {
				if src.Separation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelColumnTimeSliceType", "Separation", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeRadioEmissionType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationChannelTimeSliceTypes {
				if src.EmissionType == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationChannelTimeSliceType", "EmissionType", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeServiceATCType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeServiceATFMType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficManagementServiceTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficManagementServiceTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeServiceGroundControlType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GroundTrafficControlServiceTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GroundTrafficControlServiceTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeServiceInformationType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeServiceSARType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SearchRescueServiceTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SearchRescueServiceTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeSpecialDateType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialDateTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialDateTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeStatusOperationsType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitAvailabilityTypes {
				if src.OperationalStatus == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitAvailabilityType", "OperationalStatus", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeStatusServiceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationOperationalStatusTypes {
				if src.OperationalStatus == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationOperationalStatusType", "OperationalStatus", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ServiceOperationalStatusTypes {
				if src.OperationalStatus == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ServiceOperationalStatusType", "OperationalStatus", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeSystemActivationType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightActivationTypes {
				if src.Activation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LightActivationType", "Activation", refNames, formGroup, probe.formStage)
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
	case *aixmdatatypes.CodeUnitDependencyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitDependencyTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitDependencyType", "Type_XSD", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.CodeUnitType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitTimeSliceTypes {
				if src.Type_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitTimeSliceType", "Type_XSD", refNames, formGroup, probe.formStage)
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
	case *aixmdatatypes.CodeYesNoType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTimeSliceTypes {
				if src.CompliantICAO == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceTimeSliceType", "CompliantICAO", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTimeSliceTypes {
				if src.RadarAssisted == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceTimeSliceType", "RadarAssisted", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTimeSliceTypes {
				if src.DataLinkEnabled == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceTimeSliceType", "DataLinkEnabled", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficManagementServiceTimeSliceTypes {
				if src.CompliantICAO == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficManagementServiceTimeSliceType", "CompliantICAO", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftGroundServiceTimeSliceTypes {
				if src.CompliantICAO == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftGroundServiceTimeSliceType", "CompliantICAO", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportClearanceServiceTimeSliceTypes {
				if src.CompliantICAO == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportClearanceServiceTimeSliceType", "CompliantICAO", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportSuppliesServiceTimeSliceTypes {
				if src.CompliantICAO == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportSuppliesServiceTimeSliceType", "CompliantICAO", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FireFightingServiceTimeSliceTypes {
				if src.CompliantICAO == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FireFightingServiceTimeSliceType", "CompliantICAO", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GroundTrafficControlServiceTimeSliceTypes {
				if src.CompliantICAO == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GroundTrafficControlServiceTimeSliceType", "CompliantICAO", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GroundTrafficControlServiceTimeSliceTypes {
				if src.RadarAssisted == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GroundTrafficControlServiceTimeSliceType", "RadarAssisted", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GroundTrafficControlServiceTimeSliceTypes {
				if src.DataLinkEnabled == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GroundTrafficControlServiceTimeSliceType", "DataLinkEnabled", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTimeSliceTypes {
				if src.CompliantICAO == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceTimeSliceType", "CompliantICAO", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTimeSliceTypes {
				if src.Voice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceTimeSliceType", "Voice", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTimeSliceTypes {
				if src.DataLink == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceTimeSliceType", "DataLink", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTimeSliceTypes {
				if src.Recorded == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceTimeSliceType", "Recorded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PassengerServiceTimeSliceTypes {
				if src.CompliantICAO == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PassengerServiceTimeSliceType", "CompliantICAO", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationChannelTimeSliceTypes {
				if src.SelectiveCall == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationChannelTimeSliceType", "SelectiveCall", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationChannelTimeSliceTypes {
				if src.FlightChecked == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationChannelTimeSliceType", "FlightChecked", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SearchRescueServiceTimeSliceTypes {
				if src.CompliantICAO == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SearchRescueServiceTimeSliceType", "CompliantICAO", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelTableTimeSliceTypes {
				if src.StandardICAO == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelTableTimeSliceType", "StandardICAO", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitTimeSliceTypes {
				if src.CompliantICAO == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitTimeSliceType", "CompliantICAO", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.DateMonthDayType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialDateTimeSliceTypes {
				if src.DateDay == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialDateTimeSliceType", "DateDay", refNames, formGroup, probe.formStage)
		}

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialDateTimeSliceTypes {
				if src.DateYear == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialDateTimeSliceType", "DateYear", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.NoNumberType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightActivationTypes {
				if src.Clicks == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LightActivationType", "Clicks", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PilotControlledLightingTimeSliceTypes {
				if src.IntensitySteps == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PilotControlledLightingTimeSliceType", "IntensitySteps", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationChannelTimeSliceTypes {
				if src.Logon == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationChannelTimeSliceType", "Logon", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.TextInstructionType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportClearanceServiceTimeSliceTypes {
				if src.SnowPlan == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportClearanceServiceTimeSliceType", "SnowPlan", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PilotControlledLightingTimeSliceTypes {
				if src.ActivationInstruction == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PilotControlledLightingTimeSliceType", "ActivationInstruction", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.TextNameType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CallsignDetailTypes {
				if src.CallSign == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CallsignDetailType", "CallSign", refNames, formGroup, probe.formStage)
		}

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
	case *aixmdatatypes.ValBearingType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelSectorTimeSliceTypes {
				if src.FromTrack == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelSectorTimeSliceType", "FromTrack", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelSectorTimeSliceTypes {
				if src.ToTrack == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelSectorTimeSliceType", "ToTrack", refNames, formGroup, probe.formStage)
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
	case *aixmdatatypes.ValDistanceVerticalType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Uom", inst.Uom, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelTypes {
				if src.VerticalDistance == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelType", "VerticalDistance", refNames, formGroup, probe.formStage)
		}
	case *aixmdatatypes.ValDurationType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Uom", inst.Uom, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NilReason", inst.NilReason, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PilotControlledLightingTimeSliceTypes {
				if src.Duration == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PilotControlledLightingTimeSliceType", "Duration", refNames, formGroup, probe.formStage)
		}

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
			for src := range probe.stageOfInterest.AixmfeaturesStage.PilotControlledLightingTimeSliceTypes {
				if src.RadioFrequency == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PilotControlledLightingTimeSliceType", "RadioFrequency", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationChannelTimeSliceTypes {
				if src.FrequencyTransmission == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationChannelTimeSliceType", "FrequencyTransmission", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationChannelTimeSliceTypes {
				if src.FrequencyReception == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationChannelTimeSliceType", "FrequencyReception", refNames, formGroup, probe.formStage)
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
				for _, target := range src.AbstractLightActivationExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension", "AbstractLightActivationExtension", refNames, formGroup, probe.formStage)
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
				for _, target := range src.AbstractAircraftGroundServiceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_10", "AbstractAircraftGroundServiceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_10s {
				for _, target := range src.AbstractAirportGroundServiceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_10", "AbstractAirportGroundServiceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_10s {
				for _, target := range src.AbstractServiceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_10", "AbstractServiceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_11s {
				for _, target := range src.AbstractAirportClearanceServiceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_11", "AbstractAirportClearanceServiceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_11s {
				for _, target := range src.AbstractAirportGroundServiceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_11", "AbstractAirportGroundServiceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_11s {
				for _, target := range src.AbstractServiceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_11", "AbstractServiceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_12s {
				for _, target := range src.AbstractAirportSuppliesServiceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_12", "AbstractAirportSuppliesServiceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_12s {
				for _, target := range src.AbstractAirportGroundServiceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_12", "AbstractAirportGroundServiceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_12s {
				for _, target := range src.AbstractServiceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_12", "AbstractServiceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_13s {
				for _, target := range src.AbstractAirTrafficControlServiceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_13", "AbstractAirTrafficControlServiceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_13s {
				for _, target := range src.AbstractTrafficSeparationServiceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_13", "AbstractTrafficSeparationServiceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_13s {
				for _, target := range src.AbstractServiceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_13", "AbstractServiceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_14s {
				for _, target := range src.AbstractAirTrafficManagementServiceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_14", "AbstractAirTrafficManagementServiceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_14s {
				for _, target := range src.AbstractServiceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_14", "AbstractServiceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_15s {
				for _, target := range src.AbstractCallsignDetailExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_15", "AbstractCallsignDetailExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_16s {
				for _, target := range src.AbstractFireFightingServiceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_16", "AbstractFireFightingServiceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_16s {
				for _, target := range src.AbstractAirportGroundServiceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_16", "AbstractAirportGroundServiceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_16s {
				for _, target := range src.AbstractServiceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_16", "AbstractServiceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_17s {
				for _, target := range src.AbstractFuelExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_17", "AbstractFuelExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_18s {
				for _, target := range src.AbstractGroundTrafficControlServiceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_18", "AbstractGroundTrafficControlServiceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_18s {
				for _, target := range src.AbstractTrafficSeparationServiceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_18", "AbstractTrafficSeparationServiceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_18s {
				for _, target := range src.AbstractServiceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_18", "AbstractServiceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_19s {
				for _, target := range src.AbstractInformationServiceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_19", "AbstractInformationServiceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_19s {
				for _, target := range src.AbstractServiceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_19", "AbstractServiceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_2s {
				for _, target := range src.AbstractLinguisticNoteExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_2", "AbstractLinguisticNoteExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_20s {
				for _, target := range src.AbstractNitrogenExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_20", "AbstractNitrogenExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_21s {
				for _, target := range src.AbstractOilExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_21", "AbstractOilExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_22s {
				for _, target := range src.AbstractOxygenExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_22", "AbstractOxygenExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_23s {
				for _, target := range src.AbstractPassengerServiceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_23", "AbstractPassengerServiceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_23s {
				for _, target := range src.AbstractAirportGroundServiceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_23", "AbstractAirportGroundServiceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_23s {
				for _, target := range src.AbstractServiceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_23", "AbstractServiceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_24s {
				for _, target := range src.AbstractPilotControlledLightingExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_24", "AbstractPilotControlledLightingExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_25s {
				for _, target := range src.AbstractRadioCommunicationChannelExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_25", "AbstractRadioCommunicationChannelExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_26s {
				for _, target := range src.AbstractPropertiesWithScheduleExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_26", "AbstractPropertiesWithScheduleExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_26s {
				for _, target := range src.AbstractRadioCommunicationOperationalStatusExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_26", "AbstractRadioCommunicationOperationalStatusExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_27s {
				for _, target := range src.AbstractSearchRescueServiceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_27", "AbstractSearchRescueServiceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_27s {
				for _, target := range src.AbstractServiceExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_27", "AbstractServiceExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_28s {
				for _, target := range src.AbstractPropertiesWithScheduleExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_28", "AbstractPropertiesWithScheduleExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_28s {
				for _, target := range src.AbstractServiceOperationalStatusExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_28", "AbstractServiceOperationalStatusExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_29s {
				for _, target := range src.AbstractContactInformationExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_29", "AbstractContactInformationExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_3s {
				for _, target := range src.AbstractNoteExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_3", "AbstractNoteExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_30s {
				for _, target := range src.AbstractPropertiesWithScheduleExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_30", "AbstractPropertiesWithScheduleExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_30s {
				for _, target := range src.AbstractOnlineContactExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_30", "AbstractOnlineContactExtension", refNames, formGroup, probe.formStage)
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
				for _, target := range src.AbstractPostalAddressExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_31", "AbstractPostalAddressExtension", refNames, formGroup, probe.formStage)
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
				for _, target := range src.AbstractTelephoneContactExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_32", "AbstractTelephoneContactExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_33s {
				for _, target := range src.AbstractSpecialDateExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_33", "AbstractSpecialDateExtension", refNames, formGroup, probe.formStage)
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
				for _, target := range src.AbstractStandardLevelExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_35", "AbstractStandardLevelExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_36s {
				for _, target := range src.AbstractStandardLevelColumnExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_36", "AbstractStandardLevelColumnExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_37s {
				for _, target := range src.AbstractStandardLevelSectorExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_37", "AbstractStandardLevelSectorExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_38s {
				for _, target := range src.AbstractStandardLevelTableExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_38", "AbstractStandardLevelTableExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_4s {
				for _, target := range src.AbstractOrganisationAuthorityExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_4", "AbstractOrganisationAuthorityExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_5s {
				for _, target := range src.AbstractOrganisationAuthorityAssociationExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_5", "AbstractOrganisationAuthorityAssociationExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_6s {
				for _, target := range src.AbstractUnitExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_6", "AbstractUnitExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_7s {
				for _, target := range src.AbstractPropertiesWithScheduleExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_7", "AbstractPropertiesWithScheduleExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_7s {
				for _, target := range src.AbstractUnitAvailabilityExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_7", "AbstractUnitAvailabilityExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_8s {
				for _, target := range src.AbstractUnitDependencyExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_8", "AbstractUnitDependencyExtension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.A_extension_9s {
				for _, target := range src.AbstractRoutePortionExtension {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.A_extension_9", "AbstractRoutePortionExtension", refNames, formGroup, probe.formStage)
		}
	case *aixmabstractgmlobjecttypes.FeatureMetadataPropertyType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Owns", inst.Owns, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractAirportGroundServiceTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractAirportGroundServiceType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractServiceTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractServiceType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractTrafficSeparationServiceTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractTrafficSeparationServiceType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficManagementServiceTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficManagementServiceType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftGroundServiceTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftGroundServiceType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportClearanceServiceTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportClearanceServiceType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportSuppliesServiceTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportSuppliesServiceType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FireFightingServiceTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FireFightingServiceType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GroundTrafficControlServiceTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GroundTrafficControlServiceType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OrganisationAuthorityTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OrganisationAuthorityType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PassengerServiceTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PassengerServiceType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PilotControlledLightingTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PilotControlledLightingType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationChannelTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationChannelType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SearchRescueServiceTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SearchRescueServiceType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialDateTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialDateType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelColumnTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelColumnType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelSectorTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelSectorType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelTableTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelTableType", "FeatureMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitTypes {
				if src.FeatureMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitType", "FeatureMetadata", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficManagementServiceTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficManagementServiceTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftGroundServiceTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftGroundServiceTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportClearanceServiceTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportClearanceServiceTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportSuppliesServiceTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportSuppliesServiceTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FireFightingServiceTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FireFightingServiceTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GroundTrafficControlServiceTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GroundTrafficControlServiceTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OrganisationAuthorityTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OrganisationAuthorityTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PassengerServiceTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PassengerServiceTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PilotControlledLightingTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PilotControlledLightingTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationChannelTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationChannelTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SearchRescueServiceTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SearchRescueServiceTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialDateTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialDateTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelColumnTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelColumnTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelSectorTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelSectorTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelTableTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelTableTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitTimeSliceTypes {
				if src.TimeSliceMetadata == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitTimeSliceType", "TimeSliceMetadata", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractAirportGroundServiceTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractAirportGroundServiceType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractServiceTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractServiceType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractTrafficSeparationServiceTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractTrafficSeparationServiceType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficManagementServiceTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficManagementServiceType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftGroundServiceTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftGroundServiceType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportClearanceServiceTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportClearanceServiceType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportSuppliesServiceTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportSuppliesServiceType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FireFightingServiceTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FireFightingServiceType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GroundTrafficControlServiceTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GroundTrafficControlServiceType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OrganisationAuthorityTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OrganisationAuthorityType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PassengerServiceTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PassengerServiceType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PilotControlledLightingTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PilotControlledLightingType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationChannelTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationChannelType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SearchRescueServiceTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SearchRescueServiceType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialDateTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialDateType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelColumnTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelColumnType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelSectorTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelSectorType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelTableTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelTableType", "BoundedBy", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitTypes {
				if src.BoundedBy == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitType", "BoundedBy", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractAirportGroundServiceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractAirportGroundServiceType", "NameXSD", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractServiceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractServiceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractTrafficSeparationServiceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractTrafficSeparationServiceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficManagementServiceTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficManagementServiceTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficManagementServiceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficManagementServiceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftGroundServiceTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftGroundServiceTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftGroundServiceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftGroundServiceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportClearanceServiceTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportClearanceServiceTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportClearanceServiceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportClearanceServiceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportSuppliesServiceTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportSuppliesServiceTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportSuppliesServiceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportSuppliesServiceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CallsignDetailTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CallsignDetailType", "NameXSD", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.FireFightingServiceTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FireFightingServiceTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FireFightingServiceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FireFightingServiceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FuelTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FuelType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GroundTrafficControlServiceTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GroundTrafficControlServiceTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GroundTrafficControlServiceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GroundTrafficControlServiceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightActivationTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LightActivationType", "NameXSD", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.NitrogenTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NitrogenType", "NameXSD", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.OilTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OilType", "NameXSD", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.OrganisationAuthorityAssociationTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OrganisationAuthorityAssociationType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OrganisationAuthorityTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OrganisationAuthorityTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OrganisationAuthorityTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OrganisationAuthorityType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OxygenTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OxygenType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PassengerServiceTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PassengerServiceTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PassengerServiceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PassengerServiceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PilotControlledLightingTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PilotControlledLightingTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PilotControlledLightingTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PilotControlledLightingType", "NameXSD", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationChannelTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationChannelTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationChannelTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationChannelType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationOperationalStatusTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationOperationalStatusType", "NameXSD", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.SearchRescueServiceTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SearchRescueServiceTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SearchRescueServiceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SearchRescueServiceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ServiceOperationalStatusTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ServiceOperationalStatusType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialDateTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialDateTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialDateTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialDateType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelColumnTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelColumnTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelColumnTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelColumnType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelSectorTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelSectorTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelSectorTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelSectorType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelTableTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelTableTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelTableTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelTableType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelType", "NameXSD", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitAvailabilityTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitAvailabilityType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitDependencyTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitDependencyType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitTimeSliceTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitTimeSliceType", "NameXSD", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitTypes {
				for _, target := range src.NameXSD {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitType", "NameXSD", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractAirportGroundServiceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractAirportGroundServiceType", "Identifier", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractServiceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractServiceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AbstractTrafficSeparationServiceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AbstractTrafficSeparationServiceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficManagementServiceTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficManagementServiceTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficManagementServiceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficManagementServiceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftGroundServiceTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftGroundServiceTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftGroundServiceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftGroundServiceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportClearanceServiceTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportClearanceServiceTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportClearanceServiceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportClearanceServiceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportSuppliesServiceTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportSuppliesServiceTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportSuppliesServiceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportSuppliesServiceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.CallsignDetailTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.CallsignDetailType", "Identifier", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.FireFightingServiceTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FireFightingServiceTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FireFightingServiceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FireFightingServiceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FuelTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FuelType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GroundTrafficControlServiceTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GroundTrafficControlServiceTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GroundTrafficControlServiceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GroundTrafficControlServiceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.LightActivationTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.LightActivationType", "Identifier", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.NitrogenTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.NitrogenType", "Identifier", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.OilTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OilType", "Identifier", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.OrganisationAuthorityAssociationTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OrganisationAuthorityAssociationType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OrganisationAuthorityTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OrganisationAuthorityTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OrganisationAuthorityTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OrganisationAuthorityType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OxygenTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OxygenType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PassengerServiceTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PassengerServiceTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PassengerServiceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PassengerServiceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PilotControlledLightingTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PilotControlledLightingTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PilotControlledLightingTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PilotControlledLightingType", "Identifier", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationChannelTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationChannelTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationChannelTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationChannelType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationOperationalStatusTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationOperationalStatusType", "Identifier", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.SearchRescueServiceTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SearchRescueServiceTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SearchRescueServiceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SearchRescueServiceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.ServiceOperationalStatusTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.ServiceOperationalStatusType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialDateTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialDateTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialDateTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialDateType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelColumnTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelColumnTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelColumnTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelColumnType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelSectorTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelSectorTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelSectorTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelSectorType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelTableTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelTableTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelTableTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelTableType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelType", "Identifier", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitAvailabilityTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitAvailabilityType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitDependencyTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitDependencyType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitTimeSliceTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitTimeSliceType", "Identifier", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitTypes {
				if src.Identifier == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitType", "Identifier", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficControlServiceTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficControlServiceTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficManagementServiceTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficManagementServiceTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirTrafficManagementServiceTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirTrafficManagementServiceTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftGroundServiceTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftGroundServiceTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AircraftGroundServiceTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AircraftGroundServiceTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportClearanceServiceTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportClearanceServiceTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportClearanceServiceTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportClearanceServiceTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportSuppliesServiceTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportSuppliesServiceTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.AirportSuppliesServiceTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.AirportSuppliesServiceTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FireFightingServiceTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FireFightingServiceTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.FireFightingServiceTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.FireFightingServiceTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GroundTrafficControlServiceTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GroundTrafficControlServiceTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.GroundTrafficControlServiceTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.GroundTrafficControlServiceTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.InformationServiceTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.InformationServiceTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OrganisationAuthorityTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OrganisationAuthorityTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.OrganisationAuthorityTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.OrganisationAuthorityTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PassengerServiceTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PassengerServiceTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PassengerServiceTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PassengerServiceTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PilotControlledLightingTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PilotControlledLightingTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.PilotControlledLightingTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.PilotControlledLightingTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationChannelTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationChannelTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.RadioCommunicationChannelTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.RadioCommunicationChannelTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SearchRescueServiceTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SearchRescueServiceTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SearchRescueServiceTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SearchRescueServiceTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialDateTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialDateTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.SpecialDateTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.SpecialDateTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelColumnTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelColumnTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelColumnTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelColumnTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelSectorTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelSectorTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelSectorTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelSectorTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelTableTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelTableTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.StandardLevelTableTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.StandardLevelTableTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitTimeSliceTypes {
				if src.ValidTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitTimeSliceType", "ValidTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.AixmfeaturesStage.UnitTimeSliceTypes {
				if src.FeatureLifetime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("aixmfeatures.UnitTimeSliceType", "FeatureLifetime", refNames, formGroup, probe.formStage)
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
