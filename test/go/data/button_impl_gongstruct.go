// generated code - do not edit
package data

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	form "github.com/fullstack-lang/gongtable/go/models"
	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/gong/test/go/models"
)

type ButtonImplGongstruct struct {
	gongStruct      *gong_models.GongStruct
	Icon            gongtree_buttons.ButtonType
	formStage       *form.StageStruct
	stageOfInterest *models.StageStruct
}

func NewButtonImplGongstruct(
	gongStruct *gong_models.GongStruct,
	icon gongtree_buttons.ButtonType,
	formStage *form.StageStruct,
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

	formGroup := (&form.FormGroup{
		Name:   form.FormGroupDefaultName.ToString(),
		OnSave: NewAstructFormCallback(buttonImpl.stageOfInterest, formStage),
	}).Stage(formStage)

	switch buttonImpl.gongStruct.Name {
	case "Astruct":

		// Name field
		{
			formDiv := (&form.FormDiv{
				Name: "Name",
			}).Stage(formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, formDiv)
			formField := (&form.FormField{
				Name:        "Name",
				Label:       "Name",
				Placeholder: "Astruct",
			}).Stage(formStage)
			formDiv.FormFields = append(formDiv.FormFields, formField)

			formFieldString := (&form.FormFieldString{
				Name: "Name",
			}).Stage(formStage)
			formField.FormFieldString = formFieldString
		}

		// Date field
		{
			formDiv := (&form.FormDiv{
				Name: "Date",
			}).Stage(formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, formDiv)
			{
				formFieldPartDate := (&form.FormField{
					Name:        "Date",
					Label:       "Date",
					Placeholder: "",
				}).Stage(formStage)
				formDiv.FormFields = append(formDiv.FormFields, formFieldPartDate)

				formFieldDate := (&form.FormFieldDate{
					Name: "Date",
				}).Stage(formStage)
				formFieldPartDate.FormFieldDate = formFieldDate
			}
			{
				formFieldPartTime := (&form.FormField{
					Name:        "Time",
					Label:       "Time",
					Placeholder: "",
				}).Stage(formStage)
				formDiv.FormFields = append(formDiv.FormFields, formFieldPartTime)

				formFieldTime := (&form.FormFieldTime{
					Name: "Time",
				}).Stage(formStage)
				formFieldPartTime.FormFieldTime = formFieldTime
			}

		}

		// Booleanfield
		{
			formDiv := (&form.FormDiv{
				Name: "Booleanfield",
			}).Stage(formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, formDiv)

			checkBox := (&form.CheckBox{
				Name: "Booleanfield",
			}).Stage(formStage)
			formDiv.CheckBoxs = append(formDiv.CheckBoxs, checkBox)
		}

		// Aenum
		{
			formDiv := (&form.FormDiv{
				Name: "Aenum",
			}).Stage(formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, formDiv)
			formField := (&form.FormField{
				Name:        "Aenum",
				Label:       "Aenum",
				Placeholder: "",
			}).Stage(formStage)
			formDiv.FormFields = append(formDiv.FormFields, formField)

			formFieldSelect := (&form.FormFieldSelect{
				Name: "Aenum",
			}).Stage(formStage)
			formField.FormFieldSelect = formFieldSelect

			formField.FormFieldSelect.Options = make([]*form.Option, 0)
			{
				option := (&form.Option{
					Name: models.ENUM_VAL1.ToString(),
				}).Stage(formStage)

				formField.FormFieldSelect.Options =
					append(formField.FormFieldSelect.Options, option)
			}
			{
				option := (&form.Option{
					Name: models.ENUM_VAL2.ToString(),
				}).Stage(formStage)

				formField.FormFieldSelect.Options =
					append(formField.FormFieldSelect.Options, option)
			}
		}

		// Cenum
		{
			formDiv := (&form.FormDiv{
				Name: "Cenum",
			}).Stage(formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, formDiv)
			formField := (&form.FormField{
				Name:        "Cenum",
				Label:       "Cenum",
				Placeholder: "",
			}).Stage(formStage)
			formDiv.FormFields = append(formDiv.FormFields, formField)

			formFieldSelect := (&form.FormFieldSelect{
				Name: "Cenum",
			}).Stage(formStage)
			formField.FormFieldSelect = formFieldSelect

			formField.FormFieldSelect.Options = make([]*form.Option, 0)
			{
				option := (&form.Option{
					Name: "CENUM_VAL1",
				}).Stage(formStage)

				formField.FormFieldSelect.Options =
					append(formField.FormFieldSelect.Options, option)
			}
			{
				option := (&form.Option{
					Name: "CENUM_VAL2",
				}).Stage(formStage)

				formField.FormFieldSelect.Options =
					append(formField.FormFieldSelect.Options, option)
			}
		}

		// Intfield
		{
			formDiv := (&form.FormDiv{
				Name: "Intfield",
			}).Stage(formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, formDiv)
			formField := (&form.FormField{
				Name:  "Intfield",
				Label: "Intfield",
			}).Stage(formStage)
			formDiv.FormFields = append(formDiv.FormFields, formField)

			formFieldInt := (&form.FormFieldInt{
				Name: "Intfield",
			}).Stage(formStage)
			formField.FormFieldInt = formFieldInt
		}

		// Duration
		{
			formDiv := (&form.FormDiv{
				Name: "Duration1",
			}).Stage(formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, formDiv)

			{
				formFieldHours := (&form.FormField{
					Name:  "Hours",
					Label: "Hours",
				}).Stage(formStage)
				formDiv.FormFields = append(formDiv.FormFields, formFieldHours)
				formFieldIntHours := (&form.FormFieldInt{
					Name: "Hours",
				}).Stage(formStage)
				formFieldHours.FormFieldInt = formFieldIntHours
			}

			{
				formFieldMinutes := (&form.FormField{
					Name:  "Minutes",
					Label: "Minutes",
				}).Stage(formStage)
				formDiv.FormFields = append(formDiv.FormFields, formFieldMinutes)
				formFieldIntMinutes := (&form.FormFieldInt{
					Name: "Minutes",
				}).Stage(formStage)
				formFieldMinutes.FormFieldInt = formFieldIntMinutes
			}
			{
				formFieldSeconds := (&form.FormField{
					Name:  "Seconds",
					Label: "Seconds",
				}).Stage(formStage)
				formDiv.FormFields = append(formDiv.FormFields, formFieldSeconds)
				formFieldIntSeconds := (&form.FormFieldInt{
					Name: "Seconds",
				}).Stage(formStage)
				formFieldSeconds.FormFieldInt = formFieldIntSeconds
			}
		}
	}
	formStage.Commit()
}
