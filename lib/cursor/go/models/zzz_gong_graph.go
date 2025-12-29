// generated code - do not edit
package models

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Cursor:
		ok = stage.IsStagedCursor(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Cursor:
		ok = stage.IsStagedCursor(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedCursor(cursor *Cursor) (ok bool) {

	_, ok = stage.Cursors[cursor]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Cursor:
		stage.StageBranchCursor(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchCursor(cursor *Cursor) {

	// check if instance is already staged
	if IsStaged(stage, cursor) {
		return
	}

	cursor.Stage(stage)

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
	case *Cursor:
		toT := CopyBranchCursor(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
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

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *Cursor:
		stage.UnstageBranchCursor(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchCursor(cursor *Cursor) {

	// check if instance is already staged
	if !IsStaged(stage, cursor) {
		return
	}

	cursor.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (cursor *Cursor) GongDiff(cursorOther *Cursor) (diffs []string) {
	// insertion point for field diffs
	if cursor.Name != cursorOther.Name {
		diffs = append(diffs, "Name")
	}
	if cursor.StartX != cursorOther.StartX {
		diffs = append(diffs, "StartX")
	}
	if cursor.EndX != cursorOther.EndX {
		diffs = append(diffs, "EndX")
	}
	if cursor.Y1 != cursorOther.Y1 {
		diffs = append(diffs, "Y1")
	}
	if cursor.Y2 != cursorOther.Y2 {
		diffs = append(diffs, "Y2")
	}
	if cursor.DurationSeconds != cursorOther.DurationSeconds {
		diffs = append(diffs, "DurationSeconds")
	}
	if cursor.Color != cursorOther.Color {
		diffs = append(diffs, "Color")
	}
	if cursor.FillOpacity != cursorOther.FillOpacity {
		diffs = append(diffs, "FillOpacity")
	}
	if cursor.Stroke != cursorOther.Stroke {
		diffs = append(diffs, "Stroke")
	}
	if cursor.StrokeOpacity != cursorOther.StrokeOpacity {
		diffs = append(diffs, "StrokeOpacity")
	}
	if cursor.StrokeWidth != cursorOther.StrokeWidth {
		diffs = append(diffs, "StrokeWidth")
	}
	if cursor.StrokeDashArray != cursorOther.StrokeDashArray {
		diffs = append(diffs, "StrokeDashArray")
	}
	if cursor.StrokeDashArrayWhenSelected != cursorOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, "StrokeDashArrayWhenSelected")
	}
	if cursor.Transform != cursorOther.Transform {
		diffs = append(diffs, "Transform")
	}
	if cursor.IsPlaying != cursorOther.IsPlaying {
		diffs = append(diffs, "IsPlaying")
	}

	return
}
