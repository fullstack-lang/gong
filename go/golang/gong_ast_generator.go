package golang

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

// insertion points are places where the code is
// generated per gong struct
type ModelGongAstStructInsertionId int

const (
	ModelGongAstGenericMaps ModelGongAstStructInsertionId = iota
	ModelGongAstGenericLhs
	ModelGongAstStructInsertionsNb
)

var ModelGongAstStructSubTemplateCode map[ModelGongAstStructInsertionId]string = // new line
map[ModelGongAstStructInsertionId]string{
	ModelGongAstGenericMaps: `
var __gong__map_{{Structname}} = make(map[string]*{{Structname}})`,
	ModelGongAstGenericLhs: `
									case "{{Structname}}":
										instance{{Structname}} := (&{{Structname}}{Name: instanceName}).Stage()
										instance = any(instance{{Structname}})
										__gong__map_{{Structname}}[identifier] = instance{{Structname}}`,
}

func GongAstGenerator(modelPkg *models.ModelPkg, pkgPath string) {

	codeGO := GongAstTemplate

	subStructCodes := make(map[ModelGongAstStructInsertionId]string)
	for subStructTemplate := range ModelGongAstStructSubTemplateCode {
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

		for subStructTemplate := range ModelGongAstStructSubTemplateCode {
			fieldCode := ""
			fieldNameString := ""
			fieldNameInt := ""
			fieldNameBool := ""
			fieldNameFloat64 := ""
			fieldNameDate := ""

			generatedCodeFromSubTemplate := models.Replace8(ModelGongAstStructSubTemplateCode[subStructTemplate],
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

	for insertionPerStructId := ModelGongAstStructInsertionId(0); insertionPerStructId < ModelGongAstStructInsertionsNb; insertionPerStructId++ {
		toReplace := "{{" + string(rune(insertionPerStructId)) + "}}"
		isPresent := strings.Contains(codeGO, toReplace)
		if !isPresent {
			log.Fatal("insertion point " + toReplace + " not found in template")
		}
		codeGO = strings.ReplaceAll(codeGO, toReplace, subStructCodes[insertionPerStructId])
	}

	caserEnglish := cases.Title(language.English)
	codeGO = models.Replace4(codeGO,
		"{{PkgName}}", modelPkg.Name,
		"{{TitlePkgName}}", caserEnglish.String(modelPkg.Name),
		"{{pkgname}}", strings.ToLower(modelPkg.Name),
		"	 | ", "	", // for the replacement of the of the first bar in the Gongstruct Type def
	)

	file, err := os.Create(filepath.Join(pkgPath, "gong_ast.go"))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)
}
