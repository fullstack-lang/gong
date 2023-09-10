// generated code - do not edit
package probe

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	form "github.com/fullstack-lang/gongtable/go/models"
	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/gong/go/models"
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
	case "GongBasicField":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewGongBasicFieldFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		gongbasicfield := new(models.GongBasicField)
		FillUpForm(gongbasicfield, formGroup, buttonImpl.playground)
	case "GongEnum":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewGongEnumFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		gongenum := new(models.GongEnum)
		FillUpForm(gongenum, formGroup, buttonImpl.playground)
	case "GongEnumValue":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewGongEnumValueFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		gongenumvalue := new(models.GongEnumValue)
		FillUpForm(gongenumvalue, formGroup, buttonImpl.playground)
	case "GongLink":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewGongLinkFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		gonglink := new(models.GongLink)
		FillUpForm(gonglink, formGroup, buttonImpl.playground)
	case "GongNote":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewGongNoteFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		gongnote := new(models.GongNote)
		FillUpForm(gongnote, formGroup, buttonImpl.playground)
	case "GongStruct":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewGongStructFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		gongstruct := new(models.GongStruct)
		FillUpForm(gongstruct, formGroup, buttonImpl.playground)
	case "GongTimeField":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewGongTimeFieldFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		gongtimefield := new(models.GongTimeField)
		FillUpForm(gongtimefield, formGroup, buttonImpl.playground)
	case "Meta":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewMetaFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		meta := new(models.Meta)
		FillUpForm(meta, formGroup, buttonImpl.playground)
	case "MetaReference":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewMetaReferenceFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		metareference := new(models.MetaReference)
		FillUpForm(metareference, formGroup, buttonImpl.playground)
	case "ModelPkg":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewModelPkgFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		modelpkg := new(models.ModelPkg)
		FillUpForm(modelpkg, formGroup, buttonImpl.playground)
	case "PointerToGongStructField":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewPointerToGongStructFieldFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		pointertogongstructfield := new(models.PointerToGongStructField)
		FillUpForm(pointertogongstructfield, formGroup, buttonImpl.playground)
	case "SliceOfPointerToGongStructField":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewSliceOfPointerToGongStructFieldFormCallback(
				nil,
				buttonImpl.playground,
			),
		}).Stage(formStage)
		sliceofpointertogongstructfield := new(models.SliceOfPointerToGongStructField)
		FillUpForm(sliceofpointertogongstructfield, formGroup, buttonImpl.playground)
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

