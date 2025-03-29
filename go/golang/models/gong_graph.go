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

// insertion point for unstage branch per struct{{` + string(rune(ModelGongGraphStructInsertionUnstageBranchPerStruct)) + `}}`

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

			for _, field := range gongStruct.Fields {

				switch field := field.(type) {
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

			generatedCodeFromSubTemplate := models.Replace8(ModelGongGraphStructSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{StagingPointers}}", pointerStagingCode,
				"{{StagingSliceOfPointers}}", sliceOfPointerStagingCode,
				"{{CopyingPointers}}", pointerCopyingCode,
				"{{CopyingSliceOfPointers}}", sliceOfPointerCopyingCode,
				"{{UnstagingPointers}}", pointerUnstagingCode,
				"{{UnstagingSliceOfPointers}}", sliceOfPointerUnstagingCode)

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

	file, err := os.Create(filepath.Join(pkgPath, "gong_graph.go"))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)
}
