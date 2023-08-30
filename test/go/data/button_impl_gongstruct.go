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
		OnSave: NewAstructFormCallback(buttonImpl.stageOfInterest, formStage, nil),
	}).Stage(formStage)

	switch buttonImpl.gongStruct.Name {
	case "Astruct":
		astruct := new(models.Astruct)
		FillUpForm(astruct, formStage, formGroup)
	}
	formStage.Commit()
}

func FillUpForm[T models.Gongstruct](instance *T, formStage *form.StageStruct, formGroup *form.FormGroup) {

	switch instancesTyped := any(instance).(type) {
	case *models.Astruct:
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
				Name:  "Name",
				Value: instancesTyped.Name,
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
					Name:  "Date",
					Value: instancesTyped.Date,
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
					Name:  "Time",
					Value: instancesTyped.Date,
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
				Name:  "Booleanfield",
				Value: instancesTyped.Booleanfield,
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
				if instancesTyped.Aenum == models.ENUM_VAL1 {
					formFieldSelect.Value = option
				}

				formField.FormFieldSelect.Options =
					append(formField.FormFieldSelect.Options, option)
			}
			{
				option := (&form.Option{
					Name: models.ENUM_VAL2.ToString(),
				}).Stage(formStage)
				if instancesTyped.Aenum == models.ENUM_VAL2 {
					formFieldSelect.Value = option
				}
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
				if instancesTyped.CEnum == models.CENUM_VAL1 {
					formFieldSelect.Value = option
				}

				formField.FormFieldSelect.Options =
					append(formField.FormFieldSelect.Options, option)
			}
			{
				option := (&form.Option{
					Name: "CENUM_VAL2",
				}).Stage(formStage)
				if instancesTyped.CEnum == models.CENUM_VAL2 {
					formFieldSelect.Value = option
				}

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
				Name:  "Intfield",
				Value: instancesTyped.Intfield,
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
				formFieldHours.HasBespokeWidth = true
				formFieldHours.BespokeWidthPx = 90

				formDiv.FormFields = append(formDiv.FormFields, formFieldHours)

				formFieldIntHours := (&form.FormFieldInt{
					Name:  "Hours",
					Value: int(instancesTyped.Duration1.Hours()) % 24,
				}).Stage(formStage)
				formFieldIntHours.HasMaxValidator = true
				formFieldIntHours.MaxValue = 23
				formFieldIntHours.HasMinValidator = true
				formFieldIntHours.MinValue = 0
				formFieldHours.FormFieldInt = formFieldIntHours
			}

			{
				formFieldMinutes := (&form.FormField{
					Name:  "Minutes",
					Label: "Minutes",
				}).Stage(formStage)
				formFieldMinutes.HasBespokeWidth = true
				formFieldMinutes.BespokeWidthPx = 90
				formDiv.FormFields = append(formDiv.FormFields, formFieldMinutes)

				formFieldIntMinutes := (&form.FormFieldInt{
					Name:  "Minutes",
					Value: int(instancesTyped.Duration1.Minutes()) % 60,
				}).Stage(formStage)
				formFieldIntMinutes.HasMaxValidator = true
				formFieldIntMinutes.MaxValue = 59
				formFieldIntMinutes.HasMinValidator = true
				formFieldIntMinutes.MinValue = 0
				formFieldMinutes.FormFieldInt = formFieldIntMinutes
			}
			{
				formFieldSeconds := (&form.FormField{
					Name:  "Seconds",
					Label: "Seconds",
				}).Stage(formStage)
				formFieldSeconds.HasBespokeWidth = true
				formFieldSeconds.BespokeWidthPx = 90
				formDiv.FormFields = append(formDiv.FormFields, formFieldSeconds)

				formFieldIntSeconds := (&form.FormFieldInt{
					Name:  "Seconds",
					Value: int(instancesTyped.Duration1.Seconds()) % 60,
				}).Stage(formStage)
				formFieldIntSeconds.HasMaxValidator = true
				formFieldIntSeconds.MaxValue = 59
				formFieldIntSeconds.HasMinValidator = true
				formFieldIntSeconds.MinValue = 0
				formFieldSeconds.FormFieldInt = formFieldIntSeconds
			}
		}

		// Floatfield
		{
			formDiv := (&form.FormDiv{
				Name: "Floatfield",
			}).Stage(formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, formDiv)
			formField := (&form.FormField{
				Name:  "Floatfield",
				Label: "Floatfield",
			}).Stage(formStage)
			formDiv.FormFields = append(formDiv.FormFields, formField)

			formFieldFloat64 := (&form.FormFieldFloat64{
				Name:  "Floatfield",
				Value: instancesTyped.Floatfield,
			}).Stage(formStage)
			formField.FormFieldFloat64 = formFieldFloat64
		}
	}
}
