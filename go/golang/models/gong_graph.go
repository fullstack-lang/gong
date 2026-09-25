package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/fullstack-lang/gong/go/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const ModelGongGraphFileTemplate = `// generated code - do not edit
package {{PkgGoName}}

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

// insertion point for stage per struct{{` + string(rune(ModelGongGraphStructInsertionIsStagedPerStruct)) + `}}
// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct{{` + string(rune(ModelGongGraphStructInsertionStageBranchPerStruct)) + `}}
// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch{{` + string(rune(ModelGongGraphStructInsertionCopyBranch)) + `}}
	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct{{` + string(rune(ModelGongGraphStructInsertionCopyBranchPerStruct)) + `}}
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

// insertion point for unstage branch per struct{{` + string(rune(ModelGongGraphStructInsertionUnstageBranchPerStruct)) + `}}
// insertion point for pointer reconstruction from references{{` + string(rune(ModelGongGraphReconstructPointersFromReferences)) + `}}
// insertion point for pointer reconstruction from instances{{` + string(rune(ModelGongGraphReconstructPointersFromInstances)) + `}}
// insertion point for diff per struct{{` + string(rune(ModelGongGraphDiff)) + `}}
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
`

// insertion points are places where the code is
// generated per gong struct
type ModelGongGraphStructInsertionId int

const (
	ModelGongGraphStructInsertionIsStaged ModelGongGraphStructInsertionId = iota
	ModelGongGraphStructInsertionIsStagedPerStruct
	ModelGongGraphStructInsertionStageBranch
	ModelGongGraphStructInsertionStageBranchPerStruct
	ModelGongGraphStructInsertionCopyBranch
	ModelGongGraphStructInsertionCopyBranchPerStruct
	ModelGongGraphStructInsertionUnstageBranch
	ModelGongGraphStructInsertionUnstageBranchPerStruct
	ModelGongGraphDiff
	ModelGongGraphReconstructPointersFromInstances
	ModelGongGraphReconstructPointersFromReferences
	ModelGongGraphStructInsertionsNb
)

var ModelGongGraphStructSubTemplateCode map[ModelGongGraphStructInsertionId]string = // new line
map[ModelGongGraphStructInsertionId]string{

	ModelGongGraphStructInsertionIsStaged: "",
	ModelGongGraphStructInsertionIsStagedPerStruct: `
func ({{structname}} *{{Structname}}) GongIsStaged(stage *Stage) bool {
	_, ok := stage.{{Structname}}s[{{structname}}]
	return ok
}
`,
	ModelGongGraphStructInsertionStageBranch: "",
	ModelGongGraphStructInsertionStageBranchPerStruct: `
func ({{structname}} *{{Structname}}) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged({{structname}}) {
		return
	}

	{{structname}}.Stage(stage)

	//insertion point for the staging of instances referenced by pointers{{StagingPointers}}

	//insertion point for the staging of instances referenced by slice of pointers{{StagingSliceOfPointers}}

}
`,
	ModelGongGraphStructInsertionCopyBranch: `
	case *{{Structname}}:
		toT := GongCopyBranch{{Structname}}(mapOrigCopy, fromT)
		return any(toT).(*Type)
`,
	ModelGongGraphStructInsertionCopyBranchPerStruct: `
func GongCopyBranch{{Structname}}(mapOrigCopy map[any]any, {{structname}}From *{{Structname}}) ({{structname}}To *{{Structname}}) {
	var alreadyCopied bool
	{{structname}}To, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, {{structname}}From)
	if alreadyCopied {
		return
	}
	{{structname}}From.GongCopyBasicFields({{structname}}To)

	//insertion point for the staging of instances referenced by pointers{{CopyingPointers}}

	//insertion point for the staging of instances referenced by slice of pointers{{CopyingSliceOfPointers}}

	return
}
`,
	ModelGongGraphStructInsertionUnstageBranch: "",
	ModelGongGraphStructInsertionUnstageBranchPerStruct: `
func ({{structname}} *{{Structname}}) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged({{structname}}) {
		return
	}

	{{structname}}.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers{{UnstagingPointers}}

	//insertion point for the staging of instances referenced by slice of pointers{{UnstagingSliceOfPointers}}

}
`,
	ModelGongGraphReconstructPointersFromReferences: `
func (reference *{{Structname}}) GongReconstructPointersFromReferences(stage *Stage, instance *{{Structname}}) {
	// insertion point for pointers field{{FieldReconstructPointersFromReferences}}
	// insertion point for slice of pointers field{{FieldReconstructSliceOfPointersFromReferences}}
}
`,
	ModelGongGraphReconstructPointersFromInstances: `
func (reference *{{Structname}}) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field{{FieldReconstructPointersFromInstances}}
	// insertion point for slice of pointers fields{{FieldReconstructSliceOfPointersFromInstances}}
}
`,
	ModelGongGraphDiff: `
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func ({{structname}} *{{Structname}}) GongDiff(stage *Stage, {{structname}}Other *{{Structname}}) (diffs []string) {
	// insertion point for field diffs{{FieldDiff}}

	return
}
`,
}

type GongGraphFilePerStructSubTemplateId int

const (
	GongGraphFileFieldSubTmplStagePointerField GongGraphFilePerStructSubTemplateId = iota
	GongGraphFileFieldSubTmplStageSliceOfPointersField
	GongGraphFileFieldSubTmplCopyPointerField
	GongGraphFileFieldSubTmplCopyPointerFieldAndStop
	GongGraphFileFieldSubTmplCopySliceOfPointersField
	GongGraphFileFieldSubTmplUnstagePointerField
	GongGraphFileFieldSubTmplUnstageSliceOfPointersField
	GongGraphBasicFieldDiff

	GongGraphPointerFieldDiff
	GongGraphSliceOfPointerFieldDiff

	GongGraphPointerFieldReconstructPointersFromReferences
	GongGraphPointerFieldReconstructPointersFromInstances

	GongGraphSliceOfPointersFieldReconstructPointersFromReferences
	GongGraphSliceOfPointersFieldReconstructPointersFromInstances
)

var GongGraphFileFieldFieldSubTemplateCode map[GongGraphFilePerStructSubTemplateId]string = // declaration of the sub templates
map[GongGraphFilePerStructSubTemplateId]string{

	GongGraphFileFieldSubTmplStagePointerField: `
	if {{structname}}.{{FieldName}} != nil {
		stage.StageBranch({{structname}}.{{FieldName}})
	}`,
	GongGraphFileFieldSubTmplStageSliceOfPointersField: `
	for _, _{{assocstructname}} := range {{structname}}.{{FieldName}} {
		stage.StageBranch(_{{assocstructname}})
	}`,
	GongGraphFileFieldSubTmplCopyPointerField: `
	if {{structname}}From.{{FieldName}} != nil {
		{{structname}}To.{{FieldName}} = GongCopyBranch{{AssocStructName}}(mapOrigCopy, {{structname}}From.{{FieldName}})
	}`,
	GongGraphFileFieldSubTmplCopyPointerFieldAndStop: `
	if {{structname}}From.{{FieldName}} != nil {
		{{structname}}To.{{FieldName}} = {{structname}}From.{{FieldName}}
	}`,
	GongGraphFileFieldSubTmplCopySliceOfPointersField: `
	for _, _{{assocstructname}} := range {{structname}}From.{{FieldName}} {
		{{structname}}To.{{FieldName}} = append({{structname}}To.{{FieldName}}, GongCopyBranch{{AssocStructName}}(mapOrigCopy, _{{assocstructname}}))
	}`,
	GongGraphFileFieldSubTmplUnstagePointerField: `
	if {{structname}}.{{FieldName}} != nil {
		stage.UnstageBranch({{structname}}.{{FieldName}})
	}`,
	GongGraphFileFieldSubTmplUnstageSliceOfPointersField: `
	for _, _{{assocstructname}} := range {{structname}}.{{FieldName}} {
		stage.UnstageBranch(_{{assocstructname}})
	}`,
	GongGraphBasicFieldDiff: `
	if {{structname}}.{{FieldName}} != {{structname}}Other.{{FieldName}} {
		diffs = append(diffs, {{structname}}.GongMarshallField(stage, "{{FieldName}}"))
	}`,
	GongGraphPointerFieldDiff: `
	if {{structname}}.{{FieldName}} != {{structname}}Other.{{FieldName}} {
		diffs = append(diffs, {{structname}}.GongMarshallField(stage, "{{FieldName}}"))
	}`,
	GongGraphSliceOfPointerFieldDiff: `
	if ops := __gong__diffSliceOfPointers(stage, {{structname}}, "{{FieldName}}", {{structname}}Other.{{FieldName}}, {{structname}}.{{FieldName}}); ops != "" {
		diffs = append(diffs, ops)
	}`,

	GongGraphPointerFieldReconstructPointersFromReferences: `
	__gong__reconstructPointer(&reference.{{FieldName}}, stage.{{AssocStructName}}s_reference, instance.{{FieldName}})`,
	GongGraphPointerFieldReconstructPointersFromInstances: `
	__gong__reconstructPointerFromInstance(&reference.{{FieldName}}, stage.{{AssocStructName}}s_instance)`,
	GongGraphSliceOfPointersFieldReconstructPointersFromReferences: `
	__gong__reconstructSliceOfPointersFromReferences(&reference.{{FieldName}}, stage.{{AssocStructName}}s_reference, instance.{{FieldName}})`,
	GongGraphSliceOfPointersFieldReconstructPointersFromInstances: `
	__gong__reconstructSliceOfPointersFromInstances(&reference.{{FieldName}}, stage.{{AssocStructName}}s_instance)`,
}

func CodeGeneratorModelGongGraph(
	modelPkg *models.ModelPkg,
	pkgName string,
	pkgPath string) {

	// generate the typescript file
	codeGO := ModelGongGraphFileTemplate

	subStructCodes := make(map[ModelGongGraphStructInsertionId]string)
	for subStructTemplate := range ModelGongGraphStructSubTemplateCode {
		subStructCodes[subStructTemplate] = ""
	}

	// sort gong structs per name (for reproductibility)
	gongStructs := []*models.GongStruct{}
	for _, _struct := range modelPkg.GongStructs {
		gongStructs = append(gongStructs, _struct)
	}
	sort.Slice(gongStructs[:], func(i, j int) bool {
		return gongStructs[i].Name < gongStructs[j].Name
	})

	for _, gongStruct := range gongStructs {

		if !gongStruct.HasNameField() {
			continue
		}

		for subStructTemplate := range ModelGongGraphStructSubTemplateCode {

			pointerStagingCode := ""
			sliceOfPointerStagingCode := ""
			pointerCopyingCode := ""
			sliceOfPointerCopyingCode := ""
			pointerUnstagingCode := ""
			sliceOfPointerUnstagingCode := ""
			fieldDiff := ""
			fieldReconstructPointersFromReferences := ""
			fieldReconstructPointersFromInstances := ""
			fieldReconstructSliceOfPointersFromReferences := ""
			fieldReconstructSliceOfPointersFromInstances := ""

			for _, field := range gongStruct.Fields {

				switch field := field.(type) {
				case *models.GongBasicField, *models.GongTimeField:
					fieldDiff += models.Replace1(
						GongGraphFileFieldFieldSubTemplateCode[GongGraphBasicFieldDiff],
						"{{FieldName}}", field.GetName())
				case *models.PointerToGongStructField:
					if field.GongStruct.IsOmittedForMarshalling {
						continue
					}
					pointerStagingCode += models.Replace2(
						GongGraphFileFieldFieldSubTemplateCode[GongGraphFileFieldSubTmplStagePointerField],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name)
					if !field.IsType {
						pointerCopyingCode += models.Replace2(
							GongGraphFileFieldFieldSubTemplateCode[GongGraphFileFieldSubTmplCopyPointerField],
							"{{FieldName}}", field.Name,
							"{{AssocStructName}}", field.GongStruct.Name)
					} else {
						pointerCopyingCode += models.Replace2(
							GongGraphFileFieldFieldSubTemplateCode[GongGraphFileFieldSubTmplCopyPointerFieldAndStop],
							"{{FieldName}}", field.Name,
							"{{AssocStructName}}", field.GongStruct.Name)
					}

					pointerUnstagingCode += models.Replace2(
						GongGraphFileFieldFieldSubTemplateCode[GongGraphFileFieldSubTmplUnstagePointerField],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name)

					fieldDiff += models.Replace2(
						GongGraphFileFieldFieldSubTemplateCode[GongGraphPointerFieldDiff],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name)
					fieldReconstructPointersFromReferences += models.Replace2(
						GongGraphFileFieldFieldSubTemplateCode[GongGraphPointerFieldReconstructPointersFromReferences],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name)
					fieldReconstructPointersFromInstances += models.Replace2(
						GongGraphFileFieldFieldSubTemplateCode[GongGraphPointerFieldReconstructPointersFromInstances],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name)

				case *models.SliceOfPointerToGongStructField:
					if field.GongStruct.IsOmittedForMarshalling {
						continue
					}
					sliceOfPointerStagingCode += models.Replace3(
						GongGraphFileFieldFieldSubTemplateCode[GongGraphFileFieldSubTmplStageSliceOfPointersField],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name,
						"{{assocstructname}}", strings.ToLower(field.GongStruct.Name))
					sliceOfPointerCopyingCode += models.Replace3(
						GongGraphFileFieldFieldSubTemplateCode[GongGraphFileFieldSubTmplCopySliceOfPointersField],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name,
						"{{assocstructname}}", strings.ToLower(field.GongStruct.Name))
					sliceOfPointerUnstagingCode += models.Replace3(
						GongGraphFileFieldFieldSubTemplateCode[GongGraphFileFieldSubTmplUnstageSliceOfPointersField],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name,
						"{{assocstructname}}", strings.ToLower(field.GongStruct.Name))
					fieldDiff += models.Replace2(
						GongGraphFileFieldFieldSubTemplateCode[GongGraphSliceOfPointerFieldDiff],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name)
					fieldReconstructSliceOfPointersFromReferences += models.Replace2(
						GongGraphFileFieldFieldSubTemplateCode[GongGraphSliceOfPointersFieldReconstructPointersFromReferences],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name)
					fieldReconstructSliceOfPointersFromInstances += models.Replace2(
						GongGraphFileFieldFieldSubTemplateCode[GongGraphSliceOfPointersFieldReconstructPointersFromInstances],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name)
				default:
					_ = field
				}

			}

			sliceOfPointerStagingCode = models.Replace2(sliceOfPointerStagingCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			pointerStagingCode = models.Replace2(pointerStagingCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			sliceOfPointerCopyingCode = models.Replace2(sliceOfPointerCopyingCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			pointerCopyingCode = models.Replace2(pointerCopyingCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			sliceOfPointerUnstagingCode = models.Replace2(sliceOfPointerUnstagingCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			pointerUnstagingCode = models.Replace2(pointerUnstagingCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			fieldDiff = models.Replace2(fieldDiff,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			generatedCodeFromSubTemplate := models.Replace13(ModelGongGraphStructSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{StagingPointers}}", pointerStagingCode,
				"{{StagingSliceOfPointers}}", sliceOfPointerStagingCode,
				"{{CopyingPointers}}", pointerCopyingCode,
				"{{CopyingSliceOfPointers}}", sliceOfPointerCopyingCode,
				"{{UnstagingPointers}}", pointerUnstagingCode,
				"{{UnstagingSliceOfPointers}}", sliceOfPointerUnstagingCode,
				"{{FieldDiff}}", fieldDiff,
				"{{FieldReconstructPointersFromReferences}}", fieldReconstructPointersFromReferences,
				"{{FieldReconstructPointersFromInstances}}", fieldReconstructPointersFromInstances,
				"{{FieldReconstructSliceOfPointersFromReferences}}", fieldReconstructSliceOfPointersFromReferences,
				"{{FieldReconstructSliceOfPointersFromInstances}}", fieldReconstructSliceOfPointersFromInstances,
			)

			subStructCodes[subStructTemplate] += generatedCodeFromSubTemplate
		}

	}

	// substitutes {{<<insertionPerStructId points>>}} stuff with generated code
	for insertionPerStructId := range ModelGongGraphStructInsertionsNb {
		toReplace := "{{" + string(rune(insertionPerStructId)) + "}}"
		codeGO = strings.ReplaceAll(codeGO, toReplace, subStructCodes[insertionPerStructId])
	}

	caserEnglish := cases.Title(language.English)
	codeGO = models.Replace(codeGO,
		"{{PkgName}}", pkgName,
		"{{TitlePkgName}}", caserEnglish.String(pkgName),
		"{{pkgname}}", strings.ToLower(pkgName),
		"{{PkgGoName}}", modelPkg.PkgGoName,
		"	 | ", "	", // for the replacement of the of the first bar in the Gongstruct Type def
	)

	file, err := os.Create(filepath.Join(pkgPath, string(models.GeneratedGongGraphGoFilePath)))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)
}
