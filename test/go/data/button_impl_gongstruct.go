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
		FillUpForm(astruct, buttonImpl.stageOfInterest, formStage, formGroup)
	}
	formStage.Commit()
}

func FillUpForm[T models.Gongstruct](instance *T, stageOfInterest *models.StageStruct, formStage *form.StageStruct, formGroup *form.FormGroup) {

	switch instanceWithInferedType := any(instance).(type) {
	case *models.Astruct:
		FillUpFormDivBasicField("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		FillUpFormDivBasicField("Date", instanceWithInferedType.Date, instanceWithInferedType, formStage, formGroup)
		FillUpFormDivBasicField("Booleanfield", instanceWithInferedType.Booleanfield, instanceWithInferedType, formStage, formGroup)
		FillUpFormDivBasicField("Intfield", instanceWithInferedType.Intfield, instanceWithInferedType, formStage, formGroup)
		FillUpFormDivBasicField("Floatfield", instanceWithInferedType.Floatfield, instanceWithInferedType, formStage, formGroup)
		FillUpFormDivBasicField("Duration1", instanceWithInferedType.Duration1, instanceWithInferedType, formStage, formGroup)

		FillUpFormDivEnumStringType("Aenum", instanceWithInferedType.Aenum, instanceWithInferedType, formStage, formGroup)
		FillUpFormDivEnumStringType("Aenum_2", instanceWithInferedType.Aenum_2, instanceWithInferedType, formStage, formGroup)
		FillUpFormDivEnumIntType("Cenum", instanceWithInferedType.CEnum, instanceWithInferedType, formStage, formGroup)

		{
			formDiv := (&form.FormDiv{
				Name: "Associationtob",
			}).Stage(formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, formDiv)
			formField := (&form.FormField{
				Name:        "Association",
				Label:       "Association",
				Placeholder: "",
			}).Stage(formStage)
			formDiv.FormFields = append(formDiv.FormFields, formField)

			formFieldSelect := (&form.FormFieldSelect{
				Name:       "association",
				CanBeEmpty: true,
			}).Stage(formStage)
			formField.FormFieldSelect = formFieldSelect

			formField.FormFieldSelect.Options = make([]*form.Option, 0)
			for bstruct := range *models.GetGongstructInstancesSet[models.Bstruct](stageOfInterest) {
				option := (&form.Option{
					Name: bstruct.Name,
				}).Stage(formStage)

				if bstruct == instanceWithInferedType.Associationtob {
					formFieldSelect.Value = option
				}

				formField.FormFieldSelect.Options =
					append(formField.FormFieldSelect.Options, option)
			}
		}

	}
}
