package orm

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

const GetReverseFieldOwnerName = `// generated code - do not edit
package orm

import (
	"{{PkgPathRoot}}/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseFieldName string) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point{{` + string(rune(GetReverseFieldOwnerNameSwitch)) + `}}
	default:
		_ = inst
	}
	return
}
`

type GetReverseFieldOwnerNameId int

const (
	GetReverseFieldOwnerNameSwitch GetReverseFieldOwnerNameId = iota
	GetReverseFieldOwnerNameNb
)

var GetReverseFieldOwnerNameSubTemplateCode map[GetReverseFieldOwnerNameId]string = // new line
map[GetReverseFieldOwnerNameId]string{
	GetReverseFieldOwnerNameSwitch: `
	case *models.{{Structname}}:
		tmp := GetInstanceDBFromInstance[models.{{Structname}}, {{Structname}}DB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point{{fieldToFormCode}}
		}
`,
}

type GetReverseFieldOwnerNameSubTemplateId int

const (
	GetReverseFieldOwnerNameSwitchCode GetReverseFieldOwnerNameSubTemplateId = iota
)

var GetReverseFieldOwnerNameSubSubTemplateCode map[GetReverseFieldOwnerNameSubTemplateId]string = // declaration of the sub templates
map[GetReverseFieldOwnerNameSubTemplateId]string{

	GetReverseFieldOwnerNameSwitchCode: `
		case "{{FieldName}}":
			if tmp.{{AssocStructName}}_{{FieldName}}DBID.Int64 != 0 {
				id := uint(tmp.{{AssocStructName}}_{{FieldName}}DBID.Int64)
				reservePointerTarget := backRepo.BackRepo{{AssocStructName}}.Map_{{AssocStructName}}DBID_{{AssocStructName}}Ptr[id]
				res = reservePointerTarget.Name
			}`,
}

func CodeGeneratorGetReverseFieldOwnerName(
	mdlPkg *models.ModelPkg,
	pkgName string,
	pkgPath string,
	pkgGoPath string) {

	// generate the typescript file
	codeGO := GetReverseFieldOwnerName

	subStructCodes := make(map[GetReverseFieldOwnerNameId]string)
	for subStructTemplate := range GetReverseFieldOwnerNameSubTemplateCode {
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

		fieldNames := make(map[string]any, 0)

		for subStructTemplate := range GetReverseFieldOwnerNameSubTemplateCode {

			fieldToFormCode := ""

			//
			// Parse all fields from other structs that points to this struct
			//
			for _, __struct := range gongStructs {

				for _, field := range __struct.Fields {
					switch field := field.(type) {
					case *models.SliceOfPointerToGongStructField:

						if field.GongStruct == gongStruct {

							if _, ok := fieldNames[field.GetName()]; ok {
								fieldNames[field.GetName()] = true

								fieldToFormCode += models.Replace2(
									GetReverseFieldOwnerNameSubSubTemplateCode[GetReverseFieldOwnerNameSwitchCode],
									"{{AssocStructName}}", __struct.Name,
									"{{FieldName}}", field.GetName())
							}
						}
					}
				}
			}

			fieldToFormCode = models.Replace2(fieldToFormCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			generatedCodeFromSubTemplate := models.Replace3(GetReverseFieldOwnerNameSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{fieldToFormCode}}", fieldToFormCode)

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

	file, err := os.Create(filepath.Join(pkgPath, "../orm/get_reverse_field_owner_name.go"))
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
