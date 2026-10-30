// generated code - do not edit
package aixmdatatypes

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
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmdatatypes":
			aliasToCanonical[alias] = "aixmdatatypes"
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
					} else if pkgAlias == "" || pkgAlias == "aixmdatatypes" || len(inFile.Imports) == 0 {
						pkgAlias = "aixmdatatypes"
					}

					switch pkgAlias {
			case "aixmdatatypes":
				switch typeName {
				case "CodeATCReportingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeATCReportingType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAerialRefuellingPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAerialRefuellingPointType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAerialRefuellingPrefixType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAerialRefuellingPrefixType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAerialRefuellingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAerialRefuellingType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAircraftCategoryType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAircraftCategoryType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAircraftEngineNumberType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAircraftEngineNumberType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAircraftEngineType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAircraftEngineType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAircraftICAOType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAircraftICAOType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAircraftType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAircraftType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAircraftWingspanClassType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAircraftWingspanClassType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAltitudeAdjustmentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAltitudeAdjustmentType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAltitudeUseType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAltitudeUseType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAuthorityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAuthorityType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeCardinalDirectionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeCardinalDirectionType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeCommunicationModeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeCommunicationModeType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeComparisonType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeComparisonType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeCourseType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeCourseType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeDayType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeDayType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeDirectionTurnType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeDirectionTurnType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeDirectionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeDirectionType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeEquipmentAntiCollisionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeEquipmentAntiCollisionType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeFlightOriginType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeFlightOriginType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeFlightPurposeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeFlightPurposeType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeFlightRestrictionDesignatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeFlightRestrictionDesignatorType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeFlightRestrictionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeFlightRestrictionType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeFlightRuleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeFlightRuleType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeFlightStatusType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeFlightStatusType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeFlightType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeFlightType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeFlowConditionOperationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeFlowConditionOperationType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeFreeFlightType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeFreeFlightType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeLevelType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeLevelType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeLocationQualifierType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeLocationQualifierType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeMeteoConditionsType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeMeteoConditionsType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeMilitaryRoutePointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeMilitaryRoutePointType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeMilitaryStatusType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeMilitaryStatusType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeMilitaryTrainingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeMilitaryTrainingType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeNavigationEquipmentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeNavigationEquipmentType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeNavigationSpecificationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeNavigationSpecificationType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeNotePurposeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeNotePurposeType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeObstacleAssessmentSurfaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeObstacleAssessmentSurfaceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeObstructionIdSurfaceZoneType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeObstructionIdSurfaceZoneType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRNPType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRNPType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRVSMPointRoleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRVSMPointRoleType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRVSMType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRVSMType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeReferenceRoleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeReferenceRoleType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRouteAvailabilityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRouteAvailabilityType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRouteDesignatorLetterType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRouteDesignatorLetterType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRouteDesignatorPrefixType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRouteDesignatorPrefixType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRouteDesignatorSuffixType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRouteDesignatorSuffixType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRouteNavigationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRouteNavigationType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRouteOriginType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRouteOriginType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRouteSegmentPathType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRouteSegmentPathType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRouteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRouteType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRuleProcedureTitleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRuleProcedureTitleType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRuleProcedureType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRuleProcedureType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeSideType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeSideType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeSpeedReferenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeSpeedReferenceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeTACANChannelType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeTACANChannelType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeTelecomNetworkType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeTelecomNetworkType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeTimeEventCombinationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeTimeEventCombinationType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeTimeEventType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeTimeEventType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeTimeReferenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeTimeReferenceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeTransponderType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeTransponderType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeUpperAlphaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeUpperAlphaType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeValueInterpretationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeValueInterpretationType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeVerticalDatumType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeVerticalDatumType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeVerticalReferenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeVerticalReferenceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeWakeTurbulenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeWakeTurbulenceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeYesNoType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeYesNoType), stage, ident.Name, instanceName, preserveOrder)
				case "DateMonthDayType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DateMonthDayType), stage, ident.Name, instanceName, preserveOrder)
				case "NoNumberType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoNumberType), stage, ident.Name, instanceName, preserveOrder)
				case "NoSequenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoSequenceType), stage, ident.Name, instanceName, preserveOrder)
				case "TextAddressType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TextAddressType), stage, ident.Name, instanceName, preserveOrder)
				case "TextDesignatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TextDesignatorType), stage, ident.Name, instanceName, preserveOrder)
				case "TextInstructionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TextInstructionType), stage, ident.Name, instanceName, preserveOrder)
				case "TextNameType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TextNameType), stage, ident.Name, instanceName, preserveOrder)
				case "TextNoteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TextNoteType), stage, ident.Name, instanceName, preserveOrder)
				case "TextPhoneType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TextPhoneType), stage, ident.Name, instanceName, preserveOrder)
				case "TextPropertyNameType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TextPropertyNameType), stage, ident.Name, instanceName, preserveOrder)
				case "TimeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TimeType), stage, ident.Name, instanceName, preserveOrder)
				case "ValAngleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValAngleType), stage, ident.Name, instanceName, preserveOrder)
				case "ValBearingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValBearingType), stage, ident.Name, instanceName, preserveOrder)
				case "ValDistanceSignedType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValDistanceSignedType), stage, ident.Name, instanceName, preserveOrder)
				case "ValDistanceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValDistanceType), stage, ident.Name, instanceName, preserveOrder)
				case "ValDistanceVerticalType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValDistanceVerticalType), stage, ident.Name, instanceName, preserveOrder)
				case "ValDurationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValDurationType), stage, ident.Name, instanceName, preserveOrder)
				case "ValSlopeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValSlopeType), stage, ident.Name, instanceName, preserveOrder)
				case "ValSpeedType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValSpeedType), stage, ident.Name, instanceName, preserveOrder)
				case "ValWeightType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValWeightType), stage, ident.Name, instanceName, preserveOrder)
				case "XHTMLType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(XHTMLType), stage, ident.Name, instanceName, preserveOrder)
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
				case *CodeATCReportingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAerialRefuellingPointType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAerialRefuellingPrefixType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAerialRefuellingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAircraftCategoryType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAircraftEngineNumberType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAircraftEngineType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAircraftICAOType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAircraftType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAircraftWingspanClassType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAltitudeAdjustmentType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAltitudeUseType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAuthorityType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeCardinalDirectionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeCommunicationModeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeComparisonType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeCourseType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeDayType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeDirectionTurnType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeDirectionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeEquipmentAntiCollisionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeFlightOriginType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeFlightPurposeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeFlightRestrictionDesignatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeFlightRestrictionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeFlightRuleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeFlightStatusType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeFlightType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeFlowConditionOperationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeFreeFlightType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeLevelType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeLocationQualifierType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeMeteoConditionsType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeMilitaryRoutePointType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeMilitaryStatusType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeMilitaryTrainingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeNavigationEquipmentType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeNavigationSpecificationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeNotePurposeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeObstacleAssessmentSurfaceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeObstructionIdSurfaceZoneType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRNPType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRVSMPointRoleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRVSMType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeReferenceRoleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRouteAvailabilityType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRouteDesignatorLetterType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRouteDesignatorPrefixType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRouteDesignatorSuffixType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRouteNavigationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRouteOriginType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRouteSegmentPathType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRouteType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRuleProcedureTitleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRuleProcedureType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeSideType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeSpeedReferenceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeTACANChannelType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeTelecomNetworkType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeTimeEventCombinationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeTimeEventType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeTimeReferenceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeTransponderType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeUpperAlphaType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeValueInterpretationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeVerticalDatumType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeVerticalReferenceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeWakeTurbulenceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeYesNoType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *DateMonthDayType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *NoNumberType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *NoSequenceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *TextAddressType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *TextDesignatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *TextInstructionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *TextNameType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *TextNoteType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Lang":
						inst.Lang = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *TextPhoneType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *TextPropertyNameType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *TimeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *ValAngleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *ValBearingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *ValDistanceSignedType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Uom":
						inst.Uom = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *ValDistanceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Uom":
						inst.Uom = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *ValDistanceVerticalType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Uom":
						inst.Uom = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *ValDurationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Uom":
						inst.Uom = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *ValSlopeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *ValSpeedType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Uom":
						inst.Uom = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *ValWeightType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Uom":
						inst.Uom = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *XHTMLType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
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
type CodeATCReportingTypeUnmarshaller struct{}

func (u *CodeATCReportingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeATCReportingType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeATCReportingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeATCReportingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAerialRefuellingPointTypeUnmarshaller struct{}

func (u *CodeAerialRefuellingPointTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAerialRefuellingPointType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAerialRefuellingPointTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAerialRefuellingPointType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAerialRefuellingPrefixTypeUnmarshaller struct{}

func (u *CodeAerialRefuellingPrefixTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAerialRefuellingPrefixType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAerialRefuellingPrefixTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAerialRefuellingPrefixType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAerialRefuellingTypeUnmarshaller struct{}

func (u *CodeAerialRefuellingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAerialRefuellingType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAerialRefuellingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAerialRefuellingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAircraftCategoryTypeUnmarshaller struct{}

func (u *CodeAircraftCategoryTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAircraftCategoryType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAircraftCategoryTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAircraftCategoryType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAircraftEngineNumberTypeUnmarshaller struct{}

func (u *CodeAircraftEngineNumberTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAircraftEngineNumberType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAircraftEngineNumberTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAircraftEngineNumberType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAircraftEngineTypeUnmarshaller struct{}

func (u *CodeAircraftEngineTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAircraftEngineType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAircraftEngineTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAircraftEngineType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAircraftICAOTypeUnmarshaller struct{}

func (u *CodeAircraftICAOTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAircraftICAOType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAircraftICAOTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAircraftICAOType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAircraftTypeUnmarshaller struct{}

func (u *CodeAircraftTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAircraftType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAircraftTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAircraftType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAircraftWingspanClassTypeUnmarshaller struct{}

func (u *CodeAircraftWingspanClassTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAircraftWingspanClassType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAircraftWingspanClassTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAircraftWingspanClassType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAltitudeAdjustmentTypeUnmarshaller struct{}

func (u *CodeAltitudeAdjustmentTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAltitudeAdjustmentType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAltitudeAdjustmentTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAltitudeAdjustmentType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAltitudeUseTypeUnmarshaller struct{}

func (u *CodeAltitudeUseTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAltitudeUseType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAltitudeUseTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAltitudeUseType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAuthorityTypeUnmarshaller struct{}

func (u *CodeAuthorityTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAuthorityType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAuthorityTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAuthorityType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeCardinalDirectionTypeUnmarshaller struct{}

func (u *CodeCardinalDirectionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeCardinalDirectionType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeCardinalDirectionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeCardinalDirectionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeCommunicationModeTypeUnmarshaller struct{}

func (u *CodeCommunicationModeTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeCommunicationModeType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeCommunicationModeTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeCommunicationModeType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeComparisonTypeUnmarshaller struct{}

func (u *CodeComparisonTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeComparisonType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeComparisonTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeComparisonType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeCourseTypeUnmarshaller struct{}

func (u *CodeCourseTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeCourseType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeCourseTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeCourseType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeDayTypeUnmarshaller struct{}

func (u *CodeDayTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeDayType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeDayTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeDayType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeDirectionTurnTypeUnmarshaller struct{}

func (u *CodeDirectionTurnTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeDirectionTurnType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeDirectionTurnTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeDirectionTurnType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeDirectionTypeUnmarshaller struct{}

func (u *CodeDirectionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeDirectionType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeDirectionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeDirectionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeEquipmentAntiCollisionTypeUnmarshaller struct{}

func (u *CodeEquipmentAntiCollisionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeEquipmentAntiCollisionType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeEquipmentAntiCollisionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeEquipmentAntiCollisionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeFlightOriginTypeUnmarshaller struct{}

func (u *CodeFlightOriginTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeFlightOriginType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeFlightOriginTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeFlightOriginType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeFlightPurposeTypeUnmarshaller struct{}

func (u *CodeFlightPurposeTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeFlightPurposeType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeFlightPurposeTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeFlightPurposeType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeFlightRestrictionDesignatorTypeUnmarshaller struct{}

func (u *CodeFlightRestrictionDesignatorTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeFlightRestrictionDesignatorType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeFlightRestrictionDesignatorTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeFlightRestrictionDesignatorType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeFlightRestrictionTypeUnmarshaller struct{}

func (u *CodeFlightRestrictionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeFlightRestrictionType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeFlightRestrictionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeFlightRestrictionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeFlightRuleTypeUnmarshaller struct{}

func (u *CodeFlightRuleTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeFlightRuleType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeFlightRuleTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeFlightRuleType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeFlightStatusTypeUnmarshaller struct{}

func (u *CodeFlightStatusTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeFlightStatusType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeFlightStatusTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeFlightStatusType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeFlightTypeUnmarshaller struct{}

func (u *CodeFlightTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeFlightType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeFlightTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeFlightType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeFlowConditionOperationTypeUnmarshaller struct{}

func (u *CodeFlowConditionOperationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeFlowConditionOperationType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeFlowConditionOperationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeFlowConditionOperationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeFreeFlightTypeUnmarshaller struct{}

func (u *CodeFreeFlightTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeFreeFlightType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeFreeFlightTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeFreeFlightType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeLevelTypeUnmarshaller struct{}

func (u *CodeLevelTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeLevelType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeLevelTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeLevelType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeLocationQualifierTypeUnmarshaller struct{}

func (u *CodeLocationQualifierTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeLocationQualifierType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeLocationQualifierTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeLocationQualifierType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeMeteoConditionsTypeUnmarshaller struct{}

func (u *CodeMeteoConditionsTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeMeteoConditionsType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeMeteoConditionsTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeMeteoConditionsType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeMilitaryRoutePointTypeUnmarshaller struct{}

func (u *CodeMilitaryRoutePointTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeMilitaryRoutePointType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeMilitaryRoutePointTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeMilitaryRoutePointType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeMilitaryStatusTypeUnmarshaller struct{}

func (u *CodeMilitaryStatusTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeMilitaryStatusType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeMilitaryStatusTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeMilitaryStatusType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeMilitaryTrainingTypeUnmarshaller struct{}

func (u *CodeMilitaryTrainingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeMilitaryTrainingType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeMilitaryTrainingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeMilitaryTrainingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeNavigationEquipmentTypeUnmarshaller struct{}

func (u *CodeNavigationEquipmentTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeNavigationEquipmentType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeNavigationEquipmentTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeNavigationEquipmentType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeNavigationSpecificationTypeUnmarshaller struct{}

func (u *CodeNavigationSpecificationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeNavigationSpecificationType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeNavigationSpecificationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeNavigationSpecificationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeNotePurposeTypeUnmarshaller struct{}

func (u *CodeNotePurposeTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeNotePurposeType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeNotePurposeTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeNotePurposeType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeObstacleAssessmentSurfaceTypeUnmarshaller struct{}

func (u *CodeObstacleAssessmentSurfaceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeObstacleAssessmentSurfaceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeObstacleAssessmentSurfaceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeObstacleAssessmentSurfaceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeObstructionIdSurfaceZoneTypeUnmarshaller struct{}

func (u *CodeObstructionIdSurfaceZoneTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeObstructionIdSurfaceZoneType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeObstructionIdSurfaceZoneTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeObstructionIdSurfaceZoneType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRNPTypeUnmarshaller struct{}

func (u *CodeRNPTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRNPType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRNPTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRNPType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRVSMPointRoleTypeUnmarshaller struct{}

func (u *CodeRVSMPointRoleTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRVSMPointRoleType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRVSMPointRoleTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRVSMPointRoleType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRVSMTypeUnmarshaller struct{}

func (u *CodeRVSMTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRVSMType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRVSMTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRVSMType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeReferenceRoleTypeUnmarshaller struct{}

func (u *CodeReferenceRoleTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeReferenceRoleType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeReferenceRoleTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeReferenceRoleType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRouteAvailabilityTypeUnmarshaller struct{}

func (u *CodeRouteAvailabilityTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRouteAvailabilityType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRouteAvailabilityTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRouteAvailabilityType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRouteDesignatorLetterTypeUnmarshaller struct{}

func (u *CodeRouteDesignatorLetterTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRouteDesignatorLetterType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRouteDesignatorLetterTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRouteDesignatorLetterType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRouteDesignatorPrefixTypeUnmarshaller struct{}

func (u *CodeRouteDesignatorPrefixTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRouteDesignatorPrefixType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRouteDesignatorPrefixTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRouteDesignatorPrefixType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRouteDesignatorSuffixTypeUnmarshaller struct{}

func (u *CodeRouteDesignatorSuffixTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRouteDesignatorSuffixType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRouteDesignatorSuffixTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRouteDesignatorSuffixType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRouteNavigationTypeUnmarshaller struct{}

func (u *CodeRouteNavigationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRouteNavigationType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRouteNavigationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRouteNavigationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRouteOriginTypeUnmarshaller struct{}

func (u *CodeRouteOriginTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRouteOriginType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRouteOriginTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRouteOriginType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRouteSegmentPathTypeUnmarshaller struct{}

func (u *CodeRouteSegmentPathTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRouteSegmentPathType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRouteSegmentPathTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRouteSegmentPathType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRouteTypeUnmarshaller struct{}

func (u *CodeRouteTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRouteType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRouteTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRouteType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRuleProcedureTitleTypeUnmarshaller struct{}

func (u *CodeRuleProcedureTitleTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRuleProcedureTitleType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRuleProcedureTitleTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRuleProcedureTitleType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRuleProcedureTypeUnmarshaller struct{}

func (u *CodeRuleProcedureTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRuleProcedureType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRuleProcedureTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRuleProcedureType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeSideTypeUnmarshaller struct{}

func (u *CodeSideTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeSideType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeSideTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeSideType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeSpeedReferenceTypeUnmarshaller struct{}

func (u *CodeSpeedReferenceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeSpeedReferenceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeSpeedReferenceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeSpeedReferenceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeTACANChannelTypeUnmarshaller struct{}

func (u *CodeTACANChannelTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeTACANChannelType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeTACANChannelTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeTACANChannelType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeTelecomNetworkTypeUnmarshaller struct{}

func (u *CodeTelecomNetworkTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeTelecomNetworkType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeTelecomNetworkTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeTelecomNetworkType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeTimeEventCombinationTypeUnmarshaller struct{}

func (u *CodeTimeEventCombinationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeTimeEventCombinationType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeTimeEventCombinationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeTimeEventCombinationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeTimeEventTypeUnmarshaller struct{}

func (u *CodeTimeEventTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeTimeEventType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeTimeEventTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeTimeEventType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeTimeReferenceTypeUnmarshaller struct{}

func (u *CodeTimeReferenceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeTimeReferenceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeTimeReferenceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeTimeReferenceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeTransponderTypeUnmarshaller struct{}

func (u *CodeTransponderTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeTransponderType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeTransponderTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeTransponderType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeUpperAlphaTypeUnmarshaller struct{}

func (u *CodeUpperAlphaTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeUpperAlphaType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeUpperAlphaTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeUpperAlphaType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeValueInterpretationTypeUnmarshaller struct{}

func (u *CodeValueInterpretationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeValueInterpretationType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeValueInterpretationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeValueInterpretationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeVerticalDatumTypeUnmarshaller struct{}

func (u *CodeVerticalDatumTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeVerticalDatumType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeVerticalDatumTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeVerticalDatumType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeVerticalReferenceTypeUnmarshaller struct{}

func (u *CodeVerticalReferenceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeVerticalReferenceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeVerticalReferenceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeVerticalReferenceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeWakeTurbulenceTypeUnmarshaller struct{}

func (u *CodeWakeTurbulenceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeWakeTurbulenceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeWakeTurbulenceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeWakeTurbulenceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeYesNoTypeUnmarshaller struct{}

func (u *CodeYesNoTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeYesNoType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeYesNoTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeYesNoType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type DateMonthDayTypeUnmarshaller struct{}

func (u *DateMonthDayTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DateMonthDayType), stage, identifier, instanceName, preserveOrder)
}

func (u *DateMonthDayTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DateMonthDayType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type NoNumberTypeUnmarshaller struct{}

func (u *NoNumberTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NoNumberType), stage, identifier, instanceName, preserveOrder)
}

func (u *NoNumberTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NoNumberType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractInt(valueExpr)
	}
	return nil
}

type NoSequenceTypeUnmarshaller struct{}

func (u *NoSequenceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NoSequenceType), stage, identifier, instanceName, preserveOrder)
}

func (u *NoSequenceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NoSequenceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractInt(valueExpr)
	}
	return nil
}

type TextAddressTypeUnmarshaller struct{}

func (u *TextAddressTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TextAddressType), stage, identifier, instanceName, preserveOrder)
}

func (u *TextAddressTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TextAddressType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type TextDesignatorTypeUnmarshaller struct{}

func (u *TextDesignatorTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TextDesignatorType), stage, identifier, instanceName, preserveOrder)
}

func (u *TextDesignatorTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TextDesignatorType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type TextInstructionTypeUnmarshaller struct{}

func (u *TextInstructionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TextInstructionType), stage, identifier, instanceName, preserveOrder)
}

func (u *TextInstructionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TextInstructionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type TextNameTypeUnmarshaller struct{}

func (u *TextNameTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TextNameType), stage, identifier, instanceName, preserveOrder)
}

func (u *TextNameTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TextNameType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type TextNoteTypeUnmarshaller struct{}

func (u *TextNoteTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TextNoteType), stage, identifier, instanceName, preserveOrder)
}

func (u *TextNoteTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TextNoteType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Lang":
		instance.Lang = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type TextPhoneTypeUnmarshaller struct{}

func (u *TextPhoneTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TextPhoneType), stage, identifier, instanceName, preserveOrder)
}

func (u *TextPhoneTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TextPhoneType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type TextPropertyNameTypeUnmarshaller struct{}

func (u *TextPropertyNameTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TextPropertyNameType), stage, identifier, instanceName, preserveOrder)
}

func (u *TextPropertyNameTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TextPropertyNameType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type TimeTypeUnmarshaller struct{}

func (u *TimeTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TimeType), stage, identifier, instanceName, preserveOrder)
}

func (u *TimeTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TimeType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValAngleTypeUnmarshaller struct{}

func (u *ValAngleTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValAngleType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValAngleTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValAngleType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValBearingTypeUnmarshaller struct{}

func (u *ValBearingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValBearingType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValBearingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValBearingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValDistanceSignedTypeUnmarshaller struct{}

func (u *ValDistanceSignedTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValDistanceSignedType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValDistanceSignedTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValDistanceSignedType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Uom":
		instance.Uom = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValDistanceTypeUnmarshaller struct{}

func (u *ValDistanceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValDistanceType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValDistanceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValDistanceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Uom":
		instance.Uom = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValDistanceVerticalTypeUnmarshaller struct{}

func (u *ValDistanceVerticalTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValDistanceVerticalType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValDistanceVerticalTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValDistanceVerticalType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Uom":
		instance.Uom = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValDurationTypeUnmarshaller struct{}

func (u *ValDurationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValDurationType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValDurationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValDurationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Uom":
		instance.Uom = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValSlopeTypeUnmarshaller struct{}

func (u *ValSlopeTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValSlopeType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValSlopeTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValSlopeType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValSpeedTypeUnmarshaller struct{}

func (u *ValSpeedTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValSpeedType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValSpeedTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValSpeedType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Uom":
		instance.Uom = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValWeightTypeUnmarshaller struct{}

func (u *ValWeightTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValWeightType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValWeightTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValWeightType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Uom":
		instance.Uom = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type XHTMLTypeUnmarshaller struct{}

func (u *XHTMLTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(XHTMLType), stage, identifier, instanceName, preserveOrder)
}

func (u *XHTMLTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*XHTMLType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	}
	return nil
}
