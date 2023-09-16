package probe

const ButtonImplGongstructFileTemplate = `// generated code - do not edit
package probe

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type ButtonImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	Icon       gongtree_buttons.ButtonType
	playground *Playground
}

func NewButtonImplGongstruct(
	gongStruct *gong_models.GongStruct,
	icon gongtree_buttons.ButtonType,
	playground *Playground,
) (buttonImplGongstruct *ButtonImplGongstruct) {

	buttonImplGongstruct = new(ButtonImplGongstruct)
	buttonImplGongstruct.Icon = icon
	buttonImplGongstruct.gongStruct = gongStruct
	buttonImplGongstruct.playground = playground

	return
}

func (buttonImpl *ButtonImplGongstruct) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	log.Println("ButtonImplGongstruct: ButtonUpdated")

	FillUpFormFromGongstructName(
		buttonImpl.playground,
		buttonImpl.gongStruct.Name,
		true,
	)
}
`

type ButtonImplGongstructInsertionId int

const (
	ButtonImplPerGongstructCallToForm ButtonImplGongstructInsertionId = iota
	ButtonImplGongstructInsertionNb
)

var ButtonImplGongstructSubTemplateCode map[string]string = // new line
map[string]string{
	string(rune(ButtonImplPerGongstructCallToForm)): `
	case "{{Structname}}":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: "New {{Structname}} Form",
			OnSave: New{{Structname}}FormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		{{structname}} := new(models.{{Structname}})
		FillUpForm({{structname}}, formGroup, buttonImpl.playground)`,
}
