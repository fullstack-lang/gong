// generated code - do not edit
package gmd

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
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/gmd":
			aliasToCanonical[alias] = "gmd"
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
					} else if pkgAlias == "" || pkgAlias == "gmd" || len(inFile.Imports) == 0 {
						pkgAlias = "gmd"
					}

					switch pkgAlias {
			case "gmd":
				switch typeName {
				case "AbstractDS_Aggregate_Type":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractDS_Aggregate_Type), stage, ident.Name, instanceName, preserveOrder)
				case "DS_Aggregate_PropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DS_Aggregate_PropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "DS_DataSet_PropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DS_DataSet_PropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "DS_DataSet_Type":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DS_DataSet_Type), stage, ident.Name, instanceName, preserveOrder)
				case "DS_Initiative_PropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DS_Initiative_PropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "DS_Initiative_Type":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DS_Initiative_Type), stage, ident.Name, instanceName, preserveOrder)
				case "DS_OtherAggregate_PropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DS_OtherAggregate_PropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "DS_OtherAggregate_Type":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DS_OtherAggregate_Type), stage, ident.Name, instanceName, preserveOrder)
				case "DS_Platform_PropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DS_Platform_PropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "DS_Platform_Type":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DS_Platform_Type), stage, ident.Name, instanceName, preserveOrder)
				case "DS_ProductionSeries_PropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DS_ProductionSeries_PropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "DS_ProductionSeries_Type":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DS_ProductionSeries_Type), stage, ident.Name, instanceName, preserveOrder)
				case "DS_Sensor_PropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DS_Sensor_PropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "DS_Sensor_Type":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DS_Sensor_Type), stage, ident.Name, instanceName, preserveOrder)
				case "DS_Series_PropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DS_Series_PropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "DS_Series_Type":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DS_Series_Type), stage, ident.Name, instanceName, preserveOrder)
				case "DS_StereoMate_PropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DS_StereoMate_PropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "DS_StereoMate_Type":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DS_StereoMate_Type), stage, ident.Name, instanceName, preserveOrder)
				case "MD_Metadata_PropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MD_Metadata_PropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "MD_Metadata_Type":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MD_Metadata_Type), stage, ident.Name, instanceName, preserveOrder)
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
				case *AbstractDS_Aggregate_Type:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Uuid":
						inst.Uuid = GongExtractString(rhs)
					case "ComposedOf":
						__gong__assignSliceOfPointers(&inst.ComposedOf, rhs, identifierMap)
					case "SeriesMetadata":
						__gong__assignSliceOfPointers(&inst.SeriesMetadata, rhs, identifierMap)
					case "Subset":
						__gong__assignSliceOfPointers(&inst.Subset, rhs, identifierMap)
					case "Superset":
						__gong__assignSliceOfPointers(&inst.Superset, rhs, identifierMap)
					}
				case *DS_Aggregate_PropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Uuidref":
						inst.Uuidref = GongExtractString(rhs)
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
					case "DS_OtherAggregate":
						__gong__assignPointer(&inst.DS_OtherAggregate, rhs, identifierMap)
					case "DS_StereoMate":
						__gong__assignPointer(&inst.DS_StereoMate, rhs, identifierMap)
					case "DS_Series":
						__gong__assignPointer(&inst.DS_Series, rhs, identifierMap)
					case "DS_Platform":
						__gong__assignPointer(&inst.DS_Platform, rhs, identifierMap)
					case "DS_Sensor":
						__gong__assignPointer(&inst.DS_Sensor, rhs, identifierMap)
					case "DS_ProductionSeries":
						__gong__assignPointer(&inst.DS_ProductionSeries, rhs, identifierMap)
					case "DS_Initiative":
						__gong__assignPointer(&inst.DS_Initiative, rhs, identifierMap)
					}
				case *DS_DataSet_PropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Uuidref":
						inst.Uuidref = GongExtractString(rhs)
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
					case "DS_DataSet":
						__gong__assignPointer(&inst.DS_DataSet, rhs, identifierMap)
					}
				case *DS_DataSet_Type:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Uuid":
						inst.Uuid = GongExtractString(rhs)
					case "Has":
						__gong__assignSliceOfPointers(&inst.Has, rhs, identifierMap)
					case "PartOf":
						__gong__assignSliceOfPointers(&inst.PartOf, rhs, identifierMap)
					}
				case *DS_Initiative_PropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Uuidref":
						inst.Uuidref = GongExtractString(rhs)
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
					case "DS_Initiative":
						__gong__assignPointer(&inst.DS_Initiative, rhs, identifierMap)
					}
				case *DS_Initiative_Type:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Uuid":
						inst.Uuid = GongExtractString(rhs)
					case "ComposedOf":
						__gong__assignSliceOfPointers(&inst.ComposedOf, rhs, identifierMap)
					case "SeriesMetadata":
						__gong__assignSliceOfPointers(&inst.SeriesMetadata, rhs, identifierMap)
					case "Subset":
						__gong__assignSliceOfPointers(&inst.Subset, rhs, identifierMap)
					case "Superset":
						__gong__assignSliceOfPointers(&inst.Superset, rhs, identifierMap)
					}
				case *DS_OtherAggregate_PropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Uuidref":
						inst.Uuidref = GongExtractString(rhs)
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
					case "DS_OtherAggregate":
						__gong__assignPointer(&inst.DS_OtherAggregate, rhs, identifierMap)
					case "DS_StereoMate":
						__gong__assignPointer(&inst.DS_StereoMate, rhs, identifierMap)
					}
				case *DS_OtherAggregate_Type:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Uuid":
						inst.Uuid = GongExtractString(rhs)
					case "ComposedOf":
						__gong__assignSliceOfPointers(&inst.ComposedOf, rhs, identifierMap)
					case "SeriesMetadata":
						__gong__assignSliceOfPointers(&inst.SeriesMetadata, rhs, identifierMap)
					case "Subset":
						__gong__assignSliceOfPointers(&inst.Subset, rhs, identifierMap)
					case "Superset":
						__gong__assignSliceOfPointers(&inst.Superset, rhs, identifierMap)
					}
				case *DS_Platform_PropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Uuidref":
						inst.Uuidref = GongExtractString(rhs)
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
					case "DS_Platform":
						__gong__assignPointer(&inst.DS_Platform, rhs, identifierMap)
					}
				case *DS_Platform_Type:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Uuid":
						inst.Uuid = GongExtractString(rhs)
					case "ComposedOf":
						__gong__assignSliceOfPointers(&inst.ComposedOf, rhs, identifierMap)
					case "SeriesMetadata":
						__gong__assignSliceOfPointers(&inst.SeriesMetadata, rhs, identifierMap)
					case "Subset":
						__gong__assignSliceOfPointers(&inst.Subset, rhs, identifierMap)
					case "Superset":
						__gong__assignSliceOfPointers(&inst.Superset, rhs, identifierMap)
					}
				case *DS_ProductionSeries_PropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Uuidref":
						inst.Uuidref = GongExtractString(rhs)
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
					case "DS_ProductionSeries":
						__gong__assignPointer(&inst.DS_ProductionSeries, rhs, identifierMap)
					}
				case *DS_ProductionSeries_Type:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Uuid":
						inst.Uuid = GongExtractString(rhs)
					case "ComposedOf":
						__gong__assignSliceOfPointers(&inst.ComposedOf, rhs, identifierMap)
					case "SeriesMetadata":
						__gong__assignSliceOfPointers(&inst.SeriesMetadata, rhs, identifierMap)
					case "Subset":
						__gong__assignSliceOfPointers(&inst.Subset, rhs, identifierMap)
					case "Superset":
						__gong__assignSliceOfPointers(&inst.Superset, rhs, identifierMap)
					}
				case *DS_Sensor_PropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Uuidref":
						inst.Uuidref = GongExtractString(rhs)
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
					case "DS_Sensor":
						__gong__assignPointer(&inst.DS_Sensor, rhs, identifierMap)
					}
				case *DS_Sensor_Type:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Uuid":
						inst.Uuid = GongExtractString(rhs)
					case "ComposedOf":
						__gong__assignSliceOfPointers(&inst.ComposedOf, rhs, identifierMap)
					case "SeriesMetadata":
						__gong__assignSliceOfPointers(&inst.SeriesMetadata, rhs, identifierMap)
					case "Subset":
						__gong__assignSliceOfPointers(&inst.Subset, rhs, identifierMap)
					case "Superset":
						__gong__assignSliceOfPointers(&inst.Superset, rhs, identifierMap)
					}
				case *DS_Series_PropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Uuidref":
						inst.Uuidref = GongExtractString(rhs)
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
					case "DS_Series":
						__gong__assignPointer(&inst.DS_Series, rhs, identifierMap)
					case "DS_Platform":
						__gong__assignPointer(&inst.DS_Platform, rhs, identifierMap)
					case "DS_Sensor":
						__gong__assignPointer(&inst.DS_Sensor, rhs, identifierMap)
					case "DS_ProductionSeries":
						__gong__assignPointer(&inst.DS_ProductionSeries, rhs, identifierMap)
					}
				case *DS_Series_Type:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Uuid":
						inst.Uuid = GongExtractString(rhs)
					case "ComposedOf":
						__gong__assignSliceOfPointers(&inst.ComposedOf, rhs, identifierMap)
					case "SeriesMetadata":
						__gong__assignSliceOfPointers(&inst.SeriesMetadata, rhs, identifierMap)
					case "Subset":
						__gong__assignSliceOfPointers(&inst.Subset, rhs, identifierMap)
					case "Superset":
						__gong__assignSliceOfPointers(&inst.Superset, rhs, identifierMap)
					}
				case *DS_StereoMate_PropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Uuidref":
						inst.Uuidref = GongExtractString(rhs)
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
					case "DS_StereoMate":
						__gong__assignPointer(&inst.DS_StereoMate, rhs, identifierMap)
					}
				case *DS_StereoMate_Type:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Uuid":
						inst.Uuid = GongExtractString(rhs)
					case "ComposedOf":
						__gong__assignSliceOfPointers(&inst.ComposedOf, rhs, identifierMap)
					case "SeriesMetadata":
						__gong__assignSliceOfPointers(&inst.SeriesMetadata, rhs, identifierMap)
					case "Subset":
						__gong__assignSliceOfPointers(&inst.Subset, rhs, identifierMap)
					case "Superset":
						__gong__assignSliceOfPointers(&inst.Superset, rhs, identifierMap)
					}
				case *MD_Metadata_PropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Uuidref":
						inst.Uuidref = GongExtractString(rhs)
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
					case "MD_Metadata":
						__gong__assignPointer(&inst.MD_Metadata, rhs, identifierMap)
					}
				case *MD_Metadata_Type:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Uuid":
						inst.Uuid = GongExtractString(rhs)
					case "Series":
						__gong__assignSliceOfPointers(&inst.Series, rhs, identifierMap)
					case "Describes":
						__gong__assignSliceOfPointers(&inst.Describes, rhs, identifierMap)
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
type AbstractDS_Aggregate_TypeUnmarshaller struct{}

func (u *AbstractDS_Aggregate_TypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractDS_Aggregate_Type), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractDS_Aggregate_TypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractDS_Aggregate_Type)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Uuid":
		instance.Uuid = GongExtractString(valueExpr)
	case "ComposedOf":
		GongUnmarshallSliceOfPointers(&instance.ComposedOf, valueExpr, identifierMap)
	case "SeriesMetadata":
		GongUnmarshallSliceOfPointers(&instance.SeriesMetadata, valueExpr, identifierMap)
	case "Subset":
		GongUnmarshallSliceOfPointers(&instance.Subset, valueExpr, identifierMap)
	case "Superset":
		GongUnmarshallSliceOfPointers(&instance.Superset, valueExpr, identifierMap)
	}
	return nil
}

type DS_Aggregate_PropertyTypeUnmarshaller struct{}

func (u *DS_Aggregate_PropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DS_Aggregate_PropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *DS_Aggregate_PropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DS_Aggregate_PropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Uuidref":
		instance.Uuidref = GongExtractString(valueExpr)
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
	case "DS_OtherAggregate":
		GongUnmarshallPointer(&instance.DS_OtherAggregate, valueExpr, identifierMap)
	case "DS_StereoMate":
		GongUnmarshallPointer(&instance.DS_StereoMate, valueExpr, identifierMap)
	case "DS_Series":
		GongUnmarshallPointer(&instance.DS_Series, valueExpr, identifierMap)
	case "DS_Platform":
		GongUnmarshallPointer(&instance.DS_Platform, valueExpr, identifierMap)
	case "DS_Sensor":
		GongUnmarshallPointer(&instance.DS_Sensor, valueExpr, identifierMap)
	case "DS_ProductionSeries":
		GongUnmarshallPointer(&instance.DS_ProductionSeries, valueExpr, identifierMap)
	case "DS_Initiative":
		GongUnmarshallPointer(&instance.DS_Initiative, valueExpr, identifierMap)
	}
	return nil
}

type DS_DataSet_PropertyTypeUnmarshaller struct{}

func (u *DS_DataSet_PropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DS_DataSet_PropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *DS_DataSet_PropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DS_DataSet_PropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Uuidref":
		instance.Uuidref = GongExtractString(valueExpr)
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
	case "DS_DataSet":
		GongUnmarshallPointer(&instance.DS_DataSet, valueExpr, identifierMap)
	}
	return nil
}

type DS_DataSet_TypeUnmarshaller struct{}

func (u *DS_DataSet_TypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DS_DataSet_Type), stage, identifier, instanceName, preserveOrder)
}

func (u *DS_DataSet_TypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DS_DataSet_Type)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Uuid":
		instance.Uuid = GongExtractString(valueExpr)
	case "Has":
		GongUnmarshallSliceOfPointers(&instance.Has, valueExpr, identifierMap)
	case "PartOf":
		GongUnmarshallSliceOfPointers(&instance.PartOf, valueExpr, identifierMap)
	}
	return nil
}

type DS_Initiative_PropertyTypeUnmarshaller struct{}

func (u *DS_Initiative_PropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DS_Initiative_PropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *DS_Initiative_PropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DS_Initiative_PropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Uuidref":
		instance.Uuidref = GongExtractString(valueExpr)
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
	case "DS_Initiative":
		GongUnmarshallPointer(&instance.DS_Initiative, valueExpr, identifierMap)
	}
	return nil
}

type DS_Initiative_TypeUnmarshaller struct{}

func (u *DS_Initiative_TypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DS_Initiative_Type), stage, identifier, instanceName, preserveOrder)
}

func (u *DS_Initiative_TypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DS_Initiative_Type)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Uuid":
		instance.Uuid = GongExtractString(valueExpr)
	case "ComposedOf":
		GongUnmarshallSliceOfPointers(&instance.ComposedOf, valueExpr, identifierMap)
	case "SeriesMetadata":
		GongUnmarshallSliceOfPointers(&instance.SeriesMetadata, valueExpr, identifierMap)
	case "Subset":
		GongUnmarshallSliceOfPointers(&instance.Subset, valueExpr, identifierMap)
	case "Superset":
		GongUnmarshallSliceOfPointers(&instance.Superset, valueExpr, identifierMap)
	}
	return nil
}

type DS_OtherAggregate_PropertyTypeUnmarshaller struct{}

func (u *DS_OtherAggregate_PropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DS_OtherAggregate_PropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *DS_OtherAggregate_PropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DS_OtherAggregate_PropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Uuidref":
		instance.Uuidref = GongExtractString(valueExpr)
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
	case "DS_OtherAggregate":
		GongUnmarshallPointer(&instance.DS_OtherAggregate, valueExpr, identifierMap)
	case "DS_StereoMate":
		GongUnmarshallPointer(&instance.DS_StereoMate, valueExpr, identifierMap)
	}
	return nil
}

type DS_OtherAggregate_TypeUnmarshaller struct{}

func (u *DS_OtherAggregate_TypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DS_OtherAggregate_Type), stage, identifier, instanceName, preserveOrder)
}

func (u *DS_OtherAggregate_TypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DS_OtherAggregate_Type)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Uuid":
		instance.Uuid = GongExtractString(valueExpr)
	case "ComposedOf":
		GongUnmarshallSliceOfPointers(&instance.ComposedOf, valueExpr, identifierMap)
	case "SeriesMetadata":
		GongUnmarshallSliceOfPointers(&instance.SeriesMetadata, valueExpr, identifierMap)
	case "Subset":
		GongUnmarshallSliceOfPointers(&instance.Subset, valueExpr, identifierMap)
	case "Superset":
		GongUnmarshallSliceOfPointers(&instance.Superset, valueExpr, identifierMap)
	}
	return nil
}

type DS_Platform_PropertyTypeUnmarshaller struct{}

func (u *DS_Platform_PropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DS_Platform_PropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *DS_Platform_PropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DS_Platform_PropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Uuidref":
		instance.Uuidref = GongExtractString(valueExpr)
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
	case "DS_Platform":
		GongUnmarshallPointer(&instance.DS_Platform, valueExpr, identifierMap)
	}
	return nil
}

type DS_Platform_TypeUnmarshaller struct{}

func (u *DS_Platform_TypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DS_Platform_Type), stage, identifier, instanceName, preserveOrder)
}

func (u *DS_Platform_TypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DS_Platform_Type)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Uuid":
		instance.Uuid = GongExtractString(valueExpr)
	case "ComposedOf":
		GongUnmarshallSliceOfPointers(&instance.ComposedOf, valueExpr, identifierMap)
	case "SeriesMetadata":
		GongUnmarshallSliceOfPointers(&instance.SeriesMetadata, valueExpr, identifierMap)
	case "Subset":
		GongUnmarshallSliceOfPointers(&instance.Subset, valueExpr, identifierMap)
	case "Superset":
		GongUnmarshallSliceOfPointers(&instance.Superset, valueExpr, identifierMap)
	}
	return nil
}

type DS_ProductionSeries_PropertyTypeUnmarshaller struct{}

func (u *DS_ProductionSeries_PropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DS_ProductionSeries_PropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *DS_ProductionSeries_PropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DS_ProductionSeries_PropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Uuidref":
		instance.Uuidref = GongExtractString(valueExpr)
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
	case "DS_ProductionSeries":
		GongUnmarshallPointer(&instance.DS_ProductionSeries, valueExpr, identifierMap)
	}
	return nil
}

type DS_ProductionSeries_TypeUnmarshaller struct{}

func (u *DS_ProductionSeries_TypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DS_ProductionSeries_Type), stage, identifier, instanceName, preserveOrder)
}

func (u *DS_ProductionSeries_TypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DS_ProductionSeries_Type)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Uuid":
		instance.Uuid = GongExtractString(valueExpr)
	case "ComposedOf":
		GongUnmarshallSliceOfPointers(&instance.ComposedOf, valueExpr, identifierMap)
	case "SeriesMetadata":
		GongUnmarshallSliceOfPointers(&instance.SeriesMetadata, valueExpr, identifierMap)
	case "Subset":
		GongUnmarshallSliceOfPointers(&instance.Subset, valueExpr, identifierMap)
	case "Superset":
		GongUnmarshallSliceOfPointers(&instance.Superset, valueExpr, identifierMap)
	}
	return nil
}

type DS_Sensor_PropertyTypeUnmarshaller struct{}

func (u *DS_Sensor_PropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DS_Sensor_PropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *DS_Sensor_PropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DS_Sensor_PropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Uuidref":
		instance.Uuidref = GongExtractString(valueExpr)
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
	case "DS_Sensor":
		GongUnmarshallPointer(&instance.DS_Sensor, valueExpr, identifierMap)
	}
	return nil
}

type DS_Sensor_TypeUnmarshaller struct{}

func (u *DS_Sensor_TypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DS_Sensor_Type), stage, identifier, instanceName, preserveOrder)
}

func (u *DS_Sensor_TypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DS_Sensor_Type)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Uuid":
		instance.Uuid = GongExtractString(valueExpr)
	case "ComposedOf":
		GongUnmarshallSliceOfPointers(&instance.ComposedOf, valueExpr, identifierMap)
	case "SeriesMetadata":
		GongUnmarshallSliceOfPointers(&instance.SeriesMetadata, valueExpr, identifierMap)
	case "Subset":
		GongUnmarshallSliceOfPointers(&instance.Subset, valueExpr, identifierMap)
	case "Superset":
		GongUnmarshallSliceOfPointers(&instance.Superset, valueExpr, identifierMap)
	}
	return nil
}

type DS_Series_PropertyTypeUnmarshaller struct{}

func (u *DS_Series_PropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DS_Series_PropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *DS_Series_PropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DS_Series_PropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Uuidref":
		instance.Uuidref = GongExtractString(valueExpr)
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
	case "DS_Series":
		GongUnmarshallPointer(&instance.DS_Series, valueExpr, identifierMap)
	case "DS_Platform":
		GongUnmarshallPointer(&instance.DS_Platform, valueExpr, identifierMap)
	case "DS_Sensor":
		GongUnmarshallPointer(&instance.DS_Sensor, valueExpr, identifierMap)
	case "DS_ProductionSeries":
		GongUnmarshallPointer(&instance.DS_ProductionSeries, valueExpr, identifierMap)
	}
	return nil
}

type DS_Series_TypeUnmarshaller struct{}

func (u *DS_Series_TypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DS_Series_Type), stage, identifier, instanceName, preserveOrder)
}

func (u *DS_Series_TypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DS_Series_Type)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Uuid":
		instance.Uuid = GongExtractString(valueExpr)
	case "ComposedOf":
		GongUnmarshallSliceOfPointers(&instance.ComposedOf, valueExpr, identifierMap)
	case "SeriesMetadata":
		GongUnmarshallSliceOfPointers(&instance.SeriesMetadata, valueExpr, identifierMap)
	case "Subset":
		GongUnmarshallSliceOfPointers(&instance.Subset, valueExpr, identifierMap)
	case "Superset":
		GongUnmarshallSliceOfPointers(&instance.Superset, valueExpr, identifierMap)
	}
	return nil
}

type DS_StereoMate_PropertyTypeUnmarshaller struct{}

func (u *DS_StereoMate_PropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DS_StereoMate_PropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *DS_StereoMate_PropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DS_StereoMate_PropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Uuidref":
		instance.Uuidref = GongExtractString(valueExpr)
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
	case "DS_StereoMate":
		GongUnmarshallPointer(&instance.DS_StereoMate, valueExpr, identifierMap)
	}
	return nil
}

type DS_StereoMate_TypeUnmarshaller struct{}

func (u *DS_StereoMate_TypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DS_StereoMate_Type), stage, identifier, instanceName, preserveOrder)
}

func (u *DS_StereoMate_TypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DS_StereoMate_Type)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Uuid":
		instance.Uuid = GongExtractString(valueExpr)
	case "ComposedOf":
		GongUnmarshallSliceOfPointers(&instance.ComposedOf, valueExpr, identifierMap)
	case "SeriesMetadata":
		GongUnmarshallSliceOfPointers(&instance.SeriesMetadata, valueExpr, identifierMap)
	case "Subset":
		GongUnmarshallSliceOfPointers(&instance.Subset, valueExpr, identifierMap)
	case "Superset":
		GongUnmarshallSliceOfPointers(&instance.Superset, valueExpr, identifierMap)
	}
	return nil
}

type MD_Metadata_PropertyTypeUnmarshaller struct{}

func (u *MD_Metadata_PropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MD_Metadata_PropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *MD_Metadata_PropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MD_Metadata_PropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Uuidref":
		instance.Uuidref = GongExtractString(valueExpr)
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
	case "MD_Metadata":
		GongUnmarshallPointer(&instance.MD_Metadata, valueExpr, identifierMap)
	}
	return nil
}

type MD_Metadata_TypeUnmarshaller struct{}

func (u *MD_Metadata_TypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MD_Metadata_Type), stage, identifier, instanceName, preserveOrder)
}

func (u *MD_Metadata_TypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MD_Metadata_Type)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Uuid":
		instance.Uuid = GongExtractString(valueExpr)
	case "Series":
		GongUnmarshallSliceOfPointers(&instance.Series, valueExpr, identifierMap)
	case "Describes":
		GongUnmarshallSliceOfPointers(&instance.Describes, valueExpr, identifierMap)
	}
	return nil
}
