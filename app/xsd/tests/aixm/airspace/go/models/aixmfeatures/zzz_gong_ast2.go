// generated code - do not edit
package aixmfeatures

import (
	"embed"
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"slices"
	"strings"
	"time"

	gong_runtime "github.com/fullstack-lang/gong/pkg/runtime"
)

var (
	_ = errors.New
	_ = parser.ParseFile
	_ = filepath.Base
	_ = slices.Index[[]int, int]
	_ = strings.Trim
	_ = time.Hour
)

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
	var inFile *ast.File
	var fset *token.FileSet
	var err error
	if strings.HasPrefix(strings.TrimSpace(blob), "package ") || strings.Contains(blob, "\npackage ") {
		fset = token.NewFileSet()
		inFile, err = parser.ParseFile(fset, "", blob, parser.ParseComments)
		if err != nil {
			return errors.New("Unable to parse " + err.Error())
		}
	} else {
		inFile, fset, err = gong_runtime.ParseAstString(blob)
		if err != nil {
			return err
		}
	}
	return stage.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstFileFromAst traverses the AST and stages instances across stage and its sub-stages
func (stage *Stage) ParseAstFileFromAst(inFile *ast.File, fset *token.FileSet, preserveOrder bool) error {
	gong_runtime.CheckModuleVersion(stage.GetProbeIF(), inFile)

	identifierMap := make(map[string]any)

	if stage != nil {
		for _, instance := range stage.GetInstances() {
			identifierMap[instance.GongGetIdentifier(stage)] = instance
		}
	}

	aliasToCanonical := make(map[string]string)
	for _, imp := range inFile.Imports {
		p := strings.Trim(imp.Path.Value, "\"`")
		alias := filepath.Base(p)
		if imp.Name != nil {
			alias = imp.Name.Name
		}
		switch p {
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/aixmfeatures":
			aliasToCanonical[alias] = "aixmfeatures"
		}
	}

	ast.Inspect(inFile, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.AssignStmt:
			if len(node.Lhs) < 1 || len(node.Rhs) < 1 {
				return true
			}

			// CASE 1: Initialization ( := )
			if node.Tok == token.DEFINE {
				if ident, ok := node.Lhs[0].(*ast.Ident); ok {
					var pkgAlias string
					var typeName string
					var instanceName string

					ast.Inspect(node.Rhs[0], func(expr ast.Node) bool {
						if compLit, ok := expr.(*ast.CompositeLit); ok {
							if selExpr, ok := compLit.Type.(*ast.SelectorExpr); ok {
								if pkgId, ok := selExpr.X.(*ast.Ident); ok {
									pkgAlias = pkgId.Name
								}
								typeName = selExpr.Sel.Name
								for _, elt := range compLit.Elts {
									if kv, ok := elt.(*ast.KeyValueExpr); ok {
										if k, ok := kv.Key.(*ast.Ident); ok && k.Name == "Name" {
											if v, ok := kv.Value.(*ast.BasicLit); ok {
												instanceName = strings.Trim(v.Value, "\"`")
											}
										}
									}
								}
								return false
							}
						}
						return true
					})

					if canonical, ok := aliasToCanonical[pkgAlias]; ok {
						pkgAlias = canonical
					} else if pkgAlias == "" || pkgAlias == "aixmfeatures" || len(inFile.Imports) == 0 {
						pkgAlias = "aixmfeatures"
					}

					switch pkgAlias {
			case "aixmfeatures":
				switch typeName {
				case "A_extension":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_1":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_1), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_10":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_10), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_11":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_11), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_12":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_12), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_13":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_13), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_14":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_14), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_2":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_2), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_3":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_3), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_4":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_4), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_5":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_5), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_6":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_6), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_7":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_7), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_8":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_8), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_9":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_9), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractPropertiesWithScheduleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractPropertiesWithScheduleType), stage, ident.Name, instanceName, preserveOrder)
				case "AircraftCharacteristicPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AircraftCharacteristicPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AircraftCharacteristicType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AircraftCharacteristicType), stage, ident.Name, instanceName, preserveOrder)
				case "AirspaceActivationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirspaceActivationPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AirspaceActivationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirspaceActivationType), stage, ident.Name, instanceName, preserveOrder)
				case "AirspaceGeometryComponentPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirspaceGeometryComponentPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AirspaceGeometryComponentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirspaceGeometryComponentType), stage, ident.Name, instanceName, preserveOrder)
				case "AirspaceLayerClassPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirspaceLayerClassPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AirspaceLayerClassType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirspaceLayerClassType), stage, ident.Name, instanceName, preserveOrder)
				case "AirspaceLayerPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirspaceLayerPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AirspaceLayerType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirspaceLayerType), stage, ident.Name, instanceName, preserveOrder)
				case "AirspacePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirspacePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AirspaceTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirspaceTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AirspaceTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirspaceTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "AirspaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirspaceType), stage, ident.Name, instanceName, preserveOrder)
				case "AirspaceVolumeDependencyPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirspaceVolumeDependencyPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AirspaceVolumeDependencyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirspaceVolumeDependencyType), stage, ident.Name, instanceName, preserveOrder)
				case "AirspaceVolumePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirspaceVolumePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AirspaceVolumeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirspaceVolumeType), stage, ident.Name, instanceName, preserveOrder)
				case "AuthorityForAirspaceTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AuthorityForAirspaceTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AuthorityForAirspaceTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AuthorityForAirspaceTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "AuthorityForAirspaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AuthorityForAirspaceType), stage, ident.Name, instanceName, preserveOrder)
				case "CurvePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CurvePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "CurveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CurveType), stage, ident.Name, instanceName, preserveOrder)
				case "ElevatedCurveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ElevatedCurveType), stage, ident.Name, instanceName, preserveOrder)
				case "ElevatedSurfaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ElevatedSurfaceType), stage, ident.Name, instanceName, preserveOrder)
				case "GeoBorderTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GeoBorderTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "GeoBorderTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GeoBorderTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "GeoBorderType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GeoBorderType), stage, ident.Name, instanceName, preserveOrder)
				case "LinguisticNotePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LinguisticNotePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "LinguisticNoteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LinguisticNoteType), stage, ident.Name, instanceName, preserveOrder)
				case "NotePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NotePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "NoteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteType), stage, ident.Name, instanceName, preserveOrder)
				case "OrganisationAuthorityPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(OrganisationAuthorityPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RoutePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RoutePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "StandardLevelColumnPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StandardLevelColumnPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SurfacePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SurfacePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SurfaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SurfaceType), stage, ident.Name, instanceName, preserveOrder)
				case "TimesheetPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TimesheetPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TimesheetType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TimesheetType), stage, ident.Name, instanceName, preserveOrder)
				}
					}
				}
				return false
			}

			// CASE 2: Assignment ( = )
			if node.Tok == token.ASSIGN {
				if selExpr, ok := node.Lhs[0].(*ast.SelectorExpr); ok {
					if ident, ok := selExpr.X.(*ast.Ident); ok {
						if instance, exists := identifierMap[ident.Name]; exists {
							fieldName := selExpr.Sel.Name
							rhs := node.Rhs[0]
							switch inst := instance.(type) {
				case *A_extension:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_1:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_10:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_11:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_12:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_13:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_14:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_2:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_3:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_4:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_5:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_6:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_7:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_8:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_9:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *AbstractPropertiesWithScheduleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					}
				case *AircraftCharacteristicPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AircraftCharacteristic":
						__gong__assignPointer(&inst.AircraftCharacteristic, rhs, identifierMap)
					}
				case *AircraftCharacteristicType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AirspaceActivationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AirspaceActivation":
						__gong__assignPointer(&inst.AirspaceActivation, rhs, identifierMap)
					}
				case *AirspaceActivationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeInterval":
						__gong__assignSliceOfPointers(&inst.TimeInterval, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "SpecialDateAuthority":
						__gong__assignSliceOfPointers(&inst.SpecialDateAuthority, rhs, identifierMap)
					case "Levels":
						__gong__assignSliceOfPointers(&inst.Levels, rhs, identifierMap)
					case "User":
						__gong__assignSliceOfPointers(&inst.User, rhs, identifierMap)
					case "Aircraft":
						__gong__assignSliceOfPointers(&inst.Aircraft, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AirspaceGeometryComponentPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AirspaceGeometryComponent":
						__gong__assignPointer(&inst.AirspaceGeometryComponent, rhs, identifierMap)
					}
				case *AirspaceGeometryComponentType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "TheAirspaceVolume":
						__gong__assignPointer(&inst.TheAirspaceVolume, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AirspaceLayerClassPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AirspaceLayerClass":
						__gong__assignPointer(&inst.AirspaceLayerClass, rhs, identifierMap)
					}
				case *AirspaceLayerClassType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeInterval":
						__gong__assignSliceOfPointers(&inst.TimeInterval, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "SpecialDateAuthority":
						__gong__assignSliceOfPointers(&inst.SpecialDateAuthority, rhs, identifierMap)
					case "AssociatedLevels":
						__gong__assignSliceOfPointers(&inst.AssociatedLevels, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AirspaceLayerPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AirspaceLayer":
						__gong__assignPointer(&inst.AirspaceLayer, rhs, identifierMap)
					}
				case *AirspaceLayerType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "DiscreteLevelSeries":
						__gong__assignPointer(&inst.DiscreteLevelSeries, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AirspacePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Arcrole":
						inst.Arcrole = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					}
				case *AirspaceTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AirspaceTimeSlice":
						__gong__assignPointer(&inst.AirspaceTimeSlice, rhs, identifierMap)
					}
				case *AirspaceTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "Class":
						__gong__assignSliceOfPointers(&inst.Class, rhs, identifierMap)
					case "ProtectedRoute":
						__gong__assignPointer(&inst.ProtectedRoute, rhs, identifierMap)
					case "GeometryComponent":
						__gong__assignSliceOfPointers(&inst.GeometryComponent, rhs, identifierMap)
					case "Activation":
						__gong__assignSliceOfPointers(&inst.Activation, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AirspaceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *AirspaceVolumeDependencyPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AirspaceVolumeDependency":
						__gong__assignPointer(&inst.AirspaceVolumeDependency, rhs, identifierMap)
					}
				case *AirspaceVolumeDependencyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "TheAirspace":
						__gong__assignPointer(&inst.TheAirspace, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AirspaceVolumePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AirspaceVolume":
						__gong__assignPointer(&inst.AirspaceVolume, rhs, identifierMap)
					}
				case *AirspaceVolumeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "HorizontalProjection":
						__gong__assignPointer(&inst.HorizontalProjection, rhs, identifierMap)
					case "Centreline":
						__gong__assignPointer(&inst.Centreline, rhs, identifierMap)
					case "ContributorAirspace":
						__gong__assignPointer(&inst.ContributorAirspace, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AuthorityForAirspaceTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AuthorityForAirspaceTimeSlice":
						__gong__assignPointer(&inst.AuthorityForAirspaceTimeSlice, rhs, identifierMap)
					}
				case *AuthorityForAirspaceTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "ResponsibleOrganisation":
						__gong__assignPointer(&inst.ResponsibleOrganisation, rhs, identifierMap)
					case "AssignedAirspace":
						__gong__assignPointer(&inst.AssignedAirspace, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AuthorityForAirspaceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *CurvePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Curve":
						__gong__assignPointer(&inst.Curve, rhs, identifierMap)
					case "ElevatedCurve":
						__gong__assignPointer(&inst.ElevatedCurve, rhs, identifierMap)
					}
				case *CurveType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					}
				case *ElevatedCurveType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *ElevatedSurfaceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *GeoBorderTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "GeoBorderTimeSlice":
						__gong__assignPointer(&inst.GeoBorderTimeSlice, rhs, identifierMap)
					}
				case *GeoBorderTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "Border":
						__gong__assignPointer(&inst.Border, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *GeoBorderType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *LinguisticNotePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "LinguisticNote":
						__gong__assignPointer(&inst.LinguisticNote, rhs, identifierMap)
					}
				case *LinguisticNoteType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *NotePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Note":
						__gong__assignPointer(&inst.Note, rhs, identifierMap)
					}
				case *NoteType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TranslatedNote":
						__gong__assignSliceOfPointers(&inst.TranslatedNote, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *OrganisationAuthorityPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Arcrole":
						inst.Arcrole = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					}
				case *RoutePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Arcrole":
						inst.Arcrole = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					}
				case *StandardLevelColumnPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Arcrole":
						inst.Arcrole = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					}
				case *SurfacePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Surface":
						__gong__assignPointer(&inst.Surface, rhs, identifierMap)
					case "ElevatedSurface":
						__gong__assignPointer(&inst.ElevatedSurface, rhs, identifierMap)
					}
				case *SurfaceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					}
				case *TimesheetPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Timesheet":
						__gong__assignPointer(&inst.Timesheet, rhs, identifierMap)
					}
				case *TimesheetType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
							}
						}
					}
				}
			}
		case *ast.ExprStmt:
			if call, ok := node.X.(*ast.CallExpr); ok {
				if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
					if sel.Sel.Name == "Unstage" {
						if ident, ok := sel.X.(*ast.Ident); ok {
							if instance, ok := identifierMap[ident.Name]; ok {
								if inst, ok := instance.(GongstructIF); ok {
									inst.UnstageVoid(stage)
								}
							}
						}
					}
					if sel.Sel.Name == "Commit" {
						if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == "stage" {
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
								stage.RunBeforeCommitHooks()
								stage.RunAfterCommitHooks()
							}
						}
					}
				}
			}
		}
		return true
	})

	return nil
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

func __gong__stageSetInit[P interface {
	SetName(string)
	StageVoid(S)
	StagePreserveOrder(S, uint)
}, S any](instance P, stage S, identifier string, instanceName string, preserveOrder bool) any {
	instance.SetName(instanceName)
	if !preserveOrder {
		instance.StageVoid(stage)
	} else {
		if order, err := __gong__extractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifier", identifier)
			instance.StageVoid(stage)
		} else {
			instance.StagePreserveOrder(stage, order)
		}
	}
	return instance
}

func __gong__assignPointer[T any](targetPtr **T, rhs ast.Expr, identifierMap map[string]any) {
	if rIdent, ok := rhs.(*ast.Ident); ok {
		if rIdent.Name == "nil" {
			*targetPtr = nil
			return
		}
		if target, ok := identifierMap[rIdent.Name]; ok {
			if typedTarget, ok := target.(*T); ok {
				*targetPtr = typedTarget
			}
		}
	}
}

func __gong__assignSliceOfPointers[T any](slice *[]*T, rhs ast.Expr, identifierMap map[string]any) {
	if call, ok := rhs.(*ast.CallExpr); ok {
		funcName := ""
		var isSlices bool
		if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
			if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == "slices" {
				isSlices = true
				funcName = sel.Sel.Name
			}
		} else if ident, ok := call.Fun.(*ast.Ident); ok {
			funcName = ident.Name
		}

		if isSlices {
			if funcName == "Delete" && len(call.Args) == 3 {
				start := GongExtractInt(call.Args[1])
				end := GongExtractInt(call.Args[2])
				if end <= len(*slice) {
					*slice = slices.Delete(*slice, start, end)
				}
			} else if funcName == "Insert" && len(call.Args) == 3 {
				index := GongExtractInt(call.Args[1])
				if index <= len(*slice) {
					if rIdent, ok := call.Args[2].(*ast.Ident); ok {
						if target, ok := identifierMap[rIdent.Name]; ok {
							if typedTarget, ok := target.(*T); ok {
								*slice = slices.Insert(*slice, index, typedTarget)
							}
						}
					}
				}
			}
		} else if funcName == "append" && len(call.Args) >= 2 {
			for i := 1; i < len(call.Args); i++ {
				if rIdent, ok := call.Args[i].(*ast.Ident); ok {
					if target, ok := identifierMap[rIdent.Name]; ok {
						if typedTarget, ok := target.(*T); ok {
							*slice = append(*slice, typedTarget)
						}
					}
				}
			}
		}
	}
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

// insertion point per named struct
type A_extensionUnmarshaller struct{}

func (u *A_extensionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extensionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_1Unmarshaller struct{}

func (u *A_extension_1Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_1), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_1Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_1)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_10Unmarshaller struct{}

func (u *A_extension_10Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_10), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_10Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_10)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_11Unmarshaller struct{}

func (u *A_extension_11Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_11), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_11Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_11)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_12Unmarshaller struct{}

func (u *A_extension_12Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_12), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_12Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_12)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_13Unmarshaller struct{}

func (u *A_extension_13Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_13), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_13Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_13)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_14Unmarshaller struct{}

func (u *A_extension_14Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_14), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_14Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_14)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_2Unmarshaller struct{}

func (u *A_extension_2Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_2), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_2Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_2)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_3Unmarshaller struct{}

func (u *A_extension_3Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_3), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_3Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_3)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_4Unmarshaller struct{}

func (u *A_extension_4Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_4), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_4Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_4)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_5Unmarshaller struct{}

func (u *A_extension_5Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_5), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_5Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_5)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_6Unmarshaller struct{}

func (u *A_extension_6Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_6), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_6Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_6)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_7Unmarshaller struct{}

func (u *A_extension_7Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_7), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_7Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_7)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_8Unmarshaller struct{}

func (u *A_extension_8Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_8), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_8Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_8)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_9Unmarshaller struct{}

func (u *A_extension_9Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_9), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_9Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_9)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type AbstractPropertiesWithScheduleTypeUnmarshaller struct{}

func (u *AbstractPropertiesWithScheduleTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractPropertiesWithScheduleType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractPropertiesWithScheduleTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractPropertiesWithScheduleType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	}
	return nil
}

type AircraftCharacteristicPropertyTypeUnmarshaller struct{}

func (u *AircraftCharacteristicPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AircraftCharacteristicPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AircraftCharacteristicPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AircraftCharacteristicPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "AircraftCharacteristic":
		GongUnmarshallPointer(&instance.AircraftCharacteristic, valueExpr, identifierMap)
	}
	return nil
}

type AircraftCharacteristicTypeUnmarshaller struct{}

func (u *AircraftCharacteristicTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AircraftCharacteristicType), stage, identifier, instanceName, preserveOrder)
}

func (u *AircraftCharacteristicTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AircraftCharacteristicType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AirspaceActivationPropertyTypeUnmarshaller struct{}

func (u *AirspaceActivationPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirspaceActivationPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirspaceActivationPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirspaceActivationPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "AirspaceActivation":
		GongUnmarshallPointer(&instance.AirspaceActivation, valueExpr, identifierMap)
	}
	return nil
}

type AirspaceActivationTypeUnmarshaller struct{}

func (u *AirspaceActivationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirspaceActivationType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirspaceActivationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirspaceActivationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeInterval":
		GongUnmarshallSliceOfPointers(&instance.TimeInterval, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "SpecialDateAuthority":
		GongUnmarshallSliceOfPointers(&instance.SpecialDateAuthority, valueExpr, identifierMap)
	case "Levels":
		GongUnmarshallSliceOfPointers(&instance.Levels, valueExpr, identifierMap)
	case "User":
		GongUnmarshallSliceOfPointers(&instance.User, valueExpr, identifierMap)
	case "Aircraft":
		GongUnmarshallSliceOfPointers(&instance.Aircraft, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AirspaceGeometryComponentPropertyTypeUnmarshaller struct{}

func (u *AirspaceGeometryComponentPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirspaceGeometryComponentPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirspaceGeometryComponentPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirspaceGeometryComponentPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "AirspaceGeometryComponent":
		GongUnmarshallPointer(&instance.AirspaceGeometryComponent, valueExpr, identifierMap)
	}
	return nil
}

type AirspaceGeometryComponentTypeUnmarshaller struct{}

func (u *AirspaceGeometryComponentTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirspaceGeometryComponentType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirspaceGeometryComponentTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirspaceGeometryComponentType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "TheAirspaceVolume":
		GongUnmarshallPointer(&instance.TheAirspaceVolume, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AirspaceLayerClassPropertyTypeUnmarshaller struct{}

func (u *AirspaceLayerClassPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirspaceLayerClassPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirspaceLayerClassPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirspaceLayerClassPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "AirspaceLayerClass":
		GongUnmarshallPointer(&instance.AirspaceLayerClass, valueExpr, identifierMap)
	}
	return nil
}

type AirspaceLayerClassTypeUnmarshaller struct{}

func (u *AirspaceLayerClassTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirspaceLayerClassType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirspaceLayerClassTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirspaceLayerClassType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeInterval":
		GongUnmarshallSliceOfPointers(&instance.TimeInterval, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "SpecialDateAuthority":
		GongUnmarshallSliceOfPointers(&instance.SpecialDateAuthority, valueExpr, identifierMap)
	case "AssociatedLevels":
		GongUnmarshallSliceOfPointers(&instance.AssociatedLevels, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AirspaceLayerPropertyTypeUnmarshaller struct{}

func (u *AirspaceLayerPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirspaceLayerPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirspaceLayerPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirspaceLayerPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "AirspaceLayer":
		GongUnmarshallPointer(&instance.AirspaceLayer, valueExpr, identifierMap)
	}
	return nil
}

type AirspaceLayerTypeUnmarshaller struct{}

func (u *AirspaceLayerTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirspaceLayerType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirspaceLayerTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirspaceLayerType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "DiscreteLevelSeries":
		GongUnmarshallPointer(&instance.DiscreteLevelSeries, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AirspacePropertyTypeUnmarshaller struct{}

func (u *AirspacePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirspacePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirspacePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirspacePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type AirspaceTimeSlicePropertyTypeUnmarshaller struct{}

func (u *AirspaceTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirspaceTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirspaceTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirspaceTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "AirspaceTimeSlice":
		GongUnmarshallPointer(&instance.AirspaceTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type AirspaceTimeSliceTypeUnmarshaller struct{}

func (u *AirspaceTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirspaceTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirspaceTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirspaceTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "Class":
		GongUnmarshallSliceOfPointers(&instance.Class, valueExpr, identifierMap)
	case "ProtectedRoute":
		GongUnmarshallPointer(&instance.ProtectedRoute, valueExpr, identifierMap)
	case "GeometryComponent":
		GongUnmarshallSliceOfPointers(&instance.GeometryComponent, valueExpr, identifierMap)
	case "Activation":
		GongUnmarshallSliceOfPointers(&instance.Activation, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AirspaceTypeUnmarshaller struct{}

func (u *AirspaceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirspaceType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirspaceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirspaceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type AirspaceVolumeDependencyPropertyTypeUnmarshaller struct{}

func (u *AirspaceVolumeDependencyPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirspaceVolumeDependencyPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirspaceVolumeDependencyPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirspaceVolumeDependencyPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "AirspaceVolumeDependency":
		GongUnmarshallPointer(&instance.AirspaceVolumeDependency, valueExpr, identifierMap)
	}
	return nil
}

type AirspaceVolumeDependencyTypeUnmarshaller struct{}

func (u *AirspaceVolumeDependencyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirspaceVolumeDependencyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirspaceVolumeDependencyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirspaceVolumeDependencyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "TheAirspace":
		GongUnmarshallPointer(&instance.TheAirspace, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AirspaceVolumePropertyTypeUnmarshaller struct{}

func (u *AirspaceVolumePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirspaceVolumePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirspaceVolumePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirspaceVolumePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "AirspaceVolume":
		GongUnmarshallPointer(&instance.AirspaceVolume, valueExpr, identifierMap)
	}
	return nil
}

type AirspaceVolumeTypeUnmarshaller struct{}

func (u *AirspaceVolumeTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirspaceVolumeType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirspaceVolumeTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirspaceVolumeType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "HorizontalProjection":
		GongUnmarshallPointer(&instance.HorizontalProjection, valueExpr, identifierMap)
	case "Centreline":
		GongUnmarshallPointer(&instance.Centreline, valueExpr, identifierMap)
	case "ContributorAirspace":
		GongUnmarshallPointer(&instance.ContributorAirspace, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AuthorityForAirspaceTimeSlicePropertyTypeUnmarshaller struct{}

func (u *AuthorityForAirspaceTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AuthorityForAirspaceTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AuthorityForAirspaceTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AuthorityForAirspaceTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "AuthorityForAirspaceTimeSlice":
		GongUnmarshallPointer(&instance.AuthorityForAirspaceTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type AuthorityForAirspaceTimeSliceTypeUnmarshaller struct{}

func (u *AuthorityForAirspaceTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AuthorityForAirspaceTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *AuthorityForAirspaceTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AuthorityForAirspaceTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "ResponsibleOrganisation":
		GongUnmarshallPointer(&instance.ResponsibleOrganisation, valueExpr, identifierMap)
	case "AssignedAirspace":
		GongUnmarshallPointer(&instance.AssignedAirspace, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AuthorityForAirspaceTypeUnmarshaller struct{}

func (u *AuthorityForAirspaceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AuthorityForAirspaceType), stage, identifier, instanceName, preserveOrder)
}

func (u *AuthorityForAirspaceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AuthorityForAirspaceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type CurvePropertyTypeUnmarshaller struct{}

func (u *CurvePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CurvePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *CurvePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CurvePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Curve":
		GongUnmarshallPointer(&instance.Curve, valueExpr, identifierMap)
	case "ElevatedCurve":
		GongUnmarshallPointer(&instance.ElevatedCurve, valueExpr, identifierMap)
	}
	return nil
}

type CurveTypeUnmarshaller struct{}

func (u *CurveTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CurveType), stage, identifier, instanceName, preserveOrder)
}

func (u *CurveTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CurveType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "SrsName":
		instance.SrsName = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	}
	return nil
}

type ElevatedCurveTypeUnmarshaller struct{}

func (u *ElevatedCurveTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ElevatedCurveType), stage, identifier, instanceName, preserveOrder)
}

func (u *ElevatedCurveTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ElevatedCurveType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "SrsName":
		instance.SrsName = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type ElevatedSurfaceTypeUnmarshaller struct{}

func (u *ElevatedSurfaceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ElevatedSurfaceType), stage, identifier, instanceName, preserveOrder)
}

func (u *ElevatedSurfaceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ElevatedSurfaceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "SrsName":
		instance.SrsName = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type GeoBorderTimeSlicePropertyTypeUnmarshaller struct{}

func (u *GeoBorderTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GeoBorderTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *GeoBorderTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GeoBorderTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "GeoBorderTimeSlice":
		GongUnmarshallPointer(&instance.GeoBorderTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type GeoBorderTimeSliceTypeUnmarshaller struct{}

func (u *GeoBorderTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GeoBorderTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *GeoBorderTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GeoBorderTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "Border":
		GongUnmarshallPointer(&instance.Border, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type GeoBorderTypeUnmarshaller struct{}

func (u *GeoBorderTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GeoBorderType), stage, identifier, instanceName, preserveOrder)
}

func (u *GeoBorderTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GeoBorderType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type LinguisticNotePropertyTypeUnmarshaller struct{}

func (u *LinguisticNotePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LinguisticNotePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *LinguisticNotePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LinguisticNotePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "LinguisticNote":
		GongUnmarshallPointer(&instance.LinguisticNote, valueExpr, identifierMap)
	}
	return nil
}

type LinguisticNoteTypeUnmarshaller struct{}

func (u *LinguisticNoteTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LinguisticNoteType), stage, identifier, instanceName, preserveOrder)
}

func (u *LinguisticNoteTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LinguisticNoteType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type NotePropertyTypeUnmarshaller struct{}

func (u *NotePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NotePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *NotePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NotePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Note":
		GongUnmarshallPointer(&instance.Note, valueExpr, identifierMap)
	}
	return nil
}

type NoteTypeUnmarshaller struct{}

func (u *NoteTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NoteType), stage, identifier, instanceName, preserveOrder)
}

func (u *NoteTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NoteType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TranslatedNote":
		GongUnmarshallSliceOfPointers(&instance.TranslatedNote, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type OrganisationAuthorityPropertyTypeUnmarshaller struct{}

func (u *OrganisationAuthorityPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(OrganisationAuthorityPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *OrganisationAuthorityPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*OrganisationAuthorityPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type RoutePropertyTypeUnmarshaller struct{}

func (u *RoutePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RoutePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RoutePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RoutePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type StandardLevelColumnPropertyTypeUnmarshaller struct{}

func (u *StandardLevelColumnPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StandardLevelColumnPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *StandardLevelColumnPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StandardLevelColumnPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type SurfacePropertyTypeUnmarshaller struct{}

func (u *SurfacePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SurfacePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *SurfacePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SurfacePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Surface":
		GongUnmarshallPointer(&instance.Surface, valueExpr, identifierMap)
	case "ElevatedSurface":
		GongUnmarshallPointer(&instance.ElevatedSurface, valueExpr, identifierMap)
	}
	return nil
}

type SurfaceTypeUnmarshaller struct{}

func (u *SurfaceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SurfaceType), stage, identifier, instanceName, preserveOrder)
}

func (u *SurfaceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SurfaceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "SrsName":
		instance.SrsName = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	}
	return nil
}

type TimesheetPropertyTypeUnmarshaller struct{}

func (u *TimesheetPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TimesheetPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *TimesheetPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TimesheetPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Timesheet":
		GongUnmarshallPointer(&instance.Timesheet, valueExpr, identifierMap)
	}
	return nil
}

type TimesheetTypeUnmarshaller struct{}

func (u *TimesheetTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TimesheetType), stage, identifier, instanceName, preserveOrder)
}

func (u *TimesheetTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TimesheetType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}
