package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const NgEnumTemplateTS = `// generated from ng_file_enum.ts.go
export enum {{EnumName}} {
	// insertion point	{{` + string(rune(NgEnumInsertionPointEnumDeclaration)) + `}}
}

export interface {{EnumName}}Select {
	value: {{type}};
	viewValue: string;
}

export const {{EnumName}}List: {{EnumName}}Select[] = [ // insertion point	{{` + string(rune(NgEnumValuesInsertionPointDeclarationForPullDownSelect)) + `}}
];
`

// insertion points
type NgEnumInsertionPoint int

const (
	NgEnumInsertionPointEnumDeclaration NgEnumInsertionPoint = iota
	NgEnumValuesInsertionPointDeclarationForPullDownSelect
	NgEnumNbInsertionPoints
)

//
// Sub Templates
//
type NgEnumSubTemplate int

const (
	NgEnumDeclaration NgEnumSubTemplate = iota
	NgEnumDeclarationForPullDownSelect
)

var NgEnumHtmlSubTemplateCode map[NgEnumSubTemplate]string = map[NgEnumSubTemplate]string{
	NgEnumDeclaration: `
	{{ConstName}} = {{ConstValue}},`,

	NgEnumDeclarationForPullDownSelect: `
	{ value: {{EnumName}}.{{ConstName}}, viewValue: {{ConstValue}} },`,
}

// MultiCodeGeneratorNgEnum parses mdlPkg and generates the code for the
// Enum components
func CodeGeneratorNgEnum(
	mdlPkg *ModelPkg,
	pkgName string,
	matTargetPath string,
	pkgGoPath string) {

	for _, gongEnum := range mdlPkg.GongEnums {

		// generate the typescript file
		codeTS := NgEnumTemplateTS

		insertions := make(map[NgEnumInsertionPoint]string)
		for insertion := NgEnumInsertionPoint(0); insertion < NgEnumNbInsertionPoints; insertion++ {
			insertions[insertion] = ""
		}

		for _, value := range gongEnum.GongEnumValues {

			var viewValueOfEnum string
			if gongEnum.Type == String {
				viewValueOfEnum = value.Value
			} else {
				viewValueOfEnum = fmt.Sprintf("\"%s\"", value.Name)
			}

			insertions[NgEnumInsertionPointEnumDeclaration] += Replace2(
				NgEnumHtmlSubTemplateCode[NgEnumDeclaration],
				"{{ConstName}}", value.Name,
				"{{ConstValue}}", value.Value)

			insertions[NgEnumValuesInsertionPointDeclarationForPullDownSelect] += Replace2(
				NgEnumHtmlSubTemplateCode[NgEnumDeclarationForPullDownSelect],
				"{{ConstName}}", value.Name,
				"{{ConstValue}}", viewValueOfEnum)

		}

		// substitutes {{<<insertion points>>}} stuff with generated code
		for insertion := NgEnumInsertionPoint(0); insertion < NgEnumNbInsertionPoints; insertion++ {
			toReplace := "{{" + string(rune(insertion)) + "}}"
			codeTS = strings.ReplaceAll(codeTS, toReplace, insertions[insertion])
		}

		var typeOfEnumAsString string
		if gongEnum.Type == String {
			typeOfEnumAsString = "string"
		} else {
			typeOfEnumAsString = "number"
		}

		codeTS = Replace6(codeTS,
			"{{EnumName}}", gongEnum.Name,
			"{{PkgName}}", pkgName,
			"{{TitlePkgName}}", strings.Title(pkgName),
			"{{pkgname}}", strings.ToLower(pkgName),
			"{{PkgPathRoot}}", strings.ReplaceAll(pkgGoPath, "/models", ""),
			"{{type}}", typeOfEnumAsString)

		{
			file, err := os.Create(filepath.Join(matTargetPath, gongEnum.Name+".ts"))
			if err != nil {
				log.Panic(err)
			}
			defer file.Close()
			fmt.Fprint(file, codeTS)
		}
	}

}
