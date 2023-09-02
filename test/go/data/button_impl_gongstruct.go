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
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Date", instanceWithInferedType.Date, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Booleanfield", instanceWithInferedType.Booleanfield, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Intfield", instanceWithInferedType.Intfield, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Floatfield", instanceWithInferedType.Floatfield, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Duration1", instanceWithInferedType.Duration1, instanceWithInferedType, formStage, formGroup)

		EnumTypeStringToForm("Aenum", instanceWithInferedType.Aenum, instanceWithInferedType, formStage, formGroup)
		EnumTypeStringToForm("Aenum_2", instanceWithInferedType.Aenum_2, instanceWithInferedType, formStage, formGroup)
		EnumTypeIntToForm("Cenum", instanceWithInferedType.CEnum, instanceWithInferedType, formStage, formGroup)

		AssociationFieldToForm[*models.Astruct, *models.Bstruct, models.Bstruct]("Associationtob", instanceWithInferedType.Associationtob, stageOfInterest, instanceWithInferedType, formStage, formGroup)
		AssociationFieldToForm[*models.Astruct, *models.Bstruct, models.Bstruct]("Anotherassociationtob_2", instanceWithInferedType.Anotherassociationtob_2, stageOfInterest, instanceWithInferedType, formStage, formGroup)
		AssociationFieldToForm[*models.Astruct, *models.Astruct, models.Astruct]("AnAstruct", instanceWithInferedType.AnAstruct, stageOfInterest, instanceWithInferedType, formStage, formGroup)
	}
}
