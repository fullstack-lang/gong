// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/app/xsd/go/models"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func StageSetFillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *StageSetProbe,
) {
	switch inst := instance.(type) {
	case *models.All:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Annotation", inst.Annotation, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Annotation](), probe.formStage)
		StageSetBasicFieldtoForm("OuterElementName", inst.OuterElementName, probe.formStage, formGroup)

		{
			// Slice of pointers: Sequences
			div := (&form.FormDiv{Name: "Sequences"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Sequences {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Sequences",
				Label: "Sequences",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Alls
			div := (&form.FormDiv{Name: "Alls"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Alls {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Alls",
				Label: "Alls",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Choices
			div := (&form.FormDiv{Name: "Choices"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Choices {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Choices",
				Label: "Choices",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Groups
			div := (&form.FormDiv{Name: "Groups"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Groups {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Groups",
				Label: "Groups",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Elements
			div := (&form.FormDiv{Name: "Elements"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Elements {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Elements",
				Label: "Elements",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("Order", inst.Order, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Depth", inst.Depth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MinOccurs", inst.MinOccurs, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MaxOccurs", inst.MaxOccurs, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Alls {
				for _, target := range src.Alls {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.All", "Alls", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Choices {
				for _, target := range src.Alls {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Choice", "Alls", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ComplexTypes {
				for _, target := range src.Alls {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ComplexType", "Alls", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Extensions {
				for _, target := range src.Alls {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Extension", "Alls", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Groups {
				for _, target := range src.Alls {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Group", "Alls", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Sequences {
				for _, target := range src.Alls {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Sequence", "Alls", refNames, formGroup, probe.formStage)
		}
	case *models.Annotation:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: Documentations
			div := (&form.FormDiv{Name: "Documentations"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Documentations {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Documentations",
				Label: "Documentations",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Alls {
				if src.Annotation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.All", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Attributes {
				if src.Annotation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Attribute", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.AttributeGroups {
				if src.Annotation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AttributeGroup", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Choices {
				if src.Annotation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Choice", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ComplexTypes {
				if src.Annotation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ComplexType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Elements {
				if src.Annotation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Element", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Enumerations {
				if src.Annotation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Enumeration", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Groups {
				if src.Annotation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Group", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Lengths {
				if src.Annotation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Length", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.MaxInclusives {
				if src.Annotation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.MaxInclusive", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.MaxLengths {
				if src.Annotation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.MaxLength", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.MinInclusives {
				if src.Annotation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.MinInclusive", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.MinLengths {
				if src.Annotation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.MinLength", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Patterns {
				if src.Annotation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Pattern", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Restrictions {
				if src.Annotation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Restriction", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Schemas {
				if src.Annotation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Schema", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Sequences {
				if src.Annotation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Sequence", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.SimpleTypes {
				if src.Annotation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SimpleType", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TotalDigits {
				if src.Annotation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TotalDigit", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Unions {
				if src.Annotation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Union", "Annotation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.WhiteSpaces {
				if src.Annotation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.WhiteSpace", "Annotation", refNames, formGroup, probe.formStage)
		}
	case *models.Attribute:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NameXSD", inst.NameXSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Annotation", inst.Annotation, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Annotation](), probe.formStage)
		StageSetBasicFieldtoForm("HasNameConflict", inst.HasNameConflict, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("GoIdentifier", inst.GoIdentifier, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default", inst.Default, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Use", inst.Use, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Form", inst.Form, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Fixed", inst.Fixed, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Ref", inst.Ref, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TargetNamespace", inst.TargetNamespace, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SimpleType", inst.SimpleType, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDXSD", inst.IDXSD, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.AttributeGroups {
				for _, target := range src.Attributes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AttributeGroup", "Attributes", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ComplexTypes {
				for _, target := range src.Attributes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ComplexType", "Attributes", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Extensions {
				for _, target := range src.Attributes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Extension", "Attributes", refNames, formGroup, probe.formStage)
		}
	case *models.AttributeGroup:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NameXSD", inst.NameXSD, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Annotation", inst.Annotation, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Annotation](), probe.formStage)
		StageSetBasicFieldtoForm("HasNameConflict", inst.HasNameConflict, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("GoIdentifier", inst.GoIdentifier, probe.formStage, formGroup)

		{
			// Slice of pointers: AttributeGroups
			div := (&form.FormDiv{Name: "AttributeGroups"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AttributeGroups {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AttributeGroups",
				Label: "AttributeGroups",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("Ref", inst.Ref, probe.formStage, formGroup)

		{
			// Slice of pointers: Attributes
			div := (&form.FormDiv{Name: "Attributes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Attributes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Attributes",
				Label: "Attributes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("Order", inst.Order, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Depth", inst.Depth, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.AttributeGroups {
				for _, target := range src.AttributeGroups {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AttributeGroup", "AttributeGroups", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ComplexTypes {
				for _, target := range src.AttributeGroups {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ComplexType", "AttributeGroups", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Extensions {
				for _, target := range src.AttributeGroups {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Extension", "AttributeGroups", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Schemas {
				for _, target := range src.AttributeGroups {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Schema", "AttributeGroups", refNames, formGroup, probe.formStage)
		}
	case *models.Choice:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Annotation", inst.Annotation, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Annotation](), probe.formStage)
		StageSetBasicFieldtoForm("OuterElementName", inst.OuterElementName, probe.formStage, formGroup)

		{
			// Slice of pointers: Sequences
			div := (&form.FormDiv{Name: "Sequences"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Sequences {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Sequences",
				Label: "Sequences",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Alls
			div := (&form.FormDiv{Name: "Alls"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Alls {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Alls",
				Label: "Alls",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Choices
			div := (&form.FormDiv{Name: "Choices"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Choices {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Choices",
				Label: "Choices",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Groups
			div := (&form.FormDiv{Name: "Groups"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Groups {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Groups",
				Label: "Groups",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Elements
			div := (&form.FormDiv{Name: "Elements"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Elements {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Elements",
				Label: "Elements",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("Order", inst.Order, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Depth", inst.Depth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MinOccurs", inst.MinOccurs, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MaxOccurs", inst.MaxOccurs, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsDuplicatedInXSD", inst.IsDuplicatedInXSD, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Alls {
				for _, target := range src.Choices {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.All", "Choices", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Choices {
				for _, target := range src.Choices {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Choice", "Choices", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ComplexTypes {
				for _, target := range src.Choices {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ComplexType", "Choices", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Extensions {
				for _, target := range src.Choices {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Extension", "Choices", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Groups {
				for _, target := range src.Choices {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Group", "Choices", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Sequences {
				for _, target := range src.Choices {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Sequence", "Choices", refNames, formGroup, probe.formStage)
		}
	case *models.ComplexContent:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ComplexTypes {
				if src.ComplexContent == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ComplexType", "ComplexContent", refNames, formGroup, probe.formStage)
		}
	case *models.ComplexType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasNameConflict", inst.HasNameConflict, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("GoIdentifier", inst.GoIdentifier, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsAnonymous", inst.IsAnonymous, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("OuterElement", inst.OuterElement, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Element](), probe.formStage)
		StageSetAssociationFieldToForm("Annotation", inst.Annotation, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Annotation](), probe.formStage)
		StageSetBasicFieldtoForm("NameXSD", inst.NameXSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("OuterElementName", inst.OuterElementName, probe.formStage, formGroup)

		{
			// Slice of pointers: Sequences
			div := (&form.FormDiv{Name: "Sequences"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Sequences {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Sequences",
				Label: "Sequences",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Alls
			div := (&form.FormDiv{Name: "Alls"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Alls {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Alls",
				Label: "Alls",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Choices
			div := (&form.FormDiv{Name: "Choices"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Choices {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Choices",
				Label: "Choices",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Groups
			div := (&form.FormDiv{Name: "Groups"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Groups {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Groups",
				Label: "Groups",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Elements
			div := (&form.FormDiv{Name: "Elements"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Elements {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Elements",
				Label: "Elements",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("Order", inst.Order, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Depth", inst.Depth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MinOccurs", inst.MinOccurs, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MaxOccurs", inst.MaxOccurs, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Extension", inst.Extension, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Extension](), probe.formStage)
		StageSetAssociationFieldToForm("SimpleContent", inst.SimpleContent, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.SimpleContent](), probe.formStage)
		StageSetAssociationFieldToForm("ComplexContent", inst.ComplexContent, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ComplexContent](), probe.formStage)

		{
			// Slice of pointers: Attributes
			div := (&form.FormDiv{Name: "Attributes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Attributes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Attributes",
				Label: "Attributes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AttributeGroups
			div := (&form.FormDiv{Name: "AttributeGroups"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AttributeGroups {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AttributeGroups",
				Label: "AttributeGroups",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsDuplicatedInXSD", inst.IsDuplicatedInXSD, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Elements {
				if src.ComplexType == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Element", "ComplexType", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Schemas {
				for _, target := range src.ComplexTypes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Schema", "ComplexTypes", refNames, formGroup, probe.formStage)
		}
	case *models.Documentation:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Text", inst.Text, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Source", inst.Source, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Lang", inst.Lang, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Annotations {
				for _, target := range src.Documentations {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Annotation", "Documentations", refNames, formGroup, probe.formStage)
		}
	case *models.Element:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Order", inst.Order, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Depth", inst.Depth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasNameConflict", inst.HasNameConflict, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("GoIdentifier", inst.GoIdentifier, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Annotation", inst.Annotation, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Annotation](), probe.formStage)
		StageSetBasicFieldtoForm("NameXSD", inst.NameXSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MinOccurs", inst.MinOccurs, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MaxOccurs", inst.MaxOccurs, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Default", inst.Default, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Fixed", inst.Fixed, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Nillable", inst.Nillable, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Ref", inst.Ref, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Abstract", inst.Abstract, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Form", inst.Form, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Block", inst.Block, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Final", inst.Final, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("SimpleType", inst.SimpleType, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.SimpleType](), probe.formStage)
		StageSetAssociationFieldToForm("ComplexType", inst.ComplexType, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ComplexType](), probe.formStage)

		{
			// Slice of pointers: Groups
			div := (&form.FormDiv{Name: "Groups"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Groups {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Groups",
				Label: "Groups",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsDuplicatedInXSD", inst.IsDuplicatedInXSD, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Alls {
				for _, target := range src.Elements {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.All", "Elements", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Choices {
				for _, target := range src.Elements {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Choice", "Elements", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ComplexTypes {
				if src.OuterElement == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ComplexType", "OuterElement", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ComplexTypes {
				for _, target := range src.Elements {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ComplexType", "Elements", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Extensions {
				for _, target := range src.Elements {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Extension", "Elements", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Groups {
				if src.OuterElement == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Group", "OuterElement", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Groups {
				for _, target := range src.Elements {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Group", "Elements", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Schemas {
				for _, target := range src.Elements {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Schema", "Elements", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Sequences {
				for _, target := range src.Elements {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Sequence", "Elements", refNames, formGroup, probe.formStage)
		}
	case *models.Enumeration:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Annotation", inst.Annotation, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Annotation](), probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Restrictions {
				for _, target := range src.Enumerations {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Restriction", "Enumerations", refNames, formGroup, probe.formStage)
		}
	case *models.Extension:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("OuterElementName", inst.OuterElementName, probe.formStage, formGroup)

		{
			// Slice of pointers: Sequences
			div := (&form.FormDiv{Name: "Sequences"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Sequences {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Sequences",
				Label: "Sequences",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Alls
			div := (&form.FormDiv{Name: "Alls"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Alls {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Alls",
				Label: "Alls",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Choices
			div := (&form.FormDiv{Name: "Choices"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Choices {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Choices",
				Label: "Choices",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Groups
			div := (&form.FormDiv{Name: "Groups"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Groups {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Groups",
				Label: "Groups",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Elements
			div := (&form.FormDiv{Name: "Elements"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Elements {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Elements",
				Label: "Elements",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("Order", inst.Order, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Depth", inst.Depth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MinOccurs", inst.MinOccurs, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MaxOccurs", inst.MaxOccurs, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Base", inst.Base, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Ref", inst.Ref, probe.formStage, formGroup)

		{
			// Slice of pointers: Attributes
			div := (&form.FormDiv{Name: "Attributes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Attributes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Attributes",
				Label: "Attributes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AttributeGroups
			div := (&form.FormDiv{Name: "AttributeGroups"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AttributeGroups {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AttributeGroups",
				Label: "AttributeGroups",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ComplexTypes {
				if src.Extension == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ComplexType", "Extension", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.SimpleContents {
				if src.Extension == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SimpleContent", "Extension", refNames, formGroup, probe.formStage)
		}
	case *models.Group:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Annotation", inst.Annotation, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Annotation](), probe.formStage)
		StageSetBasicFieldtoForm("NameXSD", inst.NameXSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Ref", inst.Ref, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsAnonymous", inst.IsAnonymous, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("OuterElement", inst.OuterElement, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Element](), probe.formStage)
		StageSetBasicFieldtoForm("HasNameConflict", inst.HasNameConflict, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("GoIdentifier", inst.GoIdentifier, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("OuterElementName", inst.OuterElementName, probe.formStage, formGroup)

		{
			// Slice of pointers: Sequences
			div := (&form.FormDiv{Name: "Sequences"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Sequences {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Sequences",
				Label: "Sequences",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Alls
			div := (&form.FormDiv{Name: "Alls"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Alls {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Alls",
				Label: "Alls",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Choices
			div := (&form.FormDiv{Name: "Choices"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Choices {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Choices",
				Label: "Choices",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Groups
			div := (&form.FormDiv{Name: "Groups"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Groups {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Groups",
				Label: "Groups",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Elements
			div := (&form.FormDiv{Name: "Elements"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Elements {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Elements",
				Label: "Elements",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("Order", inst.Order, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Depth", inst.Depth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MinOccurs", inst.MinOccurs, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MaxOccurs", inst.MaxOccurs, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Alls {
				for _, target := range src.Groups {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.All", "Groups", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Choices {
				for _, target := range src.Groups {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Choice", "Groups", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ComplexTypes {
				for _, target := range src.Groups {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ComplexType", "Groups", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Elements {
				for _, target := range src.Groups {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Element", "Groups", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Extensions {
				for _, target := range src.Groups {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Extension", "Groups", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Groups {
				for _, target := range src.Groups {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Group", "Groups", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Schemas {
				for _, target := range src.Groups {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Schema", "Groups", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Sequences {
				for _, target := range src.Groups {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Sequence", "Groups", refNames, formGroup, probe.formStage)
		}
	case *models.Length:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Annotation", inst.Annotation, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Annotation](), probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Restrictions {
				if src.Length == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Restriction", "Length", refNames, formGroup, probe.formStage)
		}
	case *models.MaxInclusive:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Annotation", inst.Annotation, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Annotation](), probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Restrictions {
				if src.MaxInclusive == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Restriction", "MaxInclusive", refNames, formGroup, probe.formStage)
		}
	case *models.MaxLength:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Annotation", inst.Annotation, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Annotation](), probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Restrictions {
				if src.MaxLength == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Restriction", "MaxLength", refNames, formGroup, probe.formStage)
		}
	case *models.MinInclusive:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Annotation", inst.Annotation, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Annotation](), probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Restrictions {
				if src.MinInclusive == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Restriction", "MinInclusive", refNames, formGroup, probe.formStage)
		}
	case *models.MinLength:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Annotation", inst.Annotation, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Annotation](), probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Restrictions {
				if src.MinLength == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Restriction", "MinLength", refNames, formGroup, probe.formStage)
		}
	case *models.Pattern:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Annotation", inst.Annotation, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Annotation](), probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Restrictions {
				if src.Pattern == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Restriction", "Pattern", refNames, formGroup, probe.formStage)
		}
	case *models.Restriction:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Annotation", inst.Annotation, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Annotation](), probe.formStage)
		StageSetBasicFieldtoForm("Base", inst.Base, probe.formStage, formGroup)

		{
			// Slice of pointers: Enumerations
			div := (&form.FormDiv{Name: "Enumerations"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Enumerations {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Enumerations",
				Label: "Enumerations",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("MinInclusive", inst.MinInclusive, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.MinInclusive](), probe.formStage)
		StageSetAssociationFieldToForm("MaxInclusive", inst.MaxInclusive, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.MaxInclusive](), probe.formStage)
		StageSetAssociationFieldToForm("Pattern", inst.Pattern, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Pattern](), probe.formStage)
		StageSetAssociationFieldToForm("WhiteSpace", inst.WhiteSpace, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.WhiteSpace](), probe.formStage)
		StageSetAssociationFieldToForm("MinLength", inst.MinLength, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.MinLength](), probe.formStage)
		StageSetAssociationFieldToForm("MaxLength", inst.MaxLength, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.MaxLength](), probe.formStage)
		StageSetAssociationFieldToForm("Length", inst.Length, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Length](), probe.formStage)
		StageSetAssociationFieldToForm("TotalDigit", inst.TotalDigit, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.TotalDigit](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.SimpleContents {
				if src.Restriction == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SimpleContent", "Restriction", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.SimpleTypes {
				if src.Restriction == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SimpleType", "Restriction", refNames, formGroup, probe.formStage)
		}
	case *models.Schema:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Xs", inst.Xs, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Annotation", inst.Annotation, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Annotation](), probe.formStage)

		{
			// Slice of pointers: Elements
			div := (&form.FormDiv{Name: "Elements"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Elements {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Elements",
				Label: "Elements",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: SimpleTypes
			div := (&form.FormDiv{Name: "SimpleTypes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SimpleTypes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SimpleTypes",
				Label: "SimpleTypes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ComplexTypes
			div := (&form.FormDiv{Name: "ComplexTypes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ComplexTypes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ComplexTypes",
				Label: "ComplexTypes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AttributeGroups
			div := (&form.FormDiv{Name: "AttributeGroups"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AttributeGroups {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AttributeGroups",
				Label: "AttributeGroups",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Groups
			div := (&form.FormDiv{Name: "Groups"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Groups {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Groups",
				Label: "Groups",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("Order", inst.Order, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Depth", inst.Depth, probe.formStage, formGroup)
	case *models.Sequence:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Annotation", inst.Annotation, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Annotation](), probe.formStage)
		StageSetBasicFieldtoForm("OuterElementName", inst.OuterElementName, probe.formStage, formGroup)

		{
			// Slice of pointers: Sequences
			div := (&form.FormDiv{Name: "Sequences"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Sequences {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Sequences",
				Label: "Sequences",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Alls
			div := (&form.FormDiv{Name: "Alls"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Alls {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Alls",
				Label: "Alls",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Choices
			div := (&form.FormDiv{Name: "Choices"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Choices {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Choices",
				Label: "Choices",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Groups
			div := (&form.FormDiv{Name: "Groups"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Groups {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Groups",
				Label: "Groups",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Elements
			div := (&form.FormDiv{Name: "Elements"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Elements {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Elements",
				Label: "Elements",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("Order", inst.Order, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Depth", inst.Depth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MinOccurs", inst.MinOccurs, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MaxOccurs", inst.MaxOccurs, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Alls {
				for _, target := range src.Sequences {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.All", "Sequences", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Choices {
				for _, target := range src.Sequences {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Choice", "Sequences", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ComplexTypes {
				for _, target := range src.Sequences {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ComplexType", "Sequences", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Extensions {
				for _, target := range src.Sequences {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Extension", "Sequences", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Groups {
				for _, target := range src.Sequences {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Group", "Sequences", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Sequences {
				for _, target := range src.Sequences {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Sequence", "Sequences", refNames, formGroup, probe.formStage)
		}
	case *models.SimpleContent:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Extension", inst.Extension, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Extension](), probe.formStage)
		StageSetAssociationFieldToForm("Restriction", inst.Restriction, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Restriction](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ComplexTypes {
				if src.SimpleContent == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ComplexType", "SimpleContent", refNames, formGroup, probe.formStage)
		}
	case *models.SimpleType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Annotation", inst.Annotation, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Annotation](), probe.formStage)
		StageSetBasicFieldtoForm("NameXSD", inst.NameXSD, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Restriction", inst.Restriction, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Restriction](), probe.formStage)
		StageSetAssociationFieldToForm("Union", inst.Union, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Union](), probe.formStage)
		StageSetBasicFieldtoForm("Order", inst.Order, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Depth", inst.Depth, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Elements {
				if src.SimpleType == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Element", "SimpleType", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Schemas {
				for _, target := range src.SimpleTypes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Schema", "SimpleTypes", refNames, formGroup, probe.formStage)
		}
	case *models.TotalDigit:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Annotation", inst.Annotation, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Annotation](), probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Restrictions {
				if src.TotalDigit == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Restriction", "TotalDigit", refNames, formGroup, probe.formStage)
		}
	case *models.Union:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Annotation", inst.Annotation, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Annotation](), probe.formStage)
		StageSetBasicFieldtoForm("MemberTypes", inst.MemberTypes, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.SimpleTypes {
				if src.Union == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SimpleType", "Union", refNames, formGroup, probe.formStage)
		}
	case *models.WhiteSpace:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Annotation", inst.Annotation, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Annotation](), probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Restrictions {
				if src.WhiteSpace == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Restriction", "WhiteSpace", refNames, formGroup, probe.formStage)
		}
	default:
		_ = inst
	}
}
