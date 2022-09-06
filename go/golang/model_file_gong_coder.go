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

const ModelGongCoderFileTemplate = `package models

import "time"

// GongfieldCoder return an instance of Type where each field 
// encodes the index of the field
//
// This allows for refactorable field names
// 
func GongfieldCoder[Type Gongstruct]() Type {
	var t Type

	switch any(t).(type) {
	// insertion point for cases{{` + string(rune(ModelGongCoderGenericGongstructCoder)) + `}}
	default:
		return t
	}
}

type Gongfield interface {
	string | bool | int | float64 | time.Time | time.Duration{{` + string(rune(ModelGongCoderGenericGongstructTypes)) + `}}
}

// GongfieldName provides the name of the field by passing the instance of the coder to
// the fonction.
//
// This allows for refactorable field name
//
// fieldCoder := models.GongfieldCoder[models.Astruct]()
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Name))
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Booleanfield))
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Intfield))
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Floatfield))
// 
// limitations:
// 1. cannot encode boolean fields
// 2. for associations (pointer to gongstruct or slice of pointer to gongstruct, uses GetAssociationName)
func GongfieldName[Type PointerToGongstruct, FieldType Gongfield](field FieldType) string {
	var t Type

	switch any(t).(type) {
	// insertion point for cases{{` + string(rune(ModelGongCoderGenericGongstructNamerString)) + `}}
	default:
		return ""
	}
	_ = field
	return ""
}
`

//
// insertion points are places where the code is
// generated per gong struct
//
type ModelGongCoderStructInsertionId int

const (
	ModelGongCoderGenericGongstructTypes ModelGongCoderStructInsertionId = iota
	ModelGongCoderGenericGongstructCoder
	ModelGongCoderGenericGongstructNamerString
	ModelGongCoderStructInsertionsNb
)

var ModelGongCoderStructSubTemplateCode map[ModelGongCoderStructInsertionId]string = // new line
map[ModelGongCoderStructInsertionId]string{
	ModelGongCoderGenericGongstructTypes: ` | *{{Structname}} | []*{{Structname}}`,
	ModelGongCoderGenericGongstructCoder: `
	case {{Structname}}:
		fieldCoder := {{Structname}}{}
		// insertion point for field dependant code{{FieldCode}}
		return (any)(fieldCoder).(Type)`,
	ModelGongCoderGenericGongstructNamerString: `
	case *{{Structname}}:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name{{FieldNameString}}
		case int, int64:
			// insertion point for field dependant name{{FieldNameInt}}
		case float64:
			// insertion point for field dependant name{{FieldNameFloat64}}
		case time.Time:
			// insertion point for field dependant name{{FieldNameDate}}
		case bool:
			// insertion point for field dependant name{{FieldNameBoolean}}
		}`,
}

type ModelGongCoderFieldInsertionId int

const (
	ModelGongCoderFieldCodeString ModelGongCoderFieldInsertionId = iota
	ModelGongCoderFieldCodeInt
	ModelGongCoderFieldCodeFloat64
	ModelGongCoderFieldCodeDate
	ModelGongCoderFieldCodeBoolean
	ModelGongCoderFieldNameString
	ModelGongCoderFieldNameInt
	ModelGongCoderFieldNameFloat64
	ModelGongCoderFieldNameDate
	ModelGongCoderFieldNameBoolean
	ModelGongCoderFieldInsertionsNb
)

var ModelGongCoderFieldSubTemplateCode map[ModelGongCoderFieldInsertionId]string = // new line
map[ModelGongCoderFieldInsertionId]string{
	ModelGongCoderFieldCodeString: `
		fieldCoder.{{FieldName}} = "{{Value}}"`,
	ModelGongCoderFieldCodeInt: `
		fieldCoder.{{FieldName}} = {{Value}}`,
	ModelGongCoderFieldCodeFloat64: `
		fieldCoder.{{FieldName}} = {{Value}}`,
	ModelGongCoderFieldCodeDate: `
		fieldCoder.{{FieldName}} = time.Date({{Value}}, time.January, 0, 0, 0, 0, 0, time.UTC)`,
	ModelGongCoderFieldCodeBoolean: `
		fieldCoder.{{FieldName}} = {{Value}}`,
	ModelGongCoderFieldNameString: `
			if field == "{{Value}}" {
				return "{{FieldName}}"
			}`,
	ModelGongCoderFieldNameInt: `
			if field == {{Value}} {
				return "{{FieldName}}"
			}`,
	ModelGongCoderFieldNameFloat64: `
			if field == {{Value}} {
				return "{{FieldName}}"
			}`,
	ModelGongCoderFieldNameDate: `
			if field == time.Date({{Value}}, time.January, 0, 0, 0, 0, 0, time.UTC) {
				return "{{FieldName}}"
			}`,
	ModelGongCoderFieldNameBoolean: `
			if field == {{Value}} {
				return "{{FieldName}}"
			}`,
}

func CodeGeneratorModelGongCoder(
	mdlPkg *models.ModelPkg,
	pkgName string,
	pkgPath string) {

	// generate the typescript file
	codeGO := ModelGongCoderFileTemplate

	subStructCodes := make(map[ModelGongCoderStructInsertionId]string)
	for subStructTemplate := range ModelGongCoderStructSubTemplateCode {
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

		for subStructTemplate := range ModelGongCoderStructSubTemplateCode {
			fieldCode := ""
			fieldNameString := ""
			fieldNameInt := ""
			fieldNameBool := ""
			fieldNameFloat64 := ""
			fieldNameDate := ""

			// it is impossible to encode more than 2 values into
			// a boolean
			//
			// Therefore, only the first two boolean field can be encoded
			// boolFiedIndex is the index of the those boolean fields
			boolFiedIndex := 0
			for idx, field := range gongStruct.Fields {

				switch field := field.(type) {
				case *models.GongBasicField:
					switch field.GetBasicKind() {
					case types.String:
						fieldCode += models.Replace2(
							ModelGongCoderFieldSubTemplateCode[ModelGongCoderFieldCodeString],
							"{{FieldName}}", field.Name,
							"{{Value}}", fmt.Sprintf("%d", idx))
						fieldNameString += models.Replace2(
							ModelGongCoderFieldSubTemplateCode[ModelGongCoderFieldNameString],
							"{{FieldName}}", field.Name,
							"{{Value}}", fmt.Sprintf("%d", idx))
					case types.Int, types.Int64:
						fieldCode += models.Replace2(
							ModelGongCoderFieldSubTemplateCode[ModelGongCoderFieldCodeInt],
							"{{FieldName}}", field.Name,
							"{{Value}}", fmt.Sprintf("%d", idx))
						fieldNameInt += models.Replace2(
							ModelGongCoderFieldSubTemplateCode[ModelGongCoderFieldNameInt],
							"{{FieldName}}", field.Name,
							"{{Value}}", fmt.Sprintf("%d", idx))
					case types.Float64:
						fieldCode += models.Replace2(
							ModelGongCoderFieldSubTemplateCode[ModelGongCoderFieldCodeFloat64],
							"{{FieldName}}", field.Name,
							"{{Value}}", fmt.Sprintf("%f", float64(idx)))
						fieldNameFloat64 += models.Replace2(
							ModelGongCoderFieldSubTemplateCode[ModelGongCoderFieldNameFloat64],
							"{{FieldName}}", field.Name,
							"{{Value}}", fmt.Sprintf("%f", float64(idx)))
					case types.Bool:
						value := ""
						if boolFiedIndex > 1 {
							log.Println("WARNING : encoding of more than 2 boolean field per struct is not possible. ",
								gongStruct.Name)
						}
						if boolFiedIndex%2 == 0 {
							value = "false"
						} else {
							value = "true"
						}
						boolFiedIndex++
						fieldCode += models.Replace2(
							ModelGongCoderFieldSubTemplateCode[ModelGongCoderFieldCodeBoolean],
							"{{FieldName}}", field.Name,
							"{{Value}}", value)
						fieldNameBool += models.Replace2(
							ModelGongCoderFieldSubTemplateCode[ModelGongCoderFieldNameBoolean],
							"{{FieldName}}", field.Name,
							"{{Value}}", value)
					}
				case *models.GongTimeField:
					fieldCode += models.Replace2(
						ModelGongCoderFieldSubTemplateCode[ModelGongCoderFieldCodeDate],
						"{{FieldName}}", field.Name,
						"{{Value}}", fmt.Sprintf("%d", idx))
					fieldNameDate += models.Replace2(
						ModelGongCoderFieldSubTemplateCode[ModelGongCoderFieldNameDate],
						"{{FieldName}}", field.Name,
						"{{Value}}", fmt.Sprintf("%d", idx))
				}
			}

			generatedCodeFromSubTemplate := models.Replace8(ModelGongCoderStructSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{FieldCode}}", fieldCode,
				"{{FieldNameString}}", fieldNameString,
				"{{FieldNameInt}}", fieldNameInt,
				"{{FieldNameFloat64}}", fieldNameFloat64,
				"{{FieldNameDate}}", fieldNameDate,
				"{{FieldNameBoolean}}", fieldNameBool,
			)

			subStructCodes[subStructTemplate] += generatedCodeFromSubTemplate
		}
	}

	for insertionPerStructId := ModelGongCoderStructInsertionId(0); insertionPerStructId < ModelGongCoderStructInsertionsNb; insertionPerStructId++ {
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

	file, err := os.Create(filepath.Join(pkgPath, "gong_coder.go"))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)

}
