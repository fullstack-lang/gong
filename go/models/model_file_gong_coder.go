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
	default:
		return t
	}
}

type Gongfield interface {
	string | bool | int | float64 | time.Time | time.Duration
}

func GongfieldName[Type Gongstruct, FieldType Gongfield](field FieldType) string {
	var t Type

	switch any(t).(type) {
	}
	return ""
}
`

var ModelGongCoderStructSubTemplateCode map[ModelGongCoderStructInsertionId]string = // new line
map[ModelGongCoderStructInsertionId]string{}

//
// insertion points are places where the code is
// generated per gong struct
//
type ModelGongCoderStructInsertionId int

const (
	ModelGongCoder ModelGongCoderStructInsertionId = iota
)

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

			// replace {{ValuesInitialization}}
			valInitCode := ""
			pointerInitCode := ""
			fieldNames := `
		res = []string{`
			fieldStringValues := ``
			fieldReverseAssociationMapCreationCode := ``
			fieldReversePointerAssociationMapCode := ``
			fieldReverseSliceOfPointersAssociationMapCode := ``
			associationFieldInitialization := ``

			for idx, field := range gongStruct.Fields {

				switch field := field.(type) {
				case *GongBasicField:

					switch field.basicKind {
					case types.String:
					case types.Bool:
					case types.Float64:
					case types.Int, types.Int64:
					default:
					}
				case *GongTimeField:
				case *PointerToGongStructField:
				case *SliceOfPointerToGongStructField:
				default:
				}

				if idx > 0 {
					fieldNames += ", "
				}
				fieldNames += Replace1(
					GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringFieldName],
					"{{FieldName}}", field.GetName())
			}

			valInitCode = Replace2(valInitCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			pointerInitCode = Replace2(pointerInitCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			fieldStringValues = Replace2(fieldStringValues,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			fieldReverseAssociationMapCreationCode = Replace2(fieldReverseAssociationMapCreationCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			fieldReversePointerAssociationMapCode = Replace2(fieldReversePointerAssociationMapCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			fieldReverseSliceOfPointersAssociationMapCode = Replace2(fieldReverseSliceOfPointersAssociationMapCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			fieldNames += `}`
			generatedCodeFromSubTemplate := Replace10(ModelGongCoderStructSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{ValuesInitialization}}", valInitCode,
				"{{PointersInitialization}}", pointerInitCode,
				"{{ListOfFieldsName}}", fieldNames,
				"{{StringValueOfFields}}", fieldStringValues,
				"{{ReverseAssociationMapFunctions}}", fieldReverseAssociationMapCreationCode,
				"{{fieldReversePointerAssociationMapCode}}", fieldReversePointerAssociationMapCode,
				"{{fieldReverseSliceOfPointersAssociationMapCode}}", fieldReverseSliceOfPointersAssociationMapCode,
				"{{associationFieldInitialization}}", associationFieldInitialization,
			)

			subStructCodes[subStructTemplate] += generatedCodeFromSubTemplate
		}
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
