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
package {{PkgGoName}}

import "time"

// CleanSlice is the Stage method that removes unstaged elements from a slice of pointers.
func (stage *Stage) CleanSlice[T GongstructPtr](slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if stage.IsStaged(element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// CleanPointer is the Stage method that sets the pointer to nil if the referenced element is not staged.
func (stage *Stage) CleanPointer[T GongstructPtr](element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !stage.IsStaged(*element) {
		*element = zero
		modified = true
		return
	}
	return
}

type GongCleaner interface {
	GongClean(stage *Stage) (modified bool)
}

// insertion point per named struct{{` + string(rune(GongCleanRangeElements)) + `}}
// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		if cleaner, ok := any(instance).(GongCleaner); ok {
			modified = cleaner.GongClean(stage) || modified
		}
	}
	if modified {
		if stage.probeIF != nil {
			stage.probeIF.AddNotification(time.Now(), "Stage clean generated a modification")
		}
	}
	return
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
func ({{structname}} *{{Structname}}) GongClean(stage *Stage) (modified bool) {
	// insertion point per field{{cleanOfSliceOfPointers}}
	// insertion point per field{{cleanOfPointer}}
	return
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
	modified = stage.CleanPointer(&{{structname}}.{{FieldName}}) || modified`,
	GongCleanSubTmplCleanOfSlicePointers: `
	modified = stage.CleanSlice(&{{structname}}.{{FieldName}}) || modified`,
}

func CodeGeneratorModelGongClean(
	modelPkg *models.ModelPkg,
	pkgName string,
	pkgPath string,
	pkgGoPath string,
) {
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

			if cleanOfSliceOfPointers == "" && cleanOfPointer == "" {
				continue
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
	for insertionPerStructId := range GongCleanGongstructInsertionNb {
		toReplace := "{{" + string(rune(insertionPerStructId)) + "}}"
		codeGO = strings.ReplaceAll(codeGO, toReplace, subStructCodes[insertionPerStructId])
	}

	var pkgPathRoot string
	if before, _, ok := strings.Cut(pkgGoPath, "/go/models"); ok {
		pkgPathRoot = before + "/go"
	} else {
		pkgPathRoot = strings.ReplaceAll(pkgGoPath, "/models", "")
	}

	caserEnglish := cases.Title(language.English)
	codeGO = models.Replace(codeGO,
		"{{PkgName}}", pkgName,
		"{{TitlePkgName}}", caserEnglish.String(pkgName),
		"{{pkgname}}", strings.ToLower(pkgName),
		"{{PkgGoName}}", modelPkg.PkgGoName,
		"	 | ", "	", // for the replacement of the of the first bar in the Gongstruct Type def
		"{{PkgPathRoot}}", pkgPathRoot,
	)

	file, err := os.Create(filepath.Join(pkgPath, string(models.GeneratedGongCleanGoFilePath)))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)
}
