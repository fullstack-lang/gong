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
	ModelGongMarshallMarshallFieldMethods
	ModelGongMarshallMarshallAllFieldsMethods
	ModelGongMarshallStructInsertionsNb
)

var ModelGongMarshallStructSubTemplateCode map[ModelGongMarshallStructInsertionId]string = // new line
map[ModelGongMarshallStructInsertionId]string{

	ModelGongMarshallStructInsertionUnmarshallDeclarations: `
	gongMarshallInstances(stage, stage.{{Structname}}s, &identifiersDecl, &initializerStatements, &pointersInitializesStatements)`,

	ModelGongMarshallMarshallFieldMethods: `
func ({{structname}} *{{Structname}}) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := {{structname}}.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
{{ValuesInitialization2}}
{{PointersInitialization2}}	default:
		log.Panicf("Unknown field %s for Gongstruct {{Structname}}", fieldName)
	}
	return
}
`,
	ModelGongMarshallMarshallAllFieldsMethods: `
func ({{structname}} *{{Structname}}) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment{{ValuesInitialization}}
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
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
		res = __gong__marshallBool(ident, "{{FieldName}}", {{structname}}.{{FieldName}})
`,
	GongMarshallFileFieldSubTmplSetTimeField: `
		res = __gong__marshallTime(ident, "{{FieldName}}", {{structname}}.{{FieldName}}.String())
`,
	GongMarshallFileFieldSubTmplSetBasicFieldInt: `
		res = __gong__marshallInt(ident, "{{FieldName}}", {{structname}}.{{FieldName}})
`,
	GongMarshallFileFieldSubTmplSetBasicFieldEnumString: `
		res = __gong__marshallEnumString(ident, "{{FieldName}}", {{structname}}.{{FieldName}}.ToCodeString())
`,
	GongMarshallFileFieldSubTmplSetBasicFieldEnumInt: `
		res = __gong__marshallEnumInt(ident, "{{FieldName}}", {{structname}}.{{FieldName}}.ToCodeString())
`,
	GongMarshallFileFieldSubTmplSetBasicFieldFloat64: `
		res = __gong__marshallFloat(ident, "{{FieldName}}", {{structname}}.{{FieldName}})
`,
	GongMarshallFileFieldSubTmplSetBasicFieldString: `
		res = __gong__marshallString(ident, "{{FieldName}}", {{structname}}.{{FieldName}})
`,
	GongMarshallFileFieldSubTmplSetBasicFieldMeta: `
		if str, ok := {{structname}}.{{FieldName}}.(string); ok {
			res = __gong__marshallMeta(ident, "{{FieldName}}", str)
		}
`,
	GongMarshallFileFieldSubTmplSetPointerField: `
		if {{structname}}.{{FieldName}} != nil {
			res = __gong__marshallPointer(ident, "{{FieldName}}", {{structname}}.{{FieldName}}.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "{{FieldName}}", "nil")
		}
`,
	GongMarshallFileFieldSubTmplSetSliceOfPointersField: `
		var sb strings.Builder
		for _, _{{assocstructname}} := range {{structname}}.{{FieldName}} {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "{{FieldName}}", _{{assocstructname}}.GongGetIdentifier(stage)))
		}
		res = sb.String()
`,
	GongMarshallNonPointerFieldInitializerStatement: `
		initializerStatements.WriteString({{structname}}.GongMarshallField(stage, "{{FieldName}}"))`,
	GongMarshallPointerFieldInitializerStatement: `
		pointersInitializesStatements.WriteString({{structname}}.GongMarshallField(stage, "{{FieldName}}"))`,
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

		valInitCode := ""
		valInitCode2 := ""
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
				if field.GongStruct.IsOmittedForMarshalling {
					continue
				}
				pointerInitCode2 += `	case "` + field.GetName() + `":`
				tmp := models.Replace1(
					GongMarshallFileFieldFieldSubTemplateCode[GongMarshallFileFieldSubTmplSetPointerField],
					"{{FieldName}}", field.Name)
				valInitCode += models.Replace1(
					GongMarshallFileFieldFieldSubTemplateCode[GongMarshallPointerFieldInitializerStatement],
					"{{FieldName}}", field.Name)
				pointerInitCode2 += tmp
			case *models.SliceOfPointerToGongStructField:
				if field.GongStruct.IsOmittedForMarshalling {
					continue
				}
				pointerInitCode2 += `	case "` + field.GetName() + `":`
				tmp := models.Replace2(
					GongMarshallFileFieldFieldSubTemplateCode[GongMarshallFileFieldSubTmplSetSliceOfPointersField],
					"{{FieldName}}", field.Name,
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

		pointerInitCode2 = models.Replace2(pointerInitCode2,
			"{{structname}}", strings.ToLower(gongStruct.Name),
			"{{Structname}}", gongStruct.Name)

		for subStructTemplate := range ModelGongMarshallStructSubTemplateCode {

			if gongStruct.IsOmittedForMarshalling &&
				subStructTemplate == ModelGongMarshallStructInsertionUnmarshallDeclarations {
				continue
			}

			generatedCodeFromSubTemplate := models.Replace5(ModelGongMarshallStructSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{ValuesInitialization}}", valInitCode,
				"{{ValuesInitialization2}}", valInitCode2,
				"{{PointersInitialization2}}", pointerInitCode2,
			)

			subStructCodes[subStructTemplate] += generatedCodeFromSubTemplate
		}

	}

	// substitutes {{<<insertionPerStructId points>>}} stuff with generated code
	for insertionPerStructId := range ModelGongMarshallStructInsertionsNb {
		toReplace := "{{" + string(rune(insertionPerStructId)) + "}}"
		codeGO = strings.ReplaceAll(codeGO, toReplace, subStructCodes[insertionPerStructId])
	}

	caserEnglish := cases.Title(language.English)
	codeGO = models.Replace(codeGO,
		"{{PkgName}}", pkgName,
		"{{TitlePkgName}}", caserEnglish.String(pkgName),
		"{{pkgname}}", strings.ToLower(pkgName),
		"{{PkgGoName}}", modelPkg.PkgGoName,
		"	 | ", "	", // for the replacement of the of the first bar in the Gongstruct Type def
	)

	file, err := os.Create(filepath.Join(pkgPath, string(models.GeneratedGongMarshallGoFilePath)))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)

}
