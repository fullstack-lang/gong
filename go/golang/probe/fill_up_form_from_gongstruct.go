package probe

const FillUpFormFromGongstructTemplate = `// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"{{PkgPathRoot}}/models"
)

func FillUpFormFromGongstruct[T models.Gongstruct](instance *T, playground *Playground) {
	formStage := playground.formStage
	formStage.Reset()
	formStage.Commit()

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
	string(rune(ButtonImplPerGongstructCallToForm)): `
	case *models.{{Structname}}:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update {{Structname}} Form",
			OnSave: New{{Structname}}FormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)`,
}
