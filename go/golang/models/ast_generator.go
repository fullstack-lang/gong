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
type ModelGongAstStructInsertionId int

const (
	ModelGongAstGenericMaps ModelGongAstStructInsertionId = iota
	ModelGongAstStageProcessing
	ModelGongAstBasicLitAssignment
	ModelGongAstIdentBooleanAndPointerAssignment
	ModelGongAstIdentSelectorExprAssignment
	ModelGongAstDateAssignment
	ModelGongAstSliceOfPointersAssignment
	ModelGongAstStructInsertionsNb
)

var ModelGongAstStructSubTemplateCode map[ModelGongAstStructInsertionId]string = // new line
map[ModelGongAstStructInsertionId]string{
	ModelGongAstGenericMaps: `
var __gong__map_{{Structname}} = make(map[string]*{{Structname}})`,
	ModelGongAstStageProcessing: `
									case "{{Structname}}":
										instance{{Structname}} := new({{Structname}})
										instance{{Structname}}.Name = instanceName
										instance{{Structname}}.Stage(stage)
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
	ModelGongAstIdentSelectorExprAssignment: `
				case "{{Structname}}":
					switch fieldName {
					// insertion point for selector expr assign code{{selectorExprAssignCode}}
					}`,
	ModelGongAstDateAssignment: `
						case "{{Structname}}":
							switch fieldName {
							// insertion point for date assign code{{dateAssignCode}}
							}`,
	ModelGongAstSliceOfPointersAssignment: `
					case "{{Structname}}":
						switch fieldName {
						// insertion point for slice of pointers assign code{{sliceOfPointersAssignCode}}
						}`,
}

type ModelGongAstFieldInsertionId int

const (
	ModelGongAstFieldAssignString ModelGongAstFieldInsertionId = iota
	ModelGongAstFieldAssignMetaField
	ModelGongAstFieldAssignMetaFieldIndented
	ModelGongAstFieldAssignMetaFieldDoublyIndented
	ModelGongAstFieldAssignInt
	ModelGongAstFieldAssignFloat64
	ModelGongAstFieldAssignDate
	ModelGongAstFieldAssignDuration
	ModelGongAstFieldAssignBoolean
	ModelGongAstFieldAssignPointer
	ModelGongAstFieldAssignSliceOfPointers
	ModelGongAstFieldAssignEnum

	ModelGongAstFieldInsertionsNb
)

var ModelGongAstFieldSubTemplateCode map[ModelGongAstFieldInsertionId]string = // new line
map[ModelGongAstFieldInsertionId]string{
	ModelGongAstFieldAssignString: `
				case "{{FieldName}}":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_{{Structname}}[identifier].{{FieldName}} = fielValue`,
	ModelGongAstFieldAssignMetaField: `
				case "{{FieldName}}":
					__gong__map_{{Structname}}[identifier].{{FieldName}} = basicLit.Value`,
	ModelGongAstFieldAssignMetaFieldIndented: `
					case "{{FieldName}}":
						__gong__map_{{Structname}}[identifier].{{FieldName}} = basicLit.Value`,
	ModelGongAstFieldAssignMetaFieldDoublyIndented: `
						case "{{FieldName}}":
							__gong__map_{{Structname}}[identifier].{{FieldName}} = basicLit.Value`,
	ModelGongAstFieldAssignInt: `
				case "{{FieldName}}":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_{{Structname}}[identifier].{{FieldName}} = int(exprSign) * int(fielValue)`,
	ModelGongAstFieldAssignDuration: `
				case "{{FieldName}}":
					// convert string to duration
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_{{Structname}}[identifier].{{FieldName}} = time.Duration(int(exprSign) * int(fielValue))`,
	ModelGongAstFieldAssignFloat64: `
				case "{{FieldName}}":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_{{Structname}}[identifier].{{FieldName}} = exprSign * fielValue`,
	ModelGongAstFieldAssignBoolean: `
				case "{{FieldName}}":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_{{Structname}}[identifier].{{FieldName}} = fielValue`,
	ModelGongAstFieldAssignPointer: `
				case "{{FieldName}}":
					targetIdentifier := ident.Name
					__gong__map_{{Structname}}[identifier].{{FieldName}} = __gong__map_{{AssociationStructName}}[targetIdentifier]`,
	ModelGongAstFieldAssignEnum: `
					case "{{FieldName}}":
						var val {{EnumType}}
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_{{Structname}}[identifier].{{FieldName}} = {{EnumType}}(val)`,
	ModelGongAstFieldAssignDate: `
							case "{{FieldName}}":
								__gong__map_{{Structname}}[identifier].{{FieldName}}, _ = time.Parse(
									"2006-01-02 15:04:05.999999999 -0700 MST",
									date)`,
	ModelGongAstFieldAssignSliceOfPointers: `
						case "{{FieldName}}":
							// remove first and last char
							targetIdentifier := ident.Name
							// when parsing {{Structname}}[identifier].{{FieldName}} = append({{Structname}}[identifier].{{FieldName}}, {{AssociationStructName}} instance )
							// the map will not find the {{AssociationStructName}} instance, when parsing the first arg
							// therefore, the condition is necessary
							if target, ok := __gong__map_{{AssociationStructName}}[targetIdentifier]; ok {
								__gong__map_{{Structname}}[identifier].{{FieldName}} =
									append(__gong__map_{{Structname}}[identifier].{{FieldName}}, target)
							}`,
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
			selectorExprAssgmCode := ""
			dateAssignCode := ""
			sliceOfPointersAssignCode := ""

			for _, field := range gongStruct.Fields {
				switch field := field.(type) {
				case *models.GongBasicField:
					switch field.GetBasicKind() {
					case types.String:
						if field.GongEnum != nil {
							selectorExprAssgmCode += models.Replace2(
								ModelGongAstFieldSubTemplateCode[ModelGongAstFieldAssignEnum],
								"{{FieldName}}", field.Name,
								"{{EnumType}}", field.GongEnum.Name)
							continue
						}
						basicLitAssignCode += models.Replace1(
							ModelGongAstFieldSubTemplateCode[ModelGongAstFieldAssignString],
							"{{FieldName}}", field.Name)
					case types.Int, types.Int64:
						if field.GongEnum != nil {
							selectorExprAssgmCode += models.Replace2(
								ModelGongAstFieldSubTemplateCode[ModelGongAstFieldAssignEnum],
								"{{FieldName}}", field.Name,
								"{{EnumType}}", field.GongEnum.Name)
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
					case types.UntypedNil:
						// we have 2 possible assignments
						// one for ref_models.Astruct{}
						// one for ref_models.Astruct{}.Name
						basicLitAssignCode += models.Replace1(
							ModelGongAstFieldSubTemplateCode[ModelGongAstFieldAssignMetaField],
							"{{FieldName}}", field.Name)
						selectorExprAssgmCode += models.Replace1(
							ModelGongAstFieldSubTemplateCode[ModelGongAstFieldAssignMetaFieldIndented],
							"{{FieldName}}", field.Name)
						sliceOfPointersAssignCode += models.Replace1(
							ModelGongAstFieldSubTemplateCode[ModelGongAstFieldAssignMetaFieldDoublyIndented],
							"{{FieldName}}", field.Name)
					}
				case *models.PointerToGongStructField:
					boolAndPointerAssignCode += models.Replace2(
						ModelGongAstFieldSubTemplateCode[ModelGongAstFieldAssignPointer],
						"{{FieldName}}", field.Name,
						"{{AssociationStructName}}", field.GongStruct.Name)
				case *models.SliceOfPointerToGongStructField:
					sliceOfPointersAssignCode += models.Replace2(
						ModelGongAstFieldSubTemplateCode[ModelGongAstFieldAssignSliceOfPointers],
						"{{FieldName}}", field.Name,
						"{{AssociationStructName}}", field.GongStruct.Name)
				case *models.GongTimeField:
					dateAssignCode += models.Replace1(
						ModelGongAstFieldSubTemplateCode[ModelGongAstFieldAssignDate],
						"{{FieldName}}", field.Name)
				}
			}

			generatedCodeFromSubTemplate := models.Replace7(ModelGongAstStructSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{basicLitAssignCode}}", basicLitAssignCode,
				"{{boolAndPointerAssignCode}}", boolAndPointerAssignCode,
				"{{selectorExprAssignCode}}", selectorExprAssgmCode,
				"{{sliceOfPointersAssignCode}}", sliceOfPointersAssignCode,
				"{{dateAssignCode}}", dateAssignCode,
			)

			generatedCodeFromSubTemplate = models.Replace7(generatedCodeFromSubTemplate,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{basicLitAssignCode}}", basicLitAssignCode,
				"{{FieldNameString}}", boolAndPointerAssignCode,
				"{{selectorExprAssignCode}}", selectorExprAssgmCode,
				"{{sliceOfPointersAssignCode}}", sliceOfPointersAssignCode,
				"{{dateAssignCode}}", dateAssignCode,
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

	file, err := os.Create(filepath.Join(pkgPath, string(models.GeneratedGongAstGoFilePath)))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)
}
