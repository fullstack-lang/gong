package angular

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/fullstack-lang/gong/go/models"
)

const NgPublicApiTemplateTS = `// generated from ng_file_public_api_ts.go
/*
* Public API Surface of {{PkgName}}
*/

export * from './lib/{{pkgname}}.module'

export * from './lib/splitter/splitter.component'
export * from './lib/sidebar/sidebar.component'

export * from './lib/front-repo.service'
export * from './lib/null-int64'
export * from './lib/commitnbfromback.service'
export * from './lib/gongstruct-selection.service'
export * from './lib/push_from_front_nb.service'

{{` + string(rune(NgPublicApiInsertionStructComponentsExportDeclaration)) + `}}
{{` + string(rune(NgPublicApiInsertionEnumsExportDeclaration)) + `}}
`

// insertion points
type NgPublicApiInsertionPoint int

const (
	NgPublicApiInsertionStructComponentsExportDeclaration NgPublicApiInsertionPoint = iota
	NgPublicApiInsertionEnumsExportDeclaration
	NgPublicApiNbInsertionPoints
)

// Sub Templates
type NgPublicApiSubTemplate int

const (
	NgPublicApiStructComponentsExportDeclaration NgPublicApiSubTemplate = iota
	NgPublicApiDEnumsExportDeclaration
)

var NgPublicApiHtmlSubTemplateCode map[NgPublicApiSubTemplate]string = map[NgPublicApiSubTemplate]string{
	NgPublicApiStructComponentsExportDeclaration: `
export * from './lib/{{structname}}-detail/{{structname}}-detail.component'
export * from './lib/{{structname}}-presentation/{{structname}}-presentation.component'
export * from './lib/{{structname}}-sorting/{{structname}}-sorting.component'
export * from './lib/{{structname}}s-table/{{structname}}s-table.component'
export * from './lib/{{structname}}-db'
export * from './lib/{{structname}}.service'
`,

	NgPublicApiDEnumsExportDeclaration: `
export * from './lib/{{Enumname}}'`,
}

// MultiCodeGeneratorNgPublicApi parses mdlPkg and generates the code for the
// PublicApi components
func CodeGeneratorNgPublicApi(
	mdlPkg *models.ModelPkg,
	pkgName string,
	matTargetPath string,
	pkgGoPath string) {

	// generate the typescript file
	codeTS := NgPublicApiTemplateTS

	insertions := make(map[NgPublicApiInsertionPoint]string)
	for insertion := NgPublicApiInsertionPoint(0); insertion < NgPublicApiNbInsertionPoints; insertion++ {
		insertions[insertion] = ""
	}

	// have alphabetical order generation
	structList := []*models.GongStruct{}
	for _, _struct := range mdlPkg.GongStructs {
		if !_struct.HasNameField() {
			continue
		}
		structList = append(structList, _struct)
	}
	sort.Slice(structList[:], func(i, j int) bool {
		return structList[i].Name < structList[j].Name
	})

	for _, _struct := range structList {

		insertions[NgPublicApiInsertionStructComponentsExportDeclaration] += models.Replace2(
			NgPublicApiHtmlSubTemplateCode[NgPublicApiStructComponentsExportDeclaration],
			"{{Structname}}", _struct.Name,
			"{{structname}}", strings.ToLower(_struct.Name))
	}

	// have alphabetical order generation
	enumList := []*models.GongEnum{}
	for _, _enum := range mdlPkg.GongEnums {
		enumList = append(enumList, _enum)
	}
	sort.Slice(enumList[:], func(i, j int) bool {
		return enumList[i].Name < enumList[j].Name
	})

	for _, enum := range enumList {

		insertions[NgPublicApiInsertionEnumsExportDeclaration] += models.Replace1(
			NgPublicApiHtmlSubTemplateCode[NgPublicApiDEnumsExportDeclaration],
			"{{Enumname}}", enum.Name)
	}

	// substitutes {{<<insertion points>>}} stuff with generated code
	for insertion := NgPublicApiInsertionPoint(0); insertion < NgPublicApiNbInsertionPoints; insertion++ {
		toReplace := "{{" + string(rune(insertion)) + "}}"
		codeTS = strings.ReplaceAll(codeTS, toReplace, insertions[insertion])
	}

	codeTS = models.Replace4(codeTS,
		"{{PkgName}}", pkgName,
		"{{TitlePkgName}}", strings.Title(pkgName),
		"{{pkgname}}", strings.ToLower(pkgName),
		"{{PkgPathRoot}}", strings.ReplaceAll(pkgGoPath, "/models", ""))

	{
		file, err := os.Create(filepath.Join(matTargetPath, "..", "public-api.ts"))
		if err != nil {
			log.Panic(err)
		}
		defer file.Close()
		fmt.Fprint(file, codeTS)
	}

}
