// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/test/go/models"
	"github.com/fullstack-lang/gong/test/go/orm"
)

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
		AssociationSliceToForm("Anarrayofb", instanceWithInferedType, &instanceWithInferedType.Anarrayofb, formGroup, playground)
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
		AssociationSliceToForm("Anarrayofa", instanceWithInferedType, &instanceWithInferedType.Anarrayofa, formGroup, playground)
		AssociationSliceToForm("Anotherarrayofb", instanceWithInferedType, &instanceWithInferedType.Anotherarrayofb, formGroup, playground)
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

		reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, "Anarrayofb")
		AssociationReverseFieldToForm(
			reverseFieldOwner.(*models.Astruct),
			"Anarrayofb",
			instanceWithInferedType,
			(reverseFieldOwner.(*models.Astruct)).Anarrayofb,
			formGroup,
			playground)

	case *models.Dstruct:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)

	default:
		_ = instanceWithInferedType
	}
}
