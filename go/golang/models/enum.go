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

// insertion code for all enums
type ModelGongEnumInsertionId int

const (
	// iota + 40 is to separate the insertion code of gongstruct from insertion code of gongenum
	ModelGongEnumUtilityFunctions ModelGongEnumInsertionId = iota + 40
	ModelGongEnumInsertionsNb
)

var ModelGongEnumSubTemplateCode map[ModelGongEnumInsertionId]string = // new line
map[ModelGongEnumInsertionId]string{
	ModelGongEnumUtilityFunctions: `
// Utility function for {{EnumName}}
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func ({{enumName}} {{EnumName}}) To{{Type}}() (res {{type}}) {

	// migration of former implementation of enum
	switch {{enumName}} {
	// insertion code per enum code{{ToStringPerCodeCode}}
	}
	return
}

func ({{enumName}} *{{EnumName}}) From{{Type}}(input {{type}}) (err error) {

	switch input {
	// insertion code per enum code{{FromStringPerCodeCode}}
	default:
		return errUnkownEnum
	}
	return
}

func ({{enumName}} *{{EnumName}}) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code{{FromCodeStringPerCodeCode}}
	default:
		return errUnkownEnum
	}
	return
}

func ({{enumName}} *{{EnumName}}) ToCodeString() (res string) {

	switch *{{enumName}} {
	// insertion code per enum code{{ToCodeStringPerCodeCode}}
	}
	return
}
`,
}

// gongenum value template
type GongModelEnumValueSubTemplateId int

const (
	GongModelEnumValueFromString GongModelEnumValueSubTemplateId = iota
	GongModelEnumValueFromCodeString
	GongModelEnumValueToString
	GongModelEnumValueToCodeString
)

var GongModelEnumValueSubTemplateCode map[GongModelEnumValueSubTemplateId]string = // declaration of the sub templates
map[GongModelEnumValueSubTemplateId]string{

	GongModelEnumValueFromString: `
	case {{GongEnumValue}}:
		*{{enumName}} = {{GongEnumCode}}`,
	GongModelEnumValueFromCodeString: `
	case "{{GongEnumCode}}":
		*{{enumName}} = {{GongEnumCode}}`,
	GongModelEnumValueToString: `
	case {{GongEnumCode}}:
		res = {{GongEnumValue}}`,
	GongModelEnumValueToCodeString: `
	case {{GongEnumCode}}:
		res = "{{GongEnumCode}}"`,
}

func CodeGeneratorModelGongEnum(
	mdlPkg *models.ModelPkg,
	pkgName string,
	pkgPath string) {

	// generate the typescript file
	codeGO := ModelGongEnumFileTemplate

	subEnumCodes := make(map[ModelGongEnumInsertionId]string)
	for subEnumTemplate := range ModelGongEnumSubTemplateCode {
		subEnumCodes[subEnumTemplate] = ""
	}

	// sort gong enums per name (for reproductibility)
	gongEnums := []*models.GongEnum{}
	for _, _enum := range mdlPkg.GongEnums {
		gongEnums = append(gongEnums, _enum)
	}
	sort.Slice(gongEnums[:], func(i, j int) bool {
		return gongEnums[i].Name < gongEnums[j].Name
	})

	for _, gongEnum := range gongEnums {

		for subEnumTemplate := range ModelGongEnumSubTemplateCode {

			codeFromStringPerGongValue := ""
			codeFromCodeStringPerGongValue := ""
			codeToStringPerGongValue := ""
			codeToCodeStringPerGongValue := ""

			for _, enumValue := range gongEnum.GongEnumValues {
				codeFromStringPerGongValue += models.Replace2(GongModelEnumValueSubTemplateCode[GongModelEnumValueFromString],
					"{{GongEnumValue}}", enumValue.Value,
					"{{GongEnumCode}}", enumValue.Name)
				codeFromCodeStringPerGongValue += models.Replace1(GongModelEnumValueSubTemplateCode[GongModelEnumValueFromCodeString],
					"{{GongEnumCode}}", enumValue.Name)
				codeToStringPerGongValue += models.Replace2(GongModelEnumValueSubTemplateCode[GongModelEnumValueToString],
					"{{GongEnumValue}}", enumValue.Value,
					"{{GongEnumCode}}", enumValue.Name)

				codeToCodeStringPerGongValue += models.Replace2(GongModelEnumValueSubTemplateCode[GongModelEnumValueToCodeString],
					"{{GongEnumValue}}", enumValue.Value,
					"{{GongEnumCode}}", enumValue.Name)
			}

			generatedCodeFromSubTemplate := models.Replace4(ModelGongEnumSubTemplateCode[subEnumTemplate],
				"{{ToStringPerCodeCode}}", codeToStringPerGongValue,
				"{{FromStringPerCodeCode}}", codeFromStringPerGongValue,
				"{{FromCodeStringPerCodeCode}}", codeFromCodeStringPerGongValue,
				"{{ToCodeStringPerCodeCode}}", codeToCodeStringPerGongValue)

			var typeOfEnumAsString string
			if gongEnum.Type == models.String {
				typeOfEnumAsString = "String"
			} else {
				typeOfEnumAsString = "Int"
			}

			generatedCodeFromSubTemplate = models.Replace4(generatedCodeFromSubTemplate,
				"{{enumName}}", strings.ToLower(gongEnum.Name),
				"{{EnumName}}", gongEnum.Name,
				"{{Type}}", typeOfEnumAsString,
				"{{type}}", strings.ToLower(typeOfEnumAsString))

			subEnumCodes[subEnumTemplate] += generatedCodeFromSubTemplate
		}
	}

	// substitutes {{<<insertionPerEnumId points>>}} stuff with generated code
	for insertionPerEnumId := ModelGongEnumUtilityFunctions; insertionPerEnumId < ModelGongEnumInsertionsNb; insertionPerEnumId++ {
		toReplace := "{{" + string(rune(insertionPerEnumId)) + "}}"
		codeGO = strings.ReplaceAll(codeGO, toReplace, subEnumCodes[insertionPerEnumId])
	}

	caserEnglish := cases.Title(language.English)
	codeGO = models.Replace4(codeGO,
		"{{PkgName}}", pkgName,
		"{{TitlePkgName}}", caserEnglish.String(pkgName),
		"{{pkgname}}", strings.ToLower(pkgName),
		"	 | ", "	", // for the replacement of the of the first bar in the Gongstruct Type def
	)

	file, err := os.Create(filepath.Join(pkgPath, "gong_enum.go"))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)

}
