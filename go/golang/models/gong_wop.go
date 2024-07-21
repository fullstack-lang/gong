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

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

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

func (from *{{Structname}}) CopyBasicFields(to *{{Structname}}) {
	// insertion point{{FieldsCopyCode}}
}
`,
}

type GongWopSubSubTemplateInsertionsId int

const (
	GongWopBasicFieldDecl GongWopSubSubTemplateInsertionsId = iota + 100
	GongWopEnumFieldDecl
	GongWopTimeFieldDecl
	GongWopDurationFieldDecl
	GongWopBasicFieldCopy
)

var GongWopSubSubTemplate map[GongWopSubSubTemplateInsertionsId]string = // new line
map[GongWopSubSubTemplateInsertionsId]string{

	GongWopBasicFieldDecl: `{{DeclarationPrefixPrologue}}
	{{FieldName}} {{BasicKindName}}{{DeclarationPrefixEpilogue}}`,

	GongWopEnumFieldDecl: `{{DeclarationPrefixPrologue}}
	{{FieldName}} {{EnumType}}{{DeclarationPrefixEpilogue}}`,

	GongWopTimeFieldDecl: `{{DeclarationPrefixPrologue}}
	{{FieldName}} time.Time{{DeclarationPrefixEpilogue}}`,

	GongWopDurationFieldDecl: `{{DeclarationPrefixPrologue}}
	{{FieldName}} time.Duration{{DeclarationPrefixEpilogue}}`,

	GongWopBasicFieldCopy: `
	to.{{FieldName}} = from.{{FieldName}}`,
}

func CodeGeneratorModelGongWop(
	modelPkg *models.ModelPkg,
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

		for subStructTemplate := range ModelGongWopStructSubTemplateCode {

			fieldCode := ""
			fieldCopyCode := ""

			for _, field := range gongStruct.Fields {

				fieldName := field.GetName()
				fieldNameForDeclation := fieldName

				// in case of a field within an anonymous struct, one needs
				// to strip the prefix
				fieldNameSplitted := strings.Split(fieldName, ".")

				var _, _,
					_, _,
					prologueWOP, epilogueWOP string
				if len(fieldNameSplitted) > 1 {
					fieldNameForDeclation = fieldNameSplitted[1]
					_, _,
						_, _,
						prologueWOP, epilogueWOP =
						gongStruct.ComputeFielProloguesEpilogues(field)

				}

				switch field := field.(type) {

				case *models.GongBasicField:
					fieldCopyCode += models.Replace1(
						GongWopSubSubTemplate[GongWopBasicFieldCopy],
						"{{FieldName}}", fieldName)
					if field.GongEnum == nil {
						if field.DeclaredType == "time.Duration" {
							fieldCode += models.Replace3(
								GongWopSubSubTemplate[GongWopDurationFieldDecl],
								"{{DeclarationPrefixPrologue}}", prologueWOP,
								"{{DeclarationPrefixEpilogue}}", epilogueWOP,
								"{{FieldName}}", fieldNameForDeclation)
						} else {
							fieldCode += models.Replace4(
								GongWopSubSubTemplate[GongWopBasicFieldDecl],
								"{{DeclarationPrefixPrologue}}", prologueWOP,
								"{{DeclarationPrefixEpilogue}}", epilogueWOP,
								"{{FieldName}}", fieldNameForDeclation,
								"{{BasicKindName}}", field.BasicKindName)
						}

					} else {
						fieldCode += models.Replace4(
							GongWopSubSubTemplate[GongWopEnumFieldDecl],
							"{{DeclarationPrefixPrologue}}", prologueWOP,
							"{{DeclarationPrefixEpilogue}}", epilogueWOP,
							"{{FieldName}}", fieldNameForDeclation,
							"{{EnumType}}", field.GongEnum.GetName())
					}

				case *models.GongTimeField:
					fieldCopyCode += models.Replace1(
						GongWopSubSubTemplate[GongWopBasicFieldCopy],
						"{{FieldName}}", fieldName)
					fieldCode += models.Replace3(
						GongWopSubSubTemplate[GongWopTimeFieldDecl],
						"{{DeclarationPrefixPrologue}}", prologueWOP,
						"{{DeclarationPrefixEpilogue}}", epilogueWOP,
						"{{FieldName}}", fieldNameForDeclation)

				default:
					_ = field
				}

			}

			fieldCode = models.Replace2(fieldCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			generatedCodeFromSubTemplate := models.Replace4(ModelGongWopStructSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{FieldCode}}", fieldCode,
				"{{FieldsCopyCode}}", fieldCopyCode,
			)

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
