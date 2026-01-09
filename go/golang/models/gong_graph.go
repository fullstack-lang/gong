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
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage{{` + string(rune(ModelGongGraphStructInsertionIsStaged)) + `}}
	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage{{` + string(rune(ModelGongGraphStructInsertionIsStaged)) + `}}
	default:
		_ = target
	}
	return
}

// insertion point for stage per struct{{` + string(rune(ModelGongGraphStructInsertionIsStagedPerStruct)) + `}}
// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch{{` + string(rune(ModelGongGraphStructInsertionStageBranch)) + `}}
	default:
		_ = target
	}
}

// insertion point for stage branch per struct{{` + string(rune(ModelGongGraphStructInsertionStageBranchPerStruct)) + `}}
// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

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
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch{{` + string(rune(ModelGongGraphStructInsertionUnstageBranch)) + `}}
	default:
		_ = target
	}
}

// insertion point for unstage branch per struct{{` + string(rune(ModelGongGraphStructInsertionUnstageBranchPerStruct)) + `}}
// insertion point for diff per struct{{` + string(rune(ModelGongGraphDiff)) + `}}
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
	ModelGongGraphStructInsertionsNb
)

var ModelGongGraphStructSubTemplateCode map[ModelGongGraphStructInsertionId]string = // new line
map[ModelGongGraphStructInsertionId]string{

	ModelGongGraphStructInsertionIsStaged: `
	case *{{Structname}}:
		ok = stage.IsStaged{{Structname}}(target)
`,
	ModelGongGraphStructInsertionIsStagedPerStruct: `
func (stage *Stage) IsStaged{{Structname}}({{structname}} *{{Structname}}) (ok bool) {

	_, ok = stage.{{Structname}}s[{{structname}}]

	return
}
`,
	ModelGongGraphStructInsertionStageBranch: `
	case *{{Structname}}:
		stage.StageBranch{{Structname}}(target)
`,
	ModelGongGraphStructInsertionStageBranchPerStruct: `
func (stage *Stage) StageBranch{{Structname}}({{structname}} *{{Structname}}) {

	// check if instance is already staged
	if IsStaged(stage, {{structname}}) {
		return
	}

	{{structname}}.Stage(stage)

	//insertion point for the staging of instances referenced by pointers{{StagingPointers}}

	//insertion point for the staging of instances referenced by slice of pointers{{StagingSliceOfPointers}}

}
`,
	ModelGongGraphStructInsertionCopyBranch: `
	case *{{Structname}}:
		toT := CopyBranch{{Structname}}(mapOrigCopy, fromT)
		return any(toT).(*Type)
`,
	ModelGongGraphStructInsertionCopyBranchPerStruct: `
func CopyBranch{{Structname}}(mapOrigCopy map[any]any, {{structname}}From *{{Structname}}) ({{structname}}To *{{Structname}}) {

	// {{structname}}From has already been copied
	if _{{structname}}To, ok := mapOrigCopy[{{structname}}From]; ok {
		{{structname}}To = _{{structname}}To.(*{{Structname}})
		return
	}

	{{structname}}To = new({{Structname}})
	mapOrigCopy[{{structname}}From] = {{structname}}To
	{{structname}}From.CopyBasicFields({{structname}}To)

	//insertion point for the staging of instances referenced by pointers{{CopyingPointers}}

	//insertion point for the staging of instances referenced by slice of pointers{{CopyingSliceOfPointers}}

	return
}
`,
	ModelGongGraphStructInsertionUnstageBranch: `
	case *{{Structname}}:
		stage.UnstageBranch{{Structname}}(target)
`,
	ModelGongGraphStructInsertionUnstageBranchPerStruct: `
func (stage *Stage) UnstageBranch{{Structname}}({{structname}} *{{Structname}}) {

	// check if instance is already staged
	if !IsStaged(stage, {{structname}}) {
		return
	}

	{{structname}}.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers{{UnstagingPointers}}

	//insertion point for the staging of instances referenced by slice of pointers{{UnstagingSliceOfPointers}}

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
)

var GongGraphFileFieldFieldSubTemplateCode map[GongGraphFilePerStructSubTemplateId]string = // declaration of the sub templates
map[GongGraphFilePerStructSubTemplateId]string{

	GongGraphFileFieldSubTmplStagePointerField: `
	if {{structname}}.{{FieldName}} != nil {
		StageBranch(stage, {{structname}}.{{FieldName}})
	}`,
	GongGraphFileFieldSubTmplStageSliceOfPointersField: `
	for _, _{{assocstructname}} := range {{structname}}.{{FieldName}} {
		StageBranch(stage, _{{assocstructname}})
	}`,
	GongGraphFileFieldSubTmplCopyPointerField: `
	if {{structname}}From.{{FieldName}} != nil {
		{{structname}}To.{{FieldName}} = CopyBranch{{AssocStructName}}(mapOrigCopy, {{structname}}From.{{FieldName}})
	}`,
	GongGraphFileFieldSubTmplCopyPointerFieldAndStop: `
	if {{structname}}From.{{FieldName}} != nil {
		{{structname}}To.{{FieldName}} = {{structname}}From.{{FieldName}}
	}`,
	GongGraphFileFieldSubTmplCopySliceOfPointersField: `
	for _, _{{assocstructname}} := range {{structname}}From.{{FieldName}} {
		{{structname}}To.{{FieldName}} = append({{structname}}To.{{FieldName}}, CopyBranch{{AssocStructName}}(mapOrigCopy, _{{assocstructname}}))
	}`,
	GongGraphFileFieldSubTmplUnstagePointerField: `
	if {{structname}}.{{FieldName}} != nil {
		UnstageBranch(stage, {{structname}}.{{FieldName}})
	}`,
	GongGraphFileFieldSubTmplUnstageSliceOfPointersField: `
	for _, _{{assocstructname}} := range {{structname}}.{{FieldName}} {
		UnstageBranch(stage, _{{assocstructname}})
	}`,
	GongGraphBasicFieldDiff: `
	if {{structname}}.{{FieldName}} != {{structname}}Other.{{FieldName}} {
		diffs = append(diffs, {{structname}}.GongMarshallField(stage, "{{FieldName}}"))
	}`,
	GongGraphPointerFieldDiff: `
	if ({{structname}}.{{FieldName}} == nil) != ({{structname}}Other.{{FieldName}} == nil) {
		diffs = append(diffs, {{structname}}.GongMarshallField(stage, "{{FieldName}}"))
	} else if {{structname}}.{{FieldName}} != nil && {{structname}}Other.{{FieldName}} != nil {
		if {{structname}}.{{FieldName}} != {{structname}}Other.{{FieldName}} {
			diffs = append(diffs, {{structname}}.GongMarshallField(stage, "{{FieldName}}"))
		}
	}`,
	GongGraphSliceOfPointerFieldDiff: `
	{{FieldName}}Different := false
	if len({{structname}}.{{FieldName}}) != len({{structname}}Other.{{FieldName}}) {
		{{FieldName}}Different = true
	} else {
		for i := range {{structname}}.{{FieldName}} {
			if ({{structname}}.{{FieldName}}[i] == nil) != ({{structname}}Other.{{FieldName}}[i] == nil) {
				{{FieldName}}Different = true
				break
			} else if {{structname}}.{{FieldName}}[i] != nil && {{structname}}Other.{{FieldName}}[i] != nil {
				// this is a pointer comparaison
				if {{structname}}.{{FieldName}}[i] != {{structname}}Other.{{FieldName}}[i] {
					{{FieldName}}Different = true
					break
				}
			}
		}
	}
	if {{FieldName}}Different {
		ops := Diff(stage, {{structname}}, {{structname}}Other, "{{FieldName}}", {{structname}}Other.{{FieldName}}, {{structname}}.{{FieldName}})
		diffs = append(diffs, ops)
	}`,
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

			for _, field := range gongStruct.Fields {

				switch field := field.(type) {
				case *models.GongBasicField, *models.GongTimeField:
					fieldDiff += models.Replace1(
						GongGraphFileFieldFieldSubTemplateCode[GongGraphBasicFieldDiff],
						"{{FieldName}}", field.GetName())
				case *models.PointerToGongStructField:
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

				case *models.SliceOfPointerToGongStructField:
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

			generatedCodeFromSubTemplate := models.Replace9(ModelGongGraphStructSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{StagingPointers}}", pointerStagingCode,
				"{{StagingSliceOfPointers}}", sliceOfPointerStagingCode,
				"{{CopyingPointers}}", pointerCopyingCode,
				"{{CopyingSliceOfPointers}}", sliceOfPointerCopyingCode,
				"{{UnstagingPointers}}", pointerUnstagingCode,
				"{{UnstagingSliceOfPointers}}", sliceOfPointerUnstagingCode,
				"{{FieldDiff}}", fieldDiff,
			)

			subStructCodes[subStructTemplate] += generatedCodeFromSubTemplate
		}

	}

	// substitutes {{<<insertionPerStructId points>>}} stuff with generated code
	for insertionPerStructId := ModelGongGraphStructInsertionId(0); insertionPerStructId < ModelGongGraphStructInsertionsNb; insertionPerStructId++ {
		toReplace := "{{" + string(rune(insertionPerStructId)) + "}}"
		codeGO = strings.ReplaceAll(codeGO, toReplace, subStructCodes[insertionPerStructId])
	}

	caserEnglish := cases.Title(language.English)
	codeGO = models.Replace4(codeGO,
		"{{PkgName}}", pkgName,
		"{{TitlePkgName}}", caserEnglish.String(pkgName),
		"{{pkgname}}", strings.ToLower(pkgName),
		"	 | ", "	", // for the replacement of the of the first bar in the Gongstruct Type def
	)

	file, err := os.Create(filepath.Join(pkgPath, string(models.GeneratedGongGraphGoFilePath)))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)
}
