// generated code - do not edit
package x

import (
	"log"
	"slices"
	"time"

	"github.com/fullstack-lang/gong/test2/go/models/x"
	orm_x "github.com/fullstack-lang/gong/test2/go/orm/x"
	table "github.com/fullstack-lang/gongtable/go/models"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm_x.BackRepoStruct{}

// insertion point
func __gong__New__AFormCallback(
	a *x.A,
	playground *Playground,
) (aFormCallback *AFormCallback) {
	aFormCallback = new(AFormCallback)
	aFormCallback.playground = playground
	aFormCallback.a = a

	aFormCallback.CreationMode = (a == nil)

	return
}

type AFormCallback struct {
	a *x.A

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (aFormCallback *AFormCallback) OnSave() {

	log.Println("AFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	aFormCallback.playground.formStage.Checkout()

	if aFormCallback.a == nil {
		aFormCallback.a = new(x.A).Stage(aFormCallback.playground.stageOfInterest)
	}
	a_ := aFormCallback.a
	_ = a_

	// get the formGroup
	formGroup := aFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_.Name), formDiv)
		}
	}

	aFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[x.A](
		aFormCallback.playground,
	)
	aFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if aFormCallback.CreationMode {
		aFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__AFormCallback(
				nil,
				aFormCallback.playground,
			),
		}).Stage(aFormCallback.playground.formStage)
		a := new(x.A)
		FillUpForm(a, newFormGroup, aFormCallback.playground)
		aFormCallback.playground.formStage.Commit()
	}

	fillUpTree(aFormCallback.playground)
}
