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
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmfeatures":
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
				case "A_extension_15":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_15), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_16":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_16), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_17":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_17), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_18":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_18), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_19":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_19), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_2":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_2), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_20":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_20), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_21":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_21), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_22":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_22), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_23":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_23), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_24":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_24), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_25":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_25), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_26":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_26), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_27":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_27), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_28":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_28), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_29":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_29), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_3":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_3), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_30":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_30), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_31":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_31), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_32":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_32), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_33":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_33), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_34":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_34), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_35":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_35), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_36":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_36), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_37":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_37), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_38":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_38), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_39":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_39), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_4":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_4), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_40":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_40), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_41":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_41), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_42":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_42), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_43":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_43), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_44":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_44), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_45":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_45), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_46":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_46), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_47":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_47), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_48":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_48), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_49":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_49), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_5":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_5), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_50":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_50), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_51":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_51), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_52":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_52), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_53":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_53), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_54":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_54), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_55":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_55), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_56":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_56), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_57":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_57), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_58":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_58), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_6":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_6), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_7":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_7), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_8":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_8), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_9":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_9), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractApproachLegType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractApproachLegType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractProcedureType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractProcedureType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractPropertiesWithScheduleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractPropertiesWithScheduleType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractSegmentLegType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractSegmentLegType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractSegmentPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractSegmentPointType), stage, ident.Name, instanceName, preserveOrder)
				case "AerialRefuellingPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AerialRefuellingPointType), stage, ident.Name, instanceName, preserveOrder)
				case "AircraftCharacteristicPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AircraftCharacteristicPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AircraftCharacteristicType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AircraftCharacteristicType), stage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirportHeliportPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AltimeterSourcePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AltimeterSourcePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AltitudeAdjustmentPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AltitudeAdjustmentPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AltitudeAdjustmentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AltitudeAdjustmentType), stage, ident.Name, instanceName, preserveOrder)
				case "AngleIndicationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AngleIndicationPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AngleUsePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AngleUsePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AngleUseType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AngleUseType), stage, ident.Name, instanceName, preserveOrder)
				case "ApproachAltitudeTablePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApproachAltitudeTablePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ApproachAltitudeTableType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApproachAltitudeTableType), stage, ident.Name, instanceName, preserveOrder)
				case "ApproachConditionPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApproachConditionPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ApproachConditionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApproachConditionType), stage, ident.Name, instanceName, preserveOrder)
				case "ApproachDistanceTablePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApproachDistanceTablePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ApproachDistanceTableType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApproachDistanceTableType), stage, ident.Name, instanceName, preserveOrder)
				case "ApproachTimingTablePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApproachTimingTablePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ApproachTimingTableType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApproachTimingTableType), stage, ident.Name, instanceName, preserveOrder)
				case "ArrivalFeederLegTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ArrivalFeederLegTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ArrivalFeederLegTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ArrivalFeederLegTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "ArrivalFeederLegType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ArrivalFeederLegType), stage, ident.Name, instanceName, preserveOrder)
				case "ArrivalLegTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ArrivalLegTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ArrivalLegTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ArrivalLegTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "ArrivalLegType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ArrivalLegType), stage, ident.Name, instanceName, preserveOrder)
				case "CircleSectorPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CircleSectorPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "CircleSectorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CircleSectorType), stage, ident.Name, instanceName, preserveOrder)
				case "CirclingAreaTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CirclingAreaTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "CirclingAreaTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CirclingAreaTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "CirclingAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CirclingAreaType), stage, ident.Name, instanceName, preserveOrder)
				case "CirclingRestrictionPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CirclingRestrictionPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "CirclingRestrictionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CirclingRestrictionType), stage, ident.Name, instanceName, preserveOrder)
				case "CurvePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CurvePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "CurveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CurveType), stage, ident.Name, instanceName, preserveOrder)
				case "DMEPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DMEPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "DepartureArrivalConditionPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DepartureArrivalConditionPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "DepartureArrivalConditionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DepartureArrivalConditionType), stage, ident.Name, instanceName, preserveOrder)
				case "DepartureLegTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DepartureLegTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "DepartureLegTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DepartureLegTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "DepartureLegType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DepartureLegType), stage, ident.Name, instanceName, preserveOrder)
				case "DesignatedPointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DesignatedPointPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "DistanceIndicationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DistanceIndicationPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ElevatedCurveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ElevatedCurveType), stage, ident.Name, instanceName, preserveOrder)
				case "ElevatedPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ElevatedPointType), stage, ident.Name, instanceName, preserveOrder)
				case "ElevatedSurfaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ElevatedSurfaceType), stage, ident.Name, instanceName, preserveOrder)
				case "EnRouteSegmentPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(EnRouteSegmentPointType), stage, ident.Name, instanceName, preserveOrder)
				case "EquipmentUnavailableAdjustmentColumnPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(EquipmentUnavailableAdjustmentColumnPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "EquipmentUnavailableAdjustmentColumnType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(EquipmentUnavailableAdjustmentColumnType), stage, ident.Name, instanceName, preserveOrder)
				case "EquipmentUnavailableAdjustmentPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(EquipmentUnavailableAdjustmentPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "EquipmentUnavailableAdjustmentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(EquipmentUnavailableAdjustmentType), stage, ident.Name, instanceName, preserveOrder)
				case "FASDataBlockPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FASDataBlockPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "FASDataBlockType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FASDataBlockType), stage, ident.Name, instanceName, preserveOrder)
				case "FinalLegTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FinalLegTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "FinalLegTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FinalLegTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "FinalLegType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FinalLegType), stage, ident.Name, instanceName, preserveOrder)
				case "FinalProfilePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FinalProfilePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "FinalProfileType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FinalProfileType), stage, ident.Name, instanceName, preserveOrder)
				case "HoldingAssessmentTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(HoldingAssessmentTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "HoldingAssessmentTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(HoldingAssessmentTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "HoldingAssessmentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(HoldingAssessmentType), stage, ident.Name, instanceName, preserveOrder)
				case "HoldingPatternDistancePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(HoldingPatternDistancePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "HoldingPatternDistanceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(HoldingPatternDistanceType), stage, ident.Name, instanceName, preserveOrder)
				case "HoldingPatternDurationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(HoldingPatternDurationPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "HoldingPatternDurationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(HoldingPatternDurationType), stage, ident.Name, instanceName, preserveOrder)
				case "HoldingPatternPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(HoldingPatternPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "HoldingPatternTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(HoldingPatternTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "HoldingPatternTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(HoldingPatternTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "HoldingPatternType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(HoldingPatternType), stage, ident.Name, instanceName, preserveOrder)
				case "HoldingUsePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(HoldingUsePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "HoldingUseType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(HoldingUseType), stage, ident.Name, instanceName, preserveOrder)
				case "InitialLegTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(InitialLegTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "InitialLegTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(InitialLegTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "InitialLegType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(InitialLegType), stage, ident.Name, instanceName, preserveOrder)
				case "InstrumentApproachProcedurePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(InstrumentApproachProcedurePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "InstrumentApproachProcedureTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(InstrumentApproachProcedureTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "InstrumentApproachProcedureTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(InstrumentApproachProcedureTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "InstrumentApproachProcedureType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(InstrumentApproachProcedureType), stage, ident.Name, instanceName, preserveOrder)
				case "IntermediateLegTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(IntermediateLegTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "IntermediateLegTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(IntermediateLegTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "IntermediateLegType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(IntermediateLegType), stage, ident.Name, instanceName, preserveOrder)
				case "LandingTakeoffAreaCollectionPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LandingTakeoffAreaCollectionPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "LandingTakeoffAreaCollectionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LandingTakeoffAreaCollectionType), stage, ident.Name, instanceName, preserveOrder)
				case "LinguisticNotePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LinguisticNotePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "LinguisticNoteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LinguisticNoteType), stage, ident.Name, instanceName, preserveOrder)
				case "MinimaPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MinimaPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "MinimaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MinimaType), stage, ident.Name, instanceName, preserveOrder)
				case "MissedApproachGroupPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MissedApproachGroupPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "MissedApproachGroupType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MissedApproachGroupType), stage, ident.Name, instanceName, preserveOrder)
				case "MissedApproachLegTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MissedApproachLegTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "MissedApproachLegTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MissedApproachLegTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "MissedApproachLegType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MissedApproachLegType), stage, ident.Name, instanceName, preserveOrder)
				case "NavaidPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NavaidPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "NavigationAreaRestrictionTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NavigationAreaRestrictionTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "NavigationAreaRestrictionTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NavigationAreaRestrictionTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "NavigationAreaRestrictionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NavigationAreaRestrictionType), stage, ident.Name, instanceName, preserveOrder)
				case "NavigationAreaSectorPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NavigationAreaSectorPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "NavigationAreaSectorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NavigationAreaSectorType), stage, ident.Name, instanceName, preserveOrder)
				case "NavigationAreaTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NavigationAreaTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "NavigationAreaTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NavigationAreaTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "NavigationAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NavigationAreaType), stage, ident.Name, instanceName, preserveOrder)
				case "NotePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NotePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "NoteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteType), stage, ident.Name, instanceName, preserveOrder)
				case "ObstacleAssessmentAreaPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ObstacleAssessmentAreaPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ObstacleAssessmentAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ObstacleAssessmentAreaType), stage, ident.Name, instanceName, preserveOrder)
				case "ObstaclePlacementPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ObstaclePlacementPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ObstaclePlacementType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ObstaclePlacementType), stage, ident.Name, instanceName, preserveOrder)
				case "ObstructionPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ObstructionPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ObstructionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ObstructionType), stage, ident.Name, instanceName, preserveOrder)
				case "OrganisationAuthorityPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(OrganisationAuthorityPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "PointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PointPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "PointReferencePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PointReferencePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "PointReferenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PointReferenceType), stage, ident.Name, instanceName, preserveOrder)
				case "PointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PointType), stage, ident.Name, instanceName, preserveOrder)
				case "ProcedureAvailabilityPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ProcedureAvailabilityPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ProcedureAvailabilityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ProcedureAvailabilityType), stage, ident.Name, instanceName, preserveOrder)
				case "ProcedureDMETimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ProcedureDMETimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ProcedureDMETimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ProcedureDMETimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "ProcedureDMEType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ProcedureDMEType), stage, ident.Name, instanceName, preserveOrder)
				case "ProcedurePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ProcedurePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ProcedureTransitionLegPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ProcedureTransitionLegPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ProcedureTransitionLegType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ProcedureTransitionLegType), stage, ident.Name, instanceName, preserveOrder)
				case "ProcedureTransitionPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ProcedureTransitionPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ProcedureTransitionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ProcedureTransitionType), stage, ident.Name, instanceName, preserveOrder)
				case "RadarSystemPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RadarSystemPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RadioFrequencyAreaPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RadioFrequencyAreaPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayCentrelinePointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayCentrelinePointPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayDirectionPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayDirectionPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SafeAltitudeAreaPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SafeAltitudeAreaPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SafeAltitudeAreaSectorPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SafeAltitudeAreaSectorPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SafeAltitudeAreaSectorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SafeAltitudeAreaSectorType), stage, ident.Name, instanceName, preserveOrder)
				case "SafeAltitudeAreaTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SafeAltitudeAreaTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SafeAltitudeAreaTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SafeAltitudeAreaTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "SafeAltitudeAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SafeAltitudeAreaType), stage, ident.Name, instanceName, preserveOrder)
				case "SectorDesignPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SectorDesignPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SectorDesignType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SectorDesignType), stage, ident.Name, instanceName, preserveOrder)
				case "SegmentLegPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SegmentLegPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SegmentPointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SegmentPointPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SpecialNavigationSystemPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SpecialNavigationSystemPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "StandardInstrumentArrivalPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StandardInstrumentArrivalPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "StandardInstrumentArrivalTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StandardInstrumentArrivalTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "StandardInstrumentArrivalTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StandardInstrumentArrivalTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "StandardInstrumentArrivalType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StandardInstrumentArrivalType), stage, ident.Name, instanceName, preserveOrder)
				case "StandardInstrumentDeparturePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StandardInstrumentDeparturePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "StandardInstrumentDepartureTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StandardInstrumentDepartureTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "StandardInstrumentDepartureTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StandardInstrumentDepartureTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "StandardInstrumentDepartureType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StandardInstrumentDepartureType), stage, ident.Name, instanceName, preserveOrder)
				case "SurfacePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SurfacePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SurfaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SurfaceType), stage, ident.Name, instanceName, preserveOrder)
				case "TerminalArrivalAreaSectorPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TerminalArrivalAreaSectorPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TerminalArrivalAreaSectorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TerminalArrivalAreaSectorType), stage, ident.Name, instanceName, preserveOrder)
				case "TerminalArrivalAreaTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TerminalArrivalAreaTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TerminalArrivalAreaTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TerminalArrivalAreaTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "TerminalArrivalAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TerminalArrivalAreaType), stage, ident.Name, instanceName, preserveOrder)
				case "TerminalSegmentPointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TerminalSegmentPointPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TerminalSegmentPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TerminalSegmentPointType), stage, ident.Name, instanceName, preserveOrder)
				case "TimesheetPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TimesheetPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TimesheetType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TimesheetType), stage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TouchDownLiftOffPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "UnplannedHoldingPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(UnplannedHoldingPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "UnplannedHoldingTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(UnplannedHoldingTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "UnplannedHoldingTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(UnplannedHoldingTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "UnplannedHoldingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(UnplannedHoldingType), stage, ident.Name, instanceName, preserveOrder)
				case "VerticalStructurePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(VerticalStructurePropertyType), stage, ident.Name, instanceName, preserveOrder)
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
				case *A_extension_15:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_16:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_17:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_18:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_19:
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
				case *A_extension_20:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_21:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_22:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_23:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_24:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_25:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_26:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_27:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_28:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_29:
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
				case *A_extension_30:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_31:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_32:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_33:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_34:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_35:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_36:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_37:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_38:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_39:
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
				case *A_extension_40:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_41:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_42:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_43:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_44:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_45:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_46:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_47:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_48:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_49:
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
				case *A_extension_50:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_51:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_52:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_53:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_54:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_55:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_56:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_57:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_58:
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
				case *AbstractApproachLegType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					}
				case *AbstractProcedureType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
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
				case *AbstractSegmentLegType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					}
				case *AbstractSegmentPointType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					}
				case *AerialRefuellingPointType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "FacilityMakeup":
						__gong__assignSliceOfPointers(&inst.FacilityMakeup, rhs, identifierMap)
					case "PointChoice_fixDesignatedPoint":
						__gong__assignPointer(&inst.PointChoice_fixDesignatedPoint, rhs, identifierMap)
					case "PointChoice_navaidSystem":
						__gong__assignPointer(&inst.PointChoice_navaidSystem, rhs, identifierMap)
					case "PointChoice_position":
						__gong__assignPointer(&inst.PointChoice_position, rhs, identifierMap)
					case "PointChoice_runwayPoint":
						__gong__assignPointer(&inst.PointChoice_runwayPoint, rhs, identifierMap)
					case "PointChoice_aimingPoint":
						__gong__assignPointer(&inst.PointChoice_aimingPoint, rhs, identifierMap)
					case "PointChoice_airportReferencePoint":
						__gong__assignPointer(&inst.PointChoice_airportReferencePoint, rhs, identifierMap)
					case "ExtendedServiceVolume":
						__gong__assignPointer(&inst.ExtendedServiceVolume, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
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
				case *AirportHeliportPropertyType:
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
				case *AltimeterSourcePropertyType:
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
				case *AltitudeAdjustmentPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AltitudeAdjustment":
						__gong__assignPointer(&inst.AltitudeAdjustment, rhs, identifierMap)
					}
				case *AltitudeAdjustmentType:
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
				case *AngleIndicationPropertyType:
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
				case *AngleUsePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AngleUse":
						__gong__assignPointer(&inst.AngleUse, rhs, identifierMap)
					}
				case *AngleUseType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "TheAngleIndication":
						__gong__assignPointer(&inst.TheAngleIndication, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *ApproachAltitudeTablePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ApproachAltitudeTable":
						__gong__assignPointer(&inst.ApproachAltitudeTable, rhs, identifierMap)
					}
				case *ApproachAltitudeTableType:
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
				case *ApproachConditionPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ApproachCondition":
						__gong__assignPointer(&inst.ApproachCondition, rhs, identifierMap)
					}
				case *ApproachConditionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "MinimumSet":
						__gong__assignPointer(&inst.MinimumSet, rhs, identifierMap)
					case "CirclingRestriction":
						__gong__assignSliceOfPointers(&inst.CirclingRestriction, rhs, identifierMap)
					case "AircraftCategory":
						__gong__assignSliceOfPointers(&inst.AircraftCategory, rhs, identifierMap)
					case "LandingArea":
						__gong__assignSliceOfPointers(&inst.LandingArea, rhs, identifierMap)
					case "Altimeter":
						__gong__assignPointer(&inst.Altimeter, rhs, identifierMap)
					case "DesignSurface":
						__gong__assignSliceOfPointers(&inst.DesignSurface, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *ApproachDistanceTablePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ApproachDistanceTable":
						__gong__assignPointer(&inst.ApproachDistanceTable, rhs, identifierMap)
					}
				case *ApproachDistanceTableType:
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
				case *ApproachTimingTablePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ApproachTimingTable":
						__gong__assignPointer(&inst.ApproachTimingTable, rhs, identifierMap)
					}
				case *ApproachTimingTableType:
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
				case *ArrivalFeederLegTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "ArrivalFeederLegTimeSlice":
						__gong__assignPointer(&inst.ArrivalFeederLegTimeSlice, rhs, identifierMap)
					}
				case *ArrivalFeederLegTimeSliceType:
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
					case "StartPoint":
						__gong__assignPointer(&inst.StartPoint, rhs, identifierMap)
					case "EndPoint":
						__gong__assignPointer(&inst.EndPoint, rhs, identifierMap)
					case "Trajectory":
						__gong__assignPointer(&inst.Trajectory, rhs, identifierMap)
					case "ArcCentre":
						__gong__assignPointer(&inst.ArcCentre, rhs, identifierMap)
					case "Angle":
						__gong__assignPointer(&inst.Angle, rhs, identifierMap)
					case "Distance":
						__gong__assignPointer(&inst.Distance, rhs, identifierMap)
					case "AircraftCategory":
						__gong__assignSliceOfPointers(&inst.AircraftCategory, rhs, identifierMap)
					case "Holding":
						__gong__assignPointer(&inst.Holding, rhs, identifierMap)
					case "DesignSurface":
						__gong__assignSliceOfPointers(&inst.DesignSurface, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Approach":
						__gong__assignPointer(&inst.Approach, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *ArrivalFeederLegType:
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
				case *ArrivalLegTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "ArrivalLegTimeSlice":
						__gong__assignPointer(&inst.ArrivalLegTimeSlice, rhs, identifierMap)
					}
				case *ArrivalLegTimeSliceType:
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
					case "StartPoint":
						__gong__assignPointer(&inst.StartPoint, rhs, identifierMap)
					case "EndPoint":
						__gong__assignPointer(&inst.EndPoint, rhs, identifierMap)
					case "Trajectory":
						__gong__assignPointer(&inst.Trajectory, rhs, identifierMap)
					case "ArcCentre":
						__gong__assignPointer(&inst.ArcCentre, rhs, identifierMap)
					case "Angle":
						__gong__assignPointer(&inst.Angle, rhs, identifierMap)
					case "Distance":
						__gong__assignPointer(&inst.Distance, rhs, identifierMap)
					case "AircraftCategory":
						__gong__assignSliceOfPointers(&inst.AircraftCategory, rhs, identifierMap)
					case "Holding":
						__gong__assignPointer(&inst.Holding, rhs, identifierMap)
					case "DesignSurface":
						__gong__assignSliceOfPointers(&inst.DesignSurface, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Arrival":
						__gong__assignPointer(&inst.Arrival, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *ArrivalLegType:
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
				case *CircleSectorPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "CircleSector":
						__gong__assignPointer(&inst.CircleSector, rhs, identifierMap)
					}
				case *CircleSectorType:
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
				case *CirclingAreaTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "CirclingAreaTimeSlice":
						__gong__assignPointer(&inst.CirclingAreaTimeSlice, rhs, identifierMap)
					}
				case *CirclingAreaTimeSliceType:
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
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "Approach":
						__gong__assignPointer(&inst.Approach, rhs, identifierMap)
					case "Condition":
						__gong__assignSliceOfPointers(&inst.Condition, rhs, identifierMap)
					case "AircraftCategory":
						__gong__assignPointer(&inst.AircraftCategory, rhs, identifierMap)
					case "DesignSurface":
						__gong__assignSliceOfPointers(&inst.DesignSurface, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *CirclingAreaType:
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
				case *CirclingRestrictionPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "CirclingRestriction":
						__gong__assignPointer(&inst.CirclingRestriction, rhs, identifierMap)
					}
				case *CirclingRestrictionType:
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
					case "SectorDescription":
						__gong__assignPointer(&inst.SectorDescription, rhs, identifierMap)
					case "RestrictionArea":
						__gong__assignPointer(&inst.RestrictionArea, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
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
				case *DMEPropertyType:
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
				case *DepartureArrivalConditionPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "DepartureArrivalCondition":
						__gong__assignPointer(&inst.DepartureArrivalCondition, rhs, identifierMap)
					}
				case *DepartureArrivalConditionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "EngineType":
						__gong__assignPointer(&inst.EngineType, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *DepartureLegTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "DepartureLegTimeSlice":
						__gong__assignPointer(&inst.DepartureLegTimeSlice, rhs, identifierMap)
					}
				case *DepartureLegTimeSliceType:
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
					case "StartPoint":
						__gong__assignPointer(&inst.StartPoint, rhs, identifierMap)
					case "EndPoint":
						__gong__assignPointer(&inst.EndPoint, rhs, identifierMap)
					case "Trajectory":
						__gong__assignPointer(&inst.Trajectory, rhs, identifierMap)
					case "ArcCentre":
						__gong__assignPointer(&inst.ArcCentre, rhs, identifierMap)
					case "Angle":
						__gong__assignPointer(&inst.Angle, rhs, identifierMap)
					case "Distance":
						__gong__assignPointer(&inst.Distance, rhs, identifierMap)
					case "AircraftCategory":
						__gong__assignSliceOfPointers(&inst.AircraftCategory, rhs, identifierMap)
					case "Holding":
						__gong__assignPointer(&inst.Holding, rhs, identifierMap)
					case "DesignSurface":
						__gong__assignSliceOfPointers(&inst.DesignSurface, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Departure":
						__gong__assignPointer(&inst.Departure, rhs, identifierMap)
					case "Condition":
						__gong__assignSliceOfPointers(&inst.Condition, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *DepartureLegType:
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
				case *DesignatedPointPropertyType:
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
				case *DistanceIndicationPropertyType:
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
				case *ElevatedPointType:
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
				case *EnRouteSegmentPointType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "FacilityMakeup":
						__gong__assignSliceOfPointers(&inst.FacilityMakeup, rhs, identifierMap)
					case "PointChoice_fixDesignatedPoint":
						__gong__assignPointer(&inst.PointChoice_fixDesignatedPoint, rhs, identifierMap)
					case "PointChoice_navaidSystem":
						__gong__assignPointer(&inst.PointChoice_navaidSystem, rhs, identifierMap)
					case "PointChoice_position":
						__gong__assignPointer(&inst.PointChoice_position, rhs, identifierMap)
					case "PointChoice_runwayPoint":
						__gong__assignPointer(&inst.PointChoice_runwayPoint, rhs, identifierMap)
					case "PointChoice_aimingPoint":
						__gong__assignPointer(&inst.PointChoice_aimingPoint, rhs, identifierMap)
					case "PointChoice_airportReferencePoint":
						__gong__assignPointer(&inst.PointChoice_airportReferencePoint, rhs, identifierMap)
					case "ExtendedServiceVolume":
						__gong__assignPointer(&inst.ExtendedServiceVolume, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *EquipmentUnavailableAdjustmentColumnPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EquipmentUnavailableAdjustmentColumn":
						__gong__assignPointer(&inst.EquipmentUnavailableAdjustmentColumn, rhs, identifierMap)
					}
				case *EquipmentUnavailableAdjustmentColumnType:
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
				case *EquipmentUnavailableAdjustmentPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EquipmentUnavailableAdjustment":
						__gong__assignPointer(&inst.EquipmentUnavailableAdjustment, rhs, identifierMap)
					}
				case *EquipmentUnavailableAdjustmentType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "AdjustmentINOPCol":
						__gong__assignSliceOfPointers(&inst.AdjustmentINOPCol, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *FASDataBlockPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "FASDataBlock":
						__gong__assignPointer(&inst.FASDataBlock, rhs, identifierMap)
					}
				case *FASDataBlockType:
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
				case *FinalLegTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "FinalLegTimeSlice":
						__gong__assignPointer(&inst.FinalLegTimeSlice, rhs, identifierMap)
					}
				case *FinalLegTimeSliceType:
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
					case "StartPoint":
						__gong__assignPointer(&inst.StartPoint, rhs, identifierMap)
					case "EndPoint":
						__gong__assignPointer(&inst.EndPoint, rhs, identifierMap)
					case "Trajectory":
						__gong__assignPointer(&inst.Trajectory, rhs, identifierMap)
					case "ArcCentre":
						__gong__assignPointer(&inst.ArcCentre, rhs, identifierMap)
					case "Angle":
						__gong__assignPointer(&inst.Angle, rhs, identifierMap)
					case "Distance":
						__gong__assignPointer(&inst.Distance, rhs, identifierMap)
					case "AircraftCategory":
						__gong__assignSliceOfPointers(&inst.AircraftCategory, rhs, identifierMap)
					case "Holding":
						__gong__assignPointer(&inst.Holding, rhs, identifierMap)
					case "DesignSurface":
						__gong__assignSliceOfPointers(&inst.DesignSurface, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Approach":
						__gong__assignPointer(&inst.Approach, rhs, identifierMap)
					case "Condition":
						__gong__assignSliceOfPointers(&inst.Condition, rhs, identifierMap)
					case "FinalPathAlignmentPoint_fixDesignatedPoint":
						__gong__assignPointer(&inst.FinalPathAlignmentPoint_fixDesignatedPoint, rhs, identifierMap)
					case "FinalPathAlignmentPoint_navaidSystem":
						__gong__assignPointer(&inst.FinalPathAlignmentPoint_navaidSystem, rhs, identifierMap)
					case "FinalPathAlignmentPoint_position":
						__gong__assignPointer(&inst.FinalPathAlignmentPoint_position, rhs, identifierMap)
					case "FinalPathAlignmentPoint_runwayPoint":
						__gong__assignPointer(&inst.FinalPathAlignmentPoint_runwayPoint, rhs, identifierMap)
					case "FinalPathAlignmentPoint_aimingPoint":
						__gong__assignPointer(&inst.FinalPathAlignmentPoint_aimingPoint, rhs, identifierMap)
					case "FinalPathAlignmentPoint_airportReferencePoint":
						__gong__assignPointer(&inst.FinalPathAlignmentPoint_airportReferencePoint, rhs, identifierMap)
					case "VisualDescentPoint":
						__gong__assignPointer(&inst.VisualDescentPoint, rhs, identifierMap)
					case "FASData":
						__gong__assignPointer(&inst.FASData, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *FinalLegType:
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
				case *FinalProfilePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "FinalProfile":
						__gong__assignPointer(&inst.FinalProfile, rhs, identifierMap)
					}
				case *FinalProfileType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Altitude":
						__gong__assignSliceOfPointers(&inst.Altitude, rhs, identifierMap)
					case "Distance":
						__gong__assignSliceOfPointers(&inst.Distance, rhs, identifierMap)
					case "Timing":
						__gong__assignSliceOfPointers(&inst.Timing, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *HoldingAssessmentTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "HoldingAssessmentTimeSlice":
						__gong__assignPointer(&inst.HoldingAssessmentTimeSlice, rhs, identifierMap)
					}
				case *HoldingAssessmentTimeSliceType:
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
					case "HoldingPoint":
						__gong__assignPointer(&inst.HoldingPoint, rhs, identifierMap)
					case "UnplannedHolding":
						__gong__assignPointer(&inst.UnplannedHolding, rhs, identifierMap)
					case "AssessedHoldingPattern":
						__gong__assignPointer(&inst.AssessedHoldingPattern, rhs, identifierMap)
					case "ObstacleAssessment":
						__gong__assignSliceOfPointers(&inst.ObstacleAssessment, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *HoldingAssessmentType:
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
				case *HoldingPatternDistancePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "HoldingPatternDistance":
						__gong__assignPointer(&inst.HoldingPatternDistance, rhs, identifierMap)
					}
				case *HoldingPatternDistanceType:
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
				case *HoldingPatternDurationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "HoldingPatternDuration":
						__gong__assignPointer(&inst.HoldingPatternDuration, rhs, identifierMap)
					}
				case *HoldingPatternDurationType:
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
				case *HoldingPatternPropertyType:
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
				case *HoldingPatternTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "HoldingPatternTimeSlice":
						__gong__assignPointer(&inst.HoldingPatternTimeSlice, rhs, identifierMap)
					}
				case *HoldingPatternTimeSliceType:
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
					case "OutboundLegSpan_endPoint":
						__gong__assignPointer(&inst.OutboundLegSpan_endPoint, rhs, identifierMap)
					case "OutboundLegSpan_endDistance":
						__gong__assignPointer(&inst.OutboundLegSpan_endDistance, rhs, identifierMap)
					case "OutboundLegSpan_endTime":
						__gong__assignPointer(&inst.OutboundLegSpan_endTime, rhs, identifierMap)
					case "HoldingPoint":
						__gong__assignPointer(&inst.HoldingPoint, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *HoldingPatternType:
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
				case *HoldingUsePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "HoldingUse":
						__gong__assignPointer(&inst.HoldingUse, rhs, identifierMap)
					}
				case *HoldingUseType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "TheHoldingPattern":
						__gong__assignPointer(&inst.TheHoldingPattern, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *InitialLegTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "InitialLegTimeSlice":
						__gong__assignPointer(&inst.InitialLegTimeSlice, rhs, identifierMap)
					}
				case *InitialLegTimeSliceType:
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
					case "StartPoint":
						__gong__assignPointer(&inst.StartPoint, rhs, identifierMap)
					case "EndPoint":
						__gong__assignPointer(&inst.EndPoint, rhs, identifierMap)
					case "Trajectory":
						__gong__assignPointer(&inst.Trajectory, rhs, identifierMap)
					case "ArcCentre":
						__gong__assignPointer(&inst.ArcCentre, rhs, identifierMap)
					case "Angle":
						__gong__assignPointer(&inst.Angle, rhs, identifierMap)
					case "Distance":
						__gong__assignPointer(&inst.Distance, rhs, identifierMap)
					case "AircraftCategory":
						__gong__assignSliceOfPointers(&inst.AircraftCategory, rhs, identifierMap)
					case "Holding":
						__gong__assignPointer(&inst.Holding, rhs, identifierMap)
					case "DesignSurface":
						__gong__assignSliceOfPointers(&inst.DesignSurface, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Approach":
						__gong__assignPointer(&inst.Approach, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *InitialLegType:
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
				case *InstrumentApproachProcedurePropertyType:
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
				case *InstrumentApproachProcedureTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "InstrumentApproachProcedureTimeSlice":
						__gong__assignPointer(&inst.InstrumentApproachProcedureTimeSlice, rhs, identifierMap)
					}
				case *InstrumentApproachProcedureTimeSliceType:
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
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "AirportHeliport":
						__gong__assignSliceOfPointers(&inst.AirportHeliport, rhs, identifierMap)
					case "AircraftCharacteristic":
						__gong__assignSliceOfPointers(&inst.AircraftCharacteristic, rhs, identifierMap)
					case "FlightTransition":
						__gong__assignSliceOfPointers(&inst.FlightTransition, rhs, identifierMap)
					case "GuidanceFacility_specialNavigationSystem":
						__gong__assignPointer(&inst.GuidanceFacility_specialNavigationSystem, rhs, identifierMap)
					case "GuidanceFacility_radar":
						__gong__assignPointer(&inst.GuidanceFacility_radar, rhs, identifierMap)
					case "GuidanceFacility_navaid":
						__gong__assignPointer(&inst.GuidanceFacility_navaid, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "SafeAltitude":
						__gong__assignPointer(&inst.SafeAltitude, rhs, identifierMap)
					case "Landing":
						__gong__assignPointer(&inst.Landing, rhs, identifierMap)
					case "MissedInstruction":
						__gong__assignSliceOfPointers(&inst.MissedInstruction, rhs, identifierMap)
					case "FinalProfile":
						__gong__assignPointer(&inst.FinalProfile, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *InstrumentApproachProcedureType:
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
				case *IntermediateLegTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "IntermediateLegTimeSlice":
						__gong__assignPointer(&inst.IntermediateLegTimeSlice, rhs, identifierMap)
					}
				case *IntermediateLegTimeSliceType:
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
					case "StartPoint":
						__gong__assignPointer(&inst.StartPoint, rhs, identifierMap)
					case "EndPoint":
						__gong__assignPointer(&inst.EndPoint, rhs, identifierMap)
					case "Trajectory":
						__gong__assignPointer(&inst.Trajectory, rhs, identifierMap)
					case "ArcCentre":
						__gong__assignPointer(&inst.ArcCentre, rhs, identifierMap)
					case "Angle":
						__gong__assignPointer(&inst.Angle, rhs, identifierMap)
					case "Distance":
						__gong__assignPointer(&inst.Distance, rhs, identifierMap)
					case "AircraftCategory":
						__gong__assignSliceOfPointers(&inst.AircraftCategory, rhs, identifierMap)
					case "Holding":
						__gong__assignPointer(&inst.Holding, rhs, identifierMap)
					case "DesignSurface":
						__gong__assignSliceOfPointers(&inst.DesignSurface, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Approach":
						__gong__assignPointer(&inst.Approach, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *IntermediateLegType:
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
				case *LandingTakeoffAreaCollectionPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "LandingTakeoffAreaCollection":
						__gong__assignPointer(&inst.LandingTakeoffAreaCollection, rhs, identifierMap)
					}
				case *LandingTakeoffAreaCollectionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Runway":
						__gong__assignSliceOfPointers(&inst.Runway, rhs, identifierMap)
					case "TLOF":
						__gong__assignSliceOfPointers(&inst.TLOF, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
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
				case *MinimaPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Minima":
						__gong__assignPointer(&inst.Minima, rhs, identifierMap)
					}
				case *MinimaType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "AdjustmentINOP":
						__gong__assignSliceOfPointers(&inst.AdjustmentINOP, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *MissedApproachGroupPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "MissedApproachGroup":
						__gong__assignPointer(&inst.MissedApproachGroup, rhs, identifierMap)
					}
				case *MissedApproachGroupType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Altimeter":
						__gong__assignSliceOfPointers(&inst.Altimeter, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *MissedApproachLegTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "MissedApproachLegTimeSlice":
						__gong__assignPointer(&inst.MissedApproachLegTimeSlice, rhs, identifierMap)
					}
				case *MissedApproachLegTimeSliceType:
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
					case "StartPoint":
						__gong__assignPointer(&inst.StartPoint, rhs, identifierMap)
					case "EndPoint":
						__gong__assignPointer(&inst.EndPoint, rhs, identifierMap)
					case "Trajectory":
						__gong__assignPointer(&inst.Trajectory, rhs, identifierMap)
					case "ArcCentre":
						__gong__assignPointer(&inst.ArcCentre, rhs, identifierMap)
					case "Angle":
						__gong__assignPointer(&inst.Angle, rhs, identifierMap)
					case "Distance":
						__gong__assignPointer(&inst.Distance, rhs, identifierMap)
					case "AircraftCategory":
						__gong__assignSliceOfPointers(&inst.AircraftCategory, rhs, identifierMap)
					case "Holding":
						__gong__assignPointer(&inst.Holding, rhs, identifierMap)
					case "DesignSurface":
						__gong__assignSliceOfPointers(&inst.DesignSurface, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Approach":
						__gong__assignPointer(&inst.Approach, rhs, identifierMap)
					case "Condition":
						__gong__assignSliceOfPointers(&inst.Condition, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *MissedApproachLegType:
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
				case *NavaidPropertyType:
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
				case *NavigationAreaRestrictionTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NavigationAreaRestrictionTimeSlice":
						__gong__assignPointer(&inst.NavigationAreaRestrictionTimeSlice, rhs, identifierMap)
					}
				case *NavigationAreaRestrictionTimeSliceType:
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
					case "Procedure":
						__gong__assignSliceOfPointers(&inst.Procedure, rhs, identifierMap)
					case "DesignSurface":
						__gong__assignPointer(&inst.DesignSurface, rhs, identifierMap)
					case "SectorDefinition":
						__gong__assignPointer(&inst.SectorDefinition, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *NavigationAreaRestrictionType:
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
				case *NavigationAreaSectorPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "NavigationAreaSector":
						__gong__assignPointer(&inst.NavigationAreaSector, rhs, identifierMap)
					}
				case *NavigationAreaSectorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "SectorDefinition":
						__gong__assignPointer(&inst.SectorDefinition, rhs, identifierMap)
					case "SignificantObstacle":
						__gong__assignSliceOfPointers(&inst.SignificantObstacle, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "SectorCriteria":
						__gong__assignSliceOfPointers(&inst.SectorCriteria, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *NavigationAreaTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NavigationAreaTimeSlice":
						__gong__assignPointer(&inst.NavigationAreaTimeSlice, rhs, identifierMap)
					}
				case *NavigationAreaTimeSliceType:
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
					case "Departure":
						__gong__assignPointer(&inst.Departure, rhs, identifierMap)
					case "Sector":
						__gong__assignSliceOfPointers(&inst.Sector, rhs, identifierMap)
					case "CentrePoint_fixDesignatedPoint":
						__gong__assignPointer(&inst.CentrePoint_fixDesignatedPoint, rhs, identifierMap)
					case "CentrePoint_navaidSystem":
						__gong__assignPointer(&inst.CentrePoint_navaidSystem, rhs, identifierMap)
					case "CentrePoint_position":
						__gong__assignPointer(&inst.CentrePoint_position, rhs, identifierMap)
					case "CentrePoint_runwayPoint":
						__gong__assignPointer(&inst.CentrePoint_runwayPoint, rhs, identifierMap)
					case "CentrePoint_aimingPoint":
						__gong__assignPointer(&inst.CentrePoint_aimingPoint, rhs, identifierMap)
					case "CentrePoint_airportReferencePoint":
						__gong__assignPointer(&inst.CentrePoint_airportReferencePoint, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *NavigationAreaType:
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
				case *ObstacleAssessmentAreaPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ObstacleAssessmentArea":
						__gong__assignPointer(&inst.ObstacleAssessmentArea, rhs, identifierMap)
					}
				case *ObstacleAssessmentAreaType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "AircraftCategory":
						__gong__assignSliceOfPointers(&inst.AircraftCategory, rhs, identifierMap)
					case "SignificantObstacle":
						__gong__assignSliceOfPointers(&inst.SignificantObstacle, rhs, identifierMap)
					case "Surface":
						__gong__assignPointer(&inst.Surface, rhs, identifierMap)
					case "StartingCurve":
						__gong__assignPointer(&inst.StartingCurve, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *ObstaclePlacementPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ObstaclePlacement":
						__gong__assignPointer(&inst.ObstaclePlacement, rhs, identifierMap)
					}
				case *ObstaclePlacementType:
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
				case *ObstructionPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Obstruction":
						__gong__assignPointer(&inst.Obstruction, rhs, identifierMap)
					}
				case *ObstructionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TheVerticalStructure":
						__gong__assignPointer(&inst.TheVerticalStructure, rhs, identifierMap)
					case "Adjustment":
						__gong__assignSliceOfPointers(&inst.Adjustment, rhs, identifierMap)
					case "ObstaclePlacement":
						__gong__assignSliceOfPointers(&inst.ObstaclePlacement, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
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
				case *PointPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Point":
						__gong__assignPointer(&inst.Point, rhs, identifierMap)
					case "ElevatedPoint":
						__gong__assignPointer(&inst.ElevatedPoint, rhs, identifierMap)
					}
				case *PointReferencePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "PointReference":
						__gong__assignPointer(&inst.PointReference, rhs, identifierMap)
					}
				case *PointReferenceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Point":
						__gong__assignPointer(&inst.Point, rhs, identifierMap)
					case "FacilityAngle":
						__gong__assignSliceOfPointers(&inst.FacilityAngle, rhs, identifierMap)
					case "FacilityDistance":
						__gong__assignSliceOfPointers(&inst.FacilityDistance, rhs, identifierMap)
					case "FixToleranceArea":
						__gong__assignPointer(&inst.FixToleranceArea, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *PointType:
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
				case *ProcedureAvailabilityPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ProcedureAvailability":
						__gong__assignPointer(&inst.ProcedureAvailability, rhs, identifierMap)
					}
				case *ProcedureAvailabilityType:
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
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *ProcedureDMETimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "ProcedureDMETimeSlice":
						__gong__assignPointer(&inst.ProcedureDMETimeSlice, rhs, identifierMap)
					}
				case *ProcedureDMETimeSliceType:
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
					case "DME":
						__gong__assignPointer(&inst.DME, rhs, identifierMap)
					case "SegmentLeg":
						__gong__assignPointer(&inst.SegmentLeg, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *ProcedureDMEType:
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
				case *ProcedurePropertyType:
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
				case *ProcedureTransitionLegPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ProcedureTransitionLeg":
						__gong__assignPointer(&inst.ProcedureTransitionLeg, rhs, identifierMap)
					}
				case *ProcedureTransitionLegType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "TheSegmentLeg":
						__gong__assignPointer(&inst.TheSegmentLeg, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *ProcedureTransitionPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ProcedureTransition":
						__gong__assignPointer(&inst.ProcedureTransition, rhs, identifierMap)
					}
				case *ProcedureTransitionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "DepartureRunwayTransition":
						__gong__assignPointer(&inst.DepartureRunwayTransition, rhs, identifierMap)
					case "Trajectory":
						__gong__assignPointer(&inst.Trajectory, rhs, identifierMap)
					case "TransitionLeg":
						__gong__assignSliceOfPointers(&inst.TransitionLeg, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *RadarSystemPropertyType:
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
				case *RadioFrequencyAreaPropertyType:
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
				case *RunwayCentrelinePointPropertyType:
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
				case *RunwayDirectionPropertyType:
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
				case *SafeAltitudeAreaPropertyType:
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
				case *SafeAltitudeAreaSectorPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "SafeAltitudeAreaSector":
						__gong__assignPointer(&inst.SafeAltitudeAreaSector, rhs, identifierMap)
					}
				case *SafeAltitudeAreaSectorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "SignificantObstacle":
						__gong__assignSliceOfPointers(&inst.SignificantObstacle, rhs, identifierMap)
					case "SectorDefinition":
						__gong__assignPointer(&inst.SectorDefinition, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *SafeAltitudeAreaTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "SafeAltitudeAreaTimeSlice":
						__gong__assignPointer(&inst.SafeAltitudeAreaTimeSlice, rhs, identifierMap)
					}
				case *SafeAltitudeAreaTimeSliceType:
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
					case "CentrePoint_fixDesignatedPoint":
						__gong__assignPointer(&inst.CentrePoint_fixDesignatedPoint, rhs, identifierMap)
					case "CentrePoint_navaidSystem":
						__gong__assignPointer(&inst.CentrePoint_navaidSystem, rhs, identifierMap)
					case "CentrePoint_position":
						__gong__assignPointer(&inst.CentrePoint_position, rhs, identifierMap)
					case "CentrePoint_runwayPoint":
						__gong__assignPointer(&inst.CentrePoint_runwayPoint, rhs, identifierMap)
					case "CentrePoint_aimingPoint":
						__gong__assignPointer(&inst.CentrePoint_aimingPoint, rhs, identifierMap)
					case "CentrePoint_airportReferencePoint":
						__gong__assignPointer(&inst.CentrePoint_airportReferencePoint, rhs, identifierMap)
					case "Sector":
						__gong__assignSliceOfPointers(&inst.Sector, rhs, identifierMap)
					case "Location":
						__gong__assignSliceOfPointers(&inst.Location, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *SafeAltitudeAreaType:
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
				case *SectorDesignPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "SectorDesign":
						__gong__assignPointer(&inst.SectorDesign, rhs, identifierMap)
					}
				case *SectorDesignType:
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
				case *SegmentLegPropertyType:
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
				case *SegmentPointPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AerialRefuellingPoint":
						__gong__assignPointer(&inst.AerialRefuellingPoint, rhs, identifierMap)
					case "EnRouteSegmentPoint":
						__gong__assignPointer(&inst.EnRouteSegmentPoint, rhs, identifierMap)
					case "TerminalSegmentPoint":
						__gong__assignPointer(&inst.TerminalSegmentPoint, rhs, identifierMap)
					}
				case *SpecialNavigationSystemPropertyType:
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
				case *StandardInstrumentArrivalPropertyType:
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
				case *StandardInstrumentArrivalTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "StandardInstrumentArrivalTimeSlice":
						__gong__assignPointer(&inst.StandardInstrumentArrivalTimeSlice, rhs, identifierMap)
					}
				case *StandardInstrumentArrivalTimeSliceType:
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
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "AirportHeliport":
						__gong__assignSliceOfPointers(&inst.AirportHeliport, rhs, identifierMap)
					case "AircraftCharacteristic":
						__gong__assignSliceOfPointers(&inst.AircraftCharacteristic, rhs, identifierMap)
					case "FlightTransition":
						__gong__assignSliceOfPointers(&inst.FlightTransition, rhs, identifierMap)
					case "GuidanceFacility_specialNavigationSystem":
						__gong__assignPointer(&inst.GuidanceFacility_specialNavigationSystem, rhs, identifierMap)
					case "GuidanceFacility_radar":
						__gong__assignPointer(&inst.GuidanceFacility_radar, rhs, identifierMap)
					case "GuidanceFacility_navaid":
						__gong__assignPointer(&inst.GuidanceFacility_navaid, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "SafeAltitude":
						__gong__assignPointer(&inst.SafeAltitude, rhs, identifierMap)
					case "Arrival":
						__gong__assignPointer(&inst.Arrival, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *StandardInstrumentArrivalType:
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
				case *StandardInstrumentDeparturePropertyType:
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
				case *StandardInstrumentDepartureTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "StandardInstrumentDepartureTimeSlice":
						__gong__assignPointer(&inst.StandardInstrumentDepartureTimeSlice, rhs, identifierMap)
					}
				case *StandardInstrumentDepartureTimeSliceType:
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
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "AirportHeliport":
						__gong__assignSliceOfPointers(&inst.AirportHeliport, rhs, identifierMap)
					case "AircraftCharacteristic":
						__gong__assignSliceOfPointers(&inst.AircraftCharacteristic, rhs, identifierMap)
					case "FlightTransition":
						__gong__assignSliceOfPointers(&inst.FlightTransition, rhs, identifierMap)
					case "GuidanceFacility_specialNavigationSystem":
						__gong__assignPointer(&inst.GuidanceFacility_specialNavigationSystem, rhs, identifierMap)
					case "GuidanceFacility_radar":
						__gong__assignPointer(&inst.GuidanceFacility_radar, rhs, identifierMap)
					case "GuidanceFacility_navaid":
						__gong__assignPointer(&inst.GuidanceFacility_navaid, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "SafeAltitude":
						__gong__assignPointer(&inst.SafeAltitude, rhs, identifierMap)
					case "Takeoff":
						__gong__assignPointer(&inst.Takeoff, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *StandardInstrumentDepartureType:
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
				case *TerminalArrivalAreaSectorPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "TerminalArrivalAreaSector":
						__gong__assignPointer(&inst.TerminalArrivalAreaSector, rhs, identifierMap)
					}
				case *TerminalArrivalAreaSectorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "SectorDefinition":
						__gong__assignPointer(&inst.SectorDefinition, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "SignificantObstacle":
						__gong__assignSliceOfPointers(&inst.SignificantObstacle, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *TerminalArrivalAreaTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "TerminalArrivalAreaTimeSlice":
						__gong__assignPointer(&inst.TerminalArrivalAreaTimeSlice, rhs, identifierMap)
					}
				case *TerminalArrivalAreaTimeSliceType:
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
					case "IF_fixDesignatedPoint":
						__gong__assignPointer(&inst.IF_fixDesignatedPoint, rhs, identifierMap)
					case "IF_navaidSystem":
						__gong__assignPointer(&inst.IF_navaidSystem, rhs, identifierMap)
					case "IF_position":
						__gong__assignPointer(&inst.IF_position, rhs, identifierMap)
					case "IF_runwayPoint":
						__gong__assignPointer(&inst.IF_runwayPoint, rhs, identifierMap)
					case "IF_aimingPoint":
						__gong__assignPointer(&inst.IF_aimingPoint, rhs, identifierMap)
					case "IF_airportReferencePoint":
						__gong__assignPointer(&inst.IF_airportReferencePoint, rhs, identifierMap)
					case "IAF_fixDesignatedPoint":
						__gong__assignPointer(&inst.IAF_fixDesignatedPoint, rhs, identifierMap)
					case "IAF_navaidSystem":
						__gong__assignPointer(&inst.IAF_navaidSystem, rhs, identifierMap)
					case "IAF_position":
						__gong__assignPointer(&inst.IAF_position, rhs, identifierMap)
					case "IAF_runwayPoint":
						__gong__assignPointer(&inst.IAF_runwayPoint, rhs, identifierMap)
					case "IAF_aimingPoint":
						__gong__assignPointer(&inst.IAF_aimingPoint, rhs, identifierMap)
					case "IAF_airportReferencePoint":
						__gong__assignPointer(&inst.IAF_airportReferencePoint, rhs, identifierMap)
					case "Buffer":
						__gong__assignPointer(&inst.Buffer, rhs, identifierMap)
					case "Sector":
						__gong__assignSliceOfPointers(&inst.Sector, rhs, identifierMap)
					case "ApproachRNAV":
						__gong__assignPointer(&inst.ApproachRNAV, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *TerminalArrivalAreaType:
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
				case *TerminalSegmentPointPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "TerminalSegmentPoint":
						__gong__assignPointer(&inst.TerminalSegmentPoint, rhs, identifierMap)
					}
				case *TerminalSegmentPointType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "FacilityMakeup":
						__gong__assignSliceOfPointers(&inst.FacilityMakeup, rhs, identifierMap)
					case "PointChoice_fixDesignatedPoint":
						__gong__assignPointer(&inst.PointChoice_fixDesignatedPoint, rhs, identifierMap)
					case "PointChoice_navaidSystem":
						__gong__assignPointer(&inst.PointChoice_navaidSystem, rhs, identifierMap)
					case "PointChoice_position":
						__gong__assignPointer(&inst.PointChoice_position, rhs, identifierMap)
					case "PointChoice_runwayPoint":
						__gong__assignPointer(&inst.PointChoice_runwayPoint, rhs, identifierMap)
					case "PointChoice_aimingPoint":
						__gong__assignPointer(&inst.PointChoice_aimingPoint, rhs, identifierMap)
					case "PointChoice_airportReferencePoint":
						__gong__assignPointer(&inst.PointChoice_airportReferencePoint, rhs, identifierMap)
					case "ExtendedServiceVolume":
						__gong__assignPointer(&inst.ExtendedServiceVolume, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
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
				case *TouchDownLiftOffPropertyType:
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
				case *UnplannedHoldingPropertyType:
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
				case *UnplannedHoldingTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "UnplannedHoldingTimeSlice":
						__gong__assignPointer(&inst.UnplannedHoldingTimeSlice, rhs, identifierMap)
					}
				case *UnplannedHoldingTimeSliceType:
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
					case "HoldingPoint":
						__gong__assignPointer(&inst.HoldingPoint, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *UnplannedHoldingType:
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
				case *VerticalStructurePropertyType:
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

type A_extension_15Unmarshaller struct{}

func (u *A_extension_15Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_15), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_15Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_15)
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

type A_extension_16Unmarshaller struct{}

func (u *A_extension_16Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_16), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_16Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_16)
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

type A_extension_17Unmarshaller struct{}

func (u *A_extension_17Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_17), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_17Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_17)
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

type A_extension_18Unmarshaller struct{}

func (u *A_extension_18Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_18), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_18Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_18)
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

type A_extension_19Unmarshaller struct{}

func (u *A_extension_19Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_19), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_19Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_19)
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

type A_extension_20Unmarshaller struct{}

func (u *A_extension_20Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_20), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_20Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_20)
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

type A_extension_21Unmarshaller struct{}

func (u *A_extension_21Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_21), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_21Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_21)
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

type A_extension_22Unmarshaller struct{}

func (u *A_extension_22Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_22), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_22Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_22)
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

type A_extension_23Unmarshaller struct{}

func (u *A_extension_23Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_23), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_23Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_23)
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

type A_extension_24Unmarshaller struct{}

func (u *A_extension_24Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_24), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_24Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_24)
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

type A_extension_25Unmarshaller struct{}

func (u *A_extension_25Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_25), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_25Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_25)
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

type A_extension_26Unmarshaller struct{}

func (u *A_extension_26Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_26), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_26Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_26)
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

type A_extension_27Unmarshaller struct{}

func (u *A_extension_27Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_27), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_27Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_27)
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

type A_extension_28Unmarshaller struct{}

func (u *A_extension_28Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_28), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_28Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_28)
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

type A_extension_29Unmarshaller struct{}

func (u *A_extension_29Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_29), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_29Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_29)
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

type A_extension_30Unmarshaller struct{}

func (u *A_extension_30Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_30), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_30Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_30)
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

type A_extension_31Unmarshaller struct{}

func (u *A_extension_31Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_31), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_31Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_31)
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

type A_extension_32Unmarshaller struct{}

func (u *A_extension_32Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_32), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_32Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_32)
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

type A_extension_33Unmarshaller struct{}

func (u *A_extension_33Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_33), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_33Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_33)
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

type A_extension_34Unmarshaller struct{}

func (u *A_extension_34Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_34), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_34Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_34)
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

type A_extension_35Unmarshaller struct{}

func (u *A_extension_35Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_35), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_35Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_35)
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

type A_extension_36Unmarshaller struct{}

func (u *A_extension_36Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_36), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_36Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_36)
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

type A_extension_37Unmarshaller struct{}

func (u *A_extension_37Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_37), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_37Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_37)
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

type A_extension_38Unmarshaller struct{}

func (u *A_extension_38Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_38), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_38Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_38)
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

type A_extension_39Unmarshaller struct{}

func (u *A_extension_39Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_39), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_39Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_39)
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

type A_extension_40Unmarshaller struct{}

func (u *A_extension_40Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_40), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_40Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_40)
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

type A_extension_41Unmarshaller struct{}

func (u *A_extension_41Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_41), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_41Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_41)
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

type A_extension_42Unmarshaller struct{}

func (u *A_extension_42Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_42), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_42Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_42)
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

type A_extension_43Unmarshaller struct{}

func (u *A_extension_43Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_43), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_43Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_43)
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

type A_extension_44Unmarshaller struct{}

func (u *A_extension_44Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_44), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_44Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_44)
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

type A_extension_45Unmarshaller struct{}

func (u *A_extension_45Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_45), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_45Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_45)
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

type A_extension_46Unmarshaller struct{}

func (u *A_extension_46Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_46), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_46Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_46)
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

type A_extension_47Unmarshaller struct{}

func (u *A_extension_47Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_47), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_47Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_47)
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

type A_extension_48Unmarshaller struct{}

func (u *A_extension_48Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_48), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_48Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_48)
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

type A_extension_49Unmarshaller struct{}

func (u *A_extension_49Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_49), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_49Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_49)
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

type A_extension_50Unmarshaller struct{}

func (u *A_extension_50Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_50), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_50Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_50)
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

type A_extension_51Unmarshaller struct{}

func (u *A_extension_51Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_51), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_51Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_51)
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

type A_extension_52Unmarshaller struct{}

func (u *A_extension_52Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_52), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_52Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_52)
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

type A_extension_53Unmarshaller struct{}

func (u *A_extension_53Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_53), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_53Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_53)
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

type A_extension_54Unmarshaller struct{}

func (u *A_extension_54Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_54), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_54Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_54)
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

type A_extension_55Unmarshaller struct{}

func (u *A_extension_55Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_55), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_55Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_55)
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

type A_extension_56Unmarshaller struct{}

func (u *A_extension_56Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_56), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_56Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_56)
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

type A_extension_57Unmarshaller struct{}

func (u *A_extension_57Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_57), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_57Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_57)
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

type A_extension_58Unmarshaller struct{}

func (u *A_extension_58Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_58), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_58Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_58)
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

type AbstractApproachLegTypeUnmarshaller struct{}

func (u *AbstractApproachLegTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractApproachLegType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractApproachLegTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractApproachLegType)
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

type AbstractProcedureTypeUnmarshaller struct{}

func (u *AbstractProcedureTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractProcedureType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractProcedureTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractProcedureType)
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

type AbstractSegmentLegTypeUnmarshaller struct{}

func (u *AbstractSegmentLegTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractSegmentLegType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractSegmentLegTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractSegmentLegType)
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

type AbstractSegmentPointTypeUnmarshaller struct{}

func (u *AbstractSegmentPointTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractSegmentPointType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractSegmentPointTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractSegmentPointType)
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

type AerialRefuellingPointTypeUnmarshaller struct{}

func (u *AerialRefuellingPointTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AerialRefuellingPointType), stage, identifier, instanceName, preserveOrder)
}

func (u *AerialRefuellingPointTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AerialRefuellingPointType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "FacilityMakeup":
		GongUnmarshallSliceOfPointers(&instance.FacilityMakeup, valueExpr, identifierMap)
	case "PointChoice_fixDesignatedPoint":
		GongUnmarshallPointer(&instance.PointChoice_fixDesignatedPoint, valueExpr, identifierMap)
	case "PointChoice_navaidSystem":
		GongUnmarshallPointer(&instance.PointChoice_navaidSystem, valueExpr, identifierMap)
	case "PointChoice_position":
		GongUnmarshallPointer(&instance.PointChoice_position, valueExpr, identifierMap)
	case "PointChoice_runwayPoint":
		GongUnmarshallPointer(&instance.PointChoice_runwayPoint, valueExpr, identifierMap)
	case "PointChoice_aimingPoint":
		GongUnmarshallPointer(&instance.PointChoice_aimingPoint, valueExpr, identifierMap)
	case "PointChoice_airportReferencePoint":
		GongUnmarshallPointer(&instance.PointChoice_airportReferencePoint, valueExpr, identifierMap)
	case "ExtendedServiceVolume":
		GongUnmarshallPointer(&instance.ExtendedServiceVolume, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
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

type AirportHeliportPropertyTypeUnmarshaller struct{}

func (u *AirportHeliportPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirportHeliportPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirportHeliportPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirportHeliportPropertyType)
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

type AltimeterSourcePropertyTypeUnmarshaller struct{}

func (u *AltimeterSourcePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AltimeterSourcePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AltimeterSourcePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AltimeterSourcePropertyType)
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

type AltitudeAdjustmentPropertyTypeUnmarshaller struct{}

func (u *AltitudeAdjustmentPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AltitudeAdjustmentPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AltitudeAdjustmentPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AltitudeAdjustmentPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "AltitudeAdjustment":
		GongUnmarshallPointer(&instance.AltitudeAdjustment, valueExpr, identifierMap)
	}
	return nil
}

type AltitudeAdjustmentTypeUnmarshaller struct{}

func (u *AltitudeAdjustmentTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AltitudeAdjustmentType), stage, identifier, instanceName, preserveOrder)
}

func (u *AltitudeAdjustmentTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AltitudeAdjustmentType)
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

type AngleIndicationPropertyTypeUnmarshaller struct{}

func (u *AngleIndicationPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AngleIndicationPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AngleIndicationPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AngleIndicationPropertyType)
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

type AngleUsePropertyTypeUnmarshaller struct{}

func (u *AngleUsePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AngleUsePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AngleUsePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AngleUsePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "AngleUse":
		GongUnmarshallPointer(&instance.AngleUse, valueExpr, identifierMap)
	}
	return nil
}

type AngleUseTypeUnmarshaller struct{}

func (u *AngleUseTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AngleUseType), stage, identifier, instanceName, preserveOrder)
}

func (u *AngleUseTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AngleUseType)
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
	case "TheAngleIndication":
		GongUnmarshallPointer(&instance.TheAngleIndication, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type ApproachAltitudeTablePropertyTypeUnmarshaller struct{}

func (u *ApproachAltitudeTablePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApproachAltitudeTablePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApproachAltitudeTablePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApproachAltitudeTablePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "ApproachAltitudeTable":
		GongUnmarshallPointer(&instance.ApproachAltitudeTable, valueExpr, identifierMap)
	}
	return nil
}

type ApproachAltitudeTableTypeUnmarshaller struct{}

func (u *ApproachAltitudeTableTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApproachAltitudeTableType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApproachAltitudeTableTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApproachAltitudeTableType)
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

type ApproachConditionPropertyTypeUnmarshaller struct{}

func (u *ApproachConditionPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApproachConditionPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApproachConditionPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApproachConditionPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "ApproachCondition":
		GongUnmarshallPointer(&instance.ApproachCondition, valueExpr, identifierMap)
	}
	return nil
}

type ApproachConditionTypeUnmarshaller struct{}

func (u *ApproachConditionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApproachConditionType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApproachConditionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApproachConditionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "MinimumSet":
		GongUnmarshallPointer(&instance.MinimumSet, valueExpr, identifierMap)
	case "CirclingRestriction":
		GongUnmarshallSliceOfPointers(&instance.CirclingRestriction, valueExpr, identifierMap)
	case "AircraftCategory":
		GongUnmarshallSliceOfPointers(&instance.AircraftCategory, valueExpr, identifierMap)
	case "LandingArea":
		GongUnmarshallSliceOfPointers(&instance.LandingArea, valueExpr, identifierMap)
	case "Altimeter":
		GongUnmarshallPointer(&instance.Altimeter, valueExpr, identifierMap)
	case "DesignSurface":
		GongUnmarshallSliceOfPointers(&instance.DesignSurface, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type ApproachDistanceTablePropertyTypeUnmarshaller struct{}

func (u *ApproachDistanceTablePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApproachDistanceTablePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApproachDistanceTablePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApproachDistanceTablePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "ApproachDistanceTable":
		GongUnmarshallPointer(&instance.ApproachDistanceTable, valueExpr, identifierMap)
	}
	return nil
}

type ApproachDistanceTableTypeUnmarshaller struct{}

func (u *ApproachDistanceTableTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApproachDistanceTableType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApproachDistanceTableTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApproachDistanceTableType)
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

type ApproachTimingTablePropertyTypeUnmarshaller struct{}

func (u *ApproachTimingTablePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApproachTimingTablePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApproachTimingTablePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApproachTimingTablePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "ApproachTimingTable":
		GongUnmarshallPointer(&instance.ApproachTimingTable, valueExpr, identifierMap)
	}
	return nil
}

type ApproachTimingTableTypeUnmarshaller struct{}

func (u *ApproachTimingTableTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApproachTimingTableType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApproachTimingTableTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApproachTimingTableType)
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

type ArrivalFeederLegTimeSlicePropertyTypeUnmarshaller struct{}

func (u *ArrivalFeederLegTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ArrivalFeederLegTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ArrivalFeederLegTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ArrivalFeederLegTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "ArrivalFeederLegTimeSlice":
		GongUnmarshallPointer(&instance.ArrivalFeederLegTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type ArrivalFeederLegTimeSliceTypeUnmarshaller struct{}

func (u *ArrivalFeederLegTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ArrivalFeederLegTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *ArrivalFeederLegTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ArrivalFeederLegTimeSliceType)
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
	case "StartPoint":
		GongUnmarshallPointer(&instance.StartPoint, valueExpr, identifierMap)
	case "EndPoint":
		GongUnmarshallPointer(&instance.EndPoint, valueExpr, identifierMap)
	case "Trajectory":
		GongUnmarshallPointer(&instance.Trajectory, valueExpr, identifierMap)
	case "ArcCentre":
		GongUnmarshallPointer(&instance.ArcCentre, valueExpr, identifierMap)
	case "Angle":
		GongUnmarshallPointer(&instance.Angle, valueExpr, identifierMap)
	case "Distance":
		GongUnmarshallPointer(&instance.Distance, valueExpr, identifierMap)
	case "AircraftCategory":
		GongUnmarshallSliceOfPointers(&instance.AircraftCategory, valueExpr, identifierMap)
	case "Holding":
		GongUnmarshallPointer(&instance.Holding, valueExpr, identifierMap)
	case "DesignSurface":
		GongUnmarshallSliceOfPointers(&instance.DesignSurface, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Approach":
		GongUnmarshallPointer(&instance.Approach, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type ArrivalFeederLegTypeUnmarshaller struct{}

func (u *ArrivalFeederLegTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ArrivalFeederLegType), stage, identifier, instanceName, preserveOrder)
}

func (u *ArrivalFeederLegTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ArrivalFeederLegType)
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

type ArrivalLegTimeSlicePropertyTypeUnmarshaller struct{}

func (u *ArrivalLegTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ArrivalLegTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ArrivalLegTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ArrivalLegTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "ArrivalLegTimeSlice":
		GongUnmarshallPointer(&instance.ArrivalLegTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type ArrivalLegTimeSliceTypeUnmarshaller struct{}

func (u *ArrivalLegTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ArrivalLegTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *ArrivalLegTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ArrivalLegTimeSliceType)
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
	case "StartPoint":
		GongUnmarshallPointer(&instance.StartPoint, valueExpr, identifierMap)
	case "EndPoint":
		GongUnmarshallPointer(&instance.EndPoint, valueExpr, identifierMap)
	case "Trajectory":
		GongUnmarshallPointer(&instance.Trajectory, valueExpr, identifierMap)
	case "ArcCentre":
		GongUnmarshallPointer(&instance.ArcCentre, valueExpr, identifierMap)
	case "Angle":
		GongUnmarshallPointer(&instance.Angle, valueExpr, identifierMap)
	case "Distance":
		GongUnmarshallPointer(&instance.Distance, valueExpr, identifierMap)
	case "AircraftCategory":
		GongUnmarshallSliceOfPointers(&instance.AircraftCategory, valueExpr, identifierMap)
	case "Holding":
		GongUnmarshallPointer(&instance.Holding, valueExpr, identifierMap)
	case "DesignSurface":
		GongUnmarshallSliceOfPointers(&instance.DesignSurface, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Arrival":
		GongUnmarshallPointer(&instance.Arrival, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type ArrivalLegTypeUnmarshaller struct{}

func (u *ArrivalLegTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ArrivalLegType), stage, identifier, instanceName, preserveOrder)
}

func (u *ArrivalLegTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ArrivalLegType)
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

type CircleSectorPropertyTypeUnmarshaller struct{}

func (u *CircleSectorPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CircleSectorPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *CircleSectorPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CircleSectorPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "CircleSector":
		GongUnmarshallPointer(&instance.CircleSector, valueExpr, identifierMap)
	}
	return nil
}

type CircleSectorTypeUnmarshaller struct{}

func (u *CircleSectorTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CircleSectorType), stage, identifier, instanceName, preserveOrder)
}

func (u *CircleSectorTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CircleSectorType)
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

type CirclingAreaTimeSlicePropertyTypeUnmarshaller struct{}

func (u *CirclingAreaTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CirclingAreaTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *CirclingAreaTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CirclingAreaTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "CirclingAreaTimeSlice":
		GongUnmarshallPointer(&instance.CirclingAreaTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type CirclingAreaTimeSliceTypeUnmarshaller struct{}

func (u *CirclingAreaTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CirclingAreaTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CirclingAreaTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CirclingAreaTimeSliceType)
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
	case "Extent":
		GongUnmarshallPointer(&instance.Extent, valueExpr, identifierMap)
	case "Approach":
		GongUnmarshallPointer(&instance.Approach, valueExpr, identifierMap)
	case "Condition":
		GongUnmarshallSliceOfPointers(&instance.Condition, valueExpr, identifierMap)
	case "AircraftCategory":
		GongUnmarshallPointer(&instance.AircraftCategory, valueExpr, identifierMap)
	case "DesignSurface":
		GongUnmarshallSliceOfPointers(&instance.DesignSurface, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type CirclingAreaTypeUnmarshaller struct{}

func (u *CirclingAreaTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CirclingAreaType), stage, identifier, instanceName, preserveOrder)
}

func (u *CirclingAreaTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CirclingAreaType)
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

type CirclingRestrictionPropertyTypeUnmarshaller struct{}

func (u *CirclingRestrictionPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CirclingRestrictionPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *CirclingRestrictionPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CirclingRestrictionPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "CirclingRestriction":
		GongUnmarshallPointer(&instance.CirclingRestriction, valueExpr, identifierMap)
	}
	return nil
}

type CirclingRestrictionTypeUnmarshaller struct{}

func (u *CirclingRestrictionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CirclingRestrictionType), stage, identifier, instanceName, preserveOrder)
}

func (u *CirclingRestrictionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CirclingRestrictionType)
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
	case "SectorDescription":
		GongUnmarshallPointer(&instance.SectorDescription, valueExpr, identifierMap)
	case "RestrictionArea":
		GongUnmarshallPointer(&instance.RestrictionArea, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
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

type DMEPropertyTypeUnmarshaller struct{}

func (u *DMEPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DMEPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *DMEPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DMEPropertyType)
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

type DepartureArrivalConditionPropertyTypeUnmarshaller struct{}

func (u *DepartureArrivalConditionPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DepartureArrivalConditionPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *DepartureArrivalConditionPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DepartureArrivalConditionPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "DepartureArrivalCondition":
		GongUnmarshallPointer(&instance.DepartureArrivalCondition, valueExpr, identifierMap)
	}
	return nil
}

type DepartureArrivalConditionTypeUnmarshaller struct{}

func (u *DepartureArrivalConditionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DepartureArrivalConditionType), stage, identifier, instanceName, preserveOrder)
}

func (u *DepartureArrivalConditionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DepartureArrivalConditionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "EngineType":
		GongUnmarshallPointer(&instance.EngineType, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type DepartureLegTimeSlicePropertyTypeUnmarshaller struct{}

func (u *DepartureLegTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DepartureLegTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *DepartureLegTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DepartureLegTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "DepartureLegTimeSlice":
		GongUnmarshallPointer(&instance.DepartureLegTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type DepartureLegTimeSliceTypeUnmarshaller struct{}

func (u *DepartureLegTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DepartureLegTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *DepartureLegTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DepartureLegTimeSliceType)
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
	case "StartPoint":
		GongUnmarshallPointer(&instance.StartPoint, valueExpr, identifierMap)
	case "EndPoint":
		GongUnmarshallPointer(&instance.EndPoint, valueExpr, identifierMap)
	case "Trajectory":
		GongUnmarshallPointer(&instance.Trajectory, valueExpr, identifierMap)
	case "ArcCentre":
		GongUnmarshallPointer(&instance.ArcCentre, valueExpr, identifierMap)
	case "Angle":
		GongUnmarshallPointer(&instance.Angle, valueExpr, identifierMap)
	case "Distance":
		GongUnmarshallPointer(&instance.Distance, valueExpr, identifierMap)
	case "AircraftCategory":
		GongUnmarshallSliceOfPointers(&instance.AircraftCategory, valueExpr, identifierMap)
	case "Holding":
		GongUnmarshallPointer(&instance.Holding, valueExpr, identifierMap)
	case "DesignSurface":
		GongUnmarshallSliceOfPointers(&instance.DesignSurface, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Departure":
		GongUnmarshallPointer(&instance.Departure, valueExpr, identifierMap)
	case "Condition":
		GongUnmarshallSliceOfPointers(&instance.Condition, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type DepartureLegTypeUnmarshaller struct{}

func (u *DepartureLegTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DepartureLegType), stage, identifier, instanceName, preserveOrder)
}

func (u *DepartureLegTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DepartureLegType)
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

type DesignatedPointPropertyTypeUnmarshaller struct{}

func (u *DesignatedPointPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DesignatedPointPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *DesignatedPointPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DesignatedPointPropertyType)
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

type DistanceIndicationPropertyTypeUnmarshaller struct{}

func (u *DistanceIndicationPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DistanceIndicationPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *DistanceIndicationPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DistanceIndicationPropertyType)
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

type ElevatedPointTypeUnmarshaller struct{}

func (u *ElevatedPointTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ElevatedPointType), stage, identifier, instanceName, preserveOrder)
}

func (u *ElevatedPointTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ElevatedPointType)
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

type EnRouteSegmentPointTypeUnmarshaller struct{}

func (u *EnRouteSegmentPointTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EnRouteSegmentPointType), stage, identifier, instanceName, preserveOrder)
}

func (u *EnRouteSegmentPointTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EnRouteSegmentPointType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "FacilityMakeup":
		GongUnmarshallSliceOfPointers(&instance.FacilityMakeup, valueExpr, identifierMap)
	case "PointChoice_fixDesignatedPoint":
		GongUnmarshallPointer(&instance.PointChoice_fixDesignatedPoint, valueExpr, identifierMap)
	case "PointChoice_navaidSystem":
		GongUnmarshallPointer(&instance.PointChoice_navaidSystem, valueExpr, identifierMap)
	case "PointChoice_position":
		GongUnmarshallPointer(&instance.PointChoice_position, valueExpr, identifierMap)
	case "PointChoice_runwayPoint":
		GongUnmarshallPointer(&instance.PointChoice_runwayPoint, valueExpr, identifierMap)
	case "PointChoice_aimingPoint":
		GongUnmarshallPointer(&instance.PointChoice_aimingPoint, valueExpr, identifierMap)
	case "PointChoice_airportReferencePoint":
		GongUnmarshallPointer(&instance.PointChoice_airportReferencePoint, valueExpr, identifierMap)
	case "ExtendedServiceVolume":
		GongUnmarshallPointer(&instance.ExtendedServiceVolume, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type EquipmentUnavailableAdjustmentColumnPropertyTypeUnmarshaller struct{}

func (u *EquipmentUnavailableAdjustmentColumnPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EquipmentUnavailableAdjustmentColumnPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *EquipmentUnavailableAdjustmentColumnPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EquipmentUnavailableAdjustmentColumnPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EquipmentUnavailableAdjustmentColumn":
		GongUnmarshallPointer(&instance.EquipmentUnavailableAdjustmentColumn, valueExpr, identifierMap)
	}
	return nil
}

type EquipmentUnavailableAdjustmentColumnTypeUnmarshaller struct{}

func (u *EquipmentUnavailableAdjustmentColumnTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EquipmentUnavailableAdjustmentColumnType), stage, identifier, instanceName, preserveOrder)
}

func (u *EquipmentUnavailableAdjustmentColumnTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EquipmentUnavailableAdjustmentColumnType)
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

type EquipmentUnavailableAdjustmentPropertyTypeUnmarshaller struct{}

func (u *EquipmentUnavailableAdjustmentPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EquipmentUnavailableAdjustmentPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *EquipmentUnavailableAdjustmentPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EquipmentUnavailableAdjustmentPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EquipmentUnavailableAdjustment":
		GongUnmarshallPointer(&instance.EquipmentUnavailableAdjustment, valueExpr, identifierMap)
	}
	return nil
}

type EquipmentUnavailableAdjustmentTypeUnmarshaller struct{}

func (u *EquipmentUnavailableAdjustmentTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EquipmentUnavailableAdjustmentType), stage, identifier, instanceName, preserveOrder)
}

func (u *EquipmentUnavailableAdjustmentTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EquipmentUnavailableAdjustmentType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "AdjustmentINOPCol":
		GongUnmarshallSliceOfPointers(&instance.AdjustmentINOPCol, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type FASDataBlockPropertyTypeUnmarshaller struct{}

func (u *FASDataBlockPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FASDataBlockPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *FASDataBlockPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FASDataBlockPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "FASDataBlock":
		GongUnmarshallPointer(&instance.FASDataBlock, valueExpr, identifierMap)
	}
	return nil
}

type FASDataBlockTypeUnmarshaller struct{}

func (u *FASDataBlockTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FASDataBlockType), stage, identifier, instanceName, preserveOrder)
}

func (u *FASDataBlockTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FASDataBlockType)
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

type FinalLegTimeSlicePropertyTypeUnmarshaller struct{}

func (u *FinalLegTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FinalLegTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *FinalLegTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FinalLegTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "FinalLegTimeSlice":
		GongUnmarshallPointer(&instance.FinalLegTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type FinalLegTimeSliceTypeUnmarshaller struct{}

func (u *FinalLegTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FinalLegTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *FinalLegTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FinalLegTimeSliceType)
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
	case "StartPoint":
		GongUnmarshallPointer(&instance.StartPoint, valueExpr, identifierMap)
	case "EndPoint":
		GongUnmarshallPointer(&instance.EndPoint, valueExpr, identifierMap)
	case "Trajectory":
		GongUnmarshallPointer(&instance.Trajectory, valueExpr, identifierMap)
	case "ArcCentre":
		GongUnmarshallPointer(&instance.ArcCentre, valueExpr, identifierMap)
	case "Angle":
		GongUnmarshallPointer(&instance.Angle, valueExpr, identifierMap)
	case "Distance":
		GongUnmarshallPointer(&instance.Distance, valueExpr, identifierMap)
	case "AircraftCategory":
		GongUnmarshallSliceOfPointers(&instance.AircraftCategory, valueExpr, identifierMap)
	case "Holding":
		GongUnmarshallPointer(&instance.Holding, valueExpr, identifierMap)
	case "DesignSurface":
		GongUnmarshallSliceOfPointers(&instance.DesignSurface, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Approach":
		GongUnmarshallPointer(&instance.Approach, valueExpr, identifierMap)
	case "Condition":
		GongUnmarshallSliceOfPointers(&instance.Condition, valueExpr, identifierMap)
	case "FinalPathAlignmentPoint_fixDesignatedPoint":
		GongUnmarshallPointer(&instance.FinalPathAlignmentPoint_fixDesignatedPoint, valueExpr, identifierMap)
	case "FinalPathAlignmentPoint_navaidSystem":
		GongUnmarshallPointer(&instance.FinalPathAlignmentPoint_navaidSystem, valueExpr, identifierMap)
	case "FinalPathAlignmentPoint_position":
		GongUnmarshallPointer(&instance.FinalPathAlignmentPoint_position, valueExpr, identifierMap)
	case "FinalPathAlignmentPoint_runwayPoint":
		GongUnmarshallPointer(&instance.FinalPathAlignmentPoint_runwayPoint, valueExpr, identifierMap)
	case "FinalPathAlignmentPoint_aimingPoint":
		GongUnmarshallPointer(&instance.FinalPathAlignmentPoint_aimingPoint, valueExpr, identifierMap)
	case "FinalPathAlignmentPoint_airportReferencePoint":
		GongUnmarshallPointer(&instance.FinalPathAlignmentPoint_airportReferencePoint, valueExpr, identifierMap)
	case "VisualDescentPoint":
		GongUnmarshallPointer(&instance.VisualDescentPoint, valueExpr, identifierMap)
	case "FASData":
		GongUnmarshallPointer(&instance.FASData, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type FinalLegTypeUnmarshaller struct{}

func (u *FinalLegTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FinalLegType), stage, identifier, instanceName, preserveOrder)
}

func (u *FinalLegTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FinalLegType)
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

type FinalProfilePropertyTypeUnmarshaller struct{}

func (u *FinalProfilePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FinalProfilePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *FinalProfilePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FinalProfilePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "FinalProfile":
		GongUnmarshallPointer(&instance.FinalProfile, valueExpr, identifierMap)
	}
	return nil
}

type FinalProfileTypeUnmarshaller struct{}

func (u *FinalProfileTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FinalProfileType), stage, identifier, instanceName, preserveOrder)
}

func (u *FinalProfileTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FinalProfileType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Altitude":
		GongUnmarshallSliceOfPointers(&instance.Altitude, valueExpr, identifierMap)
	case "Distance":
		GongUnmarshallSliceOfPointers(&instance.Distance, valueExpr, identifierMap)
	case "Timing":
		GongUnmarshallSliceOfPointers(&instance.Timing, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type HoldingAssessmentTimeSlicePropertyTypeUnmarshaller struct{}

func (u *HoldingAssessmentTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(HoldingAssessmentTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *HoldingAssessmentTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*HoldingAssessmentTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "HoldingAssessmentTimeSlice":
		GongUnmarshallPointer(&instance.HoldingAssessmentTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type HoldingAssessmentTimeSliceTypeUnmarshaller struct{}

func (u *HoldingAssessmentTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(HoldingAssessmentTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *HoldingAssessmentTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*HoldingAssessmentTimeSliceType)
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
	case "HoldingPoint":
		GongUnmarshallPointer(&instance.HoldingPoint, valueExpr, identifierMap)
	case "UnplannedHolding":
		GongUnmarshallPointer(&instance.UnplannedHolding, valueExpr, identifierMap)
	case "AssessedHoldingPattern":
		GongUnmarshallPointer(&instance.AssessedHoldingPattern, valueExpr, identifierMap)
	case "ObstacleAssessment":
		GongUnmarshallSliceOfPointers(&instance.ObstacleAssessment, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type HoldingAssessmentTypeUnmarshaller struct{}

func (u *HoldingAssessmentTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(HoldingAssessmentType), stage, identifier, instanceName, preserveOrder)
}

func (u *HoldingAssessmentTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*HoldingAssessmentType)
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

type HoldingPatternDistancePropertyTypeUnmarshaller struct{}

func (u *HoldingPatternDistancePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(HoldingPatternDistancePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *HoldingPatternDistancePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*HoldingPatternDistancePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "HoldingPatternDistance":
		GongUnmarshallPointer(&instance.HoldingPatternDistance, valueExpr, identifierMap)
	}
	return nil
}

type HoldingPatternDistanceTypeUnmarshaller struct{}

func (u *HoldingPatternDistanceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(HoldingPatternDistanceType), stage, identifier, instanceName, preserveOrder)
}

func (u *HoldingPatternDistanceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*HoldingPatternDistanceType)
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

type HoldingPatternDurationPropertyTypeUnmarshaller struct{}

func (u *HoldingPatternDurationPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(HoldingPatternDurationPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *HoldingPatternDurationPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*HoldingPatternDurationPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "HoldingPatternDuration":
		GongUnmarshallPointer(&instance.HoldingPatternDuration, valueExpr, identifierMap)
	}
	return nil
}

type HoldingPatternDurationTypeUnmarshaller struct{}

func (u *HoldingPatternDurationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(HoldingPatternDurationType), stage, identifier, instanceName, preserveOrder)
}

func (u *HoldingPatternDurationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*HoldingPatternDurationType)
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

type HoldingPatternPropertyTypeUnmarshaller struct{}

func (u *HoldingPatternPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(HoldingPatternPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *HoldingPatternPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*HoldingPatternPropertyType)
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

type HoldingPatternTimeSlicePropertyTypeUnmarshaller struct{}

func (u *HoldingPatternTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(HoldingPatternTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *HoldingPatternTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*HoldingPatternTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "HoldingPatternTimeSlice":
		GongUnmarshallPointer(&instance.HoldingPatternTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type HoldingPatternTimeSliceTypeUnmarshaller struct{}

func (u *HoldingPatternTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(HoldingPatternTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *HoldingPatternTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*HoldingPatternTimeSliceType)
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
	case "OutboundLegSpan_endPoint":
		GongUnmarshallPointer(&instance.OutboundLegSpan_endPoint, valueExpr, identifierMap)
	case "OutboundLegSpan_endDistance":
		GongUnmarshallPointer(&instance.OutboundLegSpan_endDistance, valueExpr, identifierMap)
	case "OutboundLegSpan_endTime":
		GongUnmarshallPointer(&instance.OutboundLegSpan_endTime, valueExpr, identifierMap)
	case "HoldingPoint":
		GongUnmarshallPointer(&instance.HoldingPoint, valueExpr, identifierMap)
	case "Extent":
		GongUnmarshallPointer(&instance.Extent, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type HoldingPatternTypeUnmarshaller struct{}

func (u *HoldingPatternTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(HoldingPatternType), stage, identifier, instanceName, preserveOrder)
}

func (u *HoldingPatternTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*HoldingPatternType)
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

type HoldingUsePropertyTypeUnmarshaller struct{}

func (u *HoldingUsePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(HoldingUsePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *HoldingUsePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*HoldingUsePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "HoldingUse":
		GongUnmarshallPointer(&instance.HoldingUse, valueExpr, identifierMap)
	}
	return nil
}

type HoldingUseTypeUnmarshaller struct{}

func (u *HoldingUseTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(HoldingUseType), stage, identifier, instanceName, preserveOrder)
}

func (u *HoldingUseTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*HoldingUseType)
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
	case "TheHoldingPattern":
		GongUnmarshallPointer(&instance.TheHoldingPattern, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type InitialLegTimeSlicePropertyTypeUnmarshaller struct{}

func (u *InitialLegTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(InitialLegTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *InitialLegTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*InitialLegTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "InitialLegTimeSlice":
		GongUnmarshallPointer(&instance.InitialLegTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type InitialLegTimeSliceTypeUnmarshaller struct{}

func (u *InitialLegTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(InitialLegTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *InitialLegTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*InitialLegTimeSliceType)
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
	case "StartPoint":
		GongUnmarshallPointer(&instance.StartPoint, valueExpr, identifierMap)
	case "EndPoint":
		GongUnmarshallPointer(&instance.EndPoint, valueExpr, identifierMap)
	case "Trajectory":
		GongUnmarshallPointer(&instance.Trajectory, valueExpr, identifierMap)
	case "ArcCentre":
		GongUnmarshallPointer(&instance.ArcCentre, valueExpr, identifierMap)
	case "Angle":
		GongUnmarshallPointer(&instance.Angle, valueExpr, identifierMap)
	case "Distance":
		GongUnmarshallPointer(&instance.Distance, valueExpr, identifierMap)
	case "AircraftCategory":
		GongUnmarshallSliceOfPointers(&instance.AircraftCategory, valueExpr, identifierMap)
	case "Holding":
		GongUnmarshallPointer(&instance.Holding, valueExpr, identifierMap)
	case "DesignSurface":
		GongUnmarshallSliceOfPointers(&instance.DesignSurface, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Approach":
		GongUnmarshallPointer(&instance.Approach, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type InitialLegTypeUnmarshaller struct{}

func (u *InitialLegTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(InitialLegType), stage, identifier, instanceName, preserveOrder)
}

func (u *InitialLegTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*InitialLegType)
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

type InstrumentApproachProcedurePropertyTypeUnmarshaller struct{}

func (u *InstrumentApproachProcedurePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(InstrumentApproachProcedurePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *InstrumentApproachProcedurePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*InstrumentApproachProcedurePropertyType)
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

type InstrumentApproachProcedureTimeSlicePropertyTypeUnmarshaller struct{}

func (u *InstrumentApproachProcedureTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(InstrumentApproachProcedureTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *InstrumentApproachProcedureTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*InstrumentApproachProcedureTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "InstrumentApproachProcedureTimeSlice":
		GongUnmarshallPointer(&instance.InstrumentApproachProcedureTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type InstrumentApproachProcedureTimeSliceTypeUnmarshaller struct{}

func (u *InstrumentApproachProcedureTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(InstrumentApproachProcedureTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *InstrumentApproachProcedureTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*InstrumentApproachProcedureTimeSliceType)
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
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "AirportHeliport":
		GongUnmarshallSliceOfPointers(&instance.AirportHeliport, valueExpr, identifierMap)
	case "AircraftCharacteristic":
		GongUnmarshallSliceOfPointers(&instance.AircraftCharacteristic, valueExpr, identifierMap)
	case "FlightTransition":
		GongUnmarshallSliceOfPointers(&instance.FlightTransition, valueExpr, identifierMap)
	case "GuidanceFacility_specialNavigationSystem":
		GongUnmarshallPointer(&instance.GuidanceFacility_specialNavigationSystem, valueExpr, identifierMap)
	case "GuidanceFacility_radar":
		GongUnmarshallPointer(&instance.GuidanceFacility_radar, valueExpr, identifierMap)
	case "GuidanceFacility_navaid":
		GongUnmarshallPointer(&instance.GuidanceFacility_navaid, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "SafeAltitude":
		GongUnmarshallPointer(&instance.SafeAltitude, valueExpr, identifierMap)
	case "Landing":
		GongUnmarshallPointer(&instance.Landing, valueExpr, identifierMap)
	case "MissedInstruction":
		GongUnmarshallSliceOfPointers(&instance.MissedInstruction, valueExpr, identifierMap)
	case "FinalProfile":
		GongUnmarshallPointer(&instance.FinalProfile, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type InstrumentApproachProcedureTypeUnmarshaller struct{}

func (u *InstrumentApproachProcedureTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(InstrumentApproachProcedureType), stage, identifier, instanceName, preserveOrder)
}

func (u *InstrumentApproachProcedureTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*InstrumentApproachProcedureType)
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

type IntermediateLegTimeSlicePropertyTypeUnmarshaller struct{}

func (u *IntermediateLegTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(IntermediateLegTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *IntermediateLegTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*IntermediateLegTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "IntermediateLegTimeSlice":
		GongUnmarshallPointer(&instance.IntermediateLegTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type IntermediateLegTimeSliceTypeUnmarshaller struct{}

func (u *IntermediateLegTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(IntermediateLegTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *IntermediateLegTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*IntermediateLegTimeSliceType)
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
	case "StartPoint":
		GongUnmarshallPointer(&instance.StartPoint, valueExpr, identifierMap)
	case "EndPoint":
		GongUnmarshallPointer(&instance.EndPoint, valueExpr, identifierMap)
	case "Trajectory":
		GongUnmarshallPointer(&instance.Trajectory, valueExpr, identifierMap)
	case "ArcCentre":
		GongUnmarshallPointer(&instance.ArcCentre, valueExpr, identifierMap)
	case "Angle":
		GongUnmarshallPointer(&instance.Angle, valueExpr, identifierMap)
	case "Distance":
		GongUnmarshallPointer(&instance.Distance, valueExpr, identifierMap)
	case "AircraftCategory":
		GongUnmarshallSliceOfPointers(&instance.AircraftCategory, valueExpr, identifierMap)
	case "Holding":
		GongUnmarshallPointer(&instance.Holding, valueExpr, identifierMap)
	case "DesignSurface":
		GongUnmarshallSliceOfPointers(&instance.DesignSurface, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Approach":
		GongUnmarshallPointer(&instance.Approach, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type IntermediateLegTypeUnmarshaller struct{}

func (u *IntermediateLegTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(IntermediateLegType), stage, identifier, instanceName, preserveOrder)
}

func (u *IntermediateLegTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*IntermediateLegType)
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

type LandingTakeoffAreaCollectionPropertyTypeUnmarshaller struct{}

func (u *LandingTakeoffAreaCollectionPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LandingTakeoffAreaCollectionPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *LandingTakeoffAreaCollectionPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LandingTakeoffAreaCollectionPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "LandingTakeoffAreaCollection":
		GongUnmarshallPointer(&instance.LandingTakeoffAreaCollection, valueExpr, identifierMap)
	}
	return nil
}

type LandingTakeoffAreaCollectionTypeUnmarshaller struct{}

func (u *LandingTakeoffAreaCollectionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LandingTakeoffAreaCollectionType), stage, identifier, instanceName, preserveOrder)
}

func (u *LandingTakeoffAreaCollectionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LandingTakeoffAreaCollectionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Runway":
		GongUnmarshallSliceOfPointers(&instance.Runway, valueExpr, identifierMap)
	case "TLOF":
		GongUnmarshallSliceOfPointers(&instance.TLOF, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
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

type MinimaPropertyTypeUnmarshaller struct{}

func (u *MinimaPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MinimaPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *MinimaPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MinimaPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Minima":
		GongUnmarshallPointer(&instance.Minima, valueExpr, identifierMap)
	}
	return nil
}

type MinimaTypeUnmarshaller struct{}

func (u *MinimaTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MinimaType), stage, identifier, instanceName, preserveOrder)
}

func (u *MinimaTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MinimaType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "AdjustmentINOP":
		GongUnmarshallSliceOfPointers(&instance.AdjustmentINOP, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type MissedApproachGroupPropertyTypeUnmarshaller struct{}

func (u *MissedApproachGroupPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MissedApproachGroupPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *MissedApproachGroupPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MissedApproachGroupPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "MissedApproachGroup":
		GongUnmarshallPointer(&instance.MissedApproachGroup, valueExpr, identifierMap)
	}
	return nil
}

type MissedApproachGroupTypeUnmarshaller struct{}

func (u *MissedApproachGroupTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MissedApproachGroupType), stage, identifier, instanceName, preserveOrder)
}

func (u *MissedApproachGroupTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MissedApproachGroupType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Altimeter":
		GongUnmarshallSliceOfPointers(&instance.Altimeter, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type MissedApproachLegTimeSlicePropertyTypeUnmarshaller struct{}

func (u *MissedApproachLegTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MissedApproachLegTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *MissedApproachLegTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MissedApproachLegTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "MissedApproachLegTimeSlice":
		GongUnmarshallPointer(&instance.MissedApproachLegTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type MissedApproachLegTimeSliceTypeUnmarshaller struct{}

func (u *MissedApproachLegTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MissedApproachLegTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *MissedApproachLegTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MissedApproachLegTimeSliceType)
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
	case "StartPoint":
		GongUnmarshallPointer(&instance.StartPoint, valueExpr, identifierMap)
	case "EndPoint":
		GongUnmarshallPointer(&instance.EndPoint, valueExpr, identifierMap)
	case "Trajectory":
		GongUnmarshallPointer(&instance.Trajectory, valueExpr, identifierMap)
	case "ArcCentre":
		GongUnmarshallPointer(&instance.ArcCentre, valueExpr, identifierMap)
	case "Angle":
		GongUnmarshallPointer(&instance.Angle, valueExpr, identifierMap)
	case "Distance":
		GongUnmarshallPointer(&instance.Distance, valueExpr, identifierMap)
	case "AircraftCategory":
		GongUnmarshallSliceOfPointers(&instance.AircraftCategory, valueExpr, identifierMap)
	case "Holding":
		GongUnmarshallPointer(&instance.Holding, valueExpr, identifierMap)
	case "DesignSurface":
		GongUnmarshallSliceOfPointers(&instance.DesignSurface, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Approach":
		GongUnmarshallPointer(&instance.Approach, valueExpr, identifierMap)
	case "Condition":
		GongUnmarshallSliceOfPointers(&instance.Condition, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type MissedApproachLegTypeUnmarshaller struct{}

func (u *MissedApproachLegTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MissedApproachLegType), stage, identifier, instanceName, preserveOrder)
}

func (u *MissedApproachLegTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MissedApproachLegType)
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

type NavaidPropertyTypeUnmarshaller struct{}

func (u *NavaidPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NavaidPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *NavaidPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NavaidPropertyType)
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

type NavigationAreaRestrictionTimeSlicePropertyTypeUnmarshaller struct{}

func (u *NavigationAreaRestrictionTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NavigationAreaRestrictionTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *NavigationAreaRestrictionTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NavigationAreaRestrictionTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NavigationAreaRestrictionTimeSlice":
		GongUnmarshallPointer(&instance.NavigationAreaRestrictionTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type NavigationAreaRestrictionTimeSliceTypeUnmarshaller struct{}

func (u *NavigationAreaRestrictionTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NavigationAreaRestrictionTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *NavigationAreaRestrictionTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NavigationAreaRestrictionTimeSliceType)
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
	case "Procedure":
		GongUnmarshallSliceOfPointers(&instance.Procedure, valueExpr, identifierMap)
	case "DesignSurface":
		GongUnmarshallPointer(&instance.DesignSurface, valueExpr, identifierMap)
	case "SectorDefinition":
		GongUnmarshallPointer(&instance.SectorDefinition, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type NavigationAreaRestrictionTypeUnmarshaller struct{}

func (u *NavigationAreaRestrictionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NavigationAreaRestrictionType), stage, identifier, instanceName, preserveOrder)
}

func (u *NavigationAreaRestrictionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NavigationAreaRestrictionType)
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

type NavigationAreaSectorPropertyTypeUnmarshaller struct{}

func (u *NavigationAreaSectorPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NavigationAreaSectorPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *NavigationAreaSectorPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NavigationAreaSectorPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "NavigationAreaSector":
		GongUnmarshallPointer(&instance.NavigationAreaSector, valueExpr, identifierMap)
	}
	return nil
}

type NavigationAreaSectorTypeUnmarshaller struct{}

func (u *NavigationAreaSectorTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NavigationAreaSectorType), stage, identifier, instanceName, preserveOrder)
}

func (u *NavigationAreaSectorTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NavigationAreaSectorType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "SectorDefinition":
		GongUnmarshallPointer(&instance.SectorDefinition, valueExpr, identifierMap)
	case "SignificantObstacle":
		GongUnmarshallSliceOfPointers(&instance.SignificantObstacle, valueExpr, identifierMap)
	case "Extent":
		GongUnmarshallPointer(&instance.Extent, valueExpr, identifierMap)
	case "SectorCriteria":
		GongUnmarshallSliceOfPointers(&instance.SectorCriteria, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type NavigationAreaTimeSlicePropertyTypeUnmarshaller struct{}

func (u *NavigationAreaTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NavigationAreaTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *NavigationAreaTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NavigationAreaTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NavigationAreaTimeSlice":
		GongUnmarshallPointer(&instance.NavigationAreaTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type NavigationAreaTimeSliceTypeUnmarshaller struct{}

func (u *NavigationAreaTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NavigationAreaTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *NavigationAreaTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NavigationAreaTimeSliceType)
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
	case "Departure":
		GongUnmarshallPointer(&instance.Departure, valueExpr, identifierMap)
	case "Sector":
		GongUnmarshallSliceOfPointers(&instance.Sector, valueExpr, identifierMap)
	case "CentrePoint_fixDesignatedPoint":
		GongUnmarshallPointer(&instance.CentrePoint_fixDesignatedPoint, valueExpr, identifierMap)
	case "CentrePoint_navaidSystem":
		GongUnmarshallPointer(&instance.CentrePoint_navaidSystem, valueExpr, identifierMap)
	case "CentrePoint_position":
		GongUnmarshallPointer(&instance.CentrePoint_position, valueExpr, identifierMap)
	case "CentrePoint_runwayPoint":
		GongUnmarshallPointer(&instance.CentrePoint_runwayPoint, valueExpr, identifierMap)
	case "CentrePoint_aimingPoint":
		GongUnmarshallPointer(&instance.CentrePoint_aimingPoint, valueExpr, identifierMap)
	case "CentrePoint_airportReferencePoint":
		GongUnmarshallPointer(&instance.CentrePoint_airportReferencePoint, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type NavigationAreaTypeUnmarshaller struct{}

func (u *NavigationAreaTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NavigationAreaType), stage, identifier, instanceName, preserveOrder)
}

func (u *NavigationAreaTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NavigationAreaType)
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

type ObstacleAssessmentAreaPropertyTypeUnmarshaller struct{}

func (u *ObstacleAssessmentAreaPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ObstacleAssessmentAreaPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ObstacleAssessmentAreaPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ObstacleAssessmentAreaPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "ObstacleAssessmentArea":
		GongUnmarshallPointer(&instance.ObstacleAssessmentArea, valueExpr, identifierMap)
	}
	return nil
}

type ObstacleAssessmentAreaTypeUnmarshaller struct{}

func (u *ObstacleAssessmentAreaTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ObstacleAssessmentAreaType), stage, identifier, instanceName, preserveOrder)
}

func (u *ObstacleAssessmentAreaTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ObstacleAssessmentAreaType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "AircraftCategory":
		GongUnmarshallSliceOfPointers(&instance.AircraftCategory, valueExpr, identifierMap)
	case "SignificantObstacle":
		GongUnmarshallSliceOfPointers(&instance.SignificantObstacle, valueExpr, identifierMap)
	case "Surface":
		GongUnmarshallPointer(&instance.Surface, valueExpr, identifierMap)
	case "StartingCurve":
		GongUnmarshallPointer(&instance.StartingCurve, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type ObstaclePlacementPropertyTypeUnmarshaller struct{}

func (u *ObstaclePlacementPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ObstaclePlacementPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ObstaclePlacementPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ObstaclePlacementPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "ObstaclePlacement":
		GongUnmarshallPointer(&instance.ObstaclePlacement, valueExpr, identifierMap)
	}
	return nil
}

type ObstaclePlacementTypeUnmarshaller struct{}

func (u *ObstaclePlacementTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ObstaclePlacementType), stage, identifier, instanceName, preserveOrder)
}

func (u *ObstaclePlacementTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ObstaclePlacementType)
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

type ObstructionPropertyTypeUnmarshaller struct{}

func (u *ObstructionPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ObstructionPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ObstructionPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ObstructionPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Obstruction":
		GongUnmarshallPointer(&instance.Obstruction, valueExpr, identifierMap)
	}
	return nil
}

type ObstructionTypeUnmarshaller struct{}

func (u *ObstructionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ObstructionType), stage, identifier, instanceName, preserveOrder)
}

func (u *ObstructionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ObstructionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TheVerticalStructure":
		GongUnmarshallPointer(&instance.TheVerticalStructure, valueExpr, identifierMap)
	case "Adjustment":
		GongUnmarshallSliceOfPointers(&instance.Adjustment, valueExpr, identifierMap)
	case "ObstaclePlacement":
		GongUnmarshallSliceOfPointers(&instance.ObstaclePlacement, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
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
	case "Point":
		GongUnmarshallPointer(&instance.Point, valueExpr, identifierMap)
	case "ElevatedPoint":
		GongUnmarshallPointer(&instance.ElevatedPoint, valueExpr, identifierMap)
	}
	return nil
}

type PointReferencePropertyTypeUnmarshaller struct{}

func (u *PointReferencePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PointReferencePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *PointReferencePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PointReferencePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "PointReference":
		GongUnmarshallPointer(&instance.PointReference, valueExpr, identifierMap)
	}
	return nil
}

type PointReferenceTypeUnmarshaller struct{}

func (u *PointReferenceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PointReferenceType), stage, identifier, instanceName, preserveOrder)
}

func (u *PointReferenceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PointReferenceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Point":
		GongUnmarshallPointer(&instance.Point, valueExpr, identifierMap)
	case "FacilityAngle":
		GongUnmarshallSliceOfPointers(&instance.FacilityAngle, valueExpr, identifierMap)
	case "FacilityDistance":
		GongUnmarshallSliceOfPointers(&instance.FacilityDistance, valueExpr, identifierMap)
	case "FixToleranceArea":
		GongUnmarshallPointer(&instance.FixToleranceArea, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
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
	case "SrsName":
		instance.SrsName = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	}
	return nil
}

type ProcedureAvailabilityPropertyTypeUnmarshaller struct{}

func (u *ProcedureAvailabilityPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ProcedureAvailabilityPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ProcedureAvailabilityPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ProcedureAvailabilityPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "ProcedureAvailability":
		GongUnmarshallPointer(&instance.ProcedureAvailability, valueExpr, identifierMap)
	}
	return nil
}

type ProcedureAvailabilityTypeUnmarshaller struct{}

func (u *ProcedureAvailabilityTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ProcedureAvailabilityType), stage, identifier, instanceName, preserveOrder)
}

func (u *ProcedureAvailabilityTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ProcedureAvailabilityType)
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
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type ProcedureDMETimeSlicePropertyTypeUnmarshaller struct{}

func (u *ProcedureDMETimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ProcedureDMETimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ProcedureDMETimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ProcedureDMETimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "ProcedureDMETimeSlice":
		GongUnmarshallPointer(&instance.ProcedureDMETimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type ProcedureDMETimeSliceTypeUnmarshaller struct{}

func (u *ProcedureDMETimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ProcedureDMETimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *ProcedureDMETimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ProcedureDMETimeSliceType)
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
	case "DME":
		GongUnmarshallPointer(&instance.DME, valueExpr, identifierMap)
	case "SegmentLeg":
		GongUnmarshallPointer(&instance.SegmentLeg, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type ProcedureDMETypeUnmarshaller struct{}

func (u *ProcedureDMETypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ProcedureDMEType), stage, identifier, instanceName, preserveOrder)
}

func (u *ProcedureDMETypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ProcedureDMEType)
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

type ProcedurePropertyTypeUnmarshaller struct{}

func (u *ProcedurePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ProcedurePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ProcedurePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ProcedurePropertyType)
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

type ProcedureTransitionLegPropertyTypeUnmarshaller struct{}

func (u *ProcedureTransitionLegPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ProcedureTransitionLegPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ProcedureTransitionLegPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ProcedureTransitionLegPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "ProcedureTransitionLeg":
		GongUnmarshallPointer(&instance.ProcedureTransitionLeg, valueExpr, identifierMap)
	}
	return nil
}

type ProcedureTransitionLegTypeUnmarshaller struct{}

func (u *ProcedureTransitionLegTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ProcedureTransitionLegType), stage, identifier, instanceName, preserveOrder)
}

func (u *ProcedureTransitionLegTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ProcedureTransitionLegType)
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
	case "TheSegmentLeg":
		GongUnmarshallPointer(&instance.TheSegmentLeg, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type ProcedureTransitionPropertyTypeUnmarshaller struct{}

func (u *ProcedureTransitionPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ProcedureTransitionPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ProcedureTransitionPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ProcedureTransitionPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "ProcedureTransition":
		GongUnmarshallPointer(&instance.ProcedureTransition, valueExpr, identifierMap)
	}
	return nil
}

type ProcedureTransitionTypeUnmarshaller struct{}

func (u *ProcedureTransitionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ProcedureTransitionType), stage, identifier, instanceName, preserveOrder)
}

func (u *ProcedureTransitionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ProcedureTransitionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "DepartureRunwayTransition":
		GongUnmarshallPointer(&instance.DepartureRunwayTransition, valueExpr, identifierMap)
	case "Trajectory":
		GongUnmarshallPointer(&instance.Trajectory, valueExpr, identifierMap)
	case "TransitionLeg":
		GongUnmarshallSliceOfPointers(&instance.TransitionLeg, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type RadarSystemPropertyTypeUnmarshaller struct{}

func (u *RadarSystemPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RadarSystemPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RadarSystemPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RadarSystemPropertyType)
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

type RadioFrequencyAreaPropertyTypeUnmarshaller struct{}

func (u *RadioFrequencyAreaPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RadioFrequencyAreaPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RadioFrequencyAreaPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RadioFrequencyAreaPropertyType)
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

type RunwayCentrelinePointPropertyTypeUnmarshaller struct{}

func (u *RunwayCentrelinePointPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayCentrelinePointPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayCentrelinePointPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayCentrelinePointPropertyType)
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

type RunwayDirectionPropertyTypeUnmarshaller struct{}

func (u *RunwayDirectionPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayDirectionPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayDirectionPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayDirectionPropertyType)
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

type SafeAltitudeAreaPropertyTypeUnmarshaller struct{}

func (u *SafeAltitudeAreaPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SafeAltitudeAreaPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *SafeAltitudeAreaPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SafeAltitudeAreaPropertyType)
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

type SafeAltitudeAreaSectorPropertyTypeUnmarshaller struct{}

func (u *SafeAltitudeAreaSectorPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SafeAltitudeAreaSectorPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *SafeAltitudeAreaSectorPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SafeAltitudeAreaSectorPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "SafeAltitudeAreaSector":
		GongUnmarshallPointer(&instance.SafeAltitudeAreaSector, valueExpr, identifierMap)
	}
	return nil
}

type SafeAltitudeAreaSectorTypeUnmarshaller struct{}

func (u *SafeAltitudeAreaSectorTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SafeAltitudeAreaSectorType), stage, identifier, instanceName, preserveOrder)
}

func (u *SafeAltitudeAreaSectorTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SafeAltitudeAreaSectorType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Extent":
		GongUnmarshallPointer(&instance.Extent, valueExpr, identifierMap)
	case "SignificantObstacle":
		GongUnmarshallSliceOfPointers(&instance.SignificantObstacle, valueExpr, identifierMap)
	case "SectorDefinition":
		GongUnmarshallPointer(&instance.SectorDefinition, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type SafeAltitudeAreaTimeSlicePropertyTypeUnmarshaller struct{}

func (u *SafeAltitudeAreaTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SafeAltitudeAreaTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *SafeAltitudeAreaTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SafeAltitudeAreaTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "SafeAltitudeAreaTimeSlice":
		GongUnmarshallPointer(&instance.SafeAltitudeAreaTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type SafeAltitudeAreaTimeSliceTypeUnmarshaller struct{}

func (u *SafeAltitudeAreaTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SafeAltitudeAreaTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *SafeAltitudeAreaTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SafeAltitudeAreaTimeSliceType)
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
	case "CentrePoint_fixDesignatedPoint":
		GongUnmarshallPointer(&instance.CentrePoint_fixDesignatedPoint, valueExpr, identifierMap)
	case "CentrePoint_navaidSystem":
		GongUnmarshallPointer(&instance.CentrePoint_navaidSystem, valueExpr, identifierMap)
	case "CentrePoint_position":
		GongUnmarshallPointer(&instance.CentrePoint_position, valueExpr, identifierMap)
	case "CentrePoint_runwayPoint":
		GongUnmarshallPointer(&instance.CentrePoint_runwayPoint, valueExpr, identifierMap)
	case "CentrePoint_aimingPoint":
		GongUnmarshallPointer(&instance.CentrePoint_aimingPoint, valueExpr, identifierMap)
	case "CentrePoint_airportReferencePoint":
		GongUnmarshallPointer(&instance.CentrePoint_airportReferencePoint, valueExpr, identifierMap)
	case "Sector":
		GongUnmarshallSliceOfPointers(&instance.Sector, valueExpr, identifierMap)
	case "Location":
		GongUnmarshallSliceOfPointers(&instance.Location, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type SafeAltitudeAreaTypeUnmarshaller struct{}

func (u *SafeAltitudeAreaTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SafeAltitudeAreaType), stage, identifier, instanceName, preserveOrder)
}

func (u *SafeAltitudeAreaTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SafeAltitudeAreaType)
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

type SectorDesignPropertyTypeUnmarshaller struct{}

func (u *SectorDesignPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SectorDesignPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *SectorDesignPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SectorDesignPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "SectorDesign":
		GongUnmarshallPointer(&instance.SectorDesign, valueExpr, identifierMap)
	}
	return nil
}

type SectorDesignTypeUnmarshaller struct{}

func (u *SectorDesignTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SectorDesignType), stage, identifier, instanceName, preserveOrder)
}

func (u *SectorDesignTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SectorDesignType)
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

type SegmentLegPropertyTypeUnmarshaller struct{}

func (u *SegmentLegPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SegmentLegPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *SegmentLegPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SegmentLegPropertyType)
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

type SegmentPointPropertyTypeUnmarshaller struct{}

func (u *SegmentPointPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SegmentPointPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *SegmentPointPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SegmentPointPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "AerialRefuellingPoint":
		GongUnmarshallPointer(&instance.AerialRefuellingPoint, valueExpr, identifierMap)
	case "EnRouteSegmentPoint":
		GongUnmarshallPointer(&instance.EnRouteSegmentPoint, valueExpr, identifierMap)
	case "TerminalSegmentPoint":
		GongUnmarshallPointer(&instance.TerminalSegmentPoint, valueExpr, identifierMap)
	}
	return nil
}

type SpecialNavigationSystemPropertyTypeUnmarshaller struct{}

func (u *SpecialNavigationSystemPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SpecialNavigationSystemPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *SpecialNavigationSystemPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SpecialNavigationSystemPropertyType)
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

type StandardInstrumentArrivalPropertyTypeUnmarshaller struct{}

func (u *StandardInstrumentArrivalPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StandardInstrumentArrivalPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *StandardInstrumentArrivalPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StandardInstrumentArrivalPropertyType)
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

type StandardInstrumentArrivalTimeSlicePropertyTypeUnmarshaller struct{}

func (u *StandardInstrumentArrivalTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StandardInstrumentArrivalTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *StandardInstrumentArrivalTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StandardInstrumentArrivalTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "StandardInstrumentArrivalTimeSlice":
		GongUnmarshallPointer(&instance.StandardInstrumentArrivalTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type StandardInstrumentArrivalTimeSliceTypeUnmarshaller struct{}

func (u *StandardInstrumentArrivalTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StandardInstrumentArrivalTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *StandardInstrumentArrivalTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StandardInstrumentArrivalTimeSliceType)
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
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "AirportHeliport":
		GongUnmarshallSliceOfPointers(&instance.AirportHeliport, valueExpr, identifierMap)
	case "AircraftCharacteristic":
		GongUnmarshallSliceOfPointers(&instance.AircraftCharacteristic, valueExpr, identifierMap)
	case "FlightTransition":
		GongUnmarshallSliceOfPointers(&instance.FlightTransition, valueExpr, identifierMap)
	case "GuidanceFacility_specialNavigationSystem":
		GongUnmarshallPointer(&instance.GuidanceFacility_specialNavigationSystem, valueExpr, identifierMap)
	case "GuidanceFacility_radar":
		GongUnmarshallPointer(&instance.GuidanceFacility_radar, valueExpr, identifierMap)
	case "GuidanceFacility_navaid":
		GongUnmarshallPointer(&instance.GuidanceFacility_navaid, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "SafeAltitude":
		GongUnmarshallPointer(&instance.SafeAltitude, valueExpr, identifierMap)
	case "Arrival":
		GongUnmarshallPointer(&instance.Arrival, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type StandardInstrumentArrivalTypeUnmarshaller struct{}

func (u *StandardInstrumentArrivalTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StandardInstrumentArrivalType), stage, identifier, instanceName, preserveOrder)
}

func (u *StandardInstrumentArrivalTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StandardInstrumentArrivalType)
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

type StandardInstrumentDeparturePropertyTypeUnmarshaller struct{}

func (u *StandardInstrumentDeparturePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StandardInstrumentDeparturePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *StandardInstrumentDeparturePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StandardInstrumentDeparturePropertyType)
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

type StandardInstrumentDepartureTimeSlicePropertyTypeUnmarshaller struct{}

func (u *StandardInstrumentDepartureTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StandardInstrumentDepartureTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *StandardInstrumentDepartureTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StandardInstrumentDepartureTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "StandardInstrumentDepartureTimeSlice":
		GongUnmarshallPointer(&instance.StandardInstrumentDepartureTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type StandardInstrumentDepartureTimeSliceTypeUnmarshaller struct{}

func (u *StandardInstrumentDepartureTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StandardInstrumentDepartureTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *StandardInstrumentDepartureTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StandardInstrumentDepartureTimeSliceType)
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
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "AirportHeliport":
		GongUnmarshallSliceOfPointers(&instance.AirportHeliport, valueExpr, identifierMap)
	case "AircraftCharacteristic":
		GongUnmarshallSliceOfPointers(&instance.AircraftCharacteristic, valueExpr, identifierMap)
	case "FlightTransition":
		GongUnmarshallSliceOfPointers(&instance.FlightTransition, valueExpr, identifierMap)
	case "GuidanceFacility_specialNavigationSystem":
		GongUnmarshallPointer(&instance.GuidanceFacility_specialNavigationSystem, valueExpr, identifierMap)
	case "GuidanceFacility_radar":
		GongUnmarshallPointer(&instance.GuidanceFacility_radar, valueExpr, identifierMap)
	case "GuidanceFacility_navaid":
		GongUnmarshallPointer(&instance.GuidanceFacility_navaid, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "SafeAltitude":
		GongUnmarshallPointer(&instance.SafeAltitude, valueExpr, identifierMap)
	case "Takeoff":
		GongUnmarshallPointer(&instance.Takeoff, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type StandardInstrumentDepartureTypeUnmarshaller struct{}

func (u *StandardInstrumentDepartureTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StandardInstrumentDepartureType), stage, identifier, instanceName, preserveOrder)
}

func (u *StandardInstrumentDepartureTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StandardInstrumentDepartureType)
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

type TerminalArrivalAreaSectorPropertyTypeUnmarshaller struct{}

func (u *TerminalArrivalAreaSectorPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TerminalArrivalAreaSectorPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *TerminalArrivalAreaSectorPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TerminalArrivalAreaSectorPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "TerminalArrivalAreaSector":
		GongUnmarshallPointer(&instance.TerminalArrivalAreaSector, valueExpr, identifierMap)
	}
	return nil
}

type TerminalArrivalAreaSectorTypeUnmarshaller struct{}

func (u *TerminalArrivalAreaSectorTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TerminalArrivalAreaSectorType), stage, identifier, instanceName, preserveOrder)
}

func (u *TerminalArrivalAreaSectorTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TerminalArrivalAreaSectorType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "SectorDefinition":
		GongUnmarshallPointer(&instance.SectorDefinition, valueExpr, identifierMap)
	case "Extent":
		GongUnmarshallPointer(&instance.Extent, valueExpr, identifierMap)
	case "SignificantObstacle":
		GongUnmarshallSliceOfPointers(&instance.SignificantObstacle, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type TerminalArrivalAreaTimeSlicePropertyTypeUnmarshaller struct{}

func (u *TerminalArrivalAreaTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TerminalArrivalAreaTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *TerminalArrivalAreaTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TerminalArrivalAreaTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "TerminalArrivalAreaTimeSlice":
		GongUnmarshallPointer(&instance.TerminalArrivalAreaTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type TerminalArrivalAreaTimeSliceTypeUnmarshaller struct{}

func (u *TerminalArrivalAreaTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TerminalArrivalAreaTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *TerminalArrivalAreaTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TerminalArrivalAreaTimeSliceType)
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
	case "IF_fixDesignatedPoint":
		GongUnmarshallPointer(&instance.IF_fixDesignatedPoint, valueExpr, identifierMap)
	case "IF_navaidSystem":
		GongUnmarshallPointer(&instance.IF_navaidSystem, valueExpr, identifierMap)
	case "IF_position":
		GongUnmarshallPointer(&instance.IF_position, valueExpr, identifierMap)
	case "IF_runwayPoint":
		GongUnmarshallPointer(&instance.IF_runwayPoint, valueExpr, identifierMap)
	case "IF_aimingPoint":
		GongUnmarshallPointer(&instance.IF_aimingPoint, valueExpr, identifierMap)
	case "IF_airportReferencePoint":
		GongUnmarshallPointer(&instance.IF_airportReferencePoint, valueExpr, identifierMap)
	case "IAF_fixDesignatedPoint":
		GongUnmarshallPointer(&instance.IAF_fixDesignatedPoint, valueExpr, identifierMap)
	case "IAF_navaidSystem":
		GongUnmarshallPointer(&instance.IAF_navaidSystem, valueExpr, identifierMap)
	case "IAF_position":
		GongUnmarshallPointer(&instance.IAF_position, valueExpr, identifierMap)
	case "IAF_runwayPoint":
		GongUnmarshallPointer(&instance.IAF_runwayPoint, valueExpr, identifierMap)
	case "IAF_aimingPoint":
		GongUnmarshallPointer(&instance.IAF_aimingPoint, valueExpr, identifierMap)
	case "IAF_airportReferencePoint":
		GongUnmarshallPointer(&instance.IAF_airportReferencePoint, valueExpr, identifierMap)
	case "Buffer":
		GongUnmarshallPointer(&instance.Buffer, valueExpr, identifierMap)
	case "Sector":
		GongUnmarshallSliceOfPointers(&instance.Sector, valueExpr, identifierMap)
	case "ApproachRNAV":
		GongUnmarshallPointer(&instance.ApproachRNAV, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type TerminalArrivalAreaTypeUnmarshaller struct{}

func (u *TerminalArrivalAreaTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TerminalArrivalAreaType), stage, identifier, instanceName, preserveOrder)
}

func (u *TerminalArrivalAreaTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TerminalArrivalAreaType)
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

type TerminalSegmentPointPropertyTypeUnmarshaller struct{}

func (u *TerminalSegmentPointPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TerminalSegmentPointPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *TerminalSegmentPointPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TerminalSegmentPointPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "TerminalSegmentPoint":
		GongUnmarshallPointer(&instance.TerminalSegmentPoint, valueExpr, identifierMap)
	}
	return nil
}

type TerminalSegmentPointTypeUnmarshaller struct{}

func (u *TerminalSegmentPointTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TerminalSegmentPointType), stage, identifier, instanceName, preserveOrder)
}

func (u *TerminalSegmentPointTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TerminalSegmentPointType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "FacilityMakeup":
		GongUnmarshallSliceOfPointers(&instance.FacilityMakeup, valueExpr, identifierMap)
	case "PointChoice_fixDesignatedPoint":
		GongUnmarshallPointer(&instance.PointChoice_fixDesignatedPoint, valueExpr, identifierMap)
	case "PointChoice_navaidSystem":
		GongUnmarshallPointer(&instance.PointChoice_navaidSystem, valueExpr, identifierMap)
	case "PointChoice_position":
		GongUnmarshallPointer(&instance.PointChoice_position, valueExpr, identifierMap)
	case "PointChoice_runwayPoint":
		GongUnmarshallPointer(&instance.PointChoice_runwayPoint, valueExpr, identifierMap)
	case "PointChoice_aimingPoint":
		GongUnmarshallPointer(&instance.PointChoice_aimingPoint, valueExpr, identifierMap)
	case "PointChoice_airportReferencePoint":
		GongUnmarshallPointer(&instance.PointChoice_airportReferencePoint, valueExpr, identifierMap)
	case "ExtendedServiceVolume":
		GongUnmarshallPointer(&instance.ExtendedServiceVolume, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
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

type TouchDownLiftOffPropertyTypeUnmarshaller struct{}

func (u *TouchDownLiftOffPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TouchDownLiftOffPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *TouchDownLiftOffPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TouchDownLiftOffPropertyType)
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

type UnplannedHoldingPropertyTypeUnmarshaller struct{}

func (u *UnplannedHoldingPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(UnplannedHoldingPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *UnplannedHoldingPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*UnplannedHoldingPropertyType)
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

type UnplannedHoldingTimeSlicePropertyTypeUnmarshaller struct{}

func (u *UnplannedHoldingTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(UnplannedHoldingTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *UnplannedHoldingTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*UnplannedHoldingTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "UnplannedHoldingTimeSlice":
		GongUnmarshallPointer(&instance.UnplannedHoldingTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type UnplannedHoldingTimeSliceTypeUnmarshaller struct{}

func (u *UnplannedHoldingTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(UnplannedHoldingTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *UnplannedHoldingTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*UnplannedHoldingTimeSliceType)
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
	case "HoldingPoint":
		GongUnmarshallPointer(&instance.HoldingPoint, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type UnplannedHoldingTypeUnmarshaller struct{}

func (u *UnplannedHoldingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(UnplannedHoldingType), stage, identifier, instanceName, preserveOrder)
}

func (u *UnplannedHoldingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*UnplannedHoldingType)
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

type VerticalStructurePropertyTypeUnmarshaller struct{}

func (u *VerticalStructurePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(VerticalStructurePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *VerticalStructurePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*VerticalStructurePropertyType)
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
