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

const GongReverse = `// generated code - do not edit
package models

func GetReverseFieldOwnerName(
	stage *StageStruct,
	instance any,
	reverseField *ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point{{` + string(rune(GetReverseFieldOwnerNameSwitch)) + `}}
	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T Gongstruct](
	stage *StageStruct,
	instance *T,
	reverseField *ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point{{` + string(rune(GetReverseFieldOwnerSwitch)) + `}}
	default:
		_ = inst
	}
	return res
}
`

type GetReverseFieldOwnerNameId int

const (
	GetReverseFieldOwnerNameSwitch GetReverseFieldOwnerNameId = iota
	GetReverseFieldOwnerSwitch
	GetReverseFieldOwnerNameNb
)

var GetReverseFieldOwnerNameSubTemplateCode map[GetReverseFieldOwnerNameId]string = // new line
map[GetReverseFieldOwnerNameId]string{
	GetReverseFieldOwnerNameSwitch: `
	case *{{Structname}}:
		switch reverseField.GongstructName {
		// insertion point{{fieldToFormCodeName}}
		}
`,
	GetReverseFieldOwnerSwitch: `
	case *{{Structname}}:
		switch reverseField.GongstructName {
		// insertion point{{fieldToFormCode}}
		}
`,
}

type GetReverseFieldOwnerNameSubTemplateId int

const (
	GetReverseFieldOwnerNameSwitchCode GetReverseFieldOwnerNameSubTemplateId = iota
	GetReverseFieldOwnerSwitchCode
	GetReverseFieldOwnerNameMasterSwitchCodeStart
	GetReverseFieldOwnerNameMasterSwitchCodeEnd
)

var GetReverseFieldOwnerNameSubSubTemplateCode map[GetReverseFieldOwnerNameSubTemplateId]string = // declaration of the sub templates
map[GetReverseFieldOwnerNameSubTemplateId]string{

	GetReverseFieldOwnerNameMasterSwitchCodeStart: `
		case "{{AssocStructName}}":
			switch reverseField.Fieldname {`,
	GetReverseFieldOwnerNameSwitchCode: `
			case "{{FieldName}}":
				if _{{assocStructName}}, ok := stage.{{AssocStructName}}_{{FieldName}}_reverseMap[inst]; ok {
					res = _{{assocStructName}}.Name
				}`,
	GetReverseFieldOwnerSwitchCode: `
			case "{{FieldName}}":
				res = stage.{{AssocStructName}}_{{FieldName}}_reverseMap[inst]`,
	GetReverseFieldOwnerNameMasterSwitchCodeEnd: `
			}`,
}

func CodeGeneratorGongReverse(
	modelPkg *models.ModelPkg,
	pkgName string,
	pkgPath string,
	pkgGoPath string) {

	// generate the typescript file
	codeGO := GongReverse

	subStructCodes := make(map[GetReverseFieldOwnerNameId]string)
	for subStructTemplate := range GetReverseFieldOwnerNameSubTemplateCode {
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

		fieldNames := make(map[string]any, 0)
		_ = fieldNames

		for subStructTemplate := range GetReverseFieldOwnerNameSubTemplateCode {

			fieldToFormCode := ""
			fieldToFormCodeName := ""

			//
			// Parse all fields from other structs that points to this struct
			//
			for _, __struct := range gongStructs {

				// count the number of cases
				nbCases := 0
				for _, field := range __struct.Fields {
					switch field := field.(type) {
					case *models.SliceOfPointerToGongStructField:
						if field.GongStruct == gongStruct {
							nbCases += 1
						}
						if field.GongStruct == gongStruct {
							nbCases += 1
						}
					}
				}

				if nbCases == 0 {
					continue
				}

				fieldToFormCode += models.Replace1(
					GetReverseFieldOwnerNameSubSubTemplateCode[GetReverseFieldOwnerNameMasterSwitchCodeStart],
					"{{AssocStructName}}", __struct.Name)
				fieldToFormCodeName += models.Replace1(
					GetReverseFieldOwnerNameSubSubTemplateCode[GetReverseFieldOwnerNameMasterSwitchCodeStart],
					"{{AssocStructName}}", __struct.Name)

				for _, field := range __struct.Fields {

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

						if field.GongStruct == gongStruct {
							fieldToFormCode += models.Replace2(
								GetReverseFieldOwnerNameSubSubTemplateCode[GetReverseFieldOwnerSwitchCode],
								"{{AssocStructName}}", __struct.Name,
								"{{FieldName}}", fieldNameForReverseMapField)
						}
						if field.GongStruct == gongStruct {
							fieldToFormCodeName += models.Replace3(
								GetReverseFieldOwnerNameSubSubTemplateCode[GetReverseFieldOwnerNameSwitchCode],
								"{{AssocStructName}}", __struct.Name,
								"{{assocStructName}}", strings.ToLower(__struct.Name),
								"{{FieldName}}", fieldNameForReverseMapField)
						}
					}
				}

				fieldToFormCode += models.Replace1(
					GetReverseFieldOwnerNameSubSubTemplateCode[GetReverseFieldOwnerNameMasterSwitchCodeEnd],
					"{{AssocStructName}}", __struct.Name)
				fieldToFormCodeName += models.Replace1(
					GetReverseFieldOwnerNameSubSubTemplateCode[GetReverseFieldOwnerNameMasterSwitchCodeEnd],
					"{{AssocStructName}}", __struct.Name)
			}

			fieldToFormCode = models.Replace2(fieldToFormCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			fieldToFormCodeName = models.Replace2(fieldToFormCodeName,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			generatedCodeFromSubTemplate := models.Replace4(GetReverseFieldOwnerNameSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{fieldToFormCode}}", fieldToFormCode,
				"{{fieldToFormCodeName}}", fieldToFormCodeName)

			subStructCodes[subStructTemplate] += generatedCodeFromSubTemplate
		}

	}

	// substitutes {{<<insertionPerStructId points>>}} stuff with generated code
	for insertionPerStructId := GetReverseFieldOwnerNameId(0); insertionPerStructId < GetReverseFieldOwnerNameNb; insertionPerStructId++ {
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

	file, err := os.Create(filepath.Join(pkgPath, "gong_reverse.go"))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)

}

func removeDuplicates(strList []models.FieldInterface) []models.FieldInterface {
	list := []models.FieldInterface{}
	encountered := map[string]bool{}

	for _, item := range strList {
		if encountered[item.GetName()] == true {
			continue
		}
		encountered[item.GetName()] = true
		list = append(list, item)
	}

	return list
}
