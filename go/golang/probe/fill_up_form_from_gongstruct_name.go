package probe

const FillUpFormFromGongstructNameTemplate = `// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"{{PkgPathRoot}}/models"
)

// updateFillUpForm updates the current form if there is one
func (probe *Probe) updateFillUpForm() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		switch onSave := formGroup.OnSave.(type) { // insertion point{{` + string(rune(FillUpFormFromGongstructNameSwitchCase1)) + `}}
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
	// insertion point{{` + string(rune(FillUpFormFromGongstructNameSwitchCase)) + `}}
	}
	formStage.Commit()
}
`

type FillUpFormFromGongstructNameInsertionId int

const (
	FillUpFormFromGongstructNameSwitchCase FillUpFormFromGongstructNameInsertionId = iota
	FillUpFormFromGongstructNameSwitchCase1
	FillUpFormFromGongstructNameInsertionNb
)

var FillUpFormFromGongstructNameSubTemplateCode map[string]string = // new line
map[string]string{
	string(rune(FillUpFormFromGongstructNameSwitchCase)): `
	case "{{Structname}}":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "{{Structname}} Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__{{Structname}}FormCallback(
			nil,
			probe,
			formGroup,
		)
		{{structname}} := new(models.{{Structname}})
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm({{structname}}, formGroup, probe)`,
	string(rune(FillUpFormFromGongstructNameSwitchCase1)): `
		case *{{Structname}}FormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "{{Structname}}", true)
			} else {
				FillUpFormFromGongstruct(onSave.{{structname}}, probe)
			}`,
}
