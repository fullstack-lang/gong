// generated code - do not edit
package portrayalcatalogue

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
func (md_portrayalcataloguereference_propertytype *MD_PortrayalCatalogueReference_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_PortrayalCatalogueReference_PropertyTypes[md_portrayalcataloguereference_propertytype]
	return ok
}

func (md_portrayalcataloguereference_type *MD_PortrayalCatalogueReference_Type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_PortrayalCatalogueReference_Types[md_portrayalcataloguereference_type]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (md_portrayalcataloguereference_propertytype *MD_PortrayalCatalogueReference_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_portrayalcataloguereference_propertytype) {
		return
	}

	md_portrayalcataloguereference_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_portrayalcataloguereference_propertytype.MD_PortrayalCatalogueReference != nil {
		stage.StageBranch(md_portrayalcataloguereference_propertytype.MD_PortrayalCatalogueReference)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_portrayalcataloguereference_type *MD_PortrayalCatalogueReference_Type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_portrayalcataloguereference_type) {
		return
	}

	md_portrayalcataloguereference_type.Stage(stage)

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
	case *MD_PortrayalCatalogueReference_PropertyType:
		toT := GongCopyBranchMD_PortrayalCatalogueReference_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_PortrayalCatalogueReference_Type:
		toT := GongCopyBranchMD_PortrayalCatalogueReference_Type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchMD_PortrayalCatalogueReference_PropertyType(mapOrigCopy map[any]any, md_portrayalcataloguereference_propertytypeFrom *MD_PortrayalCatalogueReference_PropertyType) (md_portrayalcataloguereference_propertytypeTo *MD_PortrayalCatalogueReference_PropertyType) {
	var alreadyCopied bool
	md_portrayalcataloguereference_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_portrayalcataloguereference_propertytypeFrom)
	if alreadyCopied {
		return
	}
	md_portrayalcataloguereference_propertytypeFrom.GongCopyBasicFields(md_portrayalcataloguereference_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if md_portrayalcataloguereference_propertytypeFrom.MD_PortrayalCatalogueReference != nil {
		md_portrayalcataloguereference_propertytypeTo.MD_PortrayalCatalogueReference = GongCopyBranchMD_PortrayalCatalogueReference_Type(mapOrigCopy, md_portrayalcataloguereference_propertytypeFrom.MD_PortrayalCatalogueReference)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_PortrayalCatalogueReference_Type(mapOrigCopy map[any]any, md_portrayalcataloguereference_typeFrom *MD_PortrayalCatalogueReference_Type) (md_portrayalcataloguereference_typeTo *MD_PortrayalCatalogueReference_Type) {
	var alreadyCopied bool
	md_portrayalcataloguereference_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_portrayalcataloguereference_typeFrom)
	if alreadyCopied {
		return
	}
	md_portrayalcataloguereference_typeFrom.GongCopyBasicFields(md_portrayalcataloguereference_typeTo)

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
func (md_portrayalcataloguereference_propertytype *MD_PortrayalCatalogueReference_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_portrayalcataloguereference_propertytype) {
		return
	}

	md_portrayalcataloguereference_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_portrayalcataloguereference_propertytype.MD_PortrayalCatalogueReference != nil {
		stage.UnstageBranch(md_portrayalcataloguereference_propertytype.MD_PortrayalCatalogueReference)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_portrayalcataloguereference_type *MD_PortrayalCatalogueReference_Type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_portrayalcataloguereference_type) {
		return
	}

	md_portrayalcataloguereference_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *MD_PortrayalCatalogueReference_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *MD_PortrayalCatalogueReference_PropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.MD_PortrayalCatalogueReference, stage.MD_PortrayalCatalogueReference_Types_reference, instance.MD_PortrayalCatalogueReference)
	// insertion point for slice of pointers field
}

func (reference *MD_PortrayalCatalogueReference_Type) GongReconstructPointersFromReferences(stage *Stage, instance *MD_PortrayalCatalogueReference_Type) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *MD_PortrayalCatalogueReference_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.MD_PortrayalCatalogueReference, stage.MD_PortrayalCatalogueReference_Types_instance)
	// insertion point for slice of pointers fields
}

func (reference *MD_PortrayalCatalogueReference_Type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_portrayalcataloguereference_propertytype *MD_PortrayalCatalogueReference_PropertyType) GongDiff(stage *Stage, md_portrayalcataloguereference_propertytypeOther *MD_PortrayalCatalogueReference_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if md_portrayalcataloguereference_propertytype.Name != md_portrayalcataloguereference_propertytypeOther.Name {
		diffs = append(diffs, md_portrayalcataloguereference_propertytype.GongMarshallField(stage, "Name"))
	}
	if md_portrayalcataloguereference_propertytype.NilReason != md_portrayalcataloguereference_propertytypeOther.NilReason {
		diffs = append(diffs, md_portrayalcataloguereference_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if md_portrayalcataloguereference_propertytype.Uuidref != md_portrayalcataloguereference_propertytypeOther.Uuidref {
		diffs = append(diffs, md_portrayalcataloguereference_propertytype.GongMarshallField(stage, "Uuidref"))
	}
	if md_portrayalcataloguereference_propertytype.Type_XSD != md_portrayalcataloguereference_propertytypeOther.Type_XSD {
		diffs = append(diffs, md_portrayalcataloguereference_propertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if md_portrayalcataloguereference_propertytype.Href != md_portrayalcataloguereference_propertytypeOther.Href {
		diffs = append(diffs, md_portrayalcataloguereference_propertytype.GongMarshallField(stage, "Href"))
	}
	if md_portrayalcataloguereference_propertytype.Role != md_portrayalcataloguereference_propertytypeOther.Role {
		diffs = append(diffs, md_portrayalcataloguereference_propertytype.GongMarshallField(stage, "Role"))
	}
	if md_portrayalcataloguereference_propertytype.Arcrole != md_portrayalcataloguereference_propertytypeOther.Arcrole {
		diffs = append(diffs, md_portrayalcataloguereference_propertytype.GongMarshallField(stage, "Arcrole"))
	}
	if md_portrayalcataloguereference_propertytype.Title != md_portrayalcataloguereference_propertytypeOther.Title {
		diffs = append(diffs, md_portrayalcataloguereference_propertytype.GongMarshallField(stage, "Title"))
	}
	if md_portrayalcataloguereference_propertytype.Show != md_portrayalcataloguereference_propertytypeOther.Show {
		diffs = append(diffs, md_portrayalcataloguereference_propertytype.GongMarshallField(stage, "Show"))
	}
	if md_portrayalcataloguereference_propertytype.Actuate != md_portrayalcataloguereference_propertytypeOther.Actuate {
		diffs = append(diffs, md_portrayalcataloguereference_propertytype.GongMarshallField(stage, "Actuate"))
	}
	if md_portrayalcataloguereference_propertytype.MD_PortrayalCatalogueReference != md_portrayalcataloguereference_propertytypeOther.MD_PortrayalCatalogueReference {
		diffs = append(diffs, md_portrayalcataloguereference_propertytype.GongMarshallField(stage, "MD_PortrayalCatalogueReference"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_portrayalcataloguereference_type *MD_PortrayalCatalogueReference_Type) GongDiff(stage *Stage, md_portrayalcataloguereference_typeOther *MD_PortrayalCatalogueReference_Type) (diffs []string) {
	// insertion point for field diffs
	if md_portrayalcataloguereference_type.Name != md_portrayalcataloguereference_typeOther.Name {
		diffs = append(diffs, md_portrayalcataloguereference_type.GongMarshallField(stage, "Name"))
	}
	if md_portrayalcataloguereference_type.Id != md_portrayalcataloguereference_typeOther.Id {
		diffs = append(diffs, md_portrayalcataloguereference_type.GongMarshallField(stage, "Id"))
	}
	if md_portrayalcataloguereference_type.Uuid != md_portrayalcataloguereference_typeOther.Uuid {
		diffs = append(diffs, md_portrayalcataloguereference_type.GongMarshallField(stage, "Uuid"))
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
