// generated code - do not edit
package probe

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	form "github.com/fullstack-lang/gongtable/go/models"
	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/gong/test/go/models"
)

type ButtonImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	Icon       gongtree_buttons.ButtonType
	playground *Playground
}

func NewButtonImplGongstruct(
	gongStruct *gong_models.GongStruct,
	icon gongtree_buttons.ButtonType,
	playground *Playground,
) (buttonImplGongstruct *ButtonImplGongstruct) {

	buttonImplGongstruct = new(ButtonImplGongstruct)
	buttonImplGongstruct.Icon = icon
	buttonImplGongstruct.gongStruct = gongStruct
	buttonImplGongstruct.playground = playground

	return
}

func (buttonImpl *ButtonImplGongstruct) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	log.Println("ButtonImplGongstruct: ButtonUpdated")

	formStage := buttonImpl.playground.formStage
	formStage.Reset()
	formStage.Commit()

	switch buttonImpl.gongStruct.Name {
	// insertion point
	case "Astruct":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: "New Astruct Form",
			OnSave: NewAstructFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		astruct := new(models.Astruct)
		FillUpForm(astruct, formGroup, buttonImpl.playground)
	case "AstructBstruct2Use":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewAstructBstruct2UseFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		astructbstruct2use := new(models.AstructBstruct2Use)
		FillUpForm(astructbstruct2use, formGroup, buttonImpl.playground)
	case "AstructBstructUse":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewAstructBstructUseFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		astructbstructuse := new(models.AstructBstructUse)
		FillUpForm(astructbstructuse, formGroup, buttonImpl.playground)
	case "Bstruct":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewBstructFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		bstruct := new(models.Bstruct)
		FillUpForm(bstruct, formGroup, buttonImpl.playground)
	case "Dstruct":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewDstructFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		dstruct := new(models.Dstruct)
		FillUpForm(dstruct, formGroup, buttonImpl.playground)
	}
	formStage.Commit()
}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	playground *Playground,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Astruct:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("Associationtob", instanceWithInferedType.Associationtob, formGroup, playground)
		AssociationSliceToForm("Anotherarrayofb", instanceWithInferedType, &instanceWithInferedType.Anotherarrayofb, formGroup, playground)
		AssociationFieldToForm("Anotherassociationtob_2", instanceWithInferedType.Anotherassociationtob_2, formGroup, playground)
		BasicFieldtoForm("Date", instanceWithInferedType.Date, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Booleanfield", instanceWithInferedType.Booleanfield, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeStringToForm("Aenum", instanceWithInferedType.Aenum, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeStringToForm("Aenum_2", instanceWithInferedType.Aenum_2, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeStringToForm("Benum", instanceWithInferedType.Benum, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeIntToForm("CEnum", instanceWithInferedType.CEnum, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("CName", instanceWithInferedType.CName, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("CFloatfield", instanceWithInferedType.CFloatfield, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("Bstruct", instanceWithInferedType.Bstruct, formGroup, playground)
		AssociationFieldToForm("Bstruct2", instanceWithInferedType.Bstruct2, formGroup, playground)
		AssociationFieldToForm("Dstruct", instanceWithInferedType.Dstruct, formGroup, playground)
		AssociationFieldToForm("Dstruct2", instanceWithInferedType.Dstruct2, formGroup, playground)
		AssociationFieldToForm("Dstruct3", instanceWithInferedType.Dstruct3, formGroup, playground)
		AssociationFieldToForm("Dstruct4", instanceWithInferedType.Dstruct4, formGroup, playground)
		BasicFieldtoForm("Floatfield", instanceWithInferedType.Floatfield, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Intfield", instanceWithInferedType.Intfield, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Anotherbooleanfield", instanceWithInferedType.Anotherbooleanfield, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Duration1", instanceWithInferedType.Duration1, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("Anarrayofb", instanceWithInferedType, &instanceWithInferedType.Anarrayofb, formGroup, playground)
		AssociationSliceToForm("Anarrayofa", instanceWithInferedType, &instanceWithInferedType.Anarrayofa, formGroup, playground)
		AssociationSliceToForm("AnarrayofbUse", instanceWithInferedType, &instanceWithInferedType.AnarrayofbUse, formGroup, playground)
		AssociationSliceToForm("Anarrayofb2Use", instanceWithInferedType, &instanceWithInferedType.Anarrayofb2Use, formGroup, playground)
		AssociationFieldToForm("AnAstruct", instanceWithInferedType.AnAstruct, formGroup, playground)
		BasicFieldtoForm("StructRef", instanceWithInferedType.StructRef, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("FieldRef", instanceWithInferedType.FieldRef, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("EnumIntRef", instanceWithInferedType.EnumIntRef, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("EnumStringRef", instanceWithInferedType.EnumStringRef, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("EnumValue", instanceWithInferedType.EnumValue, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("ConstIdentifierValue", instanceWithInferedType.ConstIdentifierValue, instanceWithInferedType, playground.formStage, formGroup)

	case *models.AstructBstruct2Use:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("Bstrcut2", instanceWithInferedType.Bstrcut2, formGroup, playground)

	case *models.AstructBstructUse:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("Bstruct2", instanceWithInferedType.Bstruct2, formGroup, playground)

	case *models.Bstruct:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Floatfield", instanceWithInferedType.Floatfield, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Floatfield2", instanceWithInferedType.Floatfield2, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Intfield", instanceWithInferedType.Intfield, instanceWithInferedType, playground.formStage, formGroup)

	case *models.Dstruct:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)

	default:
		_ = instanceWithInferedType
	}
}
