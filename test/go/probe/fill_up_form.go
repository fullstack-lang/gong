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
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("Associationtob", instanceWithInferedType.Associationtob, formGroup, playground)
		AssociationSliceToForm("Anarrayofb", instanceWithInferedType, &instanceWithInferedType.Anarrayofb, formGroup, playground)
		AssociationFieldToForm("Anotherassociationtob_2", instanceWithInferedType.Anotherassociationtob_2, formGroup, playground)
		BasicFieldtoForm("Date", instanceWithInferedType.Date, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Booleanfield", instanceWithInferedType.Booleanfield, instanceWithInferedType, playground.formStage, formGroup, false)
		EnumTypeStringToForm("Aenum", instanceWithInferedType.Aenum, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeStringToForm("Aenum_2", instanceWithInferedType.Aenum_2, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeStringToForm("Benum", instanceWithInferedType.Benum, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeIntToForm("CEnum", instanceWithInferedType.CEnum, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("CName", instanceWithInferedType.CName, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("CFloatfield", instanceWithInferedType.CFloatfield, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("Bstruct", instanceWithInferedType.Bstruct, formGroup, playground)
		AssociationFieldToForm("Bstruct2", instanceWithInferedType.Bstruct2, formGroup, playground)
		AssociationFieldToForm("Dstruct", instanceWithInferedType.Dstruct, formGroup, playground)
		AssociationFieldToForm("Dstruct2", instanceWithInferedType.Dstruct2, formGroup, playground)
		AssociationFieldToForm("Dstruct3", instanceWithInferedType.Dstruct3, formGroup, playground)
		AssociationFieldToForm("Dstruct4", instanceWithInferedType.Dstruct4, formGroup, playground)
		BasicFieldtoForm("Floatfield", instanceWithInferedType.Floatfield, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Intfield", instanceWithInferedType.Intfield, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Anotherbooleanfield", instanceWithInferedType.Anotherbooleanfield, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Duration1", instanceWithInferedType.Duration1, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("Anarrayofa", instanceWithInferedType, &instanceWithInferedType.Anarrayofa, formGroup, playground)
		AssociationSliceToForm("Anotherarrayofb", instanceWithInferedType, &instanceWithInferedType.Anotherarrayofb, formGroup, playground)
		AssociationSliceToForm("AnarrayofbUse", instanceWithInferedType, &instanceWithInferedType.AnarrayofbUse, formGroup, playground)
		AssociationSliceToForm("Anarrayofb2Use", instanceWithInferedType, &instanceWithInferedType.Anarrayofb2Use, formGroup, playground)
		AssociationFieldToForm("AnAstruct", instanceWithInferedType.AnAstruct, formGroup, playground)
		BasicFieldtoForm("StructRef", instanceWithInferedType.StructRef, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("FieldRef", instanceWithInferedType.FieldRef, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("EnumIntRef", instanceWithInferedType.EnumIntRef, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("EnumStringRef", instanceWithInferedType.EnumStringRef, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("EnumValue", instanceWithInferedType.EnumValue, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("ConstIdentifierValue", instanceWithInferedType.ConstIdentifierValue, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("TextArea", instanceWithInferedType.TextArea, instanceWithInferedType, playground.formStage, formGroup, true)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Astruct"
			rf.Fieldname = "Anarrayofa"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Astruct),
					"Anarrayofa",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Astruct, *models.Astruct](
					nil,
					"Anarrayofa",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.AstructBstruct2Use:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("Bstrcut2", instanceWithInferedType.Bstrcut2, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Astruct"
			rf.Fieldname = "Anarrayofb2Use"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Astruct),
					"Anarrayofb2Use",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Astruct, *models.AstructBstruct2Use](
					nil,
					"Anarrayofb2Use",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.AstructBstructUse:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("Bstruct2", instanceWithInferedType.Bstruct2, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Astruct"
			rf.Fieldname = "AnarrayofbUse"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Astruct),
					"AnarrayofbUse",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Astruct, *models.AstructBstructUse](
					nil,
					"AnarrayofbUse",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.Bstruct:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Floatfield", instanceWithInferedType.Floatfield, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Floatfield2", instanceWithInferedType.Floatfield2, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Intfield", instanceWithInferedType.Intfield, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Astruct"
			rf.Fieldname = "Anarrayofb"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Astruct),
					"Anarrayofb",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Astruct, *models.Bstruct](
					nil,
					"Anarrayofb",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Astruct"
			rf.Fieldname = "Anotherarrayofb"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Astruct),
					"Anotherarrayofb",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Astruct, *models.Bstruct](
					nil,
					"Anotherarrayofb",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Dstruct"
			rf.Fieldname = "Anarrayofb"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Dstruct),
					"Anarrayofb",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Dstruct, *models.Bstruct](
					nil,
					"Anarrayofb",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.Dstruct:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("Anarrayofb", instanceWithInferedType, &instanceWithInferedType.Anarrayofb, formGroup, playground)

	default:
		_ = instanceWithInferedType
	}
}
