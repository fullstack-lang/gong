// generated code - do not edit
package probe

import (
	"log"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/test2/go/models"
)

const __dummmy__time = time.Nanosecond

// insertion point
func NewDummyFormCallback(
	dummy *models.Dummy,
	playground *Playground,
) (dummyFormCallback *DummyFormCallback) {
	dummyFormCallback = new(DummyFormCallback)
	dummyFormCallback.playground = playground
	dummyFormCallback.dummy = dummy
	return
}

type DummyFormCallback struct {
	dummy *models.Dummy

	playground *Playground
}

func (dummyFormCallback *DummyFormCallback) OnSave() {

	log.Println("DummyFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	dummyFormCallback.playground.formStage.Checkout()

	if dummyFormCallback.dummy == nil {
		dummyFormCallback.dummy = new(models.Dummy).Stage(dummyFormCallback.playground.stageOfInterest)
	}
	dummy_ := dummyFormCallback.dummy
	_ = dummy_

	// get the formGroup
	formGroup := dummyFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(dummy_.Name), formDiv)
		}
	}

	dummyFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Dummy](
		dummyFormCallback.playground,
	)
	dummyFormCallback.playground.tableStage.Commit()
}
