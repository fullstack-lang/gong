package models

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
	for idx, {{structname}} := range {{structname}}Ordered {

		id = generatesIdentifier("{{Structname}}", idx, {{structname}}.Name)
		map_{{Structname}}_Identifiers[{{structname}}] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "{{Structname}}")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", {{structname}}.Name)
		identifiersDecl += decl

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

type GongMarshallFilePerStructSubTemplateId int

const (
	GongMarshallFileFieldSubTmplSetBasicFieldBool GongMarshallFilePerStructSubTemplateId = iota
	GongMarshallFileFieldSubTmplSetBasicFieldInt
	GongMarshallFileFieldSubTmplSetBasicFieldEnumString
	GongMarshallFileFieldSubTmplSetBasicFieldEnumInt
	GongMarshallFileFieldSubTmplSetBasicFieldFloat64
	GongMarshallFileFieldSubTmplSetBasicFieldString
	GongMarshallFileFieldSubTmplSetBasicFieldStringDocLink
	GongMarshallFileFieldSubTmplSetTimeField
	GongMarshallFileFieldSubTmplSetPointerField
	GongMarshallFileFieldSubTmplSetSliceOfPointersField
)

var GongMarshallFileFieldFieldSubTemplateCode map[GongMarshallFilePerStructSubTemplateId]string = // declaration of the sub templates
map[GongMarshallFilePerStructSubTemplateId]string{

	GongMarshallFileFieldSubTmplSetBasicFieldBool: `
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "{{FieldName}}")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", {{structname}}.{{FieldName}}))
		initializerStatements += setValueField
`,
	GongMarshallFileFieldSubTmplSetTimeField: `
		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "{{FieldName}}")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", {{structname}}.{{FieldName}}.String())
		initializerStatements += setValueField
`,
	GongMarshallFileFieldSubTmplSetBasicFieldInt: `
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "{{FieldName}}")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", {{structname}}.{{FieldName}}))
		initializerStatements += setValueField
`,
	GongMarshallFileFieldSubTmplSetBasicFieldEnumString: `
		if {{structname}}.{{FieldName}} != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "{{FieldName}}")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+{{structname}}.{{FieldName}}.ToCodeString())
			initializerStatements += setValueField
		}
`,
	GongMarshallFileFieldSubTmplSetBasicFieldEnumInt: `
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "{{FieldName}}")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+{{structname}}.{{FieldName}}.ToCodeString())
		initializerStatements += setValueField
`,
	GongMarshallFileFieldSubTmplSetBasicFieldFloat64: `
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "{{FieldName}}")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", {{structname}}.{{FieldName}}))
		initializerStatements += setValueField
`,
	GongMarshallFileFieldSubTmplSetBasicFieldString: `
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "{{FieldName}}")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string({{structname}}.{{FieldName}}))
		initializerStatements += setValueField
`,
	GongMarshallFileFieldSubTmplSetBasicFieldStringDocLink: `
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "\n\t{{Identifier}}",
			fmt.Sprintf("\n\n\t//gong:ident [%s] comment added to overcome the problem with the comment map association\n\t{{Identifier}}",
				string({{structname}}.{{FieldName}})))
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "{{FieldName}}")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string({{structname}}.{{FieldName}}))
		initializerStatements += setValueField
`,
	GongMarshallFileFieldSubTmplSetPointerField: `
		if {{structname}}.{{FieldName}} != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "{{FieldName}}")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_{{AssocStructName}}_Identifiers[{{structname}}.{{FieldName}}])
			pointersInitializesStatements += setPointerField
		}
`,
	GongMarshallFileFieldSubTmplSetSliceOfPointersField: `
		for _, _{{assocstructname}} := range {{structname}}.{{FieldName}} {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "{{FieldName}}")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_{{AssocStructName}}_Identifiers[_{{assocstructname}}])
			pointersInitializesStatements += setPointerField
		}
`,
}

func CodeGeneratorModelGongMarshall(
	modelPkg *models.ModelPkg,
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

		for subStructTemplate := range ModelGongMarshallStructSubTemplateCode {

			valInitCode := ""
			pointerInitCode := ""

			for _, field := range gongStruct.Fields {

				switch field := field.(type) {
				case *models.GongBasicField:

					switch field.GetBasicKind() {
					case types.String:
						if field.GongEnum == nil {

							if !field.IsDocLink {
								valInitCode += models.Replace1(
									GongMarshallFileFieldFieldSubTemplateCode[GongMarshallFileFieldSubTmplSetBasicFieldString],
									"{{FieldName}}", field.Name)
							} else {
								valInitCode += models.Replace1(
									GongMarshallFileFieldFieldSubTemplateCode[GongMarshallFileFieldSubTmplSetBasicFieldStringDocLink],
									"{{FieldName}}", field.Name)
							}

						} else {
							valInitCode += models.Replace1(
								GongMarshallFileFieldFieldSubTemplateCode[GongMarshallFileFieldSubTmplSetBasicFieldEnumString],
								"{{FieldName}}", field.Name)
						}
					case types.Bool:
						valInitCode += models.Replace1(
							GongMarshallFileFieldFieldSubTemplateCode[GongMarshallFileFieldSubTmplSetBasicFieldBool],
							"{{FieldName}}", field.Name)
					case types.Float64:
						valInitCode += models.Replace1(
							GongMarshallFileFieldFieldSubTemplateCode[GongMarshallFileFieldSubTmplSetBasicFieldFloat64],
							"{{FieldName}}", field.Name)
					case types.Int, types.Int64:
						if field.GongEnum == nil {
							valInitCode += models.Replace1(
								GongMarshallFileFieldFieldSubTemplateCode[GongMarshallFileFieldSubTmplSetBasicFieldInt],
								"{{FieldName}}", field.Name)
						} else {
							valInitCode += models.Replace1(
								GongMarshallFileFieldFieldSubTemplateCode[GongMarshallFileFieldSubTmplSetBasicFieldEnumInt],
								"{{FieldName}}", field.Name)
						}
					default:
					}
				case *models.GongTimeField:
					valInitCode += models.Replace1(
						GongMarshallFileFieldFieldSubTemplateCode[GongMarshallFileFieldSubTmplSetTimeField],
						"{{FieldName}}", field.Name)
				case *models.PointerToGongStructField:
					pointerInitCode += models.Replace2(
						GongMarshallFileFieldFieldSubTemplateCode[GongMarshallFileFieldSubTmplSetPointerField],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name)
				case *models.SliceOfPointerToGongStructField:
					pointerInitCode += models.Replace3(
						GongMarshallFileFieldFieldSubTemplateCode[GongMarshallFileFieldSubTmplSetSliceOfPointersField],
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
