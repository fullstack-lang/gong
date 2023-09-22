// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/test2/go/models"
	"github.com/fullstack-lang/gong/test2/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__DummyFormCallback(
	dummy *models.Dummy,
	playground *Playground,
) (dummyFormCallback *DummyFormCallback) {
	dummyFormCallback = new(DummyFormCallback)
	dummyFormCallback.playground = playground
	dummyFormCallback.dummy = dummy

	dummyFormCallback.CreationMode = (dummy == nil)

	return
}

type DummyFormCallback struct {
	dummy *models.Dummy

	// If the form call is called on the creation of a new instnace
	CreationMode bool

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

	// display a new form by reset the form stage
	if dummyFormCallback.CreationMode {
		dummyFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__DummyFormCallback(
				nil,
				dummyFormCallback.playground,
			),
		}).Stage(dummyFormCallback.playground.formStage)
		dummy := new(models.Dummy)
		FillUpForm(dummy, newFormGroup, dummyFormCallback.playground)
		dummyFormCallback.playground.formStage.Commit()
	}

}
