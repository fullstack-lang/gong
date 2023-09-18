// generated code - do not edit
package probe

import (
	"log"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/go/models"
)

const __dummmy__time = time.Nanosecond

// insertion point
func NewGongBasicFieldFormCallback(
	gongbasicfield *models.GongBasicField,
	playground *Playground,
) (gongbasicfieldFormCallback *GongBasicFieldFormCallback) {
	gongbasicfieldFormCallback = new(GongBasicFieldFormCallback)
	gongbasicfieldFormCallback.playground = playground
	gongbasicfieldFormCallback.gongbasicfield = gongbasicfield

	gongbasicfieldFormCallback.CreationMode = (gongbasicfield == nil)

	return
}

type GongBasicFieldFormCallback struct {
	gongbasicfield *models.GongBasicField

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (gongbasicfieldFormCallback *GongBasicFieldFormCallback) OnSave() {

	log.Println("GongBasicFieldFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongbasicfieldFormCallback.playground.formStage.Checkout()

	if gongbasicfieldFormCallback.gongbasicfield == nil {
		gongbasicfieldFormCallback.gongbasicfield = new(models.GongBasicField).Stage(gongbasicfieldFormCallback.playground.stageOfInterest)
	}
	gongbasicfield_ := gongbasicfieldFormCallback.gongbasicfield
	_ = gongbasicfield_

	// get the formGroup
	formGroup := gongbasicfieldFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongbasicfield_.Name), formDiv)
		case "BasicKindName":
			FormDivBasicFieldToField(&(gongbasicfield_.BasicKindName), formDiv)
		case "GongEnum":
			FormDivSelectFieldToField(&(gongbasicfield_.GongEnum), gongbasicfieldFormCallback.playground.stageOfInterest, formDiv)
		case "DeclaredType":
			FormDivBasicFieldToField(&(gongbasicfield_.DeclaredType), formDiv)
		case "CompositeStructName":
			FormDivBasicFieldToField(&(gongbasicfield_.CompositeStructName), formDiv)
		case "Index":
			FormDivBasicFieldToField(&(gongbasicfield_.Index), formDiv)
		case "IsDocLink":
			FormDivBasicFieldToField(&(gongbasicfield_.IsDocLink), formDiv)
		}
	}

	gongbasicfieldFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.GongBasicField](
		gongbasicfieldFormCallback.playground,
	)
	gongbasicfieldFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if gongbasicfieldFormCallback.CreationMode {
		gongbasicfieldFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewGongBasicFieldFormCallback(
				nil,
				gongbasicfieldFormCallback.playground,
			),
		}).Stage(gongbasicfieldFormCallback.playground.formStage)
		gongbasicfield := new(models.GongBasicField)
		FillUpForm(gongbasicfield, newFormGroup, gongbasicfieldFormCallback.playground)
		gongbasicfieldFormCallback.playground.formStage.Commit()
	}

}
func NewGongEnumFormCallback(
	gongenum *models.GongEnum,
	playground *Playground,
) (gongenumFormCallback *GongEnumFormCallback) {
	gongenumFormCallback = new(GongEnumFormCallback)
	gongenumFormCallback.playground = playground
	gongenumFormCallback.gongenum = gongenum

	gongenumFormCallback.CreationMode = (gongenum == nil)

	return
}

type GongEnumFormCallback struct {
	gongenum *models.GongEnum

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (gongenumFormCallback *GongEnumFormCallback) OnSave() {

	log.Println("GongEnumFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongenumFormCallback.playground.formStage.Checkout()

	if gongenumFormCallback.gongenum == nil {
		gongenumFormCallback.gongenum = new(models.GongEnum).Stage(gongenumFormCallback.playground.stageOfInterest)
	}
	gongenum_ := gongenumFormCallback.gongenum
	_ = gongenum_

	// get the formGroup
	formGroup := gongenumFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongenum_.Name), formDiv)
		case "Type":
			FormDivEnumIntFieldToField(&(gongenum_.Type), formDiv)
		}
	}

	gongenumFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.GongEnum](
		gongenumFormCallback.playground,
	)
	gongenumFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if gongenumFormCallback.CreationMode {
		gongenumFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewGongEnumFormCallback(
				nil,
				gongenumFormCallback.playground,
			),
		}).Stage(gongenumFormCallback.playground.formStage)
		gongenum := new(models.GongEnum)
		FillUpForm(gongenum, newFormGroup, gongenumFormCallback.playground)
		gongenumFormCallback.playground.formStage.Commit()
	}

}
func NewGongEnumValueFormCallback(
	gongenumvalue *models.GongEnumValue,
	playground *Playground,
) (gongenumvalueFormCallback *GongEnumValueFormCallback) {
	gongenumvalueFormCallback = new(GongEnumValueFormCallback)
	gongenumvalueFormCallback.playground = playground
	gongenumvalueFormCallback.gongenumvalue = gongenumvalue

	gongenumvalueFormCallback.CreationMode = (gongenumvalue == nil)

	return
}

type GongEnumValueFormCallback struct {
	gongenumvalue *models.GongEnumValue

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (gongenumvalueFormCallback *GongEnumValueFormCallback) OnSave() {

	log.Println("GongEnumValueFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongenumvalueFormCallback.playground.formStage.Checkout()

	if gongenumvalueFormCallback.gongenumvalue == nil {
		gongenumvalueFormCallback.gongenumvalue = new(models.GongEnumValue).Stage(gongenumvalueFormCallback.playground.stageOfInterest)
	}
	gongenumvalue_ := gongenumvalueFormCallback.gongenumvalue
	_ = gongenumvalue_

	// get the formGroup
	formGroup := gongenumvalueFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongenumvalue_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(gongenumvalue_.Value), formDiv)
		}
	}

	gongenumvalueFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.GongEnumValue](
		gongenumvalueFormCallback.playground,
	)
	gongenumvalueFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if gongenumvalueFormCallback.CreationMode {
		gongenumvalueFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewGongEnumValueFormCallback(
				nil,
				gongenumvalueFormCallback.playground,
			),
		}).Stage(gongenumvalueFormCallback.playground.formStage)
		gongenumvalue := new(models.GongEnumValue)
		FillUpForm(gongenumvalue, newFormGroup, gongenumvalueFormCallback.playground)
		gongenumvalueFormCallback.playground.formStage.Commit()
	}

}
func NewGongLinkFormCallback(
	gonglink *models.GongLink,
	playground *Playground,
) (gonglinkFormCallback *GongLinkFormCallback) {
	gonglinkFormCallback = new(GongLinkFormCallback)
	gonglinkFormCallback.playground = playground
	gonglinkFormCallback.gonglink = gonglink

	gonglinkFormCallback.CreationMode = (gonglink == nil)

	return
}

type GongLinkFormCallback struct {
	gonglink *models.GongLink

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (gonglinkFormCallback *GongLinkFormCallback) OnSave() {

	log.Println("GongLinkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gonglinkFormCallback.playground.formStage.Checkout()

	if gonglinkFormCallback.gonglink == nil {
		gonglinkFormCallback.gonglink = new(models.GongLink).Stage(gonglinkFormCallback.playground.stageOfInterest)
	}
	gonglink_ := gonglinkFormCallback.gonglink
	_ = gonglink_

	// get the formGroup
	formGroup := gonglinkFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gonglink_.Name), formDiv)
		case "Recv":
			FormDivBasicFieldToField(&(gonglink_.Recv), formDiv)
		case "ImportPath":
			FormDivBasicFieldToField(&(gonglink_.ImportPath), formDiv)
		}
	}

	gonglinkFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.GongLink](
		gonglinkFormCallback.playground,
	)
	gonglinkFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if gonglinkFormCallback.CreationMode {
		gonglinkFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewGongLinkFormCallback(
				nil,
				gonglinkFormCallback.playground,
			),
		}).Stage(gonglinkFormCallback.playground.formStage)
		gonglink := new(models.GongLink)
		FillUpForm(gonglink, newFormGroup, gonglinkFormCallback.playground)
		gonglinkFormCallback.playground.formStage.Commit()
	}

}
func NewGongNoteFormCallback(
	gongnote *models.GongNote,
	playground *Playground,
) (gongnoteFormCallback *GongNoteFormCallback) {
	gongnoteFormCallback = new(GongNoteFormCallback)
	gongnoteFormCallback.playground = playground
	gongnoteFormCallback.gongnote = gongnote

	gongnoteFormCallback.CreationMode = (gongnote == nil)

	return
}

type GongNoteFormCallback struct {
	gongnote *models.GongNote

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (gongnoteFormCallback *GongNoteFormCallback) OnSave() {

	log.Println("GongNoteFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongnoteFormCallback.playground.formStage.Checkout()

	if gongnoteFormCallback.gongnote == nil {
		gongnoteFormCallback.gongnote = new(models.GongNote).Stage(gongnoteFormCallback.playground.stageOfInterest)
	}
	gongnote_ := gongnoteFormCallback.gongnote
	_ = gongnote_

	// get the formGroup
	formGroup := gongnoteFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongnote_.Name), formDiv)
		case "Body":
			FormDivBasicFieldToField(&(gongnote_.Body), formDiv)
		case "BodyHTML":
			FormDivBasicFieldToField(&(gongnote_.BodyHTML), formDiv)
		}
	}

	gongnoteFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.GongNote](
		gongnoteFormCallback.playground,
	)
	gongnoteFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if gongnoteFormCallback.CreationMode {
		gongnoteFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewGongNoteFormCallback(
				nil,
				gongnoteFormCallback.playground,
			),
		}).Stage(gongnoteFormCallback.playground.formStage)
		gongnote := new(models.GongNote)
		FillUpForm(gongnote, newFormGroup, gongnoteFormCallback.playground)
		gongnoteFormCallback.playground.formStage.Commit()
	}

}
func NewGongStructFormCallback(
	gongstruct *models.GongStruct,
	playground *Playground,
) (gongstructFormCallback *GongStructFormCallback) {
	gongstructFormCallback = new(GongStructFormCallback)
	gongstructFormCallback.playground = playground
	gongstructFormCallback.gongstruct = gongstruct

	gongstructFormCallback.CreationMode = (gongstruct == nil)

	return
}

type GongStructFormCallback struct {
	gongstruct *models.GongStruct

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (gongstructFormCallback *GongStructFormCallback) OnSave() {

	log.Println("GongStructFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongstructFormCallback.playground.formStage.Checkout()

	if gongstructFormCallback.gongstruct == nil {
		gongstructFormCallback.gongstruct = new(models.GongStruct).Stage(gongstructFormCallback.playground.stageOfInterest)
	}
	gongstruct_ := gongstructFormCallback.gongstruct
	_ = gongstruct_

	// get the formGroup
	formGroup := gongstructFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongstruct_.Name), formDiv)
		case "HasOnAfterUpdateSignature":
			FormDivBasicFieldToField(&(gongstruct_.HasOnAfterUpdateSignature), formDiv)
		}
	}

	gongstructFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.GongStruct](
		gongstructFormCallback.playground,
	)
	gongstructFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if gongstructFormCallback.CreationMode {
		gongstructFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewGongStructFormCallback(
				nil,
				gongstructFormCallback.playground,
			),
		}).Stage(gongstructFormCallback.playground.formStage)
		gongstruct := new(models.GongStruct)
		FillUpForm(gongstruct, newFormGroup, gongstructFormCallback.playground)
		gongstructFormCallback.playground.formStage.Commit()
	}

}
func NewGongTimeFieldFormCallback(
	gongtimefield *models.GongTimeField,
	playground *Playground,
) (gongtimefieldFormCallback *GongTimeFieldFormCallback) {
	gongtimefieldFormCallback = new(GongTimeFieldFormCallback)
	gongtimefieldFormCallback.playground = playground
	gongtimefieldFormCallback.gongtimefield = gongtimefield

	gongtimefieldFormCallback.CreationMode = (gongtimefield == nil)

	return
}

type GongTimeFieldFormCallback struct {
	gongtimefield *models.GongTimeField

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (gongtimefieldFormCallback *GongTimeFieldFormCallback) OnSave() {

	log.Println("GongTimeFieldFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongtimefieldFormCallback.playground.formStage.Checkout()

	if gongtimefieldFormCallback.gongtimefield == nil {
		gongtimefieldFormCallback.gongtimefield = new(models.GongTimeField).Stage(gongtimefieldFormCallback.playground.stageOfInterest)
	}
	gongtimefield_ := gongtimefieldFormCallback.gongtimefield
	_ = gongtimefield_

	// get the formGroup
	formGroup := gongtimefieldFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongtimefield_.Name), formDiv)
		case "Index":
			FormDivBasicFieldToField(&(gongtimefield_.Index), formDiv)
		case "CompositeStructName":
			FormDivBasicFieldToField(&(gongtimefield_.CompositeStructName), formDiv)
		}
	}

	gongtimefieldFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.GongTimeField](
		gongtimefieldFormCallback.playground,
	)
	gongtimefieldFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if gongtimefieldFormCallback.CreationMode {
		gongtimefieldFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewGongTimeFieldFormCallback(
				nil,
				gongtimefieldFormCallback.playground,
			),
		}).Stage(gongtimefieldFormCallback.playground.formStage)
		gongtimefield := new(models.GongTimeField)
		FillUpForm(gongtimefield, newFormGroup, gongtimefieldFormCallback.playground)
		gongtimefieldFormCallback.playground.formStage.Commit()
	}

}
func NewMetaFormCallback(
	meta *models.Meta,
	playground *Playground,
) (metaFormCallback *MetaFormCallback) {
	metaFormCallback = new(MetaFormCallback)
	metaFormCallback.playground = playground
	metaFormCallback.meta = meta

	metaFormCallback.CreationMode = (meta == nil)

	return
}

type MetaFormCallback struct {
	meta *models.Meta

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (metaFormCallback *MetaFormCallback) OnSave() {

	log.Println("MetaFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	metaFormCallback.playground.formStage.Checkout()

	if metaFormCallback.meta == nil {
		metaFormCallback.meta = new(models.Meta).Stage(metaFormCallback.playground.stageOfInterest)
	}
	meta_ := metaFormCallback.meta
	_ = meta_

	// get the formGroup
	formGroup := metaFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(meta_.Name), formDiv)
		case "Text":
			FormDivBasicFieldToField(&(meta_.Text), formDiv)
		}
	}

	metaFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Meta](
		metaFormCallback.playground,
	)
	metaFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if metaFormCallback.CreationMode {
		metaFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewMetaFormCallback(
				nil,
				metaFormCallback.playground,
			),
		}).Stage(metaFormCallback.playground.formStage)
		meta := new(models.Meta)
		FillUpForm(meta, newFormGroup, metaFormCallback.playground)
		metaFormCallback.playground.formStage.Commit()
	}

}
func NewMetaReferenceFormCallback(
	metareference *models.MetaReference,
	playground *Playground,
) (metareferenceFormCallback *MetaReferenceFormCallback) {
	metareferenceFormCallback = new(MetaReferenceFormCallback)
	metareferenceFormCallback.playground = playground
	metareferenceFormCallback.metareference = metareference

	metareferenceFormCallback.CreationMode = (metareference == nil)

	return
}

type MetaReferenceFormCallback struct {
	metareference *models.MetaReference

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (metareferenceFormCallback *MetaReferenceFormCallback) OnSave() {

	log.Println("MetaReferenceFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	metareferenceFormCallback.playground.formStage.Checkout()

	if metareferenceFormCallback.metareference == nil {
		metareferenceFormCallback.metareference = new(models.MetaReference).Stage(metareferenceFormCallback.playground.stageOfInterest)
	}
	metareference_ := metareferenceFormCallback.metareference
	_ = metareference_

	// get the formGroup
	formGroup := metareferenceFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(metareference_.Name), formDiv)
		}
	}

	metareferenceFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.MetaReference](
		metareferenceFormCallback.playground,
	)
	metareferenceFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if metareferenceFormCallback.CreationMode {
		metareferenceFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewMetaReferenceFormCallback(
				nil,
				metareferenceFormCallback.playground,
			),
		}).Stage(metareferenceFormCallback.playground.formStage)
		metareference := new(models.MetaReference)
		FillUpForm(metareference, newFormGroup, metareferenceFormCallback.playground)
		metareferenceFormCallback.playground.formStage.Commit()
	}

}
func NewModelPkgFormCallback(
	modelpkg *models.ModelPkg,
	playground *Playground,
) (modelpkgFormCallback *ModelPkgFormCallback) {
	modelpkgFormCallback = new(ModelPkgFormCallback)
	modelpkgFormCallback.playground = playground
	modelpkgFormCallback.modelpkg = modelpkg

	modelpkgFormCallback.CreationMode = (modelpkg == nil)

	return
}

type ModelPkgFormCallback struct {
	modelpkg *models.ModelPkg

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (modelpkgFormCallback *ModelPkgFormCallback) OnSave() {

	log.Println("ModelPkgFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	modelpkgFormCallback.playground.formStage.Checkout()

	if modelpkgFormCallback.modelpkg == nil {
		modelpkgFormCallback.modelpkg = new(models.ModelPkg).Stage(modelpkgFormCallback.playground.stageOfInterest)
	}
	modelpkg_ := modelpkgFormCallback.modelpkg
	_ = modelpkg_

	// get the formGroup
	formGroup := modelpkgFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(modelpkg_.Name), formDiv)
		case "PkgPath":
			FormDivBasicFieldToField(&(modelpkg_.PkgPath), formDiv)
		}
	}

	modelpkgFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.ModelPkg](
		modelpkgFormCallback.playground,
	)
	modelpkgFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if modelpkgFormCallback.CreationMode {
		modelpkgFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewModelPkgFormCallback(
				nil,
				modelpkgFormCallback.playground,
			),
		}).Stage(modelpkgFormCallback.playground.formStage)
		modelpkg := new(models.ModelPkg)
		FillUpForm(modelpkg, newFormGroup, modelpkgFormCallback.playground)
		modelpkgFormCallback.playground.formStage.Commit()
	}

}
func NewPointerToGongStructFieldFormCallback(
	pointertogongstructfield *models.PointerToGongStructField,
	playground *Playground,
) (pointertogongstructfieldFormCallback *PointerToGongStructFieldFormCallback) {
	pointertogongstructfieldFormCallback = new(PointerToGongStructFieldFormCallback)
	pointertogongstructfieldFormCallback.playground = playground
	pointertogongstructfieldFormCallback.pointertogongstructfield = pointertogongstructfield

	pointertogongstructfieldFormCallback.CreationMode = (pointertogongstructfield == nil)

	return
}

type PointerToGongStructFieldFormCallback struct {
	pointertogongstructfield *models.PointerToGongStructField

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (pointertogongstructfieldFormCallback *PointerToGongStructFieldFormCallback) OnSave() {

	log.Println("PointerToGongStructFieldFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	pointertogongstructfieldFormCallback.playground.formStage.Checkout()

	if pointertogongstructfieldFormCallback.pointertogongstructfield == nil {
		pointertogongstructfieldFormCallback.pointertogongstructfield = new(models.PointerToGongStructField).Stage(pointertogongstructfieldFormCallback.playground.stageOfInterest)
	}
	pointertogongstructfield_ := pointertogongstructfieldFormCallback.pointertogongstructfield
	_ = pointertogongstructfield_

	// get the formGroup
	formGroup := pointertogongstructfieldFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(pointertogongstructfield_.Name), formDiv)
		case "GongStruct":
			FormDivSelectFieldToField(&(pointertogongstructfield_.GongStruct), pointertogongstructfieldFormCallback.playground.stageOfInterest, formDiv)
		case "Index":
			FormDivBasicFieldToField(&(pointertogongstructfield_.Index), formDiv)
		case "CompositeStructName":
			FormDivBasicFieldToField(&(pointertogongstructfield_.CompositeStructName), formDiv)
		}
	}

	pointertogongstructfieldFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.PointerToGongStructField](
		pointertogongstructfieldFormCallback.playground,
	)
	pointertogongstructfieldFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if pointertogongstructfieldFormCallback.CreationMode {
		pointertogongstructfieldFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewPointerToGongStructFieldFormCallback(
				nil,
				pointertogongstructfieldFormCallback.playground,
			),
		}).Stage(pointertogongstructfieldFormCallback.playground.formStage)
		pointertogongstructfield := new(models.PointerToGongStructField)
		FillUpForm(pointertogongstructfield, newFormGroup, pointertogongstructfieldFormCallback.playground)
		pointertogongstructfieldFormCallback.playground.formStage.Commit()
	}

}
func NewSliceOfPointerToGongStructFieldFormCallback(
	sliceofpointertogongstructfield *models.SliceOfPointerToGongStructField,
	playground *Playground,
) (sliceofpointertogongstructfieldFormCallback *SliceOfPointerToGongStructFieldFormCallback) {
	sliceofpointertogongstructfieldFormCallback = new(SliceOfPointerToGongStructFieldFormCallback)
	sliceofpointertogongstructfieldFormCallback.playground = playground
	sliceofpointertogongstructfieldFormCallback.sliceofpointertogongstructfield = sliceofpointertogongstructfield

	sliceofpointertogongstructfieldFormCallback.CreationMode = (sliceofpointertogongstructfield == nil)

	return
}

type SliceOfPointerToGongStructFieldFormCallback struct {
	sliceofpointertogongstructfield *models.SliceOfPointerToGongStructField

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (sliceofpointertogongstructfieldFormCallback *SliceOfPointerToGongStructFieldFormCallback) OnSave() {

	log.Println("SliceOfPointerToGongStructFieldFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	sliceofpointertogongstructfieldFormCallback.playground.formStage.Checkout()

	if sliceofpointertogongstructfieldFormCallback.sliceofpointertogongstructfield == nil {
		sliceofpointertogongstructfieldFormCallback.sliceofpointertogongstructfield = new(models.SliceOfPointerToGongStructField).Stage(sliceofpointertogongstructfieldFormCallback.playground.stageOfInterest)
	}
	sliceofpointertogongstructfield_ := sliceofpointertogongstructfieldFormCallback.sliceofpointertogongstructfield
	_ = sliceofpointertogongstructfield_

	// get the formGroup
	formGroup := sliceofpointertogongstructfieldFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(sliceofpointertogongstructfield_.Name), formDiv)
		case "GongStruct":
			FormDivSelectFieldToField(&(sliceofpointertogongstructfield_.GongStruct), sliceofpointertogongstructfieldFormCallback.playground.stageOfInterest, formDiv)
		case "Index":
			FormDivBasicFieldToField(&(sliceofpointertogongstructfield_.Index), formDiv)
		case "CompositeStructName":
			FormDivBasicFieldToField(&(sliceofpointertogongstructfield_.CompositeStructName), formDiv)
		}
	}

	sliceofpointertogongstructfieldFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.SliceOfPointerToGongStructField](
		sliceofpointertogongstructfieldFormCallback.playground,
	)
	sliceofpointertogongstructfieldFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if sliceofpointertogongstructfieldFormCallback.CreationMode {
		sliceofpointertogongstructfieldFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewSliceOfPointerToGongStructFieldFormCallback(
				nil,
				sliceofpointertogongstructfieldFormCallback.playground,
			),
		}).Stage(sliceofpointertogongstructfieldFormCallback.playground.formStage)
		sliceofpointertogongstructfield := new(models.SliceOfPointerToGongStructField)
		FillUpForm(sliceofpointertogongstructfield, newFormGroup, sliceofpointertogongstructfieldFormCallback.playground)
		sliceofpointertogongstructfieldFormCallback.playground.formStage.Commit()
	}

}
