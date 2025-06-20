// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/split/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

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

	// log.Println("AsSplitFormCallback, OnSave")

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
		case "AsSplitAreas":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.AsSplitArea](assplitFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.AsSplitArea, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.AsSplitArea)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					assplitFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			assplit_.AsSplitAreas = instanceSlice

		}
	}

	// manage the suppress operation
	if assplitFormCallback.formGroup.HasSuppressButtonBeenPressed {
		assplit_.Unstage(assplitFormCallback.probe.stageOfInterest)
	}

	assplitFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.AsSplit](
		assplitFormCallback.probe,
	)
	assplitFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if assplitFormCallback.CreationMode || assplitFormCallback.formGroup.HasSuppressButtonBeenPressed {
		assplitFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
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

	updateAndCommitTree(assplitFormCallback.probe)
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

	// log.Println("AsSplitAreaFormCallback, OnSave")

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
		case "ShowNameInHeader":
			FormDivBasicFieldToField(&(assplitarea_.ShowNameInHeader), formDiv)
		case "Size":
			FormDivBasicFieldToField(&(assplitarea_.Size), formDiv)
		case "IsAny":
			FormDivBasicFieldToField(&(assplitarea_.IsAny), formDiv)
		case "AsSplit":
			FormDivSelectFieldToField(&(assplitarea_.AsSplit), assplitareaFormCallback.probe.stageOfInterest, formDiv)
		case "Button":
			FormDivSelectFieldToField(&(assplitarea_.Button), assplitareaFormCallback.probe.stageOfInterest, formDiv)
		case "Cursor":
			FormDivSelectFieldToField(&(assplitarea_.Cursor), assplitareaFormCallback.probe.stageOfInterest, formDiv)
		case "Doc":
			FormDivSelectFieldToField(&(assplitarea_.Doc), assplitareaFormCallback.probe.stageOfInterest, formDiv)
		case "Form":
			FormDivSelectFieldToField(&(assplitarea_.Form), assplitareaFormCallback.probe.stageOfInterest, formDiv)
		case "Load":
			FormDivSelectFieldToField(&(assplitarea_.Load), assplitareaFormCallback.probe.stageOfInterest, formDiv)
		case "Slider":
			FormDivSelectFieldToField(&(assplitarea_.Slider), assplitareaFormCallback.probe.stageOfInterest, formDiv)
		case "Split":
			FormDivSelectFieldToField(&(assplitarea_.Split), assplitareaFormCallback.probe.stageOfInterest, formDiv)
		case "Svg":
			FormDivSelectFieldToField(&(assplitarea_.Svg), assplitareaFormCallback.probe.stageOfInterest, formDiv)
		case "Table":
			FormDivSelectFieldToField(&(assplitarea_.Table), assplitareaFormCallback.probe.stageOfInterest, formDiv)
		case "Tone":
			FormDivSelectFieldToField(&(assplitarea_.Tone), assplitareaFormCallback.probe.stageOfInterest, formDiv)
		case "Tree":
			FormDivSelectFieldToField(&(assplitarea_.Tree), assplitareaFormCallback.probe.stageOfInterest, formDiv)
		case "Xlsx":
			FormDivSelectFieldToField(&(assplitarea_.Xlsx), assplitareaFormCallback.probe.stageOfInterest, formDiv)
		case "HasDiv":
			FormDivBasicFieldToField(&(assplitarea_.HasDiv), formDiv)
		case "DivStyle":
			FormDivBasicFieldToField(&(assplitarea_.DivStyle), formDiv)
		case "AsSplit:AsSplitAreas":
			// WARNING : this form deals with the N-N association "AsSplit.AsSplitAreas []*AsSplitArea" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of AsSplitArea). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.AsSplit
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "AsSplit"
				rf.Fieldname = "AsSplitAreas"
				formerAssociationSource := models.GetReverseFieldOwner(
					assplitareaFormCallback.probe.stageOfInterest,
					assplitarea_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.AsSplit)
					if !ok {
						log.Fatalln("Source of AsSplit.AsSplitAreas []*AsSplitArea, is not an AsSplit instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.AsSplitAreas, assplitarea_)
					formerSource.AsSplitAreas = slices.Delete(formerSource.AsSplitAreas, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.AsSplit
			for _assplit := range *models.GetGongstructInstancesSet[models.AsSplit](assplitareaFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _assplit.GetName() == newSourceName.GetName() {
					newSource = _assplit // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of AsSplit.AsSplitAreas []*AsSplitArea, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.AsSplitAreas = append(newSource.AsSplitAreas, assplitarea_)
		case "View:RootAsSplitAreas":
			// WARNING : this form deals with the N-N association "View.RootAsSplitAreas []*AsSplitArea" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of AsSplitArea). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.View
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "View"
				rf.Fieldname = "RootAsSplitAreas"
				formerAssociationSource := models.GetReverseFieldOwner(
					assplitareaFormCallback.probe.stageOfInterest,
					assplitarea_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.View)
					if !ok {
						log.Fatalln("Source of View.RootAsSplitAreas []*AsSplitArea, is not an View instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.RootAsSplitAreas, assplitarea_)
					formerSource.RootAsSplitAreas = slices.Delete(formerSource.RootAsSplitAreas, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.View
			for _view := range *models.GetGongstructInstancesSet[models.View](assplitareaFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _view.GetName() == newSourceName.GetName() {
					newSource = _view // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of View.RootAsSplitAreas []*AsSplitArea, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.RootAsSplitAreas = append(newSource.RootAsSplitAreas, assplitarea_)
		}
	}

	// manage the suppress operation
	if assplitareaFormCallback.formGroup.HasSuppressButtonBeenPressed {
		assplitarea_.Unstage(assplitareaFormCallback.probe.stageOfInterest)
	}

	assplitareaFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.AsSplitArea](
		assplitareaFormCallback.probe,
	)
	assplitareaFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if assplitareaFormCallback.CreationMode || assplitareaFormCallback.formGroup.HasSuppressButtonBeenPressed {
		assplitareaFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
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

	updateAndCommitTree(assplitareaFormCallback.probe)
}
func __gong__New__ButtonFormCallback(
	button *models.Button,
	probe *Probe,
	formGroup *table.FormGroup,
) (buttonFormCallback *ButtonFormCallback) {
	buttonFormCallback = new(ButtonFormCallback)
	buttonFormCallback.probe = probe
	buttonFormCallback.button = button
	buttonFormCallback.formGroup = formGroup

	buttonFormCallback.CreationMode = (button == nil)

	return
}

type ButtonFormCallback struct {
	button *models.Button

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (buttonFormCallback *ButtonFormCallback) OnSave() {

	// log.Println("ButtonFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	buttonFormCallback.probe.formStage.Checkout()

	if buttonFormCallback.button == nil {
		buttonFormCallback.button = new(models.Button).Stage(buttonFormCallback.probe.stageOfInterest)
	}
	button_ := buttonFormCallback.button
	_ = button_

	for _, formDiv := range buttonFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(button_.Name), formDiv)
		case "StackName":
			FormDivBasicFieldToField(&(button_.StackName), formDiv)
		}
	}

	// manage the suppress operation
	if buttonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		button_.Unstage(buttonFormCallback.probe.stageOfInterest)
	}

	buttonFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Button](
		buttonFormCallback.probe,
	)
	buttonFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if buttonFormCallback.CreationMode || buttonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		buttonFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(buttonFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ButtonFormCallback(
			nil,
			buttonFormCallback.probe,
			newFormGroup,
		)
		button := new(models.Button)
		FillUpForm(button, newFormGroup, buttonFormCallback.probe)
		buttonFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(buttonFormCallback.probe)
}
func __gong__New__CursorFormCallback(
	cursor *models.Cursor,
	probe *Probe,
	formGroup *table.FormGroup,
) (cursorFormCallback *CursorFormCallback) {
	cursorFormCallback = new(CursorFormCallback)
	cursorFormCallback.probe = probe
	cursorFormCallback.cursor = cursor
	cursorFormCallback.formGroup = formGroup

	cursorFormCallback.CreationMode = (cursor == nil)

	return
}

type CursorFormCallback struct {
	cursor *models.Cursor

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (cursorFormCallback *CursorFormCallback) OnSave() {

	// log.Println("CursorFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	cursorFormCallback.probe.formStage.Checkout()

	if cursorFormCallback.cursor == nil {
		cursorFormCallback.cursor = new(models.Cursor).Stage(cursorFormCallback.probe.stageOfInterest)
	}
	cursor_ := cursorFormCallback.cursor
	_ = cursor_

	for _, formDiv := range cursorFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(cursor_.Name), formDiv)
		case "StackName":
			FormDivBasicFieldToField(&(cursor_.StackName), formDiv)
		case "Style":
			FormDivBasicFieldToField(&(cursor_.Style), formDiv)
		}
	}

	// manage the suppress operation
	if cursorFormCallback.formGroup.HasSuppressButtonBeenPressed {
		cursor_.Unstage(cursorFormCallback.probe.stageOfInterest)
	}

	cursorFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Cursor](
		cursorFormCallback.probe,
	)
	cursorFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if cursorFormCallback.CreationMode || cursorFormCallback.formGroup.HasSuppressButtonBeenPressed {
		cursorFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(cursorFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CursorFormCallback(
			nil,
			cursorFormCallback.probe,
			newFormGroup,
		)
		cursor := new(models.Cursor)
		FillUpForm(cursor, newFormGroup, cursorFormCallback.probe)
		cursorFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(cursorFormCallback.probe)
}
func __gong__New__DocFormCallback(
	doc *models.Doc,
	probe *Probe,
	formGroup *table.FormGroup,
) (docFormCallback *DocFormCallback) {
	docFormCallback = new(DocFormCallback)
	docFormCallback.probe = probe
	docFormCallback.doc = doc
	docFormCallback.formGroup = formGroup

	docFormCallback.CreationMode = (doc == nil)

	return
}

type DocFormCallback struct {
	doc *models.Doc

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (docFormCallback *DocFormCallback) OnSave() {

	// log.Println("DocFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	docFormCallback.probe.formStage.Checkout()

	if docFormCallback.doc == nil {
		docFormCallback.doc = new(models.Doc).Stage(docFormCallback.probe.stageOfInterest)
	}
	doc_ := docFormCallback.doc
	_ = doc_

	for _, formDiv := range docFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(doc_.Name), formDiv)
		case "StackName":
			FormDivBasicFieldToField(&(doc_.StackName), formDiv)
		}
	}

	// manage the suppress operation
	if docFormCallback.formGroup.HasSuppressButtonBeenPressed {
		doc_.Unstage(docFormCallback.probe.stageOfInterest)
	}

	docFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Doc](
		docFormCallback.probe,
	)
	docFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if docFormCallback.CreationMode || docFormCallback.formGroup.HasSuppressButtonBeenPressed {
		docFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(docFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DocFormCallback(
			nil,
			docFormCallback.probe,
			newFormGroup,
		)
		doc := new(models.Doc)
		FillUpForm(doc, newFormGroup, docFormCallback.probe)
		docFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(docFormCallback.probe)
}
func __gong__New__FavIconFormCallback(
	favicon *models.FavIcon,
	probe *Probe,
	formGroup *table.FormGroup,
) (faviconFormCallback *FavIconFormCallback) {
	faviconFormCallback = new(FavIconFormCallback)
	faviconFormCallback.probe = probe
	faviconFormCallback.favicon = favicon
	faviconFormCallback.formGroup = formGroup

	faviconFormCallback.CreationMode = (favicon == nil)

	return
}

type FavIconFormCallback struct {
	favicon *models.FavIcon

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (faviconFormCallback *FavIconFormCallback) OnSave() {

	// log.Println("FavIconFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	faviconFormCallback.probe.formStage.Checkout()

	if faviconFormCallback.favicon == nil {
		faviconFormCallback.favicon = new(models.FavIcon).Stage(faviconFormCallback.probe.stageOfInterest)
	}
	favicon_ := faviconFormCallback.favicon
	_ = favicon_

	for _, formDiv := range faviconFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(favicon_.Name), formDiv)
		case "SVG":
			FormDivBasicFieldToField(&(favicon_.SVG), formDiv)
		}
	}

	// manage the suppress operation
	if faviconFormCallback.formGroup.HasSuppressButtonBeenPressed {
		favicon_.Unstage(faviconFormCallback.probe.stageOfInterest)
	}

	faviconFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.FavIcon](
		faviconFormCallback.probe,
	)
	faviconFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if faviconFormCallback.CreationMode || faviconFormCallback.formGroup.HasSuppressButtonBeenPressed {
		faviconFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(faviconFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FavIconFormCallback(
			nil,
			faviconFormCallback.probe,
			newFormGroup,
		)
		favicon := new(models.FavIcon)
		FillUpForm(favicon, newFormGroup, faviconFormCallback.probe)
		faviconFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(faviconFormCallback.probe)
}
func __gong__New__FormFormCallback(
	form *models.Form,
	probe *Probe,
	formGroup *table.FormGroup,
) (formFormCallback *FormFormCallback) {
	formFormCallback = new(FormFormCallback)
	formFormCallback.probe = probe
	formFormCallback.form = form
	formFormCallback.formGroup = formGroup

	formFormCallback.CreationMode = (form == nil)

	return
}

type FormFormCallback struct {
	form *models.Form

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (formFormCallback *FormFormCallback) OnSave() {

	// log.Println("FormFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formFormCallback.probe.formStage.Checkout()

	if formFormCallback.form == nil {
		formFormCallback.form = new(models.Form).Stage(formFormCallback.probe.stageOfInterest)
	}
	form_ := formFormCallback.form
	_ = form_

	for _, formDiv := range formFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(form_.Name), formDiv)
		case "StackName":
			FormDivBasicFieldToField(&(form_.StackName), formDiv)
		case "FormName":
			FormDivBasicFieldToField(&(form_.FormName), formDiv)
		}
	}

	// manage the suppress operation
	if formFormCallback.formGroup.HasSuppressButtonBeenPressed {
		form_.Unstage(formFormCallback.probe.stageOfInterest)
	}

	formFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Form](
		formFormCallback.probe,
	)
	formFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if formFormCallback.CreationMode || formFormCallback.formGroup.HasSuppressButtonBeenPressed {
		formFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(formFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FormFormCallback(
			nil,
			formFormCallback.probe,
			newFormGroup,
		)
		form := new(models.Form)
		FillUpForm(form, newFormGroup, formFormCallback.probe)
		formFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(formFormCallback.probe)
}
func __gong__New__LoadFormCallback(
	load *models.Load,
	probe *Probe,
	formGroup *table.FormGroup,
) (loadFormCallback *LoadFormCallback) {
	loadFormCallback = new(LoadFormCallback)
	loadFormCallback.probe = probe
	loadFormCallback.load = load
	loadFormCallback.formGroup = formGroup

	loadFormCallback.CreationMode = (load == nil)

	return
}

type LoadFormCallback struct {
	load *models.Load

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (loadFormCallback *LoadFormCallback) OnSave() {

	// log.Println("LoadFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	loadFormCallback.probe.formStage.Checkout()

	if loadFormCallback.load == nil {
		loadFormCallback.load = new(models.Load).Stage(loadFormCallback.probe.stageOfInterest)
	}
	load_ := loadFormCallback.load
	_ = load_

	for _, formDiv := range loadFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(load_.Name), formDiv)
		case "StackName":
			FormDivBasicFieldToField(&(load_.StackName), formDiv)
		}
	}

	// manage the suppress operation
	if loadFormCallback.formGroup.HasSuppressButtonBeenPressed {
		load_.Unstage(loadFormCallback.probe.stageOfInterest)
	}

	loadFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Load](
		loadFormCallback.probe,
	)
	loadFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if loadFormCallback.CreationMode || loadFormCallback.formGroup.HasSuppressButtonBeenPressed {
		loadFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(loadFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LoadFormCallback(
			nil,
			loadFormCallback.probe,
			newFormGroup,
		)
		load := new(models.Load)
		FillUpForm(load, newFormGroup, loadFormCallback.probe)
		loadFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(loadFormCallback.probe)
}
func __gong__New__LogoOnTheLeftFormCallback(
	logoontheleft *models.LogoOnTheLeft,
	probe *Probe,
	formGroup *table.FormGroup,
) (logoontheleftFormCallback *LogoOnTheLeftFormCallback) {
	logoontheleftFormCallback = new(LogoOnTheLeftFormCallback)
	logoontheleftFormCallback.probe = probe
	logoontheleftFormCallback.logoontheleft = logoontheleft
	logoontheleftFormCallback.formGroup = formGroup

	logoontheleftFormCallback.CreationMode = (logoontheleft == nil)

	return
}

type LogoOnTheLeftFormCallback struct {
	logoontheleft *models.LogoOnTheLeft

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (logoontheleftFormCallback *LogoOnTheLeftFormCallback) OnSave() {

	// log.Println("LogoOnTheLeftFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	logoontheleftFormCallback.probe.formStage.Checkout()

	if logoontheleftFormCallback.logoontheleft == nil {
		logoontheleftFormCallback.logoontheleft = new(models.LogoOnTheLeft).Stage(logoontheleftFormCallback.probe.stageOfInterest)
	}
	logoontheleft_ := logoontheleftFormCallback.logoontheleft
	_ = logoontheleft_

	for _, formDiv := range logoontheleftFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(logoontheleft_.Name), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(logoontheleft_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(logoontheleft_.Height), formDiv)
		case "SVG":
			FormDivBasicFieldToField(&(logoontheleft_.SVG), formDiv)
		}
	}

	// manage the suppress operation
	if logoontheleftFormCallback.formGroup.HasSuppressButtonBeenPressed {
		logoontheleft_.Unstage(logoontheleftFormCallback.probe.stageOfInterest)
	}

	logoontheleftFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.LogoOnTheLeft](
		logoontheleftFormCallback.probe,
	)
	logoontheleftFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if logoontheleftFormCallback.CreationMode || logoontheleftFormCallback.formGroup.HasSuppressButtonBeenPressed {
		logoontheleftFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(logoontheleftFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LogoOnTheLeftFormCallback(
			nil,
			logoontheleftFormCallback.probe,
			newFormGroup,
		)
		logoontheleft := new(models.LogoOnTheLeft)
		FillUpForm(logoontheleft, newFormGroup, logoontheleftFormCallback.probe)
		logoontheleftFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(logoontheleftFormCallback.probe)
}
func __gong__New__LogoOnTheRightFormCallback(
	logoontheright *models.LogoOnTheRight,
	probe *Probe,
	formGroup *table.FormGroup,
) (logoontherightFormCallback *LogoOnTheRightFormCallback) {
	logoontherightFormCallback = new(LogoOnTheRightFormCallback)
	logoontherightFormCallback.probe = probe
	logoontherightFormCallback.logoontheright = logoontheright
	logoontherightFormCallback.formGroup = formGroup

	logoontherightFormCallback.CreationMode = (logoontheright == nil)

	return
}

type LogoOnTheRightFormCallback struct {
	logoontheright *models.LogoOnTheRight

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (logoontherightFormCallback *LogoOnTheRightFormCallback) OnSave() {

	// log.Println("LogoOnTheRightFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	logoontherightFormCallback.probe.formStage.Checkout()

	if logoontherightFormCallback.logoontheright == nil {
		logoontherightFormCallback.logoontheright = new(models.LogoOnTheRight).Stage(logoontherightFormCallback.probe.stageOfInterest)
	}
	logoontheright_ := logoontherightFormCallback.logoontheright
	_ = logoontheright_

	for _, formDiv := range logoontherightFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(logoontheright_.Name), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(logoontheright_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(logoontheright_.Height), formDiv)
		case "SVG":
			FormDivBasicFieldToField(&(logoontheright_.SVG), formDiv)
		}
	}

	// manage the suppress operation
	if logoontherightFormCallback.formGroup.HasSuppressButtonBeenPressed {
		logoontheright_.Unstage(logoontherightFormCallback.probe.stageOfInterest)
	}

	logoontherightFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.LogoOnTheRight](
		logoontherightFormCallback.probe,
	)
	logoontherightFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if logoontherightFormCallback.CreationMode || logoontherightFormCallback.formGroup.HasSuppressButtonBeenPressed {
		logoontherightFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(logoontherightFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LogoOnTheRightFormCallback(
			nil,
			logoontherightFormCallback.probe,
			newFormGroup,
		)
		logoontheright := new(models.LogoOnTheRight)
		FillUpForm(logoontheright, newFormGroup, logoontherightFormCallback.probe)
		logoontherightFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(logoontherightFormCallback.probe)
}
func __gong__New__SliderFormCallback(
	slider *models.Slider,
	probe *Probe,
	formGroup *table.FormGroup,
) (sliderFormCallback *SliderFormCallback) {
	sliderFormCallback = new(SliderFormCallback)
	sliderFormCallback.probe = probe
	sliderFormCallback.slider = slider
	sliderFormCallback.formGroup = formGroup

	sliderFormCallback.CreationMode = (slider == nil)

	return
}

type SliderFormCallback struct {
	slider *models.Slider

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (sliderFormCallback *SliderFormCallback) OnSave() {

	// log.Println("SliderFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	sliderFormCallback.probe.formStage.Checkout()

	if sliderFormCallback.slider == nil {
		sliderFormCallback.slider = new(models.Slider).Stage(sliderFormCallback.probe.stageOfInterest)
	}
	slider_ := sliderFormCallback.slider
	_ = slider_

	for _, formDiv := range sliderFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(slider_.Name), formDiv)
		case "StackName":
			FormDivBasicFieldToField(&(slider_.StackName), formDiv)
		}
	}

	// manage the suppress operation
	if sliderFormCallback.formGroup.HasSuppressButtonBeenPressed {
		slider_.Unstage(sliderFormCallback.probe.stageOfInterest)
	}

	sliderFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Slider](
		sliderFormCallback.probe,
	)
	sliderFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if sliderFormCallback.CreationMode || sliderFormCallback.formGroup.HasSuppressButtonBeenPressed {
		sliderFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(sliderFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SliderFormCallback(
			nil,
			sliderFormCallback.probe,
			newFormGroup,
		)
		slider := new(models.Slider)
		FillUpForm(slider, newFormGroup, sliderFormCallback.probe)
		sliderFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(sliderFormCallback.probe)
}
func __gong__New__SplitFormCallback(
	split *models.Split,
	probe *Probe,
	formGroup *table.FormGroup,
) (splitFormCallback *SplitFormCallback) {
	splitFormCallback = new(SplitFormCallback)
	splitFormCallback.probe = probe
	splitFormCallback.split = split
	splitFormCallback.formGroup = formGroup

	splitFormCallback.CreationMode = (split == nil)

	return
}

type SplitFormCallback struct {
	split *models.Split

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (splitFormCallback *SplitFormCallback) OnSave() {

	// log.Println("SplitFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	splitFormCallback.probe.formStage.Checkout()

	if splitFormCallback.split == nil {
		splitFormCallback.split = new(models.Split).Stage(splitFormCallback.probe.stageOfInterest)
	}
	split_ := splitFormCallback.split
	_ = split_

	for _, formDiv := range splitFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(split_.Name), formDiv)
		case "StackName":
			FormDivBasicFieldToField(&(split_.StackName), formDiv)
		}
	}

	// manage the suppress operation
	if splitFormCallback.formGroup.HasSuppressButtonBeenPressed {
		split_.Unstage(splitFormCallback.probe.stageOfInterest)
	}

	splitFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Split](
		splitFormCallback.probe,
	)
	splitFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if splitFormCallback.CreationMode || splitFormCallback.formGroup.HasSuppressButtonBeenPressed {
		splitFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(splitFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SplitFormCallback(
			nil,
			splitFormCallback.probe,
			newFormGroup,
		)
		split := new(models.Split)
		FillUpForm(split, newFormGroup, splitFormCallback.probe)
		splitFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(splitFormCallback.probe)
}
func __gong__New__SvgFormCallback(
	svg *models.Svg,
	probe *Probe,
	formGroup *table.FormGroup,
) (svgFormCallback *SvgFormCallback) {
	svgFormCallback = new(SvgFormCallback)
	svgFormCallback.probe = probe
	svgFormCallback.svg = svg
	svgFormCallback.formGroup = formGroup

	svgFormCallback.CreationMode = (svg == nil)

	return
}

type SvgFormCallback struct {
	svg *models.Svg

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (svgFormCallback *SvgFormCallback) OnSave() {

	// log.Println("SvgFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	svgFormCallback.probe.formStage.Checkout()

	if svgFormCallback.svg == nil {
		svgFormCallback.svg = new(models.Svg).Stage(svgFormCallback.probe.stageOfInterest)
	}
	svg_ := svgFormCallback.svg
	_ = svg_

	for _, formDiv := range svgFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(svg_.Name), formDiv)
		case "StackName":
			FormDivBasicFieldToField(&(svg_.StackName), formDiv)
		case "Style":
			FormDivBasicFieldToField(&(svg_.Style), formDiv)
		}
	}

	// manage the suppress operation
	if svgFormCallback.formGroup.HasSuppressButtonBeenPressed {
		svg_.Unstage(svgFormCallback.probe.stageOfInterest)
	}

	svgFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Svg](
		svgFormCallback.probe,
	)
	svgFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if svgFormCallback.CreationMode || svgFormCallback.formGroup.HasSuppressButtonBeenPressed {
		svgFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(svgFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SvgFormCallback(
			nil,
			svgFormCallback.probe,
			newFormGroup,
		)
		svg := new(models.Svg)
		FillUpForm(svg, newFormGroup, svgFormCallback.probe)
		svgFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(svgFormCallback.probe)
}
func __gong__New__TableFormCallback(
	table *models.Table,
	probe *Probe,
	formGroup *table.FormGroup,
) (tableFormCallback *TableFormCallback) {
	tableFormCallback = new(TableFormCallback)
	tableFormCallback.probe = probe
	tableFormCallback.table = table
	tableFormCallback.formGroup = formGroup

	tableFormCallback.CreationMode = (table == nil)

	return
}

type TableFormCallback struct {
	table *models.Table

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (tableFormCallback *TableFormCallback) OnSave() {

	// log.Println("TableFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tableFormCallback.probe.formStage.Checkout()

	if tableFormCallback.table == nil {
		tableFormCallback.table = new(models.Table).Stage(tableFormCallback.probe.stageOfInterest)
	}
	table_ := tableFormCallback.table
	_ = table_

	for _, formDiv := range tableFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(table_.Name), formDiv)
		case "StackName":
			FormDivBasicFieldToField(&(table_.StackName), formDiv)
		case "TableName":
			FormDivBasicFieldToField(&(table_.TableName), formDiv)
		}
	}

	// manage the suppress operation
	if tableFormCallback.formGroup.HasSuppressButtonBeenPressed {
		table_.Unstage(tableFormCallback.probe.stageOfInterest)
	}

	tableFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Table](
		tableFormCallback.probe,
	)
	tableFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if tableFormCallback.CreationMode || tableFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tableFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(tableFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TableFormCallback(
			nil,
			tableFormCallback.probe,
			newFormGroup,
		)
		table := new(models.Table)
		FillUpForm(table, newFormGroup, tableFormCallback.probe)
		tableFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(tableFormCallback.probe)
}
func __gong__New__TitleFormCallback(
	title *models.Title,
	probe *Probe,
	formGroup *table.FormGroup,
) (titleFormCallback *TitleFormCallback) {
	titleFormCallback = new(TitleFormCallback)
	titleFormCallback.probe = probe
	titleFormCallback.title = title
	titleFormCallback.formGroup = formGroup

	titleFormCallback.CreationMode = (title == nil)

	return
}

type TitleFormCallback struct {
	title *models.Title

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (titleFormCallback *TitleFormCallback) OnSave() {

	// log.Println("TitleFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	titleFormCallback.probe.formStage.Checkout()

	if titleFormCallback.title == nil {
		titleFormCallback.title = new(models.Title).Stage(titleFormCallback.probe.stageOfInterest)
	}
	title_ := titleFormCallback.title
	_ = title_

	for _, formDiv := range titleFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(title_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if titleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		title_.Unstage(titleFormCallback.probe.stageOfInterest)
	}

	titleFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Title](
		titleFormCallback.probe,
	)
	titleFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if titleFormCallback.CreationMode || titleFormCallback.formGroup.HasSuppressButtonBeenPressed {
		titleFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(titleFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TitleFormCallback(
			nil,
			titleFormCallback.probe,
			newFormGroup,
		)
		title := new(models.Title)
		FillUpForm(title, newFormGroup, titleFormCallback.probe)
		titleFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(titleFormCallback.probe)
}
func __gong__New__ToneFormCallback(
	tone *models.Tone,
	probe *Probe,
	formGroup *table.FormGroup,
) (toneFormCallback *ToneFormCallback) {
	toneFormCallback = new(ToneFormCallback)
	toneFormCallback.probe = probe
	toneFormCallback.tone = tone
	toneFormCallback.formGroup = formGroup

	toneFormCallback.CreationMode = (tone == nil)

	return
}

type ToneFormCallback struct {
	tone *models.Tone

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (toneFormCallback *ToneFormCallback) OnSave() {

	// log.Println("ToneFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	toneFormCallback.probe.formStage.Checkout()

	if toneFormCallback.tone == nil {
		toneFormCallback.tone = new(models.Tone).Stage(toneFormCallback.probe.stageOfInterest)
	}
	tone_ := toneFormCallback.tone
	_ = tone_

	for _, formDiv := range toneFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tone_.Name), formDiv)
		case "StackName":
			FormDivBasicFieldToField(&(tone_.StackName), formDiv)
		}
	}

	// manage the suppress operation
	if toneFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tone_.Unstage(toneFormCallback.probe.stageOfInterest)
	}

	toneFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Tone](
		toneFormCallback.probe,
	)
	toneFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if toneFormCallback.CreationMode || toneFormCallback.formGroup.HasSuppressButtonBeenPressed {
		toneFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(toneFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ToneFormCallback(
			nil,
			toneFormCallback.probe,
			newFormGroup,
		)
		tone := new(models.Tone)
		FillUpForm(tone, newFormGroup, toneFormCallback.probe)
		toneFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(toneFormCallback.probe)
}
func __gong__New__TreeFormCallback(
	tree *models.Tree,
	probe *Probe,
	formGroup *table.FormGroup,
) (treeFormCallback *TreeFormCallback) {
	treeFormCallback = new(TreeFormCallback)
	treeFormCallback.probe = probe
	treeFormCallback.tree = tree
	treeFormCallback.formGroup = formGroup

	treeFormCallback.CreationMode = (tree == nil)

	return
}

type TreeFormCallback struct {
	tree *models.Tree

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (treeFormCallback *TreeFormCallback) OnSave() {

	// log.Println("TreeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	treeFormCallback.probe.formStage.Checkout()

	if treeFormCallback.tree == nil {
		treeFormCallback.tree = new(models.Tree).Stage(treeFormCallback.probe.stageOfInterest)
	}
	tree_ := treeFormCallback.tree
	_ = tree_

	for _, formDiv := range treeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tree_.Name), formDiv)
		case "StackName":
			FormDivBasicFieldToField(&(tree_.StackName), formDiv)
		case "TreeName":
			FormDivBasicFieldToField(&(tree_.TreeName), formDiv)
		}
	}

	// manage the suppress operation
	if treeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tree_.Unstage(treeFormCallback.probe.stageOfInterest)
	}

	treeFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Tree](
		treeFormCallback.probe,
	)
	treeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if treeFormCallback.CreationMode || treeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		treeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(treeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TreeFormCallback(
			nil,
			treeFormCallback.probe,
			newFormGroup,
		)
		tree := new(models.Tree)
		FillUpForm(tree, newFormGroup, treeFormCallback.probe)
		treeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(treeFormCallback.probe)
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

	// log.Println("ViewFormCallback, OnSave")

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
		case "ShowViewName":
			FormDivBasicFieldToField(&(view_.ShowViewName), formDiv)
		case "RootAsSplitAreas":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.AsSplitArea](viewFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.AsSplitArea, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.AsSplitArea)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					viewFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			view_.RootAsSplitAreas = instanceSlice

		case "IsSelectedView":
			FormDivBasicFieldToField(&(view_.IsSelectedView), formDiv)
		}
	}

	// manage the suppress operation
	if viewFormCallback.formGroup.HasSuppressButtonBeenPressed {
		view_.Unstage(viewFormCallback.probe.stageOfInterest)
	}

	viewFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.View](
		viewFormCallback.probe,
	)
	viewFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if viewFormCallback.CreationMode || viewFormCallback.formGroup.HasSuppressButtonBeenPressed {
		viewFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
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

	updateAndCommitTree(viewFormCallback.probe)
}
func __gong__New__XlsxFormCallback(
	xlsx *models.Xlsx,
	probe *Probe,
	formGroup *table.FormGroup,
) (xlsxFormCallback *XlsxFormCallback) {
	xlsxFormCallback = new(XlsxFormCallback)
	xlsxFormCallback.probe = probe
	xlsxFormCallback.xlsx = xlsx
	xlsxFormCallback.formGroup = formGroup

	xlsxFormCallback.CreationMode = (xlsx == nil)

	return
}

type XlsxFormCallback struct {
	xlsx *models.Xlsx

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (xlsxFormCallback *XlsxFormCallback) OnSave() {

	// log.Println("XlsxFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xlsxFormCallback.probe.formStage.Checkout()

	if xlsxFormCallback.xlsx == nil {
		xlsxFormCallback.xlsx = new(models.Xlsx).Stage(xlsxFormCallback.probe.stageOfInterest)
	}
	xlsx_ := xlsxFormCallback.xlsx
	_ = xlsx_

	for _, formDiv := range xlsxFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xlsx_.Name), formDiv)
		case "StackName":
			FormDivBasicFieldToField(&(xlsx_.StackName), formDiv)
		}
	}

	// manage the suppress operation
	if xlsxFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xlsx_.Unstage(xlsxFormCallback.probe.stageOfInterest)
	}

	xlsxFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Xlsx](
		xlsxFormCallback.probe,
	)
	xlsxFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xlsxFormCallback.CreationMode || xlsxFormCallback.formGroup.HasSuppressButtonBeenPressed {
		xlsxFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(xlsxFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__XlsxFormCallback(
			nil,
			xlsxFormCallback.probe,
			newFormGroup,
		)
		xlsx := new(models.Xlsx)
		FillUpForm(xlsx, newFormGroup, xlsxFormCallback.probe)
		xlsxFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(xlsxFormCallback.probe)
}
