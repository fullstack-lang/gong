package golang

import (
	"fmt"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/fullstack-lang/gong/go/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// insertion points are places where the code is
// generated per gong struct
type ModelGongMarshallStructInsertionId int

const (
	ModelGongMarshallStructInsertionUnmarshallDeclarations ModelGongMarshallStructInsertionId = iota
	ModelGongMarshallStructInsertionUnmarshallPointersInitializations
	ModelGongMarshallStructInsertionsNb
)

var ModelGongMarshallStructSubTemplateCode map[ModelGongMarshallStructInsertionId]string = // new line
map[ModelGongMarshallStructInsertionId]string{

	ModelGongMarshallStructInsertionUnmarshallDeclarations: `
	map_{{Structname}}_Identifiers := make(map[*{{Structname}}]string)
	_ = map_{{Structname}}_Identifiers

	{{structname}}Ordered := []*{{Structname}}{}
	for {{structname}} := range stage.{{Structname}}s {
		{{structname}}Ordered = append({{structname}}Ordered, {{structname}})
	}
	sort.Slice({{structname}}Ordered[:], func(i, j int) bool {
		return {{structname}}Ordered[i].Name < {{structname}}Ordered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of {{Structname}}"
	for idx, {{structname}} := range {{structname}}Ordered {

		id = generatesIdentifier("{{Structname}}", idx, {{structname}}.Name)
		map_{{Structname}}_Identifiers[{{structname}}] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "{{Structname}}")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", {{structname}}.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// {{Structname}} values setup"
		// Initialisation of values{{ValuesInitialization}}
	}
`,

	ModelGongMarshallStructInsertionUnmarshallPointersInitializations: `
	for idx, {{structname}} := range {{structname}}Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("{{Structname}}", idx, {{structname}}.Name)
		map_{{Structname}}_Identifiers[{{structname}}] = id

		// Initialisation of values{{PointersInitialization}}
	}
`,
}

func CodeGeneratorModelGongMarshall(
	mdlPkg *models.ModelPkg,
	pkgName string,
	pkgPath string) {

	// generate the typescript file
	codeGO := ModelGongMarshallFileTemplate

	subStructCodes := make(map[ModelGongMarshallStructInsertionId]string)
	for subStructTemplate := range ModelGongMarshallStructSubTemplateCode {
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

		for subStructTemplate := range ModelGongMarshallStructSubTemplateCode {

			valInitCode := ""
			pointerInitCode := ""

			for _, field := range gongStruct.Fields {

				switch field := field.(type) {
				case *models.GongBasicField:

					switch field.GetBasicKind() {
					case types.String:
						if field.GongEnum == nil {
							valInitCode += models.Replace1(
								GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetBasicFieldString],
								"{{FieldName}}", field.Name)

						} else {
							valInitCode += models.Replace1(
								GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetBasicFieldEnumString],
								"{{FieldName}}", field.Name)
						}
					case types.Bool:
						valInitCode += models.Replace1(
							GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetBasicFieldBool],
							"{{FieldName}}", field.Name)
					case types.Float64:
						valInitCode += models.Replace1(
							GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetBasicFieldFloat64],
							"{{FieldName}}", field.Name)
					case types.Int, types.Int64:
						if field.GongEnum == nil {
							valInitCode += models.Replace1(
								GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetBasicFieldInt],
								"{{FieldName}}", field.Name)
						} else {
							valInitCode += models.Replace1(
								GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetBasicFieldEnumInt],
								"{{FieldName}}", field.Name)
						}
					default:
					}
				case *models.GongTimeField:
					valInitCode += models.Replace1(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetTimeField],
						"{{FieldName}}", field.Name)
				case *models.PointerToGongStructField:
					pointerInitCode += models.Replace2(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetPointerField],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name)
				case *models.SliceOfPointerToGongStructField:
					pointerInitCode += models.Replace3(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetSliceOfPointersField],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name,
						"{{assocstructname}}", strings.ToLower(field.GongStruct.Name))
				default:
				}

			}

			valInitCode = models.Replace2(valInitCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			pointerInitCode = models.Replace2(pointerInitCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			generatedCodeFromSubTemplate := models.Replace4(ModelGongMarshallStructSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{ValuesInitialization}}", valInitCode,
				"{{PointersInitialization}}", pointerInitCode)

			subStructCodes[subStructTemplate] += generatedCodeFromSubTemplate
		}

	}

	// substitutes {{<<insertionPerStructId points>>}} stuff with generated code
	for insertionPerStructId := ModelGongMarshallStructInsertionId(0); insertionPerStructId < ModelGongMarshallStructInsertionsNb; insertionPerStructId++ {
		toReplace := "{{" + string(rune(insertionPerStructId)) + "}}"
		codeGO = strings.ReplaceAll(codeGO, toReplace, subStructCodes[insertionPerStructId])
	}

	caserEnglish := cases.Title(language.English)
	codeGO = models.Replace4(codeGO,
		"{{PkgName}}", pkgName,
		"{{TitlePkgName}}", caserEnglish.String(pkgName),
		"{{pkgname}}", strings.ToLower(pkgName),
		"	 | ", "	", // for the replacement of the of the first bar in the Gongstruct Type def
	)

	file, err := os.Create(filepath.Join(pkgPath, "gong_marshall.go"))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)

}
