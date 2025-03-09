// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *AsSplit:
		ok = stage.IsStagedAsSplit(target)

	case *AsSplitArea:
		ok = stage.IsStagedAsSplitArea(target)

	case *Doc:
		ok = stage.IsStagedDoc(target)

	case *Form:
		ok = stage.IsStagedForm(target)

	case *Svg:
		ok = stage.IsStagedSvg(target)

	case *Table:
		ok = stage.IsStagedTable(target)

	case *Tree:
		ok = stage.IsStagedTree(target)

	case *View:
		ok = stage.IsStagedView(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *StageStruct) IsStagedAsSplit(assplit *AsSplit) (ok bool) {

	_, ok = stage.AsSplits[assplit]

	return
}

func (stage *StageStruct) IsStagedAsSplitArea(assplitarea *AsSplitArea) (ok bool) {

	_, ok = stage.AsSplitAreas[assplitarea]

	return
}

func (stage *StageStruct) IsStagedDoc(doc *Doc) (ok bool) {

	_, ok = stage.Docs[doc]

	return
}

func (stage *StageStruct) IsStagedForm(form *Form) (ok bool) {

	_, ok = stage.Forms[form]

	return
}

func (stage *StageStruct) IsStagedSvg(svg *Svg) (ok bool) {

	_, ok = stage.Svgs[svg]

	return
}

func (stage *StageStruct) IsStagedTable(table *Table) (ok bool) {

	_, ok = stage.Tables[table]

	return
}

func (stage *StageStruct) IsStagedTree(tree *Tree) (ok bool) {

	_, ok = stage.Trees[tree]

	return
}

func (stage *StageStruct) IsStagedView(view *View) (ok bool) {

	_, ok = stage.Views[view]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *AsSplit:
		stage.StageBranchAsSplit(target)

	case *AsSplitArea:
		stage.StageBranchAsSplitArea(target)

	case *Doc:
		stage.StageBranchDoc(target)

	case *Form:
		stage.StageBranchForm(target)

	case *Svg:
		stage.StageBranchSvg(target)

	case *Table:
		stage.StageBranchTable(target)

	case *Tree:
		stage.StageBranchTree(target)

	case *View:
		stage.StageBranchView(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchAsSplit(assplit *AsSplit) {

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

func (stage *StageStruct) StageBranchAsSplitArea(assplitarea *AsSplitArea) {

	// check if instance is already staged
	if IsStaged(stage, assplitarea) {
		return
	}

	assplitarea.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if assplitarea.Tree != nil {
		StageBranch(stage, assplitarea.Tree)
	}
	if assplitarea.Table != nil {
		StageBranch(stage, assplitarea.Table)
	}
	if assplitarea.Form != nil {
		StageBranch(stage, assplitarea.Form)
	}
	if assplitarea.Svg != nil {
		StageBranch(stage, assplitarea.Svg)
	}
	if assplitarea.Doc != nil {
		StageBranch(stage, assplitarea.Doc)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assplit := range assplitarea.AsSplits {
		StageBranch(stage, _assplit)
	}

}

func (stage *StageStruct) StageBranchDoc(doc *Doc) {

	// check if instance is already staged
	if IsStaged(stage, doc) {
		return
	}

	doc.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchForm(form *Form) {

	// check if instance is already staged
	if IsStaged(stage, form) {
		return
	}

	form.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSvg(svg *Svg) {

	// check if instance is already staged
	if IsStaged(stage, svg) {
		return
	}

	svg.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchTable(table *Table) {

	// check if instance is already staged
	if IsStaged(stage, table) {
		return
	}

	table.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchTree(tree *Tree) {

	// check if instance is already staged
	if IsStaged(stage, tree) {
		return
	}

	tree.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchView(view *View) {

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

	case *Doc:
		toT := CopyBranchDoc(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Form:
		toT := CopyBranchForm(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Svg:
		toT := CopyBranchSvg(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Table:
		toT := CopyBranchTable(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tree:
		toT := CopyBranchTree(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *View:
		toT := CopyBranchView(mapOrigCopy, fromT)
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
	if assplitareaFrom.Tree != nil {
		assplitareaTo.Tree = CopyBranchTree(mapOrigCopy, assplitareaFrom.Tree)
	}
	if assplitareaFrom.Table != nil {
		assplitareaTo.Table = CopyBranchTable(mapOrigCopy, assplitareaFrom.Table)
	}
	if assplitareaFrom.Form != nil {
		assplitareaTo.Form = CopyBranchForm(mapOrigCopy, assplitareaFrom.Form)
	}
	if assplitareaFrom.Svg != nil {
		assplitareaTo.Svg = CopyBranchSvg(mapOrigCopy, assplitareaFrom.Svg)
	}
	if assplitareaFrom.Doc != nil {
		assplitareaTo.Doc = CopyBranchDoc(mapOrigCopy, assplitareaFrom.Doc)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assplit := range assplitareaFrom.AsSplits {
		assplitareaTo.AsSplits = append(assplitareaTo.AsSplits, CopyBranchAsSplit(mapOrigCopy, _assplit))
	}

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

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *AsSplit:
		stage.UnstageBranchAsSplit(target)

	case *AsSplitArea:
		stage.UnstageBranchAsSplitArea(target)

	case *Doc:
		stage.UnstageBranchDoc(target)

	case *Form:
		stage.UnstageBranchForm(target)

	case *Svg:
		stage.UnstageBranchSvg(target)

	case *Table:
		stage.UnstageBranchTable(target)

	case *Tree:
		stage.UnstageBranchTree(target)

	case *View:
		stage.UnstageBranchView(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchAsSplit(assplit *AsSplit) {

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

func (stage *StageStruct) UnstageBranchAsSplitArea(assplitarea *AsSplitArea) {

	// check if instance is already staged
	if !IsStaged(stage, assplitarea) {
		return
	}

	assplitarea.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if assplitarea.Tree != nil {
		UnstageBranch(stage, assplitarea.Tree)
	}
	if assplitarea.Table != nil {
		UnstageBranch(stage, assplitarea.Table)
	}
	if assplitarea.Form != nil {
		UnstageBranch(stage, assplitarea.Form)
	}
	if assplitarea.Svg != nil {
		UnstageBranch(stage, assplitarea.Svg)
	}
	if assplitarea.Doc != nil {
		UnstageBranch(stage, assplitarea.Doc)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assplit := range assplitarea.AsSplits {
		UnstageBranch(stage, _assplit)
	}

}

func (stage *StageStruct) UnstageBranchDoc(doc *Doc) {

	// check if instance is already staged
	if !IsStaged(stage, doc) {
		return
	}

	doc.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchForm(form *Form) {

	// check if instance is already staged
	if !IsStaged(stage, form) {
		return
	}

	form.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSvg(svg *Svg) {

	// check if instance is already staged
	if !IsStaged(stage, svg) {
		return
	}

	svg.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchTable(table *Table) {

	// check if instance is already staged
	if !IsStaged(stage, table) {
		return
	}

	table.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchTree(tree *Tree) {

	// check if instance is already staged
	if !IsStaged(stage, tree) {
		return
	}

	tree.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchView(view *View) {

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
