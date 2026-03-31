// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/dsm/xsd/go/models"
)

// ux_form updates the current form if there is one
func (probe *Probe) ux_form() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		switch onSave := formGroup.OnSave.(type) { // insertion point
		case *AllFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "All", true)
			} else {
				FillUpFormFromGongstruct(onSave.all, probe)
			}
		case *AnnotationFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Annotation", true)
			} else {
				FillUpFormFromGongstruct(onSave.annotation, probe)
			}
		case *AttributeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Attribute", true)
			} else {
				FillUpFormFromGongstruct(onSave.attribute, probe)
			}
		case *AttributeGroupFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "AttributeGroup", true)
			} else {
				FillUpFormFromGongstruct(onSave.attributegroup, probe)
			}
		case *ChoiceFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Choice", true)
			} else {
				FillUpFormFromGongstruct(onSave.choice, probe)
			}
		case *ComplexContentFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ComplexContent", true)
			} else {
				FillUpFormFromGongstruct(onSave.complexcontent, probe)
			}
		case *ComplexTypeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ComplexType", true)
			} else {
				FillUpFormFromGongstruct(onSave.complextype, probe)
			}
		case *DocumentationFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Documentation", true)
			} else {
				FillUpFormFromGongstruct(onSave.documentation, probe)
			}
		case *ElementFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Element", true)
			} else {
				FillUpFormFromGongstruct(onSave.element, probe)
			}
		case *EnumerationFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Enumeration", true)
			} else {
				FillUpFormFromGongstruct(onSave.enumeration, probe)
			}
		case *ExtensionFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Extension", true)
			} else {
				FillUpFormFromGongstruct(onSave.extension, probe)
			}
		case *GroupFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Group", true)
			} else {
				FillUpFormFromGongstruct(onSave.group, probe)
			}
		case *LengthFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Length", true)
			} else {
				FillUpFormFromGongstruct(onSave.length, probe)
			}
		case *MaxInclusiveFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "MaxInclusive", true)
			} else {
				FillUpFormFromGongstruct(onSave.maxinclusive, probe)
			}
		case *MaxLengthFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "MaxLength", true)
			} else {
				FillUpFormFromGongstruct(onSave.maxlength, probe)
			}
		case *MinInclusiveFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "MinInclusive", true)
			} else {
				FillUpFormFromGongstruct(onSave.mininclusive, probe)
			}
		case *MinLengthFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "MinLength", true)
			} else {
				FillUpFormFromGongstruct(onSave.minlength, probe)
			}
		case *PatternFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Pattern", true)
			} else {
				FillUpFormFromGongstruct(onSave.pattern, probe)
			}
		case *RestrictionFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Restriction", true)
			} else {
				FillUpFormFromGongstruct(onSave.restriction, probe)
			}
		case *SchemaFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Schema", true)
			} else {
				FillUpFormFromGongstruct(onSave.schema, probe)
			}
		case *SequenceFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Sequence", true)
			} else {
				FillUpFormFromGongstruct(onSave.sequence, probe)
			}
		case *SimpleContentFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "SimpleContent", true)
			} else {
				FillUpFormFromGongstruct(onSave.simplecontent, probe)
			}
		case *SimpleTypeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "SimpleType", true)
			} else {
				FillUpFormFromGongstruct(onSave.simpletype, probe)
			}
		case *TotalDigitFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "TotalDigit", true)
			} else {
				FillUpFormFromGongstruct(onSave.totaldigit, probe)
			}
		case *UnionFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Union", true)
			} else {
				FillUpFormFromGongstruct(onSave.union, probe)
			}
		case *WhiteSpaceFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "WhiteSpace", true)
			} else {
				FillUpFormFromGongstruct(onSave.whitespace, probe)
			}
		}
	}
}

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "All":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "All Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AllFormCallback(
			nil,
			probe,
			formGroup,
		)
		all := new(models.All)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(all, formGroup, probe)
	case "Annotation":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Annotation Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AnnotationFormCallback(
			nil,
			probe,
			formGroup,
		)
		annotation := new(models.Annotation)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(annotation, formGroup, probe)
	case "Attribute":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Attribute Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AttributeFormCallback(
			nil,
			probe,
			formGroup,
		)
		attribute := new(models.Attribute)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attribute, formGroup, probe)
	case "AttributeGroup":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "AttributeGroup Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AttributeGroupFormCallback(
			nil,
			probe,
			formGroup,
		)
		attributegroup := new(models.AttributeGroup)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attributegroup, formGroup, probe)
	case "Choice":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Choice Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ChoiceFormCallback(
			nil,
			probe,
			formGroup,
		)
		choice := new(models.Choice)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(choice, formGroup, probe)
	case "ComplexContent":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ComplexContent Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ComplexContentFormCallback(
			nil,
			probe,
			formGroup,
		)
		complexcontent := new(models.ComplexContent)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(complexcontent, formGroup, probe)
	case "ComplexType":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ComplexType Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ComplexTypeFormCallback(
			nil,
			probe,
			formGroup,
		)
		complextype := new(models.ComplexType)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(complextype, formGroup, probe)
	case "Documentation":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Documentation Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DocumentationFormCallback(
			nil,
			probe,
			formGroup,
		)
		documentation := new(models.Documentation)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(documentation, formGroup, probe)
	case "Element":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Element Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ElementFormCallback(
			nil,
			probe,
			formGroup,
		)
		element := new(models.Element)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(element, formGroup, probe)
	case "Enumeration":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Enumeration Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EnumerationFormCallback(
			nil,
			probe,
			formGroup,
		)
		enumeration := new(models.Enumeration)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(enumeration, formGroup, probe)
	case "Extension":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Extension Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ExtensionFormCallback(
			nil,
			probe,
			formGroup,
		)
		extension := new(models.Extension)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(extension, formGroup, probe)
	case "Group":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Group Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GroupFormCallback(
			nil,
			probe,
			formGroup,
		)
		group := new(models.Group)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(group, formGroup, probe)
	case "Length":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Length Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LengthFormCallback(
			nil,
			probe,
			formGroup,
		)
		length := new(models.Length)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(length, formGroup, probe)
	case "MaxInclusive":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "MaxInclusive Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MaxInclusiveFormCallback(
			nil,
			probe,
			formGroup,
		)
		maxinclusive := new(models.MaxInclusive)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(maxinclusive, formGroup, probe)
	case "MaxLength":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "MaxLength Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MaxLengthFormCallback(
			nil,
			probe,
			formGroup,
		)
		maxlength := new(models.MaxLength)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(maxlength, formGroup, probe)
	case "MinInclusive":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "MinInclusive Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MinInclusiveFormCallback(
			nil,
			probe,
			formGroup,
		)
		mininclusive := new(models.MinInclusive)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(mininclusive, formGroup, probe)
	case "MinLength":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "MinLength Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MinLengthFormCallback(
			nil,
			probe,
			formGroup,
		)
		minlength := new(models.MinLength)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(minlength, formGroup, probe)
	case "Pattern":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Pattern Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PatternFormCallback(
			nil,
			probe,
			formGroup,
		)
		pattern := new(models.Pattern)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(pattern, formGroup, probe)
	case "Restriction":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Restriction Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RestrictionFormCallback(
			nil,
			probe,
			formGroup,
		)
		restriction := new(models.Restriction)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(restriction, formGroup, probe)
	case "Schema":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Schema Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SchemaFormCallback(
			nil,
			probe,
			formGroup,
		)
		schema := new(models.Schema)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(schema, formGroup, probe)
	case "Sequence":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Sequence Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SequenceFormCallback(
			nil,
			probe,
			formGroup,
		)
		sequence := new(models.Sequence)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(sequence, formGroup, probe)
	case "SimpleContent":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SimpleContent Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SimpleContentFormCallback(
			nil,
			probe,
			formGroup,
		)
		simplecontent := new(models.SimpleContent)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(simplecontent, formGroup, probe)
	case "SimpleType":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SimpleType Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SimpleTypeFormCallback(
			nil,
			probe,
			formGroup,
		)
		simpletype := new(models.SimpleType)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(simpletype, formGroup, probe)
	case "TotalDigit":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TotalDigit Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TotalDigitFormCallback(
			nil,
			probe,
			formGroup,
		)
		totaldigit := new(models.TotalDigit)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(totaldigit, formGroup, probe)
	case "Union":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Union Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__UnionFormCallback(
			nil,
			probe,
			formGroup,
		)
		union := new(models.Union)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(union, formGroup, probe)
	case "WhiteSpace":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "WhiteSpace Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__WhiteSpaceFormCallback(
			nil,
			probe,
			formGroup,
		)
		whitespace := new(models.WhiteSpace)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(whitespace, formGroup, probe)
	}
	formStage.Commit()
}
