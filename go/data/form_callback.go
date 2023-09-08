// generated code - do not edit
package data

import (
	"log"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"
	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gong/go/orm"
)

const __dummmy__time = time.Nanosecond

// insertion point
func NewGongBasicFieldFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	gongbasicfield *models.GongBasicField,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (gongbasicfieldFormCallback *GongBasicFieldFormCallback) {
	gongbasicfieldFormCallback = new(GongBasicFieldFormCallback)
	gongbasicfieldFormCallback.stageOfInterest = stageOfInterest
	gongbasicfieldFormCallback.tableStage = tableStage
	gongbasicfieldFormCallback.formStage = formStage
	gongbasicfieldFormCallback.gongbasicfield = gongbasicfield
	gongbasicfieldFormCallback.r = r
	gongbasicfieldFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type GongBasicFieldFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	gongbasicfield            *models.GongBasicField
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (gongbasicfieldFormCallback *GongBasicFieldFormCallback) OnSave() {

	log.Println("GongBasicFieldFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongbasicfieldFormCallback.formStage.Checkout()

	if gongbasicfieldFormCallback.gongbasicfield == nil {
		gongbasicfieldFormCallback.gongbasicfield = new(models.GongBasicField).Stage(gongbasicfieldFormCallback.stageOfInterest)
	}
	gongbasicfield := gongbasicfieldFormCallback.gongbasicfield
	_ = gongbasicfield

	// get the formGroup
	formGroup := gongbasicfieldFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongbasicfield.Name), formDiv)
		case "BasicKindName":
			FormDivBasicFieldToField(&(gongbasicfield.BasicKindName), formDiv)
		case "GongEnum":
			FormDivSelectFieldToField(&(gongbasicfield.GongEnum), gongbasicfieldFormCallback.stageOfInterest, formDiv)
		case "DeclaredType":
			FormDivBasicFieldToField(&(gongbasicfield.DeclaredType), formDiv)
		case "CompositeStructName":
			FormDivBasicFieldToField(&(gongbasicfield.CompositeStructName), formDiv)
		case "Index":
			FormDivBasicFieldToField(&(gongbasicfield.Index), formDiv)
		case "IsDocLink":
			FormDivBasicFieldToField(&(gongbasicfield.IsDocLink), formDiv)
		}
	}

	gongbasicfieldFormCallback.stageOfInterest.Commit()
	fillUpTable[models.GongBasicField](
		gongbasicfieldFormCallback.stageOfInterest,
		gongbasicfieldFormCallback.tableStage,
		gongbasicfieldFormCallback.formStage,
		gongbasicfieldFormCallback.r,
		gongbasicfieldFormCallback.backRepoOfInterest,
	)
	gongbasicfieldFormCallback.tableStage.Commit()
}
func NewGongEnumFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	gongenum *models.GongEnum,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (gongenumFormCallback *GongEnumFormCallback) {
	gongenumFormCallback = new(GongEnumFormCallback)
	gongenumFormCallback.stageOfInterest = stageOfInterest
	gongenumFormCallback.tableStage = tableStage
	gongenumFormCallback.formStage = formStage
	gongenumFormCallback.gongenum = gongenum
	gongenumFormCallback.r = r
	gongenumFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type GongEnumFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	gongenum            *models.GongEnum
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (gongenumFormCallback *GongEnumFormCallback) OnSave() {

	log.Println("GongEnumFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongenumFormCallback.formStage.Checkout()

	if gongenumFormCallback.gongenum == nil {
		gongenumFormCallback.gongenum = new(models.GongEnum).Stage(gongenumFormCallback.stageOfInterest)
	}
	gongenum := gongenumFormCallback.gongenum
	_ = gongenum

	// get the formGroup
	formGroup := gongenumFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongenum.Name), formDiv)
		case "Type":
			FormDivEnumIntFieldToField(&(gongenum.Type), formDiv)
		}
	}

	gongenumFormCallback.stageOfInterest.Commit()
	fillUpTable[models.GongEnum](
		gongenumFormCallback.stageOfInterest,
		gongenumFormCallback.tableStage,
		gongenumFormCallback.formStage,
		gongenumFormCallback.r,
		gongenumFormCallback.backRepoOfInterest,
	)
	gongenumFormCallback.tableStage.Commit()
}
func NewGongEnumValueFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	gongenumvalue *models.GongEnumValue,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (gongenumvalueFormCallback *GongEnumValueFormCallback) {
	gongenumvalueFormCallback = new(GongEnumValueFormCallback)
	gongenumvalueFormCallback.stageOfInterest = stageOfInterest
	gongenumvalueFormCallback.tableStage = tableStage
	gongenumvalueFormCallback.formStage = formStage
	gongenumvalueFormCallback.gongenumvalue = gongenumvalue
	gongenumvalueFormCallback.r = r
	gongenumvalueFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type GongEnumValueFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	gongenumvalue            *models.GongEnumValue
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (gongenumvalueFormCallback *GongEnumValueFormCallback) OnSave() {

	log.Println("GongEnumValueFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongenumvalueFormCallback.formStage.Checkout()

	if gongenumvalueFormCallback.gongenumvalue == nil {
		gongenumvalueFormCallback.gongenumvalue = new(models.GongEnumValue).Stage(gongenumvalueFormCallback.stageOfInterest)
	}
	gongenumvalue := gongenumvalueFormCallback.gongenumvalue
	_ = gongenumvalue

	// get the formGroup
	formGroup := gongenumvalueFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongenumvalue.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(gongenumvalue.Value), formDiv)
		}
	}

	gongenumvalueFormCallback.stageOfInterest.Commit()
	fillUpTable[models.GongEnumValue](
		gongenumvalueFormCallback.stageOfInterest,
		gongenumvalueFormCallback.tableStage,
		gongenumvalueFormCallback.formStage,
		gongenumvalueFormCallback.r,
		gongenumvalueFormCallback.backRepoOfInterest,
	)
	gongenumvalueFormCallback.tableStage.Commit()
}
func NewGongLinkFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	gonglink *models.GongLink,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (gonglinkFormCallback *GongLinkFormCallback) {
	gonglinkFormCallback = new(GongLinkFormCallback)
	gonglinkFormCallback.stageOfInterest = stageOfInterest
	gonglinkFormCallback.tableStage = tableStage
	gonglinkFormCallback.formStage = formStage
	gonglinkFormCallback.gonglink = gonglink
	gonglinkFormCallback.r = r
	gonglinkFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type GongLinkFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	gonglink            *models.GongLink
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (gonglinkFormCallback *GongLinkFormCallback) OnSave() {

	log.Println("GongLinkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gonglinkFormCallback.formStage.Checkout()

	if gonglinkFormCallback.gonglink == nil {
		gonglinkFormCallback.gonglink = new(models.GongLink).Stage(gonglinkFormCallback.stageOfInterest)
	}
	gonglink := gonglinkFormCallback.gonglink
	_ = gonglink

	// get the formGroup
	formGroup := gonglinkFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gonglink.Name), formDiv)
		case "Recv":
			FormDivBasicFieldToField(&(gonglink.Recv), formDiv)
		case "ImportPath":
			FormDivBasicFieldToField(&(gonglink.ImportPath), formDiv)
		}
	}

	gonglinkFormCallback.stageOfInterest.Commit()
	fillUpTable[models.GongLink](
		gonglinkFormCallback.stageOfInterest,
		gonglinkFormCallback.tableStage,
		gonglinkFormCallback.formStage,
		gonglinkFormCallback.r,
		gonglinkFormCallback.backRepoOfInterest,
	)
	gonglinkFormCallback.tableStage.Commit()
}
func NewGongNoteFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	gongnote *models.GongNote,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (gongnoteFormCallback *GongNoteFormCallback) {
	gongnoteFormCallback = new(GongNoteFormCallback)
	gongnoteFormCallback.stageOfInterest = stageOfInterest
	gongnoteFormCallback.tableStage = tableStage
	gongnoteFormCallback.formStage = formStage
	gongnoteFormCallback.gongnote = gongnote
	gongnoteFormCallback.r = r
	gongnoteFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type GongNoteFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	gongnote            *models.GongNote
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (gongnoteFormCallback *GongNoteFormCallback) OnSave() {

	log.Println("GongNoteFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongnoteFormCallback.formStage.Checkout()

	if gongnoteFormCallback.gongnote == nil {
		gongnoteFormCallback.gongnote = new(models.GongNote).Stage(gongnoteFormCallback.stageOfInterest)
	}
	gongnote := gongnoteFormCallback.gongnote
	_ = gongnote

	// get the formGroup
	formGroup := gongnoteFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongnote.Name), formDiv)
		case "Body":
			FormDivBasicFieldToField(&(gongnote.Body), formDiv)
		case "BodyHTML":
			FormDivBasicFieldToField(&(gongnote.BodyHTML), formDiv)
		}
	}

	gongnoteFormCallback.stageOfInterest.Commit()
	fillUpTable[models.GongNote](
		gongnoteFormCallback.stageOfInterest,
		gongnoteFormCallback.tableStage,
		gongnoteFormCallback.formStage,
		gongnoteFormCallback.r,
		gongnoteFormCallback.backRepoOfInterest,
	)
	gongnoteFormCallback.tableStage.Commit()
}
func NewGongStructFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	gongstruct *models.GongStruct,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (gongstructFormCallback *GongStructFormCallback) {
	gongstructFormCallback = new(GongStructFormCallback)
	gongstructFormCallback.stageOfInterest = stageOfInterest
	gongstructFormCallback.tableStage = tableStage
	gongstructFormCallback.formStage = formStage
	gongstructFormCallback.gongstruct = gongstruct
	gongstructFormCallback.r = r
	gongstructFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type GongStructFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	gongstruct            *models.GongStruct
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (gongstructFormCallback *GongStructFormCallback) OnSave() {

	log.Println("GongStructFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongstructFormCallback.formStage.Checkout()

	if gongstructFormCallback.gongstruct == nil {
		gongstructFormCallback.gongstruct = new(models.GongStruct).Stage(gongstructFormCallback.stageOfInterest)
	}
	gongstruct := gongstructFormCallback.gongstruct
	_ = gongstruct

	// get the formGroup
	formGroup := gongstructFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongstruct.Name), formDiv)
		case "HasOnAfterUpdateSignature":
			FormDivBasicFieldToField(&(gongstruct.HasOnAfterUpdateSignature), formDiv)
		}
	}

	gongstructFormCallback.stageOfInterest.Commit()
	fillUpTable[models.GongStruct](
		gongstructFormCallback.stageOfInterest,
		gongstructFormCallback.tableStage,
		gongstructFormCallback.formStage,
		gongstructFormCallback.r,
		gongstructFormCallback.backRepoOfInterest,
	)
	gongstructFormCallback.tableStage.Commit()
}
func NewGongTimeFieldFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	gongtimefield *models.GongTimeField,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (gongtimefieldFormCallback *GongTimeFieldFormCallback) {
	gongtimefieldFormCallback = new(GongTimeFieldFormCallback)
	gongtimefieldFormCallback.stageOfInterest = stageOfInterest
	gongtimefieldFormCallback.tableStage = tableStage
	gongtimefieldFormCallback.formStage = formStage
	gongtimefieldFormCallback.gongtimefield = gongtimefield
	gongtimefieldFormCallback.r = r
	gongtimefieldFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type GongTimeFieldFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	gongtimefield            *models.GongTimeField
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (gongtimefieldFormCallback *GongTimeFieldFormCallback) OnSave() {

	log.Println("GongTimeFieldFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongtimefieldFormCallback.formStage.Checkout()

	if gongtimefieldFormCallback.gongtimefield == nil {
		gongtimefieldFormCallback.gongtimefield = new(models.GongTimeField).Stage(gongtimefieldFormCallback.stageOfInterest)
	}
	gongtimefield := gongtimefieldFormCallback.gongtimefield
	_ = gongtimefield

	// get the formGroup
	formGroup := gongtimefieldFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongtimefield.Name), formDiv)
		case "Index":
			FormDivBasicFieldToField(&(gongtimefield.Index), formDiv)
		case "CompositeStructName":
			FormDivBasicFieldToField(&(gongtimefield.CompositeStructName), formDiv)
		}
	}

	gongtimefieldFormCallback.stageOfInterest.Commit()
	fillUpTable[models.GongTimeField](
		gongtimefieldFormCallback.stageOfInterest,
		gongtimefieldFormCallback.tableStage,
		gongtimefieldFormCallback.formStage,
		gongtimefieldFormCallback.r,
		gongtimefieldFormCallback.backRepoOfInterest,
	)
	gongtimefieldFormCallback.tableStage.Commit()
}
func NewMetaFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	meta *models.Meta,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (metaFormCallback *MetaFormCallback) {
	metaFormCallback = new(MetaFormCallback)
	metaFormCallback.stageOfInterest = stageOfInterest
	metaFormCallback.tableStage = tableStage
	metaFormCallback.formStage = formStage
	metaFormCallback.meta = meta
	metaFormCallback.r = r
	metaFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type MetaFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	meta            *models.Meta
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (metaFormCallback *MetaFormCallback) OnSave() {

	log.Println("MetaFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	metaFormCallback.formStage.Checkout()

	if metaFormCallback.meta == nil {
		metaFormCallback.meta = new(models.Meta).Stage(metaFormCallback.stageOfInterest)
	}
	meta := metaFormCallback.meta
	_ = meta

	// get the formGroup
	formGroup := metaFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(meta.Name), formDiv)
		case "Text":
			FormDivBasicFieldToField(&(meta.Text), formDiv)
		}
	}

	metaFormCallback.stageOfInterest.Commit()
	fillUpTable[models.Meta](
		metaFormCallback.stageOfInterest,
		metaFormCallback.tableStage,
		metaFormCallback.formStage,
		metaFormCallback.r,
		metaFormCallback.backRepoOfInterest,
	)
	metaFormCallback.tableStage.Commit()
}
func NewMetaReferenceFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	metareference *models.MetaReference,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (metareferenceFormCallback *MetaReferenceFormCallback) {
	metareferenceFormCallback = new(MetaReferenceFormCallback)
	metareferenceFormCallback.stageOfInterest = stageOfInterest
	metareferenceFormCallback.tableStage = tableStage
	metareferenceFormCallback.formStage = formStage
	metareferenceFormCallback.metareference = metareference
	metareferenceFormCallback.r = r
	metareferenceFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type MetaReferenceFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	metareference            *models.MetaReference
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (metareferenceFormCallback *MetaReferenceFormCallback) OnSave() {

	log.Println("MetaReferenceFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	metareferenceFormCallback.formStage.Checkout()

	if metareferenceFormCallback.metareference == nil {
		metareferenceFormCallback.metareference = new(models.MetaReference).Stage(metareferenceFormCallback.stageOfInterest)
	}
	metareference := metareferenceFormCallback.metareference
	_ = metareference

	// get the formGroup
	formGroup := metareferenceFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(metareference.Name), formDiv)
		}
	}

	metareferenceFormCallback.stageOfInterest.Commit()
	fillUpTable[models.MetaReference](
		metareferenceFormCallback.stageOfInterest,
		metareferenceFormCallback.tableStage,
		metareferenceFormCallback.formStage,
		metareferenceFormCallback.r,
		metareferenceFormCallback.backRepoOfInterest,
	)
	metareferenceFormCallback.tableStage.Commit()
}
func NewModelPkgFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	modelpkg *models.ModelPkg,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (modelpkgFormCallback *ModelPkgFormCallback) {
	modelpkgFormCallback = new(ModelPkgFormCallback)
	modelpkgFormCallback.stageOfInterest = stageOfInterest
	modelpkgFormCallback.tableStage = tableStage
	modelpkgFormCallback.formStage = formStage
	modelpkgFormCallback.modelpkg = modelpkg
	modelpkgFormCallback.r = r
	modelpkgFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type ModelPkgFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	modelpkg            *models.ModelPkg
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (modelpkgFormCallback *ModelPkgFormCallback) OnSave() {

	log.Println("ModelPkgFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	modelpkgFormCallback.formStage.Checkout()

	if modelpkgFormCallback.modelpkg == nil {
		modelpkgFormCallback.modelpkg = new(models.ModelPkg).Stage(modelpkgFormCallback.stageOfInterest)
	}
	modelpkg := modelpkgFormCallback.modelpkg
	_ = modelpkg

	// get the formGroup
	formGroup := modelpkgFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(modelpkg.Name), formDiv)
		case "PkgPath":
			FormDivBasicFieldToField(&(modelpkg.PkgPath), formDiv)
		}
	}

	modelpkgFormCallback.stageOfInterest.Commit()
	fillUpTable[models.ModelPkg](
		modelpkgFormCallback.stageOfInterest,
		modelpkgFormCallback.tableStage,
		modelpkgFormCallback.formStage,
		modelpkgFormCallback.r,
		modelpkgFormCallback.backRepoOfInterest,
	)
	modelpkgFormCallback.tableStage.Commit()
}
func NewPointerToGongStructFieldFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	pointertogongstructfield *models.PointerToGongStructField,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (pointertogongstructfieldFormCallback *PointerToGongStructFieldFormCallback) {
	pointertogongstructfieldFormCallback = new(PointerToGongStructFieldFormCallback)
	pointertogongstructfieldFormCallback.stageOfInterest = stageOfInterest
	pointertogongstructfieldFormCallback.tableStage = tableStage
	pointertogongstructfieldFormCallback.formStage = formStage
	pointertogongstructfieldFormCallback.pointertogongstructfield = pointertogongstructfield
	pointertogongstructfieldFormCallback.r = r
	pointertogongstructfieldFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type PointerToGongStructFieldFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	pointertogongstructfield            *models.PointerToGongStructField
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (pointertogongstructfieldFormCallback *PointerToGongStructFieldFormCallback) OnSave() {

	log.Println("PointerToGongStructFieldFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	pointertogongstructfieldFormCallback.formStage.Checkout()

	if pointertogongstructfieldFormCallback.pointertogongstructfield == nil {
		pointertogongstructfieldFormCallback.pointertogongstructfield = new(models.PointerToGongStructField).Stage(pointertogongstructfieldFormCallback.stageOfInterest)
	}
	pointertogongstructfield := pointertogongstructfieldFormCallback.pointertogongstructfield
	_ = pointertogongstructfield

	// get the formGroup
	formGroup := pointertogongstructfieldFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(pointertogongstructfield.Name), formDiv)
		case "GongStruct":
			FormDivSelectFieldToField(&(pointertogongstructfield.GongStruct), pointertogongstructfieldFormCallback.stageOfInterest, formDiv)
		case "Index":
			FormDivBasicFieldToField(&(pointertogongstructfield.Index), formDiv)
		case "CompositeStructName":
			FormDivBasicFieldToField(&(pointertogongstructfield.CompositeStructName), formDiv)
		}
	}

	pointertogongstructfieldFormCallback.stageOfInterest.Commit()
	fillUpTable[models.PointerToGongStructField](
		pointertogongstructfieldFormCallback.stageOfInterest,
		pointertogongstructfieldFormCallback.tableStage,
		pointertogongstructfieldFormCallback.formStage,
		pointertogongstructfieldFormCallback.r,
		pointertogongstructfieldFormCallback.backRepoOfInterest,
	)
	pointertogongstructfieldFormCallback.tableStage.Commit()
}
func NewSliceOfPointerToGongStructFieldFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	sliceofpointertogongstructfield *models.SliceOfPointerToGongStructField,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (sliceofpointertogongstructfieldFormCallback *SliceOfPointerToGongStructFieldFormCallback) {
	sliceofpointertogongstructfieldFormCallback = new(SliceOfPointerToGongStructFieldFormCallback)
	sliceofpointertogongstructfieldFormCallback.stageOfInterest = stageOfInterest
	sliceofpointertogongstructfieldFormCallback.tableStage = tableStage
	sliceofpointertogongstructfieldFormCallback.formStage = formStage
	sliceofpointertogongstructfieldFormCallback.sliceofpointertogongstructfield = sliceofpointertogongstructfield
	sliceofpointertogongstructfieldFormCallback.r = r
	sliceofpointertogongstructfieldFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type SliceOfPointerToGongStructFieldFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	sliceofpointertogongstructfield            *models.SliceOfPointerToGongStructField
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (sliceofpointertogongstructfieldFormCallback *SliceOfPointerToGongStructFieldFormCallback) OnSave() {

	log.Println("SliceOfPointerToGongStructFieldFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	sliceofpointertogongstructfieldFormCallback.formStage.Checkout()

	if sliceofpointertogongstructfieldFormCallback.sliceofpointertogongstructfield == nil {
		sliceofpointertogongstructfieldFormCallback.sliceofpointertogongstructfield = new(models.SliceOfPointerToGongStructField).Stage(sliceofpointertogongstructfieldFormCallback.stageOfInterest)
	}
	sliceofpointertogongstructfield := sliceofpointertogongstructfieldFormCallback.sliceofpointertogongstructfield
	_ = sliceofpointertogongstructfield

	// get the formGroup
	formGroup := sliceofpointertogongstructfieldFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(sliceofpointertogongstructfield.Name), formDiv)
		case "GongStruct":
			FormDivSelectFieldToField(&(sliceofpointertogongstructfield.GongStruct), sliceofpointertogongstructfieldFormCallback.stageOfInterest, formDiv)
		case "Index":
			FormDivBasicFieldToField(&(sliceofpointertogongstructfield.Index), formDiv)
		case "CompositeStructName":
			FormDivBasicFieldToField(&(sliceofpointertogongstructfield.CompositeStructName), formDiv)
		}
	}

	sliceofpointertogongstructfieldFormCallback.stageOfInterest.Commit()
	fillUpTable[models.SliceOfPointerToGongStructField](
		sliceofpointertogongstructfieldFormCallback.stageOfInterest,
		sliceofpointertogongstructfieldFormCallback.tableStage,
		sliceofpointertogongstructfieldFormCallback.formStage,
		sliceofpointertogongstructfieldFormCallback.r,
		sliceofpointertogongstructfieldFormCallback.backRepoOfInterest,
	)
	sliceofpointertogongstructfieldFormCallback.tableStage.Commit()
}
