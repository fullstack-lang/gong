// generated code - do not edit
package data

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	table "github.com/fullstack-lang/gongtable/go/models"
	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type ButtonImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	Icon       gongtree_buttons.ButtonType
	formStage  *table.StageStruct
}

func NewButtonImplGongstruct(
	gongStruct *gong_models.GongStruct,
	icon gongtree_buttons.ButtonType,
	formStage *table.StageStruct,
) (buttonImplGongstruct *ButtonImplGongstruct) {

	buttonImplGongstruct = new(ButtonImplGongstruct)
	buttonImplGongstruct.Icon = icon
	buttonImplGongstruct.gongStruct = gongStruct
	buttonImplGongstruct.formStage = formStage

	return
}

func (buttonImpl *ButtonImplGongstruct) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	log.Println("ButtonImplGongstruct: ButtonUpdated")

	formStage := buttonImpl.formStage
	formStage.Reset()
	formStage.Commit()

	formGroup := (&table.FormGroup{
		Name: table.FormGroupDefaultName.ToString(),
	}).Stage(formStage)

	switch buttonImpl.gongStruct.Name {
	case "Astruct":
		formDiv := (&table.FormDiv{
			Name: "Name",
		}).Stage(formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDiv)
		formField := (&table.FormField{
			Name:        "Name",
			Label:       "Name",
			Placeholder: "Astruct",
		}).Stage(formStage)
		formDiv.FormFields = append(formDiv.FormFields, formField)

		formFieldString := (&table.FormFieldString{
			Name: "Name",
		}).Stage(formStage)
		formField.FormFieldString = formFieldString
	}
	formStage.Commit()
}
