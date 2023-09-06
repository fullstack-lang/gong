// generated code - do not edit
package data

import (
	"log"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"
	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gong/test/go/models"
	"github.com/fullstack-lang/gong/test/go/orm"
)

const __dummmy__time = time.Nanosecond

// insertion point
	func NewAstructFormCallback(
		stageOfInterest *models.StageStruct,
		tableStage *table.StageStruct,
		formStage *table.StageStruct,
		astruct *models.Astruct,
		r *gin.Engine,
		backRepoOfInterest *orm.BackRepoStruct,
	) (astructFormCallback *AstructFormCallback) {
		astructFormCallback = new(AstructFormCallback)
		astructFormCallback.stageOfInterest = stageOfInterest
		astructFormCallback.tableStage = tableStage
		astructFormCallback.formStage = formStage
		astructFormCallback.astruct = astruct
		astructFormCallback.r = r
		astructFormCallback.backRepoOfInterest = backRepoOfInterest
		return
	}
	
	type AstructFormCallback struct {
		stageOfInterest    *models.StageStruct
		tableStage         *table.StageStruct
		formStage          *table.StageStruct
		astruct            *models.Astruct
		r                  *gin.Engine
		backRepoOfInterest *orm.BackRepoStruct
	}
	
	func (astructFormCallback *AstructFormCallback) OnSave() {
	
		log.Println("AstructFormCallback, OnSave")
	
		// checkout formStage to have the form group on the stage synchronized with the
		// back repo (and front repo)
		astructFormCallback.formStage.Checkout()
	
		if astructFormCallback.astruct == nil {
			astructFormCallback.astruct = new(models.Astruct).Stage(astructFormCallback.stageOfInterest)
		}
		astruct := astructFormCallback.astruct
		_ = astruct
	
		// get the formGroup
		formGroup := astructFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]
	
		for _, formDiv := range formGroup.FormDivs {
			switch formDiv.Name {
			// insertion point per field
			case "Name":
				FormDivBasicFieldToField(&(astruct.Name), formDiv)
			case "Booleanfield":
				FormDivBasicFieldToField(&(astruct.Booleanfield), formDiv)
			case "Aenum":
				FormDivEnumStringFieldToField(&(astruct.Aenum), formDiv)
			case "Aenum_2":
				FormDivEnumStringFieldToField(&(astruct.Aenum_2), formDiv)
			case "Benum":
				FormDivEnumStringFieldToField(&(astruct.Benum), formDiv)
			case "CEnum":
				FormDivEnumIntFieldToField(&(astruct.CEnum), formDiv)
			case "CName":
				FormDivBasicFieldToField(&(astruct.CName), formDiv)
			case "CFloatfield":
				FormDivBasicFieldToField(&(astruct.CFloatfield), formDiv)
			case "Floatfield":
				FormDivBasicFieldToField(&(astruct.Floatfield), formDiv)
			case "Intfield":
				FormDivBasicFieldToField(&(astruct.Intfield), formDiv)
			case "Anotherbooleanfield":
				FormDivBasicFieldToField(&(astruct.Anotherbooleanfield), formDiv)
			case "Duration1":
				FormDivBasicFieldToField(&(astruct.Duration1), formDiv)
			case "StructRef":
				FormDivBasicFieldToField(&(astruct.StructRef), formDiv)
			case "FieldRef":
				FormDivBasicFieldToField(&(astruct.FieldRef), formDiv)
			case "EnumIntRef":
				FormDivBasicFieldToField(&(astruct.EnumIntRef), formDiv)
			case "EnumStringRef":
				FormDivBasicFieldToField(&(astruct.EnumStringRef), formDiv)
			case "EnumValue":
				FormDivBasicFieldToField(&(astruct.EnumValue), formDiv)
			case "ConstIdentifierValue":
				FormDivBasicFieldToField(&(astruct.ConstIdentifierValue), formDiv)
			}
		}
	
		astructFormCallback.stageOfInterest.Commit()
		fillUpTable[models.Astruct](
			astructFormCallback.stageOfInterest,
			astructFormCallback.tableStage,
			astructFormCallback.formStage,
			astructFormCallback.r,
			astructFormCallback.backRepoOfInterest,
		)
		astructFormCallback.tableStage.Commit()
	}
	func NewAstructBstruct2UseFormCallback(
		stageOfInterest *models.StageStruct,
		tableStage *table.StageStruct,
		formStage *table.StageStruct,
		astructbstruct2use *models.AstructBstruct2Use,
		r *gin.Engine,
		backRepoOfInterest *orm.BackRepoStruct,
	) (astructbstruct2useFormCallback *AstructBstruct2UseFormCallback) {
		astructbstruct2useFormCallback = new(AstructBstruct2UseFormCallback)
		astructbstruct2useFormCallback.stageOfInterest = stageOfInterest
		astructbstruct2useFormCallback.tableStage = tableStage
		astructbstruct2useFormCallback.formStage = formStage
		astructbstruct2useFormCallback.astructbstruct2use = astructbstruct2use
		astructbstruct2useFormCallback.r = r
		astructbstruct2useFormCallback.backRepoOfInterest = backRepoOfInterest
		return
	}
	
	type AstructBstruct2UseFormCallback struct {
		stageOfInterest    *models.StageStruct
		tableStage         *table.StageStruct
		formStage          *table.StageStruct
		astructbstruct2use            *models.AstructBstruct2Use
		r                  *gin.Engine
		backRepoOfInterest *orm.BackRepoStruct
	}
	
	func (astructbstruct2useFormCallback *AstructBstruct2UseFormCallback) OnSave() {
	
		log.Println("AstructBstruct2UseFormCallback, OnSave")
	
		// checkout formStage to have the form group on the stage synchronized with the
		// back repo (and front repo)
		astructbstruct2useFormCallback.formStage.Checkout()
	
		if astructbstruct2useFormCallback.astructbstruct2use == nil {
			astructbstruct2useFormCallback.astructbstruct2use = new(models.AstructBstruct2Use).Stage(astructbstruct2useFormCallback.stageOfInterest)
		}
		astructbstruct2use := astructbstruct2useFormCallback.astructbstruct2use
		_ = astructbstruct2use
	
		// get the formGroup
		formGroup := astructbstruct2useFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]
	
		for _, formDiv := range formGroup.FormDivs {
			switch formDiv.Name {
			// insertion point per field
			case "Name":
				FormDivBasicFieldToField(&(astructbstruct2use.Name), formDiv)
			}
		}
	
		astructbstruct2useFormCallback.stageOfInterest.Commit()
		fillUpTable[models.AstructBstruct2Use](
			astructbstruct2useFormCallback.stageOfInterest,
			astructbstruct2useFormCallback.tableStage,
			astructbstruct2useFormCallback.formStage,
			astructbstruct2useFormCallback.r,
			astructbstruct2useFormCallback.backRepoOfInterest,
		)
		astructbstruct2useFormCallback.tableStage.Commit()
	}
	func NewAstructBstructUseFormCallback(
		stageOfInterest *models.StageStruct,
		tableStage *table.StageStruct,
		formStage *table.StageStruct,
		astructbstructuse *models.AstructBstructUse,
		r *gin.Engine,
		backRepoOfInterest *orm.BackRepoStruct,
	) (astructbstructuseFormCallback *AstructBstructUseFormCallback) {
		astructbstructuseFormCallback = new(AstructBstructUseFormCallback)
		astructbstructuseFormCallback.stageOfInterest = stageOfInterest
		astructbstructuseFormCallback.tableStage = tableStage
		astructbstructuseFormCallback.formStage = formStage
		astructbstructuseFormCallback.astructbstructuse = astructbstructuse
		astructbstructuseFormCallback.r = r
		astructbstructuseFormCallback.backRepoOfInterest = backRepoOfInterest
		return
	}
	
	type AstructBstructUseFormCallback struct {
		stageOfInterest    *models.StageStruct
		tableStage         *table.StageStruct
		formStage          *table.StageStruct
		astructbstructuse            *models.AstructBstructUse
		r                  *gin.Engine
		backRepoOfInterest *orm.BackRepoStruct
	}
	
	func (astructbstructuseFormCallback *AstructBstructUseFormCallback) OnSave() {
	
		log.Println("AstructBstructUseFormCallback, OnSave")
	
		// checkout formStage to have the form group on the stage synchronized with the
		// back repo (and front repo)
		astructbstructuseFormCallback.formStage.Checkout()
	
		if astructbstructuseFormCallback.astructbstructuse == nil {
			astructbstructuseFormCallback.astructbstructuse = new(models.AstructBstructUse).Stage(astructbstructuseFormCallback.stageOfInterest)
		}
		astructbstructuse := astructbstructuseFormCallback.astructbstructuse
		_ = astructbstructuse
	
		// get the formGroup
		formGroup := astructbstructuseFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]
	
		for _, formDiv := range formGroup.FormDivs {
			switch formDiv.Name {
			// insertion point per field
			case "Name":
				FormDivBasicFieldToField(&(astructbstructuse.Name), formDiv)
			}
		}
	
		astructbstructuseFormCallback.stageOfInterest.Commit()
		fillUpTable[models.AstructBstructUse](
			astructbstructuseFormCallback.stageOfInterest,
			astructbstructuseFormCallback.tableStage,
			astructbstructuseFormCallback.formStage,
			astructbstructuseFormCallback.r,
			astructbstructuseFormCallback.backRepoOfInterest,
		)
		astructbstructuseFormCallback.tableStage.Commit()
	}
	func NewBstructFormCallback(
		stageOfInterest *models.StageStruct,
		tableStage *table.StageStruct,
		formStage *table.StageStruct,
		bstruct *models.Bstruct,
		r *gin.Engine,
		backRepoOfInterest *orm.BackRepoStruct,
	) (bstructFormCallback *BstructFormCallback) {
		bstructFormCallback = new(BstructFormCallback)
		bstructFormCallback.stageOfInterest = stageOfInterest
		bstructFormCallback.tableStage = tableStage
		bstructFormCallback.formStage = formStage
		bstructFormCallback.bstruct = bstruct
		bstructFormCallback.r = r
		bstructFormCallback.backRepoOfInterest = backRepoOfInterest
		return
	}
	
	type BstructFormCallback struct {
		stageOfInterest    *models.StageStruct
		tableStage         *table.StageStruct
		formStage          *table.StageStruct
		bstruct            *models.Bstruct
		r                  *gin.Engine
		backRepoOfInterest *orm.BackRepoStruct
	}
	
	func (bstructFormCallback *BstructFormCallback) OnSave() {
	
		log.Println("BstructFormCallback, OnSave")
	
		// checkout formStage to have the form group on the stage synchronized with the
		// back repo (and front repo)
		bstructFormCallback.formStage.Checkout()
	
		if bstructFormCallback.bstruct == nil {
			bstructFormCallback.bstruct = new(models.Bstruct).Stage(bstructFormCallback.stageOfInterest)
		}
		bstruct := bstructFormCallback.bstruct
		_ = bstruct
	
		// get the formGroup
		formGroup := bstructFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]
	
		for _, formDiv := range formGroup.FormDivs {
			switch formDiv.Name {
			// insertion point per field
			case "Name":
				FormDivBasicFieldToField(&(bstruct.Name), formDiv)
			case "Floatfield":
				FormDivBasicFieldToField(&(bstruct.Floatfield), formDiv)
			case "Floatfield2":
				FormDivBasicFieldToField(&(bstruct.Floatfield2), formDiv)
			case "Intfield":
				FormDivBasicFieldToField(&(bstruct.Intfield), formDiv)
			}
		}
	
		bstructFormCallback.stageOfInterest.Commit()
		fillUpTable[models.Bstruct](
			bstructFormCallback.stageOfInterest,
			bstructFormCallback.tableStage,
			bstructFormCallback.formStage,
			bstructFormCallback.r,
			bstructFormCallback.backRepoOfInterest,
		)
		bstructFormCallback.tableStage.Commit()
	}
	func NewDstructFormCallback(
		stageOfInterest *models.StageStruct,
		tableStage *table.StageStruct,
		formStage *table.StageStruct,
		dstruct *models.Dstruct,
		r *gin.Engine,
		backRepoOfInterest *orm.BackRepoStruct,
	) (dstructFormCallback *DstructFormCallback) {
		dstructFormCallback = new(DstructFormCallback)
		dstructFormCallback.stageOfInterest = stageOfInterest
		dstructFormCallback.tableStage = tableStage
		dstructFormCallback.formStage = formStage
		dstructFormCallback.dstruct = dstruct
		dstructFormCallback.r = r
		dstructFormCallback.backRepoOfInterest = backRepoOfInterest
		return
	}
	
	type DstructFormCallback struct {
		stageOfInterest    *models.StageStruct
		tableStage         *table.StageStruct
		formStage          *table.StageStruct
		dstruct            *models.Dstruct
		r                  *gin.Engine
		backRepoOfInterest *orm.BackRepoStruct
	}
	
	func (dstructFormCallback *DstructFormCallback) OnSave() {
	
		log.Println("DstructFormCallback, OnSave")
	
		// checkout formStage to have the form group on the stage synchronized with the
		// back repo (and front repo)
		dstructFormCallback.formStage.Checkout()
	
		if dstructFormCallback.dstruct == nil {
			dstructFormCallback.dstruct = new(models.Dstruct).Stage(dstructFormCallback.stageOfInterest)
		}
		dstruct := dstructFormCallback.dstruct
		_ = dstruct
	
		// get the formGroup
		formGroup := dstructFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]
	
		for _, formDiv := range formGroup.FormDivs {
			switch formDiv.Name {
			// insertion point per field
			case "Name":
				FormDivBasicFieldToField(&(dstruct.Name), formDiv)
			}
		}
	
		dstructFormCallback.stageOfInterest.Commit()
		fillUpTable[models.Dstruct](
			dstructFormCallback.stageOfInterest,
			dstructFormCallback.tableStage,
			dstructFormCallback.formStage,
			dstructFormCallback.r,
			dstructFormCallback.backRepoOfInterest,
		)
		dstructFormCallback.tableStage.Commit()
	}
