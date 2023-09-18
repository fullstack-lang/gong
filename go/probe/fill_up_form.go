// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/go/models"
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
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("BasicKindName", instanceWithInferedType.BasicKindName, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("GongEnum", instanceWithInferedType.GongEnum, formGroup, playground)
		BasicFieldtoForm("DeclaredType", instanceWithInferedType.DeclaredType, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("CompositeStructName", instanceWithInferedType.CompositeStructName, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Index", instanceWithInferedType.Index, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("IsDocLink", instanceWithInferedType.IsDocLink, instanceWithInferedType, playground.formStage, formGroup)

	case *models.GongEnum:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeIntToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("GongEnumValues", instanceWithInferedType, &instanceWithInferedType.GongEnumValues, formGroup, playground)

	case *models.GongEnumValue:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, playground.formStage, formGroup)

	case *models.GongLink:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Recv", instanceWithInferedType.Recv, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("ImportPath", instanceWithInferedType.ImportPath, instanceWithInferedType, playground.formStage, formGroup)

	case *models.GongNote:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Body", instanceWithInferedType.Body, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("BodyHTML", instanceWithInferedType.BodyHTML, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("Links", instanceWithInferedType, &instanceWithInferedType.Links, formGroup, playground)

	case *models.GongStruct:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("GongBasicFields", instanceWithInferedType, &instanceWithInferedType.GongBasicFields, formGroup, playground)
		AssociationSliceToForm("GongTimeFields", instanceWithInferedType, &instanceWithInferedType.GongTimeFields, formGroup, playground)
		AssociationSliceToForm("PointerToGongStructFields", instanceWithInferedType, &instanceWithInferedType.PointerToGongStructFields, formGroup, playground)
		AssociationSliceToForm("SliceOfPointerToGongStructFields", instanceWithInferedType, &instanceWithInferedType.SliceOfPointerToGongStructFields, formGroup, playground)
		BasicFieldtoForm("HasOnAfterUpdateSignature", instanceWithInferedType.HasOnAfterUpdateSignature, instanceWithInferedType, playground.formStage, formGroup)

	case *models.GongTimeField:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Index", instanceWithInferedType.Index, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("CompositeStructName", instanceWithInferedType.CompositeStructName, instanceWithInferedType, playground.formStage, formGroup)

	case *models.Meta:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("MetaReferences", instanceWithInferedType, &instanceWithInferedType.MetaReferences, formGroup, playground)

	case *models.MetaReference:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)

	case *models.ModelPkg:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("PkgPath", instanceWithInferedType.PkgPath, instanceWithInferedType, playground.formStage, formGroup)

	case *models.PointerToGongStructField:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("GongStruct", instanceWithInferedType.GongStruct, formGroup, playground)
		BasicFieldtoForm("Index", instanceWithInferedType.Index, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("CompositeStructName", instanceWithInferedType.CompositeStructName, instanceWithInferedType, playground.formStage, formGroup)

	case *models.SliceOfPointerToGongStructField:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("GongStruct", instanceWithInferedType.GongStruct, formGroup, playground)
		BasicFieldtoForm("Index", instanceWithInferedType.Index, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("CompositeStructName", instanceWithInferedType.CompositeStructName, instanceWithInferedType, playground.formStage, formGroup)

	default:
		_ = instanceWithInferedType
	}
}
