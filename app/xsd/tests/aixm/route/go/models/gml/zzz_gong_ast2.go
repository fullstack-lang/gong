// generated code - do not edit
package gml

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
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/gml":
			aliasToCanonical[alias] = "gml"
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
					} else if pkgAlias == "" || pkgAlias == "gml" || len(inFile.Imports) == 0 {
						pkgAlias = "gml"
					}

					switch pkgAlias {
			case "gml":
				switch typeName {
				case "AbstractCurveSegmentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractCurveSegmentType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractCurveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractCurveType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractFeatureMemberType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractFeatureMemberType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractFeatureType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractFeatureType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractGMLType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractGMLType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractGeometricPrimitiveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractGeometricPrimitiveType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractGeometryType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractGeometryType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractMetadataPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractMetadataPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractRingPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractRingPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractRingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractRingType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractSurfacePatchType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractSurfacePatchType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractSurfaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractSurfaceType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractTimeGeometricPrimitiveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractTimeGeometricPrimitiveType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractTimeObjectType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractTimeObjectType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractTimePrimitiveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractTimePrimitiveType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "AngleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AngleType), stage, ident.Name, instanceName, preserveOrder)
				case "ArcByCenterPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ArcByCenterPointType), stage, ident.Name, instanceName, preserveOrder)
				case "ArcStringType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ArcStringType), stage, ident.Name, instanceName, preserveOrder)
				case "ArcType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ArcType), stage, ident.Name, instanceName, preserveOrder)
				case "BoundingShapeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(BoundingShapeType), stage, ident.Name, instanceName, preserveOrder)
				case "CircleByCenterPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CircleByCenterPointType), stage, ident.Name, instanceName, preserveOrder)
				case "CircleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CircleType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeWithAuthorityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeWithAuthorityType), stage, ident.Name, instanceName, preserveOrder)
				case "CompositeCurveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CompositeCurveType), stage, ident.Name, instanceName, preserveOrder)
				case "CurvePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CurvePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "CurveSegmentArrayPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CurveSegmentArrayPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "CurveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CurveType), stage, ident.Name, instanceName, preserveOrder)
				case "DirectPositionListType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DirectPositionListType), stage, ident.Name, instanceName, preserveOrder)
				case "DirectPositionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DirectPositionType), stage, ident.Name, instanceName, preserveOrder)
				case "DynamicFeatureType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DynamicFeatureType), stage, ident.Name, instanceName, preserveOrder)
				case "EnvelopeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(EnvelopeType), stage, ident.Name, instanceName, preserveOrder)
				case "GeodesicStringType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GeodesicStringType), stage, ident.Name, instanceName, preserveOrder)
				case "GeodesicType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GeodesicType), stage, ident.Name, instanceName, preserveOrder)
				case "LengthType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LengthType), stage, ident.Name, instanceName, preserveOrder)
				case "LineStringSegmentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LineStringSegmentType), stage, ident.Name, instanceName, preserveOrder)
				case "MeasureType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MeasureType), stage, ident.Name, instanceName, preserveOrder)
				case "OrientableCurveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(OrientableCurveType), stage, ident.Name, instanceName, preserveOrder)
				case "PointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PointPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "PointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PointType), stage, ident.Name, instanceName, preserveOrder)
				case "PolygonPatchType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PolygonPatchType), stage, ident.Name, instanceName, preserveOrder)
				case "RingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RingType), stage, ident.Name, instanceName, preserveOrder)
				case "SurfacePatchArrayPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SurfacePatchArrayPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SurfaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SurfaceType), stage, ident.Name, instanceName, preserveOrder)
				case "TimeInstantType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TimeInstantType), stage, ident.Name, instanceName, preserveOrder)
				case "TimePeriodType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TimePeriodType), stage, ident.Name, instanceName, preserveOrder)
				case "TimePositionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TimePositionType), stage, ident.Name, instanceName, preserveOrder)
				case "TimePrimitivePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TimePrimitivePropertyType), stage, ident.Name, instanceName, preserveOrder)
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
				case *AbstractCurveSegmentType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NumDerivativesAtStart":
						inst.NumDerivativesAtStart = GongExtractInt(rhs)
					case "NumDerivativesAtEnd":
						inst.NumDerivativesAtEnd = GongExtractInt(rhs)
					case "NumDerivativeInterior":
						inst.NumDerivativeInterior = GongExtractInt(rhs)
					}
				case *AbstractCurveType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					}
				case *AbstractFeatureMemberType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *AbstractFeatureType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "BoundedBy":
						__gong__assignPointer(&inst.BoundedBy, rhs, identifierMap)
					}
				case *AbstractGMLType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					}
				case *AbstractGeometricPrimitiveType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					}
				case *AbstractGeometryType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					}
				case *AbstractMetadataPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *AbstractRingPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Ring":
						__gong__assignPointer(&inst.Ring, rhs, identifierMap)
					}
				case *AbstractRingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *AbstractSurfacePatchType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *AbstractSurfaceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					}
				case *AbstractTimeGeometricPrimitiveType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "Frame":
						inst.Frame = GongExtractString(rhs)
					}
				case *AbstractTimeObjectType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					}
				case *AbstractTimePrimitiveType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					}
				case *AbstractTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "ValidTime":
						__gong__assignPointer(&inst.ValidTime, rhs, identifierMap)
					}
				case *AngleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *ArcByCenterPointType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NumDerivativesAtStart":
						inst.NumDerivativesAtStart = GongExtractInt(rhs)
					case "NumDerivativesAtEnd":
						inst.NumDerivativesAtEnd = GongExtractInt(rhs)
					case "NumDerivativeInterior":
						inst.NumDerivativeInterior = GongExtractInt(rhs)
					case "Interpolation":
						GongUnmarshallEnum(&inst.Interpolation, rhs)
					case "NumArc":
						inst.NumArc = GongExtractInt(rhs)
					case "Pos":
						__gong__assignPointer(&inst.Pos, rhs, identifierMap)
					case "PointProperty":
						__gong__assignPointer(&inst.PointProperty, rhs, identifierMap)
					case "PosList":
						__gong__assignPointer(&inst.PosList, rhs, identifierMap)
					case "Radius":
						__gong__assignPointer(&inst.Radius, rhs, identifierMap)
					case "StartAngle":
						__gong__assignPointer(&inst.StartAngle, rhs, identifierMap)
					case "EndAngle":
						__gong__assignPointer(&inst.EndAngle, rhs, identifierMap)
					}
				case *ArcStringType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NumDerivativesAtStart":
						inst.NumDerivativesAtStart = GongExtractInt(rhs)
					case "NumDerivativesAtEnd":
						inst.NumDerivativesAtEnd = GongExtractInt(rhs)
					case "NumDerivativeInterior":
						inst.NumDerivativeInterior = GongExtractInt(rhs)
					case "Interpolation":
						GongUnmarshallEnum(&inst.Interpolation, rhs)
					case "NumArc":
						inst.NumArc = GongExtractInt(rhs)
					case "Pos":
						__gong__assignSliceOfPointers(&inst.Pos, rhs, identifierMap)
					case "PointProperty":
						__gong__assignSliceOfPointers(&inst.PointProperty, rhs, identifierMap)
					case "PosList":
						__gong__assignPointer(&inst.PosList, rhs, identifierMap)
					}
				case *ArcType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NumDerivativesAtStart":
						inst.NumDerivativesAtStart = GongExtractInt(rhs)
					case "NumDerivativesAtEnd":
						inst.NumDerivativesAtEnd = GongExtractInt(rhs)
					case "NumDerivativeInterior":
						inst.NumDerivativeInterior = GongExtractInt(rhs)
					case "Interpolation":
						GongUnmarshallEnum(&inst.Interpolation, rhs)
					case "NumArc":
						inst.NumArc = GongExtractInt(rhs)
					case "Pos":
						__gong__assignPointer(&inst.Pos, rhs, identifierMap)
					case "PointProperty":
						__gong__assignPointer(&inst.PointProperty, rhs, identifierMap)
					case "PosList":
						__gong__assignPointer(&inst.PosList, rhs, identifierMap)
					}
				case *BoundingShapeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Envelope":
						__gong__assignPointer(&inst.Envelope, rhs, identifierMap)
					}
				case *CircleByCenterPointType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NumDerivativesAtStart":
						inst.NumDerivativesAtStart = GongExtractInt(rhs)
					case "NumDerivativesAtEnd":
						inst.NumDerivativesAtEnd = GongExtractInt(rhs)
					case "NumDerivativeInterior":
						inst.NumDerivativeInterior = GongExtractInt(rhs)
					case "Interpolation":
						GongUnmarshallEnum(&inst.Interpolation, rhs)
					case "NumArc":
						inst.NumArc = GongExtractInt(rhs)
					case "Pos":
						__gong__assignPointer(&inst.Pos, rhs, identifierMap)
					case "PointProperty":
						__gong__assignPointer(&inst.PointProperty, rhs, identifierMap)
					case "PosList":
						__gong__assignPointer(&inst.PosList, rhs, identifierMap)
					case "Radius":
						__gong__assignPointer(&inst.Radius, rhs, identifierMap)
					}
				case *CircleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NumDerivativesAtStart":
						inst.NumDerivativesAtStart = GongExtractInt(rhs)
					case "NumDerivativesAtEnd":
						inst.NumDerivativesAtEnd = GongExtractInt(rhs)
					case "NumDerivativeInterior":
						inst.NumDerivativeInterior = GongExtractInt(rhs)
					case "Interpolation":
						GongUnmarshallEnum(&inst.Interpolation, rhs)
					case "NumArc":
						inst.NumArc = GongExtractInt(rhs)
					case "Pos":
						__gong__assignSliceOfPointers(&inst.Pos, rhs, identifierMap)
					case "PointProperty":
						__gong__assignSliceOfPointers(&inst.PointProperty, rhs, identifierMap)
					case "PosList":
						__gong__assignPointer(&inst.PosList, rhs, identifierMap)
					}
				case *CodeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "CodeSpace":
						inst.CodeSpace = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeWithAuthorityType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "CodeSpace":
						inst.CodeSpace = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CompositeCurveType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "AggregationType":
						inst.AggregationType = GongExtractString(rhs)
					case "CurveMember":
						__gong__assignSliceOfPointers(&inst.CurveMember, rhs, identifierMap)
					}
				case *CurvePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
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
					case "Curve":
						__gong__assignPointer(&inst.Curve, rhs, identifierMap)
					case "OrientableCurve":
						__gong__assignPointer(&inst.OrientableCurve, rhs, identifierMap)
					case "CompositeCurve":
						__gong__assignPointer(&inst.CompositeCurve, rhs, identifierMap)
					}
				case *CurveSegmentArrayPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "LineStringSegment":
						__gong__assignSliceOfPointers(&inst.LineStringSegment, rhs, identifierMap)
					case "ArcString":
						__gong__assignSliceOfPointers(&inst.ArcString, rhs, identifierMap)
					case "Arc":
						__gong__assignSliceOfPointers(&inst.Arc, rhs, identifierMap)
					case "Circle":
						__gong__assignSliceOfPointers(&inst.Circle, rhs, identifierMap)
					case "ArcByCenterPoint":
						__gong__assignSliceOfPointers(&inst.ArcByCenterPoint, rhs, identifierMap)
					case "CircleByCenterPoint":
						__gong__assignSliceOfPointers(&inst.CircleByCenterPoint, rhs, identifierMap)
					case "GeodesicString":
						__gong__assignSliceOfPointers(&inst.GeodesicString, rhs, identifierMap)
					case "Geodesic":
						__gong__assignSliceOfPointers(&inst.Geodesic, rhs, identifierMap)
					}
				case *CurveType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "Segments":
						__gong__assignPointer(&inst.Segments, rhs, identifierMap)
					}
				case *DirectPositionListType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *DirectPositionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *DynamicFeatureType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "BoundedBy":
						__gong__assignPointer(&inst.BoundedBy, rhs, identifierMap)
					}
				case *EnvelopeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "LowerCorner":
						__gong__assignPointer(&inst.LowerCorner, rhs, identifierMap)
					case "UpperCorner":
						__gong__assignPointer(&inst.UpperCorner, rhs, identifierMap)
					}
				case *GeodesicStringType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NumDerivativesAtStart":
						inst.NumDerivativesAtStart = GongExtractInt(rhs)
					case "NumDerivativesAtEnd":
						inst.NumDerivativesAtEnd = GongExtractInt(rhs)
					case "NumDerivativeInterior":
						inst.NumDerivativeInterior = GongExtractInt(rhs)
					case "Interpolation":
						GongUnmarshallEnum(&inst.Interpolation, rhs)
					case "PosList":
						__gong__assignPointer(&inst.PosList, rhs, identifierMap)
					case "Pos":
						__gong__assignPointer(&inst.Pos, rhs, identifierMap)
					case "PointProperty":
						__gong__assignPointer(&inst.PointProperty, rhs, identifierMap)
					}
				case *GeodesicType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NumDerivativesAtStart":
						inst.NumDerivativesAtStart = GongExtractInt(rhs)
					case "NumDerivativesAtEnd":
						inst.NumDerivativesAtEnd = GongExtractInt(rhs)
					case "NumDerivativeInterior":
						inst.NumDerivativeInterior = GongExtractInt(rhs)
					case "Interpolation":
						GongUnmarshallEnum(&inst.Interpolation, rhs)
					case "PosList":
						__gong__assignPointer(&inst.PosList, rhs, identifierMap)
					case "Pos":
						__gong__assignPointer(&inst.Pos, rhs, identifierMap)
					case "PointProperty":
						__gong__assignPointer(&inst.PointProperty, rhs, identifierMap)
					}
				case *LengthType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *LineStringSegmentType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NumDerivativesAtStart":
						inst.NumDerivativesAtStart = GongExtractInt(rhs)
					case "NumDerivativesAtEnd":
						inst.NumDerivativesAtEnd = GongExtractInt(rhs)
					case "NumDerivativeInterior":
						inst.NumDerivativeInterior = GongExtractInt(rhs)
					case "Interpolation":
						GongUnmarshallEnum(&inst.Interpolation, rhs)
					case "Pos":
						__gong__assignSliceOfPointers(&inst.Pos, rhs, identifierMap)
					case "PointProperty":
						__gong__assignSliceOfPointers(&inst.PointProperty, rhs, identifierMap)
					case "PosList":
						__gong__assignPointer(&inst.PosList, rhs, identifierMap)
					}
				case *MeasureType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Uom":
						inst.Uom = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractFloat(rhs)
					}
				case *OrientableCurveType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "Orientation":
						GongUnmarshallEnum(&inst.Orientation, rhs)
					case "BaseCurve":
						__gong__assignPointer(&inst.BaseCurve, rhs, identifierMap)
					}
				case *PointPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
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
					case "Point":
						__gong__assignPointer(&inst.Point, rhs, identifierMap)
					}
				case *PointType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "Pos":
						__gong__assignPointer(&inst.Pos, rhs, identifierMap)
					}
				case *PolygonPatchType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Interpolation":
						GongUnmarshallEnum(&inst.Interpolation, rhs)
					case "Exterior":
						__gong__assignPointer(&inst.Exterior, rhs, identifierMap)
					}
				case *RingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "AggregationType":
						inst.AggregationType = GongExtractString(rhs)
					case "CurveMember":
						__gong__assignSliceOfPointers(&inst.CurveMember, rhs, identifierMap)
					}
				case *SurfacePatchArrayPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "PolygonPatch":
						__gong__assignSliceOfPointers(&inst.PolygonPatch, rhs, identifierMap)
					}
				case *SurfaceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "Patches":
						__gong__assignPointer(&inst.Patches, rhs, identifierMap)
					}
				case *TimeInstantType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "Frame":
						inst.Frame = GongExtractString(rhs)
					case "TimePosition":
						__gong__assignPointer(&inst.TimePosition, rhs, identifierMap)
					}
				case *TimePeriodType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "Frame":
						inst.Frame = GongExtractString(rhs)
					case "BeginPosition":
						__gong__assignPointer(&inst.BeginPosition, rhs, identifierMap)
					case "EndPosition":
						__gong__assignPointer(&inst.EndPosition, rhs, identifierMap)
					}
				case *TimePositionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Frame":
						inst.Frame = GongExtractString(rhs)
					case "IndeterminatePosition":
						GongUnmarshallEnum(&inst.IndeterminatePosition, rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *TimePrimitivePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
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
					case "TimeInstant":
						__gong__assignPointer(&inst.TimeInstant, rhs, identifierMap)
					case "TimePeriod":
						__gong__assignPointer(&inst.TimePeriod, rhs, identifierMap)
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
type AbstractCurveSegmentTypeUnmarshaller struct{}

func (u *AbstractCurveSegmentTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractCurveSegmentType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractCurveSegmentTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractCurveSegmentType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NumDerivativesAtStart":
		instance.NumDerivativesAtStart = GongExtractInt(valueExpr)
	case "NumDerivativesAtEnd":
		instance.NumDerivativesAtEnd = GongExtractInt(valueExpr)
	case "NumDerivativeInterior":
		instance.NumDerivativeInterior = GongExtractInt(valueExpr)
	}
	return nil
}

type AbstractCurveTypeUnmarshaller struct{}

func (u *AbstractCurveTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractCurveType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractCurveTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractCurveType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Identifier":
		GongUnmarshallPointer(&instance.Identifier, valueExpr, identifierMap)
	case "NameXSD":
		GongUnmarshallSliceOfPointers(&instance.NameXSD, valueExpr, identifierMap)
	case "SrsName":
		instance.SrsName = GongExtractString(valueExpr)
	}
	return nil
}

type AbstractFeatureMemberTypeUnmarshaller struct{}

func (u *AbstractFeatureMemberTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractFeatureMemberType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractFeatureMemberTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractFeatureMemberType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type AbstractFeatureTypeUnmarshaller struct{}

func (u *AbstractFeatureTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractFeatureType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractFeatureTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractFeatureType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Identifier":
		GongUnmarshallPointer(&instance.Identifier, valueExpr, identifierMap)
	case "NameXSD":
		GongUnmarshallSliceOfPointers(&instance.NameXSD, valueExpr, identifierMap)
	case "BoundedBy":
		GongUnmarshallPointer(&instance.BoundedBy, valueExpr, identifierMap)
	}
	return nil
}

type AbstractGMLTypeUnmarshaller struct{}

func (u *AbstractGMLTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractGMLType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractGMLTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractGMLType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Identifier":
		GongUnmarshallPointer(&instance.Identifier, valueExpr, identifierMap)
	case "NameXSD":
		GongUnmarshallSliceOfPointers(&instance.NameXSD, valueExpr, identifierMap)
	}
	return nil
}

type AbstractGeometricPrimitiveTypeUnmarshaller struct{}

func (u *AbstractGeometricPrimitiveTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractGeometricPrimitiveType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractGeometricPrimitiveTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractGeometricPrimitiveType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Identifier":
		GongUnmarshallPointer(&instance.Identifier, valueExpr, identifierMap)
	case "NameXSD":
		GongUnmarshallSliceOfPointers(&instance.NameXSD, valueExpr, identifierMap)
	case "SrsName":
		instance.SrsName = GongExtractString(valueExpr)
	}
	return nil
}

type AbstractGeometryTypeUnmarshaller struct{}

func (u *AbstractGeometryTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractGeometryType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractGeometryTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractGeometryType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Identifier":
		GongUnmarshallPointer(&instance.Identifier, valueExpr, identifierMap)
	case "NameXSD":
		GongUnmarshallSliceOfPointers(&instance.NameXSD, valueExpr, identifierMap)
	case "SrsName":
		instance.SrsName = GongExtractString(valueExpr)
	}
	return nil
}

type AbstractMetadataPropertyTypeUnmarshaller struct{}

func (u *AbstractMetadataPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractMetadataPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractMetadataPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractMetadataPropertyType)
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

type AbstractRingPropertyTypeUnmarshaller struct{}

func (u *AbstractRingPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractRingPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractRingPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractRingPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Ring":
		GongUnmarshallPointer(&instance.Ring, valueExpr, identifierMap)
	}
	return nil
}

type AbstractRingTypeUnmarshaller struct{}

func (u *AbstractRingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractRingType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractRingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractRingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type AbstractSurfacePatchTypeUnmarshaller struct{}

func (u *AbstractSurfacePatchTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractSurfacePatchType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractSurfacePatchTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractSurfacePatchType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type AbstractSurfaceTypeUnmarshaller struct{}

func (u *AbstractSurfaceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractSurfaceType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractSurfaceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractSurfaceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Identifier":
		GongUnmarshallPointer(&instance.Identifier, valueExpr, identifierMap)
	case "NameXSD":
		GongUnmarshallSliceOfPointers(&instance.NameXSD, valueExpr, identifierMap)
	case "SrsName":
		instance.SrsName = GongExtractString(valueExpr)
	}
	return nil
}

type AbstractTimeGeometricPrimitiveTypeUnmarshaller struct{}

func (u *AbstractTimeGeometricPrimitiveTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractTimeGeometricPrimitiveType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractTimeGeometricPrimitiveTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractTimeGeometricPrimitiveType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Identifier":
		GongUnmarshallPointer(&instance.Identifier, valueExpr, identifierMap)
	case "NameXSD":
		GongUnmarshallSliceOfPointers(&instance.NameXSD, valueExpr, identifierMap)
	case "Frame":
		instance.Frame = GongExtractString(valueExpr)
	}
	return nil
}

type AbstractTimeObjectTypeUnmarshaller struct{}

func (u *AbstractTimeObjectTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractTimeObjectType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractTimeObjectTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractTimeObjectType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Identifier":
		GongUnmarshallPointer(&instance.Identifier, valueExpr, identifierMap)
	case "NameXSD":
		GongUnmarshallSliceOfPointers(&instance.NameXSD, valueExpr, identifierMap)
	}
	return nil
}

type AbstractTimePrimitiveTypeUnmarshaller struct{}

func (u *AbstractTimePrimitiveTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractTimePrimitiveType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractTimePrimitiveTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractTimePrimitiveType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Identifier":
		GongUnmarshallPointer(&instance.Identifier, valueExpr, identifierMap)
	case "NameXSD":
		GongUnmarshallSliceOfPointers(&instance.NameXSD, valueExpr, identifierMap)
	}
	return nil
}

type AbstractTimeSliceTypeUnmarshaller struct{}

func (u *AbstractTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Identifier":
		GongUnmarshallPointer(&instance.Identifier, valueExpr, identifierMap)
	case "NameXSD":
		GongUnmarshallSliceOfPointers(&instance.NameXSD, valueExpr, identifierMap)
	case "ValidTime":
		GongUnmarshallPointer(&instance.ValidTime, valueExpr, identifierMap)
	}
	return nil
}

type AngleTypeUnmarshaller struct{}

func (u *AngleTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AngleType), stage, identifier, instanceName, preserveOrder)
}

func (u *AngleTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AngleType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ArcByCenterPointTypeUnmarshaller struct{}

func (u *ArcByCenterPointTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ArcByCenterPointType), stage, identifier, instanceName, preserveOrder)
}

func (u *ArcByCenterPointTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ArcByCenterPointType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NumDerivativesAtStart":
		instance.NumDerivativesAtStart = GongExtractInt(valueExpr)
	case "NumDerivativesAtEnd":
		instance.NumDerivativesAtEnd = GongExtractInt(valueExpr)
	case "NumDerivativeInterior":
		instance.NumDerivativeInterior = GongExtractInt(valueExpr)
	case "Interpolation":
		GongUnmarshallEnum(&instance.Interpolation, valueExpr)
	case "NumArc":
		instance.NumArc = GongExtractInt(valueExpr)
	case "Pos":
		GongUnmarshallPointer(&instance.Pos, valueExpr, identifierMap)
	case "PointProperty":
		GongUnmarshallPointer(&instance.PointProperty, valueExpr, identifierMap)
	case "PosList":
		GongUnmarshallPointer(&instance.PosList, valueExpr, identifierMap)
	case "Radius":
		GongUnmarshallPointer(&instance.Radius, valueExpr, identifierMap)
	case "StartAngle":
		GongUnmarshallPointer(&instance.StartAngle, valueExpr, identifierMap)
	case "EndAngle":
		GongUnmarshallPointer(&instance.EndAngle, valueExpr, identifierMap)
	}
	return nil
}

type ArcStringTypeUnmarshaller struct{}

func (u *ArcStringTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ArcStringType), stage, identifier, instanceName, preserveOrder)
}

func (u *ArcStringTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ArcStringType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NumDerivativesAtStart":
		instance.NumDerivativesAtStart = GongExtractInt(valueExpr)
	case "NumDerivativesAtEnd":
		instance.NumDerivativesAtEnd = GongExtractInt(valueExpr)
	case "NumDerivativeInterior":
		instance.NumDerivativeInterior = GongExtractInt(valueExpr)
	case "Interpolation":
		GongUnmarshallEnum(&instance.Interpolation, valueExpr)
	case "NumArc":
		instance.NumArc = GongExtractInt(valueExpr)
	case "Pos":
		GongUnmarshallSliceOfPointers(&instance.Pos, valueExpr, identifierMap)
	case "PointProperty":
		GongUnmarshallSliceOfPointers(&instance.PointProperty, valueExpr, identifierMap)
	case "PosList":
		GongUnmarshallPointer(&instance.PosList, valueExpr, identifierMap)
	}
	return nil
}

type ArcTypeUnmarshaller struct{}

func (u *ArcTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ArcType), stage, identifier, instanceName, preserveOrder)
}

func (u *ArcTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ArcType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NumDerivativesAtStart":
		instance.NumDerivativesAtStart = GongExtractInt(valueExpr)
	case "NumDerivativesAtEnd":
		instance.NumDerivativesAtEnd = GongExtractInt(valueExpr)
	case "NumDerivativeInterior":
		instance.NumDerivativeInterior = GongExtractInt(valueExpr)
	case "Interpolation":
		GongUnmarshallEnum(&instance.Interpolation, valueExpr)
	case "NumArc":
		instance.NumArc = GongExtractInt(valueExpr)
	case "Pos":
		GongUnmarshallPointer(&instance.Pos, valueExpr, identifierMap)
	case "PointProperty":
		GongUnmarshallPointer(&instance.PointProperty, valueExpr, identifierMap)
	case "PosList":
		GongUnmarshallPointer(&instance.PosList, valueExpr, identifierMap)
	}
	return nil
}

type BoundingShapeTypeUnmarshaller struct{}

func (u *BoundingShapeTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(BoundingShapeType), stage, identifier, instanceName, preserveOrder)
}

func (u *BoundingShapeTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*BoundingShapeType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Envelope":
		GongUnmarshallPointer(&instance.Envelope, valueExpr, identifierMap)
	}
	return nil
}

type CircleByCenterPointTypeUnmarshaller struct{}

func (u *CircleByCenterPointTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CircleByCenterPointType), stage, identifier, instanceName, preserveOrder)
}

func (u *CircleByCenterPointTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CircleByCenterPointType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NumDerivativesAtStart":
		instance.NumDerivativesAtStart = GongExtractInt(valueExpr)
	case "NumDerivativesAtEnd":
		instance.NumDerivativesAtEnd = GongExtractInt(valueExpr)
	case "NumDerivativeInterior":
		instance.NumDerivativeInterior = GongExtractInt(valueExpr)
	case "Interpolation":
		GongUnmarshallEnum(&instance.Interpolation, valueExpr)
	case "NumArc":
		instance.NumArc = GongExtractInt(valueExpr)
	case "Pos":
		GongUnmarshallPointer(&instance.Pos, valueExpr, identifierMap)
	case "PointProperty":
		GongUnmarshallPointer(&instance.PointProperty, valueExpr, identifierMap)
	case "PosList":
		GongUnmarshallPointer(&instance.PosList, valueExpr, identifierMap)
	case "Radius":
		GongUnmarshallPointer(&instance.Radius, valueExpr, identifierMap)
	}
	return nil
}

type CircleTypeUnmarshaller struct{}

func (u *CircleTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CircleType), stage, identifier, instanceName, preserveOrder)
}

func (u *CircleTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CircleType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NumDerivativesAtStart":
		instance.NumDerivativesAtStart = GongExtractInt(valueExpr)
	case "NumDerivativesAtEnd":
		instance.NumDerivativesAtEnd = GongExtractInt(valueExpr)
	case "NumDerivativeInterior":
		instance.NumDerivativeInterior = GongExtractInt(valueExpr)
	case "Interpolation":
		GongUnmarshallEnum(&instance.Interpolation, valueExpr)
	case "NumArc":
		instance.NumArc = GongExtractInt(valueExpr)
	case "Pos":
		GongUnmarshallSliceOfPointers(&instance.Pos, valueExpr, identifierMap)
	case "PointProperty":
		GongUnmarshallSliceOfPointers(&instance.PointProperty, valueExpr, identifierMap)
	case "PosList":
		GongUnmarshallPointer(&instance.PosList, valueExpr, identifierMap)
	}
	return nil
}

type CodeTypeUnmarshaller struct{}

func (u *CodeTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "CodeSpace":
		instance.CodeSpace = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeWithAuthorityTypeUnmarshaller struct{}

func (u *CodeWithAuthorityTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeWithAuthorityType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeWithAuthorityTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeWithAuthorityType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "CodeSpace":
		instance.CodeSpace = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CompositeCurveTypeUnmarshaller struct{}

func (u *CompositeCurveTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CompositeCurveType), stage, identifier, instanceName, preserveOrder)
}

func (u *CompositeCurveTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CompositeCurveType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Identifier":
		GongUnmarshallPointer(&instance.Identifier, valueExpr, identifierMap)
	case "NameXSD":
		GongUnmarshallSliceOfPointers(&instance.NameXSD, valueExpr, identifierMap)
	case "SrsName":
		instance.SrsName = GongExtractString(valueExpr)
	case "AggregationType":
		instance.AggregationType = GongExtractString(valueExpr)
	case "CurveMember":
		GongUnmarshallSliceOfPointers(&instance.CurveMember, valueExpr, identifierMap)
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
	case "Curve":
		GongUnmarshallPointer(&instance.Curve, valueExpr, identifierMap)
	case "OrientableCurve":
		GongUnmarshallPointer(&instance.OrientableCurve, valueExpr, identifierMap)
	case "CompositeCurve":
		GongUnmarshallPointer(&instance.CompositeCurve, valueExpr, identifierMap)
	}
	return nil
}

type CurveSegmentArrayPropertyTypeUnmarshaller struct{}

func (u *CurveSegmentArrayPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CurveSegmentArrayPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *CurveSegmentArrayPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CurveSegmentArrayPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "LineStringSegment":
		GongUnmarshallSliceOfPointers(&instance.LineStringSegment, valueExpr, identifierMap)
	case "ArcString":
		GongUnmarshallSliceOfPointers(&instance.ArcString, valueExpr, identifierMap)
	case "Arc":
		GongUnmarshallSliceOfPointers(&instance.Arc, valueExpr, identifierMap)
	case "Circle":
		GongUnmarshallSliceOfPointers(&instance.Circle, valueExpr, identifierMap)
	case "ArcByCenterPoint":
		GongUnmarshallSliceOfPointers(&instance.ArcByCenterPoint, valueExpr, identifierMap)
	case "CircleByCenterPoint":
		GongUnmarshallSliceOfPointers(&instance.CircleByCenterPoint, valueExpr, identifierMap)
	case "GeodesicString":
		GongUnmarshallSliceOfPointers(&instance.GeodesicString, valueExpr, identifierMap)
	case "Geodesic":
		GongUnmarshallSliceOfPointers(&instance.Geodesic, valueExpr, identifierMap)
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
	case "Identifier":
		GongUnmarshallPointer(&instance.Identifier, valueExpr, identifierMap)
	case "NameXSD":
		GongUnmarshallSliceOfPointers(&instance.NameXSD, valueExpr, identifierMap)
	case "SrsName":
		instance.SrsName = GongExtractString(valueExpr)
	case "Segments":
		GongUnmarshallPointer(&instance.Segments, valueExpr, identifierMap)
	}
	return nil
}

type DirectPositionListTypeUnmarshaller struct{}

func (u *DirectPositionListTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DirectPositionListType), stage, identifier, instanceName, preserveOrder)
}

func (u *DirectPositionListTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DirectPositionListType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SrsName":
		instance.SrsName = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type DirectPositionTypeUnmarshaller struct{}

func (u *DirectPositionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DirectPositionType), stage, identifier, instanceName, preserveOrder)
}

func (u *DirectPositionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DirectPositionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SrsName":
		instance.SrsName = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type DynamicFeatureTypeUnmarshaller struct{}

func (u *DynamicFeatureTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DynamicFeatureType), stage, identifier, instanceName, preserveOrder)
}

func (u *DynamicFeatureTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DynamicFeatureType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Identifier":
		GongUnmarshallPointer(&instance.Identifier, valueExpr, identifierMap)
	case "NameXSD":
		GongUnmarshallSliceOfPointers(&instance.NameXSD, valueExpr, identifierMap)
	case "BoundedBy":
		GongUnmarshallPointer(&instance.BoundedBy, valueExpr, identifierMap)
	}
	return nil
}

type EnvelopeTypeUnmarshaller struct{}

func (u *EnvelopeTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EnvelopeType), stage, identifier, instanceName, preserveOrder)
}

func (u *EnvelopeTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EnvelopeType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SrsName":
		instance.SrsName = GongExtractString(valueExpr)
	case "LowerCorner":
		GongUnmarshallPointer(&instance.LowerCorner, valueExpr, identifierMap)
	case "UpperCorner":
		GongUnmarshallPointer(&instance.UpperCorner, valueExpr, identifierMap)
	}
	return nil
}

type GeodesicStringTypeUnmarshaller struct{}

func (u *GeodesicStringTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GeodesicStringType), stage, identifier, instanceName, preserveOrder)
}

func (u *GeodesicStringTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GeodesicStringType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NumDerivativesAtStart":
		instance.NumDerivativesAtStart = GongExtractInt(valueExpr)
	case "NumDerivativesAtEnd":
		instance.NumDerivativesAtEnd = GongExtractInt(valueExpr)
	case "NumDerivativeInterior":
		instance.NumDerivativeInterior = GongExtractInt(valueExpr)
	case "Interpolation":
		GongUnmarshallEnum(&instance.Interpolation, valueExpr)
	case "PosList":
		GongUnmarshallPointer(&instance.PosList, valueExpr, identifierMap)
	case "Pos":
		GongUnmarshallPointer(&instance.Pos, valueExpr, identifierMap)
	case "PointProperty":
		GongUnmarshallPointer(&instance.PointProperty, valueExpr, identifierMap)
	}
	return nil
}

type GeodesicTypeUnmarshaller struct{}

func (u *GeodesicTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GeodesicType), stage, identifier, instanceName, preserveOrder)
}

func (u *GeodesicTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GeodesicType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NumDerivativesAtStart":
		instance.NumDerivativesAtStart = GongExtractInt(valueExpr)
	case "NumDerivativesAtEnd":
		instance.NumDerivativesAtEnd = GongExtractInt(valueExpr)
	case "NumDerivativeInterior":
		instance.NumDerivativeInterior = GongExtractInt(valueExpr)
	case "Interpolation":
		GongUnmarshallEnum(&instance.Interpolation, valueExpr)
	case "PosList":
		GongUnmarshallPointer(&instance.PosList, valueExpr, identifierMap)
	case "Pos":
		GongUnmarshallPointer(&instance.Pos, valueExpr, identifierMap)
	case "PointProperty":
		GongUnmarshallPointer(&instance.PointProperty, valueExpr, identifierMap)
	}
	return nil
}

type LengthTypeUnmarshaller struct{}

func (u *LengthTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LengthType), stage, identifier, instanceName, preserveOrder)
}

func (u *LengthTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LengthType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type LineStringSegmentTypeUnmarshaller struct{}

func (u *LineStringSegmentTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LineStringSegmentType), stage, identifier, instanceName, preserveOrder)
}

func (u *LineStringSegmentTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LineStringSegmentType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NumDerivativesAtStart":
		instance.NumDerivativesAtStart = GongExtractInt(valueExpr)
	case "NumDerivativesAtEnd":
		instance.NumDerivativesAtEnd = GongExtractInt(valueExpr)
	case "NumDerivativeInterior":
		instance.NumDerivativeInterior = GongExtractInt(valueExpr)
	case "Interpolation":
		GongUnmarshallEnum(&instance.Interpolation, valueExpr)
	case "Pos":
		GongUnmarshallSliceOfPointers(&instance.Pos, valueExpr, identifierMap)
	case "PointProperty":
		GongUnmarshallSliceOfPointers(&instance.PointProperty, valueExpr, identifierMap)
	case "PosList":
		GongUnmarshallPointer(&instance.PosList, valueExpr, identifierMap)
	}
	return nil
}

type MeasureTypeUnmarshaller struct{}

func (u *MeasureTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MeasureType), stage, identifier, instanceName, preserveOrder)
}

func (u *MeasureTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MeasureType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Uom":
		instance.Uom = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractFloat(valueExpr)
	}
	return nil
}

type OrientableCurveTypeUnmarshaller struct{}

func (u *OrientableCurveTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(OrientableCurveType), stage, identifier, instanceName, preserveOrder)
}

func (u *OrientableCurveTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*OrientableCurveType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Identifier":
		GongUnmarshallPointer(&instance.Identifier, valueExpr, identifierMap)
	case "NameXSD":
		GongUnmarshallSliceOfPointers(&instance.NameXSD, valueExpr, identifierMap)
	case "SrsName":
		instance.SrsName = GongExtractString(valueExpr)
	case "Orientation":
		GongUnmarshallEnum(&instance.Orientation, valueExpr)
	case "BaseCurve":
		GongUnmarshallPointer(&instance.BaseCurve, valueExpr, identifierMap)
	}
	return nil
}

type PointPropertyTypeUnmarshaller struct{}

func (u *PointPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PointPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *PointPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PointPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
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
	case "Point":
		GongUnmarshallPointer(&instance.Point, valueExpr, identifierMap)
	}
	return nil
}

type PointTypeUnmarshaller struct{}

func (u *PointTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PointType), stage, identifier, instanceName, preserveOrder)
}

func (u *PointTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PointType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Identifier":
		GongUnmarshallPointer(&instance.Identifier, valueExpr, identifierMap)
	case "NameXSD":
		GongUnmarshallSliceOfPointers(&instance.NameXSD, valueExpr, identifierMap)
	case "SrsName":
		instance.SrsName = GongExtractString(valueExpr)
	case "Pos":
		GongUnmarshallPointer(&instance.Pos, valueExpr, identifierMap)
	}
	return nil
}

type PolygonPatchTypeUnmarshaller struct{}

func (u *PolygonPatchTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PolygonPatchType), stage, identifier, instanceName, preserveOrder)
}

func (u *PolygonPatchTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PolygonPatchType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Interpolation":
		GongUnmarshallEnum(&instance.Interpolation, valueExpr)
	case "Exterior":
		GongUnmarshallPointer(&instance.Exterior, valueExpr, identifierMap)
	}
	return nil
}

type RingTypeUnmarshaller struct{}

func (u *RingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RingType), stage, identifier, instanceName, preserveOrder)
}

func (u *RingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "AggregationType":
		instance.AggregationType = GongExtractString(valueExpr)
	case "CurveMember":
		GongUnmarshallSliceOfPointers(&instance.CurveMember, valueExpr, identifierMap)
	}
	return nil
}

type SurfacePatchArrayPropertyTypeUnmarshaller struct{}

func (u *SurfacePatchArrayPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SurfacePatchArrayPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *SurfacePatchArrayPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SurfacePatchArrayPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "PolygonPatch":
		GongUnmarshallSliceOfPointers(&instance.PolygonPatch, valueExpr, identifierMap)
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
	case "Identifier":
		GongUnmarshallPointer(&instance.Identifier, valueExpr, identifierMap)
	case "NameXSD":
		GongUnmarshallSliceOfPointers(&instance.NameXSD, valueExpr, identifierMap)
	case "SrsName":
		instance.SrsName = GongExtractString(valueExpr)
	case "Patches":
		GongUnmarshallPointer(&instance.Patches, valueExpr, identifierMap)
	}
	return nil
}

type TimeInstantTypeUnmarshaller struct{}

func (u *TimeInstantTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TimeInstantType), stage, identifier, instanceName, preserveOrder)
}

func (u *TimeInstantTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TimeInstantType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Identifier":
		GongUnmarshallPointer(&instance.Identifier, valueExpr, identifierMap)
	case "NameXSD":
		GongUnmarshallSliceOfPointers(&instance.NameXSD, valueExpr, identifierMap)
	case "Frame":
		instance.Frame = GongExtractString(valueExpr)
	case "TimePosition":
		GongUnmarshallPointer(&instance.TimePosition, valueExpr, identifierMap)
	}
	return nil
}

type TimePeriodTypeUnmarshaller struct{}

func (u *TimePeriodTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TimePeriodType), stage, identifier, instanceName, preserveOrder)
}

func (u *TimePeriodTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TimePeriodType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Identifier":
		GongUnmarshallPointer(&instance.Identifier, valueExpr, identifierMap)
	case "NameXSD":
		GongUnmarshallSliceOfPointers(&instance.NameXSD, valueExpr, identifierMap)
	case "Frame":
		instance.Frame = GongExtractString(valueExpr)
	case "BeginPosition":
		GongUnmarshallPointer(&instance.BeginPosition, valueExpr, identifierMap)
	case "EndPosition":
		GongUnmarshallPointer(&instance.EndPosition, valueExpr, identifierMap)
	}
	return nil
}

type TimePositionTypeUnmarshaller struct{}

func (u *TimePositionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TimePositionType), stage, identifier, instanceName, preserveOrder)
}

func (u *TimePositionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TimePositionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Frame":
		instance.Frame = GongExtractString(valueExpr)
	case "IndeterminatePosition":
		GongUnmarshallEnum(&instance.IndeterminatePosition, valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type TimePrimitivePropertyTypeUnmarshaller struct{}

func (u *TimePrimitivePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TimePrimitivePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *TimePrimitivePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TimePrimitivePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
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
	case "TimeInstant":
		GongUnmarshallPointer(&instance.TimeInstant, valueExpr, identifierMap)
	case "TimePeriod":
		GongUnmarshallPointer(&instance.TimePeriod, valueExpr, identifierMap)
	}
	return nil
}
