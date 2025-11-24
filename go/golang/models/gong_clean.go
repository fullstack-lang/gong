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

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice []T) []T {
	if slice == nil {
		return nil
	}

	var cleanedSlice []T
	for _, element := range slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	return cleanedSlice
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element T) T {
	if !IsStagedPointerToGongstruct(stage, element) {
		var zero T
		return zero
	}
	return element
}

// insertion point per named struct{{` + string(rune(GongCleanRangeElements)) + `}}
// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() {
	for _, instance := range stage.GetInstances() {
		instance.GongClean(stage)
	}
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
// Clean garbage collect unstaged instances that are referenced by {{Structname}}
func ({{structname}} *{{Structname}}) GongClean(stage *Stage) {
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
	{{structname}}.{{FieldName}} = GongCleanPointer(stage, {{structname}}.{{FieldName}})`,
	GongCleanSubTmplCleanOfSlicePointers: `
	{{structname}}.{{FieldName}} = GongCleanSlice(stage, {{structname}}.{{FieldName}})`,
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
