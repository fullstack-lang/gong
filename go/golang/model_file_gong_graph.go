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
	ModelGongGraphStructInsertionIsStageBranch
	ModelGongGraphStructInsertionIsStageBranchPerStruct
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
	ModelGongGraphStructInsertionIsStageBranch: `
	case *{{Structname}}:
		stage.StageBranch{{Structname}}(target)
`,
	ModelGongGraphStructInsertionIsStageBranchPerStruct: `
func (stage *StageStruct) StageBranch{{Structname}}({{structname}} *{{Structname}}) {

	// check if instance is already staged
	if IsStaged(stage, {{structname}}) {
		return
	}

	{{structname}}.Stage()

	//insertion point for the staging of instances referenced by pointers{{StagingPointers}}

	//insertion point for the staging of instances referenced by slice of pointers{{StagingSliceOfPointers}}

}
`,
}

type GongGraphFilePerStructSubTemplateId int

const (
	GongGraphFileFieldSubTmplSetPointerField GongGraphFilePerStructSubTemplateId = iota
	GongGraphFileFieldSubTmplSetSliceOfPointersField
)

var GongGraphFileFieldFieldSubTemplateCode map[GongGraphFilePerStructSubTemplateId]string = // declaration of the sub templates
map[GongGraphFilePerStructSubTemplateId]string{

	GongGraphFileFieldSubTmplSetPointerField: `
	if {{structname}}.{{FieldName}} != nil {
		StageBranch(stage, {{structname}}.{{FieldName}})
	}`,
	GongGraphFileFieldSubTmplSetSliceOfPointersField: `
	for _, _{{assocstructname}} := range {{structname}}.{{FieldName}} {
		StageBranch(stage, _{{assocstructname}})
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

			for _, field := range gongStruct.Fields {

				switch field := field.(type) {
				case *models.PointerToGongStructField:
					pointerStagingCode += models.Replace2(
						GongGraphFileFieldFieldSubTemplateCode[GongGraphFileFieldSubTmplSetPointerField],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name)
				case *models.SliceOfPointerToGongStructField:
					sliceOfPointerStagingCode += models.Replace3(
						GongGraphFileFieldFieldSubTemplateCode[GongGraphFileFieldSubTmplSetSliceOfPointersField],
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

			generatedCodeFromSubTemplate := models.Replace4(ModelGongGraphStructSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{StagingPointers}}", pointerStagingCode,
				"{{StagingSliceOfPointers}}", sliceOfPointerStagingCode)

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
