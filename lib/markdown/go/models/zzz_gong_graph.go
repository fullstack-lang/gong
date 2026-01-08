// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Content:
		ok = stage.IsStagedContent(target)

	case *JpgImage:
		ok = stage.IsStagedJpgImage(target)

	case *PngImage:
		ok = stage.IsStagedPngImage(target)

	case *SvgImage:
		ok = stage.IsStagedSvgImage(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Content:
		ok = stage.IsStagedContent(target)

	case *JpgImage:
		ok = stage.IsStagedJpgImage(target)

	case *PngImage:
		ok = stage.IsStagedPngImage(target)

	case *SvgImage:
		ok = stage.IsStagedSvgImage(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedContent(content *Content) (ok bool) {

	_, ok = stage.Contents[content]

	return
}

func (stage *Stage) IsStagedJpgImage(jpgimage *JpgImage) (ok bool) {

	_, ok = stage.JpgImages[jpgimage]

	return
}

func (stage *Stage) IsStagedPngImage(pngimage *PngImage) (ok bool) {

	_, ok = stage.PngImages[pngimage]

	return
}

func (stage *Stage) IsStagedSvgImage(svgimage *SvgImage) (ok bool) {

	_, ok = stage.SvgImages[svgimage]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Content:
		stage.StageBranchContent(target)

	case *JpgImage:
		stage.StageBranchJpgImage(target)

	case *PngImage:
		stage.StageBranchPngImage(target)

	case *SvgImage:
		stage.StageBranchSvgImage(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchContent(content *Content) {

	// check if instance is already staged
	if IsStaged(stage, content) {
		return
	}

	content.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchJpgImage(jpgimage *JpgImage) {

	// check if instance is already staged
	if IsStaged(stage, jpgimage) {
		return
	}

	jpgimage.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPngImage(pngimage *PngImage) {

	// check if instance is already staged
	if IsStaged(stage, pngimage) {
		return
	}

	pngimage.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchSvgImage(svgimage *SvgImage) {

	// check if instance is already staged
	if IsStaged(stage, svgimage) {
		return
	}

	svgimage.Stage(stage)

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
	case *Content:
		toT := CopyBranchContent(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *JpgImage:
		toT := CopyBranchJpgImage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PngImage:
		toT := CopyBranchPngImage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SvgImage:
		toT := CopyBranchSvgImage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchContent(mapOrigCopy map[any]any, contentFrom *Content) (contentTo *Content) {

	// contentFrom has already been copied
	if _contentTo, ok := mapOrigCopy[contentFrom]; ok {
		contentTo = _contentTo.(*Content)
		return
	}

	contentTo = new(Content)
	mapOrigCopy[contentFrom] = contentTo
	contentFrom.CopyBasicFields(contentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchJpgImage(mapOrigCopy map[any]any, jpgimageFrom *JpgImage) (jpgimageTo *JpgImage) {

	// jpgimageFrom has already been copied
	if _jpgimageTo, ok := mapOrigCopy[jpgimageFrom]; ok {
		jpgimageTo = _jpgimageTo.(*JpgImage)
		return
	}

	jpgimageTo = new(JpgImage)
	mapOrigCopy[jpgimageFrom] = jpgimageTo
	jpgimageFrom.CopyBasicFields(jpgimageTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPngImage(mapOrigCopy map[any]any, pngimageFrom *PngImage) (pngimageTo *PngImage) {

	// pngimageFrom has already been copied
	if _pngimageTo, ok := mapOrigCopy[pngimageFrom]; ok {
		pngimageTo = _pngimageTo.(*PngImage)
		return
	}

	pngimageTo = new(PngImage)
	mapOrigCopy[pngimageFrom] = pngimageTo
	pngimageFrom.CopyBasicFields(pngimageTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSvgImage(mapOrigCopy map[any]any, svgimageFrom *SvgImage) (svgimageTo *SvgImage) {

	// svgimageFrom has already been copied
	if _svgimageTo, ok := mapOrigCopy[svgimageFrom]; ok {
		svgimageTo = _svgimageTo.(*SvgImage)
		return
	}

	svgimageTo = new(SvgImage)
	mapOrigCopy[svgimageFrom] = svgimageTo
	svgimageFrom.CopyBasicFields(svgimageTo)

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
	case *Content:
		stage.UnstageBranchContent(target)

	case *JpgImage:
		stage.UnstageBranchJpgImage(target)

	case *PngImage:
		stage.UnstageBranchPngImage(target)

	case *SvgImage:
		stage.UnstageBranchSvgImage(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchContent(content *Content) {

	// check if instance is already staged
	if !IsStaged(stage, content) {
		return
	}

	content.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchJpgImage(jpgimage *JpgImage) {

	// check if instance is already staged
	if !IsStaged(stage, jpgimage) {
		return
	}

	jpgimage.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPngImage(pngimage *PngImage) {

	// check if instance is already staged
	if !IsStaged(stage, pngimage) {
		return
	}

	pngimage.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchSvgImage(svgimage *SvgImage) {

	// check if instance is already staged
	if !IsStaged(stage, svgimage) {
		return
	}

	svgimage.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (content *Content) GongDiff(stage *Stage, contentOther *Content) (diffs []string) {
	// insertion point for field diffs
	if content.Name != contentOther.Name {
		diffs = append(diffs, content.GongMarshallField(stage, "Name"))
	}
	if content.Content != contentOther.Content {
		diffs = append(diffs, content.GongMarshallField(stage, "Content"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (jpgimage *JpgImage) GongDiff(stage *Stage, jpgimageOther *JpgImage) (diffs []string) {
	// insertion point for field diffs
	if jpgimage.Name != jpgimageOther.Name {
		diffs = append(diffs, jpgimage.GongMarshallField(stage, "Name"))
	}
	if jpgimage.Base64Content != jpgimageOther.Base64Content {
		diffs = append(diffs, jpgimage.GongMarshallField(stage, "Base64Content"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (pngimage *PngImage) GongDiff(stage *Stage, pngimageOther *PngImage) (diffs []string) {
	// insertion point for field diffs
	if pngimage.Name != pngimageOther.Name {
		diffs = append(diffs, pngimage.GongMarshallField(stage, "Name"))
	}
	if pngimage.Base64Content != pngimageOther.Base64Content {
		diffs = append(diffs, pngimage.GongMarshallField(stage, "Base64Content"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (svgimage *SvgImage) GongDiff(stage *Stage, svgimageOther *SvgImage) (diffs []string) {
	// insertion point for field diffs
	if svgimage.Name != svgimageOther.Name {
		diffs = append(diffs, svgimage.GongMarshallField(stage, "Name"))
	}
	if svgimage.Content != svgimageOther.Content {
		diffs = append(diffs, svgimage.GongMarshallField(stage, "Content"))
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
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, k+1)
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
