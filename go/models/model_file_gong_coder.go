package models

import (
	"fmt"
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
	string | bool | int | float64 | time.Time | time.Duration{{` + string(rune(ModelGongCoderGenericGongstructTypes)) + `}}
}

func GongfieldName[Type Gongstruct, FieldType Gongfield](field FieldType) string {
	var t Type

	switch any(t).(type) {
	}
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
	ModelGongCoderStructInsertionsNb
)

var ModelGongCoderStructSubTemplateCode map[ModelGongCoderStructInsertionId]string = // new line
map[ModelGongCoderStructInsertionId]string{
	ModelGongCoderGenericGongstructTypes: ` | *{{Structname}} | []*{{Structname}}`,
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
			generatedCodeFromSubTemplate := Replace2(ModelGongCoderStructSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
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
