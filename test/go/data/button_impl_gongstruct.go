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
)

type ButtonImplGongstruct struct {
	gongStruct      *gong_models.GongStruct
	Icon            gongtree_buttons.ButtonType
	formStage       *form.StageStruct
	stageOfInterest *models.StageStruct
	r               *gin.Engine
}

func NewButtonImplGongstruct(
	gongStruct *gong_models.GongStruct,
	icon gongtree_buttons.ButtonType,
	formStage *form.StageStruct,
	stageOfInterest *models.StageStruct,
	r *gin.Engine,
) (buttonImplGongstruct *ButtonImplGongstruct) {

	buttonImplGongstruct = new(ButtonImplGongstruct)
	buttonImplGongstruct.Icon = icon
	buttonImplGongstruct.gongStruct = gongStruct
	buttonImplGongstruct.formStage = formStage
	buttonImplGongstruct.stageOfInterest = stageOfInterest
	buttonImplGongstruct.r = r

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

		formDiv := (&form.FormDiv{
			Name: "Anarrayofb",
		}).Stage(formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDiv)

		formEditAssocButton := (&form.FormEditAssocButton{
			Name:  "Anarrayofb",
			Label: "Anarrayofb",
		}).Stage(formStage)
		formDiv.FormEditAssocButton = formEditAssocButton
		onAssocEditon := NewOnAssocEditon(r, formStage)
		formEditAssocButton.OnAssocEditon = onAssocEditon

		formSortAssocButton := (&form.FormSortAssocButton{
			Name:  "Anarrayofb",
			Label: "Anarrayofb",
		}).Stage(formStage)
		formDiv.FormSortAssocButton = formSortAssocButton
		onSortingEditon := NewOnSortingEditon(r, formStage)
		formSortAssocButton.OnSortEdition = onSortingEditon

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
