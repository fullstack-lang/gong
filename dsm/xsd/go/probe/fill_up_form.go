// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/xsd/go/models"
)

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.All:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("OuterElementName", instanceWithInferedType.OuterElementName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Sequences", instanceWithInferedType, &instanceWithInferedType.Sequences, formGroup, probe)
		AssociationSliceToForm("Alls", instanceWithInferedType, &instanceWithInferedType.Alls, formGroup, probe)
		AssociationSliceToForm("Choices", instanceWithInferedType, &instanceWithInferedType.Choices, formGroup, probe)
		AssociationSliceToForm("Groups", instanceWithInferedType, &instanceWithInferedType.Groups, formGroup, probe)
		AssociationSliceToForm("Elements", instanceWithInferedType, &instanceWithInferedType.Elements, formGroup, probe)
		BasicFieldtoForm("Order", instanceWithInferedType.Order, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Depth", instanceWithInferedType.Depth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MinOccurs", instanceWithInferedType.MinOccurs, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MaxOccurs", instanceWithInferedType.MaxOccurs, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.All, *models.All](
				"All",
				"Alls",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.All) []*models.All {
					return owner.Alls
				})
		}
		{
			AssociationReverseSliceToForm[*models.Choice, *models.All](
				"Choice",
				"Alls",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Choice) []*models.All {
					return owner.Alls
				})
		}
		{
			AssociationReverseSliceToForm[*models.ComplexType, *models.All](
				"ComplexType",
				"Alls",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.ComplexType) []*models.All {
					return owner.Alls
				})
		}
		{
			AssociationReverseSliceToForm[*models.Extension, *models.All](
				"Extension",
				"Alls",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Extension) []*models.All {
					return owner.Alls
				})
		}
		{
			AssociationReverseSliceToForm[*models.Group, *models.All](
				"Group",
				"Alls",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Group) []*models.All {
					return owner.Alls
				})
		}
		{
			AssociationReverseSliceToForm[*models.Sequence, *models.All](
				"Sequence",
				"Alls",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Sequence) []*models.All {
					return owner.Alls
				})
		}

	case *models.Annotation:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Documentations", instanceWithInferedType, &instanceWithInferedType.Documentations, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Attribute:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("NameXSD", instanceWithInferedType.NameXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("HasNameConflict", instanceWithInferedType.HasNameConflict, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("GoIdentifier", instanceWithInferedType.GoIdentifier, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default", instanceWithInferedType.Default, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Use", instanceWithInferedType.Use, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Form", instanceWithInferedType.Form, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Fixed", instanceWithInferedType.Fixed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Ref", instanceWithInferedType.Ref, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("TargetNamespace", instanceWithInferedType.TargetNamespace, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("SimpleType", instanceWithInferedType.SimpleType, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IDXSD", instanceWithInferedType.IDXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.AttributeGroup, *models.Attribute](
				"AttributeGroup",
				"Attributes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.AttributeGroup) []*models.Attribute {
					return owner.Attributes
				})
		}
		{
			AssociationReverseSliceToForm[*models.ComplexType, *models.Attribute](
				"ComplexType",
				"Attributes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.ComplexType) []*models.Attribute {
					return owner.Attributes
				})
		}
		{
			AssociationReverseSliceToForm[*models.Extension, *models.Attribute](
				"Extension",
				"Attributes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Extension) []*models.Attribute {
					return owner.Attributes
				})
		}

	case *models.AttributeGroup:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("NameXSD", instanceWithInferedType.NameXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("HasNameConflict", instanceWithInferedType.HasNameConflict, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("GoIdentifier", instanceWithInferedType.GoIdentifier, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("AttributeGroups", instanceWithInferedType, &instanceWithInferedType.AttributeGroups, formGroup, probe)
		BasicFieldtoForm("Ref", instanceWithInferedType.Ref, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Attributes", instanceWithInferedType, &instanceWithInferedType.Attributes, formGroup, probe)
		BasicFieldtoForm("Order", instanceWithInferedType.Order, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Depth", instanceWithInferedType.Depth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.AttributeGroup, *models.AttributeGroup](
				"AttributeGroup",
				"AttributeGroups",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.AttributeGroup) []*models.AttributeGroup {
					return owner.AttributeGroups
				})
		}
		{
			AssociationReverseSliceToForm[*models.ComplexType, *models.AttributeGroup](
				"ComplexType",
				"AttributeGroups",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.ComplexType) []*models.AttributeGroup {
					return owner.AttributeGroups
				})
		}
		{
			AssociationReverseSliceToForm[*models.Extension, *models.AttributeGroup](
				"Extension",
				"AttributeGroups",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Extension) []*models.AttributeGroup {
					return owner.AttributeGroups
				})
		}
		{
			AssociationReverseSliceToForm[*models.Schema, *models.AttributeGroup](
				"Schema",
				"AttributeGroups",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Schema) []*models.AttributeGroup {
					return owner.AttributeGroups
				})
		}

	case *models.Choice:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("OuterElementName", instanceWithInferedType.OuterElementName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Sequences", instanceWithInferedType, &instanceWithInferedType.Sequences, formGroup, probe)
		AssociationSliceToForm("Alls", instanceWithInferedType, &instanceWithInferedType.Alls, formGroup, probe)
		AssociationSliceToForm("Choices", instanceWithInferedType, &instanceWithInferedType.Choices, formGroup, probe)
		AssociationSliceToForm("Groups", instanceWithInferedType, &instanceWithInferedType.Groups, formGroup, probe)
		AssociationSliceToForm("Elements", instanceWithInferedType, &instanceWithInferedType.Elements, formGroup, probe)
		BasicFieldtoForm("Order", instanceWithInferedType.Order, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Depth", instanceWithInferedType.Depth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MinOccurs", instanceWithInferedType.MinOccurs, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MaxOccurs", instanceWithInferedType.MaxOccurs, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsDuplicatedInXSD", instanceWithInferedType.IsDuplicatedInXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.All, *models.Choice](
				"All",
				"Choices",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.All) []*models.Choice {
					return owner.Choices
				})
		}
		{
			AssociationReverseSliceToForm[*models.Choice, *models.Choice](
				"Choice",
				"Choices",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Choice) []*models.Choice {
					return owner.Choices
				})
		}
		{
			AssociationReverseSliceToForm[*models.ComplexType, *models.Choice](
				"ComplexType",
				"Choices",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.ComplexType) []*models.Choice {
					return owner.Choices
				})
		}
		{
			AssociationReverseSliceToForm[*models.Extension, *models.Choice](
				"Extension",
				"Choices",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Extension) []*models.Choice {
					return owner.Choices
				})
		}
		{
			AssociationReverseSliceToForm[*models.Group, *models.Choice](
				"Group",
				"Choices",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Group) []*models.Choice {
					return owner.Choices
				})
		}
		{
			AssociationReverseSliceToForm[*models.Sequence, *models.Choice](
				"Sequence",
				"Choices",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Sequence) []*models.Choice {
					return owner.Choices
				})
		}

	case *models.ComplexContent:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.ComplexType:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("HasNameConflict", instanceWithInferedType.HasNameConflict, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("GoIdentifier", instanceWithInferedType.GoIdentifier, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsAnonymous", instanceWithInferedType.IsAnonymous, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("OuterElement", instanceWithInferedType.OuterElement, formGroup, probe)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("NameXSD", instanceWithInferedType.NameXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("OuterElementName", instanceWithInferedType.OuterElementName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Sequences", instanceWithInferedType, &instanceWithInferedType.Sequences, formGroup, probe)
		AssociationSliceToForm("Alls", instanceWithInferedType, &instanceWithInferedType.Alls, formGroup, probe)
		AssociationSliceToForm("Choices", instanceWithInferedType, &instanceWithInferedType.Choices, formGroup, probe)
		AssociationSliceToForm("Groups", instanceWithInferedType, &instanceWithInferedType.Groups, formGroup, probe)
		AssociationSliceToForm("Elements", instanceWithInferedType, &instanceWithInferedType.Elements, formGroup, probe)
		BasicFieldtoForm("Order", instanceWithInferedType.Order, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Depth", instanceWithInferedType.Depth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MinOccurs", instanceWithInferedType.MinOccurs, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MaxOccurs", instanceWithInferedType.MaxOccurs, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Extension", instanceWithInferedType.Extension, formGroup, probe)
		AssociationFieldToForm("SimpleContent", instanceWithInferedType.SimpleContent, formGroup, probe)
		AssociationFieldToForm("ComplexContent", instanceWithInferedType.ComplexContent, formGroup, probe)
		AssociationSliceToForm("Attributes", instanceWithInferedType, &instanceWithInferedType.Attributes, formGroup, probe)
		AssociationSliceToForm("AttributeGroups", instanceWithInferedType, &instanceWithInferedType.AttributeGroups, formGroup, probe)
		BasicFieldtoForm("IsDuplicatedInXSD", instanceWithInferedType.IsDuplicatedInXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Schema, *models.ComplexType](
				"Schema",
				"ComplexTypes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Schema) []*models.ComplexType {
					return owner.ComplexTypes
				})
		}

	case *models.Documentation:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Source", instanceWithInferedType.Source, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Lang", instanceWithInferedType.Lang, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Annotation, *models.Documentation](
				"Annotation",
				"Documentations",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Annotation) []*models.Documentation {
					return owner.Documentations
				})
		}

	case *models.Element:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Order", instanceWithInferedType.Order, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Depth", instanceWithInferedType.Depth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("HasNameConflict", instanceWithInferedType.HasNameConflict, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("GoIdentifier", instanceWithInferedType.GoIdentifier, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("NameXSD", instanceWithInferedType.NameXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MinOccurs", instanceWithInferedType.MinOccurs, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MaxOccurs", instanceWithInferedType.MaxOccurs, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Default", instanceWithInferedType.Default, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Fixed", instanceWithInferedType.Fixed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Nillable", instanceWithInferedType.Nillable, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Ref", instanceWithInferedType.Ref, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Abstract", instanceWithInferedType.Abstract, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Form", instanceWithInferedType.Form, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Block", instanceWithInferedType.Block, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Final", instanceWithInferedType.Final, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("SimpleType", instanceWithInferedType.SimpleType, formGroup, probe)
		AssociationFieldToForm("ComplexType", instanceWithInferedType.ComplexType, formGroup, probe)
		AssociationSliceToForm("Groups", instanceWithInferedType, &instanceWithInferedType.Groups, formGroup, probe)
		BasicFieldtoForm("IsDuplicatedInXSD", instanceWithInferedType.IsDuplicatedInXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.All, *models.Element](
				"All",
				"Elements",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.All) []*models.Element {
					return owner.Elements
				})
		}
		{
			AssociationReverseSliceToForm[*models.Choice, *models.Element](
				"Choice",
				"Elements",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Choice) []*models.Element {
					return owner.Elements
				})
		}
		{
			AssociationReverseSliceToForm[*models.ComplexType, *models.Element](
				"ComplexType",
				"Elements",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.ComplexType) []*models.Element {
					return owner.Elements
				})
		}
		{
			AssociationReverseSliceToForm[*models.Extension, *models.Element](
				"Extension",
				"Elements",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Extension) []*models.Element {
					return owner.Elements
				})
		}
		{
			AssociationReverseSliceToForm[*models.Group, *models.Element](
				"Group",
				"Elements",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Group) []*models.Element {
					return owner.Elements
				})
		}
		{
			AssociationReverseSliceToForm[*models.Schema, *models.Element](
				"Schema",
				"Elements",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Schema) []*models.Element {
					return owner.Elements
				})
		}
		{
			AssociationReverseSliceToForm[*models.Sequence, *models.Element](
				"Sequence",
				"Elements",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Sequence) []*models.Element {
					return owner.Elements
				})
		}

	case *models.Enumeration:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Restriction, *models.Enumeration](
				"Restriction",
				"Enumerations",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Restriction) []*models.Enumeration {
					return owner.Enumerations
				})
		}

	case *models.Extension:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("OuterElementName", instanceWithInferedType.OuterElementName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Sequences", instanceWithInferedType, &instanceWithInferedType.Sequences, formGroup, probe)
		AssociationSliceToForm("Alls", instanceWithInferedType, &instanceWithInferedType.Alls, formGroup, probe)
		AssociationSliceToForm("Choices", instanceWithInferedType, &instanceWithInferedType.Choices, formGroup, probe)
		AssociationSliceToForm("Groups", instanceWithInferedType, &instanceWithInferedType.Groups, formGroup, probe)
		AssociationSliceToForm("Elements", instanceWithInferedType, &instanceWithInferedType.Elements, formGroup, probe)
		BasicFieldtoForm("Order", instanceWithInferedType.Order, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Depth", instanceWithInferedType.Depth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MinOccurs", instanceWithInferedType.MinOccurs, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MaxOccurs", instanceWithInferedType.MaxOccurs, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Base", instanceWithInferedType.Base, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Ref", instanceWithInferedType.Ref, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Attributes", instanceWithInferedType, &instanceWithInferedType.Attributes, formGroup, probe)
		AssociationSliceToForm("AttributeGroups", instanceWithInferedType, &instanceWithInferedType.AttributeGroups, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Group:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("NameXSD", instanceWithInferedType.NameXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Ref", instanceWithInferedType.Ref, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsAnonymous", instanceWithInferedType.IsAnonymous, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("OuterElement", instanceWithInferedType.OuterElement, formGroup, probe)
		BasicFieldtoForm("HasNameConflict", instanceWithInferedType.HasNameConflict, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("GoIdentifier", instanceWithInferedType.GoIdentifier, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("OuterElementName", instanceWithInferedType.OuterElementName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Sequences", instanceWithInferedType, &instanceWithInferedType.Sequences, formGroup, probe)
		AssociationSliceToForm("Alls", instanceWithInferedType, &instanceWithInferedType.Alls, formGroup, probe)
		AssociationSliceToForm("Choices", instanceWithInferedType, &instanceWithInferedType.Choices, formGroup, probe)
		AssociationSliceToForm("Groups", instanceWithInferedType, &instanceWithInferedType.Groups, formGroup, probe)
		AssociationSliceToForm("Elements", instanceWithInferedType, &instanceWithInferedType.Elements, formGroup, probe)
		BasicFieldtoForm("Order", instanceWithInferedType.Order, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Depth", instanceWithInferedType.Depth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MinOccurs", instanceWithInferedType.MinOccurs, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MaxOccurs", instanceWithInferedType.MaxOccurs, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.All, *models.Group](
				"All",
				"Groups",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.All) []*models.Group {
					return owner.Groups
				})
		}
		{
			AssociationReverseSliceToForm[*models.Choice, *models.Group](
				"Choice",
				"Groups",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Choice) []*models.Group {
					return owner.Groups
				})
		}
		{
			AssociationReverseSliceToForm[*models.ComplexType, *models.Group](
				"ComplexType",
				"Groups",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.ComplexType) []*models.Group {
					return owner.Groups
				})
		}
		{
			AssociationReverseSliceToForm[*models.Element, *models.Group](
				"Element",
				"Groups",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Element) []*models.Group {
					return owner.Groups
				})
		}
		{
			AssociationReverseSliceToForm[*models.Extension, *models.Group](
				"Extension",
				"Groups",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Extension) []*models.Group {
					return owner.Groups
				})
		}
		{
			AssociationReverseSliceToForm[*models.Group, *models.Group](
				"Group",
				"Groups",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Group) []*models.Group {
					return owner.Groups
				})
		}
		{
			AssociationReverseSliceToForm[*models.Schema, *models.Group](
				"Schema",
				"Groups",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Schema) []*models.Group {
					return owner.Groups
				})
		}
		{
			AssociationReverseSliceToForm[*models.Sequence, *models.Group](
				"Sequence",
				"Groups",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Sequence) []*models.Group {
					return owner.Groups
				})
		}

	case *models.Length:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.MaxInclusive:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.MaxLength:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.MinInclusive:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.MinLength:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Pattern:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Restriction:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("Base", instanceWithInferedType.Base, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Enumerations", instanceWithInferedType, &instanceWithInferedType.Enumerations, formGroup, probe)
		AssociationFieldToForm("MinInclusive", instanceWithInferedType.MinInclusive, formGroup, probe)
		AssociationFieldToForm("MaxInclusive", instanceWithInferedType.MaxInclusive, formGroup, probe)
		AssociationFieldToForm("Pattern", instanceWithInferedType.Pattern, formGroup, probe)
		AssociationFieldToForm("WhiteSpace", instanceWithInferedType.WhiteSpace, formGroup, probe)
		AssociationFieldToForm("MinLength", instanceWithInferedType.MinLength, formGroup, probe)
		AssociationFieldToForm("MaxLength", instanceWithInferedType.MaxLength, formGroup, probe)
		AssociationFieldToForm("Length", instanceWithInferedType.Length, formGroup, probe)
		AssociationFieldToForm("TotalDigit", instanceWithInferedType.TotalDigit, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Schema:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Xs", instanceWithInferedType.Xs, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		AssociationSliceToForm("Elements", instanceWithInferedType, &instanceWithInferedType.Elements, formGroup, probe)
		AssociationSliceToForm("SimpleTypes", instanceWithInferedType, &instanceWithInferedType.SimpleTypes, formGroup, probe)
		AssociationSliceToForm("ComplexTypes", instanceWithInferedType, &instanceWithInferedType.ComplexTypes, formGroup, probe)
		AssociationSliceToForm("AttributeGroups", instanceWithInferedType, &instanceWithInferedType.AttributeGroups, formGroup, probe)
		AssociationSliceToForm("Groups", instanceWithInferedType, &instanceWithInferedType.Groups, formGroup, probe)
		BasicFieldtoForm("Order", instanceWithInferedType.Order, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Depth", instanceWithInferedType.Depth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Sequence:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("OuterElementName", instanceWithInferedType.OuterElementName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Sequences", instanceWithInferedType, &instanceWithInferedType.Sequences, formGroup, probe)
		AssociationSliceToForm("Alls", instanceWithInferedType, &instanceWithInferedType.Alls, formGroup, probe)
		AssociationSliceToForm("Choices", instanceWithInferedType, &instanceWithInferedType.Choices, formGroup, probe)
		AssociationSliceToForm("Groups", instanceWithInferedType, &instanceWithInferedType.Groups, formGroup, probe)
		AssociationSliceToForm("Elements", instanceWithInferedType, &instanceWithInferedType.Elements, formGroup, probe)
		BasicFieldtoForm("Order", instanceWithInferedType.Order, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Depth", instanceWithInferedType.Depth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MinOccurs", instanceWithInferedType.MinOccurs, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MaxOccurs", instanceWithInferedType.MaxOccurs, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.All, *models.Sequence](
				"All",
				"Sequences",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.All) []*models.Sequence {
					return owner.Sequences
				})
		}
		{
			AssociationReverseSliceToForm[*models.Choice, *models.Sequence](
				"Choice",
				"Sequences",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Choice) []*models.Sequence {
					return owner.Sequences
				})
		}
		{
			AssociationReverseSliceToForm[*models.ComplexType, *models.Sequence](
				"ComplexType",
				"Sequences",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.ComplexType) []*models.Sequence {
					return owner.Sequences
				})
		}
		{
			AssociationReverseSliceToForm[*models.Extension, *models.Sequence](
				"Extension",
				"Sequences",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Extension) []*models.Sequence {
					return owner.Sequences
				})
		}
		{
			AssociationReverseSliceToForm[*models.Group, *models.Sequence](
				"Group",
				"Sequences",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Group) []*models.Sequence {
					return owner.Sequences
				})
		}
		{
			AssociationReverseSliceToForm[*models.Sequence, *models.Sequence](
				"Sequence",
				"Sequences",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Sequence) []*models.Sequence {
					return owner.Sequences
				})
		}

	case *models.SimpleContent:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Extension", instanceWithInferedType.Extension, formGroup, probe)
		AssociationFieldToForm("Restriction", instanceWithInferedType.Restriction, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.SimpleType:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("NameXSD", instanceWithInferedType.NameXSD, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Restriction", instanceWithInferedType.Restriction, formGroup, probe)
		AssociationFieldToForm("Union", instanceWithInferedType.Union, formGroup, probe)
		BasicFieldtoForm("Order", instanceWithInferedType.Order, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Depth", instanceWithInferedType.Depth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Schema, *models.SimpleType](
				"Schema",
				"SimpleTypes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Schema) []*models.SimpleType {
					return owner.SimpleTypes
				})
		}

	case *models.TotalDigit:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Union:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("MemberTypes", instanceWithInferedType.MemberTypes, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.WhiteSpace:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Annotation", instanceWithInferedType.Annotation, formGroup, probe)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	default:
		_ = instanceWithInferedType
	}
}
