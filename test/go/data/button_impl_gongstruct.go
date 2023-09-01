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
		FillUpFormDivBasicField("Name", instancesTyped.Name, instancesTyped, formStage, formGroup)
		FillUpFormDivBasicField("Date", instancesTyped.Date, instancesTyped, formStage, formGroup)
		FillUpFormDivBasicField("Booleanfield", instancesTyped.Booleanfield, instancesTyped, formStage, formGroup)
		FillUpFormDivBasicField("Intfield", instancesTyped.Intfield, instancesTyped, formStage, formGroup)
		FillUpFormDivBasicField("Floatfield", instancesTyped.Floatfield, instancesTyped, formStage, formGroup)
		FillUpFormDivBasicField("Duration1", instancesTyped.Duration1, instancesTyped, formStage, formGroup)

		FillUpFormDivEnumStringType("Aenum", instancesTyped.Aenum, instancesTyped, formStage, formGroup)
		FillUpFormDivEnumStringType("Aenum_2", instancesTyped.Aenum_2, instancesTyped, formStage, formGroup)
		FillUpFormDivEnumIntType("Cenum", instancesTyped.CEnum, instancesTyped, formStage, formGroup)

	}
}
