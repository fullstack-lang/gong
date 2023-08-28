// generated code - do not edit
package data

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	table "github.com/fullstack-lang/gongtable/go/models"
	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/gong/test/go/models"
)

type ButtonImplGongstruct struct {
	gongStruct      *gong_models.GongStruct
	Icon            gongtree_buttons.ButtonType
	formStage       *table.StageStruct
	stageOfInterest *models.StageStruct
}

func NewButtonImplGongstruct(
	gongStruct *gong_models.GongStruct,
	icon gongtree_buttons.ButtonType,
	formStage *table.StageStruct,
	stageOfInterest *models.StageStruct,
) (buttonImplGongstruct *ButtonImplGongstruct) {

	buttonImplGongstruct = new(ButtonImplGongstruct)
	buttonImplGongstruct.Icon = icon
	buttonImplGongstruct.gongStruct = gongStruct
	buttonImplGongstruct.formStage = formStage
	buttonImplGongstruct.stageOfInterest = stageOfInterest

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
		Name:   table.FormGroupDefaultName.ToString(),
		OnSave: NewAstructFormCallback(buttonImpl.stageOfInterest, formStage),
	}).Stage(formStage)

	switch buttonImpl.gongStruct.Name {
	case "Astruct":

		// Name field
		{
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

		// Date field
		{
			formDiv := (&table.FormDiv{
				Name: "Date",
			}).Stage(formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, formDiv)
			{
				formFieldPartDate := (&table.FormField{
					Name:        "Date",
					Label:       "Date",
					Placeholder: "",
				}).Stage(formStage)
				formDiv.FormFields = append(formDiv.FormFields, formFieldPartDate)

				formFieldDate := (&table.FormFieldDate{
					Name: "Date",
				}).Stage(formStage)
				formFieldPartDate.FormFieldDate = formFieldDate
			}
			{
				formFieldPartTime := (&table.FormField{
					Name:        "Time",
					Label:       "Time",
					Placeholder: "",
				}).Stage(formStage)
				formDiv.FormFields = append(formDiv.FormFields, formFieldPartTime)

				formFieldTime := (&table.FormFieldTime{
					Name: "Time",
				}).Stage(formStage)
				formFieldPartTime.FormFieldTime = formFieldTime
			}

		}

		// Booleanfield
		{
			formDiv := (&table.FormDiv{
				Name: "Booleanfield",
			}).Stage(formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, formDiv)

			checkBox := (&table.CheckBox{
				Name: "Booleanfield",
			}).Stage(formStage)
			formDiv.CheckBoxs = append(formDiv.CheckBoxs, checkBox)
		}
	}
	formStage.Commit()
}
