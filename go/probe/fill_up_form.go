// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/go/models"
)

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.GongBasicField:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BasicKindName", instanceWithInferedType.BasicKindName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("GongEnum", instanceWithInferedType.GongEnum, formGroup, probe)
		BasicFieldtoForm("DeclaredType", instanceWithInferedType.DeclaredType, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CompositeStructName", instanceWithInferedType.CompositeStructName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Index", instanceWithInferedType.Index, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsDocLink", instanceWithInferedType.IsDocLink, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsTextArea", instanceWithInferedType.IsTextArea, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsBespokeWidth", instanceWithInferedType.IsBespokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BespokeWidth", instanceWithInferedType.BespokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsBespokeHeight", instanceWithInferedType.IsBespokeHeight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BespokeHeight", instanceWithInferedType.BespokeHeight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongStruct"
			rf.Fieldname = "GongBasicFields"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.GongStruct),
					"GongBasicFields",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.GongStruct](
					nil,
					"GongBasicFields",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.GongEnum:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		AssociationSliceToForm("GongEnumValues", instanceWithInferedType, &instanceWithInferedType.GongEnumValues, formGroup, probe)

	case *models.GongEnumValue:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongEnum"
			rf.Fieldname = "GongEnumValues"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.GongEnum),
					"GongEnumValues",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.GongEnum](
					nil,
					"GongEnumValues",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.GongLink:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Recv", instanceWithInferedType.Recv, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ImportPath", instanceWithInferedType.ImportPath, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongNote"
			rf.Fieldname = "Links"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.GongNote),
					"Links",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.GongNote](
					nil,
					"Links",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.GongNote:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Body", instanceWithInferedType.Body, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BodyHTML", instanceWithInferedType.BodyHTML, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Links", instanceWithInferedType, &instanceWithInferedType.Links, formGroup, probe)

	case *models.GongStruct:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("GongBasicFields", instanceWithInferedType, &instanceWithInferedType.GongBasicFields, formGroup, probe)
		AssociationSliceToForm("GongTimeFields", instanceWithInferedType, &instanceWithInferedType.GongTimeFields, formGroup, probe)
		AssociationSliceToForm("PointerToGongStructFields", instanceWithInferedType, &instanceWithInferedType.PointerToGongStructFields, formGroup, probe)
		AssociationSliceToForm("SliceOfPointerToGongStructFields", instanceWithInferedType, &instanceWithInferedType.SliceOfPointerToGongStructFields, formGroup, probe)
		BasicFieldtoForm("HasOnAfterUpdateSignature", instanceWithInferedType.HasOnAfterUpdateSignature, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsIgnoredForFront", instanceWithInferedType.IsIgnoredForFront, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.GongTimeField:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Index", instanceWithInferedType.Index, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CompositeStructName", instanceWithInferedType.CompositeStructName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BespokeTimeFormat", instanceWithInferedType.BespokeTimeFormat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongStruct"
			rf.Fieldname = "GongTimeFields"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.GongStruct),
					"GongTimeFields",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.GongStruct](
					nil,
					"GongTimeFields",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Meta:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("MetaReferences", instanceWithInferedType, &instanceWithInferedType.MetaReferences, formGroup, probe)

	case *models.MetaReference:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Meta"
			rf.Fieldname = "MetaReferences"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Meta),
					"MetaReferences",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Meta](
					nil,
					"MetaReferences",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ModelPkg:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("PkgPath", instanceWithInferedType.PkgPath, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("PathToGoSubDirectory", instanceWithInferedType.PathToGoSubDirectory, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("OrmPkgGenPath", instanceWithInferedType.OrmPkgGenPath, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DbOrmPkgGenPath", instanceWithInferedType.DbOrmPkgGenPath, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DbLiteOrmPkgGenPath", instanceWithInferedType.DbLiteOrmPkgGenPath, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DbPkgGenPath", instanceWithInferedType.DbPkgGenPath, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ControllersPkgGenPath", instanceWithInferedType.ControllersPkgGenPath, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FullstackPkgGenPath", instanceWithInferedType.FullstackPkgGenPath, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StackPkgGenPath", instanceWithInferedType.StackPkgGenPath, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StaticPkgGenPath", instanceWithInferedType.StaticPkgGenPath, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ProbePkgGenPath", instanceWithInferedType.ProbePkgGenPath, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NgWorkspacePath", instanceWithInferedType.NgWorkspacePath, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NgWorkspaceName", instanceWithInferedType.NgWorkspaceName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NgDataLibrarySourceCodeDirectory", instanceWithInferedType.NgDataLibrarySourceCodeDirectory, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NgSpecificLibrarySourceCodeDirectory", instanceWithInferedType.NgSpecificLibrarySourceCodeDirectory, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MaterialLibDatamodelTargetPath", instanceWithInferedType.MaterialLibDatamodelTargetPath, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.PointerToGongStructField:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("GongStruct", instanceWithInferedType.GongStruct, formGroup, probe)
		BasicFieldtoForm("Index", instanceWithInferedType.Index, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CompositeStructName", instanceWithInferedType.CompositeStructName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsType", instanceWithInferedType.IsType, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongStruct"
			rf.Fieldname = "PointerToGongStructFields"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.GongStruct),
					"PointerToGongStructFields",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.GongStruct](
					nil,
					"PointerToGongStructFields",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.SliceOfPointerToGongStructField:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("GongStruct", instanceWithInferedType.GongStruct, formGroup, probe)
		BasicFieldtoForm("Index", instanceWithInferedType.Index, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CompositeStructName", instanceWithInferedType.CompositeStructName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongStruct"
			rf.Fieldname = "SliceOfPointerToGongStructFields"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.GongStruct),
					"SliceOfPointerToGongStructFields",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.GongStruct](
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
