// generated code - do not edit
package data

import (
	"log"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"
	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gong/test2/go/models"
	"github.com/fullstack-lang/gong/test2/go/orm"
)

const __dummmy__time = time.Nanosecond

// insertion point
func NewDummyFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	dummy *models.Dummy,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (dummyFormCallback *DummyFormCallback) {
	dummyFormCallback = new(DummyFormCallback)
	dummyFormCallback.stageOfInterest = stageOfInterest
	dummyFormCallback.tableStage = tableStage
	dummyFormCallback.formStage = formStage
	dummyFormCallback.dummy = dummy
	dummyFormCallback.r = r
	dummyFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type DummyFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	dummy            *models.Dummy
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (dummyFormCallback *DummyFormCallback) OnSave() {

	log.Println("DummyFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	dummyFormCallback.formStage.Checkout()

	if dummyFormCallback.dummy == nil {
		dummyFormCallback.dummy = new(models.Dummy).Stage(dummyFormCallback.stageOfInterest)
	}
	dummy_ := dummyFormCallback.dummy
	_ = dummy_

	// get the formGroup
	formGroup := dummyFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(dummy_.Name), formDiv)
		}
	}

	dummyFormCallback.stageOfInterest.Commit()
	fillUpTable[models.Dummy](
		dummyFormCallback.stageOfInterest,
		dummyFormCallback.tableStage,
		dummyFormCallback.formStage,
		dummyFormCallback.r,
		dummyFormCallback.backRepoOfInterest,
	)
	dummyFormCallback.tableStage.Commit()
}
