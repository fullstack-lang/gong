// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *AsSplit:
		ok = stage.IsStagedAsSplit(target)

	case *AsSplitArea:
		ok = stage.IsStagedAsSplitArea(target)

	case *Button:
		ok = stage.IsStagedButton(target)

	case *Cursor:
		ok = stage.IsStagedCursor(target)

	case *Doc:
		ok = stage.IsStagedDoc(target)

	case *Form:
		ok = stage.IsStagedForm(target)

	case *Load:
		ok = stage.IsStagedLoad(target)

	case *Slider:
		ok = stage.IsStagedSlider(target)

	case *Split:
		ok = stage.IsStagedSplit(target)

	case *Svg:
		ok = stage.IsStagedSvg(target)

	case *Table:
		ok = stage.IsStagedTable(target)

	case *Tone:
		ok = stage.IsStagedTone(target)

	case *Tree:
		ok = stage.IsStagedTree(target)

	case *View:
		ok = stage.IsStagedView(target)

	case *Xlsx:
		ok = stage.IsStagedXlsx(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedAsSplit(assplit *AsSplit) (ok bool) {

	_, ok = stage.AsSplits[assplit]

	return
}

func (stage *Stage) IsStagedAsSplitArea(assplitarea *AsSplitArea) (ok bool) {

	_, ok = stage.AsSplitAreas[assplitarea]

	return
}

func (stage *Stage) IsStagedButton(button *Button) (ok bool) {

	_, ok = stage.Buttons[button]

	return
}

func (stage *Stage) IsStagedCursor(cursor *Cursor) (ok bool) {

	_, ok = stage.Cursors[cursor]

	return
}

func (stage *Stage) IsStagedDoc(doc *Doc) (ok bool) {

	_, ok = stage.Docs[doc]

	return
}

func (stage *Stage) IsStagedForm(form *Form) (ok bool) {

	_, ok = stage.Forms[form]

	return
}

func (stage *Stage) IsStagedLoad(load *Load) (ok bool) {

	_, ok = stage.Loads[load]

	return
}

func (stage *Stage) IsStagedSlider(slider *Slider) (ok bool) {

	_, ok = stage.Sliders[slider]

	return
}

func (stage *Stage) IsStagedSplit(split *Split) (ok bool) {

	_, ok = stage.Splits[split]

	return
}

func (stage *Stage) IsStagedSvg(svg *Svg) (ok bool) {

	_, ok = stage.Svgs[svg]

	return
}

func (stage *Stage) IsStagedTable(table *Table) (ok bool) {

	_, ok = stage.Tables[table]

	return
}

func (stage *Stage) IsStagedTone(tone *Tone) (ok bool) {

	_, ok = stage.Tones[tone]

	return
}

func (stage *Stage) IsStagedTree(tree *Tree) (ok bool) {

	_, ok = stage.Trees[tree]

	return
}

func (stage *Stage) IsStagedView(view *View) (ok bool) {

	_, ok = stage.Views[view]

	return
}

func (stage *Stage) IsStagedXlsx(xlsx *Xlsx) (ok bool) {

	_, ok = stage.Xlsxs[xlsx]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *AsSplit:
		stage.StageBranchAsSplit(target)

	case *AsSplitArea:
		stage.StageBranchAsSplitArea(target)

	case *Button:
		stage.StageBranchButton(target)

	case *Cursor:
		stage.StageBranchCursor(target)

	case *Doc:
		stage.StageBranchDoc(target)

	case *Form:
		stage.StageBranchForm(target)

	case *Load:
		stage.StageBranchLoad(target)

	case *Slider:
		stage.StageBranchSlider(target)

	case *Split:
		stage.StageBranchSplit(target)

	case *Svg:
		stage.StageBranchSvg(target)

	case *Table:
		stage.StageBranchTable(target)

	case *Tone:
		stage.StageBranchTone(target)

	case *Tree:
		stage.StageBranchTree(target)

	case *View:
		stage.StageBranchView(target)

	case *Xlsx:
		stage.StageBranchXlsx(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchAsSplit(assplit *AsSplit) {

	// check if instance is already staged
	if IsStaged(stage, assplit) {
		return
	}

	assplit.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assplitarea := range assplit.AsSplitAreas {
		StageBranch(stage, _assplitarea)
	}

}

func (stage *Stage) StageBranchAsSplitArea(assplitarea *AsSplitArea) {

	// check if instance is already staged
	if IsStaged(stage, assplitarea) {
		return
	}

	assplitarea.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if assplitarea.AsSplit != nil {
		StageBranch(stage, assplitarea.AsSplit)
	}
	if assplitarea.Button != nil {
		StageBranch(stage, assplitarea.Button)
	}
	if assplitarea.Cursor != nil {
		StageBranch(stage, assplitarea.Cursor)
	}
	if assplitarea.Doc != nil {
		StageBranch(stage, assplitarea.Doc)
	}
	if assplitarea.Form != nil {
		StageBranch(stage, assplitarea.Form)
	}
	if assplitarea.Load != nil {
		StageBranch(stage, assplitarea.Load)
	}
	if assplitarea.Slider != nil {
		StageBranch(stage, assplitarea.Slider)
	}
	if assplitarea.Split != nil {
		StageBranch(stage, assplitarea.Split)
	}
	if assplitarea.Svg != nil {
		StageBranch(stage, assplitarea.Svg)
	}
	if assplitarea.Table != nil {
		StageBranch(stage, assplitarea.Table)
	}
	if assplitarea.Tone != nil {
		StageBranch(stage, assplitarea.Tone)
	}
	if assplitarea.Tree != nil {
		StageBranch(stage, assplitarea.Tree)
	}
	if assplitarea.Xlsx != nil {
		StageBranch(stage, assplitarea.Xlsx)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchButton(button *Button) {

	// check if instance is already staged
	if IsStaged(stage, button) {
		return
	}

	button.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchCursor(cursor *Cursor) {

	// check if instance is already staged
	if IsStaged(stage, cursor) {
		return
	}

	cursor.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchDoc(doc *Doc) {

	// check if instance is already staged
	if IsStaged(stage, doc) {
		return
	}

	doc.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchForm(form *Form) {

	// check if instance is already staged
	if IsStaged(stage, form) {
		return
	}

	form.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchLoad(load *Load) {

	// check if instance is already staged
	if IsStaged(stage, load) {
		return
	}

	load.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchSlider(slider *Slider) {

	// check if instance is already staged
	if IsStaged(stage, slider) {
		return
	}

	slider.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchSplit(split *Split) {

	// check if instance is already staged
	if IsStaged(stage, split) {
		return
	}

	split.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchSvg(svg *Svg) {

	// check if instance is already staged
	if IsStaged(stage, svg) {
		return
	}

	svg.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTable(table *Table) {

	// check if instance is already staged
	if IsStaged(stage, table) {
		return
	}

	table.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTone(tone *Tone) {

	// check if instance is already staged
	if IsStaged(stage, tone) {
		return
	}

	tone.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTree(tree *Tree) {

	// check if instance is already staged
	if IsStaged(stage, tree) {
		return
	}

	tree.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchView(view *View) {

	// check if instance is already staged
	if IsStaged(stage, view) {
		return
	}

	view.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assplitarea := range view.RootAsSplitAreas {
		StageBranch(stage, _assplitarea)
	}

}

func (stage *Stage) StageBranchXlsx(xlsx *Xlsx) {

	// check if instance is already staged
	if IsStaged(stage, xlsx) {
		return
	}

	xlsx.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *AsSplit:
		toT := CopyBranchAsSplit(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AsSplitArea:
		toT := CopyBranchAsSplitArea(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Button:
		toT := CopyBranchButton(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Cursor:
		toT := CopyBranchCursor(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Doc:
		toT := CopyBranchDoc(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Form:
		toT := CopyBranchForm(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Load:
		toT := CopyBranchLoad(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Slider:
		toT := CopyBranchSlider(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Split:
		toT := CopyBranchSplit(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Svg:
		toT := CopyBranchSvg(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Table:
		toT := CopyBranchTable(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tone:
		toT := CopyBranchTone(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tree:
		toT := CopyBranchTree(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *View:
		toT := CopyBranchView(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xlsx:
		toT := CopyBranchXlsx(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchAsSplit(mapOrigCopy map[any]any, assplitFrom *AsSplit) (assplitTo *AsSplit) {

	// assplitFrom has already been copied
	if _assplitTo, ok := mapOrigCopy[assplitFrom]; ok {
		assplitTo = _assplitTo.(*AsSplit)
		return
	}

	assplitTo = new(AsSplit)
	mapOrigCopy[assplitFrom] = assplitTo
	assplitFrom.CopyBasicFields(assplitTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assplitarea := range assplitFrom.AsSplitAreas {
		assplitTo.AsSplitAreas = append(assplitTo.AsSplitAreas, CopyBranchAsSplitArea(mapOrigCopy, _assplitarea))
	}

	return
}

func CopyBranchAsSplitArea(mapOrigCopy map[any]any, assplitareaFrom *AsSplitArea) (assplitareaTo *AsSplitArea) {

	// assplitareaFrom has already been copied
	if _assplitareaTo, ok := mapOrigCopy[assplitareaFrom]; ok {
		assplitareaTo = _assplitareaTo.(*AsSplitArea)
		return
	}

	assplitareaTo = new(AsSplitArea)
	mapOrigCopy[assplitareaFrom] = assplitareaTo
	assplitareaFrom.CopyBasicFields(assplitareaTo)

	//insertion point for the staging of instances referenced by pointers
	if assplitareaFrom.AsSplit != nil {
		assplitareaTo.AsSplit = CopyBranchAsSplit(mapOrigCopy, assplitareaFrom.AsSplit)
	}
	if assplitareaFrom.Button != nil {
		assplitareaTo.Button = CopyBranchButton(mapOrigCopy, assplitareaFrom.Button)
	}
	if assplitareaFrom.Cursor != nil {
		assplitareaTo.Cursor = CopyBranchCursor(mapOrigCopy, assplitareaFrom.Cursor)
	}
	if assplitareaFrom.Doc != nil {
		assplitareaTo.Doc = CopyBranchDoc(mapOrigCopy, assplitareaFrom.Doc)
	}
	if assplitareaFrom.Form != nil {
		assplitareaTo.Form = CopyBranchForm(mapOrigCopy, assplitareaFrom.Form)
	}
	if assplitareaFrom.Load != nil {
		assplitareaTo.Load = CopyBranchLoad(mapOrigCopy, assplitareaFrom.Load)
	}
	if assplitareaFrom.Slider != nil {
		assplitareaTo.Slider = CopyBranchSlider(mapOrigCopy, assplitareaFrom.Slider)
	}
	if assplitareaFrom.Split != nil {
		assplitareaTo.Split = CopyBranchSplit(mapOrigCopy, assplitareaFrom.Split)
	}
	if assplitareaFrom.Svg != nil {
		assplitareaTo.Svg = CopyBranchSvg(mapOrigCopy, assplitareaFrom.Svg)
	}
	if assplitareaFrom.Table != nil {
		assplitareaTo.Table = CopyBranchTable(mapOrigCopy, assplitareaFrom.Table)
	}
	if assplitareaFrom.Tone != nil {
		assplitareaTo.Tone = CopyBranchTone(mapOrigCopy, assplitareaFrom.Tone)
	}
	if assplitareaFrom.Tree != nil {
		assplitareaTo.Tree = CopyBranchTree(mapOrigCopy, assplitareaFrom.Tree)
	}
	if assplitareaFrom.Xlsx != nil {
		assplitareaTo.Xlsx = CopyBranchXlsx(mapOrigCopy, assplitareaFrom.Xlsx)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchButton(mapOrigCopy map[any]any, buttonFrom *Button) (buttonTo *Button) {

	// buttonFrom has already been copied
	if _buttonTo, ok := mapOrigCopy[buttonFrom]; ok {
		buttonTo = _buttonTo.(*Button)
		return
	}

	buttonTo = new(Button)
	mapOrigCopy[buttonFrom] = buttonTo
	buttonFrom.CopyBasicFields(buttonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchCursor(mapOrigCopy map[any]any, cursorFrom *Cursor) (cursorTo *Cursor) {

	// cursorFrom has already been copied
	if _cursorTo, ok := mapOrigCopy[cursorFrom]; ok {
		cursorTo = _cursorTo.(*Cursor)
		return
	}

	cursorTo = new(Cursor)
	mapOrigCopy[cursorFrom] = cursorTo
	cursorFrom.CopyBasicFields(cursorTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDoc(mapOrigCopy map[any]any, docFrom *Doc) (docTo *Doc) {

	// docFrom has already been copied
	if _docTo, ok := mapOrigCopy[docFrom]; ok {
		docTo = _docTo.(*Doc)
		return
	}

	docTo = new(Doc)
	mapOrigCopy[docFrom] = docTo
	docFrom.CopyBasicFields(docTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchForm(mapOrigCopy map[any]any, formFrom *Form) (formTo *Form) {

	// formFrom has already been copied
	if _formTo, ok := mapOrigCopy[formFrom]; ok {
		formTo = _formTo.(*Form)
		return
	}

	formTo = new(Form)
	mapOrigCopy[formFrom] = formTo
	formFrom.CopyBasicFields(formTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchLoad(mapOrigCopy map[any]any, loadFrom *Load) (loadTo *Load) {

	// loadFrom has already been copied
	if _loadTo, ok := mapOrigCopy[loadFrom]; ok {
		loadTo = _loadTo.(*Load)
		return
	}

	loadTo = new(Load)
	mapOrigCopy[loadFrom] = loadTo
	loadFrom.CopyBasicFields(loadTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSlider(mapOrigCopy map[any]any, sliderFrom *Slider) (sliderTo *Slider) {

	// sliderFrom has already been copied
	if _sliderTo, ok := mapOrigCopy[sliderFrom]; ok {
		sliderTo = _sliderTo.(*Slider)
		return
	}

	sliderTo = new(Slider)
	mapOrigCopy[sliderFrom] = sliderTo
	sliderFrom.CopyBasicFields(sliderTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSplit(mapOrigCopy map[any]any, splitFrom *Split) (splitTo *Split) {

	// splitFrom has already been copied
	if _splitTo, ok := mapOrigCopy[splitFrom]; ok {
		splitTo = _splitTo.(*Split)
		return
	}

	splitTo = new(Split)
	mapOrigCopy[splitFrom] = splitTo
	splitFrom.CopyBasicFields(splitTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSvg(mapOrigCopy map[any]any, svgFrom *Svg) (svgTo *Svg) {

	// svgFrom has already been copied
	if _svgTo, ok := mapOrigCopy[svgFrom]; ok {
		svgTo = _svgTo.(*Svg)
		return
	}

	svgTo = new(Svg)
	mapOrigCopy[svgFrom] = svgTo
	svgFrom.CopyBasicFields(svgTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTable(mapOrigCopy map[any]any, tableFrom *Table) (tableTo *Table) {

	// tableFrom has already been copied
	if _tableTo, ok := mapOrigCopy[tableFrom]; ok {
		tableTo = _tableTo.(*Table)
		return
	}

	tableTo = new(Table)
	mapOrigCopy[tableFrom] = tableTo
	tableFrom.CopyBasicFields(tableTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTone(mapOrigCopy map[any]any, toneFrom *Tone) (toneTo *Tone) {

	// toneFrom has already been copied
	if _toneTo, ok := mapOrigCopy[toneFrom]; ok {
		toneTo = _toneTo.(*Tone)
		return
	}

	toneTo = new(Tone)
	mapOrigCopy[toneFrom] = toneTo
	toneFrom.CopyBasicFields(toneTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTree(mapOrigCopy map[any]any, treeFrom *Tree) (treeTo *Tree) {

	// treeFrom has already been copied
	if _treeTo, ok := mapOrigCopy[treeFrom]; ok {
		treeTo = _treeTo.(*Tree)
		return
	}

	treeTo = new(Tree)
	mapOrigCopy[treeFrom] = treeTo
	treeFrom.CopyBasicFields(treeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchView(mapOrigCopy map[any]any, viewFrom *View) (viewTo *View) {

	// viewFrom has already been copied
	if _viewTo, ok := mapOrigCopy[viewFrom]; ok {
		viewTo = _viewTo.(*View)
		return
	}

	viewTo = new(View)
	mapOrigCopy[viewFrom] = viewTo
	viewFrom.CopyBasicFields(viewTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assplitarea := range viewFrom.RootAsSplitAreas {
		viewTo.RootAsSplitAreas = append(viewTo.RootAsSplitAreas, CopyBranchAsSplitArea(mapOrigCopy, _assplitarea))
	}

	return
}

func CopyBranchXlsx(mapOrigCopy map[any]any, xlsxFrom *Xlsx) (xlsxTo *Xlsx) {

	// xlsxFrom has already been copied
	if _xlsxTo, ok := mapOrigCopy[xlsxFrom]; ok {
		xlsxTo = _xlsxTo.(*Xlsx)
		return
	}

	xlsxTo = new(Xlsx)
	mapOrigCopy[xlsxFrom] = xlsxTo
	xlsxFrom.CopyBasicFields(xlsxTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *AsSplit:
		stage.UnstageBranchAsSplit(target)

	case *AsSplitArea:
		stage.UnstageBranchAsSplitArea(target)

	case *Button:
		stage.UnstageBranchButton(target)

	case *Cursor:
		stage.UnstageBranchCursor(target)

	case *Doc:
		stage.UnstageBranchDoc(target)

	case *Form:
		stage.UnstageBranchForm(target)

	case *Load:
		stage.UnstageBranchLoad(target)

	case *Slider:
		stage.UnstageBranchSlider(target)

	case *Split:
		stage.UnstageBranchSplit(target)

	case *Svg:
		stage.UnstageBranchSvg(target)

	case *Table:
		stage.UnstageBranchTable(target)

	case *Tone:
		stage.UnstageBranchTone(target)

	case *Tree:
		stage.UnstageBranchTree(target)

	case *View:
		stage.UnstageBranchView(target)

	case *Xlsx:
		stage.UnstageBranchXlsx(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchAsSplit(assplit *AsSplit) {

	// check if instance is already staged
	if !IsStaged(stage, assplit) {
		return
	}

	assplit.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assplitarea := range assplit.AsSplitAreas {
		UnstageBranch(stage, _assplitarea)
	}

}

func (stage *Stage) UnstageBranchAsSplitArea(assplitarea *AsSplitArea) {

	// check if instance is already staged
	if !IsStaged(stage, assplitarea) {
		return
	}

	assplitarea.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if assplitarea.AsSplit != nil {
		UnstageBranch(stage, assplitarea.AsSplit)
	}
	if assplitarea.Button != nil {
		UnstageBranch(stage, assplitarea.Button)
	}
	if assplitarea.Cursor != nil {
		UnstageBranch(stage, assplitarea.Cursor)
	}
	if assplitarea.Doc != nil {
		UnstageBranch(stage, assplitarea.Doc)
	}
	if assplitarea.Form != nil {
		UnstageBranch(stage, assplitarea.Form)
	}
	if assplitarea.Load != nil {
		UnstageBranch(stage, assplitarea.Load)
	}
	if assplitarea.Slider != nil {
		UnstageBranch(stage, assplitarea.Slider)
	}
	if assplitarea.Split != nil {
		UnstageBranch(stage, assplitarea.Split)
	}
	if assplitarea.Svg != nil {
		UnstageBranch(stage, assplitarea.Svg)
	}
	if assplitarea.Table != nil {
		UnstageBranch(stage, assplitarea.Table)
	}
	if assplitarea.Tone != nil {
		UnstageBranch(stage, assplitarea.Tone)
	}
	if assplitarea.Tree != nil {
		UnstageBranch(stage, assplitarea.Tree)
	}
	if assplitarea.Xlsx != nil {
		UnstageBranch(stage, assplitarea.Xlsx)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchButton(button *Button) {

	// check if instance is already staged
	if !IsStaged(stage, button) {
		return
	}

	button.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchCursor(cursor *Cursor) {

	// check if instance is already staged
	if !IsStaged(stage, cursor) {
		return
	}

	cursor.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchDoc(doc *Doc) {

	// check if instance is already staged
	if !IsStaged(stage, doc) {
		return
	}

	doc.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchForm(form *Form) {

	// check if instance is already staged
	if !IsStaged(stage, form) {
		return
	}

	form.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchLoad(load *Load) {

	// check if instance is already staged
	if !IsStaged(stage, load) {
		return
	}

	load.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchSlider(slider *Slider) {

	// check if instance is already staged
	if !IsStaged(stage, slider) {
		return
	}

	slider.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchSplit(split *Split) {

	// check if instance is already staged
	if !IsStaged(stage, split) {
		return
	}

	split.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchSvg(svg *Svg) {

	// check if instance is already staged
	if !IsStaged(stage, svg) {
		return
	}

	svg.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTable(table *Table) {

	// check if instance is already staged
	if !IsStaged(stage, table) {
		return
	}

	table.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTone(tone *Tone) {

	// check if instance is already staged
	if !IsStaged(stage, tone) {
		return
	}

	tone.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTree(tree *Tree) {

	// check if instance is already staged
	if !IsStaged(stage, tree) {
		return
	}

	tree.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchView(view *View) {

	// check if instance is already staged
	if !IsStaged(stage, view) {
		return
	}

	view.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assplitarea := range view.RootAsSplitAreas {
		UnstageBranch(stage, _assplitarea)
	}

}

func (stage *Stage) UnstageBranchXlsx(xlsx *Xlsx) {

	// check if instance is already staged
	if !IsStaged(stage, xlsx) {
		return
	}

	xlsx.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}
