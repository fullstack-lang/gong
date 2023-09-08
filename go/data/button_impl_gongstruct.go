// generated code - do not edit
package data

import (
	"log"

	"github.com/gin-gonic/gin"

	gong_models "github.com/fullstack-lang/gong/go/models"
	form "github.com/fullstack-lang/gongtable/go/models"
	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gong/go/orm"
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
	case "GongBasicField":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewGongBasicFieldFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		gongbasicfield := new(models.GongBasicField)
		FillUpForm(gongbasicfield, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "GongEnum":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewGongEnumFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		gongenum := new(models.GongEnum)
		FillUpForm(gongenum, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "GongEnumValue":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewGongEnumValueFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		gongenumvalue := new(models.GongEnumValue)
		FillUpForm(gongenumvalue, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "GongLink":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewGongLinkFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		gonglink := new(models.GongLink)
		FillUpForm(gonglink, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "GongNote":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewGongNoteFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		gongnote := new(models.GongNote)
		FillUpForm(gongnote, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "GongStruct":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewGongStructFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		gongstruct := new(models.GongStruct)
		FillUpForm(gongstruct, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "GongTimeField":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewGongTimeFieldFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		gongtimefield := new(models.GongTimeField)
		FillUpForm(gongtimefield, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "Meta":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewMetaFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		meta := new(models.Meta)
		FillUpForm(meta, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "MetaReference":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewMetaReferenceFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		metareference := new(models.MetaReference)
		FillUpForm(metareference, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "ModelPkg":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewModelPkgFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		modelpkg := new(models.ModelPkg)
		FillUpForm(modelpkg, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "PointerToGongStructField":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewPointerToGongStructFieldFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		pointertogongstructfield := new(models.PointerToGongStructField)
		FillUpForm(pointertogongstructfield, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
	case "SliceOfPointerToGongStructField":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: NewSliceOfPointerToGongStructFieldFormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		sliceofpointertogongstructfield := new(models.SliceOfPointerToGongStructField)
		FillUpForm(sliceofpointertogongstructfield, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)
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
	case *models.GongBasicField:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("BasicKindName", instanceWithInferedType.BasicKindName, instanceWithInferedType, formStage, formGroup)
		AssociationFieldToForm("GongEnum", instanceWithInferedType.GongEnum, stageOfInterest, formStage, formGroup)
		BasicFieldtoForm("DeclaredType", instanceWithInferedType.DeclaredType, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("CompositeStructName", instanceWithInferedType.CompositeStructName, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Index", instanceWithInferedType.Index, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("IsDocLink", instanceWithInferedType.IsDocLink, instanceWithInferedType, formStage, formGroup)

	case *models.GongEnum:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		EnumTypeIntToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, formStage, formGroup)
		AssociationSliceToForm("GongEnumValues", instanceWithInferedType, &instanceWithInferedType.GongEnumValues, stageOfInterest, formStage, formGroup, r)

	case *models.GongEnumValue:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, formStage, formGroup)

	case *models.GongLink:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Recv", instanceWithInferedType.Recv, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("ImportPath", instanceWithInferedType.ImportPath, instanceWithInferedType, formStage, formGroup)

	case *models.GongNote:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Body", instanceWithInferedType.Body, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("BodyHTML", instanceWithInferedType.BodyHTML, instanceWithInferedType, formStage, formGroup)
		AssociationSliceToForm("Links", instanceWithInferedType, &instanceWithInferedType.Links, stageOfInterest, formStage, formGroup, r)

	case *models.GongStruct:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		AssociationSliceToForm("GongBasicFields", instanceWithInferedType, &instanceWithInferedType.GongBasicFields, stageOfInterest, formStage, formGroup, r)
		AssociationSliceToForm("GongTimeFields", instanceWithInferedType, &instanceWithInferedType.GongTimeFields, stageOfInterest, formStage, formGroup, r)
		AssociationSliceToForm("PointerToGongStructFields", instanceWithInferedType, &instanceWithInferedType.PointerToGongStructFields, stageOfInterest, formStage, formGroup, r)
		AssociationSliceToForm("SliceOfPointerToGongStructFields", instanceWithInferedType, &instanceWithInferedType.SliceOfPointerToGongStructFields, stageOfInterest, formStage, formGroup, r)
		BasicFieldtoForm("HasOnAfterUpdateSignature", instanceWithInferedType.HasOnAfterUpdateSignature, instanceWithInferedType, formStage, formGroup)

	case *models.GongTimeField:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Index", instanceWithInferedType.Index, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("CompositeStructName", instanceWithInferedType.CompositeStructName, instanceWithInferedType, formStage, formGroup)

	case *models.Meta:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, formStage, formGroup)
		AssociationSliceToForm("MetaReferences", instanceWithInferedType, &instanceWithInferedType.MetaReferences, stageOfInterest, formStage, formGroup, r)

	case *models.MetaReference:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)

	case *models.ModelPkg:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("PkgPath", instanceWithInferedType.PkgPath, instanceWithInferedType, formStage, formGroup)

	case *models.PointerToGongStructField:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		AssociationFieldToForm("GongStruct", instanceWithInferedType.GongStruct, stageOfInterest, formStage, formGroup)
		BasicFieldtoForm("Index", instanceWithInferedType.Index, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("CompositeStructName", instanceWithInferedType.CompositeStructName, instanceWithInferedType, formStage, formGroup)

	case *models.SliceOfPointerToGongStructField:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, formStage, formGroup)
		AssociationFieldToForm("GongStruct", instanceWithInferedType.GongStruct, stageOfInterest, formStage, formGroup)
		BasicFieldtoForm("Index", instanceWithInferedType.Index, instanceWithInferedType, formStage, formGroup)
		BasicFieldtoForm("CompositeStructName", instanceWithInferedType.CompositeStructName, instanceWithInferedType, formStage, formGroup)

	default:
		_ = instanceWithInferedType
	}
}

