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
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmfeatures":
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
				case "AbstractDirectFlightType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractDirectFlightType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractPropertiesWithScheduleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractPropertiesWithScheduleType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractSegmentPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractSegmentPointType), stage, ident.Name, instanceName, preserveOrder)
				case "AerialRefuellingAnchorPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AerialRefuellingAnchorPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AerialRefuellingAnchorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AerialRefuellingAnchorType), stage, ident.Name, instanceName, preserveOrder)
				case "AerialRefuellingPointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AerialRefuellingPointPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AerialRefuellingPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AerialRefuellingPointType), stage, ident.Name, instanceName, preserveOrder)
				case "AerialRefuellingPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AerialRefuellingPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AerialRefuellingTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AerialRefuellingTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AerialRefuellingTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AerialRefuellingTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "AerialRefuellingTrackPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AerialRefuellingTrackPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AerialRefuellingTrackType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AerialRefuellingTrackType), stage, ident.Name, instanceName, preserveOrder)
				case "AerialRefuellingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AerialRefuellingType), stage, ident.Name, instanceName, preserveOrder)
				case "AircraftCharacteristicPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AircraftCharacteristicPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AircraftCharacteristicType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AircraftCharacteristicType), stage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirportHeliportPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AirspaceBorderCrossingPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirspaceBorderCrossingPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AirspaceBorderCrossingTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirspaceBorderCrossingTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AirspaceBorderCrossingTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirspaceBorderCrossingTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "AirspaceBorderCrossingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirspaceBorderCrossingType), stage, ident.Name, instanceName, preserveOrder)
				case "AirspaceLayerPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirspaceLayerPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AirspaceLayerType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirspaceLayerType), stage, ident.Name, instanceName, preserveOrder)
				case "AirspacePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirspacePropertyType), stage, ident.Name, instanceName, preserveOrder)
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
				case "AuthorityForAerialRefuellingPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AuthorityForAerialRefuellingPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AuthorityForAerialRefuellingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AuthorityForAerialRefuellingType), stage, ident.Name, instanceName, preserveOrder)
				case "ChangeOverPointTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ChangeOverPointTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ChangeOverPointTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ChangeOverPointTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "ChangeOverPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ChangeOverPointType), stage, ident.Name, instanceName, preserveOrder)
				case "ContactInformationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ContactInformationPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ContactInformationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ContactInformationType), stage, ident.Name, instanceName, preserveOrder)
				case "CurvePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CurvePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "CurveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CurveType), stage, ident.Name, instanceName, preserveOrder)
				case "DMEPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DMEPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "DesignatedPointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DesignatedPointPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "DirectFlightClassType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DirectFlightClassType), stage, ident.Name, instanceName, preserveOrder)
				case "DirectFlightPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DirectFlightPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "DirectFlightSegmentPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DirectFlightSegmentPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "DirectFlightSegmentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DirectFlightSegmentType), stage, ident.Name, instanceName, preserveOrder)
				case "DistanceIndicationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DistanceIndicationPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ElevatedCurveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ElevatedCurveType), stage, ident.Name, instanceName, preserveOrder)
				case "ElevatedPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ElevatedPointType), stage, ident.Name, instanceName, preserveOrder)
				case "ElevatedSurfaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ElevatedSurfaceType), stage, ident.Name, instanceName, preserveOrder)
				case "EnRouteSegmentPointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(EnRouteSegmentPointPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "EnRouteSegmentPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(EnRouteSegmentPointType), stage, ident.Name, instanceName, preserveOrder)
				case "FlightCharacteristicPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FlightCharacteristicPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "FlightCharacteristicType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FlightCharacteristicType), stage, ident.Name, instanceName, preserveOrder)
				case "FlightConditionCircumstancePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FlightConditionCircumstancePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "FlightConditionCircumstanceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FlightConditionCircumstanceType), stage, ident.Name, instanceName, preserveOrder)
				case "FlightConditionCombinationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FlightConditionCombinationPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "FlightConditionCombinationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FlightConditionCombinationType), stage, ident.Name, instanceName, preserveOrder)
				case "FlightConditionElementPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FlightConditionElementPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "FlightConditionElementType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FlightConditionElementType), stage, ident.Name, instanceName, preserveOrder)
				case "FlightRestrictionLevelPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FlightRestrictionLevelPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "FlightRestrictionLevelType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FlightRestrictionLevelType), stage, ident.Name, instanceName, preserveOrder)
				case "FlightRestrictionRoutePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FlightRestrictionRoutePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "FlightRestrictionRouteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FlightRestrictionRouteType), stage, ident.Name, instanceName, preserveOrder)
				case "FlightRestrictionTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FlightRestrictionTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "FlightRestrictionTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FlightRestrictionTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "FlightRestrictionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FlightRestrictionType), stage, ident.Name, instanceName, preserveOrder)
				case "FlightRoutingElementPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FlightRoutingElementPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "FlightRoutingElementType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FlightRoutingElementType), stage, ident.Name, instanceName, preserveOrder)
				case "LinguisticNotePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LinguisticNotePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "LinguisticNoteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LinguisticNoteType), stage, ident.Name, instanceName, preserveOrder)
				case "MeteorologyPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MeteorologyPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "MeteorologyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MeteorologyType), stage, ident.Name, instanceName, preserveOrder)
				case "NavaidPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NavaidPropertyType), stage, ident.Name, instanceName, preserveOrder)
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
				case "OnlineContactPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(OnlineContactPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "OnlineContactType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(OnlineContactType), stage, ident.Name, instanceName, preserveOrder)
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
				case "PostalAddressPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PostalAddressPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "PostalAddressType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PostalAddressType), stage, ident.Name, instanceName, preserveOrder)
				case "RadioFrequencyAreaPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RadioFrequencyAreaPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RouteAvailabilityPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RouteAvailabilityPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RouteAvailabilityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RouteAvailabilityType), stage, ident.Name, instanceName, preserveOrder)
				case "RouteDMETimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RouteDMETimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RouteDMETimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RouteDMETimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "RouteDMEType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RouteDMEType), stage, ident.Name, instanceName, preserveOrder)
				case "RoutePortionPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RoutePortionPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RoutePortionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RoutePortionType), stage, ident.Name, instanceName, preserveOrder)
				case "RoutePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RoutePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RouteSegmentTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RouteSegmentTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RouteSegmentTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RouteSegmentTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "RouteSegmentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RouteSegmentType), stage, ident.Name, instanceName, preserveOrder)
				case "RouteTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RouteTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RouteTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RouteTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "RouteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RouteType), stage, ident.Name, instanceName, preserveOrder)
				case "RulesProceduresTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RulesProceduresTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RulesProceduresTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RulesProceduresTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "RulesProceduresType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RulesProceduresType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayCentrelinePointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayCentrelinePointPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "StandardInstrumentArrivalPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StandardInstrumentArrivalPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "StandardInstrumentDeparturePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StandardInstrumentDeparturePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "StandardLevelColumnPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StandardLevelColumnPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SurfacePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SurfacePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SurfaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SurfaceType), stage, ident.Name, instanceName, preserveOrder)
				case "TelephoneContactPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TelephoneContactPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TelephoneContactType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TelephoneContactType), stage, ident.Name, instanceName, preserveOrder)
				case "TimesheetPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TimesheetPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TimesheetType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TimesheetType), stage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TouchDownLiftOffPropertyType), stage, ident.Name, instanceName, preserveOrder)
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
				case *AbstractDirectFlightType:
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
				case *AbstractSegmentPointType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					}
				case *AerialRefuellingAnchorPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AerialRefuellingAnchor":
						__gong__assignPointer(&inst.AerialRefuellingAnchor, rhs, identifierMap)
					}
				case *AerialRefuellingAnchorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "VerticalExtent":
						__gong__assignSliceOfPointers(&inst.VerticalExtent, rhs, identifierMap)
					case "Point":
						__gong__assignSliceOfPointers(&inst.Point, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AerialRefuellingPointPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AerialRefuellingPoint":
						__gong__assignPointer(&inst.AerialRefuellingPoint, rhs, identifierMap)
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
				case *AerialRefuellingPropertyType:
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
				case *AerialRefuellingTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AerialRefuellingTimeSlice":
						__gong__assignPointer(&inst.AerialRefuellingTimeSlice, rhs, identifierMap)
					}
				case *AerialRefuellingTimeSliceType:
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
					case "ProtectingAirspace":
						__gong__assignSliceOfPointers(&inst.ProtectingAirspace, rhs, identifierMap)
					case "Track":
						__gong__assignSliceOfPointers(&inst.Track, rhs, identifierMap)
					case "Anchor":
						__gong__assignSliceOfPointers(&inst.Anchor, rhs, identifierMap)
					case "OppositeTrack":
						__gong__assignPointer(&inst.OppositeTrack, rhs, identifierMap)
					case "ManagingOrganisation":
						__gong__assignSliceOfPointers(&inst.ManagingOrganisation, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AerialRefuellingTrackPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AerialRefuellingTrack":
						__gong__assignPointer(&inst.AerialRefuellingTrack, rhs, identifierMap)
					}
				case *AerialRefuellingTrackType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "Point":
						__gong__assignSliceOfPointers(&inst.Point, rhs, identifierMap)
					case "VerticalExtent":
						__gong__assignSliceOfPointers(&inst.VerticalExtent, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AerialRefuellingType:
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
				case *AirspaceBorderCrossingPropertyType:
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
				case *AirspaceBorderCrossingTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AirspaceBorderCrossingTimeSlice":
						__gong__assignPointer(&inst.AirspaceBorderCrossingTimeSlice, rhs, identifierMap)
					}
				case *AirspaceBorderCrossingTimeSliceType:
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
					case "ExitedAirspace":
						__gong__assignPointer(&inst.ExitedAirspace, rhs, identifierMap)
					case "EnteredAirspace":
						__gong__assignPointer(&inst.EnteredAirspace, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AirspaceBorderCrossingType:
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
				case *AuthorityForAerialRefuellingPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AuthorityForAerialRefuelling":
						__gong__assignPointer(&inst.AuthorityForAerialRefuelling, rhs, identifierMap)
					}
				case *AuthorityForAerialRefuellingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "TheOrganisationAuthority":
						__gong__assignPointer(&inst.TheOrganisationAuthority, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *ChangeOverPointTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "ChangeOverPointTimeSlice":
						__gong__assignPointer(&inst.ChangeOverPointTimeSlice, rhs, identifierMap)
					}
				case *ChangeOverPointTimeSliceType:
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
					case "Location_fixDesignatedPoint":
						__gong__assignPointer(&inst.Location_fixDesignatedPoint, rhs, identifierMap)
					case "Location_navaidSystem":
						__gong__assignPointer(&inst.Location_navaidSystem, rhs, identifierMap)
					case "Location_position":
						__gong__assignPointer(&inst.Location_position, rhs, identifierMap)
					case "Location_runwayPoint":
						__gong__assignPointer(&inst.Location_runwayPoint, rhs, identifierMap)
					case "Location_aimingPoint":
						__gong__assignPointer(&inst.Location_aimingPoint, rhs, identifierMap)
					case "Location_airportReferencePoint":
						__gong__assignPointer(&inst.Location_airportReferencePoint, rhs, identifierMap)
					case "ApplicableRoutePortion":
						__gong__assignPointer(&inst.ApplicableRoutePortion, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *ChangeOverPointType:
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
				case *ContactInformationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ContactInformation":
						__gong__assignPointer(&inst.ContactInformation, rhs, identifierMap)
					}
				case *ContactInformationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Address":
						__gong__assignSliceOfPointers(&inst.Address, rhs, identifierMap)
					case "NetworkNode":
						__gong__assignSliceOfPointers(&inst.NetworkNode, rhs, identifierMap)
					case "PhoneFax":
						__gong__assignSliceOfPointers(&inst.PhoneFax, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
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
				case *DirectFlightClassType:
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
				case *DirectFlightPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "DirectFlightClass":
						__gong__assignPointer(&inst.DirectFlightClass, rhs, identifierMap)
					case "DirectFlightSegment":
						__gong__assignPointer(&inst.DirectFlightSegment, rhs, identifierMap)
					}
				case *DirectFlightSegmentPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "DirectFlightSegment":
						__gong__assignPointer(&inst.DirectFlightSegment, rhs, identifierMap)
					}
				case *DirectFlightSegmentType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "End_fixDesignatedPoint":
						__gong__assignPointer(&inst.End_fixDesignatedPoint, rhs, identifierMap)
					case "End_navaidSystem":
						__gong__assignPointer(&inst.End_navaidSystem, rhs, identifierMap)
					case "End_position":
						__gong__assignPointer(&inst.End_position, rhs, identifierMap)
					case "End_runwayPoint":
						__gong__assignPointer(&inst.End_runwayPoint, rhs, identifierMap)
					case "End_aimingPoint":
						__gong__assignPointer(&inst.End_aimingPoint, rhs, identifierMap)
					case "End_airportReferencePoint":
						__gong__assignPointer(&inst.End_airportReferencePoint, rhs, identifierMap)
					case "Start_fixDesignatedPoint":
						__gong__assignPointer(&inst.Start_fixDesignatedPoint, rhs, identifierMap)
					case "Start_navaidSystem":
						__gong__assignPointer(&inst.Start_navaidSystem, rhs, identifierMap)
					case "Start_position":
						__gong__assignPointer(&inst.Start_position, rhs, identifierMap)
					case "Start_runwayPoint":
						__gong__assignPointer(&inst.Start_runwayPoint, rhs, identifierMap)
					case "Start_aimingPoint":
						__gong__assignPointer(&inst.Start_aimingPoint, rhs, identifierMap)
					case "Start_airportReferencePoint":
						__gong__assignPointer(&inst.Start_airportReferencePoint, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
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
				case *EnRouteSegmentPointPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnRouteSegmentPoint":
						__gong__assignPointer(&inst.EnRouteSegmentPoint, rhs, identifierMap)
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
				case *FlightCharacteristicPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "FlightCharacteristic":
						__gong__assignPointer(&inst.FlightCharacteristic, rhs, identifierMap)
					}
				case *FlightCharacteristicType:
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
				case *FlightConditionCircumstancePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "FlightConditionCircumstance":
						__gong__assignPointer(&inst.FlightConditionCircumstance, rhs, identifierMap)
					}
				case *FlightConditionCircumstanceType:
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
				case *FlightConditionCombinationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "FlightConditionCombination":
						__gong__assignPointer(&inst.FlightConditionCombination, rhs, identifierMap)
					}
				case *FlightConditionCombinationType:
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
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *FlightConditionElementPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "FlightConditionElement":
						__gong__assignPointer(&inst.FlightConditionElement, rhs, identifierMap)
					}
				case *FlightConditionElementType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "FlightCondition_weather":
						__gong__assignPointer(&inst.FlightCondition_weather, rhs, identifierMap)
					case "FlightCondition_aircraft":
						__gong__assignPointer(&inst.FlightCondition_aircraft, rhs, identifierMap)
					case "FlightCondition_flight":
						__gong__assignPointer(&inst.FlightCondition_flight, rhs, identifierMap)
					case "FlightCondition_directFlightCondition":
						__gong__assignPointer(&inst.FlightCondition_directFlightCondition, rhs, identifierMap)
					case "FlightCondition_borderCrossingCondition":
						__gong__assignPointer(&inst.FlightCondition_borderCrossingCondition, rhs, identifierMap)
					case "FlightCondition_operand":
						__gong__assignPointer(&inst.FlightCondition_operand, rhs, identifierMap)
					case "FlightCondition_routePortionCondition":
						__gong__assignPointer(&inst.FlightCondition_routePortionCondition, rhs, identifierMap)
					case "FlightCondition_standardInstrumentDepartureCondition":
						__gong__assignPointer(&inst.FlightCondition_standardInstrumentDepartureCondition, rhs, identifierMap)
					case "FlightCondition_standardInstrumentArrivalCondition":
						__gong__assignPointer(&inst.FlightCondition_standardInstrumentArrivalCondition, rhs, identifierMap)
					case "FlightCondition_organisationCondition":
						__gong__assignPointer(&inst.FlightCondition_organisationCondition, rhs, identifierMap)
					case "SignificantPointCondition_fixDesignatedPoint":
						__gong__assignPointer(&inst.SignificantPointCondition_fixDesignatedPoint, rhs, identifierMap)
					case "SignificantPointCondition_navaidSystem":
						__gong__assignPointer(&inst.SignificantPointCondition_navaidSystem, rhs, identifierMap)
					case "SignificantPointCondition_position":
						__gong__assignPointer(&inst.SignificantPointCondition_position, rhs, identifierMap)
					case "SignificantPointCondition_runwayPoint":
						__gong__assignPointer(&inst.SignificantPointCondition_runwayPoint, rhs, identifierMap)
					case "SignificantPointCondition_aimingPoint":
						__gong__assignPointer(&inst.SignificantPointCondition_aimingPoint, rhs, identifierMap)
					case "SignificantPointCondition_airportReferencePoint":
						__gong__assignPointer(&inst.SignificantPointCondition_airportReferencePoint, rhs, identifierMap)
					case "FlightCondition_airspaceCondition":
						__gong__assignPointer(&inst.FlightCondition_airspaceCondition, rhs, identifierMap)
					case "FlightCondition_airportHeliportCondition":
						__gong__assignPointer(&inst.FlightCondition_airportHeliportCondition, rhs, identifierMap)
					case "FlightCondition_aerialRefuellingCondition":
						__gong__assignPointer(&inst.FlightCondition_aerialRefuellingCondition, rhs, identifierMap)
					case "OperationalCondition":
						__gong__assignPointer(&inst.OperationalCondition, rhs, identifierMap)
					case "FlightLevel":
						__gong__assignSliceOfPointers(&inst.FlightLevel, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *FlightRestrictionLevelPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "FlightRestrictionLevel":
						__gong__assignPointer(&inst.FlightRestrictionLevel, rhs, identifierMap)
					}
				case *FlightRestrictionLevelType:
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
				case *FlightRestrictionRoutePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "FlightRestrictionRoute":
						__gong__assignPointer(&inst.FlightRestrictionRoute, rhs, identifierMap)
					}
				case *FlightRestrictionRouteType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "RouteElement":
						__gong__assignSliceOfPointers(&inst.RouteElement, rhs, identifierMap)
					case "Contact":
						__gong__assignSliceOfPointers(&inst.Contact, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *FlightRestrictionTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "FlightRestrictionTimeSlice":
						__gong__assignPointer(&inst.FlightRestrictionTimeSlice, rhs, identifierMap)
					}
				case *FlightRestrictionTimeSliceType:
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
					case "Flight":
						__gong__assignPointer(&inst.Flight, rhs, identifierMap)
					case "RegulatedRoute":
						__gong__assignSliceOfPointers(&inst.RegulatedRoute, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *FlightRestrictionType:
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
				case *FlightRoutingElementPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "FlightRoutingElement":
						__gong__assignPointer(&inst.FlightRoutingElement, rhs, identifierMap)
					}
				case *FlightRoutingElementType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "FlightLevel":
						__gong__assignSliceOfPointers(&inst.FlightLevel, rhs, identifierMap)
					case "Element_directFlightElement":
						__gong__assignPointer(&inst.Element_directFlightElement, rhs, identifierMap)
					case "Element_routePortionElement":
						__gong__assignPointer(&inst.Element_routePortionElement, rhs, identifierMap)
					case "Element_standardInstrumentDepartureElement":
						__gong__assignPointer(&inst.Element_standardInstrumentDepartureElement, rhs, identifierMap)
					case "Element_standardInstrumentArrivalElement":
						__gong__assignPointer(&inst.Element_standardInstrumentArrivalElement, rhs, identifierMap)
					case "PointElement_fixDesignatedPoint":
						__gong__assignPointer(&inst.PointElement_fixDesignatedPoint, rhs, identifierMap)
					case "PointElement_navaidSystem":
						__gong__assignPointer(&inst.PointElement_navaidSystem, rhs, identifierMap)
					case "PointElement_position":
						__gong__assignPointer(&inst.PointElement_position, rhs, identifierMap)
					case "PointElement_runwayPoint":
						__gong__assignPointer(&inst.PointElement_runwayPoint, rhs, identifierMap)
					case "PointElement_aimingPoint":
						__gong__assignPointer(&inst.PointElement_aimingPoint, rhs, identifierMap)
					case "PointElement_airportReferencePoint":
						__gong__assignPointer(&inst.PointElement_airportReferencePoint, rhs, identifierMap)
					case "Element_airspaceElement":
						__gong__assignPointer(&inst.Element_airspaceElement, rhs, identifierMap)
					case "Element_airportHeliportElement":
						__gong__assignPointer(&inst.Element_airportHeliportElement, rhs, identifierMap)
					case "Element_aerialRefuellingElement":
						__gong__assignPointer(&inst.Element_aerialRefuellingElement, rhs, identifierMap)
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
				case *MeteorologyPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Meteorology":
						__gong__assignPointer(&inst.Meteorology, rhs, identifierMap)
					}
				case *MeteorologyType:
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
				case *OnlineContactPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "OnlineContact":
						__gong__assignPointer(&inst.OnlineContact, rhs, identifierMap)
					}
				case *OnlineContactType:
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
				case *PostalAddressPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "PostalAddress":
						__gong__assignPointer(&inst.PostalAddress, rhs, identifierMap)
					}
				case *PostalAddressType:
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
				case *RouteAvailabilityPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "RouteAvailability":
						__gong__assignPointer(&inst.RouteAvailability, rhs, identifierMap)
					}
				case *RouteAvailabilityType:
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
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *RouteDMETimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RouteDMETimeSlice":
						__gong__assignPointer(&inst.RouteDMETimeSlice, rhs, identifierMap)
					}
				case *RouteDMETimeSliceType:
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
					case "ReferencedDME":
						__gong__assignPointer(&inst.ReferencedDME, rhs, identifierMap)
					case "ApplicableRoutePortion":
						__gong__assignPointer(&inst.ApplicableRoutePortion, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *RouteDMEType:
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
				case *RoutePortionPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "RoutePortion":
						__gong__assignPointer(&inst.RoutePortion, rhs, identifierMap)
					}
				case *RoutePortionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Start_fixDesignatedPoint":
						__gong__assignPointer(&inst.Start_fixDesignatedPoint, rhs, identifierMap)
					case "Start_navaidSystem":
						__gong__assignPointer(&inst.Start_navaidSystem, rhs, identifierMap)
					case "Start_position":
						__gong__assignPointer(&inst.Start_position, rhs, identifierMap)
					case "Start_runwayPoint":
						__gong__assignPointer(&inst.Start_runwayPoint, rhs, identifierMap)
					case "Start_aimingPoint":
						__gong__assignPointer(&inst.Start_aimingPoint, rhs, identifierMap)
					case "Start_airportReferencePoint":
						__gong__assignPointer(&inst.Start_airportReferencePoint, rhs, identifierMap)
					case "IntermediatePoint_fixDesignatedPoint":
						__gong__assignPointer(&inst.IntermediatePoint_fixDesignatedPoint, rhs, identifierMap)
					case "IntermediatePoint_navaidSystem":
						__gong__assignPointer(&inst.IntermediatePoint_navaidSystem, rhs, identifierMap)
					case "IntermediatePoint_position":
						__gong__assignPointer(&inst.IntermediatePoint_position, rhs, identifierMap)
					case "IntermediatePoint_runwayPoint":
						__gong__assignPointer(&inst.IntermediatePoint_runwayPoint, rhs, identifierMap)
					case "IntermediatePoint_aimingPoint":
						__gong__assignPointer(&inst.IntermediatePoint_aimingPoint, rhs, identifierMap)
					case "IntermediatePoint_airportReferencePoint":
						__gong__assignPointer(&inst.IntermediatePoint_airportReferencePoint, rhs, identifierMap)
					case "ReferencedRoute":
						__gong__assignPointer(&inst.ReferencedRoute, rhs, identifierMap)
					case "End_fixDesignatedPoint":
						__gong__assignPointer(&inst.End_fixDesignatedPoint, rhs, identifierMap)
					case "End_navaidSystem":
						__gong__assignPointer(&inst.End_navaidSystem, rhs, identifierMap)
					case "End_position":
						__gong__assignPointer(&inst.End_position, rhs, identifierMap)
					case "End_runwayPoint":
						__gong__assignPointer(&inst.End_runwayPoint, rhs, identifierMap)
					case "End_aimingPoint":
						__gong__assignPointer(&inst.End_aimingPoint, rhs, identifierMap)
					case "End_airportReferencePoint":
						__gong__assignPointer(&inst.End_airportReferencePoint, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
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
				case *RouteSegmentTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RouteSegmentTimeSlice":
						__gong__assignPointer(&inst.RouteSegmentTimeSlice, rhs, identifierMap)
					}
				case *RouteSegmentTimeSliceType:
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
					case "Start":
						__gong__assignPointer(&inst.Start, rhs, identifierMap)
					case "RouteFormed":
						__gong__assignPointer(&inst.RouteFormed, rhs, identifierMap)
					case "EvaluationArea":
						__gong__assignPointer(&inst.EvaluationArea, rhs, identifierMap)
					case "CurveExtent":
						__gong__assignPointer(&inst.CurveExtent, rhs, identifierMap)
					case "End":
						__gong__assignPointer(&inst.End, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *RouteSegmentType:
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
				case *RouteTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RouteTimeSlice":
						__gong__assignPointer(&inst.RouteTimeSlice, rhs, identifierMap)
					}
				case *RouteTimeSliceType:
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
					case "UserOrganisation":
						__gong__assignPointer(&inst.UserOrganisation, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *RouteType:
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
				case *RulesProceduresTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RulesProceduresTimeSlice":
						__gong__assignPointer(&inst.RulesProceduresTimeSlice, rhs, identifierMap)
					}
				case *RulesProceduresTimeSliceType:
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
					case "AffectedLocation":
						__gong__assignSliceOfPointers(&inst.AffectedLocation, rhs, identifierMap)
					case "AffectedArea":
						__gong__assignSliceOfPointers(&inst.AffectedArea, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *RulesProceduresType:
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
				case *TelephoneContactPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "TelephoneContact":
						__gong__assignPointer(&inst.TelephoneContact, rhs, identifierMap)
					}
				case *TelephoneContactType:
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

type AbstractDirectFlightTypeUnmarshaller struct{}

func (u *AbstractDirectFlightTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractDirectFlightType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractDirectFlightTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractDirectFlightType)
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

type AerialRefuellingAnchorPropertyTypeUnmarshaller struct{}

func (u *AerialRefuellingAnchorPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AerialRefuellingAnchorPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AerialRefuellingAnchorPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AerialRefuellingAnchorPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "AerialRefuellingAnchor":
		GongUnmarshallPointer(&instance.AerialRefuellingAnchor, valueExpr, identifierMap)
	}
	return nil
}

type AerialRefuellingAnchorTypeUnmarshaller struct{}

func (u *AerialRefuellingAnchorTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AerialRefuellingAnchorType), stage, identifier, instanceName, preserveOrder)
}

func (u *AerialRefuellingAnchorTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AerialRefuellingAnchorType)
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
	case "VerticalExtent":
		GongUnmarshallSliceOfPointers(&instance.VerticalExtent, valueExpr, identifierMap)
	case "Point":
		GongUnmarshallSliceOfPointers(&instance.Point, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AerialRefuellingPointPropertyTypeUnmarshaller struct{}

func (u *AerialRefuellingPointPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AerialRefuellingPointPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AerialRefuellingPointPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AerialRefuellingPointPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "AerialRefuellingPoint":
		GongUnmarshallPointer(&instance.AerialRefuellingPoint, valueExpr, identifierMap)
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

type AerialRefuellingPropertyTypeUnmarshaller struct{}

func (u *AerialRefuellingPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AerialRefuellingPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AerialRefuellingPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AerialRefuellingPropertyType)
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

type AerialRefuellingTimeSlicePropertyTypeUnmarshaller struct{}

func (u *AerialRefuellingTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AerialRefuellingTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AerialRefuellingTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AerialRefuellingTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "AerialRefuellingTimeSlice":
		GongUnmarshallPointer(&instance.AerialRefuellingTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type AerialRefuellingTimeSliceTypeUnmarshaller struct{}

func (u *AerialRefuellingTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AerialRefuellingTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *AerialRefuellingTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AerialRefuellingTimeSliceType)
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
	case "ProtectingAirspace":
		GongUnmarshallSliceOfPointers(&instance.ProtectingAirspace, valueExpr, identifierMap)
	case "Track":
		GongUnmarshallSliceOfPointers(&instance.Track, valueExpr, identifierMap)
	case "Anchor":
		GongUnmarshallSliceOfPointers(&instance.Anchor, valueExpr, identifierMap)
	case "OppositeTrack":
		GongUnmarshallPointer(&instance.OppositeTrack, valueExpr, identifierMap)
	case "ManagingOrganisation":
		GongUnmarshallSliceOfPointers(&instance.ManagingOrganisation, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AerialRefuellingTrackPropertyTypeUnmarshaller struct{}

func (u *AerialRefuellingTrackPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AerialRefuellingTrackPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AerialRefuellingTrackPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AerialRefuellingTrackPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "AerialRefuellingTrack":
		GongUnmarshallPointer(&instance.AerialRefuellingTrack, valueExpr, identifierMap)
	}
	return nil
}

type AerialRefuellingTrackTypeUnmarshaller struct{}

func (u *AerialRefuellingTrackTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AerialRefuellingTrackType), stage, identifier, instanceName, preserveOrder)
}

func (u *AerialRefuellingTrackTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AerialRefuellingTrackType)
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
	case "Point":
		GongUnmarshallSliceOfPointers(&instance.Point, valueExpr, identifierMap)
	case "VerticalExtent":
		GongUnmarshallSliceOfPointers(&instance.VerticalExtent, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AerialRefuellingTypeUnmarshaller struct{}

func (u *AerialRefuellingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AerialRefuellingType), stage, identifier, instanceName, preserveOrder)
}

func (u *AerialRefuellingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AerialRefuellingType)
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

type AirspaceBorderCrossingPropertyTypeUnmarshaller struct{}

func (u *AirspaceBorderCrossingPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirspaceBorderCrossingPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirspaceBorderCrossingPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirspaceBorderCrossingPropertyType)
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

type AirspaceBorderCrossingTimeSlicePropertyTypeUnmarshaller struct{}

func (u *AirspaceBorderCrossingTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirspaceBorderCrossingTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirspaceBorderCrossingTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirspaceBorderCrossingTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "AirspaceBorderCrossingTimeSlice":
		GongUnmarshallPointer(&instance.AirspaceBorderCrossingTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type AirspaceBorderCrossingTimeSliceTypeUnmarshaller struct{}

func (u *AirspaceBorderCrossingTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirspaceBorderCrossingTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirspaceBorderCrossingTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirspaceBorderCrossingTimeSliceType)
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
	case "ExitedAirspace":
		GongUnmarshallPointer(&instance.ExitedAirspace, valueExpr, identifierMap)
	case "EnteredAirspace":
		GongUnmarshallPointer(&instance.EnteredAirspace, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AirspaceBorderCrossingTypeUnmarshaller struct{}

func (u *AirspaceBorderCrossingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirspaceBorderCrossingType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirspaceBorderCrossingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirspaceBorderCrossingType)
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

type AuthorityForAerialRefuellingPropertyTypeUnmarshaller struct{}

func (u *AuthorityForAerialRefuellingPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AuthorityForAerialRefuellingPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AuthorityForAerialRefuellingPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AuthorityForAerialRefuellingPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "AuthorityForAerialRefuelling":
		GongUnmarshallPointer(&instance.AuthorityForAerialRefuelling, valueExpr, identifierMap)
	}
	return nil
}

type AuthorityForAerialRefuellingTypeUnmarshaller struct{}

func (u *AuthorityForAerialRefuellingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AuthorityForAerialRefuellingType), stage, identifier, instanceName, preserveOrder)
}

func (u *AuthorityForAerialRefuellingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AuthorityForAerialRefuellingType)
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
	case "TheOrganisationAuthority":
		GongUnmarshallPointer(&instance.TheOrganisationAuthority, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type ChangeOverPointTimeSlicePropertyTypeUnmarshaller struct{}

func (u *ChangeOverPointTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ChangeOverPointTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ChangeOverPointTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ChangeOverPointTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "ChangeOverPointTimeSlice":
		GongUnmarshallPointer(&instance.ChangeOverPointTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type ChangeOverPointTimeSliceTypeUnmarshaller struct{}

func (u *ChangeOverPointTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ChangeOverPointTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *ChangeOverPointTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ChangeOverPointTimeSliceType)
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
	case "Location_fixDesignatedPoint":
		GongUnmarshallPointer(&instance.Location_fixDesignatedPoint, valueExpr, identifierMap)
	case "Location_navaidSystem":
		GongUnmarshallPointer(&instance.Location_navaidSystem, valueExpr, identifierMap)
	case "Location_position":
		GongUnmarshallPointer(&instance.Location_position, valueExpr, identifierMap)
	case "Location_runwayPoint":
		GongUnmarshallPointer(&instance.Location_runwayPoint, valueExpr, identifierMap)
	case "Location_aimingPoint":
		GongUnmarshallPointer(&instance.Location_aimingPoint, valueExpr, identifierMap)
	case "Location_airportReferencePoint":
		GongUnmarshallPointer(&instance.Location_airportReferencePoint, valueExpr, identifierMap)
	case "ApplicableRoutePortion":
		GongUnmarshallPointer(&instance.ApplicableRoutePortion, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type ChangeOverPointTypeUnmarshaller struct{}

func (u *ChangeOverPointTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ChangeOverPointType), stage, identifier, instanceName, preserveOrder)
}

func (u *ChangeOverPointTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ChangeOverPointType)
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

type ContactInformationPropertyTypeUnmarshaller struct{}

func (u *ContactInformationPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ContactInformationPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ContactInformationPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ContactInformationPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "ContactInformation":
		GongUnmarshallPointer(&instance.ContactInformation, valueExpr, identifierMap)
	}
	return nil
}

type ContactInformationTypeUnmarshaller struct{}

func (u *ContactInformationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ContactInformationType), stage, identifier, instanceName, preserveOrder)
}

func (u *ContactInformationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ContactInformationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Address":
		GongUnmarshallSliceOfPointers(&instance.Address, valueExpr, identifierMap)
	case "NetworkNode":
		GongUnmarshallSliceOfPointers(&instance.NetworkNode, valueExpr, identifierMap)
	case "PhoneFax":
		GongUnmarshallSliceOfPointers(&instance.PhoneFax, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
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

type DirectFlightClassTypeUnmarshaller struct{}

func (u *DirectFlightClassTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DirectFlightClassType), stage, identifier, instanceName, preserveOrder)
}

func (u *DirectFlightClassTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DirectFlightClassType)
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

type DirectFlightPropertyTypeUnmarshaller struct{}

func (u *DirectFlightPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DirectFlightPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *DirectFlightPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DirectFlightPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "DirectFlightClass":
		GongUnmarshallPointer(&instance.DirectFlightClass, valueExpr, identifierMap)
	case "DirectFlightSegment":
		GongUnmarshallPointer(&instance.DirectFlightSegment, valueExpr, identifierMap)
	}
	return nil
}

type DirectFlightSegmentPropertyTypeUnmarshaller struct{}

func (u *DirectFlightSegmentPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DirectFlightSegmentPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *DirectFlightSegmentPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DirectFlightSegmentPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "DirectFlightSegment":
		GongUnmarshallPointer(&instance.DirectFlightSegment, valueExpr, identifierMap)
	}
	return nil
}

type DirectFlightSegmentTypeUnmarshaller struct{}

func (u *DirectFlightSegmentTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DirectFlightSegmentType), stage, identifier, instanceName, preserveOrder)
}

func (u *DirectFlightSegmentTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DirectFlightSegmentType)
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
	case "End_fixDesignatedPoint":
		GongUnmarshallPointer(&instance.End_fixDesignatedPoint, valueExpr, identifierMap)
	case "End_navaidSystem":
		GongUnmarshallPointer(&instance.End_navaidSystem, valueExpr, identifierMap)
	case "End_position":
		GongUnmarshallPointer(&instance.End_position, valueExpr, identifierMap)
	case "End_runwayPoint":
		GongUnmarshallPointer(&instance.End_runwayPoint, valueExpr, identifierMap)
	case "End_aimingPoint":
		GongUnmarshallPointer(&instance.End_aimingPoint, valueExpr, identifierMap)
	case "End_airportReferencePoint":
		GongUnmarshallPointer(&instance.End_airportReferencePoint, valueExpr, identifierMap)
	case "Start_fixDesignatedPoint":
		GongUnmarshallPointer(&instance.Start_fixDesignatedPoint, valueExpr, identifierMap)
	case "Start_navaidSystem":
		GongUnmarshallPointer(&instance.Start_navaidSystem, valueExpr, identifierMap)
	case "Start_position":
		GongUnmarshallPointer(&instance.Start_position, valueExpr, identifierMap)
	case "Start_runwayPoint":
		GongUnmarshallPointer(&instance.Start_runwayPoint, valueExpr, identifierMap)
	case "Start_aimingPoint":
		GongUnmarshallPointer(&instance.Start_aimingPoint, valueExpr, identifierMap)
	case "Start_airportReferencePoint":
		GongUnmarshallPointer(&instance.Start_airportReferencePoint, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
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

type EnRouteSegmentPointPropertyTypeUnmarshaller struct{}

func (u *EnRouteSegmentPointPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EnRouteSegmentPointPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *EnRouteSegmentPointPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EnRouteSegmentPointPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnRouteSegmentPoint":
		GongUnmarshallPointer(&instance.EnRouteSegmentPoint, valueExpr, identifierMap)
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

type FlightCharacteristicPropertyTypeUnmarshaller struct{}

func (u *FlightCharacteristicPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FlightCharacteristicPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *FlightCharacteristicPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FlightCharacteristicPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "FlightCharacteristic":
		GongUnmarshallPointer(&instance.FlightCharacteristic, valueExpr, identifierMap)
	}
	return nil
}

type FlightCharacteristicTypeUnmarshaller struct{}

func (u *FlightCharacteristicTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FlightCharacteristicType), stage, identifier, instanceName, preserveOrder)
}

func (u *FlightCharacteristicTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FlightCharacteristicType)
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

type FlightConditionCircumstancePropertyTypeUnmarshaller struct{}

func (u *FlightConditionCircumstancePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FlightConditionCircumstancePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *FlightConditionCircumstancePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FlightConditionCircumstancePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "FlightConditionCircumstance":
		GongUnmarshallPointer(&instance.FlightConditionCircumstance, valueExpr, identifierMap)
	}
	return nil
}

type FlightConditionCircumstanceTypeUnmarshaller struct{}

func (u *FlightConditionCircumstanceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FlightConditionCircumstanceType), stage, identifier, instanceName, preserveOrder)
}

func (u *FlightConditionCircumstanceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FlightConditionCircumstanceType)
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

type FlightConditionCombinationPropertyTypeUnmarshaller struct{}

func (u *FlightConditionCombinationPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FlightConditionCombinationPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *FlightConditionCombinationPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FlightConditionCombinationPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "FlightConditionCombination":
		GongUnmarshallPointer(&instance.FlightConditionCombination, valueExpr, identifierMap)
	}
	return nil
}

type FlightConditionCombinationTypeUnmarshaller struct{}

func (u *FlightConditionCombinationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FlightConditionCombinationType), stage, identifier, instanceName, preserveOrder)
}

func (u *FlightConditionCombinationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FlightConditionCombinationType)
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
	case "Element":
		GongUnmarshallSliceOfPointers(&instance.Element, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type FlightConditionElementPropertyTypeUnmarshaller struct{}

func (u *FlightConditionElementPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FlightConditionElementPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *FlightConditionElementPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FlightConditionElementPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "FlightConditionElement":
		GongUnmarshallPointer(&instance.FlightConditionElement, valueExpr, identifierMap)
	}
	return nil
}

type FlightConditionElementTypeUnmarshaller struct{}

func (u *FlightConditionElementTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FlightConditionElementType), stage, identifier, instanceName, preserveOrder)
}

func (u *FlightConditionElementTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FlightConditionElementType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "FlightCondition_weather":
		GongUnmarshallPointer(&instance.FlightCondition_weather, valueExpr, identifierMap)
	case "FlightCondition_aircraft":
		GongUnmarshallPointer(&instance.FlightCondition_aircraft, valueExpr, identifierMap)
	case "FlightCondition_flight":
		GongUnmarshallPointer(&instance.FlightCondition_flight, valueExpr, identifierMap)
	case "FlightCondition_directFlightCondition":
		GongUnmarshallPointer(&instance.FlightCondition_directFlightCondition, valueExpr, identifierMap)
	case "FlightCondition_borderCrossingCondition":
		GongUnmarshallPointer(&instance.FlightCondition_borderCrossingCondition, valueExpr, identifierMap)
	case "FlightCondition_operand":
		GongUnmarshallPointer(&instance.FlightCondition_operand, valueExpr, identifierMap)
	case "FlightCondition_routePortionCondition":
		GongUnmarshallPointer(&instance.FlightCondition_routePortionCondition, valueExpr, identifierMap)
	case "FlightCondition_standardInstrumentDepartureCondition":
		GongUnmarshallPointer(&instance.FlightCondition_standardInstrumentDepartureCondition, valueExpr, identifierMap)
	case "FlightCondition_standardInstrumentArrivalCondition":
		GongUnmarshallPointer(&instance.FlightCondition_standardInstrumentArrivalCondition, valueExpr, identifierMap)
	case "FlightCondition_organisationCondition":
		GongUnmarshallPointer(&instance.FlightCondition_organisationCondition, valueExpr, identifierMap)
	case "SignificantPointCondition_fixDesignatedPoint":
		GongUnmarshallPointer(&instance.SignificantPointCondition_fixDesignatedPoint, valueExpr, identifierMap)
	case "SignificantPointCondition_navaidSystem":
		GongUnmarshallPointer(&instance.SignificantPointCondition_navaidSystem, valueExpr, identifierMap)
	case "SignificantPointCondition_position":
		GongUnmarshallPointer(&instance.SignificantPointCondition_position, valueExpr, identifierMap)
	case "SignificantPointCondition_runwayPoint":
		GongUnmarshallPointer(&instance.SignificantPointCondition_runwayPoint, valueExpr, identifierMap)
	case "SignificantPointCondition_aimingPoint":
		GongUnmarshallPointer(&instance.SignificantPointCondition_aimingPoint, valueExpr, identifierMap)
	case "SignificantPointCondition_airportReferencePoint":
		GongUnmarshallPointer(&instance.SignificantPointCondition_airportReferencePoint, valueExpr, identifierMap)
	case "FlightCondition_airspaceCondition":
		GongUnmarshallPointer(&instance.FlightCondition_airspaceCondition, valueExpr, identifierMap)
	case "FlightCondition_airportHeliportCondition":
		GongUnmarshallPointer(&instance.FlightCondition_airportHeliportCondition, valueExpr, identifierMap)
	case "FlightCondition_aerialRefuellingCondition":
		GongUnmarshallPointer(&instance.FlightCondition_aerialRefuellingCondition, valueExpr, identifierMap)
	case "OperationalCondition":
		GongUnmarshallPointer(&instance.OperationalCondition, valueExpr, identifierMap)
	case "FlightLevel":
		GongUnmarshallSliceOfPointers(&instance.FlightLevel, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type FlightRestrictionLevelPropertyTypeUnmarshaller struct{}

func (u *FlightRestrictionLevelPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FlightRestrictionLevelPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *FlightRestrictionLevelPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FlightRestrictionLevelPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "FlightRestrictionLevel":
		GongUnmarshallPointer(&instance.FlightRestrictionLevel, valueExpr, identifierMap)
	}
	return nil
}

type FlightRestrictionLevelTypeUnmarshaller struct{}

func (u *FlightRestrictionLevelTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FlightRestrictionLevelType), stage, identifier, instanceName, preserveOrder)
}

func (u *FlightRestrictionLevelTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FlightRestrictionLevelType)
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

type FlightRestrictionRoutePropertyTypeUnmarshaller struct{}

func (u *FlightRestrictionRoutePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FlightRestrictionRoutePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *FlightRestrictionRoutePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FlightRestrictionRoutePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "FlightRestrictionRoute":
		GongUnmarshallPointer(&instance.FlightRestrictionRoute, valueExpr, identifierMap)
	}
	return nil
}

type FlightRestrictionRouteTypeUnmarshaller struct{}

func (u *FlightRestrictionRouteTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FlightRestrictionRouteType), stage, identifier, instanceName, preserveOrder)
}

func (u *FlightRestrictionRouteTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FlightRestrictionRouteType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "RouteElement":
		GongUnmarshallSliceOfPointers(&instance.RouteElement, valueExpr, identifierMap)
	case "Contact":
		GongUnmarshallSliceOfPointers(&instance.Contact, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type FlightRestrictionTimeSlicePropertyTypeUnmarshaller struct{}

func (u *FlightRestrictionTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FlightRestrictionTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *FlightRestrictionTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FlightRestrictionTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "FlightRestrictionTimeSlice":
		GongUnmarshallPointer(&instance.FlightRestrictionTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type FlightRestrictionTimeSliceTypeUnmarshaller struct{}

func (u *FlightRestrictionTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FlightRestrictionTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *FlightRestrictionTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FlightRestrictionTimeSliceType)
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
	case "Flight":
		GongUnmarshallPointer(&instance.Flight, valueExpr, identifierMap)
	case "RegulatedRoute":
		GongUnmarshallSliceOfPointers(&instance.RegulatedRoute, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type FlightRestrictionTypeUnmarshaller struct{}

func (u *FlightRestrictionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FlightRestrictionType), stage, identifier, instanceName, preserveOrder)
}

func (u *FlightRestrictionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FlightRestrictionType)
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

type FlightRoutingElementPropertyTypeUnmarshaller struct{}

func (u *FlightRoutingElementPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FlightRoutingElementPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *FlightRoutingElementPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FlightRoutingElementPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "FlightRoutingElement":
		GongUnmarshallPointer(&instance.FlightRoutingElement, valueExpr, identifierMap)
	}
	return nil
}

type FlightRoutingElementTypeUnmarshaller struct{}

func (u *FlightRoutingElementTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FlightRoutingElementType), stage, identifier, instanceName, preserveOrder)
}

func (u *FlightRoutingElementTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FlightRoutingElementType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "FlightLevel":
		GongUnmarshallSliceOfPointers(&instance.FlightLevel, valueExpr, identifierMap)
	case "Element_directFlightElement":
		GongUnmarshallPointer(&instance.Element_directFlightElement, valueExpr, identifierMap)
	case "Element_routePortionElement":
		GongUnmarshallPointer(&instance.Element_routePortionElement, valueExpr, identifierMap)
	case "Element_standardInstrumentDepartureElement":
		GongUnmarshallPointer(&instance.Element_standardInstrumentDepartureElement, valueExpr, identifierMap)
	case "Element_standardInstrumentArrivalElement":
		GongUnmarshallPointer(&instance.Element_standardInstrumentArrivalElement, valueExpr, identifierMap)
	case "PointElement_fixDesignatedPoint":
		GongUnmarshallPointer(&instance.PointElement_fixDesignatedPoint, valueExpr, identifierMap)
	case "PointElement_navaidSystem":
		GongUnmarshallPointer(&instance.PointElement_navaidSystem, valueExpr, identifierMap)
	case "PointElement_position":
		GongUnmarshallPointer(&instance.PointElement_position, valueExpr, identifierMap)
	case "PointElement_runwayPoint":
		GongUnmarshallPointer(&instance.PointElement_runwayPoint, valueExpr, identifierMap)
	case "PointElement_aimingPoint":
		GongUnmarshallPointer(&instance.PointElement_aimingPoint, valueExpr, identifierMap)
	case "PointElement_airportReferencePoint":
		GongUnmarshallPointer(&instance.PointElement_airportReferencePoint, valueExpr, identifierMap)
	case "Element_airspaceElement":
		GongUnmarshallPointer(&instance.Element_airspaceElement, valueExpr, identifierMap)
	case "Element_airportHeliportElement":
		GongUnmarshallPointer(&instance.Element_airportHeliportElement, valueExpr, identifierMap)
	case "Element_aerialRefuellingElement":
		GongUnmarshallPointer(&instance.Element_aerialRefuellingElement, valueExpr, identifierMap)
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

type MeteorologyPropertyTypeUnmarshaller struct{}

func (u *MeteorologyPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MeteorologyPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *MeteorologyPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MeteorologyPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Meteorology":
		GongUnmarshallPointer(&instance.Meteorology, valueExpr, identifierMap)
	}
	return nil
}

type MeteorologyTypeUnmarshaller struct{}

func (u *MeteorologyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MeteorologyType), stage, identifier, instanceName, preserveOrder)
}

func (u *MeteorologyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MeteorologyType)
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

type OnlineContactPropertyTypeUnmarshaller struct{}

func (u *OnlineContactPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(OnlineContactPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *OnlineContactPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*OnlineContactPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "OnlineContact":
		GongUnmarshallPointer(&instance.OnlineContact, valueExpr, identifierMap)
	}
	return nil
}

type OnlineContactTypeUnmarshaller struct{}

func (u *OnlineContactTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(OnlineContactType), stage, identifier, instanceName, preserveOrder)
}

func (u *OnlineContactTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*OnlineContactType)
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

type PostalAddressPropertyTypeUnmarshaller struct{}

func (u *PostalAddressPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PostalAddressPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *PostalAddressPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PostalAddressPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "PostalAddress":
		GongUnmarshallPointer(&instance.PostalAddress, valueExpr, identifierMap)
	}
	return nil
}

type PostalAddressTypeUnmarshaller struct{}

func (u *PostalAddressTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PostalAddressType), stage, identifier, instanceName, preserveOrder)
}

func (u *PostalAddressTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PostalAddressType)
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

type RouteAvailabilityPropertyTypeUnmarshaller struct{}

func (u *RouteAvailabilityPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RouteAvailabilityPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RouteAvailabilityPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RouteAvailabilityPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "RouteAvailability":
		GongUnmarshallPointer(&instance.RouteAvailability, valueExpr, identifierMap)
	}
	return nil
}

type RouteAvailabilityTypeUnmarshaller struct{}

func (u *RouteAvailabilityTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RouteAvailabilityType), stage, identifier, instanceName, preserveOrder)
}

func (u *RouteAvailabilityTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RouteAvailabilityType)
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
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type RouteDMETimeSlicePropertyTypeUnmarshaller struct{}

func (u *RouteDMETimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RouteDMETimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RouteDMETimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RouteDMETimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "RouteDMETimeSlice":
		GongUnmarshallPointer(&instance.RouteDMETimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type RouteDMETimeSliceTypeUnmarshaller struct{}

func (u *RouteDMETimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RouteDMETimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *RouteDMETimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RouteDMETimeSliceType)
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
	case "ReferencedDME":
		GongUnmarshallPointer(&instance.ReferencedDME, valueExpr, identifierMap)
	case "ApplicableRoutePortion":
		GongUnmarshallPointer(&instance.ApplicableRoutePortion, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type RouteDMETypeUnmarshaller struct{}

func (u *RouteDMETypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RouteDMEType), stage, identifier, instanceName, preserveOrder)
}

func (u *RouteDMETypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RouteDMEType)
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

type RoutePortionPropertyTypeUnmarshaller struct{}

func (u *RoutePortionPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RoutePortionPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RoutePortionPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RoutePortionPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "RoutePortion":
		GongUnmarshallPointer(&instance.RoutePortion, valueExpr, identifierMap)
	}
	return nil
}

type RoutePortionTypeUnmarshaller struct{}

func (u *RoutePortionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RoutePortionType), stage, identifier, instanceName, preserveOrder)
}

func (u *RoutePortionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RoutePortionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Start_fixDesignatedPoint":
		GongUnmarshallPointer(&instance.Start_fixDesignatedPoint, valueExpr, identifierMap)
	case "Start_navaidSystem":
		GongUnmarshallPointer(&instance.Start_navaidSystem, valueExpr, identifierMap)
	case "Start_position":
		GongUnmarshallPointer(&instance.Start_position, valueExpr, identifierMap)
	case "Start_runwayPoint":
		GongUnmarshallPointer(&instance.Start_runwayPoint, valueExpr, identifierMap)
	case "Start_aimingPoint":
		GongUnmarshallPointer(&instance.Start_aimingPoint, valueExpr, identifierMap)
	case "Start_airportReferencePoint":
		GongUnmarshallPointer(&instance.Start_airportReferencePoint, valueExpr, identifierMap)
	case "IntermediatePoint_fixDesignatedPoint":
		GongUnmarshallPointer(&instance.IntermediatePoint_fixDesignatedPoint, valueExpr, identifierMap)
	case "IntermediatePoint_navaidSystem":
		GongUnmarshallPointer(&instance.IntermediatePoint_navaidSystem, valueExpr, identifierMap)
	case "IntermediatePoint_position":
		GongUnmarshallPointer(&instance.IntermediatePoint_position, valueExpr, identifierMap)
	case "IntermediatePoint_runwayPoint":
		GongUnmarshallPointer(&instance.IntermediatePoint_runwayPoint, valueExpr, identifierMap)
	case "IntermediatePoint_aimingPoint":
		GongUnmarshallPointer(&instance.IntermediatePoint_aimingPoint, valueExpr, identifierMap)
	case "IntermediatePoint_airportReferencePoint":
		GongUnmarshallPointer(&instance.IntermediatePoint_airportReferencePoint, valueExpr, identifierMap)
	case "ReferencedRoute":
		GongUnmarshallPointer(&instance.ReferencedRoute, valueExpr, identifierMap)
	case "End_fixDesignatedPoint":
		GongUnmarshallPointer(&instance.End_fixDesignatedPoint, valueExpr, identifierMap)
	case "End_navaidSystem":
		GongUnmarshallPointer(&instance.End_navaidSystem, valueExpr, identifierMap)
	case "End_position":
		GongUnmarshallPointer(&instance.End_position, valueExpr, identifierMap)
	case "End_runwayPoint":
		GongUnmarshallPointer(&instance.End_runwayPoint, valueExpr, identifierMap)
	case "End_aimingPoint":
		GongUnmarshallPointer(&instance.End_aimingPoint, valueExpr, identifierMap)
	case "End_airportReferencePoint":
		GongUnmarshallPointer(&instance.End_airportReferencePoint, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
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

type RouteSegmentTimeSlicePropertyTypeUnmarshaller struct{}

func (u *RouteSegmentTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RouteSegmentTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RouteSegmentTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RouteSegmentTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "RouteSegmentTimeSlice":
		GongUnmarshallPointer(&instance.RouteSegmentTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type RouteSegmentTimeSliceTypeUnmarshaller struct{}

func (u *RouteSegmentTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RouteSegmentTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *RouteSegmentTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RouteSegmentTimeSliceType)
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
	case "Start":
		GongUnmarshallPointer(&instance.Start, valueExpr, identifierMap)
	case "RouteFormed":
		GongUnmarshallPointer(&instance.RouteFormed, valueExpr, identifierMap)
	case "EvaluationArea":
		GongUnmarshallPointer(&instance.EvaluationArea, valueExpr, identifierMap)
	case "CurveExtent":
		GongUnmarshallPointer(&instance.CurveExtent, valueExpr, identifierMap)
	case "End":
		GongUnmarshallPointer(&instance.End, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type RouteSegmentTypeUnmarshaller struct{}

func (u *RouteSegmentTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RouteSegmentType), stage, identifier, instanceName, preserveOrder)
}

func (u *RouteSegmentTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RouteSegmentType)
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

type RouteTimeSlicePropertyTypeUnmarshaller struct{}

func (u *RouteTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RouteTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RouteTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RouteTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "RouteTimeSlice":
		GongUnmarshallPointer(&instance.RouteTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type RouteTimeSliceTypeUnmarshaller struct{}

func (u *RouteTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RouteTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *RouteTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RouteTimeSliceType)
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
	case "UserOrganisation":
		GongUnmarshallPointer(&instance.UserOrganisation, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type RouteTypeUnmarshaller struct{}

func (u *RouteTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RouteType), stage, identifier, instanceName, preserveOrder)
}

func (u *RouteTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RouteType)
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

type RulesProceduresTimeSlicePropertyTypeUnmarshaller struct{}

func (u *RulesProceduresTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RulesProceduresTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RulesProceduresTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RulesProceduresTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "RulesProceduresTimeSlice":
		GongUnmarshallPointer(&instance.RulesProceduresTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type RulesProceduresTimeSliceTypeUnmarshaller struct{}

func (u *RulesProceduresTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RulesProceduresTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *RulesProceduresTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RulesProceduresTimeSliceType)
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
	case "AffectedLocation":
		GongUnmarshallSliceOfPointers(&instance.AffectedLocation, valueExpr, identifierMap)
	case "AffectedArea":
		GongUnmarshallSliceOfPointers(&instance.AffectedArea, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type RulesProceduresTypeUnmarshaller struct{}

func (u *RulesProceduresTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RulesProceduresType), stage, identifier, instanceName, preserveOrder)
}

func (u *RulesProceduresTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RulesProceduresType)
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

type TelephoneContactPropertyTypeUnmarshaller struct{}

func (u *TelephoneContactPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TelephoneContactPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *TelephoneContactPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TelephoneContactPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "TelephoneContact":
		GongUnmarshallPointer(&instance.TelephoneContact, valueExpr, identifierMap)
	}
	return nil
}

type TelephoneContactTypeUnmarshaller struct{}

func (u *TelephoneContactTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TelephoneContactType), stage, identifier, instanceName, preserveOrder)
}

func (u *TelephoneContactTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TelephoneContactType)
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
