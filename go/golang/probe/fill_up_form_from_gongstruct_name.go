package probe

const FillUpFormFromGongstructNameTemplate = `// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"{{PkgPathRoot}}/models"
)

func FillUpFormFromGongstructName(
	playground *Playground,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := playground.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = "New"
	} else {
		prefix = "Update"
	}

	switch gongstructName {
	// insertion point{{` + string(rune(FillUpFormFromGongstructNameSwitchCase)) + `}}
	}
	formStage.Commit()
}
`

type FillUpFormFromGongstructNameInsertionId int

const (
	FillUpFormFromGongstructNameSwitchCase FillUpFormFromGongstructNameInsertionId = iota
	FillUpFormFromGongstructNameInsertionNb
)

var FillUpFormFromGongstructNameSubTemplateCode map[string]string = // new line
map[string]string{
	string(rune(FillUpTreeStructCase)): `
	case "{{Structname}}":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " {{Structname}} Form",
			OnSave: New{{Structname}}FormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		{{structname}} := new(models.{{Structname}})
		FillUpForm({{structname}}, formGroup, playground)`,
}
