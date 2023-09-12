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

// EvictInOtherSlices allows for adherance between
// the gong association model and go.
//
// Says you have a Astruct struct with a slice field "anarrayofb []*Bstruct"
//
// go allows many Astruct instance to have the anarrayofb field that have the
// same pointers. go slices are MANY-MANY association.
//
// With gong it is only ZERO-ONE-MANY associations, a Bstruct can be pointed only
// once by an Astruct instance through a given field. This follows the requirement
// that gong is suited for full stack programming and therefore the association
// is encoded as a reverse pointer (not as a joint table). In gong, a named struct
// is translated in a table and each table is a named struct.
//
// EvictInOtherSlices removes the fields instances from other
// fields of other instance
//
// Note : algo is in O(N)log(N) of nb of Astruct and Bstruct instances
func EvictInOtherSlices[OwningType PointerToGongstruct, FieldType PointerToGongstruct](
	stage *StageStruct,
	owningInstance OwningType,
	sliceField []FieldType,
	fieldName string) {

	// create a map of the field elements
	setOfFieldInstances := make(map[FieldType]any, 0)
	for _, fieldInstance := range sliceField {
		setOfFieldInstances[fieldInstance] = true
	}

	switch owningInstanceInfered := any(owningInstance).(type) {
	// insertion point{{` + string(rune(GongSliceCase)) + `}}
	default:
		_ = owningInstanceInfered // to avoid "declared and not used" error if no named struct has slices
	}
}
`

type GongSliceGongstructInsertionId int

const (
	GongSliceCase GongSliceGongstructInsertionId = iota
	GongSliceGongstructInsertionNb
)

var GongSliceGongstructSubTemplateCode map[GongSliceGongstructInsertionId]string = // new line
map[GongSliceGongstructInsertionId]string{
	GongSliceCase: `
	case *{{Structname}}:
		// insertion point per field{{perFieldCode}}
`,
}

type GongSliceSubTemplateId int

const (
	GongSliceSubTmplSliceOfPointersToStruct GongSliceSubTemplateId = iota
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
}

func CodeGeneratorModelGongSlice(
	mdlPkg *models.ModelPkg,
	pkgName string,
	pkgPath string,
	pkgGoPath string) {

	// generate the typescript file
	codeGO := GongSliceTemplate

	subStructCodes := make(map[GongSliceGongstructInsertionId]string)
	for subStructTemplate := range GongSliceGongstructSubTemplateCode {
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

		for subStructTemplate := range GongSliceGongstructSubTemplateCode {

			perFieldCode := ""

			for _, field := range gongStruct.Fields {

				switch field := field.(type) {
				case *models.SliceOfPointerToGongStructField:
					perFieldCode += models.Replace2(
						GongSliceFileFieldFieldSubTemplateCode[GongSliceSubTmplSliceOfPointersToStruct],
						"{{FieldName}}", field.Name,
						"{{AssociationStructName}}", field.GongStruct.Name)
				default:
				}

			}

			perFieldCode = models.Replace2(perFieldCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			generatedCodeFromSubTemplate := models.Replace3(GongSliceGongstructSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{perFieldCode}}", perFieldCode)

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
