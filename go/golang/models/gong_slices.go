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

const GongSliceTemplate = `// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct{{` + string(rune(GongSliceReverseMapCompute)) + `}}
}
`

type GongSliceGongstructInsertionId int

const (
	GongSliceCase GongSliceGongstructInsertionId = iota
	GongSliceReverseMapCompute
	GongSliceGongstructInsertionNb
)

var GongSliceGongstructSubTemplateCode map[GongSliceGongstructInsertionId]string = // new line
map[GongSliceGongstructInsertionId]string{
	GongSliceCase: `
	case *{{Structname}}:
		// insertion point per field{{perFieldCode}}
`,
	GongSliceReverseMapCompute: `
	// Compute reverse map for named struct {{Structname}}
	// insertion point per field{{sliceOfPointerFieldReverseMapComputationCode}}
`,
}

type GongSliceSubTemplateId int

const (
	GongSliceSubTmplSliceOfPointersToStruct GongSliceSubTemplateId = iota
	GongSliceSubTmplSliceOfPointersReverseMapComputation
)

var GongSliceFileFieldFieldSubTemplateCode map[GongSliceSubTemplateId]string = // declaration of the sub templates
map[GongSliceSubTemplateId]string{

	GongSliceSubTmplSliceOfPointersToStruct: `
		if fieldName == "{{FieldName}}" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*{{Structname}}) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*{{Structname}})
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.{{FieldName}}).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.{{FieldName}} = _inferedTypeInstance.{{FieldName}}[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.{{FieldName}} =
								append(_inferedTypeInstance.{{FieldName}}, any(fieldInstance).(*{{AssociationStructName}}))
						}
					}
				}
			}
		}`,
	GongSliceSubTmplSliceOfPointersReverseMapComputation: `
	clear(stage.{{Structname}}_{{FieldNameForReverseMapField}}_reverseMap)
	stage.{{Structname}}_{{FieldNameForReverseMapField}}_reverseMap = make(map[*{{AssociationStructName}}]*{{Structname}})
	for {{structname}} := range stage.{{Structname}}s {
		_ = {{structname}}
		for _, _{{associationStructName}} := range {{structname}}.{{FieldName}} {
			stage.{{Structname}}_{{FieldNameForReverseMapField}}_reverseMap[_{{associationStructName}}] = {{structname}}
		}
	}`,
}

func CodeGeneratorModelGongSlice(
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
	codeGO := GongSliceTemplate

	subStructCodes := make(map[GongSliceGongstructInsertionId]string)
	for subStructTemplate := range GongSliceGongstructSubTemplateCode {
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

		for subStructTemplate := range GongSliceGongstructSubTemplateCode {

			perFieldCode := ""
			sliceOfPointerFieldReverseMapComputationCode := ""

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
				case *models.SliceOfPointerToGongStructField:
					perFieldCode += models.Replace3(
						GongSliceFileFieldFieldSubTemplateCode[GongSliceSubTmplSliceOfPointersToStruct],
						"{{FieldName}}", field.Name,
						"{{AssociationStructName}}", field.GongStruct.Name,
						"{{associationStructName}}", strings.ToLower(field.GongStruct.Name))
					sliceOfPointerFieldReverseMapComputationCode += models.Replace4(
						GongSliceFileFieldFieldSubTemplateCode[GongSliceSubTmplSliceOfPointersReverseMapComputation],
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

			sliceOfPointerFieldReverseMapComputationCode = models.Replace2(sliceOfPointerFieldReverseMapComputationCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			generatedCodeFromSubTemplate := models.Replace4(GongSliceGongstructSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{perFieldCode}}", perFieldCode,
				"{{sliceOfPointerFieldReverseMapComputationCode}}", sliceOfPointerFieldReverseMapComputationCode)

			subStructCodes[subStructTemplate] += generatedCodeFromSubTemplate
		}

	}

	// substitutes {{<<insertionPerStructId points>>}} stuff with generated code
	for insertionPerStructId := GongSliceGongstructInsertionId(0); insertionPerStructId < GongSliceGongstructInsertionNb; insertionPerStructId++ {
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

	file, err := os.Create(filepath.Join(pkgPath, "../models/gong_slices.go"))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)

}
