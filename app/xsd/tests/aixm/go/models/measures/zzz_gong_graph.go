// generated code - do not edit
package measures

import (
	"fmt"
	"slices"
)

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (angletype *AngleType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AngleTypes[angletype]
	return ok
}

func (areatype *AreaType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AreaTypes[areatype]
	return ok
}

func (gridlengthtype *GridLengthType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GridLengthTypes[gridlengthtype]
	return ok
}

func (lengthtype *LengthType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.LengthTypes[lengthtype]
	return ok
}

func (scaletype *ScaleType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ScaleTypes[scaletype]
	return ok
}

func (speedtype *SpeedType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.SpeedTypes[speedtype]
	return ok
}

func (timetype *TimeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TimeTypes[timetype]
	return ok
}

func (volumetype *VolumeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.VolumeTypes[volumetype]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (angletype *AngleType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(angletype) {
		return
	}

	angletype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (areatype *AreaType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(areatype) {
		return
	}

	areatype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (gridlengthtype *GridLengthType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(gridlengthtype) {
		return
	}

	gridlengthtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (lengthtype *LengthType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(lengthtype) {
		return
	}

	lengthtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (scaletype *ScaleType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(scaletype) {
		return
	}

	scaletype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (speedtype *SpeedType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(speedtype) {
		return
	}

	speedtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (timetype *TimeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(timetype) {
		return
	}

	timetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (volumetype *VolumeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(volumetype) {
		return
	}

	volumetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

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
	case *AngleType:
		toT := GongCopyBranchAngleType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AreaType:
		toT := GongCopyBranchAreaType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GridLengthType:
		toT := GongCopyBranchGridLengthType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LengthType:
		toT := GongCopyBranchLengthType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ScaleType:
		toT := GongCopyBranchScaleType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SpeedType:
		toT := GongCopyBranchSpeedType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TimeType:
		toT := GongCopyBranchTimeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *VolumeType:
		toT := GongCopyBranchVolumeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchAngleType(mapOrigCopy map[any]any, angletypeFrom *AngleType) (angletypeTo *AngleType) {
	var alreadyCopied bool
	angletypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, angletypeFrom)
	if alreadyCopied {
		return
	}
	angletypeFrom.GongCopyBasicFields(angletypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAreaType(mapOrigCopy map[any]any, areatypeFrom *AreaType) (areatypeTo *AreaType) {
	var alreadyCopied bool
	areatypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, areatypeFrom)
	if alreadyCopied {
		return
	}
	areatypeFrom.GongCopyBasicFields(areatypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGridLengthType(mapOrigCopy map[any]any, gridlengthtypeFrom *GridLengthType) (gridlengthtypeTo *GridLengthType) {
	var alreadyCopied bool
	gridlengthtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, gridlengthtypeFrom)
	if alreadyCopied {
		return
	}
	gridlengthtypeFrom.GongCopyBasicFields(gridlengthtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLengthType(mapOrigCopy map[any]any, lengthtypeFrom *LengthType) (lengthtypeTo *LengthType) {
	var alreadyCopied bool
	lengthtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, lengthtypeFrom)
	if alreadyCopied {
		return
	}
	lengthtypeFrom.GongCopyBasicFields(lengthtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchScaleType(mapOrigCopy map[any]any, scaletypeFrom *ScaleType) (scaletypeTo *ScaleType) {
	var alreadyCopied bool
	scaletypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, scaletypeFrom)
	if alreadyCopied {
		return
	}
	scaletypeFrom.GongCopyBasicFields(scaletypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSpeedType(mapOrigCopy map[any]any, speedtypeFrom *SpeedType) (speedtypeTo *SpeedType) {
	var alreadyCopied bool
	speedtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, speedtypeFrom)
	if alreadyCopied {
		return
	}
	speedtypeFrom.GongCopyBasicFields(speedtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTimeType(mapOrigCopy map[any]any, timetypeFrom *TimeType) (timetypeTo *TimeType) {
	var alreadyCopied bool
	timetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, timetypeFrom)
	if alreadyCopied {
		return
	}
	timetypeFrom.GongCopyBasicFields(timetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchVolumeType(mapOrigCopy map[any]any, volumetypeFrom *VolumeType) (volumetypeTo *VolumeType) {
	var alreadyCopied bool
	volumetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, volumetypeFrom)
	if alreadyCopied {
		return
	}
	volumetypeFrom.GongCopyBasicFields(volumetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

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

// insertion point for unstage branch per struct
func (angletype *AngleType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(angletype) {
		return
	}

	angletype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (areatype *AreaType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(areatype) {
		return
	}

	areatype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (gridlengthtype *GridLengthType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(gridlengthtype) {
		return
	}

	gridlengthtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (lengthtype *LengthType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(lengthtype) {
		return
	}

	lengthtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (scaletype *ScaleType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(scaletype) {
		return
	}

	scaletype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (speedtype *SpeedType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(speedtype) {
		return
	}

	speedtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (timetype *TimeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(timetype) {
		return
	}

	timetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (volumetype *VolumeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(volumetype) {
		return
	}

	volumetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *AngleType) GongReconstructPointersFromReferences(stage *Stage, instance *AngleType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *AreaType) GongReconstructPointersFromReferences(stage *Stage, instance *AreaType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *GridLengthType) GongReconstructPointersFromReferences(stage *Stage, instance *GridLengthType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *LengthType) GongReconstructPointersFromReferences(stage *Stage, instance *LengthType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ScaleType) GongReconstructPointersFromReferences(stage *Stage, instance *ScaleType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *SpeedType) GongReconstructPointersFromReferences(stage *Stage, instance *SpeedType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TimeType) GongReconstructPointersFromReferences(stage *Stage, instance *TimeType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *VolumeType) GongReconstructPointersFromReferences(stage *Stage, instance *VolumeType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *AngleType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *AreaType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *GridLengthType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *LengthType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ScaleType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *SpeedType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TimeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *VolumeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (angletype *AngleType) GongDiff(stage *Stage, angletypeOther *AngleType) (diffs []string) {
	// insertion point for field diffs
	if angletype.Name != angletypeOther.Name {
		diffs = append(diffs, angletype.GongMarshallField(stage, "Name"))
	}
	if angletype.EnclosedText != angletypeOther.EnclosedText {
		diffs = append(diffs, angletype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (areatype *AreaType) GongDiff(stage *Stage, areatypeOther *AreaType) (diffs []string) {
	// insertion point for field diffs
	if areatype.Name != areatypeOther.Name {
		diffs = append(diffs, areatype.GongMarshallField(stage, "Name"))
	}
	if areatype.EnclosedText != areatypeOther.EnclosedText {
		diffs = append(diffs, areatype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (gridlengthtype *GridLengthType) GongDiff(stage *Stage, gridlengthtypeOther *GridLengthType) (diffs []string) {
	// insertion point for field diffs
	if gridlengthtype.Name != gridlengthtypeOther.Name {
		diffs = append(diffs, gridlengthtype.GongMarshallField(stage, "Name"))
	}
	if gridlengthtype.EnclosedText != gridlengthtypeOther.EnclosedText {
		diffs = append(diffs, gridlengthtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (lengthtype *LengthType) GongDiff(stage *Stage, lengthtypeOther *LengthType) (diffs []string) {
	// insertion point for field diffs
	if lengthtype.Name != lengthtypeOther.Name {
		diffs = append(diffs, lengthtype.GongMarshallField(stage, "Name"))
	}
	if lengthtype.EnclosedText != lengthtypeOther.EnclosedText {
		diffs = append(diffs, lengthtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (scaletype *ScaleType) GongDiff(stage *Stage, scaletypeOther *ScaleType) (diffs []string) {
	// insertion point for field diffs
	if scaletype.Name != scaletypeOther.Name {
		diffs = append(diffs, scaletype.GongMarshallField(stage, "Name"))
	}
	if scaletype.EnclosedText != scaletypeOther.EnclosedText {
		diffs = append(diffs, scaletype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (speedtype *SpeedType) GongDiff(stage *Stage, speedtypeOther *SpeedType) (diffs []string) {
	// insertion point for field diffs
	if speedtype.Name != speedtypeOther.Name {
		diffs = append(diffs, speedtype.GongMarshallField(stage, "Name"))
	}
	if speedtype.EnclosedText != speedtypeOther.EnclosedText {
		diffs = append(diffs, speedtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (timetype *TimeType) GongDiff(stage *Stage, timetypeOther *TimeType) (diffs []string) {
	// insertion point for field diffs
	if timetype.Name != timetypeOther.Name {
		diffs = append(diffs, timetype.GongMarshallField(stage, "Name"))
	}
	if timetype.EnclosedText != timetypeOther.EnclosedText {
		diffs = append(diffs, timetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (volumetype *VolumeType) GongDiff(stage *Stage, volumetypeOther *VolumeType) (diffs []string) {
	// insertion point for field diffs
	if volumetype.Name != volumetypeOther.Name {
		diffs = append(diffs, volumetype.GongMarshallField(stage, "Name"))
	}
	if volumetype.EnclosedText != volumetypeOther.EnclosedText {
		diffs = append(diffs, volumetype.GongMarshallField(stage, "EnclosedText"))
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

func __gong__copyBranchCheck[T any](mapOrigCopy map[any]any, from *T) (*T, bool) {
	if to, ok := mapOrigCopy[from]; ok {
		return to.(*T), true
	}
	to := new(T)
	mapOrigCopy[from] = to
	return to, false
}

func __gong__reconstructPointer[T comparable](field *T, refMap map[T]T, instanceField T) {
	var zero T
	if instanceField != zero {
		*field = refMap[instanceField]
	}
}

func __gong__reconstructPointerFromInstance[T comparable](field *T, instMap map[T]T) {
	ref := *field
	var zero T
	if ref != zero {
		*field = zero
		if inst, ok := instMap[ref]; ok {
			*field = inst
		}
	}
}

func __gong__reconstructSliceOfPointersFromReferences[T comparable](field *[]T, refMap map[T]T, instanceSlice []T) {
	*field = (*field)[:0]
	for _, b := range instanceSlice {
		*field = append(*field, refMap[b])
	}
}

func __gong__reconstructSliceOfPointersFromInstances[T comparable](field *[]T, instMap map[T]T) {
	var res []T
	for _, ref := range *field {
		if inst, ok := instMap[ref]; ok {
			res = append(res, inst)
		}
	}
	*field = res
}

func __gong__diffSliceOfPointers[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	instance GongstructIF,
	fieldName string,
	oldSlice, newSlice []T,
) string {
	if slices.Equal(oldSlice, newSlice) {
		return ""
	}
	return stage.Diff(
		instance,
		fieldName,
		len(oldSlice),
		len(newSlice),
		func(i, j int) bool {
			return oldSlice[i] == newSlice[j]
		},
		func(j int) string {
			return newSlice[j].GongGetIdentifier(stage)
		},
	)
}
