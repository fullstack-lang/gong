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
type ModelGongAstStructInsertionId int

const (
	ModelGongAstGenericMaps ModelGongAstStructInsertionId = iota
	ModelGongAstStageProcessing
	ModelGongAstBasicLitAssignment
	ModelGongAstIdentBooleanAndPointerAssignment
	ModelGongAstStructInsertionsNb
)

var ModelGongAstStructSubTemplateCode map[ModelGongAstStructInsertionId]string = // new line
map[ModelGongAstStructInsertionId]string{
	ModelGongAstGenericMaps: `
var __gong__map_{{Structname}} = make(map[string]*{{Structname}})`,
	ModelGongAstStageProcessing: `
									case "{{Structname}}":
										instance{{Structname}} := (&{{Structname}}{Name: instanceName}).Stage()
										instance = any(instance{{Structname}})
										__gong__map_{{Structname}}[identifier] = instance{{Structname}}`,
	ModelGongAstBasicLitAssignment: `
			case "{{Structname}}":
				switch fieldName {
				// insertion point for field dependant code{{basicLitAssignCode}}
				}`,
	ModelGongAstIdentBooleanAndPointerAssignment: `
			case "{{Structname}}":
				switch fieldName {
				// insertion point for field dependant code{{boolAndPointerAssignCode}}
				}`,
}

type ModelGongAstFieldInsertionId int

const (
	ModelGongAstFieldAssignString ModelGongAstFieldInsertionId = iota
	ModelGongAstFieldAssignInt
	ModelGongAstFieldAssignFloat64
	ModelGongAstFieldAssignDate
	ModelGongAstFieldAssignDuration
	ModelGongAstFieldAssignBoolean

	ModelGongAstFieldInsertionsNb
)

var ModelGongAstFieldSubTemplateCode map[ModelGongAstFieldInsertionId]string = // new line
map[ModelGongAstFieldInsertionId]string{
	ModelGongAstFieldAssignString: `
				case "{{FieldName}}":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_{{Structname}}[identifier].{{FieldName}} = fielValue`,
	ModelGongAstFieldAssignInt: `
				case "{{FieldName}}":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_{{Structname}}[identifier].{{FieldName}} = int(fielValue)`,
	ModelGongAstFieldAssignDuration: `
				case "{{FieldName}}":
					// convert string to duration
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_{{Structname}}[identifier].{{FieldName}} = time.Duration(fielValue)`,
	ModelGongAstFieldAssignFloat64: `
				case "{{FieldName}}":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_{{Structname}}[identifier].{{FieldName}} = fielValue`,
	ModelGongAstFieldAssignBoolean: `
				case "{{FieldName}}":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_{{Structname}}[identifier].{{FieldName}} = fielValue`,
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
			basicLitAssignCode := ""
			boolAndPointerAssignCode := ""
			fieldNameInt := ""
			fieldNameBool := ""
			fieldNameFloat64 := ""
			fieldNameDate := ""
			fieldNameDuration := ""

			for _, field := range gongStruct.Fields {
				switch field := field.(type) {
				case *models.GongBasicField:
					switch field.GetBasicKind() {
					case types.String:
						if field.GongEnum != nil {
							continue
						}
						basicLitAssignCode += models.Replace1(
							ModelGongAstFieldSubTemplateCode[ModelGongAstFieldAssignString],
							"{{FieldName}}", field.Name)
					case types.Int, types.Int64:
						if field.GongEnum != nil {
							continue
						}
						if field.DeclaredType == "time.Duration" {
							basicLitAssignCode += models.Replace1(
								ModelGongAstFieldSubTemplateCode[ModelGongAstFieldAssignDuration],
								"{{FieldName}}", field.Name)
							continue
						}
						basicLitAssignCode += models.Replace1(
							ModelGongAstFieldSubTemplateCode[ModelGongAstFieldAssignInt],
							"{{FieldName}}", field.Name)
					case types.Float64:
						basicLitAssignCode += models.Replace1(
							ModelGongAstFieldSubTemplateCode[ModelGongAstFieldAssignFloat64],
							"{{FieldName}}", field.Name)
					case types.Bool:
						boolAndPointerAssignCode += models.Replace1(
							ModelGongAstFieldSubTemplateCode[ModelGongAstFieldAssignBoolean],
							"{{FieldName}}", field.Name)
					}

				}
			}

			generatedCodeFromSubTemplate := models.Replace9(ModelGongAstStructSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{basicLitAssignCode}}", basicLitAssignCode,
				"{{boolAndPointerAssignCode}}", boolAndPointerAssignCode,
				"{{FieldNameInt}}", fieldNameInt,
				"{{FieldNameFloat64}}", fieldNameFloat64,
				"{{FieldNameDate}}", fieldNameDate,
				"{{FieldNameDuration}}", fieldNameDuration,
				"{{FieldNameBoolean}}", fieldNameBool,
			)

			generatedCodeFromSubTemplate = models.Replace8(generatedCodeFromSubTemplate,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{basicLitAssignCode}}", basicLitAssignCode,
				"{{FieldNameString}}", boolAndPointerAssignCode,
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
