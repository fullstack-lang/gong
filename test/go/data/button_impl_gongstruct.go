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

	switch buttonImpl.gongStruct.Name {
	// insertion point
	case "Astruct":
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
		astruct := new(models.Astruct)
		FillUpForm(astruct, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "AstructBstruct2Use":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewAstructBstruct2UseFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		astructbstruct2use := new(models.AstructBstruct2Use)
		FillUpForm(astructbstruct2use, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "AstructBstructUse":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewAstructBstructUseFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		astructbstructuse := new(models.AstructBstructUse)
		FillUpForm(astructbstructuse, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "Bstruct":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewBstructFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		bstruct := new(models.Bstruct)
		FillUpForm(bstruct, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "Dstruct":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewDstructFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		dstruct := new(models.Dstruct)
		FillUpForm(dstruct, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
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
	// insertion point
	case *models.Astruct:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Booleanfield", instanceWithInferedType.Booleanfield, instanceWithInferedType, formStage, formGroup)
		EnumTypeStringToForm("Aenum", instanceWithInferedType.Aenum, instanceWithInferedType, formStage, formGroup)
		EnumTypeStringToForm("Aenum_2", instanceWithInferedType.Aenum_2, instanceWithInferedType, formStage, formGroup)
		EnumTypeStringToForm("Benum", instanceWithInferedType.Benum, instanceWithInferedType, formStage, formGroup)
		EnumTypeIntToForm("CEnum", instanceWithInferedType.CEnum, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("CName", instanceWithInferedType.CName, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("CFloatfield", instanceWithInferedType.CFloatfield, instanceWithInferedType, formStage, formGroup)
		AssociationFieldToForm("Bstruct", instanceWithInferedType.Bstruct, stageOfInterest, formStage, formGroup)
		AssociationFieldToForm("Bstruct2", instanceWithInferedType.Bstruct2, stageOfInterest, formStage, formGroup)
		AssociationFieldToForm("Dstruct", instanceWithInferedType.Dstruct, stageOfInterest, formStage, formGroup)
		AssociationFieldToForm("Dstruct2", instanceWithInferedType.Dstruct2, stageOfInterest, formStage, formGroup)
		AssociationFieldToForm("Dstruct3", instanceWithInferedType.Dstruct3, stageOfInterest, formStage, formGroup)
		AssociationFieldToForm("Dstruct4", instanceWithInferedType.Dstruct4, stageOfInterest, formStage, formGroup)
		BasicFieldtoForm("Floatfield", instanceWithInferedType.Floatfield, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Intfield", instanceWithInferedType.Intfield, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Anotherbooleanfield", instanceWithInferedType.Anotherbooleanfield, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Duration1", instanceWithInferedType.Duration1, instanceWithInferedType, formStage, formGroup)
		AssociationFieldToForm("Associationtob", instanceWithInferedType.Associationtob, stageOfInterest, formStage, formGroup)
		AssociationFieldToForm("Anotherassociationtob_2", instanceWithInferedType.Anotherassociationtob_2, stageOfInterest, formStage, formGroup)
		AssociationSliceToForm("Anarrayofb", instanceWithInferedType, &instanceWithInferedType.Anarrayofb, stageOfInterest, formStage, formGroup, r)
		AssociationSliceToForm("Anotherarrayofb", instanceWithInferedType, &instanceWithInferedType.Anotherarrayofb, stageOfInterest, formStage, formGroup, r)
		AssociationSliceToForm("Anarrayofa", instanceWithInferedType, &instanceWithInferedType.Anarrayofa, stageOfInterest, formStage, formGroup, r)
		AssociationSliceToForm("AnarrayofbUse", instanceWithInferedType, &instanceWithInferedType.AnarrayofbUse, stageOfInterest, formStage, formGroup, r)
		AssociationSliceToForm("Anarrayofb2Use", instanceWithInferedType, &instanceWithInferedType.Anarrayofb2Use, stageOfInterest, formStage, formGroup, r)
		AssociationFieldToForm("AnAstruct", instanceWithInferedType.AnAstruct, stageOfInterest, formStage, formGroup)
		BasicFieldtoForm("StructRef", instanceWithInferedType.StructRef, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("FieldRef", instanceWithInferedType.FieldRef, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("EnumIntRef", instanceWithInferedType.EnumIntRef, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("EnumStringRef", instanceWithInferedType.EnumStringRef, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("EnumValue", instanceWithInferedType.EnumValue, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("ConstIdentifierValue", instanceWithInferedType.ConstIdentifierValue, instanceWithInferedType, formStage, formGroup)

	case *models.AstructBstruct2Use:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		AssociationFieldToForm("Bstrcut2", instanceWithInferedType.Bstrcut2, stageOfInterest, formStage, formGroup)

	case *models.AstructBstructUse:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		AssociationFieldToForm("Bstruct2", instanceWithInferedType.Bstruct2, stageOfInterest, formStage, formGroup)

	case *models.Bstruct:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Floatfield", instanceWithInferedType.Floatfield, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Floatfield2", instanceWithInferedType.Floatfield2, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Intfield", instanceWithInferedType.Intfield, instanceWithInferedType, formStage, formGroup)

	case *models.Dstruct:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)

	default:
		_ = instanceWithInferedType
	}
}

