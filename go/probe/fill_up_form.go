// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gong/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.GongBasicField:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("BasicKindName", instanceWithInferedType.BasicKindName, instanceWithInferedType, probe.formStage, formGroup, false)
		AssociationFieldToForm("GongEnum", instanceWithInferedType.GongEnum, formGroup, probe)
		BasicFieldtoForm("DeclaredType", instanceWithInferedType.DeclaredType, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("CompositeStructName", instanceWithInferedType.CompositeStructName, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Index", instanceWithInferedType.Index, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("IsDocLink", instanceWithInferedType.IsDocLink, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("IsTextArea", instanceWithInferedType.IsTextArea, instanceWithInferedType, probe.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongStruct"
			rf.Fieldname = "GongBasicFields"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.GongStruct),
					"GongBasicFields",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.GongStruct, *models.GongBasicField](
					nil,
					"GongBasicFields",
					instanceWithInferedType,
					formGroup,
					probe)
			}	
		}

	case *models.GongEnum:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		EnumTypeIntToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		AssociationSliceToForm("GongEnumValues", instanceWithInferedType, &instanceWithInferedType.GongEnumValues, formGroup, probe)

	case *models.GongEnumValue:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongEnum"
			rf.Fieldname = "GongEnumValues"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.GongEnum),
					"GongEnumValues",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.GongEnum, *models.GongEnumValue](
					nil,
					"GongEnumValues",
					instanceWithInferedType,
					formGroup,
					probe)
			}	
		}

	case *models.GongLink:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Recv", instanceWithInferedType.Recv, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("ImportPath", instanceWithInferedType.ImportPath, instanceWithInferedType, probe.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongNote"
			rf.Fieldname = "Links"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.GongNote),
					"Links",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.GongNote, *models.GongLink](
					nil,
					"Links",
					instanceWithInferedType,
					formGroup,
					probe)
			}	
		}

	case *models.GongNote:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Body", instanceWithInferedType.Body, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("BodyHTML", instanceWithInferedType.BodyHTML, instanceWithInferedType, probe.formStage, formGroup, false)
		AssociationSliceToForm("Links", instanceWithInferedType, &instanceWithInferedType.Links, formGroup, probe)

	case *models.GongStruct:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		AssociationSliceToForm("GongBasicFields", instanceWithInferedType, &instanceWithInferedType.GongBasicFields, formGroup, probe)
		AssociationSliceToForm("GongTimeFields", instanceWithInferedType, &instanceWithInferedType.GongTimeFields, formGroup, probe)
		AssociationSliceToForm("PointerToGongStructFields", instanceWithInferedType, &instanceWithInferedType.PointerToGongStructFields, formGroup, probe)
		AssociationSliceToForm("SliceOfPointerToGongStructFields", instanceWithInferedType, &instanceWithInferedType.SliceOfPointerToGongStructFields, formGroup, probe)
		BasicFieldtoForm("HasOnAfterUpdateSignature", instanceWithInferedType.HasOnAfterUpdateSignature, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("IsIgnoredForFront", instanceWithInferedType.IsIgnoredForFront, instanceWithInferedType, probe.formStage, formGroup, false)

	case *models.GongTimeField:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Index", instanceWithInferedType.Index, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("CompositeStructName", instanceWithInferedType.CompositeStructName, instanceWithInferedType, probe.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongStruct"
			rf.Fieldname = "GongTimeFields"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.GongStruct),
					"GongTimeFields",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.GongStruct, *models.GongTimeField](
					nil,
					"GongTimeFields",
					instanceWithInferedType,
					formGroup,
					probe)
			}	
		}

	case *models.Meta:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup, false)
		AssociationSliceToForm("MetaReferences", instanceWithInferedType, &instanceWithInferedType.MetaReferences, formGroup, probe)

	case *models.MetaReference:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Meta"
			rf.Fieldname = "MetaReferences"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Meta),
					"MetaReferences",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Meta, *models.MetaReference](
					nil,
					"MetaReferences",
					instanceWithInferedType,
					formGroup,
					probe)
			}	
		}

	case *models.ModelPkg:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("PkgPath", instanceWithInferedType.PkgPath, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("PathToGoSubDirectory", instanceWithInferedType.PathToGoSubDirectory, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("OrmPkgGenPath", instanceWithInferedType.OrmPkgGenPath, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("ControllersPkgGenPath", instanceWithInferedType.ControllersPkgGenPath, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("FullstackPkgGenPath", instanceWithInferedType.FullstackPkgGenPath, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("StackPkgGenPath", instanceWithInferedType.StackPkgGenPath, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("StaticPkgGenPath", instanceWithInferedType.StaticPkgGenPath, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("ProbePkgGenPath", instanceWithInferedType.ProbePkgGenPath, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("NgWorkspacePath", instanceWithInferedType.NgWorkspacePath, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("NgDataLibrarySourceCodeDirectory", instanceWithInferedType.NgDataLibrarySourceCodeDirectory, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("NgSpecificLibrarySourceCodeDirectory", instanceWithInferedType.NgSpecificLibrarySourceCodeDirectory, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("MaterialLibDatamodelTargetPath", instanceWithInferedType.MaterialLibDatamodelTargetPath, instanceWithInferedType, probe.formStage, formGroup, false)

	case *models.PointerToGongStructField:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		AssociationFieldToForm("GongStruct", instanceWithInferedType.GongStruct, formGroup, probe)
		BasicFieldtoForm("Index", instanceWithInferedType.Index, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("CompositeStructName", instanceWithInferedType.CompositeStructName, instanceWithInferedType, probe.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongStruct"
			rf.Fieldname = "PointerToGongStructFields"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.GongStruct),
					"PointerToGongStructFields",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.GongStruct, *models.PointerToGongStructField](
					nil,
					"PointerToGongStructFields",
					instanceWithInferedType,
					formGroup,
					probe)
			}	
		}

	case *models.SliceOfPointerToGongStructField:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		AssociationFieldToForm("GongStruct", instanceWithInferedType.GongStruct, formGroup, probe)
		BasicFieldtoForm("Index", instanceWithInferedType.Index, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("CompositeStructName", instanceWithInferedType.CompositeStructName, instanceWithInferedType, probe.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongStruct"
			rf.Fieldname = "SliceOfPointerToGongStructFields"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.GongStruct),
					"SliceOfPointerToGongStructFields",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.GongStruct, *models.SliceOfPointerToGongStructField](
					nil,
					"SliceOfPointerToGongStructFields",
					instanceWithInferedType,
					formGroup,
					probe)
			}	
		}

	default:
		_ = instanceWithInferedType
	}
}
