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
func (content *Content) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Contents[content]

	return
}

func (stage *Stage) IsStagedContent(content *Content) (ok bool) {

	return content.GongIsStaged(stage)
}

func (jpgimage *JpgImage) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.JpgImages[jpgimage]

	return
}

func (stage *Stage) IsStagedJpgImage(jpgimage *JpgImage) (ok bool) {

	return jpgimage.GongIsStaged(stage)
}

func (pngimage *PngImage) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.PngImages[pngimage]

	return
}

func (stage *Stage) IsStagedPngImage(pngimage *PngImage) (ok bool) {

	return pngimage.GongIsStaged(stage)
}

func (svgimage *SvgImage) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.SvgImages[svgimage]

	return
}

func (stage *Stage) IsStagedSvgImage(svgimage *SvgImage) (ok bool) {

	return svgimage.GongIsStaged(stage)
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
func (content *Content) GongStageBranch(stage *Stage) {
	stage.StageBranchContent(content)
}

func (stage *Stage) StageBranchContent(content *Content) {

	// check if instance is already staged
	if stage.IsStaged(content) {
		return
	}

	content.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (jpgimage *JpgImage) GongStageBranch(stage *Stage) {
	stage.StageBranchJpgImage(jpgimage)
}

func (stage *Stage) StageBranchJpgImage(jpgimage *JpgImage) {

	// check if instance is already staged
	if stage.IsStaged(jpgimage) {
		return
	}

	jpgimage.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pngimage *PngImage) GongStageBranch(stage *Stage) {
	stage.StageBranchPngImage(pngimage)
}

func (stage *Stage) StageBranchPngImage(pngimage *PngImage) {

	// check if instance is already staged
	if stage.IsStaged(pngimage) {
		return
	}

	pngimage.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (svgimage *SvgImage) GongStageBranch(stage *Stage) {
	stage.StageBranchSvgImage(svgimage)
}

func (stage *Stage) StageBranchSvgImage(svgimage *SvgImage) {

	// check if instance is already staged
	if stage.IsStaged(svgimage) {
		return
	}

	svgimage.Stage(stage)

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
	case *Content:
		toT := GongCopyBranchContent(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *JpgImage:
		toT := GongCopyBranchJpgImage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PngImage:
		toT := GongCopyBranchPngImage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SvgImage:
		toT := GongCopyBranchSvgImage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchContent(mapOrigCopy map[any]any, contentFrom *Content) (contentTo *Content) {

	// contentFrom has already been copied
	if _contentTo, ok := mapOrigCopy[contentFrom]; ok {
		contentTo = _contentTo.(*Content)
		return
	}

	contentTo = new(Content)
	mapOrigCopy[contentFrom] = contentTo
	contentFrom.GongCopyBasicFields(contentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchJpgImage(mapOrigCopy map[any]any, jpgimageFrom *JpgImage) (jpgimageTo *JpgImage) {

	// jpgimageFrom has already been copied
	if _jpgimageTo, ok := mapOrigCopy[jpgimageFrom]; ok {
		jpgimageTo = _jpgimageTo.(*JpgImage)
		return
	}

	jpgimageTo = new(JpgImage)
	mapOrigCopy[jpgimageFrom] = jpgimageTo
	jpgimageFrom.GongCopyBasicFields(jpgimageTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPngImage(mapOrigCopy map[any]any, pngimageFrom *PngImage) (pngimageTo *PngImage) {

	// pngimageFrom has already been copied
	if _pngimageTo, ok := mapOrigCopy[pngimageFrom]; ok {
		pngimageTo = _pngimageTo.(*PngImage)
		return
	}

	pngimageTo = new(PngImage)
	mapOrigCopy[pngimageFrom] = pngimageTo
	pngimageFrom.GongCopyBasicFields(pngimageTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSvgImage(mapOrigCopy map[any]any, svgimageFrom *SvgImage) (svgimageTo *SvgImage) {

	// svgimageFrom has already been copied
	if _svgimageTo, ok := mapOrigCopy[svgimageFrom]; ok {
		svgimageTo = _svgimageTo.(*SvgImage)
		return
	}

	svgimageTo = new(SvgImage)
	mapOrigCopy[svgimageFrom] = svgimageTo
	svgimageFrom.GongCopyBasicFields(svgimageTo)

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

// UnstageBranch is a backward-compatible package-level forwarder.
func UnstageBranch(stage *Stage, instance GongstructIF) {
	stage.UnstageBranch(instance)
}

// insertion point for unstage branch per struct
func (content *Content) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchContent(content)
}

func (stage *Stage) UnstageBranchContent(content *Content) {

	// check if instance is already staged
	if !stage.IsStaged(content) {
		return
	}

	content.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (jpgimage *JpgImage) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchJpgImage(jpgimage)
}

func (stage *Stage) UnstageBranchJpgImage(jpgimage *JpgImage) {

	// check if instance is already staged
	if !stage.IsStaged(jpgimage) {
		return
	}

	jpgimage.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pngimage *PngImage) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPngImage(pngimage)
}

func (stage *Stage) UnstageBranchPngImage(pngimage *PngImage) {

	// check if instance is already staged
	if !stage.IsStaged(pngimage) {
		return
	}

	pngimage.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (svgimage *SvgImage) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchSvgImage(svgimage)
}

func (stage *Stage) UnstageBranchSvgImage(svgimage *SvgImage) {

	// check if instance is already staged
	if !stage.IsStaged(svgimage) {
		return
	}

	svgimage.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *Content) GongReconstructPointersFromReferences(stage *Stage, instance *Content) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *JpgImage) GongReconstructPointersFromReferences(stage *Stage, instance *JpgImage) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PngImage) GongReconstructPointersFromReferences(stage *Stage, instance *PngImage) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *SvgImage) GongReconstructPointersFromReferences(stage *Stage, instance *SvgImage) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *Content) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *JpgImage) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PngImage) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *SvgImage) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
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
