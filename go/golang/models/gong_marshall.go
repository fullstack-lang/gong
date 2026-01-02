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
	ModelGongMarshallMarshallFieldMethods
	ModelGongMarshallMarshallAllFieldsMethods
	ModelGongMarshallStructInsertionsNb
)

var ModelGongMarshallStructSubTemplateCode map[ModelGongMarshallStructInsertionId]string = // new line
map[ModelGongMarshallStructInsertionId]string{

	ModelGongMarshallStructInsertionUnmarshallDeclarations: `
	{{structname}}Ordered := []*{{Structname}}{}
	for {{structname}} := range stage.{{Structname}}s {
		{{structname}}Ordered = append({{structname}}Ordered, {{structname}})
	}
	sort.Slice({{structname}}Ordered[:], func(i, j int) bool {
		{{structname}}i := {{structname}}Ordered[i]
		{{structname}}j := {{structname}}Ordered[j]
		{{structname}}i_order, oki := stage.{{Structname}}Map_Staged_Order[{{structname}}i]
		{{structname}}j_order, okj := stage.{{Structname}}Map_Staged_Order[{{structname}}j]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return {{structname}}i_order < {{structname}}j_order
	})
	if len({{structname}}Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for _, {{structname}} := range {{structname}}Ordered {

		identifiersDecl += {{structname}}.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment{{ValuesInitialization}}
	}
`,

	ModelGongMarshallStructInsertionUnmarshallPointersInitializations: `
	for _, {{structname}} := range {{structname}}Ordered {
		_ = {{structname}}
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization{{PointersInitialization}}
	}
`,
	ModelGongMarshallMarshallFieldMethods: `
func ({{structname}} *{{Structname}}) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
{{ValuesInitialization2}}
{{PointersInitialization2}}	default:
		log.Panicf("Unknown field %s for Gongstruct {{Structname}}", fieldName)
	}
	return
}
`,
	ModelGongMarshallMarshallAllFieldsMethods: `
func ({{structname}} *{{Structname}}) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment{{ValuesInitialization}}
	}
	return
}`,
}

type GongMarshallFilePerStructSubTemplateId int

const (
	GongMarshallFileFieldSubTmplSetBasicFieldBool GongMarshallFilePerStructSubTemplateId = iota
	GongMarshallFileFieldSubTmplSetBasicFieldInt
	GongMarshallFileFieldSubTmplSetBasicFieldEnumString
	GongMarshallFileFieldSubTmplSetBasicFieldEnumInt
	GongMarshallFileFieldSubTmplSetBasicFieldFloat64
	GongMarshallFileFieldSubTmplSetBasicFieldString
	GongMarshallFileFieldSubTmplSetBasicFieldMeta
	GongMarshallFileFieldSubTmplSetTimeField
	GongMarshallFileFieldSubTmplSetPointerField
	GongMarshallFileFieldSubTmplSetSliceOfPointersField
	GongMarshallNonPointerFieldInitializerStatement
	GongMarshallPointerFieldInitializerStatement
)

var GongMarshallFileFieldFieldSubTemplateCode map[GongMarshallFilePerStructSubTemplateId]string = // declaration of the sub templates
map[GongMarshallFilePerStructSubTemplateId]string{

	GongMarshallFileFieldSubTmplSetBasicFieldBool: `
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", {{structname}}.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "{{FieldName}}")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", {{structname}}.{{FieldName}}))
`,
	GongMarshallFileFieldSubTmplSetTimeField: `
		res = TimeInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", {{structname}}.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "{{FieldName}}")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", {{structname}}.{{FieldName}}.String())
`,
	GongMarshallFileFieldSubTmplSetBasicFieldInt: `
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", {{structname}}.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "{{FieldName}}")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", {{structname}}.{{FieldName}}))
`,
	GongMarshallFileFieldSubTmplSetBasicFieldEnumString: `
		if {{structname}}.{{FieldName}} != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", {{structname}}.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "{{FieldName}}")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+{{structname}}.{{FieldName}}.ToCodeString())
		}
`,
	GongMarshallFileFieldSubTmplSetBasicFieldEnumInt: `
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", {{structname}}.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "{{FieldName}}")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+{{structname}}.{{FieldName}}.ToCodeString())
`,
	GongMarshallFileFieldSubTmplSetBasicFieldFloat64: `
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", {{structname}}.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "{{FieldName}}")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", {{structname}}.{{FieldName}}))
`,
	GongMarshallFileFieldSubTmplSetBasicFieldString: `
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", {{structname}}.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "{{FieldName}}")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string({{structname}}.{{FieldName}}))
`,
	GongMarshallFileFieldSubTmplSetBasicFieldMeta: `
		if str, ok := {{structname}}.{{FieldName}}.(string); ok {
			res = MetaFieldStructInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", {{structname}}.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "{{FieldName}}")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", str)
		}
`,
	GongMarshallFileFieldSubTmplSetPointerField: `
		if {{structname}}.{{FieldName}} != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", {{structname}}.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "{{FieldName}}")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", {{structname}}.{{FieldName}}.GongGetIdentifier(stage))
		}
`,
	GongMarshallFileFieldSubTmplSetSliceOfPointersField: `
		for _, _{{assocstructname}} := range {{structname}}.{{FieldName}} {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", {{structname}}.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "{{FieldName}}")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _{{assocstructname}}.GongGetIdentifier(stage))
			res += tmp
		}
`,
	GongMarshallNonPointerFieldInitializerStatement: `
		initializerStatements += {{structname}}.GongMarshallField(stage, "{{FieldName}}")`,
	GongMarshallPointerFieldInitializerStatement: `
		pointersInitializesStatements += {{structname}}.GongMarshallField(stage, "{{FieldName}}")`,
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
			valInitCode2 := ""
			pointerInitCode := ""
			pointerInitCode2 := ""

			for _, field := range gongStruct.Fields {

				switch field := field.(type) {
				case *models.GongBasicField:
					valInitCode2 += `	case "` + field.GetName() + `":`
					tmp := ""

					switch field.GetBasicKind() {
					case types.String:
						if field.GongEnum == nil {
							tmp = models.Replace1(
								GongMarshallFileFieldFieldSubTemplateCode[GongMarshallFileFieldSubTmplSetBasicFieldString],
								"{{FieldName}}", field.Name)
						} else {
							tmp = models.Replace1(
								GongMarshallFileFieldFieldSubTemplateCode[GongMarshallFileFieldSubTmplSetBasicFieldEnumString],
								"{{FieldName}}", field.Name)
						}
					case types.Bool:
						tmp = models.Replace1(
							GongMarshallFileFieldFieldSubTemplateCode[GongMarshallFileFieldSubTmplSetBasicFieldBool],
							"{{FieldName}}", field.Name)
					case types.Float64:
						tmp = models.Replace1(
							GongMarshallFileFieldFieldSubTemplateCode[GongMarshallFileFieldSubTmplSetBasicFieldFloat64],
							"{{FieldName}}", field.Name)
					case types.Int, types.Int64:
						if field.GongEnum == nil {
							tmp = models.Replace1(
								GongMarshallFileFieldFieldSubTemplateCode[GongMarshallFileFieldSubTmplSetBasicFieldInt],
								"{{FieldName}}", field.Name)
						} else {
							tmp = models.Replace1(
								GongMarshallFileFieldFieldSubTemplateCode[GongMarshallFileFieldSubTmplSetBasicFieldEnumInt],
								"{{FieldName}}", field.Name)
						}
					case types.UntypedNil:
						tmp = models.Replace1(
							GongMarshallFileFieldFieldSubTemplateCode[GongMarshallFileFieldSubTmplSetBasicFieldMeta],
							"{{FieldName}}", field.Name)
					default:
					}
					valInitCode += models.Replace1(
						GongMarshallFileFieldFieldSubTemplateCode[GongMarshallNonPointerFieldInitializerStatement],
						"{{FieldName}}", field.Name)
					valInitCode2 += tmp
				case *models.GongTimeField:
					valInitCode2 += `	case "` + field.GetName() + `":`
					tmp := models.Replace1(
						GongMarshallFileFieldFieldSubTemplateCode[GongMarshallFileFieldSubTmplSetTimeField],
						"{{FieldName}}", field.Name)
					valInitCode += models.Replace1(
						GongMarshallFileFieldFieldSubTemplateCode[GongMarshallNonPointerFieldInitializerStatement],
						"{{FieldName}}", field.Name)
					valInitCode2 += tmp
				case *models.PointerToGongStructField:
					pointerInitCode2 += `	case "` + field.GetName() + `":`
					tmp := models.Replace2(
						GongMarshallFileFieldFieldSubTemplateCode[GongMarshallFileFieldSubTmplSetPointerField],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name)
					valInitCode += models.Replace1(
						GongMarshallFileFieldFieldSubTemplateCode[GongMarshallPointerFieldInitializerStatement],
						"{{FieldName}}", field.Name)
					pointerInitCode2 += tmp
				case *models.SliceOfPointerToGongStructField:
					pointerInitCode2 += `	case "` + field.GetName() + `":`
					tmp := models.Replace3(
						GongMarshallFileFieldFieldSubTemplateCode[GongMarshallFileFieldSubTmplSetSliceOfPointersField],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name,
						"{{assocstructname}}", strings.ToLower(field.GongStruct.Name))
					valInitCode += models.Replace1(
						GongMarshallFileFieldFieldSubTemplateCode[GongMarshallPointerFieldInitializerStatement],
						"{{FieldName}}", field.Name)
					pointerInitCode2 += tmp
				default:
				}

			}

			valInitCode = models.Replace2(valInitCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			valInitCode2 = models.Replace2(valInitCode2,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			pointerInitCode = models.Replace2(pointerInitCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			pointerInitCode2 = models.Replace2(pointerInitCode2,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			generatedCodeFromSubTemplate := models.Replace6(ModelGongMarshallStructSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{ValuesInitialization}}", valInitCode,
				"{{ValuesInitialization2}}", valInitCode2,
				"{{PointersInitialization}}", pointerInitCode,
				"{{PointersInitialization2}}", pointerInitCode2,
			)

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

	file, err := os.Create(filepath.Join(pkgPath, string(models.GeneratedGongMarshallGoFilePath)))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)

}
