package probe

const FillUpFormFromGongstructTemplate = `// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"{{PkgPathRoot}}/models"
)

func FillUpFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	FillUpNamedFormFromGongstruct[T](instance, probe, formStage, gongtable.FormGroupDefaultName.ToString())

}

func FillUpNamedFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe, formStage *gongtable.StageStruct, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point{{` + string(rune(FillUpFormFromGongstructSwitchCase)) + `}}
	default:
		_ = instancesTyped
	}
	formStage.Commit()
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
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "{{Structname}} Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__{{Structname}}FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)`,
}
