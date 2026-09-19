// generated code - do not edit
package models

import "fmt"

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (assplit *AsSplit) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.AsSplits[assplit]

	return
}

func (stage *Stage) IsStagedAsSplit(assplit *AsSplit) (ok bool) {

	return assplit.GongIsStaged(stage)
}

func (assplitarea *AsSplitArea) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.AsSplitAreas[assplitarea]

	return
}

func (stage *Stage) IsStagedAsSplitArea(assplitarea *AsSplitArea) (ok bool) {

	return assplitarea.GongIsStaged(stage)
}

func (button *Button) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Buttons[button]

	return
}

func (stage *Stage) IsStagedButton(button *Button) (ok bool) {

	return button.GongIsStaged(stage)
}

func (favicon *FavIcon) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.FavIcons[favicon]

	return
}

func (stage *Stage) IsStagedFavIcon(favicon *FavIcon) (ok bool) {

	return favicon.GongIsStaged(stage)
}

func (form *Form) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Forms[form]

	return
}

func (stage *Stage) IsStagedForm(form *Form) (ok bool) {

	return form.GongIsStaged(stage)
}

func (load *Load) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Loads[load]

	return
}

func (stage *Stage) IsStagedLoad(load *Load) (ok bool) {

	return load.GongIsStaged(stage)
}

func (logoontheleft *LogoOnTheLeft) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.LogoOnTheLefts[logoontheleft]

	return
}

func (stage *Stage) IsStagedLogoOnTheLeft(logoontheleft *LogoOnTheLeft) (ok bool) {

	return logoontheleft.GongIsStaged(stage)
}

func (logoontheright *LogoOnTheRight) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.LogoOnTheRights[logoontheright]

	return
}

func (stage *Stage) IsStagedLogoOnTheRight(logoontheright *LogoOnTheRight) (ok bool) {

	return logoontheright.GongIsStaged(stage)
}

func (split *Split) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Splits[split]

	return
}

func (stage *Stage) IsStagedSplit(split *Split) (ok bool) {

	return split.GongIsStaged(stage)
}

func (svg *Svg) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Svgs[svg]

	return
}

func (stage *Stage) IsStagedSvg(svg *Svg) (ok bool) {

	return svg.GongIsStaged(stage)
}

func (table *Table) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Tables[table]

	return
}

func (stage *Stage) IsStagedTable(table *Table) (ok bool) {

	return table.GongIsStaged(stage)
}

func (title *Title) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Titles[title]

	return
}

func (stage *Stage) IsStagedTitle(title *Title) (ok bool) {

	return title.GongIsStaged(stage)
}

func (tree *Tree) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Trees[tree]

	return
}

func (stage *Stage) IsStagedTree(tree *Tree) (ok bool) {

	return tree.GongIsStaged(stage)
}

func (view *View) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Views[view]

	return
}

func (stage *Stage) IsStagedView(view *View) (ok bool) {

	return view.GongIsStaged(stage)
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// StageBranch is a backward-compatible package-level forwarder.
func StageBranch(stage *Stage, instance GongstructIF) {
	stage.StageBranch(instance)
}

// insertion point for stage branch per struct
func (assplit *AsSplit) GongStageBranch(stage *Stage) {
	stage.StageBranchAsSplit(assplit)
}

func (stage *Stage) StageBranchAsSplit(assplit *AsSplit) {

	// check if instance is already staged
	if stage.IsStaged(assplit) {
		return
	}

	assplit.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assplitarea := range assplit.AsSplitAreas {
		stage.StageBranch(_assplitarea)
	}

}

func (assplitarea *AsSplitArea) GongStageBranch(stage *Stage) {
	stage.StageBranchAsSplitArea(assplitarea)
}

func (stage *Stage) StageBranchAsSplitArea(assplitarea *AsSplitArea) {

	// check if instance is already staged
	if stage.IsStaged(assplitarea) {
		return
	}

	assplitarea.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if assplitarea.AsSplit != nil {
		stage.StageBranch(assplitarea.AsSplit)
	}
	if assplitarea.Button != nil {
		stage.StageBranch(assplitarea.Button)
	}
	if assplitarea.Form != nil {
		stage.StageBranch(assplitarea.Form)
	}
	if assplitarea.Load != nil {
		stage.StageBranch(assplitarea.Load)
	}
	if assplitarea.Split != nil {
		stage.StageBranch(assplitarea.Split)
	}
	if assplitarea.Svg != nil {
		stage.StageBranch(assplitarea.Svg)
	}
	if assplitarea.Table != nil {
		stage.StageBranch(assplitarea.Table)
	}
	if assplitarea.Tree != nil {
		stage.StageBranch(assplitarea.Tree)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (button *Button) GongStageBranch(stage *Stage) {
	stage.StageBranchButton(button)
}

func (stage *Stage) StageBranchButton(button *Button) {

	// check if instance is already staged
	if stage.IsStaged(button) {
		return
	}

	button.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (favicon *FavIcon) GongStageBranch(stage *Stage) {
	stage.StageBranchFavIcon(favicon)
}

func (stage *Stage) StageBranchFavIcon(favicon *FavIcon) {

	// check if instance is already staged
	if stage.IsStaged(favicon) {
		return
	}

	favicon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (form *Form) GongStageBranch(stage *Stage) {
	stage.StageBranchForm(form)
}

func (stage *Stage) StageBranchForm(form *Form) {

	// check if instance is already staged
	if stage.IsStaged(form) {
		return
	}

	form.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (load *Load) GongStageBranch(stage *Stage) {
	stage.StageBranchLoad(load)
}

func (stage *Stage) StageBranchLoad(load *Load) {

	// check if instance is already staged
	if stage.IsStaged(load) {
		return
	}

	load.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (logoontheleft *LogoOnTheLeft) GongStageBranch(stage *Stage) {
	stage.StageBranchLogoOnTheLeft(logoontheleft)
}

func (stage *Stage) StageBranchLogoOnTheLeft(logoontheleft *LogoOnTheLeft) {

	// check if instance is already staged
	if stage.IsStaged(logoontheleft) {
		return
	}

	logoontheleft.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (logoontheright *LogoOnTheRight) GongStageBranch(stage *Stage) {
	stage.StageBranchLogoOnTheRight(logoontheright)
}

func (stage *Stage) StageBranchLogoOnTheRight(logoontheright *LogoOnTheRight) {

	// check if instance is already staged
	if stage.IsStaged(logoontheright) {
		return
	}

	logoontheright.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (split *Split) GongStageBranch(stage *Stage) {
	stage.StageBranchSplit(split)
}

func (stage *Stage) StageBranchSplit(split *Split) {

	// check if instance is already staged
	if stage.IsStaged(split) {
		return
	}

	split.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (svg *Svg) GongStageBranch(stage *Stage) {
	stage.StageBranchSvg(svg)
}

func (stage *Stage) StageBranchSvg(svg *Svg) {

	// check if instance is already staged
	if stage.IsStaged(svg) {
		return
	}

	svg.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (table *Table) GongStageBranch(stage *Stage) {
	stage.StageBranchTable(table)
}

func (stage *Stage) StageBranchTable(table *Table) {

	// check if instance is already staged
	if stage.IsStaged(table) {
		return
	}

	table.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (title *Title) GongStageBranch(stage *Stage) {
	stage.StageBranchTitle(title)
}

func (stage *Stage) StageBranchTitle(title *Title) {

	// check if instance is already staged
	if stage.IsStaged(title) {
		return
	}

	title.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tree *Tree) GongStageBranch(stage *Stage) {
	stage.StageBranchTree(tree)
}

func (stage *Stage) StageBranchTree(tree *Tree) {

	// check if instance is already staged
	if stage.IsStaged(tree) {
		return
	}

	tree.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (view *View) GongStageBranch(stage *Stage) {
	stage.StageBranchView(view)
}

func (stage *Stage) StageBranchView(view *View) {

	// check if instance is already staged
	if stage.IsStaged(view) {
		return
	}

	view.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assplitarea := range view.RootAsSplitAreas {
		stage.StageBranch(_assplitarea)
	}

}

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *AsSplit:
		toT := GongCopyBranchAsSplit(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AsSplitArea:
		toT := GongCopyBranchAsSplitArea(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Button:
		toT := GongCopyBranchButton(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FavIcon:
		toT := GongCopyBranchFavIcon(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Form:
		toT := GongCopyBranchForm(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Load:
		toT := GongCopyBranchLoad(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LogoOnTheLeft:
		toT := GongCopyBranchLogoOnTheLeft(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LogoOnTheRight:
		toT := GongCopyBranchLogoOnTheRight(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Split:
		toT := GongCopyBranchSplit(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Svg:
		toT := GongCopyBranchSvg(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Table:
		toT := GongCopyBranchTable(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Title:
		toT := GongCopyBranchTitle(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tree:
		toT := GongCopyBranchTree(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *View:
		toT := GongCopyBranchView(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchAsSplit(mapOrigCopy map[any]any, assplitFrom *AsSplit) (assplitTo *AsSplit) {

	// assplitFrom has already been copied
	if _assplitTo, ok := mapOrigCopy[assplitFrom]; ok {
		assplitTo = _assplitTo.(*AsSplit)
		return
	}

	assplitTo = new(AsSplit)
	mapOrigCopy[assplitFrom] = assplitTo
	assplitFrom.GongCopyBasicFields(assplitTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assplitarea := range assplitFrom.AsSplitAreas {
		assplitTo.AsSplitAreas = append(assplitTo.AsSplitAreas, GongCopyBranchAsSplitArea(mapOrigCopy, _assplitarea))
	}

	return
}

func GongCopyBranchAsSplitArea(mapOrigCopy map[any]any, assplitareaFrom *AsSplitArea) (assplitareaTo *AsSplitArea) {

	// assplitareaFrom has already been copied
	if _assplitareaTo, ok := mapOrigCopy[assplitareaFrom]; ok {
		assplitareaTo = _assplitareaTo.(*AsSplitArea)
		return
	}

	assplitareaTo = new(AsSplitArea)
	mapOrigCopy[assplitareaFrom] = assplitareaTo
	assplitareaFrom.GongCopyBasicFields(assplitareaTo)

	//insertion point for the staging of instances referenced by pointers
	if assplitareaFrom.AsSplit != nil {
		assplitareaTo.AsSplit = GongCopyBranchAsSplit(mapOrigCopy, assplitareaFrom.AsSplit)
	}
	if assplitareaFrom.Button != nil {
		assplitareaTo.Button = GongCopyBranchButton(mapOrigCopy, assplitareaFrom.Button)
	}
	if assplitareaFrom.Form != nil {
		assplitareaTo.Form = GongCopyBranchForm(mapOrigCopy, assplitareaFrom.Form)
	}
	if assplitareaFrom.Load != nil {
		assplitareaTo.Load = GongCopyBranchLoad(mapOrigCopy, assplitareaFrom.Load)
	}
	if assplitareaFrom.Split != nil {
		assplitareaTo.Split = GongCopyBranchSplit(mapOrigCopy, assplitareaFrom.Split)
	}
	if assplitareaFrom.Svg != nil {
		assplitareaTo.Svg = GongCopyBranchSvg(mapOrigCopy, assplitareaFrom.Svg)
	}
	if assplitareaFrom.Table != nil {
		assplitareaTo.Table = GongCopyBranchTable(mapOrigCopy, assplitareaFrom.Table)
	}
	if assplitareaFrom.Tree != nil {
		assplitareaTo.Tree = GongCopyBranchTree(mapOrigCopy, assplitareaFrom.Tree)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchButton(mapOrigCopy map[any]any, buttonFrom *Button) (buttonTo *Button) {

	// buttonFrom has already been copied
	if _buttonTo, ok := mapOrigCopy[buttonFrom]; ok {
		buttonTo = _buttonTo.(*Button)
		return
	}

	buttonTo = new(Button)
	mapOrigCopy[buttonFrom] = buttonTo
	buttonFrom.GongCopyBasicFields(buttonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFavIcon(mapOrigCopy map[any]any, faviconFrom *FavIcon) (faviconTo *FavIcon) {

	// faviconFrom has already been copied
	if _faviconTo, ok := mapOrigCopy[faviconFrom]; ok {
		faviconTo = _faviconTo.(*FavIcon)
		return
	}

	faviconTo = new(FavIcon)
	mapOrigCopy[faviconFrom] = faviconTo
	faviconFrom.GongCopyBasicFields(faviconTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchForm(mapOrigCopy map[any]any, formFrom *Form) (formTo *Form) {

	// formFrom has already been copied
	if _formTo, ok := mapOrigCopy[formFrom]; ok {
		formTo = _formTo.(*Form)
		return
	}

	formTo = new(Form)
	mapOrigCopy[formFrom] = formTo
	formFrom.GongCopyBasicFields(formTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLoad(mapOrigCopy map[any]any, loadFrom *Load) (loadTo *Load) {

	// loadFrom has already been copied
	if _loadTo, ok := mapOrigCopy[loadFrom]; ok {
		loadTo = _loadTo.(*Load)
		return
	}

	loadTo = new(Load)
	mapOrigCopy[loadFrom] = loadTo
	loadFrom.GongCopyBasicFields(loadTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLogoOnTheLeft(mapOrigCopy map[any]any, logoontheleftFrom *LogoOnTheLeft) (logoontheleftTo *LogoOnTheLeft) {

	// logoontheleftFrom has already been copied
	if _logoontheleftTo, ok := mapOrigCopy[logoontheleftFrom]; ok {
		logoontheleftTo = _logoontheleftTo.(*LogoOnTheLeft)
		return
	}

	logoontheleftTo = new(LogoOnTheLeft)
	mapOrigCopy[logoontheleftFrom] = logoontheleftTo
	logoontheleftFrom.GongCopyBasicFields(logoontheleftTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLogoOnTheRight(mapOrigCopy map[any]any, logoontherightFrom *LogoOnTheRight) (logoontherightTo *LogoOnTheRight) {

	// logoontherightFrom has already been copied
	if _logoontherightTo, ok := mapOrigCopy[logoontherightFrom]; ok {
		logoontherightTo = _logoontherightTo.(*LogoOnTheRight)
		return
	}

	logoontherightTo = new(LogoOnTheRight)
	mapOrigCopy[logoontherightFrom] = logoontherightTo
	logoontherightFrom.GongCopyBasicFields(logoontherightTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSplit(mapOrigCopy map[any]any, splitFrom *Split) (splitTo *Split) {

	// splitFrom has already been copied
	if _splitTo, ok := mapOrigCopy[splitFrom]; ok {
		splitTo = _splitTo.(*Split)
		return
	}

	splitTo = new(Split)
	mapOrigCopy[splitFrom] = splitTo
	splitFrom.GongCopyBasicFields(splitTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSvg(mapOrigCopy map[any]any, svgFrom *Svg) (svgTo *Svg) {

	// svgFrom has already been copied
	if _svgTo, ok := mapOrigCopy[svgFrom]; ok {
		svgTo = _svgTo.(*Svg)
		return
	}

	svgTo = new(Svg)
	mapOrigCopy[svgFrom] = svgTo
	svgFrom.GongCopyBasicFields(svgTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTable(mapOrigCopy map[any]any, tableFrom *Table) (tableTo *Table) {

	// tableFrom has already been copied
	if _tableTo, ok := mapOrigCopy[tableFrom]; ok {
		tableTo = _tableTo.(*Table)
		return
	}

	tableTo = new(Table)
	mapOrigCopy[tableFrom] = tableTo
	tableFrom.GongCopyBasicFields(tableTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTitle(mapOrigCopy map[any]any, titleFrom *Title) (titleTo *Title) {

	// titleFrom has already been copied
	if _titleTo, ok := mapOrigCopy[titleFrom]; ok {
		titleTo = _titleTo.(*Title)
		return
	}

	titleTo = new(Title)
	mapOrigCopy[titleFrom] = titleTo
	titleFrom.GongCopyBasicFields(titleTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTree(mapOrigCopy map[any]any, treeFrom *Tree) (treeTo *Tree) {

	// treeFrom has already been copied
	if _treeTo, ok := mapOrigCopy[treeFrom]; ok {
		treeTo = _treeTo.(*Tree)
		return
	}

	treeTo = new(Tree)
	mapOrigCopy[treeFrom] = treeTo
	treeFrom.GongCopyBasicFields(treeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchView(mapOrigCopy map[any]any, viewFrom *View) (viewTo *View) {

	// viewFrom has already been copied
	if _viewTo, ok := mapOrigCopy[viewFrom]; ok {
		viewTo = _viewTo.(*View)
		return
	}

	viewTo = new(View)
	mapOrigCopy[viewFrom] = viewTo
	viewFrom.GongCopyBasicFields(viewTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assplitarea := range viewFrom.RootAsSplitAreas {
		viewTo.RootAsSplitAreas = append(viewTo.RootAsSplitAreas, GongCopyBranchAsSplitArea(mapOrigCopy, _assplitarea))
	}

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
// UnstageBranch is the Stage method that unstages instance and applies UnstageBranch recursively.
func (stage *Stage) UnstageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongUnstageBranch(stage)
	}
}

// UnstageBranch is a backward-compatible package-level forwarder.
func UnstageBranch(stage *Stage, instance GongstructIF) {
	stage.UnstageBranch(instance)
}

// insertion point for unstage branch per struct
func (assplit *AsSplit) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchAsSplit(assplit)
}

func (stage *Stage) UnstageBranchAsSplit(assplit *AsSplit) {

	// check if instance is already staged
	if !stage.IsStaged(assplit) {
		return
	}

	assplit.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assplitarea := range assplit.AsSplitAreas {
		stage.UnstageBranch(_assplitarea)
	}

}

func (assplitarea *AsSplitArea) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchAsSplitArea(assplitarea)
}

func (stage *Stage) UnstageBranchAsSplitArea(assplitarea *AsSplitArea) {

	// check if instance is already staged
	if !stage.IsStaged(assplitarea) {
		return
	}

	assplitarea.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if assplitarea.AsSplit != nil {
		stage.UnstageBranch(assplitarea.AsSplit)
	}
	if assplitarea.Button != nil {
		stage.UnstageBranch(assplitarea.Button)
	}
	if assplitarea.Form != nil {
		stage.UnstageBranch(assplitarea.Form)
	}
	if assplitarea.Load != nil {
		stage.UnstageBranch(assplitarea.Load)
	}
	if assplitarea.Split != nil {
		stage.UnstageBranch(assplitarea.Split)
	}
	if assplitarea.Svg != nil {
		stage.UnstageBranch(assplitarea.Svg)
	}
	if assplitarea.Table != nil {
		stage.UnstageBranch(assplitarea.Table)
	}
	if assplitarea.Tree != nil {
		stage.UnstageBranch(assplitarea.Tree)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (button *Button) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchButton(button)
}

func (stage *Stage) UnstageBranchButton(button *Button) {

	// check if instance is already staged
	if !stage.IsStaged(button) {
		return
	}

	button.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (favicon *FavIcon) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchFavIcon(favicon)
}

func (stage *Stage) UnstageBranchFavIcon(favicon *FavIcon) {

	// check if instance is already staged
	if !stage.IsStaged(favicon) {
		return
	}

	favicon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (form *Form) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchForm(form)
}

func (stage *Stage) UnstageBranchForm(form *Form) {

	// check if instance is already staged
	if !stage.IsStaged(form) {
		return
	}

	form.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (load *Load) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchLoad(load)
}

func (stage *Stage) UnstageBranchLoad(load *Load) {

	// check if instance is already staged
	if !stage.IsStaged(load) {
		return
	}

	load.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (logoontheleft *LogoOnTheLeft) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchLogoOnTheLeft(logoontheleft)
}

func (stage *Stage) UnstageBranchLogoOnTheLeft(logoontheleft *LogoOnTheLeft) {

	// check if instance is already staged
	if !stage.IsStaged(logoontheleft) {
		return
	}

	logoontheleft.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (logoontheright *LogoOnTheRight) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchLogoOnTheRight(logoontheright)
}

func (stage *Stage) UnstageBranchLogoOnTheRight(logoontheright *LogoOnTheRight) {

	// check if instance is already staged
	if !stage.IsStaged(logoontheright) {
		return
	}

	logoontheright.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (split *Split) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchSplit(split)
}

func (stage *Stage) UnstageBranchSplit(split *Split) {

	// check if instance is already staged
	if !stage.IsStaged(split) {
		return
	}

	split.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (svg *Svg) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchSvg(svg)
}

func (stage *Stage) UnstageBranchSvg(svg *Svg) {

	// check if instance is already staged
	if !stage.IsStaged(svg) {
		return
	}

	svg.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (table *Table) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTable(table)
}

func (stage *Stage) UnstageBranchTable(table *Table) {

	// check if instance is already staged
	if !stage.IsStaged(table) {
		return
	}

	table.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (title *Title) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTitle(title)
}

func (stage *Stage) UnstageBranchTitle(title *Title) {

	// check if instance is already staged
	if !stage.IsStaged(title) {
		return
	}

	title.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tree *Tree) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTree(tree)
}

func (stage *Stage) UnstageBranchTree(tree *Tree) {

	// check if instance is already staged
	if !stage.IsStaged(tree) {
		return
	}

	tree.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (view *View) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchView(view)
}

func (stage *Stage) UnstageBranchView(view *View) {

	// check if instance is already staged
	if !stage.IsStaged(view) {
		return
	}

	view.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assplitarea := range view.RootAsSplitAreas {
		stage.UnstageBranch(_assplitarea)
	}

}

// insertion point for pointer reconstruction from references
func (reference *AsSplit) GongReconstructPointersFromReferences(stage *Stage, instance *AsSplit) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.AsSplitAreas = reference.AsSplitAreas[:0]
	for _, _b := range instance.AsSplitAreas {
		reference.AsSplitAreas = append(reference.AsSplitAreas, stage.AsSplitAreas_reference[_b])
	}
}

func (reference *AsSplitArea) GongReconstructPointersFromReferences(stage *Stage, instance *AsSplitArea) {
	// insertion point for pointers field
	if instance.AsSplit != nil {
		reference.AsSplit = stage.AsSplits_reference[instance.AsSplit]
	}
	if instance.Button != nil {
		reference.Button = stage.Buttons_reference[instance.Button]
	}
	if instance.Form != nil {
		reference.Form = stage.Forms_reference[instance.Form]
	}
	if instance.Load != nil {
		reference.Load = stage.Loads_reference[instance.Load]
	}
	if instance.Split != nil {
		reference.Split = stage.Splits_reference[instance.Split]
	}
	if instance.Svg != nil {
		reference.Svg = stage.Svgs_reference[instance.Svg]
	}
	if instance.Table != nil {
		reference.Table = stage.Tables_reference[instance.Table]
	}
	if instance.Tree != nil {
		reference.Tree = stage.Trees_reference[instance.Tree]
	}
	// insertion point for slice of pointers field
}

func (reference *Button) GongReconstructPointersFromReferences(stage *Stage, instance *Button) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *FavIcon) GongReconstructPointersFromReferences(stage *Stage, instance *FavIcon) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Form) GongReconstructPointersFromReferences(stage *Stage, instance *Form) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Load) GongReconstructPointersFromReferences(stage *Stage, instance *Load) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *LogoOnTheLeft) GongReconstructPointersFromReferences(stage *Stage, instance *LogoOnTheLeft) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *LogoOnTheRight) GongReconstructPointersFromReferences(stage *Stage, instance *LogoOnTheRight) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Split) GongReconstructPointersFromReferences(stage *Stage, instance *Split) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Svg) GongReconstructPointersFromReferences(stage *Stage, instance *Svg) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Table) GongReconstructPointersFromReferences(stage *Stage, instance *Table) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Title) GongReconstructPointersFromReferences(stage *Stage, instance *Title) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Tree) GongReconstructPointersFromReferences(stage *Stage, instance *Tree) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *View) GongReconstructPointersFromReferences(stage *Stage, instance *View) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.RootAsSplitAreas = reference.RootAsSplitAreas[:0]
	for _, _b := range instance.RootAsSplitAreas {
		reference.RootAsSplitAreas = append(reference.RootAsSplitAreas, stage.AsSplitAreas_reference[_b])
	}
}

// insertion point for pointer reconstruction from instances
func (reference *AsSplit) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _AsSplitAreas []*AsSplitArea
	for _, _reference := range reference.AsSplitAreas {
		if _instance, ok := stage.AsSplitAreas_instance[_reference]; ok {
			_AsSplitAreas = append(_AsSplitAreas, _instance)
		}
	}
	reference.AsSplitAreas = _AsSplitAreas
}

func (reference *AsSplitArea) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.AsSplit; _reference != nil {
		reference.AsSplit = nil
		if _instance, ok := stage.AsSplits_instance[_reference]; ok {
			reference.AsSplit = _instance
		}
	}
	if _reference := reference.Button; _reference != nil {
		reference.Button = nil
		if _instance, ok := stage.Buttons_instance[_reference]; ok {
			reference.Button = _instance
		}
	}
	if _reference := reference.Form; _reference != nil {
		reference.Form = nil
		if _instance, ok := stage.Forms_instance[_reference]; ok {
			reference.Form = _instance
		}
	}
	if _reference := reference.Load; _reference != nil {
		reference.Load = nil
		if _instance, ok := stage.Loads_instance[_reference]; ok {
			reference.Load = _instance
		}
	}
	if _reference := reference.Split; _reference != nil {
		reference.Split = nil
		if _instance, ok := stage.Splits_instance[_reference]; ok {
			reference.Split = _instance
		}
	}
	if _reference := reference.Svg; _reference != nil {
		reference.Svg = nil
		if _instance, ok := stage.Svgs_instance[_reference]; ok {
			reference.Svg = _instance
		}
	}
	if _reference := reference.Table; _reference != nil {
		reference.Table = nil
		if _instance, ok := stage.Tables_instance[_reference]; ok {
			reference.Table = _instance
		}
	}
	if _reference := reference.Tree; _reference != nil {
		reference.Tree = nil
		if _instance, ok := stage.Trees_instance[_reference]; ok {
			reference.Tree = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Button) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *FavIcon) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Form) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Load) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *LogoOnTheLeft) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *LogoOnTheRight) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Split) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Svg) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Table) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Title) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Tree) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *View) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _RootAsSplitAreas []*AsSplitArea
	for _, _reference := range reference.RootAsSplitAreas {
		if _instance, ok := stage.AsSplitAreas_instance[_reference]; ok {
			_RootAsSplitAreas = append(_RootAsSplitAreas, _instance)
		}
	}
	reference.RootAsSplitAreas = _RootAsSplitAreas
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
		ops := stage.Diff(
			assplit,
			"AsSplitAreas",
			len(assplitOther.AsSplitAreas),
			len(assplit.AsSplitAreas),
			func(i, j int) bool {
				return assplitOther.AsSplitAreas[i] == assplit.AsSplitAreas[j]
			},
			func(j int) string {
				return assplit.AsSplitAreas[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if assplit.IsSizeInPixel != assplitOther.IsSizeInPixel {
		diffs = append(diffs, assplit.GongMarshallField(stage, "IsSizeInPixel"))
	}
	if assplit.IsWithCustomGutterSize != assplitOther.IsWithCustomGutterSize {
		diffs = append(diffs, assplit.GongMarshallField(stage, "IsWithCustomGutterSize"))
	}
	if assplit.GutterSize != assplitOther.GutterSize {
		diffs = append(diffs, assplit.GongMarshallField(stage, "GutterSize"))
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
	if (assplitarea.Tree == nil) != (assplitareaOther.Tree == nil) {
		diffs = append(diffs, assplitarea.GongMarshallField(stage, "Tree"))
	} else if assplitarea.Tree != nil && assplitareaOther.Tree != nil {
		if assplitarea.Tree != assplitareaOther.Tree {
			diffs = append(diffs, assplitarea.GongMarshallField(stage, "Tree"))
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
		ops := stage.Diff(
			view,
			"RootAsSplitAreas",
			len(viewOther.RootAsSplitAreas),
			len(view.RootAsSplitAreas),
			func(i, j int) bool {
				return viewOther.RootAsSplitAreas[i] == view.RootAsSplitAreas[j]
			},
			func(j int) string {
				return view.RootAsSplitAreas[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if view.IsSelectedView != viewOther.IsSelectedView {
		diffs = append(diffs, view.GongMarshallField(stage, "IsSelectedView"))
	}
	if view.Direction != viewOther.Direction {
		diffs = append(diffs, view.GongMarshallField(stage, "Direction"))
	}
	if view.IsSecondaryView != viewOther.IsSecondaryView {
		diffs = append(diffs, view.GongMarshallField(stage, "IsSecondaryView"))
	}
	if view.IsSizeInPixel != viewOther.IsSizeInPixel {
		diffs = append(diffs, view.GongMarshallField(stage, "IsSizeInPixel"))
	}
	if view.IsWithCustomGutterSize != viewOther.IsWithCustomGutterSize {
		diffs = append(diffs, view.GongMarshallField(stage, "IsWithCustomGutterSize"))
	}
	if view.GutterSize != viewOther.GutterSize {
		diffs = append(diffs, view.GongMarshallField(stage, "GutterSize"))
	}

	return
}

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff(
	a GongstructIF,
	fieldName string,
	lenOld, lenNew int,
	equal func(i, j int) bool,
	getNewIdentifier func(j int) string,
) (ops string) {
	m, n := lenOld, lenNew

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range m {
		for j := range n {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if equal(i-1, j-1) {
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

	// Track kept indices in old slice
	keptOldIndices := make([]int, 0, len(keptIndices))
	for k := range m {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := range n {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}
