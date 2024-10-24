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

export * from './lib/front-repo.service'
export * from './lib/null-int64'
export * from './lib/commitnbfromback.service'
export * from './lib/push_from_front_nb.service'
export * from './lib/back-repo-data'
export * from './lib/web-socket-service'

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
export * from './lib/{{structname}}-api'
export * from './lib/{{structname}}'
export * from './lib/{{structname}}.service'
`,

	NgPublicApiDEnumsExportDeclaration: `
export * from './lib/{{Enumname}}'`,
}

// MultiCodeGeneratorNgPublicApi parses mdlPkg and generates the code for the
// PublicApi components
func CodeGeneratorNgPublicApi(modelPkg *models.ModelPkg) {

	pkgName := modelPkg.Name
	matTargetPath := modelPkg.NgDataLibrarySourceCodeDirectory
	pkgGoPath := modelPkg.PkgPath

	// generate the typescript file
	codeTS := NgPublicApiTemplateTS

	insertions := make(map[NgPublicApiInsertionPoint]string)
	for insertion := NgPublicApiInsertionPoint(0); insertion < NgPublicApiNbInsertionPoints; insertion++ {
		insertions[insertion] = ""
	}

	// have alphabetical order generation
	structList := []*models.GongStruct{}
	for _, _struct := range modelPkg.GongStructs {
		if !_struct.HasNameField() || _struct.IsIgnoredForFront {
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
	for _, _enum := range modelPkg.GongEnums {
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
