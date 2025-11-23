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

const GongCleanTemplate = `// generated code - do not edit
package models

// Clean computes the reverse map, for all intances, for all clean to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) Clean() {
	// insertion point per named struct{{` + string(rune(GongCleanRangeElements)) + `}}
}
`

type GongCleanGongstructInsertionId int

const (
	GongCleanRangeElements GongCleanGongstructInsertionId = iota
	GongCleanGongstructInsertionNb
)

var GongCleanGongstructSubTemplateCode map[GongCleanGongstructInsertionId]string = // new line
map[GongCleanGongstructInsertionId]string{
	GongCleanRangeElements: `
	// Compute reverse map for named struct {{Structname}}
	for {{structname}} := range stage.{{Structname}}s {
		_ = {{structname}}
		// insertion point per field{{cleanOfSliceOfPointers}}
		// insertion point per field{{cleanOfPointer}}
	}
`,
}

type GongCleanSubTemplateId int

const (
	GongCleanSubTmplCleanPointer GongCleanSubTemplateId = iota
	GongCleanSubTmplCleanOfSlicePointers
)

var GongCleanFileFieldFieldSubTemplateCode map[GongCleanSubTemplateId]string = // declaration of the sub templates
map[GongCleanSubTemplateId]string{

	GongCleanSubTmplCleanPointer: `
		if !IsStaged(stage, {{structname}}.{{FieldName}}) {
			{{structname}}.{{FieldName}} = nil
		}`,
	GongCleanSubTmplCleanOfSlicePointers: `
		var _{{FieldName}} []*{{AssociationStructName}}
		for _, _{{associationStructName}} := range {{structname}}.{{FieldName}} {
			if IsStaged(stage, _{{associationStructName}}) {
			 	_{{FieldName}} = append(_{{FieldName}}, _{{associationStructName}})
			}
		}
		{{structname}}.{{FieldName}} = _{{FieldName}}`,
}

func CodeGeneratorModelGongClean(
	modelPkg *models.ModelPkg,
	pkgName string,
	pkgPath string,
	pkgGoPath string) {

	// this code is not robust to empty models
	// map[Gongstruct]any cannot compile
	if len(modelPkg.GongStructs) == 0 {
		return
	}

	// generate the typescript file
	codeGO := GongCleanTemplate

	subStructCodes := make(map[GongCleanGongstructInsertionId]string)
	for subStructTemplate := range GongCleanGongstructSubTemplateCode {
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

		for subStructTemplate := range GongCleanGongstructSubTemplateCode {

			perFieldCode := ""
			cleanOfSliceOfPointers := ""
			cleanOfPointer := ""

			for _, field := range gongStruct.Fields {

				fieldName := field.GetName()
				fieldNameForReverseMapField := fieldName

				// in case of a field within an anonymous struct, one needs
				// to strip the prefix
				fieldNameSplitted := strings.Split(fieldName, ".")
				isWithinAnonymousStruct := len(fieldNameSplitted) > 1
				if isWithinAnonymousStruct {
					fieldNameForReverseMapField = fieldNameSplitted[0] + "_" + fieldNameSplitted[1]
				}

				switch field := field.(type) {
				case *models.PointerToGongStructField:
					cleanOfPointer += models.Replace4(
						GongCleanFileFieldFieldSubTemplateCode[GongCleanSubTmplCleanPointer],
						"{{FieldNameForReverseMapField}}", fieldNameForReverseMapField,
						"{{FieldName}}", fieldName,
						"{{AssociationStructName}}", field.GongStruct.Name,
						"{{associationStructName}}", strings.ToLower(field.GongStruct.Name))
				case *models.SliceOfPointerToGongStructField:
					cleanOfSliceOfPointers += models.Replace4(
						GongCleanFileFieldFieldSubTemplateCode[GongCleanSubTmplCleanOfSlicePointers],
						"{{FieldNameForReverseMapField}}", fieldNameForReverseMapField,
						"{{FieldName}}", fieldName,
						"{{AssociationStructName}}", field.GongStruct.Name,
						"{{associationStructName}}", strings.ToLower(field.GongStruct.Name))
				default:
				}

			}

			perFieldCode = models.Replace2(perFieldCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			cleanOfSliceOfPointers = models.Replace2(cleanOfSliceOfPointers,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			cleanOfPointer = models.Replace2(cleanOfPointer,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			generatedCodeFromSubTemplate := models.Replace5(GongCleanGongstructSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{perFieldCode}}", perFieldCode,
				"{{cleanOfSliceOfPointers}}", cleanOfSliceOfPointers,
				"{{cleanOfPointer}}", cleanOfPointer)

			subStructCodes[subStructTemplate] += generatedCodeFromSubTemplate
		}

	}

	// substitutes {{<<insertionPerStructId points>>}} stuff with generated code
	for insertionPerStructId := GongCleanGongstructInsertionId(0); insertionPerStructId < GongCleanGongstructInsertionNb; insertionPerStructId++ {
		toReplace := "{{" + string(rune(insertionPerStructId)) + "}}"
		codeGO = strings.ReplaceAll(codeGO, toReplace, subStructCodes[insertionPerStructId])
	}

	caserEnglish := cases.Title(language.English)
	codeGO = models.Replace5(codeGO,
		"{{PkgName}}", pkgName,
		"{{TitlePkgName}}", caserEnglish.String(pkgName),
		"{{pkgname}}", strings.ToLower(pkgName),
		"	 | ", "	", // for the replacement of the of the first bar in the Gongstruct Type def

		"{{PkgPathRoot}}", strings.ReplaceAll(pkgGoPath, "/models", ""),
	)

	file, err := os.Create(filepath.Join(pkgPath, string(models.GeneratedGongCleanGoFilePath)))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)

}
