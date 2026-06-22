package probe

const FillUpFormFromGongstructTemplate = `// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"{{PkgPathRoot}}/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point{{` + string(rune(FillUpFormFromGongstructSwitchCase)) + `}}
	default:
		_ = instancesTyped
	}

	if probe.GetCommitMode() {
		formStage.Commit()
	}
}
`

type FillUpFormFromGongstructInsertionId int

const (
	FillUpFormFromGongstructSwitchCase FillUpFormFromGongstructInsertionId = iota
	FillUpFormFromGongstructInsertionNb
)

var FillUpFormFromGongstructSubTemplateCode map[string]string = // new line
map[string]string{
	string(rune(FillUpTreeStructCase)): `
	case *models.{{Structname}}:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "{{Structname}}",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__{{Structname}}FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)`,
}
