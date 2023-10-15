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

const ModelGongWopFileTemplate = `// generated code - do not edit
package models

// insertion point{{` + string(rune(ModelGongWopStruct)) + `}}
`

type ModelGongWopStructInsertionId int

const (
	ModelGongWopStruct ModelGongWopStructInsertionId = iota
	ModelGongWopStructInsertionsNb
)

var ModelGongWopStructSubTemplateCode map[ModelGongWopStructInsertionId]string = // new line
map[ModelGongWopStructInsertionId]string{
	ModelGongWopStruct: `
type {{Structname}}_WOP struct {
	// insertion point{{FieldCode}}
}
`,
}

type GongWopSubSubTemplateInsertionsId int

const (
	GongWopFieldStringDecl GongWopSubSubTemplateInsertionsId = iota + 100
	GongWopFieldIntDecl
)

var GongWopSubSubTemplate map[GongWopSubSubTemplateInsertionsId]string = // new line
map[GongWopSubSubTemplateInsertionsId]string{

	GongWopFieldStringDecl: `
	{{FieldName}} string
`,
	GongWopFieldIntDecl: `
	{{FieldName}} int
`,
}

func CodeGeneratorModelGongWop(
	mdlPkg *models.ModelPkg,
	pkgName string,
	pkgPath string) {

	// generate the typescript file
	codeGO := ModelGongWopFileTemplate

	subStructCodes := make(map[ModelGongWopStructInsertionId]string)
	for subStructTemplate := range ModelGongWopStructSubTemplateCode {
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

		for subStructTemplate := range ModelGongWopStructSubTemplateCode {

			fieldCode := ""

			for _, field := range gongStruct.Fields {

				switch field := field.(type) {
				case *models.GongBasicField:
					fieldCode += models.Replace1(
						GongWopSubSubTemplate[GongWopFieldStringDecl],
						"{{FieldName}}", field.Name)
				default:
					_ = field
				}

			}

			fieldCode = models.Replace2(fieldCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			generatedCodeFromSubTemplate := models.Replace3(ModelGongWopStructSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{FieldCode}}", fieldCode)

			subStructCodes[subStructTemplate] += generatedCodeFromSubTemplate
		}

	}

	// substitutes {{<<insertionPerStructId points>>}} stuff with generated code
	for insertionPerStructId := ModelGongWopStructInsertionId(0); insertionPerStructId < ModelGongWopStructInsertionsNb; insertionPerStructId++ {
		toReplace := "{{" + string(rune(insertionPerStructId)) + "}}"
		codeGO = strings.ReplaceAll(codeGO, toReplace, subStructCodes[insertionPerStructId])
	}

	caserEnglish := cases.Title(language.English)
	codeGO = models.Replace3(codeGO,
		"{{PkgName}}", pkgName,
		"{{TitlePkgName}}", caserEnglish.String(pkgName),
		"{{pkgname}}", strings.ToLower(pkgName),
	)

	file, err := os.Create(filepath.Join(pkgPath, "gong_wop.go"))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)
}
