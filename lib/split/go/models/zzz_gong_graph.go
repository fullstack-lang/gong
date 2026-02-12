// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

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

	case *FavIcon:
		ok = stage.IsStagedFavIcon(target)

	case *Form:
		ok = stage.IsStagedForm(target)

	case *Load:
		ok = stage.IsStagedLoad(target)

	case *LogoOnTheLeft:
		ok = stage.IsStagedLogoOnTheLeft(target)

	case *LogoOnTheRight:
		ok = stage.IsStagedLogoOnTheRight(target)

	case *Markdown:
		ok = stage.IsStagedMarkdown(target)

	case *Slider:
		ok = stage.IsStagedSlider(target)

	case *Split:
		ok = stage.IsStagedSplit(target)

	case *Svg:
		ok = stage.IsStagedSvg(target)

	case *Table:
		ok = stage.IsStagedTable(target)

	case *Title:
		ok = stage.IsStagedTitle(target)

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

	case *FavIcon:
		ok = stage.IsStagedFavIcon(target)

	case *Form:
		ok = stage.IsStagedForm(target)

	case *Load:
		ok = stage.IsStagedLoad(target)

	case *LogoOnTheLeft:
		ok = stage.IsStagedLogoOnTheLeft(target)

	case *LogoOnTheRight:
		ok = stage.IsStagedLogoOnTheRight(target)

	case *Markdown:
		ok = stage.IsStagedMarkdown(target)

	case *Slider:
		ok = stage.IsStagedSlider(target)

	case *Split:
		ok = stage.IsStagedSplit(target)

	case *Svg:
		ok = stage.IsStagedSvg(target)

	case *Table:
		ok = stage.IsStagedTable(target)

	case *Title:
		ok = stage.IsStagedTitle(target)

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

func (stage *Stage) IsStagedFavIcon(favicon *FavIcon) (ok bool) {

	_, ok = stage.FavIcons[favicon]

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

func (stage *Stage) IsStagedLogoOnTheLeft(logoontheleft *LogoOnTheLeft) (ok bool) {

	_, ok = stage.LogoOnTheLefts[logoontheleft]

	return
}

func (stage *Stage) IsStagedLogoOnTheRight(logoontheright *LogoOnTheRight) (ok bool) {

	_, ok = stage.LogoOnTheRights[logoontheright]

	return
}

func (stage *Stage) IsStagedMarkdown(markdown *Markdown) (ok bool) {

	_, ok = stage.Markdowns[markdown]

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

func (stage *Stage) IsStagedTitle(title *Title) (ok bool) {

	_, ok = stage.Titles[title]

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

	case *FavIcon:
		stage.StageBranchFavIcon(target)

	case *Form:
		stage.StageBranchForm(target)

	case *Load:
		stage.StageBranchLoad(target)

	case *LogoOnTheLeft:
		stage.StageBranchLogoOnTheLeft(target)

	case *LogoOnTheRight:
		stage.StageBranchLogoOnTheRight(target)

	case *Markdown:
		stage.StageBranchMarkdown(target)

	case *Slider:
		stage.StageBranchSlider(target)

	case *Split:
		stage.StageBranchSplit(target)

	case *Svg:
		stage.StageBranchSvg(target)

	case *Table:
		stage.StageBranchTable(target)

	case *Title:
		stage.StageBranchTitle(target)

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
	if assplitarea.Form != nil {
		StageBranch(stage, assplitarea.Form)
	}
	if assplitarea.Load != nil {
		StageBranch(stage, assplitarea.Load)
	}
	if assplitarea.Markdown != nil {
		StageBranch(stage, assplitarea.Markdown)
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

func (stage *Stage) StageBranchFavIcon(favicon *FavIcon) {

	// check if instance is already staged
	if IsStaged(stage, favicon) {
		return
	}

	favicon.Stage(stage)

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

func (stage *Stage) StageBranchLogoOnTheLeft(logoontheleft *LogoOnTheLeft) {

	// check if instance is already staged
	if IsStaged(stage, logoontheleft) {
		return
	}

	logoontheleft.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchLogoOnTheRight(logoontheright *LogoOnTheRight) {

	// check if instance is already staged
	if IsStaged(stage, logoontheright) {
		return
	}

	logoontheright.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchMarkdown(markdown *Markdown) {

	// check if instance is already staged
	if IsStaged(stage, markdown) {
		return
	}

	markdown.Stage(stage)

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

func (stage *Stage) StageBranchTitle(title *Title) {

	// check if instance is already staged
	if IsStaged(stage, title) {
		return
	}

	title.Stage(stage)

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

	case *FavIcon:
		toT := CopyBranchFavIcon(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Form:
		toT := CopyBranchForm(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Load:
		toT := CopyBranchLoad(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LogoOnTheLeft:
		toT := CopyBranchLogoOnTheLeft(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LogoOnTheRight:
		toT := CopyBranchLogoOnTheRight(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Markdown:
		toT := CopyBranchMarkdown(mapOrigCopy, fromT)
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

	case *Title:
		toT := CopyBranchTitle(mapOrigCopy, fromT)
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
	if assplitareaFrom.Form != nil {
		assplitareaTo.Form = CopyBranchForm(mapOrigCopy, assplitareaFrom.Form)
	}
	if assplitareaFrom.Load != nil {
		assplitareaTo.Load = CopyBranchLoad(mapOrigCopy, assplitareaFrom.Load)
	}
	if assplitareaFrom.Markdown != nil {
		assplitareaTo.Markdown = CopyBranchMarkdown(mapOrigCopy, assplitareaFrom.Markdown)
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

func CopyBranchFavIcon(mapOrigCopy map[any]any, faviconFrom *FavIcon) (faviconTo *FavIcon) {

	// faviconFrom has already been copied
	if _faviconTo, ok := mapOrigCopy[faviconFrom]; ok {
		faviconTo = _faviconTo.(*FavIcon)
		return
	}

	faviconTo = new(FavIcon)
	mapOrigCopy[faviconFrom] = faviconTo
	faviconFrom.CopyBasicFields(faviconTo)

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

func CopyBranchLogoOnTheLeft(mapOrigCopy map[any]any, logoontheleftFrom *LogoOnTheLeft) (logoontheleftTo *LogoOnTheLeft) {

	// logoontheleftFrom has already been copied
	if _logoontheleftTo, ok := mapOrigCopy[logoontheleftFrom]; ok {
		logoontheleftTo = _logoontheleftTo.(*LogoOnTheLeft)
		return
	}

	logoontheleftTo = new(LogoOnTheLeft)
	mapOrigCopy[logoontheleftFrom] = logoontheleftTo
	logoontheleftFrom.CopyBasicFields(logoontheleftTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchLogoOnTheRight(mapOrigCopy map[any]any, logoontherightFrom *LogoOnTheRight) (logoontherightTo *LogoOnTheRight) {

	// logoontherightFrom has already been copied
	if _logoontherightTo, ok := mapOrigCopy[logoontherightFrom]; ok {
		logoontherightTo = _logoontherightTo.(*LogoOnTheRight)
		return
	}

	logoontherightTo = new(LogoOnTheRight)
	mapOrigCopy[logoontherightFrom] = logoontherightTo
	logoontherightFrom.CopyBasicFields(logoontherightTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMarkdown(mapOrigCopy map[any]any, markdownFrom *Markdown) (markdownTo *Markdown) {

	// markdownFrom has already been copied
	if _markdownTo, ok := mapOrigCopy[markdownFrom]; ok {
		markdownTo = _markdownTo.(*Markdown)
		return
	}

	markdownTo = new(Markdown)
	mapOrigCopy[markdownFrom] = markdownTo
	markdownFrom.CopyBasicFields(markdownTo)

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

func CopyBranchTitle(mapOrigCopy map[any]any, titleFrom *Title) (titleTo *Title) {

	// titleFrom has already been copied
	if _titleTo, ok := mapOrigCopy[titleFrom]; ok {
		titleTo = _titleTo.(*Title)
		return
	}

	titleTo = new(Title)
	mapOrigCopy[titleFrom] = titleTo
	titleFrom.CopyBasicFields(titleTo)

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

	case *FavIcon:
		stage.UnstageBranchFavIcon(target)

	case *Form:
		stage.UnstageBranchForm(target)

	case *Load:
		stage.UnstageBranchLoad(target)

	case *LogoOnTheLeft:
		stage.UnstageBranchLogoOnTheLeft(target)

	case *LogoOnTheRight:
		stage.UnstageBranchLogoOnTheRight(target)

	case *Markdown:
		stage.UnstageBranchMarkdown(target)

	case *Slider:
		stage.UnstageBranchSlider(target)

	case *Split:
		stage.UnstageBranchSplit(target)

	case *Svg:
		stage.UnstageBranchSvg(target)

	case *Table:
		stage.UnstageBranchTable(target)

	case *Title:
		stage.UnstageBranchTitle(target)

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
	if assplitarea.Form != nil {
		UnstageBranch(stage, assplitarea.Form)
	}
	if assplitarea.Load != nil {
		UnstageBranch(stage, assplitarea.Load)
	}
	if assplitarea.Markdown != nil {
		UnstageBranch(stage, assplitarea.Markdown)
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

func (stage *Stage) UnstageBranchFavIcon(favicon *FavIcon) {

	// check if instance is already staged
	if !IsStaged(stage, favicon) {
		return
	}

	favicon.Unstage(stage)

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

func (stage *Stage) UnstageBranchLogoOnTheLeft(logoontheleft *LogoOnTheLeft) {

	// check if instance is already staged
	if !IsStaged(stage, logoontheleft) {
		return
	}

	logoontheleft.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchLogoOnTheRight(logoontheright *LogoOnTheRight) {

	// check if instance is already staged
	if !IsStaged(stage, logoontheright) {
		return
	}

	logoontheright.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchMarkdown(markdown *Markdown) {

	// check if instance is already staged
	if !IsStaged(stage, markdown) {
		return
	}

	markdown.Unstage(stage)

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

func (stage *Stage) UnstageBranchTitle(title *Title) {

	// check if instance is already staged
	if !IsStaged(stage, title) {
		return
	}

	title.Unstage(stage)

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

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (assplit *AsSplit) GongDiff(stage *Stage, assplitOther *AsSplit) (diffs []string) {
	// insertion point for field diffs
	if assplit.Name != assplitOther.Name {
		diffs = append(diffs, assplit.GongMarshallField(stage, "Name"))
	}
	if assplit.Direction != assplitOther.Direction {
		diffs = append(diffs, assplit.GongMarshallField(stage, "Direction"))
	}
	AsSplitAreasDifferent := false
	if len(assplit.AsSplitAreas) != len(assplitOther.AsSplitAreas) {
		AsSplitAreasDifferent = true
	} else {
		for i := range assplit.AsSplitAreas {
			if (assplit.AsSplitAreas[i] == nil) != (assplitOther.AsSplitAreas[i] == nil) {
				AsSplitAreasDifferent = true
				break
			} else if assplit.AsSplitAreas[i] != nil && assplitOther.AsSplitAreas[i] != nil {
				// this is a pointer comparaison
				if assplit.AsSplitAreas[i] != assplitOther.AsSplitAreas[i] {
					AsSplitAreasDifferent = true
					break
				}
			}
		}
	}
	if AsSplitAreasDifferent {
		ops := Diff(stage, assplit, assplitOther, "AsSplitAreas", assplitOther.AsSplitAreas, assplit.AsSplitAreas)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (assplitarea *AsSplitArea) GongDiff(stage *Stage, assplitareaOther *AsSplitArea) (diffs []string) {
	// insertion point for field diffs
	if assplitarea.Name != assplitareaOther.Name {
		diffs = append(diffs, assplitarea.GongMarshallField(stage, "Name"))
	}
	if assplitarea.ShowNameInHeader != assplitareaOther.ShowNameInHeader {
		diffs = append(diffs, assplitarea.GongMarshallField(stage, "ShowNameInHeader"))
	}
	if assplitarea.Size != assplitareaOther.Size {
		diffs = append(diffs, assplitarea.GongMarshallField(stage, "Size"))
	}
	if assplitarea.IsAny != assplitareaOther.IsAny {
		diffs = append(diffs, assplitarea.GongMarshallField(stage, "IsAny"))
	}
	if (assplitarea.AsSplit == nil) != (assplitareaOther.AsSplit == nil) {
		diffs = append(diffs, assplitarea.GongMarshallField(stage, "AsSplit"))
	} else if assplitarea.AsSplit != nil && assplitareaOther.AsSplit != nil {
		if assplitarea.AsSplit != assplitareaOther.AsSplit {
			diffs = append(diffs, assplitarea.GongMarshallField(stage, "AsSplit"))
		}
	}
	if (assplitarea.Button == nil) != (assplitareaOther.Button == nil) {
		diffs = append(diffs, assplitarea.GongMarshallField(stage, "Button"))
	} else if assplitarea.Button != nil && assplitareaOther.Button != nil {
		if assplitarea.Button != assplitareaOther.Button {
			diffs = append(diffs, assplitarea.GongMarshallField(stage, "Button"))
		}
	}
	if (assplitarea.Cursor == nil) != (assplitareaOther.Cursor == nil) {
		diffs = append(diffs, assplitarea.GongMarshallField(stage, "Cursor"))
	} else if assplitarea.Cursor != nil && assplitareaOther.Cursor != nil {
		if assplitarea.Cursor != assplitareaOther.Cursor {
			diffs = append(diffs, assplitarea.GongMarshallField(stage, "Cursor"))
		}
	}
	if (assplitarea.Form == nil) != (assplitareaOther.Form == nil) {
		diffs = append(diffs, assplitarea.GongMarshallField(stage, "Form"))
	} else if assplitarea.Form != nil && assplitareaOther.Form != nil {
		if assplitarea.Form != assplitareaOther.Form {
			diffs = append(diffs, assplitarea.GongMarshallField(stage, "Form"))
		}
	}
	if (assplitarea.Load == nil) != (assplitareaOther.Load == nil) {
		diffs = append(diffs, assplitarea.GongMarshallField(stage, "Load"))
	} else if assplitarea.Load != nil && assplitareaOther.Load != nil {
		if assplitarea.Load != assplitareaOther.Load {
			diffs = append(diffs, assplitarea.GongMarshallField(stage, "Load"))
		}
	}
	if (assplitarea.Markdown == nil) != (assplitareaOther.Markdown == nil) {
		diffs = append(diffs, assplitarea.GongMarshallField(stage, "Markdown"))
	} else if assplitarea.Markdown != nil && assplitareaOther.Markdown != nil {
		if assplitarea.Markdown != assplitareaOther.Markdown {
			diffs = append(diffs, assplitarea.GongMarshallField(stage, "Markdown"))
		}
	}
	if (assplitarea.Slider == nil) != (assplitareaOther.Slider == nil) {
		diffs = append(diffs, assplitarea.GongMarshallField(stage, "Slider"))
	} else if assplitarea.Slider != nil && assplitareaOther.Slider != nil {
		if assplitarea.Slider != assplitareaOther.Slider {
			diffs = append(diffs, assplitarea.GongMarshallField(stage, "Slider"))
		}
	}
	if (assplitarea.Split == nil) != (assplitareaOther.Split == nil) {
		diffs = append(diffs, assplitarea.GongMarshallField(stage, "Split"))
	} else if assplitarea.Split != nil && assplitareaOther.Split != nil {
		if assplitarea.Split != assplitareaOther.Split {
			diffs = append(diffs, assplitarea.GongMarshallField(stage, "Split"))
		}
	}
	if (assplitarea.Svg == nil) != (assplitareaOther.Svg == nil) {
		diffs = append(diffs, assplitarea.GongMarshallField(stage, "Svg"))
	} else if assplitarea.Svg != nil && assplitareaOther.Svg != nil {
		if assplitarea.Svg != assplitareaOther.Svg {
			diffs = append(diffs, assplitarea.GongMarshallField(stage, "Svg"))
		}
	}
	if (assplitarea.Table == nil) != (assplitareaOther.Table == nil) {
		diffs = append(diffs, assplitarea.GongMarshallField(stage, "Table"))
	} else if assplitarea.Table != nil && assplitareaOther.Table != nil {
		if assplitarea.Table != assplitareaOther.Table {
			diffs = append(diffs, assplitarea.GongMarshallField(stage, "Table"))
		}
	}
	if (assplitarea.Tone == nil) != (assplitareaOther.Tone == nil) {
		diffs = append(diffs, assplitarea.GongMarshallField(stage, "Tone"))
	} else if assplitarea.Tone != nil && assplitareaOther.Tone != nil {
		if assplitarea.Tone != assplitareaOther.Tone {
			diffs = append(diffs, assplitarea.GongMarshallField(stage, "Tone"))
		}
	}
	if (assplitarea.Tree == nil) != (assplitareaOther.Tree == nil) {
		diffs = append(diffs, assplitarea.GongMarshallField(stage, "Tree"))
	} else if assplitarea.Tree != nil && assplitareaOther.Tree != nil {
		if assplitarea.Tree != assplitareaOther.Tree {
			diffs = append(diffs, assplitarea.GongMarshallField(stage, "Tree"))
		}
	}
	if (assplitarea.Xlsx == nil) != (assplitareaOther.Xlsx == nil) {
		diffs = append(diffs, assplitarea.GongMarshallField(stage, "Xlsx"))
	} else if assplitarea.Xlsx != nil && assplitareaOther.Xlsx != nil {
		if assplitarea.Xlsx != assplitareaOther.Xlsx {
			diffs = append(diffs, assplitarea.GongMarshallField(stage, "Xlsx"))
		}
	}
	if assplitarea.HasDiv != assplitareaOther.HasDiv {
		diffs = append(diffs, assplitarea.GongMarshallField(stage, "HasDiv"))
	}
	if assplitarea.DivStyle != assplitareaOther.DivStyle {
		diffs = append(diffs, assplitarea.GongMarshallField(stage, "DivStyle"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (button *Button) GongDiff(stage *Stage, buttonOther *Button) (diffs []string) {
	// insertion point for field diffs
	if button.Name != buttonOther.Name {
		diffs = append(diffs, button.GongMarshallField(stage, "Name"))
	}
	if button.StackName != buttonOther.StackName {
		diffs = append(diffs, button.GongMarshallField(stage, "StackName"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (cursor *Cursor) GongDiff(stage *Stage, cursorOther *Cursor) (diffs []string) {
	// insertion point for field diffs
	if cursor.Name != cursorOther.Name {
		diffs = append(diffs, cursor.GongMarshallField(stage, "Name"))
	}
	if cursor.StackName != cursorOther.StackName {
		diffs = append(diffs, cursor.GongMarshallField(stage, "StackName"))
	}
	if cursor.Style != cursorOther.Style {
		diffs = append(diffs, cursor.GongMarshallField(stage, "Style"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (favicon *FavIcon) GongDiff(stage *Stage, faviconOther *FavIcon) (diffs []string) {
	// insertion point for field diffs
	if favicon.Name != faviconOther.Name {
		diffs = append(diffs, favicon.GongMarshallField(stage, "Name"))
	}
	if favicon.SVG != faviconOther.SVG {
		diffs = append(diffs, favicon.GongMarshallField(stage, "SVG"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (form *Form) GongDiff(stage *Stage, formOther *Form) (diffs []string) {
	// insertion point for field diffs
	if form.Name != formOther.Name {
		diffs = append(diffs, form.GongMarshallField(stage, "Name"))
	}
	if form.StackName != formOther.StackName {
		diffs = append(diffs, form.GongMarshallField(stage, "StackName"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (load *Load) GongDiff(stage *Stage, loadOther *Load) (diffs []string) {
	// insertion point for field diffs
	if load.Name != loadOther.Name {
		diffs = append(diffs, load.GongMarshallField(stage, "Name"))
	}
	if load.StackName != loadOther.StackName {
		diffs = append(diffs, load.GongMarshallField(stage, "StackName"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (logoontheleft *LogoOnTheLeft) GongDiff(stage *Stage, logoontheleftOther *LogoOnTheLeft) (diffs []string) {
	// insertion point for field diffs
	if logoontheleft.Name != logoontheleftOther.Name {
		diffs = append(diffs, logoontheleft.GongMarshallField(stage, "Name"))
	}
	if logoontheleft.Width != logoontheleftOther.Width {
		diffs = append(diffs, logoontheleft.GongMarshallField(stage, "Width"))
	}
	if logoontheleft.Height != logoontheleftOther.Height {
		diffs = append(diffs, logoontheleft.GongMarshallField(stage, "Height"))
	}
	if logoontheleft.SVG != logoontheleftOther.SVG {
		diffs = append(diffs, logoontheleft.GongMarshallField(stage, "SVG"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (logoontheright *LogoOnTheRight) GongDiff(stage *Stage, logoontherightOther *LogoOnTheRight) (diffs []string) {
	// insertion point for field diffs
	if logoontheright.Name != logoontherightOther.Name {
		diffs = append(diffs, logoontheright.GongMarshallField(stage, "Name"))
	}
	if logoontheright.Width != logoontherightOther.Width {
		diffs = append(diffs, logoontheright.GongMarshallField(stage, "Width"))
	}
	if logoontheright.Height != logoontherightOther.Height {
		diffs = append(diffs, logoontheright.GongMarshallField(stage, "Height"))
	}
	if logoontheright.SVG != logoontherightOther.SVG {
		diffs = append(diffs, logoontheright.GongMarshallField(stage, "SVG"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (markdown *Markdown) GongDiff(stage *Stage, markdownOther *Markdown) (diffs []string) {
	// insertion point for field diffs
	if markdown.Name != markdownOther.Name {
		diffs = append(diffs, markdown.GongMarshallField(stage, "Name"))
	}
	if markdown.StackName != markdownOther.StackName {
		diffs = append(diffs, markdown.GongMarshallField(stage, "StackName"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (slider *Slider) GongDiff(stage *Stage, sliderOther *Slider) (diffs []string) {
	// insertion point for field diffs
	if slider.Name != sliderOther.Name {
		diffs = append(diffs, slider.GongMarshallField(stage, "Name"))
	}
	if slider.StackName != sliderOther.StackName {
		diffs = append(diffs, slider.GongMarshallField(stage, "StackName"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (split *Split) GongDiff(stage *Stage, splitOther *Split) (diffs []string) {
	// insertion point for field diffs
	if split.Name != splitOther.Name {
		diffs = append(diffs, split.GongMarshallField(stage, "Name"))
	}
	if split.StackName != splitOther.StackName {
		diffs = append(diffs, split.GongMarshallField(stage, "StackName"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (svg *Svg) GongDiff(stage *Stage, svgOther *Svg) (diffs []string) {
	// insertion point for field diffs
	if svg.Name != svgOther.Name {
		diffs = append(diffs, svg.GongMarshallField(stage, "Name"))
	}
	if svg.StackName != svgOther.StackName {
		diffs = append(diffs, svg.GongMarshallField(stage, "StackName"))
	}
	if svg.Style != svgOther.Style {
		diffs = append(diffs, svg.GongMarshallField(stage, "Style"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (table *Table) GongDiff(stage *Stage, tableOther *Table) (diffs []string) {
	// insertion point for field diffs
	if table.Name != tableOther.Name {
		diffs = append(diffs, table.GongMarshallField(stage, "Name"))
	}
	if table.StackName != tableOther.StackName {
		diffs = append(diffs, table.GongMarshallField(stage, "StackName"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (title *Title) GongDiff(stage *Stage, titleOther *Title) (diffs []string) {
	// insertion point for field diffs
	if title.Name != titleOther.Name {
		diffs = append(diffs, title.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tone *Tone) GongDiff(stage *Stage, toneOther *Tone) (diffs []string) {
	// insertion point for field diffs
	if tone.Name != toneOther.Name {
		diffs = append(diffs, tone.GongMarshallField(stage, "Name"))
	}
	if tone.StackName != toneOther.StackName {
		diffs = append(diffs, tone.GongMarshallField(stage, "StackName"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tree *Tree) GongDiff(stage *Stage, treeOther *Tree) (diffs []string) {
	// insertion point for field diffs
	if tree.Name != treeOther.Name {
		diffs = append(diffs, tree.GongMarshallField(stage, "Name"))
	}
	if tree.StackName != treeOther.StackName {
		diffs = append(diffs, tree.GongMarshallField(stage, "StackName"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (view *View) GongDiff(stage *Stage, viewOther *View) (diffs []string) {
	// insertion point for field diffs
	if view.Name != viewOther.Name {
		diffs = append(diffs, view.GongMarshallField(stage, "Name"))
	}
	if view.ShowViewName != viewOther.ShowViewName {
		diffs = append(diffs, view.GongMarshallField(stage, "ShowViewName"))
	}
	RootAsSplitAreasDifferent := false
	if len(view.RootAsSplitAreas) != len(viewOther.RootAsSplitAreas) {
		RootAsSplitAreasDifferent = true
	} else {
		for i := range view.RootAsSplitAreas {
			if (view.RootAsSplitAreas[i] == nil) != (viewOther.RootAsSplitAreas[i] == nil) {
				RootAsSplitAreasDifferent = true
				break
			} else if view.RootAsSplitAreas[i] != nil && viewOther.RootAsSplitAreas[i] != nil {
				// this is a pointer comparaison
				if view.RootAsSplitAreas[i] != viewOther.RootAsSplitAreas[i] {
					RootAsSplitAreasDifferent = true
					break
				}
			}
		}
	}
	if RootAsSplitAreasDifferent {
		ops := Diff(stage, view, viewOther, "RootAsSplitAreas", viewOther.RootAsSplitAreas, view.RootAsSplitAreas)
		diffs = append(diffs, ops)
	}
	if view.IsSelectedView != viewOther.IsSelectedView {
		diffs = append(diffs, view.GongMarshallField(stage, "IsSelectedView"))
	}
	if view.Direction != viewOther.Direction {
		diffs = append(diffs, view.GongMarshallField(stage, "Direction"))
	}
	if view.IsSecondatyView != viewOther.IsSecondatyView {
		diffs = append(diffs, view.GongMarshallField(stage, "IsSecondatyView"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (xlsx *Xlsx) GongDiff(stage *Stage, xlsxOther *Xlsx) (diffs []string) {
	// insertion point for field diffs
	if xlsx.Name != xlsxOther.Name {
		diffs = append(diffs, xlsx.GongMarshallField(stage, "Name"))
	}
	if xlsx.StackName != xlsxOther.StackName {
		diffs = append(diffs, xlsx.GongMarshallField(stage, "StackName"))
	}

	return
}

// Diff returns the sequence of operations to transform oldSlice into newSlice.
// It requires type T to be comparable (e.g., pointers, ints, strings).
func Diff[T1, T2 PointerToGongstruct](stage *Stage, a, b T1, fieldName string, oldSlice, newSlice []T2) (ops string) {
	m, n := len(oldSlice), len(newSlice)

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if oldSlice[i] == newSlice[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
				} else {
					dp[i+1][j+1] = dp[i+1][j]
				}
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if oldSlice[i-1] == newSlice[j-1] {
			keptIndices[i-1] = true
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// 3. PHASE 1: Generate Deletions
	// MUST go from High Index -> Low Index to preserve validity of lower indices.
	for k := m - 1; k >= 0; k-- {
		if !keptIndices[k] {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetReferenceIdentifier(stage), fieldName, a.GongGetReferenceIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Create a temporary view of what's left after deletions for tracking matches
	var currentLCS []T2
	for k := 0; k < m; k++ {
		if keptIndices[k] {
			currentLCS = append(currentLCS, oldSlice[k])
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k, targetVal := range newSlice {
		if lcsIdx < len(currentLCS) && currentLCS[lcsIdx] == targetVal {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, targetVal.GongGetIdentifier(stage))
		}
	}

	return ops
}
