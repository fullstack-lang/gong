// generated code - do not edit
package data

import (
	"log"

	"github.com/gin-gonic/gin"

	gong_models "github.com/fullstack-lang/gong/go/models"
	form "github.com/fullstack-lang/gongtable/go/models"
	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/gong/test/go/models"
	"github.com/fullstack-lang/gong/test/go/orm"
)

type ButtonImplGongstruct struct {
	gongStruct         *gong_models.GongStruct
	Icon               gongtree_buttons.ButtonType
	tableStage         *form.StageStruct
	formStage          *form.StageStruct
	stageOfInterest    *models.StageStruct
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func NewButtonImplGongstruct(
	gongStruct *gong_models.GongStruct,
	icon gongtree_buttons.ButtonType,
	tableStage *form.StageStruct,
	formStage *form.StageStruct,
	stageOfInterest *models.StageStruct,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (buttonImplGongstruct *ButtonImplGongstruct) {

	buttonImplGongstruct = new(ButtonImplGongstruct)
	buttonImplGongstruct.Icon = icon
	buttonImplGongstruct.gongStruct = gongStruct
	buttonImplGongstruct.tableStage = tableStage
	buttonImplGongstruct.formStage = formStage
	buttonImplGongstruct.stageOfInterest = stageOfInterest
	buttonImplGongstruct.r = r
	buttonImplGongstruct.backRepoOfInterest = backRepoOfInterest

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
		Name: form.FormGroupDefaultName.ToString(),
		OnSave: NewAstructFormCallback(
			buttonImpl.stageOfInterest,
			buttonImpl.tableStage,
			formStage,
			nil,
			buttonImpl.r,
			buttonImpl.backRepoOfInterest,
		),
	}).Stage(formStage)

	switch buttonImpl.gongStruct.Name {
	case "Astruct":
		astruct := new(models.Astruct)
		FillUpForm(astruct, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	}
	formStage.Commit()
}

func FillUpForm[T models.Gongstruct](
	instance *T,
	stageOfInterest *models.StageStruct,
	formStage *form.StageStruct,
	formGroup *form.FormGroup,
	r *gin.Engine,
) {

	switch instanceWithInferedType := any(instance).(type) {
	case *models.Astruct:
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)

		AssociationSliceToForm("Anarrayofb", &instanceWithInferedType.Anarrayofb, stageOfInterest, formStage, formGroup, r)

		BasicFieldtoForm("Date", instanceWithInferedType.Date, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Booleanfield", instanceWithInferedType.Booleanfield, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Intfield", instanceWithInferedType.Intfield, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Floatfield", instanceWithInferedType.Floatfield, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Duration1", instanceWithInferedType.Duration1, instanceWithInferedType, formStage, formGroup)

		EnumTypeStringToForm("Aenum", instanceWithInferedType.Aenum, instanceWithInferedType, formStage, formGroup)
		EnumTypeStringToForm("Aenum_2", instanceWithInferedType.Aenum_2, instanceWithInferedType, formStage, formGroup)
		EnumTypeIntToForm("Cenum", instanceWithInferedType.CEnum, instanceWithInferedType, formStage, formGroup)

		AssociationFieldToForm("Associationtob", instanceWithInferedType.Associationtob, stageOfInterest, formStage, formGroup)
		AssociationFieldToForm("Anotherassociationtob_2", instanceWithInferedType.Anotherassociationtob_2, stageOfInterest, formStage, formGroup)
		AssociationFieldToForm("AnAstruct", instanceWithInferedType.AnAstruct, stageOfInterest, formStage, formGroup)

	}
}
