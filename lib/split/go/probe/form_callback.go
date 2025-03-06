// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/split/go/models"
	"github.com/fullstack-lang/gong/lib/split/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__AsSplitFormCallback(
	assplit *models.AsSplit,
	probe *Probe,
	formGroup *table.FormGroup,
) (assplitFormCallback *AsSplitFormCallback) {
	assplitFormCallback = new(AsSplitFormCallback)
	assplitFormCallback.probe = probe
	assplitFormCallback.assplit = assplit
	assplitFormCallback.formGroup = formGroup

	assplitFormCallback.CreationMode = (assplit == nil)

	return
}

type AsSplitFormCallback struct {
	assplit *models.AsSplit

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (assplitFormCallback *AsSplitFormCallback) OnSave() {

	log.Println("AsSplitFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	assplitFormCallback.probe.formStage.Checkout()

	if assplitFormCallback.assplit == nil {
		assplitFormCallback.assplit = new(models.AsSplit).Stage(assplitFormCallback.probe.stageOfInterest)
	}
	assplit_ := assplitFormCallback.assplit
	_ = assplit_

	for _, formDiv := range assplitFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(assplit_.Name), formDiv)
		case "Direction":
			FormDivEnumStringFieldToField(&(assplit_.Direction), formDiv)
		case "AsSplitArea:AsSplits":
			// we need to retrieve the field owner before the change
			var pastAsSplitAreaOwner *models.AsSplitArea
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "AsSplitArea"
			rf.Fieldname = "AsSplits"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				assplitFormCallback.probe.stageOfInterest,
				assplitFormCallback.probe.backRepoOfInterest,
				assplit_,
				&rf)

			if reverseFieldOwner != nil {
				pastAsSplitAreaOwner = reverseFieldOwner.(*models.AsSplitArea)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAsSplitAreaOwner != nil {
					idx := slices.Index(pastAsSplitAreaOwner.AsSplits, assplit_)
					pastAsSplitAreaOwner.AsSplits = slices.Delete(pastAsSplitAreaOwner.AsSplits, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _assplitarea := range *models.GetGongstructInstancesSet[models.AsSplitArea](assplitFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _assplitarea.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAsSplitAreaOwner := _assplitarea // we have a match
						if pastAsSplitAreaOwner != nil {
							if newAsSplitAreaOwner != pastAsSplitAreaOwner {
								idx := slices.Index(pastAsSplitAreaOwner.AsSplits, assplit_)
								pastAsSplitAreaOwner.AsSplits = slices.Delete(pastAsSplitAreaOwner.AsSplits, idx, idx+1)
								newAsSplitAreaOwner.AsSplits = append(newAsSplitAreaOwner.AsSplits, assplit_)
							}
						} else {
							newAsSplitAreaOwner.AsSplits = append(newAsSplitAreaOwner.AsSplits, assplit_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if assplitFormCallback.formGroup.HasSuppressButtonBeenPressed {
		assplit_.Unstage(assplitFormCallback.probe.stageOfInterest)
	}

	assplitFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.AsSplit](
		assplitFormCallback.probe,
	)
	assplitFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if assplitFormCallback.CreationMode || assplitFormCallback.formGroup.HasSuppressButtonBeenPressed {
		assplitFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(assplitFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AsSplitFormCallback(
			nil,
			assplitFormCallback.probe,
			newFormGroup,
		)
		assplit := new(models.AsSplit)
		FillUpForm(assplit, newFormGroup, assplitFormCallback.probe)
		assplitFormCallback.probe.formStage.Commit()
	}

	fillUpTree(assplitFormCallback.probe)
}
func __gong__New__AsSplitAreaFormCallback(
	assplitarea *models.AsSplitArea,
	probe *Probe,
	formGroup *table.FormGroup,
) (assplitareaFormCallback *AsSplitAreaFormCallback) {
	assplitareaFormCallback = new(AsSplitAreaFormCallback)
	assplitareaFormCallback.probe = probe
	assplitareaFormCallback.assplitarea = assplitarea
	assplitareaFormCallback.formGroup = formGroup

	assplitareaFormCallback.CreationMode = (assplitarea == nil)

	return
}

type AsSplitAreaFormCallback struct {
	assplitarea *models.AsSplitArea

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (assplitareaFormCallback *AsSplitAreaFormCallback) OnSave() {

	log.Println("AsSplitAreaFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	assplitareaFormCallback.probe.formStage.Checkout()

	if assplitareaFormCallback.assplitarea == nil {
		assplitareaFormCallback.assplitarea = new(models.AsSplitArea).Stage(assplitareaFormCallback.probe.stageOfInterest)
	}
	assplitarea_ := assplitareaFormCallback.assplitarea
	_ = assplitarea_

	for _, formDiv := range assplitareaFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(assplitarea_.Name), formDiv)
		case "Size":
			FormDivBasicFieldToField(&(assplitarea_.Size), formDiv)
		case "IsAny":
			FormDivBasicFieldToField(&(assplitarea_.IsAny), formDiv)
		case "AsSplit:AsSplitAreas":
			// we need to retrieve the field owner before the change
			var pastAsSplitOwner *models.AsSplit
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "AsSplit"
			rf.Fieldname = "AsSplitAreas"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				assplitareaFormCallback.probe.stageOfInterest,
				assplitareaFormCallback.probe.backRepoOfInterest,
				assplitarea_,
				&rf)

			if reverseFieldOwner != nil {
				pastAsSplitOwner = reverseFieldOwner.(*models.AsSplit)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastAsSplitOwner != nil {
					idx := slices.Index(pastAsSplitOwner.AsSplitAreas, assplitarea_)
					pastAsSplitOwner.AsSplitAreas = slices.Delete(pastAsSplitOwner.AsSplitAreas, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _assplit := range *models.GetGongstructInstancesSet[models.AsSplit](assplitareaFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _assplit.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newAsSplitOwner := _assplit // we have a match
						if pastAsSplitOwner != nil {
							if newAsSplitOwner != pastAsSplitOwner {
								idx := slices.Index(pastAsSplitOwner.AsSplitAreas, assplitarea_)
								pastAsSplitOwner.AsSplitAreas = slices.Delete(pastAsSplitOwner.AsSplitAreas, idx, idx+1)
								newAsSplitOwner.AsSplitAreas = append(newAsSplitOwner.AsSplitAreas, assplitarea_)
							}
						} else {
							newAsSplitOwner.AsSplitAreas = append(newAsSplitOwner.AsSplitAreas, assplitarea_)
						}
					}
				}
			}
		case "View:RootAsSplitAreas":
			// we need to retrieve the field owner before the change
			var pastViewOwner *models.View
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "View"
			rf.Fieldname = "RootAsSplitAreas"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				assplitareaFormCallback.probe.stageOfInterest,
				assplitareaFormCallback.probe.backRepoOfInterest,
				assplitarea_,
				&rf)

			if reverseFieldOwner != nil {
				pastViewOwner = reverseFieldOwner.(*models.View)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastViewOwner != nil {
					idx := slices.Index(pastViewOwner.RootAsSplitAreas, assplitarea_)
					pastViewOwner.RootAsSplitAreas = slices.Delete(pastViewOwner.RootAsSplitAreas, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _view := range *models.GetGongstructInstancesSet[models.View](assplitareaFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _view.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newViewOwner := _view // we have a match
						if pastViewOwner != nil {
							if newViewOwner != pastViewOwner {
								idx := slices.Index(pastViewOwner.RootAsSplitAreas, assplitarea_)
								pastViewOwner.RootAsSplitAreas = slices.Delete(pastViewOwner.RootAsSplitAreas, idx, idx+1)
								newViewOwner.RootAsSplitAreas = append(newViewOwner.RootAsSplitAreas, assplitarea_)
							}
						} else {
							newViewOwner.RootAsSplitAreas = append(newViewOwner.RootAsSplitAreas, assplitarea_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if assplitareaFormCallback.formGroup.HasSuppressButtonBeenPressed {
		assplitarea_.Unstage(assplitareaFormCallback.probe.stageOfInterest)
	}

	assplitareaFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.AsSplitArea](
		assplitareaFormCallback.probe,
	)
	assplitareaFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if assplitareaFormCallback.CreationMode || assplitareaFormCallback.formGroup.HasSuppressButtonBeenPressed {
		assplitareaFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(assplitareaFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AsSplitAreaFormCallback(
			nil,
			assplitareaFormCallback.probe,
			newFormGroup,
		)
		assplitarea := new(models.AsSplitArea)
		FillUpForm(assplitarea, newFormGroup, assplitareaFormCallback.probe)
		assplitareaFormCallback.probe.formStage.Commit()
	}

	fillUpTree(assplitareaFormCallback.probe)
}
func __gong__New__ViewFormCallback(
	view *models.View,
	probe *Probe,
	formGroup *table.FormGroup,
) (viewFormCallback *ViewFormCallback) {
	viewFormCallback = new(ViewFormCallback)
	viewFormCallback.probe = probe
	viewFormCallback.view = view
	viewFormCallback.formGroup = formGroup

	viewFormCallback.CreationMode = (view == nil)

	return
}

type ViewFormCallback struct {
	view *models.View

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (viewFormCallback *ViewFormCallback) OnSave() {

	log.Println("ViewFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	viewFormCallback.probe.formStage.Checkout()

	if viewFormCallback.view == nil {
		viewFormCallback.view = new(models.View).Stage(viewFormCallback.probe.stageOfInterest)
	}
	view_ := viewFormCallback.view
	_ = view_

	for _, formDiv := range viewFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(view_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if viewFormCallback.formGroup.HasSuppressButtonBeenPressed {
		view_.Unstage(viewFormCallback.probe.stageOfInterest)
	}

	viewFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.View](
		viewFormCallback.probe,
	)
	viewFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if viewFormCallback.CreationMode || viewFormCallback.formGroup.HasSuppressButtonBeenPressed {
		viewFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(viewFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ViewFormCallback(
			nil,
			viewFormCallback.probe,
			newFormGroup,
		)
		view := new(models.View)
		FillUpForm(view, newFormGroup, viewFormCallback.probe)
		viewFormCallback.probe.formStage.Commit()
	}

	fillUpTree(viewFormCallback.probe)
}
