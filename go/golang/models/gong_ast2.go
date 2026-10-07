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

const ModelGongAst2Template = `// generated code - do not edit
package {{PkgGoName}}

import (
	"embed"
	"go/ast"
	"go/token"
	"log"
	"time"

	gong_runtime "github.com/fullstack-lang/gong/pkg/runtime"
)

var _ = time.Hour

// swagger:ignore
type GONG__ExpressionType = gong_runtime.GONG__ExpressionType

const (
	GONG__STRUCT_INSTANCE      = gong_runtime.GONG__STRUCT_INSTANCE
	GONG__FIELD_OR_CONST_VALUE = gong_runtime.GONG__FIELD_OR_CONST_VALUE
	GONG__FIELD_VALUE          = gong_runtime.GONG__FIELD_VALUE
	GONG__ENUM_CAST_INT        = gong_runtime.GONG__ENUM_CAST_INT
	GONG__ENUM_CAST_STRING     = gong_runtime.GONG__ENUM_CAST_STRING
	GONG__IDENTIFIER_CONST     = gong_runtime.GONG__IDENTIFIER_CONST
)

// ------------------------------------------------------------------------------------------------
// STATIC AST PARSING LOGIC
// ------------------------------------------------------------------------------------------------

// GongModelUnmarshaller abstracts the logic for setting fields on a staged instance
type GongModelUnmarshaller interface {
	// Initialize creates the struct, stages it, and returns the pointer as 'any'
	Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error)

	// UnmarshallField sets a field's value based on the AST expression
	UnmarshallField(stage *Stage, instance GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error
}

type ModelUnmarshaller = GongModelUnmarshaller

// ParseAstFile Parse pathToFile and stages all instances declared in the file
func (stage *Stage) ParseAstFile(pathToFile string, preserveOrder bool) error {
	inFile, fset, err := gong_runtime.ParseAstFile(pathToFile)
	if err != nil {
		return err
	}
	return stage.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstEmbeddedFile parses the Go source code from an embedded file
func (stage *Stage) ParseAstEmbeddedFile(directory embed.FS, pathToFile string) error {
	inFile, fset, err := gong_runtime.ParseAstEmbeddedFile(directory, pathToFile, stage.GetName())
	if err != nil {
		return err
	}
	return stage.ParseAstFileFromAst(inFile, fset, false)
}

// ParseAstString parses the Go source code from a string
func (stage *Stage) ParseAstString(blob string, preserveOrder bool) error {
	inFile, fset, err := gong_runtime.ParseAstString(blob)
	if err != nil {
		return err
	}
	return stage.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstFileFromAst traverses the AST and stages instances using the Unmarshaller registry
func (stage *Stage) ParseAstFileFromAst(inFile *ast.File, fset *token.FileSet, preserveOrder bool) error {
	gong_runtime.CheckModuleVersion(stage.GetProbeIF(), inFile)

	// 1. Remove Global Variables: Use a local map to track variable names to instances
	identifierMap := make(map[string]GongstructIF)

	for _, instance := range stage.GetInstances() {
		identifierMap[instance.GongGetIdentifier(stage)] = instance
	}

	return gong_runtime.WalkAstFile(inFile,
		func(identName string, typeName string, instanceName string) {
			if typeName != "" {
				if unmarshaller, exists := stage.GongUnmarshallers[typeName]; exists {
					instance, err := unmarshaller.Initialize(stage, identName, instanceName, preserveOrder)
					if err == nil {
						identifierMap[identName] = instance
					}
				}
			}
		},
		func(identName string, fieldName string, valueExpr ast.Expr) {
			if instance, exists := identifierMap[identName]; exists {
				typeName := instance.GongGetGongstructName()
				if unmarshaller, exists := stage.GongUnmarshallers[typeName]; exists {
					unmarshaller.UnmarshallField(stage, instance, fieldName, valueExpr, identifierMap)
				}
			}
		},
		func(identName string) {
			if instance, ok := identifierMap[identName]; ok {
				instance.UnstageVoid(stage)
			}
		},
		func() {
			if stage.IsInDeltaMode() && stage.GetNavigationMode() != GongNavigationModeNavigating {
				stage.Commit()
			} else {
				stage.ComputeInstancesNb()
				stage.ComputeReferenceAndOrders()
				if stage.OnInitCommitCallback != nil {
					stage.OnInitCommitCallback.BeforeCommit(stage)
				}
				if stage.OnInitCommitFromBackCallback != nil {
					stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
				}
				// 1. Run all Before Commit hooks
				stage.RunBeforeCommitHooks()

				// 2. Run all After Commit hooks
				stage.RunAfterCommitHooks()
			}
		},
	)
}

// --- Generic Helpers for Unmarshallers (delegating to gong_runtime) ---

var GongExtractString = gong_runtime.ExtractString
var GongExtractInt = gong_runtime.ExtractInt
var GongExtractFloat = gong_runtime.ExtractFloat
var GongExtractBool = gong_runtime.ExtractBool
var GongExtractExpr = gong_runtime.ExtractExpr
var GongExtractDate = gong_runtime.ExtractDate

func __gong__extractMiddleUint(input string) (uint, error) {
	return gong_runtime.ExtractMiddleUint(input)
}

// GongUnmarshallSliceOfPointers handles append, slices.Delete, and slices.Insert for slice fields
func GongUnmarshallSliceOfPointers[T GongstructPtr](
	slice *[]T,
	valueExpr ast.Expr,
	identifierMap map[string]GongstructIF) (err error) {
	return gong_runtime.UnmarshallSliceOfPointers(slice, valueExpr, identifierMap)
}

// GongUnmarshallPointer handles assignment of a single pointer field
func GongUnmarshallPointer[T GongstructPtr](
	ptr *T,
	valueExpr ast.Expr,
	identifierMap map[string]GongstructIF) {
	gong_runtime.UnmarshallPointer(ptr, valueExpr, identifierMap)
}

// GongUnmarshallEnum handles assignment of enum fields (via SelectorExpr or String fallback)
func GongUnmarshallEnum[T interface{ FromCodeString(string) error }](
	ptr T,
	valueExpr ast.Expr) {
	gong_runtime.UnmarshallEnum(ptr, valueExpr)
}

// GongInitialize initializes a staged instance, sets its name, and stages it
func GongInitialize[P interface {
	GongstructIF
	StagePreserveOrder(stage *Stage, order uint)
}](instance P, stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance.SetName(instanceName)
	if !preserveOrder {
		instance.StageVoid(stage)
	} else {
		if newOrder, err := __gong__extractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.StageVoid(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

// insertion point per named struct{{` + string(rune(GongAst2Unmarshaller)) + `}}`

type GongAst2GongstructInsertionId int

const (
	GongAst2Unmarshaller GongAst2GongstructInsertionId = iota
	GongAst2GongstructInsertionNb
)

var GongAst2GongstructSubTemplateCode map[GongAst2GongstructInsertionId]string = // new line
map[GongAst2GongstructInsertionId]string{
	GongAst2Unmarshaller: `
type {{Structname}}Unmarshaller struct{}

func (u *{{Structname}}Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new({{Structname}}), stage, identifier, instanceName, preserveOrder)
}

func (u *{{Structname}}Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*{{Structname}})
	_ = instance
	switch fieldName {
	// insertion point per field{{perFieldCode}}
	}
	return nil
}
`,
}

type GongAst2SubTemplateId int

const (
	GongAst2SubTmplStringField GongAst2SubTemplateId = iota
	GongAst2SubTmplIntField
	GongAst2SubTmplFloatField
	GongAst2SubTmplBoolField
	GongAst2SubTmplDateField
	GongAst2SubTmplDurationField
	GongAst2SubTmplBasicFieldEnumString
	GongAst2SubTmplBasicFieldEnumInt
	GongAst2SubTmplPointerToStruct
	GongAst2SubTmplSliceOfPointers
	GongAst2SubTmplAnyField
)

var GongAst2FileFieldFieldSubTemplateCode map[GongAst2SubTemplateId]string = // declaration of the sub templates
map[GongAst2SubTemplateId]string{

	GongAst2SubTmplStringField: `
	case "{{FieldName}}":
		instance.{{FieldName}} = GongExtractString(valueExpr)`,
	GongAst2SubTmplDateField: `
	case "{{FieldName}}":
		instance.{{FieldName}} = GongExtractDate(valueExpr)`,
	GongAst2SubTmplDurationField: `
	case "{{FieldName}}":
		instance.{{FieldName}} = time.Duration(GongExtractInt(valueExpr))`,
	GongAst2SubTmplIntField: `
	case "{{FieldName}}":
		instance.{{FieldName}} = GongExtractInt(valueExpr)`,
	GongAst2SubTmplFloatField: `
	case "{{FieldName}}":
		instance.{{FieldName}} = GongExtractFloat(valueExpr)`,
	GongAst2SubTmplBoolField: `
	case "{{FieldName}}":
		instance.{{FieldName}} = GongExtractBool(valueExpr)`,
	GongAst2SubTmplBasicFieldEnumString: `
	case "{{FieldName}}":
		GongUnmarshallEnum(&instance.{{FieldName}}, valueExpr)`,
	GongAst2SubTmplBasicFieldEnumInt: `
	case "{{FieldName}}":
		GongUnmarshallEnum(&instance.{{FieldName}}, valueExpr)`,
	GongAst2SubTmplPointerToStruct: `
	case "{{FieldName}}":
		GongUnmarshallPointer(&instance.{{FieldName}}, valueExpr, identifierMap)`,
	GongAst2SubTmplSliceOfPointers: `
	case "{{FieldName}}":
		GongUnmarshallSliceOfPointers(&instance.{{FieldName}}, valueExpr, identifierMap)`,
	GongAst2SubTmplAnyField: `
	case "{{FieldName}}":
		instance.{{FieldName}} = GongExtractExpr(valueExpr)`,
}

func GongAst2(modelPkg *models.ModelPkg, pkgPath string) {

	codeGO := ModelGongAst2Template

	subStructCodes := make(map[GongAst2GongstructInsertionId]string)
	for subStructTemplate := range GongAst2GongstructSubTemplateCode {
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

		for subStructTemplate := range GongAst2GongstructSubTemplateCode {

			perFieldCode := ""
			for _, field := range gongStruct.Fields {

				switch field := field.(type) {
				case *models.GongBasicField:

					switch field.GetBasicKind() {
					case types.String:
						if field.GongEnum == nil {
							perFieldCode += models.Replace1(
								GongAst2FileFieldFieldSubTemplateCode[GongAst2SubTmplStringField],
								"{{FieldName}}", field.Name)
						} else {
							perFieldCode += models.Replace1(
								GongAst2FileFieldFieldSubTemplateCode[GongAst2SubTmplBasicFieldEnumString],
								"{{FieldName}}", field.Name)
						}
					case types.Int, types.Int8, types.Int16, types.Int32, types.Int64:
						if field.GongEnum != nil {
							perFieldCode += models.Replace1(
								GongAst2FileFieldFieldSubTemplateCode[GongAst2SubTmplBasicFieldEnumInt],
								"{{FieldName}}", field.Name)
							break

						}
						if field.DeclaredType == "time.Duration" {
							perFieldCode += models.Replace1(
								GongAst2FileFieldFieldSubTemplateCode[GongAst2SubTmplDurationField],
								"{{FieldName}}", field.Name)
							continue
						}
						perFieldCode += models.Replace1(
							GongAst2FileFieldFieldSubTemplateCode[GongAst2SubTmplIntField],
							"{{FieldName}}", field.Name)
					case types.Float32, types.Float64:
						perFieldCode += models.Replace1(
							GongAst2FileFieldFieldSubTemplateCode[GongAst2SubTmplFloatField],
							"{{FieldName}}", field.Name)
					case types.Bool:
						perFieldCode += models.Replace1(
							GongAst2FileFieldFieldSubTemplateCode[GongAst2SubTmplBoolField],
							"{{FieldName}}", field.Name)
					case types.UntypedNil:
						perFieldCode += models.Replace1(
							GongAst2FileFieldFieldSubTemplateCode[GongAst2SubTmplAnyField],
							"{{FieldName}}", field.Name)

					default:
					}
				case *models.GongTimeField:
					perFieldCode += models.Replace1(
						GongAst2FileFieldFieldSubTemplateCode[GongAst2SubTmplDateField],
						"{{FieldName}}", field.Name)
				case *models.PointerToGongStructField:
					perFieldCode += models.Replace1(
						GongAst2FileFieldFieldSubTemplateCode[GongAst2SubTmplPointerToStruct],
						"{{FieldName}}", field.Name)
				case *models.SliceOfPointerToGongStructField:
					perFieldCode += models.Replace2(
						GongAst2FileFieldFieldSubTemplateCode[GongAst2SubTmplSliceOfPointers],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name)
				default:
				}

			}

			generatedCodeFromSubTemplate := models.Replace3(GongAst2GongstructSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,

				"{{perFieldCode}}", perFieldCode,
				// "{{cleanOfSliceOfPointers}}", cleanOfSliceOfPointers,
				// "{{cleanOfPointer}}", cleanOfPointer)
			)

			subStructCodes[subStructTemplate] += generatedCodeFromSubTemplate
		}

	}

	// substitutes {{<<insertionPerStructId points>>}} stuff with generated code
	for insertionPerStructId := range GongAst2GongstructInsertionNb {
		toReplace := "{{" + string(rune(insertionPerStructId)) + "}}"
		codeGO = strings.ReplaceAll(codeGO, toReplace, subStructCodes[insertionPerStructId])
	}

	caserEnglish := cases.Title(language.English)
	codeGO = models.Replace(codeGO,
		"{{PkgName}}", modelPkg.Name,
		"{{TitlePkgName}}", caserEnglish.String(modelPkg.Name),
		"{{pkgname}}", strings.ToLower(modelPkg.Name),
		"{{PkgGoName}}", modelPkg.PkgGoName,
		"	 | ", "	", // for the replacement of the of the first bar in the Gongstruct Type def
	)

	file, err := os.Create(filepath.Join(pkgPath, string(models.GeneratedGongAstGo2FilePath)))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)
}
