// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gong/go/orm"
)

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	playground *Playground,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.GongBasicField:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("BasicKindName", instanceWithInferedType.BasicKindName, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("GongEnum", instanceWithInferedType.GongEnum, formGroup, playground)
		BasicFieldtoForm("DeclaredType", instanceWithInferedType.DeclaredType, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("CompositeStructName", instanceWithInferedType.CompositeStructName, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Index", instanceWithInferedType.Index, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("IsDocLink", instanceWithInferedType.IsDocLink, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("IsTextArea", instanceWithInferedType.IsTextArea, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongStruct"
			rf.Fieldname = "GongBasicFields"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.GongStruct),
					"GongBasicFields",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.GongStruct, *models.GongBasicField](
					nil,
					"GongBasicFields",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.GongEnum:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		EnumTypeIntToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("GongEnumValues", instanceWithInferedType, &instanceWithInferedType.GongEnumValues, formGroup, playground)

	case *models.GongEnumValue:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongEnum"
			rf.Fieldname = "GongEnumValues"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.GongEnum),
					"GongEnumValues",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.GongEnum, *models.GongEnumValue](
					nil,
					"GongEnumValues",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.GongLink:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Recv", instanceWithInferedType.Recv, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("ImportPath", instanceWithInferedType.ImportPath, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongNote"
			rf.Fieldname = "Links"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.GongNote),
					"Links",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.GongNote, *models.GongLink](
					nil,
					"Links",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.GongNote:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Body", instanceWithInferedType.Body, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("BodyHTML", instanceWithInferedType.BodyHTML, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("Links", instanceWithInferedType, &instanceWithInferedType.Links, formGroup, playground)

	case *models.GongStruct:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("GongBasicFields", instanceWithInferedType, &instanceWithInferedType.GongBasicFields, formGroup, playground)
		AssociationSliceToForm("GongTimeFields", instanceWithInferedType, &instanceWithInferedType.GongTimeFields, formGroup, playground)
		AssociationSliceToForm("PointerToGongStructFields", instanceWithInferedType, &instanceWithInferedType.PointerToGongStructFields, formGroup, playground)
		AssociationSliceToForm("SliceOfPointerToGongStructFields", instanceWithInferedType, &instanceWithInferedType.SliceOfPointerToGongStructFields, formGroup, playground)
		BasicFieldtoForm("HasOnAfterUpdateSignature", instanceWithInferedType.HasOnAfterUpdateSignature, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.GongTimeField:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Index", instanceWithInferedType.Index, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("CompositeStructName", instanceWithInferedType.CompositeStructName, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongStruct"
			rf.Fieldname = "GongTimeFields"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.GongStruct),
					"GongTimeFields",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.GongStruct, *models.GongTimeField](
					nil,
					"GongTimeFields",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.Meta:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("MetaReferences", instanceWithInferedType, &instanceWithInferedType.MetaReferences, formGroup, playground)

	case *models.MetaReference:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Meta"
			rf.Fieldname = "MetaReferences"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Meta),
					"MetaReferences",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Meta, *models.MetaReference](
					nil,
					"MetaReferences",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.ModelPkg:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("PkgPath", instanceWithInferedType.PkgPath, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.PointerToGongStructField:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("GongStruct", instanceWithInferedType.GongStruct, formGroup, playground)
		BasicFieldtoForm("Index", instanceWithInferedType.Index, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("CompositeStructName", instanceWithInferedType.CompositeStructName, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongStruct"
			rf.Fieldname = "PointerToGongStructFields"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.GongStruct),
					"PointerToGongStructFields",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.GongStruct, *models.PointerToGongStructField](
					nil,
					"PointerToGongStructFields",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.SliceOfPointerToGongStructField:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("GongStruct", instanceWithInferedType.GongStruct, formGroup, playground)
		BasicFieldtoForm("Index", instanceWithInferedType.Index, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("CompositeStructName", instanceWithInferedType.CompositeStructName, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongStruct"
			rf.Fieldname = "SliceOfPointerToGongStructFields"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.GongStruct),
					"SliceOfPointerToGongStructFields",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.GongStruct, *models.SliceOfPointerToGongStructField](
					nil,
					"SliceOfPointerToGongStructFields",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	default:
		_ = instanceWithInferedType
	}
}
