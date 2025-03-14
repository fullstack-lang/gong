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
		case "ShowNameInHeader":
			FormDivBasicFieldToField(&(assplitarea_.ShowNameInHeader), formDiv)
		case "Size":
			FormDivBasicFieldToField(&(assplitarea_.Size), formDiv)
		case "IsAny":
			FormDivBasicFieldToField(&(assplitarea_.IsAny), formDiv)
		case "Tree":
			FormDivSelectFieldToField(&(assplitarea_.Tree), assplitareaFormCallback.probe.stageOfInterest, formDiv)
		case "Table":
			FormDivSelectFieldToField(&(assplitarea_.Table), assplitareaFormCallback.probe.stageOfInterest, formDiv)
		case "Form":
			FormDivSelectFieldToField(&(assplitarea_.Form), assplitareaFormCallback.probe.stageOfInterest, formDiv)
		case "Svg":
			FormDivSelectFieldToField(&(assplitarea_.Svg), assplitareaFormCallback.probe.stageOfInterest, formDiv)
		case "Doc":
			FormDivSelectFieldToField(&(assplitarea_.Doc), assplitareaFormCallback.probe.stageOfInterest, formDiv)
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

	fillUpTree(assplitareaFormCallback.probe)
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

	log.Println("DocFormCallback, OnSave")

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
	fillUpTable[models.Doc](
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

	fillUpTree(docFormCallback.probe)
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

	log.Println("FormFormCallback, OnSave")

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
	fillUpTable[models.Form](
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

	fillUpTree(formFormCallback.probe)
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

	log.Println("SvgFormCallback, OnSave")

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
		}
	}

	// manage the suppress operation
	if svgFormCallback.formGroup.HasSuppressButtonBeenPressed {
		svg_.Unstage(svgFormCallback.probe.stageOfInterest)
	}

	svgFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Svg](
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

	fillUpTree(svgFormCallback.probe)
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

	log.Println("TableFormCallback, OnSave")

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
	fillUpTable[models.Table](
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

	fillUpTree(tableFormCallback.probe)
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

	log.Println("TreeFormCallback, OnSave")

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
	fillUpTable[models.Tree](
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

	fillUpTree(treeFormCallback.probe)
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

	fillUpTree(viewFormCallback.probe)
}
