package models

import (
	"fmt"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const ModelGongCoderFileTemplate = `package models

import "time"

// GongfieldCoder return an instance of Type where each field
// encodes the index of the field.
// This allows for refactorable field name
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

func GongfieldName[Type Gongstruct, FieldType Gongfield, AssociationFieldType Gongstruct](field FieldType) string {
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
	case {{Structname}}:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name{{FieldNameString}}
		case int, int64:
			// insertion point for field dependant name{{FieldNameInt}}
		case *AssociationFieldType:
			// insertion point for field dependant name{{FieldNamePointerToStruct}}
		case []*AssociationFieldType:
			// insertion point for field dependant name{{FieldNamePointerToStruct}}
		}`,
}

type ModelGongCoderFieldInsertionId int

const (
	ModelGongCoderFieldCodeString ModelGongCoderFieldInsertionId = iota
	ModelGongCoderFieldCodeInt
	ModelGongCoderFieldCodeFloat64
	ModelGongCoderFieldCodeDate
	ModelGongCoderFieldCodePointerToStruct
	ModelGongCoderFieldCodeSliceOfPointersToStruct
	ModelGongCoderFieldNameString
	ModelGongCoderFieldNameInt
	ModelGongCoderFieldNameFloat64
	ModelGongCoderFieldNameDate
	ModelGongCoderFieldNamePointerToStruct
	ModelGongCoderFieldNameSliceOfPointersToStruct
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
	ModelGongCoderFieldCodePointerToStruct: `
		fieldCoder.{{FieldName}} = &{{AssocStructName}}{Name: "{{Value}}"}`,
	ModelGongCoderFieldCodeSliceOfPointersToStruct: `
		fieldCoder.{{FieldName}} = []*{{AssocStructName}}{
			{
				Name: "{{Value}}",
			},
		}`,
	ModelGongCoderFieldNameString: `
			if field == "{{Value}}" {
				return "{{FieldName}}"
			}`,
	ModelGongCoderFieldNameInt: `
			if field == {{Value}} {
				return "{{FieldName}}"
			}`,
}

func CodeGeneratorModelGongCoder(
	mdlPkg *ModelPkg,
	pkgName string,
	pkgPath string) {

	// generate the typescript file
	codeGO := ModelGongCoderFileTemplate

	subStructCodes := make(map[ModelGongCoderStructInsertionId]string)
	for subStructTemplate := range ModelGongCoderStructSubTemplateCode {
		subStructCodes[subStructTemplate] = ""
	}

	// sort gong structs per name (for reproductibility)
	gongStructs := []*GongStruct{}
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

			for idx, field := range gongStruct.Fields {

				switch field := field.(type) {
				case *GongBasicField:
					switch field.basicKind {
					case types.String:
						fieldCode += Replace2(
							ModelGongCoderFieldSubTemplateCode[ModelGongCoderFieldCodeString],
							"{{FieldName}}", field.Name,
							"{{Value}}", fmt.Sprintf("%d", idx))
						fieldNameString += Replace2(
							ModelGongCoderFieldSubTemplateCode[ModelGongCoderFieldNameString],
							"{{FieldName}}", field.Name,
							"{{Value}}", fmt.Sprintf("%d", idx))
					case types.Int, types.Int64:
						fieldCode += Replace2(
							ModelGongCoderFieldSubTemplateCode[ModelGongCoderFieldCodeInt],
							"{{FieldName}}", field.Name,
							"{{Value}}", fmt.Sprintf("%d", idx))
						fieldNameInt += Replace2(
							ModelGongCoderFieldSubTemplateCode[ModelGongCoderFieldNameString],
							"{{FieldName}}", field.Name,
							"{{Value}}", fmt.Sprintf("%d", idx))
					case types.Float64:
						fieldCode += Replace2(
							ModelGongCoderFieldSubTemplateCode[ModelGongCoderFieldCodeInt],
							"{{FieldName}}", field.Name,
							"{{Value}}", fmt.Sprintf("%f", float64(idx)))
					}
				case *GongTimeField:
					fieldCode += Replace2(
						ModelGongCoderFieldSubTemplateCode[ModelGongCoderFieldCodeDate],
						"{{FieldName}}", field.Name,
						"{{Value}}", fmt.Sprintf("%d", idx))
				case *PointerToGongStructField:
					fieldCode += Replace3(
						ModelGongCoderFieldSubTemplateCode[ModelGongCoderFieldCodePointerToStruct],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name,
						"{{Value}}", fmt.Sprintf("%d", idx))
				case *SliceOfPointerToGongStructField:
					fieldCode += Replace3(
						ModelGongCoderFieldSubTemplateCode[ModelGongCoderFieldCodeSliceOfPointersToStruct],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name,
						"{{Value}}", fmt.Sprintf("%d", idx))
				}
			}

			generatedCodeFromSubTemplate := Replace5(ModelGongCoderStructSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{FieldCode}}", fieldCode,
				"{{FieldNameString}}", fieldNameString,
				"{{FieldNameInt}}", fieldNameInt,
			)

			subStructCodes[subStructTemplate] += generatedCodeFromSubTemplate
		}
	}

	for insertionPerStructId := ModelGongCoderStructInsertionId(0); insertionPerStructId < ModelGongCoderStructInsertionsNb; insertionPerStructId++ {
		toReplace := "{{" + string(rune(insertionPerStructId)) + "}}"
		codeGO = strings.ReplaceAll(codeGO, toReplace, subStructCodes[insertionPerStructId])
	}

	codeGO = Replace4(codeGO,
		"{{PkgName}}", pkgName,
		"{{TitlePkgName}}", strings.Title(pkgName),
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
