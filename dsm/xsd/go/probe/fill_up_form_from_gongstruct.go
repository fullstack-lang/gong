// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/xsd/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.All:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "All",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AllFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Annotation:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Annotation",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AnnotationFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Attribute:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Attribute",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AttributeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.AttributeGroup:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "AttributeGroup",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AttributeGroupFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Choice:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Choice",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ChoiceFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ComplexContent:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ComplexContent",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ComplexContentFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ComplexType:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ComplexType",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ComplexTypeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Documentation:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Documentation",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DocumentationFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Element:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Element",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ElementFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Enumeration:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Enumeration",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EnumerationFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Extension:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Extension",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ExtensionFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Group:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Group",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GroupFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Length:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Length",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LengthFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.MaxInclusive:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "MaxInclusive",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MaxInclusiveFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.MaxLength:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "MaxLength",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MaxLengthFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.MinInclusive:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "MinInclusive",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MinInclusiveFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.MinLength:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "MinLength",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MinLengthFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Pattern:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Pattern",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PatternFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Restriction:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Restriction",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RestrictionFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Schema:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Schema",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SchemaFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Sequence:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Sequence",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SequenceFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SimpleContent:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "SimpleContent",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SimpleContentFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SimpleType:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "SimpleType",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SimpleTypeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TotalDigit:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "TotalDigit",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TotalDigitFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Union:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Union",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__UnionFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.WhiteSpace:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "WhiteSpace",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__WhiteSpaceFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}

	if probe.GetCommitMode() {
		formStage.Commit()
	}
}
