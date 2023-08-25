package golang

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

// insertion points are places where the code is
// generated per gong struct
type ModelGongGraphStructInsertionId int

const (
	ModelGongGraphStructInsertionIsStaged ModelGongGraphStructInsertionId = iota
	ModelGongGraphStructInsertionIsStagedPerStruct
	ModelGongGraphStructInsertionStageBranch
	ModelGongGraphStructInsertionStageBranchPerStruct
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
	func (stage *StageStruct) IsStaged{{Structname}}({{structname}} *{{Structname}}) (ok bool) {

		_, ok = stage.{{Structname}}s[{{structname}}]
	
		return
	}
`,
	ModelGongGraphStructInsertionStageBranch: `
	case *{{Structname}}:
		stage.StageBranch{{Structname}}(target)
`,
	ModelGongGraphStructInsertionStageBranchPerStruct: `
func (stage *StageStruct) StageBranch{{Structname}}({{structname}} *{{Structname}}) {

	// check if instance is already staged
	if IsStaged(stage, {{structname}}) {
		return
	}

	{{structname}}.Stage(stage)

	//insertion point for the staging of instances referenced by pointers{{StagingPointers}}

	//insertion point for the staging of instances referenced by slice of pointers{{StagingSliceOfPointers}}

}
`,
	ModelGongGraphStructInsertionUnstageBranch: `
	case *{{Structname}}:
		stage.UnstageBranch{{Structname}}(target)
`,
	ModelGongGraphStructInsertionUnstageBranchPerStruct: `
func (stage *StageStruct) UnstageBranch{{Structname}}({{structname}} *{{Structname}}) {

	// check if instance is already staged
	if ! IsStaged(stage, {{structname}}) {
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
	mdlPkg *models.ModelPkg,
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
	for _, _struct := range mdlPkg.GongStructs {
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
			pointerUnstagingCode := ""
			sliceOfPointerUnstagingCode := ""

			for _, field := range gongStruct.Fields {

				switch field := field.(type) {
				case *models.PointerToGongStructField:
					pointerStagingCode += models.Replace2(
						GongGraphFileFieldFieldSubTemplateCode[GongGraphFileFieldSubTmplStagePointerField],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name)
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

			sliceOfPointerUnstagingCode = models.Replace2(sliceOfPointerUnstagingCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			pointerUnstagingCode = models.Replace2(pointerUnstagingCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			generatedCodeFromSubTemplate := models.Replace6(ModelGongGraphStructSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{StagingPointers}}", pointerStagingCode,
				"{{StagingSliceOfPointers}}", sliceOfPointerStagingCode,
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
